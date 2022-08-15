


package edgeorder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrderItemByName(ctx *pulumi.Context, args *LookupOrderItemByNameArgs, opts ...pulumi.InvokeOption) (*LookupOrderItemByNameResult, error) {
	var rv LookupOrderItemByNameResult
	err := ctx.Invoke("azure-native:edgeorder:getOrderItemByName", args, &rv, opts...)
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

func LookupOrderItemByNameOutput(ctx *pulumi.Context, args LookupOrderItemByNameOutputArgs, opts ...pulumi.InvokeOption) LookupOrderItemByNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrderItemByNameResult, error) {
			args := v.(LookupOrderItemByNameArgs)
			r, err := LookupOrderItemByName(ctx, &args, opts...)
			var s LookupOrderItemByNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrderItemByNameResultOutput)
}

type LookupOrderItemByNameOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	OrderItemName     pulumi.StringInput    `pulumi:"orderItemName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupOrderItemByNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderItemByNameArgs)(nil)).Elem()
}


type LookupOrderItemByNameResultOutput struct{ *pulumi.OutputState }

func (LookupOrderItemByNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderItemByNameResult)(nil)).Elem()
}

func (o LookupOrderItemByNameResultOutput) ToLookupOrderItemByNameResultOutput() LookupOrderItemByNameResultOutput {
	return o
}

func (o LookupOrderItemByNameResultOutput) ToLookupOrderItemByNameResultOutputWithContext(ctx context.Context) LookupOrderItemByNameResultOutput {
	return o
}

func (o LookupOrderItemByNameResultOutput) AddressDetails() AddressDetailsResponseOutput {
	return o.ApplyT(func(v LookupOrderItemByNameResult) AddressDetailsResponse { return v.AddressDetails }).(AddressDetailsResponseOutput)
}

func (o LookupOrderItemByNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemByNameResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrderItemByNameResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemByNameResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupOrderItemByNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemByNameResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOrderItemByNameResultOutput) OrderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemByNameResult) string { return v.OrderId }).(pulumi.StringOutput)
}

func (o LookupOrderItemByNameResultOutput) OrderItemDetails() OrderItemDetailsResponseOutput {
	return o.ApplyT(func(v LookupOrderItemByNameResult) OrderItemDetailsResponse { return v.OrderItemDetails }).(OrderItemDetailsResponseOutput)
}

func (o LookupOrderItemByNameResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemByNameResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o LookupOrderItemByNameResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOrderItemByNameResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOrderItemByNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOrderItemByNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupOrderItemByNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemByNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrderItemByNameResultOutput{})
}
