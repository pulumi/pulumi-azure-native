


package machinelearningservices

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
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:CodeContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:CodeContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource CodeContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices:CodeContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCodeContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodeContainerState, opts ...pulumi.ResourceOption) (*CodeContainer, error) {
	var resource CodeContainer
	err := ctx.ReadResource("azure-native:machinelearningservices:CodeContainer", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**CodeContainer)(nil)).Elem()
}

func (i *CodeContainer) ToCodeContainerOutput() CodeContainerOutput {
	return i.ToCodeContainerOutputWithContext(context.Background())
}

func (i *CodeContainer) ToCodeContainerOutputWithContext(ctx context.Context) CodeContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeContainerOutput)
}

type CodeContainerOutput struct{ *pulumi.OutputState }

func (CodeContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeContainer)(nil)).Elem()
}

func (o CodeContainerOutput) ToCodeContainerOutput() CodeContainerOutput {
	return o
}

func (o CodeContainerOutput) ToCodeContainerOutputWithContext(ctx context.Context) CodeContainerOutput {
	return o
}

func (o CodeContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CodeContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CodeContainerOutput) Properties() CodeContainerResponseOutput {
	return o.ApplyT(func(v *CodeContainer) CodeContainerResponseOutput { return v.Properties }).(CodeContainerResponseOutput)
}

func (o CodeContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CodeContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CodeContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CodeContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CodeContainerOutput{})
}
