


package azureactivedirectory

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupB2CTenant(ctx *pulumi.Context, args *LookupB2CTenantArgs, opts ...pulumi.InvokeOption) (*LookupB2CTenantResult, error) {
	var rv LookupB2CTenantResult
	err := ctx.Invoke("azure-native:azureactivedirectory:getB2CTenant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupB2CTenantArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}

type LookupB2CTenantResult struct {
	BillingConfig *B2CTenantResourcePropertiesResponseBillingConfig `pulumi:"billingConfig"`
	Id            string                                            `pulumi:"id"`
	Location      string                                            `pulumi:"location"`
	Name          string                                            `pulumi:"name"`
	Sku           B2CResourceSKUResponse                            `pulumi:"sku"`
	Tags          map[string]string                                 `pulumi:"tags"`
	TenantId      *string                                           `pulumi:"tenantId"`
	Type          string                                            `pulumi:"type"`
}
