


package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Tag struct {
	pulumi.CustomResourceState

	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	Name        pulumi.StringOutput `pulumi:"name"`
	Type        pulumi.StringOutput `pulumi:"type"`
}


func NewTag(ctx *pulumi.Context,
	name string, args *TagArgs, opts ...pulumi.ResourceOption) (*Tag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Tag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Tag"),
		},
	})
	opts = append(opts, aliases)
	var resource Tag
	err := ctx.RegisterResource("azure-native:apimanagement/v20200601preview:Tag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagState, opts ...pulumi.ResourceOption) (*Tag, error) {
	var resource Tag
	err := ctx.ReadResource("azure-native:apimanagement/v20200601preview:Tag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tagState struct {
}

type TagState struct {
}

func (TagState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagState)(nil)).Elem()
}

type tagArgs struct {
	DisplayName       string  `pulumi:"displayName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	TagId             *string `pulumi:"tagId"`
}


type TagArgs struct {
	DisplayName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	TagId             pulumi.StringPtrInput
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagArgs)(nil)).Elem()
}

type TagInput interface {
	pulumi.Input

	ToTagOutput() TagOutput
	ToTagOutputWithContext(ctx context.Context) TagOutput
}

func (*Tag) ElementType() reflect.Type {
	return reflect.TypeOf((**Tag)(nil)).Elem()
}

func (i *Tag) ToTagOutput() TagOutput {
	return i.ToTagOutputWithContext(context.Background())
}

func (i *Tag) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOutput)
}

type TagOutput struct{ *pulumi.OutputState }

func (TagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tag)(nil)).Elem()
}

func (o TagOutput) ToTagOutput() TagOutput {
	return o
}

func (o TagOutput) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagOutput{})
}
