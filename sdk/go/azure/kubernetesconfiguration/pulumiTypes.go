


package kubernetesconfiguration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComplianceStatusResponse struct {
	ComplianceState   string  `pulumi:"complianceState"`
	LastConfigApplied *string `pulumi:"lastConfigApplied"`
	Message           *string `pulumi:"message"`
	MessageLevel      *string `pulumi:"messageLevel"`
}

type ComplianceStatusResponseOutput struct{ *pulumi.OutputState }

func (ComplianceStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComplianceStatusResponse)(nil)).Elem()
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponseOutput() ComplianceStatusResponseOutput {
	return o
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponseOutputWithContext(ctx context.Context) ComplianceStatusResponseOutput {
	return o
}

func (o ComplianceStatusResponseOutput) ComplianceState() pulumi.StringOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) string { return v.ComplianceState }).(pulumi.StringOutput)
}

func (o ComplianceStatusResponseOutput) LastConfigApplied() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.LastConfigApplied }).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponseOutput) MessageLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.MessageLevel }).(pulumi.StringPtrOutput)
}

type ConfigurationIdentity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type ConfigurationIdentityInput interface {
	pulumi.Input

	ToConfigurationIdentityOutput() ConfigurationIdentityOutput
	ToConfigurationIdentityOutputWithContext(context.Context) ConfigurationIdentityOutput
}

type ConfigurationIdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (ConfigurationIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationIdentity)(nil)).Elem()
}

func (i ConfigurationIdentityArgs) ToConfigurationIdentityOutput() ConfigurationIdentityOutput {
	return i.ToConfigurationIdentityOutputWithContext(context.Background())
}

func (i ConfigurationIdentityArgs) ToConfigurationIdentityOutputWithContext(ctx context.Context) ConfigurationIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationIdentityOutput)
}

func (i ConfigurationIdentityArgs) ToConfigurationIdentityPtrOutput() ConfigurationIdentityPtrOutput {
	return i.ToConfigurationIdentityPtrOutputWithContext(context.Background())
}

func (i ConfigurationIdentityArgs) ToConfigurationIdentityPtrOutputWithContext(ctx context.Context) ConfigurationIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationIdentityOutput).ToConfigurationIdentityPtrOutputWithContext(ctx)
}









type ConfigurationIdentityPtrInput interface {
	pulumi.Input

	ToConfigurationIdentityPtrOutput() ConfigurationIdentityPtrOutput
	ToConfigurationIdentityPtrOutputWithContext(context.Context) ConfigurationIdentityPtrOutput
}

type configurationIdentityPtrType ConfigurationIdentityArgs

func ConfigurationIdentityPtr(v *ConfigurationIdentityArgs) ConfigurationIdentityPtrInput {
	return (*configurationIdentityPtrType)(v)
}

func (*configurationIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationIdentity)(nil)).Elem()
}

func (i *configurationIdentityPtrType) ToConfigurationIdentityPtrOutput() ConfigurationIdentityPtrOutput {
	return i.ToConfigurationIdentityPtrOutputWithContext(context.Background())
}

func (i *configurationIdentityPtrType) ToConfigurationIdentityPtrOutputWithContext(ctx context.Context) ConfigurationIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationIdentityPtrOutput)
}

type ConfigurationIdentityOutput struct{ *pulumi.OutputState }

func (ConfigurationIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationIdentity)(nil)).Elem()
}

func (o ConfigurationIdentityOutput) ToConfigurationIdentityOutput() ConfigurationIdentityOutput {
	return o
}

func (o ConfigurationIdentityOutput) ToConfigurationIdentityOutputWithContext(ctx context.Context) ConfigurationIdentityOutput {
	return o
}

func (o ConfigurationIdentityOutput) ToConfigurationIdentityPtrOutput() ConfigurationIdentityPtrOutput {
	return o.ToConfigurationIdentityPtrOutputWithContext(context.Background())
}

func (o ConfigurationIdentityOutput) ToConfigurationIdentityPtrOutputWithContext(ctx context.Context) ConfigurationIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationIdentity) *ConfigurationIdentity {
		return &v
	}).(ConfigurationIdentityPtrOutput)
}

func (o ConfigurationIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ConfigurationIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type ConfigurationIdentityPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationIdentity)(nil)).Elem()
}

func (o ConfigurationIdentityPtrOutput) ToConfigurationIdentityPtrOutput() ConfigurationIdentityPtrOutput {
	return o
}

func (o ConfigurationIdentityPtrOutput) ToConfigurationIdentityPtrOutputWithContext(ctx context.Context) ConfigurationIdentityPtrOutput {
	return o
}

func (o ConfigurationIdentityPtrOutput) Elem() ConfigurationIdentityOutput {
	return o.ApplyT(func(v *ConfigurationIdentity) ConfigurationIdentity {
		if v != nil {
			return *v
		}
		var ret ConfigurationIdentity
		return ret
	}).(ConfigurationIdentityOutput)
}

func (o ConfigurationIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ConfigurationIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type ConfigurationIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ConfigurationIdentityResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationIdentityResponse)(nil)).Elem()
}

func (o ConfigurationIdentityResponseOutput) ToConfigurationIdentityResponseOutput() ConfigurationIdentityResponseOutput {
	return o
}

func (o ConfigurationIdentityResponseOutput) ToConfigurationIdentityResponseOutputWithContext(ctx context.Context) ConfigurationIdentityResponseOutput {
	return o
}

func (o ConfigurationIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ConfigurationIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ConfigurationIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConfigurationIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationIdentityResponse)(nil)).Elem()
}

func (o ConfigurationIdentityResponsePtrOutput) ToConfigurationIdentityResponsePtrOutput() ConfigurationIdentityResponsePtrOutput {
	return o
}

func (o ConfigurationIdentityResponsePtrOutput) ToConfigurationIdentityResponsePtrOutputWithContext(ctx context.Context) ConfigurationIdentityResponsePtrOutput {
	return o
}

func (o ConfigurationIdentityResponsePtrOutput) Elem() ConfigurationIdentityResponseOutput {
	return o.ApplyT(func(v *ConfigurationIdentityResponse) ConfigurationIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationIdentityResponse
		return ret
	}).(ConfigurationIdentityResponseOutput)
}

func (o ConfigurationIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type DependsOnDefinition struct {
	KustomizationName *string `pulumi:"kustomizationName"`
}





type DependsOnDefinitionInput interface {
	pulumi.Input

	ToDependsOnDefinitionOutput() DependsOnDefinitionOutput
	ToDependsOnDefinitionOutputWithContext(context.Context) DependsOnDefinitionOutput
}

type DependsOnDefinitionArgs struct {
	KustomizationName pulumi.StringPtrInput `pulumi:"kustomizationName"`
}

func (DependsOnDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DependsOnDefinition)(nil)).Elem()
}

func (i DependsOnDefinitionArgs) ToDependsOnDefinitionOutput() DependsOnDefinitionOutput {
	return i.ToDependsOnDefinitionOutputWithContext(context.Background())
}

func (i DependsOnDefinitionArgs) ToDependsOnDefinitionOutputWithContext(ctx context.Context) DependsOnDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependsOnDefinitionOutput)
}





type DependsOnDefinitionArrayInput interface {
	pulumi.Input

	ToDependsOnDefinitionArrayOutput() DependsOnDefinitionArrayOutput
	ToDependsOnDefinitionArrayOutputWithContext(context.Context) DependsOnDefinitionArrayOutput
}

type DependsOnDefinitionArray []DependsOnDefinitionInput

func (DependsOnDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DependsOnDefinition)(nil)).Elem()
}

func (i DependsOnDefinitionArray) ToDependsOnDefinitionArrayOutput() DependsOnDefinitionArrayOutput {
	return i.ToDependsOnDefinitionArrayOutputWithContext(context.Background())
}

func (i DependsOnDefinitionArray) ToDependsOnDefinitionArrayOutputWithContext(ctx context.Context) DependsOnDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependsOnDefinitionArrayOutput)
}

type DependsOnDefinitionOutput struct{ *pulumi.OutputState }

func (DependsOnDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DependsOnDefinition)(nil)).Elem()
}

func (o DependsOnDefinitionOutput) ToDependsOnDefinitionOutput() DependsOnDefinitionOutput {
	return o
}

func (o DependsOnDefinitionOutput) ToDependsOnDefinitionOutputWithContext(ctx context.Context) DependsOnDefinitionOutput {
	return o
}

func (o DependsOnDefinitionOutput) KustomizationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DependsOnDefinition) *string { return v.KustomizationName }).(pulumi.StringPtrOutput)
}

type DependsOnDefinitionArrayOutput struct{ *pulumi.OutputState }

func (DependsOnDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DependsOnDefinition)(nil)).Elem()
}

func (o DependsOnDefinitionArrayOutput) ToDependsOnDefinitionArrayOutput() DependsOnDefinitionArrayOutput {
	return o
}

func (o DependsOnDefinitionArrayOutput) ToDependsOnDefinitionArrayOutputWithContext(ctx context.Context) DependsOnDefinitionArrayOutput {
	return o
}

func (o DependsOnDefinitionArrayOutput) Index(i pulumi.IntInput) DependsOnDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DependsOnDefinition {
		return vs[0].([]DependsOnDefinition)[vs[1].(int)]
	}).(DependsOnDefinitionOutput)
}

type DependsOnDefinitionResponse struct {
	KustomizationName *string `pulumi:"kustomizationName"`
}

type DependsOnDefinitionResponseOutput struct{ *pulumi.OutputState }

func (DependsOnDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DependsOnDefinitionResponse)(nil)).Elem()
}

func (o DependsOnDefinitionResponseOutput) ToDependsOnDefinitionResponseOutput() DependsOnDefinitionResponseOutput {
	return o
}

func (o DependsOnDefinitionResponseOutput) ToDependsOnDefinitionResponseOutputWithContext(ctx context.Context) DependsOnDefinitionResponseOutput {
	return o
}

