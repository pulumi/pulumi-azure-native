


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseVulnerabilityAssessmentRuleBaselineItem struct {
	Result []string `pulumi:"result"`
}





type DatabaseVulnerabilityAssessmentRuleBaselineItemInput interface {
	pulumi.Input

	ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemOutput
	ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutputWithContext(context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemOutput
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemArgs struct {
	Result pulumi.StringArrayInput `pulumi:"result"`
}

func (DatabaseVulnerabilityAssessmentRuleBaselineItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemArgs) ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemOutput {
	return i.ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutputWithContext(context.Background())
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemArgs) ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseVulnerabilityAssessmentRuleBaselineItemOutput)
}





type DatabaseVulnerabilityAssessmentRuleBaselineItemArrayInput interface {
	pulumi.Input

	ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput
	ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemArray []DatabaseVulnerabilityAssessmentRuleBaselineItemInput

func (DatabaseVulnerabilityAssessmentRuleBaselineItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemArray) ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return i.ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(context.Background())
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemArray) ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemOutput struct{ *pulumi.OutputState }

func (DatabaseVulnerabilityAssessmentRuleBaselineItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemOutput) Result() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatabaseVulnerabilityAssessmentRuleBaselineItem) []string { return v.Result }).(pulumi.StringArrayOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput struct{ *pulumi.OutputState }

func (DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput) Index(i pulumi.IntInput) DatabaseVulnerabilityAssessmentRuleBaselineItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseVulnerabilityAssessmentRuleBaselineItem {
		return vs[0].([]DatabaseVulnerabilityAssessmentRuleBaselineItem)[vs[1].(int)]
	}).(DatabaseVulnerabilityAssessmentRuleBaselineItemOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemResponse struct {
	Result []string `pulumi:"result"`
}





type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseInput interface {
	pulumi.Input

	ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput
	ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutputWithContext(context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArgs struct {
	Result pulumi.StringArrayInput `pulumi:"result"`
}

func (DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemResponse)(nil)).Elem()
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArgs) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return i.ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutputWithContext(context.Background())
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArgs) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput)
}





type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayInput interface {
	pulumi.Input

	ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput
	ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutputWithContext(context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArray []DatabaseVulnerabilityAssessmentRuleBaselineItemResponseInput

func (DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseVulnerabilityAssessmentRuleBaselineItemResponse)(nil)).Elem()
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArray) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput {
	return i.ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutputWithContext(context.Background())
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArray) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput struct{ *pulumi.OutputState }

func (DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemResponse)(nil)).Elem()
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput) Result() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatabaseVulnerabilityAssessmentRuleBaselineItemResponse) []string { return v.Result }).(pulumi.StringArrayOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseVulnerabilityAssessmentRuleBaselineItemResponse)(nil)).Elem()
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) Index(i pulumi.IntInput) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseVulnerabilityAssessmentRuleBaselineItemResponse {
		return vs[0].([]DatabaseVulnerabilityAssessmentRuleBaselineItemResponse)[vs[1].(int)]
	}).(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput)
}

type ElasticPoolPerDatabaseSettings struct {
	MaxCapacity *float64 `pulumi:"maxCapacity"`
	MinCapacity *float64 `pulumi:"minCapacity"`
}





type ElasticPoolPerDatabaseSettingsInput interface {
	pulumi.Input

	ToElasticPoolPerDatabaseSettingsOutput() ElasticPoolPerDatabaseSettingsOutput
	ToElasticPoolPerDatabaseSettingsOutputWithContext(context.Context) ElasticPoolPerDatabaseSettingsOutput
}

type ElasticPoolPerDatabaseSettingsArgs struct {
	MaxCapacity pulumi.Float64PtrInput `pulumi:"maxCapacity"`
	MinCapacity pulumi.Float64PtrInput `pulumi:"minCapacity"`
}

func (ElasticPoolPerDatabaseSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolPerDatabaseSettings)(nil)).Elem()
}

func (i ElasticPoolPerDatabaseSettingsArgs) ToElasticPoolPerDatabaseSettingsOutput() ElasticPoolPerDatabaseSettingsOutput {
	return i.ToElasticPoolPerDatabaseSettingsOutputWithContext(context.Background())
}

func (i ElasticPoolPerDatabaseSettingsArgs) ToElasticPoolPerDatabaseSettingsOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPerDatabaseSettingsOutput)
}

func (i ElasticPoolPerDatabaseSettingsArgs) ToElasticPoolPerDatabaseSettingsPtrOutput() ElasticPoolPerDatabaseSettingsPtrOutput {
	return i.ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(context.Background())
}

func (i ElasticPoolPerDatabaseSettingsArgs) ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPerDatabaseSettingsOutput).ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(ctx)
}









type ElasticPoolPerDatabaseSettingsPtrInput interface {
	pulumi.Input

	ToElasticPoolPerDatabaseSettingsPtrOutput() ElasticPoolPerDatabaseSettingsPtrOutput
	ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(context.Context) ElasticPoolPerDatabaseSettingsPtrOutput
}

type elasticPoolPerDatabaseSettingsPtrType ElasticPoolPerDatabaseSettingsArgs

func ElasticPoolPerDatabaseSettingsPtr(v *ElasticPoolPerDatabaseSettingsArgs) ElasticPoolPerDatabaseSettingsPtrInput {
	return (*elasticPoolPerDatabaseSettingsPtrType)(v)
}

func (*elasticPoolPerDatabaseSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPoolPerDatabaseSettings)(nil)).Elem()
}

func (i *elasticPoolPerDatabaseSettingsPtrType) ToElasticPoolPerDatabaseSettingsPtrOutput() ElasticPoolPerDatabaseSettingsPtrOutput {
	return i.ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(context.Background())
}

func (i *elasticPoolPerDatabaseSettingsPtrType) ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPerDatabaseSettingsPtrOutput)
}

type ElasticPoolPerDatabaseSettingsOutput struct{ *pulumi.OutputState }

func (ElasticPoolPerDatabaseSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolPerDatabaseSettings)(nil)).Elem()
}

func (o ElasticPoolPerDatabaseSettingsOutput) ToElasticPoolPerDatabaseSettingsOutput() ElasticPoolPerDatabaseSettingsOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsOutput) ToElasticPoolPerDatabaseSettingsOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsOutput) ToElasticPoolPerDatabaseSettingsPtrOutput() ElasticPoolPerDatabaseSettingsPtrOutput {
	return o.ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(context.Background())
}

func (o ElasticPoolPerDatabaseSettingsOutput) ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ElasticPoolPerDatabaseSettings) *ElasticPoolPerDatabaseSettings {
		return &v
	}).(ElasticPoolPerDatabaseSettingsPtrOutput)
}

func (o ElasticPoolPerDatabaseSettingsOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ElasticPoolPerDatabaseSettings) *float64 { return v.MaxCapacity }).(pulumi.Float64PtrOutput)
}

func (o ElasticPoolPerDatabaseSettingsOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ElasticPoolPerDatabaseSettings) *float64 { return v.MinCapacity }).(pulumi.Float64PtrOutput)
}

type ElasticPoolPerDatabaseSettingsPtrOutput struct{ *pulumi.OutputState }

func (ElasticPoolPerDatabaseSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPoolPerDatabaseSettings)(nil)).Elem()
}

func (o ElasticPoolPerDatabaseSettingsPtrOutput) ToElasticPoolPerDatabaseSettingsPtrOutput() ElasticPoolPerDatabaseSettingsPtrOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsPtrOutput) ToElasticPoolPerDatabaseSettingsPtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsPtrOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsPtrOutput) Elem() ElasticPoolPerDatabaseSettingsOutput {
	return o.ApplyT(func(v *ElasticPoolPerDatabaseSettings) ElasticPoolPerDatabaseSettings {
		if v != nil {
			return *v
		}
		var ret ElasticPoolPerDatabaseSettings
		return ret
	}).(ElasticPoolPerDatabaseSettingsOutput)
}

func (o ElasticPoolPerDatabaseSettingsPtrOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ElasticPoolPerDatabaseSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxCapacity
	}).(pulumi.Float64PtrOutput)
}

func (o ElasticPoolPerDatabaseSettingsPtrOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ElasticPoolPerDatabaseSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.MinCapacity
	}).(pulumi.Float64PtrOutput)
}

type ElasticPoolPerDatabaseSettingsResponse struct {
	MaxCapacity *float64 `pulumi:"maxCapacity"`
	MinCapacity *float64 `pulumi:"minCapacity"`
}





type ElasticPoolPerDatabaseSettingsResponseInput interface {
	pulumi.Input

	ToElasticPoolPerDatabaseSettingsResponseOutput() ElasticPoolPerDatabaseSettingsResponseOutput
	ToElasticPoolPerDatabaseSettingsResponseOutputWithContext(context.Context) ElasticPoolPerDatabaseSettingsResponseOutput
}

type ElasticPoolPerDatabaseSettingsResponseArgs struct {
	MaxCapacity pulumi.Float64PtrInput `pulumi:"maxCapacity"`
	MinCapacity pulumi.Float64PtrInput `pulumi:"minCapacity"`
}

func (ElasticPoolPerDatabaseSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolPerDatabaseSettingsResponse)(nil)).Elem()
}

func (i ElasticPoolPerDatabaseSettingsResponseArgs) ToElasticPoolPerDatabaseSettingsResponseOutput() ElasticPoolPerDatabaseSettingsResponseOutput {
	return i.ToElasticPoolPerDatabaseSettingsResponseOutputWithContext(context.Background())
}

func (i ElasticPoolPerDatabaseSettingsResponseArgs) ToElasticPoolPerDatabaseSettingsResponseOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPerDatabaseSettingsResponseOutput)
}

func (i ElasticPoolPerDatabaseSettingsResponseArgs) ToElasticPoolPerDatabaseSettingsResponsePtrOutput() ElasticPoolPerDatabaseSettingsResponsePtrOutput {
	return i.ToElasticPoolPerDatabaseSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ElasticPoolPerDatabaseSettingsResponseArgs) ToElasticPoolPerDatabaseSettingsResponsePtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPerDatabaseSettingsResponseOutput).ToElasticPoolPerDatabaseSettingsResponsePtrOutputWithContext(ctx)
}









type ElasticPoolPerDatabaseSettingsResponsePtrInput interface {
	pulumi.Input

	ToElasticPoolPerDatabaseSettingsResponsePtrOutput() ElasticPoolPerDatabaseSettingsResponsePtrOutput
	ToElasticPoolPerDatabaseSettingsResponsePtrOutputWithContext(context.Context) ElasticPoolPerDatabaseSettingsResponsePtrOutput
}

type elasticPoolPerDatabaseSettingsResponsePtrType ElasticPoolPerDatabaseSettingsResponseArgs

func ElasticPoolPerDatabaseSettingsResponsePtr(v *ElasticPoolPerDatabaseSettingsResponseArgs) ElasticPoolPerDatabaseSettingsResponsePtrInput {
	return (*elasticPoolPerDatabaseSettingsResponsePtrType)(v)
}

func (*elasticPoolPerDatabaseSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPoolPerDatabaseSettingsResponse)(nil)).Elem()
}

func (i *elasticPoolPerDatabaseSettingsResponsePtrType) ToElasticPoolPerDatabaseSettingsResponsePtrOutput() ElasticPoolPerDatabaseSettingsResponsePtrOutput {
	return i.ToElasticPoolPerDatabaseSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *elasticPoolPerDatabaseSettingsResponsePtrType) ToElasticPoolPerDatabaseSettingsResponsePtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolPerDatabaseSettingsResponsePtrOutput)
}

type ElasticPoolPerDatabaseSettingsResponseOutput struct{ *pulumi.OutputState }

func (ElasticPoolPerDatabaseSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolPerDatabaseSettingsResponse)(nil)).Elem()
}

func (o ElasticPoolPerDatabaseSettingsResponseOutput) ToElasticPoolPerDatabaseSettingsResponseOutput() ElasticPoolPerDatabaseSettingsResponseOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsResponseOutput) ToElasticPoolPerDatabaseSettingsResponseOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsResponseOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsResponseOutput) ToElasticPoolPerDatabaseSettingsResponsePtrOutput() ElasticPoolPerDatabaseSettingsResponsePtrOutput {
	return o.ToElasticPoolPerDatabaseSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ElasticPoolPerDatabaseSettingsResponseOutput) ToElasticPoolPerDatabaseSettingsResponsePtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ElasticPoolPerDatabaseSettingsResponse) *ElasticPoolPerDatabaseSettingsResponse {
		return &v
	}).(ElasticPoolPerDatabaseSettingsResponsePtrOutput)
}

func (o ElasticPoolPerDatabaseSettingsResponseOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ElasticPoolPerDatabaseSettingsResponse) *float64 { return v.MaxCapacity }).(pulumi.Float64PtrOutput)
}

func (o ElasticPoolPerDatabaseSettingsResponseOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ElasticPoolPerDatabaseSettingsResponse) *float64 { return v.MinCapacity }).(pulumi.Float64PtrOutput)
}

type ElasticPoolPerDatabaseSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ElasticPoolPerDatabaseSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPoolPerDatabaseSettingsResponse)(nil)).Elem()
}

func (o ElasticPoolPerDatabaseSettingsResponsePtrOutput) ToElasticPoolPerDatabaseSettingsResponsePtrOutput() ElasticPoolPerDatabaseSettingsResponsePtrOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsResponsePtrOutput) ToElasticPoolPerDatabaseSettingsResponsePtrOutputWithContext(ctx context.Context) ElasticPoolPerDatabaseSettingsResponsePtrOutput {
	return o
}

func (o ElasticPoolPerDatabaseSettingsResponsePtrOutput) Elem() ElasticPoolPerDatabaseSettingsResponseOutput {
	return o.ApplyT(func(v *ElasticPoolPerDatabaseSettingsResponse) ElasticPoolPerDatabaseSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ElasticPoolPerDatabaseSettingsResponse
		return ret
	}).(ElasticPoolPerDatabaseSettingsResponseOutput)
}

func (o ElasticPoolPerDatabaseSettingsResponsePtrOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ElasticPoolPerDatabaseSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxCapacity
	}).(pulumi.Float64PtrOutput)
}

func (o ElasticPoolPerDatabaseSettingsResponsePtrOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ElasticPoolPerDatabaseSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MinCapacity
	}).(pulumi.Float64PtrOutput)
}

type FailoverGroupReadOnlyEndpoint struct {
	FailoverPolicy *string `pulumi:"failoverPolicy"`
}





type FailoverGroupReadOnlyEndpointInput interface {
	pulumi.Input

	ToFailoverGroupReadOnlyEndpointOutput() FailoverGroupReadOnlyEndpointOutput
	ToFailoverGroupReadOnlyEndpointOutputWithContext(context.Context) FailoverGroupReadOnlyEndpointOutput
}

type FailoverGroupReadOnlyEndpointArgs struct {
	FailoverPolicy pulumi.StringPtrInput `pulumi:"failoverPolicy"`
}

func (FailoverGroupReadOnlyEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (i FailoverGroupReadOnlyEndpointArgs) ToFailoverGroupReadOnlyEndpointOutput() FailoverGroupReadOnlyEndpointOutput {
	return i.ToFailoverGroupReadOnlyEndpointOutputWithContext(context.Background())
}

func (i FailoverGroupReadOnlyEndpointArgs) ToFailoverGroupReadOnlyEndpointOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointOutput)
}

func (i FailoverGroupReadOnlyEndpointArgs) ToFailoverGroupReadOnlyEndpointPtrOutput() FailoverGroupReadOnlyEndpointPtrOutput {
	return i.ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Background())
}

func (i FailoverGroupReadOnlyEndpointArgs) ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointOutput).ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx)
}









type FailoverGroupReadOnlyEndpointPtrInput interface {
	pulumi.Input

	ToFailoverGroupReadOnlyEndpointPtrOutput() FailoverGroupReadOnlyEndpointPtrOutput
	ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Context) FailoverGroupReadOnlyEndpointPtrOutput
}

type failoverGroupReadOnlyEndpointPtrType FailoverGroupReadOnlyEndpointArgs

func FailoverGroupReadOnlyEndpointPtr(v *FailoverGroupReadOnlyEndpointArgs) FailoverGroupReadOnlyEndpointPtrInput {
	return (*failoverGroupReadOnlyEndpointPtrType)(v)
}

func (*failoverGroupReadOnlyEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (i *failoverGroupReadOnlyEndpointPtrType) ToFailoverGroupReadOnlyEndpointPtrOutput() FailoverGroupReadOnlyEndpointPtrOutput {
	return i.ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Background())
}

func (i *failoverGroupReadOnlyEndpointPtrType) ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointPtrOutput)
}

type FailoverGroupReadOnlyEndpointOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadOnlyEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (o FailoverGroupReadOnlyEndpointOutput) ToFailoverGroupReadOnlyEndpointOutput() FailoverGroupReadOnlyEndpointOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointOutput) ToFailoverGroupReadOnlyEndpointOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointOutput) ToFailoverGroupReadOnlyEndpointPtrOutput() FailoverGroupReadOnlyEndpointPtrOutput {
	return o.ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Background())
}

func (o FailoverGroupReadOnlyEndpointOutput) ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailoverGroupReadOnlyEndpoint) *FailoverGroupReadOnlyEndpoint {
		return &v
	}).(FailoverGroupReadOnlyEndpointPtrOutput)
}

func (o FailoverGroupReadOnlyEndpointOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverGroupReadOnlyEndpoint) *string { return v.FailoverPolicy }).(pulumi.StringPtrOutput)
}

type FailoverGroupReadOnlyEndpointPtrOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadOnlyEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (o FailoverGroupReadOnlyEndpointPtrOutput) ToFailoverGroupReadOnlyEndpointPtrOutput() FailoverGroupReadOnlyEndpointPtrOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointPtrOutput) ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointPtrOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointPtrOutput) Elem() FailoverGroupReadOnlyEndpointOutput {
	return o.ApplyT(func(v *FailoverGroupReadOnlyEndpoint) FailoverGroupReadOnlyEndpoint {
		if v != nil {
			return *v
		}
		var ret FailoverGroupReadOnlyEndpoint
		return ret
	}).(FailoverGroupReadOnlyEndpointOutput)
}

func (o FailoverGroupReadOnlyEndpointPtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadOnlyEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

type FailoverGroupReadOnlyEndpointResponse struct {
	FailoverPolicy *string `pulumi:"failoverPolicy"`
}





type FailoverGroupReadOnlyEndpointResponseInput interface {
	pulumi.Input

	ToFailoverGroupReadOnlyEndpointResponseOutput() FailoverGroupReadOnlyEndpointResponseOutput
	ToFailoverGroupReadOnlyEndpointResponseOutputWithContext(context.Context) FailoverGroupReadOnlyEndpointResponseOutput
}

type FailoverGroupReadOnlyEndpointResponseArgs struct {
	FailoverPolicy pulumi.StringPtrInput `pulumi:"failoverPolicy"`
}

func (FailoverGroupReadOnlyEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (i FailoverGroupReadOnlyEndpointResponseArgs) ToFailoverGroupReadOnlyEndpointResponseOutput() FailoverGroupReadOnlyEndpointResponseOutput {
	return i.ToFailoverGroupReadOnlyEndpointResponseOutputWithContext(context.Background())
}

func (i FailoverGroupReadOnlyEndpointResponseArgs) ToFailoverGroupReadOnlyEndpointResponseOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointResponseOutput)
}

func (i FailoverGroupReadOnlyEndpointResponseArgs) ToFailoverGroupReadOnlyEndpointResponsePtrOutput() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return i.ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Background())
}

func (i FailoverGroupReadOnlyEndpointResponseArgs) ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointResponseOutput).ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx)
}









type FailoverGroupReadOnlyEndpointResponsePtrInput interface {
	pulumi.Input

	ToFailoverGroupReadOnlyEndpointResponsePtrOutput() FailoverGroupReadOnlyEndpointResponsePtrOutput
	ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Context) FailoverGroupReadOnlyEndpointResponsePtrOutput
}

type failoverGroupReadOnlyEndpointResponsePtrType FailoverGroupReadOnlyEndpointResponseArgs

func FailoverGroupReadOnlyEndpointResponsePtr(v *FailoverGroupReadOnlyEndpointResponseArgs) FailoverGroupReadOnlyEndpointResponsePtrInput {
	return (*failoverGroupReadOnlyEndpointResponsePtrType)(v)
}

func (*failoverGroupReadOnlyEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (i *failoverGroupReadOnlyEndpointResponsePtrType) ToFailoverGroupReadOnlyEndpointResponsePtrOutput() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return i.ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *failoverGroupReadOnlyEndpointResponsePtrType) ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointResponsePtrOutput)
}

type FailoverGroupReadOnlyEndpointResponseOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadOnlyEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (o FailoverGroupReadOnlyEndpointResponseOutput) ToFailoverGroupReadOnlyEndpointResponseOutput() FailoverGroupReadOnlyEndpointResponseOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointResponseOutput) ToFailoverGroupReadOnlyEndpointResponseOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponseOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointResponseOutput) ToFailoverGroupReadOnlyEndpointResponsePtrOutput() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Background())
}

func (o FailoverGroupReadOnlyEndpointResponseOutput) ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailoverGroupReadOnlyEndpointResponse) *FailoverGroupReadOnlyEndpointResponse {
		return &v
	}).(FailoverGroupReadOnlyEndpointResponsePtrOutput)
}

func (o FailoverGroupReadOnlyEndpointResponseOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverGroupReadOnlyEndpointResponse) *string { return v.FailoverPolicy }).(pulumi.StringPtrOutput)
}

type FailoverGroupReadOnlyEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadOnlyEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (o FailoverGroupReadOnlyEndpointResponsePtrOutput) ToFailoverGroupReadOnlyEndpointResponsePtrOutput() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointResponsePtrOutput) ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointResponsePtrOutput) Elem() FailoverGroupReadOnlyEndpointResponseOutput {
	return o.ApplyT(func(v *FailoverGroupReadOnlyEndpointResponse) FailoverGroupReadOnlyEndpointResponse {
		if v != nil {
			return *v
		}
		var ret FailoverGroupReadOnlyEndpointResponse
		return ret
	}).(FailoverGroupReadOnlyEndpointResponseOutput)
}

func (o FailoverGroupReadOnlyEndpointResponsePtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadOnlyEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

type FailoverGroupReadWriteEndpoint struct {
	FailoverPolicy                         string `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes *int   `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}





type FailoverGroupReadWriteEndpointInput interface {
	pulumi.Input

	ToFailoverGroupReadWriteEndpointOutput() FailoverGroupReadWriteEndpointOutput
	ToFailoverGroupReadWriteEndpointOutputWithContext(context.Context) FailoverGroupReadWriteEndpointOutput
}

type FailoverGroupReadWriteEndpointArgs struct {
	FailoverPolicy                         pulumi.StringInput `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes pulumi.IntPtrInput `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}

func (FailoverGroupReadWriteEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (i FailoverGroupReadWriteEndpointArgs) ToFailoverGroupReadWriteEndpointOutput() FailoverGroupReadWriteEndpointOutput {
	return i.ToFailoverGroupReadWriteEndpointOutputWithContext(context.Background())
}

func (i FailoverGroupReadWriteEndpointArgs) ToFailoverGroupReadWriteEndpointOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointOutput)
}

func (i FailoverGroupReadWriteEndpointArgs) ToFailoverGroupReadWriteEndpointPtrOutput() FailoverGroupReadWriteEndpointPtrOutput {
	return i.ToFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Background())
}

func (i FailoverGroupReadWriteEndpointArgs) ToFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointOutput).ToFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx)
}









type FailoverGroupReadWriteEndpointPtrInput interface {
	pulumi.Input

	ToFailoverGroupReadWriteEndpointPtrOutput() FailoverGroupReadWriteEndpointPtrOutput
	ToFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Context) FailoverGroupReadWriteEndpointPtrOutput
}

type failoverGroupReadWriteEndpointPtrType FailoverGroupReadWriteEndpointArgs

func FailoverGroupReadWriteEndpointPtr(v *FailoverGroupReadWriteEndpointArgs) FailoverGroupReadWriteEndpointPtrInput {
	return (*failoverGroupReadWriteEndpointPtrType)(v)
}

func (*failoverGroupReadWriteEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (i *failoverGroupReadWriteEndpointPtrType) ToFailoverGroupReadWriteEndpointPtrOutput() FailoverGroupReadWriteEndpointPtrOutput {
	return i.ToFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Background())
}

func (i *failoverGroupReadWriteEndpointPtrType) ToFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointPtrOutput)
}

type FailoverGroupReadWriteEndpointOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadWriteEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (o FailoverGroupReadWriteEndpointOutput) ToFailoverGroupReadWriteEndpointOutput() FailoverGroupReadWriteEndpointOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointOutput) ToFailoverGroupReadWriteEndpointOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointOutput) ToFailoverGroupReadWriteEndpointPtrOutput() FailoverGroupReadWriteEndpointPtrOutput {
	return o.ToFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Background())
}

func (o FailoverGroupReadWriteEndpointOutput) ToFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailoverGroupReadWriteEndpoint) *FailoverGroupReadWriteEndpoint {
		return &v
	}).(FailoverGroupReadWriteEndpointPtrOutput)
}

func (o FailoverGroupReadWriteEndpointOutput) FailoverPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpoint) string { return v.FailoverPolicy }).(pulumi.StringOutput)
}

func (o FailoverGroupReadWriteEndpointOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpoint) *int { return v.FailoverWithDataLossGracePeriodMinutes }).(pulumi.IntPtrOutput)
}

type FailoverGroupReadWriteEndpointPtrOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadWriteEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (o FailoverGroupReadWriteEndpointPtrOutput) ToFailoverGroupReadWriteEndpointPtrOutput() FailoverGroupReadWriteEndpointPtrOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointPtrOutput) ToFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointPtrOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointPtrOutput) Elem() FailoverGroupReadWriteEndpointOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpoint) FailoverGroupReadWriteEndpoint {
		if v != nil {
			return *v
		}
		var ret FailoverGroupReadWriteEndpoint
		return ret
	}).(FailoverGroupReadWriteEndpointOutput)
}

func (o FailoverGroupReadWriteEndpointPtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpoint) *string {
		if v == nil {
			return nil
		}
		return &v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

func (o FailoverGroupReadWriteEndpointPtrOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpoint) *int {
		if v == nil {
			return nil
		}
		return v.FailoverWithDataLossGracePeriodMinutes
	}).(pulumi.IntPtrOutput)
}

type FailoverGroupReadWriteEndpointResponse struct {
	FailoverPolicy                         string `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes *int   `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}





type FailoverGroupReadWriteEndpointResponseInput interface {
	pulumi.Input

	ToFailoverGroupReadWriteEndpointResponseOutput() FailoverGroupReadWriteEndpointResponseOutput
	ToFailoverGroupReadWriteEndpointResponseOutputWithContext(context.Context) FailoverGroupReadWriteEndpointResponseOutput
}

type FailoverGroupReadWriteEndpointResponseArgs struct {
	FailoverPolicy                         pulumi.StringInput `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes pulumi.IntPtrInput `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}

func (FailoverGroupReadWriteEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (i FailoverGroupReadWriteEndpointResponseArgs) ToFailoverGroupReadWriteEndpointResponseOutput() FailoverGroupReadWriteEndpointResponseOutput {
	return i.ToFailoverGroupReadWriteEndpointResponseOutputWithContext(context.Background())
}

func (i FailoverGroupReadWriteEndpointResponseArgs) ToFailoverGroupReadWriteEndpointResponseOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointResponseOutput)
}

func (i FailoverGroupReadWriteEndpointResponseArgs) ToFailoverGroupReadWriteEndpointResponsePtrOutput() FailoverGroupReadWriteEndpointResponsePtrOutput {
	return i.ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Background())
}

func (i FailoverGroupReadWriteEndpointResponseArgs) ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointResponseOutput).ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx)
}









type FailoverGroupReadWriteEndpointResponsePtrInput interface {
	pulumi.Input

	ToFailoverGroupReadWriteEndpointResponsePtrOutput() FailoverGroupReadWriteEndpointResponsePtrOutput
	ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Context) FailoverGroupReadWriteEndpointResponsePtrOutput
}

type failoverGroupReadWriteEndpointResponsePtrType FailoverGroupReadWriteEndpointResponseArgs

func FailoverGroupReadWriteEndpointResponsePtr(v *FailoverGroupReadWriteEndpointResponseArgs) FailoverGroupReadWriteEndpointResponsePtrInput {
	return (*failoverGroupReadWriteEndpointResponsePtrType)(v)
}

func (*failoverGroupReadWriteEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (i *failoverGroupReadWriteEndpointResponsePtrType) ToFailoverGroupReadWriteEndpointResponsePtrOutput() FailoverGroupReadWriteEndpointResponsePtrOutput {
	return i.ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *failoverGroupReadWriteEndpointResponsePtrType) ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointResponsePtrOutput)
}

type FailoverGroupReadWriteEndpointResponseOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadWriteEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (o FailoverGroupReadWriteEndpointResponseOutput) ToFailoverGroupReadWriteEndpointResponseOutput() FailoverGroupReadWriteEndpointResponseOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointResponseOutput) ToFailoverGroupReadWriteEndpointResponseOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponseOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointResponseOutput) ToFailoverGroupReadWriteEndpointResponsePtrOutput() FailoverGroupReadWriteEndpointResponsePtrOutput {
	return o.ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Background())
}

func (o FailoverGroupReadWriteEndpointResponseOutput) ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailoverGroupReadWriteEndpointResponse) *FailoverGroupReadWriteEndpointResponse {
		return &v
	}).(FailoverGroupReadWriteEndpointResponsePtrOutput)
}

func (o FailoverGroupReadWriteEndpointResponseOutput) FailoverPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpointResponse) string { return v.FailoverPolicy }).(pulumi.StringOutput)
}

func (o FailoverGroupReadWriteEndpointResponseOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpointResponse) *int { return v.FailoverWithDataLossGracePeriodMinutes }).(pulumi.IntPtrOutput)
}

type FailoverGroupReadWriteEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadWriteEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (o FailoverGroupReadWriteEndpointResponsePtrOutput) ToFailoverGroupReadWriteEndpointResponsePtrOutput() FailoverGroupReadWriteEndpointResponsePtrOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointResponsePtrOutput) ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponsePtrOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointResponsePtrOutput) Elem() FailoverGroupReadWriteEndpointResponseOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpointResponse) FailoverGroupReadWriteEndpointResponse {
		if v != nil {
			return *v
		}
		var ret FailoverGroupReadWriteEndpointResponse
		return ret
	}).(FailoverGroupReadWriteEndpointResponseOutput)
}

func (o FailoverGroupReadWriteEndpointResponsePtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

func (o FailoverGroupReadWriteEndpointResponsePtrOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpointResponse) *int {
		if v == nil {
			return nil
		}
		return v.FailoverWithDataLossGracePeriodMinutes
	}).(pulumi.IntPtrOutput)
}

