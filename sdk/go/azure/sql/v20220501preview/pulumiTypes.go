


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseIdentity struct {
	Type                   *string                `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type DatabaseIdentityInput interface {
	pulumi.Input

	ToDatabaseIdentityOutput() DatabaseIdentityOutput
	ToDatabaseIdentityOutputWithContext(context.Context) DatabaseIdentityOutput
}

type DatabaseIdentityArgs struct {
	Type                   pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput       `pulumi:"userAssignedIdentities"`
}

func (DatabaseIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseIdentity)(nil)).Elem()
}

func (i DatabaseIdentityArgs) ToDatabaseIdentityOutput() DatabaseIdentityOutput {
	return i.ToDatabaseIdentityOutputWithContext(context.Background())
}

func (i DatabaseIdentityArgs) ToDatabaseIdentityOutputWithContext(ctx context.Context) DatabaseIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIdentityOutput)
}

func (i DatabaseIdentityArgs) ToDatabaseIdentityPtrOutput() DatabaseIdentityPtrOutput {
	return i.ToDatabaseIdentityPtrOutputWithContext(context.Background())
}

func (i DatabaseIdentityArgs) ToDatabaseIdentityPtrOutputWithContext(ctx context.Context) DatabaseIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIdentityOutput).ToDatabaseIdentityPtrOutputWithContext(ctx)
}









type DatabaseIdentityPtrInput interface {
	pulumi.Input

	ToDatabaseIdentityPtrOutput() DatabaseIdentityPtrOutput
	ToDatabaseIdentityPtrOutputWithContext(context.Context) DatabaseIdentityPtrOutput
}

type databaseIdentityPtrType DatabaseIdentityArgs

func DatabaseIdentityPtr(v *DatabaseIdentityArgs) DatabaseIdentityPtrInput {
	return (*databaseIdentityPtrType)(v)
}

func (*databaseIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseIdentity)(nil)).Elem()
}

func (i *databaseIdentityPtrType) ToDatabaseIdentityPtrOutput() DatabaseIdentityPtrOutput {
	return i.ToDatabaseIdentityPtrOutputWithContext(context.Background())
}

func (i *databaseIdentityPtrType) ToDatabaseIdentityPtrOutputWithContext(ctx context.Context) DatabaseIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseIdentityPtrOutput)
}

type DatabaseIdentityOutput struct{ *pulumi.OutputState }

func (DatabaseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseIdentity)(nil)).Elem()
}

func (o DatabaseIdentityOutput) ToDatabaseIdentityOutput() DatabaseIdentityOutput {
	return o
}

func (o DatabaseIdentityOutput) ToDatabaseIdentityOutputWithContext(ctx context.Context) DatabaseIdentityOutput {
	return o
}

func (o DatabaseIdentityOutput) ToDatabaseIdentityPtrOutput() DatabaseIdentityPtrOutput {
	return o.ToDatabaseIdentityPtrOutputWithContext(context.Background())
}

func (o DatabaseIdentityOutput) ToDatabaseIdentityPtrOutputWithContext(ctx context.Context) DatabaseIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseIdentity) *DatabaseIdentity {
		return &v
	}).(DatabaseIdentityPtrOutput)
}

func (o DatabaseIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o DatabaseIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v DatabaseIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type DatabaseIdentityPtrOutput struct{ *pulumi.OutputState }

func (DatabaseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseIdentity)(nil)).Elem()
}

func (o DatabaseIdentityPtrOutput) ToDatabaseIdentityPtrOutput() DatabaseIdentityPtrOutput {
	return o
}

func (o DatabaseIdentityPtrOutput) ToDatabaseIdentityPtrOutputWithContext(ctx context.Context) DatabaseIdentityPtrOutput {
	return o
}

func (o DatabaseIdentityPtrOutput) Elem() DatabaseIdentityOutput {
	return o.ApplyT(func(v *DatabaseIdentity) DatabaseIdentity {
		if v != nil {
			return *v
		}
		var ret DatabaseIdentity
		return ret
	}).(DatabaseIdentityOutput)
}

func (o DatabaseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o DatabaseIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *DatabaseIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type DatabaseIdentityResponse struct {
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]DatabaseUserIdentityResponse `pulumi:"userAssignedIdentities"`
}

type DatabaseIdentityResponseOutput struct{ *pulumi.OutputState }

func (DatabaseIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseIdentityResponse)(nil)).Elem()
}

func (o DatabaseIdentityResponseOutput) ToDatabaseIdentityResponseOutput() DatabaseIdentityResponseOutput {
	return o
}

func (o DatabaseIdentityResponseOutput) ToDatabaseIdentityResponseOutputWithContext(ctx context.Context) DatabaseIdentityResponseOutput {
	return o
}

func (o DatabaseIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o DatabaseIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o DatabaseIdentityResponseOutput) UserAssignedIdentities() DatabaseUserIdentityResponseMapOutput {
	return o.ApplyT(func(v DatabaseIdentityResponse) map[string]DatabaseUserIdentityResponse {
		return v.UserAssignedIdentities
	}).(DatabaseUserIdentityResponseMapOutput)
}

type DatabaseIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (DatabaseIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseIdentityResponse)(nil)).Elem()
}

func (o DatabaseIdentityResponsePtrOutput) ToDatabaseIdentityResponsePtrOutput() DatabaseIdentityResponsePtrOutput {
	return o
}

func (o DatabaseIdentityResponsePtrOutput) ToDatabaseIdentityResponsePtrOutputWithContext(ctx context.Context) DatabaseIdentityResponsePtrOutput {
	return o
}

func (o DatabaseIdentityResponsePtrOutput) Elem() DatabaseIdentityResponseOutput {
	return o.ApplyT(func(v *DatabaseIdentityResponse) DatabaseIdentityResponse {
		if v != nil {
			return *v
		}
		var ret DatabaseIdentityResponse
		return ret
	}).(DatabaseIdentityResponseOutput)
}

func (o DatabaseIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o DatabaseIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o DatabaseIdentityResponsePtrOutput) UserAssignedIdentities() DatabaseUserIdentityResponseMapOutput {
	return o.ApplyT(func(v *DatabaseIdentityResponse) map[string]DatabaseUserIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(DatabaseUserIdentityResponseMapOutput)
}

type DatabaseUserIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type DatabaseUserIdentityResponseOutput struct{ *pulumi.OutputState }

func (DatabaseUserIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseUserIdentityResponse)(nil)).Elem()
}

func (o DatabaseUserIdentityResponseOutput) ToDatabaseUserIdentityResponseOutput() DatabaseUserIdentityResponseOutput {
	return o
}

func (o DatabaseUserIdentityResponseOutput) ToDatabaseUserIdentityResponseOutputWithContext(ctx context.Context) DatabaseUserIdentityResponseOutput {
	return o
}

func (o DatabaseUserIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseUserIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o DatabaseUserIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseUserIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type DatabaseUserIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (DatabaseUserIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatabaseUserIdentityResponse)(nil)).Elem()
}

func (o DatabaseUserIdentityResponseMapOutput) ToDatabaseUserIdentityResponseMapOutput() DatabaseUserIdentityResponseMapOutput {
	return o
}

func (o DatabaseUserIdentityResponseMapOutput) ToDatabaseUserIdentityResponseMapOutputWithContext(ctx context.Context) DatabaseUserIdentityResponseMapOutput {
	return o
}

func (o DatabaseUserIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) DatabaseUserIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatabaseUserIdentityResponse {
		return vs[0].(map[string]DatabaseUserIdentityResponse)[vs[1].(string)]
	}).(DatabaseUserIdentityResponseOutput)
}

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

func (o FailoverGroupReadWriteEndpointOutput) FailoverPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpoint) string { return v.FailoverPolicy }).(pulumi.StringOutput)
}

func (o FailoverGroupReadWriteEndpointOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpoint) *int { return v.FailoverWithDataLossGracePeriodMinutes }).(pulumi.IntPtrOutput)
}

type FailoverGroupReadWriteEndpointResponse struct {
	FailoverPolicy                         string `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes *int   `pulumi:"failoverWithDataLossGracePeriodMinutes"`
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

func (o FailoverGroupReadWriteEndpointResponseOutput) FailoverPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpointResponse) string { return v.FailoverPolicy }).(pulumi.StringOutput)
}

