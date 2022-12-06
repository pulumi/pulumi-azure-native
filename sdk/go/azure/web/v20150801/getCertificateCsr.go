


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupCertificateCsr(ctx *pulumi.Context, args *LookupCertificateCsrArgs, opts ...pulumi.InvokeOption) (*LookupCertificateCsrResult, error) {
	var rv LookupCertificateCsrResult
	err := ctx.Invoke("azure-native:web/v20150801:getCertificateCsr", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateCsrArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCertificateCsrResult struct {
	CsrString          *string           `pulumi:"csrString"`
	DistinguishedName  *string           `pulumi:"distinguishedName"`
	HostingEnvironment *string           `pulumi:"hostingEnvironment"`
	Id                 *string           `pulumi:"id"`
	Kind               *string           `pulumi:"kind"`
	Location           string            `pulumi:"location"`
	Name               *string           `pulumi:"name"`
	Password           *string           `pulumi:"password"`
	PfxBlob            *string           `pulumi:"pfxBlob"`
	PublicKeyHash      *string           `pulumi:"publicKeyHash"`
	Tags               map[string]string `pulumi:"tags"`
	Type               *string           `pulumi:"type"`
}

func LookupCertificateCsrOutput(ctx *pulumi.Context, args LookupCertificateCsrOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateCsrResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateCsrResult, error) {
			args := v.(LookupCertificateCsrArgs)
			r, err := LookupCertificateCsr(ctx, &args, opts...)
			var s LookupCertificateCsrResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCertificateCsrResultOutput)
}

type LookupCertificateCsrOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCertificateCsrOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateCsrArgs)(nil)).Elem()
}


type LookupCertificateCsrResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateCsrResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateCsrResult)(nil)).Elem()
}

func (o LookupCertificateCsrResultOutput) ToLookupCertificateCsrResultOutput() LookupCertificateCsrResultOutput {
	return o
}

func (o LookupCertificateCsrResultOutput) ToLookupCertificateCsrResultOutputWithContext(ctx context.Context) LookupCertificateCsrResultOutput {
	return o
}

func (o LookupCertificateCsrResultOutput) CsrString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) *string { return v.CsrString }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateCsrResultOutput) DistinguishedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) *string { return v.DistinguishedName }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateCsrResultOutput) HostingEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) *string { return v.HostingEnvironment }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateCsrResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateCsrResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateCsrResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCertificateCsrResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateCsrResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateCsrResultOutput) PfxBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) *string { return v.PfxBlob }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateCsrResultOutput) PublicKeyHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) *string { return v.PublicKeyHash }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateCsrResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCertificateCsrResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateCsrResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateCsrResultOutput{})
}
