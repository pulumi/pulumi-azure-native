


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationStore struct {
	pulumi.CustomResourceState

	CreationDate               pulumi.StringOutput                                   `pulumi:"creationDate"`
	DisableLocalAuth           pulumi.BoolPtrOutput                                  `pulumi:"disableLocalAuth"`
	Encryption                 EncryptionPropertiesResponsePtrOutput                 `pulumi:"encryption"`
	Endpoint                   pulumi.StringOutput                                   `pulumi:"endpoint"`
	Identity                   ResourceIdentityResponsePtrOutput                     `pulumi:"identity"`
	Location                   pulumi.StringOutput                                   `pulumi:"location"`
	Name                       pulumi.StringOutput                                   `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionReferenceResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                                   `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                                `pulumi:"publicNetworkAccess"`
	Sku                        SkuResponseOutput                                     `pulumi:"sku"`
	SystemData                 SystemDataResponseOutput                              `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                                `pulumi:"tags"`
	Type                       pulumi.StringOutput                                   `pulumi:"type"`
}


func NewConfigurationStore(ctx *pulumi.Context,
	name string, args *ConfigurationStoreArgs, opts ...pulumi.ResourceOption) (*ConfigurationStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:appconfiguration/v20210301preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-nextgen:appconfiguration:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20190201preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-nextgen:appconfiguration/v20190201preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20191001:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-nextgen:appconfiguration/v20191001:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20191101preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-nextgen:appconfiguration/v20191101preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20200601:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-nextgen:appconfiguration/v20200601:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20200701preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-nextgen:appconfiguration/v20200701preview:ConfigurationStore"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationStore
	err := ctx.RegisterResource("azure-native:appconfiguration/v20210301preview:ConfigurationStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationStoreState, opts ...pulumi.ResourceOption) (*ConfigurationStore, error) {
	var resource ConfigurationStore
	err := ctx.ReadResource("azure-native:appconfiguration/v20210301preview:ConfigurationStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationStoreState struct {
}

type ConfigurationStoreState struct {
}

func (ConfigurationStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationStoreState)(nil)).Elem()
}

type configurationStoreArgs struct {
	ConfigStoreName     *string               `pulumi:"configStoreName"`
	DisableLocalAuth    *bool                 `pulumi:"disableLocalAuth"`
	Encryption          *EncryptionProperties `pulumi:"encryption"`
	Identity            *ResourceIdentity     `pulumi:"identity"`
	Location            *string               `pulumi:"location"`
	PublicNetworkAccess *string               `pulumi:"publicNetworkAccess"`
	ResourceGroupName   string                `pulumi:"resourceGroupName"`
	Sku                 Sku                   `pulumi:"sku"`
	Tags                map[string]string     `pulumi:"tags"`
}


type ConfigurationStoreArgs struct {
	ConfigStoreName     pulumi.StringPtrInput
	DisableLocalAuth    pulumi.BoolPtrInput
	Encryption          EncryptionPropertiesPtrInput
	Identity            ResourceIdentityPtrInput
	Location            pulumi.StringPtrInput
	PublicNetworkAccess pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Sku                 SkuInput
	Tags                pulumi.StringMapInput
}

func (ConfigurationStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationStoreArgs)(nil)).Elem()
}

type ConfigurationStoreInput interface {
	pulumi.Input

	ToConfigurationStoreOutput() ConfigurationStoreOutput
	ToConfigurationStoreOutputWithContext(ctx context.Context) ConfigurationStoreOutput
}

func (*ConfigurationStore) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationStore)(nil))
}

func (i *ConfigurationStore) ToConfigurationStoreOutput() ConfigurationStoreOutput {
	return i.ToConfigurationStoreOutputWithContext(context.Background())
}

func (i *ConfigurationStore) ToConfigurationStoreOutputWithContext(ctx context.Context) ConfigurationStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationStoreOutput)
}

type ConfigurationStoreOutput struct{ *pulumi.OutputState }

func (ConfigurationStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationStore)(nil))
}

func (o ConfigurationStoreOutput) ToConfigurationStoreOutput() ConfigurationStoreOutput {
	return o
}

func (o ConfigurationStoreOutput) ToConfigurationStoreOutputWithContext(ctx context.Context) ConfigurationStoreOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigurationStoreOutput{})
}
