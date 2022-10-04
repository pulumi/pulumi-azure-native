


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotificationRegistration struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                              `pulumi:"name"`
	Properties NotificationRegistrationResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                         `pulumi:"systemData"`
	Type       pulumi.StringOutput                              `pulumi:"type"`
}


func NewNotificationRegistration(ctx *pulumi.Context,
	name string, args *NotificationRegistrationArgs, opts ...pulumi.ResourceOption) (*NotificationRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:providerhub:NotificationRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:NotificationRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:NotificationRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210901preview:NotificationRegistration"),
		},
	})
	opts = append(opts, aliases)
	var resource NotificationRegistration
	err := ctx.RegisterResource("azure-native:providerhub/v20210601preview:NotificationRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNotificationRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationRegistrationState, opts ...pulumi.ResourceOption) (*NotificationRegistration, error) {
	var resource NotificationRegistration
	err := ctx.ReadResource("azure-native:providerhub/v20210601preview:NotificationRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type notificationRegistrationState struct {
}

type NotificationRegistrationState struct {
}

func (NotificationRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRegistrationState)(nil)).Elem()
}

type notificationRegistrationArgs struct {
	NotificationRegistrationName *string                             `pulumi:"notificationRegistrationName"`
	Properties                   *NotificationRegistrationProperties `pulumi:"properties"`
	ProviderNamespace            string                              `pulumi:"providerNamespace"`
}


type NotificationRegistrationArgs struct {
	NotificationRegistrationName pulumi.StringPtrInput
	Properties                   NotificationRegistrationPropertiesPtrInput
	ProviderNamespace            pulumi.StringInput
}

func (NotificationRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRegistrationArgs)(nil)).Elem()
}

type NotificationRegistrationInput interface {
	pulumi.Input

	ToNotificationRegistrationOutput() NotificationRegistrationOutput
	ToNotificationRegistrationOutputWithContext(ctx context.Context) NotificationRegistrationOutput
}

func (*NotificationRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationRegistration)(nil)).Elem()
}

func (i *NotificationRegistration) ToNotificationRegistrationOutput() NotificationRegistrationOutput {
	return i.ToNotificationRegistrationOutputWithContext(context.Background())
}

func (i *NotificationRegistration) ToNotificationRegistrationOutputWithContext(ctx context.Context) NotificationRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRegistrationOutput)
}

type NotificationRegistrationOutput struct{ *pulumi.OutputState }

func (NotificationRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationRegistration)(nil)).Elem()
}

func (o NotificationRegistrationOutput) ToNotificationRegistrationOutput() NotificationRegistrationOutput {
	return o
}

func (o NotificationRegistrationOutput) ToNotificationRegistrationOutputWithContext(ctx context.Context) NotificationRegistrationOutput {
	return o
}

func (o NotificationRegistrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationRegistration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NotificationRegistrationOutput) Properties() NotificationRegistrationResponsePropertiesOutput {
	return o.ApplyT(func(v *NotificationRegistration) NotificationRegistrationResponsePropertiesOutput {
		return v.Properties
	}).(NotificationRegistrationResponsePropertiesOutput)
}

func (o NotificationRegistrationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NotificationRegistration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o NotificationRegistrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationRegistration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationRegistrationOutput{})
}
