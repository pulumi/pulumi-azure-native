


package operationsmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ArmTemplateParameter struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type ArmTemplateParameterInput interface {
	pulumi.Input

	ToArmTemplateParameterOutput() ArmTemplateParameterOutput
	ToArmTemplateParameterOutputWithContext(context.Context) ArmTemplateParameterOutput
}

type ArmTemplateParameterArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ArmTemplateParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmTemplateParameter)(nil)).Elem()
}

func (i ArmTemplateParameterArgs) ToArmTemplateParameterOutput() ArmTemplateParameterOutput {
	return i.ToArmTemplateParameterOutputWithContext(context.Background())
}

func (i ArmTemplateParameterArgs) ToArmTemplateParameterOutputWithContext(ctx context.Context) ArmTemplateParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmTemplateParameterOutput)
}





type ArmTemplateParameterArrayInput interface {
	pulumi.Input

	ToArmTemplateParameterArrayOutput() ArmTemplateParameterArrayOutput
	ToArmTemplateParameterArrayOutputWithContext(context.Context) ArmTemplateParameterArrayOutput
}

type ArmTemplateParameterArray []ArmTemplateParameterInput

func (ArmTemplateParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArmTemplateParameter)(nil)).Elem()
}

func (i ArmTemplateParameterArray) ToArmTemplateParameterArrayOutput() ArmTemplateParameterArrayOutput {
	return i.ToArmTemplateParameterArrayOutputWithContext(context.Background())
}

func (i ArmTemplateParameterArray) ToArmTemplateParameterArrayOutputWithContext(ctx context.Context) ArmTemplateParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmTemplateParameterArrayOutput)
}

type ArmTemplateParameterOutput struct{ *pulumi.OutputState }

func (ArmTemplateParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmTemplateParameter)(nil)).Elem()
}

func (o ArmTemplateParameterOutput) ToArmTemplateParameterOutput() ArmTemplateParameterOutput {
	return o
}

func (o ArmTemplateParameterOutput) ToArmTemplateParameterOutputWithContext(ctx context.Context) ArmTemplateParameterOutput {
	return o
}

func (o ArmTemplateParameterOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmTemplateParameter) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArmTemplateParameterOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmTemplateParameter) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ArmTemplateParameterArrayOutput struct{ *pulumi.OutputState }

func (ArmTemplateParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArmTemplateParameter)(nil)).Elem()
}

func (o ArmTemplateParameterArrayOutput) ToArmTemplateParameterArrayOutput() ArmTemplateParameterArrayOutput {
	return o
}

func (o ArmTemplateParameterArrayOutput) ToArmTemplateParameterArrayOutputWithContext(ctx context.Context) ArmTemplateParameterArrayOutput {
	return o
}

func (o ArmTemplateParameterArrayOutput) Index(i pulumi.IntInput) ArmTemplateParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArmTemplateParameter {
		return vs[0].([]ArmTemplateParameter)[vs[1].(int)]
	}).(ArmTemplateParameterOutput)
}

type ArmTemplateParameterResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type ArmTemplateParameterResponseOutput struct{ *pulumi.OutputState }

func (ArmTemplateParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmTemplateParameterResponse)(nil)).Elem()
}

func (o ArmTemplateParameterResponseOutput) ToArmTemplateParameterResponseOutput() ArmTemplateParameterResponseOutput {
	return o
}

func (o ArmTemplateParameterResponseOutput) ToArmTemplateParameterResponseOutputWithContext(ctx context.Context) ArmTemplateParameterResponseOutput {
	return o
}

func (o ArmTemplateParameterResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmTemplateParameterResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArmTemplateParameterResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmTemplateParameterResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ArmTemplateParameterResponseArrayOutput struct{ *pulumi.OutputState }

func (ArmTemplateParameterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArmTemplateParameterResponse)(nil)).Elem()
}

func (o ArmTemplateParameterResponseArrayOutput) ToArmTemplateParameterResponseArrayOutput() ArmTemplateParameterResponseArrayOutput {
	return o
}

func (o ArmTemplateParameterResponseArrayOutput) ToArmTemplateParameterResponseArrayOutputWithContext(ctx context.Context) ArmTemplateParameterResponseArrayOutput {
	return o
}

