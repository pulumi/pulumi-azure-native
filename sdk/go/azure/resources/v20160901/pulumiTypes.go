


package v20160901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AliasPathTypeResponse struct {
	ApiVersions []string `pulumi:"apiVersions"`
	Path        *string  `pulumi:"path"`
}





type AliasPathTypeResponseInput interface {
	pulumi.Input

	ToAliasPathTypeResponseOutput() AliasPathTypeResponseOutput
	ToAliasPathTypeResponseOutputWithContext(context.Context) AliasPathTypeResponseOutput
}

type AliasPathTypeResponseArgs struct {
	ApiVersions pulumi.StringArrayInput `pulumi:"apiVersions"`
	Path        pulumi.StringPtrInput   `pulumi:"path"`
}

func (AliasPathTypeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasPathTypeResponse)(nil)).Elem()
}

func (i AliasPathTypeResponseArgs) ToAliasPathTypeResponseOutput() AliasPathTypeResponseOutput {
	return i.ToAliasPathTypeResponseOutputWithContext(context.Background())
}

func (i AliasPathTypeResponseArgs) ToAliasPathTypeResponseOutputWithContext(ctx context.Context) AliasPathTypeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPathTypeResponseOutput)
}





type AliasPathTypeResponseArrayInput interface {
	pulumi.Input

	ToAliasPathTypeResponseArrayOutput() AliasPathTypeResponseArrayOutput
	ToAliasPathTypeResponseArrayOutputWithContext(context.Context) AliasPathTypeResponseArrayOutput
}

type AliasPathTypeResponseArray []AliasPathTypeResponseInput

func (AliasPathTypeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AliasPathTypeResponse)(nil)).Elem()
}

func (i AliasPathTypeResponseArray) ToAliasPathTypeResponseArrayOutput() AliasPathTypeResponseArrayOutput {
	return i.ToAliasPathTypeResponseArrayOutputWithContext(context.Background())
}

func (i AliasPathTypeResponseArray) ToAliasPathTypeResponseArrayOutputWithContext(ctx context.Context) AliasPathTypeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPathTypeResponseArrayOutput)
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





type AliasTypeResponseInput interface {
	pulumi.Input

	ToAliasTypeResponseOutput() AliasTypeResponseOutput
	ToAliasTypeResponseOutputWithContext(context.Context) AliasTypeResponseOutput
}

type AliasTypeResponseArgs struct {
	Name  pulumi.StringPtrInput           `pulumi:"name"`
	Paths AliasPathTypeResponseArrayInput `pulumi:"paths"`
}

func (AliasTypeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasTypeResponse)(nil)).Elem()
}

func (i AliasTypeResponseArgs) ToAliasTypeResponseOutput() AliasTypeResponseOutput {
	return i.ToAliasTypeResponseOutputWithContext(context.Background())
}

func (i AliasTypeResponseArgs) ToAliasTypeResponseOutputWithContext(ctx context.Context) AliasTypeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasTypeResponseOutput)
}





type AliasTypeResponseArrayInput interface {
	pulumi.Input

	ToAliasTypeResponseArrayOutput() AliasTypeResponseArrayOutput
	ToAliasTypeResponseArrayOutputWithContext(context.Context) AliasTypeResponseArrayOutput
}

type AliasTypeResponseArray []AliasTypeResponseInput

func (AliasTypeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AliasTypeResponse)(nil)).Elem()
}

func (i AliasTypeResponseArray) ToAliasTypeResponseArrayOutput() AliasTypeResponseArrayOutput {
	return i.ToAliasTypeResponseArrayOutputWithContext(context.Background())
}

func (i AliasTypeResponseArray) ToAliasTypeResponseArrayOutputWithContext(ctx context.Context) AliasTypeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasTypeResponseArrayOutput)
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





type BasicDependencyResponseInput interface {
	pulumi.Input

	ToBasicDependencyResponseOutput() BasicDependencyResponseOutput
	ToBasicDependencyResponseOutputWithContext(context.Context) BasicDependencyResponseOutput
}

type BasicDependencyResponseArgs struct {
	Id           pulumi.StringPtrInput `pulumi:"id"`
	ResourceName pulumi.StringPtrInput `pulumi:"resourceName"`
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
}

func (BasicDependencyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BasicDependencyResponse)(nil)).Elem()
}

func (i BasicDependencyResponseArgs) ToBasicDependencyResponseOutput() BasicDependencyResponseOutput {
	return i.ToBasicDependencyResponseOutputWithContext(context.Background())
}

func (i BasicDependencyResponseArgs) ToBasicDependencyResponseOutputWithContext(ctx context.Context) BasicDependencyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicDependencyResponseOutput)
}





type BasicDependencyResponseArrayInput interface {
	pulumi.Input

	ToBasicDependencyResponseArrayOutput() BasicDependencyResponseArrayOutput
	ToBasicDependencyResponseArrayOutputWithContext(context.Context) BasicDependencyResponseArrayOutput
}

type BasicDependencyResponseArray []BasicDependencyResponseInput

func (BasicDependencyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BasicDependencyResponse)(nil)).Elem()
}

func (i BasicDependencyResponseArray) ToBasicDependencyResponseArrayOutput() BasicDependencyResponseArrayOutput {
	return i.ToBasicDependencyResponseArrayOutputWithContext(context.Background())
}

