


package v20181201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationGateway struct {
	pulumi.CustomResourceState

	AuthenticationCertificates          ApplicationGatewayAuthenticationCertificateResponseArrayOutput         `pulumi:"authenticationCertificates"`
	AutoscaleConfiguration              ApplicationGatewayAutoscaleConfigurationResponsePtrOutput              `pulumi:"autoscaleConfiguration"`
	BackendAddressPools                 ApplicationGatewayBackendAddressPoolResponseArrayOutput                `pulumi:"backendAddressPools"`
	BackendHttpSettingsCollection       ApplicationGatewayBackendHttpSettingsResponseArrayOutput               `pulumi:"backendHttpSettingsCollection"`
	CustomErrorConfigurations           ApplicationGatewayCustomErrorResponseArrayOutput                       `pulumi:"customErrorConfigurations"`
	EnableFips                          pulumi.BoolPtrOutput                                                   `pulumi:"enableFips"`
	EnableHttp2                         pulumi.BoolPtrOutput                                                   `pulumi:"enableHttp2"`
	Etag                                pulumi.StringPtrOutput                                                 `pulumi:"etag"`
	FirewallPolicy                      SubResourceResponsePtrOutput                                           `pulumi:"firewallPolicy"`
	FrontendIPConfigurations            ApplicationGatewayFrontendIPConfigurationResponseArrayOutput           `pulumi:"frontendIPConfigurations"`
	FrontendPorts                       ApplicationGatewayFrontendPortResponseArrayOutput                      `pulumi:"frontendPorts"`
	GatewayIPConfigurations             ApplicationGatewayIPConfigurationResponseArrayOutput                   `pulumi:"gatewayIPConfigurations"`
	HttpListeners                       ApplicationGatewayHttpListenerResponseArrayOutput                      `pulumi:"httpListeners"`
	Identity                            ManagedServiceIdentityResponsePtrOutput                                `pulumi:"identity"`
	Location                            pulumi.StringPtrOutput                                                 `pulumi:"location"`
	Name                                pulumi.StringOutput                                                    `pulumi:"name"`
	OperationalState                    pulumi.StringOutput                                                    `pulumi:"operationalState"`
	Probes                              ApplicationGatewayProbeResponseArrayOutput                             `pulumi:"probes"`
	ProvisioningState                   pulumi.StringPtrOutput                                                 `pulumi:"provisioningState"`
	RedirectConfigurations              ApplicationGatewayRedirectConfigurationResponseArrayOutput             `pulumi:"redirectConfigurations"`
	RequestRoutingRules                 ApplicationGatewayRequestRoutingRuleResponseArrayOutput                `pulumi:"requestRoutingRules"`
	ResourceGuid                        pulumi.StringPtrOutput                                                 `pulumi:"resourceGuid"`
	RewriteRuleSets                     ApplicationGatewayRewriteRuleSetResponseArrayOutput                    `pulumi:"rewriteRuleSets"`
	Sku                                 ApplicationGatewaySkuResponsePtrOutput                                 `pulumi:"sku"`
	SslCertificates                     ApplicationGatewaySslCertificateResponseArrayOutput                    `pulumi:"sslCertificates"`
	SslPolicy                           ApplicationGatewaySslPolicyResponsePtrOutput                           `pulumi:"sslPolicy"`
	Tags                                pulumi.StringMapOutput                                                 `pulumi:"tags"`
	TrustedRootCertificates             ApplicationGatewayTrustedRootCertificateResponseArrayOutput            `pulumi:"trustedRootCertificates"`
	Type                                pulumi.StringOutput                                                    `pulumi:"type"`
	UrlPathMaps                         ApplicationGatewayUrlPathMapResponseArrayOutput                        `pulumi:"urlPathMaps"`
	WebApplicationFirewallConfiguration ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput `pulumi:"webApplicationFirewallConfiguration"`
	Zones                               pulumi.StringArrayOutput                                               `pulumi:"zones"`
}


