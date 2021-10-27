


package v20190101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiManagementService struct {
	pulumi.CustomResourceState

	AdditionalLocations         AdditionalLocationResponseArrayOutput           `pulumi:"additionalLocations"`
	Certificates                CertificateConfigurationResponseArrayOutput     `pulumi:"certificates"`
	CreatedAtUtc                pulumi.StringOutput                             `pulumi:"createdAtUtc"`
	CustomProperties            pulumi.StringMapOutput                          `pulumi:"customProperties"`
	DeveloperPortalUrl          pulumi.StringOutput                             `pulumi:"developerPortalUrl"`
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
	ScmUrl                      pulumi.StringOutput                             `pulumi:"scmUrl"`
	Sku                         ApiManagementServiceSkuPropertiesResponseOutput `pulumi:"sku"`
	Tags                        pulumi.StringMapOutput                          `pulumi:"tags"`
	TargetProvisioningState     pulumi.StringOutput                             `pulumi:"targetProvisioningState"`
	Type                        pulumi.StringOutput                             `pulumi:"type"`
	VirtualNetworkConfiguration VirtualNetworkConfigurationResponsePtrOutput    `pulumi:"virtualNetworkConfiguration"`
	VirtualNetworkType          pulumi.StringPtrOutput                          `pulumi:"virtualNetworkType"`
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
	if args.EnableClientCertificate == nil {
		args.EnableClientCertificate = pulumi.BoolPtr(false)
	}
	if args.VirtualNetworkType == nil {
		args.VirtualNetworkType = pulumi.StringPtr("None")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20160707:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20161010:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiManagementService"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210801:ApiManagementService"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiManagementService
	err := ctx.RegisterResource("azure-native:apimanagement/v20190101:ApiManagementService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiManagementService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiManagementServiceState, opts ...pulumi.ResourceOption) (*ApiManagementService, error) {
	var resource ApiManagementService
	err := ctx.ReadResource("azure-native:apimanagement/v20190101:ApiManagementService", name, id, state, &resource, opts...)
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
	Certificates                []CertificateConfiguration        `pulumi:"certificates"`
	CustomProperties            map[string]string                 `pulumi:"customProperties"`
	EnableClientCertificate     *bool                             `pulumi:"enableClientCertificate"`
	HostnameConfigurations      []HostnameConfiguration           `pulumi:"hostnameConfigurations"`
	Identity                    *ApiManagementServiceIdentity     `pulumi:"identity"`
	Location                    *string                           `pulumi:"location"`
	NotificationSenderEmail     *string                           `pulumi:"notificationSenderEmail"`
	PublisherEmail              string                            `pulumi:"publisherEmail"`
	PublisherName               string                            `pulumi:"publisherName"`
	ResourceGroupName           string                            `pulumi:"resourceGroupName"`
	ServiceName                 *string                           `pulumi:"serviceName"`
	Sku                         ApiManagementServiceSkuProperties `pulumi:"sku"`
	Tags                        map[string]string                 `pulumi:"tags"`
	VirtualNetworkConfiguration *VirtualNetworkConfiguration      `pulumi:"virtualNetworkConfiguration"`
	VirtualNetworkType          *string                           `pulumi:"virtualNetworkType"`
}


type ApiManagementServiceArgs struct {
	AdditionalLocations         AdditionalLocationArrayInput
	Certificates                CertificateConfigurationArrayInput
	CustomProperties            pulumi.StringMapInput
	EnableClientCertificate     pulumi.BoolPtrInput
	HostnameConfigurations      HostnameConfigurationArrayInput
	Identity                    ApiManagementServiceIdentityPtrInput
	Location                    pulumi.StringPtrInput
	NotificationSenderEmail     pulumi.StringPtrInput
	PublisherEmail              pulumi.StringInput
	PublisherName               pulumi.StringInput
	ResourceGroupName           pulumi.StringInput
	ServiceName                 pulumi.StringPtrInput
	Sku                         ApiManagementServiceSkuPropertiesInput
	Tags                        pulumi.StringMapInput
	VirtualNetworkConfiguration VirtualNetworkConfigurationPtrInput
	VirtualNetworkType          pulumi.StringPtrInput
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
	return reflect.TypeOf((*ApiManagementService)(nil))
}

func (i *ApiManagementService) ToApiManagementServiceOutput() ApiManagementServiceOutput {
	return i.ToApiManagementServiceOutputWithContext(context.Background())
}

func (i *ApiManagementService) ToApiManagementServiceOutputWithContext(ctx context.Context) ApiManagementServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceOutput)
}

type ApiManagementServiceOutput struct{ *pulumi.OutputState }

func (ApiManagementServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementService)(nil))
}

func (o ApiManagementServiceOutput) ToApiManagementServiceOutput() ApiManagementServiceOutput {
	return o
}

func (o ApiManagementServiceOutput) ToApiManagementServiceOutputWithContext(ctx context.Context) ApiManagementServiceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiManagementServiceInput)(nil)).Elem(), &ApiManagementService{})
	pulumi.RegisterOutputType(ApiManagementServiceOutput{})
}
