


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationGateway(ctx *pulumi.Context, args *LookupApplicationGatewayArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGatewayResult, error) {
	var rv LookupApplicationGatewayResult
	err := ctx.Invoke("azure-native:network/v20220701:getApplicationGateway", args, &rv, opts...)
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
	BackendSettingsCollection           []ApplicationGatewayBackendSettingsResponse                    `pulumi:"backendSettingsCollection"`
	CustomErrorConfigurations           []ApplicationGatewayCustomErrorResponse                        `pulumi:"customErrorConfigurations"`
	EnableFips                          *bool                                                          `pulumi:"enableFips"`
	EnableHttp2                         *bool                                                          `pulumi:"enableHttp2"`
	Etag                                string                                                         `pulumi:"etag"`
	FirewallPolicy                      *SubResourceResponse                                           `pulumi:"firewallPolicy"`
	ForceFirewallPolicyAssociation      *bool                                                          `pulumi:"forceFirewallPolicyAssociation"`
	FrontendIPConfigurations            []ApplicationGatewayFrontendIPConfigurationResponse            `pulumi:"frontendIPConfigurations"`
	FrontendPorts                       []ApplicationGatewayFrontendPortResponse                       `pulumi:"frontendPorts"`
	GatewayIPConfigurations             []ApplicationGatewayIPConfigurationResponse                    `pulumi:"gatewayIPConfigurations"`
	GlobalConfiguration                 *ApplicationGatewayGlobalConfigurationResponse                 `pulumi:"globalConfiguration"`
	HttpListeners                       []ApplicationGatewayHttpListenerResponse                       `pulumi:"httpListeners"`
	Id                                  *string                                                        `pulumi:"id"`
	Identity                            *ManagedServiceIdentityResponse                                `pulumi:"identity"`
	Listeners                           []ApplicationGatewayListenerResponse                           `pulumi:"listeners"`
	LoadDistributionPolicies            []ApplicationGatewayLoadDistributionPolicyResponse             `pulumi:"loadDistributionPolicies"`
	Location                            *string                                                        `pulumi:"location"`
	Name                                string                                                         `pulumi:"name"`
	OperationalState                    string                                                         `pulumi:"operationalState"`
	PrivateEndpointConnections          []ApplicationGatewayPrivateEndpointConnectionResponse          `pulumi:"privateEndpointConnections"`
	PrivateLinkConfigurations           []ApplicationGatewayPrivateLinkConfigurationResponse           `pulumi:"privateLinkConfigurations"`
	Probes                              []ApplicationGatewayProbeResponse                              `pulumi:"probes"`
	ProvisioningState                   string                                                         `pulumi:"provisioningState"`
	RedirectConfigurations              []ApplicationGatewayRedirectConfigurationResponse              `pulumi:"redirectConfigurations"`
	RequestRoutingRules                 []ApplicationGatewayRequestRoutingRuleResponse                 `pulumi:"requestRoutingRules"`
	ResourceGuid                        string                                                         `pulumi:"resourceGuid"`
	RewriteRuleSets                     []ApplicationGatewayRewriteRuleSetResponse                     `pulumi:"rewriteRuleSets"`
	RoutingRules                        []ApplicationGatewayRoutingRuleResponse                        `pulumi:"routingRules"`
	Sku                                 *ApplicationGatewaySkuResponse                                 `pulumi:"sku"`
	SslCertificates                     []ApplicationGatewaySslCertificateResponse                     `pulumi:"sslCertificates"`
	SslPolicy                           *ApplicationGatewaySslPolicyResponse                           `pulumi:"sslPolicy"`
	SslProfiles                         []ApplicationGatewaySslProfileResponse                         `pulumi:"sslProfiles"`
	Tags                                map[string]string                                              `pulumi:"tags"`
	TrustedClientCertificates           []ApplicationGatewayTrustedClientCertificateResponse           `pulumi:"trustedClientCertificates"`
	TrustedRootCertificates             []ApplicationGatewayTrustedRootCertificateResponse             `pulumi:"trustedRootCertificates"`
	Type                                string                                                         `pulumi:"type"`
	UrlPathMaps                         []ApplicationGatewayUrlPathMapResponse                         `pulumi:"urlPathMaps"`
	WebApplicationFirewallConfiguration *ApplicationGatewayWebApplicationFirewallConfigurationResponse `pulumi:"webApplicationFirewallConfiguration"`
	Zones                               []string                                                       `pulumi:"zones"`
}

