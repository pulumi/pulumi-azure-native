


package v20190901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:databox/v20190901:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupJobArgs struct {
	Expand            *string `pulumi:"expand"`
	JobName           string  `pulumi:"jobName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupJobResult struct {
	CancellationReason        string                   `pulumi:"cancellationReason"`
	DeliveryInfo              *JobDeliveryInfoResponse `pulumi:"deliveryInfo"`
	DeliveryType              *string                  `pulumi:"deliveryType"`
	Details                   interface{}              `pulumi:"details"`
	Error                     ErrorResponse            `pulumi:"error"`
	Id                        string                   `pulumi:"id"`
	IsCancellable             bool                     `pulumi:"isCancellable"`
	IsCancellableWithoutFee   bool                     `pulumi:"isCancellableWithoutFee"`
	IsDeletable               bool                     `pulumi:"isDeletable"`
	IsShippingAddressEditable bool                     `pulumi:"isShippingAddressEditable"`
	Location                  string                   `pulumi:"location"`
	Name                      string                   `pulumi:"name"`
	Sku                       SkuResponse              `pulumi:"sku"`
	StartTime                 string                   `pulumi:"startTime"`
	Status                    string                   `pulumi:"status"`
	Tags                      map[string]string        `pulumi:"tags"`
	Type                      string                   `pulumi:"type"`
}


func (val *LookupJobResult) Defaults() *LookupJobResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DeliveryType) {
		deliveryType_ := "NonScheduled"
		tmp.DeliveryType = &deliveryType_
	}
	return &tmp
}
