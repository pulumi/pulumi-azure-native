


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupP2sVpnServerConfiguration(ctx *pulumi.Context, args *LookupP2sVpnServerConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupP2sVpnServerConfigurationResult, error) {
	var rv LookupP2sVpnServerConfigurationResult
	err := ctx.Invoke("azure-native:network:getP2sVpnServerConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupP2sVpnServerConfigurationArgs struct {
	P2SVpnServerConfigurationName string `pulumi:"p2SVpnServerConfigurationName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	VirtualWanName                string `pulumi:"virtualWanName"`
}


type LookupP2sVpnServerConfigurationResult struct {
	Etag                                           string                                                  `pulumi:"etag"`
	Id                                             *string                                                 `pulumi:"id"`
	Name                                           *string                                                 `pulumi:"name"`
	P2SVpnGateways                                 []SubResourceResponse                                   `pulumi:"p2SVpnGateways"`
	P2SVpnServerConfigRadiusClientRootCertificates []P2SVpnServerConfigRadiusClientRootCertificateResponse `pulumi:"p2SVpnServerConfigRadiusClientRootCertificates"`
	P2SVpnServerConfigRadiusServerRootCertificates []P2SVpnServerConfigRadiusServerRootCertificateResponse `pulumi:"p2SVpnServerConfigRadiusServerRootCertificates"`
	P2SVpnServerConfigVpnClientRevokedCertificates []P2SVpnServerConfigVpnClientRevokedCertificateResponse `pulumi:"p2SVpnServerConfigVpnClientRevokedCertificates"`
	P2SVpnServerConfigVpnClientRootCertificates    []P2SVpnServerConfigVpnClientRootCertificateResponse    `pulumi:"p2SVpnServerConfigVpnClientRootCertificates"`
	ProvisioningState                              string                                                  `pulumi:"provisioningState"`
	RadiusServerAddress                            *string                                                 `pulumi:"radiusServerAddress"`
	RadiusServerSecret                             *string                                                 `pulumi:"radiusServerSecret"`
	VpnClientIpsecPolicies                         []IpsecPolicyResponse                                   `pulumi:"vpnClientIpsecPolicies"`
	VpnProtocols                                   []string                                                `pulumi:"vpnProtocols"`
}

func LookupP2sVpnServerConfigurationOutput(ctx *pulumi.Context, args LookupP2sVpnServerConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupP2sVpnServerConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupP2sVpnServerConfigurationResult, error) {
			args := v.(LookupP2sVpnServerConfigurationArgs)
			r, err := LookupP2sVpnServerConfiguration(ctx, &args, opts...)
			var s LookupP2sVpnServerConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupP2sVpnServerConfigurationResultOutput)
}

type LookupP2sVpnServerConfigurationOutputArgs struct {
	P2SVpnServerConfigurationName pulumi.StringInput `pulumi:"p2SVpnServerConfigurationName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualWanName                pulumi.StringInput `pulumi:"virtualWanName"`
}

func (LookupP2sVpnServerConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupP2sVpnServerConfigurationArgs)(nil)).Elem()
}


type LookupP2sVpnServerConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupP2sVpnServerConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupP2sVpnServerConfigurationResult)(nil)).Elem()
}

func (o LookupP2sVpnServerConfigurationResultOutput) ToLookupP2sVpnServerConfigurationResultOutput() LookupP2sVpnServerConfigurationResultOutput {
	return o
}

func (o LookupP2sVpnServerConfigurationResultOutput) ToLookupP2sVpnServerConfigurationResultOutputWithContext(ctx context.Context) LookupP2sVpnServerConfigurationResultOutput {
	return o
}

func (o LookupP2sVpnServerConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) P2SVpnGateways() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) []SubResourceResponse { return v.P2SVpnGateways }).(SubResourceResponseArrayOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) P2SVpnServerConfigRadiusClientRootCertificates() P2SVpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) []P2SVpnServerConfigRadiusClientRootCertificateResponse {
		return v.P2SVpnServerConfigRadiusClientRootCertificates
	}).(P2SVpnServerConfigRadiusClientRootCertificateResponseArrayOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) P2SVpnServerConfigRadiusServerRootCertificates() P2SVpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) []P2SVpnServerConfigRadiusServerRootCertificateResponse {
		return v.P2SVpnServerConfigRadiusServerRootCertificates
	}).(P2SVpnServerConfigRadiusServerRootCertificateResponseArrayOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) P2SVpnServerConfigVpnClientRevokedCertificates() P2SVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) []P2SVpnServerConfigVpnClientRevokedCertificateResponse {
		return v.P2SVpnServerConfigVpnClientRevokedCertificates
	}).(P2SVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) P2SVpnServerConfigVpnClientRootCertificates() P2SVpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) []P2SVpnServerConfigVpnClientRootCertificateResponse {
		return v.P2SVpnServerConfigVpnClientRootCertificates
	}).(P2SVpnServerConfigVpnClientRootCertificateResponseArrayOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) *string { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) *string { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) VpnClientIpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) []IpsecPolicyResponse { return v.VpnClientIpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o LookupP2sVpnServerConfigurationResultOutput) VpnProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupP2sVpnServerConfigurationResult) []string { return v.VpnProtocols }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupP2sVpnServerConfigurationResultOutput{})
}
