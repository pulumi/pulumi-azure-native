


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVpnServerConfiguration(ctx *pulumi.Context, args *LookupVpnServerConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupVpnServerConfigurationResult, error) {
	var rv LookupVpnServerConfigurationResult
	err := ctx.Invoke("azure-native:network/v20200301:getVpnServerConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnServerConfigurationArgs struct {
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	VpnServerConfigurationName string `pulumi:"vpnServerConfigurationName"`
}


type LookupVpnServerConfigurationResult struct {
	AadAuthenticationParameters  *AadAuthenticationParametersResponse                 `pulumi:"aadAuthenticationParameters"`
	Etag                         string                                               `pulumi:"etag"`
	Id                           *string                                              `pulumi:"id"`
	Location                     *string                                              `pulumi:"location"`
	Name                         string                                               `pulumi:"name"`
	P2SVpnGateways               []P2SVpnGatewayResponse                              `pulumi:"p2SVpnGateways"`
	ProvisioningState            string                                               `pulumi:"provisioningState"`
	RadiusClientRootCertificates []VpnServerConfigRadiusClientRootCertificateResponse `pulumi:"radiusClientRootCertificates"`
	RadiusServerAddress          *string                                              `pulumi:"radiusServerAddress"`
	RadiusServerRootCertificates []VpnServerConfigRadiusServerRootCertificateResponse `pulumi:"radiusServerRootCertificates"`
	RadiusServerSecret           *string                                              `pulumi:"radiusServerSecret"`
	RadiusServers                []RadiusServerResponse                               `pulumi:"radiusServers"`
	Tags                         map[string]string                                    `pulumi:"tags"`
	Type                         string                                               `pulumi:"type"`
	VpnAuthenticationTypes       []string                                             `pulumi:"vpnAuthenticationTypes"`
	VpnClientIpsecPolicies       []IpsecPolicyResponse                                `pulumi:"vpnClientIpsecPolicies"`
	VpnClientRevokedCertificates []VpnServerConfigVpnClientRevokedCertificateResponse `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    []VpnServerConfigVpnClientRootCertificateResponse    `pulumi:"vpnClientRootCertificates"`
	VpnProtocols                 []string                                             `pulumi:"vpnProtocols"`
}

func LookupVpnServerConfigurationOutput(ctx *pulumi.Context, args LookupVpnServerConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupVpnServerConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnServerConfigurationResult, error) {
			args := v.(LookupVpnServerConfigurationArgs)
			r, err := LookupVpnServerConfiguration(ctx, &args, opts...)
			var s LookupVpnServerConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpnServerConfigurationResultOutput)
}

type LookupVpnServerConfigurationOutputArgs struct {
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
	VpnServerConfigurationName pulumi.StringInput `pulumi:"vpnServerConfigurationName"`
}

func (LookupVpnServerConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnServerConfigurationArgs)(nil)).Elem()
}


type LookupVpnServerConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupVpnServerConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnServerConfigurationResult)(nil)).Elem()
}

func (o LookupVpnServerConfigurationResultOutput) ToLookupVpnServerConfigurationResultOutput() LookupVpnServerConfigurationResultOutput {
	return o
}

func (o LookupVpnServerConfigurationResultOutput) ToLookupVpnServerConfigurationResultOutputWithContext(ctx context.Context) LookupVpnServerConfigurationResultOutput {
	return o
}

func (o LookupVpnServerConfigurationResultOutput) AadAuthenticationParameters() AadAuthenticationParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) *AadAuthenticationParametersResponse {
		return v.AadAuthenticationParameters
	}).(AadAuthenticationParametersResponsePtrOutput)
}

func (o LookupVpnServerConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVpnServerConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVpnServerConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVpnServerConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVpnServerConfigurationResultOutput) P2SVpnGateways() P2SVpnGatewayResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) []P2SVpnGatewayResponse { return v.P2SVpnGateways }).(P2SVpnGatewayResponseArrayOutput)
}

func (o LookupVpnServerConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVpnServerConfigurationResultOutput) RadiusClientRootCertificates() VpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) []VpnServerConfigRadiusClientRootCertificateResponse {
		return v.RadiusClientRootCertificates
	}).(VpnServerConfigRadiusClientRootCertificateResponseArrayOutput)
}

func (o LookupVpnServerConfigurationResultOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) *string { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

func (o LookupVpnServerConfigurationResultOutput) RadiusServerRootCertificates() VpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) []VpnServerConfigRadiusServerRootCertificateResponse {
		return v.RadiusServerRootCertificates
	}).(VpnServerConfigRadiusServerRootCertificateResponseArrayOutput)
}

func (o LookupVpnServerConfigurationResultOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) *string { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

func (o LookupVpnServerConfigurationResultOutput) RadiusServers() RadiusServerResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) []RadiusServerResponse { return v.RadiusServers }).(RadiusServerResponseArrayOutput)
}

func (o LookupVpnServerConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVpnServerConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVpnServerConfigurationResultOutput) VpnAuthenticationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) []string { return v.VpnAuthenticationTypes }).(pulumi.StringArrayOutput)
}

func (o LookupVpnServerConfigurationResultOutput) VpnClientIpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) []IpsecPolicyResponse { return v.VpnClientIpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o LookupVpnServerConfigurationResultOutput) VpnClientRevokedCertificates() VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) []VpnServerConfigVpnClientRevokedCertificateResponse {
		return v.VpnClientRevokedCertificates
	}).(VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput)
}

func (o LookupVpnServerConfigurationResultOutput) VpnClientRootCertificates() VpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) []VpnServerConfigVpnClientRootCertificateResponse {
		return v.VpnClientRootCertificates
	}).(VpnServerConfigVpnClientRootCertificateResponseArrayOutput)
}

func (o LookupVpnServerConfigurationResultOutput) VpnProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpnServerConfigurationResult) []string { return v.VpnProtocols }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnServerConfigurationResultOutput{})
}
