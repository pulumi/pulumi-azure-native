


package v20190301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AliasPathTypeResponse struct {
	ApiVersions []string `pulumi:"apiVersions"`
	Path        *string  `pulumi:"path"`
}

type AliasPathTypeResponseOutput struct{ *pulumi.OutputState }

func (AliasPathTypeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasPathTypeResponse)(nil)).Elem()
}

func (o AliasPathTypeResponseOutput) ToAliasPathTypeResponseOutput() AliasPathTypeResponseOutput {
	return o
}

func (o AliasPathTypeResponseOutput) ToAliasPathTypeResponseOutputWithContext(ctx context.Context) AliasPathTypeResponseOutput {
	return o
}

func (o AliasPathTypeResponseOutput) ApiVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AliasPathTypeResponse) []string { return v.ApiVersions }).(pulumi.StringArrayOutput)
}

func (o AliasPathTypeResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AliasPathTypeResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type AliasPathTypeResponseArrayOutput struct{ *pulumi.OutputState }

func (AliasPathTypeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AliasPathTypeResponse)(nil)).Elem()
}

func (o AliasPathTypeResponseArrayOutput) ToAliasPathTypeResponseArrayOutput() AliasPathTypeResponseArrayOutput {
	return o
}

func (o AliasPathTypeResponseArrayOutput) ToAliasPathTypeResponseArrayOutputWithContext(ctx context.Context) AliasPathTypeResponseArrayOutput {
	return o
}

func (o AliasPathTypeResponseArrayOutput) Index(i pulumi.IntInput) AliasPathTypeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AliasPathTypeResponse {
		return vs[0].([]AliasPathTypeResponse)[vs[1].(int)]
	}).(AliasPathTypeResponseOutput)
}

type AliasTypeResponse struct {
	Name  *string                 `pulumi:"name"`
	Paths []AliasPathTypeResponse `pulumi:"paths"`
}

type AliasTypeResponseOutput struct{ *pulumi.OutputState }

func (AliasTypeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasTypeResponse)(nil)).Elem()
}

func (o AliasTypeResponseOutput) ToAliasTypeResponseOutput() AliasTypeResponseOutput {
	return o
}

func (o AliasTypeResponseOutput) ToAliasTypeResponseOutputWithContext(ctx context.Context) AliasTypeResponseOutput {
	return o
}

func (o AliasTypeResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AliasTypeResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AliasTypeResponseOutput) Paths() AliasPathTypeResponseArrayOutput {
	return o.ApplyT(func(v AliasTypeResponse) []AliasPathTypeResponse { return v.Paths }).(AliasPathTypeResponseArrayOutput)
}

type AliasTypeResponseArrayOutput struct{ *pulumi.OutputState }

func (AliasTypeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AliasTypeResponse)(nil)).Elem()
}

func (o AliasTypeResponseArrayOutput) ToAliasTypeResponseArrayOutput() AliasTypeResponseArrayOutput {
	return o
}

func (o AliasTypeResponseArrayOutput) ToAliasTypeResponseArrayOutputWithContext(ctx context.Context) AliasTypeResponseArrayOutput {
	return o
}

func (o AliasTypeResponseArrayOutput) Index(i pulumi.IntInput) AliasTypeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AliasTypeResponse {
		return vs[0].([]AliasTypeResponse)[vs[1].(int)]
	}).(AliasTypeResponseOutput)
}

type BasicDependencyResponse struct {
	Id           *string `pulumi:"id"`
	ResourceName *string `pulumi:"resourceName"`
	ResourceType *string `pulumi:"resourceType"`
}

type BasicDependencyResponseOutput struct{ *pulumi.OutputState }

func (BasicDependencyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BasicDependencyResponse)(nil)).Elem()
}

func (o BasicDependencyResponseOutput) ToBasicDependencyResponseOutput() BasicDependencyResponseOutput {
	return o
}

func (o BasicDependencyResponseOutput) ToBasicDependencyResponseOutputWithContext(ctx context.Context) BasicDependencyResponseOutput {
	return o
}

func (o BasicDependencyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicDependencyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BasicDependencyResponseOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicDependencyResponse) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o BasicDependencyResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicDependencyResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

type BasicDependencyResponseArrayOutput struct{ *pulumi.OutputState }

func (BasicDependencyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BasicDependencyResponse)(nil)).Elem()
}

func (o BasicDependencyResponseArrayOutput) ToBasicDependencyResponseArrayOutput() BasicDependencyResponseArrayOutput {
	return o
}

func (o BasicDependencyResponseArrayOutput) ToBasicDependencyResponseArrayOutputWithContext(ctx context.Context) BasicDependencyResponseArrayOutput {
	return o
}

func (o BasicDependencyResponseArrayOutput) Index(i pulumi.IntInput) BasicDependencyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BasicDependencyResponse {
		return vs[0].([]BasicDependencyResponse)[vs[1].(int)]
	}).(BasicDependencyResponseOutput)
}

type DebugSetting struct {
	DetailLevel *string `pulumi:"detailLevel"`
}





type DebugSettingInput interface {
	pulumi.Input

	ToDebugSettingOutput() DebugSettingOutput
	ToDebugSettingOutputWithContext(context.Context) DebugSettingOutput
}

type DebugSettingArgs struct {
	DetailLevel pulumi.StringPtrInput `pulumi:"detailLevel"`
}

func (DebugSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DebugSetting)(nil)).Elem()
}

func (i DebugSettingArgs) ToDebugSettingOutput() DebugSettingOutput {
	return i.ToDebugSettingOutputWithContext(context.Background())
}

func (i DebugSettingArgs) ToDebugSettingOutputWithContext(ctx context.Context) DebugSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebugSettingOutput)
}

func (i DebugSettingArgs) ToDebugSettingPtrOutput() DebugSettingPtrOutput {
	return i.ToDebugSettingPtrOutputWithContext(context.Background())
}

func (i DebugSettingArgs) ToDebugSettingPtrOutputWithContext(ctx context.Context) DebugSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebugSettingOutput).ToDebugSettingPtrOutputWithContext(ctx)
}









type DebugSettingPtrInput interface {
	pulumi.Input

	ToDebugSettingPtrOutput() DebugSettingPtrOutput
	ToDebugSettingPtrOutputWithContext(context.Context) DebugSettingPtrOutput
}