func NewApplicationGateway(ctx *pulumi.Context,
	name string, args *ApplicationGatewayArgs, opts ...pulumi.ResourceOption) (*ApplicationGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ApplicationGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationGateway
	err := ctx.RegisterResource("azure-native:network/v20181201:ApplicationGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGatewayState, opts ...pulumi.ResourceOption) (*ApplicationGateway, error) {
	var resource ApplicationGateway
	err := ctx.ReadResource("azure-native:network/v20181201:ApplicationGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationGatewayState struct {
}

type ApplicationGatewayState struct {
}

func (ApplicationGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayState)(nil)).Elem()
}

type applicationGatewayArgs struct {
	ApplicationGatewayName              *string                                                `pulumi:"applicationGatewayName"`
	AuthenticationCertificates          []ApplicationGatewayAuthenticationCertificate          `pulumi:"authenticationCertificates"`
	AutoscaleConfiguration              *ApplicationGatewayAutoscaleConfiguration              `pulumi:"autoscaleConfiguration"`
	BackendAddressPools                 []ApplicationGatewayBackendAddressPool                 `pulumi:"backendAddressPools"`
	BackendHttpSettingsCollection       []ApplicationGatewayBackendHttpSettings                `pulumi:"backendHttpSettingsCollection"`
	CustomErrorConfigurations           []ApplicationGatewayCustomError                        `pulumi:"customErrorConfigurations"`
	EnableFips                          *bool                                                  `pulumi:"enableFips"`
	EnableHttp2                         *bool                                                  `pulumi:"enableHttp2"`
	FirewallPolicy                      *SubResource                                           `pulumi:"firewallPolicy"`
	FrontendIPConfigurations            []ApplicationGatewayFrontendIPConfiguration            `pulumi:"frontendIPConfigurations"`
	FrontendPorts                       []ApplicationGatewayFrontendPort                       `pulumi:"frontendPorts"`
	GatewayIPConfigurations             []ApplicationGatewayIPConfiguration                    `pulumi:"gatewayIPConfigurations"`
	HttpListeners                       []ApplicationGatewayHttpListener                       `pulumi:"httpListeners"`
	Id                                  *string                                                `pulumi:"id"`
	Identity                            *ManagedServiceIdentity                                `pulumi:"identity"`
	Location                            *string                                                `pulumi:"location"`
	Probes                              []ApplicationGatewayProbe                              `pulumi:"probes"`
	ProvisioningState                   *string                                                `pulumi:"provisioningState"`
	RedirectConfigurations              []ApplicationGatewayRedirectConfiguration              `pulumi:"redirectConfigurations"`
	RequestRoutingRules                 []ApplicationGatewayRequestRoutingRule                 `pulumi:"requestRoutingRules"`
	ResourceGroupName                   string                                                 `pulumi:"resourceGroupName"`
	ResourceGuid                        *string                                                `pulumi:"resourceGuid"`
	RewriteRuleSets                     []ApplicationGatewayRewriteRuleSet                     `pulumi:"rewriteRuleSets"`
	Sku                                 *ApplicationGatewaySku                                 `pulumi:"sku"`
	SslCertificates                     []ApplicationGatewaySslCertificate                     `pulumi:"sslCertificates"`
	SslPolicy                           *ApplicationGatewaySslPolicy                           `pulumi:"sslPolicy"`
	Tags                                map[string]string                                      `pulumi:"tags"`
	TrustedRootCertificates             []ApplicationGatewayTrustedRootCertificate             `pulumi:"trustedRootCertificates"`
	UrlPathMaps                         []ApplicationGatewayUrlPathMap                         `pulumi:"urlPathMaps"`
	WebApplicationFirewallConfiguration *ApplicationGatewayWebApplicationFirewallConfiguration `pulumi:"webApplicationFirewallConfiguration"`
	Zones                               []string                                               `pulumi:"zones"`
}


