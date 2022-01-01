/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package grpccloudprovider

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
)

func ignoreTestMultipassClusterScaling(t *testing.T) {
	provider := testProviderWithEmbed(t, 300, false)

	defer provider.grpcCloudProvider.Cleanup()

	var err error
	var nodeGroup cloudprovider.NodeGroup
	var instances []cloudprovider.Instance

	if assert.NoError(t, provider.Refresh()) {
		systemLabels := map[string]string{
			apiv1.LabelZoneFailureDomain: "zone-test",
		}
		labels := map[string]string{
			"monitor":  "true",
			"database": "true",
		}

		nodeGroup, err = provider.NewNodeGroup("medium", labels, systemLabels, nil, nil)

		if assert.NoError(t, err, "NewNodeGroup failed") {

			if assert.NotNil(t, nodeGroup) {
				nodeGroup, err = nodeGroup.Create()

				targetSize, _ := nodeGroup.TargetSize()

				assert.Equal(t, nodeGroup.MinSize(), testMinNodeSize)
				assert.Equal(t, nodeGroup.MaxSize(), testMaxNodeSize)
				assert.Equal(t, targetSize, 0)

				if assert.NoError(t, err, "Create node group failed") {

					if assert.Equal(t, testGroupID, nodeGroup.(*GrpcNodeGroup).name, "Created node group no found") {

						provider.Refresh()

						if assert.NoError(t, nodeGroup.IncreaseSize(1), "Increase size failed") {
							instances, err = nodeGroup.Nodes()

							if assert.NoError(t, err, "Nodes failed") && assert.NotNil(t, instances, "Nodes doesn't return instance") {

								if assert.True(t, len(instances) > 0) {
									instance := instances[0]

									if assert.Equal(t, instance.Id, providerIDForNode(testGroupID, testNodeName), "Unexpected instance") {
										if instance.Status != nil {
											if instance.Status.State != cloudprovider.InstanceRunning {
												log.Printf("Current status of node is:%d", instance.Status.State)
											}
										}

										nodeInfo, err := nodeGroup.TemplateNodeInfo()

										if assert.NoError(t, err, "TemplateNodeInfo failed") {

											if assert.NotNil(t, nodeInfo, "TemplateNodeInfo return nil") {
												node := &apiv1.Node{
													Spec: apiv1.NodeSpec{
														ProviderID: instance.Id,
													},
												}

												ng, err := provider.NodeGroupForNode(node)

												if assert.NoError(t, err, "NodeGroupForNode failed") {
													if assert.NotNil(t, ng, "NodeGroupForNode return nil") {
														assert.Equal(t, ng.Id(), nodeGroup.Id(), "NodeGroupForNode not same group id")
													}
												}

												if assert.NoError(t, nodeGroup.DeleteNodes([]*apiv1.Node{node}), "DeleteNodes failed") {

												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
