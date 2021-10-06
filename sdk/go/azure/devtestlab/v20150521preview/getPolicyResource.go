


package v20150521preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyResource(ctx *pulumi.Context, args *LookupPolicyResourceArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResourceResult, error) {
	var rv LookupPolicyResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getPolicyResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyResourceArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	PolicySetName     string `pulumi:"policySetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPolicyResourceResult struct {
	Description       *string           `pulumi:"description"`
	EvaluatorType     *string           `pulumi:"evaluatorType"`
	FactData          *string           `pulumi:"factData"`
	FactName          *string           `pulumi:"factName"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Threshold         *string           `pulumi:"threshold"`
	Type              *string           `pulumi:"type"`
}
