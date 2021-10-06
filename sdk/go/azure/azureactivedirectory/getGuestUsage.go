


package azureactivedirectory

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGuestUsage(ctx *pulumi.Context, args *LookupGuestUsageArgs, opts ...pulumi.InvokeOption) (*LookupGuestUsageResult, error) {
	var rv LookupGuestUsageResult
	err := ctx.Invoke("azure-native:azureactivedirectory:getGuestUsage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGuestUsageArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupGuestUsageResult struct {
	Id       string            `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	TenantId *string           `pulumi:"tenantId"`
	Type     string            `pulumi:"type"`
}
