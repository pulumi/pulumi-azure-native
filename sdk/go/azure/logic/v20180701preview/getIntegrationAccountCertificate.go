


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountCertificate(ctx *pulumi.Context, args *LookupIntegrationAccountCertificateArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountCertificateResult, error) {
	var rv LookupIntegrationAccountCertificateResult
	err := ctx.Invoke("azure-native:logic/v20180701preview:getIntegrationAccountCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountCertificateArgs struct {
	CertificateName        string `pulumi:"certificateName"`
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountCertificateResult struct {
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

func LookupIntegrationAccountCertificateOutput(ctx *pulumi.Context, args LookupIntegrationAccountCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountCertificateResult, error) {
			args := v.(LookupIntegrationAccountCertificateArgs)
			r, err := LookupIntegrationAccountCertificate(ctx, &args, opts...)
			var s LookupIntegrationAccountCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountCertificateResultOutput)
}

type LookupIntegrationAccountCertificateOutputArgs struct {
	CertificateName        pulumi.StringInput `pulumi:"certificateName"`
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationAccountCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountCertificateArgs)(nil)).Elem()
}


type LookupIntegrationAccountCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountCertificateResult)(nil)).Elem()
}

func (o LookupIntegrationAccountCertificateResultOutput) ToLookupIntegrationAccountCertificateResultOutput() LookupIntegrationAccountCertificateResultOutput {
	return o
}

func (o LookupIntegrationAccountCertificateResultOutput) ToLookupIntegrationAccountCertificateResultOutputWithContext(ctx context.Context) LookupIntegrationAccountCertificateResultOutput {
	return o
}

func (o LookupIntegrationAccountCertificateResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountCertificateResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountCertificateResultOutput) Key() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) *KeyVaultKeyReferenceResponse { return v.Key }).(KeyVaultKeyReferenceResponsePtrOutput)
}

func (o LookupIntegrationAccountCertificateResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountCertificateResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupIntegrationAccountCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountCertificateResultOutput) PublicCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) *string { return v.PublicCertificate }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountCertificateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIntegrationAccountCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountCertificateResultOutput{})
}
