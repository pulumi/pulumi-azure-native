


package v20211115preview

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
	autoShutdownProfileApplier := func(v AutoShutdownProfile) *AutoShutdownProfile { return v.Defaults() }
	args.AutoShutdownProfile = args.AutoShutdownProfile.ToAutoShutdownProfileOutput().ApplyT(autoShutdownProfileApplier).(AutoShutdownProfilePtrOutput).Elem()
	connectionProfileApplier := func(v ConnectionProfile) *ConnectionProfile { return v.Defaults() }
	args.ConnectionProfile = args.ConnectionProfile.ToConnectionProfileOutput().ApplyT(connectionProfileApplier).(ConnectionProfilePtrOutput).Elem()
	virtualMachineProfileApplier := func(v VirtualMachineProfile) *VirtualMachineProfile { return v.Defaults() }
	args.VirtualMachineProfile = args.VirtualMachineProfile.ToVirtualMachineProfileOutput().ApplyT(virtualMachineProfileApplier).(VirtualMachineProfilePtrOutput).Elem()
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices/v20211001preview:Lab"),
		},
	})
	opts = append(opts, aliases)
	var resource Lab
	err := ctx.RegisterResource("azure-native:labservices/v20211115preview:Lab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabState, opts ...pulumi.ResourceOption) (*Lab, error) {
	var resource Lab
	err := ctx.ReadResource("azure-native:labservices/v20211115preview:Lab", name, id, state, &resource, opts...)
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

func init() {
	pulumi.RegisterOutputType(LabOutput{})
}
