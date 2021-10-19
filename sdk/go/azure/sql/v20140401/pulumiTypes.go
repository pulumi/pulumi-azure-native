


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





type OperationImpactResponseInput interface {
	pulumi.Input

	ToOperationImpactResponseOutput() OperationImpactResponseOutput
	ToOperationImpactResponseOutputWithContext(context.Context) OperationImpactResponseOutput
}

type OperationImpactResponseArgs struct {
	ChangeValueAbsolute pulumi.Float64Input `pulumi:"changeValueAbsolute"`
	ChangeValueRelative pulumi.Float64Input `pulumi:"changeValueRelative"`
	Name                pulumi.StringInput  `pulumi:"name"`
	Unit                pulumi.StringInput  `pulumi:"unit"`
}

func (OperationImpactResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationImpactResponse)(nil)).Elem()
}

func (i OperationImpactResponseArgs) ToOperationImpactResponseOutput() OperationImpactResponseOutput {
	return i.ToOperationImpactResponseOutputWithContext(context.Background())
}

func (i OperationImpactResponseArgs) ToOperationImpactResponseOutputWithContext(ctx context.Context) OperationImpactResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperationImpactResponseOutput)
}





type OperationImpactResponseArrayInput interface {
	pulumi.Input

	ToOperationImpactResponseArrayOutput() OperationImpactResponseArrayOutput
	ToOperationImpactResponseArrayOutputWithContext(context.Context) OperationImpactResponseArrayOutput
}

type OperationImpactResponseArray []OperationImpactResponseInput

func (OperationImpactResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OperationImpactResponse)(nil)).Elem()
}

func (i OperationImpactResponseArray) ToOperationImpactResponseArrayOutput() OperationImpactResponseArrayOutput {
	return i.ToOperationImpactResponseArrayOutputWithContext(context.Background())
}

func (i OperationImpactResponseArray) ToOperationImpactResponseArrayOutputWithContext(ctx context.Context) OperationImpactResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperationImpactResponseArrayOutput)
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





type RecommendedIndexResponseInput interface {
	pulumi.Input

	ToRecommendedIndexResponseOutput() RecommendedIndexResponseOutput
	ToRecommendedIndexResponseOutputWithContext(context.Context) RecommendedIndexResponseOutput
}

type RecommendedIndexResponseArgs struct {
	Action          pulumi.StringInput                `pulumi:"action"`
	Columns         pulumi.StringArrayInput           `pulumi:"columns"`
	Created         pulumi.StringInput                `pulumi:"created"`
	EstimatedImpact OperationImpactResponseArrayInput `pulumi:"estimatedImpact"`
	Id              pulumi.StringInput                `pulumi:"id"`
	IncludedColumns pulumi.StringArrayInput           `pulumi:"includedColumns"`
	IndexScript     pulumi.StringInput                `pulumi:"indexScript"`
	IndexType       pulumi.StringInput                `pulumi:"indexType"`
	LastModified    pulumi.StringInput                `pulumi:"lastModified"`
	Name            pulumi.StringInput                `pulumi:"name"`
	ReportedImpact  OperationImpactResponseArrayInput `pulumi:"reportedImpact"`
	Schema          pulumi.StringInput                `pulumi:"schema"`
	State           pulumi.StringInput                `pulumi:"state"`
	Table           pulumi.StringInput                `pulumi:"table"`
	Type            pulumi.StringInput                `pulumi:"type"`
}

func (RecommendedIndexResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedIndexResponse)(nil)).Elem()
}

func (i RecommendedIndexResponseArgs) ToRecommendedIndexResponseOutput() RecommendedIndexResponseOutput {
	return i.ToRecommendedIndexResponseOutputWithContext(context.Background())
}

func (i RecommendedIndexResponseArgs) ToRecommendedIndexResponseOutputWithContext(ctx context.Context) RecommendedIndexResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedIndexResponseOutput)
}





type RecommendedIndexResponseArrayInput interface {
	pulumi.Input

	ToRecommendedIndexResponseArrayOutput() RecommendedIndexResponseArrayOutput
	ToRecommendedIndexResponseArrayOutputWithContext(context.Context) RecommendedIndexResponseArrayOutput
}

type RecommendedIndexResponseArray []RecommendedIndexResponseInput

func (RecommendedIndexResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedIndexResponse)(nil)).Elem()
}

func (i RecommendedIndexResponseArray) ToRecommendedIndexResponseArrayOutput() RecommendedIndexResponseArrayOutput {
	return i.ToRecommendedIndexResponseArrayOutputWithContext(context.Background())
}

func (i RecommendedIndexResponseArray) ToRecommendedIndexResponseArrayOutputWithContext(ctx context.Context) RecommendedIndexResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedIndexResponseArrayOutput)
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