type InstanceFailoverGroupReadOnlyEndpoint struct {
	FailoverPolicy *string `pulumi:"failoverPolicy"`
}





type InstanceFailoverGroupReadOnlyEndpointInput interface {
	pulumi.Input

	ToInstanceFailoverGroupReadOnlyEndpointOutput() InstanceFailoverGroupReadOnlyEndpointOutput
	ToInstanceFailoverGroupReadOnlyEndpointOutputWithContext(context.Context) InstanceFailoverGroupReadOnlyEndpointOutput
}

type InstanceFailoverGroupReadOnlyEndpointArgs struct {
	FailoverPolicy pulumi.StringPtrInput `pulumi:"failoverPolicy"`
}

func (InstanceFailoverGroupReadOnlyEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (i InstanceFailoverGroupReadOnlyEndpointArgs) ToInstanceFailoverGroupReadOnlyEndpointOutput() InstanceFailoverGroupReadOnlyEndpointOutput {
	return i.ToInstanceFailoverGroupReadOnlyEndpointOutputWithContext(context.Background())
}

func (i InstanceFailoverGroupReadOnlyEndpointArgs) ToInstanceFailoverGroupReadOnlyEndpointOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadOnlyEndpointOutput)
}

func (i InstanceFailoverGroupReadOnlyEndpointArgs) ToInstanceFailoverGroupReadOnlyEndpointPtrOutput() InstanceFailoverGroupReadOnlyEndpointPtrOutput {
	return i.ToInstanceFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Background())
}

func (i InstanceFailoverGroupReadOnlyEndpointArgs) ToInstanceFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadOnlyEndpointOutput).ToInstanceFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx)
}









type InstanceFailoverGroupReadOnlyEndpointPtrInput interface {
	pulumi.Input

	ToInstanceFailoverGroupReadOnlyEndpointPtrOutput() InstanceFailoverGroupReadOnlyEndpointPtrOutput
	ToInstanceFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Context) InstanceFailoverGroupReadOnlyEndpointPtrOutput
}

type instanceFailoverGroupReadOnlyEndpointPtrType InstanceFailoverGroupReadOnlyEndpointArgs

func InstanceFailoverGroupReadOnlyEndpointPtr(v *InstanceFailoverGroupReadOnlyEndpointArgs) InstanceFailoverGroupReadOnlyEndpointPtrInput {
	return (*instanceFailoverGroupReadOnlyEndpointPtrType)(v)
}

func (*instanceFailoverGroupReadOnlyEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (i *instanceFailoverGroupReadOnlyEndpointPtrType) ToInstanceFailoverGroupReadOnlyEndpointPtrOutput() InstanceFailoverGroupReadOnlyEndpointPtrOutput {
	return i.ToInstanceFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Background())
}

func (i *instanceFailoverGroupReadOnlyEndpointPtrType) ToInstanceFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadOnlyEndpointPtrOutput)
}

type InstanceFailoverGroupReadOnlyEndpointOutput struct{ *pulumi.OutputState }

func (InstanceFailoverGroupReadOnlyEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (o InstanceFailoverGroupReadOnlyEndpointOutput) ToInstanceFailoverGroupReadOnlyEndpointOutput() InstanceFailoverGroupReadOnlyEndpointOutput {
	return o
}

func (o InstanceFailoverGroupReadOnlyEndpointOutput) ToInstanceFailoverGroupReadOnlyEndpointOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointOutput {
	return o
}

func (o InstanceFailoverGroupReadOnlyEndpointOutput) ToInstanceFailoverGroupReadOnlyEndpointPtrOutput() InstanceFailoverGroupReadOnlyEndpointPtrOutput {
	return o.ToInstanceFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Background())
}

func (o InstanceFailoverGroupReadOnlyEndpointOutput) ToInstanceFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceFailoverGroupReadOnlyEndpoint) *InstanceFailoverGroupReadOnlyEndpoint {
		return &v
	}).(InstanceFailoverGroupReadOnlyEndpointPtrOutput)
}

func (o InstanceFailoverGroupReadOnlyEndpointOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceFailoverGroupReadOnlyEndpoint) *string { return v.FailoverPolicy }).(pulumi.StringPtrOutput)
}

type InstanceFailoverGroupReadOnlyEndpointPtrOutput struct{ *pulumi.OutputState }

func (InstanceFailoverGroupReadOnlyEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (o InstanceFailoverGroupReadOnlyEndpointPtrOutput) ToInstanceFailoverGroupReadOnlyEndpointPtrOutput() InstanceFailoverGroupReadOnlyEndpointPtrOutput {
	return o
}

func (o InstanceFailoverGroupReadOnlyEndpointPtrOutput) ToInstanceFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointPtrOutput {
	return o
}

func (o InstanceFailoverGroupReadOnlyEndpointPtrOutput) Elem() InstanceFailoverGroupReadOnlyEndpointOutput {
	return o.ApplyT(func(v *InstanceFailoverGroupReadOnlyEndpoint) InstanceFailoverGroupReadOnlyEndpoint {
		if v != nil {
			return *v
		}
		var ret InstanceFailoverGroupReadOnlyEndpoint
		return ret
	}).(InstanceFailoverGroupReadOnlyEndpointOutput)
}

func (o InstanceFailoverGroupReadOnlyEndpointPtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceFailoverGroupReadOnlyEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

type InstanceFailoverGroupReadOnlyEndpointResponse struct {
	FailoverPolicy *string `pulumi:"failoverPolicy"`
}





type InstanceFailoverGroupReadOnlyEndpointResponseInput interface {
	pulumi.Input

	ToInstanceFailoverGroupReadOnlyEndpointResponseOutput() InstanceFailoverGroupReadOnlyEndpointResponseOutput
	ToInstanceFailoverGroupReadOnlyEndpointResponseOutputWithContext(context.Context) InstanceFailoverGroupReadOnlyEndpointResponseOutput
}

type InstanceFailoverGroupReadOnlyEndpointResponseArgs struct {
	FailoverPolicy pulumi.StringPtrInput `pulumi:"failoverPolicy"`
}

func (InstanceFailoverGroupReadOnlyEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (i InstanceFailoverGroupReadOnlyEndpointResponseArgs) ToInstanceFailoverGroupReadOnlyEndpointResponseOutput() InstanceFailoverGroupReadOnlyEndpointResponseOutput {
	return i.ToInstanceFailoverGroupReadOnlyEndpointResponseOutputWithContext(context.Background())
}

func (i InstanceFailoverGroupReadOnlyEndpointResponseArgs) ToInstanceFailoverGroupReadOnlyEndpointResponseOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadOnlyEndpointResponseOutput)
}

func (i InstanceFailoverGroupReadOnlyEndpointResponseArgs) ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutput() InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
	return i.ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Background())
}

func (i InstanceFailoverGroupReadOnlyEndpointResponseArgs) ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadOnlyEndpointResponseOutput).ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx)
}









type InstanceFailoverGroupReadOnlyEndpointResponsePtrInput interface {
	pulumi.Input

	ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutput() InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput
	ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Context) InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput
}

type instanceFailoverGroupReadOnlyEndpointResponsePtrType InstanceFailoverGroupReadOnlyEndpointResponseArgs

func InstanceFailoverGroupReadOnlyEndpointResponsePtr(v *InstanceFailoverGroupReadOnlyEndpointResponseArgs) InstanceFailoverGroupReadOnlyEndpointResponsePtrInput {
	return (*instanceFailoverGroupReadOnlyEndpointResponsePtrType)(v)
}

func (*instanceFailoverGroupReadOnlyEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (i *instanceFailoverGroupReadOnlyEndpointResponsePtrType) ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutput() InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
	return i.ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *instanceFailoverGroupReadOnlyEndpointResponsePtrType) ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput)
}

type InstanceFailoverGroupReadOnlyEndpointResponseOutput struct{ *pulumi.OutputState }

func (InstanceFailoverGroupReadOnlyEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (o InstanceFailoverGroupReadOnlyEndpointResponseOutput) ToInstanceFailoverGroupReadOnlyEndpointResponseOutput() InstanceFailoverGroupReadOnlyEndpointResponseOutput {
	return o
}

func (o InstanceFailoverGroupReadOnlyEndpointResponseOutput) ToInstanceFailoverGroupReadOnlyEndpointResponseOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointResponseOutput {
	return o
}

func (o InstanceFailoverGroupReadOnlyEndpointResponseOutput) ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutput() InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Background())
}

func (o InstanceFailoverGroupReadOnlyEndpointResponseOutput) ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceFailoverGroupReadOnlyEndpointResponse) *InstanceFailoverGroupReadOnlyEndpointResponse {
		return &v
	}).(InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput)
}

func (o InstanceFailoverGroupReadOnlyEndpointResponseOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceFailoverGroupReadOnlyEndpointResponse) *string { return v.FailoverPolicy }).(pulumi.StringPtrOutput)
}

type InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (o InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput) ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutput() InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o
}

func (o InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput) ToInstanceFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o
}

func (o InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput) Elem() InstanceFailoverGroupReadOnlyEndpointResponseOutput {
	return o.ApplyT(func(v *InstanceFailoverGroupReadOnlyEndpointResponse) InstanceFailoverGroupReadOnlyEndpointResponse {
		if v != nil {
			return *v
		}
		var ret InstanceFailoverGroupReadOnlyEndpointResponse
		return ret
	}).(InstanceFailoverGroupReadOnlyEndpointResponseOutput)
}

func (o InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceFailoverGroupReadOnlyEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

type InstanceFailoverGroupReadWriteEndpoint struct {
	FailoverPolicy                         string `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes *int   `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}





type InstanceFailoverGroupReadWriteEndpointInput interface {
	pulumi.Input

	ToInstanceFailoverGroupReadWriteEndpointOutput() InstanceFailoverGroupReadWriteEndpointOutput
	ToInstanceFailoverGroupReadWriteEndpointOutputWithContext(context.Context) InstanceFailoverGroupReadWriteEndpointOutput
}

type InstanceFailoverGroupReadWriteEndpointArgs struct {
	FailoverPolicy                         pulumi.StringInput `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes pulumi.IntPtrInput `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}

func (InstanceFailoverGroupReadWriteEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (i InstanceFailoverGroupReadWriteEndpointArgs) ToInstanceFailoverGroupReadWriteEndpointOutput() InstanceFailoverGroupReadWriteEndpointOutput {
	return i.ToInstanceFailoverGroupReadWriteEndpointOutputWithContext(context.Background())
}

func (i InstanceFailoverGroupReadWriteEndpointArgs) ToInstanceFailoverGroupReadWriteEndpointOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadWriteEndpointOutput)
}

func (i InstanceFailoverGroupReadWriteEndpointArgs) ToInstanceFailoverGroupReadWriteEndpointPtrOutput() InstanceFailoverGroupReadWriteEndpointPtrOutput {
	return i.ToInstanceFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Background())
}

func (i InstanceFailoverGroupReadWriteEndpointArgs) ToInstanceFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadWriteEndpointOutput).ToInstanceFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx)
}









type InstanceFailoverGroupReadWriteEndpointPtrInput interface {
	pulumi.Input

	ToInstanceFailoverGroupReadWriteEndpointPtrOutput() InstanceFailoverGroupReadWriteEndpointPtrOutput
	ToInstanceFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Context) InstanceFailoverGroupReadWriteEndpointPtrOutput
}

type instanceFailoverGroupReadWriteEndpointPtrType InstanceFailoverGroupReadWriteEndpointArgs

func InstanceFailoverGroupReadWriteEndpointPtr(v *InstanceFailoverGroupReadWriteEndpointArgs) InstanceFailoverGroupReadWriteEndpointPtrInput {
	return (*instanceFailoverGroupReadWriteEndpointPtrType)(v)
}

func (*instanceFailoverGroupReadWriteEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (i *instanceFailoverGroupReadWriteEndpointPtrType) ToInstanceFailoverGroupReadWriteEndpointPtrOutput() InstanceFailoverGroupReadWriteEndpointPtrOutput {
	return i.ToInstanceFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Background())
}

func (i *instanceFailoverGroupReadWriteEndpointPtrType) ToInstanceFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadWriteEndpointPtrOutput)
}

type InstanceFailoverGroupReadWriteEndpointOutput struct{ *pulumi.OutputState }

func (InstanceFailoverGroupReadWriteEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (o InstanceFailoverGroupReadWriteEndpointOutput) ToInstanceFailoverGroupReadWriteEndpointOutput() InstanceFailoverGroupReadWriteEndpointOutput {
	return o
}

func (o InstanceFailoverGroupReadWriteEndpointOutput) ToInstanceFailoverGroupReadWriteEndpointOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointOutput {
	return o
}

func (o InstanceFailoverGroupReadWriteEndpointOutput) ToInstanceFailoverGroupReadWriteEndpointPtrOutput() InstanceFailoverGroupReadWriteEndpointPtrOutput {
	return o.ToInstanceFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Background())
}

func (o InstanceFailoverGroupReadWriteEndpointOutput) ToInstanceFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceFailoverGroupReadWriteEndpoint) *InstanceFailoverGroupReadWriteEndpoint {
		return &v
	}).(InstanceFailoverGroupReadWriteEndpointPtrOutput)
}

func (o InstanceFailoverGroupReadWriteEndpointOutput) FailoverPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceFailoverGroupReadWriteEndpoint) string { return v.FailoverPolicy }).(pulumi.StringOutput)
}

func (o InstanceFailoverGroupReadWriteEndpointOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InstanceFailoverGroupReadWriteEndpoint) *int { return v.FailoverWithDataLossGracePeriodMinutes }).(pulumi.IntPtrOutput)
}

type InstanceFailoverGroupReadWriteEndpointPtrOutput struct{ *pulumi.OutputState }

func (InstanceFailoverGroupReadWriteEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (o InstanceFailoverGroupReadWriteEndpointPtrOutput) ToInstanceFailoverGroupReadWriteEndpointPtrOutput() InstanceFailoverGroupReadWriteEndpointPtrOutput {
	return o
}

func (o InstanceFailoverGroupReadWriteEndpointPtrOutput) ToInstanceFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointPtrOutput {
	return o
}

func (o InstanceFailoverGroupReadWriteEndpointPtrOutput) Elem() InstanceFailoverGroupReadWriteEndpointOutput {
	return o.ApplyT(func(v *InstanceFailoverGroupReadWriteEndpoint) InstanceFailoverGroupReadWriteEndpoint {
		if v != nil {
			return *v
		}
		var ret InstanceFailoverGroupReadWriteEndpoint
		return ret
	}).(InstanceFailoverGroupReadWriteEndpointOutput)
}

func (o InstanceFailoverGroupReadWriteEndpointPtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceFailoverGroupReadWriteEndpoint) *string {
		if v == nil {
			return nil
		}
		return &v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

func (o InstanceFailoverGroupReadWriteEndpointPtrOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InstanceFailoverGroupReadWriteEndpoint) *int {
		if v == nil {
			return nil
		}
		return v.FailoverWithDataLossGracePeriodMinutes
	}).(pulumi.IntPtrOutput)
}

type InstanceFailoverGroupReadWriteEndpointResponse struct {
	FailoverPolicy                         string `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes *int   `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}





type InstanceFailoverGroupReadWriteEndpointResponseInput interface {
	pulumi.Input

	ToInstanceFailoverGroupReadWriteEndpointResponseOutput() InstanceFailoverGroupReadWriteEndpointResponseOutput
	ToInstanceFailoverGroupReadWriteEndpointResponseOutputWithContext(context.Context) InstanceFailoverGroupReadWriteEndpointResponseOutput
}

type InstanceFailoverGroupReadWriteEndpointResponseArgs struct {
	FailoverPolicy                         pulumi.StringInput `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes pulumi.IntPtrInput `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}

func (InstanceFailoverGroupReadWriteEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (i InstanceFailoverGroupReadWriteEndpointResponseArgs) ToInstanceFailoverGroupReadWriteEndpointResponseOutput() InstanceFailoverGroupReadWriteEndpointResponseOutput {
	return i.ToInstanceFailoverGroupReadWriteEndpointResponseOutputWithContext(context.Background())
}

func (i InstanceFailoverGroupReadWriteEndpointResponseArgs) ToInstanceFailoverGroupReadWriteEndpointResponseOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadWriteEndpointResponseOutput)
}

func (i InstanceFailoverGroupReadWriteEndpointResponseArgs) ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutput() InstanceFailoverGroupReadWriteEndpointResponsePtrOutput {
	return i.ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Background())
}

func (i InstanceFailoverGroupReadWriteEndpointResponseArgs) ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadWriteEndpointResponseOutput).ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx)
}









type InstanceFailoverGroupReadWriteEndpointResponsePtrInput interface {
	pulumi.Input

	ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutput() InstanceFailoverGroupReadWriteEndpointResponsePtrOutput
	ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Context) InstanceFailoverGroupReadWriteEndpointResponsePtrOutput
}

type instanceFailoverGroupReadWriteEndpointResponsePtrType InstanceFailoverGroupReadWriteEndpointResponseArgs

func InstanceFailoverGroupReadWriteEndpointResponsePtr(v *InstanceFailoverGroupReadWriteEndpointResponseArgs) InstanceFailoverGroupReadWriteEndpointResponsePtrInput {
	return (*instanceFailoverGroupReadWriteEndpointResponsePtrType)(v)
}

func (*instanceFailoverGroupReadWriteEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (i *instanceFailoverGroupReadWriteEndpointResponsePtrType) ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutput() InstanceFailoverGroupReadWriteEndpointResponsePtrOutput {
	return i.ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *instanceFailoverGroupReadWriteEndpointResponsePtrType) ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupReadWriteEndpointResponsePtrOutput)
}

type InstanceFailoverGroupReadWriteEndpointResponseOutput struct{ *pulumi.OutputState }

func (InstanceFailoverGroupReadWriteEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (o InstanceFailoverGroupReadWriteEndpointResponseOutput) ToInstanceFailoverGroupReadWriteEndpointResponseOutput() InstanceFailoverGroupReadWriteEndpointResponseOutput {
	return o
}

func (o InstanceFailoverGroupReadWriteEndpointResponseOutput) ToInstanceFailoverGroupReadWriteEndpointResponseOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointResponseOutput {
	return o
}

func (o InstanceFailoverGroupReadWriteEndpointResponseOutput) ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutput() InstanceFailoverGroupReadWriteEndpointResponsePtrOutput {
	return o.ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Background())
}

func (o InstanceFailoverGroupReadWriteEndpointResponseOutput) ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceFailoverGroupReadWriteEndpointResponse) *InstanceFailoverGroupReadWriteEndpointResponse {
		return &v
	}).(InstanceFailoverGroupReadWriteEndpointResponsePtrOutput)
}

func (o InstanceFailoverGroupReadWriteEndpointResponseOutput) FailoverPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceFailoverGroupReadWriteEndpointResponse) string { return v.FailoverPolicy }).(pulumi.StringOutput)
}

func (o InstanceFailoverGroupReadWriteEndpointResponseOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InstanceFailoverGroupReadWriteEndpointResponse) *int {
		return v.FailoverWithDataLossGracePeriodMinutes
	}).(pulumi.IntPtrOutput)
}

type InstanceFailoverGroupReadWriteEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (InstanceFailoverGroupReadWriteEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (o InstanceFailoverGroupReadWriteEndpointResponsePtrOutput) ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutput() InstanceFailoverGroupReadWriteEndpointResponsePtrOutput {
	return o
}

func (o InstanceFailoverGroupReadWriteEndpointResponsePtrOutput) ToInstanceFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) InstanceFailoverGroupReadWriteEndpointResponsePtrOutput {
	return o
}

func (o InstanceFailoverGroupReadWriteEndpointResponsePtrOutput) Elem() InstanceFailoverGroupReadWriteEndpointResponseOutput {
	return o.ApplyT(func(v *InstanceFailoverGroupReadWriteEndpointResponse) InstanceFailoverGroupReadWriteEndpointResponse {
		if v != nil {
			return *v
		}
		var ret InstanceFailoverGroupReadWriteEndpointResponse
		return ret
	}).(InstanceFailoverGroupReadWriteEndpointResponseOutput)
}

func (o InstanceFailoverGroupReadWriteEndpointResponsePtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceFailoverGroupReadWriteEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

func (o InstanceFailoverGroupReadWriteEndpointResponsePtrOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InstanceFailoverGroupReadWriteEndpointResponse) *int {
		if v == nil {
			return nil
		}
		return v.FailoverWithDataLossGracePeriodMinutes
	}).(pulumi.IntPtrOutput)
}

type JobSchedule struct {
	Enabled   *bool            `pulumi:"enabled"`
	EndTime   *string          `pulumi:"endTime"`
	Interval  *string          `pulumi:"interval"`
	StartTime *string          `pulumi:"startTime"`
	Type      *JobScheduleType `pulumi:"type"`
}





type JobScheduleInput interface {
	pulumi.Input

	ToJobScheduleOutput() JobScheduleOutput
	ToJobScheduleOutputWithContext(context.Context) JobScheduleOutput
}

type JobScheduleArgs struct {
	Enabled   pulumi.BoolPtrInput     `pulumi:"enabled"`
	EndTime   pulumi.StringPtrInput   `pulumi:"endTime"`
	Interval  pulumi.StringPtrInput   `pulumi:"interval"`
	StartTime pulumi.StringPtrInput   `pulumi:"startTime"`
	Type      JobScheduleTypePtrInput `pulumi:"type"`
}

func (JobScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobSchedule)(nil)).Elem()
}

func (i JobScheduleArgs) ToJobScheduleOutput() JobScheduleOutput {
	return i.ToJobScheduleOutputWithContext(context.Background())
}

func (i JobScheduleArgs) ToJobScheduleOutputWithContext(ctx context.Context) JobScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleOutput)
}

func (i JobScheduleArgs) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return i.ToJobSchedulePtrOutputWithContext(context.Background())
}

func (i JobScheduleArgs) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleOutput).ToJobSchedulePtrOutputWithContext(ctx)
}









type JobSchedulePtrInput interface {
	pulumi.Input

	ToJobSchedulePtrOutput() JobSchedulePtrOutput
	ToJobSchedulePtrOutputWithContext(context.Context) JobSchedulePtrOutput
}

type jobSchedulePtrType JobScheduleArgs

func JobSchedulePtr(v *JobScheduleArgs) JobSchedulePtrInput {
	return (*jobSchedulePtrType)(v)
}

func (*jobSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSchedule)(nil)).Elem()
}

func (i *jobSchedulePtrType) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return i.ToJobSchedulePtrOutputWithContext(context.Background())
}

func (i *jobSchedulePtrType) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobSchedulePtrOutput)
}

type JobScheduleOutput struct{ *pulumi.OutputState }

func (JobScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobSchedule)(nil)).Elem()
}

func (o JobScheduleOutput) ToJobScheduleOutput() JobScheduleOutput {
	return o
}

func (o JobScheduleOutput) ToJobScheduleOutputWithContext(ctx context.Context) JobScheduleOutput {
	return o
}

func (o JobScheduleOutput) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return o.ToJobSchedulePtrOutputWithContext(context.Background())
}

func (o JobScheduleOutput) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobSchedule) *JobSchedule {
		return &v
	}).(JobSchedulePtrOutput)
}

func (o JobScheduleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v JobSchedule) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o JobScheduleOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobSchedule) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o JobScheduleOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobSchedule) *string { return v.Interval }).(pulumi.StringPtrOutput)
}

func (o JobScheduleOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobSchedule) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o JobScheduleOutput) Type() JobScheduleTypePtrOutput {
	return o.ApplyT(func(v JobSchedule) *JobScheduleType { return v.Type }).(JobScheduleTypePtrOutput)
}

type JobSchedulePtrOutput struct{ *pulumi.OutputState }

func (JobSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSchedule)(nil)).Elem()
}

func (o JobSchedulePtrOutput) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return o
}

func (o JobSchedulePtrOutput) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return o
}

func (o JobSchedulePtrOutput) Elem() JobScheduleOutput {
	return o.ApplyT(func(v *JobSchedule) JobSchedule {
		if v != nil {
			return *v
		}
		var ret JobSchedule
		return ret
	}).(JobScheduleOutput)
}

func (o JobSchedulePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *JobSchedule) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o JobSchedulePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobSchedule) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o JobSchedulePtrOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobSchedule) *string {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.StringPtrOutput)
}

func (o JobSchedulePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobSchedule) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o JobSchedulePtrOutput) Type() JobScheduleTypePtrOutput {
	return o.ApplyT(func(v *JobSchedule) *JobScheduleType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(JobScheduleTypePtrOutput)
}

type JobScheduleResponse struct {
	Enabled   *bool   `pulumi:"enabled"`
	EndTime   *string `pulumi:"endTime"`
	Interval  *string `pulumi:"interval"`
	StartTime *string `pulumi:"startTime"`
	Type      *string `pulumi:"type"`
}





type JobScheduleResponseInput interface {
	pulumi.Input

	ToJobScheduleResponseOutput() JobScheduleResponseOutput
	ToJobScheduleResponseOutputWithContext(context.Context) JobScheduleResponseOutput
}

type JobScheduleResponseArgs struct {
	Enabled   pulumi.BoolPtrInput   `pulumi:"enabled"`
	EndTime   pulumi.StringPtrInput `pulumi:"endTime"`
	Interval  pulumi.StringPtrInput `pulumi:"interval"`
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
	Type      pulumi.StringPtrInput `pulumi:"type"`
}

func (JobScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobScheduleResponse)(nil)).Elem()
}

func (i JobScheduleResponseArgs) ToJobScheduleResponseOutput() JobScheduleResponseOutput {
	return i.ToJobScheduleResponseOutputWithContext(context.Background())
}

func (i JobScheduleResponseArgs) ToJobScheduleResponseOutputWithContext(ctx context.Context) JobScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleResponseOutput)
}

func (i JobScheduleResponseArgs) ToJobScheduleResponsePtrOutput() JobScheduleResponsePtrOutput {
	return i.ToJobScheduleResponsePtrOutputWithContext(context.Background())
}

func (i JobScheduleResponseArgs) ToJobScheduleResponsePtrOutputWithContext(ctx context.Context) JobScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleResponseOutput).ToJobScheduleResponsePtrOutputWithContext(ctx)
}









type JobScheduleResponsePtrInput interface {
	pulumi.Input

	ToJobScheduleResponsePtrOutput() JobScheduleResponsePtrOutput
	ToJobScheduleResponsePtrOutputWithContext(context.Context) JobScheduleResponsePtrOutput
}

type jobScheduleResponsePtrType JobScheduleResponseArgs

func JobScheduleResponsePtr(v *JobScheduleResponseArgs) JobScheduleResponsePtrInput {
	return (*jobScheduleResponsePtrType)(v)
}

func (*jobScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobScheduleResponse)(nil)).Elem()
}

func (i *jobScheduleResponsePtrType) ToJobScheduleResponsePtrOutput() JobScheduleResponsePtrOutput {
	return i.ToJobScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *jobScheduleResponsePtrType) ToJobScheduleResponsePtrOutputWithContext(ctx context.Context) JobScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleResponsePtrOutput)
}

type JobScheduleResponseOutput struct{ *pulumi.OutputState }

func (JobScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobScheduleResponse)(nil)).Elem()
}

func (o JobScheduleResponseOutput) ToJobScheduleResponseOutput() JobScheduleResponseOutput {
	return o
}

func (o JobScheduleResponseOutput) ToJobScheduleResponseOutputWithContext(ctx context.Context) JobScheduleResponseOutput {
	return o
}

func (o JobScheduleResponseOutput) ToJobScheduleResponsePtrOutput() JobScheduleResponsePtrOutput {
	return o.ToJobScheduleResponsePtrOutputWithContext(context.Background())
}

func (o JobScheduleResponseOutput) ToJobScheduleResponsePtrOutputWithContext(ctx context.Context) JobScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobScheduleResponse) *JobScheduleResponse {
		return &v
	}).(JobScheduleResponsePtrOutput)
}

func (o JobScheduleResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v JobScheduleResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o JobScheduleResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobScheduleResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponseOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobScheduleResponse) *string { return v.Interval }).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobScheduleResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobScheduleResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JobScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (JobScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobScheduleResponse)(nil)).Elem()
}

func (o JobScheduleResponsePtrOutput) ToJobScheduleResponsePtrOutput() JobScheduleResponsePtrOutput {
	return o
}

func (o JobScheduleResponsePtrOutput) ToJobScheduleResponsePtrOutputWithContext(ctx context.Context) JobScheduleResponsePtrOutput {
	return o
}

func (o JobScheduleResponsePtrOutput) Elem() JobScheduleResponseOutput {
	return o.ApplyT(func(v *JobScheduleResponse) JobScheduleResponse {
		if v != nil {
			return *v
		}
		var ret JobScheduleResponse
		return ret
	}).(JobScheduleResponseOutput)
}

func (o JobScheduleResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *JobScheduleResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o JobScheduleResponsePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponsePtrOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type JobStepAction struct {
	Source *string `pulumi:"source"`
	Type   *string `pulumi:"type"`
	Value  string  `pulumi:"value"`
}





type JobStepActionInput interface {
	pulumi.Input

	ToJobStepActionOutput() JobStepActionOutput
	ToJobStepActionOutputWithContext(context.Context) JobStepActionOutput
}

type JobStepActionArgs struct {
	Source pulumi.StringPtrInput `pulumi:"source"`
	Type   pulumi.StringPtrInput `pulumi:"type"`
	Value  pulumi.StringInput    `pulumi:"value"`
}

func (JobStepActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepAction)(nil)).Elem()
}

func (i JobStepActionArgs) ToJobStepActionOutput() JobStepActionOutput {
	return i.ToJobStepActionOutputWithContext(context.Background())
}

func (i JobStepActionArgs) ToJobStepActionOutputWithContext(ctx context.Context) JobStepActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionOutput)
}

func (i JobStepActionArgs) ToJobStepActionPtrOutput() JobStepActionPtrOutput {
	return i.ToJobStepActionPtrOutputWithContext(context.Background())
}

func (i JobStepActionArgs) ToJobStepActionPtrOutputWithContext(ctx context.Context) JobStepActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionOutput).ToJobStepActionPtrOutputWithContext(ctx)
}









type JobStepActionPtrInput interface {
	pulumi.Input

	ToJobStepActionPtrOutput() JobStepActionPtrOutput
	ToJobStepActionPtrOutputWithContext(context.Context) JobStepActionPtrOutput
}

type jobStepActionPtrType JobStepActionArgs

func JobStepActionPtr(v *JobStepActionArgs) JobStepActionPtrInput {
	return (*jobStepActionPtrType)(v)
}

func (*jobStepActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepAction)(nil)).Elem()
}

func (i *jobStepActionPtrType) ToJobStepActionPtrOutput() JobStepActionPtrOutput {
	return i.ToJobStepActionPtrOutputWithContext(context.Background())
}

