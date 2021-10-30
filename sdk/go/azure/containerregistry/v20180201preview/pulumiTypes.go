


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





type BaseImageDependencyResponseInput interface {
	pulumi.Input

	ToBaseImageDependencyResponseOutput() BaseImageDependencyResponseOutput
	ToBaseImageDependencyResponseOutputWithContext(context.Context) BaseImageDependencyResponseOutput
}

type BaseImageDependencyResponseArgs struct {
	Digest     pulumi.StringPtrInput `pulumi:"digest"`
	Registry   pulumi.StringPtrInput `pulumi:"registry"`
	Repository pulumi.StringPtrInput `pulumi:"repository"`
	Tag        pulumi.StringPtrInput `pulumi:"tag"`
	Type       pulumi.StringPtrInput `pulumi:"type"`
}

func (BaseImageDependencyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BaseImageDependencyResponse)(nil)).Elem()
}

func (i BaseImageDependencyResponseArgs) ToBaseImageDependencyResponseOutput() BaseImageDependencyResponseOutput {
	return i.ToBaseImageDependencyResponseOutputWithContext(context.Background())
}

func (i BaseImageDependencyResponseArgs) ToBaseImageDependencyResponseOutputWithContext(ctx context.Context) BaseImageDependencyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseImageDependencyResponseOutput)
}





type BaseImageDependencyResponseArrayInput interface {
	pulumi.Input

	ToBaseImageDependencyResponseArrayOutput() BaseImageDependencyResponseArrayOutput
	ToBaseImageDependencyResponseArrayOutputWithContext(context.Context) BaseImageDependencyResponseArrayOutput
}

type BaseImageDependencyResponseArray []BaseImageDependencyResponseInput

func (BaseImageDependencyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BaseImageDependencyResponse)(nil)).Elem()
}

func (i BaseImageDependencyResponseArray) ToBaseImageDependencyResponseArrayOutput() BaseImageDependencyResponseArrayOutput {
	return i.ToBaseImageDependencyResponseArrayOutputWithContext(context.Background())
}

func (i BaseImageDependencyResponseArray) ToBaseImageDependencyResponseArrayOutputWithContext(ctx context.Context) BaseImageDependencyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaseImageDependencyResponseArrayOutput)
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





type BuildArgumentResponseInput interface {
	pulumi.Input

	ToBuildArgumentResponseOutput() BuildArgumentResponseOutput
	ToBuildArgumentResponseOutputWithContext(context.Context) BuildArgumentResponseOutput
}

type BuildArgumentResponseArgs struct {
	IsSecret pulumi.BoolPtrInput `pulumi:"isSecret"`
	Name     pulumi.StringInput  `pulumi:"name"`
	Type     pulumi.StringInput  `pulumi:"type"`
	Value    pulumi.StringInput  `pulumi:"value"`
}

func (BuildArgumentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildArgumentResponse)(nil)).Elem()
}

func (i BuildArgumentResponseArgs) ToBuildArgumentResponseOutput() BuildArgumentResponseOutput {
	return i.ToBuildArgumentResponseOutputWithContext(context.Background())
}

func (i BuildArgumentResponseArgs) ToBuildArgumentResponseOutputWithContext(ctx context.Context) BuildArgumentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildArgumentResponseOutput)
}





type BuildArgumentResponseArrayInput interface {
	pulumi.Input

	ToBuildArgumentResponseArrayOutput() BuildArgumentResponseArrayOutput
	ToBuildArgumentResponseArrayOutputWithContext(context.Context) BuildArgumentResponseArrayOutput
}

type BuildArgumentResponseArray []BuildArgumentResponseInput

func (BuildArgumentResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildArgumentResponse)(nil)).Elem()
}

func (i BuildArgumentResponseArray) ToBuildArgumentResponseArrayOutput() BuildArgumentResponseArrayOutput {
	return i.ToBuildArgumentResponseArrayOutputWithContext(context.Background())
}

func (i BuildArgumentResponseArray) ToBuildArgumentResponseArrayOutputWithContext(ctx context.Context) BuildArgumentResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildArgumentResponseArrayOutput)
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





