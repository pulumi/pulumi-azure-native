


package media

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Transform struct {
	pulumi.CustomResourceState

	Created      pulumi.StringOutput                `pulumi:"created"`
	Description  pulumi.StringPtrOutput             `pulumi:"description"`
	LastModified pulumi.StringOutput                `pulumi:"lastModified"`
	Name         pulumi.StringOutput                `pulumi:"name"`
	Outputs      TransformOutputResponseArrayOutput `pulumi:"outputs"`
	SystemData   SystemDataResponseOutput           `pulumi:"systemData"`
	Type         pulumi.StringOutput                `pulumi:"type"`
}


func NewTransform(ctx *pulumi.Context,
	name string, args *TransformArgs, opts ...pulumi.ResourceOption) (*Transform, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Outputs == nil {
		return nil, errors.New("invalid value for required argument 'Outputs'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media/v20180330preview:Transform"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:Transform"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:Transform"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:Transform"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:Transform"),
		},
	})
	opts = append(opts, aliases)
	var resource Transform
	err := ctx.RegisterResource("azure-native:media:Transform", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTransform(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransformState, opts ...pulumi.ResourceOption) (*Transform, error) {
	var resource Transform
	err := ctx.ReadResource("azure-native:media:Transform", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type transformState struct {
}

type TransformState struct {
}

func (TransformState) ElementType() reflect.Type {
	return reflect.TypeOf((*transformState)(nil)).Elem()
}

type transformArgs struct {
	AccountName       string                `pulumi:"accountName"`
	Description       *string               `pulumi:"description"`
	Outputs           []TransformOutputType `pulumi:"outputs"`
	ResourceGroupName string                `pulumi:"resourceGroupName"`
	TransformName     *string               `pulumi:"transformName"`
}


type TransformArgs struct {
	AccountName       pulumi.StringInput
	Description       pulumi.StringPtrInput
	Outputs           TransformOutputTypeArrayInput
	ResourceGroupName pulumi.StringInput
	TransformName     pulumi.StringPtrInput
}

func (TransformArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transformArgs)(nil)).Elem()
}

type TransformInput interface {
	pulumi.Input

	ToTransformOutput() TransformOutput
	ToTransformOutputWithContext(ctx context.Context) TransformOutput
}

func (*Transform) ElementType() reflect.Type {
	return reflect.TypeOf((**Transform)(nil)).Elem()
}

func (i *Transform) ToTransformOutput() TransformOutput {
	return i.ToTransformOutputWithContext(context.Background())
}

func (i *Transform) ToTransformOutputWithContext(ctx context.Context) TransformOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformOutput)
}

type TransformOutput struct{ *pulumi.OutputState }

func (TransformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Transform)(nil)).Elem()
}

func (o TransformOutput) ToTransformOutput() TransformOutput {
	return o
}

func (o TransformOutput) ToTransformOutputWithContext(ctx context.Context) TransformOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TransformOutput{})
}
