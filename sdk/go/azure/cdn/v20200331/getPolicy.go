


package v20200331

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	var rv LookupPolicyResult
	err := ctx.Invoke("azure-native:cdn/v20200331:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyArgs struct {
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPolicyResult struct {
	CustomRules       *CustomRuleListResponse     `pulumi:"customRules"`
	EndpointLinks     []CdnEndpointResponse       `pulumi:"endpointLinks"`
	Etag              *string                     `pulumi:"etag"`
	Id                string                      `pulumi:"id"`
	Location          string                      `pulumi:"location"`
	ManagedRules      *ManagedRuleSetListResponse `pulumi:"managedRules"`
	Name              string                      `pulumi:"name"`
	PolicySettings    *PolicySettingsResponse     `pulumi:"policySettings"`
	ProvisioningState string                      `pulumi:"provisioningState"`
	RateLimitRules    *RateLimitRuleListResponse  `pulumi:"rateLimitRules"`
	ResourceState     string                      `pulumi:"resourceState"`
	Sku               SkuResponse                 `pulumi:"sku"`
	Tags              map[string]string           `pulumi:"tags"`
	Type              string                      `pulumi:"type"`
}
