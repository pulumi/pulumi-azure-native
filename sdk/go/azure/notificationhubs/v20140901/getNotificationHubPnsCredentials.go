


package v20140901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNotificationHubPnsCredentials(ctx *pulumi.Context, args *GetNotificationHubPnsCredentialsArgs, opts ...pulumi.InvokeOption) (*GetNotificationHubPnsCredentialsResult, error) {
	var rv GetNotificationHubPnsCredentialsResult
	err := ctx.Invoke("azure-native:notificationhubs/v20140901:getNotificationHubPnsCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetNotificationHubPnsCredentialsArgs struct {
	NamespaceName       string `pulumi:"namespaceName"`
	NotificationHubName string `pulumi:"notificationHubName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type GetNotificationHubPnsCredentialsResult struct {
	Id         *string                           `pulumi:"id"`
	Location   *string                           `pulumi:"location"`
	Name       *string                           `pulumi:"name"`
	Properties NotificationHubPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       *string                           `pulumi:"type"`
}
