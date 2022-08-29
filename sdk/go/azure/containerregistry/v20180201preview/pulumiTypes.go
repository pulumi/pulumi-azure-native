


package v20180201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BaseImageDependencyResponse struct {
	Digest     *string `pulumi:"digest"`
	Registry   *string `pulumi:"registry"`
	Repository *string `pulumi:"repository"`
	Tag        *string `pulumi:"tag"`
	Type       *string `pulumi:"type"`
}

type BaseImageDependencyResponseOutput struct{ *pulumi.OutputState }

func (BaseImageDependencyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BaseImageDependencyResponse)(nil)).Elem()
}

func (o BaseImageDependencyResponseOutput) ToBaseImageDependencyResponseOutput() BaseImageDependencyResponseOutput {
	return o
}

func (o BaseImageDependencyResponseOutput) ToBaseImageDependencyResponseOutputWithContext(ctx context.Context) BaseImageDependencyResponseOutput {
	return o
}

func (o BaseImageDependencyResponseOutput) Digest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageDependencyResponse) *string { return v.Digest }).(pulumi.StringPtrOutput)
}

func (o BaseImageDependencyResponseOutput) Registry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageDependencyResponse) *string { return v.Registry }).(pulumi.StringPtrOutput)
}

func (o BaseImageDependencyResponseOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageDependencyResponse) *string { return v.Repository }).(pulumi.StringPtrOutput)
}

func (o BaseImageDependencyResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageDependencyResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func (o BaseImageDependencyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BaseImageDependencyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type BaseImageDependencyResponseArrayOutput struct{ *pulumi.OutputState }

func (BaseImageDependencyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BaseImageDependencyResponse)(nil)).Elem()
}

func (o BaseImageDependencyResponseArrayOutput) ToBaseImageDependencyResponseArrayOutput() BaseImageDependencyResponseArrayOutput {
	return o
}

func (o BaseImageDependencyResponseArrayOutput) ToBaseImageDependencyResponseArrayOutputWithContext(ctx context.Context) BaseImageDependencyResponseArrayOutput {
	return o
}

func (o BaseImageDependencyResponseArrayOutput) Index(i pulumi.IntInput) BaseImageDependencyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BaseImageDependencyResponse {
		return vs[0].([]BaseImageDependencyResponse)[vs[1].(int)]
	}).(BaseImageDependencyResponseOutput)
}

type BuildArgumentResponse struct {
	IsSecret *bool  `pulumi:"isSecret"`
	Name     string `pulumi:"name"`
	Type     string `pulumi:"type"`
	Value    string `pulumi:"value"`
}


func (val *BuildArgumentResponse) Defaults() *BuildArgumentResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSecret) {
		isSecret_ := false
		tmp.IsSecret = &isSecret_
	}
	return &tmp
}

type BuildArgumentResponseOutput struct{ *pulumi.OutputState }

func (BuildArgumentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildArgumentResponse)(nil)).Elem()
}

func (o BuildArgumentResponseOutput) ToBuildArgumentResponseOutput() BuildArgumentResponseOutput {
	return o
}

func (o BuildArgumentResponseOutput) ToBuildArgumentResponseOutputWithContext(ctx context.Context) BuildArgumentResponseOutput {
	return o
}

func (o BuildArgumentResponseOutput) GetIsSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BuildArgumentResponse) *bool { return v.IsSecret }).(pulumi.BoolPtrOutput)
}

func (o BuildArgumentResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BuildArgumentResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o BuildArgumentResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BuildArgumentResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o BuildArgumentResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v BuildArgumentResponse) string { return v.Value }).(pulumi.StringOutput)
}

type BuildArgumentResponseArrayOutput struct{ *pulumi.OutputState }

func (BuildArgumentResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildArgumentResponse)(nil)).Elem()
}

func (o BuildArgumentResponseArrayOutput) ToBuildArgumentResponseArrayOutput() BuildArgumentResponseArrayOutput {
	return o
}

func (o BuildArgumentResponseArrayOutput) ToBuildArgumentResponseArrayOutputWithContext(ctx context.Context) BuildArgumentResponseArrayOutput {
	return o
}

func (o BuildArgumentResponseArrayOutput) Index(i pulumi.IntInput) BuildArgumentResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BuildArgumentResponse {
		return vs[0].([]BuildArgumentResponse)[vs[1].(int)]
	}).(BuildArgumentResponseOutput)
}

