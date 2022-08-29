


package authorization

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetClientConfig(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetClientConfigResult, error) {
	var rv GetClientConfigResult
	err := ctx.Invoke("azure-native:authorization:getClientConfig", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}


type GetClientConfigResult struct {
	ClientId       string `pulumi:"clientId"`
	ObjectId       string `pulumi:"objectId"`
	SubscriptionId string `pulumi:"subscriptionId"`
	TenantId       string `pulumi:"tenantId"`
}
