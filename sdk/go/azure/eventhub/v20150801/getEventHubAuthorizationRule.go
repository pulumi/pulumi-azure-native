


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventHubAuthorizationRule(ctx *pulumi.Context, args *LookupEventHubAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupEventHubAuthorizationRuleResult, error) {
	var rv LookupEventHubAuthorizationRuleResult
	err := ctx.Invoke("azure-native:eventhub/v20150801:getEventHubAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	EventHubName          string `pulumi:"eventHubName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupEventHubAuthorizationRuleResult struct {
	Id       string   `pulumi:"id"`
	Location *string  `pulumi:"location"`
	Name     string   `pulumi:"name"`
	Rights   []string `pulumi:"rights"`
	Type     string   `pulumi:"type"`
}
