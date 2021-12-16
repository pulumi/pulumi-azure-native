


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContentType struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Schema      pulumi.AnyOutput       `pulumi:"schema"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	Version     pulumi.StringPtrOutput `pulumi:"version"`
}


func NewContentType(ctx *pulumi.Context,
	name string, args *ContentTypeArgs, opts ...pulumi.ResourceOption) (*ContentType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ContentType"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ContentType"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ContentType"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ContentType"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ContentType"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ContentType"),
		},
	})
	opts = append(opts, aliases)
	var resource ContentType
	err := ctx.RegisterResource("azure-native:apimanagement/v20210401preview:ContentType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContentType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentTypeState, opts ...pulumi.ResourceOption) (*ContentType, error) {
	var resource ContentType
	err := ctx.ReadResource("azure-native:apimanagement/v20210401preview:ContentType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type contentTypeState struct {
}

type ContentTypeState struct {
}

func (ContentTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentTypeState)(nil)).Elem()
}

type contentTypeArgs struct {
	ContentTypeId     *string `pulumi:"contentTypeId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type ContentTypeArgs struct {
	ContentTypeId     pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (ContentTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentTypeArgs)(nil)).Elem()
}

type ContentTypeInput interface {
	pulumi.Input

	ToContentTypeOutput() ContentTypeOutput
	ToContentTypeOutputWithContext(ctx context.Context) ContentTypeOutput
}

func (*ContentType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentType)(nil)).Elem()
}

func (i *ContentType) ToContentTypeOutput() ContentTypeOutput {
	return i.ToContentTypeOutputWithContext(context.Background())
}

func (i *ContentType) ToContentTypeOutputWithContext(ctx context.Context) ContentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentTypeOutput)
}

type ContentTypeOutput struct{ *pulumi.OutputState }

func (ContentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentType)(nil)).Elem()
}

func (o ContentTypeOutput) ToContentTypeOutput() ContentTypeOutput {
	return o
}

func (o ContentTypeOutput) ToContentTypeOutputWithContext(ctx context.Context) ContentTypeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ContentTypeOutput{})
}
