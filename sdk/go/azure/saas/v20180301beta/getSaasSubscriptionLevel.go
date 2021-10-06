


package v20180301beta

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSaasSubscriptionLevel(ctx *pulumi.Context, args *LookupSaasSubscriptionLevelArgs, opts ...pulumi.InvokeOption) (*LookupSaasSubscriptionLevelResult, error) {
	var rv LookupSaasSubscriptionLevelResult
	err := ctx.Invoke("azure-native:saas/v20180301beta:getSaasSubscriptionLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSaasSubscriptionLevelArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSaasSubscriptionLevelResult struct {
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties SaasResourceResponseProperties `pulumi:"properties"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}