type ServiceTierAdvisorResponseInput interface {
	pulumi.Input

	ToServiceTierAdvisorResponseOutput() ServiceTierAdvisorResponseOutput
	ToServiceTierAdvisorResponseOutputWithContext(context.Context) ServiceTierAdvisorResponseOutput
}

type ServiceTierAdvisorResponseArgs struct {
	ActiveTimeRatio                                        pulumi.Float64Input              `pulumi:"activeTimeRatio"`
	AvgDtu                                                 pulumi.Float64Input              `pulumi:"avgDtu"`
	Confidence                                             pulumi.Float64Input              `pulumi:"confidence"`
	CurrentServiceLevelObjective                           pulumi.StringInput               `pulumi:"currentServiceLevelObjective"`
	CurrentServiceLevelObjectiveId                         pulumi.StringInput               `pulumi:"currentServiceLevelObjectiveId"`
	DatabaseSizeBasedRecommendationServiceLevelObjective   pulumi.StringInput               `pulumi:"databaseSizeBasedRecommendationServiceLevelObjective"`
	DatabaseSizeBasedRecommendationServiceLevelObjectiveId pulumi.StringInput               `pulumi:"databaseSizeBasedRecommendationServiceLevelObjectiveId"`
	DisasterPlanBasedRecommendationServiceLevelObjective   pulumi.StringInput               `pulumi:"disasterPlanBasedRecommendationServiceLevelObjective"`
	DisasterPlanBasedRecommendationServiceLevelObjectiveId pulumi.StringInput               `pulumi:"disasterPlanBasedRecommendationServiceLevelObjectiveId"`
	Id                                                     pulumi.StringInput               `pulumi:"id"`
	MaxDtu                                                 pulumi.Float64Input              `pulumi:"maxDtu"`
	MaxSizeInGB                                            pulumi.Float64Input              `pulumi:"maxSizeInGB"`
	MinDtu                                                 pulumi.Float64Input              `pulumi:"minDtu"`
	Name                                                   pulumi.StringInput               `pulumi:"name"`
	ObservationPeriodEnd                                   pulumi.StringInput               `pulumi:"observationPeriodEnd"`
	ObservationPeriodStart                                 pulumi.StringInput               `pulumi:"observationPeriodStart"`
	OverallRecommendationServiceLevelObjective             pulumi.StringInput               `pulumi:"overallRecommendationServiceLevelObjective"`
	OverallRecommendationServiceLevelObjectiveId           pulumi.StringInput               `pulumi:"overallRecommendationServiceLevelObjectiveId"`
	ServiceLevelObjectiveUsageMetrics                      SloUsageMetricResponseArrayInput `pulumi:"serviceLevelObjectiveUsageMetrics"`
	Type                                                   pulumi.StringInput               `pulumi:"type"`
	UsageBasedRecommendationServiceLevelObjective          pulumi.StringInput               `pulumi:"usageBasedRecommendationServiceLevelObjective"`
	UsageBasedRecommendationServiceLevelObjectiveId        pulumi.StringInput               `pulumi:"usageBasedRecommendationServiceLevelObjectiveId"`
}

func (ServiceTierAdvisorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTierAdvisorResponse)(nil)).Elem()
}

func (i ServiceTierAdvisorResponseArgs) ToServiceTierAdvisorResponseOutput() ServiceTierAdvisorResponseOutput {
	return i.ToServiceTierAdvisorResponseOutputWithContext(context.Background())
}

func (i ServiceTierAdvisorResponseArgs) ToServiceTierAdvisorResponseOutputWithContext(ctx context.Context) ServiceTierAdvisorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTierAdvisorResponseOutput)
}





type ServiceTierAdvisorResponseArrayInput interface {
	pulumi.Input

	ToServiceTierAdvisorResponseArrayOutput() ServiceTierAdvisorResponseArrayOutput
	ToServiceTierAdvisorResponseArrayOutputWithContext(context.Context) ServiceTierAdvisorResponseArrayOutput
}

type ServiceTierAdvisorResponseArray []ServiceTierAdvisorResponseInput

func (ServiceTierAdvisorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceTierAdvisorResponse)(nil)).Elem()
}

func (i ServiceTierAdvisorResponseArray) ToServiceTierAdvisorResponseArrayOutput() ServiceTierAdvisorResponseArrayOutput {
	return i.ToServiceTierAdvisorResponseArrayOutputWithContext(context.Background())
}

func (i ServiceTierAdvisorResponseArray) ToServiceTierAdvisorResponseArrayOutputWithContext(ctx context.Context) ServiceTierAdvisorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTierAdvisorResponseArrayOutput)
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





type SloUsageMetricResponseInput interface {
	pulumi.Input

	ToSloUsageMetricResponseOutput() SloUsageMetricResponseOutput
	ToSloUsageMetricResponseOutputWithContext(context.Context) SloUsageMetricResponseOutput
}

