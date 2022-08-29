


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContentItem struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput `pulumi:"name"`
	Properties pulumi.AnyOutput    `pulumi:"properties"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewContentItem(ctx *pulumi.Context,
	name string, args *ContentItemArgs, opts ...pulumi.ResourceOption) (*ContentItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentTypeId == nil {
		return nil, errors.New("invalid value for required argument 'ContentTypeId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ContentItem"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ContentItem"),
		},
	})
	opts = append(opts, aliases)
	var resource ContentItem
	err := ctx.RegisterResource("azure-native:apimanagement/v20201201:ContentItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContentItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentItemState, opts ...pulumi.ResourceOption) (*ContentItem, error) {
	var resource ContentItem
	err := ctx.ReadResource("azure-native:apimanagement/v20201201:ContentItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type contentItemState struct {
}

type ContentItemState struct {
}

func (ContentItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentItemState)(nil)).Elem()
}

type contentItemArgs struct {
	ContentItemId     *string `pulumi:"contentItemId"`
	ContentTypeId     string  `pulumi:"contentTypeId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type ContentItemArgs struct {
	ContentItemId     pulumi.StringPtrInput
	ContentTypeId     pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (ContentItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentItemArgs)(nil)).Elem()
}

type ContentItemInput interface {
	pulumi.Input

	ToContentItemOutput() ContentItemOutput
	ToContentItemOutputWithContext(ctx context.Context) ContentItemOutput
}

func (*ContentItem) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentItem)(nil)).Elem()
}

func (i *ContentItem) ToContentItemOutput() ContentItemOutput {
	return i.ToContentItemOutputWithContext(context.Background())
}

func (i *ContentItem) ToContentItemOutputWithContext(ctx context.Context) ContentItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentItemOutput)
}

type ContentItemOutput struct{ *pulumi.OutputState }

func (ContentItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentItem)(nil)).Elem()
}

func (o ContentItemOutput) ToContentItemOutput() ContentItemOutput {
	return o
}

func (o ContentItemOutput) ToContentItemOutputWithContext(ctx context.Context) ContentItemOutput {
	return o
}

func (o ContentItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ContentItemOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ContentItem) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o ContentItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContentItemOutput{})
}