type debugSettingPtrType DebugSettingArgs

func DebugSettingPtr(v *DebugSettingArgs) DebugSettingPtrInput {
	return (*debugSettingPtrType)(v)
}

func (*debugSettingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DebugSetting)(nil)).Elem()
}

func (i *debugSettingPtrType) ToDebugSettingPtrOutput() DebugSettingPtrOutput {
	return i.ToDebugSettingPtrOutputWithContext(context.Background())
}

func (i *debugSettingPtrType) ToDebugSettingPtrOutputWithContext(ctx context.Context) DebugSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebugSettingPtrOutput)
}

type DebugSettingOutput struct{ *pulumi.OutputState }

func (DebugSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DebugSetting)(nil)).Elem()
}

func (o DebugSettingOutput) ToDebugSettingOutput() DebugSettingOutput {
	return o
}

func (o DebugSettingOutput) ToDebugSettingOutputWithContext(ctx context.Context) DebugSettingOutput {
	return o
}

func (o DebugSettingOutput) ToDebugSettingPtrOutput() DebugSettingPtrOutput {
	return o.ToDebugSettingPtrOutputWithContext(context.Background())
}

func (o DebugSettingOutput) ToDebugSettingPtrOutputWithContext(ctx context.Context) DebugSettingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DebugSetting) *DebugSetting {
		return &v
	}).(DebugSettingPtrOutput)
}

func (o DebugSettingOutput) DetailLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DebugSetting) *string { return v.DetailLevel }).(pulumi.StringPtrOutput)
}

type DebugSettingPtrOutput struct{ *pulumi.OutputState }

func (DebugSettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DebugSetting)(nil)).Elem()
}

func (o DebugSettingPtrOutput) ToDebugSettingPtrOutput() DebugSettingPtrOutput {
	return o
}

func (o DebugSettingPtrOutput) ToDebugSettingPtrOutputWithContext(ctx context.Context) DebugSettingPtrOutput {
	return o
}

func (o DebugSettingPtrOutput) Elem() DebugSettingOutput {
	return o.ApplyT(func(v *DebugSetting) DebugSetting {
		if v != nil {
			return *v
		}
		var ret DebugSetting
		return ret
	}).(DebugSettingOutput)
}

func (o DebugSettingPtrOutput) DetailLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DebugSetting) *string {
		if v == nil {
			return nil
		}
		return v.DetailLevel
	}).(pulumi.StringPtrOutput)
}

type DebugSettingResponse struct {
	DetailLevel *string `pulumi:"detailLevel"`
}

type DebugSettingResponseOutput struct{ *pulumi.OutputState }

func (DebugSettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DebugSettingResponse)(nil)).Elem()
}

func (o DebugSettingResponseOutput) ToDebugSettingResponseOutput() DebugSettingResponseOutput {
	return o
}

func (o DebugSettingResponseOutput) ToDebugSettingResponseOutputWithContext(ctx context.Context) DebugSettingResponseOutput {
	return o
}

func (o DebugSettingResponseOutput) DetailLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DebugSettingResponse) *string { return v.DetailLevel }).(pulumi.StringPtrOutput)
}

type DebugSettingResponsePtrOutput struct{ *pulumi.OutputState }

func (DebugSettingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DebugSettingResponse)(nil)).Elem()
}

func (o DebugSettingResponsePtrOutput) ToDebugSettingResponsePtrOutput() DebugSettingResponsePtrOutput {
	return o
}

func (o DebugSettingResponsePtrOutput) ToDebugSettingResponsePtrOutputWithContext(ctx context.Context) DebugSettingResponsePtrOutput {
	return o
}

func (o DebugSettingResponsePtrOutput) Elem() DebugSettingResponseOutput {
	return o.ApplyT(func(v *DebugSettingResponse) DebugSettingResponse {
		if v != nil {
			return *v
		}
		var ret DebugSettingResponse
		return ret
	}).(DebugSettingResponseOutput)
}

func (o DebugSettingResponsePtrOutput) DetailLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DebugSettingResponse) *string {
		if v == nil {
			return nil
		}
		return v.DetailLevel
	}).(pulumi.StringPtrOutput)
}

type DependencyResponse struct {
	DependsOn    []BasicDependencyResponse `pulumi:"dependsOn"`
	Id           *string                   `pulumi:"id"`
	ResourceName *string                   `pulumi:"resourceName"`
	ResourceType *string                   `pulumi:"resourceType"`
}

type DependencyResponseOutput struct{ *pulumi.OutputState }

func (DependencyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DependencyResponse)(nil)).Elem()
}

func (o DependencyResponseOutput) ToDependencyResponseOutput() DependencyResponseOutput {
	return o
}

func (o DependencyResponseOutput) ToDependencyResponseOutputWithContext(ctx context.Context) DependencyResponseOutput {
	return o
}

func (o DependencyResponseOutput) DependsOn() BasicDependencyResponseArrayOutput {
	return o.ApplyT(func(v DependencyResponse) []BasicDependencyResponse { return v.DependsOn }).(BasicDependencyResponseArrayOutput)
}

func (o DependencyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DependencyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o DependencyResponseOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DependencyResponse) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o DependencyResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DependencyResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

type DependencyResponseArrayOutput struct{ *pulumi.OutputState }

func (DependencyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DependencyResponse)(nil)).Elem()
}

func (o DependencyResponseArrayOutput) ToDependencyResponseArrayOutput() DependencyResponseArrayOutput {
	return o
}

func (o DependencyResponseArrayOutput) ToDependencyResponseArrayOutputWithContext(ctx context.Context) DependencyResponseArrayOutput {
	return o
}

func (o DependencyResponseArrayOutput) Index(i pulumi.IntInput) DependencyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DependencyResponse {
		return vs[0].([]DependencyResponse)[vs[1].(int)]
	}).(DependencyResponseOutput)
}

type DeploymentProperties struct {
	DebugSetting      *DebugSetting      `pulumi:"debugSetting"`
	Mode              DeploymentMode     `pulumi:"mode"`
	OnErrorDeployment *OnErrorDeployment `pulumi:"onErrorDeployment"`
	Parameters        interface{}        `pulumi:"parameters"`
	ParametersLink    *ParametersLink    `pulumi:"parametersLink"`
	Template          interface{}        `pulumi:"template"`
	TemplateLink      *TemplateLink      `pulumi:"templateLink"`
}





