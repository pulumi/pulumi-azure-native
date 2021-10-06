


package authorization

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyPricing(ctx *pulumi.Context, args *LookupPolicyPricingArgs, opts ...pulumi.InvokeOption) (*LookupPolicyPricingResult, error) {
	var rv LookupPolicyPricingResult
	err := ctx.Invoke("azure-native:authorization:getPolicyPricing", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyPricingArgs struct {
	PolicyPricingName string `pulumi:"policyPricingName"`
	Scope             string `pulumi:"scope"`
}


type LookupPolicyPricingResult struct {
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	PricingTier       string             `pulumi:"pricingTier"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}
