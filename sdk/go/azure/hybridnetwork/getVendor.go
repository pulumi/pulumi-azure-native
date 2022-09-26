


package hybridnetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVendor(ctx *pulumi.Context, args *LookupVendorArgs, opts ...pulumi.InvokeOption) (*LookupVendorResult, error) {
	var rv LookupVendorResult
	err := ctx.Invoke("azure-native:hybridnetwork:getVendor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVendorArgs struct {
	VendorName string `pulumi:"vendorName"`
}


type LookupVendorResult struct {
	Id                string                `pulumi:"id"`
	Name              string                `pulumi:"name"`
	ProvisioningState string                `pulumi:"provisioningState"`
	Skus              []SubResourceResponse `pulumi:"skus"`
	Type              string                `pulumi:"type"`
}

func LookupVendorOutput(ctx *pulumi.Context, args LookupVendorOutputArgs, opts ...pulumi.InvokeOption) LookupVendorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVendorResult, error) {
			args := v.(LookupVendorArgs)
			r, err := LookupVendor(ctx, &args, opts...)
			var s LookupVendorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVendorResultOutput)
}

type LookupVendorOutputArgs struct {
	VendorName pulumi.StringInput `pulumi:"vendorName"`
}

func (LookupVendorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVendorArgs)(nil)).Elem()
}


type LookupVendorResultOutput struct{ *pulumi.OutputState }

func (LookupVendorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVendorResult)(nil)).Elem()
}

func (o LookupVendorResultOutput) ToLookupVendorResultOutput() LookupVendorResultOutput {
	return o
}

func (o LookupVendorResultOutput) ToLookupVendorResultOutputWithContext(ctx context.Context) LookupVendorResultOutput {
	return o
}

func (o LookupVendorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVendorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVendorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVendorResultOutput) Skus() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVendorResult) []SubResourceResponse { return v.Skus }).(SubResourceResponseArrayOutput)
}

func (o LookupVendorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVendorResultOutput{})
}