type DockerBuildStepResponseInput interface {
	pulumi.Input

	ToDockerBuildStepResponseOutput() DockerBuildStepResponseOutput
	ToDockerBuildStepResponseOutputWithContext(context.Context) DockerBuildStepResponseOutput
}

type DockerBuildStepResponseArgs struct {
	BaseImageDependencies BaseImageDependencyResponseArrayInput `pulumi:"baseImageDependencies"`
	BaseImageTrigger      pulumi.StringPtrInput                 `pulumi:"baseImageTrigger"`
	Branch                pulumi.StringPtrInput                 `pulumi:"branch"`
	BuildArguments        BuildArgumentResponseArrayInput       `pulumi:"buildArguments"`
	ContextPath           pulumi.StringPtrInput                 `pulumi:"contextPath"`
	DockerFilePath        pulumi.StringPtrInput                 `pulumi:"dockerFilePath"`
	ImageNames            pulumi.StringArrayInput               `pulumi:"imageNames"`
	IsPushEnabled         pulumi.BoolPtrInput                   `pulumi:"isPushEnabled"`
	NoCache               pulumi.BoolPtrInput                   `pulumi:"noCache"`
	ProvisioningState     pulumi.StringInput                    `pulumi:"provisioningState"`
	Type                  pulumi.StringInput                    `pulumi:"type"`
}

func (DockerBuildStepResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuildStepResponse)(nil)).Elem()
}

func (i DockerBuildStepResponseArgs) ToDockerBuildStepResponseOutput() DockerBuildStepResponseOutput {
	return i.ToDockerBuildStepResponseOutputWithContext(context.Background())
}

func (i DockerBuildStepResponseArgs) ToDockerBuildStepResponseOutputWithContext(ctx context.Context) DockerBuildStepResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerBuildStepResponseOutput)
}

func (i DockerBuildStepResponseArgs) ToDockerBuildStepResponsePtrOutput() DockerBuildStepResponsePtrOutput {
	return i.ToDockerBuildStepResponsePtrOutputWithContext(context.Background())
}

func (i DockerBuildStepResponseArgs) ToDockerBuildStepResponsePtrOutputWithContext(ctx context.Context) DockerBuildStepResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerBuildStepResponseOutput).ToDockerBuildStepResponsePtrOutputWithContext(ctx)
}









type DockerBuildStepResponsePtrInput interface {
	pulumi.Input

	ToDockerBuildStepResponsePtrOutput() DockerBuildStepResponsePtrOutput
	ToDockerBuildStepResponsePtrOutputWithContext(context.Context) DockerBuildStepResponsePtrOutput
}

type dockerBuildStepResponsePtrType DockerBuildStepResponseArgs

func DockerBuildStepResponsePtr(v *DockerBuildStepResponseArgs) DockerBuildStepResponsePtrInput {
	return (*dockerBuildStepResponsePtrType)(v)
}

func (*dockerBuildStepResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerBuildStepResponse)(nil)).Elem()
}

func (i *dockerBuildStepResponsePtrType) ToDockerBuildStepResponsePtrOutput() DockerBuildStepResponsePtrOutput {
	return i.ToDockerBuildStepResponsePtrOutputWithContext(context.Background())
}

func (i *dockerBuildStepResponsePtrType) ToDockerBuildStepResponsePtrOutputWithContext(ctx context.Context) DockerBuildStepResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerBuildStepResponsePtrOutput)
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

func (o DockerBuildStepResponseOutput) ToDockerBuildStepResponsePtrOutput() DockerBuildStepResponsePtrOutput {
	return o.ToDockerBuildStepResponsePtrOutputWithContext(context.Background())
}

func (o DockerBuildStepResponseOutput) ToDockerBuildStepResponsePtrOutputWithContext(ctx context.Context) DockerBuildStepResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DockerBuildStepResponse) *DockerBuildStepResponse {
		return &v
	}).(DockerBuildStepResponsePtrOutput)
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

type DockerBuildStepResponsePtrOutput struct{ *pulumi.OutputState }

func (DockerBuildStepResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerBuildStepResponse)(nil)).Elem()
}

