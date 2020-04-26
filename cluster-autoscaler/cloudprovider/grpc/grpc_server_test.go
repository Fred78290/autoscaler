package grpccloudprovider

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	apiv1 "k8s.io/api/core/v1"
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
	wg              sync.WaitGroup
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
	log.Printf("Call server Connect: %v", request)

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
	log.Printf("Call server Name: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &NameReply{
		Name: ProviderName,
	}, nil
}

func (s *grpcServer) NodeGroups(ctx context.Context, request *CloudProviderServiceRequest) (*NodeGroupsReply, error) {
	log.Printf("Call server NodeGroups: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &NodeGroupsReply{
		NodeGroups: []*NodeGroup{
			&NodeGroup{
				Id: s.nodeGroupID,
			},
		},
	}, nil
}

func (s *grpcServer) NodeGroupForNode(ctx context.Context, request *NodeGroupForNodeRequest) (*NodeGroupForNodeReply, error) {
	log.Printf("Call server NodeGroupForNode: %v", request)

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
	log.Printf("Call server Pricing: %v", request)

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
	log.Printf("Call server GetAvailableMachineTypes: %v", request)

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
	log.Printf("Call server NewNodeGroup: %v", request)

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
	log.Printf("Call server GetResourceLimiter: %v", request)

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
	log.Printf("Call server GPULabel: %v", request)

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
	log.Printf("Call server GPULabel: %v", request)

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
	log.Printf("Call server Cleanup: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &CleanupReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) Refresh(ctx context.Context, request *CloudProviderServiceRequest) (*RefreshReply, error) {
	log.Printf("Call server Refresh: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &RefreshReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) MaxSize(ctx context.Context, request *NodeGroupServiceRequest) (*MaxSizeReply, error) {
	log.Printf("Call server MaxSize: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &MaxSizeReply{
		MaxSize: s.maxNodeSize,
	}, nil
}

func (s *grpcServer) MinSize(ctx context.Context, request *NodeGroupServiceRequest) (*MinSizeReply, error) {
	log.Printf("Call server MinSize: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &MinSizeReply{
		MinSize: s.minNodeSize,
	}, nil
}

func (s *grpcServer) TargetSize(ctx context.Context, request *NodeGroupServiceRequest) (*TargetSizeReply, error) {
	log.Printf("Call server TargetSize: %v", request)

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
	log.Printf("Call server IncreaseSize: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &IncreaseSizeReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) DeleteNodes(ctx context.Context, request *DeleteNodesRequest) (*DeleteNodesReply, error) {
	log.Printf("Call server DeleteNodes: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &DeleteNodesReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) DecreaseTargetSize(ctx context.Context, request *DecreaseTargetSizeRequest) (*DecreaseTargetSizeReply, error) {
	log.Printf("Call server DecreaseTargetSize: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &DecreaseTargetSizeReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) Id(ctx context.Context, request *NodeGroupServiceRequest) (*IdReply, error) {
	log.Printf("Call server Id: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &IdReply{
		Response: "test",
	}, nil
}

func (s *grpcServer) Debug(ctx context.Context, request *NodeGroupServiceRequest) (*DebugReply, error) {
	log.Printf("Call server Debug: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &DebugReply{
		Response: "test",
	}, nil
}

func (s *grpcServer) Nodes(ctx context.Context, request *NodeGroupServiceRequest) (*NodesReply, error) {
	log.Printf("Call server Nodes: %v", request)

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
	log.Printf("Call server TemplateNodeInfo: %v", request)

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
	log.Printf("Call server Exist: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &ExistReply{
		//Exists: request.NodeGroupID == providerIDForNode(testGroupID, testNodeName),
		Exists: request.NodeGroupID == testGroupID,
	}, nil
}

func (s *grpcServer) Create(ctx context.Context, request *NodeGroupServiceRequest) (*CreateReply, error) {
	log.Printf("Call server Create: %v", request)

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
	log.Printf("Call server Delete: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &DeleteReply{
		Error: nil,
	}, nil
}

func (s *grpcServer) Autoprovisioned(ctx context.Context, request *NodeGroupServiceRequest) (*AutoprovisionedReply, error) {
	log.Printf("Call server Autoprovisioned: %v", request)

	if request.GetProviderID() != serverProviderIdentifier {
		return nil, ErrMismatchingProvider
	}

	return &AutoprovisionedReply{
		Autoprovisioned: true,
	}, nil
}

func (s *grpcServer) Belongs(ctx context.Context, request *BelongsRequest) (*BelongsReply, error) {
	log.Printf("Call server Belongs: %v", request)

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
	log.Printf("Call server NodePrice: %v", request)

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
	log.Printf("Call server PodPrice: %v", request)

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
	log.Printf("Stop listening test server")

	testServer.GracefulStop()
}

func startTestServer(wg *sync.WaitGroup) {
	log.Printf("Start listening test server")

	lis, err := net.Listen("unix", "/tmp/cluster-autoscaler-grpc.sock")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	testServer = grpc.NewServer()

	RegisterCloudProviderServiceServer(testServer, testGrpcServer)
	RegisterNodeGroupServiceServer(testServer, testGrpcServer)
	RegisterPricingModelServiceServer(testServer, testGrpcServer)

	reflection.Register(testServer)

	wg.Done()

	if err := testServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("End listening test server")
}
