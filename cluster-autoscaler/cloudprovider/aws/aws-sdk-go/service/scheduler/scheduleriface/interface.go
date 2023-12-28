// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package scheduleriface provides an interface to enable mocking the Amazon EventBridge Scheduler service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package scheduleriface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/scheduler"
)

// SchedulerAPI provides an interface to enable mocking the
// scheduler.Scheduler service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon EventBridge Scheduler.
//	func myFunc(svc scheduleriface.SchedulerAPI) bool {
//	    // Make svc.CreateSchedule request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := scheduler.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockSchedulerClient struct {
//	    scheduleriface.SchedulerAPI
//	}
//	func (m *mockSchedulerClient) CreateSchedule(input *scheduler.CreateScheduleInput) (*scheduler.CreateScheduleOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockSchedulerClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SchedulerAPI interface {
	CreateSchedule(*scheduler.CreateScheduleInput) (*scheduler.CreateScheduleOutput, error)
	CreateScheduleWithContext(aws.Context, *scheduler.CreateScheduleInput, ...request.Option) (*scheduler.CreateScheduleOutput, error)
	CreateScheduleRequest(*scheduler.CreateScheduleInput) (*request.Request, *scheduler.CreateScheduleOutput)

	CreateScheduleGroup(*scheduler.CreateScheduleGroupInput) (*scheduler.CreateScheduleGroupOutput, error)
	CreateScheduleGroupWithContext(aws.Context, *scheduler.CreateScheduleGroupInput, ...request.Option) (*scheduler.CreateScheduleGroupOutput, error)
	CreateScheduleGroupRequest(*scheduler.CreateScheduleGroupInput) (*request.Request, *scheduler.CreateScheduleGroupOutput)

	DeleteSchedule(*scheduler.DeleteScheduleInput) (*scheduler.DeleteScheduleOutput, error)
	DeleteScheduleWithContext(aws.Context, *scheduler.DeleteScheduleInput, ...request.Option) (*scheduler.DeleteScheduleOutput, error)
	DeleteScheduleRequest(*scheduler.DeleteScheduleInput) (*request.Request, *scheduler.DeleteScheduleOutput)

	DeleteScheduleGroup(*scheduler.DeleteScheduleGroupInput) (*scheduler.DeleteScheduleGroupOutput, error)
	DeleteScheduleGroupWithContext(aws.Context, *scheduler.DeleteScheduleGroupInput, ...request.Option) (*scheduler.DeleteScheduleGroupOutput, error)
	DeleteScheduleGroupRequest(*scheduler.DeleteScheduleGroupInput) (*request.Request, *scheduler.DeleteScheduleGroupOutput)

	GetSchedule(*scheduler.GetScheduleInput) (*scheduler.GetScheduleOutput, error)
	GetScheduleWithContext(aws.Context, *scheduler.GetScheduleInput, ...request.Option) (*scheduler.GetScheduleOutput, error)
	GetScheduleRequest(*scheduler.GetScheduleInput) (*request.Request, *scheduler.GetScheduleOutput)

	GetScheduleGroup(*scheduler.GetScheduleGroupInput) (*scheduler.GetScheduleGroupOutput, error)
	GetScheduleGroupWithContext(aws.Context, *scheduler.GetScheduleGroupInput, ...request.Option) (*scheduler.GetScheduleGroupOutput, error)
	GetScheduleGroupRequest(*scheduler.GetScheduleGroupInput) (*request.Request, *scheduler.GetScheduleGroupOutput)

	ListScheduleGroups(*scheduler.ListScheduleGroupsInput) (*scheduler.ListScheduleGroupsOutput, error)
	ListScheduleGroupsWithContext(aws.Context, *scheduler.ListScheduleGroupsInput, ...request.Option) (*scheduler.ListScheduleGroupsOutput, error)
	ListScheduleGroupsRequest(*scheduler.ListScheduleGroupsInput) (*request.Request, *scheduler.ListScheduleGroupsOutput)

	ListScheduleGroupsPages(*scheduler.ListScheduleGroupsInput, func(*scheduler.ListScheduleGroupsOutput, bool) bool) error
	ListScheduleGroupsPagesWithContext(aws.Context, *scheduler.ListScheduleGroupsInput, func(*scheduler.ListScheduleGroupsOutput, bool) bool, ...request.Option) error

	ListSchedules(*scheduler.ListSchedulesInput) (*scheduler.ListSchedulesOutput, error)
	ListSchedulesWithContext(aws.Context, *scheduler.ListSchedulesInput, ...request.Option) (*scheduler.ListSchedulesOutput, error)
	ListSchedulesRequest(*scheduler.ListSchedulesInput) (*request.Request, *scheduler.ListSchedulesOutput)

	ListSchedulesPages(*scheduler.ListSchedulesInput, func(*scheduler.ListSchedulesOutput, bool) bool) error
	ListSchedulesPagesWithContext(aws.Context, *scheduler.ListSchedulesInput, func(*scheduler.ListSchedulesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*scheduler.ListTagsForResourceInput) (*scheduler.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *scheduler.ListTagsForResourceInput, ...request.Option) (*scheduler.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*scheduler.ListTagsForResourceInput) (*request.Request, *scheduler.ListTagsForResourceOutput)

	TagResource(*scheduler.TagResourceInput) (*scheduler.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *scheduler.TagResourceInput, ...request.Option) (*scheduler.TagResourceOutput, error)
	TagResourceRequest(*scheduler.TagResourceInput) (*request.Request, *scheduler.TagResourceOutput)

	UntagResource(*scheduler.UntagResourceInput) (*scheduler.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *scheduler.UntagResourceInput, ...request.Option) (*scheduler.UntagResourceOutput, error)
	UntagResourceRequest(*scheduler.UntagResourceInput) (*request.Request, *scheduler.UntagResourceOutput)

	UpdateSchedule(*scheduler.UpdateScheduleInput) (*scheduler.UpdateScheduleOutput, error)
	UpdateScheduleWithContext(aws.Context, *scheduler.UpdateScheduleInput, ...request.Option) (*scheduler.UpdateScheduleOutput, error)
	UpdateScheduleRequest(*scheduler.UpdateScheduleInput) (*request.Request, *scheduler.UpdateScheduleOutput)
}

var _ SchedulerAPI = (*scheduler.Scheduler)(nil)
