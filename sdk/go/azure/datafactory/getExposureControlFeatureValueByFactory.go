


package datafactory

import (
	"context"
	"reflect"

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

func GetExposureControlFeatureValueByFactoryOutput(ctx *pulumi.Context, args GetExposureControlFeatureValueByFactoryOutputArgs, opts ...pulumi.InvokeOption) GetExposureControlFeatureValueByFactoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetExposureControlFeatureValueByFactoryResult, error) {
			args := v.(GetExposureControlFeatureValueByFactoryArgs)
			r, err := GetExposureControlFeatureValueByFactory(ctx, &args, opts...)
			var s GetExposureControlFeatureValueByFactoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetExposureControlFeatureValueByFactoryResultOutput)
}

type GetExposureControlFeatureValueByFactoryOutputArgs struct {
	FactoryName       pulumi.StringInput    `pulumi:"factoryName"`
	FeatureName       pulumi.StringPtrInput `pulumi:"featureName"`
	FeatureType       pulumi.StringPtrInput `pulumi:"featureType"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (GetExposureControlFeatureValueByFactoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExposureControlFeatureValueByFactoryArgs)(nil)).Elem()
}


type GetExposureControlFeatureValueByFactoryResultOutput struct{ *pulumi.OutputState }

func (GetExposureControlFeatureValueByFactoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExposureControlFeatureValueByFactoryResult)(nil)).Elem()
}

func (o GetExposureControlFeatureValueByFactoryResultOutput) ToGetExposureControlFeatureValueByFactoryResultOutput() GetExposureControlFeatureValueByFactoryResultOutput {
	return o
}

func (o GetExposureControlFeatureValueByFactoryResultOutput) ToGetExposureControlFeatureValueByFactoryResultOutputWithContext(ctx context.Context) GetExposureControlFeatureValueByFactoryResultOutput {
	return o
}

func (o GetExposureControlFeatureValueByFactoryResultOutput) FeatureName() pulumi.StringOutput {
	return o.ApplyT(func(v GetExposureControlFeatureValueByFactoryResult) string { return v.FeatureName }).(pulumi.StringOutput)
}

func (o GetExposureControlFeatureValueByFactoryResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetExposureControlFeatureValueByFactoryResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetExposureControlFeatureValueByFactoryResultOutput{})
}
