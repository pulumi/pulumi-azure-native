


package v20190201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationGateway(ctx *pulumi.Context, args *LookupApplicationGatewayArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGatewayResult, error) {
	var rv LookupApplicationGatewayResult
	err := ctx.Invoke("azure-native:network/v20190201:getApplicationGateway", args, &rv, opts...)
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
	FirewallPolicy                      *SubResourceResponse                                           `pulumi:"firewallPolicy"`
	FrontendIPConfigurations            []ApplicationGatewayFrontendIPConfigurationResponse            `pulumi:"frontendIPConfigurations"`
	FrontendPorts                       []ApplicationGatewayFrontendPortResponse                       `pulumi:"frontendPorts"`
	GatewayIPConfigurations             []ApplicationGatewayIPConfigurationResponse                    `pulumi:"gatewayIPConfigurations"`
	HttpListeners                       []ApplicationGatewayHttpListenerResponse                       `pulumi:"httpListeners"`
	Id                                  *string                                                        `pulumi:"id"`
	Identity                            *ManagedServiceIdentityResponse                                `pulumi:"identity"`
	Location                            *string                                                        `pulumi:"location"`
	Name                                string                                                         `pulumi:"name"`
	OperationalState                    string                                                         `pulumi:"operationalState"`
	Probes                              []ApplicationGatewayProbeResponse                              `pulumi:"probes"`
	ProvisioningState                   *string                                                        `pulumi:"provisioningState"`
	RedirectConfigurations              []ApplicationGatewayRedirectConfigurationResponse              `pulumi:"redirectConfigurations"`
	RequestRoutingRules                 []ApplicationGatewayRequestRoutingRuleResponse                 `pulumi:"requestRoutingRules"`
	ResourceGuid                        *string                                                        `pulumi:"resourceGuid"`
	RewriteRuleSets                     []ApplicationGatewayRewriteRuleSetResponse                     `pulumi:"rewriteRuleSets"`
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

func (o LookupApplicationGatewayResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGatewayResultOutput) FirewallPolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *SubResourceResponse { return v.FirewallPolicy }).(SubResourceResponsePtrOutput)
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

func (o LookupApplicationGatewayResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationGatewayResultOutput) OperationalState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.OperationalState }).(pulumi.StringOutput)
}

func (o LookupApplicationGatewayResultOutput) Probes() ApplicationGatewayProbeResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayProbeResponse { return v.Probes }).(ApplicationGatewayProbeResponseArrayOutput)
}

func (o LookupApplicationGatewayResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
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

func (o LookupApplicationGatewayResultOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationGatewayResultOutput) RewriteRuleSets() ApplicationGatewayRewriteRuleSetResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayRewriteRuleSetResponse {
		return v.RewriteRuleSets
	}).(ApplicationGatewayRewriteRuleSetResponseArrayOutput)
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

func (o LookupApplicationGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
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
