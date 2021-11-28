


package v20211201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAddressByName(ctx *pulumi.Context, args *LookupAddressByNameArgs, opts ...pulumi.InvokeOption) (*LookupAddressByNameResult, error) {
	var rv LookupAddressByNameResult
	err := ctx.Invoke("azure-native:edgeorder/v20211201:getAddressByName", args, &rv, opts...)
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
