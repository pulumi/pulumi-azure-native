


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ApplicationGateway struct {
	pulumi.CustomResourceState

	BackendAddressPools           ApplicationGatewayBackendAddressPoolResponseArrayOutput      `pulumi:"backendAddressPools"`
	BackendHttpSettingsCollection ApplicationGatewayBackendHttpSettingsResponseArrayOutput     `pulumi:"backendHttpSettingsCollection"`
	Etag                          pulumi.StringPtrOutput                                       `pulumi:"etag"`
	FrontendIPConfigurations      ApplicationGatewayFrontendIPConfigurationResponseArrayOutput `pulumi:"frontendIPConfigurations"`
	FrontendPorts                 ApplicationGatewayFrontendPortResponseArrayOutput            `pulumi:"frontendPorts"`
	GatewayIPConfigurations       ApplicationGatewayIPConfigurationResponseArrayOutput         `pulumi:"gatewayIPConfigurations"`
	HttpListeners                 ApplicationGatewayHttpListenerResponseArrayOutput            `pulumi:"httpListeners"`
	Location                      pulumi.StringOutput                                          `pulumi:"location"`
	Name                          pulumi.StringOutput                                          `pulumi:"name"`
	OperationalState              pulumi.StringOutput                                          `pulumi:"operationalState"`
	ProvisioningState             pulumi.StringPtrOutput                                       `pulumi:"provisioningState"`
	RequestRoutingRules           ApplicationGatewayRequestRoutingRuleResponseArrayOutput      `pulumi:"requestRoutingRules"`
	ResourceGuid                  pulumi.StringPtrOutput                                       `pulumi:"resourceGuid"`
	Sku                           ApplicationGatewaySkuResponsePtrOutput                       `pulumi:"sku"`
	SslCertificates               ApplicationGatewaySslCertificateResponseArrayOutput          `pulumi:"sslCertificates"`
	Tags                          pulumi.StringMapOutput                                       `pulumi:"tags"`
	Type                          pulumi.StringOutput                                          `pulumi:"type"`
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
			Type: pulumi.String("azure-native:network/v20181201:ApplicationGateway"),
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
	})
	opts = append(opts, aliases)
	var resource ApplicationGateway
	err := ctx.RegisterResource("azure-native:network/v20150501preview:ApplicationGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGatewayState, opts ...pulumi.ResourceOption) (*ApplicationGateway, error) {
	var resource ApplicationGateway
	err := ctx.ReadResource("azure-native:network/v20150501preview:ApplicationGateway", name, id, state, &resource, opts...)
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
	ApplicationGatewayName        *string                                     `pulumi:"applicationGatewayName"`
	BackendAddressPools           []ApplicationGatewayBackendAddressPool      `pulumi:"backendAddressPools"`
	BackendHttpSettingsCollection []ApplicationGatewayBackendHttpSettings     `pulumi:"backendHttpSettingsCollection"`
	FrontendIPConfigurations      []ApplicationGatewayFrontendIPConfiguration `pulumi:"frontendIPConfigurations"`
	FrontendPorts                 []ApplicationGatewayFrontendPort            `pulumi:"frontendPorts"`
	GatewayIPConfigurations       []ApplicationGatewayIPConfiguration         `pulumi:"gatewayIPConfigurations"`
	HttpListeners                 []ApplicationGatewayHttpListener            `pulumi:"httpListeners"`
	Location                      *string                                     `pulumi:"location"`
	ProvisioningState             *string                                     `pulumi:"provisioningState"`
	RequestRoutingRules           []ApplicationGatewayRequestRoutingRule      `pulumi:"requestRoutingRules"`
	ResourceGroupName             string                                      `pulumi:"resourceGroupName"`
	ResourceGuid                  *string                                     `pulumi:"resourceGuid"`
	Sku                           *ApplicationGatewaySku                      `pulumi:"sku"`
	SslCertificates               []ApplicationGatewaySslCertificate          `pulumi:"sslCertificates"`
	Tags                          map[string]string                           `pulumi:"tags"`
}


type ApplicationGatewayArgs struct {
	ApplicationGatewayName        pulumi.StringPtrInput
	BackendAddressPools           ApplicationGatewayBackendAddressPoolArrayInput
	BackendHttpSettingsCollection ApplicationGatewayBackendHttpSettingsArrayInput
	FrontendIPConfigurations      ApplicationGatewayFrontendIPConfigurationArrayInput
	FrontendPorts                 ApplicationGatewayFrontendPortArrayInput
	GatewayIPConfigurations       ApplicationGatewayIPConfigurationArrayInput
	HttpListeners                 ApplicationGatewayHttpListenerArrayInput
	Location                      pulumi.StringPtrInput
	ProvisioningState             pulumi.StringPtrInput
	RequestRoutingRules           ApplicationGatewayRequestRoutingRuleArrayInput
	ResourceGroupName             pulumi.StringInput
	ResourceGuid                  pulumi.StringPtrInput
	Sku                           ApplicationGatewaySkuPtrInput
	SslCertificates               ApplicationGatewaySslCertificateArrayInput
	Tags                          pulumi.StringMapInput
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

func (o ApplicationGatewayOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
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

func (o ApplicationGatewayOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ApplicationGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationGatewayOutput) OperationalState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringOutput { return v.OperationalState }).(pulumi.StringOutput)
}

func (o ApplicationGatewayOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayOutput) RequestRoutingRules() ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
		return v.RequestRoutingRules
	}).(ApplicationGatewayRequestRoutingRuleResponseArrayOutput)
}

func (o ApplicationGatewayOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayOutput) Sku() ApplicationGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewaySkuResponsePtrOutput { return v.Sku }).(ApplicationGatewaySkuResponsePtrOutput)
}

func (o ApplicationGatewayOutput) SslCertificates() ApplicationGatewaySslCertificateResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewaySslCertificateResponseArrayOutput {
		return v.SslCertificates
	}).(ApplicationGatewaySslCertificateResponseArrayOutput)
}

func (o ApplicationGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicationGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGatewayOutput{})
}