type DeploymentPropertiesInput interface {
	pulumi.Input

	ToDeploymentPropertiesOutput() DeploymentPropertiesOutput
	ToDeploymentPropertiesOutputWithContext(context.Context) DeploymentPropertiesOutput
}

type DeploymentPropertiesArgs struct {
	DebugSetting      DebugSettingPtrInput      `pulumi:"debugSetting"`
	Mode              DeploymentModeInput       `pulumi:"mode"`
	OnErrorDeployment OnErrorDeploymentPtrInput `pulumi:"onErrorDeployment"`
	Parameters        pulumi.Input              `pulumi:"parameters"`
	ParametersLink    ParametersLinkPtrInput    `pulumi:"parametersLink"`
	Template          pulumi.Input              `pulumi:"template"`
	TemplateLink      TemplateLinkPtrInput      `pulumi:"templateLink"`
}

func (DeploymentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentProperties)(nil)).Elem()
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesOutput() DeploymentPropertiesOutput {
	return i.ToDeploymentPropertiesOutputWithContext(context.Background())
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesOutputWithContext(ctx context.Context) DeploymentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesOutput)
}

type DeploymentPropertiesOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentProperties)(nil)).Elem()
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesOutput() DeploymentPropertiesOutput {
	return o
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesOutputWithContext(ctx context.Context) DeploymentPropertiesOutput {
	return o
}

func (o DeploymentPropertiesOutput) DebugSetting() DebugSettingPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *DebugSetting { return v.DebugSetting }).(DebugSettingPtrOutput)
}

func (o DeploymentPropertiesOutput) Mode() DeploymentModeOutput {
	return o.ApplyT(func(v DeploymentProperties) DeploymentMode { return v.Mode }).(DeploymentModeOutput)
}

func (o DeploymentPropertiesOutput) OnErrorDeployment() OnErrorDeploymentPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *OnErrorDeployment { return v.OnErrorDeployment }).(OnErrorDeploymentPtrOutput)
}

func (o DeploymentPropertiesOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentProperties) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesOutput) ParametersLink() ParametersLinkPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *ParametersLink { return v.ParametersLink }).(ParametersLinkPtrOutput)
}

func (o DeploymentPropertiesOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentProperties) interface{} { return v.Template }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesOutput) TemplateLink() TemplateLinkPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *TemplateLink { return v.TemplateLink }).(TemplateLinkPtrOutput)
}

