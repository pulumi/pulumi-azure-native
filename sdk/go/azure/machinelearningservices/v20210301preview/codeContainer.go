


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CodeContainer struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput         `pulumi:"name"`
	Properties CodeContainerResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput    `pulumi:"systemData"`
	Type       pulumi.StringOutput         `pulumi:"type"`
}


func NewCodeContainer(ctx *pulumi.Context,
	name string, args *CodeContainerArgs, opts ...pulumi.ResourceOption) (*CodeContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210301preview:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices:CodeContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource CodeContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210301preview:CodeContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCodeContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodeContainerState, opts ...pulumi.ResourceOption) (*CodeContainer, error) {
	var resource CodeContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210301preview:CodeContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type codeContainerState struct {
}

type CodeContainerState struct {
}

func (CodeContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*codeContainerState)(nil)).Elem()
}

type codeContainerArgs struct {
	Name              *string           `pulumi:"name"`
	Properties        CodeContainerType `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	WorkspaceName     string            `pulumi:"workspaceName"`
}


type CodeContainerArgs struct {
	Name              pulumi.StringPtrInput
	Properties        CodeContainerTypeInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (CodeContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*codeContainerArgs)(nil)).Elem()
}

type CodeContainerInput interface {
	pulumi.Input

	ToCodeContainerOutput() CodeContainerOutput
	ToCodeContainerOutputWithContext(ctx context.Context) CodeContainerOutput
}

func (*CodeContainer) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeContainer)(nil))
}

func (i *CodeContainer) ToCodeContainerOutput() CodeContainerOutput {
	return i.ToCodeContainerOutputWithContext(context.Background())
}

func (i *CodeContainer) ToCodeContainerOutputWithContext(ctx context.Context) CodeContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeContainerOutput)
}

type CodeContainerOutput struct{ *pulumi.OutputState }

func (CodeContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeContainer)(nil))
}

func (o CodeContainerOutput) ToCodeContainerOutput() CodeContainerOutput {
	return o
}

func (o CodeContainerOutput) ToCodeContainerOutputWithContext(ctx context.Context) CodeContainerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CodeContainerOutput{})
}
