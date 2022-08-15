


package devices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDpsCertificate(ctx *pulumi.Context, args *LookupDpsCertificateArgs, opts ...pulumi.InvokeOption) (*LookupDpsCertificateResult, error) {
	var rv LookupDpsCertificateResult
	err := ctx.Invoke("azure-native:devices:getDpsCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDpsCertificateArgs struct {
	CertificateName         string `pulumi:"certificateName"`
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupDpsCertificateResult struct {
	Etag       string                        `pulumi:"etag"`
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties CertificatePropertiesResponse `pulumi:"properties"`
	Type       string                        `pulumi:"type"`
}

func LookupDpsCertificateOutput(ctx *pulumi.Context, args LookupDpsCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupDpsCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDpsCertificateResult, error) {
			args := v.(LookupDpsCertificateArgs)
			r, err := LookupDpsCertificate(ctx, &args, opts...)
			var s LookupDpsCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDpsCertificateResultOutput)
}

type LookupDpsCertificateOutputArgs struct {
	CertificateName         pulumi.StringInput `pulumi:"certificateName"`
	ProvisioningServiceName pulumi.StringInput `pulumi:"provisioningServiceName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDpsCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDpsCertificateArgs)(nil)).Elem()
}


type LookupDpsCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupDpsCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDpsCertificateResult)(nil)).Elem()
}

func (o LookupDpsCertificateResultOutput) ToLookupDpsCertificateResultOutput() LookupDpsCertificateResultOutput {
	return o
}

func (o LookupDpsCertificateResultOutput) ToLookupDpsCertificateResultOutputWithContext(ctx context.Context) LookupDpsCertificateResultOutput {
	return o
}

func (o LookupDpsCertificateResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDpsCertificateResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDpsCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDpsCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDpsCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDpsCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDpsCertificateResultOutput) Properties() CertificatePropertiesResponseOutput {
	return o.ApplyT(func(v LookupDpsCertificateResult) CertificatePropertiesResponse { return v.Properties }).(CertificatePropertiesResponseOutput)
}

func (o LookupDpsCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDpsCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDpsCertificateResultOutput{})
}