type DeploymentPropertiesExtendedResponse struct {
	CorrelationId     string                             `pulumi:"correlationId"`
	DebugSetting      *DebugSettingResponse              `pulumi:"debugSetting"`
	Dependencies      []DependencyResponse               `pulumi:"dependencies"`
	Error             ErrorResponseResponse              `pulumi:"error"`
	Mode              *string                            `pulumi:"mode"`
	OnErrorDeployment *OnErrorDeploymentExtendedResponse `pulumi:"onErrorDeployment"`
	Outputs           interface{}                        `pulumi:"outputs"`
	Parameters        interface{}                        `pulumi:"parameters"`
	ParametersLink    *ParametersLinkResponse            `pulumi:"parametersLink"`
	Providers         []ProviderResponse                 `pulumi:"providers"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	Template          interface{}                        `pulumi:"template"`
	TemplateLink      *TemplateLinkResponse              `pulumi:"templateLink"`
	Timestamp         string                             `pulumi:"timestamp"`
}

type DeploymentPropertiesExtendedResponseOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesExtendedResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentPropertiesExtendedResponse)(nil)).Elem()
}

func (o DeploymentPropertiesExtendedResponseOutput) ToDeploymentPropertiesExtendedResponseOutput() DeploymentPropertiesExtendedResponseOutput {
	return o
}

func (o DeploymentPropertiesExtendedResponseOutput) ToDeploymentPropertiesExtendedResponseOutputWithContext(ctx context.Context) DeploymentPropertiesExtendedResponseOutput {
	return o
}

func (o DeploymentPropertiesExtendedResponseOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) DebugSetting() DebugSettingResponsePtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *DebugSettingResponse { return v.DebugSetting }).(DebugSettingResponsePtrOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Dependencies() DependencyResponseArrayOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) []DependencyResponse { return v.Dependencies }).(DependencyResponseArrayOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Error() ErrorResponseResponseOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) ErrorResponseResponse { return v.Error }).(ErrorResponseResponseOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) OnErrorDeployment() OnErrorDeploymentExtendedResponsePtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *OnErrorDeploymentExtendedResponse {
		return v.OnErrorDeployment
	}).(OnErrorDeploymentExtendedResponsePtrOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) interface{} { return v.Outputs }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) ParametersLink() ParametersLinkResponsePtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *ParametersLinkResponse { return v.ParametersLink }).(ParametersLinkResponsePtrOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Providers() ProviderResponseArrayOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) []ProviderResponse { return v.Providers }).(ProviderResponseArrayOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) interface{} { return v.Template }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) TemplateLink() TemplateLinkResponsePtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *TemplateLinkResponse { return v.TemplateLink }).(TemplateLinkResponsePtrOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}

type ErrorAdditionalInfoResponseOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o ErrorAdditionalInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ErrorAdditionalInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) Index(i pulumi.IntInput) ErrorAdditionalInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorAdditionalInfoResponse {
		return vs[0].([]ErrorAdditionalInfoResponse)[vs[1].(int)]
	}).(ErrorAdditionalInfoResponseOutput)
}

type ErrorResponseResponse struct {
	AdditionalInfo []ErrorAdditionalInfoResponse `pulumi:"additionalInfo"`
	Code           string                        `pulumi:"code"`
	Details        []ErrorResponseResponse       `pulumi:"details"`
	Message        string                        `pulumi:"message"`
	Target         string                        `pulumi:"target"`
}

type ErrorResponseResponseOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutput() ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutputWithContext(ctx context.Context) ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorResponseResponse) []ErrorAdditionalInfoResponse { return v.AdditionalInfo }).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorResponseResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorResponseResponseOutput) Details() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v ErrorResponseResponse) []ErrorResponseResponse { return v.Details }).(ErrorResponseResponseArrayOutput)
}

func (o ErrorResponseResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorResponseResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorResponseResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponseArrayOutput) ToErrorResponseResponseArrayOutput() ErrorResponseResponseArrayOutput {
	return o
}

func (o ErrorResponseResponseArrayOutput) ToErrorResponseResponseArrayOutputWithContext(ctx context.Context) ErrorResponseResponseArrayOutput {
	return o
}

func (o ErrorResponseResponseArrayOutput) Index(i pulumi.IntInput) ErrorResponseResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorResponseResponse {
		return vs[0].([]ErrorResponseResponse)[vs[1].(int)]
	}).(ErrorResponseResponseOutput)
}

type Identity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                            `pulumi:"principalId"`
	TenantId               string                                            `pulumi:"tenantId"`
	Type                   *string                                           `pulumi:"type"`
	UserAssignedIdentities map[string]IdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
}

type IdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type IdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutput() IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o IdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type IdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutput() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) IdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]IdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(IdentityResponseUserAssignedIdentitiesOutput)
}

type OnErrorDeployment struct {
	DeploymentName *string                `pulumi:"deploymentName"`
	Type           *OnErrorDeploymentType `pulumi:"type"`
}





type OnErrorDeploymentInput interface {
	pulumi.Input

	ToOnErrorDeploymentOutput() OnErrorDeploymentOutput
	ToOnErrorDeploymentOutputWithContext(context.Context) OnErrorDeploymentOutput
}

type OnErrorDeploymentArgs struct {
	DeploymentName pulumi.StringPtrInput         `pulumi:"deploymentName"`
	Type           OnErrorDeploymentTypePtrInput `pulumi:"type"`
}

func (OnErrorDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnErrorDeployment)(nil)).Elem()
}

func (i OnErrorDeploymentArgs) ToOnErrorDeploymentOutput() OnErrorDeploymentOutput {
	return i.ToOnErrorDeploymentOutputWithContext(context.Background())
}

func (i OnErrorDeploymentArgs) ToOnErrorDeploymentOutputWithContext(ctx context.Context) OnErrorDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnErrorDeploymentOutput)
}

func (i OnErrorDeploymentArgs) ToOnErrorDeploymentPtrOutput() OnErrorDeploymentPtrOutput {
	return i.ToOnErrorDeploymentPtrOutputWithContext(context.Background())
}

func (i OnErrorDeploymentArgs) ToOnErrorDeploymentPtrOutputWithContext(ctx context.Context) OnErrorDeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnErrorDeploymentOutput).ToOnErrorDeploymentPtrOutputWithContext(ctx)
}









type OnErrorDeploymentPtrInput interface {
	pulumi.Input

	ToOnErrorDeploymentPtrOutput() OnErrorDeploymentPtrOutput
	ToOnErrorDeploymentPtrOutputWithContext(context.Context) OnErrorDeploymentPtrOutput
}

type onErrorDeploymentPtrType OnErrorDeploymentArgs

func OnErrorDeploymentPtr(v *OnErrorDeploymentArgs) OnErrorDeploymentPtrInput {
	return (*onErrorDeploymentPtrType)(v)
}

func (*onErrorDeploymentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OnErrorDeployment)(nil)).Elem()
}

func (i *onErrorDeploymentPtrType) ToOnErrorDeploymentPtrOutput() OnErrorDeploymentPtrOutput {
	return i.ToOnErrorDeploymentPtrOutputWithContext(context.Background())
}

func (i *onErrorDeploymentPtrType) ToOnErrorDeploymentPtrOutputWithContext(ctx context.Context) OnErrorDeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnErrorDeploymentPtrOutput)
}

type OnErrorDeploymentOutput struct{ *pulumi.OutputState }

func (OnErrorDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnErrorDeployment)(nil)).Elem()
}

func (o OnErrorDeploymentOutput) ToOnErrorDeploymentOutput() OnErrorDeploymentOutput {
	return o
}

func (o OnErrorDeploymentOutput) ToOnErrorDeploymentOutputWithContext(ctx context.Context) OnErrorDeploymentOutput {
	return o
}

func (o OnErrorDeploymentOutput) ToOnErrorDeploymentPtrOutput() OnErrorDeploymentPtrOutput {
	return o.ToOnErrorDeploymentPtrOutputWithContext(context.Background())
}

func (o OnErrorDeploymentOutput) ToOnErrorDeploymentPtrOutputWithContext(ctx context.Context) OnErrorDeploymentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OnErrorDeployment) *OnErrorDeployment {
		return &v
	}).(OnErrorDeploymentPtrOutput)
}

func (o OnErrorDeploymentOutput) DeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnErrorDeployment) *string { return v.DeploymentName }).(pulumi.StringPtrOutput)
}

func (o OnErrorDeploymentOutput) Type() OnErrorDeploymentTypePtrOutput {
	return o.ApplyT(func(v OnErrorDeployment) *OnErrorDeploymentType { return v.Type }).(OnErrorDeploymentTypePtrOutput)
}

type OnErrorDeploymentPtrOutput struct{ *pulumi.OutputState }

func (OnErrorDeploymentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnErrorDeployment)(nil)).Elem()
}

func (o OnErrorDeploymentPtrOutput) ToOnErrorDeploymentPtrOutput() OnErrorDeploymentPtrOutput {
	return o
}

func (o OnErrorDeploymentPtrOutput) ToOnErrorDeploymentPtrOutputWithContext(ctx context.Context) OnErrorDeploymentPtrOutput {
	return o
}

func (o OnErrorDeploymentPtrOutput) Elem() OnErrorDeploymentOutput {
	return o.ApplyT(func(v *OnErrorDeployment) OnErrorDeployment {
		if v != nil {
			return *v
		}
		var ret OnErrorDeployment
		return ret
	}).(OnErrorDeploymentOutput)
}

func (o OnErrorDeploymentPtrOutput) DeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnErrorDeployment) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentName
	}).(pulumi.StringPtrOutput)
}

func (o OnErrorDeploymentPtrOutput) Type() OnErrorDeploymentTypePtrOutput {
	return o.ApplyT(func(v *OnErrorDeployment) *OnErrorDeploymentType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(OnErrorDeploymentTypePtrOutput)
}

type OnErrorDeploymentExtendedResponse struct {
	DeploymentName    *string `pulumi:"deploymentName"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Type              *string `pulumi:"type"`
}