func (o DockerBuildStepResponsePtrOutput) ToDockerBuildStepResponsePtrOutput() DockerBuildStepResponsePtrOutput {
	return o
}

func (o DockerBuildStepResponsePtrOutput) ToDockerBuildStepResponsePtrOutputWithContext(ctx context.Context) DockerBuildStepResponsePtrOutput {
	return o
}

func (o DockerBuildStepResponsePtrOutput) Elem() DockerBuildStepResponseOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) DockerBuildStepResponse {
		if v != nil {
			return *v
		}
		var ret DockerBuildStepResponse
		return ret
	}).(DockerBuildStepResponseOutput)
}

func (o DockerBuildStepResponsePtrOutput) BaseImageDependencies() BaseImageDependencyResponseArrayOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) []BaseImageDependencyResponse {
		if v == nil {
			return nil
		}
		return v.BaseImageDependencies
	}).(BaseImageDependencyResponseArrayOutput)
}

func (o DockerBuildStepResponsePtrOutput) BaseImageTrigger() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) *string {
		if v == nil {
			return nil
		}
		return v.BaseImageTrigger
	}).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponsePtrOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) *string {
		if v == nil {
			return nil
		}
		return v.Branch
	}).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponsePtrOutput) BuildArguments() BuildArgumentResponseArrayOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) []BuildArgumentResponse {
		if v == nil {
			return nil
		}
		return v.BuildArguments
	}).(BuildArgumentResponseArrayOutput)
}

func (o DockerBuildStepResponsePtrOutput) ContextPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContextPath
	}).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponsePtrOutput) DockerFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) *string {
		if v == nil {
			return nil
		}
		return v.DockerFilePath
	}).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponsePtrOutput) ImageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) []string {
		if v == nil {
			return nil
		}
		return v.ImageNames
	}).(pulumi.StringArrayOutput)
}

func (o DockerBuildStepResponsePtrOutput) IsPushEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsPushEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o DockerBuildStepResponsePtrOutput) NoCache() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) *bool {
		if v == nil {
			return nil
		}
		return v.NoCache
	}).(pulumi.BoolPtrOutput)
}

func (o DockerBuildStepResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o DockerBuildStepResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DockerBuildStepResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
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

func (i PlatformPropertiesArgs) ToPlatformPropertiesPtrOutput() PlatformPropertiesPtrOutput {
	return i.ToPlatformPropertiesPtrOutputWithContext(context.Background())
}

func (i PlatformPropertiesArgs) ToPlatformPropertiesPtrOutputWithContext(ctx context.Context) PlatformPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesOutput).ToPlatformPropertiesPtrOutputWithContext(ctx)
}









type PlatformPropertiesPtrInput interface {
	pulumi.Input

	ToPlatformPropertiesPtrOutput() PlatformPropertiesPtrOutput
	ToPlatformPropertiesPtrOutputWithContext(context.Context) PlatformPropertiesPtrOutput
}

type platformPropertiesPtrType PlatformPropertiesArgs

func PlatformPropertiesPtr(v *PlatformPropertiesArgs) PlatformPropertiesPtrInput {
	return (*platformPropertiesPtrType)(v)
}

func (*platformPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformProperties)(nil)).Elem()
}

func (i *platformPropertiesPtrType) ToPlatformPropertiesPtrOutput() PlatformPropertiesPtrOutput {
	return i.ToPlatformPropertiesPtrOutputWithContext(context.Background())
}

func (i *platformPropertiesPtrType) ToPlatformPropertiesPtrOutputWithContext(ctx context.Context) PlatformPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesPtrOutput)
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

func (o PlatformPropertiesOutput) ToPlatformPropertiesPtrOutput() PlatformPropertiesPtrOutput {
	return o.ToPlatformPropertiesPtrOutputWithContext(context.Background())
}

func (o PlatformPropertiesOutput) ToPlatformPropertiesPtrOutputWithContext(ctx context.Context) PlatformPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlatformProperties) *PlatformProperties {
		return &v
	}).(PlatformPropertiesPtrOutput)
}

