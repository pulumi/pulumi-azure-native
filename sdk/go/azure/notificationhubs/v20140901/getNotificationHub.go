


package v20140901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNotificationHub(ctx *pulumi.Context, args *LookupNotificationHubArgs, opts ...pulumi.InvokeOption) (*LookupNotificationHubResult, error) {
	var rv LookupNotificationHubResult
	err := ctx.Invoke("azure-native:notificationhubs/v20140901:getNotificationHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationHubArgs struct {
	NamespaceName       string `pulumi:"namespaceName"`
	NotificationHubName string `pulumi:"notificationHubName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupNotificationHubResult struct {
	Id         *string                           `pulumi:"id"`
	Location   *string                           `pulumi:"location"`
	Name       *string                           `pulumi:"name"`
	Properties NotificationHubPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       *string                           `pulumi:"type"`
}
