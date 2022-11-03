


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebPubSubCustomCertificate(ctx *pulumi.Context, args *LookupWebPubSubCustomCertificateArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubCustomCertificateResult, error) {
	var rv LookupWebPubSubCustomCertificateResult
	err := ctx.Invoke("azure-native:webpubsub/v20220801preview:getWebPubSubCustomCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebPubSubCustomCertificateArgs struct {
	CertificateName   string `pulumi:"certificateName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWebPubSubCustomCertificateResult struct {
	Id                    string             `pulumi:"id"`
	KeyVaultBaseUri       string             `pulumi:"keyVaultBaseUri"`
	KeyVaultSecretName    string             `pulumi:"keyVaultSecretName"`
	KeyVaultSecretVersion *string            `pulumi:"keyVaultSecretVersion"`
	Name                  string             `pulumi:"name"`
	ProvisioningState     string             `pulumi:"provisioningState"`
	SystemData            SystemDataResponse `pulumi:"systemData"`
	Type                  string             `pulumi:"type"`
}

func LookupWebPubSubCustomCertificateOutput(ctx *pulumi.Context, args LookupWebPubSubCustomCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubCustomCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubCustomCertificateResult, error) {
			args := v.(LookupWebPubSubCustomCertificateArgs)
			r, err := LookupWebPubSubCustomCertificate(ctx, &args, opts...)
			var s LookupWebPubSubCustomCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebPubSubCustomCertificateResultOutput)
}

type LookupWebPubSubCustomCertificateOutputArgs struct {
	CertificateName   pulumi.StringInput `pulumi:"certificateName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWebPubSubCustomCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubCustomCertificateArgs)(nil)).Elem()
}


type LookupWebPubSubCustomCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubCustomCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubCustomCertificateResult)(nil)).Elem()
}

func (o LookupWebPubSubCustomCertificateResultOutput) ToLookupWebPubSubCustomCertificateResultOutput() LookupWebPubSubCustomCertificateResultOutput {
	return o
}

func (o LookupWebPubSubCustomCertificateResultOutput) ToLookupWebPubSubCustomCertificateResultOutputWithContext(ctx context.Context) LookupWebPubSubCustomCertificateResultOutput {
	return o
}

func (o LookupWebPubSubCustomCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebPubSubCustomCertificateResultOutput) KeyVaultBaseUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.KeyVaultBaseUri }).(pulumi.StringOutput)
}

func (o LookupWebPubSubCustomCertificateResultOutput) KeyVaultSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.KeyVaultSecretName }).(pulumi.StringOutput)
}

func (o LookupWebPubSubCustomCertificateResultOutput) KeyVaultSecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) *string { return v.KeyVaultSecretVersion }).(pulumi.StringPtrOutput)
}

func (o LookupWebPubSubCustomCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebPubSubCustomCertificateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWebPubSubCustomCertificateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWebPubSubCustomCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubCustomCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubCustomCertificateResultOutput{})
}