func (o FailoverGroupReadWriteEndpointResponseOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpointResponse) *int { return v.FailoverWithDataLossGracePeriodMinutes }).(pulumi.IntPtrOutput)
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

type JobSchedule struct {
	Enabled   *bool            `pulumi:"enabled"`
	EndTime   *string          `pulumi:"endTime"`
	Interval  *string          `pulumi:"interval"`
	StartTime *string          `pulumi:"startTime"`
	Type      *JobScheduleType `pulumi:"type"`
}


func (val *JobSchedule) Defaults() *JobSchedule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EndTime) {
		endTime_ := "9999-12-31T03:59:59-08:00"
		tmp.EndTime = &endTime_
	}
	if isZero(tmp.StartTime) {
		startTime_ := "0001-01-01T16:00:00-08:00"
		tmp.StartTime = &startTime_
	}
	if isZero(tmp.Type) {
		type_ := JobScheduleType("Once")
		tmp.Type = &type_
	}
	return &tmp
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


func (val *JobScheduleArgs) Defaults() *JobScheduleArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EndTime) {
		tmp.EndTime = pulumi.StringPtr("9999-12-31T03:59:59-08:00")
	}
	if isZero(tmp.StartTime) {
		tmp.StartTime = pulumi.StringPtr("0001-01-01T16:00:00-08:00")
	}
	if isZero(tmp.Type) {
		tmp.Type = JobScheduleType("Once")
	}
	return &tmp
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


