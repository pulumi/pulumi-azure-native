


package v20160301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNotificationHubAuthorizationRule(ctx *pulumi.Context, args *LookupNotificationHubAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNotificationHubAuthorizationRuleResult, error) {
	var rv LookupNotificationHubAuthorizationRuleResult
	err := ctx.Invoke("azure-native:notificationhubs/v20160301:getNotificationHubAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationHubAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	NotificationHubName   string `pulumi:"notificationHubName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupNotificationHubAuthorizationRuleResult struct {
	Id       string            `pulumi:"id"`
	Location string            `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Rights   []string          `pulumi:"rights"`
	Sku      *SkuResponse      `pulumi:"sku"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}
