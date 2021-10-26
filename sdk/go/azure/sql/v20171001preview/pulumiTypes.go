


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
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpointInput)(nil)).Elem(), InstanceFailoverGroupReadOnlyEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpointPtrInput)(nil)).Elem(), InstanceFailoverGroupReadOnlyEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpointResponseInput)(nil)).Elem(), InstanceFailoverGroupReadOnlyEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadOnlyEndpointResponsePtrInput)(nil)).Elem(), InstanceFailoverGroupReadOnlyEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpointInput)(nil)).Elem(), InstanceFailoverGroupReadWriteEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpointPtrInput)(nil)).Elem(), InstanceFailoverGroupReadWriteEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpointResponseInput)(nil)).Elem(), InstanceFailoverGroupReadWriteEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceFailoverGroupReadWriteEndpointResponsePtrInput)(nil)).Elem(), InstanceFailoverGroupReadWriteEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePairInfoInput)(nil)).Elem(), ManagedInstancePairInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePairInfoArrayInput)(nil)).Elem(), ManagedInstancePairInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePairInfoResponseInput)(nil)).Elem(), ManagedInstancePairInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedInstancePairInfoResponseArrayInput)(nil)).Elem(), ManagedInstancePairInfoResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerRegionInfoInput)(nil)).Elem(), PartnerRegionInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerRegionInfoArrayInput)(nil)).Elem(), PartnerRegionInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerRegionInfoResponseInput)(nil)).Elem(), PartnerRegionInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerRegionInfoResponseArrayInput)(nil)).Elem(), PartnerRegionInfoResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
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
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointPtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointResponseOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointPtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointResponseOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointResponsePtrOutput{})
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
