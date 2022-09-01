


package v20170426

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPredictionModelStatus(ctx *pulumi.Context, args *GetPredictionModelStatusArgs, opts ...pulumi.InvokeOption) (*GetPredictionModelStatusResult, error) {
	var rv GetPredictionModelStatusResult
	err := ctx.Invoke("azure-native:customerinsights/v20170426:getPredictionModelStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetPredictionModelStatusArgs struct {
	HubName           string `pulumi:"hubName"`
	PredictionName    string `pulumi:"predictionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetPredictionModelStatusResult struct {
	Message            string `pulumi:"message"`
	ModelVersion       string `pulumi:"modelVersion"`
	PredictionGuidId   string `pulumi:"predictionGuidId"`
	PredictionName     string `pulumi:"predictionName"`
	SignalsUsed        int    `pulumi:"signalsUsed"`
	Status             string `pulumi:"status"`
	TenantId           string `pulumi:"tenantId"`
	TestSetCount       int    `pulumi:"testSetCount"`
	TrainingAccuracy   int    `pulumi:"trainingAccuracy"`
	TrainingSetCount   int    `pulumi:"trainingSetCount"`
	ValidationSetCount int    `pulumi:"validationSetCount"`
}

func GetPredictionModelStatusOutput(ctx *pulumi.Context, args GetPredictionModelStatusOutputArgs, opts ...pulumi.InvokeOption) GetPredictionModelStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPredictionModelStatusResult, error) {
			args := v.(GetPredictionModelStatusArgs)
			r, err := GetPredictionModelStatus(ctx, &args, opts...)
			var s GetPredictionModelStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPredictionModelStatusResultOutput)
}

type GetPredictionModelStatusOutputArgs struct {
	HubName           pulumi.StringInput `pulumi:"hubName"`
	PredictionName    pulumi.StringInput `pulumi:"predictionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetPredictionModelStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPredictionModelStatusArgs)(nil)).Elem()
}


type GetPredictionModelStatusResultOutput struct{ *pulumi.OutputState }

func (GetPredictionModelStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPredictionModelStatusResult)(nil)).Elem()
}

func (o GetPredictionModelStatusResultOutput) ToGetPredictionModelStatusResultOutput() GetPredictionModelStatusResultOutput {
	return o
}

func (o GetPredictionModelStatusResultOutput) ToGetPredictionModelStatusResultOutputWithContext(ctx context.Context) GetPredictionModelStatusResultOutput {
	return o
}

func (o GetPredictionModelStatusResultOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v GetPredictionModelStatusResult) string { return v.Message }).(pulumi.StringOutput)
}

func (o GetPredictionModelStatusResultOutput) ModelVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetPredictionModelStatusResult) string { return v.ModelVersion }).(pulumi.StringOutput)
}

func (o GetPredictionModelStatusResultOutput) PredictionGuidId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPredictionModelStatusResult) string { return v.PredictionGuidId }).(pulumi.StringOutput)
}

func (o GetPredictionModelStatusResultOutput) PredictionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetPredictionModelStatusResult) string { return v.PredictionName }).(pulumi.StringOutput)
}

func (o GetPredictionModelStatusResultOutput) SignalsUsed() pulumi.IntOutput {
	return o.ApplyT(func(v GetPredictionModelStatusResult) int { return v.SignalsUsed }).(pulumi.IntOutput)
}

func (o GetPredictionModelStatusResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetPredictionModelStatusResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetPredictionModelStatusResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPredictionModelStatusResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o GetPredictionModelStatusResultOutput) TestSetCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetPredictionModelStatusResult) int { return v.TestSetCount }).(pulumi.IntOutput)
}

func (o GetPredictionModelStatusResultOutput) TrainingAccuracy() pulumi.IntOutput {
	return o.ApplyT(func(v GetPredictionModelStatusResult) int { return v.TrainingAccuracy }).(pulumi.IntOutput)
}

func (o GetPredictionModelStatusResultOutput) TrainingSetCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetPredictionModelStatusResult) int { return v.TrainingSetCount }).(pulumi.IntOutput)
}

func (o GetPredictionModelStatusResultOutput) ValidationSetCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetPredictionModelStatusResult) int { return v.ValidationSetCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPredictionModelStatusResultOutput{})
}
