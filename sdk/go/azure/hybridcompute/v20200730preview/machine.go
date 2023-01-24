


package v20200730preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Machine struct {
	pulumi.CustomResourceState

	AdFqdn            pulumi.StringOutput                             `pulumi:"adFqdn"`
	AgentVersion      pulumi.StringOutput                             `pulumi:"agentVersion"`
	ClientPublicKey   pulumi.StringPtrOutput                          `pulumi:"clientPublicKey"`
	DisplayName       pulumi.StringOutput                             `pulumi:"displayName"`
	DnsFqdn           pulumi.StringOutput                             `pulumi:"dnsFqdn"`
	DomainName        pulumi.StringOutput                             `pulumi:"domainName"`
	ErrorDetails      ErrorDetailResponseArrayOutput                  `pulumi:"errorDetails"`
	Extensions        MachineExtensionInstanceViewResponseArrayOutput `pulumi:"extensions"`
	Identity          MachineResponseIdentityPtrOutput                `pulumi:"identity"`
	LastStatusChange  pulumi.StringOutput                             `pulumi:"lastStatusChange"`
	Location          pulumi.StringOutput                             `pulumi:"location"`
	LocationData      LocationDataResponsePtrOutput                   `pulumi:"locationData"`
	MachineFqdn       pulumi.StringOutput                             `pulumi:"machineFqdn"`
	Name              pulumi.StringOutput                             `pulumi:"name"`
	OsName            pulumi.StringOutput                             `pulumi:"osName"`
	OsProfile         MachinePropertiesResponseOsProfilePtrOutput     `pulumi:"osProfile"`
	OsSku             pulumi.StringOutput                             `pulumi:"osSku"`
	OsVersion         pulumi.StringOutput                             `pulumi:"osVersion"`
	ProvisioningState pulumi.StringOutput                             `pulumi:"provisioningState"`
	Status            pulumi.StringOutput                             `pulumi:"status"`
	Tags              pulumi.StringMapOutput                          `pulumi:"tags"`
	Type              pulumi.StringOutput                             `pulumi:"type"`
	VmId              pulumi.StringPtrOutput                          `pulumi:"vmId"`
	VmUuid            pulumi.StringOutput                             `pulumi:"vmUuid"`
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
		{
			Type: pulumi.String("azure-native:hybridcompute/v20221110:Machine"),
		},
	})
	opts = append(opts, aliases)
	var resource Machine
	err := ctx.RegisterResource("azure-native:hybridcompute/v20200730preview:Machine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineState, opts ...pulumi.ResourceOption) (*Machine, error) {
	var resource Machine
	err := ctx.ReadResource("azure-native:hybridcompute/v20200730preview:Machine", name, id, state, &resource, opts...)
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
	ClientPublicKey   *string                        `pulumi:"clientPublicKey"`
	Extensions        []MachineExtensionInstanceView `pulumi:"extensions"`
	Identity          *MachineIdentity               `pulumi:"identity"`
	Location          *string                        `pulumi:"location"`
	LocationData      *LocationData                  `pulumi:"locationData"`
	Name              *string                        `pulumi:"name"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	Tags              map[string]string              `pulumi:"tags"`
	VmId              *string                        `pulumi:"vmId"`
}


type MachineArgs struct {
	ClientPublicKey   pulumi.StringPtrInput
	Extensions        MachineExtensionInstanceViewArrayInput
	Identity          MachineIdentityPtrInput
	Location          pulumi.StringPtrInput
	LocationData      LocationDataPtrInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VmId              pulumi.StringPtrInput
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

func (o MachineOutput) AdFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.AdFqdn }).(pulumi.StringOutput)
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

func (o MachineOutput) DnsFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.DnsFqdn }).(pulumi.StringOutput)
}

func (o MachineOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o MachineOutput) ErrorDetails() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *Machine) ErrorDetailResponseArrayOutput { return v.ErrorDetails }).(ErrorDetailResponseArrayOutput)
}

func (o MachineOutput) Extensions() MachineExtensionInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v *Machine) MachineExtensionInstanceViewResponseArrayOutput { return v.Extensions }).(MachineExtensionInstanceViewResponseArrayOutput)
}

func (o MachineOutput) Identity() MachineResponseIdentityPtrOutput {
	return o.ApplyT(func(v *Machine) MachineResponseIdentityPtrOutput { return v.Identity }).(MachineResponseIdentityPtrOutput)
}

func (o MachineOutput) LastStatusChange() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.LastStatusChange }).(pulumi.StringOutput)
}

func (o MachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MachineOutput) LocationData() LocationDataResponsePtrOutput {
	return o.ApplyT(func(v *Machine) LocationDataResponsePtrOutput { return v.LocationData }).(LocationDataResponsePtrOutput)
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

func (o MachineOutput) OsProfile() MachinePropertiesResponseOsProfilePtrOutput {
	return o.ApplyT(func(v *Machine) MachinePropertiesResponseOsProfilePtrOutput { return v.OsProfile }).(MachinePropertiesResponseOsProfilePtrOutput)
}

func (o MachineOutput) OsSku() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.OsSku }).(pulumi.StringOutput)
}

func (o MachineOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.OsVersion }).(pulumi.StringOutput)
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

func (o MachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o MachineOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.VmId }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) VmUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.VmUuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineOutput{})
}
