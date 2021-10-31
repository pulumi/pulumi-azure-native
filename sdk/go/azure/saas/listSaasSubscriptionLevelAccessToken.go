


package saas

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSaasSubscriptionLevelAccessToken(ctx *pulumi.Context, args *ListSaasSubscriptionLevelAccessTokenArgs, opts ...pulumi.InvokeOption) (*ListSaasSubscriptionLevelAccessTokenResult, error) {
	var rv ListSaasSubscriptionLevelAccessTokenResult
	err := ctx.Invoke("azure-native:saas:listSaasSubscriptionLevelAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSaasSubscriptionLevelAccessTokenArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListSaasSubscriptionLevelAccessTokenResult struct {
	PublisherOfferBaseUri *string `pulumi:"publisherOfferBaseUri"`
	Token                 *string `pulumi:"token"`
}
