


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiManagementService struct {
	pulumi.CustomResourceState

	AdditionalLocations         AdditionalLocationResponseArrayOutput           `pulumi:"additionalLocations"`
	ApiVersionConstraint        ApiVersionConstraintResponsePtrOutput           `pulumi:"apiVersionConstraint"`
	Certificates                CertificateConfigurationResponseArrayOutput     `pulumi:"certificates"`
	CreatedAtUtc                pulumi.StringOutput                             `pulumi:"createdAtUtc"`
	CustomProperties            pulumi.StringMapOutput                          `pulumi:"customProperties"`
	DeveloperPortalUrl          pulumi.StringOutput                             `pulumi:"developerPortalUrl"`
	DisableGateway              pulumi.BoolPtrOutput                            `pulumi:"disableGateway"`
	EnableClientCertificate     pulumi.BoolPtrOutput                            `pulumi:"enableClientCertificate"`
	Etag                        pulumi.StringOutput                             `pulumi:"etag"`
	GatewayRegionalUrl          pulumi.StringOutput                             `pulumi:"gatewayRegionalUrl"`
	GatewayUrl                  pulumi.StringOutput                             `pulumi:"gatewayUrl"`
	HostnameConfigurations      HostnameConfigurationResponseArrayOutput        `pulumi:"hostnameConfigurations"`
	Identity                    ApiManagementServiceIdentityResponsePtrOutput   `pulumi:"identity"`
	Location                    pulumi.StringOutput                             `pulumi:"location"`
	ManagementApiUrl            pulumi.StringOutput                             `pulumi:"managementApiUrl"`
	Name                        pulumi.StringOutput                             `pulumi:"name"`
	NotificationSenderEmail     pulumi.StringPtrOutput                          `pulumi:"notificationSenderEmail"`
	PortalUrl                   pulumi.StringOutput                             `pulumi:"portalUrl"`
	PrivateIPAddresses          pulumi.StringArrayOutput                        `pulumi:"privateIPAddresses"`
	ProvisioningState           pulumi.StringOutput                             `pulumi:"provisioningState"`
	PublicIPAddresses           pulumi.StringArrayOutput                        `pulumi:"publicIPAddresses"`
	PublisherEmail              pulumi.StringOutput                             `pulumi:"publisherEmail"`
	PublisherName               pulumi.StringOutput                             `pulumi:"publisherName"`
	Restore                     pulumi.BoolPtrOutput                            `pulumi:"restore"`
	ScmUrl                      pulumi.StringOutput                             `pulumi:"scmUrl"`
	Sku                         ApiManagementServiceSkuPropertiesResponseOutput `pulumi:"sku"`
	Tags                        pulumi.StringMapOutput                          `pulumi:"tags"`
	TargetProvisioningState     pulumi.StringOutput                             `pulumi:"targetProvisioningState"`
	Type                        pulumi.StringOutput                             `pulumi:"type"`
	VirtualNetworkConfiguration VirtualNetworkConfigurationResponsePtrOutput    `pulumi:"virtualNetworkConfiguration"`
	VirtualNetworkType          pulumi.StringPtrOutput                          `pulumi:"virtualNetworkType"`
	Zones                       pulumi.StringArrayOutput                        `pulumi:"zones"`
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
	if isZero(args.DisableGateway) {
		args.DisableGateway = pulumi.BoolPtr(false)
	}
	if isZero(args.EnableClientCertificate) {
		args.EnableClientCertificate = pulumi.BoolPtr(false)
	}
	if isZero(args.Restore) {
		args.Restore = pulumi.BoolPtr(false)
	}
	if isZero(args.VirtualNetworkType) {
		args.VirtualNetworkType = pulumi.StringPtr("None")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:ApiManagementService"),
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
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ApiManagementService"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiManagementService
	err := ctx.RegisterResource("azure-native:apimanagement/v20201201:ApiManagementService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiManagementService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiManagementServiceState, opts ...pulumi.ResourceOption) (*ApiManagementService, error) {
	var resource ApiManagementService
	err := ctx.ReadResource("azure-native:apimanagement/v20201201:ApiManagementService", name, id, state, &resource, opts...)
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
	AdditionalLocations         []AdditionalLocation              `pulumi:"additionalLocations"`
	ApiVersionConstraint        *ApiVersionConstraint             `pulumi:"apiVersionConstraint"`
	Certificates                []CertificateConfiguration        `pulumi:"certificates"`
	CustomProperties            map[string]string                 `pulumi:"customProperties"`
	DisableGateway              *bool                             `pulumi:"disableGateway"`
	EnableClientCertificate     *bool                             `pulumi:"enableClientCertificate"`
	HostnameConfigurations      []HostnameConfiguration           `pulumi:"hostnameConfigurations"`
	Identity                    *ApiManagementServiceIdentity     `pulumi:"identity"`
	Location                    *string                           `pulumi:"location"`
	NotificationSenderEmail     *string                           `pulumi:"notificationSenderEmail"`
	PublisherEmail              string                            `pulumi:"publisherEmail"`
	PublisherName               string                            `pulumi:"publisherName"`
	ResourceGroupName           string                            `pulumi:"resourceGroupName"`
	Restore                     *bool                             `pulumi:"restore"`
	ServiceName                 *string                           `pulumi:"serviceName"`
	Sku                         ApiManagementServiceSkuProperties `pulumi:"sku"`
	Tags                        map[string]string                 `pulumi:"tags"`
	VirtualNetworkConfiguration *VirtualNetworkConfiguration      `pulumi:"virtualNetworkConfiguration"`
	VirtualNetworkType          *string                           `pulumi:"virtualNetworkType"`
	Zones                       []string                          `pulumi:"zones"`
}


type ApiManagementServiceArgs struct {
	AdditionalLocations         AdditionalLocationArrayInput
	ApiVersionConstraint        ApiVersionConstraintPtrInput
	Certificates                CertificateConfigurationArrayInput
	CustomProperties            pulumi.StringMapInput
	DisableGateway              pulumi.BoolPtrInput
	EnableClientCertificate     pulumi.BoolPtrInput
	HostnameConfigurations      HostnameConfigurationArrayInput
	Identity                    ApiManagementServiceIdentityPtrInput
	Location                    pulumi.StringPtrInput
	NotificationSenderEmail     pulumi.StringPtrInput
	PublisherEmail              pulumi.StringInput
	PublisherName               pulumi.StringInput
	ResourceGroupName           pulumi.StringInput
	Restore                     pulumi.BoolPtrInput
	ServiceName                 pulumi.StringPtrInput
	Sku                         ApiManagementServiceSkuPropertiesInput
	Tags                        pulumi.StringMapInput
	VirtualNetworkConfiguration VirtualNetworkConfigurationPtrInput
	VirtualNetworkType          pulumi.StringPtrInput
	Zones                       pulumi.StringArrayInput
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

func (o ApiManagementServiceOutput) AdditionalLocations() AdditionalLocationResponseArrayOutput {
	return o.ApplyT(func(v *ApiManagementService) AdditionalLocationResponseArrayOutput { return v.AdditionalLocations }).(AdditionalLocationResponseArrayOutput)
}

func (o ApiManagementServiceOutput) ApiVersionConstraint() ApiVersionConstraintResponsePtrOutput {
	return o.ApplyT(func(v *ApiManagementService) ApiVersionConstraintResponsePtrOutput { return v.ApiVersionConstraint }).(ApiVersionConstraintResponsePtrOutput)
}

func (o ApiManagementServiceOutput) Certificates() CertificateConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *ApiManagementService) CertificateConfigurationResponseArrayOutput { return v.Certificates }).(CertificateConfigurationResponseArrayOutput)
}

