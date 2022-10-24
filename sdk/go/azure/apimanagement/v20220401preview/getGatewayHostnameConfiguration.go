


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayHostnameConfiguration(ctx *pulumi.Context, args *LookupGatewayHostnameConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupGatewayHostnameConfigurationResult, error) {
	var rv LookupGatewayHostnameConfigurationResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:getGatewayHostnameConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayHostnameConfigurationArgs struct {
	GatewayId         string `pulumi:"gatewayId"`
	HcId              string `pulumi:"hcId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupGatewayHostnameConfigurationResult struct {
	CertificateId              *string `pulumi:"certificateId"`
	Hostname                   *string `pulumi:"hostname"`
	Http2Enabled               *bool   `pulumi:"http2Enabled"`
	Id                         string  `pulumi:"id"`
	Name                       string  `pulumi:"name"`
	NegotiateClientCertificate *bool   `pulumi:"negotiateClientCertificate"`
	Tls10Enabled               *bool   `pulumi:"tls10Enabled"`
	Tls11Enabled               *bool   `pulumi:"tls11Enabled"`
	Type                       string  `pulumi:"type"`
}

func LookupGatewayHostnameConfigurationOutput(ctx *pulumi.Context, args LookupGatewayHostnameConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayHostnameConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayHostnameConfigurationResult, error) {
			args := v.(LookupGatewayHostnameConfigurationArgs)
			r, err := LookupGatewayHostnameConfiguration(ctx, &args, opts...)
			var s LookupGatewayHostnameConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayHostnameConfigurationResultOutput)
}

type LookupGatewayHostnameConfigurationOutputArgs struct {
	GatewayId         pulumi.StringInput `pulumi:"gatewayId"`
	HcId              pulumi.StringInput `pulumi:"hcId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGatewayHostnameConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayHostnameConfigurationArgs)(nil)).Elem()
}


type LookupGatewayHostnameConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayHostnameConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayHostnameConfigurationResult)(nil)).Elem()
}

func (o LookupGatewayHostnameConfigurationResultOutput) ToLookupGatewayHostnameConfigurationResultOutput() LookupGatewayHostnameConfigurationResultOutput {
	return o
}

func (o LookupGatewayHostnameConfigurationResultOutput) ToLookupGatewayHostnameConfigurationResultOutputWithContext(ctx context.Context) LookupGatewayHostnameConfigurationResultOutput {
	return o
}

func (o LookupGatewayHostnameConfigurationResultOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *string { return v.CertificateId }).(pulumi.StringPtrOutput)
}

func (o LookupGatewayHostnameConfigurationResultOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o LookupGatewayHostnameConfigurationResultOutput) Http2Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *bool { return v.Http2Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupGatewayHostnameConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayHostnameConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGatewayHostnameConfigurationResultOutput) NegotiateClientCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *bool { return v.NegotiateClientCertificate }).(pulumi.BoolPtrOutput)
}

func (o LookupGatewayHostnameConfigurationResultOutput) Tls10Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *bool { return v.Tls10Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupGatewayHostnameConfigurationResultOutput) Tls11Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *bool { return v.Tls11Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupGatewayHostnameConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayHostnameConfigurationResultOutput{})
}
