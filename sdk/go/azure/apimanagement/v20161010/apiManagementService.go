


package v20161010

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiManagementService struct {
	pulumi.CustomResourceState

	AdditionalLocations     AdditionalRegionResponseArrayOutput             `pulumi:"additionalLocations"`
	AddresserEmail          pulumi.StringPtrOutput                          `pulumi:"addresserEmail"`
	CreatedAtUtc            pulumi.StringOutput                             `pulumi:"createdAtUtc"`
	CustomProperties        pulumi.StringMapOutput                          `pulumi:"customProperties"`
	Etag                    pulumi.StringOutput                             `pulumi:"etag"`
	HostnameConfigurations  HostnameConfigurationResponseArrayOutput        `pulumi:"hostnameConfigurations"`
	Location                pulumi.StringOutput                             `pulumi:"location"`
	ManagementApiUrl        pulumi.StringOutput                             `pulumi:"managementApiUrl"`
	Name                    pulumi.StringPtrOutput                          `pulumi:"name"`
	PortalUrl               pulumi.StringOutput                             `pulumi:"portalUrl"`
	ProvisioningState       pulumi.StringOutput                             `pulumi:"provisioningState"`
	PublisherEmail          pulumi.StringOutput                             `pulumi:"publisherEmail"`
	PublisherName           pulumi.StringOutput                             `pulumi:"publisherName"`
	RuntimeUrl              pulumi.StringOutput                             `pulumi:"runtimeUrl"`
	ScmUrl                  pulumi.StringOutput                             `pulumi:"scmUrl"`
	Sku                     ApiManagementServiceSkuPropertiesResponseOutput `pulumi:"sku"`
	StaticIPs               pulumi.StringArrayOutput                        `pulumi:"staticIPs"`
	Tags                    pulumi.StringMapOutput                          `pulumi:"tags"`
	TargetProvisioningState pulumi.StringOutput                             `pulumi:"targetProvisioningState"`
	Type                    pulumi.StringOutput                             `pulumi:"type"`
	VpnType                 pulumi.StringPtrOutput                          `pulumi:"vpnType"`
	Vpnconfiguration        VirtualNetworkConfigurationResponsePtrOutput    `pulumi:"vpnconfiguration"`
}


func NewApiManagementService(ctx *pulumi.Context,
	name string, args *ApiManagementServiceArgs, opts ...pulumi.ResourceOption) (*ApiManagementService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublisherEmail == nil {
		return nil, errors.New("invalid value for required argument 'PublisherEmail'")
	}
	if args.PublisherName == nil {
		return nil, errors.New("invalid value for required argument 'PublisherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	args.Sku = args.Sku.ToApiManagementServiceSkuPropertiesOutput().ApplyT(func(v ApiManagementServiceSkuProperties) ApiManagementServiceSkuProperties { return *v.Defaults() }).(ApiManagementServiceSkuPropertiesOutput)
	if isZero(args.VpnType) {
		args.VpnType = VirtualNetworkType("None")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiManagementService"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiManagementService
	err := ctx.RegisterResource("azure-native:apimanagement/v20161010:ApiManagementService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiManagementService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiManagementServiceState, opts ...pulumi.ResourceOption) (*ApiManagementService, error) {
	var resource ApiManagementService
	err := ctx.ReadResource("azure-native:apimanagement/v20161010:ApiManagementService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiManagementServiceState struct {
}

type ApiManagementServiceState struct {
}

func (ApiManagementServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiManagementServiceState)(nil)).Elem()
}

type apiManagementServiceArgs struct {
	AdditionalLocations    []AdditionalRegion                `pulumi:"additionalLocations"`
	AddresserEmail         *string                           `pulumi:"addresserEmail"`
	CustomProperties       map[string]string                 `pulumi:"customProperties"`
	HostnameConfigurations []HostnameConfiguration           `pulumi:"hostnameConfigurations"`
	Location               *string                           `pulumi:"location"`
	Name                   *string                           `pulumi:"name"`
	PublisherEmail         string                            `pulumi:"publisherEmail"`
	PublisherName          string                            `pulumi:"publisherName"`
	ResourceGroupName      string                            `pulumi:"resourceGroupName"`
	ServiceName            *string                           `pulumi:"serviceName"`
	Sku                    ApiManagementServiceSkuProperties `pulumi:"sku"`
	Tags                   map[string]string                 `pulumi:"tags"`
	VpnType                *VirtualNetworkType               `pulumi:"vpnType"`
	Vpnconfiguration       *VirtualNetworkConfiguration      `pulumi:"vpnconfiguration"`
}


type ApiManagementServiceArgs struct {
	AdditionalLocations    AdditionalRegionArrayInput
	AddresserEmail         pulumi.StringPtrInput
	CustomProperties       pulumi.StringMapInput
	HostnameConfigurations HostnameConfigurationArrayInput
	Location               pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	PublisherEmail         pulumi.StringInput
	PublisherName          pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	ServiceName            pulumi.StringPtrInput
	Sku                    ApiManagementServiceSkuPropertiesInput
	Tags                   pulumi.StringMapInput
	VpnType                VirtualNetworkTypePtrInput
	Vpnconfiguration       VirtualNetworkConfigurationPtrInput
}

func (ApiManagementServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiManagementServiceArgs)(nil)).Elem()
}

type ApiManagementServiceInput interface {
	pulumi.Input

	ToApiManagementServiceOutput() ApiManagementServiceOutput
	ToApiManagementServiceOutputWithContext(ctx context.Context) ApiManagementServiceOutput
}

func (*ApiManagementService) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementService)(nil)).Elem()
}