type SloUsageMetricResponseArgs struct {
	InRangeTimeRatio        pulumi.Float64Input `pulumi:"inRangeTimeRatio"`
	ServiceLevelObjective   pulumi.StringInput  `pulumi:"serviceLevelObjective"`
	ServiceLevelObjectiveId pulumi.StringInput  `pulumi:"serviceLevelObjectiveId"`
}

func (SloUsageMetricResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SloUsageMetricResponse)(nil)).Elem()
}

func (i SloUsageMetricResponseArgs) ToSloUsageMetricResponseOutput() SloUsageMetricResponseOutput {
	return i.ToSloUsageMetricResponseOutputWithContext(context.Background())
}

func (i SloUsageMetricResponseArgs) ToSloUsageMetricResponseOutputWithContext(ctx context.Context) SloUsageMetricResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SloUsageMetricResponseOutput)
}





type SloUsageMetricResponseArrayInput interface {
	pulumi.Input

	ToSloUsageMetricResponseArrayOutput() SloUsageMetricResponseArrayOutput
	ToSloUsageMetricResponseArrayOutputWithContext(context.Context) SloUsageMetricResponseArrayOutput
}

type SloUsageMetricResponseArray []SloUsageMetricResponseInput

func (SloUsageMetricResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SloUsageMetricResponse)(nil)).Elem()
}

func (i SloUsageMetricResponseArray) ToSloUsageMetricResponseArrayOutput() SloUsageMetricResponseArrayOutput {
	return i.ToSloUsageMetricResponseArrayOutputWithContext(context.Background())
}

func (i SloUsageMetricResponseArray) ToSloUsageMetricResponseArrayOutputWithContext(ctx context.Context) SloUsageMetricResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SloUsageMetricResponseArrayOutput)
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





type TransparentDataEncryptionResponseInput interface {
	pulumi.Input

	ToTransparentDataEncryptionResponseOutput() TransparentDataEncryptionResponseOutput
	ToTransparentDataEncryptionResponseOutputWithContext(context.Context) TransparentDataEncryptionResponseOutput
}

type TransparentDataEncryptionResponseArgs struct {
	Id       pulumi.StringInput    `pulumi:"id"`
	Location pulumi.StringInput    `pulumi:"location"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Status   pulumi.StringPtrInput `pulumi:"status"`
	Type     pulumi.StringInput    `pulumi:"type"`
}

func (TransparentDataEncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransparentDataEncryptionResponse)(nil)).Elem()
}

func (i TransparentDataEncryptionResponseArgs) ToTransparentDataEncryptionResponseOutput() TransparentDataEncryptionResponseOutput {
	return i.ToTransparentDataEncryptionResponseOutputWithContext(context.Background())
}

func (i TransparentDataEncryptionResponseArgs) ToTransparentDataEncryptionResponseOutputWithContext(ctx context.Context) TransparentDataEncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransparentDataEncryptionResponseOutput)
}





type TransparentDataEncryptionResponseArrayInput interface {
	pulumi.Input

	ToTransparentDataEncryptionResponseArrayOutput() TransparentDataEncryptionResponseArrayOutput
	ToTransparentDataEncryptionResponseArrayOutputWithContext(context.Context) TransparentDataEncryptionResponseArrayOutput
}

type TransparentDataEncryptionResponseArray []TransparentDataEncryptionResponseInput

func (TransparentDataEncryptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransparentDataEncryptionResponse)(nil)).Elem()
}

func (i TransparentDataEncryptionResponseArray) ToTransparentDataEncryptionResponseArrayOutput() TransparentDataEncryptionResponseArrayOutput {
	return i.ToTransparentDataEncryptionResponseArrayOutputWithContext(context.Background())
}

func (i TransparentDataEncryptionResponseArray) ToTransparentDataEncryptionResponseArrayOutputWithContext(ctx context.Context) TransparentDataEncryptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransparentDataEncryptionResponseArrayOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*OperationImpactResponseInput)(nil)).Elem(), OperationImpactResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OperationImpactResponseArrayInput)(nil)).Elem(), OperationImpactResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecommendedIndexResponseInput)(nil)).Elem(), RecommendedIndexResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecommendedIndexResponseArrayInput)(nil)).Elem(), RecommendedIndexResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTierAdvisorResponseInput)(nil)).Elem(), ServiceTierAdvisorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTierAdvisorResponseArrayInput)(nil)).Elem(), ServiceTierAdvisorResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SloUsageMetricResponseInput)(nil)).Elem(), SloUsageMetricResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SloUsageMetricResponseArrayInput)(nil)).Elem(), SloUsageMetricResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransparentDataEncryptionResponseInput)(nil)).Elem(), TransparentDataEncryptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransparentDataEncryptionResponseArrayInput)(nil)).Elem(), TransparentDataEncryptionResponseArray{})
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
