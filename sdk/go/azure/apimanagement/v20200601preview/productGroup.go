


package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProductGroup struct {
	pulumi.CustomResourceState

	BuiltIn     pulumi.BoolOutput      `pulumi:"builtIn"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName pulumi.StringOutput    `pulumi:"displayName"`
	ExternalId  pulumi.StringPtrOutput `pulumi:"externalId"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewProductGroup(ctx *pulumi.Context,
	name string, args *ProductGroupArgs, opts ...pulumi.ResourceOption) (*ProductGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ProductGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ProductGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ProductGroup
	err := ctx.RegisterResource("azure-native:apimanagement/v20200601preview:ProductGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProductGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductGroupState, opts ...pulumi.ResourceOption) (*ProductGroup, error) {
	var resource ProductGroup
	err := ctx.ReadResource("azure-native:apimanagement/v20200601preview:ProductGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type productGroupState struct {
}

type ProductGroupState struct {
}

func (ProductGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*productGroupState)(nil)).Elem()
}

type productGroupArgs struct {
	GroupId           *string `pulumi:"groupId"`
	ProductId         string  `pulumi:"productId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type ProductGroupArgs struct {
	GroupId           pulumi.StringPtrInput
	ProductId         pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (ProductGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productGroupArgs)(nil)).Elem()
}

type ProductGroupInput interface {
	pulumi.Input

	ToProductGroupOutput() ProductGroupOutput
	ToProductGroupOutputWithContext(ctx context.Context) ProductGroupOutput
}

func (*ProductGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductGroup)(nil)).Elem()
}

func (i *ProductGroup) ToProductGroupOutput() ProductGroupOutput {
	return i.ToProductGroupOutputWithContext(context.Background())
}

func (i *ProductGroup) ToProductGroupOutputWithContext(ctx context.Context) ProductGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductGroupOutput)
}

type ProductGroupOutput struct{ *pulumi.OutputState }

func (ProductGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductGroup)(nil)).Elem()
}

func (o ProductGroupOutput) ToProductGroupOutput() ProductGroupOutput {
	return o
}

func (o ProductGroupOutput) ToProductGroupOutputWithContext(ctx context.Context) ProductGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProductGroupOutput{})
}