func (o DependsOnDefinitionResponseOutput) KustomizationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DependsOnDefinitionResponse) *string { return v.KustomizationName }).(pulumi.StringPtrOutput)
}

type DependsOnDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (DependsOnDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DependsOnDefinitionResponse)(nil)).Elem()
}

func (o DependsOnDefinitionResponseArrayOutput) ToDependsOnDefinitionResponseArrayOutput() DependsOnDefinitionResponseArrayOutput {
	return o
}

func (o DependsOnDefinitionResponseArrayOutput) ToDependsOnDefinitionResponseArrayOutputWithContext(ctx context.Context) DependsOnDefinitionResponseArrayOutput {
	return o
}

func (o DependsOnDefinitionResponseArrayOutput) Index(i pulumi.IntInput) DependsOnDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DependsOnDefinitionResponse {
		return vs[0].([]DependsOnDefinitionResponse)[vs[1].(int)]
	}).(DependsOnDefinitionResponseOutput)
}

type ErrorDefinitionResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}

type ErrorDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutput() ErrorDefinitionResponseOutput {
	return o
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutputWithContext(ctx context.Context) ErrorDefinitionResponseOutput {
	return o
}

func (o ErrorDefinitionResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDefinitionResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Message }).(pulumi.StringOutput)
}

type ExtensionStatus struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}


func (val *ExtensionStatus) Defaults() *ExtensionStatus {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		level_ := "Information"
		tmp.Level = &level_
	}
	return &tmp
}





type ExtensionStatusInput interface {
	pulumi.Input

	ToExtensionStatusOutput() ExtensionStatusOutput
	ToExtensionStatusOutputWithContext(context.Context) ExtensionStatusOutput
}

type ExtensionStatusArgs struct {
	Code          pulumi.StringPtrInput `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput `pulumi:"displayStatus"`
	Level         pulumi.StringPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput `pulumi:"message"`
	Time          pulumi.StringPtrInput `pulumi:"time"`
}


func (val *ExtensionStatusArgs) Defaults() *ExtensionStatusArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		tmp.Level = pulumi.StringPtr("Information")
	}
	return &tmp
}
func (ExtensionStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionStatus)(nil)).Elem()
}

func (i ExtensionStatusArgs) ToExtensionStatusOutput() ExtensionStatusOutput {
	return i.ToExtensionStatusOutputWithContext(context.Background())
}

func (i ExtensionStatusArgs) ToExtensionStatusOutputWithContext(ctx context.Context) ExtensionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionStatusOutput)
}





type ExtensionStatusArrayInput interface {
	pulumi.Input

	ToExtensionStatusArrayOutput() ExtensionStatusArrayOutput
	ToExtensionStatusArrayOutputWithContext(context.Context) ExtensionStatusArrayOutput
}

type ExtensionStatusArray []ExtensionStatusInput

func (ExtensionStatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionStatus)(nil)).Elem()
}

func (i ExtensionStatusArray) ToExtensionStatusArrayOutput() ExtensionStatusArrayOutput {
	return i.ToExtensionStatusArrayOutputWithContext(context.Background())
}

func (i ExtensionStatusArray) ToExtensionStatusArrayOutputWithContext(ctx context.Context) ExtensionStatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionStatusArrayOutput)
}

type ExtensionStatusOutput struct{ *pulumi.OutputState }

func (ExtensionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionStatus)(nil)).Elem()
}

func (o ExtensionStatusOutput) ToExtensionStatusOutput() ExtensionStatusOutput {
	return o
}

func (o ExtensionStatusOutput) ToExtensionStatusOutputWithContext(ctx context.Context) ExtensionStatusOutput {
	return o
}

func (o ExtensionStatusOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type ExtensionStatusArrayOutput struct{ *pulumi.OutputState }

func (ExtensionStatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionStatus)(nil)).Elem()
}

func (o ExtensionStatusArrayOutput) ToExtensionStatusArrayOutput() ExtensionStatusArrayOutput {
	return o
}

func (o ExtensionStatusArrayOutput) ToExtensionStatusArrayOutputWithContext(ctx context.Context) ExtensionStatusArrayOutput {
	return o
}

func (o ExtensionStatusArrayOutput) Index(i pulumi.IntInput) ExtensionStatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtensionStatus {
		return vs[0].([]ExtensionStatus)[vs[1].(int)]
	}).(ExtensionStatusOutput)
}

type ExtensionStatusResponse struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}


func (val *ExtensionStatusResponse) Defaults() *ExtensionStatusResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		level_ := "Information"
		tmp.Level = &level_
	}
	return &tmp
}

type ExtensionStatusResponseOutput struct{ *pulumi.OutputState }

func (ExtensionStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionStatusResponse)(nil)).Elem()
}

func (o ExtensionStatusResponseOutput) ToExtensionStatusResponseOutput() ExtensionStatusResponseOutput {
	return o
}

func (o ExtensionStatusResponseOutput) ToExtensionStatusResponseOutputWithContext(ctx context.Context) ExtensionStatusResponseOutput {
	return o
}

func (o ExtensionStatusResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type ExtensionStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (ExtensionStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionStatusResponse)(nil)).Elem()
}

func (o ExtensionStatusResponseArrayOutput) ToExtensionStatusResponseArrayOutput() ExtensionStatusResponseArrayOutput {
	return o
}

func (o ExtensionStatusResponseArrayOutput) ToExtensionStatusResponseArrayOutputWithContext(ctx context.Context) ExtensionStatusResponseArrayOutput {
	return o
}

func (o ExtensionStatusResponseArrayOutput) Index(i pulumi.IntInput) ExtensionStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtensionStatusResponse {
		return vs[0].([]ExtensionStatusResponse)[vs[1].(int)]
	}).(ExtensionStatusResponseOutput)
}

type GitRepositoryDefinition struct {
	HttpsCAFile           *string                  `pulumi:"httpsCAFile"`
	HttpsUser             *string                  `pulumi:"httpsUser"`
	LocalAuthRef          *string                  `pulumi:"localAuthRef"`
	RepositoryRef         *RepositoryRefDefinition `pulumi:"repositoryRef"`
	SshKnownHosts         *string                  `pulumi:"sshKnownHosts"`
	SyncIntervalInSeconds *float64                 `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      *float64                 `pulumi:"timeoutInSeconds"`
	Url                   *string                  `pulumi:"url"`
}


func (val *GitRepositoryDefinition) Defaults() *GitRepositoryDefinition {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}





type GitRepositoryDefinitionInput interface {
	pulumi.Input

	ToGitRepositoryDefinitionOutput() GitRepositoryDefinitionOutput
	ToGitRepositoryDefinitionOutputWithContext(context.Context) GitRepositoryDefinitionOutput
}

