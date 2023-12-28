// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chimesdkidentity

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// The input parameters don't match the service's restrictions.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The request could not be processed because of conflict in the current state
	// of the resource.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	//
	// The client is permanently forbidden from making the request.
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// One or more of the resources in the request does not exist in the system.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeResourceLimitExceededException for service response error code
	// "ResourceLimitExceededException".
	//
	// The request exceeds the resource limit.
	ErrCodeResourceLimitExceededException = "ResourceLimitExceededException"

	// ErrCodeServiceFailureException for service response error code
	// "ServiceFailureException".
	//
	// The service encountered an unexpected error.
	ErrCodeServiceFailureException = "ServiceFailureException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is currently unavailable.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeThrottledClientException for service response error code
	// "ThrottledClientException".
	//
	// The client exceeded its request rate limit.
	ErrCodeThrottledClientException = "ThrottledClientException"

	// ErrCodeUnauthorizedClientException for service response error code
	// "UnauthorizedClientException".
	//
	// The client is not currently authorized to make the request.
	ErrCodeUnauthorizedClientException = "UnauthorizedClientException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":            newErrorBadRequestException,
	"ConflictException":              newErrorConflictException,
	"ForbiddenException":             newErrorForbiddenException,
	"NotFoundException":              newErrorNotFoundException,
	"ResourceLimitExceededException": newErrorResourceLimitExceededException,
	"ServiceFailureException":        newErrorServiceFailureException,
	"ServiceUnavailableException":    newErrorServiceUnavailableException,
	"ThrottledClientException":       newErrorThrottledClientException,
	"UnauthorizedClientException":    newErrorUnauthorizedClientException,
}