func (i BasicDependencyResponseArray) ToBasicDependencyResponseArrayOutputWithContext(ctx context.Context) BasicDependencyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicDependencyResponseArrayOutput)
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





type DebugSettingResponseInput interface {
	pulumi.Input

	ToDebugSettingResponseOutput() DebugSettingResponseOutput
	ToDebugSettingResponseOutputWithContext(context.Context) DebugSettingResponseOutput
}

type DebugSettingResponseArgs struct {
	DetailLevel pulumi.StringPtrInput `pulumi:"detailLevel"`
}

func (DebugSettingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DebugSettingResponse)(nil)).Elem()
}

func (i DebugSettingResponseArgs) ToDebugSettingResponseOutput() DebugSettingResponseOutput {
	return i.ToDebugSettingResponseOutputWithContext(context.Background())
}

func (i DebugSettingResponseArgs) ToDebugSettingResponseOutputWithContext(ctx context.Context) DebugSettingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebugSettingResponseOutput)
}

func (i DebugSettingResponseArgs) ToDebugSettingResponsePtrOutput() DebugSettingResponsePtrOutput {
	return i.ToDebugSettingResponsePtrOutputWithContext(context.Background())
}

func (i DebugSettingResponseArgs) ToDebugSettingResponsePtrOutputWithContext(ctx context.Context) DebugSettingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebugSettingResponseOutput).ToDebugSettingResponsePtrOutputWithContext(ctx)
}









type DebugSettingResponsePtrInput interface {
	pulumi.Input

	ToDebugSettingResponsePtrOutput() DebugSettingResponsePtrOutput
	ToDebugSettingResponsePtrOutputWithContext(context.Context) DebugSettingResponsePtrOutput
}

type debugSettingResponsePtrType DebugSettingResponseArgs

func DebugSettingResponsePtr(v *DebugSettingResponseArgs) DebugSettingResponsePtrInput {
	return (*debugSettingResponsePtrType)(v)
}

func (*debugSettingResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DebugSettingResponse)(nil)).Elem()
}

func (i *debugSettingResponsePtrType) ToDebugSettingResponsePtrOutput() DebugSettingResponsePtrOutput {
	return i.ToDebugSettingResponsePtrOutputWithContext(context.Background())
}

func (i *debugSettingResponsePtrType) ToDebugSettingResponsePtrOutputWithContext(ctx context.Context) DebugSettingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DebugSettingResponsePtrOutput)
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

func (o DebugSettingResponseOutput) ToDebugSettingResponsePtrOutput() DebugSettingResponsePtrOutput {
	return o.ToDebugSettingResponsePtrOutputWithContext(context.Background())
}

func (o DebugSettingResponseOutput) ToDebugSettingResponsePtrOutputWithContext(ctx context.Context) DebugSettingResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DebugSettingResponse) *DebugSettingResponse {
		return &v
	}).(DebugSettingResponsePtrOutput)
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





type DependencyResponseInput interface {
	pulumi.Input

	ToDependencyResponseOutput() DependencyResponseOutput
	ToDependencyResponseOutputWithContext(context.Context) DependencyResponseOutput
}

type DependencyResponseArgs struct {
	DependsOn    BasicDependencyResponseArrayInput `pulumi:"dependsOn"`
	Id           pulumi.StringPtrInput             `pulumi:"id"`
	ResourceName pulumi.StringPtrInput             `pulumi:"resourceName"`
	ResourceType pulumi.StringPtrInput             `pulumi:"resourceType"`
}

func (DependencyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DependencyResponse)(nil)).Elem()
}

func (i DependencyResponseArgs) ToDependencyResponseOutput() DependencyResponseOutput {
	return i.ToDependencyResponseOutputWithContext(context.Background())
}

func (i DependencyResponseArgs) ToDependencyResponseOutputWithContext(ctx context.Context) DependencyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependencyResponseOutput)
}





type DependencyResponseArrayInput interface {
	pulumi.Input

	ToDependencyResponseArrayOutput() DependencyResponseArrayOutput
	ToDependencyResponseArrayOutputWithContext(context.Context) DependencyResponseArrayOutput
}

type DependencyResponseArray []DependencyResponseInput

func (DependencyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DependencyResponse)(nil)).Elem()
}

func (i DependencyResponseArray) ToDependencyResponseArrayOutput() DependencyResponseArrayOutput {
	return i.ToDependencyResponseArrayOutputWithContext(context.Background())
}

func (i DependencyResponseArray) ToDependencyResponseArrayOutputWithContext(ctx context.Context) DependencyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependencyResponseArrayOutput)
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
	DebugSetting   *DebugSetting   `pulumi:"debugSetting"`
	Mode           DeploymentMode  `pulumi:"mode"`
	Parameters     interface{}     `pulumi:"parameters"`
	ParametersLink *ParametersLink `pulumi:"parametersLink"`
	Template       interface{}     `pulumi:"template"`
	TemplateLink   *TemplateLink   `pulumi:"templateLink"`
}





