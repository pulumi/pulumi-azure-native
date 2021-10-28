


package v20190802preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Machine struct {
	pulumi.CustomResourceState

	AgentVersion      pulumi.StringOutput                             `pulumi:"agentVersion"`
	ClientPublicKey   pulumi.StringPtrOutput                          `pulumi:"clientPublicKey"`
	DisplayName       pulumi.StringOutput                             `pulumi:"displayName"`
	ErrorDetails      ErrorDetailResponseArrayOutput                  `pulumi:"errorDetails"`
	Extensions        MachineExtensionInstanceViewResponseArrayOutput `pulumi:"extensions"`
	LastStatusChange  pulumi.StringOutput                             `pulumi:"lastStatusChange"`
	Location          pulumi.StringOutput                             `pulumi:"location"`
	MachineFqdn       pulumi.StringOutput                             `pulumi:"machineFqdn"`
	Name              pulumi.StringOutput                             `pulumi:"name"`
	OsName            pulumi.StringPtrOutput                          `pulumi:"osName"`
	OsProfile         OSProfileResponseOutput                         `pulumi:"osProfile"`
	OsVersion         pulumi.StringPtrOutput                          `pulumi:"osVersion"`
	PhysicalLocation  pulumi.StringPtrOutput                          `pulumi:"physicalLocation"`
	PrincipalId       pulumi.StringOutput                             `pulumi:"principalId"`
	ProvisioningState pulumi.StringOutput                             `pulumi:"provisioningState"`
	Status            pulumi.StringOutput                             `pulumi:"status"`
	Tags              pulumi.StringMapOutput                          `pulumi:"tags"`
	TenantId          pulumi.StringOutput                             `pulumi:"tenantId"`
	Type              pulumi.StringOutput                             `pulumi:"type"`
	VmId              pulumi.StringOutput                             `pulumi:"vmId"`
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
			Type: pulumi.String("azure-nextgen:hybridcompute/v20190802preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20190318preview:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute/v20190318preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20191212:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute/v20191212:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200730preview:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute/v20200730preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200802:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute/v20200802:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200815preview:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute/v20200815preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210128preview:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute/v20210128preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210325preview:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute/v20210325preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210422preview:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute/v20210422preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210517preview:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute/v20210517preview:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210520:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute/v20210520:Machine"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210610preview:Machine"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridcompute/v20210610preview:Machine"),
		},
	})
	opts = append(opts, aliases)
	var resource Machine
	err := ctx.RegisterResource("azure-native:hybridcompute/v20190802preview:Machine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineState, opts ...pulumi.ResourceOption) (*Machine, error) {
	var resource Machine
	err := ctx.ReadResource("azure-native:hybridcompute/v20190802preview:Machine", name, id, state, &resource, opts...)
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
	Location          *string                        `pulumi:"location"`
	Name              *string                        `pulumi:"name"`
	OsName            *string                        `pulumi:"osName"`
	OsVersion         *string                        `pulumi:"osVersion"`
	PhysicalLocation  *string                        `pulumi:"physicalLocation"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	Tags              map[string]string              `pulumi:"tags"`
	Type              *string                        `pulumi:"type"`
}


type MachineArgs struct {
	ClientPublicKey   pulumi.StringPtrInput
	Extensions        MachineExtensionInstanceViewArrayInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	OsName            pulumi.StringPtrInput
	OsVersion         pulumi.StringPtrInput
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
	return reflect.TypeOf((*Machine)(nil))
}

func (i *Machine) ToMachineOutput() MachineOutput {
	return i.ToMachineOutputWithContext(context.Background())
}

func (i *Machine) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineOutput)
}

type MachineOutput struct{ *pulumi.OutputState }

func (MachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Machine)(nil))
}

func (o MachineOutput) ToMachineOutput() MachineOutput {
	return o
}

func (o MachineOutput) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineInput)(nil)).Elem(), &Machine{})
	pulumi.RegisterOutputType(MachineOutput{})
}
