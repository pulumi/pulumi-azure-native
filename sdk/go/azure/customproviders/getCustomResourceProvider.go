


package customproviders

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomResourceProvider(ctx *pulumi.Context, args *LookupCustomResourceProviderArgs, opts ...pulumi.InvokeOption) (*LookupCustomResourceProviderResult, error) {
	var rv LookupCustomResourceProviderResult
	err := ctx.Invoke("azure-native:customproviders:getCustomResourceProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomResourceProviderArgs struct {
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	ResourceProviderName string `pulumi:"resourceProviderName"`
}


type LookupCustomResourceProviderResult struct {
	Actions           []CustomRPActionRouteDefinitionResponse       `pulumi:"actions"`
	Id                string                                        `pulumi:"id"`
	Location          string                                        `pulumi:"location"`
	Name              string                                        `pulumi:"name"`
	ProvisioningState string                                        `pulumi:"provisioningState"`
	ResourceTypes     []CustomRPResourceTypeRouteDefinitionResponse `pulumi:"resourceTypes"`
	Tags              map[string]string                             `pulumi:"tags"`
	Type              string                                        `pulumi:"type"`
	Validations       []CustomRPValidationsResponse                 `pulumi:"validations"`
}
