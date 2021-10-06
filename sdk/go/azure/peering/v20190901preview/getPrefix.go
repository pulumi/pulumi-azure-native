


package v20190901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrefix(ctx *pulumi.Context, args *LookupPrefixArgs, opts ...pulumi.InvokeOption) (*LookupPrefixResult, error) {
	var rv LookupPrefixResult
	err := ctx.Invoke("azure-native:peering/v20190901preview:getPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrefixArgs struct {
	Expand             *string `pulumi:"expand"`
	PeeringServiceName string  `pulumi:"peeringServiceName"`
	PrefixName         string  `pulumi:"prefixName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupPrefixResult struct {
	ErrorMessage          string                              `pulumi:"errorMessage"`
	Events                []PeeringServicePrefixEventResponse `pulumi:"events"`
	Id                    string                              `pulumi:"id"`
	LearnedType           string                              `pulumi:"learnedType"`
	Name                  string                              `pulumi:"name"`
	Prefix                *string                             `pulumi:"prefix"`
	PrefixValidationState string                              `pulumi:"prefixValidationState"`
	ProvisioningState     string                              `pulumi:"provisioningState"`
	Type                  string                              `pulumi:"type"`
}
