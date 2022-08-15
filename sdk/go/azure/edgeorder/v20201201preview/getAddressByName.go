


package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupAddressByName(ctx *pulumi.Context, args *LookupAddressByNameArgs, opts ...pulumi.InvokeOption) (*LookupAddressByNameResult, error) {
	var rv LookupAddressByNameResult
	err := ctx.Invoke("azure-native:edgeorder/v20201201preview:getAddressByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAddressByNameArgs struct {
	AddressName       string `pulumi:"addressName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAddressByNameResult struct {
	ContactDetails  ContactDetailsResponse   `pulumi:"contactDetails"`
	Id              string                   `pulumi:"id"`
	Location        string                   `pulumi:"location"`
	Name            string                   `pulumi:"name"`
	ShippingAddress *ShippingAddressResponse `pulumi:"shippingAddress"`
	SystemData      SystemDataResponse       `pulumi:"systemData"`
	Tags            map[string]string        `pulumi:"tags"`
	Type            string                   `pulumi:"type"`
}

func LookupAddressByNameOutput(ctx *pulumi.Context, args LookupAddressByNameOutputArgs, opts ...pulumi.InvokeOption) LookupAddressByNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAddressByNameResult, error) {
			args := v.(LookupAddressByNameArgs)
			r, err := LookupAddressByName(ctx, &args, opts...)
			var s LookupAddressByNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAddressByNameResultOutput)
}

type LookupAddressByNameOutputArgs struct {
	AddressName       pulumi.StringInput `pulumi:"addressName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAddressByNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddressByNameArgs)(nil)).Elem()
}


type LookupAddressByNameResultOutput struct{ *pulumi.OutputState }

func (LookupAddressByNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddressByNameResult)(nil)).Elem()
}

func (o LookupAddressByNameResultOutput) ToLookupAddressByNameResultOutput() LookupAddressByNameResultOutput {
	return o
}

func (o LookupAddressByNameResultOutput) ToLookupAddressByNameResultOutputWithContext(ctx context.Context) LookupAddressByNameResultOutput {
	return o
}

func (o LookupAddressByNameResultOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v LookupAddressByNameResult) ContactDetailsResponse { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o LookupAddressByNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressByNameResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAddressByNameResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressByNameResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAddressByNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressByNameResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAddressByNameResultOutput) ShippingAddress() ShippingAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupAddressByNameResult) *ShippingAddressResponse { return v.ShippingAddress }).(ShippingAddressResponsePtrOutput)
}

func (o LookupAddressByNameResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAddressByNameResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAddressByNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAddressByNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAddressByNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressByNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAddressByNameResultOutput{})
}
