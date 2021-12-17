


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkGroup struct {
	pulumi.CustomResourceState

	ConditionalMembership pulumi.StringPtrOutput              `pulumi:"conditionalMembership"`
	Description           pulumi.StringPtrOutput              `pulumi:"description"`
	DisplayName           pulumi.StringPtrOutput              `pulumi:"displayName"`
	Etag                  pulumi.StringOutput                 `pulumi:"etag"`
	GroupMembers          GroupMembersItemResponseArrayOutput `pulumi:"groupMembers"`
	MemberType            pulumi.StringPtrOutput              `pulumi:"memberType"`
	Name                  pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput                 `pulumi:"provisioningState"`
	SystemData            SystemDataResponseOutput            `pulumi:"systemData"`
	Type                  pulumi.StringOutput                 `pulumi:"type"`
}


func NewNetworkGroup(ctx *pulumi.Context,
	name string, args *NetworkGroupArgs, opts ...pulumi.ResourceOption) (*NetworkGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210201preview:NetworkGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkGroup
	err := ctx.RegisterResource("azure-native:network:NetworkGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkGroupState, opts ...pulumi.ResourceOption) (*NetworkGroup, error) {
	var resource NetworkGroup
	err := ctx.ReadResource("azure-native:network:NetworkGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkGroupState struct {
}

type NetworkGroupState struct {
}

func (NetworkGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkGroupState)(nil)).Elem()
}

type networkGroupArgs struct {
	ConditionalMembership *string            `pulumi:"conditionalMembership"`
	Description           *string            `pulumi:"description"`
	DisplayName           *string            `pulumi:"displayName"`
	GroupMembers          []GroupMembersItem `pulumi:"groupMembers"`
	MemberType            *string            `pulumi:"memberType"`
	NetworkGroupName      *string            `pulumi:"networkGroupName"`
	NetworkManagerName    string             `pulumi:"networkManagerName"`
	ResourceGroupName     string             `pulumi:"resourceGroupName"`
}


type NetworkGroupArgs struct {
	ConditionalMembership pulumi.StringPtrInput
	Description           pulumi.StringPtrInput
	DisplayName           pulumi.StringPtrInput
	GroupMembers          GroupMembersItemArrayInput
	MemberType            pulumi.StringPtrInput
	NetworkGroupName      pulumi.StringPtrInput
	NetworkManagerName    pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
}

func (NetworkGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkGroupArgs)(nil)).Elem()
}

type NetworkGroupInput interface {
	pulumi.Input

	ToNetworkGroupOutput() NetworkGroupOutput
	ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput
}

func (*NetworkGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkGroup)(nil)).Elem()
}

func (i *NetworkGroup) ToNetworkGroupOutput() NetworkGroupOutput {
	return i.ToNetworkGroupOutputWithContext(context.Background())
}

func (i *NetworkGroup) ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkGroupOutput)
}

type NetworkGroupOutput struct{ *pulumi.OutputState }

func (NetworkGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkGroup)(nil)).Elem()
}

func (o NetworkGroupOutput) ToNetworkGroupOutput() NetworkGroupOutput {
	return o
}

func (o NetworkGroupOutput) ToNetworkGroupOutputWithContext(ctx context.Context) NetworkGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkGroupOutput{})
}
