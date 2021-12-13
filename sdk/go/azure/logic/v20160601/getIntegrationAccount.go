


package v20160601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccount(ctx *pulumi.Context, args *LookupIntegrationAccountArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountResult, error) {
	var rv LookupIntegrationAccountResult
	err := ctx.Invoke("azure-native:logic/v20160601:getIntegrationAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountResult struct {
	Id       string                         `pulumi:"id"`
	Location *string                        `pulumi:"location"`
	Name     string                         `pulumi:"name"`
	Sku      *IntegrationAccountSkuResponse `pulumi:"sku"`
	Tags     map[string]string              `pulumi:"tags"`
	Type     string                         `pulumi:"type"`
}
