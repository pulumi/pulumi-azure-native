


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrder(ctx *pulumi.Context, args *LookupOrderArgs, opts ...pulumi.InvokeOption) (*LookupOrderResult, error) {
	var rv LookupOrderResult
	err := ctx.Invoke("azure-native:databoxedge/v20210201:getOrder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrderArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOrderResult struct {
	ContactInformation   ContactDetailsResponse `pulumi:"contactInformation"`
	CurrentStatus        OrderStatusResponse    `pulumi:"currentStatus"`
	DeliveryTrackingInfo []TrackingInfoResponse `pulumi:"deliveryTrackingInfo"`
	Id                   string                 `pulumi:"id"`
	Name                 string                 `pulumi:"name"`
	OrderHistory         []OrderStatusResponse  `pulumi:"orderHistory"`
	ReturnTrackingInfo   []TrackingInfoResponse `pulumi:"returnTrackingInfo"`
	SerialNumber         string                 `pulumi:"serialNumber"`
	ShipmentType         *string                `pulumi:"shipmentType"`
	ShippingAddress      *AddressResponse       `pulumi:"shippingAddress"`
	SystemData           SystemDataResponse     `pulumi:"systemData"`
	Type                 string                 `pulumi:"type"`
}
