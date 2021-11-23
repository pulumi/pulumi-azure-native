


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TagByOperation struct {
	pulumi.CustomResourceState

	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	Name        pulumi.StringOutput `pulumi:"name"`
	Type        pulumi.StringOutput `pulumi:"type"`
}


func NewTagByOperation(ctx *pulumi.Context,
	name string, args *TagByOperationArgs, opts ...pulumi.ResourceOption) (*TagByOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.OperationId == nil {
		return nil, errors.New("invalid value for required argument 'OperationId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:TagByOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:TagByOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:TagByOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:TagByOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:TagByOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:TagByOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:TagByOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:TagByOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:TagByOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:TagByOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:TagByOperation"),
		},
	})
	opts = append(opts, aliases)
	var resource TagByOperation
	err := ctx.RegisterResource("azure-native:apimanagement/v20201201:TagByOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTagByOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagByOperationState, opts ...pulumi.ResourceOption) (*TagByOperation, error) {
	var resource TagByOperation
	err := ctx.ReadResource("azure-native:apimanagement/v20201201:TagByOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tagByOperationState struct {
}

type TagByOperationState struct {
}

func (TagByOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByOperationState)(nil)).Elem()
}

type tagByOperationArgs struct {
	ApiId             string  `pulumi:"apiId"`
	OperationId       string  `pulumi:"operationId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	TagId             *string `pulumi:"tagId"`
}


type TagByOperationArgs struct {
	ApiId             pulumi.StringInput
	OperationId       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	TagId             pulumi.StringPtrInput
}

func (TagByOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByOperationArgs)(nil)).Elem()
}

type TagByOperationInput interface {
	pulumi.Input

	ToTagByOperationOutput() TagByOperationOutput
	ToTagByOperationOutputWithContext(ctx context.Context) TagByOperationOutput
}

func (*TagByOperation) ElementType() reflect.Type {
	return reflect.TypeOf((*TagByOperation)(nil))
}

func (i *TagByOperation) ToTagByOperationOutput() TagByOperationOutput {
	return i.ToTagByOperationOutputWithContext(context.Background())
}

func (i *TagByOperation) ToTagByOperationOutputWithContext(ctx context.Context) TagByOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagByOperationOutput)
}

type TagByOperationOutput struct{ *pulumi.OutputState }

func (TagByOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagByOperation)(nil))
}

func (o TagByOperationOutput) ToTagByOperationOutput() TagByOperationOutput {
	return o
}

func (o TagByOperationOutput) ToTagByOperationOutputWithContext(ctx context.Context) TagByOperationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagByOperationOutput{})
}
