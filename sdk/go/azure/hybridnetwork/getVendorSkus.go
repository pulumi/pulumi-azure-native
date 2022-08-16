


package hybridnetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVendorSkus(ctx *pulumi.Context, args *LookupVendorSkusArgs, opts ...pulumi.InvokeOption) (*LookupVendorSkusResult, error) {
	var rv LookupVendorSkusResult
	err := ctx.Invoke("azure-native:hybridnetwork:getVendorSkus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVendorSkusArgs struct {
	SkuName    string `pulumi:"skuName"`
	VendorName string `pulumi:"vendorName"`
}


type LookupVendorSkusResult struct {
	DeploymentMode               *string                          `pulumi:"deploymentMode"`
	Id                           string                           `pulumi:"id"`
	ManagedApplicationParameters interface{}                      `pulumi:"managedApplicationParameters"`
	ManagedApplicationTemplate   interface{}                      `pulumi:"managedApplicationTemplate"`
	Name                         string                           `pulumi:"name"`
	NetworkFunctionTemplate      *NetworkFunctionTemplateResponse `pulumi:"networkFunctionTemplate"`
	Preview                      *bool                            `pulumi:"preview"`
	ProvisioningState            string                           `pulumi:"provisioningState"`
	SkuType                      *string                          `pulumi:"skuType"`
	Type                         string                           `pulumi:"type"`
}

func LookupVendorSkusOutput(ctx *pulumi.Context, args LookupVendorSkusOutputArgs, opts ...pulumi.InvokeOption) LookupVendorSkusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVendorSkusResult, error) {
			args := v.(LookupVendorSkusArgs)
			r, err := LookupVendorSkus(ctx, &args, opts...)
			var s LookupVendorSkusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVendorSkusResultOutput)
}

type LookupVendorSkusOutputArgs struct {
	SkuName    pulumi.StringInput `pulumi:"skuName"`
	VendorName pulumi.StringInput `pulumi:"vendorName"`
}

func (LookupVendorSkusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVendorSkusArgs)(nil)).Elem()
}


type LookupVendorSkusResultOutput struct{ *pulumi.OutputState }

func (LookupVendorSkusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVendorSkusResult)(nil)).Elem()
}

func (o LookupVendorSkusResultOutput) ToLookupVendorSkusResultOutput() LookupVendorSkusResultOutput {
	return o
}

func (o LookupVendorSkusResultOutput) ToLookupVendorSkusResultOutputWithContext(ctx context.Context) LookupVendorSkusResultOutput {
	return o
}

func (o LookupVendorSkusResultOutput) DeploymentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) *string { return v.DeploymentMode }).(pulumi.StringPtrOutput)
}

func (o LookupVendorSkusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVendorSkusResultOutput) ManagedApplicationParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) interface{} { return v.ManagedApplicationParameters }).(pulumi.AnyOutput)
}

func (o LookupVendorSkusResultOutput) ManagedApplicationTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) interface{} { return v.ManagedApplicationTemplate }).(pulumi.AnyOutput)
}

func (o LookupVendorSkusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVendorSkusResultOutput) NetworkFunctionTemplate() NetworkFunctionTemplateResponsePtrOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) *NetworkFunctionTemplateResponse { return v.NetworkFunctionTemplate }).(NetworkFunctionTemplateResponsePtrOutput)
}

func (o LookupVendorSkusResultOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) *bool { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o LookupVendorSkusResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVendorSkusResultOutput) SkuType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) *string { return v.SkuType }).(pulumi.StringPtrOutput)
}

func (o LookupVendorSkusResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVendorSkusResultOutput{})
}
