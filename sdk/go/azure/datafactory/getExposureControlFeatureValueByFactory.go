


package datafactory

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetExposureControlFeatureValueByFactory(ctx *pulumi.Context, args *GetExposureControlFeatureValueByFactoryArgs, opts ...pulumi.InvokeOption) (*GetExposureControlFeatureValueByFactoryResult, error) {
	var rv GetExposureControlFeatureValueByFactoryResult
	err := ctx.Invoke("azure-native:datafactory:getExposureControlFeatureValueByFactory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetExposureControlFeatureValueByFactoryArgs struct {
	FactoryName       string  `pulumi:"factoryName"`
	FeatureName       *string `pulumi:"featureName"`
	FeatureType       *string `pulumi:"featureType"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type GetExposureControlFeatureValueByFactoryResult struct {
	FeatureName string `pulumi:"featureName"`
	Value       string `pulumi:"value"`
}