type OnErrorDeploymentExtendedResponseOutput struct{ *pulumi.OutputState }

func (OnErrorDeploymentExtendedResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnErrorDeploymentExtendedResponse)(nil)).Elem()
}

func (o OnErrorDeploymentExtendedResponseOutput) ToOnErrorDeploymentExtendedResponseOutput() OnErrorDeploymentExtendedResponseOutput {
	return o
}

func (o OnErrorDeploymentExtendedResponseOutput) ToOnErrorDeploymentExtendedResponseOutputWithContext(ctx context.Context) OnErrorDeploymentExtendedResponseOutput {
	return o
}

func (o OnErrorDeploymentExtendedResponseOutput) DeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnErrorDeploymentExtendedResponse) *string { return v.DeploymentName }).(pulumi.StringPtrOutput)
}

func (o OnErrorDeploymentExtendedResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v OnErrorDeploymentExtendedResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o OnErrorDeploymentExtendedResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnErrorDeploymentExtendedResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type OnErrorDeploymentExtendedResponsePtrOutput struct{ *pulumi.OutputState }

func (OnErrorDeploymentExtendedResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnErrorDeploymentExtendedResponse)(nil)).Elem()
}

func (o OnErrorDeploymentExtendedResponsePtrOutput) ToOnErrorDeploymentExtendedResponsePtrOutput() OnErrorDeploymentExtendedResponsePtrOutput {
	return o
}

func (o OnErrorDeploymentExtendedResponsePtrOutput) ToOnErrorDeploymentExtendedResponsePtrOutputWithContext(ctx context.Context) OnErrorDeploymentExtendedResponsePtrOutput {
	return o
}

func (o OnErrorDeploymentExtendedResponsePtrOutput) Elem() OnErrorDeploymentExtendedResponseOutput {
	return o.ApplyT(func(v *OnErrorDeploymentExtendedResponse) OnErrorDeploymentExtendedResponse {
		if v != nil {
			return *v
		}
		var ret OnErrorDeploymentExtendedResponse
		return ret
	}).(OnErrorDeploymentExtendedResponseOutput)
}

func (o OnErrorDeploymentExtendedResponsePtrOutput) DeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnErrorDeploymentExtendedResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentName
	}).(pulumi.StringPtrOutput)
}

func (o OnErrorDeploymentExtendedResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnErrorDeploymentExtendedResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o OnErrorDeploymentExtendedResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OnErrorDeploymentExtendedResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ParametersLink struct {
	ContentVersion *string `pulumi:"contentVersion"`
	Uri            string  `pulumi:"uri"`
}





type ParametersLinkInput interface {
	pulumi.Input

	ToParametersLinkOutput() ParametersLinkOutput
	ToParametersLinkOutputWithContext(context.Context) ParametersLinkOutput
}

type ParametersLinkArgs struct {
	ContentVersion pulumi.StringPtrInput `pulumi:"contentVersion"`
	Uri            pulumi.StringInput    `pulumi:"uri"`
}

func (ParametersLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParametersLink)(nil)).Elem()
}

func (i ParametersLinkArgs) ToParametersLinkOutput() ParametersLinkOutput {
	return i.ToParametersLinkOutputWithContext(context.Background())
}

func (i ParametersLinkArgs) ToParametersLinkOutputWithContext(ctx context.Context) ParametersLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersLinkOutput)
}

func (i ParametersLinkArgs) ToParametersLinkPtrOutput() ParametersLinkPtrOutput {
	return i.ToParametersLinkPtrOutputWithContext(context.Background())
}

func (i ParametersLinkArgs) ToParametersLinkPtrOutputWithContext(ctx context.Context) ParametersLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersLinkOutput).ToParametersLinkPtrOutputWithContext(ctx)
}









type ParametersLinkPtrInput interface {
	pulumi.Input

	ToParametersLinkPtrOutput() ParametersLinkPtrOutput
	ToParametersLinkPtrOutputWithContext(context.Context) ParametersLinkPtrOutput
}

type parametersLinkPtrType ParametersLinkArgs

func ParametersLinkPtr(v *ParametersLinkArgs) ParametersLinkPtrInput {
	return (*parametersLinkPtrType)(v)
}

func (*parametersLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ParametersLink)(nil)).Elem()
}

func (i *parametersLinkPtrType) ToParametersLinkPtrOutput() ParametersLinkPtrOutput {
	return i.ToParametersLinkPtrOutputWithContext(context.Background())
}

func (i *parametersLinkPtrType) ToParametersLinkPtrOutputWithContext(ctx context.Context) ParametersLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersLinkPtrOutput)
}

type ParametersLinkOutput struct{ *pulumi.OutputState }

func (ParametersLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParametersLink)(nil)).Elem()
}

func (o ParametersLinkOutput) ToParametersLinkOutput() ParametersLinkOutput {
	return o
}

func (o ParametersLinkOutput) ToParametersLinkOutputWithContext(ctx context.Context) ParametersLinkOutput {
	return o
}

func (o ParametersLinkOutput) ToParametersLinkPtrOutput() ParametersLinkPtrOutput {
	return o.ToParametersLinkPtrOutputWithContext(context.Background())
}

func (o ParametersLinkOutput) ToParametersLinkPtrOutputWithContext(ctx context.Context) ParametersLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParametersLink) *ParametersLink {
		return &v
	}).(ParametersLinkPtrOutput)
}

func (o ParametersLinkOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParametersLink) *string { return v.ContentVersion }).(pulumi.StringPtrOutput)
}

func (o ParametersLinkOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ParametersLink) string { return v.Uri }).(pulumi.StringOutput)
}

type ParametersLinkPtrOutput struct{ *pulumi.OutputState }

func (ParametersLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParametersLink)(nil)).Elem()
}

func (o ParametersLinkPtrOutput) ToParametersLinkPtrOutput() ParametersLinkPtrOutput {
	return o
}

func (o ParametersLinkPtrOutput) ToParametersLinkPtrOutputWithContext(ctx context.Context) ParametersLinkPtrOutput {
	return o
}

func (o ParametersLinkPtrOutput) Elem() ParametersLinkOutput {
	return o.ApplyT(func(v *ParametersLink) ParametersLink {
		if v != nil {
			return *v
		}
		var ret ParametersLink
		return ret
	}).(ParametersLinkOutput)
}