func (i *jobStepActionPtrType) ToJobStepActionPtrOutputWithContext(ctx context.Context) JobStepActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionPtrOutput)
}

type JobStepActionOutput struct{ *pulumi.OutputState }

func (JobStepActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepAction)(nil)).Elem()
}

func (o JobStepActionOutput) ToJobStepActionOutput() JobStepActionOutput {
	return o
}

func (o JobStepActionOutput) ToJobStepActionOutputWithContext(ctx context.Context) JobStepActionOutput {
	return o
}

func (o JobStepActionOutput) ToJobStepActionPtrOutput() JobStepActionPtrOutput {
	return o.ToJobStepActionPtrOutputWithContext(context.Background())
}

func (o JobStepActionOutput) ToJobStepActionPtrOutputWithContext(ctx context.Context) JobStepActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepAction) *JobStepAction {
		return &v
	}).(JobStepActionPtrOutput)
}

func (o JobStepActionOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepAction) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o JobStepActionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepAction) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o JobStepActionOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepAction) string { return v.Value }).(pulumi.StringOutput)
}

type JobStepActionPtrOutput struct{ *pulumi.OutputState }

func (JobStepActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepAction)(nil)).Elem()
}

func (o JobStepActionPtrOutput) ToJobStepActionPtrOutput() JobStepActionPtrOutput {
	return o
}

func (o JobStepActionPtrOutput) ToJobStepActionPtrOutputWithContext(ctx context.Context) JobStepActionPtrOutput {
	return o
}

func (o JobStepActionPtrOutput) Elem() JobStepActionOutput {
	return o.ApplyT(func(v *JobStepAction) JobStepAction {
		if v != nil {
			return *v
		}
		var ret JobStepAction
		return ret
	}).(JobStepActionOutput)
}

func (o JobStepActionPtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepAction) *string {
		if v == nil {
			return nil
		}
		return v.Source
	}).(pulumi.StringPtrOutput)
}

func (o JobStepActionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepAction) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o JobStepActionPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepAction) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type JobStepActionResponse struct {
	Source *string `pulumi:"source"`
	Type   *string `pulumi:"type"`
	Value  string  `pulumi:"value"`
}





type JobStepActionResponseInput interface {
	pulumi.Input

	ToJobStepActionResponseOutput() JobStepActionResponseOutput
	ToJobStepActionResponseOutputWithContext(context.Context) JobStepActionResponseOutput
}

type JobStepActionResponseArgs struct {
	Source pulumi.StringPtrInput `pulumi:"source"`
	Type   pulumi.StringPtrInput `pulumi:"type"`
	Value  pulumi.StringInput    `pulumi:"value"`
}

func (JobStepActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepActionResponse)(nil)).Elem()
}

func (i JobStepActionResponseArgs) ToJobStepActionResponseOutput() JobStepActionResponseOutput {
	return i.ToJobStepActionResponseOutputWithContext(context.Background())
}

func (i JobStepActionResponseArgs) ToJobStepActionResponseOutputWithContext(ctx context.Context) JobStepActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionResponseOutput)
}

func (i JobStepActionResponseArgs) ToJobStepActionResponsePtrOutput() JobStepActionResponsePtrOutput {
	return i.ToJobStepActionResponsePtrOutputWithContext(context.Background())
}

func (i JobStepActionResponseArgs) ToJobStepActionResponsePtrOutputWithContext(ctx context.Context) JobStepActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionResponseOutput).ToJobStepActionResponsePtrOutputWithContext(ctx)
}









type JobStepActionResponsePtrInput interface {
	pulumi.Input

	ToJobStepActionResponsePtrOutput() JobStepActionResponsePtrOutput
	ToJobStepActionResponsePtrOutputWithContext(context.Context) JobStepActionResponsePtrOutput
}

type jobStepActionResponsePtrType JobStepActionResponseArgs

func JobStepActionResponsePtr(v *JobStepActionResponseArgs) JobStepActionResponsePtrInput {
	return (*jobStepActionResponsePtrType)(v)
}

func (*jobStepActionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepActionResponse)(nil)).Elem()
}

func (i *jobStepActionResponsePtrType) ToJobStepActionResponsePtrOutput() JobStepActionResponsePtrOutput {
	return i.ToJobStepActionResponsePtrOutputWithContext(context.Background())
}

func (i *jobStepActionResponsePtrType) ToJobStepActionResponsePtrOutputWithContext(ctx context.Context) JobStepActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionResponsePtrOutput)
}

type JobStepActionResponseOutput struct{ *pulumi.OutputState }

func (JobStepActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepActionResponse)(nil)).Elem()
}

func (o JobStepActionResponseOutput) ToJobStepActionResponseOutput() JobStepActionResponseOutput {
	return o
}

func (o JobStepActionResponseOutput) ToJobStepActionResponseOutputWithContext(ctx context.Context) JobStepActionResponseOutput {
	return o
}

func (o JobStepActionResponseOutput) ToJobStepActionResponsePtrOutput() JobStepActionResponsePtrOutput {
	return o.ToJobStepActionResponsePtrOutputWithContext(context.Background())
}

func (o JobStepActionResponseOutput) ToJobStepActionResponsePtrOutputWithContext(ctx context.Context) JobStepActionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepActionResponse) *JobStepActionResponse {
		return &v
	}).(JobStepActionResponsePtrOutput)
}

func (o JobStepActionResponseOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepActionResponse) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o JobStepActionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepActionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o JobStepActionResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepActionResponse) string { return v.Value }).(pulumi.StringOutput)
}

type JobStepActionResponsePtrOutput struct{ *pulumi.OutputState }

func (JobStepActionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepActionResponse)(nil)).Elem()
}

func (o JobStepActionResponsePtrOutput) ToJobStepActionResponsePtrOutput() JobStepActionResponsePtrOutput {
	return o
}

func (o JobStepActionResponsePtrOutput) ToJobStepActionResponsePtrOutputWithContext(ctx context.Context) JobStepActionResponsePtrOutput {
	return o
}

func (o JobStepActionResponsePtrOutput) Elem() JobStepActionResponseOutput {
	return o.ApplyT(func(v *JobStepActionResponse) JobStepActionResponse {
		if v != nil {
			return *v
		}
		var ret JobStepActionResponse
		return ret
	}).(JobStepActionResponseOutput)
}

func (o JobStepActionResponsePtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Source
	}).(pulumi.StringPtrOutput)
}

func (o JobStepActionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o JobStepActionResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepActionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type JobStepExecutionOptions struct {
	InitialRetryIntervalSeconds    *int     `pulumi:"initialRetryIntervalSeconds"`
	MaximumRetryIntervalSeconds    *int     `pulumi:"maximumRetryIntervalSeconds"`
	RetryAttempts                  *int     `pulumi:"retryAttempts"`
	RetryIntervalBackoffMultiplier *float64 `pulumi:"retryIntervalBackoffMultiplier"`
	TimeoutSeconds                 *int     `pulumi:"timeoutSeconds"`
}





type JobStepExecutionOptionsInput interface {
	pulumi.Input

	ToJobStepExecutionOptionsOutput() JobStepExecutionOptionsOutput
	ToJobStepExecutionOptionsOutputWithContext(context.Context) JobStepExecutionOptionsOutput
}

type JobStepExecutionOptionsArgs struct {
	InitialRetryIntervalSeconds    pulumi.IntPtrInput     `pulumi:"initialRetryIntervalSeconds"`
	MaximumRetryIntervalSeconds    pulumi.IntPtrInput     `pulumi:"maximumRetryIntervalSeconds"`
	RetryAttempts                  pulumi.IntPtrInput     `pulumi:"retryAttempts"`
	RetryIntervalBackoffMultiplier pulumi.Float64PtrInput `pulumi:"retryIntervalBackoffMultiplier"`
	TimeoutSeconds                 pulumi.IntPtrInput     `pulumi:"timeoutSeconds"`
}

func (JobStepExecutionOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepExecutionOptions)(nil)).Elem()
}

func (i JobStepExecutionOptionsArgs) ToJobStepExecutionOptionsOutput() JobStepExecutionOptionsOutput {
	return i.ToJobStepExecutionOptionsOutputWithContext(context.Background())
}

func (i JobStepExecutionOptionsArgs) ToJobStepExecutionOptionsOutputWithContext(ctx context.Context) JobStepExecutionOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsOutput)
}

func (i JobStepExecutionOptionsArgs) ToJobStepExecutionOptionsPtrOutput() JobStepExecutionOptionsPtrOutput {
	return i.ToJobStepExecutionOptionsPtrOutputWithContext(context.Background())
}

func (i JobStepExecutionOptionsArgs) ToJobStepExecutionOptionsPtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsOutput).ToJobStepExecutionOptionsPtrOutputWithContext(ctx)
}









type JobStepExecutionOptionsPtrInput interface {
	pulumi.Input

	ToJobStepExecutionOptionsPtrOutput() JobStepExecutionOptionsPtrOutput
	ToJobStepExecutionOptionsPtrOutputWithContext(context.Context) JobStepExecutionOptionsPtrOutput
}

type jobStepExecutionOptionsPtrType JobStepExecutionOptionsArgs

func JobStepExecutionOptionsPtr(v *JobStepExecutionOptionsArgs) JobStepExecutionOptionsPtrInput {
	return (*jobStepExecutionOptionsPtrType)(v)
}

func (*jobStepExecutionOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepExecutionOptions)(nil)).Elem()
}

func (i *jobStepExecutionOptionsPtrType) ToJobStepExecutionOptionsPtrOutput() JobStepExecutionOptionsPtrOutput {
	return i.ToJobStepExecutionOptionsPtrOutputWithContext(context.Background())
}

func (i *jobStepExecutionOptionsPtrType) ToJobStepExecutionOptionsPtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsPtrOutput)
}

type JobStepExecutionOptionsOutput struct{ *pulumi.OutputState }

func (JobStepExecutionOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepExecutionOptions)(nil)).Elem()
}

func (o JobStepExecutionOptionsOutput) ToJobStepExecutionOptionsOutput() JobStepExecutionOptionsOutput {
	return o
}

func (o JobStepExecutionOptionsOutput) ToJobStepExecutionOptionsOutputWithContext(ctx context.Context) JobStepExecutionOptionsOutput {
	return o
}

func (o JobStepExecutionOptionsOutput) ToJobStepExecutionOptionsPtrOutput() JobStepExecutionOptionsPtrOutput {
	return o.ToJobStepExecutionOptionsPtrOutputWithContext(context.Background())
}

func (o JobStepExecutionOptionsOutput) ToJobStepExecutionOptionsPtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepExecutionOptions) *JobStepExecutionOptions {
		return &v
	}).(JobStepExecutionOptionsPtrOutput)
}

func (o JobStepExecutionOptionsOutput) InitialRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptions) *int { return v.InitialRetryIntervalSeconds }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsOutput) MaximumRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptions) *int { return v.MaximumRetryIntervalSeconds }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsOutput) RetryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptions) *int { return v.RetryAttempts }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsOutput) RetryIntervalBackoffMultiplier() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptions) *float64 { return v.RetryIntervalBackoffMultiplier }).(pulumi.Float64PtrOutput)
}

func (o JobStepExecutionOptionsOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptions) *int { return v.TimeoutSeconds }).(pulumi.IntPtrOutput)
}

type JobStepExecutionOptionsPtrOutput struct{ *pulumi.OutputState }

func (JobStepExecutionOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepExecutionOptions)(nil)).Elem()
}

func (o JobStepExecutionOptionsPtrOutput) ToJobStepExecutionOptionsPtrOutput() JobStepExecutionOptionsPtrOutput {
	return o
}

func (o JobStepExecutionOptionsPtrOutput) ToJobStepExecutionOptionsPtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsPtrOutput {
	return o
}

func (o JobStepExecutionOptionsPtrOutput) Elem() JobStepExecutionOptionsOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) JobStepExecutionOptions {
		if v != nil {
			return *v
		}
		var ret JobStepExecutionOptions
		return ret
	}).(JobStepExecutionOptionsOutput)
}

func (o JobStepExecutionOptionsPtrOutput) InitialRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) *int {
		if v == nil {
			return nil
		}
		return v.InitialRetryIntervalSeconds
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsPtrOutput) MaximumRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) *int {
		if v == nil {
			return nil
		}
		return v.MaximumRetryIntervalSeconds
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsPtrOutput) RetryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) *int {
		if v == nil {
			return nil
		}
		return v.RetryAttempts
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsPtrOutput) RetryIntervalBackoffMultiplier() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryIntervalBackoffMultiplier
	}).(pulumi.Float64PtrOutput)
}

func (o JobStepExecutionOptionsPtrOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) *int {
		if v == nil {
			return nil
		}
		return v.TimeoutSeconds
	}).(pulumi.IntPtrOutput)
}

type JobStepExecutionOptionsResponse struct {
	InitialRetryIntervalSeconds    *int     `pulumi:"initialRetryIntervalSeconds"`
	MaximumRetryIntervalSeconds    *int     `pulumi:"maximumRetryIntervalSeconds"`
	RetryAttempts                  *int     `pulumi:"retryAttempts"`
	RetryIntervalBackoffMultiplier *float64 `pulumi:"retryIntervalBackoffMultiplier"`
	TimeoutSeconds                 *int     `pulumi:"timeoutSeconds"`
}





type JobStepExecutionOptionsResponseInput interface {
	pulumi.Input

	ToJobStepExecutionOptionsResponseOutput() JobStepExecutionOptionsResponseOutput
	ToJobStepExecutionOptionsResponseOutputWithContext(context.Context) JobStepExecutionOptionsResponseOutput
}

type JobStepExecutionOptionsResponseArgs struct {
	InitialRetryIntervalSeconds    pulumi.IntPtrInput     `pulumi:"initialRetryIntervalSeconds"`
	MaximumRetryIntervalSeconds    pulumi.IntPtrInput     `pulumi:"maximumRetryIntervalSeconds"`
	RetryAttempts                  pulumi.IntPtrInput     `pulumi:"retryAttempts"`
	RetryIntervalBackoffMultiplier pulumi.Float64PtrInput `pulumi:"retryIntervalBackoffMultiplier"`
	TimeoutSeconds                 pulumi.IntPtrInput     `pulumi:"timeoutSeconds"`
}

func (JobStepExecutionOptionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepExecutionOptionsResponse)(nil)).Elem()
}

func (i JobStepExecutionOptionsResponseArgs) ToJobStepExecutionOptionsResponseOutput() JobStepExecutionOptionsResponseOutput {
	return i.ToJobStepExecutionOptionsResponseOutputWithContext(context.Background())
}

func (i JobStepExecutionOptionsResponseArgs) ToJobStepExecutionOptionsResponseOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsResponseOutput)
}

func (i JobStepExecutionOptionsResponseArgs) ToJobStepExecutionOptionsResponsePtrOutput() JobStepExecutionOptionsResponsePtrOutput {
	return i.ToJobStepExecutionOptionsResponsePtrOutputWithContext(context.Background())
}

func (i JobStepExecutionOptionsResponseArgs) ToJobStepExecutionOptionsResponsePtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsResponseOutput).ToJobStepExecutionOptionsResponsePtrOutputWithContext(ctx)
}









type JobStepExecutionOptionsResponsePtrInput interface {
	pulumi.Input

	ToJobStepExecutionOptionsResponsePtrOutput() JobStepExecutionOptionsResponsePtrOutput
	ToJobStepExecutionOptionsResponsePtrOutputWithContext(context.Context) JobStepExecutionOptionsResponsePtrOutput
}

type jobStepExecutionOptionsResponsePtrType JobStepExecutionOptionsResponseArgs

func JobStepExecutionOptionsResponsePtr(v *JobStepExecutionOptionsResponseArgs) JobStepExecutionOptionsResponsePtrInput {
	return (*jobStepExecutionOptionsResponsePtrType)(v)
}

func (*jobStepExecutionOptionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepExecutionOptionsResponse)(nil)).Elem()
}

func (i *jobStepExecutionOptionsResponsePtrType) ToJobStepExecutionOptionsResponsePtrOutput() JobStepExecutionOptionsResponsePtrOutput {
	return i.ToJobStepExecutionOptionsResponsePtrOutputWithContext(context.Background())
}

func (i *jobStepExecutionOptionsResponsePtrType) ToJobStepExecutionOptionsResponsePtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsResponsePtrOutput)
}

type JobStepExecutionOptionsResponseOutput struct{ *pulumi.OutputState }

func (JobStepExecutionOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepExecutionOptionsResponse)(nil)).Elem()
}

func (o JobStepExecutionOptionsResponseOutput) ToJobStepExecutionOptionsResponseOutput() JobStepExecutionOptionsResponseOutput {
	return o
}

func (o JobStepExecutionOptionsResponseOutput) ToJobStepExecutionOptionsResponseOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponseOutput {
	return o
}

func (o JobStepExecutionOptionsResponseOutput) ToJobStepExecutionOptionsResponsePtrOutput() JobStepExecutionOptionsResponsePtrOutput {
	return o.ToJobStepExecutionOptionsResponsePtrOutputWithContext(context.Background())
}

func (o JobStepExecutionOptionsResponseOutput) ToJobStepExecutionOptionsResponsePtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepExecutionOptionsResponse) *JobStepExecutionOptionsResponse {
		return &v
	}).(JobStepExecutionOptionsResponsePtrOutput)
}

func (o JobStepExecutionOptionsResponseOutput) InitialRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptionsResponse) *int { return v.InitialRetryIntervalSeconds }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponseOutput) MaximumRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptionsResponse) *int { return v.MaximumRetryIntervalSeconds }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponseOutput) RetryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptionsResponse) *int { return v.RetryAttempts }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponseOutput) RetryIntervalBackoffMultiplier() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptionsResponse) *float64 { return v.RetryIntervalBackoffMultiplier }).(pulumi.Float64PtrOutput)
}

func (o JobStepExecutionOptionsResponseOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptionsResponse) *int { return v.TimeoutSeconds }).(pulumi.IntPtrOutput)
}

type JobStepExecutionOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (JobStepExecutionOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepExecutionOptionsResponse)(nil)).Elem()
}

func (o JobStepExecutionOptionsResponsePtrOutput) ToJobStepExecutionOptionsResponsePtrOutput() JobStepExecutionOptionsResponsePtrOutput {
	return o
}

func (o JobStepExecutionOptionsResponsePtrOutput) ToJobStepExecutionOptionsResponsePtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponsePtrOutput {
	return o
}

func (o JobStepExecutionOptionsResponsePtrOutput) Elem() JobStepExecutionOptionsResponseOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) JobStepExecutionOptionsResponse {
		if v != nil {
			return *v
		}
		var ret JobStepExecutionOptionsResponse
		return ret
	}).(JobStepExecutionOptionsResponseOutput)
}

func (o JobStepExecutionOptionsResponsePtrOutput) InitialRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.InitialRetryIntervalSeconds
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponsePtrOutput) MaximumRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaximumRetryIntervalSeconds
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponsePtrOutput) RetryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetryAttempts
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponsePtrOutput) RetryIntervalBackoffMultiplier() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryIntervalBackoffMultiplier
	}).(pulumi.Float64PtrOutput)
}

func (o JobStepExecutionOptionsResponsePtrOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.TimeoutSeconds
	}).(pulumi.IntPtrOutput)
}

type JobStepOutputType struct {
	Credential        string  `pulumi:"credential"`
	DatabaseName      string  `pulumi:"databaseName"`
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	SchemaName        *string `pulumi:"schemaName"`
	ServerName        string  `pulumi:"serverName"`
	SubscriptionId    *string `pulumi:"subscriptionId"`
	TableName         string  `pulumi:"tableName"`
	Type              *string `pulumi:"type"`
}





type JobStepOutputTypeInput interface {
	pulumi.Input

	ToJobStepOutputTypeOutput() JobStepOutputTypeOutput
	ToJobStepOutputTypeOutputWithContext(context.Context) JobStepOutputTypeOutput
}

type JobStepOutputTypeArgs struct {
	Credential        pulumi.StringInput    `pulumi:"credential"`
	DatabaseName      pulumi.StringInput    `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringPtrInput `pulumi:"resourceGroupName"`
	SchemaName        pulumi.StringPtrInput `pulumi:"schemaName"`
	ServerName        pulumi.StringInput    `pulumi:"serverName"`
	SubscriptionId    pulumi.StringPtrInput `pulumi:"subscriptionId"`
	TableName         pulumi.StringInput    `pulumi:"tableName"`
	Type              pulumi.StringPtrInput `pulumi:"type"`
}

func (JobStepOutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepOutputType)(nil)).Elem()
}

func (i JobStepOutputTypeArgs) ToJobStepOutputTypeOutput() JobStepOutputTypeOutput {
	return i.ToJobStepOutputTypeOutputWithContext(context.Background())
}

func (i JobStepOutputTypeArgs) ToJobStepOutputTypeOutputWithContext(ctx context.Context) JobStepOutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputTypeOutput)
}

func (i JobStepOutputTypeArgs) ToJobStepOutputTypePtrOutput() JobStepOutputTypePtrOutput {
	return i.ToJobStepOutputTypePtrOutputWithContext(context.Background())
}

func (i JobStepOutputTypeArgs) ToJobStepOutputTypePtrOutputWithContext(ctx context.Context) JobStepOutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputTypeOutput).ToJobStepOutputTypePtrOutputWithContext(ctx)
}









type JobStepOutputTypePtrInput interface {
	pulumi.Input

	ToJobStepOutputTypePtrOutput() JobStepOutputTypePtrOutput
	ToJobStepOutputTypePtrOutputWithContext(context.Context) JobStepOutputTypePtrOutput
}

type jobStepOutputTypePtrType JobStepOutputTypeArgs

func JobStepOutputTypePtr(v *JobStepOutputTypeArgs) JobStepOutputTypePtrInput {
	return (*jobStepOutputTypePtrType)(v)
}

func (*jobStepOutputTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepOutputType)(nil)).Elem()
}

func (i *jobStepOutputTypePtrType) ToJobStepOutputTypePtrOutput() JobStepOutputTypePtrOutput {
	return i.ToJobStepOutputTypePtrOutputWithContext(context.Background())
}

func (i *jobStepOutputTypePtrType) ToJobStepOutputTypePtrOutputWithContext(ctx context.Context) JobStepOutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputTypePtrOutput)
}

type JobStepOutputTypeOutput struct{ *pulumi.OutputState }

func (JobStepOutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepOutputType)(nil)).Elem()
}

func (o JobStepOutputTypeOutput) ToJobStepOutputTypeOutput() JobStepOutputTypeOutput {
	return o
}

func (o JobStepOutputTypeOutput) ToJobStepOutputTypeOutputWithContext(ctx context.Context) JobStepOutputTypeOutput {
	return o
}

func (o JobStepOutputTypeOutput) ToJobStepOutputTypePtrOutput() JobStepOutputTypePtrOutput {
	return o.ToJobStepOutputTypePtrOutputWithContext(context.Background())
}

func (o JobStepOutputTypeOutput) ToJobStepOutputTypePtrOutputWithContext(ctx context.Context) JobStepOutputTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepOutputType) *JobStepOutputType {
		return &v
	}).(JobStepOutputTypePtrOutput)
}

func (o JobStepOutputTypeOutput) Credential() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputType) string { return v.Credential }).(pulumi.StringOutput)
}

func (o JobStepOutputTypeOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputType) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o JobStepOutputTypeOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputType) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypeOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputType) *string { return v.SchemaName }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypeOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputType) string { return v.ServerName }).(pulumi.StringOutput)
}

func (o JobStepOutputTypeOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputType) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypeOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputType) string { return v.TableName }).(pulumi.StringOutput)
}

func (o JobStepOutputTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JobStepOutputTypePtrOutput struct{ *pulumi.OutputState }

func (JobStepOutputTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepOutputType)(nil)).Elem()
}

func (o JobStepOutputTypePtrOutput) ToJobStepOutputTypePtrOutput() JobStepOutputTypePtrOutput {
	return o
}

func (o JobStepOutputTypePtrOutput) ToJobStepOutputTypePtrOutputWithContext(ctx context.Context) JobStepOutputTypePtrOutput {
	return o
}

func (o JobStepOutputTypePtrOutput) Elem() JobStepOutputTypeOutput {
	return o.ApplyT(func(v *JobStepOutputType) JobStepOutputType {
		if v != nil {
			return *v
		}
		var ret JobStepOutputType
		return ret
	}).(JobStepOutputTypeOutput)
}

func (o JobStepOutputTypePtrOutput) Credential() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return &v.Credential
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return &v.DatabaseName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return v.SchemaName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return &v.ServerName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return &v.TableName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type JobStepOutputResponse struct {
	Credential        string  `pulumi:"credential"`
	DatabaseName      string  `pulumi:"databaseName"`
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	SchemaName        *string `pulumi:"schemaName"`
	ServerName        string  `pulumi:"serverName"`
	SubscriptionId    *string `pulumi:"subscriptionId"`
	TableName         string  `pulumi:"tableName"`
	Type              *string `pulumi:"type"`
}





type JobStepOutputResponseInput interface {
	pulumi.Input

	ToJobStepOutputResponseOutput() JobStepOutputResponseOutput
	ToJobStepOutputResponseOutputWithContext(context.Context) JobStepOutputResponseOutput
}

type JobStepOutputResponseArgs struct {
	Credential        pulumi.StringInput    `pulumi:"credential"`
	DatabaseName      pulumi.StringInput    `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringPtrInput `pulumi:"resourceGroupName"`
	SchemaName        pulumi.StringPtrInput `pulumi:"schemaName"`
	ServerName        pulumi.StringInput    `pulumi:"serverName"`
	SubscriptionId    pulumi.StringPtrInput `pulumi:"subscriptionId"`
	TableName         pulumi.StringInput    `pulumi:"tableName"`
	Type              pulumi.StringPtrInput `pulumi:"type"`
}

func (JobStepOutputResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepOutputResponse)(nil)).Elem()
}

func (i JobStepOutputResponseArgs) ToJobStepOutputResponseOutput() JobStepOutputResponseOutput {
	return i.ToJobStepOutputResponseOutputWithContext(context.Background())
}

func (i JobStepOutputResponseArgs) ToJobStepOutputResponseOutputWithContext(ctx context.Context) JobStepOutputResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputResponseOutput)
}

func (i JobStepOutputResponseArgs) ToJobStepOutputResponsePtrOutput() JobStepOutputResponsePtrOutput {
	return i.ToJobStepOutputResponsePtrOutputWithContext(context.Background())
}

func (i JobStepOutputResponseArgs) ToJobStepOutputResponsePtrOutputWithContext(ctx context.Context) JobStepOutputResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputResponseOutput).ToJobStepOutputResponsePtrOutputWithContext(ctx)
}









type JobStepOutputResponsePtrInput interface {
	pulumi.Input

	ToJobStepOutputResponsePtrOutput() JobStepOutputResponsePtrOutput
	ToJobStepOutputResponsePtrOutputWithContext(context.Context) JobStepOutputResponsePtrOutput
}

type jobStepOutputResponsePtrType JobStepOutputResponseArgs

func JobStepOutputResponsePtr(v *JobStepOutputResponseArgs) JobStepOutputResponsePtrInput {
	return (*jobStepOutputResponsePtrType)(v)
}

func (*jobStepOutputResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepOutputResponse)(nil)).Elem()
}

func (i *jobStepOutputResponsePtrType) ToJobStepOutputResponsePtrOutput() JobStepOutputResponsePtrOutput {
	return i.ToJobStepOutputResponsePtrOutputWithContext(context.Background())
}

func (i *jobStepOutputResponsePtrType) ToJobStepOutputResponsePtrOutputWithContext(ctx context.Context) JobStepOutputResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputResponsePtrOutput)
}

type JobStepOutputResponseOutput struct{ *pulumi.OutputState }

func (JobStepOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepOutputResponse)(nil)).Elem()
}

func (o JobStepOutputResponseOutput) ToJobStepOutputResponseOutput() JobStepOutputResponseOutput {
	return o
}

func (o JobStepOutputResponseOutput) ToJobStepOutputResponseOutputWithContext(ctx context.Context) JobStepOutputResponseOutput {
	return o
}

func (o JobStepOutputResponseOutput) ToJobStepOutputResponsePtrOutput() JobStepOutputResponsePtrOutput {
	return o.ToJobStepOutputResponsePtrOutputWithContext(context.Background())
}

func (o JobStepOutputResponseOutput) ToJobStepOutputResponsePtrOutputWithContext(ctx context.Context) JobStepOutputResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepOutputResponse) *JobStepOutputResponse {
		return &v
	}).(JobStepOutputResponsePtrOutput)
}

func (o JobStepOutputResponseOutput) Credential() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputResponse) string { return v.Credential }).(pulumi.StringOutput)
}

func (o JobStepOutputResponseOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputResponse) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o JobStepOutputResponseOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputResponse) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponseOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputResponse) *string { return v.SchemaName }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponseOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputResponse) string { return v.ServerName }).(pulumi.StringOutput)
}

func (o JobStepOutputResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponseOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputResponse) string { return v.TableName }).(pulumi.StringOutput)
}

func (o JobStepOutputResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JobStepOutputResponsePtrOutput struct{ *pulumi.OutputState }

func (JobStepOutputResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepOutputResponse)(nil)).Elem()
}

func (o JobStepOutputResponsePtrOutput) ToJobStepOutputResponsePtrOutput() JobStepOutputResponsePtrOutput {
	return o
}

func (o JobStepOutputResponsePtrOutput) ToJobStepOutputResponsePtrOutputWithContext(ctx context.Context) JobStepOutputResponsePtrOutput {
	return o
}

func (o JobStepOutputResponsePtrOutput) Elem() JobStepOutputResponseOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) JobStepOutputResponse {
		if v != nil {
			return *v
		}
		var ret JobStepOutputResponse
		return ret
	}).(JobStepOutputResponseOutput)
}

func (o JobStepOutputResponsePtrOutput) Credential() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Credential
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DatabaseName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return v.SchemaName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServerName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TableName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type JobTarget struct {
	DatabaseName      *string                       `pulumi:"databaseName"`
	ElasticPoolName   *string                       `pulumi:"elasticPoolName"`
	MembershipType    *JobTargetGroupMembershipType `pulumi:"membershipType"`
	RefreshCredential *string                       `pulumi:"refreshCredential"`
	ServerName        *string                       `pulumi:"serverName"`
	ShardMapName      *string                       `pulumi:"shardMapName"`
	Type              string                        `pulumi:"type"`
}





type JobTargetInput interface {
	pulumi.Input

	ToJobTargetOutput() JobTargetOutput
	ToJobTargetOutputWithContext(context.Context) JobTargetOutput
}