func (o PlatformPropertiesOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PlatformProperties) *int { return v.Cpu }).(pulumi.IntPtrOutput)
}

func (o PlatformPropertiesOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformProperties) string { return v.OsType }).(pulumi.StringOutput)
}

type PlatformPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PlatformPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformProperties)(nil)).Elem()
}

func (o PlatformPropertiesPtrOutput) ToPlatformPropertiesPtrOutput() PlatformPropertiesPtrOutput {
	return o
}

func (o PlatformPropertiesPtrOutput) ToPlatformPropertiesPtrOutputWithContext(ctx context.Context) PlatformPropertiesPtrOutput {
	return o
}

func (o PlatformPropertiesPtrOutput) Elem() PlatformPropertiesOutput {
	return o.ApplyT(func(v *PlatformProperties) PlatformProperties {
		if v != nil {
			return *v
		}
		var ret PlatformProperties
		return ret
	}).(PlatformPropertiesOutput)
}

func (o PlatformPropertiesPtrOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PlatformProperties) *int {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.IntPtrOutput)
}

func (o PlatformPropertiesPtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformProperties) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

type PlatformPropertiesResponse struct {
	Cpu    *int   `pulumi:"cpu"`
	OsType string `pulumi:"osType"`
}





type PlatformPropertiesResponseInput interface {
	pulumi.Input

	ToPlatformPropertiesResponseOutput() PlatformPropertiesResponseOutput
	ToPlatformPropertiesResponseOutputWithContext(context.Context) PlatformPropertiesResponseOutput
}

type PlatformPropertiesResponseArgs struct {
	Cpu    pulumi.IntPtrInput `pulumi:"cpu"`
	OsType pulumi.StringInput `pulumi:"osType"`
}

func (PlatformPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformPropertiesResponse)(nil)).Elem()
}

func (i PlatformPropertiesResponseArgs) ToPlatformPropertiesResponseOutput() PlatformPropertiesResponseOutput {
	return i.ToPlatformPropertiesResponseOutputWithContext(context.Background())
}

func (i PlatformPropertiesResponseArgs) ToPlatformPropertiesResponseOutputWithContext(ctx context.Context) PlatformPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesResponseOutput)
}

func (i PlatformPropertiesResponseArgs) ToPlatformPropertiesResponsePtrOutput() PlatformPropertiesResponsePtrOutput {
	return i.ToPlatformPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PlatformPropertiesResponseArgs) ToPlatformPropertiesResponsePtrOutputWithContext(ctx context.Context) PlatformPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesResponseOutput).ToPlatformPropertiesResponsePtrOutputWithContext(ctx)
}









type PlatformPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPlatformPropertiesResponsePtrOutput() PlatformPropertiesResponsePtrOutput
	ToPlatformPropertiesResponsePtrOutputWithContext(context.Context) PlatformPropertiesResponsePtrOutput
}

type platformPropertiesResponsePtrType PlatformPropertiesResponseArgs

func PlatformPropertiesResponsePtr(v *PlatformPropertiesResponseArgs) PlatformPropertiesResponsePtrInput {
	return (*platformPropertiesResponsePtrType)(v)
}

func (*platformPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformPropertiesResponse)(nil)).Elem()
}

func (i *platformPropertiesResponsePtrType) ToPlatformPropertiesResponsePtrOutput() PlatformPropertiesResponsePtrOutput {
	return i.ToPlatformPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *platformPropertiesResponsePtrType) ToPlatformPropertiesResponsePtrOutputWithContext(ctx context.Context) PlatformPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformPropertiesResponsePtrOutput)
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

func (o PlatformPropertiesResponseOutput) ToPlatformPropertiesResponsePtrOutput() PlatformPropertiesResponsePtrOutput {
	return o.ToPlatformPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PlatformPropertiesResponseOutput) ToPlatformPropertiesResponsePtrOutputWithContext(ctx context.Context) PlatformPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlatformPropertiesResponse) *PlatformPropertiesResponse {
		return &v
	}).(PlatformPropertiesResponsePtrOutput)
}

func (o PlatformPropertiesResponseOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PlatformPropertiesResponse) *int { return v.Cpu }).(pulumi.IntPtrOutput)
}

func (o PlatformPropertiesResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformPropertiesResponse) string { return v.OsType }).(pulumi.StringOutput)
}

type PlatformPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PlatformPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformPropertiesResponse)(nil)).Elem()
}

func (o PlatformPropertiesResponsePtrOutput) ToPlatformPropertiesResponsePtrOutput() PlatformPropertiesResponsePtrOutput {
	return o
}

func (o PlatformPropertiesResponsePtrOutput) ToPlatformPropertiesResponsePtrOutputWithContext(ctx context.Context) PlatformPropertiesResponsePtrOutput {
	return o
}

func (o PlatformPropertiesResponsePtrOutput) Elem() PlatformPropertiesResponseOutput {
	return o.ApplyT(func(v *PlatformPropertiesResponse) PlatformPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PlatformPropertiesResponse
		return ret
	}).(PlatformPropertiesResponseOutput)
}

func (o PlatformPropertiesResponsePtrOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PlatformPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.IntPtrOutput)
}

func (o PlatformPropertiesResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
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





type SourceControlAuthInfoResponseInput interface {
	pulumi.Input

	ToSourceControlAuthInfoResponseOutput() SourceControlAuthInfoResponseOutput
	ToSourceControlAuthInfoResponseOutputWithContext(context.Context) SourceControlAuthInfoResponseOutput
}

type SourceControlAuthInfoResponseArgs struct {
	ExpiresIn    pulumi.IntPtrInput    `pulumi:"expiresIn"`
	RefreshToken pulumi.StringPtrInput `pulumi:"refreshToken"`
	Scope        pulumi.StringPtrInput `pulumi:"scope"`
	Token        pulumi.StringInput    `pulumi:"token"`
	TokenType    pulumi.StringPtrInput `pulumi:"tokenType"`
}

func (SourceControlAuthInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControlAuthInfoResponse)(nil)).Elem()
}

func (i SourceControlAuthInfoResponseArgs) ToSourceControlAuthInfoResponseOutput() SourceControlAuthInfoResponseOutput {
	return i.ToSourceControlAuthInfoResponseOutputWithContext(context.Background())
}

func (i SourceControlAuthInfoResponseArgs) ToSourceControlAuthInfoResponseOutputWithContext(ctx context.Context) SourceControlAuthInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlAuthInfoResponseOutput)
}

func (i SourceControlAuthInfoResponseArgs) ToSourceControlAuthInfoResponsePtrOutput() SourceControlAuthInfoResponsePtrOutput {
	return i.ToSourceControlAuthInfoResponsePtrOutputWithContext(context.Background())
}

func (i SourceControlAuthInfoResponseArgs) ToSourceControlAuthInfoResponsePtrOutputWithContext(ctx context.Context) SourceControlAuthInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlAuthInfoResponseOutput).ToSourceControlAuthInfoResponsePtrOutputWithContext(ctx)
}









type SourceControlAuthInfoResponsePtrInput interface {
	pulumi.Input

	ToSourceControlAuthInfoResponsePtrOutput() SourceControlAuthInfoResponsePtrOutput
	ToSourceControlAuthInfoResponsePtrOutputWithContext(context.Context) SourceControlAuthInfoResponsePtrOutput
}

type sourceControlAuthInfoResponsePtrType SourceControlAuthInfoResponseArgs

func SourceControlAuthInfoResponsePtr(v *SourceControlAuthInfoResponseArgs) SourceControlAuthInfoResponsePtrInput {
	return (*sourceControlAuthInfoResponsePtrType)(v)
}

func (*sourceControlAuthInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControlAuthInfoResponse)(nil)).Elem()
}

func (i *sourceControlAuthInfoResponsePtrType) ToSourceControlAuthInfoResponsePtrOutput() SourceControlAuthInfoResponsePtrOutput {
	return i.ToSourceControlAuthInfoResponsePtrOutputWithContext(context.Background())
}

