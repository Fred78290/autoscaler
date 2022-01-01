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
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/config"
	errors "k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	klog "k8s.io/klog/v2"
)

const (
	// ProviderName is the cloud provider name for AWS
	ProviderName = "grpc"
)

// ErrMissingConfig is returned if GRPC config is missing.
var ErrMissingConfig = errors.NewAutoscalerError(errors.InternalError, "Missing GRPC config")

// ErrMismatchingProvider is returned if Provider ID does'nt match with target server.
var ErrMismatchingProvider = errors.NewAutoscalerError(errors.InternalError, "Provider ID doesn't match with target server")

// grpcCloudProvider implements CloudProvider interface.
type grpcCloudProvider struct {
	manager         *GrpcManager
	resourceLimiter *cloudprovider.ResourceLimiter
	nodeGroups      []*NodeGroupDef
}

func (grpc *grpcCloudProvider) GetManager() *GrpcManager {
	return grpc.manager
}

// Cleanup stops the go routine that is handling the current view of the ASGs in the form of a cache
func (grpc *grpcCloudProvider) Cleanup() error {
	manager := grpc.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	cloudProviderService, err := manager.GetCloudProviderServiceClient()

	if err == nil {
		r, err := cloudProviderService.Cleanup(ctx, &CloudProviderServiceRequest{ProviderID: manager.GetCloudProviderID()})

		if err != nil {
			log.Printf("Could not cleanup cloud provider:%s error: %v", manager.GetCloudProviderID(), err)
		} else if r.Error != nil {
			log.Printf("Cloud provider:%s cleanup return error: %v", manager.GetCloudProviderID(), err)
		}

		return grpc.GetManager().Close()
	}

	return err
}

// Name returns name of the cloud provider.
func (grpc *grpcCloudProvider) Name() string {
	return ProviderName
}

// NodeGroups returns all node groups configured for this cloud provider.
func (grpc *grpcCloudProvider) NodeGroups() []cloudprovider.NodeGroup {
	manager := grpc.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	cloudProviderService, err := manager.GetCloudProviderServiceClient()

	if err == nil {
		r, err := cloudProviderService.NodeGroups(ctx, &CloudProviderServiceRequest{ProviderID: manager.GetCloudProviderID()})

		if err != nil {
			log.Printf("Could not get NodeGroupForNode for cloud provider:%s error: %v", manager.GetCloudProviderID(), err)
		} else {
			var nodeGroups []cloudprovider.NodeGroup

			for _, nodeGroup := range r.GetNodeGroups() {
				nodeGroups = append(nodeGroups, &GrpcNodeGroup{
					name:    nodeGroup.GetId(),
					manager: manager,
				})
			}

			return nodeGroups
		}
	}

	return nil
}

// NodeGroupForNode returns the node group for the given node.
func (grpc *grpcCloudProvider) NodeGroupForNode(node *apiv1.Node) (cloudprovider.NodeGroup, error) {
	manager := grpc.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	cloudProviderService, err := manager.GetCloudProviderServiceClient()

	if err == nil {
		r, err := cloudProviderService.NodeGroupForNode(ctx, &NodeGroupForNodeRequest{ProviderID: manager.GetCloudProviderID(), Node: toJSON(node)})

		if err != nil {
			log.Printf("Could not get NodeGroupForNode for cloud provider:%s error: %v", manager.GetCloudProviderID(), err)

			return nil, err
		} else if rerr := r.GetError(); rerr != nil {
			log.Printf("Cloud provider:%s call NodeGroupForNode got error: %v", manager.GetCloudProviderID(), rerr)

			return nil, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		} else if r.GetNodeGroup() == nil || len(r.GetNodeGroup().GetId()) == 0 {
			log.Printf("NodeGroup for node:%s not found", node.Spec.ProviderID)
			return nil, nil
		}

		return &GrpcNodeGroup{
			name:    r.GetNodeGroup().GetId(),
			manager: manager,
		}, nil
	}

	return nil, err
}

// Pricing returns pricing model for this cloud provider or error if not available.
func (grpc *grpcCloudProvider) Pricing() (cloudprovider.PricingModel, errors.AutoscalerError) {
	manager := grpc.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	cloudProviderService, err := manager.GetCloudProviderServiceClient()

	if err == nil {
		r, err := cloudProviderService.Pricing(ctx, &CloudProviderServiceRequest{ProviderID: manager.GetCloudProviderID()})

		if err != nil {
			log.Printf("Could not get Pricing for cloud provider:%s error: %v", manager.GetCloudProviderID(), err)

			return nil, errors.ToAutoscalerError(errors.InternalError, err)
		} else if rerr := r.GetError(); rerr != nil {
			log.Printf("Cloud provider:%s call Pricing got error: %v", manager.GetCloudProviderID(), rerr)

			return nil, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}

		return &GrpcPriceModel{
			name:    r.GetPriceModel().GetId(),
			manager: manager,
		}, nil
	}

	return nil, errors.NewAutoscalerError(errors.ApiCallError, err.Error())
}

