package grpccloudprovider

import (
	"fmt"
	"log"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
)

type server struct{}

type testGrpcCloudProvider struct {
	grpcCloudProvider
	useEmbeddedTestServer bool
	wg                    sync.WaitGroup
}

var (
	availableGPUTypes = map[string]struct{}{
		"Tesla-P4": {},
		"M40":      {},
		"P100":     {},
		"V100":     {},
	}
)

func (grpc *testGrpcCloudProvider) Cleanup() error {

	if grpc.useEmbeddedTestServer {
		err := grpc.grpcCloudProvider.Cleanup()

		stopTestServer()

		return err
	}

	return nil
}

func (grpc *testGrpcCloudProvider) Start() {
	if grpc.useEmbeddedTestServer {
		grpc.wg.Add(1)
		go startTestServer(&grpc.wg)

		grpc.wg.Wait()
	}
}

func testProvider(t *testing.T) *testGrpcCloudProvider {
	return testProviderWithEmbed(t, 30, true)
}

func testProviderWithEmbed(t *testing.T, timeout int, embedded bool) *testGrpcCloudProvider {
	resourceLimiter := cloudprovider.NewResourceLimiter(
		map[string]int64{cloudprovider.ResourceNameCores: 1, cloudprovider.ResourceNameMemory: 10000000},
		map[string]int64{cloudprovider.ResourceNameCores: 5, cloudprovider.ResourceNameMemory: 100000000})

	manager := &GrpcManager{
		resourceLimiter: resourceLimiter,
		config: GrpcConfig{
			Address:    "unix:/tmp/cluster-autoscaler-grpc.sock",
			Identifier: testProviderID,
			Timeout:    timeout,
		},
	}

	grpc := &testGrpcCloudProvider{
		useEmbeddedTestServer: embedded,
		grpcCloudProvider: grpcCloudProvider{
			manager:         manager,
			resourceLimiter: resourceLimiter,
			nodeGroups: []*NodeGroupDef{
				&NodeGroupDef{
					Provisionned:        false,
					IncludeExistingNode: false,
					NodeGroupID:         testGroupID,
					MinSize:             testMinNodeSize,
					MaxSize:             testMaxNodeSize,
					Labels: map[string]string{
						"monitor":  "true",
						"database": "true",
					},
				},
			},
		},
	}

	if embedded {
		grpc.Start()
	}

	if manager.isConnected() == false {
		err := manager.connect()

		assert.NoError(t, err)
	}

	return grpc
}

func TestName(t *testing.T) {
	provider := testProvider(t)
	defer provider.Cleanup()

	assert.NoError(t, provider.Refresh())

	assert.Equal(t, provider.Name(), ProviderName)
}

func TestCleanup(t *testing.T) {
	provider := testProvider(t)

	assert.NoError(t, provider.Refresh())

	err := provider.Cleanup()

	assert.NoError(t, err)
}

func TestNodeGroups(t *testing.T) {
	provider := testProvider(t)
	defer provider.Cleanup()

	if assert.NoError(t, provider.Refresh()) {

		nodeGroups := provider.NodeGroups()

		log.Printf("Called NodeGroups got: %v", nodeGroups)

		if assert.True(t, len(nodeGroups) > 0) {
			assert.Equal(t, nodeGroups[0].Id(), testGroupID)
			assert.Equal(t, nodeGroups[0].MinSize(), testMinNodeSize)
			assert.Equal(t, nodeGroups[0].MaxSize(), testMaxNodeSize)
		}
	}
}

func TestNodeGroupForNode(t *testing.T) {
	provider := testProvider(t)
	defer provider.Cleanup()

	node := &apiv1.Node{
		Spec: apiv1.NodeSpec{
			ProviderID: fmt.Sprintf("%s://%s/object?type=node&name=%s", testProviderID, testGroupID, testNodeName),
		},
	}

	if assert.NoError(t, provider.Refresh()) {

		group, err := provider.NodeGroupForNode(node)

		if assert.NoError(t, err) {
			if assert.NotNil(t, group) {

				log.Printf("Called NodeGroupForNode got: %v", group)

				assert.Equal(t, group.Id(), testGroupID)
				assert.Equal(t, group.MinSize(), 0)
				assert.Equal(t, group.MaxSize(), 5)

				// test node in cluster that is not in a group managed by cluster autoscaler
				nodeNotInGroup := &apiv1.Node{
					Spec: apiv1.NodeSpec{
						ProviderID: "grpc:///localhost/test-instance-id-not-in-group",
					},
				}

				group, err = provider.NodeGroupForNode(nodeNotInGroup)

				assert.Error(t, err)
				assert.Nil(t, group)
			}
		}
	}
}

func TestPricing(t *testing.T) {
	provider := testProvider(t)
	defer provider.Cleanup()

	assert.NoError(t, provider.Refresh())

	pricing, err := provider.Pricing()

	if assert.NoError(t, err) {
		log.Printf("Called Pricing got: %v", pricing)

		assert.Equal(t, pricing.(*GrpcPriceModel).name, testProviderID)
	}
}

func TestGetAvailableMachineTypes(t *testing.T) {
	provider := testProvider(t)
	defer provider.Cleanup()

	if assert.NoError(t, provider.Refresh()) {

		available, err := provider.GetAvailableMachineTypes()

		if assert.NoError(t, err) {

			log.Printf("Called GetAvailableMachineTypes got: %v", available)

			if assert.True(t, len(available) > 0) {
				assert.Contains(t, []string{"tiny", "medium", "large", "extra-large"}, available[0])
			}
		}
	}
}

func TestNewNodeGroup(t *testing.T) {
	provider := testProvider(t)
	defer provider.Cleanup()

	if assert.NoError(t, provider.Refresh()) {
		systemLabels := map[string]string{
			apiv1.LabelZoneFailureDomain: "zone-test",
		}
		labels := map[string]string{
			"monitor":  "true",
			"database": "true",
		}
		nodeGroup, err := provider.NewNodeGroup("medium", labels, systemLabels, nil, nil)

		if assert.NoError(t, err) {
			log.Printf("Called NewNodeGroup got: %v", nodeGroup)

			assert.NotNil(t, nodeGroup)
		}
	}
}

func TestGetResourceLimiter(t *testing.T) {
	provider := testProvider(t)
	defer provider.Cleanup()

	if assert.NoError(t, provider.Refresh()) {

		rl, err := provider.GetResourceLimiter()

		if assert.NoError(t, err) {

			log.Printf("Called GetResourceLimiter got: %v", rl)

			assert.NotNil(t, rl)
		}
	}
}

func TestGPULabel(t *testing.T) {
	provider := testProvider(t)
	defer provider.Cleanup()

	if assert.NoError(t, provider.Refresh()) {

		label := provider.GPULabel()

		if assert.NotEmpty(t, label) {

			log.Printf("Called GPULabel got: %v", label)

			assert.Equal(t, label, "turing")
		}
	}
}

func TestGetAvailableGPUTypes(t *testing.T) {
	provider := testProvider(t)
	defer provider.Cleanup()

	if assert.NoError(t, provider.Refresh()) {

		gpus := provider.GetAvailableGPUTypes()

		if assert.NotNil(t, gpus) {

			log.Printf("Called GetAvailableGPUTypes got: %v", gpus)
		}
	}
}

func TestRefresh(t *testing.T) {
	provider := testProvider(t)
	defer provider.Cleanup()

	assert.NoError(t, provider.Refresh())
}
