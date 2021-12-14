


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotificationChannel struct {
	pulumi.CustomResourceState

	CreatedDate        pulumi.StringOutput      `pulumi:"createdDate"`
	Description        pulumi.StringPtrOutput   `pulumi:"description"`
	EmailRecipient     pulumi.StringPtrOutput   `pulumi:"emailRecipient"`
	Events             EventResponseArrayOutput `pulumi:"events"`
	Location           pulumi.StringPtrOutput   `pulumi:"location"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	NotificationLocale pulumi.StringPtrOutput   `pulumi:"notificationLocale"`
	ProvisioningState  pulumi.StringOutput      `pulumi:"provisioningState"`
	Tags               pulumi.StringMapOutput   `pulumi:"tags"`
	Type               pulumi.StringOutput      `pulumi:"type"`
	UniqueIdentifier   pulumi.StringOutput      `pulumi:"uniqueIdentifier"`
	WebHookUrl         pulumi.StringPtrOutput   `pulumi:"webHookUrl"`
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
			Type: pulumi.String("azure-native:devtestlab/v20160515:NotificationChannel"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:NotificationChannel"),
		},
	})
	opts = append(opts, aliases)
	var resource NotificationChannel
	err := ctx.RegisterResource("azure-native:devtestlab:NotificationChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNotificationChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationChannelState, opts ...pulumi.ResourceOption) (*NotificationChannel, error) {
	var resource NotificationChannel
	err := ctx.ReadResource("azure-native:devtestlab:NotificationChannel", name, id, state, &resource, opts...)
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
	Description        *string           `pulumi:"description"`
	EmailRecipient     *string           `pulumi:"emailRecipient"`
	Events             []Event           `pulumi:"events"`
	LabName            string            `pulumi:"labName"`
	Location           *string           `pulumi:"location"`
	Name               *string           `pulumi:"name"`
	NotificationLocale *string           `pulumi:"notificationLocale"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	Tags               map[string]string `pulumi:"tags"`
	WebHookUrl         *string           `pulumi:"webHookUrl"`
}


type NotificationChannelArgs struct {
	Description        pulumi.StringPtrInput
	EmailRecipient     pulumi.StringPtrInput
	Events             EventArrayInput
	LabName            pulumi.StringInput
	Location           pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	NotificationLocale pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	Tags               pulumi.StringMapInput
	WebHookUrl         pulumi.StringPtrInput
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