func (o ParametersLinkPtrOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParametersLink) *string {
		if v == nil {
			return nil
		}
		return v.ContentVersion
	}).(pulumi.StringPtrOutput)
}

func (o ParametersLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParametersLink) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type ParametersLinkResponse struct {
	ContentVersion *string `pulumi:"contentVersion"`
	Uri            string  `pulumi:"uri"`
}

type ParametersLinkResponseOutput struct{ *pulumi.OutputState }

func (ParametersLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParametersLinkResponse)(nil)).Elem()
}

func (o ParametersLinkResponseOutput) ToParametersLinkResponseOutput() ParametersLinkResponseOutput {
	return o
}

func (o ParametersLinkResponseOutput) ToParametersLinkResponseOutputWithContext(ctx context.Context) ParametersLinkResponseOutput {
	return o
}

func (o ParametersLinkResponseOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParametersLinkResponse) *string { return v.ContentVersion }).(pulumi.StringPtrOutput)
}

func (o ParametersLinkResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ParametersLinkResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type ParametersLinkResponsePtrOutput struct{ *pulumi.OutputState }

func (ParametersLinkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParametersLinkResponse)(nil)).Elem()
}

func (o ParametersLinkResponsePtrOutput) ToParametersLinkResponsePtrOutput() ParametersLinkResponsePtrOutput {
	return o
}

func (o ParametersLinkResponsePtrOutput) ToParametersLinkResponsePtrOutputWithContext(ctx context.Context) ParametersLinkResponsePtrOutput {
	return o
}

func (o ParametersLinkResponsePtrOutput) Elem() ParametersLinkResponseOutput {
	return o.ApplyT(func(v *ParametersLinkResponse) ParametersLinkResponse {
		if v != nil {
			return *v
		}
		var ret ParametersLinkResponse
		return ret
	}).(ParametersLinkResponseOutput)
}

func (o ParametersLinkResponsePtrOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParametersLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentVersion
	}).(pulumi.StringPtrOutput)
}

func (o ParametersLinkResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParametersLinkResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type Plan struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Product       pulumi.StringPtrInput `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringPtrInput `pulumi:"publisher"`
	Version       pulumi.StringPtrInput `pulumi:"version"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

func (i PlanArgs) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput).ToPlanPtrOutputWithContext(ctx)
}









type PlanPtrInput interface {
	pulumi.Input

	ToPlanPtrOutput() PlanPtrOutput
	ToPlanPtrOutputWithContext(context.Context) PlanPtrOutput
}

type planPtrType PlanArgs

func PlanPtr(v *PlanArgs) PlanPtrInput {
	return (*planPtrType)(v)
}

func (*planPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (i *planPtrType) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i *planPtrType) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanPtrOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o.ToPlanPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Plan) *Plan {
		return &v
	}).(PlanPtrOutput)
}

func (o PlanOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type PlanPtrOutput struct{ *pulumi.OutputState }

func (PlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (o PlanPtrOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) Elem() PlanOutput {
	return o.ApplyT(func(v *Plan) Plan {
		if v != nil {
			return *v
		}
		var ret Plan
		return ret
	}).(PlanOutput)
}

func (o PlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type PlanResponse struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type PlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) Elem() PlanResponseOutput {
	return o.ApplyT(func(v *PlanResponse) PlanResponse {
		if v != nil {
			return *v
		}
		var ret PlanResponse
		return ret
	}).(PlanResponseOutput)
}

func (o PlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ProviderResourceTypeResponse struct {
	Aliases      []AliasTypeResponse   `pulumi:"aliases"`
	ApiVersions  []string              `pulumi:"apiVersions"`
	Capabilities *string               `pulumi:"capabilities"`
	Locations    []string              `pulumi:"locations"`
	Properties   map[string]string     `pulumi:"properties"`
	ResourceType *string               `pulumi:"resourceType"`
	ZoneMappings []ZoneMappingResponse `pulumi:"zoneMappings"`
}

type ProviderResourceTypeResponseOutput struct{ *pulumi.OutputState }

func (ProviderResourceTypeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderResourceTypeResponse)(nil)).Elem()
}

func (o ProviderResourceTypeResponseOutput) ToProviderResourceTypeResponseOutput() ProviderResourceTypeResponseOutput {
	return o
}

func (o ProviderResourceTypeResponseOutput) ToProviderResourceTypeResponseOutputWithContext(ctx context.Context) ProviderResourceTypeResponseOutput {
	return o
}

func (o ProviderResourceTypeResponseOutput) Aliases() AliasTypeResponseArrayOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) []AliasTypeResponse { return v.Aliases }).(AliasTypeResponseArrayOutput)
}

func (o ProviderResourceTypeResponseOutput) ApiVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) []string { return v.ApiVersions }).(pulumi.StringArrayOutput)
}

func (o ProviderResourceTypeResponseOutput) Capabilities() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) *string { return v.Capabilities }).(pulumi.StringPtrOutput)
}

func (o ProviderResourceTypeResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o ProviderResourceTypeResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ProviderResourceTypeResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o ProviderResourceTypeResponseOutput) ZoneMappings() ZoneMappingResponseArrayOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) []ZoneMappingResponse { return v.ZoneMappings }).(ZoneMappingResponseArrayOutput)
}

type ProviderResourceTypeResponseArrayOutput struct{ *pulumi.OutputState }

func (ProviderResourceTypeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderResourceTypeResponse)(nil)).Elem()
}

func (o ProviderResourceTypeResponseArrayOutput) ToProviderResourceTypeResponseArrayOutput() ProviderResourceTypeResponseArrayOutput {
	return o
}

func (o ProviderResourceTypeResponseArrayOutput) ToProviderResourceTypeResponseArrayOutputWithContext(ctx context.Context) ProviderResourceTypeResponseArrayOutput {
	return o
}

func (o ProviderResourceTypeResponseArrayOutput) Index(i pulumi.IntInput) ProviderResourceTypeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderResourceTypeResponse {
		return vs[0].([]ProviderResourceTypeResponse)[vs[1].(int)]
	}).(ProviderResourceTypeResponseOutput)
}

