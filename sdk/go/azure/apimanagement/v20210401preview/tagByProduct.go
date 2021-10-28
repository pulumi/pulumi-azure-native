


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TagByProduct struct {
	pulumi.CustomResourceState

	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	Name        pulumi.StringOutput `pulumi:"name"`
	Type        pulumi.StringOutput `pulumi:"type"`
}


func NewTagByProduct(ctx *pulumi.Context,
	name string, args *TagByProductArgs, opts ...pulumi.ResourceOption) (*TagByProduct, error) {
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
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210801:TagByProduct"),
		},
	})
	opts = append(opts, aliases)
	var resource TagByProduct
	err := ctx.RegisterResource("azure-native:apimanagement/v20210401preview:TagByProduct", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTagByProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagByProductState, opts ...pulumi.ResourceOption) (*TagByProduct, error) {
	var resource TagByProduct
	err := ctx.ReadResource("azure-native:apimanagement/v20210401preview:TagByProduct", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tagByProductState struct {
}

type TagByProductState struct {
}

func (TagByProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByProductState)(nil)).Elem()
}

type tagByProductArgs struct {
	ProductId         string  `pulumi:"productId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	TagId             *string `pulumi:"tagId"`
}


type TagByProductArgs struct {
	ProductId         pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	TagId             pulumi.StringPtrInput
}

func (TagByProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByProductArgs)(nil)).Elem()
}

type TagByProductInput interface {
	pulumi.Input

	ToTagByProductOutput() TagByProductOutput
	ToTagByProductOutputWithContext(ctx context.Context) TagByProductOutput
}

func (*TagByProduct) ElementType() reflect.Type {
	return reflect.TypeOf((*TagByProduct)(nil))
}

func (i *TagByProduct) ToTagByProductOutput() TagByProductOutput {
	return i.ToTagByProductOutputWithContext(context.Background())
}

func (i *TagByProduct) ToTagByProductOutputWithContext(ctx context.Context) TagByProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagByProductOutput)
}

type TagByProductOutput struct{ *pulumi.OutputState }

func (TagByProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagByProduct)(nil))
}

func (o TagByProductOutput) ToTagByProductOutput() TagByProductOutput {
	return o
}

func (o TagByProductOutput) ToTagByProductOutputWithContext(ctx context.Context) TagByProductOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagByProductInput)(nil)).Elem(), &TagByProduct{})
	pulumi.RegisterOutputType(TagByProductOutput{})
}
