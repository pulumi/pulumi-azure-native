


package v20170301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Registry struct {
	pulumi.CustomResourceState

	AdminUserEnabled  pulumi.BoolPtrOutput                      `pulumi:"adminUserEnabled"`
	CreationDate      pulumi.StringOutput                       `pulumi:"creationDate"`
	Location          pulumi.StringOutput                       `pulumi:"location"`
	LoginServer       pulumi.StringOutput                       `pulumi:"loginServer"`
	Name              pulumi.StringOutput                       `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                       `pulumi:"provisioningState"`
	Sku               SkuResponseOutput                         `pulumi:"sku"`
	StorageAccount    StorageAccountPropertiesResponsePtrOutput `pulumi:"storageAccount"`
	Tags              pulumi.StringMapOutput                    `pulumi:"tags"`
	Type              pulumi.StringOutput                       `pulumi:"type"`
}


func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.StorageAccount == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccount'")
	}
	if args.AdminUserEnabled == nil {
		args.AdminUserEnabled = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20170301:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry:Registry"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20160627preview:Registry"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20160627preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20170601preview:Registry"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20170601preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20171001:Registry"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20171001:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190501:Registry"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20190501:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20191201preview:Registry"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20191201preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:Registry"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20201101preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:Registry"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20210601preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:Registry"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20210801preview:Registry"),
		},
	})
	opts = append(opts, aliases)
	var resource Registry
	err := ctx.RegisterResource("azure-native:containerregistry/v20170301:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("azure-native:containerregistry/v20170301:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registryState struct {
}

type RegistryState struct {
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	AdminUserEnabled  *bool                    `pulumi:"adminUserEnabled"`
	Location          *string                  `pulumi:"location"`
	RegistryName      *string                  `pulumi:"registryName"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	Sku               Sku                      `pulumi:"sku"`
	StorageAccount    StorageAccountParameters `pulumi:"storageAccount"`
	Tags              map[string]string        `pulumi:"tags"`
}


type RegistryArgs struct {
	AdminUserEnabled  pulumi.BoolPtrInput
	Location          pulumi.StringPtrInput
	RegistryName      pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuInput
	StorageAccount    StorageAccountParametersInput
	Tags              pulumi.StringMapInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}

type RegistryInput interface {
	pulumi.Input

	ToRegistryOutput() RegistryOutput
	ToRegistryOutputWithContext(ctx context.Context) RegistryOutput
}

func (*Registry) ElementType() reflect.Type {
	return reflect.TypeOf((*Registry)(nil))
}

func (i *Registry) ToRegistryOutput() RegistryOutput {
	return i.ToRegistryOutputWithContext(context.Background())
}

func (i *Registry) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryOutput)
}

type RegistryOutput struct{ *pulumi.OutputState }

func (RegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Registry)(nil))
}

func (o RegistryOutput) ToRegistryOutput() RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryInput)(nil)).Elem(), &Registry{})
	pulumi.RegisterOutputType(RegistryOutput{})
}
