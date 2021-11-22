


package v20180801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDdosProtectionPlan(ctx *pulumi.Context, args *LookupDdosProtectionPlanArgs, opts ...pulumi.InvokeOption) (*LookupDdosProtectionPlanResult, error) {
	var rv LookupDdosProtectionPlanResult
	err := ctx.Invoke("azure-native:network/v20180801:getDdosProtectionPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDdosProtectionPlanArgs struct {
	DdosProtectionPlanName string `pulumi:"ddosProtectionPlanName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupDdosProtectionPlanResult struct {
	Etag              string                `pulumi:"etag"`
	Id                string                `pulumi:"id"`
	Location          *string               `pulumi:"location"`
	Name              string                `pulumi:"name"`
	ProvisioningState string                `pulumi:"provisioningState"`
	ResourceGuid      string                `pulumi:"resourceGuid"`
	Tags              map[string]string     `pulumi:"tags"`
	Type              string                `pulumi:"type"`
	VirtualNetworks   []SubResourceResponse `pulumi:"virtualNetworks"`
}