func (val *JobScheduleResponse) Defaults() *JobScheduleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EndTime) {
		endTime_ := "9999-12-31T03:59:59-08:00"
		tmp.EndTime = &endTime_
	}
	if isZero(tmp.StartTime) {
		startTime_ := "0001-01-01T16:00:00-08:00"
		tmp.StartTime = &startTime_
	}
	if isZero(tmp.Type) {
		type_ := "Once"
		tmp.Type = &type_
	}
	return &tmp
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


func (val *JobStepAction) Defaults() *JobStepAction {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Source) {
		source_ := "Inline"
		tmp.Source = &source_
	}
	if isZero(tmp.Type) {
		type_ := "TSql"
		tmp.Type = &type_
	}
	return &tmp
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


func (val *JobStepActionArgs) Defaults() *JobStepActionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Source) {
		tmp.Source = pulumi.StringPtr("Inline")
	}
	if isZero(tmp.Type) {
		tmp.Type = pulumi.StringPtr("TSql")
	}
	return &tmp
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

func (o JobStepActionOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepAction) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o JobStepActionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepAction) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o JobStepActionOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepAction) string { return v.Value }).(pulumi.StringOutput)
}

type JobStepActionResponse struct {
	Source *string `pulumi:"source"`
	Type   *string `pulumi:"type"`
	Value  string  `pulumi:"value"`
}


func (val *JobStepActionResponse) Defaults() *JobStepActionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Source) {
		source_ := "Inline"
		tmp.Source = &source_
	}
	if isZero(tmp.Type) {
		type_ := "TSql"
		tmp.Type = &type_
	}
	return &tmp
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

func (o JobStepActionResponseOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepActionResponse) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o JobStepActionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepActionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o JobStepActionResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepActionResponse) string { return v.Value }).(pulumi.StringOutput)
}