type DeploymentPropertiesInput interface {
	pulumi.Input

	ToDeploymentPropertiesOutput() DeploymentPropertiesOutput
	ToDeploymentPropertiesOutputWithContext(context.Context) DeploymentPropertiesOutput
}

type DeploymentPropertiesArgs struct {
	DebugSetting   DebugSettingPtrInput   `pulumi:"debugSetting"`
	Mode           DeploymentModeInput    `pulumi:"mode"`
	Parameters     pulumi.Input           `pulumi:"parameters"`
	ParametersLink ParametersLinkPtrInput `pulumi:"parametersLink"`
	Template       pulumi.Input           `pulumi:"template"`
	TemplateLink   TemplateLinkPtrInput   `pulumi:"templateLink"`
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

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return i.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesOutput).ToDeploymentPropertiesPtrOutputWithContext(ctx)
}









type DeploymentPropertiesPtrInput interface {
	pulumi.Input

	ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput
	ToDeploymentPropertiesPtrOutputWithContext(context.Context) DeploymentPropertiesPtrOutput
}

type deploymentPropertiesPtrType DeploymentPropertiesArgs

func DeploymentPropertiesPtr(v *DeploymentPropertiesArgs) DeploymentPropertiesPtrInput {
	return (*deploymentPropertiesPtrType)(v)
}

func (*deploymentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentProperties)(nil)).Elem()
}

func (i *deploymentPropertiesPtrType) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return i.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (i *deploymentPropertiesPtrType) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesPtrOutput)
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

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return o.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentProperties) *DeploymentProperties {
		return &v
	}).(DeploymentPropertiesPtrOutput)
}

func (o DeploymentPropertiesOutput) DebugSetting() DebugSettingPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *DebugSetting { return v.DebugSetting }).(DebugSettingPtrOutput)
}

func (o DeploymentPropertiesOutput) Mode() DeploymentModeOutput {
	return o.ApplyT(func(v DeploymentProperties) DeploymentMode { return v.Mode }).(DeploymentModeOutput)
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

type DeploymentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentProperties)(nil)).Elem()
}

func (o DeploymentPropertiesPtrOutput) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return o
}

func (o DeploymentPropertiesPtrOutput) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return o
}

func (o DeploymentPropertiesPtrOutput) Elem() DeploymentPropertiesOutput {
	return o.ApplyT(func(v *DeploymentProperties) DeploymentProperties {
		if v != nil {
			return *v
		}
		var ret DeploymentProperties
		return ret
	}).(DeploymentPropertiesOutput)
}

func (o DeploymentPropertiesPtrOutput) DebugSetting() DebugSettingPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *DebugSetting {
		if v == nil {
			return nil
		}
		return v.DebugSetting
	}).(DebugSettingPtrOutput)
}

func (o DeploymentPropertiesPtrOutput) Mode() DeploymentModePtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *DeploymentMode {
		if v == nil {
			return nil
		}
		return &v.Mode
	}).(DeploymentModePtrOutput)
}

func (o DeploymentPropertiesPtrOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesPtrOutput) ParametersLink() ParametersLinkPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *ParametersLink {
		if v == nil {
			return nil
		}
		return v.ParametersLink
	}).(ParametersLinkPtrOutput)
}

func (o DeploymentPropertiesPtrOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Template
	}).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesPtrOutput) TemplateLink() TemplateLinkPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *TemplateLink {
		if v == nil {
			return nil
		}
		return v.TemplateLink
	}).(TemplateLinkPtrOutput)
}

type DeploymentPropertiesExtendedResponse struct {
	CorrelationId     string                  `pulumi:"correlationId"`
	DebugSetting      *DebugSettingResponse   `pulumi:"debugSetting"`
	Dependencies      []DependencyResponse    `pulumi:"dependencies"`
	Mode              *string                 `pulumi:"mode"`
	Outputs           interface{}             `pulumi:"outputs"`
	Parameters        interface{}             `pulumi:"parameters"`
	ParametersLink    *ParametersLinkResponse `pulumi:"parametersLink"`
	Providers         []ProviderResponse      `pulumi:"providers"`
	ProvisioningState string                  `pulumi:"provisioningState"`
	Template          interface{}             `pulumi:"template"`
	TemplateLink      *TemplateLinkResponse   `pulumi:"templateLink"`
	Timestamp         string                  `pulumi:"timestamp"`
}





type DeploymentPropertiesExtendedResponseInput interface {
	pulumi.Input

	ToDeploymentPropertiesExtendedResponseOutput() DeploymentPropertiesExtendedResponseOutput
	ToDeploymentPropertiesExtendedResponseOutputWithContext(context.Context) DeploymentPropertiesExtendedResponseOutput
}

type DeploymentPropertiesExtendedResponseArgs struct {
	CorrelationId     pulumi.StringInput             `pulumi:"correlationId"`
	DebugSetting      DebugSettingResponsePtrInput   `pulumi:"debugSetting"`
	Dependencies      DependencyResponseArrayInput   `pulumi:"dependencies"`
	Mode              pulumi.StringPtrInput          `pulumi:"mode"`
	Outputs           pulumi.Input                   `pulumi:"outputs"`
	Parameters        pulumi.Input                   `pulumi:"parameters"`
	ParametersLink    ParametersLinkResponsePtrInput `pulumi:"parametersLink"`
	Providers         ProviderResponseArrayInput     `pulumi:"providers"`
	ProvisioningState pulumi.StringInput             `pulumi:"provisioningState"`
	Template          pulumi.Input                   `pulumi:"template"`
	TemplateLink      TemplateLinkResponsePtrInput   `pulumi:"templateLink"`
	Timestamp         pulumi.StringInput             `pulumi:"timestamp"`
}

func (DeploymentPropertiesExtendedResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentPropertiesExtendedResponse)(nil)).Elem()
}

func (i DeploymentPropertiesExtendedResponseArgs) ToDeploymentPropertiesExtendedResponseOutput() DeploymentPropertiesExtendedResponseOutput {
	return i.ToDeploymentPropertiesExtendedResponseOutputWithContext(context.Background())
}

func (i DeploymentPropertiesExtendedResponseArgs) ToDeploymentPropertiesExtendedResponseOutputWithContext(ctx context.Context) DeploymentPropertiesExtendedResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesExtendedResponseOutput)
}

func (i DeploymentPropertiesExtendedResponseArgs) ToDeploymentPropertiesExtendedResponsePtrOutput() DeploymentPropertiesExtendedResponsePtrOutput {
	return i.ToDeploymentPropertiesExtendedResponsePtrOutputWithContext(context.Background())
}

func (i DeploymentPropertiesExtendedResponseArgs) ToDeploymentPropertiesExtendedResponsePtrOutputWithContext(ctx context.Context) DeploymentPropertiesExtendedResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesExtendedResponseOutput).ToDeploymentPropertiesExtendedResponsePtrOutputWithContext(ctx)
}









type DeploymentPropertiesExtendedResponsePtrInput interface {
	pulumi.Input

	ToDeploymentPropertiesExtendedResponsePtrOutput() DeploymentPropertiesExtendedResponsePtrOutput
	ToDeploymentPropertiesExtendedResponsePtrOutputWithContext(context.Context) DeploymentPropertiesExtendedResponsePtrOutput
}

type deploymentPropertiesExtendedResponsePtrType DeploymentPropertiesExtendedResponseArgs

func DeploymentPropertiesExtendedResponsePtr(v *DeploymentPropertiesExtendedResponseArgs) DeploymentPropertiesExtendedResponsePtrInput {
	return (*deploymentPropertiesExtendedResponsePtrType)(v)
}

func (*deploymentPropertiesExtendedResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentPropertiesExtendedResponse)(nil)).Elem()
}

func (i *deploymentPropertiesExtendedResponsePtrType) ToDeploymentPropertiesExtendedResponsePtrOutput() DeploymentPropertiesExtendedResponsePtrOutput {
	return i.ToDeploymentPropertiesExtendedResponsePtrOutputWithContext(context.Background())
}

func (i *deploymentPropertiesExtendedResponsePtrType) ToDeploymentPropertiesExtendedResponsePtrOutputWithContext(ctx context.Context) DeploymentPropertiesExtendedResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesExtendedResponsePtrOutput)
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

func (o DeploymentPropertiesExtendedResponseOutput) ToDeploymentPropertiesExtendedResponsePtrOutput() DeploymentPropertiesExtendedResponsePtrOutput {
	return o.ToDeploymentPropertiesExtendedResponsePtrOutputWithContext(context.Background())
}

func (o DeploymentPropertiesExtendedResponseOutput) ToDeploymentPropertiesExtendedResponsePtrOutputWithContext(ctx context.Context) DeploymentPropertiesExtendedResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentPropertiesExtendedResponse) *DeploymentPropertiesExtendedResponse {
		return &v
	}).(DeploymentPropertiesExtendedResponsePtrOutput)
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

func (o DeploymentPropertiesExtendedResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
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

type DeploymentPropertiesExtendedResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesExtendedResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentPropertiesExtendedResponse)(nil)).Elem()
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) ToDeploymentPropertiesExtendedResponsePtrOutput() DeploymentPropertiesExtendedResponsePtrOutput {
	return o
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) ToDeploymentPropertiesExtendedResponsePtrOutputWithContext(ctx context.Context) DeploymentPropertiesExtendedResponsePtrOutput {
	return o
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) Elem() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) DeploymentPropertiesExtendedResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentPropertiesExtendedResponse
		return ret
	}).(DeploymentPropertiesExtendedResponseOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) DebugSetting() DebugSettingResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *DebugSettingResponse {
		if v == nil {
			return nil
		}
		return v.DebugSetting
	}).(DebugSettingResponsePtrOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) Dependencies() DependencyResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) []DependencyResponse {
		if v == nil {
			return nil
		}
		return v.Dependencies
	}).(DependencyResponseArrayOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Outputs
	}).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) ParametersLink() ParametersLinkResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *ParametersLinkResponse {
		if v == nil {
			return nil
		}
		return v.ParametersLink
	}).(ParametersLinkResponsePtrOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) Providers() ProviderResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) []ProviderResponse {
		if v == nil {
			return nil
		}
		return v.Providers
	}).(ProviderResponseArrayOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Template
	}).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) TemplateLink() TemplateLinkResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *TemplateLinkResponse {
		if v == nil {
			return nil
		}
		return v.TemplateLink
	}).(TemplateLinkResponsePtrOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Timestamp
	}).(pulumi.StringPtrOutput)
}