type JobTargetArgs struct {
	DatabaseName      pulumi.StringPtrInput                `pulumi:"databaseName"`
	ElasticPoolName   pulumi.StringPtrInput                `pulumi:"elasticPoolName"`
	MembershipType    JobTargetGroupMembershipTypePtrInput `pulumi:"membershipType"`
	RefreshCredential pulumi.StringPtrInput                `pulumi:"refreshCredential"`
	ServerName        pulumi.StringPtrInput                `pulumi:"serverName"`
	ShardMapName      pulumi.StringPtrInput                `pulumi:"shardMapName"`
	Type              pulumi.StringInput                   `pulumi:"type"`
}

func (JobTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTarget)(nil)).Elem()
}

func (i JobTargetArgs) ToJobTargetOutput() JobTargetOutput {
	return i.ToJobTargetOutputWithContext(context.Background())
}

func (i JobTargetArgs) ToJobTargetOutputWithContext(ctx context.Context) JobTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTargetOutput)
}





type JobTargetArrayInput interface {
	pulumi.Input

	ToJobTargetArrayOutput() JobTargetArrayOutput
	ToJobTargetArrayOutputWithContext(context.Context) JobTargetArrayOutput
}

type JobTargetArray []JobTargetInput

func (JobTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobTarget)(nil)).Elem()
}

func (i JobTargetArray) ToJobTargetArrayOutput() JobTargetArrayOutput {
	return i.ToJobTargetArrayOutputWithContext(context.Background())
}

func (i JobTargetArray) ToJobTargetArrayOutputWithContext(ctx context.Context) JobTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTargetArrayOutput)
}

type JobTargetOutput struct{ *pulumi.OutputState }

func (JobTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTarget)(nil)).Elem()
}

func (o JobTargetOutput) ToJobTargetOutput() JobTargetOutput {
	return o
}

func (o JobTargetOutput) ToJobTargetOutputWithContext(ctx context.Context) JobTargetOutput {
	return o
}

func (o JobTargetOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTarget) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o JobTargetOutput) ElasticPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTarget) *string { return v.ElasticPoolName }).(pulumi.StringPtrOutput)
}

func (o JobTargetOutput) MembershipType() JobTargetGroupMembershipTypePtrOutput {
	return o.ApplyT(func(v JobTarget) *JobTargetGroupMembershipType { return v.MembershipType }).(JobTargetGroupMembershipTypePtrOutput)
}

func (o JobTargetOutput) RefreshCredential() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTarget) *string { return v.RefreshCredential }).(pulumi.StringPtrOutput)
}

func (o JobTargetOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTarget) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o JobTargetOutput) ShardMapName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTarget) *string { return v.ShardMapName }).(pulumi.StringPtrOutput)
}

func (o JobTargetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JobTarget) string { return v.Type }).(pulumi.StringOutput)
}

type JobTargetArrayOutput struct{ *pulumi.OutputState }

func (JobTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobTarget)(nil)).Elem()
}

func (o JobTargetArrayOutput) ToJobTargetArrayOutput() JobTargetArrayOutput {
	return o
}

func (o JobTargetArrayOutput) ToJobTargetArrayOutputWithContext(ctx context.Context) JobTargetArrayOutput {
	return o
}

func (o JobTargetArrayOutput) Index(i pulumi.IntInput) JobTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobTarget {
		return vs[0].([]JobTarget)[vs[1].(int)]
	}).(JobTargetOutput)
}

type JobTargetResponse struct {
	DatabaseName      *string `pulumi:"databaseName"`
	ElasticPoolName   *string `pulumi:"elasticPoolName"`
	MembershipType    *string `pulumi:"membershipType"`
	RefreshCredential *string `pulumi:"refreshCredential"`
	ServerName        *string `pulumi:"serverName"`
	ShardMapName      *string `pulumi:"shardMapName"`
	Type              string  `pulumi:"type"`
}





type JobTargetResponseInput interface {
	pulumi.Input

	ToJobTargetResponseOutput() JobTargetResponseOutput
	ToJobTargetResponseOutputWithContext(context.Context) JobTargetResponseOutput
}

type JobTargetResponseArgs struct {
	DatabaseName      pulumi.StringPtrInput `pulumi:"databaseName"`
	ElasticPoolName   pulumi.StringPtrInput `pulumi:"elasticPoolName"`
	MembershipType    pulumi.StringPtrInput `pulumi:"membershipType"`
	RefreshCredential pulumi.StringPtrInput `pulumi:"refreshCredential"`
	ServerName        pulumi.StringPtrInput `pulumi:"serverName"`
	ShardMapName      pulumi.StringPtrInput `pulumi:"shardMapName"`
	Type              pulumi.StringInput    `pulumi:"type"`
}

func (JobTargetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTargetResponse)(nil)).Elem()
}

func (i JobTargetResponseArgs) ToJobTargetResponseOutput() JobTargetResponseOutput {
	return i.ToJobTargetResponseOutputWithContext(context.Background())
}

func (i JobTargetResponseArgs) ToJobTargetResponseOutputWithContext(ctx context.Context) JobTargetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTargetResponseOutput)
}





type JobTargetResponseArrayInput interface {
	pulumi.Input

	ToJobTargetResponseArrayOutput() JobTargetResponseArrayOutput
	ToJobTargetResponseArrayOutputWithContext(context.Context) JobTargetResponseArrayOutput
}

type JobTargetResponseArray []JobTargetResponseInput

func (JobTargetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobTargetResponse)(nil)).Elem()
}

func (i JobTargetResponseArray) ToJobTargetResponseArrayOutput() JobTargetResponseArrayOutput {
	return i.ToJobTargetResponseArrayOutputWithContext(context.Background())
}

func (i JobTargetResponseArray) ToJobTargetResponseArrayOutputWithContext(ctx context.Context) JobTargetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTargetResponseArrayOutput)
}

type JobTargetResponseOutput struct{ *pulumi.OutputState }

func (JobTargetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTargetResponse)(nil)).Elem()
}

func (o JobTargetResponseOutput) ToJobTargetResponseOutput() JobTargetResponseOutput {
	return o
}

func (o JobTargetResponseOutput) ToJobTargetResponseOutputWithContext(ctx context.Context) JobTargetResponseOutput {
	return o
}

func (o JobTargetResponseOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) ElasticPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.ElasticPoolName }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) MembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.MembershipType }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) RefreshCredential() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.RefreshCredential }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) ShardMapName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.ShardMapName }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JobTargetResponse) string { return v.Type }).(pulumi.StringOutput)
}

type JobTargetResponseArrayOutput struct{ *pulumi.OutputState }

func (JobTargetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobTargetResponse)(nil)).Elem()
}

func (o JobTargetResponseArrayOutput) ToJobTargetResponseArrayOutput() JobTargetResponseArrayOutput {
	return o
}

func (o JobTargetResponseArrayOutput) ToJobTargetResponseArrayOutputWithContext(ctx context.Context) JobTargetResponseArrayOutput {
	return o
}

func (o JobTargetResponseArrayOutput) Index(i pulumi.IntInput) JobTargetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobTargetResponse {
		return vs[0].([]JobTargetResponse)[vs[1].(int)]
	}).(JobTargetResponseOutput)
}

type ManagedInstanceExternalAdministrator struct {
	AdministratorType         *string `pulumi:"administratorType"`
	AzureADOnlyAuthentication *bool   `pulumi:"azureADOnlyAuthentication"`
	Login                     *string `pulumi:"login"`
	PrincipalType             *string `pulumi:"principalType"`
	Sid                       *string `pulumi:"sid"`
	TenantId                  *string `pulumi:"tenantId"`
}





type ManagedInstanceExternalAdministratorInput interface {
	pulumi.Input

	ToManagedInstanceExternalAdministratorOutput() ManagedInstanceExternalAdministratorOutput
	ToManagedInstanceExternalAdministratorOutputWithContext(context.Context) ManagedInstanceExternalAdministratorOutput
}

type ManagedInstanceExternalAdministratorArgs struct {
	AdministratorType         pulumi.StringPtrInput `pulumi:"administratorType"`
	AzureADOnlyAuthentication pulumi.BoolPtrInput   `pulumi:"azureADOnlyAuthentication"`
	Login                     pulumi.StringPtrInput `pulumi:"login"`
	PrincipalType             pulumi.StringPtrInput `pulumi:"principalType"`
	Sid                       pulumi.StringPtrInput `pulumi:"sid"`
	TenantId                  pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (ManagedInstanceExternalAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceExternalAdministrator)(nil)).Elem()
}

func (i ManagedInstanceExternalAdministratorArgs) ToManagedInstanceExternalAdministratorOutput() ManagedInstanceExternalAdministratorOutput {
	return i.ToManagedInstanceExternalAdministratorOutputWithContext(context.Background())
}

func (i ManagedInstanceExternalAdministratorArgs) ToManagedInstanceExternalAdministratorOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceExternalAdministratorOutput)
}

func (i ManagedInstanceExternalAdministratorArgs) ToManagedInstanceExternalAdministratorPtrOutput() ManagedInstanceExternalAdministratorPtrOutput {
	return i.ToManagedInstanceExternalAdministratorPtrOutputWithContext(context.Background())
}

func (i ManagedInstanceExternalAdministratorArgs) ToManagedInstanceExternalAdministratorPtrOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceExternalAdministratorOutput).ToManagedInstanceExternalAdministratorPtrOutputWithContext(ctx)
}









type ManagedInstanceExternalAdministratorPtrInput interface {
	pulumi.Input

	ToManagedInstanceExternalAdministratorPtrOutput() ManagedInstanceExternalAdministratorPtrOutput
	ToManagedInstanceExternalAdministratorPtrOutputWithContext(context.Context) ManagedInstanceExternalAdministratorPtrOutput
}

type managedInstanceExternalAdministratorPtrType ManagedInstanceExternalAdministratorArgs

func ManagedInstanceExternalAdministratorPtr(v *ManagedInstanceExternalAdministratorArgs) ManagedInstanceExternalAdministratorPtrInput {
	return (*managedInstanceExternalAdministratorPtrType)(v)
}

func (*managedInstanceExternalAdministratorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceExternalAdministrator)(nil)).Elem()
}

func (i *managedInstanceExternalAdministratorPtrType) ToManagedInstanceExternalAdministratorPtrOutput() ManagedInstanceExternalAdministratorPtrOutput {
	return i.ToManagedInstanceExternalAdministratorPtrOutputWithContext(context.Background())
}

func (i *managedInstanceExternalAdministratorPtrType) ToManagedInstanceExternalAdministratorPtrOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceExternalAdministratorPtrOutput)
}

type ManagedInstanceExternalAdministratorOutput struct{ *pulumi.OutputState }

func (ManagedInstanceExternalAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceExternalAdministrator)(nil)).Elem()
}

func (o ManagedInstanceExternalAdministratorOutput) ToManagedInstanceExternalAdministratorOutput() ManagedInstanceExternalAdministratorOutput {
	return o
}

func (o ManagedInstanceExternalAdministratorOutput) ToManagedInstanceExternalAdministratorOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorOutput {
	return o
}

func (o ManagedInstanceExternalAdministratorOutput) ToManagedInstanceExternalAdministratorPtrOutput() ManagedInstanceExternalAdministratorPtrOutput {
	return o.ToManagedInstanceExternalAdministratorPtrOutputWithContext(context.Background())
}

func (o ManagedInstanceExternalAdministratorOutput) ToManagedInstanceExternalAdministratorPtrOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedInstanceExternalAdministrator) *ManagedInstanceExternalAdministrator {
		return &v
	}).(ManagedInstanceExternalAdministratorPtrOutput)
}

func (o ManagedInstanceExternalAdministratorOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministrator) *string { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorOutput) AzureADOnlyAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministrator) *bool { return v.AzureADOnlyAuthentication }).(pulumi.BoolPtrOutput)
}

func (o ManagedInstanceExternalAdministratorOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministrator) *string { return v.Login }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministrator) *string { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministrator) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministrator) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ManagedInstanceExternalAdministratorPtrOutput struct{ *pulumi.OutputState }

func (ManagedInstanceExternalAdministratorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceExternalAdministrator)(nil)).Elem()
}

func (o ManagedInstanceExternalAdministratorPtrOutput) ToManagedInstanceExternalAdministratorPtrOutput() ManagedInstanceExternalAdministratorPtrOutput {
	return o
}

func (o ManagedInstanceExternalAdministratorPtrOutput) ToManagedInstanceExternalAdministratorPtrOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorPtrOutput {
	return o
}

func (o ManagedInstanceExternalAdministratorPtrOutput) Elem() ManagedInstanceExternalAdministratorOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministrator) ManagedInstanceExternalAdministrator {
		if v != nil {
			return *v
		}
		var ret ManagedInstanceExternalAdministrator
		return ret
	}).(ManagedInstanceExternalAdministratorOutput)
}

func (o ManagedInstanceExternalAdministratorPtrOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministrator) *string {
		if v == nil {
			return nil
		}
		return v.AdministratorType
	}).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorPtrOutput) AzureADOnlyAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministrator) *bool {
		if v == nil {
			return nil
		}
		return v.AzureADOnlyAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedInstanceExternalAdministratorPtrOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministrator) *string {
		if v == nil {
			return nil
		}
		return v.Login
	}).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorPtrOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministrator) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalType
	}).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorPtrOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministrator) *string {
		if v == nil {
			return nil
		}
		return v.Sid
	}).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministrator) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ManagedInstanceExternalAdministratorResponse struct {
	AdministratorType         *string `pulumi:"administratorType"`
	AzureADOnlyAuthentication *bool   `pulumi:"azureADOnlyAuthentication"`
	Login                     *string `pulumi:"login"`
	PrincipalType             *string `pulumi:"principalType"`
	Sid                       *string `pulumi:"sid"`
	TenantId                  *string `pulumi:"tenantId"`
}





type ManagedInstanceExternalAdministratorResponseInput interface {
	pulumi.Input

	ToManagedInstanceExternalAdministratorResponseOutput() ManagedInstanceExternalAdministratorResponseOutput
	ToManagedInstanceExternalAdministratorResponseOutputWithContext(context.Context) ManagedInstanceExternalAdministratorResponseOutput
}

type ManagedInstanceExternalAdministratorResponseArgs struct {
	AdministratorType         pulumi.StringPtrInput `pulumi:"administratorType"`
	AzureADOnlyAuthentication pulumi.BoolPtrInput   `pulumi:"azureADOnlyAuthentication"`
	Login                     pulumi.StringPtrInput `pulumi:"login"`
	PrincipalType             pulumi.StringPtrInput `pulumi:"principalType"`
	Sid                       pulumi.StringPtrInput `pulumi:"sid"`
	TenantId                  pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (ManagedInstanceExternalAdministratorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceExternalAdministratorResponse)(nil)).Elem()
}

func (i ManagedInstanceExternalAdministratorResponseArgs) ToManagedInstanceExternalAdministratorResponseOutput() ManagedInstanceExternalAdministratorResponseOutput {
	return i.ToManagedInstanceExternalAdministratorResponseOutputWithContext(context.Background())
}

func (i ManagedInstanceExternalAdministratorResponseArgs) ToManagedInstanceExternalAdministratorResponseOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceExternalAdministratorResponseOutput)
}

func (i ManagedInstanceExternalAdministratorResponseArgs) ToManagedInstanceExternalAdministratorResponsePtrOutput() ManagedInstanceExternalAdministratorResponsePtrOutput {
	return i.ToManagedInstanceExternalAdministratorResponsePtrOutputWithContext(context.Background())
}

func (i ManagedInstanceExternalAdministratorResponseArgs) ToManagedInstanceExternalAdministratorResponsePtrOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceExternalAdministratorResponseOutput).ToManagedInstanceExternalAdministratorResponsePtrOutputWithContext(ctx)
}









type ManagedInstanceExternalAdministratorResponsePtrInput interface {
	pulumi.Input

	ToManagedInstanceExternalAdministratorResponsePtrOutput() ManagedInstanceExternalAdministratorResponsePtrOutput
	ToManagedInstanceExternalAdministratorResponsePtrOutputWithContext(context.Context) ManagedInstanceExternalAdministratorResponsePtrOutput
}

type managedInstanceExternalAdministratorResponsePtrType ManagedInstanceExternalAdministratorResponseArgs

func ManagedInstanceExternalAdministratorResponsePtr(v *ManagedInstanceExternalAdministratorResponseArgs) ManagedInstanceExternalAdministratorResponsePtrInput {
	return (*managedInstanceExternalAdministratorResponsePtrType)(v)
}

func (*managedInstanceExternalAdministratorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceExternalAdministratorResponse)(nil)).Elem()
}

func (i *managedInstanceExternalAdministratorResponsePtrType) ToManagedInstanceExternalAdministratorResponsePtrOutput() ManagedInstanceExternalAdministratorResponsePtrOutput {
	return i.ToManagedInstanceExternalAdministratorResponsePtrOutputWithContext(context.Background())
}

func (i *managedInstanceExternalAdministratorResponsePtrType) ToManagedInstanceExternalAdministratorResponsePtrOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceExternalAdministratorResponsePtrOutput)
}

type ManagedInstanceExternalAdministratorResponseOutput struct{ *pulumi.OutputState }

func (ManagedInstanceExternalAdministratorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceExternalAdministratorResponse)(nil)).Elem()
}

func (o ManagedInstanceExternalAdministratorResponseOutput) ToManagedInstanceExternalAdministratorResponseOutput() ManagedInstanceExternalAdministratorResponseOutput {
	return o
}

func (o ManagedInstanceExternalAdministratorResponseOutput) ToManagedInstanceExternalAdministratorResponseOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorResponseOutput {
	return o
}

func (o ManagedInstanceExternalAdministratorResponseOutput) ToManagedInstanceExternalAdministratorResponsePtrOutput() ManagedInstanceExternalAdministratorResponsePtrOutput {
	return o.ToManagedInstanceExternalAdministratorResponsePtrOutputWithContext(context.Background())
}

func (o ManagedInstanceExternalAdministratorResponseOutput) ToManagedInstanceExternalAdministratorResponsePtrOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedInstanceExternalAdministratorResponse) *ManagedInstanceExternalAdministratorResponse {
		return &v
	}).(ManagedInstanceExternalAdministratorResponsePtrOutput)
}

func (o ManagedInstanceExternalAdministratorResponseOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministratorResponse) *string { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorResponseOutput) AzureADOnlyAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministratorResponse) *bool { return v.AzureADOnlyAuthentication }).(pulumi.BoolPtrOutput)
}

func (o ManagedInstanceExternalAdministratorResponseOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministratorResponse) *string { return v.Login }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorResponseOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministratorResponse) *string { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorResponseOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministratorResponse) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstanceExternalAdministratorResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ManagedInstanceExternalAdministratorResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedInstanceExternalAdministratorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceExternalAdministratorResponse)(nil)).Elem()
}

func (o ManagedInstanceExternalAdministratorResponsePtrOutput) ToManagedInstanceExternalAdministratorResponsePtrOutput() ManagedInstanceExternalAdministratorResponsePtrOutput {
	return o
}

func (o ManagedInstanceExternalAdministratorResponsePtrOutput) ToManagedInstanceExternalAdministratorResponsePtrOutputWithContext(ctx context.Context) ManagedInstanceExternalAdministratorResponsePtrOutput {
	return o
}

func (o ManagedInstanceExternalAdministratorResponsePtrOutput) Elem() ManagedInstanceExternalAdministratorResponseOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministratorResponse) ManagedInstanceExternalAdministratorResponse {
		if v != nil {
			return *v
		}
		var ret ManagedInstanceExternalAdministratorResponse
		return ret
	}).(ManagedInstanceExternalAdministratorResponseOutput)
}

func (o ManagedInstanceExternalAdministratorResponsePtrOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministratorResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdministratorType
	}).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorResponsePtrOutput) AzureADOnlyAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministratorResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AzureADOnlyAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedInstanceExternalAdministratorResponsePtrOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministratorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Login
	}).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorResponsePtrOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministratorResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalType
	}).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorResponsePtrOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministratorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sid
	}).(pulumi.StringPtrOutput)
}

func (o ManagedInstanceExternalAdministratorResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceExternalAdministratorResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ManagedInstancePairInfo struct {
	PartnerManagedInstanceId *string `pulumi:"partnerManagedInstanceId"`
	PrimaryManagedInstanceId *string `pulumi:"primaryManagedInstanceId"`
}





type ManagedInstancePairInfoInput interface {
	pulumi.Input

	ToManagedInstancePairInfoOutput() ManagedInstancePairInfoOutput
	ToManagedInstancePairInfoOutputWithContext(context.Context) ManagedInstancePairInfoOutput
}

type ManagedInstancePairInfoArgs struct {
	PartnerManagedInstanceId pulumi.StringPtrInput `pulumi:"partnerManagedInstanceId"`
	PrimaryManagedInstanceId pulumi.StringPtrInput `pulumi:"primaryManagedInstanceId"`
}

func (ManagedInstancePairInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePairInfo)(nil)).Elem()
}

func (i ManagedInstancePairInfoArgs) ToManagedInstancePairInfoOutput() ManagedInstancePairInfoOutput {
	return i.ToManagedInstancePairInfoOutputWithContext(context.Background())
}

func (i ManagedInstancePairInfoArgs) ToManagedInstancePairInfoOutputWithContext(ctx context.Context) ManagedInstancePairInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePairInfoOutput)
}





type ManagedInstancePairInfoArrayInput interface {
	pulumi.Input

	ToManagedInstancePairInfoArrayOutput() ManagedInstancePairInfoArrayOutput
	ToManagedInstancePairInfoArrayOutputWithContext(context.Context) ManagedInstancePairInfoArrayOutput
}

type ManagedInstancePairInfoArray []ManagedInstancePairInfoInput

func (ManagedInstancePairInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedInstancePairInfo)(nil)).Elem()
}

func (i ManagedInstancePairInfoArray) ToManagedInstancePairInfoArrayOutput() ManagedInstancePairInfoArrayOutput {
	return i.ToManagedInstancePairInfoArrayOutputWithContext(context.Background())
}

func (i ManagedInstancePairInfoArray) ToManagedInstancePairInfoArrayOutputWithContext(ctx context.Context) ManagedInstancePairInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePairInfoArrayOutput)
}

type ManagedInstancePairInfoOutput struct{ *pulumi.OutputState }

func (ManagedInstancePairInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePairInfo)(nil)).Elem()
}

func (o ManagedInstancePairInfoOutput) ToManagedInstancePairInfoOutput() ManagedInstancePairInfoOutput {
	return o
}

func (o ManagedInstancePairInfoOutput) ToManagedInstancePairInfoOutputWithContext(ctx context.Context) ManagedInstancePairInfoOutput {
	return o
}

func (o ManagedInstancePairInfoOutput) PartnerManagedInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstancePairInfo) *string { return v.PartnerManagedInstanceId }).(pulumi.StringPtrOutput)
}

func (o ManagedInstancePairInfoOutput) PrimaryManagedInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstancePairInfo) *string { return v.PrimaryManagedInstanceId }).(pulumi.StringPtrOutput)
}

type ManagedInstancePairInfoArrayOutput struct{ *pulumi.OutputState }

func (ManagedInstancePairInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedInstancePairInfo)(nil)).Elem()
}

func (o ManagedInstancePairInfoArrayOutput) ToManagedInstancePairInfoArrayOutput() ManagedInstancePairInfoArrayOutput {
	return o
}

func (o ManagedInstancePairInfoArrayOutput) ToManagedInstancePairInfoArrayOutputWithContext(ctx context.Context) ManagedInstancePairInfoArrayOutput {
	return o
}

func (o ManagedInstancePairInfoArrayOutput) Index(i pulumi.IntInput) ManagedInstancePairInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedInstancePairInfo {
		return vs[0].([]ManagedInstancePairInfo)[vs[1].(int)]
	}).(ManagedInstancePairInfoOutput)
}

type ManagedInstancePairInfoResponse struct {
	PartnerManagedInstanceId *string `pulumi:"partnerManagedInstanceId"`
	PrimaryManagedInstanceId *string `pulumi:"primaryManagedInstanceId"`
}





type ManagedInstancePairInfoResponseInput interface {
	pulumi.Input

	ToManagedInstancePairInfoResponseOutput() ManagedInstancePairInfoResponseOutput
	ToManagedInstancePairInfoResponseOutputWithContext(context.Context) ManagedInstancePairInfoResponseOutput
}

type ManagedInstancePairInfoResponseArgs struct {
	PartnerManagedInstanceId pulumi.StringPtrInput `pulumi:"partnerManagedInstanceId"`
	PrimaryManagedInstanceId pulumi.StringPtrInput `pulumi:"primaryManagedInstanceId"`
}

func (ManagedInstancePairInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePairInfoResponse)(nil)).Elem()
}

func (i ManagedInstancePairInfoResponseArgs) ToManagedInstancePairInfoResponseOutput() ManagedInstancePairInfoResponseOutput {
	return i.ToManagedInstancePairInfoResponseOutputWithContext(context.Background())
}

func (i ManagedInstancePairInfoResponseArgs) ToManagedInstancePairInfoResponseOutputWithContext(ctx context.Context) ManagedInstancePairInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePairInfoResponseOutput)
}





type ManagedInstancePairInfoResponseArrayInput interface {
	pulumi.Input

	ToManagedInstancePairInfoResponseArrayOutput() ManagedInstancePairInfoResponseArrayOutput
	ToManagedInstancePairInfoResponseArrayOutputWithContext(context.Context) ManagedInstancePairInfoResponseArrayOutput
}

type ManagedInstancePairInfoResponseArray []ManagedInstancePairInfoResponseInput

func (ManagedInstancePairInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedInstancePairInfoResponse)(nil)).Elem()
}

func (i ManagedInstancePairInfoResponseArray) ToManagedInstancePairInfoResponseArrayOutput() ManagedInstancePairInfoResponseArrayOutput {
	return i.ToManagedInstancePairInfoResponseArrayOutputWithContext(context.Background())
}

func (i ManagedInstancePairInfoResponseArray) ToManagedInstancePairInfoResponseArrayOutputWithContext(ctx context.Context) ManagedInstancePairInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePairInfoResponseArrayOutput)
}

type ManagedInstancePairInfoResponseOutput struct{ *pulumi.OutputState }

func (ManagedInstancePairInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePairInfoResponse)(nil)).Elem()
}

func (o ManagedInstancePairInfoResponseOutput) ToManagedInstancePairInfoResponseOutput() ManagedInstancePairInfoResponseOutput {
	return o
}

func (o ManagedInstancePairInfoResponseOutput) ToManagedInstancePairInfoResponseOutputWithContext(ctx context.Context) ManagedInstancePairInfoResponseOutput {
	return o
}

func (o ManagedInstancePairInfoResponseOutput) PartnerManagedInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstancePairInfoResponse) *string { return v.PartnerManagedInstanceId }).(pulumi.StringPtrOutput)
}

func (o ManagedInstancePairInfoResponseOutput) PrimaryManagedInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstancePairInfoResponse) *string { return v.PrimaryManagedInstanceId }).(pulumi.StringPtrOutput)
}

type ManagedInstancePairInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedInstancePairInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedInstancePairInfoResponse)(nil)).Elem()
}

func (o ManagedInstancePairInfoResponseArrayOutput) ToManagedInstancePairInfoResponseArrayOutput() ManagedInstancePairInfoResponseArrayOutput {
	return o
}

func (o ManagedInstancePairInfoResponseArrayOutput) ToManagedInstancePairInfoResponseArrayOutputWithContext(ctx context.Context) ManagedInstancePairInfoResponseArrayOutput {
	return o
}

func (o ManagedInstancePairInfoResponseArrayOutput) Index(i pulumi.IntInput) ManagedInstancePairInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedInstancePairInfoResponse {
		return vs[0].([]ManagedInstancePairInfoResponse)[vs[1].(int)]
	}).(ManagedInstancePairInfoResponseOutput)
}

type ManagedInstancePecPropertyResponse struct {
	Id         string                                                     `pulumi:"id"`
	Properties ManagedInstancePrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
}





type ManagedInstancePecPropertyResponseInput interface {
	pulumi.Input

	ToManagedInstancePecPropertyResponseOutput() ManagedInstancePecPropertyResponseOutput
	ToManagedInstancePecPropertyResponseOutputWithContext(context.Context) ManagedInstancePecPropertyResponseOutput
}

type ManagedInstancePecPropertyResponseArgs struct {
	Id         pulumi.StringInput                                              `pulumi:"id"`
	Properties ManagedInstancePrivateEndpointConnectionPropertiesResponseInput `pulumi:"properties"`
}

func (ManagedInstancePecPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePecPropertyResponse)(nil)).Elem()
}

func (i ManagedInstancePecPropertyResponseArgs) ToManagedInstancePecPropertyResponseOutput() ManagedInstancePecPropertyResponseOutput {
	return i.ToManagedInstancePecPropertyResponseOutputWithContext(context.Background())
}

func (i ManagedInstancePecPropertyResponseArgs) ToManagedInstancePecPropertyResponseOutputWithContext(ctx context.Context) ManagedInstancePecPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePecPropertyResponseOutput)
}





type ManagedInstancePecPropertyResponseArrayInput interface {
	pulumi.Input

	ToManagedInstancePecPropertyResponseArrayOutput() ManagedInstancePecPropertyResponseArrayOutput
	ToManagedInstancePecPropertyResponseArrayOutputWithContext(context.Context) ManagedInstancePecPropertyResponseArrayOutput
}

type ManagedInstancePecPropertyResponseArray []ManagedInstancePecPropertyResponseInput

func (ManagedInstancePecPropertyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedInstancePecPropertyResponse)(nil)).Elem()
}

func (i ManagedInstancePecPropertyResponseArray) ToManagedInstancePecPropertyResponseArrayOutput() ManagedInstancePecPropertyResponseArrayOutput {
	return i.ToManagedInstancePecPropertyResponseArrayOutputWithContext(context.Background())
}

func (i ManagedInstancePecPropertyResponseArray) ToManagedInstancePecPropertyResponseArrayOutputWithContext(ctx context.Context) ManagedInstancePecPropertyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePecPropertyResponseArrayOutput)
}

type ManagedInstancePecPropertyResponseOutput struct{ *pulumi.OutputState }

func (ManagedInstancePecPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePecPropertyResponse)(nil)).Elem()
}

func (o ManagedInstancePecPropertyResponseOutput) ToManagedInstancePecPropertyResponseOutput() ManagedInstancePecPropertyResponseOutput {
	return o
}

func (o ManagedInstancePecPropertyResponseOutput) ToManagedInstancePecPropertyResponseOutputWithContext(ctx context.Context) ManagedInstancePecPropertyResponseOutput {
	return o
}

func (o ManagedInstancePecPropertyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedInstancePecPropertyResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ManagedInstancePecPropertyResponseOutput) Properties() ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v ManagedInstancePecPropertyResponse) ManagedInstancePrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput)
}

type ManagedInstancePecPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedInstancePecPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedInstancePecPropertyResponse)(nil)).Elem()
}

func (o ManagedInstancePecPropertyResponseArrayOutput) ToManagedInstancePecPropertyResponseArrayOutput() ManagedInstancePecPropertyResponseArrayOutput {
	return o
}

func (o ManagedInstancePecPropertyResponseArrayOutput) ToManagedInstancePecPropertyResponseArrayOutputWithContext(ctx context.Context) ManagedInstancePecPropertyResponseArrayOutput {
	return o
}

func (o ManagedInstancePecPropertyResponseArrayOutput) Index(i pulumi.IntInput) ManagedInstancePecPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedInstancePecPropertyResponse {
		return vs[0].([]ManagedInstancePecPropertyResponse)[vs[1].(int)]
	}).(ManagedInstancePecPropertyResponseOutput)
}

type ManagedInstancePrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *ManagedInstancePrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                                            `pulumi:"provisioningState"`
}





type ManagedInstancePrivateEndpointConnectionPropertiesResponseInput interface {
	pulumi.Input

	ToManagedInstancePrivateEndpointConnectionPropertiesResponseOutput() ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput
	ToManagedInstancePrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Context) ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput
}

type ManagedInstancePrivateEndpointConnectionPropertiesResponseArgs struct {
	PrivateEndpoint                   ManagedInstancePrivateEndpointPropertyResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                                       `pulumi:"provisioningState"`
}

func (ManagedInstancePrivateEndpointConnectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i ManagedInstancePrivateEndpointConnectionPropertiesResponseArgs) ToManagedInstancePrivateEndpointConnectionPropertiesResponseOutput() ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput {
	return i.ToManagedInstancePrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Background())
}

func (i ManagedInstancePrivateEndpointConnectionPropertiesResponseArgs) ToManagedInstancePrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput)
}

type ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput) ToManagedInstancePrivateEndpointConnectionPropertiesResponseOutput() ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput) ToManagedInstancePrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v ManagedInstancePrivateEndpointConnectionPropertiesResponse) *ManagedInstancePrivateEndpointPropertyResponse {
		return v.PrivateEndpoint
	}).(ManagedInstancePrivateEndpointPropertyResponsePtrOutput)
}

func (o ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v ManagedInstancePrivateEndpointConnectionPropertiesResponse) *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedInstancePrivateEndpointConnectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ManagedInstancePrivateEndpointProperty struct {
	Id *string `pulumi:"id"`
}





type ManagedInstancePrivateEndpointPropertyInput interface {
	pulumi.Input

	ToManagedInstancePrivateEndpointPropertyOutput() ManagedInstancePrivateEndpointPropertyOutput
	ToManagedInstancePrivateEndpointPropertyOutputWithContext(context.Context) ManagedInstancePrivateEndpointPropertyOutput
}

type ManagedInstancePrivateEndpointPropertyArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ManagedInstancePrivateEndpointPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePrivateEndpointProperty)(nil)).Elem()
}

func (i ManagedInstancePrivateEndpointPropertyArgs) ToManagedInstancePrivateEndpointPropertyOutput() ManagedInstancePrivateEndpointPropertyOutput {
	return i.ToManagedInstancePrivateEndpointPropertyOutputWithContext(context.Background())
}

func (i ManagedInstancePrivateEndpointPropertyArgs) ToManagedInstancePrivateEndpointPropertyOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateEndpointPropertyOutput)
}

func (i ManagedInstancePrivateEndpointPropertyArgs) ToManagedInstancePrivateEndpointPropertyPtrOutput() ManagedInstancePrivateEndpointPropertyPtrOutput {
	return i.ToManagedInstancePrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i ManagedInstancePrivateEndpointPropertyArgs) ToManagedInstancePrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateEndpointPropertyOutput).ToManagedInstancePrivateEndpointPropertyPtrOutputWithContext(ctx)
}









type ManagedInstancePrivateEndpointPropertyPtrInput interface {
	pulumi.Input

	ToManagedInstancePrivateEndpointPropertyPtrOutput() ManagedInstancePrivateEndpointPropertyPtrOutput
	ToManagedInstancePrivateEndpointPropertyPtrOutputWithContext(context.Context) ManagedInstancePrivateEndpointPropertyPtrOutput
}

type managedInstancePrivateEndpointPropertyPtrType ManagedInstancePrivateEndpointPropertyArgs

func ManagedInstancePrivateEndpointPropertyPtr(v *ManagedInstancePrivateEndpointPropertyArgs) ManagedInstancePrivateEndpointPropertyPtrInput {
	return (*managedInstancePrivateEndpointPropertyPtrType)(v)
}

func (*managedInstancePrivateEndpointPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateEndpointProperty)(nil)).Elem()
}

func (i *managedInstancePrivateEndpointPropertyPtrType) ToManagedInstancePrivateEndpointPropertyPtrOutput() ManagedInstancePrivateEndpointPropertyPtrOutput {
	return i.ToManagedInstancePrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i *managedInstancePrivateEndpointPropertyPtrType) ToManagedInstancePrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateEndpointPropertyPtrOutput)
}

type ManagedInstancePrivateEndpointPropertyOutput struct{ *pulumi.OutputState }

func (ManagedInstancePrivateEndpointPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePrivateEndpointProperty)(nil)).Elem()
}

func (o ManagedInstancePrivateEndpointPropertyOutput) ToManagedInstancePrivateEndpointPropertyOutput() ManagedInstancePrivateEndpointPropertyOutput {
	return o
}

func (o ManagedInstancePrivateEndpointPropertyOutput) ToManagedInstancePrivateEndpointPropertyOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyOutput {
	return o
}

func (o ManagedInstancePrivateEndpointPropertyOutput) ToManagedInstancePrivateEndpointPropertyPtrOutput() ManagedInstancePrivateEndpointPropertyPtrOutput {
	return o.ToManagedInstancePrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (o ManagedInstancePrivateEndpointPropertyOutput) ToManagedInstancePrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedInstancePrivateEndpointProperty) *ManagedInstancePrivateEndpointProperty {
		return &v
	}).(ManagedInstancePrivateEndpointPropertyPtrOutput)
}

func (o ManagedInstancePrivateEndpointPropertyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstancePrivateEndpointProperty) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ManagedInstancePrivateEndpointPropertyPtrOutput struct{ *pulumi.OutputState }

func (ManagedInstancePrivateEndpointPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateEndpointProperty)(nil)).Elem()
}

func (o ManagedInstancePrivateEndpointPropertyPtrOutput) ToManagedInstancePrivateEndpointPropertyPtrOutput() ManagedInstancePrivateEndpointPropertyPtrOutput {
	return o
}

func (o ManagedInstancePrivateEndpointPropertyPtrOutput) ToManagedInstancePrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyPtrOutput {
	return o
}

func (o ManagedInstancePrivateEndpointPropertyPtrOutput) Elem() ManagedInstancePrivateEndpointPropertyOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointProperty) ManagedInstancePrivateEndpointProperty {
		if v != nil {
			return *v
		}
		var ret ManagedInstancePrivateEndpointProperty
		return ret
	}).(ManagedInstancePrivateEndpointPropertyOutput)
}

func (o ManagedInstancePrivateEndpointPropertyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointProperty) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ManagedInstancePrivateEndpointPropertyResponse struct {
	Id *string `pulumi:"id"`
}





type ManagedInstancePrivateEndpointPropertyResponseInput interface {
	pulumi.Input

	ToManagedInstancePrivateEndpointPropertyResponseOutput() ManagedInstancePrivateEndpointPropertyResponseOutput
	ToManagedInstancePrivateEndpointPropertyResponseOutputWithContext(context.Context) ManagedInstancePrivateEndpointPropertyResponseOutput
}

type ManagedInstancePrivateEndpointPropertyResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ManagedInstancePrivateEndpointPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i ManagedInstancePrivateEndpointPropertyResponseArgs) ToManagedInstancePrivateEndpointPropertyResponseOutput() ManagedInstancePrivateEndpointPropertyResponseOutput {
	return i.ToManagedInstancePrivateEndpointPropertyResponseOutputWithContext(context.Background())
}

func (i ManagedInstancePrivateEndpointPropertyResponseArgs) ToManagedInstancePrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateEndpointPropertyResponseOutput)
}

func (i ManagedInstancePrivateEndpointPropertyResponseArgs) ToManagedInstancePrivateEndpointPropertyResponsePtrOutput() ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return i.ToManagedInstancePrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i ManagedInstancePrivateEndpointPropertyResponseArgs) ToManagedInstancePrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateEndpointPropertyResponseOutput).ToManagedInstancePrivateEndpointPropertyResponsePtrOutputWithContext(ctx)
}









type ManagedInstancePrivateEndpointPropertyResponsePtrInput interface {
	pulumi.Input

	ToManagedInstancePrivateEndpointPropertyResponsePtrOutput() ManagedInstancePrivateEndpointPropertyResponsePtrOutput
	ToManagedInstancePrivateEndpointPropertyResponsePtrOutputWithContext(context.Context) ManagedInstancePrivateEndpointPropertyResponsePtrOutput
}

type managedInstancePrivateEndpointPropertyResponsePtrType ManagedInstancePrivateEndpointPropertyResponseArgs

func ManagedInstancePrivateEndpointPropertyResponsePtr(v *ManagedInstancePrivateEndpointPropertyResponseArgs) ManagedInstancePrivateEndpointPropertyResponsePtrInput {
	return (*managedInstancePrivateEndpointPropertyResponsePtrType)(v)
}

func (*managedInstancePrivateEndpointPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i *managedInstancePrivateEndpointPropertyResponsePtrType) ToManagedInstancePrivateEndpointPropertyResponsePtrOutput() ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return i.ToManagedInstancePrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *managedInstancePrivateEndpointPropertyResponsePtrType) ToManagedInstancePrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateEndpointPropertyResponsePtrOutput)
}

type ManagedInstancePrivateEndpointPropertyResponseOutput struct{ *pulumi.OutputState }

func (ManagedInstancePrivateEndpointPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o ManagedInstancePrivateEndpointPropertyResponseOutput) ToManagedInstancePrivateEndpointPropertyResponseOutput() ManagedInstancePrivateEndpointPropertyResponseOutput {
	return o
}

func (o ManagedInstancePrivateEndpointPropertyResponseOutput) ToManagedInstancePrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyResponseOutput {
	return o
}

func (o ManagedInstancePrivateEndpointPropertyResponseOutput) ToManagedInstancePrivateEndpointPropertyResponsePtrOutput() ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return o.ToManagedInstancePrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (o ManagedInstancePrivateEndpointPropertyResponseOutput) ToManagedInstancePrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedInstancePrivateEndpointPropertyResponse) *ManagedInstancePrivateEndpointPropertyResponse {
		return &v
	}).(ManagedInstancePrivateEndpointPropertyResponsePtrOutput)
}

func (o ManagedInstancePrivateEndpointPropertyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedInstancePrivateEndpointPropertyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ManagedInstancePrivateEndpointPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedInstancePrivateEndpointPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o ManagedInstancePrivateEndpointPropertyResponsePtrOutput) ToManagedInstancePrivateEndpointPropertyResponsePtrOutput() ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o ManagedInstancePrivateEndpointPropertyResponsePtrOutput) ToManagedInstancePrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o ManagedInstancePrivateEndpointPropertyResponsePtrOutput) Elem() ManagedInstancePrivateEndpointPropertyResponseOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointPropertyResponse) ManagedInstancePrivateEndpointPropertyResponse {
		if v != nil {
			return *v
		}
		var ret ManagedInstancePrivateEndpointPropertyResponse
		return ret
	}).(ManagedInstancePrivateEndpointPropertyResponseOutput)
}

func (o ManagedInstancePrivateEndpointPropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ManagedInstancePrivateLinkServiceConnectionStateProperty struct {
	Description string `pulumi:"description"`
	Status      string `pulumi:"status"`
}





type ManagedInstancePrivateLinkServiceConnectionStatePropertyInput interface {
	pulumi.Input

	ToManagedInstancePrivateLinkServiceConnectionStatePropertyOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput
	ToManagedInstancePrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput
}

type ManagedInstancePrivateLinkServiceConnectionStatePropertyArgs struct {
	Description pulumi.StringInput `pulumi:"description"`
	Status      pulumi.StringInput `pulumi:"status"`
}

func (ManagedInstancePrivateLinkServiceConnectionStatePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i ManagedInstancePrivateLinkServiceConnectionStatePropertyArgs) ToManagedInstancePrivateLinkServiceConnectionStatePropertyOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput {
	return i.ToManagedInstancePrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Background())
}

func (i ManagedInstancePrivateLinkServiceConnectionStatePropertyArgs) ToManagedInstancePrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput)
}

func (i ManagedInstancePrivateLinkServiceConnectionStatePropertyArgs) ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i ManagedInstancePrivateLinkServiceConnectionStatePropertyArgs) ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput).ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx)
}









type ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrInput interface {
	pulumi.Input

	ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput
	ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput
}

type managedInstancePrivateLinkServiceConnectionStatePropertyPtrType ManagedInstancePrivateLinkServiceConnectionStatePropertyArgs

func ManagedInstancePrivateLinkServiceConnectionStatePropertyPtr(v *ManagedInstancePrivateLinkServiceConnectionStatePropertyArgs) ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrInput {
	return (*managedInstancePrivateLinkServiceConnectionStatePropertyPtrType)(v)
}

func (*managedInstancePrivateLinkServiceConnectionStatePropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i *managedInstancePrivateLinkServiceConnectionStatePropertyPtrType) ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i *managedInstancePrivateLinkServiceConnectionStatePropertyPtrType) ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput struct{ *pulumi.OutputState }

func (ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedInstancePrivateLinkServiceConnectionStateProperty) *ManagedInstancePrivateLinkServiceConnectionStateProperty {
		return &v
	}).(ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput)
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedInstancePrivateLinkServiceConnectionStateProperty) string { return v.Description }).(pulumi.StringOutput)
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedInstancePrivateLinkServiceConnectionStateProperty) string { return v.Status }).(pulumi.StringOutput)
}

type ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput struct{ *pulumi.OutputState }

func (ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput) Elem() ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateLinkServiceConnectionStateProperty) ManagedInstancePrivateLinkServiceConnectionStateProperty {
		if v != nil {
			return *v
		}
		var ret ManagedInstancePrivateLinkServiceConnectionStateProperty
		return ret
	}).(ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput)
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}





type ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseInput interface {
	pulumi.Input

	ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput
	ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput
}

type ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseArgs struct {
	ActionsRequired pulumi.StringInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput `pulumi:"description"`
	Status          pulumi.StringInput `pulumi:"status"`
}

func (ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseArgs) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput {
	return i.ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Background())
}

func (i ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseArgs) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (i ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseArgs) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseArgs) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput).ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx)
}









type ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrInput interface {
	pulumi.Input

	ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput
	ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput
}

type managedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrType ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseArgs

func ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtr(v *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseArgs) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrInput {
	return (*managedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrType)(v)
}

func (*managedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i *managedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrType) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i *managedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrType) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

type ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse) *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse {
		return &v
	}).(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse) string {
		return v.ActionsRequired
	}).(pulumi.StringOutput)
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse) string { return v.Status }).(pulumi.StringOutput)
}

type ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Elem() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse {
		if v != nil {
			return *v
		}
		var ret ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse
		return ret
	}).(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PartnerInfo struct {
	Id string `pulumi:"id"`
}





type PartnerInfoInput interface {
	pulumi.Input

	ToPartnerInfoOutput() PartnerInfoOutput
	ToPartnerInfoOutputWithContext(context.Context) PartnerInfoOutput
}

type PartnerInfoArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PartnerInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerInfo)(nil)).Elem()
}

func (i PartnerInfoArgs) ToPartnerInfoOutput() PartnerInfoOutput {
	return i.ToPartnerInfoOutputWithContext(context.Background())
}

func (i PartnerInfoArgs) ToPartnerInfoOutputWithContext(ctx context.Context) PartnerInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerInfoOutput)
}





type PartnerInfoArrayInput interface {
	pulumi.Input

	ToPartnerInfoArrayOutput() PartnerInfoArrayOutput
	ToPartnerInfoArrayOutputWithContext(context.Context) PartnerInfoArrayOutput
}

type PartnerInfoArray []PartnerInfoInput

func (PartnerInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerInfo)(nil)).Elem()
}

func (i PartnerInfoArray) ToPartnerInfoArrayOutput() PartnerInfoArrayOutput {
	return i.ToPartnerInfoArrayOutputWithContext(context.Background())
}

func (i PartnerInfoArray) ToPartnerInfoArrayOutputWithContext(ctx context.Context) PartnerInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerInfoArrayOutput)
}

type PartnerInfoOutput struct{ *pulumi.OutputState }

func (PartnerInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerInfo)(nil)).Elem()
}

func (o PartnerInfoOutput) ToPartnerInfoOutput() PartnerInfoOutput {
	return o
}

func (o PartnerInfoOutput) ToPartnerInfoOutputWithContext(ctx context.Context) PartnerInfoOutput {
	return o
}

func (o PartnerInfoOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PartnerInfo) string { return v.Id }).(pulumi.StringOutput)
}

type PartnerInfoArrayOutput struct{ *pulumi.OutputState }

func (PartnerInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerInfo)(nil)).Elem()
}

func (o PartnerInfoArrayOutput) ToPartnerInfoArrayOutput() PartnerInfoArrayOutput {
	return o
}

func (o PartnerInfoArrayOutput) ToPartnerInfoArrayOutputWithContext(ctx context.Context) PartnerInfoArrayOutput {
	return o
}

func (o PartnerInfoArrayOutput) Index(i pulumi.IntInput) PartnerInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PartnerInfo {
		return vs[0].([]PartnerInfo)[vs[1].(int)]
	}).(PartnerInfoOutput)
}

type PartnerInfoResponse struct {
	Id              string `pulumi:"id"`
	Location        string `pulumi:"location"`
	ReplicationRole string `pulumi:"replicationRole"`
}





type PartnerInfoResponseInput interface {
	pulumi.Input

	ToPartnerInfoResponseOutput() PartnerInfoResponseOutput
	ToPartnerInfoResponseOutputWithContext(context.Context) PartnerInfoResponseOutput
}

type PartnerInfoResponseArgs struct {
	Id              pulumi.StringInput `pulumi:"id"`
	Location        pulumi.StringInput `pulumi:"location"`
	ReplicationRole pulumi.StringInput `pulumi:"replicationRole"`
}

func (PartnerInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerInfoResponse)(nil)).Elem()
}

func (i PartnerInfoResponseArgs) ToPartnerInfoResponseOutput() PartnerInfoResponseOutput {
	return i.ToPartnerInfoResponseOutputWithContext(context.Background())
}

func (i PartnerInfoResponseArgs) ToPartnerInfoResponseOutputWithContext(ctx context.Context) PartnerInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerInfoResponseOutput)
}





type PartnerInfoResponseArrayInput interface {
	pulumi.Input

	ToPartnerInfoResponseArrayOutput() PartnerInfoResponseArrayOutput
	ToPartnerInfoResponseArrayOutputWithContext(context.Context) PartnerInfoResponseArrayOutput
}

type PartnerInfoResponseArray []PartnerInfoResponseInput

func (PartnerInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerInfoResponse)(nil)).Elem()
}

func (i PartnerInfoResponseArray) ToPartnerInfoResponseArrayOutput() PartnerInfoResponseArrayOutput {
	return i.ToPartnerInfoResponseArrayOutputWithContext(context.Background())
}

func (i PartnerInfoResponseArray) ToPartnerInfoResponseArrayOutputWithContext(ctx context.Context) PartnerInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerInfoResponseArrayOutput)
}

type PartnerInfoResponseOutput struct{ *pulumi.OutputState }

func (PartnerInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerInfoResponse)(nil)).Elem()
}

func (o PartnerInfoResponseOutput) ToPartnerInfoResponseOutput() PartnerInfoResponseOutput {
	return o
}

func (o PartnerInfoResponseOutput) ToPartnerInfoResponseOutputWithContext(ctx context.Context) PartnerInfoResponseOutput {
	return o
}

func (o PartnerInfoResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PartnerInfoResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PartnerInfoResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v PartnerInfoResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o PartnerInfoResponseOutput) ReplicationRole() pulumi.StringOutput {
	return o.ApplyT(func(v PartnerInfoResponse) string { return v.ReplicationRole }).(pulumi.StringOutput)
}

type PartnerInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (PartnerInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerInfoResponse)(nil)).Elem()
}

func (o PartnerInfoResponseArrayOutput) ToPartnerInfoResponseArrayOutput() PartnerInfoResponseArrayOutput {
	return o
}

func (o PartnerInfoResponseArrayOutput) ToPartnerInfoResponseArrayOutputWithContext(ctx context.Context) PartnerInfoResponseArrayOutput {
	return o
}

func (o PartnerInfoResponseArrayOutput) Index(i pulumi.IntInput) PartnerInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PartnerInfoResponse {
		return vs[0].([]PartnerInfoResponse)[vs[1].(int)]
	}).(PartnerInfoResponseOutput)
}

type PartnerRegionInfo struct {
	Location *string `pulumi:"location"`
}





type PartnerRegionInfoInput interface {
	pulumi.Input

	ToPartnerRegionInfoOutput() PartnerRegionInfoOutput
	ToPartnerRegionInfoOutputWithContext(context.Context) PartnerRegionInfoOutput
}

type PartnerRegionInfoArgs struct {
	Location pulumi.StringPtrInput `pulumi:"location"`
}

func (PartnerRegionInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerRegionInfo)(nil)).Elem()
}

func (i PartnerRegionInfoArgs) ToPartnerRegionInfoOutput() PartnerRegionInfoOutput {
	return i.ToPartnerRegionInfoOutputWithContext(context.Background())
}

func (i PartnerRegionInfoArgs) ToPartnerRegionInfoOutputWithContext(ctx context.Context) PartnerRegionInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerRegionInfoOutput)
}





type PartnerRegionInfoArrayInput interface {
	pulumi.Input

	ToPartnerRegionInfoArrayOutput() PartnerRegionInfoArrayOutput
	ToPartnerRegionInfoArrayOutputWithContext(context.Context) PartnerRegionInfoArrayOutput
}

type PartnerRegionInfoArray []PartnerRegionInfoInput

func (PartnerRegionInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerRegionInfo)(nil)).Elem()
}

func (i PartnerRegionInfoArray) ToPartnerRegionInfoArrayOutput() PartnerRegionInfoArrayOutput {
	return i.ToPartnerRegionInfoArrayOutputWithContext(context.Background())
}

func (i PartnerRegionInfoArray) ToPartnerRegionInfoArrayOutputWithContext(ctx context.Context) PartnerRegionInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerRegionInfoArrayOutput)
}

type PartnerRegionInfoOutput struct{ *pulumi.OutputState }

func (PartnerRegionInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerRegionInfo)(nil)).Elem()
}

func (o PartnerRegionInfoOutput) ToPartnerRegionInfoOutput() PartnerRegionInfoOutput {
	return o
}

func (o PartnerRegionInfoOutput) ToPartnerRegionInfoOutputWithContext(ctx context.Context) PartnerRegionInfoOutput {
	return o
}

func (o PartnerRegionInfoOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerRegionInfo) *string { return v.Location }).(pulumi.StringPtrOutput)
}

type PartnerRegionInfoArrayOutput struct{ *pulumi.OutputState }

func (PartnerRegionInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerRegionInfo)(nil)).Elem()
}

func (o PartnerRegionInfoArrayOutput) ToPartnerRegionInfoArrayOutput() PartnerRegionInfoArrayOutput {
	return o
}

func (o PartnerRegionInfoArrayOutput) ToPartnerRegionInfoArrayOutputWithContext(ctx context.Context) PartnerRegionInfoArrayOutput {
	return o
}

func (o PartnerRegionInfoArrayOutput) Index(i pulumi.IntInput) PartnerRegionInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PartnerRegionInfo {
		return vs[0].([]PartnerRegionInfo)[vs[1].(int)]
	}).(PartnerRegionInfoOutput)
}

type PartnerRegionInfoResponse struct {
	Location        *string `pulumi:"location"`
	ReplicationRole string  `pulumi:"replicationRole"`
}





type PartnerRegionInfoResponseInput interface {
	pulumi.Input

	ToPartnerRegionInfoResponseOutput() PartnerRegionInfoResponseOutput
	ToPartnerRegionInfoResponseOutputWithContext(context.Context) PartnerRegionInfoResponseOutput
}

type PartnerRegionInfoResponseArgs struct {
	Location        pulumi.StringPtrInput `pulumi:"location"`
	ReplicationRole pulumi.StringInput    `pulumi:"replicationRole"`
}

func (PartnerRegionInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerRegionInfoResponse)(nil)).Elem()
}

func (i PartnerRegionInfoResponseArgs) ToPartnerRegionInfoResponseOutput() PartnerRegionInfoResponseOutput {
	return i.ToPartnerRegionInfoResponseOutputWithContext(context.Background())
}

func (i PartnerRegionInfoResponseArgs) ToPartnerRegionInfoResponseOutputWithContext(ctx context.Context) PartnerRegionInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerRegionInfoResponseOutput)
}





type PartnerRegionInfoResponseArrayInput interface {
	pulumi.Input

	ToPartnerRegionInfoResponseArrayOutput() PartnerRegionInfoResponseArrayOutput
	ToPartnerRegionInfoResponseArrayOutputWithContext(context.Context) PartnerRegionInfoResponseArrayOutput
}

type PartnerRegionInfoResponseArray []PartnerRegionInfoResponseInput

func (PartnerRegionInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerRegionInfoResponse)(nil)).Elem()
}

func (i PartnerRegionInfoResponseArray) ToPartnerRegionInfoResponseArrayOutput() PartnerRegionInfoResponseArrayOutput {
	return i.ToPartnerRegionInfoResponseArrayOutputWithContext(context.Background())
}

func (i PartnerRegionInfoResponseArray) ToPartnerRegionInfoResponseArrayOutputWithContext(ctx context.Context) PartnerRegionInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerRegionInfoResponseArrayOutput)
}

type PartnerRegionInfoResponseOutput struct{ *pulumi.OutputState }

func (PartnerRegionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerRegionInfoResponse)(nil)).Elem()
}

func (o PartnerRegionInfoResponseOutput) ToPartnerRegionInfoResponseOutput() PartnerRegionInfoResponseOutput {
	return o
}

func (o PartnerRegionInfoResponseOutput) ToPartnerRegionInfoResponseOutputWithContext(ctx context.Context) PartnerRegionInfoResponseOutput {
	return o
}

func (o PartnerRegionInfoResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartnerRegionInfoResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PartnerRegionInfoResponseOutput) ReplicationRole() pulumi.StringOutput {
	return o.ApplyT(func(v PartnerRegionInfoResponse) string { return v.ReplicationRole }).(pulumi.StringOutput)
}

type PartnerRegionInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (PartnerRegionInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerRegionInfoResponse)(nil)).Elem()
}

func (o PartnerRegionInfoResponseArrayOutput) ToPartnerRegionInfoResponseArrayOutput() PartnerRegionInfoResponseArrayOutput {
	return o
}

func (o PartnerRegionInfoResponseArrayOutput) ToPartnerRegionInfoResponseArrayOutputWithContext(ctx context.Context) PartnerRegionInfoResponseArrayOutput {
	return o
}

func (o PartnerRegionInfoResponseArrayOutput) Index(i pulumi.IntInput) PartnerRegionInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PartnerRegionInfoResponse {
		return vs[0].([]PartnerRegionInfoResponse)[vs[1].(int)]
	}).(PartnerRegionInfoResponseOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                             `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionPropertiesResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput
	ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponseOutput
}

type PrivateEndpointConnectionPropertiesResponseArgs struct {
	PrivateEndpoint                   PrivateEndpointPropertyResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                        `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput)
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointPropertyResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PrivateEndpointProperty struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointPropertyInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput
	ToPrivateEndpointPropertyOutputWithContext(context.Context) PrivateEndpointPropertyOutput
}

type PrivateEndpointPropertyArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return i.ToPrivateEndpointPropertyOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput)
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput).ToPrivateEndpointPropertyPtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput
	ToPrivateEndpointPropertyPtrOutputWithContext(context.Context) PrivateEndpointPropertyPtrOutput
}

type privateEndpointPropertyPtrType PrivateEndpointPropertyArgs

func PrivateEndpointPropertyPtr(v *PrivateEndpointPropertyArgs) PrivateEndpointPropertyPtrInput {
	return (*privateEndpointPropertyPtrType)(v)
}

func (*privateEndpointPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyPtrOutput)
}

type PrivateEndpointPropertyOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointProperty) *PrivateEndpointProperty {
		return &v
	}).(PrivateEndpointPropertyPtrOutput)
}

func (o PrivateEndpointPropertyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointProperty) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) Elem() PrivateEndpointPropertyOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) PrivateEndpointProperty {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointProperty
		return ret
	}).(PrivateEndpointPropertyOutput)
}

func (o PrivateEndpointPropertyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponse struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointPropertyResponseInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput
	ToPrivateEndpointPropertyResponseOutputWithContext(context.Context) PrivateEndpointPropertyResponseOutput
}

type PrivateEndpointPropertyResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return i.ToPrivateEndpointPropertyResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput)
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput).ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput
	ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Context) PrivateEndpointPropertyResponsePtrOutput
}

type privateEndpointPropertyResponsePtrType PrivateEndpointPropertyResponseArgs

func PrivateEndpointPropertyResponsePtr(v *PrivateEndpointPropertyResponseArgs) PrivateEndpointPropertyResponsePtrInput {
	return (*privateEndpointPropertyResponsePtrType)(v)
}

func (*privateEndpointPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponsePtrOutput)
}

type PrivateEndpointPropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointPropertyResponse) *PrivateEndpointPropertyResponse {
		return &v
	}).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointPropertyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointPropertyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) Elem() PrivateEndpointPropertyResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) PrivateEndpointPropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointPropertyResponse
		return ret
	}).(PrivateEndpointPropertyResponseOutput)
}

func (o PrivateEndpointPropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateProperty struct {
	Description string `pulumi:"description"`
	Status      string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput
	ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyOutput
}

type PrivateLinkServiceConnectionStatePropertyArgs struct {
	Description pulumi.StringInput `pulumi:"description"`
	Status      pulumi.StringInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput).ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput
	ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput
}

type privateLinkServiceConnectionStatePropertyPtrType PrivateLinkServiceConnectionStatePropertyArgs