func (o ArmTemplateParameterResponseArrayOutput) Index(i pulumi.IntInput) ArmTemplateParameterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArmTemplateParameterResponse {
		return vs[0].([]ArmTemplateParameterResponse)[vs[1].(int)]
	}).(ArmTemplateParameterResponseOutput)
}

type ManagementAssociationProperties struct {
	ApplicationId string `pulumi:"applicationId"`
}





type ManagementAssociationPropertiesInput interface {
	pulumi.Input

	ToManagementAssociationPropertiesOutput() ManagementAssociationPropertiesOutput
	ToManagementAssociationPropertiesOutputWithContext(context.Context) ManagementAssociationPropertiesOutput
}

type ManagementAssociationPropertiesArgs struct {
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
}

func (ManagementAssociationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementAssociationProperties)(nil)).Elem()
}

func (i ManagementAssociationPropertiesArgs) ToManagementAssociationPropertiesOutput() ManagementAssociationPropertiesOutput {
	return i.ToManagementAssociationPropertiesOutputWithContext(context.Background())
}

func (i ManagementAssociationPropertiesArgs) ToManagementAssociationPropertiesOutputWithContext(ctx context.Context) ManagementAssociationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementAssociationPropertiesOutput)
}

func (i ManagementAssociationPropertiesArgs) ToManagementAssociationPropertiesPtrOutput() ManagementAssociationPropertiesPtrOutput {
	return i.ToManagementAssociationPropertiesPtrOutputWithContext(context.Background())
}

func (i ManagementAssociationPropertiesArgs) ToManagementAssociationPropertiesPtrOutputWithContext(ctx context.Context) ManagementAssociationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementAssociationPropertiesOutput).ToManagementAssociationPropertiesPtrOutputWithContext(ctx)
}









type ManagementAssociationPropertiesPtrInput interface {
	pulumi.Input

	ToManagementAssociationPropertiesPtrOutput() ManagementAssociationPropertiesPtrOutput
	ToManagementAssociationPropertiesPtrOutputWithContext(context.Context) ManagementAssociationPropertiesPtrOutput
}

type managementAssociationPropertiesPtrType ManagementAssociationPropertiesArgs

func ManagementAssociationPropertiesPtr(v *ManagementAssociationPropertiesArgs) ManagementAssociationPropertiesPtrInput {
	return (*managementAssociationPropertiesPtrType)(v)
}

func (*managementAssociationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementAssociationProperties)(nil)).Elem()
}

func (i *managementAssociationPropertiesPtrType) ToManagementAssociationPropertiesPtrOutput() ManagementAssociationPropertiesPtrOutput {
	return i.ToManagementAssociationPropertiesPtrOutputWithContext(context.Background())
}

func (i *managementAssociationPropertiesPtrType) ToManagementAssociationPropertiesPtrOutputWithContext(ctx context.Context) ManagementAssociationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementAssociationPropertiesPtrOutput)
}

type ManagementAssociationPropertiesOutput struct{ *pulumi.OutputState }

func (ManagementAssociationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementAssociationProperties)(nil)).Elem()
}

func (o ManagementAssociationPropertiesOutput) ToManagementAssociationPropertiesOutput() ManagementAssociationPropertiesOutput {
	return o
}

func (o ManagementAssociationPropertiesOutput) ToManagementAssociationPropertiesOutputWithContext(ctx context.Context) ManagementAssociationPropertiesOutput {
	return o
}

func (o ManagementAssociationPropertiesOutput) ToManagementAssociationPropertiesPtrOutput() ManagementAssociationPropertiesPtrOutput {
	return o.ToManagementAssociationPropertiesPtrOutputWithContext(context.Background())
}

func (o ManagementAssociationPropertiesOutput) ToManagementAssociationPropertiesPtrOutputWithContext(ctx context.Context) ManagementAssociationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementAssociationProperties) *ManagementAssociationProperties {
		return &v
	}).(ManagementAssociationPropertiesPtrOutput)
}

func (o ManagementAssociationPropertiesOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementAssociationProperties) string { return v.ApplicationId }).(pulumi.StringOutput)
}

type ManagementAssociationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ManagementAssociationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementAssociationProperties)(nil)).Elem()
}

func (o ManagementAssociationPropertiesPtrOutput) ToManagementAssociationPropertiesPtrOutput() ManagementAssociationPropertiesPtrOutput {
	return o
}

func (o ManagementAssociationPropertiesPtrOutput) ToManagementAssociationPropertiesPtrOutputWithContext(ctx context.Context) ManagementAssociationPropertiesPtrOutput {
	return o
}

func (o ManagementAssociationPropertiesPtrOutput) Elem() ManagementAssociationPropertiesOutput {
	return o.ApplyT(func(v *ManagementAssociationProperties) ManagementAssociationProperties {
		if v != nil {
			return *v
		}
		var ret ManagementAssociationProperties
		return ret
	}).(ManagementAssociationPropertiesOutput)
}

func (o ManagementAssociationPropertiesPtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementAssociationProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

type ManagementAssociationPropertiesResponse struct {
	ApplicationId string `pulumi:"applicationId"`
}

type ManagementAssociationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManagementAssociationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementAssociationPropertiesResponse)(nil)).Elem()
}

func (o ManagementAssociationPropertiesResponseOutput) ToManagementAssociationPropertiesResponseOutput() ManagementAssociationPropertiesResponseOutput {
	return o
}

func (o ManagementAssociationPropertiesResponseOutput) ToManagementAssociationPropertiesResponseOutputWithContext(ctx context.Context) ManagementAssociationPropertiesResponseOutput {
	return o
}

func (o ManagementAssociationPropertiesResponseOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementAssociationPropertiesResponse) string { return v.ApplicationId }).(pulumi.StringOutput)
}

type ManagementConfigurationProperties struct {
	ApplicationId      *string                `pulumi:"applicationId"`
	Parameters         []ArmTemplateParameter `pulumi:"parameters"`
	ParentResourceType string                 `pulumi:"parentResourceType"`
	Template           interface{}            `pulumi:"template"`
}





type ManagementConfigurationPropertiesInput interface {
	pulumi.Input

	ToManagementConfigurationPropertiesOutput() ManagementConfigurationPropertiesOutput
	ToManagementConfigurationPropertiesOutputWithContext(context.Context) ManagementConfigurationPropertiesOutput
}

type ManagementConfigurationPropertiesArgs struct {
	ApplicationId      pulumi.StringPtrInput          `pulumi:"applicationId"`
	Parameters         ArmTemplateParameterArrayInput `pulumi:"parameters"`
	ParentResourceType pulumi.StringInput             `pulumi:"parentResourceType"`
	Template           pulumi.Input                   `pulumi:"template"`
}

func (ManagementConfigurationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementConfigurationProperties)(nil)).Elem()
}

func (i ManagementConfigurationPropertiesArgs) ToManagementConfigurationPropertiesOutput() ManagementConfigurationPropertiesOutput {
	return i.ToManagementConfigurationPropertiesOutputWithContext(context.Background())
}

func (i ManagementConfigurationPropertiesArgs) ToManagementConfigurationPropertiesOutputWithContext(ctx context.Context) ManagementConfigurationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementConfigurationPropertiesOutput)
}

func (i ManagementConfigurationPropertiesArgs) ToManagementConfigurationPropertiesPtrOutput() ManagementConfigurationPropertiesPtrOutput {
	return i.ToManagementConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (i ManagementConfigurationPropertiesArgs) ToManagementConfigurationPropertiesPtrOutputWithContext(ctx context.Context) ManagementConfigurationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementConfigurationPropertiesOutput).ToManagementConfigurationPropertiesPtrOutputWithContext(ctx)
}









type ManagementConfigurationPropertiesPtrInput interface {
	pulumi.Input

	ToManagementConfigurationPropertiesPtrOutput() ManagementConfigurationPropertiesPtrOutput
	ToManagementConfigurationPropertiesPtrOutputWithContext(context.Context) ManagementConfigurationPropertiesPtrOutput
}

type managementConfigurationPropertiesPtrType ManagementConfigurationPropertiesArgs

func ManagementConfigurationPropertiesPtr(v *ManagementConfigurationPropertiesArgs) ManagementConfigurationPropertiesPtrInput {
	return (*managementConfigurationPropertiesPtrType)(v)
}