func (o ApiManagementServiceOutput) CreatedAtUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.CreatedAtUtc }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) CustomProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringMapOutput { return v.CustomProperties }).(pulumi.StringMapOutput)
}

func (o ApiManagementServiceOutput) DeveloperPortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.DeveloperPortalUrl }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) DisableGateway() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.BoolPtrOutput { return v.DisableGateway }).(pulumi.BoolPtrOutput)
}

func (o ApiManagementServiceOutput) EnableClientCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.BoolPtrOutput { return v.EnableClientCertificate }).(pulumi.BoolPtrOutput)
}

func (o ApiManagementServiceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) GatewayRegionalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.GatewayRegionalUrl }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) GatewayUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.GatewayUrl }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) HostnameConfigurations() HostnameConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *ApiManagementService) HostnameConfigurationResponseArrayOutput {
		return v.HostnameConfigurations
	}).(HostnameConfigurationResponseArrayOutput)
}

func (o ApiManagementServiceOutput) Identity() ApiManagementServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ApiManagementService) ApiManagementServiceIdentityResponsePtrOutput { return v.Identity }).(ApiManagementServiceIdentityResponsePtrOutput)
}

func (o ApiManagementServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) ManagementApiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.ManagementApiUrl }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) NotificationSenderEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringPtrOutput { return v.NotificationSenderEmail }).(pulumi.StringPtrOutput)
}

func (o ApiManagementServiceOutput) PortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.PortalUrl }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) PrivateIPAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringArrayOutput { return v.PrivateIPAddresses }).(pulumi.StringArrayOutput)
}

func (o ApiManagementServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) PublicIPAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringArrayOutput { return v.PublicIPAddresses }).(pulumi.StringArrayOutput)
}

func (o ApiManagementServiceOutput) PublisherEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.PublisherEmail }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) PublisherName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.PublisherName }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) Restore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.BoolPtrOutput { return v.Restore }).(pulumi.BoolPtrOutput)
}

func (o ApiManagementServiceOutput) ScmUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringOutput { return v.ScmUrl }).(pulumi.StringOutput)
}

func (o ApiManagementServiceOutput) Sku() ApiManagementServiceSkuPropertiesResponseOutput {
	return o.ApplyT(func(v *ApiManagementService) ApiManagementServiceSkuPropertiesResponseOutput { return v.Sku }).(ApiManagementServiceSkuPropertiesResponseOutput)
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

func (o ApiManagementServiceOutput) VirtualNetworkConfiguration() VirtualNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ApiManagementService) VirtualNetworkConfigurationResponsePtrOutput {
		return v.VirtualNetworkConfiguration
	}).(VirtualNetworkConfigurationResponsePtrOutput)
}

func (o ApiManagementServiceOutput) VirtualNetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringPtrOutput { return v.VirtualNetworkType }).(pulumi.StringPtrOutput)
}

func (o ApiManagementServiceOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiManagementService) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiManagementServiceOutput{})
}
