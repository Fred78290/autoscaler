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
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	klog "k8s.io/klog/v2"
)

const serverProviderIdentifier = "multipass"

type serverResourceLimiter struct {
	minLimits map[string]int64
	maxLimits map[string]int64
}

type grpcServer struct {
	resourceLimiter serverResourceLimiter
	nodeGroupID     string
	minNodeSize     int32
	maxNodeSize     int32
}

var testServer *grpc.Server
var testGrpcServer = &grpcServer{
	nodeGroupID: "ca-grpc-multipass",
	minNodeSize: 0,
	maxNodeSize: 5,
}

func providerID(groupID string) string {
	return fmt.Sprintf("%s://%s/object?type=group", testProviderID, groupID)
}

func providerIDForNode(groupID, nodeName string) string {
	return fmt.Sprintf("%s://%s/object?type=node&name=%s", testProviderID, groupID, nodeName)
}

func (s *grpcServer) Connect(ctx context.Context, request *ConnectRequest) (*ConnectReply, error) {
	klog.V(4).Infof("Call server Connect: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	s.resourceLimiter.minLimits = request.ResourceLimiter.MinLimits
	s.resourceLimiter.maxLimits = request.ResourceLimiter.MaxLimits

	return &ConnectReply{
		Response: &ConnectReply_Connected{
			Connected: true,
		},
	}, nil
}

func (s *grpcServer) Name(ctx context.Context, request *CloudProviderServiceRequest) (*NameReply, error) {
	klog.V(4).Infof("Call server Name: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &NameReply{
		Name: cloudprovider.GrpcProviderName,
	}, nil
}

func (s *grpcServer) NodeGroups(ctx context.Context, request *CloudProviderServiceRequest) (*NodeGroupsReply, error) {
	klog.V(4).Infof("Call server NodeGroups: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &NodeGroupsReply{
		NodeGroups: []*NodeGroup{
			{
				Id: s.nodeGroupID,
			},
		},
	}, nil
}

func (s *grpcServer) NodeGroupForNode(ctx context.Context, request *NodeGroupForNodeRequest) (*NodeGroupForNodeReply, error) {
	klog.V(4).Infof("Call server NodeGroupForNode: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	node, err := NodeFromJSON(request.GetNode())

	if err != nil {
		return nil, err
	}

	nodeID := providerIDForNode(testGroupID, testNodeName)

	if node.Spec.ProviderID != nodeID {
		return &NodeGroupForNodeReply{
			Response: &NodeGroupForNodeReply_Error{
				Error: &Error{
					Code:   "cloudProviderError",
					Reason: "Node not found",
				},
			},
		}, nil
	}

	return &NodeGroupForNodeReply{
		Response: &NodeGroupForNodeReply_NodeGroup{
			NodeGroup: &NodeGroup{
				Id: s.nodeGroupID,
			},
		},
	}, nil
}

func (s *grpcServer) Pricing(ctx context.Context, request *CloudProviderServiceRequest) (*PricingModelReply, error) {
	klog.V(4).Infof("Call server Pricing: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &PricingModelReply{
		Response: &PricingModelReply_PriceModel{
			PriceModel: &PricingModel{
				Id: testProviderID,
			},
		},
	}, nil
}

func (s *grpcServer) GetAvailableMachineTypes(ctx context.Context, request *CloudProviderServiceRequest) (*AvailableMachineTypesReply, error) {
	klog.V(4).Infof("Call server GetAvailableMachineTypes: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &AvailableMachineTypesReply{
		Response: &AvailableMachineTypesReply_AvailableMachineTypes{
			AvailableMachineTypes: &AvailableMachineTypes{
				MachineType: []string{"tiny"},
			},
		},
	}, nil
}

func (s *grpcServer) NewNodeGroup(ctx context.Context, request *NewNodeGroupRequest) (*NewNodeGroupReply, error) {
	klog.V(4).Infof("Call server NewNodeGroup: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	if request.GetMachineType() != "medium" {
		return &NewNodeGroupReply{
			Response: &NewNodeGroupReply_Error{
				Error: &Error{
					Code:   "cloudProviderError",
					Reason: "Wrong machine type",
				},
			},
		}, nil
	}

	s.nodeGroupID = request.NodeGroupID
	s.minNodeSize = request.MinNodeSize
	s.maxNodeSize = request.MaxNodeSize

	return &NewNodeGroupReply{
		Response: &NewNodeGroupReply_NodeGroup{
			NodeGroup: &NodeGroup{
				Id: s.nodeGroupID,
			},
		},
	}, nil
}

func (s *grpcServer) GetResourceLimiter(ctx context.Context, request *CloudProviderServiceRequest) (*ResourceLimiterReply, error) {
	klog.V(4).Infof("Call server GetResourceLimiter: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &ResourceLimiterReply{
		Response: &ResourceLimiterReply_ResourceLimiter{
			ResourceLimiter: &ResourceLimiter{
				MinLimits: s.resourceLimiter.minLimits,
				MaxLimits: s.resourceLimiter.maxLimits,
			},
		},
	}, nil
}

func (s *grpcServer) GPULabel(ctx context.Context, request *CloudProviderServiceRequest) (*GPULabelReply, error) {
	klog.V(4).Infof("Call server GPULabel: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &GPULabelReply{
		Response: &GPULabelReply_Gpulabel{
			Gpulabel: "turing",
		},
	}, nil
}

func (s *grpcServer) GetAvailableGPUTypes(ctx context.Context, request *CloudProviderServiceRequest) (*GetAvailableGPUTypesReply, error) {
	klog.V(4).Infof("Call server GPULabel: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	gpus := make(map[string]string)

	for name, value := range availableGPUTypes {
		gpus[name] = toJSON(value)
	}

	return &GetAvailableGPUTypesReply{
		AvailableGpuTypes: gpus,
	}, nil
}

func (s *grpcServer) Cleanup(ctx context.Context, request *CloudProviderServiceRequest) (*CleanupReply, error) {
	klog.V(4).Infof("Call server Cleanup: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &CleanupReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) Refresh(ctx context.Context, request *CloudProviderServiceRequest) (*RefreshReply, error) {
	klog.V(4).Infof("Call server Refresh: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &RefreshReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) MaxSize(ctx context.Context, request *NodeGroupServiceRequest) (*MaxSizeReply, error) {
	klog.V(4).Infof("Call server MaxSize: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &MaxSizeReply{
		MaxSize: s.maxNodeSize,
	}, nil
}

func (s *grpcServer) MinSize(ctx context.Context, request *NodeGroupServiceRequest) (*MinSizeReply, error) {
	klog.V(4).Infof("Call server MinSize: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &MinSizeReply{
		MinSize: s.minNodeSize,
	}, nil
}

func (s *grpcServer) TargetSize(ctx context.Context, request *NodeGroupServiceRequest) (*TargetSizeReply, error) {
	klog.V(4).Infof("Call server TargetSize: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &TargetSizeReply{
		Response: &TargetSizeReply_TargetSize{
			TargetSize: 0,
		},
	}, nil
}

func (s *grpcServer) IncreaseSize(ctx context.Context, request *IncreaseSizeRequest) (*IncreaseSizeReply, error) {
	klog.V(4).Infof("Call server IncreaseSize: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &IncreaseSizeReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) DeleteNodes(ctx context.Context, request *DeleteNodesRequest) (*DeleteNodesReply, error) {
	klog.V(4).Infof("Call server DeleteNodes: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &DeleteNodesReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) DecreaseTargetSize(ctx context.Context, request *DecreaseTargetSizeRequest) (*DecreaseTargetSizeReply, error) {
	klog.V(4).Infof("Call server DecreaseTargetSize: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &DecreaseTargetSizeReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) Id(ctx context.Context, request *NodeGroupServiceRequest) (*IdReply, error) {
	klog.V(4).Infof("Call server Id: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &IdReply{
		Response: "test",
	}, nil
}

func (s *grpcServer) Debug(ctx context.Context, request *NodeGroupServiceRequest) (*DebugReply, error) {
	klog.V(4).Infof("Call server Debug: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &DebugReply{
		Response: "test",
	}, nil
}

func (s *grpcServer) Nodes(ctx context.Context, request *NodeGroupServiceRequest) (*NodesReply, error) {
	klog.V(4).Infof("Call server Nodes: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	instances := &Instance{
		Id: testGroupID,
		Status: &InstanceStatus{
			State:     InstanceState_STATE_RUNNING,
			ErrorInfo: nil,
		},
	}

	return &NodesReply{
		Response: &NodesReply_Instances{
			Instances: &Instances{
				Items: []*Instance{instances},
			},
		},
	}, nil
}

func (s *grpcServer) TemplateNodeInfo(ctx context.Context, request *NodeGroupServiceRequest) (*TemplateNodeInfoReply, error) {
	klog.V(4).Infof("Call server TemplateNodeInfo: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	node := &apiv1.Node{
		Spec: apiv1.NodeSpec{
			ProviderID: providerIDForNode(testGroupID, testNodeName),
		},
	}

	return &TemplateNodeInfoReply{
		Response: &TemplateNodeInfoReply_NodeInfo{NodeInfo: &NodeInfo{
			Node: toJSON(node),
		}},
	}, nil
}

func (s *grpcServer) Exist(ctx context.Context, request *NodeGroupServiceRequest) (*ExistReply, error) {
	klog.V(4).Infof("Call server Exist: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &ExistReply{
		//Exists: request.NodeGroupID == providerIDForNode(testGroupID, testNodeName),
		Exists: request.NodeGroupID == testGroupID,
	}, nil
}

func (s *grpcServer) Create(ctx context.Context, request *NodeGroupServiceRequest) (*CreateReply, error) {
	klog.V(4).Infof("Call server Create: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &CreateReply{
		Response: &CreateReply_NodeGroup{
			NodeGroup: &NodeGroup{
				Id: s.nodeGroupID,
			},
		},
	}, nil
}

func (s *grpcServer) Delete(ctx context.Context, request *NodeGroupServiceRequest) (*DeleteReply, error) {
	klog.V(4).Infof("Call server Delete: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &DeleteReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) Autoprovisioned(ctx context.Context, request *NodeGroupServiceRequest) (*AutoprovisionedReply, error) {
	klog.V(4).Infof("Call server Autoprovisioned: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &AutoprovisionedReply{
		Autoprovisioned: true,
	}, nil
}

func (s *grpcServer) Belongs(ctx context.Context, request *BelongsRequest) (*BelongsReply, error) {
	klog.V(4).Infof("Call server Belongs: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	node, err := NodeFromJSON(request.GetNode())

	if err != nil {
		return &BelongsReply{
			Response: &BelongsReply_Error{
				Error: &Error{
					Code:   "cloundProviderError",
					Reason: "Node not found",
				},
			},
		}, nil
	}

	return &BelongsReply{
		Response: &BelongsReply_Belongs{
			Belongs: node.Spec.ProviderID == providerIDForNode(testGroupID, testNodeName),
		},
	}, nil
}

func (s *grpcServer) NodePrice(ctx context.Context, request *NodePriceRequest) (*NodePriceReply, error) {
	klog.V(4).Infof("Call server NodePrice: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &NodePriceReply{
		Response: &NodePriceReply_Price{
			Price: 0,
		},
	}, nil
}

func (s *grpcServer) PodPrice(ctx context.Context, request *PodPriceRequest) (*PodPriceReply, error) {
	klog.V(4).Infof("Call server PodPrice: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	pod, err := PodFromJSON(request.GetPod())

	if err != nil {
		return &PodPriceReply{
			Response: &PodPriceReply_Error{
				Error: &Error{
					Code:   "cloudProviderError",
					Reason: "Can't unmarshall pod",
				},
			},
		}, nil
	}

	if pod.Spec.NodeName != "test-instance-id" {
		return &PodPriceReply{
			Response: &PodPriceReply_Error{
				Error: &Error{
					Code:   "cloudProviderError",
					Reason: "Pod not found",
				},
			},
		}, nil
	}

	return &PodPriceReply{
		Response: &PodPriceReply_Price{
			Price: 0,
		},
	}, nil
}

func stopTestServer() {
	klog.V(4).Infof("Stop listening test server")

	testServer.GracefulStop()
}

func startTestServer(wg *sync.WaitGroup) {
	klog.V(4).Infof("Start listening test server")

	lis, err := net.Listen("unix", "/tmp/cluster-autoscaler-grpc.sock")

	if err != nil {
		klog.Fatalf("failed to listen: %v", err)
	}

	testServer = grpc.NewServer()

	RegisterCloudProviderServiceServer(testServer, testGrpcServer)
	RegisterNodeGroupServiceServer(testServer, testGrpcServer)
	RegisterPricingModelServiceServer(testServer, testGrpcServer)

	reflection.Register(testServer)

	wg.Done()

	if err := testServer.Serve(lis); err != nil {
		klog.Fatalf("failed to serve: %v", err)
	}

	klog.V(4).Infof("End listening test server")
}
