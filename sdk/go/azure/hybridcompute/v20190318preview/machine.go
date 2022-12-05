


package v20190318preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Machine struct {
	pulumi.CustomResourceState

	AgentVersion      pulumi.StringOutput            `pulumi:"agentVersion"`
	ClientPublicKey   pulumi.StringPtrOutput         `pulumi:"clientPublicKey"`
	DisplayName       pulumi.StringOutput            `pulumi:"displayName"`
	ErrorDetails      ErrorDetailResponseArrayOutput `pulumi:"errorDetails"`
	LastStatusChange  pulumi.StringOutput            `pulumi:"lastStatusChange"`
	Location          pulumi.StringOutput            `pulumi:"location"`
	MachineFqdn       pulumi.StringOutput            `pulumi:"machineFqdn"`
	Name              pulumi.StringOutput            `pulumi:"name"`
	OsName            pulumi.StringOutput            `pulumi:"osName"`
	OsProfile         OSProfileResponseOutput        `pulumi:"osProfile"`
	OsVersion         pulumi.StringOutput            `pulumi:"osVersion"`
	PhysicalLocation  pulumi.StringPtrOutput         `pulumi:"physicalLocation"`
	PrincipalId       pulumi.StringOutput            `pulumi:"principalId"`
	ProvisioningState pulumi.StringOutput            `pulumi:"provisioningState"`
	Status            pulumi.StringOutput            `pulumi:"status"`
	Tags              pulumi.StringMapOutput         `pulumi:"tags"`
	TenantId          pulumi.StringOutput            `pulumi:"tenantId"`
	Type              pulumi.StringOutput            `pulumi:"type"`
	VmId              pulumi.StringOutput            `pulumi:"vmId"`
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
			Type: pulumi.String("azure-native:hybridcompute/v20210517preview:Machine"),
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
	err := ctx.RegisterResource("azure-native:hybridcompute/v20190318preview:Machine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineState, opts ...pulumi.ResourceOption) (*Machine, error) {
	var resource Machine
	err := ctx.ReadResource("azure-native:hybridcompute/v20190318preview:Machine", name, id, state, &resource, opts...)
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
	ClientPublicKey   *string           `pulumi:"clientPublicKey"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	PhysicalLocation  *string           `pulumi:"physicalLocation"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
}


type MachineArgs struct {
	ClientPublicKey   pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	PhysicalLocation  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
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

func (o MachineOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o MachineOutput) ClientPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.ClientPublicKey }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o MachineOutput) ErrorDetails() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *Machine) ErrorDetailResponseArrayOutput { return v.ErrorDetails }).(ErrorDetailResponseArrayOutput)
}

func (o MachineOutput) LastStatusChange() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.LastStatusChange }).(pulumi.StringOutput)
}

func (o MachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MachineOutput) MachineFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.MachineFqdn }).(pulumi.StringOutput)
}

func (o MachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MachineOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.OsName }).(pulumi.StringOutput)
}

func (o MachineOutput) OsProfile() OSProfileResponseOutput {
	return o.ApplyT(func(v *Machine) OSProfileResponseOutput { return v.OsProfile }).(OSProfileResponseOutput)
}

func (o MachineOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.OsVersion }).(pulumi.StringOutput)
}

func (o MachineOutput) PhysicalLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.PhysicalLocation }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o MachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MachineOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o MachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MachineOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o MachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o MachineOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.VmId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineOutput{})
}