type Identity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
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

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
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

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
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





type ParametersLinkResponseInput interface {
	pulumi.Input

	ToParametersLinkResponseOutput() ParametersLinkResponseOutput
	ToParametersLinkResponseOutputWithContext(context.Context) ParametersLinkResponseOutput
}

type ParametersLinkResponseArgs struct {
	ContentVersion pulumi.StringPtrInput `pulumi:"contentVersion"`
	Uri            pulumi.StringInput    `pulumi:"uri"`
}

func (ParametersLinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParametersLinkResponse)(nil)).Elem()
}

func (i ParametersLinkResponseArgs) ToParametersLinkResponseOutput() ParametersLinkResponseOutput {
	return i.ToParametersLinkResponseOutputWithContext(context.Background())
}

func (i ParametersLinkResponseArgs) ToParametersLinkResponseOutputWithContext(ctx context.Context) ParametersLinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersLinkResponseOutput)
}

func (i ParametersLinkResponseArgs) ToParametersLinkResponsePtrOutput() ParametersLinkResponsePtrOutput {
	return i.ToParametersLinkResponsePtrOutputWithContext(context.Background())
}

func (i ParametersLinkResponseArgs) ToParametersLinkResponsePtrOutputWithContext(ctx context.Context) ParametersLinkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersLinkResponseOutput).ToParametersLinkResponsePtrOutputWithContext(ctx)
}









type ParametersLinkResponsePtrInput interface {
	pulumi.Input

	ToParametersLinkResponsePtrOutput() ParametersLinkResponsePtrOutput
	ToParametersLinkResponsePtrOutputWithContext(context.Context) ParametersLinkResponsePtrOutput
}

type parametersLinkResponsePtrType ParametersLinkResponseArgs

func ParametersLinkResponsePtr(v *ParametersLinkResponseArgs) ParametersLinkResponsePtrInput {
	return (*parametersLinkResponsePtrType)(v)
}

func (*parametersLinkResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ParametersLinkResponse)(nil)).Elem()
}

func (i *parametersLinkResponsePtrType) ToParametersLinkResponsePtrOutput() ParametersLinkResponsePtrOutput {
	return i.ToParametersLinkResponsePtrOutputWithContext(context.Background())
}

func (i *parametersLinkResponsePtrType) ToParametersLinkResponsePtrOutputWithContext(ctx context.Context) ParametersLinkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersLinkResponsePtrOutput)
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

func (o ParametersLinkResponseOutput) ToParametersLinkResponsePtrOutput() ParametersLinkResponsePtrOutput {
	return o.ToParametersLinkResponsePtrOutputWithContext(context.Background())
}

func (o ParametersLinkResponseOutput) ToParametersLinkResponsePtrOutputWithContext(ctx context.Context) ParametersLinkResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParametersLinkResponse) *ParametersLinkResponse {
		return &v
	}).(ParametersLinkResponsePtrOutput)
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

type PlanResponse struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
}





type PlanResponseInput interface {
	pulumi.Input

	ToPlanResponseOutput() PlanResponseOutput
	ToPlanResponseOutputWithContext(context.Context) PlanResponseOutput
}

