


package v20170401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupQueueAuthorizationRule(ctx *pulumi.Context, args *LookupQueueAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupQueueAuthorizationRuleResult, error) {
	var rv LookupQueueAuthorizationRuleResult
	err := ctx.Invoke("azure-native:servicebus/v20170401:getQueueAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	QueueName             string `pulumi:"queueName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupQueueAuthorizationRuleResult struct {
	Id     string   `pulumi:"id"`
	Name   string   `pulumi:"name"`
	Rights []string `pulumi:"rights"`
	Type   string   `pulumi:"type"`
}
