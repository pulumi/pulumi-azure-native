


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedNetworkPeeringPolicy struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                              `pulumi:"location"`
	Name       pulumi.StringOutput                                 `pulumi:"name"`
	Properties ManagedNetworkPeeringPolicyPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                                 `pulumi:"type"`
}


func NewManagedNetworkPeeringPolicy(ctx *pulumi.Context,
	name string, args *ManagedNetworkPeeringPolicyArgs, opts ...pulumi.ResourceOption) (*ManagedNetworkPeeringPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedNetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetwork:ManagedNetworkPeeringPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedNetworkPeeringPolicy
	err := ctx.RegisterResource("azure-native:managednetwork/v20190601preview:ManagedNetworkPeeringPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedNetworkPeeringPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedNetworkPeeringPolicyState, opts ...pulumi.ResourceOption) (*ManagedNetworkPeeringPolicy, error) {
	var resource ManagedNetworkPeeringPolicy
	err := ctx.ReadResource("azure-native:managednetwork/v20190601preview:ManagedNetworkPeeringPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedNetworkPeeringPolicyState struct {
}

type ManagedNetworkPeeringPolicyState struct {
}

func (ManagedNetworkPeeringPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkPeeringPolicyState)(nil)).Elem()
}

type managedNetworkPeeringPolicyArgs struct {
	Location                        *string                                `pulumi:"location"`
	ManagedNetworkName              string                                 `pulumi:"managedNetworkName"`
	ManagedNetworkPeeringPolicyName *string                                `pulumi:"managedNetworkPeeringPolicyName"`
	Properties                      *ManagedNetworkPeeringPolicyProperties `pulumi:"properties"`
	ResourceGroupName               string                                 `pulumi:"resourceGroupName"`
}


type ManagedNetworkPeeringPolicyArgs struct {
	Location                        pulumi.StringPtrInput
	ManagedNetworkName              pulumi.StringInput
	ManagedNetworkPeeringPolicyName pulumi.StringPtrInput
	Properties                      ManagedNetworkPeeringPolicyPropertiesPtrInput
	ResourceGroupName               pulumi.StringInput
}

func (ManagedNetworkPeeringPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkPeeringPolicyArgs)(nil)).Elem()
}

type ManagedNetworkPeeringPolicyInput interface {
	pulumi.Input

	ToManagedNetworkPeeringPolicyOutput() ManagedNetworkPeeringPolicyOutput
	ToManagedNetworkPeeringPolicyOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyOutput
}

func (*ManagedNetworkPeeringPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetworkPeeringPolicy)(nil)).Elem()
}

func (i *ManagedNetworkPeeringPolicy) ToManagedNetworkPeeringPolicyOutput() ManagedNetworkPeeringPolicyOutput {
	return i.ToManagedNetworkPeeringPolicyOutputWithContext(context.Background())
}

func (i *ManagedNetworkPeeringPolicy) ToManagedNetworkPeeringPolicyOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkPeeringPolicyOutput)
}

type ManagedNetworkPeeringPolicyOutput struct{ *pulumi.OutputState }

func (ManagedNetworkPeeringPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetworkPeeringPolicy)(nil)).Elem()
}

func (o ManagedNetworkPeeringPolicyOutput) ToManagedNetworkPeeringPolicyOutput() ManagedNetworkPeeringPolicyOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyOutput) ToManagedNetworkPeeringPolicyOutputWithContext(ctx context.Context) ManagedNetworkPeeringPolicyOutput {
	return o
}

func (o ManagedNetworkPeeringPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ManagedNetworkPeeringPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedNetworkPeeringPolicyOutput) Properties() ManagedNetworkPeeringPolicyPropertiesResponseOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicy) ManagedNetworkPeeringPolicyPropertiesResponseOutput {
		return v.Properties
	}).(ManagedNetworkPeeringPolicyPropertiesResponseOutput)
}

func (o ManagedNetworkPeeringPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetworkPeeringPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedNetworkPeeringPolicyOutput{})
}
