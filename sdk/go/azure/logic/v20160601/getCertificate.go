


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:logic/v20160601:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	CertificateName        string `pulumi:"certificateName"`
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupCertificateResult struct {
	ChangedTime       string                        `pulumi:"changedTime"`
	CreatedTime       string                        `pulumi:"createdTime"`
	Id                string                        `pulumi:"id"`
	Key               *KeyVaultKeyReferenceResponse `pulumi:"key"`
	Location          *string                       `pulumi:"location"`
	Metadata          interface{}                   `pulumi:"metadata"`
	Name              string                        `pulumi:"name"`
	PublicCertificate *string                       `pulumi:"publicCertificate"`
	Tags              map[string]string             `pulumi:"tags"`
	Type              string                        `pulumi:"type"`
}

func LookupCertificateOutput(ctx *pulumi.Context, args LookupCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateResult, error) {
			args := v.(LookupCertificateArgs)
			r, err := LookupCertificate(ctx, &args, opts...)
			var s LookupCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCertificateResultOutput)
}

type LookupCertificateOutputArgs struct {
	CertificateName        pulumi.StringInput `pulumi:"certificateName"`
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}


type LookupCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateResult)(nil)).Elem()
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutput() LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutputWithContext(ctx context.Context) LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Key() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *KeyVaultKeyReferenceResponse { return v.Key }).(KeyVaultKeyReferenceResponsePtrOutput)
}

func (o LookupCertificateResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupCertificateResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) PublicCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.PublicCertificate }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCertificateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
