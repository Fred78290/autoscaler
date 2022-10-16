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
	"encoding/json"
	"io"
	"sync"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/config"
	errors "k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	klog "k8s.io/klog/v2"
)

// GrpcManager is handles grpc communication and data caching
type GrpcManager struct {
	sync.Mutex
	resourceLimiter            *cloudprovider.ResourceLimiter
	config                     GrpcConfig
	autoscalingOptions         config.AutoscalingOptions
	nodes                      []*NodeGroupDef
	connection                 *grpc.ClientConn
	cloudProviderServiceClient CloudProviderServiceClient
	nodeGroupServiceClient     NodeGroupServiceClient
	pricingModelServiceClient  PricingModelServiceClient
}

// Close the grpc connection
func (t *GrpcManager) Close() error {
	return t.connection.Close()
}

func (t *GrpcManager) isConnected() bool {
	if t.connection == nil {
		return false
	}

	switch t.connection.GetState() {
	case connectivity.Idle:
		return false
	case connectivity.Connecting:
		return true
	case connectivity.Ready:
		return true
	case connectivity.TransientFailure:
		return false
	case connectivity.Shutdown:
		return false
	default:
		grpclog.Errorf("unknown connectivity state: %d", t.connection.GetState())
		return false
	}
}

// GetCloudProviderServiceClient return nodeGroupServiceClient
func (t *GrpcManager) GetCloudProviderServiceClient() (CloudProviderServiceClient, error) {

	if !t.isConnected() {
		if err := t.connect(); err != nil {
			return nil, err
		}
	}

	return t.cloudProviderServiceClient, nil
}

// GetNodeGroupServiceClient return nodeGroupServiceClient
func (t *GrpcManager) GetNodeGroupServiceClient() (NodeGroupServiceClient, error) {
	if !t.isConnected() {
		if err := t.connect(); err != nil {
			return nil, err
		}
	}

	return t.nodeGroupServiceClient, nil
}

// GetPricingModelServiceClient return pricingModelServiceClient
func (t *GrpcManager) GetPricingModelServiceClient() (PricingModelServiceClient, error) {
	if !t.isConnected() {
		if err := t.connect(); err != nil {
			return nil, err
		}
	}

	return t.pricingModelServiceClient, nil
}

// GetConfig return grpc config
func (t *GrpcManager) GetConfig() *GrpcConfig {
	return &t.config
}

// GetGrpcServerAddress returns the address to connect
func (t *GrpcManager) GetGrpcServerAddress() string {
	return t.GetConfig().GetAddress()
}

// GetGrpcTimeout return grpc connection timeout
func (t *GrpcManager) GetGrpcTimeout() time.Duration {
	return t.GetConfig().GetTimeout()
}

// GetCloudProviderID returns the cloud provider identifier
func (t *GrpcManager) GetCloudProviderID() string {
	return t.GetConfig().GetIdentifier()
}

func (t *GrpcManager) connect() error {
	var err error
	var reply *ConnectReply

	// Set up a connection to the server.
	options := grpc.WithTransportCredentials(insecure.NewCredentials())

	t.connection, err = grpc.Dial(t.GetGrpcServerAddress(), options)
	if err != nil {
		klog.Errorf("did not connect to %s Error: %v", t.GetGrpcServerAddress(), err)

		return err
	}

	t.cloudProviderServiceClient = NewCloudProviderServiceClient(t.connection)
	t.nodeGroupServiceClient = NewNodeGroupServiceClient(t.connection)
	t.pricingModelServiceClient = NewPricingModelServiceClient(t.connection)

	ctx, cancel := context.WithTimeout(context.Background(), t.GetGrpcTimeout())
	defer cancel()

	resourceLimiter := toResourceLimiter(t.resourceLimiter)

	reply, err = t.cloudProviderServiceClient.Connect(ctx, &ConnectRequest{
		ProviderID:           t.GetCloudProviderID(),
		ResourceLimiter:      resourceLimiter,
		KubeAdmConfiguration: t.config.KubeAdmConfiguration,
		Nodes:                t.nodes,
		AutoProvisionned:     t.autoscalingOptions.NodeAutoprovisioningEnabled,
	})

	if err != nil {
		klog.Errorf("did not connect: %v", err)
	} else if reply.GetError() != nil {
		err = errors.NewAutoscalerError(errors.CloudProviderError, reply.GetError().GetReason())

		klog.Errorf("did not connect: %v", err)
	} else if !reply.GetConnected() {
		klog.Errorf("Unable to connect")

		err = errors.NewAutoscalerError(errors.CloudProviderError, "Unable to connect")
	}

	return err
}

// newGrpcManager constructs grpcManager object.
func newGrpcManager(configReader io.Reader, opts config.AutoscalingOptions, discoveryOpts cloudprovider.NodeGroupDiscoveryOptions, rl *cloudprovider.ResourceLimiter, nodeGroups []*NodeGroupDef) (*GrpcManager, error) {
	var cfg GrpcConfig

	if configReader != nil {
		decoder := json.NewDecoder(configReader)

		if err := decoder.Decode(&cfg); err != nil {
			klog.Errorf("Couldn't read config: %v", err)
			return nil, err
		}

	} else {
		klog.Errorf("Any GRPC config defined")

		return nil, ErrMissingConfig
	}

	manager := &GrpcManager{
		resourceLimiter:    rl,
		config:             cfg,
		autoscalingOptions: opts,
		nodes:              nodeGroups,
	}

	return manager, nil
}