type DockerBuildStepResponse struct {
	BaseImageDependencies []BaseImageDependencyResponse `pulumi:"baseImageDependencies"`
	BaseImageTrigger      *string                       `pulumi:"baseImageTrigger"`
	Branch                *string                       `pulumi:"branch"`
	BuildArguments        []BuildArgumentResponse       `pulumi:"buildArguments"`
	ContextPath           *string                       `pulumi:"contextPath"`
	DockerFilePath        *string                       `pulumi:"dockerFilePath"`
	ImageNames            []string                      `pulumi:"imageNames"`
	IsPushEnabled         *bool                         `pulumi:"isPushEnabled"`
	NoCache               *bool                         `pulumi:"noCache"`
	ProvisioningState     string                        `pulumi:"provisioningState"`
	Type                  string                        `pulumi:"type"`
}


func (val *DockerBuildStepResponse) Defaults() *DockerBuildStepResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsPushEnabled) {
		isPushEnabled_ := true
		tmp.IsPushEnabled = &isPushEnabled_
	}
	if isZero(tmp.NoCache) {
		noCache_ := false
		tmp.NoCache = &noCache_
	}
	return &tmp
}

type DockerBuildStepResponseOutput struct{ *pulumi.OutputState }

func (DockerBuildStepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuildStepResponse)(nil)).Elem()
}

func (o DockerBuildStepResponseOutput) ToDockerBuildStepResponseOutput() DockerBuildStepResponseOutput {
	return o
}

func (o DockerBuildStepResponseOutput) ToDockerBuildStepResponseOutputWithContext(ctx context.Context) DockerBuildStepResponseOutput {
	return o
}

func (o DockerBuildStepResponseOutput) BaseImageDependencies() BaseImageDependencyResponseArrayOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) []BaseImageDependencyResponse { return v.BaseImageDependencies }).(BaseImageDependencyResponseArrayOutput)
}

func (o DockerBuildStepResponseOutput) BaseImageTrigger() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) *string { return v.BaseImageTrigger }).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponseOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponseOutput) BuildArguments() BuildArgumentResponseArrayOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) []BuildArgumentResponse { return v.BuildArguments }).(BuildArgumentResponseArrayOutput)
}

func (o DockerBuildStepResponseOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) *string { return v.ContextPath }).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponseOutput) DockerFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) *string { return v.DockerFilePath }).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponseOutput) ImageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) []string { return v.ImageNames }).(pulumi.StringArrayOutput)
}

func (o DockerBuildStepResponseOutput) IsPushEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) *bool { return v.IsPushEnabled }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildStepResponseOutput) NoCache() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) *bool { return v.NoCache }).(pulumi.BoolPtrOutput)
}

func (o DockerBuildStepResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DockerBuildStepResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DockerBuildStepResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PlatformProperties struct {
	Cpu    *int   `pulumi:"cpu"`
	OsType string `pulumi:"osType"`
}





type PlatformPropertiesInput interface {
	pulumi.Input

	ToPlatformPropertiesOutput() PlatformPropertiesOutput
	ToPlatformPropertiesOutputWithContext(context.Context) PlatformPropertiesOutput
}

type PlatformPropertiesArgs struct {
	Cpu    pulumi.IntPtrInput `pulumi:"cpu"`
	OsType pulumi.StringInput `pulumi:"osType"`
}

func (PlatformPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformProperties)(nil)).Elem()
}

func (i PlatformPropertiesArgs) ToPlatformPropertiesOutput() PlatformPropertiesOutput {
	return i.ToPlatformPropertiesOutputWithContext(context.Background())
}

func (i PlatformPropertiesArgs) ToPlatformPropertiesOutputWithContext(ctx context.Context) PlatformPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesOutput)
}

type PlatformPropertiesOutput struct{ *pulumi.OutputState }

func (PlatformPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformProperties)(nil)).Elem()
}

func (o PlatformPropertiesOutput) ToPlatformPropertiesOutput() PlatformPropertiesOutput {
	return o
}

func (o PlatformPropertiesOutput) ToPlatformPropertiesOutputWithContext(ctx context.Context) PlatformPropertiesOutput {
	return o
}

func (o PlatformPropertiesOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PlatformProperties) *int { return v.Cpu }).(pulumi.IntPtrOutput)
}