func (*managementConfigurationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementConfigurationProperties)(nil)).Elem()
}

func (i *managementConfigurationPropertiesPtrType) ToManagementConfigurationPropertiesPtrOutput() ManagementConfigurationPropertiesPtrOutput {
	return i.ToManagementConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (i *managementConfigurationPropertiesPtrType) ToManagementConfigurationPropertiesPtrOutputWithContext(ctx context.Context) ManagementConfigurationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementConfigurationPropertiesPtrOutput)
}

type ManagementConfigurationPropertiesOutput struct{ *pulumi.OutputState }

func (ManagementConfigurationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementConfigurationProperties)(nil)).Elem()
}

func (o ManagementConfigurationPropertiesOutput) ToManagementConfigurationPropertiesOutput() ManagementConfigurationPropertiesOutput {
	return o
}

func (o ManagementConfigurationPropertiesOutput) ToManagementConfigurationPropertiesOutputWithContext(ctx context.Context) ManagementConfigurationPropertiesOutput {
	return o
}

func (o ManagementConfigurationPropertiesOutput) ToManagementConfigurationPropertiesPtrOutput() ManagementConfigurationPropertiesPtrOutput {
	return o.ToManagementConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (o ManagementConfigurationPropertiesOutput) ToManagementConfigurationPropertiesPtrOutputWithContext(ctx context.Context) ManagementConfigurationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementConfigurationProperties) *ManagementConfigurationProperties {
		return &v
	}).(ManagementConfigurationPropertiesPtrOutput)
}

func (o ManagementConfigurationPropertiesOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementConfigurationProperties) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o ManagementConfigurationPropertiesOutput) Parameters() ArmTemplateParameterArrayOutput {
	return o.ApplyT(func(v ManagementConfigurationProperties) []ArmTemplateParameter { return v.Parameters }).(ArmTemplateParameterArrayOutput)
}

func (o ManagementConfigurationPropertiesOutput) ParentResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementConfigurationProperties) string { return v.ParentResourceType }).(pulumi.StringOutput)
}

func (o ManagementConfigurationPropertiesOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v ManagementConfigurationProperties) interface{} { return v.Template }).(pulumi.AnyOutput)
}

type ManagementConfigurationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ManagementConfigurationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementConfigurationProperties)(nil)).Elem()
}

func (o ManagementConfigurationPropertiesPtrOutput) ToManagementConfigurationPropertiesPtrOutput() ManagementConfigurationPropertiesPtrOutput {
	return o
}

func (o ManagementConfigurationPropertiesPtrOutput) ToManagementConfigurationPropertiesPtrOutputWithContext(ctx context.Context) ManagementConfigurationPropertiesPtrOutput {
	return o
}

func (o ManagementConfigurationPropertiesPtrOutput) Elem() ManagementConfigurationPropertiesOutput {
	return o.ApplyT(func(v *ManagementConfigurationProperties) ManagementConfigurationProperties {
		if v != nil {
			return *v
		}
		var ret ManagementConfigurationProperties
		return ret
	}).(ManagementConfigurationPropertiesOutput)
}

func (o ManagementConfigurationPropertiesPtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o ManagementConfigurationPropertiesPtrOutput) Parameters() ArmTemplateParameterArrayOutput {
	return o.ApplyT(func(v *ManagementConfigurationProperties) []ArmTemplateParameter {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(ArmTemplateParameterArrayOutput)
}

func (o ManagementConfigurationPropertiesPtrOutput) ParentResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ParentResourceType
	}).(pulumi.StringPtrOutput)
}

func (o ManagementConfigurationPropertiesPtrOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v *ManagementConfigurationProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Template
	}).(pulumi.AnyOutput)
}

type ManagementConfigurationPropertiesResponse struct {
	ApplicationId      *string                        `pulumi:"applicationId"`
	Parameters         []ArmTemplateParameterResponse `pulumi:"parameters"`
	ParentResourceType string                         `pulumi:"parentResourceType"`
	ProvisioningState  string                         `pulumi:"provisioningState"`
	Template           interface{}                    `pulumi:"template"`
}

type ManagementConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManagementConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementConfigurationPropertiesResponse)(nil)).Elem()
}

func (o ManagementConfigurationPropertiesResponseOutput) ToManagementConfigurationPropertiesResponseOutput() ManagementConfigurationPropertiesResponseOutput {
	return o
}

func (o ManagementConfigurationPropertiesResponseOutput) ToManagementConfigurationPropertiesResponseOutputWithContext(ctx context.Context) ManagementConfigurationPropertiesResponseOutput {
	return o
}

func (o ManagementConfigurationPropertiesResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementConfigurationPropertiesResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o ManagementConfigurationPropertiesResponseOutput) Parameters() ArmTemplateParameterResponseArrayOutput {
	return o.ApplyT(func(v ManagementConfigurationPropertiesResponse) []ArmTemplateParameterResponse { return v.Parameters }).(ArmTemplateParameterResponseArrayOutput)
}

func (o ManagementConfigurationPropertiesResponseOutput) ParentResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementConfigurationPropertiesResponse) string { return v.ParentResourceType }).(pulumi.StringOutput)
}

func (o ManagementConfigurationPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementConfigurationPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagementConfigurationPropertiesResponseOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v ManagementConfigurationPropertiesResponse) interface{} { return v.Template }).(pulumi.AnyOutput)
}

type SolutionPlan struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
}





type SolutionPlanInput interface {
	pulumi.Input

	ToSolutionPlanOutput() SolutionPlanOutput
	ToSolutionPlanOutputWithContext(context.Context) SolutionPlanOutput
}

