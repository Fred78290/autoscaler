package grpccloudprovider

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
)

type testPriceModel struct {
	GrpcPriceModel
	provider *testGrpcCloudProvider
}

func newTestPriceModel(t *testing.T) *testPriceModel {
	provider := testProvider(t)

	return &testPriceModel{
		provider: provider,
		GrpcPriceModel: GrpcPriceModel{
			name:    testGroupID,
			manager: provider.GetManager(),
		},
	}
}

func (pm *testPriceModel) providerID(providerID string) string {
	return fmt.Sprintf("%s://%s/object?type=group", providerID, pm.name)
}

func (pm *testPriceModel) providerIDForNode(nodeName string) string {
	return fmt.Sprintf("%s://%s/object?type=node&name=%s", testProviderID, pm.name, nodeName)
}

func (pm *testPriceModel) Cleanup() {
	pm.provider.Cleanup()
}

func (pm *testPriceModel) Refresh() error {
	return pm.provider.Refresh()
}

func TestNodePrice(t *testing.T) {
	pm := newTestPriceModel(t)
	defer pm.Cleanup()

	node := &apiv1.Node{
		Spec: apiv1.NodeSpec{
			ProviderID: pm.providerIDForNode(testNodeName),
		},
	}

	now := time.Now()
	price, err := pm.NodePrice(node, now, now.Add(time.Hour))

	if assert.NoError(t, err) {
		assert.Equal(t, price, float64(0))
	}
}

func TestPodPrice(t *testing.T) {
	pm := newTestPriceModel(t)
	defer pm.Cleanup()

	now := time.Now()

	pod := &apiv1.Pod{
		Spec: apiv1.PodSpec{
			NodeName: "test-instance-id",
		},
	}

	price, err := pm.PodPrice(pod, now, now.Add(time.Hour))

	if assert.NoError(t, err) {
		assert.Equal(t, price, float64(0))
	}
}
