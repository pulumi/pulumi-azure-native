


package v20220201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AvailabilityGroupListener struct {
	pulumi.CustomResourceState

	AvailabilityGroupConfiguration           AgConfigurationResponsePtrOutput              `pulumi:"availabilityGroupConfiguration"`
	AvailabilityGroupName                    pulumi.StringPtrOutput                        `pulumi:"availabilityGroupName"`
	CreateDefaultAvailabilityGroupIfNotExist pulumi.BoolPtrOutput                          `pulumi:"createDefaultAvailabilityGroupIfNotExist"`
	LoadBalancerConfigurations               LoadBalancerConfigurationResponseArrayOutput  `pulumi:"loadBalancerConfigurations"`
	MultiSubnetIpConfigurations              MultiSubnetIpConfigurationResponseArrayOutput `pulumi:"multiSubnetIpConfigurations"`
	Name                                     pulumi.StringOutput                           `pulumi:"name"`
	Port                                     pulumi.IntPtrOutput                           `pulumi:"port"`
	ProvisioningState                        pulumi.StringOutput                           `pulumi:"provisioningState"`
	SystemData                               SystemDataResponseOutput                      `pulumi:"systemData"`
	Type                                     pulumi.StringOutput                           `pulumi:"type"`
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
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20170301preview:AvailabilityGroupListener"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20211101preview:AvailabilityGroupListener"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220201preview:AvailabilityGroupListener"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220701preview:AvailabilityGroupListener"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220801preview:AvailabilityGroupListener"),
		},
	})
	opts = append(opts, aliases)
	var resource AvailabilityGroupListener
	err := ctx.RegisterResource("azure-native:sqlvirtualmachine/v20220201:AvailabilityGroupListener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAvailabilityGroupListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilityGroupListenerState, opts ...pulumi.ResourceOption) (*AvailabilityGroupListener, error) {
	var resource AvailabilityGroupListener
	err := ctx.ReadResource("azure-native:sqlvirtualmachine/v20220201:AvailabilityGroupListener", name, id, state, &resource, opts...)
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
	AvailabilityGroupConfiguration           *AgConfiguration             `pulumi:"availabilityGroupConfiguration"`
	AvailabilityGroupListenerName            *string                      `pulumi:"availabilityGroupListenerName"`
	AvailabilityGroupName                    *string                      `pulumi:"availabilityGroupName"`
	CreateDefaultAvailabilityGroupIfNotExist *bool                        `pulumi:"createDefaultAvailabilityGroupIfNotExist"`
	LoadBalancerConfigurations               []LoadBalancerConfiguration  `pulumi:"loadBalancerConfigurations"`
	MultiSubnetIpConfigurations              []MultiSubnetIpConfiguration `pulumi:"multiSubnetIpConfigurations"`
	Port                                     *int                         `pulumi:"port"`
	ResourceGroupName                        string                       `pulumi:"resourceGroupName"`
	SqlVirtualMachineGroupName               string                       `pulumi:"sqlVirtualMachineGroupName"`
}


type AvailabilityGroupListenerArgs struct {
	AvailabilityGroupConfiguration           AgConfigurationPtrInput
	AvailabilityGroupListenerName            pulumi.StringPtrInput
	AvailabilityGroupName                    pulumi.StringPtrInput
	CreateDefaultAvailabilityGroupIfNotExist pulumi.BoolPtrInput
	LoadBalancerConfigurations               LoadBalancerConfigurationArrayInput
	MultiSubnetIpConfigurations              MultiSubnetIpConfigurationArrayInput
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
	return reflect.TypeOf((**AvailabilityGroupListener)(nil)).Elem()
}

func (i *AvailabilityGroupListener) ToAvailabilityGroupListenerOutput() AvailabilityGroupListenerOutput {
	return i.ToAvailabilityGroupListenerOutputWithContext(context.Background())
}

func (i *AvailabilityGroupListener) ToAvailabilityGroupListenerOutputWithContext(ctx context.Context) AvailabilityGroupListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityGroupListenerOutput)
}

type AvailabilityGroupListenerOutput struct{ *pulumi.OutputState }

func (AvailabilityGroupListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilityGroupListener)(nil)).Elem()
}

func (o AvailabilityGroupListenerOutput) ToAvailabilityGroupListenerOutput() AvailabilityGroupListenerOutput {
	return o
}

func (o AvailabilityGroupListenerOutput) ToAvailabilityGroupListenerOutputWithContext(ctx context.Context) AvailabilityGroupListenerOutput {
	return o
}

func (o AvailabilityGroupListenerOutput) AvailabilityGroupConfiguration() AgConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AvailabilityGroupListener) AgConfigurationResponsePtrOutput {
		return v.AvailabilityGroupConfiguration
	}).(AgConfigurationResponsePtrOutput)
}

func (o AvailabilityGroupListenerOutput) AvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailabilityGroupListener) pulumi.StringPtrOutput { return v.AvailabilityGroupName }).(pulumi.StringPtrOutput)
}

func (o AvailabilityGroupListenerOutput) CreateDefaultAvailabilityGroupIfNotExist() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AvailabilityGroupListener) pulumi.BoolPtrOutput {
		return v.CreateDefaultAvailabilityGroupIfNotExist
	}).(pulumi.BoolPtrOutput)
}

func (o AvailabilityGroupListenerOutput) LoadBalancerConfigurations() LoadBalancerConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *AvailabilityGroupListener) LoadBalancerConfigurationResponseArrayOutput {
		return v.LoadBalancerConfigurations
	}).(LoadBalancerConfigurationResponseArrayOutput)
}

func (o AvailabilityGroupListenerOutput) MultiSubnetIpConfigurations() MultiSubnetIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *AvailabilityGroupListener) MultiSubnetIpConfigurationResponseArrayOutput {
		return v.MultiSubnetIpConfigurations
	}).(MultiSubnetIpConfigurationResponseArrayOutput)
}

func (o AvailabilityGroupListenerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilityGroupListener) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AvailabilityGroupListenerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilityGroupListener) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o AvailabilityGroupListenerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilityGroupListener) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AvailabilityGroupListenerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AvailabilityGroupListener) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AvailabilityGroupListenerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilityGroupListener) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AvailabilityGroupListenerOutput{})
}