type ApplicationGatewayArgs struct {
	ApplicationGatewayName              pulumi.StringPtrInput
	AuthenticationCertificates          ApplicationGatewayAuthenticationCertificateArrayInput
	AutoscaleConfiguration              ApplicationGatewayAutoscaleConfigurationPtrInput
	BackendAddressPools                 ApplicationGatewayBackendAddressPoolArrayInput
	BackendHttpSettingsCollection       ApplicationGatewayBackendHttpSettingsArrayInput
	CustomErrorConfigurations           ApplicationGatewayCustomErrorArrayInput
	EnableFips                          pulumi.BoolPtrInput
	EnableHttp2                         pulumi.BoolPtrInput
	FirewallPolicy                      SubResourcePtrInput
	FrontendIPConfigurations            ApplicationGatewayFrontendIPConfigurationArrayInput
	FrontendPorts                       ApplicationGatewayFrontendPortArrayInput
	GatewayIPConfigurations             ApplicationGatewayIPConfigurationArrayInput
	HttpListeners                       ApplicationGatewayHttpListenerArrayInput
	Id                                  pulumi.StringPtrInput
	Identity                            ManagedServiceIdentityPtrInput
	Location                            pulumi.StringPtrInput
	Probes                              ApplicationGatewayProbeArrayInput
	ProvisioningState                   pulumi.StringPtrInput
	RedirectConfigurations              ApplicationGatewayRedirectConfigurationArrayInput
	RequestRoutingRules                 ApplicationGatewayRequestRoutingRuleArrayInput
	ResourceGroupName                   pulumi.StringInput
	ResourceGuid                        pulumi.StringPtrInput
	RewriteRuleSets                     ApplicationGatewayRewriteRuleSetArrayInput
	Sku                                 ApplicationGatewaySkuPtrInput
	SslCertificates                     ApplicationGatewaySslCertificateArrayInput
	SslPolicy                           ApplicationGatewaySslPolicyPtrInput
	Tags                                pulumi.StringMapInput
	TrustedRootCertificates             ApplicationGatewayTrustedRootCertificateArrayInput
	UrlPathMaps                         ApplicationGatewayUrlPathMapArrayInput
	WebApplicationFirewallConfiguration ApplicationGatewayWebApplicationFirewallConfigurationPtrInput
	Zones                               pulumi.StringArrayInput
}

func (ApplicationGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayArgs)(nil)).Elem()
}

type ApplicationGatewayInput interface {
	pulumi.Input

	ToApplicationGatewayOutput() ApplicationGatewayOutput
	ToApplicationGatewayOutputWithContext(ctx context.Context) ApplicationGatewayOutput
}

func (*ApplicationGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGateway)(nil)).Elem()
}

func (i *ApplicationGateway) ToApplicationGatewayOutput() ApplicationGatewayOutput {
	return i.ToApplicationGatewayOutputWithContext(context.Background())
}

func (i *ApplicationGateway) ToApplicationGatewayOutputWithContext(ctx context.Context) ApplicationGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayOutput)
}

type ApplicationGatewayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGateway)(nil)).Elem()
}

func (o ApplicationGatewayOutput) ToApplicationGatewayOutput() ApplicationGatewayOutput {
	return o
}

func (o ApplicationGatewayOutput) ToApplicationGatewayOutputWithContext(ctx context.Context) ApplicationGatewayOutput {
	return o
}

func (o ApplicationGatewayOutput) AuthenticationCertificates() ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
		return v.AuthenticationCertificates
	}).(ApplicationGatewayAuthenticationCertificateResponseArrayOutput)
}

func (o ApplicationGatewayOutput) AutoscaleConfiguration() ApplicationGatewayAutoscaleConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayAutoscaleConfigurationResponsePtrOutput {
		return v.AutoscaleConfiguration
	}).(ApplicationGatewayAutoscaleConfigurationResponsePtrOutput)
}

func (o ApplicationGatewayOutput) BackendAddressPools() ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayBackendAddressPoolResponseArrayOutput {
		return v.BackendAddressPools
	}).(ApplicationGatewayBackendAddressPoolResponseArrayOutput)
}

func (o ApplicationGatewayOutput) BackendHttpSettingsCollection() ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
		return v.BackendHttpSettingsCollection
	}).(ApplicationGatewayBackendHttpSettingsResponseArrayOutput)
}

