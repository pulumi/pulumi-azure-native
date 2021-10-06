


package notificationhubs

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNotificationHubPnsCredentials(ctx *pulumi.Context, args *GetNotificationHubPnsCredentialsArgs, opts ...pulumi.InvokeOption) (*GetNotificationHubPnsCredentialsResult, error) {
	var rv GetNotificationHubPnsCredentialsResult
	err := ctx.Invoke("azure-native:notificationhubs:getNotificationHubPnsCredentials", args, &rv, opts...)
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
	AdmCredential   *AdmCredentialResponse   `pulumi:"admCredential"`
	ApnsCredential  *ApnsCredentialResponse  `pulumi:"apnsCredential"`
	BaiduCredential *BaiduCredentialResponse `pulumi:"baiduCredential"`
	GcmCredential   *GcmCredentialResponse   `pulumi:"gcmCredential"`
	Id              string                   `pulumi:"id"`
	Location        *string                  `pulumi:"location"`
	MpnsCredential  *MpnsCredentialResponse  `pulumi:"mpnsCredential"`
	Name            string                   `pulumi:"name"`
	Sku             *SkuResponse             `pulumi:"sku"`
	Tags            map[string]string        `pulumi:"tags"`
	Type            string                   `pulumi:"type"`
	WnsCredential   *WnsCredentialResponse   `pulumi:"wnsCredential"`
}
