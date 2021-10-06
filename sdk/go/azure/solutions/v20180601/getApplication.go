


package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:solutions/v20180601:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	ApplicationDefinitionId *string           `pulumi:"applicationDefinitionId"`
	Id                      string            `pulumi:"id"`
	Identity                *IdentityResponse `pulumi:"identity"`
	Kind                    string            `pulumi:"kind"`
	Location                *string           `pulumi:"location"`
	ManagedBy               *string           `pulumi:"managedBy"`
	ManagedResourceGroupId  string            `pulumi:"managedResourceGroupId"`
	Name                    string            `pulumi:"name"`
	Outputs                 interface{}       `pulumi:"outputs"`
	Parameters              interface{}       `pulumi:"parameters"`
	Plan                    *PlanResponse     `pulumi:"plan"`
	ProvisioningState       string            `pulumi:"provisioningState"`
	Sku                     *SkuResponse      `pulumi:"sku"`
	Tags                    map[string]string `pulumi:"tags"`
	Type                    string            `pulumi:"type"`
}
