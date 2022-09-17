// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package wafiface provides an interface to enable mocking the AWS WAF service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package wafiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/waf"
)

// WAFAPI provides an interface to enable mocking the
// waf.WAF service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS WAF.
//    func myFunc(svc wafiface.WAFAPI) bool {
//        // Make svc.CreateByteMatchSet request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := waf.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockWAFClient struct {
//        wafiface.WAFAPI
//    }
//    func (m *mockWAFClient) CreateByteMatchSet(input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockWAFClient{}
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
type WAFAPI interface {
	CreateByteMatchSet(*waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error)
	CreateByteMatchSetWithContext(aws.Context, *waf.CreateByteMatchSetInput, ...request.Option) (*waf.CreateByteMatchSetOutput, error)
	CreateByteMatchSetRequest(*waf.CreateByteMatchSetInput) (*request.Request, *waf.CreateByteMatchSetOutput)

	CreateGeoMatchSet(*waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error)
	CreateGeoMatchSetWithContext(aws.Context, *waf.CreateGeoMatchSetInput, ...request.Option) (*waf.CreateGeoMatchSetOutput, error)
	CreateGeoMatchSetRequest(*waf.CreateGeoMatchSetInput) (*request.Request, *waf.CreateGeoMatchSetOutput)

	CreateIPSet(*waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error)
	CreateIPSetWithContext(aws.Context, *waf.CreateIPSetInput, ...request.Option) (*waf.CreateIPSetOutput, error)
	CreateIPSetRequest(*waf.CreateIPSetInput) (*request.Request, *waf.CreateIPSetOutput)

	CreateRateBasedRule(*waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error)
	CreateRateBasedRuleWithContext(aws.Context, *waf.CreateRateBasedRuleInput, ...request.Option) (*waf.CreateRateBasedRuleOutput, error)
	CreateRateBasedRuleRequest(*waf.CreateRateBasedRuleInput) (*request.Request, *waf.CreateRateBasedRuleOutput)

	CreateRegexMatchSet(*waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error)
	CreateRegexMatchSetWithContext(aws.Context, *waf.CreateRegexMatchSetInput, ...request.Option) (*waf.CreateRegexMatchSetOutput, error)
	CreateRegexMatchSetRequest(*waf.CreateRegexMatchSetInput) (*request.Request, *waf.CreateRegexMatchSetOutput)

	CreateRegexPatternSet(*waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error)
	CreateRegexPatternSetWithContext(aws.Context, *waf.CreateRegexPatternSetInput, ...request.Option) (*waf.CreateRegexPatternSetOutput, error)
	CreateRegexPatternSetRequest(*waf.CreateRegexPatternSetInput) (*request.Request, *waf.CreateRegexPatternSetOutput)

	CreateRule(*waf.CreateRuleInput) (*waf.CreateRuleOutput, error)
	CreateRuleWithContext(aws.Context, *waf.CreateRuleInput, ...request.Option) (*waf.CreateRuleOutput, error)
	CreateRuleRequest(*waf.CreateRuleInput) (*request.Request, *waf.CreateRuleOutput)

	CreateRuleGroup(*waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error)
	CreateRuleGroupWithContext(aws.Context, *waf.CreateRuleGroupInput, ...request.Option) (*waf.CreateRuleGroupOutput, error)
	CreateRuleGroupRequest(*waf.CreateRuleGroupInput) (*request.Request, *waf.CreateRuleGroupOutput)

	CreateSizeConstraintSet(*waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error)
	CreateSizeConstraintSetWithContext(aws.Context, *waf.CreateSizeConstraintSetInput, ...request.Option) (*waf.CreateSizeConstraintSetOutput, error)
	CreateSizeConstraintSetRequest(*waf.CreateSizeConstraintSetInput) (*request.Request, *waf.CreateSizeConstraintSetOutput)

	CreateSqlInjectionMatchSet(*waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error)
	CreateSqlInjectionMatchSetWithContext(aws.Context, *waf.CreateSqlInjectionMatchSetInput, ...request.Option) (*waf.CreateSqlInjectionMatchSetOutput, error)
	CreateSqlInjectionMatchSetRequest(*waf.CreateSqlInjectionMatchSetInput) (*request.Request, *waf.CreateSqlInjectionMatchSetOutput)

	CreateWebACL(*waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error)
	CreateWebACLWithContext(aws.Context, *waf.CreateWebACLInput, ...request.Option) (*waf.CreateWebACLOutput, error)
	CreateWebACLRequest(*waf.CreateWebACLInput) (*request.Request, *waf.CreateWebACLOutput)

	CreateWebACLMigrationStack(*waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error)
	CreateWebACLMigrationStackWithContext(aws.Context, *waf.CreateWebACLMigrationStackInput, ...request.Option) (*waf.CreateWebACLMigrationStackOutput, error)
	CreateWebACLMigrationStackRequest(*waf.CreateWebACLMigrationStackInput) (*request.Request, *waf.CreateWebACLMigrationStackOutput)

	CreateXssMatchSet(*waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error)
	CreateXssMatchSetWithContext(aws.Context, *waf.CreateXssMatchSetInput, ...request.Option) (*waf.CreateXssMatchSetOutput, error)
	CreateXssMatchSetRequest(*waf.CreateXssMatchSetInput) (*request.Request, *waf.CreateXssMatchSetOutput)

	DeleteByteMatchSet(*waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error)
	DeleteByteMatchSetWithContext(aws.Context, *waf.DeleteByteMatchSetInput, ...request.Option) (*waf.DeleteByteMatchSetOutput, error)
	DeleteByteMatchSetRequest(*waf.DeleteByteMatchSetInput) (*request.Request, *waf.DeleteByteMatchSetOutput)

	DeleteGeoMatchSet(*waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error)
	DeleteGeoMatchSetWithContext(aws.Context, *waf.DeleteGeoMatchSetInput, ...request.Option) (*waf.DeleteGeoMatchSetOutput, error)
	DeleteGeoMatchSetRequest(*waf.DeleteGeoMatchSetInput) (*request.Request, *waf.DeleteGeoMatchSetOutput)

	DeleteIPSet(*waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error)
	DeleteIPSetWithContext(aws.Context, *waf.DeleteIPSetInput, ...request.Option) (*waf.DeleteIPSetOutput, error)
	DeleteIPSetRequest(*waf.DeleteIPSetInput) (*request.Request, *waf.DeleteIPSetOutput)

	DeleteLoggingConfiguration(*waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error)
	DeleteLoggingConfigurationWithContext(aws.Context, *waf.DeleteLoggingConfigurationInput, ...request.Option) (*waf.DeleteLoggingConfigurationOutput, error)
	DeleteLoggingConfigurationRequest(*waf.DeleteLoggingConfigurationInput) (*request.Request, *waf.DeleteLoggingConfigurationOutput)

	DeletePermissionPolicy(*waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error)
	DeletePermissionPolicyWithContext(aws.Context, *waf.DeletePermissionPolicyInput, ...request.Option) (*waf.DeletePermissionPolicyOutput, error)
	DeletePermissionPolicyRequest(*waf.DeletePermissionPolicyInput) (*request.Request, *waf.DeletePermissionPolicyOutput)

	DeleteRateBasedRule(*waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error)
	DeleteRateBasedRuleWithContext(aws.Context, *waf.DeleteRateBasedRuleInput, ...request.Option) (*waf.DeleteRateBasedRuleOutput, error)
	DeleteRateBasedRuleRequest(*waf.DeleteRateBasedRuleInput) (*request.Request, *waf.DeleteRateBasedRuleOutput)

	DeleteRegexMatchSet(*waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error)
	DeleteRegexMatchSetWithContext(aws.Context, *waf.DeleteRegexMatchSetInput, ...request.Option) (*waf.DeleteRegexMatchSetOutput, error)
	DeleteRegexMatchSetRequest(*waf.DeleteRegexMatchSetInput) (*request.Request, *waf.DeleteRegexMatchSetOutput)

	DeleteRegexPatternSet(*waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error)
	DeleteRegexPatternSetWithContext(aws.Context, *waf.DeleteRegexPatternSetInput, ...request.Option) (*waf.DeleteRegexPatternSetOutput, error)
	DeleteRegexPatternSetRequest(*waf.DeleteRegexPatternSetInput) (*request.Request, *waf.DeleteRegexPatternSetOutput)

	DeleteRule(*waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error)
	DeleteRuleWithContext(aws.Context, *waf.DeleteRuleInput, ...request.Option) (*waf.DeleteRuleOutput, error)
	DeleteRuleRequest(*waf.DeleteRuleInput) (*request.Request, *waf.DeleteRuleOutput)

	DeleteRuleGroup(*waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error)
	DeleteRuleGroupWithContext(aws.Context, *waf.DeleteRuleGroupInput, ...request.Option) (*waf.DeleteRuleGroupOutput, error)
	DeleteRuleGroupRequest(*waf.DeleteRuleGroupInput) (*request.Request, *waf.DeleteRuleGroupOutput)

	DeleteSizeConstraintSet(*waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error)
	DeleteSizeConstraintSetWithContext(aws.Context, *waf.DeleteSizeConstraintSetInput, ...request.Option) (*waf.DeleteSizeConstraintSetOutput, error)
	DeleteSizeConstraintSetRequest(*waf.DeleteSizeConstraintSetInput) (*request.Request, *waf.DeleteSizeConstraintSetOutput)

	DeleteSqlInjectionMatchSet(*waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error)
	DeleteSqlInjectionMatchSetWithContext(aws.Context, *waf.DeleteSqlInjectionMatchSetInput, ...request.Option) (*waf.DeleteSqlInjectionMatchSetOutput, error)
	DeleteSqlInjectionMatchSetRequest(*waf.DeleteSqlInjectionMatchSetInput) (*request.Request, *waf.DeleteSqlInjectionMatchSetOutput)

	DeleteWebACL(*waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error)
	DeleteWebACLWithContext(aws.Context, *waf.DeleteWebACLInput, ...request.Option) (*waf.DeleteWebACLOutput, error)
	DeleteWebACLRequest(*waf.DeleteWebACLInput) (*request.Request, *waf.DeleteWebACLOutput)

	DeleteXssMatchSet(*waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error)
	DeleteXssMatchSetWithContext(aws.Context, *waf.DeleteXssMatchSetInput, ...request.Option) (*waf.DeleteXssMatchSetOutput, error)
	DeleteXssMatchSetRequest(*waf.DeleteXssMatchSetInput) (*request.Request, *waf.DeleteXssMatchSetOutput)

	GetByteMatchSet(*waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error)
	GetByteMatchSetWithContext(aws.Context, *waf.GetByteMatchSetInput, ...request.Option) (*waf.GetByteMatchSetOutput, error)
	GetByteMatchSetRequest(*waf.GetByteMatchSetInput) (*request.Request, *waf.GetByteMatchSetOutput)

	GetChangeToken(*waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error)
	GetChangeTokenWithContext(aws.Context, *waf.GetChangeTokenInput, ...request.Option) (*waf.GetChangeTokenOutput, error)
	GetChangeTokenRequest(*waf.GetChangeTokenInput) (*request.Request, *waf.GetChangeTokenOutput)

	GetChangeTokenStatus(*waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error)
	GetChangeTokenStatusWithContext(aws.Context, *waf.GetChangeTokenStatusInput, ...request.Option) (*waf.GetChangeTokenStatusOutput, error)
	GetChangeTokenStatusRequest(*waf.GetChangeTokenStatusInput) (*request.Request, *waf.GetChangeTokenStatusOutput)

	GetGeoMatchSet(*waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error)
	GetGeoMatchSetWithContext(aws.Context, *waf.GetGeoMatchSetInput, ...request.Option) (*waf.GetGeoMatchSetOutput, error)
	GetGeoMatchSetRequest(*waf.GetGeoMatchSetInput) (*request.Request, *waf.GetGeoMatchSetOutput)

	GetIPSet(*waf.GetIPSetInput) (*waf.GetIPSetOutput, error)
	GetIPSetWithContext(aws.Context, *waf.GetIPSetInput, ...request.Option) (*waf.GetIPSetOutput, error)
	GetIPSetRequest(*waf.GetIPSetInput) (*request.Request, *waf.GetIPSetOutput)

	GetLoggingConfiguration(*waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error)
	GetLoggingConfigurationWithContext(aws.Context, *waf.GetLoggingConfigurationInput, ...request.Option) (*waf.GetLoggingConfigurationOutput, error)
	GetLoggingConfigurationRequest(*waf.GetLoggingConfigurationInput) (*request.Request, *waf.GetLoggingConfigurationOutput)

	GetPermissionPolicy(*waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error)
	GetPermissionPolicyWithContext(aws.Context, *waf.GetPermissionPolicyInput, ...request.Option) (*waf.GetPermissionPolicyOutput, error)
	GetPermissionPolicyRequest(*waf.GetPermissionPolicyInput) (*request.Request, *waf.GetPermissionPolicyOutput)

	GetRateBasedRule(*waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error)
	GetRateBasedRuleWithContext(aws.Context, *waf.GetRateBasedRuleInput, ...request.Option) (*waf.GetRateBasedRuleOutput, error)
	GetRateBasedRuleRequest(*waf.GetRateBasedRuleInput) (*request.Request, *waf.GetRateBasedRuleOutput)

	GetRateBasedRuleManagedKeys(*waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error)
	GetRateBasedRuleManagedKeysWithContext(aws.Context, *waf.GetRateBasedRuleManagedKeysInput, ...request.Option) (*waf.GetRateBasedRuleManagedKeysOutput, error)
	GetRateBasedRuleManagedKeysRequest(*waf.GetRateBasedRuleManagedKeysInput) (*request.Request, *waf.GetRateBasedRuleManagedKeysOutput)

	GetRegexMatchSet(*waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error)
	GetRegexMatchSetWithContext(aws.Context, *waf.GetRegexMatchSetInput, ...request.Option) (*waf.GetRegexMatchSetOutput, error)
	GetRegexMatchSetRequest(*waf.GetRegexMatchSetInput) (*request.Request, *waf.GetRegexMatchSetOutput)

	GetRegexPatternSet(*waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error)
	GetRegexPatternSetWithContext(aws.Context, *waf.GetRegexPatternSetInput, ...request.Option) (*waf.GetRegexPatternSetOutput, error)
	GetRegexPatternSetRequest(*waf.GetRegexPatternSetInput) (*request.Request, *waf.GetRegexPatternSetOutput)

	GetRule(*waf.GetRuleInput) (*waf.GetRuleOutput, error)
	GetRuleWithContext(aws.Context, *waf.GetRuleInput, ...request.Option) (*waf.GetRuleOutput, error)
	GetRuleRequest(*waf.GetRuleInput) (*request.Request, *waf.GetRuleOutput)

	GetRuleGroup(*waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error)
	GetRuleGroupWithContext(aws.Context, *waf.GetRuleGroupInput, ...request.Option) (*waf.GetRuleGroupOutput, error)
	GetRuleGroupRequest(*waf.GetRuleGroupInput) (*request.Request, *waf.GetRuleGroupOutput)

	GetSampledRequests(*waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error)
	GetSampledRequestsWithContext(aws.Context, *waf.GetSampledRequestsInput, ...request.Option) (*waf.GetSampledRequestsOutput, error)
	GetSampledRequestsRequest(*waf.GetSampledRequestsInput) (*request.Request, *waf.GetSampledRequestsOutput)

	GetSizeConstraintSet(*waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error)
	GetSizeConstraintSetWithContext(aws.Context, *waf.GetSizeConstraintSetInput, ...request.Option) (*waf.GetSizeConstraintSetOutput, error)
	GetSizeConstraintSetRequest(*waf.GetSizeConstraintSetInput) (*request.Request, *waf.GetSizeConstraintSetOutput)

	GetSqlInjectionMatchSet(*waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error)
	GetSqlInjectionMatchSetWithContext(aws.Context, *waf.GetSqlInjectionMatchSetInput, ...request.Option) (*waf.GetSqlInjectionMatchSetOutput, error)
	GetSqlInjectionMatchSetRequest(*waf.GetSqlInjectionMatchSetInput) (*request.Request, *waf.GetSqlInjectionMatchSetOutput)

	GetWebACL(*waf.GetWebACLInput) (*waf.GetWebACLOutput, error)
	GetWebACLWithContext(aws.Context, *waf.GetWebACLInput, ...request.Option) (*waf.GetWebACLOutput, error)
	GetWebACLRequest(*waf.GetWebACLInput) (*request.Request, *waf.GetWebACLOutput)

	GetXssMatchSet(*waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error)
	GetXssMatchSetWithContext(aws.Context, *waf.GetXssMatchSetInput, ...request.Option) (*waf.GetXssMatchSetOutput, error)
	GetXssMatchSetRequest(*waf.GetXssMatchSetInput) (*request.Request, *waf.GetXssMatchSetOutput)

	ListActivatedRulesInRuleGroup(*waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error)
	ListActivatedRulesInRuleGroupWithContext(aws.Context, *waf.ListActivatedRulesInRuleGroupInput, ...request.Option) (*waf.ListActivatedRulesInRuleGroupOutput, error)
	ListActivatedRulesInRuleGroupRequest(*waf.ListActivatedRulesInRuleGroupInput) (*request.Request, *waf.ListActivatedRulesInRuleGroupOutput)

	ListByteMatchSets(*waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error)
	ListByteMatchSetsWithContext(aws.Context, *waf.ListByteMatchSetsInput, ...request.Option) (*waf.ListByteMatchSetsOutput, error)
	ListByteMatchSetsRequest(*waf.ListByteMatchSetsInput) (*request.Request, *waf.ListByteMatchSetsOutput)

	ListGeoMatchSets(*waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error)
	ListGeoMatchSetsWithContext(aws.Context, *waf.ListGeoMatchSetsInput, ...request.Option) (*waf.ListGeoMatchSetsOutput, error)
	ListGeoMatchSetsRequest(*waf.ListGeoMatchSetsInput) (*request.Request, *waf.ListGeoMatchSetsOutput)

	ListIPSets(*waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error)
	ListIPSetsWithContext(aws.Context, *waf.ListIPSetsInput, ...request.Option) (*waf.ListIPSetsOutput, error)
	ListIPSetsRequest(*waf.ListIPSetsInput) (*request.Request, *waf.ListIPSetsOutput)

	ListLoggingConfigurations(*waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error)
	ListLoggingConfigurationsWithContext(aws.Context, *waf.ListLoggingConfigurationsInput, ...request.Option) (*waf.ListLoggingConfigurationsOutput, error)
	ListLoggingConfigurationsRequest(*waf.ListLoggingConfigurationsInput) (*request.Request, *waf.ListLoggingConfigurationsOutput)

	ListRateBasedRules(*waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error)
	ListRateBasedRulesWithContext(aws.Context, *waf.ListRateBasedRulesInput, ...request.Option) (*waf.ListRateBasedRulesOutput, error)
	ListRateBasedRulesRequest(*waf.ListRateBasedRulesInput) (*request.Request, *waf.ListRateBasedRulesOutput)

	ListRegexMatchSets(*waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error)
	ListRegexMatchSetsWithContext(aws.Context, *waf.ListRegexMatchSetsInput, ...request.Option) (*waf.ListRegexMatchSetsOutput, error)
	ListRegexMatchSetsRequest(*waf.ListRegexMatchSetsInput) (*request.Request, *waf.ListRegexMatchSetsOutput)

	ListRegexPatternSets(*waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error)
	ListRegexPatternSetsWithContext(aws.Context, *waf.ListRegexPatternSetsInput, ...request.Option) (*waf.ListRegexPatternSetsOutput, error)
	ListRegexPatternSetsRequest(*waf.ListRegexPatternSetsInput) (*request.Request, *waf.ListRegexPatternSetsOutput)

	ListRuleGroups(*waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error)
	ListRuleGroupsWithContext(aws.Context, *waf.ListRuleGroupsInput, ...request.Option) (*waf.ListRuleGroupsOutput, error)
	ListRuleGroupsRequest(*waf.ListRuleGroupsInput) (*request.Request, *waf.ListRuleGroupsOutput)

	ListRules(*waf.ListRulesInput) (*waf.ListRulesOutput, error)
	ListRulesWithContext(aws.Context, *waf.ListRulesInput, ...request.Option) (*waf.ListRulesOutput, error)
	ListRulesRequest(*waf.ListRulesInput) (*request.Request, *waf.ListRulesOutput)

	ListSizeConstraintSets(*waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error)
	ListSizeConstraintSetsWithContext(aws.Context, *waf.ListSizeConstraintSetsInput, ...request.Option) (*waf.ListSizeConstraintSetsOutput, error)
	ListSizeConstraintSetsRequest(*waf.ListSizeConstraintSetsInput) (*request.Request, *waf.ListSizeConstraintSetsOutput)

	ListSqlInjectionMatchSets(*waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error)
	ListSqlInjectionMatchSetsWithContext(aws.Context, *waf.ListSqlInjectionMatchSetsInput, ...request.Option) (*waf.ListSqlInjectionMatchSetsOutput, error)
	ListSqlInjectionMatchSetsRequest(*waf.ListSqlInjectionMatchSetsInput) (*request.Request, *waf.ListSqlInjectionMatchSetsOutput)

	ListSubscribedRuleGroups(*waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error)
	ListSubscribedRuleGroupsWithContext(aws.Context, *waf.ListSubscribedRuleGroupsInput, ...request.Option) (*waf.ListSubscribedRuleGroupsOutput, error)
	ListSubscribedRuleGroupsRequest(*waf.ListSubscribedRuleGroupsInput) (*request.Request, *waf.ListSubscribedRuleGroupsOutput)

	ListTagsForResource(*waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *waf.ListTagsForResourceInput, ...request.Option) (*waf.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*waf.ListTagsForResourceInput) (*request.Request, *waf.ListTagsForResourceOutput)

	ListWebACLs(*waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error)
	ListWebACLsWithContext(aws.Context, *waf.ListWebACLsInput, ...request.Option) (*waf.ListWebACLsOutput, error)
	ListWebACLsRequest(*waf.ListWebACLsInput) (*request.Request, *waf.ListWebACLsOutput)

	ListXssMatchSets(*waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error)
	ListXssMatchSetsWithContext(aws.Context, *waf.ListXssMatchSetsInput, ...request.Option) (*waf.ListXssMatchSetsOutput, error)
	ListXssMatchSetsRequest(*waf.ListXssMatchSetsInput) (*request.Request, *waf.ListXssMatchSetsOutput)

	PutLoggingConfiguration(*waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error)
	PutLoggingConfigurationWithContext(aws.Context, *waf.PutLoggingConfigurationInput, ...request.Option) (*waf.PutLoggingConfigurationOutput, error)
	PutLoggingConfigurationRequest(*waf.PutLoggingConfigurationInput) (*request.Request, *waf.PutLoggingConfigurationOutput)

	PutPermissionPolicy(*waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error)
	PutPermissionPolicyWithContext(aws.Context, *waf.PutPermissionPolicyInput, ...request.Option) (*waf.PutPermissionPolicyOutput, error)
	PutPermissionPolicyRequest(*waf.PutPermissionPolicyInput) (*request.Request, *waf.PutPermissionPolicyOutput)

	TagResource(*waf.TagResourceInput) (*waf.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *waf.TagResourceInput, ...request.Option) (*waf.TagResourceOutput, error)
	TagResourceRequest(*waf.TagResourceInput) (*request.Request, *waf.TagResourceOutput)

	UntagResource(*waf.UntagResourceInput) (*waf.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *waf.UntagResourceInput, ...request.Option) (*waf.UntagResourceOutput, error)
	UntagResourceRequest(*waf.UntagResourceInput) (*request.Request, *waf.UntagResourceOutput)

	UpdateByteMatchSet(*waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error)
	UpdateByteMatchSetWithContext(aws.Context, *waf.UpdateByteMatchSetInput, ...request.Option) (*waf.UpdateByteMatchSetOutput, error)
	UpdateByteMatchSetRequest(*waf.UpdateByteMatchSetInput) (*request.Request, *waf.UpdateByteMatchSetOutput)

	UpdateGeoMatchSet(*waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error)
	UpdateGeoMatchSetWithContext(aws.Context, *waf.UpdateGeoMatchSetInput, ...request.Option) (*waf.UpdateGeoMatchSetOutput, error)
	UpdateGeoMatchSetRequest(*waf.UpdateGeoMatchSetInput) (*request.Request, *waf.UpdateGeoMatchSetOutput)

	UpdateIPSet(*waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error)
	UpdateIPSetWithContext(aws.Context, *waf.UpdateIPSetInput, ...request.Option) (*waf.UpdateIPSetOutput, error)
	UpdateIPSetRequest(*waf.UpdateIPSetInput) (*request.Request, *waf.UpdateIPSetOutput)

	UpdateRateBasedRule(*waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error)
	UpdateRateBasedRuleWithContext(aws.Context, *waf.UpdateRateBasedRuleInput, ...request.Option) (*waf.UpdateRateBasedRuleOutput, error)
	UpdateRateBasedRuleRequest(*waf.UpdateRateBasedRuleInput) (*request.Request, *waf.UpdateRateBasedRuleOutput)

	UpdateRegexMatchSet(*waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error)
	UpdateRegexMatchSetWithContext(aws.Context, *waf.UpdateRegexMatchSetInput, ...request.Option) (*waf.UpdateRegexMatchSetOutput, error)
	UpdateRegexMatchSetRequest(*waf.UpdateRegexMatchSetInput) (*request.Request, *waf.UpdateRegexMatchSetOutput)

	UpdateRegexPatternSet(*waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error)
	UpdateRegexPatternSetWithContext(aws.Context, *waf.UpdateRegexPatternSetInput, ...request.Option) (*waf.UpdateRegexPatternSetOutput, error)
	UpdateRegexPatternSetRequest(*waf.UpdateRegexPatternSetInput) (*request.Request, *waf.UpdateRegexPatternSetOutput)

	UpdateRule(*waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error)
	UpdateRuleWithContext(aws.Context, *waf.UpdateRuleInput, ...request.Option) (*waf.UpdateRuleOutput, error)
	UpdateRuleRequest(*waf.UpdateRuleInput) (*request.Request, *waf.UpdateRuleOutput)

	UpdateRuleGroup(*waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error)
	UpdateRuleGroupWithContext(aws.Context, *waf.UpdateRuleGroupInput, ...request.Option) (*waf.UpdateRuleGroupOutput, error)
	UpdateRuleGroupRequest(*waf.UpdateRuleGroupInput) (*request.Request, *waf.UpdateRuleGroupOutput)

	UpdateSizeConstraintSet(*waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error)
	UpdateSizeConstraintSetWithContext(aws.Context, *waf.UpdateSizeConstraintSetInput, ...request.Option) (*waf.UpdateSizeConstraintSetOutput, error)
	UpdateSizeConstraintSetRequest(*waf.UpdateSizeConstraintSetInput) (*request.Request, *waf.UpdateSizeConstraintSetOutput)

	UpdateSqlInjectionMatchSet(*waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error)
	UpdateSqlInjectionMatchSetWithContext(aws.Context, *waf.UpdateSqlInjectionMatchSetInput, ...request.Option) (*waf.UpdateSqlInjectionMatchSetOutput, error)
	UpdateSqlInjectionMatchSetRequest(*waf.UpdateSqlInjectionMatchSetInput) (*request.Request, *waf.UpdateSqlInjectionMatchSetOutput)

	UpdateWebACL(*waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error)
	UpdateWebACLWithContext(aws.Context, *waf.UpdateWebACLInput, ...request.Option) (*waf.UpdateWebACLOutput, error)
	UpdateWebACLRequest(*waf.UpdateWebACLInput) (*request.Request, *waf.UpdateWebACLOutput)

	UpdateXssMatchSet(*waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error)
	UpdateXssMatchSetWithContext(aws.Context, *waf.UpdateXssMatchSetInput, ...request.Option) (*waf.UpdateXssMatchSetOutput, error)
	UpdateXssMatchSetRequest(*waf.UpdateXssMatchSetInput) (*request.Request, *waf.UpdateXssMatchSetOutput)
}

var _ WAFAPI = (*waf.WAF)(nil)