func (o PlatformPropertiesOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformProperties) string { return v.OsType }).(pulumi.StringOutput)
}

type PlatformPropertiesResponse struct {
	Cpu    *int   `pulumi:"cpu"`
	OsType string `pulumi:"osType"`
}

type PlatformPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PlatformPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformPropertiesResponse)(nil)).Elem()
}

func (o PlatformPropertiesResponseOutput) ToPlatformPropertiesResponseOutput() PlatformPropertiesResponseOutput {
	return o
}

func (o PlatformPropertiesResponseOutput) ToPlatformPropertiesResponseOutputWithContext(ctx context.Context) PlatformPropertiesResponseOutput {
	return o
}

func (o PlatformPropertiesResponseOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PlatformPropertiesResponse) *int { return v.Cpu }).(pulumi.IntPtrOutput)
}

func (o PlatformPropertiesResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformPropertiesResponse) string { return v.OsType }).(pulumi.StringOutput)
}

type SourceControlAuthInfo struct {
	ExpiresIn    *int    `pulumi:"expiresIn"`
	RefreshToken *string `pulumi:"refreshToken"`
	Scope        *string `pulumi:"scope"`
	Token        string  `pulumi:"token"`
	TokenType    *string `pulumi:"tokenType"`
}





type SourceControlAuthInfoInput interface {
	pulumi.Input

	ToSourceControlAuthInfoOutput() SourceControlAuthInfoOutput
	ToSourceControlAuthInfoOutputWithContext(context.Context) SourceControlAuthInfoOutput
}

type SourceControlAuthInfoArgs struct {
	ExpiresIn    pulumi.IntPtrInput    `pulumi:"expiresIn"`
	RefreshToken pulumi.StringPtrInput `pulumi:"refreshToken"`
	Scope        pulumi.StringPtrInput `pulumi:"scope"`
	Token        pulumi.StringInput    `pulumi:"token"`
	TokenType    pulumi.StringPtrInput `pulumi:"tokenType"`
}

func (SourceControlAuthInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControlAuthInfo)(nil)).Elem()
}

func (i SourceControlAuthInfoArgs) ToSourceControlAuthInfoOutput() SourceControlAuthInfoOutput {
	return i.ToSourceControlAuthInfoOutputWithContext(context.Background())
}

func (i SourceControlAuthInfoArgs) ToSourceControlAuthInfoOutputWithContext(ctx context.Context) SourceControlAuthInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlAuthInfoOutput)
}

func (i SourceControlAuthInfoArgs) ToSourceControlAuthInfoPtrOutput() SourceControlAuthInfoPtrOutput {
	return i.ToSourceControlAuthInfoPtrOutputWithContext(context.Background())
}

func (i SourceControlAuthInfoArgs) ToSourceControlAuthInfoPtrOutputWithContext(ctx context.Context) SourceControlAuthInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlAuthInfoOutput).ToSourceControlAuthInfoPtrOutputWithContext(ctx)
}









type SourceControlAuthInfoPtrInput interface {
	pulumi.Input

	ToSourceControlAuthInfoPtrOutput() SourceControlAuthInfoPtrOutput
	ToSourceControlAuthInfoPtrOutputWithContext(context.Context) SourceControlAuthInfoPtrOutput
}

type sourceControlAuthInfoPtrType SourceControlAuthInfoArgs

func SourceControlAuthInfoPtr(v *SourceControlAuthInfoArgs) SourceControlAuthInfoPtrInput {
	return (*sourceControlAuthInfoPtrType)(v)
}

func (*sourceControlAuthInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControlAuthInfo)(nil)).Elem()
}

func (i *sourceControlAuthInfoPtrType) ToSourceControlAuthInfoPtrOutput() SourceControlAuthInfoPtrOutput {
	return i.ToSourceControlAuthInfoPtrOutputWithContext(context.Background())
}

func (i *sourceControlAuthInfoPtrType) ToSourceControlAuthInfoPtrOutputWithContext(ctx context.Context) SourceControlAuthInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlAuthInfoPtrOutput)
}

type SourceControlAuthInfoOutput struct{ *pulumi.OutputState }

func (SourceControlAuthInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControlAuthInfo)(nil)).Elem()
}

func (o SourceControlAuthInfoOutput) ToSourceControlAuthInfoOutput() SourceControlAuthInfoOutput {
	return o
}

