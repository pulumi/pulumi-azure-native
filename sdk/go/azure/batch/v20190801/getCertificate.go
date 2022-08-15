


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:batch/v20190801:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	AccountName       string `pulumi:"accountName"`
	CertificateName   string `pulumi:"certificateName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCertificateResult struct {
	DeleteCertificateError                  DeleteCertificateErrorResponse `pulumi:"deleteCertificateError"`
	Etag                                    string                         `pulumi:"etag"`
	Format                                  *string                        `pulumi:"format"`
	Id                                      string                         `pulumi:"id"`
	Name                                    string                         `pulumi:"name"`
	PreviousProvisioningState               string                         `pulumi:"previousProvisioningState"`
	PreviousProvisioningStateTransitionTime string                         `pulumi:"previousProvisioningStateTransitionTime"`
	ProvisioningState                       string                         `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime         string                         `pulumi:"provisioningStateTransitionTime"`
	PublicData                              string                         `pulumi:"publicData"`
	Thumbprint                              *string                        `pulumi:"thumbprint"`
	ThumbprintAlgorithm                     *string                        `pulumi:"thumbprintAlgorithm"`
	Type                                    string                         `pulumi:"type"`
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
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	CertificateName   pulumi.StringInput `pulumi:"certificateName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
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

func (o LookupCertificateResultOutput) DeleteCertificateError() DeleteCertificateErrorResponseOutput {
	return o.ApplyT(func(v LookupCertificateResult) DeleteCertificateErrorResponse { return v.DeleteCertificateError }).(DeleteCertificateErrorResponseOutput)
}

func (o LookupCertificateResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) PreviousProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.PreviousProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) PreviousProvisioningStateTransitionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.PreviousProvisioningStateTransitionTime }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) ProvisioningStateTransitionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ProvisioningStateTransitionTime }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) PublicData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.PublicData }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) ThumbprintAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.ThumbprintAlgorithm }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