type ProviderResponse struct {
	Id                 string                         `pulumi:"id"`
	Namespace          *string                        `pulumi:"namespace"`
	RegistrationPolicy string                         `pulumi:"registrationPolicy"`
	RegistrationState  string                         `pulumi:"registrationState"`
	ResourceTypes      []ProviderResourceTypeResponse `pulumi:"resourceTypes"`
}

type ProviderResponseOutput struct{ *pulumi.OutputState }

func (ProviderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderResponse)(nil)).Elem()
}

func (o ProviderResponseOutput) ToProviderResponseOutput() ProviderResponseOutput {
	return o
}

func (o ProviderResponseOutput) ToProviderResponseOutputWithContext(ctx context.Context) ProviderResponseOutput {
	return o
}

func (o ProviderResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ProviderResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o ProviderResponseOutput) RegistrationPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderResponse) string { return v.RegistrationPolicy }).(pulumi.StringOutput)
}

func (o ProviderResponseOutput) RegistrationState() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderResponse) string { return v.RegistrationState }).(pulumi.StringOutput)
}

func (o ProviderResponseOutput) ResourceTypes() ProviderResourceTypeResponseArrayOutput {
	return o.ApplyT(func(v ProviderResponse) []ProviderResourceTypeResponse { return v.ResourceTypes }).(ProviderResourceTypeResponseArrayOutput)
}

type ProviderResponseArrayOutput struct{ *pulumi.OutputState }

func (ProviderResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderResponse)(nil)).Elem()
}

func (o ProviderResponseArrayOutput) ToProviderResponseArrayOutput() ProviderResponseArrayOutput {
	return o
}

func (o ProviderResponseArrayOutput) ToProviderResponseArrayOutputWithContext(ctx context.Context) ProviderResponseArrayOutput {
	return o
}

func (o ProviderResponseArrayOutput) Index(i pulumi.IntInput) ProviderResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderResponse {
		return vs[0].([]ProviderResponse)[vs[1].(int)]
	}).(ProviderResponseOutput)
}

type ResourceGroupPropertiesResponse struct {
	ProvisioningState string `pulumi:"provisioningState"`
}

type ResourceGroupPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ResourceGroupPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupPropertiesResponse)(nil)).Elem()
}

func (o ResourceGroupPropertiesResponseOutput) ToResourceGroupPropertiesResponseOutput() ResourceGroupPropertiesResponseOutput {
	return o
}

func (o ResourceGroupPropertiesResponseOutput) ToResourceGroupPropertiesResponseOutputWithContext(ctx context.Context) ResourceGroupPropertiesResponseOutput {
	return o
}

