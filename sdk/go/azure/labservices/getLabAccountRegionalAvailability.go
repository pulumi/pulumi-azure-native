


package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLabAccountRegionalAvailability(ctx *pulumi.Context, args *GetLabAccountRegionalAvailabilityArgs, opts ...pulumi.InvokeOption) (*GetLabAccountRegionalAvailabilityResult, error) {
	var rv GetLabAccountRegionalAvailabilityResult
	err := ctx.Invoke("azure-native:labservices:getLabAccountRegionalAvailability", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLabAccountRegionalAvailabilityArgs struct {
	LabAccountName    string `pulumi:"labAccountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetLabAccountRegionalAvailabilityResult struct {
	RegionalAvailability []RegionalAvailabilityResponse `pulumi:"regionalAvailability"`
}
