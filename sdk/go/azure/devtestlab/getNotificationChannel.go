


package devtestlab

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNotificationChannel(ctx *pulumi.Context, args *LookupNotificationChannelArgs, opts ...pulumi.InvokeOption) (*LookupNotificationChannelResult, error) {
	var rv LookupNotificationChannelResult
	err := ctx.Invoke("azure-native:devtestlab:getNotificationChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationChannelArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupNotificationChannelResult struct {
	CreatedDate        string            `pulumi:"createdDate"`
	Description        *string           `pulumi:"description"`
	EmailRecipient     *string           `pulumi:"emailRecipient"`
	Events             []EventResponse   `pulumi:"events"`
	Id                 string            `pulumi:"id"`
	Location           *string           `pulumi:"location"`
	Name               string            `pulumi:"name"`
	NotificationLocale *string           `pulumi:"notificationLocale"`
	ProvisioningState  string            `pulumi:"provisioningState"`
	Tags               map[string]string `pulumi:"tags"`
	Type               string            `pulumi:"type"`
	UniqueIdentifier   string            `pulumi:"uniqueIdentifier"`
	WebHookUrl         *string           `pulumi:"webHookUrl"`
}
