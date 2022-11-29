


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

func (o NotificationChannelOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o NotificationChannelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NotificationChannelOutput) EmailRecipient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringPtrOutput { return v.EmailRecipient }).(pulumi.StringPtrOutput)
}

func (o NotificationChannelOutput) Events() EventResponseArrayOutput {
	return o.ApplyT(func(v *NotificationChannel) EventResponseArrayOutput { return v.Events }).(EventResponseArrayOutput)
}

func (o NotificationChannelOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NotificationChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NotificationChannelOutput) NotificationLocale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringPtrOutput { return v.NotificationLocale }).(pulumi.StringPtrOutput)
}

func (o NotificationChannelOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NotificationChannelOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NotificationChannelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NotificationChannelOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o NotificationChannelOutput) WebHookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringPtrOutput { return v.WebHookUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationChannelOutput{})
}