type JobStepExecutionOptions struct {
	InitialRetryIntervalSeconds    *int     `pulumi:"initialRetryIntervalSeconds"`
	MaximumRetryIntervalSeconds    *int     `pulumi:"maximumRetryIntervalSeconds"`
	RetryAttempts                  *int     `pulumi:"retryAttempts"`
	RetryIntervalBackoffMultiplier *float64 `pulumi:"retryIntervalBackoffMultiplier"`
	TimeoutSeconds                 *int     `pulumi:"timeoutSeconds"`
}


func (val *JobStepExecutionOptions) Defaults() *JobStepExecutionOptions {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.InitialRetryIntervalSeconds) {
		initialRetryIntervalSeconds_ := 1
		tmp.InitialRetryIntervalSeconds = &initialRetryIntervalSeconds_
	}
	if isZero(tmp.MaximumRetryIntervalSeconds) {
		maximumRetryIntervalSeconds_ := 120
		tmp.MaximumRetryIntervalSeconds = &maximumRetryIntervalSeconds_
	}
	if isZero(tmp.RetryAttempts) {
		retryAttempts_ := 10
		tmp.RetryAttempts = &retryAttempts_
	}
	if isZero(tmp.RetryIntervalBackoffMultiplier) {
		retryIntervalBackoffMultiplier_ := 2.0
		tmp.RetryIntervalBackoffMultiplier = &retryIntervalBackoffMultiplier_
	}
	if isZero(tmp.TimeoutSeconds) {
		timeoutSeconds_ := 43200
		tmp.TimeoutSeconds = &timeoutSeconds_
	}
	return &tmp
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


func (val *JobStepExecutionOptionsArgs) Defaults() *JobStepExecutionOptionsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.InitialRetryIntervalSeconds) {
		tmp.InitialRetryIntervalSeconds = pulumi.IntPtr(1)
	}
	if isZero(tmp.MaximumRetryIntervalSeconds) {
		tmp.MaximumRetryIntervalSeconds = pulumi.IntPtr(120)
	}
	if isZero(tmp.RetryAttempts) {
		tmp.RetryAttempts = pulumi.IntPtr(10)
	}
	if isZero(tmp.RetryIntervalBackoffMultiplier) {
		tmp.RetryIntervalBackoffMultiplier = pulumi.Float64Ptr(2.0)
	}
	if isZero(tmp.TimeoutSeconds) {
		tmp.TimeoutSeconds = pulumi.IntPtr(43200)
	}
	return &tmp
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


func (val *JobStepExecutionOptionsResponse) Defaults() *JobStepExecutionOptionsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.InitialRetryIntervalSeconds) {
		initialRetryIntervalSeconds_ := 1
		tmp.InitialRetryIntervalSeconds = &initialRetryIntervalSeconds_
	}
	if isZero(tmp.MaximumRetryIntervalSeconds) {
		maximumRetryIntervalSeconds_ := 120
		tmp.MaximumRetryIntervalSeconds = &maximumRetryIntervalSeconds_
	}
	if isZero(tmp.RetryAttempts) {
		retryAttempts_ := 10
		tmp.RetryAttempts = &retryAttempts_
	}
	if isZero(tmp.RetryIntervalBackoffMultiplier) {
		retryIntervalBackoffMultiplier_ := 2.0
		tmp.RetryIntervalBackoffMultiplier = &retryIntervalBackoffMultiplier_
	}
	if isZero(tmp.TimeoutSeconds) {
		timeoutSeconds_ := 43200
		tmp.TimeoutSeconds = &timeoutSeconds_
	}
	return &tmp
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


func (val *JobStepOutputType) Defaults() *JobStepOutputType {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SchemaName) {
		schemaName_ := "dbo"
		tmp.SchemaName = &schemaName_
	}
	if isZero(tmp.Type) {
		type_ := "SqlDatabase"
		tmp.Type = &type_
	}
	return &tmp
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


func (val *JobStepOutputTypeArgs) Defaults() *JobStepOutputTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SchemaName) {
		tmp.SchemaName = pulumi.StringPtr("dbo")
	}
	if isZero(tmp.Type) {
		tmp.Type = pulumi.StringPtr("SqlDatabase")
	}
	return &tmp
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


