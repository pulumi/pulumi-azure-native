


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServiceCertificateOrderCertificate(ctx *pulumi.Context, args *LookupAppServiceCertificateOrderCertificateArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceCertificateOrderCertificateResult, error) {
	var rv LookupAppServiceCertificateOrderCertificateResult
	err := ctx.Invoke("azure-native:certificateregistration/v20210101:getAppServiceCertificateOrderCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceCertificateOrderCertificateArgs struct {
	CertificateOrderName string `pulumi:"certificateOrderName"`
	Name                 string `pulumi:"name"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupAppServiceCertificateOrderCertificateResult struct {
	Id                 string            `pulumi:"id"`
	KeyVaultId         *string           `pulumi:"keyVaultId"`
	KeyVaultSecretName *string           `pulumi:"keyVaultSecretName"`
	Kind               *string           `pulumi:"kind"`
	Location           string            `pulumi:"location"`
	Name               string            `pulumi:"name"`
	ProvisioningState  string            `pulumi:"provisioningState"`
	Tags               map[string]string `pulumi:"tags"`
	Type               string            `pulumi:"type"`
}

func LookupAppServiceCertificateOrderCertificateOutput(ctx *pulumi.Context, args LookupAppServiceCertificateOrderCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupAppServiceCertificateOrderCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppServiceCertificateOrderCertificateResult, error) {
			args := v.(LookupAppServiceCertificateOrderCertificateArgs)
			r, err := LookupAppServiceCertificateOrderCertificate(ctx, &args, opts...)
			var s LookupAppServiceCertificateOrderCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppServiceCertificateOrderCertificateResultOutput)
}

type LookupAppServiceCertificateOrderCertificateOutputArgs struct {
	CertificateOrderName pulumi.StringInput `pulumi:"certificateOrderName"`
	Name                 pulumi.StringInput `pulumi:"name"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAppServiceCertificateOrderCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceCertificateOrderCertificateArgs)(nil)).Elem()
}


type LookupAppServiceCertificateOrderCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupAppServiceCertificateOrderCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServiceCertificateOrderCertificateResult)(nil)).Elem()
}

func (o LookupAppServiceCertificateOrderCertificateResultOutput) ToLookupAppServiceCertificateOrderCertificateResultOutput() LookupAppServiceCertificateOrderCertificateResultOutput {
	return o
}

func (o LookupAppServiceCertificateOrderCertificateResultOutput) ToLookupAppServiceCertificateOrderCertificateResultOutputWithContext(ctx context.Context) LookupAppServiceCertificateOrderCertificateResultOutput {
	return o
}

func (o LookupAppServiceCertificateOrderCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderCertificateResultOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderCertificateResult) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceCertificateOrderCertificateResultOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderCertificateResult) *string { return v.KeyVaultSecretName }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceCertificateOrderCertificateResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderCertificateResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupAppServiceCertificateOrderCertificateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderCertificateResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderCertificateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderCertificateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAppServiceCertificateOrderCertificateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderCertificateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAppServiceCertificateOrderCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServiceCertificateOrderCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppServiceCertificateOrderCertificateResultOutput{})
}
