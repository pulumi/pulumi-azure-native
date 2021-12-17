


package v20160515

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotificationChannel struct {
	pulumi.CustomResourceState

	CreatedDate       pulumi.StringOutput      `pulumi:"createdDate"`
	Description       pulumi.StringPtrOutput   `pulumi:"description"`
	Events            EventResponseArrayOutput `pulumi:"events"`
	Location          pulumi.StringPtrOutput   `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringPtrOutput   `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
	UniqueIdentifier  pulumi.StringPtrOutput   `pulumi:"uniqueIdentifier"`
	WebHookUrl        pulumi.StringPtrOutput   `pulumi:"webHookUrl"`
}


func NewNotificationChannel(ctx *pulumi.Context,
	name string, args *NotificationChannelArgs, opts ...pulumi.ResourceOption) (*NotificationChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:NotificationChannel"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:NotificationChannel"),
		},
	})
	opts = append(opts, aliases)
	var resource NotificationChannel
	err := ctx.RegisterResource("azure-native:devtestlab/v20160515:NotificationChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNotificationChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationChannelState, opts ...pulumi.ResourceOption) (*NotificationChannel, error) {
	var resource NotificationChannel
	err := ctx.ReadResource("azure-native:devtestlab/v20160515:NotificationChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type notificationChannelState struct {
}

type NotificationChannelState struct {
}

func (NotificationChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelState)(nil)).Elem()
}

type notificationChannelArgs struct {
	Description       *string           `pulumi:"description"`
	Events            []Event           `pulumi:"events"`
	LabName           string            `pulumi:"labName"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	UniqueIdentifier  *string           `pulumi:"uniqueIdentifier"`
	WebHookUrl        *string           `pulumi:"webHookUrl"`
}


type NotificationChannelArgs struct {
	Description       pulumi.StringPtrInput
	Events            EventArrayInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	UniqueIdentifier  pulumi.StringPtrInput
	WebHookUrl        pulumi.StringPtrInput
}

func (NotificationChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelArgs)(nil)).Elem()
}

type NotificationChannelInput interface {
	pulumi.Input

	ToNotificationChannelOutput() NotificationChannelOutput
	ToNotificationChannelOutputWithContext(ctx context.Context) NotificationChannelOutput
}

func (*NotificationChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannel)(nil)).Elem()
}

func (i *NotificationChannel) ToNotificationChannelOutput() NotificationChannelOutput {
	return i.ToNotificationChannelOutputWithContext(context.Background())
}

func (i *NotificationChannel) ToNotificationChannelOutputWithContext(ctx context.Context) NotificationChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelOutput)
}

type NotificationChannelOutput struct{ *pulumi.OutputState }

func (NotificationChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannel)(nil)).Elem()
}

func (o NotificationChannelOutput) ToNotificationChannelOutput() NotificationChannelOutput {
	return o
}

func (o NotificationChannelOutput) ToNotificationChannelOutputWithContext(ctx context.Context) NotificationChannelOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NotificationChannelOutput{})
}