type GitRepositoryDefinitionArgs struct {
	HttpsCAFile           pulumi.StringPtrInput           `pulumi:"httpsCAFile"`
	HttpsUser             pulumi.StringPtrInput           `pulumi:"httpsUser"`
	LocalAuthRef          pulumi.StringPtrInput           `pulumi:"localAuthRef"`
	RepositoryRef         RepositoryRefDefinitionPtrInput `pulumi:"repositoryRef"`
	SshKnownHosts         pulumi.StringPtrInput           `pulumi:"sshKnownHosts"`
	SyncIntervalInSeconds pulumi.Float64PtrInput          `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      pulumi.Float64PtrInput          `pulumi:"timeoutInSeconds"`
	Url                   pulumi.StringPtrInput           `pulumi:"url"`
}


func (val *GitRepositoryDefinitionArgs) Defaults() *GitRepositoryDefinitionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SyncIntervalInSeconds) {
		tmp.SyncIntervalInSeconds = pulumi.Float64Ptr(600.0)
	}
	if isZero(tmp.TimeoutInSeconds) {
		tmp.TimeoutInSeconds = pulumi.Float64Ptr(600.0)
	}
	return &tmp
}
func (GitRepositoryDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitRepositoryDefinition)(nil)).Elem()
}

func (i GitRepositoryDefinitionArgs) ToGitRepositoryDefinitionOutput() GitRepositoryDefinitionOutput {
	return i.ToGitRepositoryDefinitionOutputWithContext(context.Background())
}

func (i GitRepositoryDefinitionArgs) ToGitRepositoryDefinitionOutputWithContext(ctx context.Context) GitRepositoryDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryDefinitionOutput)
}

func (i GitRepositoryDefinitionArgs) ToGitRepositoryDefinitionPtrOutput() GitRepositoryDefinitionPtrOutput {
	return i.ToGitRepositoryDefinitionPtrOutputWithContext(context.Background())
}

func (i GitRepositoryDefinitionArgs) ToGitRepositoryDefinitionPtrOutputWithContext(ctx context.Context) GitRepositoryDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryDefinitionOutput).ToGitRepositoryDefinitionPtrOutputWithContext(ctx)
}









type GitRepositoryDefinitionPtrInput interface {
	pulumi.Input

	ToGitRepositoryDefinitionPtrOutput() GitRepositoryDefinitionPtrOutput
	ToGitRepositoryDefinitionPtrOutputWithContext(context.Context) GitRepositoryDefinitionPtrOutput
}

type gitRepositoryDefinitionPtrType GitRepositoryDefinitionArgs

func GitRepositoryDefinitionPtr(v *GitRepositoryDefinitionArgs) GitRepositoryDefinitionPtrInput {
	return (*gitRepositoryDefinitionPtrType)(v)
}

func (*gitRepositoryDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepositoryDefinition)(nil)).Elem()
}

func (i *gitRepositoryDefinitionPtrType) ToGitRepositoryDefinitionPtrOutput() GitRepositoryDefinitionPtrOutput {
	return i.ToGitRepositoryDefinitionPtrOutputWithContext(context.Background())
}

func (i *gitRepositoryDefinitionPtrType) ToGitRepositoryDefinitionPtrOutputWithContext(ctx context.Context) GitRepositoryDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryDefinitionPtrOutput)
}

type GitRepositoryDefinitionOutput struct{ *pulumi.OutputState }

func (GitRepositoryDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitRepositoryDefinition)(nil)).Elem()
}

func (o GitRepositoryDefinitionOutput) ToGitRepositoryDefinitionOutput() GitRepositoryDefinitionOutput {
	return o
}

func (o GitRepositoryDefinitionOutput) ToGitRepositoryDefinitionOutputWithContext(ctx context.Context) GitRepositoryDefinitionOutput {
	return o
}

func (o GitRepositoryDefinitionOutput) ToGitRepositoryDefinitionPtrOutput() GitRepositoryDefinitionPtrOutput {
	return o.ToGitRepositoryDefinitionPtrOutputWithContext(context.Background())
}

func (o GitRepositoryDefinitionOutput) ToGitRepositoryDefinitionPtrOutputWithContext(ctx context.Context) GitRepositoryDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitRepositoryDefinition) *GitRepositoryDefinition {
		return &v
	}).(GitRepositoryDefinitionPtrOutput)
}

func (o GitRepositoryDefinitionOutput) HttpsCAFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *string { return v.HttpsCAFile }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionOutput) HttpsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *string { return v.HttpsUser }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *string { return v.LocalAuthRef }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionOutput) RepositoryRef() RepositoryRefDefinitionPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *RepositoryRefDefinition { return v.RepositoryRef }).(RepositoryRefDefinitionPtrOutput)
}

func (o GitRepositoryDefinitionOutput) SshKnownHosts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *string { return v.SshKnownHosts }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type GitRepositoryDefinitionPtrOutput struct{ *pulumi.OutputState }

func (GitRepositoryDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepositoryDefinition)(nil)).Elem()
}

func (o GitRepositoryDefinitionPtrOutput) ToGitRepositoryDefinitionPtrOutput() GitRepositoryDefinitionPtrOutput {
	return o
}

func (o GitRepositoryDefinitionPtrOutput) ToGitRepositoryDefinitionPtrOutputWithContext(ctx context.Context) GitRepositoryDefinitionPtrOutput {
	return o
}

func (o GitRepositoryDefinitionPtrOutput) Elem() GitRepositoryDefinitionOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) GitRepositoryDefinition {
		if v != nil {
			return *v
		}
		var ret GitRepositoryDefinition
		return ret
	}).(GitRepositoryDefinitionOutput)
}

func (o GitRepositoryDefinitionPtrOutput) HttpsCAFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *string {
		if v == nil {
			return nil
		}
		return v.HttpsCAFile
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) HttpsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *string {
		if v == nil {
			return nil
		}
		return v.HttpsUser
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *string {
		if v == nil {
			return nil
		}
		return v.LocalAuthRef
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) RepositoryRef() RepositoryRefDefinitionPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *RepositoryRefDefinition {
		if v == nil {
			return nil
		}
		return v.RepositoryRef
	}).(RepositoryRefDefinitionPtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) SshKnownHosts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *string {
		if v == nil {
			return nil
		}
		return v.SshKnownHosts
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *float64 {
		if v == nil {
			return nil
		}
		return v.SyncIntervalInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type GitRepositoryDefinitionResponse struct {
	HttpsCAFile           *string                          `pulumi:"httpsCAFile"`
	HttpsUser             *string                          `pulumi:"httpsUser"`
	LocalAuthRef          *string                          `pulumi:"localAuthRef"`
	RepositoryRef         *RepositoryRefDefinitionResponse `pulumi:"repositoryRef"`
	SshKnownHosts         *string                          `pulumi:"sshKnownHosts"`
	SyncIntervalInSeconds *float64                         `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      *float64                         `pulumi:"timeoutInSeconds"`
	Url                   *string                          `pulumi:"url"`
}


func (val *GitRepositoryDefinitionResponse) Defaults() *GitRepositoryDefinitionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}

type GitRepositoryDefinitionResponseOutput struct{ *pulumi.OutputState }

func (GitRepositoryDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitRepositoryDefinitionResponse)(nil)).Elem()
}

func (o GitRepositoryDefinitionResponseOutput) ToGitRepositoryDefinitionResponseOutput() GitRepositoryDefinitionResponseOutput {
	return o
}

func (o GitRepositoryDefinitionResponseOutput) ToGitRepositoryDefinitionResponseOutputWithContext(ctx context.Context) GitRepositoryDefinitionResponseOutput {
	return o
}

func (o GitRepositoryDefinitionResponseOutput) HttpsCAFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *string { return v.HttpsCAFile }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) HttpsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *string { return v.HttpsUser }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *string { return v.LocalAuthRef }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) RepositoryRef() RepositoryRefDefinitionResponsePtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *RepositoryRefDefinitionResponse { return v.RepositoryRef }).(RepositoryRefDefinitionResponsePtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) SshKnownHosts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *string { return v.SshKnownHosts }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type GitRepositoryDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (GitRepositoryDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepositoryDefinitionResponse)(nil)).Elem()
}

func (o GitRepositoryDefinitionResponsePtrOutput) ToGitRepositoryDefinitionResponsePtrOutput() GitRepositoryDefinitionResponsePtrOutput {
	return o
}

func (o GitRepositoryDefinitionResponsePtrOutput) ToGitRepositoryDefinitionResponsePtrOutputWithContext(ctx context.Context) GitRepositoryDefinitionResponsePtrOutput {
	return o
}

func (o GitRepositoryDefinitionResponsePtrOutput) Elem() GitRepositoryDefinitionResponseOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) GitRepositoryDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret GitRepositoryDefinitionResponse
		return ret
	}).(GitRepositoryDefinitionResponseOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) HttpsCAFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.HttpsCAFile
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) HttpsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.HttpsUser
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.LocalAuthRef
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) RepositoryRef() RepositoryRefDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *RepositoryRefDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.RepositoryRef
	}).(RepositoryRefDefinitionResponsePtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) SshKnownHosts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SshKnownHosts
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SyncIntervalInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type HelmOperatorProperties struct {
	ChartValues  *string `pulumi:"chartValues"`
	ChartVersion *string `pulumi:"chartVersion"`
}





type HelmOperatorPropertiesInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput
	ToHelmOperatorPropertiesOutputWithContext(context.Context) HelmOperatorPropertiesOutput
}

type HelmOperatorPropertiesArgs struct {
	ChartValues  pulumi.StringPtrInput `pulumi:"chartValues"`
	ChartVersion pulumi.StringPtrInput `pulumi:"chartVersion"`
}

func (HelmOperatorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorProperties)(nil)).Elem()
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput {
	return i.ToHelmOperatorPropertiesOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesOutputWithContext(ctx context.Context) HelmOperatorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesOutput)
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return i.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesOutput).ToHelmOperatorPropertiesPtrOutputWithContext(ctx)
}









type HelmOperatorPropertiesPtrInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput
	ToHelmOperatorPropertiesPtrOutputWithContext(context.Context) HelmOperatorPropertiesPtrOutput
}

type helmOperatorPropertiesPtrType HelmOperatorPropertiesArgs

func HelmOperatorPropertiesPtr(v *HelmOperatorPropertiesArgs) HelmOperatorPropertiesPtrInput {
	return (*helmOperatorPropertiesPtrType)(v)
}

func (*helmOperatorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorProperties)(nil)).Elem()
}

func (i *helmOperatorPropertiesPtrType) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return i.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (i *helmOperatorPropertiesPtrType) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesPtrOutput)
}

type HelmOperatorPropertiesOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorProperties)(nil)).Elem()
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput {
	return o
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesOutputWithContext(ctx context.Context) HelmOperatorPropertiesOutput {
	return o
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return o.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HelmOperatorProperties) *HelmOperatorProperties {
		return &v
	}).(HelmOperatorPropertiesPtrOutput)
}

func (o HelmOperatorPropertiesOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorProperties) *string { return v.ChartValues }).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorProperties) *string { return v.ChartVersion }).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorProperties)(nil)).Elem()
}

func (o HelmOperatorPropertiesPtrOutput) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return o
}

func (o HelmOperatorPropertiesPtrOutput) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return o
}

func (o HelmOperatorPropertiesPtrOutput) Elem() HelmOperatorPropertiesOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) HelmOperatorProperties {
		if v != nil {
			return *v
		}
		var ret HelmOperatorProperties
		return ret
	}).(HelmOperatorPropertiesOutput)
}

func (o HelmOperatorPropertiesPtrOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) *string {
		if v == nil {
			return nil
		}
		return v.ChartValues
	}).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesPtrOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) *string {
		if v == nil {
			return nil
		}
		return v.ChartVersion
	}).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesResponse struct {
	ChartValues  *string `pulumi:"chartValues"`
	ChartVersion *string `pulumi:"chartVersion"`
}

type HelmOperatorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponseOutput() HelmOperatorPropertiesResponseOutput {
	return o
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponseOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponseOutput {
	return o
}

func (o HelmOperatorPropertiesResponseOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorPropertiesResponse) *string { return v.ChartValues }).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesResponseOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorPropertiesResponse) *string { return v.ChartVersion }).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (o HelmOperatorPropertiesResponsePtrOutput) ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput {
	return o
}

func (o HelmOperatorPropertiesResponsePtrOutput) ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponsePtrOutput {
	return o
}

func (o HelmOperatorPropertiesResponsePtrOutput) Elem() HelmOperatorPropertiesResponseOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) HelmOperatorPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret HelmOperatorPropertiesResponse
		return ret
	}).(HelmOperatorPropertiesResponseOutput)
}

func (o HelmOperatorPropertiesResponsePtrOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ChartValues
	}).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesResponsePtrOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ChartVersion
	}).(pulumi.StringPtrOutput)
}

