


package v20140401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OperationImpactResponse struct {
	ChangeValueAbsolute float64 `pulumi:"changeValueAbsolute"`
	ChangeValueRelative float64 `pulumi:"changeValueRelative"`
	Name                string  `pulumi:"name"`
	Unit                string  `pulumi:"unit"`
}

type OperationImpactResponseOutput struct{ *pulumi.OutputState }

func (OperationImpactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationImpactResponse)(nil)).Elem()
}

func (o OperationImpactResponseOutput) ToOperationImpactResponseOutput() OperationImpactResponseOutput {
	return o
}

func (o OperationImpactResponseOutput) ToOperationImpactResponseOutputWithContext(ctx context.Context) OperationImpactResponseOutput {
	return o
}

func (o OperationImpactResponseOutput) ChangeValueAbsolute() pulumi.Float64Output {
	return o.ApplyT(func(v OperationImpactResponse) float64 { return v.ChangeValueAbsolute }).(pulumi.Float64Output)
}

func (o OperationImpactResponseOutput) ChangeValueRelative() pulumi.Float64Output {
	return o.ApplyT(func(v OperationImpactResponse) float64 { return v.ChangeValueRelative }).(pulumi.Float64Output)
}

func (o OperationImpactResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v OperationImpactResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o OperationImpactResponseOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v OperationImpactResponse) string { return v.Unit }).(pulumi.StringOutput)
}

type OperationImpactResponseArrayOutput struct{ *pulumi.OutputState }

func (OperationImpactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OperationImpactResponse)(nil)).Elem()
}

func (o OperationImpactResponseArrayOutput) ToOperationImpactResponseArrayOutput() OperationImpactResponseArrayOutput {
	return o
}

func (o OperationImpactResponseArrayOutput) ToOperationImpactResponseArrayOutputWithContext(ctx context.Context) OperationImpactResponseArrayOutput {
	return o
}

func (o OperationImpactResponseArrayOutput) Index(i pulumi.IntInput) OperationImpactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OperationImpactResponse {
		return vs[0].([]OperationImpactResponse)[vs[1].(int)]
	}).(OperationImpactResponseOutput)
}

type RecommendedIndexResponse struct {
	Action          string                    `pulumi:"action"`
	Columns         []string                  `pulumi:"columns"`
	Created         string                    `pulumi:"created"`
	EstimatedImpact []OperationImpactResponse `pulumi:"estimatedImpact"`
	Id              string                    `pulumi:"id"`
	IncludedColumns []string                  `pulumi:"includedColumns"`
	IndexScript     string                    `pulumi:"indexScript"`
	IndexType       string                    `pulumi:"indexType"`
	LastModified    string                    `pulumi:"lastModified"`
	Name            string                    `pulumi:"name"`
	ReportedImpact  []OperationImpactResponse `pulumi:"reportedImpact"`
	Schema          string                    `pulumi:"schema"`
	State           string                    `pulumi:"state"`
	Table           string                    `pulumi:"table"`
	Type            string                    `pulumi:"type"`
}

type RecommendedIndexResponseOutput struct{ *pulumi.OutputState }

func (RecommendedIndexResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedIndexResponse)(nil)).Elem()
}

func (o RecommendedIndexResponseOutput) ToRecommendedIndexResponseOutput() RecommendedIndexResponseOutput {
	return o
}

func (o RecommendedIndexResponseOutput) ToRecommendedIndexResponseOutputWithContext(ctx context.Context) RecommendedIndexResponseOutput {
	return o
}

func (o RecommendedIndexResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o RecommendedIndexResponseOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

func (o RecommendedIndexResponseOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) string { return v.Created }).(pulumi.StringOutput)
}

func (o RecommendedIndexResponseOutput) EstimatedImpact() OperationImpactResponseArrayOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) []OperationImpactResponse { return v.EstimatedImpact }).(OperationImpactResponseArrayOutput)
}

func (o RecommendedIndexResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o RecommendedIndexResponseOutput) IncludedColumns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) []string { return v.IncludedColumns }).(pulumi.StringArrayOutput)
}

func (o RecommendedIndexResponseOutput) IndexScript() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) string { return v.IndexScript }).(pulumi.StringOutput)
}

func (o RecommendedIndexResponseOutput) IndexType() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) string { return v.IndexType }).(pulumi.StringOutput)
}

