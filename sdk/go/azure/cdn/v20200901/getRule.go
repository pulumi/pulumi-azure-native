


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRule(ctx *pulumi.Context, args *LookupRuleArgs, opts ...pulumi.InvokeOption) (*LookupRuleResult, error) {
	var rv LookupRuleResult
	err := ctx.Invoke("azure-native:cdn/v20200901:getRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuleArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
	RuleSetName       string `pulumi:"ruleSetName"`
}


type LookupRuleResult struct {
	Actions                 []interface{}      `pulumi:"actions"`
	Conditions              []interface{}      `pulumi:"conditions"`
	DeploymentStatus        string             `pulumi:"deploymentStatus"`
	Id                      string             `pulumi:"id"`
	MatchProcessingBehavior *string            `pulumi:"matchProcessingBehavior"`
	Name                    string             `pulumi:"name"`
	Order                   int                `pulumi:"order"`
	ProvisioningState       string             `pulumi:"provisioningState"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	Type                    string             `pulumi:"type"`
}
