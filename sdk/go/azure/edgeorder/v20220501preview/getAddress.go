


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAddress(ctx *pulumi.Context, args *LookupAddressArgs, opts ...pulumi.InvokeOption) (*LookupAddressResult, error) {
	var rv LookupAddressResult
	err := ctx.Invoke("azure-native:edgeorder/v20220501preview:getAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAddressArgs struct {
	AddressName       string `pulumi:"addressName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAddressResult struct {
	AddressValidationStatus string                   `pulumi:"addressValidationStatus"`
	ContactDetails          ContactDetailsResponse   `pulumi:"contactDetails"`
	Id                      string                   `pulumi:"id"`
	Location                string                   `pulumi:"location"`
	Name                    string                   `pulumi:"name"`
	ShippingAddress         *ShippingAddressResponse `pulumi:"shippingAddress"`
	SystemData              SystemDataResponse       `pulumi:"systemData"`
	Tags                    map[string]string        `pulumi:"tags"`
	Type                    string                   `pulumi:"type"`
}

func LookupAddressOutput(ctx *pulumi.Context, args LookupAddressOutputArgs, opts ...pulumi.InvokeOption) LookupAddressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAddressResult, error) {
			args := v.(LookupAddressArgs)
			r, err := LookupAddress(ctx, &args, opts...)
			var s LookupAddressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAddressResultOutput)
}

type LookupAddressOutputArgs struct {
	AddressName       pulumi.StringInput `pulumi:"addressName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAddressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddressArgs)(nil)).Elem()
}


type LookupAddressResultOutput struct{ *pulumi.OutputState }

func (LookupAddressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddressResult)(nil)).Elem()
}

func (o LookupAddressResultOutput) ToLookupAddressResultOutput() LookupAddressResultOutput {
	return o
}

func (o LookupAddressResultOutput) ToLookupAddressResultOutputWithContext(ctx context.Context) LookupAddressResultOutput {
	return o
}

func (o LookupAddressResultOutput) AddressValidationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressResult) string { return v.AddressValidationStatus }).(pulumi.StringOutput)
}

func (o LookupAddressResultOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v LookupAddressResult) ContactDetailsResponse { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o LookupAddressResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAddressResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAddressResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAddressResultOutput) ShippingAddress() ShippingAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupAddressResult) *ShippingAddressResponse { return v.ShippingAddress }).(ShippingAddressResponsePtrOutput)
}

func (o LookupAddressResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAddressResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAddressResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAddressResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAddressResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAddressResultOutput{})
}
