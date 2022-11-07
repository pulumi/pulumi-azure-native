


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubscriptionNetworkManagerConnection struct {
	pulumi.CustomResourceState

	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	Etag             pulumi.StringOutput      `pulumi:"etag"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	NetworkManagerId pulumi.StringPtrOutput   `pulumi:"networkManagerId"`
	SystemData       SystemDataResponseOutput `pulumi:"systemData"`
	Type             pulumi.StringOutput      `pulumi:"type"`
}


func NewSubscriptionNetworkManagerConnection(ctx *pulumi.Context,
	name string, args *SubscriptionNetworkManagerConnectionArgs, opts ...pulumi.ResourceOption) (*SubscriptionNetworkManagerConnection, error) {
	if args == nil {
		args = &SubscriptionNetworkManagerConnectionArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210501preview:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:SubscriptionNetworkManagerConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource SubscriptionNetworkManagerConnection
	err := ctx.RegisterResource("azure-native:network:SubscriptionNetworkManagerConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubscriptionNetworkManagerConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionNetworkManagerConnectionState, opts ...pulumi.ResourceOption) (*SubscriptionNetworkManagerConnection, error) {
	var resource SubscriptionNetworkManagerConnection
	err := ctx.ReadResource("azure-native:network:SubscriptionNetworkManagerConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type subscriptionNetworkManagerConnectionState struct {
}

type SubscriptionNetworkManagerConnectionState struct {
}

func (SubscriptionNetworkManagerConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionNetworkManagerConnectionState)(nil)).Elem()
}

type subscriptionNetworkManagerConnectionArgs struct {
	Description                  *string `pulumi:"description"`
	NetworkManagerConnectionName *string `pulumi:"networkManagerConnectionName"`
	NetworkManagerId             *string `pulumi:"networkManagerId"`
}


type SubscriptionNetworkManagerConnectionArgs struct {
	Description                  pulumi.StringPtrInput
	NetworkManagerConnectionName pulumi.StringPtrInput
	NetworkManagerId             pulumi.StringPtrInput
}

func (SubscriptionNetworkManagerConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionNetworkManagerConnectionArgs)(nil)).Elem()
}

type SubscriptionNetworkManagerConnectionInput interface {
	pulumi.Input

	ToSubscriptionNetworkManagerConnectionOutput() SubscriptionNetworkManagerConnectionOutput
	ToSubscriptionNetworkManagerConnectionOutputWithContext(ctx context.Context) SubscriptionNetworkManagerConnectionOutput
}

func (*SubscriptionNetworkManagerConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionNetworkManagerConnection)(nil)).Elem()
}

func (i *SubscriptionNetworkManagerConnection) ToSubscriptionNetworkManagerConnectionOutput() SubscriptionNetworkManagerConnectionOutput {
	return i.ToSubscriptionNetworkManagerConnectionOutputWithContext(context.Background())
}

func (i *SubscriptionNetworkManagerConnection) ToSubscriptionNetworkManagerConnectionOutputWithContext(ctx context.Context) SubscriptionNetworkManagerConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionNetworkManagerConnectionOutput)
}

type SubscriptionNetworkManagerConnectionOutput struct{ *pulumi.OutputState }

func (SubscriptionNetworkManagerConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionNetworkManagerConnection)(nil)).Elem()
}

func (o SubscriptionNetworkManagerConnectionOutput) ToSubscriptionNetworkManagerConnectionOutput() SubscriptionNetworkManagerConnectionOutput {
	return o
}

func (o SubscriptionNetworkManagerConnectionOutput) ToSubscriptionNetworkManagerConnectionOutputWithContext(ctx context.Context) SubscriptionNetworkManagerConnectionOutput {
	return o
}

func (o SubscriptionNetworkManagerConnectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SubscriptionNetworkManagerConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o SubscriptionNetworkManagerConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SubscriptionNetworkManagerConnectionOutput) NetworkManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) pulumi.StringPtrOutput { return v.NetworkManagerId }).(pulumi.StringPtrOutput)
}

func (o SubscriptionNetworkManagerConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SubscriptionNetworkManagerConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionNetworkManagerConnectionOutput{})
}
