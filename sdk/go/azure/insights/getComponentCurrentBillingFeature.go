


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComponentCurrentBillingFeature(ctx *pulumi.Context, args *LookupComponentCurrentBillingFeatureArgs, opts ...pulumi.InvokeOption) (*LookupComponentCurrentBillingFeatureResult, error) {
	var rv LookupComponentCurrentBillingFeatureResult
	err := ctx.Invoke("azure-native:insights:getComponentCurrentBillingFeature", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComponentCurrentBillingFeatureArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupComponentCurrentBillingFeatureResult struct {
	CurrentBillingFeatures []string                                           `pulumi:"currentBillingFeatures"`
	DataVolumeCap          *ApplicationInsightsComponentDataVolumeCapResponse `pulumi:"dataVolumeCap"`
}