func (o SourceControlAuthInfoOutput) ToSourceControlAuthInfoOutputWithContext(ctx context.Context) SourceControlAuthInfoOutput {
	return o
}

func (o SourceControlAuthInfoOutput) ToSourceControlAuthInfoPtrOutput() SourceControlAuthInfoPtrOutput {
	return o.ToSourceControlAuthInfoPtrOutputWithContext(context.Background())
}

func (o SourceControlAuthInfoOutput) ToSourceControlAuthInfoPtrOutputWithContext(ctx context.Context) SourceControlAuthInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceControlAuthInfo) *SourceControlAuthInfo {
		return &v
	}).(SourceControlAuthInfoPtrOutput)
}

func (o SourceControlAuthInfoOutput) ExpiresIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SourceControlAuthInfo) *int { return v.ExpiresIn }).(pulumi.IntPtrOutput)
}

func (o SourceControlAuthInfoOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceControlAuthInfo) *string { return v.RefreshToken }).(pulumi.StringPtrOutput)
}

func (o SourceControlAuthInfoOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceControlAuthInfo) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o SourceControlAuthInfoOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v SourceControlAuthInfo) string { return v.Token }).(pulumi.StringOutput)
}

func (o SourceControlAuthInfoOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceControlAuthInfo) *string { return v.TokenType }).(pulumi.StringPtrOutput)
}

type SourceControlAuthInfoPtrOutput struct{ *pulumi.OutputState }

func (SourceControlAuthInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControlAuthInfo)(nil)).Elem()
}

func (o SourceControlAuthInfoPtrOutput) ToSourceControlAuthInfoPtrOutput() SourceControlAuthInfoPtrOutput {
	return o
}

func (o SourceControlAuthInfoPtrOutput) ToSourceControlAuthInfoPtrOutputWithContext(ctx context.Context) SourceControlAuthInfoPtrOutput {
	return o
}

func (o SourceControlAuthInfoPtrOutput) Elem() SourceControlAuthInfoOutput {
	return o.ApplyT(func(v *SourceControlAuthInfo) SourceControlAuthInfo {
		if v != nil {
			return *v
		}
		var ret SourceControlAuthInfo
		return ret
	}).(SourceControlAuthInfoOutput)
}

func (o SourceControlAuthInfoPtrOutput) ExpiresIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SourceControlAuthInfo) *int {
		if v == nil {
			return nil
		}
		return v.ExpiresIn
	}).(pulumi.IntPtrOutput)
}

func (o SourceControlAuthInfoPtrOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlAuthInfo) *string {
		if v == nil {
			return nil
		}
		return v.RefreshToken
	}).(pulumi.StringPtrOutput)
}

func (o SourceControlAuthInfoPtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlAuthInfo) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

func (o SourceControlAuthInfoPtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlAuthInfo) *string {
		if v == nil {
			return nil
		}
		return &v.Token
	}).(pulumi.StringPtrOutput)
}

func (o SourceControlAuthInfoPtrOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlAuthInfo) *string {
		if v == nil {
			return nil
		}
		return v.TokenType
	}).(pulumi.StringPtrOutput)
}

type SourceControlAuthInfoResponse struct {
	ExpiresIn    *int    `pulumi:"expiresIn"`
	RefreshToken *string `pulumi:"refreshToken"`
	Scope        *string `pulumi:"scope"`
	Token        string  `pulumi:"token"`
	TokenType    *string `pulumi:"tokenType"`
}

type SourceControlAuthInfoResponseOutput struct{ *pulumi.OutputState }

func (SourceControlAuthInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControlAuthInfoResponse)(nil)).Elem()
}

func (o SourceControlAuthInfoResponseOutput) ToSourceControlAuthInfoResponseOutput() SourceControlAuthInfoResponseOutput {
	return o
}

func (o SourceControlAuthInfoResponseOutput) ToSourceControlAuthInfoResponseOutputWithContext(ctx context.Context) SourceControlAuthInfoResponseOutput {
	return o
}

func (o SourceControlAuthInfoResponseOutput) ExpiresIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SourceControlAuthInfoResponse) *int { return v.ExpiresIn }).(pulumi.IntPtrOutput)
}

func (o SourceControlAuthInfoResponseOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceControlAuthInfoResponse) *string { return v.RefreshToken }).(pulumi.StringPtrOutput)
}

