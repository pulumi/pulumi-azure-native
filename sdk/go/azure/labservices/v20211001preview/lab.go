


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Lab struct {
	pulumi.CustomResourceState

	AutoShutdownProfile   AutoShutdownProfileResponseOutput   `pulumi:"autoShutdownProfile"`
	ConnectionProfile     ConnectionProfileResponseOutput     `pulumi:"connectionProfile"`
	Description           pulumi.StringPtrOutput              `pulumi:"description"`
	LabPlanId             pulumi.StringPtrOutput              `pulumi:"labPlanId"`
	Location              pulumi.StringOutput                 `pulumi:"location"`
	Name                  pulumi.StringOutput                 `pulumi:"name"`
	NetworkProfile        LabNetworkProfileResponsePtrOutput  `pulumi:"networkProfile"`
	ProvisioningState     pulumi.StringOutput                 `pulumi:"provisioningState"`
	RosterProfile         RosterProfileResponsePtrOutput      `pulumi:"rosterProfile"`
	SecurityProfile       SecurityProfileResponseOutput       `pulumi:"securityProfile"`
	State                 pulumi.StringOutput                 `pulumi:"state"`
	SystemData            SystemDataResponseOutput            `pulumi:"systemData"`
	Tags                  pulumi.StringMapOutput              `pulumi:"tags"`
	Title                 pulumi.StringPtrOutput              `pulumi:"title"`
	Type                  pulumi.StringOutput                 `pulumi:"type"`
	VirtualMachineProfile VirtualMachineProfileResponseOutput `pulumi:"virtualMachineProfile"`
}


func NewLab(ctx *pulumi.Context,
	name string, args *LabArgs, opts ...pulumi.ResourceOption) (*Lab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoShutdownProfile == nil {
		return nil, errors.New("invalid value for required argument 'AutoShutdownProfile'")
	}
	if args.ConnectionProfile == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionProfile'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SecurityProfile == nil {
		return nil, errors.New("invalid value for required argument 'SecurityProfile'")
	}
	if args.VirtualMachineProfile == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineProfile'")
	}
	args.AutoShutdownProfile = args.AutoShutdownProfile.ToAutoShutdownProfileOutput().ApplyT(func(v AutoShutdownProfile) AutoShutdownProfile { return *v.Defaults() }).(AutoShutdownProfileOutput)
	args.ConnectionProfile = args.ConnectionProfile.ToConnectionProfileOutput().ApplyT(func(v ConnectionProfile) ConnectionProfile { return *v.Defaults() }).(ConnectionProfileOutput)
	args.VirtualMachineProfile = args.VirtualMachineProfile.ToVirtualMachineProfileOutput().ApplyT(func(v VirtualMachineProfile) VirtualMachineProfile { return *v.Defaults() }).(VirtualMachineProfileOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices/v20211115preview:Lab"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20220801:Lab"),
		},
	})
	opts = append(opts, aliases)
	var resource Lab
	err := ctx.RegisterResource("azure-native:labservices/v20211001preview:Lab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabState, opts ...pulumi.ResourceOption) (*Lab, error) {
	var resource Lab
	err := ctx.ReadResource("azure-native:labservices/v20211001preview:Lab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type labState struct {
}

type LabState struct {
}

func (LabState) ElementType() reflect.Type {
	return reflect.TypeOf((*labState)(nil)).Elem()
}

type labArgs struct {
	AutoShutdownProfile   AutoShutdownProfile   `pulumi:"autoShutdownProfile"`
	ConnectionProfile     ConnectionProfile     `pulumi:"connectionProfile"`
	Description           *string               `pulumi:"description"`
	LabName               *string               `pulumi:"labName"`
	LabPlanId             *string               `pulumi:"labPlanId"`
	Location              *string               `pulumi:"location"`
	NetworkProfile        *LabNetworkProfile    `pulumi:"networkProfile"`
	ResourceGroupName     string                `pulumi:"resourceGroupName"`
	RosterProfile         *RosterProfile        `pulumi:"rosterProfile"`
	SecurityProfile       SecurityProfile       `pulumi:"securityProfile"`
	Tags                  map[string]string     `pulumi:"tags"`
	Title                 *string               `pulumi:"title"`
	VirtualMachineProfile VirtualMachineProfile `pulumi:"virtualMachineProfile"`
}


type LabArgs struct {
	AutoShutdownProfile   AutoShutdownProfileInput
	ConnectionProfile     ConnectionProfileInput
	Description           pulumi.StringPtrInput
	LabName               pulumi.StringPtrInput
	LabPlanId             pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	NetworkProfile        LabNetworkProfilePtrInput
	ResourceGroupName     pulumi.StringInput
	RosterProfile         RosterProfilePtrInput
	SecurityProfile       SecurityProfileInput
	Tags                  pulumi.StringMapInput
	Title                 pulumi.StringPtrInput
	VirtualMachineProfile VirtualMachineProfileInput
}

func (LabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labArgs)(nil)).Elem()
}

type LabInput interface {
	pulumi.Input

	ToLabOutput() LabOutput
	ToLabOutputWithContext(ctx context.Context) LabOutput
}

func (*Lab) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil)).Elem()
}

func (i *Lab) ToLabOutput() LabOutput {
	return i.ToLabOutputWithContext(context.Background())
}

func (i *Lab) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabOutput)
}

type LabOutput struct{ *pulumi.OutputState }

func (LabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil)).Elem()
}

func (o LabOutput) ToLabOutput() LabOutput {
	return o
}

func (o LabOutput) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return o
}

func (o LabOutput) AutoShutdownProfile() AutoShutdownProfileResponseOutput {
	return o.ApplyT(func(v *Lab) AutoShutdownProfileResponseOutput { return v.AutoShutdownProfile }).(AutoShutdownProfileResponseOutput)
}

func (o LabOutput) ConnectionProfile() ConnectionProfileResponseOutput {
	return o.ApplyT(func(v *Lab) ConnectionProfileResponseOutput { return v.ConnectionProfile }).(ConnectionProfileResponseOutput)
}

func (o LabOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LabOutput) LabPlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.LabPlanId }).(pulumi.StringPtrOutput)
}

func (o LabOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o LabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LabOutput) NetworkProfile() LabNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *Lab) LabNetworkProfileResponsePtrOutput { return v.NetworkProfile }).(LabNetworkProfileResponsePtrOutput)
}

func (o LabOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LabOutput) RosterProfile() RosterProfileResponsePtrOutput {
	return o.ApplyT(func(v *Lab) RosterProfileResponsePtrOutput { return v.RosterProfile }).(RosterProfileResponsePtrOutput)
}

func (o LabOutput) SecurityProfile() SecurityProfileResponseOutput {
	return o.ApplyT(func(v *Lab) SecurityProfileResponseOutput { return v.SecurityProfile }).(SecurityProfileResponseOutput)
}

func (o LabOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o LabOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Lab) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LabOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LabOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

func (o LabOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o LabOutput) VirtualMachineProfile() VirtualMachineProfileResponseOutput {
	return o.ApplyT(func(v *Lab) VirtualMachineProfileResponseOutput { return v.VirtualMachineProfile }).(VirtualMachineProfileResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LabOutput{})
}
