package grpccloudprovider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
)

const (
	testProviderID  = "multipass"
	testGroupID     = "ca-grpc-multipass"
	testNodeName    = "ca-grpc-multipass-vm-00"
	testMinNodeSize = 0
	testMaxNodeSize = 5
)

type testNodeGroup struct {
	GrpcNodeGroup
	provider *testGrpcCloudProvider
}

func newTestNodeGroup(t *testing.T) *testNodeGroup {
	provider := testProvider(t)

	return &testNodeGroup{
		provider: provider,
		GrpcNodeGroup: GrpcNodeGroup{
			name:    testGroupID,
			manager: provider.GetManager(),
		},
	}
}

func (ng *testNodeGroup) Cleanup() {
	ng.provider.Cleanup()
}

func (ng *testNodeGroup) Refresh() error {
	return ng.provider.Refresh()
}

func (ng *testNodeGroup) providerIDForNode(nodeName string) string {
	return fmt.Sprintf("%s://%s/object?type=node&name=%s", testProviderID, ng.name, nodeName)
}

func TestMaxSize(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	if assert.NoError(t, ng.Refresh()) {
		assert.Equal(t, ng.MaxSize(), testMaxNodeSize)
	}
}

func TestMinSize(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	if assert.NoError(t, ng.Refresh()) {
		assert.Equal(t, ng.MinSize(), testMinNodeSize)
	}
}

func TestTargetSize(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	if assert.NoError(t, ng.Refresh()) {
		targetSize, err := ng.TargetSize()

		assert.NoError(t, err)
		assert.Equal(t, targetSize, testMaxNodeSize)
	}
}

func TestExist(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	if assert.NoError(t, ng.Refresh()) {
		assert.Equal(t, ng.Exist(), true)

		unexistsNodeGroup := &GrpcNodeGroup{
			name:    "not",
			manager: ng.GetManager(),
		}

		assert.Equal(t, unexistsNodeGroup.Exist(), false)
	}
}

func TestCreate(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	if assert.NoError(t, ng.Refresh()) {
		nodeGroup, err := ng.Create()

		if assert.NoError(t, err) {
			assert.Equal(t, testGroupID, nodeGroup.(*GrpcNodeGroup).name)
		}
	}
}

func TestAutoprovisioned(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	if assert.NoError(t, ng.Refresh()) {
		assert.Equal(t, ng.Autoprovisioned(), true)
	}
}

func TestDelete(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	if assert.NoError(t, ng.Refresh()) {
		assert.NoError(t, ng.Delete())
	}
}

func TestIncreaseSize(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	if assert.NoError(t, ng.Refresh()) {
		assert.NoError(t, ng.IncreaseSize(1))
	}
}

func TestDecreaseTargetSize(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	if assert.NoError(t, ng.Refresh()) {
		assert.NoError(t, ng.DecreaseTargetSize(1))
	}
}

func TestBelongs(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	if assert.NoError(t, ng.Refresh()) {
		node := &apiv1.Node{
			Spec: apiv1.NodeSpec{
				ProviderID: ng.providerIDForNode(testNodeName),
			},
		}

		belong, err := ng.Belongs(node)

		if assert.NoError(t, err) {
			assert.Equal(t, belong, true)

			newnode := &apiv1.Node{
				Spec: apiv1.NodeSpec{
					ProviderID: ng.providerIDForNode("ca-grpc-multipass-vm-01"),
				},
			}

			notbelong, err := ng.Belongs(newnode)

			if assert.NoError(t, err) {
				assert.Equal(t, notbelong, false)
			}
		}
	}
}

func TestDeleteNodes(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	assert.NoError(t, ng.Refresh())

	node := &apiv1.Node{
		Spec: apiv1.NodeSpec{
			ProviderID: ng.providerIDForNode(testNodeName),
		},
	}

	assert.NoError(t, ng.DeleteNodes([]*apiv1.Node{node}))
}

func TestNodes(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	assert.NoError(t, ng.Refresh())

	instance, err := ng.Nodes()

	if assert.NoError(t, err) {
		assert.Equal(t, len(instance), 1)
	}
}

func TestTemplateNodeInfo(t *testing.T) {
	ng := newTestNodeGroup(t)
	defer ng.Cleanup()

	assert.NoError(t, ng.Refresh())

	nodeInfo, err := ng.TemplateNodeInfo()

	if assert.NoError(t, err) {
		assert.NotNil(t, nodeInfo)
	}
}
