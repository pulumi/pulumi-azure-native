


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRule(ctx *pulumi.Context, args *LookupRuleArgs, opts ...pulumi.InvokeOption) (*LookupRuleResult, error) {
	var rv LookupRuleResult
	err := ctx.Invoke("azure-native:servicebus/v20210601preview:getRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuleArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
	SubscriptionName  string `pulumi:"subscriptionName"`
	TopicName         string `pulumi:"topicName"`
}


type LookupRuleResult struct {
	Action            *ActionResponse            `pulumi:"action"`
	CorrelationFilter *CorrelationFilterResponse `pulumi:"correlationFilter"`
	FilterType        *string                    `pulumi:"filterType"`
	Id                string                     `pulumi:"id"`
	Name              string                     `pulumi:"name"`
	SqlFilter         *SqlFilterResponse         `pulumi:"sqlFilter"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	Type              string                     `pulumi:"type"`
}