func (val *JobStepOutputResponse) Defaults() *JobStepOutputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SchemaName) {
		schemaName_ := "dbo"
		tmp.SchemaName = &schemaName_
	}
	if isZero(tmp.Type) {
		type_ := "SqlDatabase"
		tmp.Type = &type_
	}
	return &tmp
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


func (val *JobTarget) Defaults() *JobTarget {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MembershipType) {
		membershipType_ := JobTargetGroupMembershipType("Include")
		tmp.MembershipType = &membershipType_
	}
	return &tmp
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


func (val *JobTargetArgs) Defaults() *JobTargetArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MembershipType) {
		tmp.MembershipType = JobTargetGroupMembershipType("Include")
	}
	return &tmp
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


func (val *JobTargetResponse) Defaults() *JobTargetResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MembershipType) {
		membershipType_ := "Include"
		tmp.MembershipType = &membershipType_
	}
	return &tmp
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
	GroupIds                          []string                                           `pulumi:"groupIds"`
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                             `pulumi:"provisioningState"`
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

func (o PrivateEndpointConnectionPropertiesResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
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

type ServicePrincipal struct {
	Type *string `pulumi:"type"`
}





type ServicePrincipalInput interface {
	pulumi.Input

	ToServicePrincipalOutput() ServicePrincipalOutput
	ToServicePrincipalOutputWithContext(context.Context) ServicePrincipalOutput
}

type ServicePrincipalArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ServicePrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipal)(nil)).Elem()
}

func (i ServicePrincipalArgs) ToServicePrincipalOutput() ServicePrincipalOutput {
	return i.ToServicePrincipalOutputWithContext(context.Background())
}

func (i ServicePrincipalArgs) ToServicePrincipalOutputWithContext(ctx context.Context) ServicePrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalOutput)
}

func (i ServicePrincipalArgs) ToServicePrincipalPtrOutput() ServicePrincipalPtrOutput {
	return i.ToServicePrincipalPtrOutputWithContext(context.Background())
}

func (i ServicePrincipalArgs) ToServicePrincipalPtrOutputWithContext(ctx context.Context) ServicePrincipalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalOutput).ToServicePrincipalPtrOutputWithContext(ctx)
}









type ServicePrincipalPtrInput interface {
	pulumi.Input

	ToServicePrincipalPtrOutput() ServicePrincipalPtrOutput
	ToServicePrincipalPtrOutputWithContext(context.Context) ServicePrincipalPtrOutput
}

type servicePrincipalPtrType ServicePrincipalArgs

func ServicePrincipalPtr(v *ServicePrincipalArgs) ServicePrincipalPtrInput {
	return (*servicePrincipalPtrType)(v)
}

func (*servicePrincipalPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipal)(nil)).Elem()
}

func (i *servicePrincipalPtrType) ToServicePrincipalPtrOutput() ServicePrincipalPtrOutput {
	return i.ToServicePrincipalPtrOutputWithContext(context.Background())
}

func (i *servicePrincipalPtrType) ToServicePrincipalPtrOutputWithContext(ctx context.Context) ServicePrincipalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPtrOutput)
}

type ServicePrincipalOutput struct{ *pulumi.OutputState }

func (ServicePrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipal)(nil)).Elem()
}

func (o ServicePrincipalOutput) ToServicePrincipalOutput() ServicePrincipalOutput {
	return o
}

func (o ServicePrincipalOutput) ToServicePrincipalOutputWithContext(ctx context.Context) ServicePrincipalOutput {
	return o
}

func (o ServicePrincipalOutput) ToServicePrincipalPtrOutput() ServicePrincipalPtrOutput {
	return o.ToServicePrincipalPtrOutputWithContext(context.Background())
}