func PrivateLinkServiceConnectionStatePropertyPtr(v *PrivateLinkServiceConnectionStatePropertyArgs) PrivateLinkServiceConnectionStatePropertyPtrInput {
	return (*privateLinkServiceConnectionStatePropertyPtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateProperty) *PrivateLinkServiceConnectionStateProperty {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) PrivateLinkServiceConnectionStateProperty {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateProperty
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput
	ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput
}

type PrivateLinkServiceConnectionStatePropertyResponseArgs struct {
	ActionsRequired pulumi.StringInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput `pulumi:"description"`
	Status          pulumi.StringInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput).ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
}

type privateLinkServiceConnectionStatePropertyResponsePtrType PrivateLinkServiceConnectionStatePropertyResponseArgs

func PrivateLinkServiceConnectionStatePropertyResponsePtr(v *PrivateLinkServiceConnectionStatePropertyResponseArgs) PrivateLinkServiceConnectionStatePropertyResponsePtrInput {
	return (*privateLinkServiceConnectionStatePropertyResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStatePropertyResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) PrivateLinkServiceConnectionStatePropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatePropertyResponse
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type RecommendedActionErrorInfoResponse struct {
	ErrorCode   string `pulumi:"errorCode"`
	IsRetryable string `pulumi:"isRetryable"`
}





type RecommendedActionErrorInfoResponseInput interface {
	pulumi.Input

	ToRecommendedActionErrorInfoResponseOutput() RecommendedActionErrorInfoResponseOutput
	ToRecommendedActionErrorInfoResponseOutputWithContext(context.Context) RecommendedActionErrorInfoResponseOutput
}

type RecommendedActionErrorInfoResponseArgs struct {
	ErrorCode   pulumi.StringInput `pulumi:"errorCode"`
	IsRetryable pulumi.StringInput `pulumi:"isRetryable"`
}

func (RecommendedActionErrorInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionErrorInfoResponse)(nil)).Elem()
}

func (i RecommendedActionErrorInfoResponseArgs) ToRecommendedActionErrorInfoResponseOutput() RecommendedActionErrorInfoResponseOutput {
	return i.ToRecommendedActionErrorInfoResponseOutputWithContext(context.Background())
}

func (i RecommendedActionErrorInfoResponseArgs) ToRecommendedActionErrorInfoResponseOutputWithContext(ctx context.Context) RecommendedActionErrorInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionErrorInfoResponseOutput)
}

type RecommendedActionErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionErrorInfoResponse)(nil)).Elem()
}

func (o RecommendedActionErrorInfoResponseOutput) ToRecommendedActionErrorInfoResponseOutput() RecommendedActionErrorInfoResponseOutput {
	return o
}

func (o RecommendedActionErrorInfoResponseOutput) ToRecommendedActionErrorInfoResponseOutputWithContext(ctx context.Context) RecommendedActionErrorInfoResponseOutput {
	return o
}

func (o RecommendedActionErrorInfoResponseOutput) ErrorCode() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionErrorInfoResponse) string { return v.ErrorCode }).(pulumi.StringOutput)
}

func (o RecommendedActionErrorInfoResponseOutput) IsRetryable() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionErrorInfoResponse) string { return v.IsRetryable }).(pulumi.StringOutput)
}

type RecommendedActionImpactRecordResponse struct {
	AbsoluteValue       float64 `pulumi:"absoluteValue"`
	ChangeValueAbsolute float64 `pulumi:"changeValueAbsolute"`
	ChangeValueRelative float64 `pulumi:"changeValueRelative"`
	DimensionName       string  `pulumi:"dimensionName"`
	Unit                string  `pulumi:"unit"`
}





type RecommendedActionImpactRecordResponseInput interface {
	pulumi.Input

	ToRecommendedActionImpactRecordResponseOutput() RecommendedActionImpactRecordResponseOutput
	ToRecommendedActionImpactRecordResponseOutputWithContext(context.Context) RecommendedActionImpactRecordResponseOutput
}

type RecommendedActionImpactRecordResponseArgs struct {
	AbsoluteValue       pulumi.Float64Input `pulumi:"absoluteValue"`
	ChangeValueAbsolute pulumi.Float64Input `pulumi:"changeValueAbsolute"`
	ChangeValueRelative pulumi.Float64Input `pulumi:"changeValueRelative"`
	DimensionName       pulumi.StringInput  `pulumi:"dimensionName"`
	Unit                pulumi.StringInput  `pulumi:"unit"`
}

func (RecommendedActionImpactRecordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionImpactRecordResponse)(nil)).Elem()
}

func (i RecommendedActionImpactRecordResponseArgs) ToRecommendedActionImpactRecordResponseOutput() RecommendedActionImpactRecordResponseOutput {
	return i.ToRecommendedActionImpactRecordResponseOutputWithContext(context.Background())
}

func (i RecommendedActionImpactRecordResponseArgs) ToRecommendedActionImpactRecordResponseOutputWithContext(ctx context.Context) RecommendedActionImpactRecordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionImpactRecordResponseOutput)
}





type RecommendedActionImpactRecordResponseArrayInput interface {
	pulumi.Input

	ToRecommendedActionImpactRecordResponseArrayOutput() RecommendedActionImpactRecordResponseArrayOutput
	ToRecommendedActionImpactRecordResponseArrayOutputWithContext(context.Context) RecommendedActionImpactRecordResponseArrayOutput
}

type RecommendedActionImpactRecordResponseArray []RecommendedActionImpactRecordResponseInput

func (RecommendedActionImpactRecordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionImpactRecordResponse)(nil)).Elem()
}

func (i RecommendedActionImpactRecordResponseArray) ToRecommendedActionImpactRecordResponseArrayOutput() RecommendedActionImpactRecordResponseArrayOutput {
	return i.ToRecommendedActionImpactRecordResponseArrayOutputWithContext(context.Background())
}

func (i RecommendedActionImpactRecordResponseArray) ToRecommendedActionImpactRecordResponseArrayOutputWithContext(ctx context.Context) RecommendedActionImpactRecordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionImpactRecordResponseArrayOutput)
}

type RecommendedActionImpactRecordResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionImpactRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionImpactRecordResponse)(nil)).Elem()
}

func (o RecommendedActionImpactRecordResponseOutput) ToRecommendedActionImpactRecordResponseOutput() RecommendedActionImpactRecordResponseOutput {
	return o
}

func (o RecommendedActionImpactRecordResponseOutput) ToRecommendedActionImpactRecordResponseOutputWithContext(ctx context.Context) RecommendedActionImpactRecordResponseOutput {
	return o
}

func (o RecommendedActionImpactRecordResponseOutput) AbsoluteValue() pulumi.Float64Output {
	return o.ApplyT(func(v RecommendedActionImpactRecordResponse) float64 { return v.AbsoluteValue }).(pulumi.Float64Output)
}

func (o RecommendedActionImpactRecordResponseOutput) ChangeValueAbsolute() pulumi.Float64Output {
	return o.ApplyT(func(v RecommendedActionImpactRecordResponse) float64 { return v.ChangeValueAbsolute }).(pulumi.Float64Output)
}

func (o RecommendedActionImpactRecordResponseOutput) ChangeValueRelative() pulumi.Float64Output {
	return o.ApplyT(func(v RecommendedActionImpactRecordResponse) float64 { return v.ChangeValueRelative }).(pulumi.Float64Output)
}

func (o RecommendedActionImpactRecordResponseOutput) DimensionName() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionImpactRecordResponse) string { return v.DimensionName }).(pulumi.StringOutput)
}

func (o RecommendedActionImpactRecordResponseOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionImpactRecordResponse) string { return v.Unit }).(pulumi.StringOutput)
}

type RecommendedActionImpactRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (RecommendedActionImpactRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionImpactRecordResponse)(nil)).Elem()
}

func (o RecommendedActionImpactRecordResponseArrayOutput) ToRecommendedActionImpactRecordResponseArrayOutput() RecommendedActionImpactRecordResponseArrayOutput {
	return o
}

func (o RecommendedActionImpactRecordResponseArrayOutput) ToRecommendedActionImpactRecordResponseArrayOutputWithContext(ctx context.Context) RecommendedActionImpactRecordResponseArrayOutput {
	return o
}

func (o RecommendedActionImpactRecordResponseArrayOutput) Index(i pulumi.IntInput) RecommendedActionImpactRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecommendedActionImpactRecordResponse {
		return vs[0].([]RecommendedActionImpactRecordResponse)[vs[1].(int)]
	}).(RecommendedActionImpactRecordResponseOutput)
}

type RecommendedActionImplementationInfoResponse struct {
	Method string `pulumi:"method"`
	Script string `pulumi:"script"`
}





type RecommendedActionImplementationInfoResponseInput interface {
	pulumi.Input

	ToRecommendedActionImplementationInfoResponseOutput() RecommendedActionImplementationInfoResponseOutput
	ToRecommendedActionImplementationInfoResponseOutputWithContext(context.Context) RecommendedActionImplementationInfoResponseOutput
}

type RecommendedActionImplementationInfoResponseArgs struct {
	Method pulumi.StringInput `pulumi:"method"`
	Script pulumi.StringInput `pulumi:"script"`
}

func (RecommendedActionImplementationInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionImplementationInfoResponse)(nil)).Elem()
}

func (i RecommendedActionImplementationInfoResponseArgs) ToRecommendedActionImplementationInfoResponseOutput() RecommendedActionImplementationInfoResponseOutput {
	return i.ToRecommendedActionImplementationInfoResponseOutputWithContext(context.Background())
}

func (i RecommendedActionImplementationInfoResponseArgs) ToRecommendedActionImplementationInfoResponseOutputWithContext(ctx context.Context) RecommendedActionImplementationInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionImplementationInfoResponseOutput)
}

type RecommendedActionImplementationInfoResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionImplementationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionImplementationInfoResponse)(nil)).Elem()
}

func (o RecommendedActionImplementationInfoResponseOutput) ToRecommendedActionImplementationInfoResponseOutput() RecommendedActionImplementationInfoResponseOutput {
	return o
}

func (o RecommendedActionImplementationInfoResponseOutput) ToRecommendedActionImplementationInfoResponseOutputWithContext(ctx context.Context) RecommendedActionImplementationInfoResponseOutput {
	return o
}

func (o RecommendedActionImplementationInfoResponseOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionImplementationInfoResponse) string { return v.Method }).(pulumi.StringOutput)
}

func (o RecommendedActionImplementationInfoResponseOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionImplementationInfoResponse) string { return v.Script }).(pulumi.StringOutput)
}

type RecommendedActionMetricInfoResponse struct {
	MetricName string  `pulumi:"metricName"`
	StartTime  string  `pulumi:"startTime"`
	TimeGrain  string  `pulumi:"timeGrain"`
	Unit       string  `pulumi:"unit"`
	Value      float64 `pulumi:"value"`
}





type RecommendedActionMetricInfoResponseInput interface {
	pulumi.Input

	ToRecommendedActionMetricInfoResponseOutput() RecommendedActionMetricInfoResponseOutput
	ToRecommendedActionMetricInfoResponseOutputWithContext(context.Context) RecommendedActionMetricInfoResponseOutput
}

type RecommendedActionMetricInfoResponseArgs struct {
	MetricName pulumi.StringInput  `pulumi:"metricName"`
	StartTime  pulumi.StringInput  `pulumi:"startTime"`
	TimeGrain  pulumi.StringInput  `pulumi:"timeGrain"`
	Unit       pulumi.StringInput  `pulumi:"unit"`
	Value      pulumi.Float64Input `pulumi:"value"`
}

func (RecommendedActionMetricInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionMetricInfoResponse)(nil)).Elem()
}

func (i RecommendedActionMetricInfoResponseArgs) ToRecommendedActionMetricInfoResponseOutput() RecommendedActionMetricInfoResponseOutput {
	return i.ToRecommendedActionMetricInfoResponseOutputWithContext(context.Background())
}

func (i RecommendedActionMetricInfoResponseArgs) ToRecommendedActionMetricInfoResponseOutputWithContext(ctx context.Context) RecommendedActionMetricInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionMetricInfoResponseOutput)
}





type RecommendedActionMetricInfoResponseArrayInput interface {
	pulumi.Input

	ToRecommendedActionMetricInfoResponseArrayOutput() RecommendedActionMetricInfoResponseArrayOutput
	ToRecommendedActionMetricInfoResponseArrayOutputWithContext(context.Context) RecommendedActionMetricInfoResponseArrayOutput
}

type RecommendedActionMetricInfoResponseArray []RecommendedActionMetricInfoResponseInput

func (RecommendedActionMetricInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionMetricInfoResponse)(nil)).Elem()
}

func (i RecommendedActionMetricInfoResponseArray) ToRecommendedActionMetricInfoResponseArrayOutput() RecommendedActionMetricInfoResponseArrayOutput {
	return i.ToRecommendedActionMetricInfoResponseArrayOutputWithContext(context.Background())
}

func (i RecommendedActionMetricInfoResponseArray) ToRecommendedActionMetricInfoResponseArrayOutputWithContext(ctx context.Context) RecommendedActionMetricInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionMetricInfoResponseArrayOutput)
}

type RecommendedActionMetricInfoResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionMetricInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionMetricInfoResponse)(nil)).Elem()
}

func (o RecommendedActionMetricInfoResponseOutput) ToRecommendedActionMetricInfoResponseOutput() RecommendedActionMetricInfoResponseOutput {
	return o
}

func (o RecommendedActionMetricInfoResponseOutput) ToRecommendedActionMetricInfoResponseOutputWithContext(ctx context.Context) RecommendedActionMetricInfoResponseOutput {
	return o
}

func (o RecommendedActionMetricInfoResponseOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionMetricInfoResponse) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o RecommendedActionMetricInfoResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionMetricInfoResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o RecommendedActionMetricInfoResponseOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionMetricInfoResponse) string { return v.TimeGrain }).(pulumi.StringOutput)
}

func (o RecommendedActionMetricInfoResponseOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionMetricInfoResponse) string { return v.Unit }).(pulumi.StringOutput)
}

func (o RecommendedActionMetricInfoResponseOutput) Value() pulumi.Float64Output {
	return o.ApplyT(func(v RecommendedActionMetricInfoResponse) float64 { return v.Value }).(pulumi.Float64Output)
}

type RecommendedActionMetricInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (RecommendedActionMetricInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionMetricInfoResponse)(nil)).Elem()
}

func (o RecommendedActionMetricInfoResponseArrayOutput) ToRecommendedActionMetricInfoResponseArrayOutput() RecommendedActionMetricInfoResponseArrayOutput {
	return o
}

func (o RecommendedActionMetricInfoResponseArrayOutput) ToRecommendedActionMetricInfoResponseArrayOutputWithContext(ctx context.Context) RecommendedActionMetricInfoResponseArrayOutput {
	return o
}

func (o RecommendedActionMetricInfoResponseArrayOutput) Index(i pulumi.IntInput) RecommendedActionMetricInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecommendedActionMetricInfoResponse {
		return vs[0].([]RecommendedActionMetricInfoResponse)[vs[1].(int)]
	}).(RecommendedActionMetricInfoResponseOutput)
}

type RecommendedActionResponse struct {
	Details                    map[string]interface{}                      `pulumi:"details"`
	ErrorDetails               RecommendedActionErrorInfoResponse          `pulumi:"errorDetails"`
	EstimatedImpact            []RecommendedActionImpactRecordResponse     `pulumi:"estimatedImpact"`
	ExecuteActionDuration      string                                      `pulumi:"executeActionDuration"`
	ExecuteActionInitiatedBy   string                                      `pulumi:"executeActionInitiatedBy"`
	ExecuteActionInitiatedTime string                                      `pulumi:"executeActionInitiatedTime"`
	ExecuteActionStartTime     string                                      `pulumi:"executeActionStartTime"`
	Id                         string                                      `pulumi:"id"`
	ImplementationDetails      RecommendedActionImplementationInfoResponse `pulumi:"implementationDetails"`
	IsArchivedAction           bool                                        `pulumi:"isArchivedAction"`
	IsExecutableAction         bool                                        `pulumi:"isExecutableAction"`
	IsRevertableAction         bool                                        `pulumi:"isRevertableAction"`
	Kind                       string                                      `pulumi:"kind"`
	LastRefresh                string                                      `pulumi:"lastRefresh"`
	LinkedObjects              []string                                    `pulumi:"linkedObjects"`
	Location                   string                                      `pulumi:"location"`
	Name                       string                                      `pulumi:"name"`
	ObservedImpact             []RecommendedActionImpactRecordResponse     `pulumi:"observedImpact"`
	RecommendationReason       string                                      `pulumi:"recommendationReason"`
	RevertActionDuration       string                                      `pulumi:"revertActionDuration"`
	RevertActionInitiatedBy    string                                      `pulumi:"revertActionInitiatedBy"`
	RevertActionInitiatedTime  string                                      `pulumi:"revertActionInitiatedTime"`
	RevertActionStartTime      string                                      `pulumi:"revertActionStartTime"`
	Score                      int                                         `pulumi:"score"`
	State                      RecommendedActionStateInfoResponse          `pulumi:"state"`
	TimeSeries                 []RecommendedActionMetricInfoResponse       `pulumi:"timeSeries"`
	Type                       string                                      `pulumi:"type"`
	ValidSince                 string                                      `pulumi:"validSince"`
}





type RecommendedActionResponseInput interface {
	pulumi.Input

	ToRecommendedActionResponseOutput() RecommendedActionResponseOutput
	ToRecommendedActionResponseOutputWithContext(context.Context) RecommendedActionResponseOutput
}

type RecommendedActionResponseArgs struct {
	Details                    pulumi.MapInput                                  `pulumi:"details"`
	ErrorDetails               RecommendedActionErrorInfoResponseInput          `pulumi:"errorDetails"`
	EstimatedImpact            RecommendedActionImpactRecordResponseArrayInput  `pulumi:"estimatedImpact"`
	ExecuteActionDuration      pulumi.StringInput                               `pulumi:"executeActionDuration"`
	ExecuteActionInitiatedBy   pulumi.StringInput                               `pulumi:"executeActionInitiatedBy"`
	ExecuteActionInitiatedTime pulumi.StringInput                               `pulumi:"executeActionInitiatedTime"`
	ExecuteActionStartTime     pulumi.StringInput                               `pulumi:"executeActionStartTime"`
	Id                         pulumi.StringInput                               `pulumi:"id"`
	ImplementationDetails      RecommendedActionImplementationInfoResponseInput `pulumi:"implementationDetails"`
	IsArchivedAction           pulumi.BoolInput                                 `pulumi:"isArchivedAction"`
	IsExecutableAction         pulumi.BoolInput                                 `pulumi:"isExecutableAction"`
	IsRevertableAction         pulumi.BoolInput                                 `pulumi:"isRevertableAction"`
	Kind                       pulumi.StringInput                               `pulumi:"kind"`
	LastRefresh                pulumi.StringInput                               `pulumi:"lastRefresh"`
	LinkedObjects              pulumi.StringArrayInput                          `pulumi:"linkedObjects"`
	Location                   pulumi.StringInput                               `pulumi:"location"`
	Name                       pulumi.StringInput                               `pulumi:"name"`
	ObservedImpact             RecommendedActionImpactRecordResponseArrayInput  `pulumi:"observedImpact"`
	RecommendationReason       pulumi.StringInput                               `pulumi:"recommendationReason"`
	RevertActionDuration       pulumi.StringInput                               `pulumi:"revertActionDuration"`
	RevertActionInitiatedBy    pulumi.StringInput                               `pulumi:"revertActionInitiatedBy"`
	RevertActionInitiatedTime  pulumi.StringInput                               `pulumi:"revertActionInitiatedTime"`
	RevertActionStartTime      pulumi.StringInput                               `pulumi:"revertActionStartTime"`
	Score                      pulumi.IntInput                                  `pulumi:"score"`
	State                      RecommendedActionStateInfoResponseInput          `pulumi:"state"`
	TimeSeries                 RecommendedActionMetricInfoResponseArrayInput    `pulumi:"timeSeries"`
	Type                       pulumi.StringInput                               `pulumi:"type"`
	ValidSince                 pulumi.StringInput                               `pulumi:"validSince"`
}

func (RecommendedActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionResponse)(nil)).Elem()
}

func (i RecommendedActionResponseArgs) ToRecommendedActionResponseOutput() RecommendedActionResponseOutput {
	return i.ToRecommendedActionResponseOutputWithContext(context.Background())
}

func (i RecommendedActionResponseArgs) ToRecommendedActionResponseOutputWithContext(ctx context.Context) RecommendedActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionResponseOutput)
}





type RecommendedActionResponseArrayInput interface {
	pulumi.Input

	ToRecommendedActionResponseArrayOutput() RecommendedActionResponseArrayOutput
	ToRecommendedActionResponseArrayOutputWithContext(context.Context) RecommendedActionResponseArrayOutput
}

type RecommendedActionResponseArray []RecommendedActionResponseInput

func (RecommendedActionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionResponse)(nil)).Elem()
}

func (i RecommendedActionResponseArray) ToRecommendedActionResponseArrayOutput() RecommendedActionResponseArrayOutput {
	return i.ToRecommendedActionResponseArrayOutputWithContext(context.Background())
}

func (i RecommendedActionResponseArray) ToRecommendedActionResponseArrayOutputWithContext(ctx context.Context) RecommendedActionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionResponseArrayOutput)
}

type RecommendedActionResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionResponse)(nil)).Elem()
}

func (o RecommendedActionResponseOutput) ToRecommendedActionResponseOutput() RecommendedActionResponseOutput {
	return o
}

func (o RecommendedActionResponseOutput) ToRecommendedActionResponseOutputWithContext(ctx context.Context) RecommendedActionResponseOutput {
	return o
}

func (o RecommendedActionResponseOutput) Details() pulumi.MapOutput {
	return o.ApplyT(func(v RecommendedActionResponse) map[string]interface{} { return v.Details }).(pulumi.MapOutput)
}

func (o RecommendedActionResponseOutput) ErrorDetails() RecommendedActionErrorInfoResponseOutput {
	return o.ApplyT(func(v RecommendedActionResponse) RecommendedActionErrorInfoResponse { return v.ErrorDetails }).(RecommendedActionErrorInfoResponseOutput)
}

func (o RecommendedActionResponseOutput) EstimatedImpact() RecommendedActionImpactRecordResponseArrayOutput {
	return o.ApplyT(func(v RecommendedActionResponse) []RecommendedActionImpactRecordResponse { return v.EstimatedImpact }).(RecommendedActionImpactRecordResponseArrayOutput)
}

func (o RecommendedActionResponseOutput) ExecuteActionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.ExecuteActionDuration }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ExecuteActionInitiatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.ExecuteActionInitiatedBy }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ExecuteActionInitiatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.ExecuteActionInitiatedTime }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ExecuteActionStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.ExecuteActionStartTime }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ImplementationDetails() RecommendedActionImplementationInfoResponseOutput {
	return o.ApplyT(func(v RecommendedActionResponse) RecommendedActionImplementationInfoResponse {
		return v.ImplementationDetails
	}).(RecommendedActionImplementationInfoResponseOutput)
}

func (o RecommendedActionResponseOutput) IsArchivedAction() pulumi.BoolOutput {
	return o.ApplyT(func(v RecommendedActionResponse) bool { return v.IsArchivedAction }).(pulumi.BoolOutput)
}

func (o RecommendedActionResponseOutput) IsExecutableAction() pulumi.BoolOutput {
	return o.ApplyT(func(v RecommendedActionResponse) bool { return v.IsExecutableAction }).(pulumi.BoolOutput)
}

func (o RecommendedActionResponseOutput) IsRevertableAction() pulumi.BoolOutput {
	return o.ApplyT(func(v RecommendedActionResponse) bool { return v.IsRevertableAction }).(pulumi.BoolOutput)
}

func (o RecommendedActionResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) LastRefresh() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.LastRefresh }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) LinkedObjects() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecommendedActionResponse) []string { return v.LinkedObjects }).(pulumi.StringArrayOutput)
}

func (o RecommendedActionResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ObservedImpact() RecommendedActionImpactRecordResponseArrayOutput {
	return o.ApplyT(func(v RecommendedActionResponse) []RecommendedActionImpactRecordResponse { return v.ObservedImpact }).(RecommendedActionImpactRecordResponseArrayOutput)
}

func (o RecommendedActionResponseOutput) RecommendationReason() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.RecommendationReason }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) RevertActionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.RevertActionDuration }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) RevertActionInitiatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.RevertActionInitiatedBy }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) RevertActionInitiatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.RevertActionInitiatedTime }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) RevertActionStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.RevertActionStartTime }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) Score() pulumi.IntOutput {
	return o.ApplyT(func(v RecommendedActionResponse) int { return v.Score }).(pulumi.IntOutput)
}

func (o RecommendedActionResponseOutput) State() RecommendedActionStateInfoResponseOutput {
	return o.ApplyT(func(v RecommendedActionResponse) RecommendedActionStateInfoResponse { return v.State }).(RecommendedActionStateInfoResponseOutput)
}

func (o RecommendedActionResponseOutput) TimeSeries() RecommendedActionMetricInfoResponseArrayOutput {
	return o.ApplyT(func(v RecommendedActionResponse) []RecommendedActionMetricInfoResponse { return v.TimeSeries }).(RecommendedActionMetricInfoResponseArrayOutput)
}

func (o RecommendedActionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ValidSince() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.ValidSince }).(pulumi.StringOutput)
}

type RecommendedActionResponseArrayOutput struct{ *pulumi.OutputState }

func (RecommendedActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionResponse)(nil)).Elem()
}

func (o RecommendedActionResponseArrayOutput) ToRecommendedActionResponseArrayOutput() RecommendedActionResponseArrayOutput {
	return o
}

func (o RecommendedActionResponseArrayOutput) ToRecommendedActionResponseArrayOutputWithContext(ctx context.Context) RecommendedActionResponseArrayOutput {
	return o
}

func (o RecommendedActionResponseArrayOutput) Index(i pulumi.IntInput) RecommendedActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecommendedActionResponse {
		return vs[0].([]RecommendedActionResponse)[vs[1].(int)]
	}).(RecommendedActionResponseOutput)
}

type RecommendedActionStateInfoResponse struct {
	ActionInitiatedBy string `pulumi:"actionInitiatedBy"`
	CurrentValue      string `pulumi:"currentValue"`
	LastModified      string `pulumi:"lastModified"`
}





type RecommendedActionStateInfoResponseInput interface {
	pulumi.Input

	ToRecommendedActionStateInfoResponseOutput() RecommendedActionStateInfoResponseOutput
	ToRecommendedActionStateInfoResponseOutputWithContext(context.Context) RecommendedActionStateInfoResponseOutput
}

type RecommendedActionStateInfoResponseArgs struct {
	ActionInitiatedBy pulumi.StringInput `pulumi:"actionInitiatedBy"`
	CurrentValue      pulumi.StringInput `pulumi:"currentValue"`
	LastModified      pulumi.StringInput `pulumi:"lastModified"`
}

func (RecommendedActionStateInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionStateInfoResponse)(nil)).Elem()
}

func (i RecommendedActionStateInfoResponseArgs) ToRecommendedActionStateInfoResponseOutput() RecommendedActionStateInfoResponseOutput {
	return i.ToRecommendedActionStateInfoResponseOutputWithContext(context.Background())
}

func (i RecommendedActionStateInfoResponseArgs) ToRecommendedActionStateInfoResponseOutputWithContext(ctx context.Context) RecommendedActionStateInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionStateInfoResponseOutput)
}

type RecommendedActionStateInfoResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionStateInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionStateInfoResponse)(nil)).Elem()
}

func (o RecommendedActionStateInfoResponseOutput) ToRecommendedActionStateInfoResponseOutput() RecommendedActionStateInfoResponseOutput {
	return o
}

func (o RecommendedActionStateInfoResponseOutput) ToRecommendedActionStateInfoResponseOutputWithContext(ctx context.Context) RecommendedActionStateInfoResponseOutput {
	return o
}

func (o RecommendedActionStateInfoResponseOutput) ActionInitiatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionStateInfoResponse) string { return v.ActionInitiatedBy }).(pulumi.StringOutput)
}

func (o RecommendedActionStateInfoResponseOutput) CurrentValue() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionStateInfoResponse) string { return v.CurrentValue }).(pulumi.StringOutput)
}

func (o RecommendedActionStateInfoResponseOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionStateInfoResponse) string { return v.LastModified }).(pulumi.StringOutput)
}

type ResourceIdentity struct {
	Type                   *string                `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	Type                   pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput       `pulumi:"userAssignedIdentities"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}









type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

func (o ResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ResourceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

func (o ResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ResourceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ResourceIdentityResponse struct {
	PrincipalId            string                          `pulumi:"principalId"`
	TenantId               string                          `pulumi:"tenantId"`
	Type                   *string                         `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityResponse `pulumi:"userAssignedIdentities"`
}





type ResourceIdentityResponseInput interface {
	pulumi.Input

	ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput
	ToResourceIdentityResponseOutputWithContext(context.Context) ResourceIdentityResponseOutput
}

type ResourceIdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput           `pulumi:"principalId"`
	TenantId               pulumi.StringInput           `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput        `pulumi:"type"`
	UserAssignedIdentities UserIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (ResourceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return i.ToResourceIdentityResponseOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput)
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput).ToResourceIdentityResponsePtrOutputWithContext(ctx)
}









type ResourceIdentityResponsePtrInput interface {
	pulumi.Input

	ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput
	ToResourceIdentityResponsePtrOutputWithContext(context.Context) ResourceIdentityResponsePtrOutput
}

type resourceIdentityResponsePtrType ResourceIdentityResponseArgs

func ResourceIdentityResponsePtr(v *ResourceIdentityResponseArgs) ResourceIdentityResponsePtrInput {
	return (*resourceIdentityResponsePtrType)(v)
}

func (*resourceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponsePtrOutput)
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityResponse) *ResourceIdentityResponse {
		return &v
	}).(ResourceIdentityResponsePtrOutput)
}

func (o ResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponseOutput) UserAssignedIdentities() UserIdentityResponseMapOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) map[string]UserIdentityResponse { return v.UserAssignedIdentities }).(UserIdentityResponseMapOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

func (o ResourceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) UserAssignedIdentities() UserIdentityResponseMapOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) map[string]UserIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityResponseMapOutput)
}

type ServerExternalAdministrator struct {
	AdministratorType         *string `pulumi:"administratorType"`
	AzureADOnlyAuthentication *bool   `pulumi:"azureADOnlyAuthentication"`
	Login                     *string `pulumi:"login"`
	PrincipalType             *string `pulumi:"principalType"`
	Sid                       *string `pulumi:"sid"`
	TenantId                  *string `pulumi:"tenantId"`
}





type ServerExternalAdministratorInput interface {
	pulumi.Input

	ToServerExternalAdministratorOutput() ServerExternalAdministratorOutput
	ToServerExternalAdministratorOutputWithContext(context.Context) ServerExternalAdministratorOutput
}