type PlanResponseArgs struct {
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Product       pulumi.StringPtrInput `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringPtrInput `pulumi:"publisher"`
}

func (PlanResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (i PlanResponseArgs) ToPlanResponseOutput() PlanResponseOutput {
	return i.ToPlanResponseOutputWithContext(context.Background())
}

func (i PlanResponseArgs) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponseOutput)
}

func (i PlanResponseArgs) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return i.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (i PlanResponseArgs) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponseOutput).ToPlanResponsePtrOutputWithContext(ctx)
}









type PlanResponsePtrInput interface {
	pulumi.Input

	ToPlanResponsePtrOutput() PlanResponsePtrOutput
	ToPlanResponsePtrOutputWithContext(context.Context) PlanResponsePtrOutput
}

type planResponsePtrType PlanResponseArgs

func PlanResponsePtr(v *PlanResponseArgs) PlanResponsePtrInput {
	return (*planResponsePtrType)(v)
}

func (*planResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (i *planResponsePtrType) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return i.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (i *planResponsePtrType) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponsePtrOutput)
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

func (o PlanResponseOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (o PlanResponseOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlanResponse) *PlanResponse {
		return &v
	}).(PlanResponsePtrOutput)
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

type ProviderResourceTypeResponse struct {
	Aliases      []AliasTypeResponse   `pulumi:"aliases"`
	ApiVersions  []string              `pulumi:"apiVersions"`
	Locations    []string              `pulumi:"locations"`
	Properties   map[string]string     `pulumi:"properties"`
	ResourceType *string               `pulumi:"resourceType"`
	ZoneMappings []ZoneMappingResponse `pulumi:"zoneMappings"`
}





type ProviderResourceTypeResponseInput interface {
	pulumi.Input

	ToProviderResourceTypeResponseOutput() ProviderResourceTypeResponseOutput
	ToProviderResourceTypeResponseOutputWithContext(context.Context) ProviderResourceTypeResponseOutput
}

type ProviderResourceTypeResponseArgs struct {
	Aliases      AliasTypeResponseArrayInput   `pulumi:"aliases"`
	ApiVersions  pulumi.StringArrayInput       `pulumi:"apiVersions"`
	Locations    pulumi.StringArrayInput       `pulumi:"locations"`
	Properties   pulumi.StringMapInput         `pulumi:"properties"`
	ResourceType pulumi.StringPtrInput         `pulumi:"resourceType"`
	ZoneMappings ZoneMappingResponseArrayInput `pulumi:"zoneMappings"`
}

func (ProviderResourceTypeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderResourceTypeResponse)(nil)).Elem()
}

func (i ProviderResourceTypeResponseArgs) ToProviderResourceTypeResponseOutput() ProviderResourceTypeResponseOutput {
	return i.ToProviderResourceTypeResponseOutputWithContext(context.Background())
}

func (i ProviderResourceTypeResponseArgs) ToProviderResourceTypeResponseOutputWithContext(ctx context.Context) ProviderResourceTypeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderResourceTypeResponseOutput)
}





type ProviderResourceTypeResponseArrayInput interface {
	pulumi.Input

	ToProviderResourceTypeResponseArrayOutput() ProviderResourceTypeResponseArrayOutput
	ToProviderResourceTypeResponseArrayOutputWithContext(context.Context) ProviderResourceTypeResponseArrayOutput
}

type ProviderResourceTypeResponseArray []ProviderResourceTypeResponseInput

func (ProviderResourceTypeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderResourceTypeResponse)(nil)).Elem()
}

func (i ProviderResourceTypeResponseArray) ToProviderResourceTypeResponseArrayOutput() ProviderResourceTypeResponseArrayOutput {
	return i.ToProviderResourceTypeResponseArrayOutputWithContext(context.Background())
}

func (i ProviderResourceTypeResponseArray) ToProviderResourceTypeResponseArrayOutputWithContext(ctx context.Context) ProviderResourceTypeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderResourceTypeResponseArrayOutput)
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
	Id                string                         `pulumi:"id"`
	Namespace         *string                        `pulumi:"namespace"`
	RegistrationState string                         `pulumi:"registrationState"`
	ResourceTypes     []ProviderResourceTypeResponse `pulumi:"resourceTypes"`
}





type ProviderResponseInput interface {
	pulumi.Input

	ToProviderResponseOutput() ProviderResponseOutput
	ToProviderResponseOutputWithContext(context.Context) ProviderResponseOutput
}

type ProviderResponseArgs struct {
	Id                pulumi.StringInput                     `pulumi:"id"`
	Namespace         pulumi.StringPtrInput                  `pulumi:"namespace"`
	RegistrationState pulumi.StringInput                     `pulumi:"registrationState"`
	ResourceTypes     ProviderResourceTypeResponseArrayInput `pulumi:"resourceTypes"`
}

func (ProviderResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderResponse)(nil)).Elem()
}

func (i ProviderResponseArgs) ToProviderResponseOutput() ProviderResponseOutput {
	return i.ToProviderResponseOutputWithContext(context.Background())
}

func (i ProviderResponseArgs) ToProviderResponseOutputWithContext(ctx context.Context) ProviderResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderResponseOutput)
}





type ProviderResponseArrayInput interface {
	pulumi.Input

	ToProviderResponseArrayOutput() ProviderResponseArrayOutput
	ToProviderResponseArrayOutputWithContext(context.Context) ProviderResponseArrayOutput
}

type ProviderResponseArray []ProviderResponseInput

func (ProviderResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderResponse)(nil)).Elem()
}

func (i ProviderResponseArray) ToProviderResponseArrayOutput() ProviderResponseArrayOutput {
	return i.ToProviderResponseArrayOutputWithContext(context.Background())
}

func (i ProviderResponseArray) ToProviderResponseArrayOutputWithContext(ctx context.Context) ProviderResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderResponseArrayOutput)
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





type ResourceGroupPropertiesResponseInput interface {
	pulumi.Input

	ToResourceGroupPropertiesResponseOutput() ResourceGroupPropertiesResponseOutput
	ToResourceGroupPropertiesResponseOutputWithContext(context.Context) ResourceGroupPropertiesResponseOutput
}

type ResourceGroupPropertiesResponseArgs struct {
	ProvisioningState pulumi.StringInput `pulumi:"provisioningState"`
}

func (ResourceGroupPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupPropertiesResponse)(nil)).Elem()
}

func (i ResourceGroupPropertiesResponseArgs) ToResourceGroupPropertiesResponseOutput() ResourceGroupPropertiesResponseOutput {
	return i.ToResourceGroupPropertiesResponseOutputWithContext(context.Background())
}

func (i ResourceGroupPropertiesResponseArgs) ToResourceGroupPropertiesResponseOutputWithContext(ctx context.Context) ResourceGroupPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupPropertiesResponseOutput)
}

func (i ResourceGroupPropertiesResponseArgs) ToResourceGroupPropertiesResponsePtrOutput() ResourceGroupPropertiesResponsePtrOutput {
	return i.ToResourceGroupPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ResourceGroupPropertiesResponseArgs) ToResourceGroupPropertiesResponsePtrOutputWithContext(ctx context.Context) ResourceGroupPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupPropertiesResponseOutput).ToResourceGroupPropertiesResponsePtrOutputWithContext(ctx)
}









type ResourceGroupPropertiesResponsePtrInput interface {
	pulumi.Input

	ToResourceGroupPropertiesResponsePtrOutput() ResourceGroupPropertiesResponsePtrOutput
	ToResourceGroupPropertiesResponsePtrOutputWithContext(context.Context) ResourceGroupPropertiesResponsePtrOutput
}

type resourceGroupPropertiesResponsePtrType ResourceGroupPropertiesResponseArgs

func ResourceGroupPropertiesResponsePtr(v *ResourceGroupPropertiesResponseArgs) ResourceGroupPropertiesResponsePtrInput {
	return (*resourceGroupPropertiesResponsePtrType)(v)
}

func (*resourceGroupPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroupPropertiesResponse)(nil)).Elem()
}

func (i *resourceGroupPropertiesResponsePtrType) ToResourceGroupPropertiesResponsePtrOutput() ResourceGroupPropertiesResponsePtrOutput {
	return i.ToResourceGroupPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *resourceGroupPropertiesResponsePtrType) ToResourceGroupPropertiesResponsePtrOutputWithContext(ctx context.Context) ResourceGroupPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupPropertiesResponsePtrOutput)
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

func (o ResourceGroupPropertiesResponseOutput) ToResourceGroupPropertiesResponsePtrOutput() ResourceGroupPropertiesResponsePtrOutput {
	return o.ToResourceGroupPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ResourceGroupPropertiesResponseOutput) ToResourceGroupPropertiesResponsePtrOutputWithContext(ctx context.Context) ResourceGroupPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceGroupPropertiesResponse) *ResourceGroupPropertiesResponse {
		return &v
	}).(ResourceGroupPropertiesResponsePtrOutput)
}

func (o ResourceGroupPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceGroupPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ResourceGroupPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceGroupPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroupPropertiesResponse)(nil)).Elem()
}

func (o ResourceGroupPropertiesResponsePtrOutput) ToResourceGroupPropertiesResponsePtrOutput() ResourceGroupPropertiesResponsePtrOutput {
	return o
}

func (o ResourceGroupPropertiesResponsePtrOutput) ToResourceGroupPropertiesResponsePtrOutputWithContext(ctx context.Context) ResourceGroupPropertiesResponsePtrOutput {
	return o
}

func (o ResourceGroupPropertiesResponsePtrOutput) Elem() ResourceGroupPropertiesResponseOutput {
	return o.ApplyT(func(v *ResourceGroupPropertiesResponse) ResourceGroupPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ResourceGroupPropertiesResponse
		return ret
	}).(ResourceGroupPropertiesResponseOutput)
}

func (o ResourceGroupPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGroupPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
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





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Model    pulumi.StringPtrInput `pulumi:"model"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
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





type TemplateLinkResponseInput interface {
	pulumi.Input

	ToTemplateLinkResponseOutput() TemplateLinkResponseOutput
	ToTemplateLinkResponseOutputWithContext(context.Context) TemplateLinkResponseOutput
}

type TemplateLinkResponseArgs struct {
	ContentVersion pulumi.StringPtrInput `pulumi:"contentVersion"`
	Uri            pulumi.StringInput    `pulumi:"uri"`
}

func (TemplateLinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateLinkResponse)(nil)).Elem()
}