func (i *ApiManagementService) ToApiManagementServiceOutput() ApiManagementServiceOutput {
	return i.ToApiManagementServiceOutputWithContext(context.Background())
}

func (i *ApiManagementService) ToApiManagementServiceOutputWithContext(ctx context.Context) ApiManagementServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceOutput)
}

type ApiManagementServiceOutput struct{ *pulumi.OutputState }

func (ApiManagementServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementService)(nil)).Elem()
}

func (o ApiManagementServiceOutput) ToApiManagementServiceOutput() ApiManagementServiceOutput {
	return o
}

func (o ApiManagementServiceOutput) ToApiManagementServiceOutputWithContext(ctx context.Context) ApiManagementServiceOutput {
	return o
}

func (o ApiManagementServiceOutput) AdditionalLocations() AdditionalRegionResponseArrayOutput {
	return o.ApplyT(func(v *ApiManagementService) AdditionalRegionResponseArrayOutput { return v.AdditionalLocations }).(AdditionalRegionResponseArrayOutput)
}

func (o ApiManagementServiceOutput) AddresserEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringPtrOutput { return v.AddresserEmail }).(pulumi.StringPtrOutput)
}

func (o ApiManagementServiceOutput) CreatedAtUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.CreatedAtUtc }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) CustomProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringMapOutput { return v.CustomProperties }).(pulumi.StringMapOutput)
}

func (o ApiManagementServiceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) HostnameConfigurations() HostnameConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *ApiManagementService) HostnameConfigurationResponseArrayOutput {
		return v.HostnameConfigurations
	}).(HostnameConfigurationResponseArrayOutput)
}

func (o ApiManagementServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) ManagementApiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.ManagementApiUrl }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiManagementServiceOutput) PortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.PortalUrl }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) PublisherEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.PublisherEmail }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) PublisherName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.PublisherName }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) RuntimeUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.RuntimeUrl }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) ScmUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.ScmUrl }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) Sku() ApiManagementServiceSkuPropertiesResponseOutput {
	return o.ApplyT(func(v *ApiManagementService) ApiManagementServiceSkuPropertiesResponseOutput { return v.Sku }).(ApiManagementServiceSkuPropertiesResponseOutput)
}

func (o ApiManagementServiceOutput) StaticIPs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringArrayOutput { return v.StaticIPs }).(pulumi.StringArrayOutput)
}

func (o ApiManagementServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApiManagementServiceOutput) TargetProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.TargetProvisioningState }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringPtrOutput { return v.VpnType }).(pulumi.StringPtrOutput)
}

func (o ApiManagementServiceOutput) Vpnconfiguration() VirtualNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ApiManagementService) VirtualNetworkConfigurationResponsePtrOutput { return v.Vpnconfiguration }).(VirtualNetworkConfigurationResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiManagementServiceOutput{})
}
