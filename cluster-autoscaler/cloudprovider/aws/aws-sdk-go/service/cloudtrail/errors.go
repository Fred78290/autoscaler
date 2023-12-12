// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeARNInvalidException for service response error code
	// "CloudTrailARNInvalidException".
	//
	// This exception is thrown when an operation is called with an ARN that is
	// not valid.
	//
	// The following is the format of a trail ARN: arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// The following is the format of an event data store ARN: arn:aws:cloudtrail:us-east-2:123456789012:eventdatastore/EXAMPLE-f852-4e8f-8bd1-bcf6cEXAMPLE
	//
	// The following is the format of a channel ARN: arn:aws:cloudtrail:us-east-2:123456789012:channel/01234567890
	ErrCodeARNInvalidException = "CloudTrailARNInvalidException"

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAccessNotEnabledException for service response error code
	// "CloudTrailAccessNotEnabledException".
	//
	// This exception is thrown when trusted access has not been enabled between
	// CloudTrail and Organizations. For more information, see Enabling Trusted
	// Access with Other Amazon Web Services Services (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html)
	// and Prepare For Creating a Trail For Your Organization (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/creating-an-organizational-trail-prepare.html).
	ErrCodeAccessNotEnabledException = "CloudTrailAccessNotEnabledException"

	// ErrCodeAccountHasOngoingImportException for service response error code
	// "AccountHasOngoingImportException".
	//
	// This exception is thrown when you start a new import and a previous import
	// is still in progress.
	ErrCodeAccountHasOngoingImportException = "AccountHasOngoingImportException"

	// ErrCodeAccountNotFoundException for service response error code
	// "AccountNotFoundException".
	//
	// This exception is thrown when the specified account is not found or not part
	// of an organization.
	ErrCodeAccountNotFoundException = "AccountNotFoundException"

	// ErrCodeAccountNotRegisteredException for service response error code
	// "AccountNotRegisteredException".
	//
	// This exception is thrown when the specified account is not registered as
	// the CloudTrail delegated administrator.
	ErrCodeAccountNotRegisteredException = "AccountNotRegisteredException"

	// ErrCodeAccountRegisteredException for service response error code
	// "AccountRegisteredException".
	//
	// This exception is thrown when the account is already registered as the CloudTrail
	// delegated administrator.
	ErrCodeAccountRegisteredException = "AccountRegisteredException"

	// ErrCodeCannotDelegateManagementAccountException for service response error code
	// "CannotDelegateManagementAccountException".
	//
	// This exception is thrown when the management account of an organization is
	// registered as the CloudTrail delegated administrator.
	ErrCodeCannotDelegateManagementAccountException = "CannotDelegateManagementAccountException"

	// ErrCodeChannelARNInvalidException for service response error code
	// "ChannelARNInvalidException".
	//
	// This exception is thrown when the specified value of ChannelARN is not valid.
	ErrCodeChannelARNInvalidException = "ChannelARNInvalidException"

	// ErrCodeChannelAlreadyExistsException for service response error code
	// "ChannelAlreadyExistsException".
	//
	// This exception is thrown when the provided channel already exists.
	ErrCodeChannelAlreadyExistsException = "ChannelAlreadyExistsException"

	// ErrCodeChannelExistsForEDSException for service response error code
	// "ChannelExistsForEDSException".
	//
	// This exception is thrown when the specified event data store cannot yet be
	// deleted because it is in use by a channel.
	ErrCodeChannelExistsForEDSException = "ChannelExistsForEDSException"

	// ErrCodeChannelMaxLimitExceededException for service response error code
	// "ChannelMaxLimitExceededException".
	//
	// This exception is thrown when the maximum number of channels limit is exceeded.
	ErrCodeChannelMaxLimitExceededException = "ChannelMaxLimitExceededException"

	// ErrCodeChannelNotFoundException for service response error code
	// "ChannelNotFoundException".
	//
	// This exception is thrown when CloudTrail cannot find the specified channel.
	ErrCodeChannelNotFoundException = "ChannelNotFoundException"

	// ErrCodeCloudTrailInvalidClientTokenIdException for service response error code
	// "CloudTrailInvalidClientTokenIdException".
	//
	// This exception is thrown when a call results in the InvalidClientTokenId
	// error code. This can occur when you are creating or updating a trail to send
	// notifications to an Amazon SNS topic that is in a suspended Amazon Web Services
	// account.
	ErrCodeCloudTrailInvalidClientTokenIdException = "CloudTrailInvalidClientTokenIdException"

	// ErrCodeCloudWatchLogsDeliveryUnavailableException for service response error code
	// "CloudWatchLogsDeliveryUnavailableException".
	//
	// Cannot set a CloudWatch Logs delivery for this Region.
	ErrCodeCloudWatchLogsDeliveryUnavailableException = "CloudWatchLogsDeliveryUnavailableException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// You are trying to update a resource when another request is in progress.
	// Allow sufficient wait time for the previous request to complete, then retry
	// your request.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// This exception is thrown when the specified resource is not ready for an
	// operation. This can occur when you try to run an operation on a resource
	// before CloudTrail has time to fully load the resource, or because another
	// operation is modifying the resource. If this exception occurs, wait a few
	// minutes, and then try the operation again.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeDelegatedAdminAccountLimitExceededException for service response error code
	// "DelegatedAdminAccountLimitExceededException".
	//
	// This exception is thrown when the maximum number of CloudTrail delegated
	// administrators is reached.
	ErrCodeDelegatedAdminAccountLimitExceededException = "DelegatedAdminAccountLimitExceededException"

	// ErrCodeEventDataStoreARNInvalidException for service response error code
	// "EventDataStoreARNInvalidException".
	//
	// The specified event data store ARN is not valid or does not map to an event
	// data store in your account.
	ErrCodeEventDataStoreARNInvalidException = "EventDataStoreARNInvalidException"

	// ErrCodeEventDataStoreAlreadyExistsException for service response error code
	// "EventDataStoreAlreadyExistsException".
	//
	// An event data store with that name already exists.
	ErrCodeEventDataStoreAlreadyExistsException = "EventDataStoreAlreadyExistsException"

	// ErrCodeEventDataStoreFederationEnabledException for service response error code
	// "EventDataStoreFederationEnabledException".
	//
	// You cannot delete the event data store because Lake query federation is enabled.
	// To delete the event data store, run the DisableFederation operation to disable
	// Lake query federation on the event data store.
	ErrCodeEventDataStoreFederationEnabledException = "EventDataStoreFederationEnabledException"

	// ErrCodeEventDataStoreHasOngoingImportException for service response error code
	// "EventDataStoreHasOngoingImportException".
	//
	// This exception is thrown when you try to update or delete an event data store
	// that currently has an import in progress.
	ErrCodeEventDataStoreHasOngoingImportException = "EventDataStoreHasOngoingImportException"

	// ErrCodeEventDataStoreMaxLimitExceededException for service response error code
	// "EventDataStoreMaxLimitExceededException".
	//
	// Your account has used the maximum number of event data stores.
	ErrCodeEventDataStoreMaxLimitExceededException = "EventDataStoreMaxLimitExceededException"

	// ErrCodeEventDataStoreNotFoundException for service response error code
	// "EventDataStoreNotFoundException".
	//
	// The specified event data store was not found.
	ErrCodeEventDataStoreNotFoundException = "EventDataStoreNotFoundException"

	// ErrCodeEventDataStoreTerminationProtectedException for service response error code
	// "EventDataStoreTerminationProtectedException".
	//
	// The event data store cannot be deleted because termination protection is
	// enabled for it.
	ErrCodeEventDataStoreTerminationProtectedException = "EventDataStoreTerminationProtectedException"

	// ErrCodeImportNotFoundException for service response error code
	// "ImportNotFoundException".
	//
	// The specified import was not found.
	ErrCodeImportNotFoundException = "ImportNotFoundException"

	// ErrCodeInactiveEventDataStoreException for service response error code
	// "InactiveEventDataStoreException".
	//
	// The event data store is inactive.
	ErrCodeInactiveEventDataStoreException = "InactiveEventDataStoreException"

	// ErrCodeInactiveQueryException for service response error code
	// "InactiveQueryException".
	//
	// The specified query cannot be canceled because it is in the FINISHED, FAILED,
	// TIMED_OUT, or CANCELLED state.
	ErrCodeInactiveQueryException = "InactiveQueryException"

	// ErrCodeInsightNotEnabledException for service response error code
	// "InsightNotEnabledException".
	//
	// If you run GetInsightSelectors on a trail or event data store that does not
	// have Insights events enabled, the operation throws the exception InsightNotEnabledException.
	ErrCodeInsightNotEnabledException = "InsightNotEnabledException"

	// ErrCodeInsufficientDependencyServiceAccessPermissionException for service response error code
	// "InsufficientDependencyServiceAccessPermissionException".
	//
	// This exception is thrown when the IAM identity that is used to create the
	// organization resource lacks one or more required permissions for creating
	// an organization resource in a required service.
	ErrCodeInsufficientDependencyServiceAccessPermissionException = "InsufficientDependencyServiceAccessPermissionException"

	// ErrCodeInsufficientEncryptionPolicyException for service response error code
	// "InsufficientEncryptionPolicyException".
	//
	// This exception is thrown when the policy on the S3 bucket or KMS key does
	// not have sufficient permissions for the operation.
	ErrCodeInsufficientEncryptionPolicyException = "InsufficientEncryptionPolicyException"

	// ErrCodeInsufficientS3BucketPolicyException for service response error code
	// "InsufficientS3BucketPolicyException".
	//
	// This exception is thrown when the policy on the S3 bucket is not sufficient.
	ErrCodeInsufficientS3BucketPolicyException = "InsufficientS3BucketPolicyException"

	// ErrCodeInsufficientSnsTopicPolicyException for service response error code
	// "InsufficientSnsTopicPolicyException".
	//
	// This exception is thrown when the policy on the Amazon SNS topic is not sufficient.
	ErrCodeInsufficientSnsTopicPolicyException = "InsufficientSnsTopicPolicyException"

	// ErrCodeInvalidCloudWatchLogsLogGroupArnException for service response error code
	// "InvalidCloudWatchLogsLogGroupArnException".
	//
	// This exception is thrown when the provided CloudWatch Logs log group is not
	// valid.
	ErrCodeInvalidCloudWatchLogsLogGroupArnException = "InvalidCloudWatchLogsLogGroupArnException"

	// ErrCodeInvalidCloudWatchLogsRoleArnException for service response error code
	// "InvalidCloudWatchLogsRoleArnException".
	//
	// This exception is thrown when the provided role is not valid.
	ErrCodeInvalidCloudWatchLogsRoleArnException = "InvalidCloudWatchLogsRoleArnException"

	// ErrCodeInvalidDateRangeException for service response error code
	// "InvalidDateRangeException".
	//
	// A date range for the query was specified that is not valid. Be sure that
	// the start time is chronologically before the end time. For more information
	// about writing a query, see Create or edit a query (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-create-edit-query.html)
	// in the CloudTrail User Guide.
	ErrCodeInvalidDateRangeException = "InvalidDateRangeException"

	// ErrCodeInvalidEventCategoryException for service response error code
	// "InvalidEventCategoryException".
	//
	// Occurs if an event category that is not valid is specified as a value of
	// EventCategory.
	ErrCodeInvalidEventCategoryException = "InvalidEventCategoryException"

	// ErrCodeInvalidEventDataStoreCategoryException for service response error code
	// "InvalidEventDataStoreCategoryException".
	//
	// This exception is thrown when event categories of specified event data stores
	// are not valid.
	ErrCodeInvalidEventDataStoreCategoryException = "InvalidEventDataStoreCategoryException"

	// ErrCodeInvalidEventDataStoreStatusException for service response error code
	// "InvalidEventDataStoreStatusException".
	//
	// The event data store is not in a status that supports the operation.
	ErrCodeInvalidEventDataStoreStatusException = "InvalidEventDataStoreStatusException"

	// ErrCodeInvalidEventSelectorsException for service response error code
	// "InvalidEventSelectorsException".
	//
	// This exception is thrown when the PutEventSelectors operation is called with
	// a number of event selectors, advanced event selectors, or data resources
	// that is not valid. The combination of event selectors or advanced event selectors
	// and data resources is not valid. A trail can have up to 5 event selectors.
	// If a trail uses advanced event selectors, a maximum of 500 total values for
	// all conditions in all advanced event selectors is allowed. A trail is limited
	// to 250 data resources. These data resources can be distributed across event
	// selectors, but the overall total cannot exceed 250.
	//
	// You can:
	//
	//    * Specify a valid number of event selectors (1 to 5) for a trail.
	//
	//    * Specify a valid number of data resources (1 to 250) for an event selector.
	//    The limit of number of resources on an individual event selector is configurable
	//    up to 250. However, this upper limit is allowed only if the total number
	//    of data resources does not exceed 250 across all event selectors for a
	//    trail.
	//
	//    * Specify up to 500 values for all conditions in all advanced event selectors
	//    for a trail.
	//
	//    * Specify a valid value for a parameter. For example, specifying the ReadWriteType
	//    parameter with a value of read-only is not valid.
	ErrCodeInvalidEventSelectorsException = "InvalidEventSelectorsException"

	// ErrCodeInvalidHomeRegionException for service response error code
	// "InvalidHomeRegionException".
	//
	// This exception is thrown when an operation is called on a trail from a Region
	// other than the Region in which the trail was created.
	ErrCodeInvalidHomeRegionException = "InvalidHomeRegionException"

	// ErrCodeInvalidImportSourceException for service response error code
	// "InvalidImportSourceException".
	//
	// This exception is thrown when the provided source S3 bucket is not valid
	// for import.
	ErrCodeInvalidImportSourceException = "InvalidImportSourceException"

	// ErrCodeInvalidInsightSelectorsException for service response error code
	// "InvalidInsightSelectorsException".
	//
	// For PutInsightSelectors, this exception is thrown when the formatting or
	// syntax of the InsightSelectors JSON statement is not valid, or the specified
	// InsightType in the InsightSelectors statement is not valid. Valid values
	// for InsightType are ApiCallRateInsight and ApiErrorRateInsight. To enable
	// Insights on an event data store, the destination event data store specified
	// by the InsightsDestination parameter must log Insights events and the source
	// event data store specified by the EventDataStore parameter must log management
	// events.
	//
	// For UpdateEventDataStore, this exception is thrown if Insights are enabled
	// on the event data store and the updated advanced event selectors are not
	// compatible with the configured InsightSelectors. If the InsightSelectors
	// includes an InsightType of ApiCallRateInsight, the source event data store
	// must log write management events. If the InsightSelectors includes an InsightType
	// of ApiErrorRateInsight, the source event data store must log management events.
	ErrCodeInvalidInsightSelectorsException = "InvalidInsightSelectorsException"

	// ErrCodeInvalidKmsKeyIdException for service response error code
	// "InvalidKmsKeyIdException".
	//
	// This exception is thrown when the KMS key ARN is not valid.
	ErrCodeInvalidKmsKeyIdException = "InvalidKmsKeyIdException"

	// ErrCodeInvalidLookupAttributesException for service response error code
	// "InvalidLookupAttributesException".
	//
	// Occurs when a lookup attribute is specified that is not valid.
	ErrCodeInvalidLookupAttributesException = "InvalidLookupAttributesException"

	// ErrCodeInvalidMaxResultsException for service response error code
	// "InvalidMaxResultsException".
	//
	// This exception is thrown if the limit specified is not valid.
	ErrCodeInvalidMaxResultsException = "InvalidMaxResultsException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// A token that is not valid, or a token that was previously used in a request
	// with different parameters. This exception is thrown if the token is not valid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidParameterCombinationException for service response error code
	// "InvalidParameterCombinationException".
	//
	// This exception is thrown when the combination of parameters provided is not
	// valid.
	ErrCodeInvalidParameterCombinationException = "InvalidParameterCombinationException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// The request includes a parameter that is not valid.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidQueryStatementException for service response error code
	// "InvalidQueryStatementException".
	//
	// The query that was submitted has validation errors, or uses incorrect syntax
	// or unsupported keywords. For more information about writing a query, see
	// Create or edit a query (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-create-edit-query.html)
	// in the CloudTrail User Guide.
	ErrCodeInvalidQueryStatementException = "InvalidQueryStatementException"

	// ErrCodeInvalidQueryStatusException for service response error code
	// "InvalidQueryStatusException".
	//
	// The query status is not valid for the operation.
	ErrCodeInvalidQueryStatusException = "InvalidQueryStatusException"

	// ErrCodeInvalidS3BucketNameException for service response error code
	// "InvalidS3BucketNameException".
	//
	// This exception is thrown when the provided S3 bucket name is not valid.
	ErrCodeInvalidS3BucketNameException = "InvalidS3BucketNameException"

	// ErrCodeInvalidS3PrefixException for service response error code
	// "InvalidS3PrefixException".
	//
	// This exception is thrown when the provided S3 prefix is not valid.
	ErrCodeInvalidS3PrefixException = "InvalidS3PrefixException"

	// ErrCodeInvalidSnsTopicNameException for service response error code
	// "InvalidSnsTopicNameException".
	//
	// This exception is thrown when the provided SNS topic name is not valid.
	ErrCodeInvalidSnsTopicNameException = "InvalidSnsTopicNameException"

	// ErrCodeInvalidSourceException for service response error code
	// "InvalidSourceException".
	//
	// This exception is thrown when the specified value of Source is not valid.
	ErrCodeInvalidSourceException = "InvalidSourceException"

	// ErrCodeInvalidTagParameterException for service response error code
	// "InvalidTagParameterException".
	//
	// This exception is thrown when the specified tag key or values are not valid.
	// It can also occur if there are duplicate tags or too many tags on the resource.
	ErrCodeInvalidTagParameterException = "InvalidTagParameterException"

	// ErrCodeInvalidTimeRangeException for service response error code
	// "InvalidTimeRangeException".
	//
	// Occurs if the timestamp values are not valid. Either the start time occurs
	// after the end time, or the time range is outside the range of possible values.
	ErrCodeInvalidTimeRangeException = "InvalidTimeRangeException"

	// ErrCodeInvalidTokenException for service response error code
	// "InvalidTokenException".
	//
	// Reserved for future use.
	ErrCodeInvalidTokenException = "InvalidTokenException"

	// ErrCodeInvalidTrailNameException for service response error code
	// "InvalidTrailNameException".
	//
	// This exception is thrown when the provided trail name is not valid. Trail
	// names must meet the following requirements:
	//
	//    * Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores
	//    (_), or dashes (-)
	//
	//    * Start with a letter or number, and end with a letter or number
	//
	//    * Be between 3 and 128 characters
	//
	//    * Have no adjacent periods, underscores or dashes. Names like my-_namespace
	//    and my--namespace are not valid.
	//
	//    * Not be in IP address format (for example, 192.168.5.4)
	ErrCodeInvalidTrailNameException = "InvalidTrailNameException"

	// ErrCodeKmsException for service response error code
	// "KmsException".
	//
	// This exception is thrown when there is an issue with the specified KMS key
	// and the trail or event data store can't be updated.
	ErrCodeKmsException = "KmsException"

	// ErrCodeKmsKeyDisabledException for service response error code
	// "KmsKeyDisabledException".
	//
	// This exception is no longer in use.
	ErrCodeKmsKeyDisabledException = "KmsKeyDisabledException"

	// ErrCodeKmsKeyNotFoundException for service response error code
	// "KmsKeyNotFoundException".
	//
	// This exception is thrown when the KMS key does not exist, when the S3 bucket
	// and the KMS key are not in the same Region, or when the KMS key associated
	// with the Amazon SNS topic either does not exist or is not in the same Region.
	ErrCodeKmsKeyNotFoundException = "KmsKeyNotFoundException"

	// ErrCodeMaxConcurrentQueriesException for service response error code
	// "MaxConcurrentQueriesException".
	//
	// You are already running the maximum number of concurrent queries. The maximum
	// number of concurrent queries is 10. Wait a minute for some queries to finish,
	// and then run the query again.
	ErrCodeMaxConcurrentQueriesException = "MaxConcurrentQueriesException"

	// ErrCodeMaximumNumberOfTrailsExceededException for service response error code
	// "MaximumNumberOfTrailsExceededException".
	//
	// This exception is thrown when the maximum number of trails is reached.
	ErrCodeMaximumNumberOfTrailsExceededException = "MaximumNumberOfTrailsExceededException"

	// ErrCodeNoManagementAccountSLRExistsException for service response error code
	// "NoManagementAccountSLRExistsException".
	//
	// This exception is thrown when the management account does not have a service-linked
	// role.
	ErrCodeNoManagementAccountSLRExistsException = "NoManagementAccountSLRExistsException"

	// ErrCodeNotOrganizationManagementAccountException for service response error code
	// "NotOrganizationManagementAccountException".
	//
	// This exception is thrown when the account making the request is not the organization's
	// management account.
	ErrCodeNotOrganizationManagementAccountException = "NotOrganizationManagementAccountException"

	// ErrCodeNotOrganizationMasterAccountException for service response error code
	// "NotOrganizationMasterAccountException".
	//
	// This exception is thrown when the Amazon Web Services account making the
	// request to create or update an organization trail or event data store is
	// not the management account for an organization in Organizations. For more
	// information, see Prepare For Creating a Trail For Your Organization (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/creating-an-organizational-trail-prepare.html)
	// or Create an event data store (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store.html).
	ErrCodeNotOrganizationMasterAccountException = "NotOrganizationMasterAccountException"

	// ErrCodeOperationNotPermittedException for service response error code
	// "OperationNotPermittedException".
	//
	// This exception is thrown when the requested operation is not permitted.
	ErrCodeOperationNotPermittedException = "OperationNotPermittedException"

	// ErrCodeOrganizationNotInAllFeaturesModeException for service response error code
	// "OrganizationNotInAllFeaturesModeException".
	//
	// This exception is thrown when Organizations is not configured to support
	// all features. All features must be enabled in Organizations to support creating
	// an organization trail or event data store.
	ErrCodeOrganizationNotInAllFeaturesModeException = "OrganizationNotInAllFeaturesModeException"

	// ErrCodeOrganizationsNotInUseException for service response error code
	// "OrganizationsNotInUseException".
	//
	// This exception is thrown when the request is made from an Amazon Web Services
	// account that is not a member of an organization. To make this request, sign
	// in using the credentials of an account that belongs to an organization.
	ErrCodeOrganizationsNotInUseException = "OrganizationsNotInUseException"

	// ErrCodeQueryIdNotFoundException for service response error code
	// "QueryIdNotFoundException".
	//
	// The query ID does not exist or does not map to a query.
	ErrCodeQueryIdNotFoundException = "QueryIdNotFoundException"

	// ErrCodeResourceARNNotValidException for service response error code
	// "ResourceARNNotValidException".
	//
	// This exception is thrown when the provided resource does not exist, or the
	// ARN format of the resource is not valid. The following is the valid format
	// for a resource ARN: arn:aws:cloudtrail:us-east-2:123456789012:channel/MyChannel.
	ErrCodeResourceARNNotValidException = "ResourceARNNotValidException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// This exception is thrown when the specified resource is not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourcePolicyNotFoundException for service response error code
	// "ResourcePolicyNotFoundException".
	//
	// This exception is thrown when the specified resource policy is not found.
	ErrCodeResourcePolicyNotFoundException = "ResourcePolicyNotFoundException"

	// ErrCodeResourcePolicyNotValidException for service response error code
	// "ResourcePolicyNotValidException".
	//
	// This exception is thrown when the resouce-based policy has syntax errors,
	// or contains a principal that is not valid.
	//
	// The following are requirements for the resource policy:
	//
	//    * Contains only one action: cloudtrail-data:PutAuditEvents
	//
	//    * Contains at least one statement. The policy can have a maximum of 20
	//    statements.
	//
	//    * Each statement contains at least one principal. A statement can have
	//    a maximum of 50 principals.
	ErrCodeResourcePolicyNotValidException = "ResourcePolicyNotValidException"

	// ErrCodeResourceTypeNotSupportedException for service response error code
	// "ResourceTypeNotSupportedException".
	//
	// This exception is thrown when the specified resource type is not supported
	// by CloudTrail.
	ErrCodeResourceTypeNotSupportedException = "ResourceTypeNotSupportedException"

	// ErrCodeS3BucketDoesNotExistException for service response error code
	// "S3BucketDoesNotExistException".
	//
	// This exception is thrown when the specified S3 bucket does not exist.
	ErrCodeS3BucketDoesNotExistException = "S3BucketDoesNotExistException"

	// ErrCodeTagsLimitExceededException for service response error code
	// "TagsLimitExceededException".
	//
	// The number of tags per trail, event data store, or channel has exceeded the
	// permitted amount. Currently, the limit is 50.
	ErrCodeTagsLimitExceededException = "TagsLimitExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// This exception is thrown when the request rate exceeds the limit.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeTrailAlreadyExistsException for service response error code
	// "TrailAlreadyExistsException".
	//
	// This exception is thrown when the specified trail already exists.
	ErrCodeTrailAlreadyExistsException = "TrailAlreadyExistsException"

	// ErrCodeTrailNotFoundException for service response error code
	// "TrailNotFoundException".
	//
	// This exception is thrown when the trail with the given name is not found.
	ErrCodeTrailNotFoundException = "TrailNotFoundException"

	// ErrCodeTrailNotProvidedException for service response error code
	// "TrailNotProvidedException".
	//
	// This exception is no longer in use.
	ErrCodeTrailNotProvidedException = "TrailNotProvidedException"

	// ErrCodeUnsupportedOperationException for service response error code
	// "UnsupportedOperationException".
	//
	// This exception is thrown when the requested operation is not supported.
	ErrCodeUnsupportedOperationException = "UnsupportedOperationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"CloudTrailARNInvalidException":                          newErrorARNInvalidException,
	"AccessDeniedException":                                  newErrorAccessDeniedException,
	"CloudTrailAccessNotEnabledException":                    newErrorAccessNotEnabledException,
	"AccountHasOngoingImportException":                       newErrorAccountHasOngoingImportException,
	"AccountNotFoundException":                               newErrorAccountNotFoundException,
	"AccountNotRegisteredException":                          newErrorAccountNotRegisteredException,
	"AccountRegisteredException":                             newErrorAccountRegisteredException,
	"CannotDelegateManagementAccountException":               newErrorCannotDelegateManagementAccountException,
	"ChannelARNInvalidException":                             newErrorChannelARNInvalidException,
	"ChannelAlreadyExistsException":                          newErrorChannelAlreadyExistsException,
	"ChannelExistsForEDSException":                           newErrorChannelExistsForEDSException,
	"ChannelMaxLimitExceededException":                       newErrorChannelMaxLimitExceededException,
	"ChannelNotFoundException":                               newErrorChannelNotFoundException,
	"CloudTrailInvalidClientTokenIdException":                newErrorCloudTrailInvalidClientTokenIdException,
	"CloudWatchLogsDeliveryUnavailableException":             newErrorCloudWatchLogsDeliveryUnavailableException,
	"ConcurrentModificationException":                        newErrorConcurrentModificationException,
	"ConflictException":                                      newErrorConflictException,
	"DelegatedAdminAccountLimitExceededException":            newErrorDelegatedAdminAccountLimitExceededException,
	"EventDataStoreARNInvalidException":                      newErrorEventDataStoreARNInvalidException,
	"EventDataStoreAlreadyExistsException":                   newErrorEventDataStoreAlreadyExistsException,
	"EventDataStoreFederationEnabledException":               newErrorEventDataStoreFederationEnabledException,
	"EventDataStoreHasOngoingImportException":                newErrorEventDataStoreHasOngoingImportException,
	"EventDataStoreMaxLimitExceededException":                newErrorEventDataStoreMaxLimitExceededException,
	"EventDataStoreNotFoundException":                        newErrorEventDataStoreNotFoundException,
	"EventDataStoreTerminationProtectedException":            newErrorEventDataStoreTerminationProtectedException,
	"ImportNotFoundException":                                newErrorImportNotFoundException,
	"InactiveEventDataStoreException":                        newErrorInactiveEventDataStoreException,
	"InactiveQueryException":                                 newErrorInactiveQueryException,
	"InsightNotEnabledException":                             newErrorInsightNotEnabledException,
	"InsufficientDependencyServiceAccessPermissionException": newErrorInsufficientDependencyServiceAccessPermissionException,
	"InsufficientEncryptionPolicyException":                  newErrorInsufficientEncryptionPolicyException,
	"InsufficientS3BucketPolicyException":                    newErrorInsufficientS3BucketPolicyException,
	"InsufficientSnsTopicPolicyException":                    newErrorInsufficientSnsTopicPolicyException,
	"InvalidCloudWatchLogsLogGroupArnException":              newErrorInvalidCloudWatchLogsLogGroupArnException,
	"InvalidCloudWatchLogsRoleArnException":                  newErrorInvalidCloudWatchLogsRoleArnException,
	"InvalidDateRangeException":                              newErrorInvalidDateRangeException,
	"InvalidEventCategoryException":                          newErrorInvalidEventCategoryException,
	"InvalidEventDataStoreCategoryException":                 newErrorInvalidEventDataStoreCategoryException,
	"InvalidEventDataStoreStatusException":                   newErrorInvalidEventDataStoreStatusException,
	"InvalidEventSelectorsException":                         newErrorInvalidEventSelectorsException,
	"InvalidHomeRegionException":                             newErrorInvalidHomeRegionException,
	"InvalidImportSourceException":                           newErrorInvalidImportSourceException,
	"InvalidInsightSelectorsException":                       newErrorInvalidInsightSelectorsException,
	"InvalidKmsKeyIdException":                               newErrorInvalidKmsKeyIdException,
	"InvalidLookupAttributesException":                       newErrorInvalidLookupAttributesException,
	"InvalidMaxResultsException":                             newErrorInvalidMaxResultsException,
	"InvalidNextTokenException":                              newErrorInvalidNextTokenException,
	"InvalidParameterCombinationException":                   newErrorInvalidParameterCombinationException,
	"InvalidParameterException":                              newErrorInvalidParameterException,
	"InvalidQueryStatementException":                         newErrorInvalidQueryStatementException,
	"InvalidQueryStatusException":                            newErrorInvalidQueryStatusException,
	"InvalidS3BucketNameException":                           newErrorInvalidS3BucketNameException,
	"InvalidS3PrefixException":                               newErrorInvalidS3PrefixException,
	"InvalidSnsTopicNameException":                           newErrorInvalidSnsTopicNameException,
	"InvalidSourceException":                                 newErrorInvalidSourceException,
	"InvalidTagParameterException":                           newErrorInvalidTagParameterException,
	"InvalidTimeRangeException":                              newErrorInvalidTimeRangeException,
	"InvalidTokenException":                                  newErrorInvalidTokenException,
	"InvalidTrailNameException":                              newErrorInvalidTrailNameException,
	"KmsException":                                           newErrorKmsException,
	"KmsKeyDisabledException":                                newErrorKmsKeyDisabledException,
	"KmsKeyNotFoundException":                                newErrorKmsKeyNotFoundException,
	"MaxConcurrentQueriesException":                          newErrorMaxConcurrentQueriesException,
	"MaximumNumberOfTrailsExceededException":                 newErrorMaximumNumberOfTrailsExceededException,
	"NoManagementAccountSLRExistsException":                  newErrorNoManagementAccountSLRExistsException,
	"NotOrganizationManagementAccountException":              newErrorNotOrganizationManagementAccountException,
	"NotOrganizationMasterAccountException":                  newErrorNotOrganizationMasterAccountException,
	"OperationNotPermittedException":                         newErrorOperationNotPermittedException,
	"OrganizationNotInAllFeaturesModeException":              newErrorOrganizationNotInAllFeaturesModeException,
	"OrganizationsNotInUseException":                         newErrorOrganizationsNotInUseException,
	"QueryIdNotFoundException":                               newErrorQueryIdNotFoundException,
	"ResourceARNNotValidException":                           newErrorResourceARNNotValidException,
	"ResourceNotFoundException":                              newErrorResourceNotFoundException,
	"ResourcePolicyNotFoundException":                        newErrorResourcePolicyNotFoundException,
	"ResourcePolicyNotValidException":                        newErrorResourcePolicyNotValidException,
	"ResourceTypeNotSupportedException":                      newErrorResourceTypeNotSupportedException,
	"S3BucketDoesNotExistException":                          newErrorS3BucketDoesNotExistException,
	"TagsLimitExceededException":                             newErrorTagsLimitExceededException,
	"ThrottlingException":                                    newErrorThrottlingException,
	"TrailAlreadyExistsException":                            newErrorTrailAlreadyExistsException,
	"TrailNotFoundException":                                 newErrorTrailNotFoundException,
	"TrailNotProvidedException":                              newErrorTrailNotProvidedException,
	"UnsupportedOperationException":                          newErrorUnsupportedOperationException,
}
