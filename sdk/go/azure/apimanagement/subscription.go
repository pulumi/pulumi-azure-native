


package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Subscription struct {
	pulumi.CustomResourceState

	AllowTracing     pulumi.BoolPtrOutput   `pulumi:"allowTracing"`
	CreatedDate      pulumi.StringOutput    `pulumi:"createdDate"`
	DisplayName      pulumi.StringPtrOutput `pulumi:"displayName"`
	EndDate          pulumi.StringPtrOutput `pulumi:"endDate"`
	ExpirationDate   pulumi.StringPtrOutput `pulumi:"expirationDate"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	NotificationDate pulumi.StringPtrOutput `pulumi:"notificationDate"`
	OwnerId          pulumi.StringPtrOutput `pulumi:"ownerId"`
	PrimaryKey       pulumi.StringPtrOutput `pulumi:"primaryKey"`
	Scope            pulumi.StringOutput    `pulumi:"scope"`
	SecondaryKey     pulumi.StringPtrOutput `pulumi:"secondaryKey"`
	StartDate        pulumi.StringPtrOutput `pulumi:"startDate"`
	State            pulumi.StringOutput    `pulumi:"state"`
	StateComment     pulumi.StringPtrOutput `pulumi:"stateComment"`
	Type             pulumi.StringOutput    `pulumi:"type"`
}


func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20160707:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20161010:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210801:Subscription"),
		},
	})
	opts = append(opts, aliases)
	var resource Subscription
	err := ctx.RegisterResource("azure-native:apimanagement:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure-native:apimanagement:Subscription", name, id, state, &resource, opts...)
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
	AllowTracing      *bool                  `pulumi:"allowTracing"`
	AppType           *string                `pulumi:"appType"`
	DisplayName       string                 `pulumi:"displayName"`
	Notify            *bool                  `pulumi:"notify"`
	OwnerId           *string                `pulumi:"ownerId"`
	PrimaryKey        *string                `pulumi:"primaryKey"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	Scope             string                 `pulumi:"scope"`
	SecondaryKey      *string                `pulumi:"secondaryKey"`
	ServiceName       string                 `pulumi:"serviceName"`
	Sid               *string                `pulumi:"sid"`
	State             *SubscriptionStateEnum `pulumi:"state"`
}


type SubscriptionArgs struct {
	AllowTracing      pulumi.BoolPtrInput
	AppType           pulumi.StringPtrInput
	DisplayName       pulumi.StringInput
	Notify            pulumi.BoolPtrInput
	OwnerId           pulumi.StringPtrInput
	PrimaryKey        pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Scope             pulumi.StringInput
	SecondaryKey      pulumi.StringPtrInput
	ServiceName       pulumi.StringInput
	Sid               pulumi.StringPtrInput
	State             SubscriptionStateEnumPtrInput
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
	return reflect.TypeOf((*Subscription)(nil))
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

type SubscriptionOutput struct{ *pulumi.OutputState }

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subscription)(nil))
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubscriptionOutput{})
}
