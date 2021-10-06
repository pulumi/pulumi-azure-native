


package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotificationRecipientEmail struct {
	pulumi.CustomResourceState

	Email pulumi.StringPtrOutput `pulumi:"email"`
	Name  pulumi.StringOutput    `pulumi:"name"`
	Type  pulumi.StringOutput    `pulumi:"type"`
}


func NewNotificationRecipientEmail(ctx *pulumi.Context,
	name string, args *NotificationRecipientEmailArgs, opts ...pulumi.ResourceOption) (*NotificationRecipientEmail, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NotificationName == nil {
		return nil, errors.New("invalid value for required argument 'NotificationName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:NotificationRecipientEmail"),
		},
	})
	opts = append(opts, aliases)
	var resource NotificationRecipientEmail
	err := ctx.RegisterResource("azure-native:apimanagement/v20191201:NotificationRecipientEmail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNotificationRecipientEmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationRecipientEmailState, opts ...pulumi.ResourceOption) (*NotificationRecipientEmail, error) {
	var resource NotificationRecipientEmail
	err := ctx.ReadResource("azure-native:apimanagement/v20191201:NotificationRecipientEmail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type notificationRecipientEmailState struct {
}

type NotificationRecipientEmailState struct {
}

func (NotificationRecipientEmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRecipientEmailState)(nil)).Elem()
}

type notificationRecipientEmailArgs struct {
	Email             *string `pulumi:"email"`
	NotificationName  string  `pulumi:"notificationName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type NotificationRecipientEmailArgs struct {
	Email             pulumi.StringPtrInput
	NotificationName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (NotificationRecipientEmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRecipientEmailArgs)(nil)).Elem()
}

type NotificationRecipientEmailInput interface {
	pulumi.Input

	ToNotificationRecipientEmailOutput() NotificationRecipientEmailOutput
	ToNotificationRecipientEmailOutputWithContext(ctx context.Context) NotificationRecipientEmailOutput
}

func (*NotificationRecipientEmail) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationRecipientEmail)(nil))
}

func (i *NotificationRecipientEmail) ToNotificationRecipientEmailOutput() NotificationRecipientEmailOutput {
	return i.ToNotificationRecipientEmailOutputWithContext(context.Background())
}

func (i *NotificationRecipientEmail) ToNotificationRecipientEmailOutputWithContext(ctx context.Context) NotificationRecipientEmailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRecipientEmailOutput)
}

type NotificationRecipientEmailOutput struct{ *pulumi.OutputState }

func (NotificationRecipientEmailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationRecipientEmail)(nil))
}

func (o NotificationRecipientEmailOutput) ToNotificationRecipientEmailOutput() NotificationRecipientEmailOutput {
	return o
}

func (o NotificationRecipientEmailOutput) ToNotificationRecipientEmailOutputWithContext(ctx context.Context) NotificationRecipientEmailOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NotificationRecipientEmailOutput{})
}