func (i TemplateLinkResponseArgs) ToTemplateLinkResponseOutput() TemplateLinkResponseOutput {
	return i.ToTemplateLinkResponseOutputWithContext(context.Background())
}

func (i TemplateLinkResponseArgs) ToTemplateLinkResponseOutputWithContext(ctx context.Context) TemplateLinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateLinkResponseOutput)
}

func (i TemplateLinkResponseArgs) ToTemplateLinkResponsePtrOutput() TemplateLinkResponsePtrOutput {
	return i.ToTemplateLinkResponsePtrOutputWithContext(context.Background())
}

func (i TemplateLinkResponseArgs) ToTemplateLinkResponsePtrOutputWithContext(ctx context.Context) TemplateLinkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateLinkResponseOutput).ToTemplateLinkResponsePtrOutputWithContext(ctx)
}









type TemplateLinkResponsePtrInput interface {
	pulumi.Input

	ToTemplateLinkResponsePtrOutput() TemplateLinkResponsePtrOutput
	ToTemplateLinkResponsePtrOutputWithContext(context.Context) TemplateLinkResponsePtrOutput
}

type templateLinkResponsePtrType TemplateLinkResponseArgs

func TemplateLinkResponsePtr(v *TemplateLinkResponseArgs) TemplateLinkResponsePtrInput {
	return (*templateLinkResponsePtrType)(v)
}

func (*templateLinkResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateLinkResponse)(nil)).Elem()
}

func (i *templateLinkResponsePtrType) ToTemplateLinkResponsePtrOutput() TemplateLinkResponsePtrOutput {
	return i.ToTemplateLinkResponsePtrOutputWithContext(context.Background())
}

