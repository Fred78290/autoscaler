// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codestarconnectionsiface provides an interface to enable mocking the AWS CodeStar connections service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codestarconnectionsiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/codestarconnections"
)

// CodeStarConnectionsAPI provides an interface to enable mocking the
// codestarconnections.CodeStarConnections service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CodeStar connections.
//    func myFunc(svc codestarconnectionsiface.CodeStarConnectionsAPI) bool {
//        // Make svc.CreateConnection request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := codestarconnections.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodeStarConnectionsClient struct {
//        codestarconnectionsiface.CodeStarConnectionsAPI
//    }
//    func (m *mockCodeStarConnectionsClient) CreateConnection(input *codestarconnections.CreateConnectionInput) (*codestarconnections.CreateConnectionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodeStarConnectionsClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CodeStarConnectionsAPI interface {
	CreateConnection(*codestarconnections.CreateConnectionInput) (*codestarconnections.CreateConnectionOutput, error)
	CreateConnectionWithContext(aws.Context, *codestarconnections.CreateConnectionInput, ...request.Option) (*codestarconnections.CreateConnectionOutput, error)
	CreateConnectionRequest(*codestarconnections.CreateConnectionInput) (*request.Request, *codestarconnections.CreateConnectionOutput)

	CreateHost(*codestarconnections.CreateHostInput) (*codestarconnections.CreateHostOutput, error)
	CreateHostWithContext(aws.Context, *codestarconnections.CreateHostInput, ...request.Option) (*codestarconnections.CreateHostOutput, error)
	CreateHostRequest(*codestarconnections.CreateHostInput) (*request.Request, *codestarconnections.CreateHostOutput)

	DeleteConnection(*codestarconnections.DeleteConnectionInput) (*codestarconnections.DeleteConnectionOutput, error)
	DeleteConnectionWithContext(aws.Context, *codestarconnections.DeleteConnectionInput, ...request.Option) (*codestarconnections.DeleteConnectionOutput, error)
	DeleteConnectionRequest(*codestarconnections.DeleteConnectionInput) (*request.Request, *codestarconnections.DeleteConnectionOutput)

	DeleteHost(*codestarconnections.DeleteHostInput) (*codestarconnections.DeleteHostOutput, error)
	DeleteHostWithContext(aws.Context, *codestarconnections.DeleteHostInput, ...request.Option) (*codestarconnections.DeleteHostOutput, error)
	DeleteHostRequest(*codestarconnections.DeleteHostInput) (*request.Request, *codestarconnections.DeleteHostOutput)

	GetConnection(*codestarconnections.GetConnectionInput) (*codestarconnections.GetConnectionOutput, error)
	GetConnectionWithContext(aws.Context, *codestarconnections.GetConnectionInput, ...request.Option) (*codestarconnections.GetConnectionOutput, error)
	GetConnectionRequest(*codestarconnections.GetConnectionInput) (*request.Request, *codestarconnections.GetConnectionOutput)

	GetHost(*codestarconnections.GetHostInput) (*codestarconnections.GetHostOutput, error)
	GetHostWithContext(aws.Context, *codestarconnections.GetHostInput, ...request.Option) (*codestarconnections.GetHostOutput, error)
	GetHostRequest(*codestarconnections.GetHostInput) (*request.Request, *codestarconnections.GetHostOutput)

	ListConnections(*codestarconnections.ListConnectionsInput) (*codestarconnections.ListConnectionsOutput, error)
	ListConnectionsWithContext(aws.Context, *codestarconnections.ListConnectionsInput, ...request.Option) (*codestarconnections.ListConnectionsOutput, error)
	ListConnectionsRequest(*codestarconnections.ListConnectionsInput) (*request.Request, *codestarconnections.ListConnectionsOutput)

	ListConnectionsPages(*codestarconnections.ListConnectionsInput, func(*codestarconnections.ListConnectionsOutput, bool) bool) error
	ListConnectionsPagesWithContext(aws.Context, *codestarconnections.ListConnectionsInput, func(*codestarconnections.ListConnectionsOutput, bool) bool, ...request.Option) error

	ListHosts(*codestarconnections.ListHostsInput) (*codestarconnections.ListHostsOutput, error)
	ListHostsWithContext(aws.Context, *codestarconnections.ListHostsInput, ...request.Option) (*codestarconnections.ListHostsOutput, error)
	ListHostsRequest(*codestarconnections.ListHostsInput) (*request.Request, *codestarconnections.ListHostsOutput)

	ListHostsPages(*codestarconnections.ListHostsInput, func(*codestarconnections.ListHostsOutput, bool) bool) error
	ListHostsPagesWithContext(aws.Context, *codestarconnections.ListHostsInput, func(*codestarconnections.ListHostsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*codestarconnections.ListTagsForResourceInput) (*codestarconnections.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *codestarconnections.ListTagsForResourceInput, ...request.Option) (*codestarconnections.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*codestarconnections.ListTagsForResourceInput) (*request.Request, *codestarconnections.ListTagsForResourceOutput)

	TagResource(*codestarconnections.TagResourceInput) (*codestarconnections.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *codestarconnections.TagResourceInput, ...request.Option) (*codestarconnections.TagResourceOutput, error)
	TagResourceRequest(*codestarconnections.TagResourceInput) (*request.Request, *codestarconnections.TagResourceOutput)

	UntagResource(*codestarconnections.UntagResourceInput) (*codestarconnections.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *codestarconnections.UntagResourceInput, ...request.Option) (*codestarconnections.UntagResourceOutput, error)
	UntagResourceRequest(*codestarconnections.UntagResourceInput) (*request.Request, *codestarconnections.UntagResourceOutput)

	UpdateHost(*codestarconnections.UpdateHostInput) (*codestarconnections.UpdateHostOutput, error)
	UpdateHostWithContext(aws.Context, *codestarconnections.UpdateHostInput, ...request.Option) (*codestarconnections.UpdateHostOutput, error)
	UpdateHostRequest(*codestarconnections.UpdateHostInput) (*request.Request, *codestarconnections.UpdateHostOutput)
}

var _ CodeStarConnectionsAPI = (*codestarconnections.CodeStarConnections)(nil)