type HelmReleasePropertiesDefinitionResponse struct {
	FailureCount        *float64                           `pulumi:"failureCount"`
	HelmChartRef        *ObjectReferenceDefinitionResponse `pulumi:"helmChartRef"`
	InstallFailureCount *float64                           `pulumi:"installFailureCount"`
	LastRevisionApplied *float64                           `pulumi:"lastRevisionApplied"`
	UpgradeFailureCount *float64                           `pulumi:"upgradeFailureCount"`
}

type HelmReleasePropertiesDefinitionResponseOutput struct{ *pulumi.OutputState }

func (HelmReleasePropertiesDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmReleasePropertiesDefinitionResponse)(nil)).Elem()
}

func (o HelmReleasePropertiesDefinitionResponseOutput) ToHelmReleasePropertiesDefinitionResponseOutput() HelmReleasePropertiesDefinitionResponseOutput {
	return o
}

func (o HelmReleasePropertiesDefinitionResponseOutput) ToHelmReleasePropertiesDefinitionResponseOutputWithContext(ctx context.Context) HelmReleasePropertiesDefinitionResponseOutput {
	return o
}

func (o HelmReleasePropertiesDefinitionResponseOutput) FailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v HelmReleasePropertiesDefinitionResponse) *float64 { return v.FailureCount }).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponseOutput) HelmChartRef() ObjectReferenceDefinitionResponsePtrOutput {
	return o.ApplyT(func(v HelmReleasePropertiesDefinitionResponse) *ObjectReferenceDefinitionResponse {
		return v.HelmChartRef
	}).(ObjectReferenceDefinitionResponsePtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponseOutput) InstallFailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v HelmReleasePropertiesDefinitionResponse) *float64 { return v.InstallFailureCount }).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponseOutput) LastRevisionApplied() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v HelmReleasePropertiesDefinitionResponse) *float64 { return v.LastRevisionApplied }).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponseOutput) UpgradeFailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v HelmReleasePropertiesDefinitionResponse) *float64 { return v.UpgradeFailureCount }).(pulumi.Float64PtrOutput)
}

type HelmReleasePropertiesDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (HelmReleasePropertiesDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmReleasePropertiesDefinitionResponse)(nil)).Elem()
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) ToHelmReleasePropertiesDefinitionResponsePtrOutput() HelmReleasePropertiesDefinitionResponsePtrOutput {
	return o
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) ToHelmReleasePropertiesDefinitionResponsePtrOutputWithContext(ctx context.Context) HelmReleasePropertiesDefinitionResponsePtrOutput {
	return o
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) Elem() HelmReleasePropertiesDefinitionResponseOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) HelmReleasePropertiesDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret HelmReleasePropertiesDefinitionResponse
		return ret
	}).(HelmReleasePropertiesDefinitionResponseOutput)
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) FailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.FailureCount
	}).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) HelmChartRef() ObjectReferenceDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) *ObjectReferenceDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.HelmChartRef
	}).(ObjectReferenceDefinitionResponsePtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) InstallFailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.InstallFailureCount
	}).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) LastRevisionApplied() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.LastRevisionApplied
	}).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) UpgradeFailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.UpgradeFailureCount
	}).(pulumi.Float64PtrOutput)
}

type KubernetesConfigurationPrivateLinkScopeProperties struct {
	ClusterResourceId   string  `pulumi:"clusterResourceId"`
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}





type KubernetesConfigurationPrivateLinkScopePropertiesInput interface {
	pulumi.Input

	ToKubernetesConfigurationPrivateLinkScopePropertiesOutput() KubernetesConfigurationPrivateLinkScopePropertiesOutput
	ToKubernetesConfigurationPrivateLinkScopePropertiesOutputWithContext(context.Context) KubernetesConfigurationPrivateLinkScopePropertiesOutput
}

type KubernetesConfigurationPrivateLinkScopePropertiesArgs struct {
	ClusterResourceId   pulumi.StringInput    `pulumi:"clusterResourceId"`
	PublicNetworkAccess pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (KubernetesConfigurationPrivateLinkScopePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesConfigurationPrivateLinkScopeProperties)(nil)).Elem()
}

func (i KubernetesConfigurationPrivateLinkScopePropertiesArgs) ToKubernetesConfigurationPrivateLinkScopePropertiesOutput() KubernetesConfigurationPrivateLinkScopePropertiesOutput {
	return i.ToKubernetesConfigurationPrivateLinkScopePropertiesOutputWithContext(context.Background())
}

func (i KubernetesConfigurationPrivateLinkScopePropertiesArgs) ToKubernetesConfigurationPrivateLinkScopePropertiesOutputWithContext(ctx context.Context) KubernetesConfigurationPrivateLinkScopePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesConfigurationPrivateLinkScopePropertiesOutput)
}

func (i KubernetesConfigurationPrivateLinkScopePropertiesArgs) ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutput() KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput {
	return i.ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutputWithContext(context.Background())
}

func (i KubernetesConfigurationPrivateLinkScopePropertiesArgs) ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesConfigurationPrivateLinkScopePropertiesOutput).ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutputWithContext(ctx)
}









type KubernetesConfigurationPrivateLinkScopePropertiesPtrInput interface {
	pulumi.Input

	ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutput() KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput
	ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutputWithContext(context.Context) KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput
}

type kubernetesConfigurationPrivateLinkScopePropertiesPtrType KubernetesConfigurationPrivateLinkScopePropertiesArgs

func KubernetesConfigurationPrivateLinkScopePropertiesPtr(v *KubernetesConfigurationPrivateLinkScopePropertiesArgs) KubernetesConfigurationPrivateLinkScopePropertiesPtrInput {
	return (*kubernetesConfigurationPrivateLinkScopePropertiesPtrType)(v)
}

func (*kubernetesConfigurationPrivateLinkScopePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesConfigurationPrivateLinkScopeProperties)(nil)).Elem()
}

func (i *kubernetesConfigurationPrivateLinkScopePropertiesPtrType) ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutput() KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput {
	return i.ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutputWithContext(context.Background())
}

func (i *kubernetesConfigurationPrivateLinkScopePropertiesPtrType) ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput)
}

type KubernetesConfigurationPrivateLinkScopePropertiesOutput struct{ *pulumi.OutputState }

func (KubernetesConfigurationPrivateLinkScopePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesConfigurationPrivateLinkScopeProperties)(nil)).Elem()
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesOutput) ToKubernetesConfigurationPrivateLinkScopePropertiesOutput() KubernetesConfigurationPrivateLinkScopePropertiesOutput {
	return o
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesOutput) ToKubernetesConfigurationPrivateLinkScopePropertiesOutputWithContext(ctx context.Context) KubernetesConfigurationPrivateLinkScopePropertiesOutput {
	return o
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesOutput) ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutput() KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput {
	return o.ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutputWithContext(context.Background())
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesOutput) ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesConfigurationPrivateLinkScopeProperties) *KubernetesConfigurationPrivateLinkScopeProperties {
		return &v
	}).(KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput)
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesOutput) ClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesConfigurationPrivateLinkScopeProperties) string { return v.ClusterResourceId }).(pulumi.StringOutput)
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesConfigurationPrivateLinkScopeProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput struct{ *pulumi.OutputState }

func (KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesConfigurationPrivateLinkScopeProperties)(nil)).Elem()
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput) ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutput() KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput {
	return o
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput) ToKubernetesConfigurationPrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput {
	return o
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput) Elem() KubernetesConfigurationPrivateLinkScopePropertiesOutput {
	return o.ApplyT(func(v *KubernetesConfigurationPrivateLinkScopeProperties) KubernetesConfigurationPrivateLinkScopeProperties {
		if v != nil {
			return *v
		}
		var ret KubernetesConfigurationPrivateLinkScopeProperties
		return ret
	}).(KubernetesConfigurationPrivateLinkScopePropertiesOutput)
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput) ClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesConfigurationPrivateLinkScopeProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ClusterResourceId
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesConfigurationPrivateLinkScopeProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type KubernetesConfigurationPrivateLinkScopePropertiesResponse struct {
	ClusterResourceId          string                              `pulumi:"clusterResourceId"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	PrivateLinkScopeId         string                              `pulumi:"privateLinkScopeId"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
}

type KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput struct{ *pulumi.OutputState }

func (KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesConfigurationPrivateLinkScopePropertiesResponse)(nil)).Elem()
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput) ToKubernetesConfigurationPrivateLinkScopePropertiesResponseOutput() KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput {
	return o
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput) ToKubernetesConfigurationPrivateLinkScopePropertiesResponseOutputWithContext(ctx context.Context) KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput {
	return o
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput) ClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesConfigurationPrivateLinkScopePropertiesResponse) string { return v.ClusterResourceId }).(pulumi.StringOutput)
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v KubernetesConfigurationPrivateLinkScopePropertiesResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput) PrivateLinkScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesConfigurationPrivateLinkScopePropertiesResponse) string { return v.PrivateLinkScopeId }).(pulumi.StringOutput)
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesConfigurationPrivateLinkScopePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesConfigurationPrivateLinkScopePropertiesResponse) *string {
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type KustomizationDefinition struct {
	DependsOn              []DependsOnDefinition `pulumi:"dependsOn"`
	Force                  *bool                 `pulumi:"force"`
	Path                   *string               `pulumi:"path"`
	Prune                  *bool                 `pulumi:"prune"`
	RetryIntervalInSeconds *float64              `pulumi:"retryIntervalInSeconds"`
	SyncIntervalInSeconds  *float64              `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds       *float64              `pulumi:"timeoutInSeconds"`
	Validation             *string               `pulumi:"validation"`
}


func (val *KustomizationDefinition) Defaults() *KustomizationDefinition {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Force) {
		force_ := false
		tmp.Force = &force_
	}
	if isZero(tmp.Path) {
		path_ := ""
		tmp.Path = &path_
	}
	if isZero(tmp.Prune) {
		prune_ := false
		tmp.Prune = &prune_
	}
	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}





type KustomizationDefinitionInput interface {
	pulumi.Input

	ToKustomizationDefinitionOutput() KustomizationDefinitionOutput
	ToKustomizationDefinitionOutputWithContext(context.Context) KustomizationDefinitionOutput
}

type KustomizationDefinitionArgs struct {
	DependsOn              DependsOnDefinitionArrayInput `pulumi:"dependsOn"`
	Force                  pulumi.BoolPtrInput           `pulumi:"force"`
	Path                   pulumi.StringPtrInput         `pulumi:"path"`
	Prune                  pulumi.BoolPtrInput           `pulumi:"prune"`
	RetryIntervalInSeconds pulumi.Float64PtrInput        `pulumi:"retryIntervalInSeconds"`
	SyncIntervalInSeconds  pulumi.Float64PtrInput        `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds       pulumi.Float64PtrInput        `pulumi:"timeoutInSeconds"`
	Validation             pulumi.StringPtrInput         `pulumi:"validation"`
}


func (val *KustomizationDefinitionArgs) Defaults() *KustomizationDefinitionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Force) {
		tmp.Force = pulumi.BoolPtr(false)
	}
	if isZero(tmp.Path) {
		tmp.Path = pulumi.StringPtr("")
	}
	if isZero(tmp.Prune) {
		tmp.Prune = pulumi.BoolPtr(false)
	}
	if isZero(tmp.SyncIntervalInSeconds) {
		tmp.SyncIntervalInSeconds = pulumi.Float64Ptr(600.0)
	}
	if isZero(tmp.TimeoutInSeconds) {
		tmp.TimeoutInSeconds = pulumi.Float64Ptr(600.0)
	}
	return &tmp
}
func (KustomizationDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KustomizationDefinition)(nil)).Elem()
}