type ServerExternalAdministratorArgs struct {
	AdministratorType         pulumi.StringPtrInput `pulumi:"administratorType"`
	AzureADOnlyAuthentication pulumi.BoolPtrInput   `pulumi:"azureADOnlyAuthentication"`
	Login                     pulumi.StringPtrInput `pulumi:"login"`
	PrincipalType             pulumi.StringPtrInput `pulumi:"principalType"`
	Sid                       pulumi.StringPtrInput `pulumi:"sid"`
	TenantId                  pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (ServerExternalAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerExternalAdministrator)(nil)).Elem()
}

func (i ServerExternalAdministratorArgs) ToServerExternalAdministratorOutput() ServerExternalAdministratorOutput {
	return i.ToServerExternalAdministratorOutputWithContext(context.Background())
}

func (i ServerExternalAdministratorArgs) ToServerExternalAdministratorOutputWithContext(ctx context.Context) ServerExternalAdministratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerExternalAdministratorOutput)
}

func (i ServerExternalAdministratorArgs) ToServerExternalAdministratorPtrOutput() ServerExternalAdministratorPtrOutput {
	return i.ToServerExternalAdministratorPtrOutputWithContext(context.Background())
}

func (i ServerExternalAdministratorArgs) ToServerExternalAdministratorPtrOutputWithContext(ctx context.Context) ServerExternalAdministratorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerExternalAdministratorOutput).ToServerExternalAdministratorPtrOutputWithContext(ctx)
}









type ServerExternalAdministratorPtrInput interface {
	pulumi.Input

	ToServerExternalAdministratorPtrOutput() ServerExternalAdministratorPtrOutput
	ToServerExternalAdministratorPtrOutputWithContext(context.Context) ServerExternalAdministratorPtrOutput
}

type serverExternalAdministratorPtrType ServerExternalAdministratorArgs

func ServerExternalAdministratorPtr(v *ServerExternalAdministratorArgs) ServerExternalAdministratorPtrInput {
	return (*serverExternalAdministratorPtrType)(v)
}

func (*serverExternalAdministratorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerExternalAdministrator)(nil)).Elem()
}

func (i *serverExternalAdministratorPtrType) ToServerExternalAdministratorPtrOutput() ServerExternalAdministratorPtrOutput {
	return i.ToServerExternalAdministratorPtrOutputWithContext(context.Background())
}

func (i *serverExternalAdministratorPtrType) ToServerExternalAdministratorPtrOutputWithContext(ctx context.Context) ServerExternalAdministratorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerExternalAdministratorPtrOutput)
}

type ServerExternalAdministratorOutput struct{ *pulumi.OutputState }

func (ServerExternalAdministratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerExternalAdministrator)(nil)).Elem()
}

func (o ServerExternalAdministratorOutput) ToServerExternalAdministratorOutput() ServerExternalAdministratorOutput {
	return o
}

func (o ServerExternalAdministratorOutput) ToServerExternalAdministratorOutputWithContext(ctx context.Context) ServerExternalAdministratorOutput {
	return o
}

func (o ServerExternalAdministratorOutput) ToServerExternalAdministratorPtrOutput() ServerExternalAdministratorPtrOutput {
	return o.ToServerExternalAdministratorPtrOutputWithContext(context.Background())
}

func (o ServerExternalAdministratorOutput) ToServerExternalAdministratorPtrOutputWithContext(ctx context.Context) ServerExternalAdministratorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerExternalAdministrator) *ServerExternalAdministrator {
		return &v
	}).(ServerExternalAdministratorPtrOutput)
}

func (o ServerExternalAdministratorOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministrator) *string { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorOutput) AzureADOnlyAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministrator) *bool { return v.AzureADOnlyAuthentication }).(pulumi.BoolPtrOutput)
}

func (o ServerExternalAdministratorOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministrator) *string { return v.Login }).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministrator) *string { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministrator) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministrator) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ServerExternalAdministratorPtrOutput struct{ *pulumi.OutputState }

func (ServerExternalAdministratorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerExternalAdministrator)(nil)).Elem()
}

func (o ServerExternalAdministratorPtrOutput) ToServerExternalAdministratorPtrOutput() ServerExternalAdministratorPtrOutput {
	return o
}

func (o ServerExternalAdministratorPtrOutput) ToServerExternalAdministratorPtrOutputWithContext(ctx context.Context) ServerExternalAdministratorPtrOutput {
	return o
}

func (o ServerExternalAdministratorPtrOutput) Elem() ServerExternalAdministratorOutput {
	return o.ApplyT(func(v *ServerExternalAdministrator) ServerExternalAdministrator {
		if v != nil {
			return *v
		}
		var ret ServerExternalAdministrator
		return ret
	}).(ServerExternalAdministratorOutput)
}

func (o ServerExternalAdministratorPtrOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministrator) *string {
		if v == nil {
			return nil
		}
		return v.AdministratorType
	}).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorPtrOutput) AzureADOnlyAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministrator) *bool {
		if v == nil {
			return nil
		}
		return v.AzureADOnlyAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o ServerExternalAdministratorPtrOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministrator) *string {
		if v == nil {
			return nil
		}
		return v.Login
	}).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorPtrOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministrator) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalType
	}).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorPtrOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministrator) *string {
		if v == nil {
			return nil
		}
		return v.Sid
	}).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministrator) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ServerExternalAdministratorResponse struct {
	AdministratorType         *string `pulumi:"administratorType"`
	AzureADOnlyAuthentication *bool   `pulumi:"azureADOnlyAuthentication"`
	Login                     *string `pulumi:"login"`
	PrincipalType             *string `pulumi:"principalType"`
	Sid                       *string `pulumi:"sid"`
	TenantId                  *string `pulumi:"tenantId"`
}





type ServerExternalAdministratorResponseInput interface {
	pulumi.Input

	ToServerExternalAdministratorResponseOutput() ServerExternalAdministratorResponseOutput
	ToServerExternalAdministratorResponseOutputWithContext(context.Context) ServerExternalAdministratorResponseOutput
}

type ServerExternalAdministratorResponseArgs struct {
	AdministratorType         pulumi.StringPtrInput `pulumi:"administratorType"`
	AzureADOnlyAuthentication pulumi.BoolPtrInput   `pulumi:"azureADOnlyAuthentication"`
	Login                     pulumi.StringPtrInput `pulumi:"login"`
	PrincipalType             pulumi.StringPtrInput `pulumi:"principalType"`
	Sid                       pulumi.StringPtrInput `pulumi:"sid"`
	TenantId                  pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (ServerExternalAdministratorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerExternalAdministratorResponse)(nil)).Elem()
}

func (i ServerExternalAdministratorResponseArgs) ToServerExternalAdministratorResponseOutput() ServerExternalAdministratorResponseOutput {
	return i.ToServerExternalAdministratorResponseOutputWithContext(context.Background())
}

func (i ServerExternalAdministratorResponseArgs) ToServerExternalAdministratorResponseOutputWithContext(ctx context.Context) ServerExternalAdministratorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerExternalAdministratorResponseOutput)
}

func (i ServerExternalAdministratorResponseArgs) ToServerExternalAdministratorResponsePtrOutput() ServerExternalAdministratorResponsePtrOutput {
	return i.ToServerExternalAdministratorResponsePtrOutputWithContext(context.Background())
}

func (i ServerExternalAdministratorResponseArgs) ToServerExternalAdministratorResponsePtrOutputWithContext(ctx context.Context) ServerExternalAdministratorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerExternalAdministratorResponseOutput).ToServerExternalAdministratorResponsePtrOutputWithContext(ctx)
}









type ServerExternalAdministratorResponsePtrInput interface {
	pulumi.Input

	ToServerExternalAdministratorResponsePtrOutput() ServerExternalAdministratorResponsePtrOutput
	ToServerExternalAdministratorResponsePtrOutputWithContext(context.Context) ServerExternalAdministratorResponsePtrOutput
}

type serverExternalAdministratorResponsePtrType ServerExternalAdministratorResponseArgs

func ServerExternalAdministratorResponsePtr(v *ServerExternalAdministratorResponseArgs) ServerExternalAdministratorResponsePtrInput {
	return (*serverExternalAdministratorResponsePtrType)(v)
}

func (*serverExternalAdministratorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerExternalAdministratorResponse)(nil)).Elem()
}

func (i *serverExternalAdministratorResponsePtrType) ToServerExternalAdministratorResponsePtrOutput() ServerExternalAdministratorResponsePtrOutput {
	return i.ToServerExternalAdministratorResponsePtrOutputWithContext(context.Background())
}

func (i *serverExternalAdministratorResponsePtrType) ToServerExternalAdministratorResponsePtrOutputWithContext(ctx context.Context) ServerExternalAdministratorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerExternalAdministratorResponsePtrOutput)
}

type ServerExternalAdministratorResponseOutput struct{ *pulumi.OutputState }

func (ServerExternalAdministratorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerExternalAdministratorResponse)(nil)).Elem()
}

func (o ServerExternalAdministratorResponseOutput) ToServerExternalAdministratorResponseOutput() ServerExternalAdministratorResponseOutput {
	return o
}

func (o ServerExternalAdministratorResponseOutput) ToServerExternalAdministratorResponseOutputWithContext(ctx context.Context) ServerExternalAdministratorResponseOutput {
	return o
}

func (o ServerExternalAdministratorResponseOutput) ToServerExternalAdministratorResponsePtrOutput() ServerExternalAdministratorResponsePtrOutput {
	return o.ToServerExternalAdministratorResponsePtrOutputWithContext(context.Background())
}

func (o ServerExternalAdministratorResponseOutput) ToServerExternalAdministratorResponsePtrOutputWithContext(ctx context.Context) ServerExternalAdministratorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerExternalAdministratorResponse) *ServerExternalAdministratorResponse {
		return &v
	}).(ServerExternalAdministratorResponsePtrOutput)
}

func (o ServerExternalAdministratorResponseOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministratorResponse) *string { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorResponseOutput) AzureADOnlyAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministratorResponse) *bool { return v.AzureADOnlyAuthentication }).(pulumi.BoolPtrOutput)
}

func (o ServerExternalAdministratorResponseOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministratorResponse) *string { return v.Login }).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorResponseOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministratorResponse) *string { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorResponseOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministratorResponse) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerExternalAdministratorResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ServerExternalAdministratorResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerExternalAdministratorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerExternalAdministratorResponse)(nil)).Elem()
}

func (o ServerExternalAdministratorResponsePtrOutput) ToServerExternalAdministratorResponsePtrOutput() ServerExternalAdministratorResponsePtrOutput {
	return o
}

func (o ServerExternalAdministratorResponsePtrOutput) ToServerExternalAdministratorResponsePtrOutputWithContext(ctx context.Context) ServerExternalAdministratorResponsePtrOutput {
	return o
}

func (o ServerExternalAdministratorResponsePtrOutput) Elem() ServerExternalAdministratorResponseOutput {
	return o.ApplyT(func(v *ServerExternalAdministratorResponse) ServerExternalAdministratorResponse {
		if v != nil {
			return *v
		}
		var ret ServerExternalAdministratorResponse
		return ret
	}).(ServerExternalAdministratorResponseOutput)
}

func (o ServerExternalAdministratorResponsePtrOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministratorResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdministratorType
	}).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorResponsePtrOutput) AzureADOnlyAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministratorResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AzureADOnlyAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o ServerExternalAdministratorResponsePtrOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministratorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Login
	}).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorResponsePtrOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministratorResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalType
	}).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorResponsePtrOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministratorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sid
	}).(pulumi.StringPtrOutput)
}

func (o ServerExternalAdministratorResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerExternalAdministratorResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ServerInfo struct {
	ServerId string `pulumi:"serverId"`
}





type ServerInfoInput interface {
	pulumi.Input

	ToServerInfoOutput() ServerInfoOutput
	ToServerInfoOutputWithContext(context.Context) ServerInfoOutput
}

type ServerInfoArgs struct {
	ServerId pulumi.StringInput `pulumi:"serverId"`
}

func (ServerInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerInfo)(nil)).Elem()
}

func (i ServerInfoArgs) ToServerInfoOutput() ServerInfoOutput {
	return i.ToServerInfoOutputWithContext(context.Background())
}

func (i ServerInfoArgs) ToServerInfoOutputWithContext(ctx context.Context) ServerInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerInfoOutput)
}





type ServerInfoArrayInput interface {
	pulumi.Input

	ToServerInfoArrayOutput() ServerInfoArrayOutput
	ToServerInfoArrayOutputWithContext(context.Context) ServerInfoArrayOutput
}

type ServerInfoArray []ServerInfoInput

func (ServerInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerInfo)(nil)).Elem()
}

func (i ServerInfoArray) ToServerInfoArrayOutput() ServerInfoArrayOutput {
	return i.ToServerInfoArrayOutputWithContext(context.Background())
}

func (i ServerInfoArray) ToServerInfoArrayOutputWithContext(ctx context.Context) ServerInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerInfoArrayOutput)
}

type ServerInfoOutput struct{ *pulumi.OutputState }

func (ServerInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerInfo)(nil)).Elem()
}

func (o ServerInfoOutput) ToServerInfoOutput() ServerInfoOutput {
	return o
}

func (o ServerInfoOutput) ToServerInfoOutputWithContext(ctx context.Context) ServerInfoOutput {
	return o
}

func (o ServerInfoOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v ServerInfo) string { return v.ServerId }).(pulumi.StringOutput)
}

type ServerInfoArrayOutput struct{ *pulumi.OutputState }

func (ServerInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerInfo)(nil)).Elem()
}

func (o ServerInfoArrayOutput) ToServerInfoArrayOutput() ServerInfoArrayOutput {
	return o
}

func (o ServerInfoArrayOutput) ToServerInfoArrayOutputWithContext(ctx context.Context) ServerInfoArrayOutput {
	return o
}

func (o ServerInfoArrayOutput) Index(i pulumi.IntInput) ServerInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerInfo {
		return vs[0].([]ServerInfo)[vs[1].(int)]
	}).(ServerInfoOutput)
}

type ServerInfoResponse struct {
	ServerId string `pulumi:"serverId"`
}





type ServerInfoResponseInput interface {
	pulumi.Input

	ToServerInfoResponseOutput() ServerInfoResponseOutput
	ToServerInfoResponseOutputWithContext(context.Context) ServerInfoResponseOutput
}

type ServerInfoResponseArgs struct {
	ServerId pulumi.StringInput `pulumi:"serverId"`
}

func (ServerInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerInfoResponse)(nil)).Elem()
}

func (i ServerInfoResponseArgs) ToServerInfoResponseOutput() ServerInfoResponseOutput {
	return i.ToServerInfoResponseOutputWithContext(context.Background())
}

func (i ServerInfoResponseArgs) ToServerInfoResponseOutputWithContext(ctx context.Context) ServerInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerInfoResponseOutput)
}





type ServerInfoResponseArrayInput interface {
	pulumi.Input

	ToServerInfoResponseArrayOutput() ServerInfoResponseArrayOutput
	ToServerInfoResponseArrayOutputWithContext(context.Context) ServerInfoResponseArrayOutput
}

type ServerInfoResponseArray []ServerInfoResponseInput

func (ServerInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerInfoResponse)(nil)).Elem()
}

func (i ServerInfoResponseArray) ToServerInfoResponseArrayOutput() ServerInfoResponseArrayOutput {
	return i.ToServerInfoResponseArrayOutputWithContext(context.Background())
}

func (i ServerInfoResponseArray) ToServerInfoResponseArrayOutputWithContext(ctx context.Context) ServerInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerInfoResponseArrayOutput)
}

type ServerInfoResponseOutput struct{ *pulumi.OutputState }

func (ServerInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerInfoResponse)(nil)).Elem()
}

func (o ServerInfoResponseOutput) ToServerInfoResponseOutput() ServerInfoResponseOutput {
	return o
}

func (o ServerInfoResponseOutput) ToServerInfoResponseOutputWithContext(ctx context.Context) ServerInfoResponseOutput {
	return o
}

func (o ServerInfoResponseOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v ServerInfoResponse) string { return v.ServerId }).(pulumi.StringOutput)
}

type ServerInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ServerInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerInfoResponse)(nil)).Elem()
}

func (o ServerInfoResponseArrayOutput) ToServerInfoResponseArrayOutput() ServerInfoResponseArrayOutput {
	return o
}

func (o ServerInfoResponseArrayOutput) ToServerInfoResponseArrayOutputWithContext(ctx context.Context) ServerInfoResponseArrayOutput {
	return o
}

func (o ServerInfoResponseArrayOutput) Index(i pulumi.IntInput) ServerInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerInfoResponse {
		return vs[0].([]ServerInfoResponse)[vs[1].(int)]
	}).(ServerInfoResponseOutput)
}