func (o SourceControlAuthInfoResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceControlAuthInfoResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o SourceControlAuthInfoResponseOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v SourceControlAuthInfoResponse) string { return v.Token }).(pulumi.StringOutput)
}

func (o SourceControlAuthInfoResponseOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceControlAuthInfoResponse) *string { return v.TokenType }).(pulumi.StringPtrOutput)
}

type SourceControlAuthInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (SourceControlAuthInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControlAuthInfoResponse)(nil)).Elem()
}

func (o SourceControlAuthInfoResponsePtrOutput) ToSourceControlAuthInfoResponsePtrOutput() SourceControlAuthInfoResponsePtrOutput {
	return o
}

func (o SourceControlAuthInfoResponsePtrOutput) ToSourceControlAuthInfoResponsePtrOutputWithContext(ctx context.Context) SourceControlAuthInfoResponsePtrOutput {
	return o
}

func (o SourceControlAuthInfoResponsePtrOutput) Elem() SourceControlAuthInfoResponseOutput {
	return o.ApplyT(func(v *SourceControlAuthInfoResponse) SourceControlAuthInfoResponse {
		if v != nil {
			return *v
		}
		var ret SourceControlAuthInfoResponse
		return ret
	}).(SourceControlAuthInfoResponseOutput)
}

func (o SourceControlAuthInfoResponsePtrOutput) ExpiresIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SourceControlAuthInfoResponse) *int {
		if v == nil {
			return nil
		}
		return v.ExpiresIn
	}).(pulumi.IntPtrOutput)
}

func (o SourceControlAuthInfoResponsePtrOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlAuthInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.RefreshToken
	}).(pulumi.StringPtrOutput)
}

func (o SourceControlAuthInfoResponsePtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlAuthInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

func (o SourceControlAuthInfoResponsePtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlAuthInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Token
	}).(pulumi.StringPtrOutput)
}

func (o SourceControlAuthInfoResponsePtrOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlAuthInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.TokenType
	}).(pulumi.StringPtrOutput)
}

type SourceRepositoryProperties struct {
	IsCommitTriggerEnabled      *bool                  `pulumi:"isCommitTriggerEnabled"`
	RepositoryUrl               string                 `pulumi:"repositoryUrl"`
	SourceControlAuthProperties *SourceControlAuthInfo `pulumi:"sourceControlAuthProperties"`
	SourceControlType           string                 `pulumi:"sourceControlType"`
}


func (val *SourceRepositoryProperties) Defaults() *SourceRepositoryProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsCommitTriggerEnabled) {
		isCommitTriggerEnabled_ := false
		tmp.IsCommitTriggerEnabled = &isCommitTriggerEnabled_
	}
	return &tmp
}





type SourceRepositoryPropertiesInput interface {
	pulumi.Input

	ToSourceRepositoryPropertiesOutput() SourceRepositoryPropertiesOutput
	ToSourceRepositoryPropertiesOutputWithContext(context.Context) SourceRepositoryPropertiesOutput
}

type SourceRepositoryPropertiesArgs struct {
	IsCommitTriggerEnabled      pulumi.BoolPtrInput           `pulumi:"isCommitTriggerEnabled"`
	RepositoryUrl               pulumi.StringInput            `pulumi:"repositoryUrl"`
	SourceControlAuthProperties SourceControlAuthInfoPtrInput `pulumi:"sourceControlAuthProperties"`
	SourceControlType           pulumi.StringInput            `pulumi:"sourceControlType"`
}


func (val *SourceRepositoryPropertiesArgs) Defaults() *SourceRepositoryPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsCommitTriggerEnabled) {
		tmp.IsCommitTriggerEnabled = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (SourceRepositoryPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceRepositoryProperties)(nil)).Elem()
}

func (i SourceRepositoryPropertiesArgs) ToSourceRepositoryPropertiesOutput() SourceRepositoryPropertiesOutput {
	return i.ToSourceRepositoryPropertiesOutputWithContext(context.Background())
}

func (i SourceRepositoryPropertiesArgs) ToSourceRepositoryPropertiesOutputWithContext(ctx context.Context) SourceRepositoryPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRepositoryPropertiesOutput)
}

type SourceRepositoryPropertiesOutput struct{ *pulumi.OutputState }

func (SourceRepositoryPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceRepositoryProperties)(nil)).Elem()
}

