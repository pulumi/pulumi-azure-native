


package v20171001preview

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

func (o InstanceFailoverGroupReadWriteEndpointOutput) FailoverPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceFailoverGroupReadWriteEndpoint) string { return v.FailoverPolicy }).(pulumi.StringOutput)
}

func (o InstanceFailoverGroupReadWriteEndpointOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InstanceFailoverGroupReadWriteEndpoint) *int { return v.FailoverWithDataLossGracePeriodMinutes }).(pulumi.IntPtrOutput)
}

type InstanceFailoverGroupReadWriteEndpointResponse struct {
	FailoverPolicy                         string `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes *int   `pulumi:"failoverWithDataLossGracePeriodMinutes"`
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

func (o InstanceFailoverGroupReadWriteEndpointResponseOutput) FailoverPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceFailoverGroupReadWriteEndpointResponse) string { return v.FailoverPolicy }).(pulumi.StringOutput)
}

func (o InstanceFailoverGroupReadWriteEndpointResponseOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InstanceFailoverGroupReadWriteEndpointResponse) *int {
		return v.FailoverWithDataLossGracePeriodMinutes
	}).(pulumi.IntPtrOutput)
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

type VulnerabilityAssessmentRecurringScansProperties struct {
	EmailSubscriptionAdmins *bool    `pulumi:"emailSubscriptionAdmins"`
	Emails                  []string `pulumi:"emails"`
	IsEnabled               *bool    `pulumi:"isEnabled"`
}


func (val *VulnerabilityAssessmentRecurringScansProperties) Defaults() *VulnerabilityAssessmentRecurringScansProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EmailSubscriptionAdmins) {
		emailSubscriptionAdmins_ := true
		tmp.EmailSubscriptionAdmins = &emailSubscriptionAdmins_
	}
	return &tmp
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


func (val *VulnerabilityAssessmentRecurringScansPropertiesResponse) Defaults() *VulnerabilityAssessmentRecurringScansPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EmailSubscriptionAdmins) {
		emailSubscriptionAdmins_ := true
		tmp.EmailSubscriptionAdmins = &emailSubscriptionAdmins_
	}
	return &tmp
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
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemOutput{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput{})
	pulumi.RegisterOutputType(ElasticPoolPerDatabaseSettingsOutput{})
	pulumi.RegisterOutputType(ElasticPoolPerDatabaseSettingsPtrOutput{})
	pulumi.RegisterOutputType(ElasticPoolPerDatabaseSettingsResponseOutput{})
	pulumi.RegisterOutputType(ElasticPoolPerDatabaseSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointPtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointResponseOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointResponseOutput{})
	pulumi.RegisterOutputType(ManagedInstancePairInfoOutput{})
	pulumi.RegisterOutputType(ManagedInstancePairInfoArrayOutput{})
	pulumi.RegisterOutputType(ManagedInstancePairInfoResponseOutput{})
	pulumi.RegisterOutputType(ManagedInstancePairInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(PartnerRegionInfoOutput{})
	pulumi.RegisterOutputType(PartnerRegionInfoArrayOutput{})
	pulumi.RegisterOutputType(PartnerRegionInfoResponseOutput{})
	pulumi.RegisterOutputType(PartnerRegionInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput{})
}
