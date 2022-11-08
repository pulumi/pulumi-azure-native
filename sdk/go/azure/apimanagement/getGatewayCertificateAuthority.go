


package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayCertificateAuthority(ctx *pulumi.Context, args *LookupGatewayCertificateAuthorityArgs, opts ...pulumi.InvokeOption) (*LookupGatewayCertificateAuthorityResult, error) {
	var rv LookupGatewayCertificateAuthorityResult
	err := ctx.Invoke("azure-native:apimanagement:getGatewayCertificateAuthority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayCertificateAuthorityArgs struct {
	CertificateId     string `pulumi:"certificateId"`
	GatewayId         string `pulumi:"gatewayId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupGatewayCertificateAuthorityResult struct {
	Id        string `pulumi:"id"`
	IsTrusted *bool  `pulumi:"isTrusted"`
	Name      string `pulumi:"name"`
	Type      string `pulumi:"type"`
}

func LookupGatewayCertificateAuthorityOutput(ctx *pulumi.Context, args LookupGatewayCertificateAuthorityOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayCertificateAuthorityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayCertificateAuthorityResult, error) {
			args := v.(LookupGatewayCertificateAuthorityArgs)
			r, err := LookupGatewayCertificateAuthority(ctx, &args, opts...)
			var s LookupGatewayCertificateAuthorityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayCertificateAuthorityResultOutput)
}

type LookupGatewayCertificateAuthorityOutputArgs struct {
	CertificateId     pulumi.StringInput `pulumi:"certificateId"`
	GatewayId         pulumi.StringInput `pulumi:"gatewayId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGatewayCertificateAuthorityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCertificateAuthorityArgs)(nil)).Elem()
}


type LookupGatewayCertificateAuthorityResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayCertificateAuthorityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCertificateAuthorityResult)(nil)).Elem()
}

func (o LookupGatewayCertificateAuthorityResultOutput) ToLookupGatewayCertificateAuthorityResultOutput() LookupGatewayCertificateAuthorityResultOutput {
	return o
}

func (o LookupGatewayCertificateAuthorityResultOutput) ToLookupGatewayCertificateAuthorityResultOutputWithContext(ctx context.Context) LookupGatewayCertificateAuthorityResultOutput {
	return o
}

func (o LookupGatewayCertificateAuthorityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCertificateAuthorityResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayCertificateAuthorityResultOutput) IsTrusted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGatewayCertificateAuthorityResult) *bool { return v.IsTrusted }).(pulumi.BoolPtrOutput)
}

func (o LookupGatewayCertificateAuthorityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCertificateAuthorityResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGatewayCertificateAuthorityResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCertificateAuthorityResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayCertificateAuthorityResultOutput{})
}
