


package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AvailabilityGroupListener struct {
	pulumi.CustomResourceState

	AvailabilityGroupName                    pulumi.StringPtrOutput                       `pulumi:"availabilityGroupName"`
	CreateDefaultAvailabilityGroupIfNotExist pulumi.BoolPtrOutput                         `pulumi:"createDefaultAvailabilityGroupIfNotExist"`
	LoadBalancerConfigurations               LoadBalancerConfigurationResponseArrayOutput `pulumi:"loadBalancerConfigurations"`
	Name                                     pulumi.StringOutput                          `pulumi:"name"`
	Port                                     pulumi.IntPtrOutput                          `pulumi:"port"`
	ProvisioningState                        pulumi.StringOutput                          `pulumi:"provisioningState"`
	Type                                     pulumi.StringOutput                          `pulumi:"type"`
}


func NewAvailabilityGroupListener(ctx *pulumi.Context,
	name string, args *AvailabilityGroupListenerArgs, opts ...pulumi.ResourceOption) (*AvailabilityGroupListener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlVirtualMachineGroupName == nil {
		return nil, errors.New("invalid value for required argument 'SqlVirtualMachineGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine:AvailabilityGroupListener"),
		},
	})
	opts = append(opts, aliases)
	var resource AvailabilityGroupListener
	err := ctx.RegisterResource("azure-native:sqlvirtualmachine/v20170301preview:AvailabilityGroupListener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAvailabilityGroupListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilityGroupListenerState, opts ...pulumi.ResourceOption) (*AvailabilityGroupListener, error) {
	var resource AvailabilityGroupListener
	err := ctx.ReadResource("azure-native:sqlvirtualmachine/v20170301preview:AvailabilityGroupListener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type availabilityGroupListenerState struct {
}

type AvailabilityGroupListenerState struct {
}

func (AvailabilityGroupListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilityGroupListenerState)(nil)).Elem()
}

type availabilityGroupListenerArgs struct {
	AvailabilityGroupListenerName            *string                     `pulumi:"availabilityGroupListenerName"`
	AvailabilityGroupName                    *string                     `pulumi:"availabilityGroupName"`
	CreateDefaultAvailabilityGroupIfNotExist *bool                       `pulumi:"createDefaultAvailabilityGroupIfNotExist"`
	LoadBalancerConfigurations               []LoadBalancerConfiguration `pulumi:"loadBalancerConfigurations"`
	Port                                     *int                        `pulumi:"port"`
	ResourceGroupName                        string                      `pulumi:"resourceGroupName"`
	SqlVirtualMachineGroupName               string                      `pulumi:"sqlVirtualMachineGroupName"`
}


type AvailabilityGroupListenerArgs struct {
	AvailabilityGroupListenerName            pulumi.StringPtrInput
	AvailabilityGroupName                    pulumi.StringPtrInput
	CreateDefaultAvailabilityGroupIfNotExist pulumi.BoolPtrInput
	LoadBalancerConfigurations               LoadBalancerConfigurationArrayInput
	Port                                     pulumi.IntPtrInput
	ResourceGroupName                        pulumi.StringInput
	SqlVirtualMachineGroupName               pulumi.StringInput
}

func (AvailabilityGroupListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilityGroupListenerArgs)(nil)).Elem()
}

type AvailabilityGroupListenerInput interface {
	pulumi.Input

	ToAvailabilityGroupListenerOutput() AvailabilityGroupListenerOutput
	ToAvailabilityGroupListenerOutputWithContext(ctx context.Context) AvailabilityGroupListenerOutput
}

func (*AvailabilityGroupListener) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityGroupListener)(nil))
}

func (i *AvailabilityGroupListener) ToAvailabilityGroupListenerOutput() AvailabilityGroupListenerOutput {
	return i.ToAvailabilityGroupListenerOutputWithContext(context.Background())
}

func (i *AvailabilityGroupListener) ToAvailabilityGroupListenerOutputWithContext(ctx context.Context) AvailabilityGroupListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityGroupListenerOutput)
}

type AvailabilityGroupListenerOutput struct{ *pulumi.OutputState }

func (AvailabilityGroupListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityGroupListener)(nil))
}

func (o AvailabilityGroupListenerOutput) ToAvailabilityGroupListenerOutput() AvailabilityGroupListenerOutput {
	return o
}

func (o AvailabilityGroupListenerOutput) ToAvailabilityGroupListenerOutputWithContext(ctx context.Context) AvailabilityGroupListenerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AvailabilityGroupListenerOutput{})
}
