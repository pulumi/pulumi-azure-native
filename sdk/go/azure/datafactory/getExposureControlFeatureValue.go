


package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetExposureControlFeatureValue(ctx *pulumi.Context, args *GetExposureControlFeatureValueArgs, opts ...pulumi.InvokeOption) (*GetExposureControlFeatureValueResult, error) {
	var rv GetExposureControlFeatureValueResult
	err := ctx.Invoke("azure-native:datafactory:getExposureControlFeatureValue", args, &rv, opts...)
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

func GetExposureControlFeatureValueOutput(ctx *pulumi.Context, args GetExposureControlFeatureValueOutputArgs, opts ...pulumi.InvokeOption) GetExposureControlFeatureValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetExposureControlFeatureValueResult, error) {
			args := v.(GetExposureControlFeatureValueArgs)
			r, err := GetExposureControlFeatureValue(ctx, &args, opts...)
			var s GetExposureControlFeatureValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetExposureControlFeatureValueResultOutput)
}

type GetExposureControlFeatureValueOutputArgs struct {
	FeatureName pulumi.StringPtrInput `pulumi:"featureName"`
	FeatureType pulumi.StringPtrInput `pulumi:"featureType"`
	LocationId  pulumi.StringInput    `pulumi:"locationId"`
}

func (GetExposureControlFeatureValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExposureControlFeatureValueArgs)(nil)).Elem()
}


type GetExposureControlFeatureValueResultOutput struct{ *pulumi.OutputState }

func (GetExposureControlFeatureValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExposureControlFeatureValueResult)(nil)).Elem()
}

func (o GetExposureControlFeatureValueResultOutput) ToGetExposureControlFeatureValueResultOutput() GetExposureControlFeatureValueResultOutput {
	return o
}

func (o GetExposureControlFeatureValueResultOutput) ToGetExposureControlFeatureValueResultOutputWithContext(ctx context.Context) GetExposureControlFeatureValueResultOutput {
	return o
}

func (o GetExposureControlFeatureValueResultOutput) FeatureName() pulumi.StringOutput {
	return o.ApplyT(func(v GetExposureControlFeatureValueResult) string { return v.FeatureName }).(pulumi.StringOutput)
}

func (o GetExposureControlFeatureValueResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetExposureControlFeatureValueResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetExposureControlFeatureValueResultOutput{})
}
