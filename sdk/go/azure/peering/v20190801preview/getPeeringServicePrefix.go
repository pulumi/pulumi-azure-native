


package v20190801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPeeringServicePrefix(ctx *pulumi.Context, args *LookupPeeringServicePrefixArgs, opts ...pulumi.InvokeOption) (*LookupPeeringServicePrefixResult, error) {
	var rv LookupPeeringServicePrefixResult
	err := ctx.Invoke("azure-native:peering/v20190801preview:getPeeringServicePrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPeeringServicePrefixArgs struct {
	PeeringServiceName string `pulumi:"peeringServiceName"`
	PrefixName         string `pulumi:"prefixName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupPeeringServicePrefixResult struct {
	Id                    string  `pulumi:"id"`
	LearnedType           *string `pulumi:"learnedType"`
	Name                  string  `pulumi:"name"`
	Prefix                *string `pulumi:"prefix"`
	PrefixValidationState *string `pulumi:"prefixValidationState"`
	ProvisioningState     string  `pulumi:"provisioningState"`
	Type                  string  `pulumi:"type"`
}
