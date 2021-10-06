


package v20190101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProduct(ctx *pulumi.Context, args *LookupProductArgs, opts ...pulumi.InvokeOption) (*LookupProductResult, error) {
	var rv LookupProductResult
	err := ctx.Invoke("azure-native:apimanagement/v20190101:getProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProductArgs struct {
	ProductId         string `pulumi:"productId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupProductResult struct {
	ApprovalRequired     *bool   `pulumi:"approvalRequired"`
	Description          *string `pulumi:"description"`
	DisplayName          string  `pulumi:"displayName"`
	Id                   string  `pulumi:"id"`
	Name                 string  `pulumi:"name"`
	State                *string `pulumi:"state"`
	SubscriptionRequired *bool   `pulumi:"subscriptionRequired"`
	SubscriptionsLimit   *int    `pulumi:"subscriptionsLimit"`
	Terms                *string `pulumi:"terms"`
	Type                 string  `pulumi:"type"`
}
