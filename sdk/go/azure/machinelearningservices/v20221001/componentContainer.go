


package v20221001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComponentContainer struct {
	pulumi.CustomResourceState

	ComponentContainerProperties ComponentContainerResponseOutput `pulumi:"componentContainerProperties"`
	Name                         pulumi.StringOutput              `pulumi:"name"`
	SystemData                   SystemDataResponseOutput         `pulumi:"systemData"`
	Type                         pulumi.StringOutput              `pulumi:"type"`
}


func NewComponentContainer(ctx *pulumi.Context,
	name string, args *ComponentContainerArgs, opts ...pulumi.ResourceOption) (*ComponentContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComponentContainerProperties == nil {
		return nil, errors.New("invalid value for required argument 'ComponentContainerProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.ComponentContainerProperties = args.ComponentContainerProperties.ToComponentContainerTypeOutput().ApplyT(func(v ComponentContainerType) ComponentContainerType { return *v.Defaults() }).(ComponentContainerTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:ComponentContainer"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:ComponentContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource ComponentContainer
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001:ComponentContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetComponentContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentContainerState, opts ...pulumi.ResourceOption) (*ComponentContainer, error) {
	var resource ComponentContainer
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001:ComponentContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type componentContainerState struct {
}

type ComponentContainerState struct {
}

func (ComponentContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentContainerState)(nil)).Elem()
}

type componentContainerArgs struct {
	ComponentContainerProperties ComponentContainerType `pulumi:"componentContainerProperties"`
	Name                         *string                `pulumi:"name"`
	ResourceGroupName            string                 `pulumi:"resourceGroupName"`
	WorkspaceName                string                 `pulumi:"workspaceName"`
}


type ComponentContainerArgs struct {
	ComponentContainerProperties ComponentContainerTypeInput
	Name                         pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	WorkspaceName                pulumi.StringInput
}

func (ComponentContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentContainerArgs)(nil)).Elem()
}

type ComponentContainerInput interface {
	pulumi.Input

	ToComponentContainerOutput() ComponentContainerOutput
	ToComponentContainerOutputWithContext(ctx context.Context) ComponentContainerOutput
}

func (*ComponentContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**ComponentContainer)(nil)).Elem()
}

func (i *ComponentContainer) ToComponentContainerOutput() ComponentContainerOutput {
	return i.ToComponentContainerOutputWithContext(context.Background())
}

func (i *ComponentContainer) ToComponentContainerOutputWithContext(ctx context.Context) ComponentContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentContainerOutput)
}

type ComponentContainerOutput struct{ *pulumi.OutputState }

func (ComponentContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComponentContainer)(nil)).Elem()
}

func (o ComponentContainerOutput) ToComponentContainerOutput() ComponentContainerOutput {
	return o
}

func (o ComponentContainerOutput) ToComponentContainerOutputWithContext(ctx context.Context) ComponentContainerOutput {
	return o
}

func (o ComponentContainerOutput) ComponentContainerProperties() ComponentContainerResponseOutput {
	return o.ApplyT(func(v *ComponentContainer) ComponentContainerResponseOutput { return v.ComponentContainerProperties }).(ComponentContainerResponseOutput)
}

func (o ComponentContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComponentContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComponentContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ComponentContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ComponentContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ComponentContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ComponentContainerOutput{})
}