func (i *sourceControlAuthInfoResponsePtrType) ToSourceControlAuthInfoResponsePtrOutputWithContext(ctx context.Context) SourceControlAuthInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlAuthInfoResponsePtrOutput)
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

func (o SourceControlAuthInfoResponseOutput) ToSourceControlAuthInfoResponsePtrOutput() SourceControlAuthInfoResponsePtrOutput {
	return o.ToSourceControlAuthInfoResponsePtrOutputWithContext(context.Background())
}

func (o SourceControlAuthInfoResponseOutput) ToSourceControlAuthInfoResponsePtrOutputWithContext(ctx context.Context) SourceControlAuthInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceControlAuthInfoResponse) *SourceControlAuthInfoResponse {
		return &v
	}).(SourceControlAuthInfoResponsePtrOutput)
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

func (SourceRepositoryPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceRepositoryProperties)(nil)).Elem()
}

func (i SourceRepositoryPropertiesArgs) ToSourceRepositoryPropertiesOutput() SourceRepositoryPropertiesOutput {
	return i.ToSourceRepositoryPropertiesOutputWithContext(context.Background())
}

func (i SourceRepositoryPropertiesArgs) ToSourceRepositoryPropertiesOutputWithContext(ctx context.Context) SourceRepositoryPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRepositoryPropertiesOutput)
}

func (i SourceRepositoryPropertiesArgs) ToSourceRepositoryPropertiesPtrOutput() SourceRepositoryPropertiesPtrOutput {
	return i.ToSourceRepositoryPropertiesPtrOutputWithContext(context.Background())
}

func (i SourceRepositoryPropertiesArgs) ToSourceRepositoryPropertiesPtrOutputWithContext(ctx context.Context) SourceRepositoryPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRepositoryPropertiesOutput).ToSourceRepositoryPropertiesPtrOutputWithContext(ctx)
}









type SourceRepositoryPropertiesPtrInput interface {
	pulumi.Input

	ToSourceRepositoryPropertiesPtrOutput() SourceRepositoryPropertiesPtrOutput
	ToSourceRepositoryPropertiesPtrOutputWithContext(context.Context) SourceRepositoryPropertiesPtrOutput
}

type sourceRepositoryPropertiesPtrType SourceRepositoryPropertiesArgs

func SourceRepositoryPropertiesPtr(v *SourceRepositoryPropertiesArgs) SourceRepositoryPropertiesPtrInput {
	return (*sourceRepositoryPropertiesPtrType)(v)
}

func (*sourceRepositoryPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRepositoryProperties)(nil)).Elem()
}

func (i *sourceRepositoryPropertiesPtrType) ToSourceRepositoryPropertiesPtrOutput() SourceRepositoryPropertiesPtrOutput {
	return i.ToSourceRepositoryPropertiesPtrOutputWithContext(context.Background())
}

func (i *sourceRepositoryPropertiesPtrType) ToSourceRepositoryPropertiesPtrOutputWithContext(ctx context.Context) SourceRepositoryPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRepositoryPropertiesPtrOutput)
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

func (o SourceRepositoryPropertiesOutput) ToSourceRepositoryPropertiesPtrOutput() SourceRepositoryPropertiesPtrOutput {
	return o.ToSourceRepositoryPropertiesPtrOutputWithContext(context.Background())
}

func (o SourceRepositoryPropertiesOutput) ToSourceRepositoryPropertiesPtrOutputWithContext(ctx context.Context) SourceRepositoryPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceRepositoryProperties) *SourceRepositoryProperties {
		return &v
	}).(SourceRepositoryPropertiesPtrOutput)
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

type SourceRepositoryPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SourceRepositoryPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRepositoryProperties)(nil)).Elem()
}

func (o SourceRepositoryPropertiesPtrOutput) ToSourceRepositoryPropertiesPtrOutput() SourceRepositoryPropertiesPtrOutput {
	return o
}

func (o SourceRepositoryPropertiesPtrOutput) ToSourceRepositoryPropertiesPtrOutputWithContext(ctx context.Context) SourceRepositoryPropertiesPtrOutput {
	return o
}