type ServerPrivateEndpointConnectionResponse struct {
	Id         string                                      `pulumi:"id"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
}





type ServerPrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToServerPrivateEndpointConnectionResponseOutput() ServerPrivateEndpointConnectionResponseOutput
	ToServerPrivateEndpointConnectionResponseOutputWithContext(context.Context) ServerPrivateEndpointConnectionResponseOutput
}

type ServerPrivateEndpointConnectionResponseArgs struct {
	Id         pulumi.StringInput                               `pulumi:"id"`
	Properties PrivateEndpointConnectionPropertiesResponseInput `pulumi:"properties"`
}

func (ServerPrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i ServerPrivateEndpointConnectionResponseArgs) ToServerPrivateEndpointConnectionResponseOutput() ServerPrivateEndpointConnectionResponseOutput {
	return i.ToServerPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i ServerPrivateEndpointConnectionResponseArgs) ToServerPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPrivateEndpointConnectionResponseOutput)
}





type ServerPrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToServerPrivateEndpointConnectionResponseArrayOutput() ServerPrivateEndpointConnectionResponseArrayOutput
	ToServerPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) ServerPrivateEndpointConnectionResponseArrayOutput
}

type ServerPrivateEndpointConnectionResponseArray []ServerPrivateEndpointConnectionResponseInput

func (ServerPrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i ServerPrivateEndpointConnectionResponseArray) ToServerPrivateEndpointConnectionResponseArrayOutput() ServerPrivateEndpointConnectionResponseArrayOutput {
	return i.ToServerPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i ServerPrivateEndpointConnectionResponseArray) ToServerPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPrivateEndpointConnectionResponseArrayOutput)
}

type ServerPrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (ServerPrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o ServerPrivateEndpointConnectionResponseOutput) ToServerPrivateEndpointConnectionResponseOutput() ServerPrivateEndpointConnectionResponseOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseOutput) ToServerPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionResponseOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ServerPrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v ServerPrivateEndpointConnectionResponse) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

type ServerPrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (ServerPrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o ServerPrivateEndpointConnectionResponseArrayOutput) ToServerPrivateEndpointConnectionResponseArrayOutput() ServerPrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseArrayOutput) ToServerPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) ServerPrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerPrivateEndpointConnectionResponse {
		return vs[0].([]ServerPrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(ServerPrivateEndpointConnectionResponseOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SyncGroupSchema struct {
	MasterSyncMemberName *string                `pulumi:"masterSyncMemberName"`
	Tables               []SyncGroupSchemaTable `pulumi:"tables"`
}





type SyncGroupSchemaInput interface {
	pulumi.Input

	ToSyncGroupSchemaOutput() SyncGroupSchemaOutput
	ToSyncGroupSchemaOutputWithContext(context.Context) SyncGroupSchemaOutput
}

type SyncGroupSchemaArgs struct {
	MasterSyncMemberName pulumi.StringPtrInput          `pulumi:"masterSyncMemberName"`
	Tables               SyncGroupSchemaTableArrayInput `pulumi:"tables"`
}

func (SyncGroupSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchema)(nil)).Elem()
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaOutput() SyncGroupSchemaOutput {
	return i.ToSyncGroupSchemaOutputWithContext(context.Background())
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaOutputWithContext(ctx context.Context) SyncGroupSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaOutput)
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return i.ToSyncGroupSchemaPtrOutputWithContext(context.Background())
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaOutput).ToSyncGroupSchemaPtrOutputWithContext(ctx)
}









type SyncGroupSchemaPtrInput interface {
	pulumi.Input

	ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput
	ToSyncGroupSchemaPtrOutputWithContext(context.Context) SyncGroupSchemaPtrOutput
}

type syncGroupSchemaPtrType SyncGroupSchemaArgs

func SyncGroupSchemaPtr(v *SyncGroupSchemaArgs) SyncGroupSchemaPtrInput {
	return (*syncGroupSchemaPtrType)(v)
}

func (*syncGroupSchemaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroupSchema)(nil)).Elem()
}

func (i *syncGroupSchemaPtrType) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return i.ToSyncGroupSchemaPtrOutputWithContext(context.Background())
}

func (i *syncGroupSchemaPtrType) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaPtrOutput)
}

type SyncGroupSchemaOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchema)(nil)).Elem()
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaOutput() SyncGroupSchemaOutput {
	return o
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaOutputWithContext(ctx context.Context) SyncGroupSchemaOutput {
	return o
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return o.ToSyncGroupSchemaPtrOutputWithContext(context.Background())
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncGroupSchema) *SyncGroupSchema {
		return &v
	}).(SyncGroupSchemaPtrOutput)
}

func (o SyncGroupSchemaOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchema) *string { return v.MasterSyncMemberName }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaOutput) Tables() SyncGroupSchemaTableArrayOutput {
	return o.ApplyT(func(v SyncGroupSchema) []SyncGroupSchemaTable { return v.Tables }).(SyncGroupSchemaTableArrayOutput)
}

type SyncGroupSchemaPtrOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroupSchema)(nil)).Elem()
}

func (o SyncGroupSchemaPtrOutput) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return o
}

func (o SyncGroupSchemaPtrOutput) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return o
}

func (o SyncGroupSchemaPtrOutput) Elem() SyncGroupSchemaOutput {
	return o.ApplyT(func(v *SyncGroupSchema) SyncGroupSchema {
		if v != nil {
			return *v
		}
		var ret SyncGroupSchema
		return ret
	}).(SyncGroupSchemaOutput)
}

func (o SyncGroupSchemaPtrOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncGroupSchema) *string {
		if v == nil {
			return nil
		}
		return v.MasterSyncMemberName
	}).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaPtrOutput) Tables() SyncGroupSchemaTableArrayOutput {
	return o.ApplyT(func(v *SyncGroupSchema) []SyncGroupSchemaTable {
		if v == nil {
			return nil
		}
		return v.Tables
	}).(SyncGroupSchemaTableArrayOutput)
}

type SyncGroupSchemaResponse struct {
	MasterSyncMemberName *string                        `pulumi:"masterSyncMemberName"`
	Tables               []SyncGroupSchemaTableResponse `pulumi:"tables"`
}





type SyncGroupSchemaResponseInput interface {
	pulumi.Input

	ToSyncGroupSchemaResponseOutput() SyncGroupSchemaResponseOutput
	ToSyncGroupSchemaResponseOutputWithContext(context.Context) SyncGroupSchemaResponseOutput
}

type SyncGroupSchemaResponseArgs struct {
	MasterSyncMemberName pulumi.StringPtrInput                  `pulumi:"masterSyncMemberName"`
	Tables               SyncGroupSchemaTableResponseArrayInput `pulumi:"tables"`
}

func (SyncGroupSchemaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaResponse)(nil)).Elem()
}

func (i SyncGroupSchemaResponseArgs) ToSyncGroupSchemaResponseOutput() SyncGroupSchemaResponseOutput {
	return i.ToSyncGroupSchemaResponseOutputWithContext(context.Background())
}

func (i SyncGroupSchemaResponseArgs) ToSyncGroupSchemaResponseOutputWithContext(ctx context.Context) SyncGroupSchemaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaResponseOutput)
}

func (i SyncGroupSchemaResponseArgs) ToSyncGroupSchemaResponsePtrOutput() SyncGroupSchemaResponsePtrOutput {
	return i.ToSyncGroupSchemaResponsePtrOutputWithContext(context.Background())
}

func (i SyncGroupSchemaResponseArgs) ToSyncGroupSchemaResponsePtrOutputWithContext(ctx context.Context) SyncGroupSchemaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaResponseOutput).ToSyncGroupSchemaResponsePtrOutputWithContext(ctx)
}









type SyncGroupSchemaResponsePtrInput interface {
	pulumi.Input

	ToSyncGroupSchemaResponsePtrOutput() SyncGroupSchemaResponsePtrOutput
	ToSyncGroupSchemaResponsePtrOutputWithContext(context.Context) SyncGroupSchemaResponsePtrOutput
}

type syncGroupSchemaResponsePtrType SyncGroupSchemaResponseArgs

func SyncGroupSchemaResponsePtr(v *SyncGroupSchemaResponseArgs) SyncGroupSchemaResponsePtrInput {
	return (*syncGroupSchemaResponsePtrType)(v)
}

func (*syncGroupSchemaResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroupSchemaResponse)(nil)).Elem()
}

func (i *syncGroupSchemaResponsePtrType) ToSyncGroupSchemaResponsePtrOutput() SyncGroupSchemaResponsePtrOutput {
	return i.ToSyncGroupSchemaResponsePtrOutputWithContext(context.Background())
}

func (i *syncGroupSchemaResponsePtrType) ToSyncGroupSchemaResponsePtrOutputWithContext(ctx context.Context) SyncGroupSchemaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaResponsePtrOutput)
}

type SyncGroupSchemaResponseOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaResponse)(nil)).Elem()
}

func (o SyncGroupSchemaResponseOutput) ToSyncGroupSchemaResponseOutput() SyncGroupSchemaResponseOutput {
	return o
}

func (o SyncGroupSchemaResponseOutput) ToSyncGroupSchemaResponseOutputWithContext(ctx context.Context) SyncGroupSchemaResponseOutput {
	return o
}

func (o SyncGroupSchemaResponseOutput) ToSyncGroupSchemaResponsePtrOutput() SyncGroupSchemaResponsePtrOutput {
	return o.ToSyncGroupSchemaResponsePtrOutputWithContext(context.Background())
}

func (o SyncGroupSchemaResponseOutput) ToSyncGroupSchemaResponsePtrOutputWithContext(ctx context.Context) SyncGroupSchemaResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncGroupSchemaResponse) *SyncGroupSchemaResponse {
		return &v
	}).(SyncGroupSchemaResponsePtrOutput)
}

func (o SyncGroupSchemaResponseOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaResponse) *string { return v.MasterSyncMemberName }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaResponseOutput) Tables() SyncGroupSchemaTableResponseArrayOutput {
	return o.ApplyT(func(v SyncGroupSchemaResponse) []SyncGroupSchemaTableResponse { return v.Tables }).(SyncGroupSchemaTableResponseArrayOutput)
}

type SyncGroupSchemaResponsePtrOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroupSchemaResponse)(nil)).Elem()
}

func (o SyncGroupSchemaResponsePtrOutput) ToSyncGroupSchemaResponsePtrOutput() SyncGroupSchemaResponsePtrOutput {
	return o
}

func (o SyncGroupSchemaResponsePtrOutput) ToSyncGroupSchemaResponsePtrOutputWithContext(ctx context.Context) SyncGroupSchemaResponsePtrOutput {
	return o
}

func (o SyncGroupSchemaResponsePtrOutput) Elem() SyncGroupSchemaResponseOutput {
	return o.ApplyT(func(v *SyncGroupSchemaResponse) SyncGroupSchemaResponse {
		if v != nil {
			return *v
		}
		var ret SyncGroupSchemaResponse
		return ret
	}).(SyncGroupSchemaResponseOutput)
}

func (o SyncGroupSchemaResponsePtrOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncGroupSchemaResponse) *string {
		if v == nil {
			return nil
		}
		return v.MasterSyncMemberName
	}).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaResponsePtrOutput) Tables() SyncGroupSchemaTableResponseArrayOutput {
	return o.ApplyT(func(v *SyncGroupSchemaResponse) []SyncGroupSchemaTableResponse {
		if v == nil {
			return nil
		}
		return v.Tables
	}).(SyncGroupSchemaTableResponseArrayOutput)
}

type SyncGroupSchemaTable struct {
	Columns    []SyncGroupSchemaTableColumn `pulumi:"columns"`
	QuotedName *string                      `pulumi:"quotedName"`
}





type SyncGroupSchemaTableInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableOutput() SyncGroupSchemaTableOutput
	ToSyncGroupSchemaTableOutputWithContext(context.Context) SyncGroupSchemaTableOutput
}

type SyncGroupSchemaTableArgs struct {
	Columns    SyncGroupSchemaTableColumnArrayInput `pulumi:"columns"`
	QuotedName pulumi.StringPtrInput                `pulumi:"quotedName"`
}

func (SyncGroupSchemaTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTable)(nil)).Elem()
}

func (i SyncGroupSchemaTableArgs) ToSyncGroupSchemaTableOutput() SyncGroupSchemaTableOutput {
	return i.ToSyncGroupSchemaTableOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableArgs) ToSyncGroupSchemaTableOutputWithContext(ctx context.Context) SyncGroupSchemaTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableOutput)
}





type SyncGroupSchemaTableArrayInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableArrayOutput() SyncGroupSchemaTableArrayOutput
	ToSyncGroupSchemaTableArrayOutputWithContext(context.Context) SyncGroupSchemaTableArrayOutput
}

type SyncGroupSchemaTableArray []SyncGroupSchemaTableInput

func (SyncGroupSchemaTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTable)(nil)).Elem()
}

func (i SyncGroupSchemaTableArray) ToSyncGroupSchemaTableArrayOutput() SyncGroupSchemaTableArrayOutput {
	return i.ToSyncGroupSchemaTableArrayOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableArray) ToSyncGroupSchemaTableArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableArrayOutput)
}

type SyncGroupSchemaTableOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTable)(nil)).Elem()
}

func (o SyncGroupSchemaTableOutput) ToSyncGroupSchemaTableOutput() SyncGroupSchemaTableOutput {
	return o
}

func (o SyncGroupSchemaTableOutput) ToSyncGroupSchemaTableOutputWithContext(ctx context.Context) SyncGroupSchemaTableOutput {
	return o
}

func (o SyncGroupSchemaTableOutput) Columns() SyncGroupSchemaTableColumnArrayOutput {
	return o.ApplyT(func(v SyncGroupSchemaTable) []SyncGroupSchemaTableColumn { return v.Columns }).(SyncGroupSchemaTableColumnArrayOutput)
}

func (o SyncGroupSchemaTableOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTable) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTable)(nil)).Elem()
}

func (o SyncGroupSchemaTableArrayOutput) ToSyncGroupSchemaTableArrayOutput() SyncGroupSchemaTableArrayOutput {
	return o
}

func (o SyncGroupSchemaTableArrayOutput) ToSyncGroupSchemaTableArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableArrayOutput {
	return o
}

func (o SyncGroupSchemaTableArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTable {
		return vs[0].([]SyncGroupSchemaTable)[vs[1].(int)]
	}).(SyncGroupSchemaTableOutput)
}

type SyncGroupSchemaTableColumn struct {
	DataSize   *string `pulumi:"dataSize"`
	DataType   *string `pulumi:"dataType"`
	QuotedName *string `pulumi:"quotedName"`
}





type SyncGroupSchemaTableColumnInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableColumnOutput() SyncGroupSchemaTableColumnOutput
	ToSyncGroupSchemaTableColumnOutputWithContext(context.Context) SyncGroupSchemaTableColumnOutput
}

type SyncGroupSchemaTableColumnArgs struct {
	DataSize   pulumi.StringPtrInput `pulumi:"dataSize"`
	DataType   pulumi.StringPtrInput `pulumi:"dataType"`
	QuotedName pulumi.StringPtrInput `pulumi:"quotedName"`
}

func (SyncGroupSchemaTableColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (i SyncGroupSchemaTableColumnArgs) ToSyncGroupSchemaTableColumnOutput() SyncGroupSchemaTableColumnOutput {
	return i.ToSyncGroupSchemaTableColumnOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableColumnArgs) ToSyncGroupSchemaTableColumnOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableColumnOutput)
}





type SyncGroupSchemaTableColumnArrayInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableColumnArrayOutput() SyncGroupSchemaTableColumnArrayOutput
	ToSyncGroupSchemaTableColumnArrayOutputWithContext(context.Context) SyncGroupSchemaTableColumnArrayOutput
}

type SyncGroupSchemaTableColumnArray []SyncGroupSchemaTableColumnInput

func (SyncGroupSchemaTableColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (i SyncGroupSchemaTableColumnArray) ToSyncGroupSchemaTableColumnArrayOutput() SyncGroupSchemaTableColumnArrayOutput {
	return i.ToSyncGroupSchemaTableColumnArrayOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableColumnArray) ToSyncGroupSchemaTableColumnArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableColumnArrayOutput)
}

type SyncGroupSchemaTableColumnOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnOutput) ToSyncGroupSchemaTableColumnOutput() SyncGroupSchemaTableColumnOutput {
	return o
}

func (o SyncGroupSchemaTableColumnOutput) ToSyncGroupSchemaTableColumnOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnOutput {
	return o
}

func (o SyncGroupSchemaTableColumnOutput) DataSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumn) *string { return v.DataSize }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumn) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumn) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableColumnArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnArrayOutput) ToSyncGroupSchemaTableColumnArrayOutput() SyncGroupSchemaTableColumnArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnArrayOutput) ToSyncGroupSchemaTableColumnArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTableColumn {
		return vs[0].([]SyncGroupSchemaTableColumn)[vs[1].(int)]
	}).(SyncGroupSchemaTableColumnOutput)
}

type SyncGroupSchemaTableColumnResponse struct {
	DataSize   *string `pulumi:"dataSize"`
	DataType   *string `pulumi:"dataType"`
	QuotedName *string `pulumi:"quotedName"`
}





type SyncGroupSchemaTableColumnResponseInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableColumnResponseOutput() SyncGroupSchemaTableColumnResponseOutput
	ToSyncGroupSchemaTableColumnResponseOutputWithContext(context.Context) SyncGroupSchemaTableColumnResponseOutput
}

type SyncGroupSchemaTableColumnResponseArgs struct {
	DataSize   pulumi.StringPtrInput `pulumi:"dataSize"`
	DataType   pulumi.StringPtrInput `pulumi:"dataType"`
	QuotedName pulumi.StringPtrInput `pulumi:"quotedName"`
}

func (SyncGroupSchemaTableColumnResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableColumnResponse)(nil)).Elem()
}

func (i SyncGroupSchemaTableColumnResponseArgs) ToSyncGroupSchemaTableColumnResponseOutput() SyncGroupSchemaTableColumnResponseOutput {
	return i.ToSyncGroupSchemaTableColumnResponseOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableColumnResponseArgs) ToSyncGroupSchemaTableColumnResponseOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableColumnResponseOutput)
}





type SyncGroupSchemaTableColumnResponseArrayInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableColumnResponseArrayOutput() SyncGroupSchemaTableColumnResponseArrayOutput
	ToSyncGroupSchemaTableColumnResponseArrayOutputWithContext(context.Context) SyncGroupSchemaTableColumnResponseArrayOutput
}

type SyncGroupSchemaTableColumnResponseArray []SyncGroupSchemaTableColumnResponseInput

func (SyncGroupSchemaTableColumnResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableColumnResponse)(nil)).Elem()
}

func (i SyncGroupSchemaTableColumnResponseArray) ToSyncGroupSchemaTableColumnResponseArrayOutput() SyncGroupSchemaTableColumnResponseArrayOutput {
	return i.ToSyncGroupSchemaTableColumnResponseArrayOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableColumnResponseArray) ToSyncGroupSchemaTableColumnResponseArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableColumnResponseArrayOutput)
}

type SyncGroupSchemaTableColumnResponseOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableColumnResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnResponseOutput) ToSyncGroupSchemaTableColumnResponseOutput() SyncGroupSchemaTableColumnResponseOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseOutput) ToSyncGroupSchemaTableColumnResponseOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnResponseOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseOutput) DataSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumnResponse) *string { return v.DataSize }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumnResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnResponseOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumnResponse) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableColumnResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnResponseArrayOutput) ToSyncGroupSchemaTableColumnResponseArrayOutput() SyncGroupSchemaTableColumnResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseArrayOutput) ToSyncGroupSchemaTableColumnResponseArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTableColumnResponse {
		return vs[0].([]SyncGroupSchemaTableColumnResponse)[vs[1].(int)]
	}).(SyncGroupSchemaTableColumnResponseOutput)
}

type SyncGroupSchemaTableResponse struct {
	Columns    []SyncGroupSchemaTableColumnResponse `pulumi:"columns"`
	QuotedName *string                              `pulumi:"quotedName"`
}





type SyncGroupSchemaTableResponseInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableResponseOutput() SyncGroupSchemaTableResponseOutput
	ToSyncGroupSchemaTableResponseOutputWithContext(context.Context) SyncGroupSchemaTableResponseOutput
}

type SyncGroupSchemaTableResponseArgs struct {
	Columns    SyncGroupSchemaTableColumnResponseArrayInput `pulumi:"columns"`
	QuotedName pulumi.StringPtrInput                        `pulumi:"quotedName"`
}

func (SyncGroupSchemaTableResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableResponse)(nil)).Elem()
}

func (i SyncGroupSchemaTableResponseArgs) ToSyncGroupSchemaTableResponseOutput() SyncGroupSchemaTableResponseOutput {
	return i.ToSyncGroupSchemaTableResponseOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableResponseArgs) ToSyncGroupSchemaTableResponseOutputWithContext(ctx context.Context) SyncGroupSchemaTableResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableResponseOutput)
}





type SyncGroupSchemaTableResponseArrayInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableResponseArrayOutput() SyncGroupSchemaTableResponseArrayOutput
	ToSyncGroupSchemaTableResponseArrayOutputWithContext(context.Context) SyncGroupSchemaTableResponseArrayOutput
}

type SyncGroupSchemaTableResponseArray []SyncGroupSchemaTableResponseInput

func (SyncGroupSchemaTableResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableResponse)(nil)).Elem()
}

func (i SyncGroupSchemaTableResponseArray) ToSyncGroupSchemaTableResponseArrayOutput() SyncGroupSchemaTableResponseArrayOutput {
	return i.ToSyncGroupSchemaTableResponseArrayOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableResponseArray) ToSyncGroupSchemaTableResponseArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableResponseArrayOutput)
}

type SyncGroupSchemaTableResponseOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableResponseOutput) ToSyncGroupSchemaTableResponseOutput() SyncGroupSchemaTableResponseOutput {
	return o
}

func (o SyncGroupSchemaTableResponseOutput) ToSyncGroupSchemaTableResponseOutputWithContext(ctx context.Context) SyncGroupSchemaTableResponseOutput {
	return o
}

func (o SyncGroupSchemaTableResponseOutput) Columns() SyncGroupSchemaTableColumnResponseArrayOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableResponse) []SyncGroupSchemaTableColumnResponse { return v.Columns }).(SyncGroupSchemaTableColumnResponseArrayOutput)
}

func (o SyncGroupSchemaTableResponseOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableResponse) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableResponseArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableResponseArrayOutput) ToSyncGroupSchemaTableResponseArrayOutput() SyncGroupSchemaTableResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableResponseArrayOutput) ToSyncGroupSchemaTableResponseArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableResponseArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTableResponse {
		return vs[0].([]SyncGroupSchemaTableResponse)[vs[1].(int)]
	}).(SyncGroupSchemaTableResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type UserIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type UserIdentityResponseInput interface {
	pulumi.Input

	ToUserIdentityResponseOutput() UserIdentityResponseOutput
	ToUserIdentityResponseOutputWithContext(context.Context) UserIdentityResponseOutput
}

type UserIdentityResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (UserIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityResponse)(nil)).Elem()
}

func (i UserIdentityResponseArgs) ToUserIdentityResponseOutput() UserIdentityResponseOutput {
	return i.ToUserIdentityResponseOutputWithContext(context.Background())
}

func (i UserIdentityResponseArgs) ToUserIdentityResponseOutputWithContext(ctx context.Context) UserIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityResponseOutput)
}





type UserIdentityResponseMapInput interface {
	pulumi.Input

	ToUserIdentityResponseMapOutput() UserIdentityResponseMapOutput
	ToUserIdentityResponseMapOutputWithContext(context.Context) UserIdentityResponseMapOutput
}

type UserIdentityResponseMap map[string]UserIdentityResponseInput

func (UserIdentityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityResponse)(nil)).Elem()
}

func (i UserIdentityResponseMap) ToUserIdentityResponseMapOutput() UserIdentityResponseMapOutput {
	return i.ToUserIdentityResponseMapOutputWithContext(context.Background())
}

func (i UserIdentityResponseMap) ToUserIdentityResponseMapOutputWithContext(ctx context.Context) UserIdentityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityResponseMapOutput)
}

type UserIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutput() UserIdentityResponseOutput {
	return o
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutputWithContext(ctx context.Context) UserIdentityResponseOutput {
	return o
}

func (o UserIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponseMapOutput) ToUserIdentityResponseMapOutput() UserIdentityResponseMapOutput {
	return o
}

func (o UserIdentityResponseMapOutput) ToUserIdentityResponseMapOutputWithContext(ctx context.Context) UserIdentityResponseMapOutput {
	return o
}

func (o UserIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityResponse {
		return vs[0].(map[string]UserIdentityResponse)[vs[1].(string)]
	}).(UserIdentityResponseOutput)
}

type VulnerabilityAssessmentRecurringScansProperties struct {
	EmailSubscriptionAdmins *bool    `pulumi:"emailSubscriptionAdmins"`
	Emails                  []string `pulumi:"emails"`
	IsEnabled               *bool    `pulumi:"isEnabled"`
}





type VulnerabilityAssessmentRecurringScansPropertiesInput interface {
	pulumi.Input

	ToVulnerabilityAssessmentRecurringScansPropertiesOutput() VulnerabilityAssessmentRecurringScansPropertiesOutput
	ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(context.Context) VulnerabilityAssessmentRecurringScansPropertiesOutput
}

type VulnerabilityAssessmentRecurringScansPropertiesArgs struct {
	EmailSubscriptionAdmins pulumi.BoolPtrInput     `pulumi:"emailSubscriptionAdmins"`
	Emails                  pulumi.StringArrayInput `pulumi:"emails"`
	IsEnabled               pulumi.BoolPtrInput     `pulumi:"isEnabled"`
}

func (VulnerabilityAssessmentRecurringScansPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesOutput() VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(context.Background())
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesOutput)
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Background())
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesOutput).ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx)
}









type VulnerabilityAssessmentRecurringScansPropertiesPtrInput interface {
	pulumi.Input

	ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput
	ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput
}

type vulnerabilityAssessmentRecurringScansPropertiesPtrType VulnerabilityAssessmentRecurringScansPropertiesArgs

func VulnerabilityAssessmentRecurringScansPropertiesPtr(v *VulnerabilityAssessmentRecurringScansPropertiesArgs) VulnerabilityAssessmentRecurringScansPropertiesPtrInput {
	return (*vulnerabilityAssessmentRecurringScansPropertiesPtrType)(v)
}

func (*vulnerabilityAssessmentRecurringScansPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (i *vulnerabilityAssessmentRecurringScansPropertiesPtrType) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Background())
}

func (i *vulnerabilityAssessmentRecurringScansPropertiesPtrType) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesOutput() VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o.ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Background())
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VulnerabilityAssessmentRecurringScansProperties) *VulnerabilityAssessmentRecurringScansProperties {
		return &v
	}).(VulnerabilityAssessmentRecurringScansPropertiesPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansProperties) *bool { return v.EmailSubscriptionAdmins }).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansProperties) []string { return v.Emails }).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansProperties) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) Elem() VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) VulnerabilityAssessmentRecurringScansProperties {
		if v != nil {
			return *v
		}
		var ret VulnerabilityAssessmentRecurringScansProperties
		return ret
	}).(VulnerabilityAssessmentRecurringScansPropertiesOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EmailSubscriptionAdmins
	}).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) []string {
		if v == nil {
			return nil
		}
		return v.Emails
	}).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesResponse struct {
	EmailSubscriptionAdmins *bool    `pulumi:"emailSubscriptionAdmins"`
	Emails                  []string `pulumi:"emails"`
	IsEnabled               *bool    `pulumi:"isEnabled"`
}





type VulnerabilityAssessmentRecurringScansPropertiesResponseInput interface {
	pulumi.Input

	ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutput() VulnerabilityAssessmentRecurringScansPropertiesResponseOutput
	ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutputWithContext(context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponseOutput
}

type VulnerabilityAssessmentRecurringScansPropertiesResponseArgs struct {
	EmailSubscriptionAdmins pulumi.BoolPtrInput     `pulumi:"emailSubscriptionAdmins"`
	Emails                  pulumi.StringArrayInput `pulumi:"emails"`
	IsEnabled               pulumi.BoolPtrInput     `pulumi:"isEnabled"`
}

func (VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesResponse)(nil)).Elem()
}

func (i VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutput() VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutputWithContext(context.Background())
}

func (i VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput)
}

func (i VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput() VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput).ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(ctx)
}









type VulnerabilityAssessmentRecurringScansPropertiesResponsePtrInput interface {
	pulumi.Input

	ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput() VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput
	ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput
}

type vulnerabilityAssessmentRecurringScansPropertiesResponsePtrType VulnerabilityAssessmentRecurringScansPropertiesResponseArgs

func VulnerabilityAssessmentRecurringScansPropertiesResponsePtr(v *VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrInput {
	return (*vulnerabilityAssessmentRecurringScansPropertiesResponsePtrType)(v)
}

func (*vulnerabilityAssessmentRecurringScansPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VulnerabilityAssessmentRecurringScansPropertiesResponse)(nil)).Elem()
}

func (i *vulnerabilityAssessmentRecurringScansPropertiesResponsePtrType) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput() VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *vulnerabilityAssessmentRecurringScansPropertiesResponsePtrType) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesResponse)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutput() VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput() VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return o.ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VulnerabilityAssessmentRecurringScansPropertiesResponse) *VulnerabilityAssessmentRecurringScansPropertiesResponse {
		return &v
	}).(VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool {
		return v.EmailSubscriptionAdmins
	}).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansPropertiesResponse) []string { return v.Emails }).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VulnerabilityAssessmentRecurringScansPropertiesResponse)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput() VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) Elem() VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) VulnerabilityAssessmentRecurringScansPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VulnerabilityAssessmentRecurringScansPropertiesResponse
		return ret
	}).(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EmailSubscriptionAdmins
	}).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Emails
	}).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemInput)(nil)).Elem(), DatabaseVulnerabilityAssessmentRuleBaselineItemArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemArrayInput)(nil)).Elem(), DatabaseVulnerabilityAssessmentRuleBaselineItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemResponseInput)(nil)).Elem(), DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayInput)(nil)).Elem(), DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolPerDatabaseSettingsInput)(nil)).Elem(), ElasticPoolPerDatabaseSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolPerDatabaseSettingsPtrInput)(nil)).Elem(), ElasticPoolPerDatabaseSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolPerDatabaseSettingsResponseInput)(nil)).Elem(), ElasticPoolPerDatabaseSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolPerDatabaseSettingsResponsePtrInput)(nil)).Elem(), ElasticPoolPerDatabaseSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FailoverGroupReadOnlyEndpointInput)(nil)).Elem(), FailoverGroupReadOnlyEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FailoverGroupReadOnlyEndpointPtrInput)(nil)).Elem(), FailoverGroupReadOnlyEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FailoverGroupReadOnlyEndpointResponseInput)(nil)).Elem(), FailoverGroupReadOnlyEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FailoverGroupReadOnlyEndpointResponsePtrInput)(nil)).Elem(), FailoverGroupReadOnlyEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FailoverGroupReadWriteEndpointInput)(nil)).Elem(), FailoverGroupReadWriteEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FailoverGroupReadWriteEndpointPtrInput)(nil)).Elem(), FailoverGroupReadWriteEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FailoverGroupReadWriteEndpointResponseInput)(nil)).Elem(), FailoverGroupReadWriteEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FailoverGroupReadWriteEndpointResponsePtrInput)(nil)).Elem(), FailoverGroupReadWriteEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpointInput)(nil)).Elem(), InstanceFailoverGroupReadOnlyEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpointPtrInput)(nil)).Elem(), InstanceFailoverGroupReadOnlyEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpointResponseInput)(nil)).Elem(), InstanceFailoverGroupReadOnlyEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpointResponsePtrInput)(nil)).Elem(), InstanceFailoverGroupReadOnlyEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpointInput)(nil)).Elem(), InstanceFailoverGroupReadWriteEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpointPtrInput)(nil)).Elem(), InstanceFailoverGroupReadWriteEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpointResponseInput)(nil)).Elem(), InstanceFailoverGroupReadWriteEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpointResponsePtrInput)(nil)).Elem(), InstanceFailoverGroupReadWriteEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobScheduleInput)(nil)).Elem(), JobScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobSchedulePtrInput)(nil)).Elem(), JobScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobScheduleResponseInput)(nil)).Elem(), JobScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobScheduleResponsePtrInput)(nil)).Elem(), JobScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepActionInput)(nil)).Elem(), JobStepActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepActionPtrInput)(nil)).Elem(), JobStepActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepActionResponseInput)(nil)).Elem(), JobStepActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepActionResponsePtrInput)(nil)).Elem(), JobStepActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepExecutionOptionsInput)(nil)).Elem(), JobStepExecutionOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepExecutionOptionsPtrInput)(nil)).Elem(), JobStepExecutionOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepExecutionOptionsResponseInput)(nil)).Elem(), JobStepExecutionOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepExecutionOptionsResponsePtrInput)(nil)).Elem(), JobStepExecutionOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepOutputTypeInput)(nil)).Elem(), JobStepOutputTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepOutputTypePtrInput)(nil)).Elem(), JobStepOutputTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepOutputResponseInput)(nil)).Elem(), JobStepOutputResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepOutputResponsePtrInput)(nil)).Elem(), JobStepOutputResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobTargetInput)(nil)).Elem(), JobTargetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobTargetArrayInput)(nil)).Elem(), JobTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobTargetResponseInput)(nil)).Elem(), JobTargetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobTargetResponseArrayInput)(nil)).Elem(), JobTargetResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstanceExternalAdministratorInput)(nil)).Elem(), ManagedInstanceExternalAdministratorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstanceExternalAdministratorPtrInput)(nil)).Elem(), ManagedInstanceExternalAdministratorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstanceExternalAdministratorResponseInput)(nil)).Elem(), ManagedInstanceExternalAdministratorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstanceExternalAdministratorResponsePtrInput)(nil)).Elem(), ManagedInstanceExternalAdministratorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePairInfoInput)(nil)).Elem(), ManagedInstancePairInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePairInfoArrayInput)(nil)).Elem(), ManagedInstancePairInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePairInfoResponseInput)(nil)).Elem(), ManagedInstancePairInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePairInfoResponseArrayInput)(nil)).Elem(), ManagedInstancePairInfoResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePecPropertyResponseInput)(nil)).Elem(), ManagedInstancePecPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePecPropertyResponseArrayInput)(nil)).Elem(), ManagedInstancePecPropertyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePrivateEndpointConnectionPropertiesResponseInput)(nil)).Elem(), ManagedInstancePrivateEndpointConnectionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePrivateEndpointPropertyInput)(nil)).Elem(), ManagedInstancePrivateEndpointPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePrivateEndpointPropertyPtrInput)(nil)).Elem(), ManagedInstancePrivateEndpointPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePrivateEndpointPropertyResponseInput)(nil)).Elem(), ManagedInstancePrivateEndpointPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePrivateEndpointPropertyResponsePtrInput)(nil)).Elem(), ManagedInstancePrivateEndpointPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePrivateLinkServiceConnectionStatePropertyInput)(nil)).Elem(), ManagedInstancePrivateLinkServiceConnectionStatePropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrInput)(nil)).Elem(), ManagedInstancePrivateLinkServiceConnectionStatePropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseInput)(nil)).Elem(), ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrInput)(nil)).Elem(), ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerInfoInput)(nil)).Elem(), PartnerInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerInfoArrayInput)(nil)).Elem(), PartnerInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerInfoResponseInput)(nil)).Elem(), PartnerInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerInfoResponseArrayInput)(nil)).Elem(), PartnerInfoResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerRegionInfoInput)(nil)).Elem(), PartnerRegionInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerRegionInfoArrayInput)(nil)).Elem(), PartnerRegionInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerRegionInfoResponseInput)(nil)).Elem(), PartnerRegionInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerRegionInfoResponseArrayInput)(nil)).Elem(), PartnerRegionInfoResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponseInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointPropertyInput)(nil)).Elem(), PrivateEndpointPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointPropertyPtrInput)(nil)).Elem(), PrivateEndpointPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointPropertyResponseInput)(nil)).Elem(), PrivateEndpointPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointPropertyResponsePtrInput)(nil)).Elem(), PrivateEndpointPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyInput)(nil)).Elem(), PrivateLinkServiceConnectionStatePropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyPtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStatePropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponseInput)(nil)).Elem(), PrivateLinkServiceConnectionStatePropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponsePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStatePropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecommendedActionErrorInfoResponseInput)(nil)).Elem(), RecommendedActionErrorInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecommendedActionImpactRecordResponseInput)(nil)).Elem(), RecommendedActionImpactRecordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecommendedActionImpactRecordResponseArrayInput)(nil)).Elem(), RecommendedActionImpactRecordResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecommendedActionImplementationInfoResponseInput)(nil)).Elem(), RecommendedActionImplementationInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecommendedActionMetricInfoResponseInput)(nil)).Elem(), RecommendedActionMetricInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecommendedActionMetricInfoResponseArrayInput)(nil)).Elem(), RecommendedActionMetricInfoResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecommendedActionResponseInput)(nil)).Elem(), RecommendedActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecommendedActionResponseArrayInput)(nil)).Elem(), RecommendedActionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecommendedActionStateInfoResponseInput)(nil)).Elem(), RecommendedActionStateInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityInput)(nil)).Elem(), ResourceIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityPtrInput)(nil)).Elem(), ResourceIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityResponseInput)(nil)).Elem(), ResourceIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityResponsePtrInput)(nil)).Elem(), ResourceIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerExternalAdministratorInput)(nil)).Elem(), ServerExternalAdministratorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerExternalAdministratorPtrInput)(nil)).Elem(), ServerExternalAdministratorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerExternalAdministratorResponseInput)(nil)).Elem(), ServerExternalAdministratorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerExternalAdministratorResponsePtrInput)(nil)).Elem(), ServerExternalAdministratorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerInfoInput)(nil)).Elem(), ServerInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerInfoArrayInput)(nil)).Elem(), ServerInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerInfoResponseInput)(nil)).Elem(), ServerInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerInfoResponseArrayInput)(nil)).Elem(), ServerInfoResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerPrivateEndpointConnectionResponseInput)(nil)).Elem(), ServerPrivateEndpointConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerPrivateEndpointConnectionResponseArrayInput)(nil)).Elem(), ServerPrivateEndpointConnectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaInput)(nil)).Elem(), SyncGroupSchemaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaPtrInput)(nil)).Elem(), SyncGroupSchemaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaResponseInput)(nil)).Elem(), SyncGroupSchemaResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaResponsePtrInput)(nil)).Elem(), SyncGroupSchemaResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaTableInput)(nil)).Elem(), SyncGroupSchemaTableArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaTableArrayInput)(nil)).Elem(), SyncGroupSchemaTableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaTableColumnInput)(nil)).Elem(), SyncGroupSchemaTableColumnArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaTableColumnArrayInput)(nil)).Elem(), SyncGroupSchemaTableColumnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaTableColumnResponseInput)(nil)).Elem(), SyncGroupSchemaTableColumnResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaTableColumnResponseArrayInput)(nil)).Elem(), SyncGroupSchemaTableColumnResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaTableResponseInput)(nil)).Elem(), SyncGroupSchemaTableResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncGroupSchemaTableResponseArrayInput)(nil)).Elem(), SyncGroupSchemaTableResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserIdentityResponseInput)(nil)).Elem(), UserIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserIdentityResponseMapInput)(nil)).Elem(), UserIdentityResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesInput)(nil)).Elem(), VulnerabilityAssessmentRecurringScansPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesPtrInput)(nil)).Elem(), VulnerabilityAssessmentRecurringScansPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesResponseInput)(nil)).Elem(), VulnerabilityAssessmentRecurringScansPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesResponsePtrInput)(nil)).Elem(), VulnerabilityAssessmentRecurringScansPropertiesResponseArgs{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemOutput{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput{})
	pulumi.RegisterOutputType(ElasticPoolPerDatabaseSettingsOutput{})
	pulumi.RegisterOutputType(ElasticPoolPerDatabaseSettingsPtrOutput{})
	pulumi.RegisterOutputType(ElasticPoolPerDatabaseSettingsResponseOutput{})
	pulumi.RegisterOutputType(ElasticPoolPerDatabaseSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadOnlyEndpointOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadOnlyEndpointPtrOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadOnlyEndpointResponseOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadOnlyEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadWriteEndpointOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadWriteEndpointPtrOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadWriteEndpointResponseOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadWriteEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointPtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointResponseOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointPtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointResponseOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(JobScheduleOutput{})
	pulumi.RegisterOutputType(JobSchedulePtrOutput{})
	pulumi.RegisterOutputType(JobScheduleResponseOutput{})
	pulumi.RegisterOutputType(JobScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStepActionOutput{})
	pulumi.RegisterOutputType(JobStepActionPtrOutput{})
	pulumi.RegisterOutputType(JobStepActionResponseOutput{})
	pulumi.RegisterOutputType(JobStepActionResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStepExecutionOptionsOutput{})
	pulumi.RegisterOutputType(JobStepExecutionOptionsPtrOutput{})
	pulumi.RegisterOutputType(JobStepExecutionOptionsResponseOutput{})
	pulumi.RegisterOutputType(JobStepExecutionOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStepOutputTypeOutput{})
	pulumi.RegisterOutputType(JobStepOutputTypePtrOutput{})
	pulumi.RegisterOutputType(JobStepOutputResponseOutput{})
	pulumi.RegisterOutputType(JobStepOutputResponsePtrOutput{})
	pulumi.RegisterOutputType(JobTargetOutput{})
	pulumi.RegisterOutputType(JobTargetArrayOutput{})
	pulumi.RegisterOutputType(JobTargetResponseOutput{})
	pulumi.RegisterOutputType(JobTargetResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedInstanceExternalAdministratorOutput{})
	pulumi.RegisterOutputType(ManagedInstanceExternalAdministratorPtrOutput{})
	pulumi.RegisterOutputType(ManagedInstanceExternalAdministratorResponseOutput{})
	pulumi.RegisterOutputType(ManagedInstanceExternalAdministratorResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedInstancePairInfoOutput{})
	pulumi.RegisterOutputType(ManagedInstancePairInfoArrayOutput{})
	pulumi.RegisterOutputType(ManagedInstancePairInfoResponseOutput{})
	pulumi.RegisterOutputType(ManagedInstancePairInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedInstancePecPropertyResponseOutput{})
	pulumi.RegisterOutputType(ManagedInstancePecPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedInstancePrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagedInstancePrivateEndpointPropertyOutput{})
	pulumi.RegisterOutputType(ManagedInstancePrivateEndpointPropertyPtrOutput{})
	pulumi.RegisterOutputType(ManagedInstancePrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(ManagedInstancePrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedInstancePrivateLinkServiceConnectionStatePropertyOutput{})
	pulumi.RegisterOutputType(ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrOutput{})
	pulumi.RegisterOutputType(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PartnerInfoOutput{})
	pulumi.RegisterOutputType(PartnerInfoArrayOutput{})
	pulumi.RegisterOutputType(PartnerInfoResponseOutput{})
	pulumi.RegisterOutputType(PartnerInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(PartnerRegionInfoOutput{})
	pulumi.RegisterOutputType(PartnerRegionInfoArrayOutput{})
	pulumi.RegisterOutputType(PartnerRegionInfoResponseOutput{})
	pulumi.RegisterOutputType(PartnerRegionInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(RecommendedActionErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(RecommendedActionImpactRecordResponseOutput{})
	pulumi.RegisterOutputType(RecommendedActionImpactRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(RecommendedActionImplementationInfoResponseOutput{})
	pulumi.RegisterOutputType(RecommendedActionMetricInfoResponseOutput{})
	pulumi.RegisterOutputType(RecommendedActionMetricInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(RecommendedActionResponseOutput{})
	pulumi.RegisterOutputType(RecommendedActionResponseArrayOutput{})
	pulumi.RegisterOutputType(RecommendedActionStateInfoResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerExternalAdministratorOutput{})
	pulumi.RegisterOutputType(ServerExternalAdministratorPtrOutput{})
	pulumi.RegisterOutputType(ServerExternalAdministratorResponseOutput{})
	pulumi.RegisterOutputType(ServerExternalAdministratorResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerInfoOutput{})
	pulumi.RegisterOutputType(ServerInfoArrayOutput{})
	pulumi.RegisterOutputType(ServerInfoResponseOutput{})
	pulumi.RegisterOutputType(ServerInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerPrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(ServerPrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaPtrOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaResponseOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableArrayOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnArrayOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnResponseOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableResponseOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput{})
}
