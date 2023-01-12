


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityContact struct {
	pulumi.CustomResourceState

	AlertNotifications  SecurityContactPropertiesResponseAlertNotificationsPtrOutput  `pulumi:"alertNotifications"`
	Emails              pulumi.StringPtrOutput                                        `pulumi:"emails"`
	Name                pulumi.StringOutput                                           `pulumi:"name"`
	NotificationsByRole SecurityContactPropertiesResponseNotificationsByRolePtrOutput `pulumi:"notificationsByRole"`
	Phone               pulumi.StringPtrOutput                                        `pulumi:"phone"`
	Type                pulumi.StringOutput                                           `pulumi:"type"`
}


func NewSecurityContact(ctx *pulumi.Context,
	name string, args *SecurityContactArgs, opts ...pulumi.ResourceOption) (*SecurityContact, error) {
	if args == nil {
		args = &SecurityContactArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:SecurityContact"),
		},
		{
			Type: pulumi.String("azure-native:security/v20170801preview:SecurityContact"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityContact
	err := ctx.RegisterResource("azure-native:security/v20200101preview:SecurityContact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecurityContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityContactState, opts ...pulumi.ResourceOption) (*SecurityContact, error) {
	var resource SecurityContact
	err := ctx.ReadResource("azure-native:security/v20200101preview:SecurityContact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type securityContactState struct {
}

type SecurityContactState struct {
}

func (SecurityContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityContactState)(nil)).Elem()
}

type securityContactArgs struct {
	AlertNotifications  *SecurityContactPropertiesAlertNotifications  `pulumi:"alertNotifications"`
	Emails              *string                                       `pulumi:"emails"`
	NotificationsByRole *SecurityContactPropertiesNotificationsByRole `pulumi:"notificationsByRole"`
	Phone               *string                                       `pulumi:"phone"`
	SecurityContactName *string                                       `pulumi:"securityContactName"`
}


type SecurityContactArgs struct {
	AlertNotifications  SecurityContactPropertiesAlertNotificationsPtrInput
	Emails              pulumi.StringPtrInput
	NotificationsByRole SecurityContactPropertiesNotificationsByRolePtrInput
	Phone               pulumi.StringPtrInput
	SecurityContactName pulumi.StringPtrInput
}

func (SecurityContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityContactArgs)(nil)).Elem()
}

type SecurityContactInput interface {
	pulumi.Input

	ToSecurityContactOutput() SecurityContactOutput
	ToSecurityContactOutputWithContext(ctx context.Context) SecurityContactOutput
}

func (*SecurityContact) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContact)(nil)).Elem()
}

func (i *SecurityContact) ToSecurityContactOutput() SecurityContactOutput {
	return i.ToSecurityContactOutputWithContext(context.Background())
}

func (i *SecurityContact) ToSecurityContactOutputWithContext(ctx context.Context) SecurityContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactOutput)
}

type SecurityContactOutput struct{ *pulumi.OutputState }

func (SecurityContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContact)(nil)).Elem()
}

func (o SecurityContactOutput) ToSecurityContactOutput() SecurityContactOutput {
	return o
}

func (o SecurityContactOutput) ToSecurityContactOutputWithContext(ctx context.Context) SecurityContactOutput {
	return o
}

func (o SecurityContactOutput) AlertNotifications() SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return o.ApplyT(func(v *SecurityContact) SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
		return v.AlertNotifications
	}).(SecurityContactPropertiesResponseAlertNotificationsPtrOutput)
}

func (o SecurityContactOutput) Emails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContact) pulumi.StringPtrOutput { return v.Emails }).(pulumi.StringPtrOutput)
}

func (o SecurityContactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityContact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SecurityContactOutput) NotificationsByRole() SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return o.ApplyT(func(v *SecurityContact) SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
		return v.NotificationsByRole
	}).(SecurityContactPropertiesResponseNotificationsByRolePtrOutput)
}

func (o SecurityContactOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContact) pulumi.StringPtrOutput { return v.Phone }).(pulumi.StringPtrOutput)
}

func (o SecurityContactOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityContact) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityContactOutput{})
}
