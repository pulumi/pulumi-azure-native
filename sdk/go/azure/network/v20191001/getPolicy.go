


package v20191001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	var rv LookupPolicyResult
	err := ctx.Invoke("azure-native:network/v20191001:getPolicy", args, &rv, opts...)
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
	CustomRules           *CustomRuleListResponse          `pulumi:"customRules"`
	Etag                  *string                          `pulumi:"etag"`
	FrontendEndpointLinks []FrontendEndpointLinkResponse   `pulumi:"frontendEndpointLinks"`
	Id                    string                           `pulumi:"id"`
	Location              *string                          `pulumi:"location"`
	ManagedRules          *ManagedRuleSetListResponse      `pulumi:"managedRules"`
	Name                  string                           `pulumi:"name"`
	PolicySettings        *FrontDoorPolicySettingsResponse `pulumi:"policySettings"`
	ProvisioningState     string                           `pulumi:"provisioningState"`
	ResourceState         string                           `pulumi:"resourceState"`
	Tags                  map[string]string                `pulumi:"tags"`
	Type                  string                           `pulumi:"type"`
}
