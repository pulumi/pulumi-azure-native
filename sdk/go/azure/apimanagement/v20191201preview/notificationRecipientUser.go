


package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotificationRecipientUser struct {
	pulumi.CustomResourceState

	Name   pulumi.StringOutput    `pulumi:"name"`
	Type   pulumi.StringOutput    `pulumi:"type"`
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
}


func NewNotificationRecipientUser(ctx *pulumi.Context,
	name string, args *NotificationRecipientUserArgs, opts ...pulumi.ResourceOption) (*NotificationRecipientUser, error) {
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
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210801:NotificationRecipientUser"),
		},
	})
	opts = append(opts, aliases)
	var resource NotificationRecipientUser
	err := ctx.RegisterResource("azure-native:apimanagement/v20191201preview:NotificationRecipientUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNotificationRecipientUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationRecipientUserState, opts ...pulumi.ResourceOption) (*NotificationRecipientUser, error) {
	var resource NotificationRecipientUser
	err := ctx.ReadResource("azure-native:apimanagement/v20191201preview:NotificationRecipientUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type notificationRecipientUserState struct {
}

type NotificationRecipientUserState struct {
}

func (NotificationRecipientUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRecipientUserState)(nil)).Elem()
}

type notificationRecipientUserArgs struct {
	NotificationName  string  `pulumi:"notificationName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	UserId            *string `pulumi:"userId"`
}


type NotificationRecipientUserArgs struct {
	NotificationName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	UserId            pulumi.StringPtrInput
}

func (NotificationRecipientUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRecipientUserArgs)(nil)).Elem()
}

type NotificationRecipientUserInput interface {
	pulumi.Input

	ToNotificationRecipientUserOutput() NotificationRecipientUserOutput
	ToNotificationRecipientUserOutputWithContext(ctx context.Context) NotificationRecipientUserOutput
}

func (*NotificationRecipientUser) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationRecipientUser)(nil))
}

func (i *NotificationRecipientUser) ToNotificationRecipientUserOutput() NotificationRecipientUserOutput {
	return i.ToNotificationRecipientUserOutputWithContext(context.Background())
}

func (i *NotificationRecipientUser) ToNotificationRecipientUserOutputWithContext(ctx context.Context) NotificationRecipientUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRecipientUserOutput)
}

type NotificationRecipientUserOutput struct{ *pulumi.OutputState }

func (NotificationRecipientUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationRecipientUser)(nil))
}

func (o NotificationRecipientUserOutput) ToNotificationRecipientUserOutput() NotificationRecipientUserOutput {
	return o
}

func (o NotificationRecipientUserOutput) ToNotificationRecipientUserOutputWithContext(ctx context.Context) NotificationRecipientUserOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationRecipientUserInput)(nil)).Elem(), &NotificationRecipientUser{})
	pulumi.RegisterOutputType(NotificationRecipientUserOutput{})
}
