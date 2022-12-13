


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectedEnvironmentsCertificate(ctx *pulumi.Context, args *LookupConnectedEnvironmentsCertificateArgs, opts ...pulumi.InvokeOption) (*LookupConnectedEnvironmentsCertificateResult, error) {
	var rv LookupConnectedEnvironmentsCertificateResult
	err := ctx.Invoke("azure-native:app/v20221001:getConnectedEnvironmentsCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectedEnvironmentsCertificateArgs struct {
	CertificateName          string `pulumi:"certificateName"`
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type LookupConnectedEnvironmentsCertificateResult struct {
	Id         string                        `pulumi:"id"`
	Location   string                        `pulumi:"location"`
	Name       string                        `pulumi:"name"`
	Properties CertificateResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Tags       map[string]string             `pulumi:"tags"`
	Type       string                        `pulumi:"type"`
}

func LookupConnectedEnvironmentsCertificateOutput(ctx *pulumi.Context, args LookupConnectedEnvironmentsCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupConnectedEnvironmentsCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectedEnvironmentsCertificateResult, error) {
			args := v.(LookupConnectedEnvironmentsCertificateArgs)
			r, err := LookupConnectedEnvironmentsCertificate(ctx, &args, opts...)
			var s LookupConnectedEnvironmentsCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectedEnvironmentsCertificateResultOutput)
}

type LookupConnectedEnvironmentsCertificateOutputArgs struct {
	CertificateName          pulumi.StringInput `pulumi:"certificateName"`
	ConnectedEnvironmentName pulumi.StringInput `pulumi:"connectedEnvironmentName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectedEnvironmentsCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsCertificateArgs)(nil)).Elem()
}


type LookupConnectedEnvironmentsCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupConnectedEnvironmentsCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsCertificateResult)(nil)).Elem()
}

func (o LookupConnectedEnvironmentsCertificateResultOutput) ToLookupConnectedEnvironmentsCertificateResultOutput() LookupConnectedEnvironmentsCertificateResultOutput {
	return o
}

func (o LookupConnectedEnvironmentsCertificateResultOutput) ToLookupConnectedEnvironmentsCertificateResultOutputWithContext(ctx context.Context) LookupConnectedEnvironmentsCertificateResultOutput {
	return o
}

func (o LookupConnectedEnvironmentsCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentsCertificateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentsCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentsCertificateResultOutput) Properties() CertificateResponsePropertiesOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) CertificateResponseProperties {
		return v.Properties
	}).(CertificateResponsePropertiesOutput)
}

func (o LookupConnectedEnvironmentsCertificateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConnectedEnvironmentsCertificateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConnectedEnvironmentsCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectedEnvironmentsCertificateResultOutput{})
}
