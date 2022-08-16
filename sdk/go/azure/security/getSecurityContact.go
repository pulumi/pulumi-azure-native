


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityContact(ctx *pulumi.Context, args *LookupSecurityContactArgs, opts ...pulumi.InvokeOption) (*LookupSecurityContactResult, error) {
	var rv LookupSecurityContactResult
	err := ctx.Invoke("azure-native:security:getSecurityContact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityContactArgs struct {
	SecurityContactName string `pulumi:"securityContactName"`
}


type LookupSecurityContactResult struct {
	AlertNotifications  *SecurityContactPropertiesResponseAlertNotifications  `pulumi:"alertNotifications"`
	Emails              *string                                               `pulumi:"emails"`
	Id                  string                                                `pulumi:"id"`
	Name                string                                                `pulumi:"name"`
	NotificationsByRole *SecurityContactPropertiesResponseNotificationsByRole `pulumi:"notificationsByRole"`
	Phone               *string                                               `pulumi:"phone"`
	Type                string                                                `pulumi:"type"`
}

func LookupSecurityContactOutput(ctx *pulumi.Context, args LookupSecurityContactOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityContactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityContactResult, error) {
			args := v.(LookupSecurityContactArgs)
			r, err := LookupSecurityContact(ctx, &args, opts...)
			var s LookupSecurityContactResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityContactResultOutput)
}

type LookupSecurityContactOutputArgs struct {
	SecurityContactName pulumi.StringInput `pulumi:"securityContactName"`
}

func (LookupSecurityContactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityContactArgs)(nil)).Elem()
}


type LookupSecurityContactResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityContactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityContactResult)(nil)).Elem()
}

func (o LookupSecurityContactResultOutput) ToLookupSecurityContactResultOutput() LookupSecurityContactResultOutput {
	return o
}

func (o LookupSecurityContactResultOutput) ToLookupSecurityContactResultOutputWithContext(ctx context.Context) LookupSecurityContactResultOutput {
	return o
}

func (o LookupSecurityContactResultOutput) AlertNotifications() SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return o.ApplyT(func(v LookupSecurityContactResult) *SecurityContactPropertiesResponseAlertNotifications {
		return v.AlertNotifications
	}).(SecurityContactPropertiesResponseAlertNotificationsPtrOutput)
}

func (o LookupSecurityContactResultOutput) Emails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityContactResult) *string { return v.Emails }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityContactResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityContactResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecurityContactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityContactResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecurityContactResultOutput) NotificationsByRole() SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return o.ApplyT(func(v LookupSecurityContactResult) *SecurityContactPropertiesResponseNotificationsByRole {
		return v.NotificationsByRole
	}).(SecurityContactPropertiesResponseNotificationsByRolePtrOutput)
}

func (o LookupSecurityContactResultOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityContactResult) *string { return v.Phone }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityContactResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityContactResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityContactResultOutput{})
}
