/*
Copyright 2020 The Kubernetes Authors.

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
	"context"
	fmt "fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/config"
	errors "k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	klog "k8s.io/klog/v2"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
)

// GrpcNodeGroup implements NodeGroup interface.
type GrpcNodeGroup struct {
	name    string
	manager *GrpcManager
}

// GetManager return the manager
func (ng *GrpcNodeGroup) GetManager() *GrpcManager {
	return ng.manager
}

// MaxSize returns maximum size of the node group.
func (ng *GrpcNodeGroup) MaxSize() int {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.MaxSize(ctx, &NodeGroupServiceRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::MaxSize for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return 0
		}

		return int(r.GetMaxSize())
	}

	return 0
}

// MinSize returns minimum size of the node group.
func (ng *GrpcNodeGroup) MinSize() int {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())

	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.MinSize(ctx, &NodeGroupServiceRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::MinSize for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return 0
		}

		return int(r.GetMinSize())
	}

	return 0
}

// TargetSize returns the current TARGET size of the node group. It is possible that the
// number is different from the number of nodes registered in Kubernetes.
func (ng *GrpcNodeGroup) TargetSize() (int, error) {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.TargetSize(ctx, &NodeGroupServiceRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::TargetSize for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return 0, err
		} else if rerr := r.GetError(); rerr != nil {
			klog.Errorf("Cloud provider:%s:%s call NodeGroup::TargetSize got error: %v", manager.GetCloudProviderID(), ng.name, rerr)

			return 0, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}

		return int(r.GetTargetSize()), nil
	}

	return 0, err
}

// Exist checks if the node group really exists on the cloud provider side. Allows to tell the
// theoretical node group from the real one.
func (ng *GrpcNodeGroup) Exist() bool {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.Exist(ctx, &NodeGroupServiceRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::Exist for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return false
		}

		return r.GetExists()
	}

	return false
}

// Create creates the node group on the cloud provider side.
func (ng *GrpcNodeGroup) Create() (cloudprovider.NodeGroup, error) {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.Create(ctx, &NodeGroupServiceRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::Create for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return nil, err
		} else if rerr := r.GetError(); rerr != nil {
			klog.Errorf("Cloud provider:%s:%s call NodeGroup::Create got error: %v", manager.GetCloudProviderID(), ng.name, rerr)

			return nil, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}

		return &GrpcNodeGroup{
			name:    r.GetNodeGroup().GetId(),
			manager: manager,
		}, nil
	}

	return nil, err
}

// Autoprovisioned returns true if the node group is autoprovisioned.
func (ng *GrpcNodeGroup) Autoprovisioned() bool {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.Autoprovisioned(ctx, &NodeGroupServiceRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::Autoprovisioned for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return false
		}

		return r.GetAutoprovisioned()
	}

	return false
}

// Delete deletes the node group on the cloud provider side.
// This will be executed only for autoprovisioned node groups, once their size drops to 0.
func (ng *GrpcNodeGroup) Delete() error {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.Delete(ctx, &NodeGroupServiceRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::Create for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return err
		} else if rerr := r.GetError(); rerr != nil {
			klog.Errorf("Cloud provider:%s:%s call NodeGroup::Create got error: %v", manager.GetCloudProviderID(), ng.name, rerr)

			return errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}
	}

	return err
}

// IncreaseSize increases Asg size
func (ng *GrpcNodeGroup) IncreaseSize(delta int) error {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.IncreaseSize(ctx, &IncreaseSizeRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name, Delta: int32(delta)})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::IncreaseSize for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return err
		} else if rerr := r.GetError(); rerr != nil {
			klog.Errorf("Cloud provider:%s:%s call NodeGroup::IncreaseSize got error: %v", manager.GetCloudProviderID(), ng.name, rerr)

			return errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}
	}

	return err
}

// DecreaseTargetSize decreases the target size of the node group. This function
// doesn't permit to delete any existing node and can be used only to reduce the
// request for new nodes that have not been yet fulfilled. Delta should be negative.
// It is assumed that cloud provider will not delete the existing nodes if the size
// when there is an option to just decrease the target.
func (ng *GrpcNodeGroup) DecreaseTargetSize(delta int) error {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())

	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.DecreaseTargetSize(ctx, &DecreaseTargetSizeRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name, Delta: int32(delta)})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::DecreaseTargetSize for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return err
		} else if rerr := r.GetError(); rerr != nil {
			klog.Errorf("Cloud provider:%s:%s call NodeGroup::DecreaseTargetSize got error: %v", manager.GetCloudProviderID(), ng.name, rerr)

			return errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}
	}

	return err
}

// Belongs returns true if the given node belongs to the NodeGroup.
func (ng *GrpcNodeGroup) Belongs(node *v1.Node) (bool, error) {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.Belongs(ctx, &BelongsRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name, Node: toJSON(node)})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::Belongs for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return false, err
		} else if rerr := r.GetError(); rerr != nil {
			klog.Errorf("Cloud provider:%s:%s call NodeGroup::Belongs got error: %v", manager.GetCloudProviderID(), ng.name, rerr)

			return false, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}

		return r.GetBelongs(), nil
	}

	return false, err
}

// DeleteNodes deletes the nodes from the group.
func (ng *GrpcNodeGroup) DeleteNodes(nodes []*v1.Node) error {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.DeleteNodes(ctx, &DeleteNodesRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name, Node: nodesAsJSON(nodes)})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::DeleteNodes for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return err
		} else if rerr := r.GetError(); rerr != nil {
			klog.Errorf("Cloud provider:%s:%s call NodeGroup::DeleteNodes got error: %v", manager.GetCloudProviderID(), ng.name, rerr)

			return errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}
	}

	return err
}

// Id returns asg id.
func (ng *GrpcNodeGroup) Id() string {
	return ng.name
}

// Debug returns a debug string for the Asg.
func (ng *GrpcNodeGroup) Debug() string {
	return fmt.Sprintf("%s (%d:%d)", ng.Id(), ng.MinSize(), ng.MaxSize())
}

// Nodes returns a list of all nodes that belong to this node group.
func (ng *GrpcNodeGroup) Nodes() ([]cloudprovider.Instance, error) {
	var instances []cloudprovider.Instance
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.Nodes(ctx, &NodeGroupServiceRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::Nodes for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return nil, err
		} else if rerr := r.GetError(); rerr != nil {
			klog.Errorf("Cloud provider:%s:%s call NodeGroup::Nodes got error: %v", manager.GetCloudProviderID(), ng.name, rerr)

			return nil, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}

		if r.GetInstances() != nil && r.GetInstances().GetItems() != nil {

			instances = make([]cloudprovider.Instance, len(r.GetInstances().GetItems()))

			for i, instance := range r.GetInstances().GetItems() {
				errorInfo := instance.GetStatus().GetErrorInfo()

				var instanceErrorInfo *cloudprovider.InstanceErrorInfo

				if errorInfo != nil {
					instanceErrorInfo = &cloudprovider.InstanceErrorInfo{
						ErrorClass:   cloudprovider.InstanceErrorClass(errorInfo.GetErrorClass()),
						ErrorCode:    errorInfo.GetErrorCode(),
						ErrorMessage: errorInfo.GetErrorMessage(),
					}
				}

				cloudInstance := cloudprovider.Instance{
					Id: instance.GetId(),
					Status: &cloudprovider.InstanceStatus{
						State:     cloudprovider.InstanceState(instance.GetStatus().GetState()),
						ErrorInfo: instanceErrorInfo,
					},
				}

				instances[i] = cloudInstance
			}
		}
	}

	return instances, err
}

func newNodeInfoResource(resource *Resource) *schedulerframework.Resource {
	res := &schedulerframework.Resource{}

	res.MilliCPU = resource.MilliCPU
	res.Memory = resource.Memory
	res.EphemeralStorage = resource.EphemeralStorage
	res.AllowedPodNumber = int(resource.AllowedPodNumber)

	if resource.ScalarResources != nil {
		res.ScalarResources = make(map[v1.ResourceName]int64, len(resource.ScalarResources))

		for n, r := range resource.ScalarResources {
			res.ScalarResources[v1.ResourceName(n)] = r
		}
	}

	return res
}

// TemplateNodeInfo returns a node template for this node group.
func (ng *GrpcNodeGroup) TemplateNodeInfo() (*schedulerframework.NodeInfo, error) {
	manager := ng.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	nodeGroupService, err := manager.GetNodeGroupServiceClient()

	if err == nil {
		r, err := nodeGroupService.TemplateNodeInfo(ctx, &NodeGroupServiceRequest{ProviderID: manager.GetCloudProviderID(), NodeGroupID: ng.name})

		if err != nil {
			klog.Errorf("Could not get NodeGroup::TemplateNodeInfo for cloud provider:%s:%s error: %v", manager.GetCloudProviderID(), ng.name, err)

			return nil, err
		} else if rerr := r.GetError(); rerr != nil {
			klog.Errorf("Cloud provider:%s:%s call NodeGroup::TemplateNodeInfo got error: %v", manager.GetCloudProviderID(), ng.name, rerr)

			return nil, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}

		return newNodeInfo(r.GetNodeInfo()), nil
	}

	return nil, err
}

// GetOptions returns NodeGroupAutoscalingOptions that should be used for this particular
// NodeGroup. Returning a nil will result in using default options.
func (ng *GrpcNodeGroup) GetOptions(defaults config.NodeGroupAutoscalingOptions) (*config.NodeGroupAutoscalingOptions, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (ng *GrpcNodeGroup) String() string {
	return fmt.Sprintf("{ %s }", ng.Debug())
}
