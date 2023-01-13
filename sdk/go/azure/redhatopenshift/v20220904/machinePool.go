


package v20220904

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MachinePool struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput      `pulumi:"name"`
	Resources  pulumi.StringPtrOutput   `pulumi:"resources"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewMachinePool(ctx *pulumi.Context,
	name string, args *MachinePoolArgs, opts ...pulumi.ResourceOption) (*MachinePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	var resource MachinePool
	err := ctx.RegisterResource("azure-native:redhatopenshift/v20220904:MachinePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachinePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachinePoolState, opts ...pulumi.ResourceOption) (*MachinePool, error) {
	var resource MachinePool
	err := ctx.ReadResource("azure-native:redhatopenshift/v20220904:MachinePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type machinePoolState struct {
}

type MachinePoolState struct {
}

func (MachinePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*machinePoolState)(nil)).Elem()
}

type machinePoolArgs struct {
	ChildResourceName *string `pulumi:"childResourceName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
	Resources         *string `pulumi:"resources"`
}


type MachinePoolArgs struct {
	ChildResourceName pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
	Resources         pulumi.StringPtrInput
}

func (MachinePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machinePoolArgs)(nil)).Elem()
}

type MachinePoolInput interface {
	pulumi.Input

	ToMachinePoolOutput() MachinePoolOutput
	ToMachinePoolOutputWithContext(ctx context.Context) MachinePoolOutput
}

func (*MachinePool) ElementType() reflect.Type {
	return reflect.TypeOf((**MachinePool)(nil)).Elem()
}

func (i *MachinePool) ToMachinePoolOutput() MachinePoolOutput {
	return i.ToMachinePoolOutputWithContext(context.Background())
}

func (i *MachinePool) ToMachinePoolOutputWithContext(ctx context.Context) MachinePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachinePoolOutput)
}

type MachinePoolOutput struct{ *pulumi.OutputState }

func (MachinePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachinePool)(nil)).Elem()
}

func (o MachinePoolOutput) ToMachinePoolOutput() MachinePoolOutput {
	return o
}

func (o MachinePoolOutput) ToMachinePoolOutputWithContext(ctx context.Context) MachinePoolOutput {
	return o
}

func (o MachinePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachinePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MachinePoolOutput) Resources() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePool) pulumi.StringPtrOutput { return v.Resources }).(pulumi.StringPtrOutput)
}

func (o MachinePoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MachinePool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MachinePoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MachinePool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachinePoolOutput{})
}
