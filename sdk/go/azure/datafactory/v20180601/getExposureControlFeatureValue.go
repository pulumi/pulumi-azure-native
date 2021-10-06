


package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetExposureControlFeatureValue(ctx *pulumi.Context, args *GetExposureControlFeatureValueArgs, opts ...pulumi.InvokeOption) (*GetExposureControlFeatureValueResult, error) {
	var rv GetExposureControlFeatureValueResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getExposureControlFeatureValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetExposureControlFeatureValueArgs struct {
	FeatureName *string `pulumi:"featureName"`
	FeatureType *string `pulumi:"featureType"`
	LocationId  string  `pulumi:"locationId"`
}


type GetExposureControlFeatureValueResult struct {
	FeatureName string `pulumi:"featureName"`
	Value       string `pulumi:"value"`
}
