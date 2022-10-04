


package signalrservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSignalRCustomCertificate(ctx *pulumi.Context, args *LookupSignalRCustomCertificateArgs, opts ...pulumi.InvokeOption) (*LookupSignalRCustomCertificateResult, error) {
	var rv LookupSignalRCustomCertificateResult
	err := ctx.Invoke("azure-native:signalrservice:getSignalRCustomCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRCustomCertificateArgs struct {
	CertificateName   string `pulumi:"certificateName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSignalRCustomCertificateResult struct {
	Id                    string             `pulumi:"id"`
	KeyVaultBaseUri       string             `pulumi:"keyVaultBaseUri"`
	KeyVaultSecretName    string             `pulumi:"keyVaultSecretName"`
	KeyVaultSecretVersion *string            `pulumi:"keyVaultSecretVersion"`
	Name                  string             `pulumi:"name"`
	ProvisioningState     string             `pulumi:"provisioningState"`
	SystemData            SystemDataResponse `pulumi:"systemData"`
	Type                  string             `pulumi:"type"`
}

func LookupSignalRCustomCertificateOutput(ctx *pulumi.Context, args LookupSignalRCustomCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupSignalRCustomCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRCustomCertificateResult, error) {
			args := v.(LookupSignalRCustomCertificateArgs)
			r, err := LookupSignalRCustomCertificate(ctx, &args, opts...)
			var s LookupSignalRCustomCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRCustomCertificateResultOutput)
}

type LookupSignalRCustomCertificateOutputArgs struct {
	CertificateName   pulumi.StringInput `pulumi:"certificateName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSignalRCustomCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRCustomCertificateArgs)(nil)).Elem()
}


type LookupSignalRCustomCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRCustomCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRCustomCertificateResult)(nil)).Elem()
}

func (o LookupSignalRCustomCertificateResultOutput) ToLookupSignalRCustomCertificateResultOutput() LookupSignalRCustomCertificateResultOutput {
	return o
}

func (o LookupSignalRCustomCertificateResultOutput) ToLookupSignalRCustomCertificateResultOutputWithContext(ctx context.Context) LookupSignalRCustomCertificateResultOutput {
	return o
}

func (o LookupSignalRCustomCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSignalRCustomCertificateResultOutput) KeyVaultBaseUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.KeyVaultBaseUri }).(pulumi.StringOutput)
}

func (o LookupSignalRCustomCertificateResultOutput) KeyVaultSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.KeyVaultSecretName }).(pulumi.StringOutput)
}

func (o LookupSignalRCustomCertificateResultOutput) KeyVaultSecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) *string { return v.KeyVaultSecretVersion }).(pulumi.StringPtrOutput)
}

func (o LookupSignalRCustomCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSignalRCustomCertificateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSignalRCustomCertificateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSignalRCustomCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRCustomCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRCustomCertificateResultOutput{})
}