// GetAvailableMachineTypes get all machine types that can be requested from the cloud provider.
func (grpc *grpcCloudProvider) GetAvailableMachineTypes() ([]string, error) {
	manager := grpc.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	cloudProviderService, err := manager.GetCloudProviderServiceClient()

	if err == nil {
		r, err := cloudProviderService.GetAvailableMachineTypes(ctx, &CloudProviderServiceRequest{ProviderID: manager.GetCloudProviderID()})

		if err != nil {
			log.Printf("Could not get GetAvailableMachineTypes for cloud provider:%s error: %v", manager.GetCloudProviderID(), err)

			return nil, err
		} else if rerr := r.GetError(); rerr != nil {
			log.Printf("Cloud provider:%s call GetAvailableMachineTypes got error: %v", manager.GetCloudProviderID(), rerr)

			return nil, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}

		return r.GetAvailableMachineTypes().GetMachineType(), nil
	}

	return nil, err
}

// NewNodeGroup builds a theoretical node group based on the node definition provided. The node group is not automatically
// created on the cloud provider side. The node group is not returned by NodeGroups() until it is created.
func (grpc *grpcCloudProvider) NewNodeGroup(machineType string, labels map[string]string, systemLabels map[string]string,
	taints []apiv1.Taint, extraResources map[string]resource.Quantity) (cloudprovider.NodeGroup, error) {

	manager := grpc.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	var taintsRequest []*v1.Taint

	if taints != nil {
		taintsRequest = make([]*v1.Taint, len(taints))

		for i := range taints {
			taintsRequest = append(taintsRequest, &taints[i])
		}
	}

	var extraResourcesRequest map[string]string

	if extraResources != nil && len(extraResources) > 0 {
		extraResourcesRequest = make(map[string]string)

		for key, value := range extraResources {
			extraResourcesRequest[key] = value.String()
		}
	}

	cloudProviderService, err := manager.GetCloudProviderServiceClient()

	if err == nil {
		for index, ng := range grpc.nodeGroups {
			// Find a group not already provisionned
			if ng.Provisionned == false {
				r, err := cloudProviderService.NewNodeGroup(ctx, &NewNodeGroupRequest{
					ProviderID:     manager.GetCloudProviderID(),
					MachineType:    machineType,
					NodeGroupID:    ng.NodeGroupID,
					MinNodeSize:    int32(ng.MinSize),
					MaxNodeSize:    int32(ng.MaxSize),
					Labels:         labels,
					SystemLabels:   systemLabels,
					Taints:         taintsRequest,
					ExtraResources: extraResourcesRequest,
				})

				if err != nil {
					log.Printf("Could not get NewNodeGroup for cloud provider:%s error: %v", manager.GetCloudProviderID(), err)

					return nil, err
				} else if rerr := r.GetError(); rerr != nil {
					log.Printf("Cloud provider:%s call NewNodeGroup got error: %v", manager.GetCloudProviderID(), rerr)

					return nil, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
				}

				grpc.nodeGroups[index].Provisionned = true

				return &GrpcNodeGroup{
					name:    r.GetNodeGroup().GetId(),
					manager: manager,
				}, nil
			}

			log.Printf("All nodes group altready provisionned")

			err = errors.NewAutoscalerError(errors.InternalError, "All node group already provisionned")
		}
	}

	return nil, err
}

// GetResourceLimiter returns struct containing limits (max, min) for resources (cores, memory etc.).
func (grpc *grpcCloudProvider) GetResourceLimiter() (*cloudprovider.ResourceLimiter, error) {
	//return grpc.resourceLimiter, nil
	manager := grpc.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	cloudProviderService, err := manager.GetCloudProviderServiceClient()

	if err == nil {
		r, err := cloudProviderService.GetResourceLimiter(ctx, &CloudProviderServiceRequest{ProviderID: manager.GetCloudProviderID()})

		if err != nil {
			log.Printf("Could not get GetResourceLimiter for cloud provider:%s error: %v", manager.GetCloudProviderID(), err)

			return nil, err
		} else if rerr := r.GetError(); rerr != nil {
			log.Printf("Cloud provider:%s call GetResourceLimiter got error: %v", manager.GetCloudProviderID(), rerr)

			return nil, errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}

		resourceLimiter := r.GetResourceLimiter()

		return cloudprovider.NewResourceLimiter(resourceLimiter.GetMinLimits(), resourceLimiter.GetMaxLimits()), nil
	}

	return nil, err
}

// GPULabel returns the label added to nodes with GPU resource.
func (grpc *grpcCloudProvider) GPULabel() string {
	manager := grpc.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	cloudProviderService, err := manager.GetCloudProviderServiceClient()

	if err == nil {
		r, err := cloudProviderService.GPULabel(ctx, &CloudProviderServiceRequest{ProviderID: manager.GetCloudProviderID()})

		if err != nil {
			log.Printf("Could not get GPULabel for cloud provider:%s error: %v", manager.GetCloudProviderID(), err)

			return ""
		}

		return r.GetGpulabel()
	}

	return ""
}

