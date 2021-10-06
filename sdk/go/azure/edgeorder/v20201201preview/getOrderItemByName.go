


package v20201201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrderItemByName(ctx *pulumi.Context, args *LookupOrderItemByNameArgs, opts ...pulumi.InvokeOption) (*LookupOrderItemByNameResult, error) {
	var rv LookupOrderItemByNameResult
	err := ctx.Invoke("azure-native:edgeorder/v20201201preview:getOrderItemByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrderItemByNameArgs struct {
	Expand            *string `pulumi:"expand"`
	OrderItemName     string  `pulumi:"orderItemName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupOrderItemByNameResult struct {
	AddressDetails   AddressDetailsResponse   `pulumi:"addressDetails"`
	Id               string                   `pulumi:"id"`
	Location         string                   `pulumi:"location"`
	Name             string                   `pulumi:"name"`
	OrderId          string                   `pulumi:"orderId"`
	OrderItemDetails OrderItemDetailsResponse `pulumi:"orderItemDetails"`
	StartTime        string                   `pulumi:"startTime"`
	SystemData       SystemDataResponse       `pulumi:"systemData"`
	Tags             map[string]string        `pulumi:"tags"`
	Type             string                   `pulumi:"type"`
}
