


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVpnServerConfiguration(ctx *pulumi.Context, args *LookupVpnServerConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupVpnServerConfigurationResult, error) {
	var rv LookupVpnServerConfigurationResult
	err := ctx.Invoke("azure-native:network/v20200601:getVpnServerConfiguration", args, &rv, opts...)
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
