


package v20180801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationGateway(ctx *pulumi.Context, args *LookupApplicationGatewayArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGatewayResult, error) {
	var rv LookupApplicationGatewayResult
	err := ctx.Invoke("azure-native:network/v20180801:getApplicationGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationGatewayArgs struct {
	ApplicationGatewayName string `pulumi:"applicationGatewayName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupApplicationGatewayResult struct {
	AuthenticationCertificates          []ApplicationGatewayAuthenticationCertificateResponse          `pulumi:"authenticationCertificates"`
	AutoscaleConfiguration              *ApplicationGatewayAutoscaleConfigurationResponse              `pulumi:"autoscaleConfiguration"`
	BackendAddressPools                 []ApplicationGatewayBackendAddressPoolResponse                 `pulumi:"backendAddressPools"`
	BackendHttpSettingsCollection       []ApplicationGatewayBackendHttpSettingsResponse                `pulumi:"backendHttpSettingsCollection"`
	CustomErrorConfigurations           []ApplicationGatewayCustomErrorResponse                        `pulumi:"customErrorConfigurations"`
	EnableFips                          *bool                                                          `pulumi:"enableFips"`
	EnableHttp2                         *bool                                                          `pulumi:"enableHttp2"`
	Etag                                *string                                                        `pulumi:"etag"`
	FrontendIPConfigurations            []ApplicationGatewayFrontendIPConfigurationResponse            `pulumi:"frontendIPConfigurations"`
	FrontendPorts                       []ApplicationGatewayFrontendPortResponse                       `pulumi:"frontendPorts"`
	GatewayIPConfigurations             []ApplicationGatewayIPConfigurationResponse                    `pulumi:"gatewayIPConfigurations"`
	HttpListeners                       []ApplicationGatewayHttpListenerResponse                       `pulumi:"httpListeners"`
	Id                                  *string                                                        `pulumi:"id"`
	Location                            *string                                                        `pulumi:"location"`
	Name                                string                                                         `pulumi:"name"`
	OperationalState                    string                                                         `pulumi:"operationalState"`
	Probes                              []ApplicationGatewayProbeResponse                              `pulumi:"probes"`
	ProvisioningState                   *string                                                        `pulumi:"provisioningState"`
	RedirectConfigurations              []ApplicationGatewayRedirectConfigurationResponse              `pulumi:"redirectConfigurations"`
	RequestRoutingRules                 []ApplicationGatewayRequestRoutingRuleResponse                 `pulumi:"requestRoutingRules"`
	ResourceGuid                        *string                                                        `pulumi:"resourceGuid"`
	Sku                                 *ApplicationGatewaySkuResponse                                 `pulumi:"sku"`
	SslCertificates                     []ApplicationGatewaySslCertificateResponse                     `pulumi:"sslCertificates"`
	SslPolicy                           *ApplicationGatewaySslPolicyResponse                           `pulumi:"sslPolicy"`
	Tags                                map[string]string                                              `pulumi:"tags"`
	TrustedRootCertificates             []ApplicationGatewayTrustedRootCertificateResponse             `pulumi:"trustedRootCertificates"`
	Type                                string                                                         `pulumi:"type"`
	UrlPathMaps                         []ApplicationGatewayUrlPathMapResponse                         `pulumi:"urlPathMaps"`
	WebApplicationFirewallConfiguration *ApplicationGatewayWebApplicationFirewallConfigurationResponse `pulumi:"webApplicationFirewallConfiguration"`
	Zones                               []string                                                       `pulumi:"zones"`
}
