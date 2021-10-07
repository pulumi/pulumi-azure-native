


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityContact(ctx *pulumi.Context, args *LookupSecurityContactArgs, opts ...pulumi.InvokeOption) (*LookupSecurityContactResult, error) {
	var rv LookupSecurityContactResult
	err := ctx.Invoke("azure-native:security/v20200101preview:getSecurityContact", args, &rv, opts...)
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