func (o SourceRepositoryPropertiesOutput) ToSourceRepositoryPropertiesOutput() SourceRepositoryPropertiesOutput {
	return o
}

func (o SourceRepositoryPropertiesOutput) ToSourceRepositoryPropertiesOutputWithContext(ctx context.Context) SourceRepositoryPropertiesOutput {
	return o
}

func (o SourceRepositoryPropertiesOutput) IsCommitTriggerEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SourceRepositoryProperties) *bool { return v.IsCommitTriggerEnabled }).(pulumi.BoolPtrOutput)
}

func (o SourceRepositoryPropertiesOutput) RepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v SourceRepositoryProperties) string { return v.RepositoryUrl }).(pulumi.StringOutput)
}

func (o SourceRepositoryPropertiesOutput) SourceControlAuthProperties() SourceControlAuthInfoPtrOutput {
	return o.ApplyT(func(v SourceRepositoryProperties) *SourceControlAuthInfo { return v.SourceControlAuthProperties }).(SourceControlAuthInfoPtrOutput)
}

func (o SourceRepositoryPropertiesOutput) SourceControlType() pulumi.StringOutput {
	return o.ApplyT(func(v SourceRepositoryProperties) string { return v.SourceControlType }).(pulumi.StringOutput)
}

type SourceRepositoryPropertiesResponse struct {
	IsCommitTriggerEnabled      *bool                          `pulumi:"isCommitTriggerEnabled"`
	RepositoryUrl               string                         `pulumi:"repositoryUrl"`
	SourceControlAuthProperties *SourceControlAuthInfoResponse `pulumi:"sourceControlAuthProperties"`
	SourceControlType           string                         `pulumi:"sourceControlType"`
}


func (val *SourceRepositoryPropertiesResponse) Defaults() *SourceRepositoryPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsCommitTriggerEnabled) {
		isCommitTriggerEnabled_ := false
		tmp.IsCommitTriggerEnabled = &isCommitTriggerEnabled_
	}
	return &tmp
}

type SourceRepositoryPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SourceRepositoryPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceRepositoryPropertiesResponse)(nil)).Elem()
}

func (o SourceRepositoryPropertiesResponseOutput) ToSourceRepositoryPropertiesResponseOutput() SourceRepositoryPropertiesResponseOutput {
	return o
}

func (o SourceRepositoryPropertiesResponseOutput) ToSourceRepositoryPropertiesResponseOutputWithContext(ctx context.Context) SourceRepositoryPropertiesResponseOutput {
	return o
}

func (o SourceRepositoryPropertiesResponseOutput) IsCommitTriggerEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SourceRepositoryPropertiesResponse) *bool { return v.IsCommitTriggerEnabled }).(pulumi.BoolPtrOutput)
}

func (o SourceRepositoryPropertiesResponseOutput) RepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v SourceRepositoryPropertiesResponse) string { return v.RepositoryUrl }).(pulumi.StringOutput)
}

func (o SourceRepositoryPropertiesResponseOutput) SourceControlAuthProperties() SourceControlAuthInfoResponsePtrOutput {
	return o.ApplyT(func(v SourceRepositoryPropertiesResponse) *SourceControlAuthInfoResponse {
		return v.SourceControlAuthProperties
	}).(SourceControlAuthInfoResponsePtrOutput)
}

func (o SourceRepositoryPropertiesResponseOutput) SourceControlType() pulumi.StringOutput {
	return o.ApplyT(func(v SourceRepositoryPropertiesResponse) string { return v.SourceControlType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BaseImageDependencyResponseOutput{})
	pulumi.RegisterOutputType(BaseImageDependencyResponseArrayOutput{})
	pulumi.RegisterOutputType(BuildArgumentResponseOutput{})
	pulumi.RegisterOutputType(BuildArgumentResponseArrayOutput{})
	pulumi.RegisterOutputType(DockerBuildStepResponseOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SourceControlAuthInfoOutput{})
	pulumi.RegisterOutputType(SourceControlAuthInfoPtrOutput{})
	pulumi.RegisterOutputType(SourceControlAuthInfoResponseOutput{})
	pulumi.RegisterOutputType(SourceControlAuthInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceRepositoryPropertiesOutput{})
	pulumi.RegisterOutputType(SourceRepositoryPropertiesResponseOutput{})
}
