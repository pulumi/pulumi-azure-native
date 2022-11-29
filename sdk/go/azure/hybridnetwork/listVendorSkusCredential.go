


package hybridnetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVendorSkusCredential(ctx *pulumi.Context, args *ListVendorSkusCredentialArgs, opts ...pulumi.InvokeOption) (*ListVendorSkusCredentialResult, error) {
	var rv ListVendorSkusCredentialResult
	err := ctx.Invoke("azure-native:hybridnetwork:listVendorSkusCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVendorSkusCredentialArgs struct {
	SkuName    string `pulumi:"skuName"`
	VendorName string `pulumi:"vendorName"`
}


type ListVendorSkusCredentialResult struct {
	AcrServerUrl *string  `pulumi:"acrServerUrl"`
	AcrToken     *string  `pulumi:"acrToken"`
	Expiry       *string  `pulumi:"expiry"`
	Repositories []string `pulumi:"repositories"`
	Username     *string  `pulumi:"username"`
}

func ListVendorSkusCredentialOutput(ctx *pulumi.Context, args ListVendorSkusCredentialOutputArgs, opts ...pulumi.InvokeOption) ListVendorSkusCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVendorSkusCredentialResult, error) {
			args := v.(ListVendorSkusCredentialArgs)
			r, err := ListVendorSkusCredential(ctx, &args, opts...)
			var s ListVendorSkusCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVendorSkusCredentialResultOutput)
}

type ListVendorSkusCredentialOutputArgs struct {
	SkuName    pulumi.StringInput `pulumi:"skuName"`
	VendorName pulumi.StringInput `pulumi:"vendorName"`
}

func (ListVendorSkusCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVendorSkusCredentialArgs)(nil)).Elem()
}


type ListVendorSkusCredentialResultOutput struct{ *pulumi.OutputState }

func (ListVendorSkusCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVendorSkusCredentialResult)(nil)).Elem()
}

func (o ListVendorSkusCredentialResultOutput) ToListVendorSkusCredentialResultOutput() ListVendorSkusCredentialResultOutput {
	return o
}

func (o ListVendorSkusCredentialResultOutput) ToListVendorSkusCredentialResultOutputWithContext(ctx context.Context) ListVendorSkusCredentialResultOutput {
	return o
}

func (o ListVendorSkusCredentialResultOutput) AcrServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVendorSkusCredentialResult) *string { return v.AcrServerUrl }).(pulumi.StringPtrOutput)
}

func (o ListVendorSkusCredentialResultOutput) AcrToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVendorSkusCredentialResult) *string { return v.AcrToken }).(pulumi.StringPtrOutput)
}

func (o ListVendorSkusCredentialResultOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVendorSkusCredentialResult) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o ListVendorSkusCredentialResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListVendorSkusCredentialResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func (o ListVendorSkusCredentialResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVendorSkusCredentialResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVendorSkusCredentialResultOutput{})
}
