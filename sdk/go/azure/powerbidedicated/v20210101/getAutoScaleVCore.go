


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAutoScaleVCore(ctx *pulumi.Context, args *LookupAutoScaleVCoreArgs, opts ...pulumi.InvokeOption) (*LookupAutoScaleVCoreResult, error) {
	var rv LookupAutoScaleVCoreResult
	err := ctx.Invoke("azure-native:powerbidedicated/v20210101:getAutoScaleVCore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutoScaleVCoreArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VcoreName         string `pulumi:"vcoreName"`
}


type LookupAutoScaleVCoreResult struct {
	CapacityLimit     *int                      `pulumi:"capacityLimit"`
	CapacityObjectId  *string                   `pulumi:"capacityObjectId"`
	Id                string                    `pulumi:"id"`
	Location          string                    `pulumi:"location"`
	Name              string                    `pulumi:"name"`
	ProvisioningState string                    `pulumi:"provisioningState"`
	Sku               AutoScaleVCoreSkuResponse `pulumi:"sku"`
	SystemData        *SystemDataResponse       `pulumi:"systemData"`
	Tags              map[string]string         `pulumi:"tags"`
	Type              string                    `pulumi:"type"`
}
