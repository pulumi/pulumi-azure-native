


package v20190601preview

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
			Type: pulumi.String("azure-native:managednetwork:ManagedNetworkGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedNetworkGroup
	err := ctx.RegisterResource("azure-native:managednetwork/v20190601preview:ManagedNetworkGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedNetworkGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedNetworkGroupState, opts ...pulumi.ResourceOption) (*ManagedNetworkGroup, error) {
	var resource ManagedNetworkGroup
	err := ctx.ReadResource("azure-native:managednetwork/v20190601preview:ManagedNetworkGroup", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*ManagedNetworkGroup)(nil))
}

func (i *ManagedNetworkGroup) ToManagedNetworkGroupOutput() ManagedNetworkGroupOutput {
	return i.ToManagedNetworkGroupOutputWithContext(context.Background())
}

func (i *ManagedNetworkGroup) ToManagedNetworkGroupOutputWithContext(ctx context.Context) ManagedNetworkGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkGroupOutput)
}

type ManagedNetworkGroupOutput struct{ *pulumi.OutputState }

func (ManagedNetworkGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetworkGroup)(nil))
}

func (o ManagedNetworkGroupOutput) ToManagedNetworkGroupOutput() ManagedNetworkGroupOutput {
	return o
}

func (o ManagedNetworkGroupOutput) ToManagedNetworkGroupOutputWithContext(ctx context.Context) ManagedNetworkGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedNetworkGroupOutput{})
}
