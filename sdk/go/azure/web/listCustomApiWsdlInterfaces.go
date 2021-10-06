


package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListCustomApiWsdlInterfaces(ctx *pulumi.Context, args *ListCustomApiWsdlInterfacesArgs, opts ...pulumi.InvokeOption) (*ListCustomApiWsdlInterfacesResult, error) {
	var rv ListCustomApiWsdlInterfacesResult
	err := ctx.Invoke("azure-native:web:listCustomApiWsdlInterfaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListCustomApiWsdlInterfacesArgs struct {
	Content        *string      `pulumi:"content"`
	ImportMethod   *string      `pulumi:"importMethod"`
	Location       string       `pulumi:"location"`
	Service        *WsdlService `pulumi:"service"`
	SubscriptionId *string      `pulumi:"subscriptionId"`
	Url            *string      `pulumi:"url"`
}


type ListCustomApiWsdlInterfacesResult struct {
	Value []WsdlServiceResponse `pulumi:"value"`
}
