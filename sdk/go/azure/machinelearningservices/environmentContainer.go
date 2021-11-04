


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentContainer struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties EnvironmentContainerResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput           `pulumi:"systemData"`
	Type       pulumi.StringOutput                `pulumi:"type"`
}


func NewEnvironmentContainer(ctx *pulumi.Context,
	name string, args *EnvironmentContainerArgs, opts ...pulumi.ResourceOption) (*EnvironmentContainer, error) {
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
			Type: pulumi.String("azure-nextgen:machinelearningservices:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210301preview:EnvironmentContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource EnvironmentContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices:EnvironmentContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnvironmentContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentContainerState, opts ...pulumi.ResourceOption) (*EnvironmentContainer, error) {
	var resource EnvironmentContainer
	err := ctx.ReadResource("azure-native:machinelearningservices:EnvironmentContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type environmentContainerState struct {
}

type EnvironmentContainerState struct {
}

func (EnvironmentContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentContainerState)(nil)).Elem()
}

type environmentContainerArgs struct {
	Name              *string                  `pulumi:"name"`
	Properties        EnvironmentContainerType `pulumi:"properties"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	WorkspaceName     string                   `pulumi:"workspaceName"`
}


type EnvironmentContainerArgs struct {
	Name              pulumi.StringPtrInput
	Properties        EnvironmentContainerTypeInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (EnvironmentContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentContainerArgs)(nil)).Elem()
}

type EnvironmentContainerInput interface {
	pulumi.Input

	ToEnvironmentContainerOutput() EnvironmentContainerOutput
	ToEnvironmentContainerOutputWithContext(ctx context.Context) EnvironmentContainerOutput
}

func (*EnvironmentContainer) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentContainer)(nil))
}

func (i *EnvironmentContainer) ToEnvironmentContainerOutput() EnvironmentContainerOutput {
	return i.ToEnvironmentContainerOutputWithContext(context.Background())
}

func (i *EnvironmentContainer) ToEnvironmentContainerOutputWithContext(ctx context.Context) EnvironmentContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentContainerOutput)
}

type EnvironmentContainerOutput struct{ *pulumi.OutputState }

func (EnvironmentContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentContainer)(nil))
}

func (o EnvironmentContainerOutput) ToEnvironmentContainerOutput() EnvironmentContainerOutput {
	return o
}

func (o EnvironmentContainerOutput) ToEnvironmentContainerOutputWithContext(ctx context.Context) EnvironmentContainerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EnvironmentContainerOutput{})
}
