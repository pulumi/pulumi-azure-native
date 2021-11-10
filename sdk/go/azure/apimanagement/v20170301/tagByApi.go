


package v20170301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TagByApi struct {
	pulumi.CustomResourceState

	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	Name        pulumi.StringOutput `pulumi:"name"`
	Type        pulumi.StringOutput `pulumi:"type"`
}


func NewTagByApi(ctx *pulumi.Context,
	name string, args *TagByApiArgs, opts ...pulumi.ResourceOption) (*TagByApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:TagByApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:TagByApi"),
		},
	})
	opts = append(opts, aliases)
	var resource TagByApi
	err := ctx.RegisterResource("azure-native:apimanagement/v20170301:TagByApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTagByApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagByApiState, opts ...pulumi.ResourceOption) (*TagByApi, error) {
	var resource TagByApi
	err := ctx.ReadResource("azure-native:apimanagement/v20170301:TagByApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tagByApiState struct {
}

type TagByApiState struct {
}

func (TagByApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByApiState)(nil)).Elem()
}

type tagByApiArgs struct {
	ApiId             string  `pulumi:"apiId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	TagId             *string `pulumi:"tagId"`
}


type TagByApiArgs struct {
	ApiId             pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	TagId             pulumi.StringPtrInput
}

func (TagByApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByApiArgs)(nil)).Elem()
}

type TagByApiInput interface {
	pulumi.Input

	ToTagByApiOutput() TagByApiOutput
	ToTagByApiOutputWithContext(ctx context.Context) TagByApiOutput
}

func (*TagByApi) ElementType() reflect.Type {
	return reflect.TypeOf((*TagByApi)(nil))
}

func (i *TagByApi) ToTagByApiOutput() TagByApiOutput {
	return i.ToTagByApiOutputWithContext(context.Background())
}

func (i *TagByApi) ToTagByApiOutputWithContext(ctx context.Context) TagByApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagByApiOutput)
}

type TagByApiOutput struct{ *pulumi.OutputState }

func (TagByApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagByApi)(nil))
}

func (o TagByApiOutput) ToTagByApiOutput() TagByApiOutput {
	return o
}

func (o TagByApiOutput) ToTagByApiOutputWithContext(ctx context.Context) TagByApiOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagByApiOutput{})
}
