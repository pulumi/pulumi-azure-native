


package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegisteredPrefix(ctx *pulumi.Context, args *LookupRegisteredPrefixArgs, opts ...pulumi.InvokeOption) (*LookupRegisteredPrefixResult, error) {
	var rv LookupRegisteredPrefixResult
	err := ctx.Invoke("azure-native:peering/v20200401:getRegisteredPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegisteredPrefixArgs struct {
	PeeringName          string `pulumi:"peeringName"`
	RegisteredPrefixName string `pulumi:"registeredPrefixName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupRegisteredPrefixResult struct {
	ErrorMessage            string  `pulumi:"errorMessage"`
	Id                      string  `pulumi:"id"`
	Name                    string  `pulumi:"name"`
	PeeringServicePrefixKey string  `pulumi:"peeringServicePrefixKey"`
	Prefix                  *string `pulumi:"prefix"`
	PrefixValidationState   string  `pulumi:"prefixValidationState"`
	ProvisioningState       string  `pulumi:"provisioningState"`
	Type                    string  `pulumi:"type"`
}
