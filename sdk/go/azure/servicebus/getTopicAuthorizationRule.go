


package servicebus

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTopicAuthorizationRule(ctx *pulumi.Context, args *LookupTopicAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupTopicAuthorizationRuleResult, error) {
	var rv LookupTopicAuthorizationRuleResult
	err := ctx.Invoke("azure-native:servicebus:getTopicAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TopicName             string `pulumi:"topicName"`
}


type LookupTopicAuthorizationRuleResult struct {
	Id     string   `pulumi:"id"`
	Name   string   `pulumi:"name"`
	Rights []string `pulumi:"rights"`
	Type   string   `pulumi:"type"`
}
