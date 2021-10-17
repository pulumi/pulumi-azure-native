


package v20170515

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSupportPlanType(ctx *pulumi.Context, args *LookupSupportPlanTypeArgs, opts ...pulumi.InvokeOption) (*LookupSupportPlanTypeResult, error) {
	var rv LookupSupportPlanTypeResult
	err := ctx.Invoke("azure-native:addons/v20170515:getSupportPlanType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSupportPlanTypeArgs struct {
	PlanTypeName string `pulumi:"planTypeName"`
	ProviderName string `pulumi:"providerName"`
}


type LookupSupportPlanTypeResult struct {
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
	Type              string  `pulumi:"type"`
}
