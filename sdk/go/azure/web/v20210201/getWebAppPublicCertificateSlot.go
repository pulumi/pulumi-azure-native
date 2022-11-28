


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPublicCertificateSlot(ctx *pulumi.Context, args *LookupWebAppPublicCertificateSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPublicCertificateSlotResult, error) {
	var rv LookupWebAppPublicCertificateSlotResult
	err := ctx.Invoke("azure-native:web/v20210201:getWebAppPublicCertificateSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPublicCertificateSlotArgs struct {
	Name                  string `pulumi:"name"`
	PublicCertificateName string `pulumi:"publicCertificateName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	Slot                  string `pulumi:"slot"`
}


type LookupWebAppPublicCertificateSlotResult struct {
	Blob                      *string `pulumi:"blob"`
	Id                        string  `pulumi:"id"`
	Kind                      *string `pulumi:"kind"`
	Name                      string  `pulumi:"name"`
	PublicCertificateLocation *string `pulumi:"publicCertificateLocation"`
	Thumbprint                string  `pulumi:"thumbprint"`
	Type                      string  `pulumi:"type"`
}

func LookupWebAppPublicCertificateSlotOutput(ctx *pulumi.Context, args LookupWebAppPublicCertificateSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppPublicCertificateSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppPublicCertificateSlotResult, error) {
			args := v.(LookupWebAppPublicCertificateSlotArgs)
			r, err := LookupWebAppPublicCertificateSlot(ctx, &args, opts...)
			var s LookupWebAppPublicCertificateSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppPublicCertificateSlotResultOutput)
}

type LookupWebAppPublicCertificateSlotOutputArgs struct {
	Name                  pulumi.StringInput `pulumi:"name"`
	PublicCertificateName pulumi.StringInput `pulumi:"publicCertificateName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot                  pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppPublicCertificateSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPublicCertificateSlotArgs)(nil)).Elem()
}


type LookupWebAppPublicCertificateSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppPublicCertificateSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPublicCertificateSlotResult)(nil)).Elem()
}

func (o LookupWebAppPublicCertificateSlotResultOutput) ToLookupWebAppPublicCertificateSlotResultOutput() LookupWebAppPublicCertificateSlotResultOutput {
	return o
}

func (o LookupWebAppPublicCertificateSlotResultOutput) ToLookupWebAppPublicCertificateSlotResultOutputWithContext(ctx context.Context) LookupWebAppPublicCertificateSlotResultOutput {
	return o
}

func (o LookupWebAppPublicCertificateSlotResultOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateSlotResult) *string { return v.Blob }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPublicCertificateSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppPublicCertificateSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPublicCertificateSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppPublicCertificateSlotResultOutput) PublicCertificateLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateSlotResult) *string { return v.PublicCertificateLocation }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPublicCertificateSlotResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateSlotResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o LookupWebAppPublicCertificateSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppPublicCertificateSlotResultOutput{})
}