func (o ResourceGroupPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceGroupPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Model    *string `pulumi:"model"`
	Name     *string `pulumi:"name"`
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
	Model    pulumi.StringPtrInput `pulumi:"model"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
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

func (o SkuOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Model }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Model
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
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
	Model    *string `pulumi:"model"`
	Name     *string `pulumi:"name"`
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

func (o SkuResponseOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Model }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Model
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
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

type TemplateLink struct {
	ContentVersion *string `pulumi:"contentVersion"`
	Uri            string  `pulumi:"uri"`
}





type TemplateLinkInput interface {
	pulumi.Input

	ToTemplateLinkOutput() TemplateLinkOutput
	ToTemplateLinkOutputWithContext(context.Context) TemplateLinkOutput
}

type TemplateLinkArgs struct {
	ContentVersion pulumi.StringPtrInput `pulumi:"contentVersion"`
	Uri            pulumi.StringInput    `pulumi:"uri"`
}

func (TemplateLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateLink)(nil)).Elem()
}

func (i TemplateLinkArgs) ToTemplateLinkOutput() TemplateLinkOutput {
	return i.ToTemplateLinkOutputWithContext(context.Background())
}

func (i TemplateLinkArgs) ToTemplateLinkOutputWithContext(ctx context.Context) TemplateLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateLinkOutput)
}

func (i TemplateLinkArgs) ToTemplateLinkPtrOutput() TemplateLinkPtrOutput {
	return i.ToTemplateLinkPtrOutputWithContext(context.Background())
}

func (i TemplateLinkArgs) ToTemplateLinkPtrOutputWithContext(ctx context.Context) TemplateLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateLinkOutput).ToTemplateLinkPtrOutputWithContext(ctx)
}









type TemplateLinkPtrInput interface {
	pulumi.Input

	ToTemplateLinkPtrOutput() TemplateLinkPtrOutput
	ToTemplateLinkPtrOutputWithContext(context.Context) TemplateLinkPtrOutput
}

type templateLinkPtrType TemplateLinkArgs

func TemplateLinkPtr(v *TemplateLinkArgs) TemplateLinkPtrInput {
	return (*templateLinkPtrType)(v)
}

func (*templateLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateLink)(nil)).Elem()
}

func (i *templateLinkPtrType) ToTemplateLinkPtrOutput() TemplateLinkPtrOutput {
	return i.ToTemplateLinkPtrOutputWithContext(context.Background())
}

func (i *templateLinkPtrType) ToTemplateLinkPtrOutputWithContext(ctx context.Context) TemplateLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateLinkPtrOutput)
}

type TemplateLinkOutput struct{ *pulumi.OutputState }

func (TemplateLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateLink)(nil)).Elem()
}

func (o TemplateLinkOutput) ToTemplateLinkOutput() TemplateLinkOutput {
	return o
}

func (o TemplateLinkOutput) ToTemplateLinkOutputWithContext(ctx context.Context) TemplateLinkOutput {
	return o
}

func (o TemplateLinkOutput) ToTemplateLinkPtrOutput() TemplateLinkPtrOutput {
	return o.ToTemplateLinkPtrOutputWithContext(context.Background())
}

func (o TemplateLinkOutput) ToTemplateLinkPtrOutputWithContext(ctx context.Context) TemplateLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TemplateLink) *TemplateLink {
		return &v
	}).(TemplateLinkPtrOutput)
}

func (o TemplateLinkOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLink) *string { return v.ContentVersion }).(pulumi.StringPtrOutput)
}

func (o TemplateLinkOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateLink) string { return v.Uri }).(pulumi.StringOutput)
}

type TemplateLinkPtrOutput struct{ *pulumi.OutputState }

func (TemplateLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateLink)(nil)).Elem()
}

func (o TemplateLinkPtrOutput) ToTemplateLinkPtrOutput() TemplateLinkPtrOutput {
	return o
}

func (o TemplateLinkPtrOutput) ToTemplateLinkPtrOutputWithContext(ctx context.Context) TemplateLinkPtrOutput {
	return o
}

func (o TemplateLinkPtrOutput) Elem() TemplateLinkOutput {
	return o.ApplyT(func(v *TemplateLink) TemplateLink {
		if v != nil {
			return *v
		}
		var ret TemplateLink
		return ret
	}).(TemplateLinkOutput)
}

func (o TemplateLinkPtrOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLink) *string {
		if v == nil {
			return nil
		}
		return v.ContentVersion
	}).(pulumi.StringPtrOutput)
}

func (o TemplateLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLink) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type TemplateLinkResponse struct {
	ContentVersion *string `pulumi:"contentVersion"`
	Uri            string  `pulumi:"uri"`
}

type TemplateLinkResponseOutput struct{ *pulumi.OutputState }

func (TemplateLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateLinkResponse)(nil)).Elem()
}

func (o TemplateLinkResponseOutput) ToTemplateLinkResponseOutput() TemplateLinkResponseOutput {
	return o
}

func (o TemplateLinkResponseOutput) ToTemplateLinkResponseOutputWithContext(ctx context.Context) TemplateLinkResponseOutput {
	return o
}

func (o TemplateLinkResponseOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLinkResponse) *string { return v.ContentVersion }).(pulumi.StringPtrOutput)
}

func (o TemplateLinkResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateLinkResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type TemplateLinkResponsePtrOutput struct{ *pulumi.OutputState }

func (TemplateLinkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateLinkResponse)(nil)).Elem()
}

func (o TemplateLinkResponsePtrOutput) ToTemplateLinkResponsePtrOutput() TemplateLinkResponsePtrOutput {
	return o
}

func (o TemplateLinkResponsePtrOutput) ToTemplateLinkResponsePtrOutputWithContext(ctx context.Context) TemplateLinkResponsePtrOutput {
	return o
}

func (o TemplateLinkResponsePtrOutput) Elem() TemplateLinkResponseOutput {
	return o.ApplyT(func(v *TemplateLinkResponse) TemplateLinkResponse {
		if v != nil {
			return *v
		}
		var ret TemplateLinkResponse
		return ret
	}).(TemplateLinkResponseOutput)
}

func (o TemplateLinkResponsePtrOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentVersion
	}).(pulumi.StringPtrOutput)
}

func (o TemplateLinkResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLinkResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type ZoneMappingResponse struct {
	Location *string  `pulumi:"location"`
	Zones    []string `pulumi:"zones"`
}

type ZoneMappingResponseOutput struct{ *pulumi.OutputState }

func (ZoneMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneMappingResponse)(nil)).Elem()
}

func (o ZoneMappingResponseOutput) ToZoneMappingResponseOutput() ZoneMappingResponseOutput {
	return o
}

func (o ZoneMappingResponseOutput) ToZoneMappingResponseOutputWithContext(ctx context.Context) ZoneMappingResponseOutput {
	return o
}

func (o ZoneMappingResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ZoneMappingResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ZoneMappingResponseOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ZoneMappingResponse) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

type ZoneMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (ZoneMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneMappingResponse)(nil)).Elem()
}

func (o ZoneMappingResponseArrayOutput) ToZoneMappingResponseArrayOutput() ZoneMappingResponseArrayOutput {
	return o
}

func (o ZoneMappingResponseArrayOutput) ToZoneMappingResponseArrayOutputWithContext(ctx context.Context) ZoneMappingResponseArrayOutput {
	return o
}

func (o ZoneMappingResponseArrayOutput) Index(i pulumi.IntInput) ZoneMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ZoneMappingResponse {
		return vs[0].([]ZoneMappingResponse)[vs[1].(int)]
	}).(ZoneMappingResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AliasPathTypeResponseOutput{})
	pulumi.RegisterOutputType(AliasPathTypeResponseArrayOutput{})
	pulumi.RegisterOutputType(AliasTypeResponseOutput{})
	pulumi.RegisterOutputType(AliasTypeResponseArrayOutput{})
	pulumi.RegisterOutputType(BasicDependencyResponseOutput{})
	pulumi.RegisterOutputType(BasicDependencyResponseArrayOutput{})
	pulumi.RegisterOutputType(DebugSettingOutput{})
	pulumi.RegisterOutputType(DebugSettingPtrOutput{})
	pulumi.RegisterOutputType(DebugSettingResponseOutput{})
	pulumi.RegisterOutputType(DebugSettingResponsePtrOutput{})
	pulumi.RegisterOutputType(DependencyResponseOutput{})
	pulumi.RegisterOutputType(DependencyResponseArrayOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesExtendedResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(OnErrorDeploymentOutput{})
	pulumi.RegisterOutputType(OnErrorDeploymentPtrOutput{})
	pulumi.RegisterOutputType(OnErrorDeploymentExtendedResponseOutput{})
	pulumi.RegisterOutputType(OnErrorDeploymentExtendedResponsePtrOutput{})
	pulumi.RegisterOutputType(ParametersLinkOutput{})
	pulumi.RegisterOutputType(ParametersLinkPtrOutput{})
	pulumi.RegisterOutputType(ParametersLinkResponseOutput{})
	pulumi.RegisterOutputType(ParametersLinkResponsePtrOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponsePtrOutput{})
	pulumi.RegisterOutputType(ProviderResourceTypeResponseOutput{})
	pulumi.RegisterOutputType(ProviderResourceTypeResponseArrayOutput{})
	pulumi.RegisterOutputType(ProviderResponseOutput{})
	pulumi.RegisterOutputType(ProviderResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceGroupPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(TemplateLinkOutput{})
	pulumi.RegisterOutputType(TemplateLinkPtrOutput{})
	pulumi.RegisterOutputType(TemplateLinkResponseOutput{})
	pulumi.RegisterOutputType(TemplateLinkResponsePtrOutput{})
	pulumi.RegisterOutputType(ZoneMappingResponseOutput{})
	pulumi.RegisterOutputType(ZoneMappingResponseArrayOutput{})
}