type SolutionPlanArgs struct {
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Product       pulumi.StringPtrInput `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringPtrInput `pulumi:"publisher"`
}

func (SolutionPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionPlan)(nil)).Elem()
}

func (i SolutionPlanArgs) ToSolutionPlanOutput() SolutionPlanOutput {
	return i.ToSolutionPlanOutputWithContext(context.Background())
}

func (i SolutionPlanArgs) ToSolutionPlanOutputWithContext(ctx context.Context) SolutionPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPlanOutput)
}

func (i SolutionPlanArgs) ToSolutionPlanPtrOutput() SolutionPlanPtrOutput {
	return i.ToSolutionPlanPtrOutputWithContext(context.Background())
}

func (i SolutionPlanArgs) ToSolutionPlanPtrOutputWithContext(ctx context.Context) SolutionPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPlanOutput).ToSolutionPlanPtrOutputWithContext(ctx)
}









type SolutionPlanPtrInput interface {
	pulumi.Input

	ToSolutionPlanPtrOutput() SolutionPlanPtrOutput
	ToSolutionPlanPtrOutputWithContext(context.Context) SolutionPlanPtrOutput
}

type solutionPlanPtrType SolutionPlanArgs

func SolutionPlanPtr(v *SolutionPlanArgs) SolutionPlanPtrInput {
	return (*solutionPlanPtrType)(v)
}

func (*solutionPlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionPlan)(nil)).Elem()
}

func (i *solutionPlanPtrType) ToSolutionPlanPtrOutput() SolutionPlanPtrOutput {
	return i.ToSolutionPlanPtrOutputWithContext(context.Background())
}

func (i *solutionPlanPtrType) ToSolutionPlanPtrOutputWithContext(ctx context.Context) SolutionPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPlanPtrOutput)
}

type SolutionPlanOutput struct{ *pulumi.OutputState }

func (SolutionPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionPlan)(nil)).Elem()
}

func (o SolutionPlanOutput) ToSolutionPlanOutput() SolutionPlanOutput {
	return o
}

func (o SolutionPlanOutput) ToSolutionPlanOutputWithContext(ctx context.Context) SolutionPlanOutput {
	return o
}

func (o SolutionPlanOutput) ToSolutionPlanPtrOutput() SolutionPlanPtrOutput {
	return o.ToSolutionPlanPtrOutputWithContext(context.Background())
}

func (o SolutionPlanOutput) ToSolutionPlanPtrOutputWithContext(ctx context.Context) SolutionPlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SolutionPlan) *SolutionPlan {
		return &v
	}).(SolutionPlanPtrOutput)
}

func (o SolutionPlanOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPlan) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SolutionPlanOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPlan) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o SolutionPlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPlan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o SolutionPlanOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPlan) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type SolutionPlanPtrOutput struct{ *pulumi.OutputState }

func (SolutionPlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionPlan)(nil)).Elem()
}

func (o SolutionPlanPtrOutput) ToSolutionPlanPtrOutput() SolutionPlanPtrOutput {
	return o
}

func (o SolutionPlanPtrOutput) ToSolutionPlanPtrOutputWithContext(ctx context.Context) SolutionPlanPtrOutput {
	return o
}

func (o SolutionPlanPtrOutput) Elem() SolutionPlanOutput {
	return o.ApplyT(func(v *SolutionPlan) SolutionPlan {
		if v != nil {
			return *v
		}
		var ret SolutionPlan
		return ret
	}).(SolutionPlanOutput)
}

func (o SolutionPlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPlan) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPlan) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPlan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPlan) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type SolutionPlanResponse struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
}

type SolutionPlanResponseOutput struct{ *pulumi.OutputState }

func (SolutionPlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionPlanResponse)(nil)).Elem()
}

func (o SolutionPlanResponseOutput) ToSolutionPlanResponseOutput() SolutionPlanResponseOutput {
	return o
}

func (o SolutionPlanResponseOutput) ToSolutionPlanResponseOutputWithContext(ctx context.Context) SolutionPlanResponseOutput {
	return o
}

func (o SolutionPlanResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPlanResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SolutionPlanResponseOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPlanResponse) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o SolutionPlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o SolutionPlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type SolutionPlanResponsePtrOutput struct{ *pulumi.OutputState }

func (SolutionPlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionPlanResponse)(nil)).Elem()
}

func (o SolutionPlanResponsePtrOutput) ToSolutionPlanResponsePtrOutput() SolutionPlanResponsePtrOutput {
	return o
}

func (o SolutionPlanResponsePtrOutput) ToSolutionPlanResponsePtrOutputWithContext(ctx context.Context) SolutionPlanResponsePtrOutput {
	return o
}

func (o SolutionPlanResponsePtrOutput) Elem() SolutionPlanResponseOutput {
	return o.ApplyT(func(v *SolutionPlanResponse) SolutionPlanResponse {
		if v != nil {
			return *v
		}
		var ret SolutionPlanResponse
		return ret
	}).(SolutionPlanResponseOutput)
}

func (o SolutionPlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type SolutionProperties struct {
	ContainedResources  []string `pulumi:"containedResources"`
	ReferencedResources []string `pulumi:"referencedResources"`
	WorkspaceResourceId string   `pulumi:"workspaceResourceId"`
}





type SolutionPropertiesInput interface {
	pulumi.Input

	ToSolutionPropertiesOutput() SolutionPropertiesOutput
	ToSolutionPropertiesOutputWithContext(context.Context) SolutionPropertiesOutput
}

type SolutionPropertiesArgs struct {
	ContainedResources  pulumi.StringArrayInput `pulumi:"containedResources"`
	ReferencedResources pulumi.StringArrayInput `pulumi:"referencedResources"`
	WorkspaceResourceId pulumi.StringInput      `pulumi:"workspaceResourceId"`
}

func (SolutionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionProperties)(nil)).Elem()
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesOutput() SolutionPropertiesOutput {
	return i.ToSolutionPropertiesOutputWithContext(context.Background())
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesOutputWithContext(ctx context.Context) SolutionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesOutput)
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return i.ToSolutionPropertiesPtrOutputWithContext(context.Background())
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesOutput).ToSolutionPropertiesPtrOutputWithContext(ctx)
}









type SolutionPropertiesPtrInput interface {
	pulumi.Input

	ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput
	ToSolutionPropertiesPtrOutputWithContext(context.Context) SolutionPropertiesPtrOutput
}

type solutionPropertiesPtrType SolutionPropertiesArgs

func SolutionPropertiesPtr(v *SolutionPropertiesArgs) SolutionPropertiesPtrInput {
	return (*solutionPropertiesPtrType)(v)
}

func (*solutionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionProperties)(nil)).Elem()
}

func (i *solutionPropertiesPtrType) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return i.ToSolutionPropertiesPtrOutputWithContext(context.Background())
}

func (i *solutionPropertiesPtrType) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesPtrOutput)
}

type SolutionPropertiesOutput struct{ *pulumi.OutputState }

func (SolutionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionProperties)(nil)).Elem()
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesOutput() SolutionPropertiesOutput {
	return o
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesOutputWithContext(ctx context.Context) SolutionPropertiesOutput {
	return o
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return o.ToSolutionPropertiesPtrOutputWithContext(context.Background())
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SolutionProperties) *SolutionProperties {
		return &v
	}).(SolutionPropertiesPtrOutput)
}

func (o SolutionPropertiesOutput) ContainedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SolutionProperties) []string { return v.ContainedResources }).(pulumi.StringArrayOutput)
}

func (o SolutionPropertiesOutput) ReferencedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SolutionProperties) []string { return v.ReferencedResources }).(pulumi.StringArrayOutput)
}

func (o SolutionPropertiesOutput) WorkspaceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SolutionProperties) string { return v.WorkspaceResourceId }).(pulumi.StringOutput)
}

type SolutionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SolutionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionProperties)(nil)).Elem()
}

func (o SolutionPropertiesPtrOutput) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return o
}

func (o SolutionPropertiesPtrOutput) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return o
}

func (o SolutionPropertiesPtrOutput) Elem() SolutionPropertiesOutput {
	return o.ApplyT(func(v *SolutionProperties) SolutionProperties {
		if v != nil {
			return *v
		}
		var ret SolutionProperties
		return ret
	}).(SolutionPropertiesOutput)
}

func (o SolutionPropertiesPtrOutput) ContainedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SolutionProperties) []string {
		if v == nil {
			return nil
		}
		return v.ContainedResources
	}).(pulumi.StringArrayOutput)
}

func (o SolutionPropertiesPtrOutput) ReferencedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SolutionProperties) []string {
		if v == nil {
			return nil
		}
		return v.ReferencedResources
	}).(pulumi.StringArrayOutput)
}

func (o SolutionPropertiesPtrOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.WorkspaceResourceId
	}).(pulumi.StringPtrOutput)
}

type SolutionPropertiesResponse struct {
	ContainedResources  []string `pulumi:"containedResources"`
	ProvisioningState   string   `pulumi:"provisioningState"`
	ReferencedResources []string `pulumi:"referencedResources"`
	WorkspaceResourceId string   `pulumi:"workspaceResourceId"`
}

type SolutionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SolutionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionPropertiesResponse)(nil)).Elem()
}

func (o SolutionPropertiesResponseOutput) ToSolutionPropertiesResponseOutput() SolutionPropertiesResponseOutput {
	return o
}

func (o SolutionPropertiesResponseOutput) ToSolutionPropertiesResponseOutputWithContext(ctx context.Context) SolutionPropertiesResponseOutput {
	return o
}

func (o SolutionPropertiesResponseOutput) ContainedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) []string { return v.ContainedResources }).(pulumi.StringArrayOutput)
}

func (o SolutionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SolutionPropertiesResponseOutput) ReferencedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) []string { return v.ReferencedResources }).(pulumi.StringArrayOutput)
}

func (o SolutionPropertiesResponseOutput) WorkspaceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) string { return v.WorkspaceResourceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ArmTemplateParameterOutput{})
	pulumi.RegisterOutputType(ArmTemplateParameterArrayOutput{})
	pulumi.RegisterOutputType(ArmTemplateParameterResponseOutput{})
	pulumi.RegisterOutputType(ArmTemplateParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementAssociationPropertiesOutput{})
	pulumi.RegisterOutputType(ManagementAssociationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagementAssociationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagementConfigurationPropertiesOutput{})
	pulumi.RegisterOutputType(ManagementConfigurationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagementConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SolutionPlanOutput{})
	pulumi.RegisterOutputType(SolutionPlanPtrOutput{})
	pulumi.RegisterOutputType(SolutionPlanResponseOutput{})
	pulumi.RegisterOutputType(SolutionPlanResponsePtrOutput{})
	pulumi.RegisterOutputType(SolutionPropertiesOutput{})
	pulumi.RegisterOutputType(SolutionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SolutionPropertiesResponseOutput{})
}
