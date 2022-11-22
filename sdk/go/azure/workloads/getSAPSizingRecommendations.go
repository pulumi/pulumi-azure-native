


package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSAPSizingRecommendations(ctx *pulumi.Context, args *GetSAPSizingRecommendationsArgs, opts ...pulumi.InvokeOption) (*GetSAPSizingRecommendationsResult, error) {
	var rv GetSAPSizingRecommendationsResult
	err := ctx.Invoke("azure-native:workloads:getSAPSizingRecommendations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSAPSizingRecommendationsArgs struct {
	AppLocation          string  `pulumi:"appLocation"`
	DatabaseType         string  `pulumi:"databaseType"`
	DbMemory             float64 `pulumi:"dbMemory"`
	DbScaleMethod        *string `pulumi:"dbScaleMethod"`
	DeploymentType       string  `pulumi:"deploymentType"`
	Environment          string  `pulumi:"environment"`
	HighAvailabilityType *string `pulumi:"highAvailabilityType"`
	Location             string  `pulumi:"location"`
	SapProduct           string  `pulumi:"sapProduct"`
	Saps                 float64 `pulumi:"saps"`
}


type GetSAPSizingRecommendationsResult struct {
	DeploymentType string `pulumi:"deploymentType"`
}

func GetSAPSizingRecommendationsOutput(ctx *pulumi.Context, args GetSAPSizingRecommendationsOutputArgs, opts ...pulumi.InvokeOption) GetSAPSizingRecommendationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSAPSizingRecommendationsResult, error) {
			args := v.(GetSAPSizingRecommendationsArgs)
			r, err := GetSAPSizingRecommendations(ctx, &args, opts...)
			var s GetSAPSizingRecommendationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSAPSizingRecommendationsResultOutput)
}

type GetSAPSizingRecommendationsOutputArgs struct {
	AppLocation          pulumi.StringInput    `pulumi:"appLocation"`
	DatabaseType         pulumi.StringInput    `pulumi:"databaseType"`
	DbMemory             pulumi.Float64Input   `pulumi:"dbMemory"`
	DbScaleMethod        pulumi.StringPtrInput `pulumi:"dbScaleMethod"`
	DeploymentType       pulumi.StringInput    `pulumi:"deploymentType"`
	Environment          pulumi.StringInput    `pulumi:"environment"`
	HighAvailabilityType pulumi.StringPtrInput `pulumi:"highAvailabilityType"`
	Location             pulumi.StringInput    `pulumi:"location"`
	SapProduct           pulumi.StringInput    `pulumi:"sapProduct"`
	Saps                 pulumi.Float64Input   `pulumi:"saps"`
}

func (GetSAPSizingRecommendationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPSizingRecommendationsArgs)(nil)).Elem()
}


type GetSAPSizingRecommendationsResultOutput struct{ *pulumi.OutputState }

func (GetSAPSizingRecommendationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPSizingRecommendationsResult)(nil)).Elem()
}

func (o GetSAPSizingRecommendationsResultOutput) ToGetSAPSizingRecommendationsResultOutput() GetSAPSizingRecommendationsResultOutput {
	return o
}

func (o GetSAPSizingRecommendationsResultOutput) ToGetSAPSizingRecommendationsResultOutputWithContext(ctx context.Context) GetSAPSizingRecommendationsResultOutput {
	return o
}

func (o GetSAPSizingRecommendationsResultOutput) DeploymentType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSAPSizingRecommendationsResult) string { return v.DeploymentType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSAPSizingRecommendationsResultOutput{})
}