func (o RecommendedIndexResponseOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o RecommendedIndexResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RecommendedIndexResponseOutput) ReportedImpact() OperationImpactResponseArrayOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) []OperationImpactResponse { return v.ReportedImpact }).(OperationImpactResponseArrayOutput)
}

func (o RecommendedIndexResponseOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) string { return v.Schema }).(pulumi.StringOutput)
}

func (o RecommendedIndexResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o RecommendedIndexResponseOutput) Table() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) string { return v.Table }).(pulumi.StringOutput)
}

func (o RecommendedIndexResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedIndexResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RecommendedIndexResponseArrayOutput struct{ *pulumi.OutputState }

func (RecommendedIndexResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedIndexResponse)(nil)).Elem()
}

func (o RecommendedIndexResponseArrayOutput) ToRecommendedIndexResponseArrayOutput() RecommendedIndexResponseArrayOutput {
	return o
}

func (o RecommendedIndexResponseArrayOutput) ToRecommendedIndexResponseArrayOutputWithContext(ctx context.Context) RecommendedIndexResponseArrayOutput {
	return o
}

func (o RecommendedIndexResponseArrayOutput) Index(i pulumi.IntInput) RecommendedIndexResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecommendedIndexResponse {
		return vs[0].([]RecommendedIndexResponse)[vs[1].(int)]
	}).(RecommendedIndexResponseOutput)
}

type ServiceTierAdvisorResponse struct {
	ActiveTimeRatio                                        float64                  `pulumi:"activeTimeRatio"`
	AvgDtu                                                 float64                  `pulumi:"avgDtu"`
	Confidence                                             float64                  `pulumi:"confidence"`
	CurrentServiceLevelObjective                           string                   `pulumi:"currentServiceLevelObjective"`
	CurrentServiceLevelObjectiveId                         string                   `pulumi:"currentServiceLevelObjectiveId"`
	DatabaseSizeBasedRecommendationServiceLevelObjective   string                   `pulumi:"databaseSizeBasedRecommendationServiceLevelObjective"`
	DatabaseSizeBasedRecommendationServiceLevelObjectiveId string                   `pulumi:"databaseSizeBasedRecommendationServiceLevelObjectiveId"`
	DisasterPlanBasedRecommendationServiceLevelObjective   string                   `pulumi:"disasterPlanBasedRecommendationServiceLevelObjective"`
	DisasterPlanBasedRecommendationServiceLevelObjectiveId string                   `pulumi:"disasterPlanBasedRecommendationServiceLevelObjectiveId"`
	Id                                                     string                   `pulumi:"id"`
	MaxDtu                                                 float64                  `pulumi:"maxDtu"`
	MaxSizeInGB                                            float64                  `pulumi:"maxSizeInGB"`
	MinDtu                                                 float64                  `pulumi:"minDtu"`
	Name                                                   string                   `pulumi:"name"`
	ObservationPeriodEnd                                   string                   `pulumi:"observationPeriodEnd"`
	ObservationPeriodStart                                 string                   `pulumi:"observationPeriodStart"`
	OverallRecommendationServiceLevelObjective             string                   `pulumi:"overallRecommendationServiceLevelObjective"`
	OverallRecommendationServiceLevelObjectiveId           string                   `pulumi:"overallRecommendationServiceLevelObjectiveId"`
	ServiceLevelObjectiveUsageMetrics                      []SloUsageMetricResponse `pulumi:"serviceLevelObjectiveUsageMetrics"`
	Type                                                   string                   `pulumi:"type"`
	UsageBasedRecommendationServiceLevelObjective          string                   `pulumi:"usageBasedRecommendationServiceLevelObjective"`
	UsageBasedRecommendationServiceLevelObjectiveId        string                   `pulumi:"usageBasedRecommendationServiceLevelObjectiveId"`
}

type ServiceTierAdvisorResponseOutput struct{ *pulumi.OutputState }

func (ServiceTierAdvisorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTierAdvisorResponse)(nil)).Elem()
}

func (o ServiceTierAdvisorResponseOutput) ToServiceTierAdvisorResponseOutput() ServiceTierAdvisorResponseOutput {
	return o
}

func (o ServiceTierAdvisorResponseOutput) ToServiceTierAdvisorResponseOutputWithContext(ctx context.Context) ServiceTierAdvisorResponseOutput {
	return o
}

func (o ServiceTierAdvisorResponseOutput) ActiveTimeRatio() pulumi.Float64Output {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) float64 { return v.ActiveTimeRatio }).(pulumi.Float64Output)
}