func LookupApplicationGatewayOutput(ctx *pulumi.Context, args LookupApplicationGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationGatewayResult, error) {
			args := v.(LookupApplicationGatewayArgs)
			r, err := LookupApplicationGateway(ctx, &args, opts...)
			var s LookupApplicationGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationGatewayResultOutput)
}

type LookupApplicationGatewayOutputArgs struct {
	ApplicationGatewayName pulumi.StringInput `pulumi:"applicationGatewayName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGatewayArgs)(nil)).Elem()
}


type LookupApplicationGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGatewayResult)(nil)).Elem()
}

func (o LookupApplicationGatewayResultOutput) ToLookupApplicationGatewayResultOutput() LookupApplicationGatewayResultOutput {
	return o
}

func (o LookupApplicationGatewayResultOutput) ToLookupApplicationGatewayResultOutputWithContext(ctx context.Context) LookupApplicationGatewayResultOutput {
	return o
}

func (o LookupApplicationGatewayResultOutput) AuthenticationCertificates() ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayAuthenticationCertificateResponse {
		return v.AuthenticationCertificates
	}).(ApplicationGatewayAuthenticationCertificateResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) AutoscaleConfiguration() ApplicationGatewayAutoscaleConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ApplicationGatewayAutoscaleConfigurationResponse {
		return v.AutoscaleConfiguration
	}).(ApplicationGatewayAutoscaleConfigurationResponsePtrOutput)
}

func (o LookupApplicationGatewayResultOutput) BackendAddressPools() ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayBackendAddressPoolResponse {
		return v.BackendAddressPools
	}).(ApplicationGatewayBackendAddressPoolResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) BackendHttpSettingsCollection() ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayBackendHttpSettingsResponse {
		return v.BackendHttpSettingsCollection
	}).(ApplicationGatewayBackendHttpSettingsResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) BackendSettingsCollection() ApplicationGatewayBackendSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayBackendSettingsResponse {
		return v.BackendSettingsCollection
	}).(ApplicationGatewayBackendSettingsResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) CustomErrorConfigurations() ApplicationGatewayCustomErrorResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayCustomErrorResponse {
		return v.CustomErrorConfigurations
	}).(ApplicationGatewayCustomErrorResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) EnableFips() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *bool { return v.EnableFips }).(pulumi.BoolPtrOutput)
}

func (o LookupApplicationGatewayResultOutput) EnableHttp2() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *bool { return v.EnableHttp2 }).(pulumi.BoolPtrOutput)
}

func (o LookupApplicationGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApplicationGatewayResultOutput) FirewallPolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *SubResourceResponse { return v.FirewallPolicy }).(SubResourceResponsePtrOutput)
}

func (o LookupApplicationGatewayResultOutput) ForceFirewallPolicyAssociation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *bool { return v.ForceFirewallPolicyAssociation }).(pulumi.BoolPtrOutput)
}

func (o LookupApplicationGatewayResultOutput) FrontendIPConfigurations() ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayFrontendIPConfigurationResponse {
		return v.FrontendIPConfigurations
	}).(ApplicationGatewayFrontendIPConfigurationResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) FrontendPorts() ApplicationGatewayFrontendPortResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayFrontendPortResponse {
		return v.FrontendPorts
	}).(ApplicationGatewayFrontendPortResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) GatewayIPConfigurations() ApplicationGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayIPConfigurationResponse {
		return v.GatewayIPConfigurations
	}).(ApplicationGatewayIPConfigurationResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) GlobalConfiguration() ApplicationGatewayGlobalConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ApplicationGatewayGlobalConfigurationResponse {
		return v.GlobalConfiguration
	}).(ApplicationGatewayGlobalConfigurationResponsePtrOutput)
}

func (o LookupApplicationGatewayResultOutput) HttpListeners() ApplicationGatewayHttpListenerResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayHttpListenerResponse {
		return v.HttpListeners
	}).(ApplicationGatewayHttpListenerResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGatewayResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupApplicationGatewayResultOutput) Listeners() ApplicationGatewayListenerResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayListenerResponse { return v.Listeners }).(ApplicationGatewayListenerResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) LoadDistributionPolicies() ApplicationGatewayLoadDistributionPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayLoadDistributionPolicyResponse {
		return v.LoadDistributionPolicies
	}).(ApplicationGatewayLoadDistributionPolicyResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationGatewayResultOutput) OperationalState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.OperationalState }).(pulumi.StringOutput)
}

func (o LookupApplicationGatewayResultOutput) PrivateEndpointConnections() ApplicationGatewayPrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayPrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(ApplicationGatewayPrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) PrivateLinkConfigurations() ApplicationGatewayPrivateLinkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayPrivateLinkConfigurationResponse {
		return v.PrivateLinkConfigurations
	}).(ApplicationGatewayPrivateLinkConfigurationResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) Probes() ApplicationGatewayProbeResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayProbeResponse { return v.Probes }).(ApplicationGatewayProbeResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApplicationGatewayResultOutput) RedirectConfigurations() ApplicationGatewayRedirectConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayRedirectConfigurationResponse {
		return v.RedirectConfigurations
	}).(ApplicationGatewayRedirectConfigurationResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) RequestRoutingRules() ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayRequestRoutingRuleResponse {
		return v.RequestRoutingRules
	}).(ApplicationGatewayRequestRoutingRuleResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupApplicationGatewayResultOutput) RewriteRuleSets() ApplicationGatewayRewriteRuleSetResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayRewriteRuleSetResponse {
		return v.RewriteRuleSets
	}).(ApplicationGatewayRewriteRuleSetResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) RoutingRules() ApplicationGatewayRoutingRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayRoutingRuleResponse { return v.RoutingRules }).(ApplicationGatewayRoutingRuleResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) Sku() ApplicationGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ApplicationGatewaySkuResponse { return v.Sku }).(ApplicationGatewaySkuResponsePtrOutput)
}

func (o LookupApplicationGatewayResultOutput) SslCertificates() ApplicationGatewaySslCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewaySslCertificateResponse {
		return v.SslCertificates
	}).(ApplicationGatewaySslCertificateResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) SslPolicy() ApplicationGatewaySslPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ApplicationGatewaySslPolicyResponse { return v.SslPolicy }).(ApplicationGatewaySslPolicyResponsePtrOutput)
}

func (o LookupApplicationGatewayResultOutput) SslProfiles() ApplicationGatewaySslProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewaySslProfileResponse { return v.SslProfiles }).(ApplicationGatewaySslProfileResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplicationGatewayResultOutput) TrustedClientCertificates() ApplicationGatewayTrustedClientCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayTrustedClientCertificateResponse {
		return v.TrustedClientCertificates
	}).(ApplicationGatewayTrustedClientCertificateResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) TrustedRootCertificates() ApplicationGatewayTrustedRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayTrustedRootCertificateResponse {
		return v.TrustedRootCertificates
	}).(ApplicationGatewayTrustedRootCertificateResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApplicationGatewayResultOutput) UrlPathMaps() ApplicationGatewayUrlPathMapResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayUrlPathMapResponse { return v.UrlPathMaps }).(ApplicationGatewayUrlPathMapResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) WebApplicationFirewallConfiguration() ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ApplicationGatewayWebApplicationFirewallConfigurationResponse {
		return v.WebApplicationFirewallConfiguration
	}).(ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput)
}

func (o LookupApplicationGatewayResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationGatewayResultOutput{})
}