// GetAvailableGPUTypes return all available GPU types cloud provider supports.
func (grpc *grpcCloudProvider) GetAvailableGPUTypes() map[string]struct{} {
	manager := grpc.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	cloudProviderService, err := manager.GetCloudProviderServiceClient()

	if err == nil {
		r, err := cloudProviderService.GetAvailableGPUTypes(ctx, &CloudProviderServiceRequest{ProviderID: manager.GetCloudProviderID()})

		if err != nil {
			log.Printf("Could not get GetAvailableGPUTypes for cloud provider:%s error: %v", manager.GetCloudProviderID(), err)

			return nil
		}

		availGpuTypes := make(map[string]struct{})

		for name, value := range r.GetAvailableGpuTypes() {
			availGpuTypes[name] = fromJSON(value)
		}

		return availGpuTypes
	}

	return nil
}

// Refresh is called before every main loop and can be used to dynamically update cloud provider state.
// In particular the list of node groups returned by NodeGroups can change as a result of CloudProvider.Refresh().
func (grpc *grpcCloudProvider) Refresh() error {
	manager := grpc.GetManager()

	manager.Lock()
	defer manager.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), manager.GetGrpcTimeout())
	defer cancel()

	cloudProviderService, err := manager.GetCloudProviderServiceClient()

	if err == nil {
		r, err := cloudProviderService.Refresh(ctx, &CloudProviderServiceRequest{ProviderID: manager.GetCloudProviderID()})

		if err != nil {
			log.Printf("Could not get Refresh for cloud provider:%s error: %v", manager.GetCloudProviderID(), err)

			return err
		} else if rerr := r.GetError(); rerr != nil {
			log.Printf("Cloud provider:%s call Refresh got error: %v", manager.GetCloudProviderID(), rerr)

			return errors.NewAutoscalerError((errors.AutoscalerErrorType)(rerr.Code), rerr.Reason)
		}

		return nil
	}

	return err
}

// buildGrpcCloudProvider builds external CloudProvider implementation for GRPc.
func newGrpcCloudProvider(grpcManager *GrpcManager, resourceLimiter *cloudprovider.ResourceLimiter, nodeGroups []*NodeGroupDef) (cloudprovider.CloudProvider, error) {

	grpc := &grpcCloudProvider{
		manager:         grpcManager,
		resourceLimiter: resourceLimiter,
		nodeGroups:      nodeGroups,
	}

	if err := grpcManager.connect(); err != nil {
		return nil, err
	}

	return grpc, nil
}

func parseNodeGroupDefs(nodeGroupDefs []string) ([]*NodeGroupDef, error) {
	result := make([]*NodeGroupDef, 0, len(nodeGroupDefs))

	for _, s := range nodeGroupDefs {
		var tailStr string
		ng := &NodeGroupDef{
			Labels: make(map[string]string),
		}

		if _, err := fmt.Sscanf(s, "%d:%d:%t/%s", &ng.MinSize, &ng.MaxSize, &ng.IncludeExistingNode, &tailStr); err != nil {
			return nil, err
		}

		labels := strings.Split(tailStr, "|")

		ng.NodeGroupID = labels[0]

		if len(labels) > 1 {
			labels := strings.Split(tailStr, ";")

			for i := 1; i < len(labels); i++ {
				values := strings.Split(labels[i], "=")

				if len(values) != 2 {
					return nil, fmt.Errorf("Misformatted label definition: %s extracted from:%s", tailStr, labels[i])
				}

				ng.Labels[values[0]] = values[1]
			}
		}

		if ng.MaxSize < ng.MinSize {
			return nil, fmt.Errorf("MaxSize < MinSize")
		}

		if ng.MaxSize < 0 {
			return nil, fmt.Errorf("MaxSize < 0")
		}

		if ng.MinSize < 0 {
			return nil, fmt.Errorf("MinSize < 0")
		}

		if len(ng.NodeGroupID) == 0 {
			return nil, fmt.Errorf("NodeGroupID is empty")
		}

		result = append(result, ng)
	}

	return result, nil
}

// BuildGrpc builds an external cloud provider piloted thru Grpc, manager etc.
func BuildGrpc(opts config.AutoscalingOptions, do cloudprovider.NodeGroupDiscoveryOptions, rl *cloudprovider.ResourceLimiter) cloudprovider.CloudProvider {
	var config io.ReadCloser

	if opts.CloudConfig != "" {
		var err error
		config, err = os.Open(opts.CloudConfig)

		if err != nil {
			klog.Fatalf("Couldn't open cloud provider configuration %s: %#v", opts.CloudConfig, err)
		}

		defer config.Close()
	}

	if nodeGroups, err := parseNodeGroupDefs(opts.NodeGroups); err == nil {
		if manager, err := newGrpcManager(config, opts, do, rl, nodeGroups); err == nil {
			if provider, err := newGrpcCloudProvider(manager, rl, nodeGroups); err == nil {
				return provider
			} else {
				klog.Fatalf("Failed to create GRPC cloud provider: %v", err)
			}
		} else {
			klog.Fatalf("Failed to create GRPC cloud provider: %v", err)
		}
	} else {
		klog.Fatalf("Failed to create GRPC Manager: %v", err)
	}

	return nil
}
