


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRuleSet(ctx *pulumi.Context, args *LookupRuleSetArgs, opts ...pulumi.InvokeOption) (*LookupRuleSetResult, error) {
	var rv LookupRuleSetResult
	err := ctx.Invoke("azure-native:cdn/v20200901:getRuleSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuleSetArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleSetName       string `pulumi:"ruleSetName"`
}


type LookupRuleSetResult struct {
	DeploymentStatus  string             `pulumi:"deploymentStatus"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}