func (o ServiceTierAdvisorResponseOutput) AvgDtu() pulumi.Float64Output {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) float64 { return v.AvgDtu }).(pulumi.Float64Output)
}

func (o ServiceTierAdvisorResponseOutput) Confidence() pulumi.Float64Output {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) float64 { return v.Confidence }).(pulumi.Float64Output)
}

func (o ServiceTierAdvisorResponseOutput) CurrentServiceLevelObjective() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string { return v.CurrentServiceLevelObjective }).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) CurrentServiceLevelObjectiveId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string { return v.CurrentServiceLevelObjectiveId }).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) DatabaseSizeBasedRecommendationServiceLevelObjective() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string {
		return v.DatabaseSizeBasedRecommendationServiceLevelObjective
	}).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) DatabaseSizeBasedRecommendationServiceLevelObjectiveId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string {
		return v.DatabaseSizeBasedRecommendationServiceLevelObjectiveId
	}).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) DisasterPlanBasedRecommendationServiceLevelObjective() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string {
		return v.DisasterPlanBasedRecommendationServiceLevelObjective
	}).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) DisasterPlanBasedRecommendationServiceLevelObjectiveId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string {
		return v.DisasterPlanBasedRecommendationServiceLevelObjectiveId
	}).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) MaxDtu() pulumi.Float64Output {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) float64 { return v.MaxDtu }).(pulumi.Float64Output)
}

func (o ServiceTierAdvisorResponseOutput) MaxSizeInGB() pulumi.Float64Output {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) float64 { return v.MaxSizeInGB }).(pulumi.Float64Output)
}

func (o ServiceTierAdvisorResponseOutput) MinDtu() pulumi.Float64Output {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) float64 { return v.MinDtu }).(pulumi.Float64Output)
}

func (o ServiceTierAdvisorResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) ObservationPeriodEnd() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string { return v.ObservationPeriodEnd }).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) ObservationPeriodStart() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string { return v.ObservationPeriodStart }).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) OverallRecommendationServiceLevelObjective() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string { return v.OverallRecommendationServiceLevelObjective }).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) OverallRecommendationServiceLevelObjectiveId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string { return v.OverallRecommendationServiceLevelObjectiveId }).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) ServiceLevelObjectiveUsageMetrics() SloUsageMetricResponseArrayOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) []SloUsageMetricResponse {
		return v.ServiceLevelObjectiveUsageMetrics
	}).(SloUsageMetricResponseArrayOutput)
}

func (o ServiceTierAdvisorResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) UsageBasedRecommendationServiceLevelObjective() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string { return v.UsageBasedRecommendationServiceLevelObjective }).(pulumi.StringOutput)
}

func (o ServiceTierAdvisorResponseOutput) UsageBasedRecommendationServiceLevelObjectiveId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTierAdvisorResponse) string { return v.UsageBasedRecommendationServiceLevelObjectiveId }).(pulumi.StringOutput)
}

type ServiceTierAdvisorResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceTierAdvisorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceTierAdvisorResponse)(nil)).Elem()
}

func (o ServiceTierAdvisorResponseArrayOutput) ToServiceTierAdvisorResponseArrayOutput() ServiceTierAdvisorResponseArrayOutput {
	return o
}

func (o ServiceTierAdvisorResponseArrayOutput) ToServiceTierAdvisorResponseArrayOutputWithContext(ctx context.Context) ServiceTierAdvisorResponseArrayOutput {
	return o
}

func (o ServiceTierAdvisorResponseArrayOutput) Index(i pulumi.IntInput) ServiceTierAdvisorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceTierAdvisorResponse {
		return vs[0].([]ServiceTierAdvisorResponse)[vs[1].(int)]
	}).(ServiceTierAdvisorResponseOutput)
}

type SloUsageMetricResponse struct {
	InRangeTimeRatio        float64 `pulumi:"inRangeTimeRatio"`
	ServiceLevelObjective   string  `pulumi:"serviceLevelObjective"`
	ServiceLevelObjectiveId string  `pulumi:"serviceLevelObjectiveId"`
}

type SloUsageMetricResponseOutput struct{ *pulumi.OutputState }

func (SloUsageMetricResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SloUsageMetricResponse)(nil)).Elem()
}

func (o SloUsageMetricResponseOutput) ToSloUsageMetricResponseOutput() SloUsageMetricResponseOutput {
	return o
}

func (o SloUsageMetricResponseOutput) ToSloUsageMetricResponseOutputWithContext(ctx context.Context) SloUsageMetricResponseOutput {
	return o
}