func (i KustomizationDefinitionArgs) ToKustomizationDefinitionOutput() KustomizationDefinitionOutput {
	return i.ToKustomizationDefinitionOutputWithContext(context.Background())
}

func (i KustomizationDefinitionArgs) ToKustomizationDefinitionOutputWithContext(ctx context.Context) KustomizationDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustomizationDefinitionOutput)
}





type KustomizationDefinitionMapInput interface {
	pulumi.Input

	ToKustomizationDefinitionMapOutput() KustomizationDefinitionMapOutput
	ToKustomizationDefinitionMapOutputWithContext(context.Context) KustomizationDefinitionMapOutput
}

type KustomizationDefinitionMap map[string]KustomizationDefinitionInput

func (KustomizationDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KustomizationDefinition)(nil)).Elem()
}

func (i KustomizationDefinitionMap) ToKustomizationDefinitionMapOutput() KustomizationDefinitionMapOutput {
	return i.ToKustomizationDefinitionMapOutputWithContext(context.Background())
}

func (i KustomizationDefinitionMap) ToKustomizationDefinitionMapOutputWithContext(ctx context.Context) KustomizationDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustomizationDefinitionMapOutput)
}

type KustomizationDefinitionOutput struct{ *pulumi.OutputState }

func (KustomizationDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KustomizationDefinition)(nil)).Elem()
}

func (o KustomizationDefinitionOutput) ToKustomizationDefinitionOutput() KustomizationDefinitionOutput {
	return o
}

func (o KustomizationDefinitionOutput) ToKustomizationDefinitionOutputWithContext(ctx context.Context) KustomizationDefinitionOutput {
	return o
}

func (o KustomizationDefinitionOutput) DependsOn() DependsOnDefinitionArrayOutput {
	return o.ApplyT(func(v KustomizationDefinition) []DependsOnDefinition { return v.DependsOn }).(DependsOnDefinitionArrayOutput)
}

func (o KustomizationDefinitionOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *bool { return v.Force }).(pulumi.BoolPtrOutput)
}

func (o KustomizationDefinitionOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o KustomizationDefinitionOutput) Prune() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *bool { return v.Prune }).(pulumi.BoolPtrOutput)
}

func (o KustomizationDefinitionOutput) RetryIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *float64 { return v.RetryIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o KustomizationDefinitionOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o KustomizationDefinitionOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o KustomizationDefinitionOutput) Validation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *string { return v.Validation }).(pulumi.StringPtrOutput)
}

type KustomizationDefinitionMapOutput struct{ *pulumi.OutputState }

func (KustomizationDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KustomizationDefinition)(nil)).Elem()
}

func (o KustomizationDefinitionMapOutput) ToKustomizationDefinitionMapOutput() KustomizationDefinitionMapOutput {
	return o
}

func (o KustomizationDefinitionMapOutput) ToKustomizationDefinitionMapOutputWithContext(ctx context.Context) KustomizationDefinitionMapOutput {
	return o
}

func (o KustomizationDefinitionMapOutput) MapIndex(k pulumi.StringInput) KustomizationDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KustomizationDefinition {
		return vs[0].(map[string]KustomizationDefinition)[vs[1].(string)]
	}).(KustomizationDefinitionOutput)
}

