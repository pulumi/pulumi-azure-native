


package v20180101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Subscription struct {
	pulumi.CustomResourceState

	CreatedDate      pulumi.StringOutput    `pulumi:"createdDate"`
	DisplayName      pulumi.StringPtrOutput `pulumi:"displayName"`
	EndDate          pulumi.StringPtrOutput `pulumi:"endDate"`
	ExpirationDate   pulumi.StringPtrOutput `pulumi:"expirationDate"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	NotificationDate pulumi.StringPtrOutput `pulumi:"notificationDate"`
	PrimaryKey       pulumi.StringOutput    `pulumi:"primaryKey"`
	ProductId        pulumi.StringOutput    `pulumi:"productId"`
	SecondaryKey     pulumi.StringOutput    `pulumi:"secondaryKey"`
	StartDate        pulumi.StringPtrOutput `pulumi:"startDate"`
	State            pulumi.StringOutput    `pulumi:"state"`
	StateComment     pulumi.StringPtrOutput `pulumi:"stateComment"`
	Type             pulumi.StringOutput    `pulumi:"type"`
	UserId           pulumi.StringOutput    `pulumi:"userId"`
}


func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Subscription"),
		},
	})
	opts = append(opts, aliases)
	var resource Subscription
	err := ctx.RegisterResource("azure-native:apimanagement/v20180101:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure-native:apimanagement/v20180101:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type subscriptionState struct {
}

type SubscriptionState struct {
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	DisplayName       string                 `pulumi:"displayName"`
	Notify            *bool                  `pulumi:"notify"`
	PrimaryKey        *string                `pulumi:"primaryKey"`
	ProductId         string                 `pulumi:"productId"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	SecondaryKey      *string                `pulumi:"secondaryKey"`
	ServiceName       string                 `pulumi:"serviceName"`
	Sid               *string                `pulumi:"sid"`
	State             *SubscriptionStateEnum `pulumi:"state"`
	UserId            string                 `pulumi:"userId"`
}


type SubscriptionArgs struct {
	DisplayName       pulumi.StringInput
	Notify            pulumi.BoolPtrInput
	PrimaryKey        pulumi.StringPtrInput
	ProductId         pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SecondaryKey      pulumi.StringPtrInput
	ServiceName       pulumi.StringInput
	Sid               pulumi.StringPtrInput
	State             SubscriptionStateEnumPtrInput
	UserId            pulumi.StringInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (*Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

type SubscriptionOutput struct{ *pulumi.OutputState }

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o SubscriptionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o SubscriptionOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o SubscriptionOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o SubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SubscriptionOutput) NotificationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.NotificationDate }).(pulumi.StringPtrOutput)
}

func (o SubscriptionOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o SubscriptionOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

func (o SubscriptionOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.SecondaryKey }).(pulumi.StringOutput)
}

func (o SubscriptionOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o SubscriptionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o SubscriptionOutput) StateComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.StateComment }).(pulumi.StringPtrOutput)
}

func (o SubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SubscriptionOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionOutput{})
}
