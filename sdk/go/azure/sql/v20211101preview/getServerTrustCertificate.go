


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerTrustCertificate(ctx *pulumi.Context, args *LookupServerTrustCertificateArgs, opts ...pulumi.InvokeOption) (*LookupServerTrustCertificateResult, error) {
	var rv LookupServerTrustCertificateResult
	err := ctx.Invoke("azure-native:sql/v20211101preview:getServerTrustCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerTrustCertificateArgs struct {
	CertificateName     string `pulumi:"certificateName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupServerTrustCertificateResult struct {
	CertificateName string  `pulumi:"certificateName"`
	Id              string  `pulumi:"id"`
	Name            string  `pulumi:"name"`
	PublicBlob      *string `pulumi:"publicBlob"`
	Thumbprint      string  `pulumi:"thumbprint"`
	Type            string  `pulumi:"type"`
}

func LookupServerTrustCertificateOutput(ctx *pulumi.Context, args LookupServerTrustCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupServerTrustCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerTrustCertificateResult, error) {
			args := v.(LookupServerTrustCertificateArgs)
			r, err := LookupServerTrustCertificate(ctx, &args, opts...)
			var s LookupServerTrustCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerTrustCertificateResultOutput)
}

type LookupServerTrustCertificateOutputArgs struct {
	CertificateName     pulumi.StringInput `pulumi:"certificateName"`
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupServerTrustCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerTrustCertificateArgs)(nil)).Elem()
}


type LookupServerTrustCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupServerTrustCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerTrustCertificateResult)(nil)).Elem()
}

func (o LookupServerTrustCertificateResultOutput) ToLookupServerTrustCertificateResultOutput() LookupServerTrustCertificateResultOutput {
	return o
}

func (o LookupServerTrustCertificateResultOutput) ToLookupServerTrustCertificateResultOutputWithContext(ctx context.Context) LookupServerTrustCertificateResultOutput {
	return o
}

func (o LookupServerTrustCertificateResultOutput) CertificateName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTrustCertificateResult) string { return v.CertificateName }).(pulumi.StringOutput)
}

func (o LookupServerTrustCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTrustCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerTrustCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTrustCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerTrustCertificateResultOutput) PublicBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerTrustCertificateResult) *string { return v.PublicBlob }).(pulumi.StringPtrOutput)
}

func (o LookupServerTrustCertificateResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTrustCertificateResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o LookupServerTrustCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTrustCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerTrustCertificateResultOutput{})
}
