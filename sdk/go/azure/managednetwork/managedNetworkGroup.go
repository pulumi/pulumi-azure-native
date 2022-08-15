


package managednetwork

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedNetworkGroup struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput           `pulumi:"etag"`
	Kind              pulumi.StringPtrOutput        `pulumi:"kind"`
	Location          pulumi.StringPtrOutput        `pulumi:"location"`
	ManagementGroups  ResourceIdResponseArrayOutput `pulumi:"managementGroups"`
	Name              pulumi.StringOutput           `pulumi:"name"`
	ProvisioningState pulumi.StringOutput           `pulumi:"provisioningState"`
	Subnets           ResourceIdResponseArrayOutput `pulumi:"subnets"`
	Subscriptions     ResourceIdResponseArrayOutput `pulumi:"subscriptions"`
	Type              pulumi.StringOutput           `pulumi:"type"`
	VirtualNetworks   ResourceIdResponseArrayOutput `pulumi:"virtualNetworks"`
}


func NewManagedNetworkGroup(ctx *pulumi.Context,
	name string, args *ManagedNetworkGroupArgs, opts ...pulumi.ResourceOption) (*ManagedNetworkGroup, error) {
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
			Type: pulumi.String("azure-native:managednetwork/v20190601preview:ManagedNetworkGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedNetworkGroup
	err := ctx.RegisterResource("azure-native:managednetwork:ManagedNetworkGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedNetworkGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedNetworkGroupState, opts ...pulumi.ResourceOption) (*ManagedNetworkGroup, error) {
	var resource ManagedNetworkGroup
	err := ctx.ReadResource("azure-native:managednetwork:ManagedNetworkGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedNetworkGroupState struct {
}

type ManagedNetworkGroupState struct {
}

func (ManagedNetworkGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkGroupState)(nil)).Elem()
}

type managedNetworkGroupArgs struct {
	Kind                    *string      `pulumi:"kind"`
	Location                *string      `pulumi:"location"`
	ManagedNetworkGroupName *string      `pulumi:"managedNetworkGroupName"`
	ManagedNetworkName      string       `pulumi:"managedNetworkName"`
	ManagementGroups        []ResourceId `pulumi:"managementGroups"`
	ResourceGroupName       string       `pulumi:"resourceGroupName"`
	Subnets                 []ResourceId `pulumi:"subnets"`
	Subscriptions           []ResourceId `pulumi:"subscriptions"`
	VirtualNetworks         []ResourceId `pulumi:"virtualNetworks"`
}


type ManagedNetworkGroupArgs struct {
	Kind                    pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	ManagedNetworkGroupName pulumi.StringPtrInput
	ManagedNetworkName      pulumi.StringInput
	ManagementGroups        ResourceIdArrayInput
	ResourceGroupName       pulumi.StringInput
	Subnets                 ResourceIdArrayInput
	Subscriptions           ResourceIdArrayInput
	VirtualNetworks         ResourceIdArrayInput
}

func (ManagedNetworkGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkGroupArgs)(nil)).Elem()
}

type ManagedNetworkGroupInput interface {
	pulumi.Input

	ToManagedNetworkGroupOutput() ManagedNetworkGroupOutput
	ToManagedNetworkGroupOutputWithContext(ctx context.Context) ManagedNetworkGroupOutput
}

func (*ManagedNetworkGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetworkGroup)(nil)).Elem()
}

func (i *ManagedNetworkGroup) ToManagedNetworkGroupOutput() ManagedNetworkGroupOutput {
	return i.ToManagedNetworkGroupOutputWithContext(context.Background())
}

func (i *ManagedNetworkGroup) ToManagedNetworkGroupOutputWithContext(ctx context.Context) ManagedNetworkGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkGroupOutput)
}

type ManagedNetworkGroupOutput struct{ *pulumi.OutputState }

func (ManagedNetworkGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedNetworkGroup)(nil)).Elem()
}

func (o ManagedNetworkGroupOutput) ToManagedNetworkGroupOutput() ManagedNetworkGroupOutput {
	return o
}

func (o ManagedNetworkGroupOutput) ToManagedNetworkGroupOutputWithContext(ctx context.Context) ManagedNetworkGroupOutput {
	return o
}

func (o ManagedNetworkGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetworkGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ManagedNetworkGroupOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedNetworkGroup) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ManagedNetworkGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedNetworkGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ManagedNetworkGroupOutput) ManagementGroups() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v *ManagedNetworkGroup) ResourceIdResponseArrayOutput { return v.ManagementGroups }).(ResourceIdResponseArrayOutput)
}

func (o ManagedNetworkGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetworkGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedNetworkGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetworkGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedNetworkGroupOutput) Subnets() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v *ManagedNetworkGroup) ResourceIdResponseArrayOutput { return v.Subnets }).(ResourceIdResponseArrayOutput)
}

func (o ManagedNetworkGroupOutput) Subscriptions() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v *ManagedNetworkGroup) ResourceIdResponseArrayOutput { return v.Subscriptions }).(ResourceIdResponseArrayOutput)
}

func (o ManagedNetworkGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedNetworkGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedNetworkGroupOutput) VirtualNetworks() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v *ManagedNetworkGroup) ResourceIdResponseArrayOutput { return v.VirtualNetworks }).(ResourceIdResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedNetworkGroupOutput{})
}