func (o SourceRepositoryPropertiesPtrOutput) Elem() SourceRepositoryPropertiesOutput {
	return o.ApplyT(func(v *SourceRepositoryProperties) SourceRepositoryProperties {
		if v != nil {
			return *v
		}
		var ret SourceRepositoryProperties
		return ret
	}).(SourceRepositoryPropertiesOutput)
}

func (o SourceRepositoryPropertiesPtrOutput) IsCommitTriggerEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourceRepositoryProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsCommitTriggerEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SourceRepositoryPropertiesPtrOutput) RepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceRepositoryProperties) *string {
		if v == nil {
			return nil
		}
		return &v.RepositoryUrl
	}).(pulumi.StringPtrOutput)
}

func (o SourceRepositoryPropertiesPtrOutput) SourceControlAuthProperties() SourceControlAuthInfoPtrOutput {
	return o.ApplyT(func(v *SourceRepositoryProperties) *SourceControlAuthInfo {
		if v == nil {
			return nil
		}
		return v.SourceControlAuthProperties
	}).(SourceControlAuthInfoPtrOutput)
}

func (o SourceRepositoryPropertiesPtrOutput) SourceControlType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceRepositoryProperties) *string {
		if v == nil {
			return nil
		}
		return &v.SourceControlType
	}).(pulumi.StringPtrOutput)
}

type SourceRepositoryPropertiesResponse struct {
	IsCommitTriggerEnabled      *bool                          `pulumi:"isCommitTriggerEnabled"`
	RepositoryUrl               string                         `pulumi:"repositoryUrl"`
	SourceControlAuthProperties *SourceControlAuthInfoResponse `pulumi:"sourceControlAuthProperties"`
	SourceControlType           string                         `pulumi:"sourceControlType"`
}





type SourceRepositoryPropertiesResponseInput interface {
	pulumi.Input

	ToSourceRepositoryPropertiesResponseOutput() SourceRepositoryPropertiesResponseOutput
	ToSourceRepositoryPropertiesResponseOutputWithContext(context.Context) SourceRepositoryPropertiesResponseOutput
}

type SourceRepositoryPropertiesResponseArgs struct {
	IsCommitTriggerEnabled      pulumi.BoolPtrInput                   `pulumi:"isCommitTriggerEnabled"`
	RepositoryUrl               pulumi.StringInput                    `pulumi:"repositoryUrl"`
	SourceControlAuthProperties SourceControlAuthInfoResponsePtrInput `pulumi:"sourceControlAuthProperties"`
	SourceControlType           pulumi.StringInput                    `pulumi:"sourceControlType"`
}

func (SourceRepositoryPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceRepositoryPropertiesResponse)(nil)).Elem()
}

func (i SourceRepositoryPropertiesResponseArgs) ToSourceRepositoryPropertiesResponseOutput() SourceRepositoryPropertiesResponseOutput {
	return i.ToSourceRepositoryPropertiesResponseOutputWithContext(context.Background())
}

func (i SourceRepositoryPropertiesResponseArgs) ToSourceRepositoryPropertiesResponseOutputWithContext(ctx context.Context) SourceRepositoryPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRepositoryPropertiesResponseOutput)
}

func (i SourceRepositoryPropertiesResponseArgs) ToSourceRepositoryPropertiesResponsePtrOutput() SourceRepositoryPropertiesResponsePtrOutput {
	return i.ToSourceRepositoryPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SourceRepositoryPropertiesResponseArgs) ToSourceRepositoryPropertiesResponsePtrOutputWithContext(ctx context.Context) SourceRepositoryPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRepositoryPropertiesResponseOutput).ToSourceRepositoryPropertiesResponsePtrOutputWithContext(ctx)
}









type SourceRepositoryPropertiesResponsePtrInput interface {
	pulumi.Input

	ToSourceRepositoryPropertiesResponsePtrOutput() SourceRepositoryPropertiesResponsePtrOutput
	ToSourceRepositoryPropertiesResponsePtrOutputWithContext(context.Context) SourceRepositoryPropertiesResponsePtrOutput
}

type sourceRepositoryPropertiesResponsePtrType SourceRepositoryPropertiesResponseArgs