func (o ApplicationGatewayOutput) CustomErrorConfigurations() ApplicationGatewayCustomErrorResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayCustomErrorResponseArrayOutput {
		return v.CustomErrorConfigurations
	}).(ApplicationGatewayCustomErrorResponseArrayOutput)
}

func (o ApplicationGatewayOutput) EnableFips() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.BoolPtrOutput { return v.EnableFips }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayOutput) EnableHttp2() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.BoolPtrOutput { return v.EnableHttp2 }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayOutput) FirewallPolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) SubResourceResponsePtrOutput { return v.FirewallPolicy }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayOutput) FrontendIPConfigurations() ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
		return v.FrontendIPConfigurations
	}).(ApplicationGatewayFrontendIPConfigurationResponseArrayOutput)
}

func (o ApplicationGatewayOutput) FrontendPorts() ApplicationGatewayFrontendPortResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayFrontendPortResponseArrayOutput { return v.FrontendPorts }).(ApplicationGatewayFrontendPortResponseArrayOutput)
}

func (o ApplicationGatewayOutput) GatewayIPConfigurations() ApplicationGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayIPConfigurationResponseArrayOutput {
		return v.GatewayIPConfigurations
	}).(ApplicationGatewayIPConfigurationResponseArrayOutput)
}

func (o ApplicationGatewayOutput) HttpListeners() ApplicationGatewayHttpListenerResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayHttpListenerResponseArrayOutput { return v.HttpListeners }).(ApplicationGatewayHttpListenerResponseArrayOutput)
}

func (o ApplicationGatewayOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o ApplicationGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationGatewayOutput) OperationalState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringOutput { return v.OperationalState }).(pulumi.StringOutput)
}

func (o ApplicationGatewayOutput) Probes() ApplicationGatewayProbeResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayProbeResponseArrayOutput { return v.Probes }).(ApplicationGatewayProbeResponseArrayOutput)
}

func (o ApplicationGatewayOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayOutput) RedirectConfigurations() ApplicationGatewayRedirectConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayRedirectConfigurationResponseArrayOutput {
		return v.RedirectConfigurations
	}).(ApplicationGatewayRedirectConfigurationResponseArrayOutput)
}

func (o ApplicationGatewayOutput) RequestRoutingRules() ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
		return v.RequestRoutingRules
	}).(ApplicationGatewayRequestRoutingRuleResponseArrayOutput)
}

func (o ApplicationGatewayOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayOutput) RewriteRuleSets() ApplicationGatewayRewriteRuleSetResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayRewriteRuleSetResponseArrayOutput {
		return v.RewriteRuleSets
	}).(ApplicationGatewayRewriteRuleSetResponseArrayOutput)
}

func (o ApplicationGatewayOutput) Sku() ApplicationGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewaySkuResponsePtrOutput { return v.Sku }).(ApplicationGatewaySkuResponsePtrOutput)
}

func (o ApplicationGatewayOutput) SslCertificates() ApplicationGatewaySslCertificateResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewaySslCertificateResponseArrayOutput {
		return v.SslCertificates
	}).(ApplicationGatewaySslCertificateResponseArrayOutput)
}

func (o ApplicationGatewayOutput) SslPolicy() ApplicationGatewaySslPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewaySslPolicyResponsePtrOutput { return v.SslPolicy }).(ApplicationGatewaySslPolicyResponsePtrOutput)
}

func (o ApplicationGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicationGatewayOutput) TrustedRootCertificates() ApplicationGatewayTrustedRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayTrustedRootCertificateResponseArrayOutput {
		return v.TrustedRootCertificates
	}).(ApplicationGatewayTrustedRootCertificateResponseArrayOutput)
}

func (o ApplicationGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApplicationGatewayOutput) UrlPathMaps() ApplicationGatewayUrlPathMapResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayUrlPathMapResponseArrayOutput { return v.UrlPathMaps }).(ApplicationGatewayUrlPathMapResponseArrayOutput)
}

func (o ApplicationGatewayOutput) WebApplicationFirewallConfiguration() ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
		return v.WebApplicationFirewallConfiguration
	}).(ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput)
}

func (o ApplicationGatewayOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGatewayOutput{})
}