func (o SloUsageMetricResponseOutput) InRangeTimeRatio() pulumi.Float64Output {
	return o.ApplyT(func(v SloUsageMetricResponse) float64 { return v.InRangeTimeRatio }).(pulumi.Float64Output)
}

func (o SloUsageMetricResponseOutput) ServiceLevelObjective() pulumi.StringOutput {
	return o.ApplyT(func(v SloUsageMetricResponse) string { return v.ServiceLevelObjective }).(pulumi.StringOutput)
}

func (o SloUsageMetricResponseOutput) ServiceLevelObjectiveId() pulumi.StringOutput {
	return o.ApplyT(func(v SloUsageMetricResponse) string { return v.ServiceLevelObjectiveId }).(pulumi.StringOutput)
}

type SloUsageMetricResponseArrayOutput struct{ *pulumi.OutputState }

func (SloUsageMetricResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SloUsageMetricResponse)(nil)).Elem()
}

func (o SloUsageMetricResponseArrayOutput) ToSloUsageMetricResponseArrayOutput() SloUsageMetricResponseArrayOutput {
	return o
}

func (o SloUsageMetricResponseArrayOutput) ToSloUsageMetricResponseArrayOutputWithContext(ctx context.Context) SloUsageMetricResponseArrayOutput {
	return o
}

func (o SloUsageMetricResponseArrayOutput) Index(i pulumi.IntInput) SloUsageMetricResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SloUsageMetricResponse {
		return vs[0].([]SloUsageMetricResponse)[vs[1].(int)]
	}).(SloUsageMetricResponseOutput)
}

type TransparentDataEncryptionResponse struct {
	Id       string  `pulumi:"id"`
	Location string  `pulumi:"location"`
	Name     string  `pulumi:"name"`
	Status   *string `pulumi:"status"`
	Type     string  `pulumi:"type"`
}

type TransparentDataEncryptionResponseOutput struct{ *pulumi.OutputState }

func (TransparentDataEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransparentDataEncryptionResponse)(nil)).Elem()
}

func (o TransparentDataEncryptionResponseOutput) ToTransparentDataEncryptionResponseOutput() TransparentDataEncryptionResponseOutput {
	return o
}

func (o TransparentDataEncryptionResponseOutput) ToTransparentDataEncryptionResponseOutputWithContext(ctx context.Context) TransparentDataEncryptionResponseOutput {
	return o
}

func (o TransparentDataEncryptionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TransparentDataEncryptionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o TransparentDataEncryptionResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v TransparentDataEncryptionResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o TransparentDataEncryptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TransparentDataEncryptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TransparentDataEncryptionResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransparentDataEncryptionResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o TransparentDataEncryptionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TransparentDataEncryptionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type TransparentDataEncryptionResponseArrayOutput struct{ *pulumi.OutputState }

func (TransparentDataEncryptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransparentDataEncryptionResponse)(nil)).Elem()
}

func (o TransparentDataEncryptionResponseArrayOutput) ToTransparentDataEncryptionResponseArrayOutput() TransparentDataEncryptionResponseArrayOutput {
	return o
}

func (o TransparentDataEncryptionResponseArrayOutput) ToTransparentDataEncryptionResponseArrayOutputWithContext(ctx context.Context) TransparentDataEncryptionResponseArrayOutput {
	return o
}

func (o TransparentDataEncryptionResponseArrayOutput) Index(i pulumi.IntInput) TransparentDataEncryptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TransparentDataEncryptionResponse {
		return vs[0].([]TransparentDataEncryptionResponse)[vs[1].(int)]
	}).(TransparentDataEncryptionResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(OperationImpactResponseOutput{})
	pulumi.RegisterOutputType(OperationImpactResponseArrayOutput{})
	pulumi.RegisterOutputType(RecommendedIndexResponseOutput{})
	pulumi.RegisterOutputType(RecommendedIndexResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceTierAdvisorResponseOutput{})
	pulumi.RegisterOutputType(ServiceTierAdvisorResponseArrayOutput{})
	pulumi.RegisterOutputType(SloUsageMetricResponseOutput{})
	pulumi.RegisterOutputType(SloUsageMetricResponseArrayOutput{})
	pulumi.RegisterOutputType(TransparentDataEncryptionResponseOutput{})
	pulumi.RegisterOutputType(TransparentDataEncryptionResponseArrayOutput{})
}
