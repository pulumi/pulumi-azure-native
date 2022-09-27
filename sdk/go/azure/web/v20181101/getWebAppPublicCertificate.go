


package v20181101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPublicCertificate(ctx *pulumi.Context, args *LookupWebAppPublicCertificateArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPublicCertificateResult, error) {
	var rv LookupWebAppPublicCertificateResult
	err := ctx.Invoke("azure-native:web/v20181101:getWebAppPublicCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPublicCertificateArgs struct {
	Name                  string `pulumi:"name"`
	PublicCertificateName string `pulumi:"publicCertificateName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupWebAppPublicCertificateResult struct {
	Blob                      *string `pulumi:"blob"`
	Id                        string  `pulumi:"id"`
	Kind                      *string `pulumi:"kind"`
	Name                      string  `pulumi:"name"`
	PublicCertificateLocation *string `pulumi:"publicCertificateLocation"`
	Thumbprint                string  `pulumi:"thumbprint"`
	Type                      string  `pulumi:"type"`
}

func LookupWebAppPublicCertificateOutput(ctx *pulumi.Context, args LookupWebAppPublicCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppPublicCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppPublicCertificateResult, error) {
			args := v.(LookupWebAppPublicCertificateArgs)
			r, err := LookupWebAppPublicCertificate(ctx, &args, opts...)
			var s LookupWebAppPublicCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppPublicCertificateResultOutput)
}

type LookupWebAppPublicCertificateOutputArgs struct {
	Name                  pulumi.StringInput `pulumi:"name"`
	PublicCertificateName pulumi.StringInput `pulumi:"publicCertificateName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppPublicCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPublicCertificateArgs)(nil)).Elem()
}


type LookupWebAppPublicCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppPublicCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPublicCertificateResult)(nil)).Elem()
}

func (o LookupWebAppPublicCertificateResultOutput) ToLookupWebAppPublicCertificateResultOutput() LookupWebAppPublicCertificateResultOutput {
	return o
}

func (o LookupWebAppPublicCertificateResultOutput) ToLookupWebAppPublicCertificateResultOutputWithContext(ctx context.Context) LookupWebAppPublicCertificateResultOutput {
	return o
}

func (o LookupWebAppPublicCertificateResultOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateResult) *string { return v.Blob }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPublicCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppPublicCertificateResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPublicCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppPublicCertificateResultOutput) PublicCertificateLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateResult) *string { return v.PublicCertificateLocation }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPublicCertificateResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o LookupWebAppPublicCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPublicCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppPublicCertificateResultOutput{})
}