func (o ServicePrincipalOutput) ToServicePrincipalPtrOutputWithContext(ctx context.Context) ServicePrincipalPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePrincipal) *ServicePrincipal {
		return &v
	}).(ServicePrincipalPtrOutput)
}

func (o ServicePrincipalOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipal) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ServicePrincipalPtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipal)(nil)).Elem()
}

func (o ServicePrincipalPtrOutput) ToServicePrincipalPtrOutput() ServicePrincipalPtrOutput {
	return o
}

func (o ServicePrincipalPtrOutput) ToServicePrincipalPtrOutputWithContext(ctx context.Context) ServicePrincipalPtrOutput {
	return o
}

func (o ServicePrincipalPtrOutput) Elem() ServicePrincipalOutput {
	return o.ApplyT(func(v *ServicePrincipal) ServicePrincipal {
		if v != nil {
			return *v
		}
		var ret ServicePrincipal
		return ret
	}).(ServicePrincipalOutput)
}

func (o ServicePrincipalPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipal) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalResponse struct {
	ClientId    string  `pulumi:"clientId"`
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ServicePrincipalResponseOutput struct{ *pulumi.OutputState }

func (ServicePrincipalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalResponse)(nil)).Elem()
}

func (o ServicePrincipalResponseOutput) ToServicePrincipalResponseOutput() ServicePrincipalResponseOutput {
	return o
}

func (o ServicePrincipalResponseOutput) ToServicePrincipalResponseOutputWithContext(ctx context.Context) ServicePrincipalResponseOutput {
	return o
}

func (o ServicePrincipalResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ServicePrincipalResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ServicePrincipalResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ServicePrincipalResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ServicePrincipalResponsePtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalResponse)(nil)).Elem()
}

func (o ServicePrincipalResponsePtrOutput) ToServicePrincipalResponsePtrOutput() ServicePrincipalResponsePtrOutput {
	return o
}

func (o ServicePrincipalResponsePtrOutput) ToServicePrincipalResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalResponsePtrOutput {
	return o
}

func (o ServicePrincipalResponsePtrOutput) Elem() ServicePrincipalResponseOutput {
	return o.ApplyT(func(v *ServicePrincipalResponse) ServicePrincipalResponse {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalResponse
		return ret
	}).(ServicePrincipalResponseOutput)
}

func (o ServicePrincipalResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
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

type UserIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
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


func (val *VulnerabilityAssessmentRecurringScansPropertiesArgs) Defaults() *VulnerabilityAssessmentRecurringScansPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EmailSubscriptionAdmins) {
		tmp.EmailSubscriptionAdmins = pulumi.BoolPtr(true)
	}
	return &tmp
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
	pulumi.RegisterOutputType(DatabaseIdentityOutput{})
	pulumi.RegisterOutputType(DatabaseIdentityPtrOutput{})
	pulumi.RegisterOutputType(DatabaseIdentityResponseOutput{})
	pulumi.RegisterOutputType(DatabaseIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(DatabaseUserIdentityResponseOutput{})
	pulumi.RegisterOutputType(DatabaseUserIdentityResponseMapOutput{})
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
	pulumi.RegisterOutputType(FailoverGroupReadWriteEndpointResponseOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointPtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointResponseOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointOutput{})
	pulumi.RegisterOutputType(InstanceFailoverGroupReadWriteEndpointResponseOutput{})
	pulumi.RegisterOutputType(JobScheduleOutput{})
	pulumi.RegisterOutputType(JobSchedulePtrOutput{})
	pulumi.RegisterOutputType(JobScheduleResponseOutput{})
	pulumi.RegisterOutputType(JobScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStepActionOutput{})
	pulumi.RegisterOutputType(JobStepActionResponseOutput{})
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
	pulumi.RegisterOutputType(ServicePrincipalOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalResponseOutput{})
	pulumi.RegisterOutputType(ServicePrincipalResponsePtrOutput{})
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
	pulumi.RegisterOutputType(UserIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput{})
}
