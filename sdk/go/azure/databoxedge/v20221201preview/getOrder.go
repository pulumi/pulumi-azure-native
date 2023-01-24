


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrder(ctx *pulumi.Context, args *LookupOrderArgs, opts ...pulumi.InvokeOption) (*LookupOrderResult, error) {
	var rv LookupOrderResult
	err := ctx.Invoke("azure-native:databoxedge/v20221201preview:getOrder", args, &rv, opts...)
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
	Kind                 string                 `pulumi:"kind"`
	Name                 string                 `pulumi:"name"`
	OrderHistory         []OrderStatusResponse  `pulumi:"orderHistory"`
	OrderId              string                 `pulumi:"orderId"`
	ReturnTrackingInfo   []TrackingInfoResponse `pulumi:"returnTrackingInfo"`
	SerialNumber         string                 `pulumi:"serialNumber"`
	ShipmentType         *string                `pulumi:"shipmentType"`
	ShippingAddress      *AddressResponse       `pulumi:"shippingAddress"`
	SystemData           SystemDataResponse     `pulumi:"systemData"`
	Type                 string                 `pulumi:"type"`
}

func LookupOrderOutput(ctx *pulumi.Context, args LookupOrderOutputArgs, opts ...pulumi.InvokeOption) LookupOrderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrderResult, error) {
			args := v.(LookupOrderArgs)
			r, err := LookupOrder(ctx, &args, opts...)
			var s LookupOrderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrderResultOutput)
}

type LookupOrderOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOrderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderArgs)(nil)).Elem()
}


type LookupOrderResultOutput struct{ *pulumi.OutputState }

func (LookupOrderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderResult)(nil)).Elem()
}

func (o LookupOrderResultOutput) ToLookupOrderResultOutput() LookupOrderResultOutput {
	return o
}

func (o LookupOrderResultOutput) ToLookupOrderResultOutputWithContext(ctx context.Context) LookupOrderResultOutput {
	return o
}

func (o LookupOrderResultOutput) ContactInformation() ContactDetailsResponseOutput {
	return o.ApplyT(func(v LookupOrderResult) ContactDetailsResponse { return v.ContactInformation }).(ContactDetailsResponseOutput)
}

func (o LookupOrderResultOutput) CurrentStatus() OrderStatusResponseOutput {
	return o.ApplyT(func(v LookupOrderResult) OrderStatusResponse { return v.CurrentStatus }).(OrderStatusResponseOutput)
}

func (o LookupOrderResultOutput) DeliveryTrackingInfo() TrackingInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupOrderResult) []TrackingInfoResponse { return v.DeliveryTrackingInfo }).(TrackingInfoResponseArrayOutput)
}

func (o LookupOrderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrderResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupOrderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOrderResultOutput) OrderHistory() OrderStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupOrderResult) []OrderStatusResponse { return v.OrderHistory }).(OrderStatusResponseArrayOutput)
}

func (o LookupOrderResultOutput) OrderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.OrderId }).(pulumi.StringOutput)
}

func (o LookupOrderResultOutput) ReturnTrackingInfo() TrackingInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupOrderResult) []TrackingInfoResponse { return v.ReturnTrackingInfo }).(TrackingInfoResponseArrayOutput)
}

func (o LookupOrderResultOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o LookupOrderResultOutput) ShipmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrderResult) *string { return v.ShipmentType }).(pulumi.StringPtrOutput)
}

func (o LookupOrderResultOutput) ShippingAddress() AddressResponsePtrOutput {
	return o.ApplyT(func(v LookupOrderResult) *AddressResponse { return v.ShippingAddress }).(AddressResponsePtrOutput)
}

func (o LookupOrderResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOrderResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOrderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrderResultOutput{})
}
