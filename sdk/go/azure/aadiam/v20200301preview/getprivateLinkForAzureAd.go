


package v20200301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkForAzureAd(ctx *pulumi.Context, args *GetprivateLinkForAzureAdArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkForAzureAdResult, error) {
	var rv GetprivateLinkForAzureAdResult
	err := ctx.Invoke("azure-native:aadiam/v20200301preview:getprivateLinkForAzureAd", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkForAzureAdArgs struct {
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetprivateLinkForAzureAdResult struct {
	AllTenants     *bool             `pulumi:"allTenants"`
	Id             string            `pulumi:"id"`
	Name           *string           `pulumi:"name"`
	OwnerTenantId  *string           `pulumi:"ownerTenantId"`
	ResourceGroup  *string           `pulumi:"resourceGroup"`
	ResourceName   *string           `pulumi:"resourceName"`
	SubscriptionId *string           `pulumi:"subscriptionId"`
	Tags           map[string]string `pulumi:"tags"`
	Tenants        []string          `pulumi:"tenants"`
	Type           string            `pulumi:"type"`
}
