


package v20210517preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Machine struct {
	pulumi.CustomResourceState

	Identity   IdentityResponsePtrOutput       `pulumi:"identity"`
	Location   pulumi.StringOutput             `pulumi:"location"`
	Name       pulumi.StringOutput             `pulumi:"name"`
	Properties MachinePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput        `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput          `pulumi:"tags"`
	Type       pulumi.StringOutput             `pulumi:"type"`
}


func NewMachine(ctx *pulumi.Context,
	name string, args *MachineArgs, opts ...pulumi.ResourceOption) (*Machine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcompute:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20190318preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20190802preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20191212:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200730preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200802:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200815preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210128preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210325preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210422preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210520:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210610preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20211210preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220310:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220510preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220811preview:Machine"),
		},
	})
	opts = append(opts, aliases)
	var resource Machine
	err := ctx.RegisterResource("azure-native:hybridcompute/v20210517preview:Machine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineState, opts ...pulumi.ResourceOption) (*Machine, error) {
	var resource Machine
	err := ctx.ReadResource("azure-native:hybridcompute/v20210517preview:Machine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type machineState struct {
}

type MachineState struct {
}

func (MachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineState)(nil)).Elem()
}

type machineArgs struct {
	Identity          *Identity          `pulumi:"identity"`
	Location          *string            `pulumi:"location"`
	MachineName       *string            `pulumi:"machineName"`
	Properties        *MachineProperties `pulumi:"properties"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	Tags              map[string]string  `pulumi:"tags"`
}


type MachineArgs struct {
	Identity          IdentityPtrInput
	Location          pulumi.StringPtrInput
	MachineName       pulumi.StringPtrInput
	Properties        MachinePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (MachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineArgs)(nil)).Elem()
}

type MachineInput interface {
	pulumi.Input

	ToMachineOutput() MachineOutput
	ToMachineOutputWithContext(ctx context.Context) MachineOutput
}

func (*Machine) ElementType() reflect.Type {
	return reflect.TypeOf((**Machine)(nil)).Elem()
}

func (i *Machine) ToMachineOutput() MachineOutput {
	return i.ToMachineOutputWithContext(context.Background())
}

func (i *Machine) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineOutput)
}

type MachineOutput struct{ *pulumi.OutputState }

func (MachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Machine)(nil)).Elem()
}

func (o MachineOutput) ToMachineOutput() MachineOutput {
	return o
}

func (o MachineOutput) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return o
}

func (o MachineOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Machine) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o MachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MachineOutput) Properties() MachinePropertiesResponseOutput {
	return o.ApplyT(func(v *Machine) MachinePropertiesResponseOutput { return v.Properties }).(MachinePropertiesResponseOutput)
}

func (o MachineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Machine) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineOutput{})
}