func SourceRepositoryPropertiesResponsePtr(v *SourceRepositoryPropertiesResponseArgs) SourceRepositoryPropertiesResponsePtrInput {
	return (*sourceRepositoryPropertiesResponsePtrType)(v)
}

func (*sourceRepositoryPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRepositoryPropertiesResponse)(nil)).Elem()
}

func (i *sourceRepositoryPropertiesResponsePtrType) ToSourceRepositoryPropertiesResponsePtrOutput() SourceRepositoryPropertiesResponsePtrOutput {
	return i.ToSourceRepositoryPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *sourceRepositoryPropertiesResponsePtrType) ToSourceRepositoryPropertiesResponsePtrOutputWithContext(ctx context.Context) SourceRepositoryPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceRepositoryPropertiesResponsePtrOutput)
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

func (o SourceRepositoryPropertiesResponseOutput) ToSourceRepositoryPropertiesResponsePtrOutput() SourceRepositoryPropertiesResponsePtrOutput {
	return o.ToSourceRepositoryPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SourceRepositoryPropertiesResponseOutput) ToSourceRepositoryPropertiesResponsePtrOutputWithContext(ctx context.Context) SourceRepositoryPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceRepositoryPropertiesResponse) *SourceRepositoryPropertiesResponse {
		return &v
	}).(SourceRepositoryPropertiesResponsePtrOutput)
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

type SourceRepositoryPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SourceRepositoryPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRepositoryPropertiesResponse)(nil)).Elem()
}

func (o SourceRepositoryPropertiesResponsePtrOutput) ToSourceRepositoryPropertiesResponsePtrOutput() SourceRepositoryPropertiesResponsePtrOutput {
	return o
}

func (o SourceRepositoryPropertiesResponsePtrOutput) ToSourceRepositoryPropertiesResponsePtrOutputWithContext(ctx context.Context) SourceRepositoryPropertiesResponsePtrOutput {
	return o
}

func (o SourceRepositoryPropertiesResponsePtrOutput) Elem() SourceRepositoryPropertiesResponseOutput {
	return o.ApplyT(func(v *SourceRepositoryPropertiesResponse) SourceRepositoryPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SourceRepositoryPropertiesResponse
		return ret
	}).(SourceRepositoryPropertiesResponseOutput)
}

func (o SourceRepositoryPropertiesResponsePtrOutput) IsCommitTriggerEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourceRepositoryPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCommitTriggerEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SourceRepositoryPropertiesResponsePtrOutput) RepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceRepositoryPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RepositoryUrl
	}).(pulumi.StringPtrOutput)
}

func (o SourceRepositoryPropertiesResponsePtrOutput) SourceControlAuthProperties() SourceControlAuthInfoResponsePtrOutput {
	return o.ApplyT(func(v *SourceRepositoryPropertiesResponse) *SourceControlAuthInfoResponse {
		if v == nil {
			return nil
		}
		return v.SourceControlAuthProperties
	}).(SourceControlAuthInfoResponsePtrOutput)
}

func (o SourceRepositoryPropertiesResponsePtrOutput) SourceControlType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceRepositoryPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceControlType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BaseImageDependencyResponseOutput{})
	pulumi.RegisterOutputType(BaseImageDependencyResponseArrayOutput{})
	pulumi.RegisterOutputType(BuildArgumentResponseOutput{})
	pulumi.RegisterOutputType(BuildArgumentResponseArrayOutput{})
	pulumi.RegisterOutputType(DockerBuildStepResponseOutput{})
	pulumi.RegisterOutputType(DockerBuildStepResponsePtrOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PlatformPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceControlAuthInfoOutput{})
	pulumi.RegisterOutputType(SourceControlAuthInfoPtrOutput{})
	pulumi.RegisterOutputType(SourceControlAuthInfoResponseOutput{})
	pulumi.RegisterOutputType(SourceControlAuthInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceRepositoryPropertiesOutput{})
	pulumi.RegisterOutputType(SourceRepositoryPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SourceRepositoryPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SourceRepositoryPropertiesResponsePtrOutput{})
}
