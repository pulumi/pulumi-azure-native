


package v20151031

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:automation/v20151031:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	CertificateName       string `pulumi:"certificateName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupCertificateResult struct {
	CreationTime     string  `pulumi:"creationTime"`
	Description      *string `pulumi:"description"`
	ExpiryTime       string  `pulumi:"expiryTime"`
	Id               string  `pulumi:"id"`
	IsExportable     bool    `pulumi:"isExportable"`
	LastModifiedTime string  `pulumi:"lastModifiedTime"`
	Name             string  `pulumi:"name"`
	Thumbprint       string  `pulumi:"thumbprint"`
	Type             string  `pulumi:"type"`
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
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	CertificateName       pulumi.StringInput `pulumi:"certificateName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
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

func (o LookupCertificateResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupCertificateResultOutput) ExpiryTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ExpiryTime }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) IsExportable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCertificateResult) bool { return v.IsExportable }).(pulumi.BoolOutput)
}

func (o LookupCertificateResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
