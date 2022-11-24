


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrderItem(ctx *pulumi.Context, args *LookupOrderItemArgs, opts ...pulumi.InvokeOption) (*LookupOrderItemResult, error) {
	var rv LookupOrderItemResult
	err := ctx.Invoke("azure-native:edgeorder/v20220501preview:getOrderItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrderItemArgs struct {
	Expand            *string `pulumi:"expand"`
	OrderItemName     string  `pulumi:"orderItemName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupOrderItemResult struct {
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

func LookupOrderItemOutput(ctx *pulumi.Context, args LookupOrderItemOutputArgs, opts ...pulumi.InvokeOption) LookupOrderItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrderItemResult, error) {
			args := v.(LookupOrderItemArgs)
			r, err := LookupOrderItem(ctx, &args, opts...)
			var s LookupOrderItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrderItemResultOutput)
}

type LookupOrderItemOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	OrderItemName     pulumi.StringInput    `pulumi:"orderItemName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupOrderItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderItemArgs)(nil)).Elem()
}


type LookupOrderItemResultOutput struct{ *pulumi.OutputState }

func (LookupOrderItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderItemResult)(nil)).Elem()
}

func (o LookupOrderItemResultOutput) ToLookupOrderItemResultOutput() LookupOrderItemResultOutput {
	return o
}

func (o LookupOrderItemResultOutput) ToLookupOrderItemResultOutputWithContext(ctx context.Context) LookupOrderItemResultOutput {
	return o
}

func (o LookupOrderItemResultOutput) AddressDetails() AddressDetailsResponseOutput {
	return o.ApplyT(func(v LookupOrderItemResult) AddressDetailsResponse { return v.AddressDetails }).(AddressDetailsResponseOutput)
}

func (o LookupOrderItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrderItemResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupOrderItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOrderItemResultOutput) OrderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.OrderId }).(pulumi.StringOutput)
}

func (o LookupOrderItemResultOutput) OrderItemDetails() OrderItemDetailsResponseOutput {
	return o.ApplyT(func(v LookupOrderItemResult) OrderItemDetailsResponse { return v.OrderItemDetails }).(OrderItemDetailsResponseOutput)
}

func (o LookupOrderItemResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o LookupOrderItemResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOrderItemResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOrderItemResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOrderItemResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupOrderItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrderItemResultOutput{})
}