type KustomizationDefinitionResponse struct {
	DependsOn              []DependsOnDefinitionResponse `pulumi:"dependsOn"`
	Force                  *bool                         `pulumi:"force"`
	Path                   *string                       `pulumi:"path"`
	Prune                  *bool                         `pulumi:"prune"`
	RetryIntervalInSeconds *float64                      `pulumi:"retryIntervalInSeconds"`
	SyncIntervalInSeconds  *float64                      `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds       *float64                      `pulumi:"timeoutInSeconds"`
	Validation             *string                       `pulumi:"validation"`
}


func (val *KustomizationDefinitionResponse) Defaults() *KustomizationDefinitionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Force) {
		force_ := false
		tmp.Force = &force_
	}
	if isZero(tmp.Path) {
		path_ := ""
		tmp.Path = &path_
	}
	if isZero(tmp.Prune) {
		prune_ := false
		tmp.Prune = &prune_
	}
	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}

type KustomizationDefinitionResponseOutput struct{ *pulumi.OutputState }

func (KustomizationDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KustomizationDefinitionResponse)(nil)).Elem()
}

func (o KustomizationDefinitionResponseOutput) ToKustomizationDefinitionResponseOutput() KustomizationDefinitionResponseOutput {
	return o
}

func (o KustomizationDefinitionResponseOutput) ToKustomizationDefinitionResponseOutputWithContext(ctx context.Context) KustomizationDefinitionResponseOutput {
	return o
}

func (o KustomizationDefinitionResponseOutput) DependsOn() DependsOnDefinitionResponseArrayOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) []DependsOnDefinitionResponse { return v.DependsOn }).(DependsOnDefinitionResponseArrayOutput)
}

func (o KustomizationDefinitionResponseOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *bool { return v.Force }).(pulumi.BoolPtrOutput)
}

func (o KustomizationDefinitionResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o KustomizationDefinitionResponseOutput) Prune() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *bool { return v.Prune }).(pulumi.BoolPtrOutput)
}

func (o KustomizationDefinitionResponseOutput) RetryIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *float64 { return v.RetryIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o KustomizationDefinitionResponseOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o KustomizationDefinitionResponseOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o KustomizationDefinitionResponseOutput) Validation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *string { return v.Validation }).(pulumi.StringPtrOutput)
}

type KustomizationDefinitionResponseMapOutput struct{ *pulumi.OutputState }

func (KustomizationDefinitionResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KustomizationDefinitionResponse)(nil)).Elem()
}

func (o KustomizationDefinitionResponseMapOutput) ToKustomizationDefinitionResponseMapOutput() KustomizationDefinitionResponseMapOutput {
	return o
}

func (o KustomizationDefinitionResponseMapOutput) ToKustomizationDefinitionResponseMapOutputWithContext(ctx context.Context) KustomizationDefinitionResponseMapOutput {
	return o
}

func (o KustomizationDefinitionResponseMapOutput) MapIndex(k pulumi.StringInput) KustomizationDefinitionResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KustomizationDefinitionResponse {
		return vs[0].(map[string]KustomizationDefinitionResponse)[vs[1].(string)]
	}).(KustomizationDefinitionResponseOutput)
}

type ObjectReferenceDefinitionResponse struct {
	Name      *string `pulumi:"name"`
	Namespace *string `pulumi:"namespace"`
}

type ObjectReferenceDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ObjectReferenceDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectReferenceDefinitionResponse)(nil)).Elem()
}

func (o ObjectReferenceDefinitionResponseOutput) ToObjectReferenceDefinitionResponseOutput() ObjectReferenceDefinitionResponseOutput {
	return o
}

func (o ObjectReferenceDefinitionResponseOutput) ToObjectReferenceDefinitionResponseOutputWithContext(ctx context.Context) ObjectReferenceDefinitionResponseOutput {
	return o
}

func (o ObjectReferenceDefinitionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectReferenceDefinitionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ObjectReferenceDefinitionResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectReferenceDefinitionResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

type ObjectReferenceDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (ObjectReferenceDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectReferenceDefinitionResponse)(nil)).Elem()
}

func (o ObjectReferenceDefinitionResponsePtrOutput) ToObjectReferenceDefinitionResponsePtrOutput() ObjectReferenceDefinitionResponsePtrOutput {
	return o
}

func (o ObjectReferenceDefinitionResponsePtrOutput) ToObjectReferenceDefinitionResponsePtrOutputWithContext(ctx context.Context) ObjectReferenceDefinitionResponsePtrOutput {
	return o
}

func (o ObjectReferenceDefinitionResponsePtrOutput) Elem() ObjectReferenceDefinitionResponseOutput {
	return o.ApplyT(func(v *ObjectReferenceDefinitionResponse) ObjectReferenceDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret ObjectReferenceDefinitionResponse
		return ret
	}).(ObjectReferenceDefinitionResponseOutput)
}

func (o ObjectReferenceDefinitionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectReferenceDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ObjectReferenceDefinitionResponsePtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectReferenceDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

type ObjectStatusConditionDefinitionResponse struct {
	LastTransitionTime *string `pulumi:"lastTransitionTime"`
	Message            *string `pulumi:"message"`
	Reason             *string `pulumi:"reason"`
	Status             *string `pulumi:"status"`
	Type               *string `pulumi:"type"`
}

type ObjectStatusConditionDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ObjectStatusConditionDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectStatusConditionDefinitionResponse)(nil)).Elem()
}

func (o ObjectStatusConditionDefinitionResponseOutput) ToObjectStatusConditionDefinitionResponseOutput() ObjectStatusConditionDefinitionResponseOutput {
	return o
}

func (o ObjectStatusConditionDefinitionResponseOutput) ToObjectStatusConditionDefinitionResponseOutputWithContext(ctx context.Context) ObjectStatusConditionDefinitionResponseOutput {
	return o
}

func (o ObjectStatusConditionDefinitionResponseOutput) LastTransitionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusConditionDefinitionResponse) *string { return v.LastTransitionTime }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusConditionDefinitionResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusConditionDefinitionResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusConditionDefinitionResponseOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusConditionDefinitionResponse) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusConditionDefinitionResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusConditionDefinitionResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusConditionDefinitionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusConditionDefinitionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ObjectStatusConditionDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ObjectStatusConditionDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ObjectStatusConditionDefinitionResponse)(nil)).Elem()
}

func (o ObjectStatusConditionDefinitionResponseArrayOutput) ToObjectStatusConditionDefinitionResponseArrayOutput() ObjectStatusConditionDefinitionResponseArrayOutput {
	return o
}

func (o ObjectStatusConditionDefinitionResponseArrayOutput) ToObjectStatusConditionDefinitionResponseArrayOutputWithContext(ctx context.Context) ObjectStatusConditionDefinitionResponseArrayOutput {
	return o
}

func (o ObjectStatusConditionDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ObjectStatusConditionDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ObjectStatusConditionDefinitionResponse {
		return vs[0].([]ObjectStatusConditionDefinitionResponse)[vs[1].(int)]
	}).(ObjectStatusConditionDefinitionResponseOutput)
}

type ObjectStatusDefinitionResponse struct {
	AppliedBy             *ObjectReferenceDefinitionResponse        `pulumi:"appliedBy"`
	ComplianceState       *string                                   `pulumi:"complianceState"`
	HelmReleaseProperties *HelmReleasePropertiesDefinitionResponse  `pulumi:"helmReleaseProperties"`
	Kind                  *string                                   `pulumi:"kind"`
	Name                  *string                                   `pulumi:"name"`
	Namespace             *string                                   `pulumi:"namespace"`
	StatusConditions      []ObjectStatusConditionDefinitionResponse `pulumi:"statusConditions"`
}

type ObjectStatusDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ObjectStatusDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectStatusDefinitionResponse)(nil)).Elem()
}

func (o ObjectStatusDefinitionResponseOutput) ToObjectStatusDefinitionResponseOutput() ObjectStatusDefinitionResponseOutput {
	return o
}

func (o ObjectStatusDefinitionResponseOutput) ToObjectStatusDefinitionResponseOutputWithContext(ctx context.Context) ObjectStatusDefinitionResponseOutput {
	return o
}

func (o ObjectStatusDefinitionResponseOutput) AppliedBy() ObjectReferenceDefinitionResponsePtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *ObjectReferenceDefinitionResponse { return v.AppliedBy }).(ObjectReferenceDefinitionResponsePtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) ComplianceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *string { return v.ComplianceState }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) HelmReleaseProperties() HelmReleasePropertiesDefinitionResponsePtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *HelmReleasePropertiesDefinitionResponse {
		return v.HelmReleaseProperties
	}).(HelmReleasePropertiesDefinitionResponsePtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) StatusConditions() ObjectStatusConditionDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) []ObjectStatusConditionDefinitionResponse {
		return v.StatusConditions
	}).(ObjectStatusConditionDefinitionResponseArrayOutput)
}

type ObjectStatusDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ObjectStatusDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ObjectStatusDefinitionResponse)(nil)).Elem()
}

func (o ObjectStatusDefinitionResponseArrayOutput) ToObjectStatusDefinitionResponseArrayOutput() ObjectStatusDefinitionResponseArrayOutput {
	return o
}

func (o ObjectStatusDefinitionResponseArrayOutput) ToObjectStatusDefinitionResponseArrayOutputWithContext(ctx context.Context) ObjectStatusDefinitionResponseArrayOutput {
	return o
}

func (o ObjectStatusDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ObjectStatusDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ObjectStatusDefinitionResponse {
		return vs[0].([]ObjectStatusDefinitionResponse)[vs[1].(int)]
	}).(ObjectStatusDefinitionResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type RepositoryRefDefinition struct {
	Branch *string `pulumi:"branch"`
	Commit *string `pulumi:"commit"`
	Semver *string `pulumi:"semver"`
	Tag    *string `pulumi:"tag"`
}





type RepositoryRefDefinitionInput interface {
	pulumi.Input

	ToRepositoryRefDefinitionOutput() RepositoryRefDefinitionOutput
	ToRepositoryRefDefinitionOutputWithContext(context.Context) RepositoryRefDefinitionOutput
}

type RepositoryRefDefinitionArgs struct {
	Branch pulumi.StringPtrInput `pulumi:"branch"`
	Commit pulumi.StringPtrInput `pulumi:"commit"`
	Semver pulumi.StringPtrInput `pulumi:"semver"`
	Tag    pulumi.StringPtrInput `pulumi:"tag"`
}

func (RepositoryRefDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryRefDefinition)(nil)).Elem()
}

func (i RepositoryRefDefinitionArgs) ToRepositoryRefDefinitionOutput() RepositoryRefDefinitionOutput {
	return i.ToRepositoryRefDefinitionOutputWithContext(context.Background())
}

func (i RepositoryRefDefinitionArgs) ToRepositoryRefDefinitionOutputWithContext(ctx context.Context) RepositoryRefDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryRefDefinitionOutput)
}

func (i RepositoryRefDefinitionArgs) ToRepositoryRefDefinitionPtrOutput() RepositoryRefDefinitionPtrOutput {
	return i.ToRepositoryRefDefinitionPtrOutputWithContext(context.Background())
}

func (i RepositoryRefDefinitionArgs) ToRepositoryRefDefinitionPtrOutputWithContext(ctx context.Context) RepositoryRefDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryRefDefinitionOutput).ToRepositoryRefDefinitionPtrOutputWithContext(ctx)
}









type RepositoryRefDefinitionPtrInput interface {
	pulumi.Input

	ToRepositoryRefDefinitionPtrOutput() RepositoryRefDefinitionPtrOutput
	ToRepositoryRefDefinitionPtrOutputWithContext(context.Context) RepositoryRefDefinitionPtrOutput
}

type repositoryRefDefinitionPtrType RepositoryRefDefinitionArgs

func RepositoryRefDefinitionPtr(v *RepositoryRefDefinitionArgs) RepositoryRefDefinitionPtrInput {
	return (*repositoryRefDefinitionPtrType)(v)
}

func (*repositoryRefDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryRefDefinition)(nil)).Elem()
}

func (i *repositoryRefDefinitionPtrType) ToRepositoryRefDefinitionPtrOutput() RepositoryRefDefinitionPtrOutput {
	return i.ToRepositoryRefDefinitionPtrOutputWithContext(context.Background())
}

func (i *repositoryRefDefinitionPtrType) ToRepositoryRefDefinitionPtrOutputWithContext(ctx context.Context) RepositoryRefDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryRefDefinitionPtrOutput)
}

type RepositoryRefDefinitionOutput struct{ *pulumi.OutputState }

func (RepositoryRefDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryRefDefinition)(nil)).Elem()
}

func (o RepositoryRefDefinitionOutput) ToRepositoryRefDefinitionOutput() RepositoryRefDefinitionOutput {
	return o
}

func (o RepositoryRefDefinitionOutput) ToRepositoryRefDefinitionOutputWithContext(ctx context.Context) RepositoryRefDefinitionOutput {
	return o
}

func (o RepositoryRefDefinitionOutput) ToRepositoryRefDefinitionPtrOutput() RepositoryRefDefinitionPtrOutput {
	return o.ToRepositoryRefDefinitionPtrOutputWithContext(context.Background())
}

func (o RepositoryRefDefinitionOutput) ToRepositoryRefDefinitionPtrOutputWithContext(ctx context.Context) RepositoryRefDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RepositoryRefDefinition) *RepositoryRefDefinition {
		return &v
	}).(RepositoryRefDefinitionPtrOutput)
}

func (o RepositoryRefDefinitionOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinition) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinition) *string { return v.Commit }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionOutput) Semver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinition) *string { return v.Semver }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinition) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type RepositoryRefDefinitionPtrOutput struct{ *pulumi.OutputState }

func (RepositoryRefDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryRefDefinition)(nil)).Elem()
}

func (o RepositoryRefDefinitionPtrOutput) ToRepositoryRefDefinitionPtrOutput() RepositoryRefDefinitionPtrOutput {
	return o
}

func (o RepositoryRefDefinitionPtrOutput) ToRepositoryRefDefinitionPtrOutputWithContext(ctx context.Context) RepositoryRefDefinitionPtrOutput {
	return o
}

func (o RepositoryRefDefinitionPtrOutput) Elem() RepositoryRefDefinitionOutput {
	return o.ApplyT(func(v *RepositoryRefDefinition) RepositoryRefDefinition {
		if v != nil {
			return *v
		}
		var ret RepositoryRefDefinition
		return ret
	}).(RepositoryRefDefinitionOutput)
}

func (o RepositoryRefDefinitionPtrOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Branch
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionPtrOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Commit
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionPtrOutput) Semver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Semver
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionPtrOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(pulumi.StringPtrOutput)
}

type RepositoryRefDefinitionResponse struct {
	Branch *string `pulumi:"branch"`
	Commit *string `pulumi:"commit"`
	Semver *string `pulumi:"semver"`
	Tag    *string `pulumi:"tag"`
}

type RepositoryRefDefinitionResponseOutput struct{ *pulumi.OutputState }

func (RepositoryRefDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryRefDefinitionResponse)(nil)).Elem()
}

func (o RepositoryRefDefinitionResponseOutput) ToRepositoryRefDefinitionResponseOutput() RepositoryRefDefinitionResponseOutput {
	return o
}

func (o RepositoryRefDefinitionResponseOutput) ToRepositoryRefDefinitionResponseOutputWithContext(ctx context.Context) RepositoryRefDefinitionResponseOutput {
	return o
}

func (o RepositoryRefDefinitionResponseOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinitionResponse) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponseOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinitionResponse) *string { return v.Commit }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponseOutput) Semver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinitionResponse) *string { return v.Semver }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinitionResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type RepositoryRefDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (RepositoryRefDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryRefDefinitionResponse)(nil)).Elem()
}

func (o RepositoryRefDefinitionResponsePtrOutput) ToRepositoryRefDefinitionResponsePtrOutput() RepositoryRefDefinitionResponsePtrOutput {
	return o
}

func (o RepositoryRefDefinitionResponsePtrOutput) ToRepositoryRefDefinitionResponsePtrOutputWithContext(ctx context.Context) RepositoryRefDefinitionResponsePtrOutput {
	return o
}

func (o RepositoryRefDefinitionResponsePtrOutput) Elem() RepositoryRefDefinitionResponseOutput {
	return o.ApplyT(func(v *RepositoryRefDefinitionResponse) RepositoryRefDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret RepositoryRefDefinitionResponse
		return ret
	}).(RepositoryRefDefinitionResponseOutput)
}

func (o RepositoryRefDefinitionResponsePtrOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Branch
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponsePtrOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Commit
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponsePtrOutput) Semver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Semver
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponsePtrOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(pulumi.StringPtrOutput)
}

type Scope struct {
	Cluster   *ScopeCluster   `pulumi:"cluster"`
	Namespace *ScopeNamespace `pulumi:"namespace"`
}





type ScopeInput interface {
	pulumi.Input

	ToScopeOutput() ScopeOutput
	ToScopeOutputWithContext(context.Context) ScopeOutput
}

type ScopeArgs struct {
	Cluster   ScopeClusterPtrInput   `pulumi:"cluster"`
	Namespace ScopeNamespacePtrInput `pulumi:"namespace"`
}

func (ScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Scope)(nil)).Elem()
}

func (i ScopeArgs) ToScopeOutput() ScopeOutput {
	return i.ToScopeOutputWithContext(context.Background())
}

func (i ScopeArgs) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeOutput)
}

func (i ScopeArgs) ToScopePtrOutput() ScopePtrOutput {
	return i.ToScopePtrOutputWithContext(context.Background())
}

func (i ScopeArgs) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeOutput).ToScopePtrOutputWithContext(ctx)
}









type ScopePtrInput interface {
	pulumi.Input

	ToScopePtrOutput() ScopePtrOutput
	ToScopePtrOutputWithContext(context.Context) ScopePtrOutput
}

type scopePtrType ScopeArgs

func ScopePtr(v *ScopeArgs) ScopePtrInput {
	return (*scopePtrType)(v)
}

func (*scopePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (i *scopePtrType) ToScopePtrOutput() ScopePtrOutput {
	return i.ToScopePtrOutputWithContext(context.Background())
}

func (i *scopePtrType) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopePtrOutput)
}

type ScopeOutput struct{ *pulumi.OutputState }

func (ScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Scope)(nil)).Elem()
}

func (o ScopeOutput) ToScopeOutput() ScopeOutput {
	return o
}

func (o ScopeOutput) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return o
}

func (o ScopeOutput) ToScopePtrOutput() ScopePtrOutput {
	return o.ToScopePtrOutputWithContext(context.Background())
}

func (o ScopeOutput) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Scope) *Scope {
		return &v
	}).(ScopePtrOutput)
}

func (o ScopeOutput) Cluster() ScopeClusterPtrOutput {
	return o.ApplyT(func(v Scope) *ScopeCluster { return v.Cluster }).(ScopeClusterPtrOutput)
}

func (o ScopeOutput) Namespace() ScopeNamespacePtrOutput {
	return o.ApplyT(func(v Scope) *ScopeNamespace { return v.Namespace }).(ScopeNamespacePtrOutput)
}

type ScopePtrOutput struct{ *pulumi.OutputState }

func (ScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (o ScopePtrOutput) ToScopePtrOutput() ScopePtrOutput {
	return o
}

func (o ScopePtrOutput) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return o
}

func (o ScopePtrOutput) Elem() ScopeOutput {
	return o.ApplyT(func(v *Scope) Scope {
		if v != nil {
			return *v
		}
		var ret Scope
		return ret
	}).(ScopeOutput)
}

func (o ScopePtrOutput) Cluster() ScopeClusterPtrOutput {
	return o.ApplyT(func(v *Scope) *ScopeCluster {
		if v == nil {
			return nil
		}
		return v.Cluster
	}).(ScopeClusterPtrOutput)
}

func (o ScopePtrOutput) Namespace() ScopeNamespacePtrOutput {
	return o.ApplyT(func(v *Scope) *ScopeNamespace {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(ScopeNamespacePtrOutput)
}

type ScopeCluster struct {
	ReleaseNamespace *string `pulumi:"releaseNamespace"`
}





type ScopeClusterInput interface {
	pulumi.Input

	ToScopeClusterOutput() ScopeClusterOutput
	ToScopeClusterOutputWithContext(context.Context) ScopeClusterOutput
}

type ScopeClusterArgs struct {
	ReleaseNamespace pulumi.StringPtrInput `pulumi:"releaseNamespace"`
}

func (ScopeClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeCluster)(nil)).Elem()
}

func (i ScopeClusterArgs) ToScopeClusterOutput() ScopeClusterOutput {
	return i.ToScopeClusterOutputWithContext(context.Background())
}

func (i ScopeClusterArgs) ToScopeClusterOutputWithContext(ctx context.Context) ScopeClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterOutput)
}

func (i ScopeClusterArgs) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return i.ToScopeClusterPtrOutputWithContext(context.Background())
}

func (i ScopeClusterArgs) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterOutput).ToScopeClusterPtrOutputWithContext(ctx)
}









type ScopeClusterPtrInput interface {
	pulumi.Input

	ToScopeClusterPtrOutput() ScopeClusterPtrOutput
	ToScopeClusterPtrOutputWithContext(context.Context) ScopeClusterPtrOutput
}

type scopeClusterPtrType ScopeClusterArgs

func ScopeClusterPtr(v *ScopeClusterArgs) ScopeClusterPtrInput {
	return (*scopeClusterPtrType)(v)
}

func (*scopeClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeCluster)(nil)).Elem()
}

func (i *scopeClusterPtrType) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return i.ToScopeClusterPtrOutputWithContext(context.Background())
}

func (i *scopeClusterPtrType) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterPtrOutput)
}

type ScopeClusterOutput struct{ *pulumi.OutputState }

func (ScopeClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeCluster)(nil)).Elem()
}

func (o ScopeClusterOutput) ToScopeClusterOutput() ScopeClusterOutput {
	return o
}

func (o ScopeClusterOutput) ToScopeClusterOutputWithContext(ctx context.Context) ScopeClusterOutput {
	return o
}

func (o ScopeClusterOutput) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return o.ToScopeClusterPtrOutputWithContext(context.Background())
}

func (o ScopeClusterOutput) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeCluster) *ScopeCluster {
		return &v
	}).(ScopeClusterPtrOutput)
}

func (o ScopeClusterOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeCluster) *string { return v.ReleaseNamespace }).(pulumi.StringPtrOutput)
}

type ScopeClusterPtrOutput struct{ *pulumi.OutputState }

func (ScopeClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeCluster)(nil)).Elem()
}

func (o ScopeClusterPtrOutput) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return o
}

func (o ScopeClusterPtrOutput) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return o
}

func (o ScopeClusterPtrOutput) Elem() ScopeClusterOutput {
	return o.ApplyT(func(v *ScopeCluster) ScopeCluster {
		if v != nil {
			return *v
		}
		var ret ScopeCluster
		return ret
	}).(ScopeClusterOutput)
}

func (o ScopeClusterPtrOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeCluster) *string {
		if v == nil {
			return nil
		}
		return v.ReleaseNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeClusterResponse struct {
	ReleaseNamespace *string `pulumi:"releaseNamespace"`
}

type ScopeClusterResponseOutput struct{ *pulumi.OutputState }

func (ScopeClusterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeClusterResponse)(nil)).Elem()
}

func (o ScopeClusterResponseOutput) ToScopeClusterResponseOutput() ScopeClusterResponseOutput {
	return o
}

func (o ScopeClusterResponseOutput) ToScopeClusterResponseOutputWithContext(ctx context.Context) ScopeClusterResponseOutput {
	return o
}

func (o ScopeClusterResponseOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeClusterResponse) *string { return v.ReleaseNamespace }).(pulumi.StringPtrOutput)
}

type ScopeClusterResponsePtrOutput struct{ *pulumi.OutputState }

func (ScopeClusterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeClusterResponse)(nil)).Elem()
}

func (o ScopeClusterResponsePtrOutput) ToScopeClusterResponsePtrOutput() ScopeClusterResponsePtrOutput {
	return o
}

func (o ScopeClusterResponsePtrOutput) ToScopeClusterResponsePtrOutputWithContext(ctx context.Context) ScopeClusterResponsePtrOutput {
	return o
}

func (o ScopeClusterResponsePtrOutput) Elem() ScopeClusterResponseOutput {
	return o.ApplyT(func(v *ScopeClusterResponse) ScopeClusterResponse {
		if v != nil {
			return *v
		}
		var ret ScopeClusterResponse
		return ret
	}).(ScopeClusterResponseOutput)
}

func (o ScopeClusterResponsePtrOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeClusterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReleaseNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeNamespace struct {
	TargetNamespace *string `pulumi:"targetNamespace"`
}





type ScopeNamespaceInput interface {
	pulumi.Input

	ToScopeNamespaceOutput() ScopeNamespaceOutput
	ToScopeNamespaceOutputWithContext(context.Context) ScopeNamespaceOutput
}

type ScopeNamespaceArgs struct {
	TargetNamespace pulumi.StringPtrInput `pulumi:"targetNamespace"`
}

func (ScopeNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeNamespace)(nil)).Elem()
}

func (i ScopeNamespaceArgs) ToScopeNamespaceOutput() ScopeNamespaceOutput {
	return i.ToScopeNamespaceOutputWithContext(context.Background())
}

func (i ScopeNamespaceArgs) ToScopeNamespaceOutputWithContext(ctx context.Context) ScopeNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceOutput)
}

func (i ScopeNamespaceArgs) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return i.ToScopeNamespacePtrOutputWithContext(context.Background())
}

func (i ScopeNamespaceArgs) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceOutput).ToScopeNamespacePtrOutputWithContext(ctx)
}









type ScopeNamespacePtrInput interface {
	pulumi.Input

	ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput
	ToScopeNamespacePtrOutputWithContext(context.Context) ScopeNamespacePtrOutput
}

type scopeNamespacePtrType ScopeNamespaceArgs

func ScopeNamespacePtr(v *ScopeNamespaceArgs) ScopeNamespacePtrInput {
	return (*scopeNamespacePtrType)(v)
}

func (*scopeNamespacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeNamespace)(nil)).Elem()
}

func (i *scopeNamespacePtrType) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return i.ToScopeNamespacePtrOutputWithContext(context.Background())
}

func (i *scopeNamespacePtrType) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespacePtrOutput)
}

type ScopeNamespaceOutput struct{ *pulumi.OutputState }

func (ScopeNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeNamespace)(nil)).Elem()
}

func (o ScopeNamespaceOutput) ToScopeNamespaceOutput() ScopeNamespaceOutput {
	return o
}

func (o ScopeNamespaceOutput) ToScopeNamespaceOutputWithContext(ctx context.Context) ScopeNamespaceOutput {
	return o
}

func (o ScopeNamespaceOutput) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return o.ToScopeNamespacePtrOutputWithContext(context.Background())
}

func (o ScopeNamespaceOutput) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeNamespace) *ScopeNamespace {
		return &v
	}).(ScopeNamespacePtrOutput)
}

func (o ScopeNamespaceOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeNamespace) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type ScopeNamespacePtrOutput struct{ *pulumi.OutputState }

func (ScopeNamespacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeNamespace)(nil)).Elem()
}

func (o ScopeNamespacePtrOutput) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return o
}

func (o ScopeNamespacePtrOutput) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return o
}

func (o ScopeNamespacePtrOutput) Elem() ScopeNamespaceOutput {
	return o.ApplyT(func(v *ScopeNamespace) ScopeNamespace {
		if v != nil {
			return *v
		}
		var ret ScopeNamespace
		return ret
	}).(ScopeNamespaceOutput)
}

func (o ScopeNamespacePtrOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeNamespace) *string {
		if v == nil {
			return nil
		}
		return v.TargetNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeNamespaceResponse struct {
	TargetNamespace *string `pulumi:"targetNamespace"`
}

type ScopeNamespaceResponseOutput struct{ *pulumi.OutputState }

func (ScopeNamespaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeNamespaceResponse)(nil)).Elem()
}

func (o ScopeNamespaceResponseOutput) ToScopeNamespaceResponseOutput() ScopeNamespaceResponseOutput {
	return o
}

func (o ScopeNamespaceResponseOutput) ToScopeNamespaceResponseOutputWithContext(ctx context.Context) ScopeNamespaceResponseOutput {
	return o
}

func (o ScopeNamespaceResponseOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeNamespaceResponse) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type ScopeNamespaceResponsePtrOutput struct{ *pulumi.OutputState }

func (ScopeNamespaceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeNamespaceResponse)(nil)).Elem()
}

func (o ScopeNamespaceResponsePtrOutput) ToScopeNamespaceResponsePtrOutput() ScopeNamespaceResponsePtrOutput {
	return o
}

func (o ScopeNamespaceResponsePtrOutput) ToScopeNamespaceResponsePtrOutputWithContext(ctx context.Context) ScopeNamespaceResponsePtrOutput {
	return o
}

func (o ScopeNamespaceResponsePtrOutput) Elem() ScopeNamespaceResponseOutput {
	return o.ApplyT(func(v *ScopeNamespaceResponse) ScopeNamespaceResponse {
		if v != nil {
			return *v
		}
		var ret ScopeNamespaceResponse
		return ret
	}).(ScopeNamespaceResponseOutput)
}

func (o ScopeNamespaceResponsePtrOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeNamespaceResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeResponse struct {
	Cluster   *ScopeClusterResponse   `pulumi:"cluster"`
	Namespace *ScopeNamespaceResponse `pulumi:"namespace"`
}

type ScopeResponseOutput struct{ *pulumi.OutputState }

func (ScopeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeResponse)(nil)).Elem()
}

func (o ScopeResponseOutput) ToScopeResponseOutput() ScopeResponseOutput {
	return o
}

func (o ScopeResponseOutput) ToScopeResponseOutputWithContext(ctx context.Context) ScopeResponseOutput {
	return o
}

func (o ScopeResponseOutput) Cluster() ScopeClusterResponsePtrOutput {
	return o.ApplyT(func(v ScopeResponse) *ScopeClusterResponse { return v.Cluster }).(ScopeClusterResponsePtrOutput)
}

func (o ScopeResponseOutput) Namespace() ScopeNamespaceResponsePtrOutput {
	return o.ApplyT(func(v ScopeResponse) *ScopeNamespaceResponse { return v.Namespace }).(ScopeNamespaceResponsePtrOutput)
}

type ScopeResponsePtrOutput struct{ *pulumi.OutputState }

func (ScopeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeResponse)(nil)).Elem()
}

func (o ScopeResponsePtrOutput) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return o
}

func (o ScopeResponsePtrOutput) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return o
}

func (o ScopeResponsePtrOutput) Elem() ScopeResponseOutput {
	return o.ApplyT(func(v *ScopeResponse) ScopeResponse {
		if v != nil {
			return *v
		}
		var ret ScopeResponse
		return ret
	}).(ScopeResponseOutput)
}

func (o ScopeResponsePtrOutput) Cluster() ScopeClusterResponsePtrOutput {
	return o.ApplyT(func(v *ScopeResponse) *ScopeClusterResponse {
		if v == nil {
			return nil
		}
		return v.Cluster
	}).(ScopeClusterResponsePtrOutput)
}

func (o ScopeResponsePtrOutput) Namespace() ScopeNamespaceResponsePtrOutput {
	return o.ApplyT(func(v *ScopeResponse) *ScopeNamespaceResponse {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(ScopeNamespaceResponsePtrOutput)
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

func init() {
	pulumi.RegisterOutputType(ComplianceStatusResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationIdentityOutput{})
	pulumi.RegisterOutputType(ConfigurationIdentityPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationIdentityResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(DependsOnDefinitionOutput{})
	pulumi.RegisterOutputType(DependsOnDefinitionArrayOutput{})
	pulumi.RegisterOutputType(DependsOnDefinitionResponseOutput{})
	pulumi.RegisterOutputType(DependsOnDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ExtensionStatusOutput{})
	pulumi.RegisterOutputType(ExtensionStatusArrayOutput{})
	pulumi.RegisterOutputType(ExtensionStatusResponseOutput{})
	pulumi.RegisterOutputType(ExtensionStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(GitRepositoryDefinitionOutput{})
	pulumi.RegisterOutputType(GitRepositoryDefinitionPtrOutput{})
	pulumi.RegisterOutputType(GitRepositoryDefinitionResponseOutput{})
	pulumi.RegisterOutputType(GitRepositoryDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(HelmReleasePropertiesDefinitionResponseOutput{})
	pulumi.RegisterOutputType(HelmReleasePropertiesDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(KubernetesConfigurationPrivateLinkScopePropertiesOutput{})
	pulumi.RegisterOutputType(KubernetesConfigurationPrivateLinkScopePropertiesPtrOutput{})
	pulumi.RegisterOutputType(KubernetesConfigurationPrivateLinkScopePropertiesResponseOutput{})
	pulumi.RegisterOutputType(KustomizationDefinitionOutput{})
	pulumi.RegisterOutputType(KustomizationDefinitionMapOutput{})
	pulumi.RegisterOutputType(KustomizationDefinitionResponseOutput{})
	pulumi.RegisterOutputType(KustomizationDefinitionResponseMapOutput{})
	pulumi.RegisterOutputType(ObjectReferenceDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ObjectReferenceDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ObjectStatusConditionDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ObjectStatusConditionDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(ObjectStatusDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ObjectStatusDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(RepositoryRefDefinitionOutput{})
	pulumi.RegisterOutputType(RepositoryRefDefinitionPtrOutput{})
	pulumi.RegisterOutputType(RepositoryRefDefinitionResponseOutput{})
	pulumi.RegisterOutputType(RepositoryRefDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ScopeOutput{})
	pulumi.RegisterOutputType(ScopePtrOutput{})
	pulumi.RegisterOutputType(ScopeClusterOutput{})
	pulumi.RegisterOutputType(ScopeClusterPtrOutput{})
	pulumi.RegisterOutputType(ScopeClusterResponseOutput{})
	pulumi.RegisterOutputType(ScopeClusterResponsePtrOutput{})
	pulumi.RegisterOutputType(ScopeNamespaceOutput{})
	pulumi.RegisterOutputType(ScopeNamespacePtrOutput{})
	pulumi.RegisterOutputType(ScopeNamespaceResponseOutput{})
	pulumi.RegisterOutputType(ScopeNamespaceResponsePtrOutput{})
	pulumi.RegisterOutputType(ScopeResponseOutput{})
	pulumi.RegisterOutputType(ScopeResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
