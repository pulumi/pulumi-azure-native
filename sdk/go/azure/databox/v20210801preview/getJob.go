


package v20210801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:databox/v20210801preview:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	Expand            *string `pulumi:"expand"`
	JobName           string  `pulumi:"jobName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupJobResult struct {
	CancellationReason        string                    `pulumi:"cancellationReason"`
	DeliveryInfo              *JobDeliveryInfoResponse  `pulumi:"deliveryInfo"`
	DeliveryType              *string                   `pulumi:"deliveryType"`
	Details                   interface{}               `pulumi:"details"`
	Error                     CloudErrorResponse        `pulumi:"error"`
	Id                        string                    `pulumi:"id"`
	Identity                  *ResourceIdentityResponse `pulumi:"identity"`
	IsCancellable             bool                      `pulumi:"isCancellable"`
	IsCancellableWithoutFee   bool                      `pulumi:"isCancellableWithoutFee"`
	IsDeletable               bool                      `pulumi:"isDeletable"`
	IsPrepareToShipEnabled    bool                      `pulumi:"isPrepareToShipEnabled"`
	IsShippingAddressEditable bool                      `pulumi:"isShippingAddressEditable"`
	Location                  string                    `pulumi:"location"`
	Name                      string                    `pulumi:"name"`
	Sku                       SkuResponse               `pulumi:"sku"`
	StartTime                 string                    `pulumi:"startTime"`
	Status                    string                    `pulumi:"status"`
	SystemData                SystemDataResponse        `pulumi:"systemData"`
	Tags                      map[string]string         `pulumi:"tags"`
	TransferType              string                    `pulumi:"transferType"`
	Type                      string                    `pulumi:"type"`
}
