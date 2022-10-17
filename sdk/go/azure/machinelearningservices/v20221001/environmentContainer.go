


package v20221001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentContainer struct {
	pulumi.CustomResourceState

	EnvironmentContainerProperties EnvironmentContainerResponseOutput `pulumi:"environmentContainerProperties"`
	Name                           pulumi.StringOutput                `pulumi:"name"`
	SystemData                     SystemDataResponseOutput           `pulumi:"systemData"`
	Type                           pulumi.StringOutput                `pulumi:"type"`
}


func NewEnvironmentContainer(ctx *pulumi.Context,
	name string, args *EnvironmentContainerArgs, opts ...pulumi.ResourceOption) (*EnvironmentContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentContainerProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.EnvironmentContainerProperties = args.EnvironmentContainerProperties.ToEnvironmentContainerTypeOutput().ApplyT(func(v EnvironmentContainerType) EnvironmentContainerType { return *v.Defaults() }).(EnvironmentContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:EnvironmentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:EnvironmentContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource EnvironmentContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001:EnvironmentContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnvironmentContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentContainerState, opts ...pulumi.ResourceOption) (*EnvironmentContainer, error) {
	var resource EnvironmentContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001:EnvironmentContainer", name, id, state, &resource, opts...)
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
	EnvironmentContainerProperties EnvironmentContainerType `pulumi:"environmentContainerProperties"`
	Name                           *string                  `pulumi:"name"`
	ResourceGroupName              string                   `pulumi:"resourceGroupName"`
	WorkspaceName                  string                   `pulumi:"workspaceName"`
}


type EnvironmentContainerArgs struct {
	EnvironmentContainerProperties EnvironmentContainerTypeInput
	Name                           pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	WorkspaceName                  pulumi.StringInput
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
	return reflect.TypeOf((**EnvironmentContainer)(nil)).Elem()
}

func (i *EnvironmentContainer) ToEnvironmentContainerOutput() EnvironmentContainerOutput {
	return i.ToEnvironmentContainerOutputWithContext(context.Background())
}

func (i *EnvironmentContainer) ToEnvironmentContainerOutputWithContext(ctx context.Context) EnvironmentContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentContainerOutput)
}

type EnvironmentContainerOutput struct{ *pulumi.OutputState }

func (EnvironmentContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentContainer)(nil)).Elem()
}

func (o EnvironmentContainerOutput) ToEnvironmentContainerOutput() EnvironmentContainerOutput {
	return o
}

func (o EnvironmentContainerOutput) ToEnvironmentContainerOutputWithContext(ctx context.Context) EnvironmentContainerOutput {
	return o
}

func (o EnvironmentContainerOutput) EnvironmentContainerProperties() EnvironmentContainerResponseOutput {
	return o.ApplyT(func(v *EnvironmentContainer) EnvironmentContainerResponseOutput {
		return v.EnvironmentContainerProperties
	}).(EnvironmentContainerResponseOutput)
}

func (o EnvironmentContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EnvironmentContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EnvironmentContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EnvironmentContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentContainerOutput{})
}