func (i *templateLinkResponsePtrType) ToTemplateLinkResponsePtrOutputWithContext(ctx context.Context) TemplateLinkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateLinkResponsePtrOutput)
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

func (o TemplateLinkResponseOutput) ToTemplateLinkResponsePtrOutput() TemplateLinkResponsePtrOutput {
	return o.ToTemplateLinkResponsePtrOutputWithContext(context.Background())
}

func (o TemplateLinkResponseOutput) ToTemplateLinkResponsePtrOutputWithContext(ctx context.Context) TemplateLinkResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TemplateLinkResponse) *TemplateLinkResponse {
		return &v
	}).(TemplateLinkResponsePtrOutput)
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





type ZoneMappingResponseInput interface {
	pulumi.Input

	ToZoneMappingResponseOutput() ZoneMappingResponseOutput
	ToZoneMappingResponseOutputWithContext(context.Context) ZoneMappingResponseOutput
}

type ZoneMappingResponseArgs struct {
	Location pulumi.StringPtrInput   `pulumi:"location"`
	Zones    pulumi.StringArrayInput `pulumi:"zones"`
}

func (ZoneMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneMappingResponse)(nil)).Elem()
}

func (i ZoneMappingResponseArgs) ToZoneMappingResponseOutput() ZoneMappingResponseOutput {
	return i.ToZoneMappingResponseOutputWithContext(context.Background())
}

func (i ZoneMappingResponseArgs) ToZoneMappingResponseOutputWithContext(ctx context.Context) ZoneMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneMappingResponseOutput)
}





type ZoneMappingResponseArrayInput interface {
	pulumi.Input

	ToZoneMappingResponseArrayOutput() ZoneMappingResponseArrayOutput
	ToZoneMappingResponseArrayOutputWithContext(context.Context) ZoneMappingResponseArrayOutput
}

type ZoneMappingResponseArray []ZoneMappingResponseInput

func (ZoneMappingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneMappingResponse)(nil)).Elem()
}

func (i ZoneMappingResponseArray) ToZoneMappingResponseArrayOutput() ZoneMappingResponseArrayOutput {
	return i.ToZoneMappingResponseArrayOutputWithContext(context.Background())
}

func (i ZoneMappingResponseArray) ToZoneMappingResponseArrayOutputWithContext(ctx context.Context) ZoneMappingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneMappingResponseArrayOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*AliasPathTypeResponseInput)(nil)).Elem(), AliasPathTypeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasPathTypeResponseArrayInput)(nil)).Elem(), AliasPathTypeResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasTypeResponseInput)(nil)).Elem(), AliasTypeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasTypeResponseArrayInput)(nil)).Elem(), AliasTypeResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicDependencyResponseInput)(nil)).Elem(), BasicDependencyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicDependencyResponseArrayInput)(nil)).Elem(), BasicDependencyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DebugSettingInput)(nil)).Elem(), DebugSettingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DebugSettingPtrInput)(nil)).Elem(), DebugSettingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DebugSettingResponseInput)(nil)).Elem(), DebugSettingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DebugSettingResponsePtrInput)(nil)).Elem(), DebugSettingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DependencyResponseInput)(nil)).Elem(), DependencyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DependencyResponseArrayInput)(nil)).Elem(), DependencyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentPropertiesInput)(nil)).Elem(), DeploymentPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentPropertiesPtrInput)(nil)).Elem(), DeploymentPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentPropertiesExtendedResponseInput)(nil)).Elem(), DeploymentPropertiesExtendedResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentPropertiesExtendedResponsePtrInput)(nil)).Elem(), DeploymentPropertiesExtendedResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPtrInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponseInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponsePtrInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersLinkInput)(nil)).Elem(), ParametersLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersLinkPtrInput)(nil)).Elem(), ParametersLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersLinkResponseInput)(nil)).Elem(), ParametersLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersLinkResponsePtrInput)(nil)).Elem(), ParametersLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanInput)(nil)).Elem(), PlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanPtrInput)(nil)).Elem(), PlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanResponseInput)(nil)).Elem(), PlanResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanResponsePtrInput)(nil)).Elem(), PlanResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderResourceTypeResponseInput)(nil)).Elem(), ProviderResourceTypeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderResourceTypeResponseArrayInput)(nil)).Elem(), ProviderResourceTypeResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderResponseInput)(nil)).Elem(), ProviderResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderResponseArrayInput)(nil)).Elem(), ProviderResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupPropertiesResponseInput)(nil)).Elem(), ResourceGroupPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupPropertiesResponsePtrInput)(nil)).Elem(), ResourceGroupPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateLinkInput)(nil)).Elem(), TemplateLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateLinkPtrInput)(nil)).Elem(), TemplateLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateLinkResponseInput)(nil)).Elem(), TemplateLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateLinkResponsePtrInput)(nil)).Elem(), TemplateLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneMappingResponseInput)(nil)).Elem(), ZoneMappingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneMappingResponseArrayInput)(nil)).Elem(), ZoneMappingResponseArray{})
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
	pulumi.RegisterOutputType(DeploymentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesExtendedResponseOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesExtendedResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ResourceGroupPropertiesResponsePtrOutput{})
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
