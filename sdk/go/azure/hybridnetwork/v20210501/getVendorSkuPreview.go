


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVendorSkuPreview(ctx *pulumi.Context, args *LookupVendorSkuPreviewArgs, opts ...pulumi.InvokeOption) (*LookupVendorSkuPreviewResult, error) {
	var rv LookupVendorSkuPreviewResult
	err := ctx.Invoke("azure-native:hybridnetwork/v20210501:getVendorSkuPreview", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVendorSkuPreviewArgs struct {
	PreviewSubscription string `pulumi:"previewSubscription"`
	SkuName             string `pulumi:"skuName"`
	VendorName          string `pulumi:"vendorName"`
}


type LookupVendorSkuPreviewResult struct {
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupVendorSkuPreviewOutput(ctx *pulumi.Context, args LookupVendorSkuPreviewOutputArgs, opts ...pulumi.InvokeOption) LookupVendorSkuPreviewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVendorSkuPreviewResult, error) {
			args := v.(LookupVendorSkuPreviewArgs)
			r, err := LookupVendorSkuPreview(ctx, &args, opts...)
			var s LookupVendorSkuPreviewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVendorSkuPreviewResultOutput)
}

type LookupVendorSkuPreviewOutputArgs struct {
	PreviewSubscription pulumi.StringInput `pulumi:"previewSubscription"`
	SkuName             pulumi.StringInput `pulumi:"skuName"`
	VendorName          pulumi.StringInput `pulumi:"vendorName"`
}

func (LookupVendorSkuPreviewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVendorSkuPreviewArgs)(nil)).Elem()
}


type LookupVendorSkuPreviewResultOutput struct{ *pulumi.OutputState }

func (LookupVendorSkuPreviewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVendorSkuPreviewResult)(nil)).Elem()
}

func (o LookupVendorSkuPreviewResultOutput) ToLookupVendorSkuPreviewResultOutput() LookupVendorSkuPreviewResultOutput {
	return o
}

func (o LookupVendorSkuPreviewResultOutput) ToLookupVendorSkuPreviewResultOutputWithContext(ctx context.Context) LookupVendorSkuPreviewResultOutput {
	return o
}

func (o LookupVendorSkuPreviewResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkuPreviewResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVendorSkuPreviewResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkuPreviewResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVendorSkuPreviewResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkuPreviewResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVendorSkuPreviewResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVendorSkuPreviewResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVendorSkuPreviewResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkuPreviewResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVendorSkuPreviewResultOutput{})
}
