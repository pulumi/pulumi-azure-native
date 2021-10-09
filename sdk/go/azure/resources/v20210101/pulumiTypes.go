


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AliasPathMetadataResponse struct {
	Attributes string `pulumi:"attributes"`
	Type       string `pulumi:"type"`
}





type AliasPathMetadataResponseInput interface {
	pulumi.Input

	ToAliasPathMetadataResponseOutput() AliasPathMetadataResponseOutput
	ToAliasPathMetadataResponseOutputWithContext(context.Context) AliasPathMetadataResponseOutput
}

type AliasPathMetadataResponseArgs struct {
	Attributes pulumi.StringInput `pulumi:"attributes"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (AliasPathMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasPathMetadataResponse)(nil)).Elem()
}

func (i AliasPathMetadataResponseArgs) ToAliasPathMetadataResponseOutput() AliasPathMetadataResponseOutput {
	return i.ToAliasPathMetadataResponseOutputWithContext(context.Background())
}

func (i AliasPathMetadataResponseArgs) ToAliasPathMetadataResponseOutputWithContext(ctx context.Context) AliasPathMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPathMetadataResponseOutput)
}

type AliasPathMetadataResponseOutput struct{ *pulumi.OutputState }

func (AliasPathMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasPathMetadataResponse)(nil)).Elem()
}

func (o AliasPathMetadataResponseOutput) ToAliasPathMetadataResponseOutput() AliasPathMetadataResponseOutput {
	return o
}

func (o AliasPathMetadataResponseOutput) ToAliasPathMetadataResponseOutputWithContext(ctx context.Context) AliasPathMetadataResponseOutput {
	return o
}

func (o AliasPathMetadataResponseOutput) Attributes() pulumi.StringOutput {
	return o.ApplyT(func(v AliasPathMetadataResponse) string { return v.Attributes }).(pulumi.StringOutput)
}

func (o AliasPathMetadataResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AliasPathMetadataResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AliasPathResponse struct {
	ApiVersions []string                  `pulumi:"apiVersions"`
	Metadata    AliasPathMetadataResponse `pulumi:"metadata"`
	Path        *string                   `pulumi:"path"`
	Pattern     *AliasPatternResponse     `pulumi:"pattern"`
}





type AliasPathResponseInput interface {
	pulumi.Input

	ToAliasPathResponseOutput() AliasPathResponseOutput
	ToAliasPathResponseOutputWithContext(context.Context) AliasPathResponseOutput
}

type AliasPathResponseArgs struct {
	ApiVersions pulumi.StringArrayInput        `pulumi:"apiVersions"`
	Metadata    AliasPathMetadataResponseInput `pulumi:"metadata"`
	Path        pulumi.StringPtrInput          `pulumi:"path"`
	Pattern     AliasPatternResponsePtrInput   `pulumi:"pattern"`
}

func (AliasPathResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasPathResponse)(nil)).Elem()
}

func (i AliasPathResponseArgs) ToAliasPathResponseOutput() AliasPathResponseOutput {
	return i.ToAliasPathResponseOutputWithContext(context.Background())
}

func (i AliasPathResponseArgs) ToAliasPathResponseOutputWithContext(ctx context.Context) AliasPathResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPathResponseOutput)
}





type AliasPathResponseArrayInput interface {
	pulumi.Input

	ToAliasPathResponseArrayOutput() AliasPathResponseArrayOutput
	ToAliasPathResponseArrayOutputWithContext(context.Context) AliasPathResponseArrayOutput
}

type AliasPathResponseArray []AliasPathResponseInput

func (AliasPathResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AliasPathResponse)(nil)).Elem()
}

func (i AliasPathResponseArray) ToAliasPathResponseArrayOutput() AliasPathResponseArrayOutput {
	return i.ToAliasPathResponseArrayOutputWithContext(context.Background())
}

func (i AliasPathResponseArray) ToAliasPathResponseArrayOutputWithContext(ctx context.Context) AliasPathResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPathResponseArrayOutput)
}

type AliasPathResponseOutput struct{ *pulumi.OutputState }

func (AliasPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasPathResponse)(nil)).Elem()
}

func (o AliasPathResponseOutput) ToAliasPathResponseOutput() AliasPathResponseOutput {
	return o
}

func (o AliasPathResponseOutput) ToAliasPathResponseOutputWithContext(ctx context.Context) AliasPathResponseOutput {
	return o
}

func (o AliasPathResponseOutput) ApiVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AliasPathResponse) []string { return v.ApiVersions }).(pulumi.StringArrayOutput)
}

func (o AliasPathResponseOutput) Metadata() AliasPathMetadataResponseOutput {
	return o.ApplyT(func(v AliasPathResponse) AliasPathMetadataResponse { return v.Metadata }).(AliasPathMetadataResponseOutput)
}

func (o AliasPathResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AliasPathResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o AliasPathResponseOutput) Pattern() AliasPatternResponsePtrOutput {
	return o.ApplyT(func(v AliasPathResponse) *AliasPatternResponse { return v.Pattern }).(AliasPatternResponsePtrOutput)
}

type AliasPathResponseArrayOutput struct{ *pulumi.OutputState }

func (AliasPathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AliasPathResponse)(nil)).Elem()
}

func (o AliasPathResponseArrayOutput) ToAliasPathResponseArrayOutput() AliasPathResponseArrayOutput {
	return o
}

func (o AliasPathResponseArrayOutput) ToAliasPathResponseArrayOutputWithContext(ctx context.Context) AliasPathResponseArrayOutput {
	return o
}

func (o AliasPathResponseArrayOutput) Index(i pulumi.IntInput) AliasPathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AliasPathResponse {
		return vs[0].([]AliasPathResponse)[vs[1].(int)]
	}).(AliasPathResponseOutput)
}

type AliasPatternResponse struct {
	Phrase   *string `pulumi:"phrase"`
	Type     *string `pulumi:"type"`
	Variable *string `pulumi:"variable"`
}





type AliasPatternResponseInput interface {
	pulumi.Input

	ToAliasPatternResponseOutput() AliasPatternResponseOutput
	ToAliasPatternResponseOutputWithContext(context.Context) AliasPatternResponseOutput
}

type AliasPatternResponseArgs struct {
	Phrase   pulumi.StringPtrInput `pulumi:"phrase"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
	Variable pulumi.StringPtrInput `pulumi:"variable"`
}

func (AliasPatternResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasPatternResponse)(nil)).Elem()
}

func (i AliasPatternResponseArgs) ToAliasPatternResponseOutput() AliasPatternResponseOutput {
	return i.ToAliasPatternResponseOutputWithContext(context.Background())
}

func (i AliasPatternResponseArgs) ToAliasPatternResponseOutputWithContext(ctx context.Context) AliasPatternResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPatternResponseOutput)
}

func (i AliasPatternResponseArgs) ToAliasPatternResponsePtrOutput() AliasPatternResponsePtrOutput {
	return i.ToAliasPatternResponsePtrOutputWithContext(context.Background())
}

func (i AliasPatternResponseArgs) ToAliasPatternResponsePtrOutputWithContext(ctx context.Context) AliasPatternResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPatternResponseOutput).ToAliasPatternResponsePtrOutputWithContext(ctx)
}









type AliasPatternResponsePtrInput interface {
	pulumi.Input

	ToAliasPatternResponsePtrOutput() AliasPatternResponsePtrOutput
	ToAliasPatternResponsePtrOutputWithContext(context.Context) AliasPatternResponsePtrOutput
}

type aliasPatternResponsePtrType AliasPatternResponseArgs

func AliasPatternResponsePtr(v *AliasPatternResponseArgs) AliasPatternResponsePtrInput {
	return (*aliasPatternResponsePtrType)(v)
}

func (*aliasPatternResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AliasPatternResponse)(nil)).Elem()
}

func (i *aliasPatternResponsePtrType) ToAliasPatternResponsePtrOutput() AliasPatternResponsePtrOutput {
	return i.ToAliasPatternResponsePtrOutputWithContext(context.Background())
}

func (i *aliasPatternResponsePtrType) ToAliasPatternResponsePtrOutputWithContext(ctx context.Context) AliasPatternResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPatternResponsePtrOutput)
}

type AliasPatternResponseOutput struct{ *pulumi.OutputState }

func (AliasPatternResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasPatternResponse)(nil)).Elem()
}

func (o AliasPatternResponseOutput) ToAliasPatternResponseOutput() AliasPatternResponseOutput {
	return o
}

func (o AliasPatternResponseOutput) ToAliasPatternResponseOutputWithContext(ctx context.Context) AliasPatternResponseOutput {
	return o
}

func (o AliasPatternResponseOutput) ToAliasPatternResponsePtrOutput() AliasPatternResponsePtrOutput {
	return o.ToAliasPatternResponsePtrOutputWithContext(context.Background())
}

func (o AliasPatternResponseOutput) ToAliasPatternResponsePtrOutputWithContext(ctx context.Context) AliasPatternResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AliasPatternResponse) *AliasPatternResponse {
		return &v
	}).(AliasPatternResponsePtrOutput)
}

func (o AliasPatternResponseOutput) Phrase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AliasPatternResponse) *string { return v.Phrase }).(pulumi.StringPtrOutput)
}

func (o AliasPatternResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AliasPatternResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o AliasPatternResponseOutput) Variable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AliasPatternResponse) *string { return v.Variable }).(pulumi.StringPtrOutput)
}

type AliasPatternResponsePtrOutput struct{ *pulumi.OutputState }

func (AliasPatternResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AliasPatternResponse)(nil)).Elem()
}

func (o AliasPatternResponsePtrOutput) ToAliasPatternResponsePtrOutput() AliasPatternResponsePtrOutput {
	return o
}

func (o AliasPatternResponsePtrOutput) ToAliasPatternResponsePtrOutputWithContext(ctx context.Context) AliasPatternResponsePtrOutput {
	return o
}

func (o AliasPatternResponsePtrOutput) Elem() AliasPatternResponseOutput {
	return o.ApplyT(func(v *AliasPatternResponse) AliasPatternResponse {
		if v != nil {
			return *v
		}
		var ret AliasPatternResponse
		return ret
	}).(AliasPatternResponseOutput)
}

func (o AliasPatternResponsePtrOutput) Phrase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AliasPatternResponse) *string {
		if v == nil {
			return nil
		}
		return v.Phrase
	}).(pulumi.StringPtrOutput)
}

func (o AliasPatternResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AliasPatternResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o AliasPatternResponsePtrOutput) Variable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AliasPatternResponse) *string {
		if v == nil {
			return nil
		}
		return v.Variable
	}).(pulumi.StringPtrOutput)
}

type AliasResponse struct {
	DefaultMetadata AliasPathMetadataResponse `pulumi:"defaultMetadata"`
	DefaultPath     *string                   `pulumi:"defaultPath"`
	DefaultPattern  *AliasPatternResponse     `pulumi:"defaultPattern"`
	Name            *string                   `pulumi:"name"`
	Paths           []AliasPathResponse       `pulumi:"paths"`
	Type            *string                   `pulumi:"type"`
}





type AliasResponseInput interface {
	pulumi.Input

	ToAliasResponseOutput() AliasResponseOutput
	ToAliasResponseOutputWithContext(context.Context) AliasResponseOutput
}

type AliasResponseArgs struct {
	DefaultMetadata AliasPathMetadataResponseInput `pulumi:"defaultMetadata"`
	DefaultPath     pulumi.StringPtrInput          `pulumi:"defaultPath"`
	DefaultPattern  AliasPatternResponsePtrInput   `pulumi:"defaultPattern"`
	Name            pulumi.StringPtrInput          `pulumi:"name"`
	Paths           AliasPathResponseArrayInput    `pulumi:"paths"`
	Type            pulumi.StringPtrInput          `pulumi:"type"`
}

func (AliasResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasResponse)(nil)).Elem()
}

func (i AliasResponseArgs) ToAliasResponseOutput() AliasResponseOutput {
	return i.ToAliasResponseOutputWithContext(context.Background())
}

func (i AliasResponseArgs) ToAliasResponseOutputWithContext(ctx context.Context) AliasResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasResponseOutput)
}





type AliasResponseArrayInput interface {
	pulumi.Input

	ToAliasResponseArrayOutput() AliasResponseArrayOutput
	ToAliasResponseArrayOutputWithContext(context.Context) AliasResponseArrayOutput
}

type AliasResponseArray []AliasResponseInput

func (AliasResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AliasResponse)(nil)).Elem()
}

func (i AliasResponseArray) ToAliasResponseArrayOutput() AliasResponseArrayOutput {
	return i.ToAliasResponseArrayOutputWithContext(context.Background())
}

func (i AliasResponseArray) ToAliasResponseArrayOutputWithContext(ctx context.Context) AliasResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasResponseArrayOutput)
}

type AliasResponseOutput struct{ *pulumi.OutputState }

func (AliasResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasResponse)(nil)).Elem()
}

func (o AliasResponseOutput) ToAliasResponseOutput() AliasResponseOutput {
	return o
}

func (o AliasResponseOutput) ToAliasResponseOutputWithContext(ctx context.Context) AliasResponseOutput {
	return o
}

func (o AliasResponseOutput) DefaultMetadata() AliasPathMetadataResponseOutput {
	return o.ApplyT(func(v AliasResponse) AliasPathMetadataResponse { return v.DefaultMetadata }).(AliasPathMetadataResponseOutput)
}

func (o AliasResponseOutput) DefaultPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AliasResponse) *string { return v.DefaultPath }).(pulumi.StringPtrOutput)
}

func (o AliasResponseOutput) DefaultPattern() AliasPatternResponsePtrOutput {
	return o.ApplyT(func(v AliasResponse) *AliasPatternResponse { return v.DefaultPattern }).(AliasPatternResponsePtrOutput)
}

func (o AliasResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AliasResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AliasResponseOutput) Paths() AliasPathResponseArrayOutput {
	return o.ApplyT(func(v AliasResponse) []AliasPathResponse { return v.Paths }).(AliasPathResponseArrayOutput)
}

func (o AliasResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AliasResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AliasResponseArrayOutput struct{ *pulumi.OutputState }

func (AliasResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AliasResponse)(nil)).Elem()
}

func (o AliasResponseArrayOutput) ToAliasResponseArrayOutput() AliasResponseArrayOutput {
	return o
}

func (o AliasResponseArrayOutput) ToAliasResponseArrayOutputWithContext(ctx context.Context) AliasResponseArrayOutput {
	return o
}

func (o AliasResponseArrayOutput) Index(i pulumi.IntInput) AliasResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AliasResponse {
		return vs[0].([]AliasResponse)[vs[1].(int)]
	}).(AliasResponseOutput)
}

type ApiProfileResponse struct {
	ApiVersion     string `pulumi:"apiVersion"`
	ProfileVersion string `pulumi:"profileVersion"`
}





type ApiProfileResponseInput interface {
	pulumi.Input

	ToApiProfileResponseOutput() ApiProfileResponseOutput
	ToApiProfileResponseOutputWithContext(context.Context) ApiProfileResponseOutput
}

type ApiProfileResponseArgs struct {
	ApiVersion     pulumi.StringInput `pulumi:"apiVersion"`
	ProfileVersion pulumi.StringInput `pulumi:"profileVersion"`
}

func (ApiProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiProfileResponse)(nil)).Elem()
}

func (i ApiProfileResponseArgs) ToApiProfileResponseOutput() ApiProfileResponseOutput {
	return i.ToApiProfileResponseOutputWithContext(context.Background())
}

func (i ApiProfileResponseArgs) ToApiProfileResponseOutputWithContext(ctx context.Context) ApiProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiProfileResponseOutput)
}





type ApiProfileResponseArrayInput interface {
	pulumi.Input

	ToApiProfileResponseArrayOutput() ApiProfileResponseArrayOutput
	ToApiProfileResponseArrayOutputWithContext(context.Context) ApiProfileResponseArrayOutput
}

type ApiProfileResponseArray []ApiProfileResponseInput

func (ApiProfileResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiProfileResponse)(nil)).Elem()
}

func (i ApiProfileResponseArray) ToApiProfileResponseArrayOutput() ApiProfileResponseArrayOutput {
	return i.ToApiProfileResponseArrayOutputWithContext(context.Background())
}

func (i ApiProfileResponseArray) ToApiProfileResponseArrayOutputWithContext(ctx context.Context) ApiProfileResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiProfileResponseArrayOutput)
}

type ApiProfileResponseOutput struct{ *pulumi.OutputState }

func (ApiProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiProfileResponse)(nil)).Elem()
}

func (o ApiProfileResponseOutput) ToApiProfileResponseOutput() ApiProfileResponseOutput {
	return o
}

func (o ApiProfileResponseOutput) ToApiProfileResponseOutputWithContext(ctx context.Context) ApiProfileResponseOutput {
	return o
}

func (o ApiProfileResponseOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ApiProfileResponse) string { return v.ApiVersion }).(pulumi.StringOutput)
}

func (o ApiProfileResponseOutput) ProfileVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ApiProfileResponse) string { return v.ProfileVersion }).(pulumi.StringOutput)
}

type ApiProfileResponseArrayOutput struct{ *pulumi.OutputState }

func (ApiProfileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiProfileResponse)(nil)).Elem()
}

func (o ApiProfileResponseArrayOutput) ToApiProfileResponseArrayOutput() ApiProfileResponseArrayOutput {
	return o
}

func (o ApiProfileResponseArrayOutput) ToApiProfileResponseArrayOutputWithContext(ctx context.Context) ApiProfileResponseArrayOutput {
	return o
}

func (o ApiProfileResponseArrayOutput) Index(i pulumi.IntInput) ApiProfileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiProfileResponse {
		return vs[0].([]ApiProfileResponse)[vs[1].(int)]
	}).(ApiProfileResponseOutput)
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
	DebugSetting                *DebugSetting                `pulumi:"debugSetting"`
	ExpressionEvaluationOptions *ExpressionEvaluationOptions `pulumi:"expressionEvaluationOptions"`
	Mode                        DeploymentMode               `pulumi:"mode"`
	OnErrorDeployment           *OnErrorDeployment           `pulumi:"onErrorDeployment"`
	Parameters                  interface{}                  `pulumi:"parameters"`
	ParametersLink              *ParametersLink              `pulumi:"parametersLink"`
	Template                    interface{}                  `pulumi:"template"`
	TemplateLink                *TemplateLink                `pulumi:"templateLink"`
}





type DeploymentPropertiesInput interface {
	pulumi.Input

	ToDeploymentPropertiesOutput() DeploymentPropertiesOutput
	ToDeploymentPropertiesOutputWithContext(context.Context) DeploymentPropertiesOutput
}

type DeploymentPropertiesArgs struct {
	DebugSetting                DebugSettingPtrInput                `pulumi:"debugSetting"`
	ExpressionEvaluationOptions ExpressionEvaluationOptionsPtrInput `pulumi:"expressionEvaluationOptions"`
	Mode                        DeploymentModeInput                 `pulumi:"mode"`
	OnErrorDeployment           OnErrorDeploymentPtrInput           `pulumi:"onErrorDeployment"`
	Parameters                  pulumi.Input                        `pulumi:"parameters"`
	ParametersLink              ParametersLinkPtrInput              `pulumi:"parametersLink"`
	Template                    pulumi.Input                        `pulumi:"template"`
	TemplateLink                TemplateLinkPtrInput                `pulumi:"templateLink"`
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

func (o DeploymentPropertiesOutput) ExpressionEvaluationOptions() ExpressionEvaluationOptionsPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *ExpressionEvaluationOptions { return v.ExpressionEvaluationOptions }).(ExpressionEvaluationOptionsPtrOutput)
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

func (o DeploymentPropertiesPtrOutput) ExpressionEvaluationOptions() ExpressionEvaluationOptionsPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *ExpressionEvaluationOptions {
		if v == nil {
			return nil
		}
		return v.ExpressionEvaluationOptions
	}).(ExpressionEvaluationOptionsPtrOutput)
}

func (o DeploymentPropertiesPtrOutput) Mode() DeploymentModePtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *DeploymentMode {
		if v == nil {
			return nil
		}
		return &v.Mode
	}).(DeploymentModePtrOutput)
}

func (o DeploymentPropertiesPtrOutput) OnErrorDeployment() OnErrorDeploymentPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *OnErrorDeployment {
		if v == nil {
			return nil
		}
		return v.OnErrorDeployment
	}).(OnErrorDeploymentPtrOutput)
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
	CorrelationId      string                            `pulumi:"correlationId"`
	DebugSetting       DebugSettingResponse              `pulumi:"debugSetting"`
	Dependencies       []DependencyResponse              `pulumi:"dependencies"`
	Duration           string                            `pulumi:"duration"`
	Error              ErrorResponseResponse             `pulumi:"error"`
	Mode               string                            `pulumi:"mode"`
	OnErrorDeployment  OnErrorDeploymentExtendedResponse `pulumi:"onErrorDeployment"`
	OutputResources    []ResourceReferenceResponse       `pulumi:"outputResources"`
	Outputs            interface{}                       `pulumi:"outputs"`
	Parameters         interface{}                       `pulumi:"parameters"`
	ParametersLink     ParametersLinkResponse            `pulumi:"parametersLink"`
	Providers          []ProviderResponse                `pulumi:"providers"`
	ProvisioningState  string                            `pulumi:"provisioningState"`
	TemplateHash       string                            `pulumi:"templateHash"`
	TemplateLink       TemplateLinkResponse              `pulumi:"templateLink"`
	Timestamp          string                            `pulumi:"timestamp"`
	ValidatedResources []ResourceReferenceResponse       `pulumi:"validatedResources"`
}





type DeploymentPropertiesExtendedResponseInput interface {
	pulumi.Input

	ToDeploymentPropertiesExtendedResponseOutput() DeploymentPropertiesExtendedResponseOutput
	ToDeploymentPropertiesExtendedResponseOutputWithContext(context.Context) DeploymentPropertiesExtendedResponseOutput
}

type DeploymentPropertiesExtendedResponseArgs struct {
	CorrelationId      pulumi.StringInput                     `pulumi:"correlationId"`
	DebugSetting       DebugSettingResponseInput              `pulumi:"debugSetting"`
	Dependencies       DependencyResponseArrayInput           `pulumi:"dependencies"`
	Duration           pulumi.StringInput                     `pulumi:"duration"`
	Error              ErrorResponseResponseInput             `pulumi:"error"`
	Mode               pulumi.StringInput                     `pulumi:"mode"`
	OnErrorDeployment  OnErrorDeploymentExtendedResponseInput `pulumi:"onErrorDeployment"`
	OutputResources    ResourceReferenceResponseArrayInput    `pulumi:"outputResources"`
	Outputs            pulumi.Input                           `pulumi:"outputs"`
	Parameters         pulumi.Input                           `pulumi:"parameters"`
	ParametersLink     ParametersLinkResponseInput            `pulumi:"parametersLink"`
	Providers          ProviderResponseArrayInput             `pulumi:"providers"`
	ProvisioningState  pulumi.StringInput                     `pulumi:"provisioningState"`
	TemplateHash       pulumi.StringInput                     `pulumi:"templateHash"`
	TemplateLink       TemplateLinkResponseInput              `pulumi:"templateLink"`
	Timestamp          pulumi.StringInput                     `pulumi:"timestamp"`
	ValidatedResources ResourceReferenceResponseArrayInput    `pulumi:"validatedResources"`
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

func (o DeploymentPropertiesExtendedResponseOutput) DebugSetting() DebugSettingResponseOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) DebugSettingResponse { return v.DebugSetting }).(DebugSettingResponseOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Dependencies() DependencyResponseArrayOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) []DependencyResponse { return v.Dependencies }).(DependencyResponseArrayOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) string { return v.Duration }).(pulumi.StringOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Error() ErrorResponseResponseOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) ErrorResponseResponse { return v.Error }).(ErrorResponseResponseOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) string { return v.Mode }).(pulumi.StringOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) OnErrorDeployment() OnErrorDeploymentExtendedResponseOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) OnErrorDeploymentExtendedResponse {
		return v.OnErrorDeployment
	}).(OnErrorDeploymentExtendedResponseOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) OutputResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) []ResourceReferenceResponse { return v.OutputResources }).(ResourceReferenceResponseArrayOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) interface{} { return v.Outputs }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) ParametersLink() ParametersLinkResponseOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) ParametersLinkResponse { return v.ParametersLink }).(ParametersLinkResponseOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Providers() ProviderResponseArrayOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) []ProviderResponse { return v.Providers }).(ProviderResponseArrayOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) TemplateHash() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) string { return v.TemplateHash }).(pulumi.StringOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) TemplateLink() TemplateLinkResponseOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) TemplateLinkResponse { return v.TemplateLink }).(TemplateLinkResponseOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) ValidatedResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) []ResourceReferenceResponse { return v.ValidatedResources }).(ResourceReferenceResponseArrayOutput)
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
		return &v.DebugSetting
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

func (o DeploymentPropertiesExtendedResponsePtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Duration
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) Error() ErrorResponseResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *ErrorResponseResponse {
		if v == nil {
			return nil
		}
		return &v.Error
	}).(ErrorResponseResponsePtrOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) OnErrorDeployment() OnErrorDeploymentExtendedResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *OnErrorDeploymentExtendedResponse {
		if v == nil {
			return nil
		}
		return &v.OnErrorDeployment
	}).(OnErrorDeploymentExtendedResponsePtrOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) OutputResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) []ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.OutputResources
	}).(ResourceReferenceResponseArrayOutput)
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
		return &v.ParametersLink
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

func (o DeploymentPropertiesExtendedResponsePtrOutput) TemplateHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TemplateHash
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesExtendedResponsePtrOutput) TemplateLink() TemplateLinkResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) *TemplateLinkResponse {
		if v == nil {
			return nil
		}
		return &v.TemplateLink
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

func (o DeploymentPropertiesExtendedResponsePtrOutput) ValidatedResources() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentPropertiesExtendedResponse) []ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.ValidatedResources
	}).(ResourceReferenceResponseArrayOutput)
}

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}





type ErrorAdditionalInfoResponseInput interface {
	pulumi.Input

	ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput
	ToErrorAdditionalInfoResponseOutputWithContext(context.Context) ErrorAdditionalInfoResponseOutput
}

type ErrorAdditionalInfoResponseArgs struct {
	Info pulumi.Input       `pulumi:"info"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (ErrorAdditionalInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (i ErrorAdditionalInfoResponseArgs) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return i.ToErrorAdditionalInfoResponseOutputWithContext(context.Background())
}

func (i ErrorAdditionalInfoResponseArgs) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorAdditionalInfoResponseOutput)
}





type ErrorAdditionalInfoResponseArrayInput interface {
	pulumi.Input

	ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput
	ToErrorAdditionalInfoResponseArrayOutputWithContext(context.Context) ErrorAdditionalInfoResponseArrayOutput
}

type ErrorAdditionalInfoResponseArray []ErrorAdditionalInfoResponseInput

func (ErrorAdditionalInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (i ErrorAdditionalInfoResponseArray) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return i.ToErrorAdditionalInfoResponseArrayOutputWithContext(context.Background())
}

func (i ErrorAdditionalInfoResponseArray) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorAdditionalInfoResponseArrayOutput)
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





type ErrorResponseResponseInput interface {
	pulumi.Input

	ToErrorResponseResponseOutput() ErrorResponseResponseOutput
	ToErrorResponseResponseOutputWithContext(context.Context) ErrorResponseResponseOutput
}

type ErrorResponseResponseArgs struct {
	AdditionalInfo ErrorAdditionalInfoResponseArrayInput `pulumi:"additionalInfo"`
	Code           pulumi.StringInput                    `pulumi:"code"`
	Details        ErrorResponseResponseArrayInput       `pulumi:"details"`
	Message        pulumi.StringInput                    `pulumi:"message"`
	Target         pulumi.StringInput                    `pulumi:"target"`
}

func (ErrorResponseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponseResponse)(nil)).Elem()
}

func (i ErrorResponseResponseArgs) ToErrorResponseResponseOutput() ErrorResponseResponseOutput {
	return i.ToErrorResponseResponseOutputWithContext(context.Background())
}

func (i ErrorResponseResponseArgs) ToErrorResponseResponseOutputWithContext(ctx context.Context) ErrorResponseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseResponseOutput)
}

func (i ErrorResponseResponseArgs) ToErrorResponseResponsePtrOutput() ErrorResponseResponsePtrOutput {
	return i.ToErrorResponseResponsePtrOutputWithContext(context.Background())
}

func (i ErrorResponseResponseArgs) ToErrorResponseResponsePtrOutputWithContext(ctx context.Context) ErrorResponseResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseResponseOutput).ToErrorResponseResponsePtrOutputWithContext(ctx)
}









type ErrorResponseResponsePtrInput interface {
	pulumi.Input

	ToErrorResponseResponsePtrOutput() ErrorResponseResponsePtrOutput
	ToErrorResponseResponsePtrOutputWithContext(context.Context) ErrorResponseResponsePtrOutput
}

type errorResponseResponsePtrType ErrorResponseResponseArgs

func ErrorResponseResponsePtr(v *ErrorResponseResponseArgs) ErrorResponseResponsePtrInput {
	return (*errorResponseResponsePtrType)(v)
}

func (*errorResponseResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponseResponse)(nil)).Elem()
}

func (i *errorResponseResponsePtrType) ToErrorResponseResponsePtrOutput() ErrorResponseResponsePtrOutput {
	return i.ToErrorResponseResponsePtrOutputWithContext(context.Background())
}

func (i *errorResponseResponsePtrType) ToErrorResponseResponsePtrOutputWithContext(ctx context.Context) ErrorResponseResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseResponsePtrOutput)
}





type ErrorResponseResponseArrayInput interface {
	pulumi.Input

	ToErrorResponseResponseArrayOutput() ErrorResponseResponseArrayOutput
	ToErrorResponseResponseArrayOutputWithContext(context.Context) ErrorResponseResponseArrayOutput
}

type ErrorResponseResponseArray []ErrorResponseResponseInput

func (ErrorResponseResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorResponseResponse)(nil)).Elem()
}

func (i ErrorResponseResponseArray) ToErrorResponseResponseArrayOutput() ErrorResponseResponseArrayOutput {
	return i.ToErrorResponseResponseArrayOutputWithContext(context.Background())
}

func (i ErrorResponseResponseArray) ToErrorResponseResponseArrayOutputWithContext(ctx context.Context) ErrorResponseResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseResponseArrayOutput)
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

func (o ErrorResponseResponseOutput) ToErrorResponseResponsePtrOutput() ErrorResponseResponsePtrOutput {
	return o.ToErrorResponseResponsePtrOutputWithContext(context.Background())
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponsePtrOutputWithContext(ctx context.Context) ErrorResponseResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ErrorResponseResponse) *ErrorResponseResponse {
		return &v
	}).(ErrorResponseResponsePtrOutput)
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

type ErrorResponseResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponsePtrOutput) ToErrorResponseResponsePtrOutput() ErrorResponseResponsePtrOutput {
	return o
}

func (o ErrorResponseResponsePtrOutput) ToErrorResponseResponsePtrOutputWithContext(ctx context.Context) ErrorResponseResponsePtrOutput {
	return o
}

func (o ErrorResponseResponsePtrOutput) Elem() ErrorResponseResponseOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) ErrorResponseResponse {
		if v != nil {
			return *v
		}
		var ret ErrorResponseResponse
		return ret
	}).(ErrorResponseResponseOutput)
}

func (o ErrorResponseResponsePtrOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) []ErrorAdditionalInfoResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalInfo
	}).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorResponseResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorResponseResponsePtrOutput) Details() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) []ErrorResponseResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ErrorResponseResponseArrayOutput)
}

func (o ErrorResponseResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ErrorResponseResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Target
	}).(pulumi.StringPtrOutput)
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

type ExpressionEvaluationOptions struct {
	Scope *string `pulumi:"scope"`
}





type ExpressionEvaluationOptionsInput interface {
	pulumi.Input

	ToExpressionEvaluationOptionsOutput() ExpressionEvaluationOptionsOutput
	ToExpressionEvaluationOptionsOutputWithContext(context.Context) ExpressionEvaluationOptionsOutput
}

type ExpressionEvaluationOptionsArgs struct {
	Scope pulumi.StringPtrInput `pulumi:"scope"`
}

func (ExpressionEvaluationOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionEvaluationOptions)(nil)).Elem()
}

func (i ExpressionEvaluationOptionsArgs) ToExpressionEvaluationOptionsOutput() ExpressionEvaluationOptionsOutput {
	return i.ToExpressionEvaluationOptionsOutputWithContext(context.Background())
}

func (i ExpressionEvaluationOptionsArgs) ToExpressionEvaluationOptionsOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressionEvaluationOptionsOutput)
}

func (i ExpressionEvaluationOptionsArgs) ToExpressionEvaluationOptionsPtrOutput() ExpressionEvaluationOptionsPtrOutput {
	return i.ToExpressionEvaluationOptionsPtrOutputWithContext(context.Background())
}

func (i ExpressionEvaluationOptionsArgs) ToExpressionEvaluationOptionsPtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressionEvaluationOptionsOutput).ToExpressionEvaluationOptionsPtrOutputWithContext(ctx)
}









type ExpressionEvaluationOptionsPtrInput interface {
	pulumi.Input

	ToExpressionEvaluationOptionsPtrOutput() ExpressionEvaluationOptionsPtrOutput
	ToExpressionEvaluationOptionsPtrOutputWithContext(context.Context) ExpressionEvaluationOptionsPtrOutput
}

type expressionEvaluationOptionsPtrType ExpressionEvaluationOptionsArgs

func ExpressionEvaluationOptionsPtr(v *ExpressionEvaluationOptionsArgs) ExpressionEvaluationOptionsPtrInput {
	return (*expressionEvaluationOptionsPtrType)(v)
}

func (*expressionEvaluationOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressionEvaluationOptions)(nil)).Elem()
}

func (i *expressionEvaluationOptionsPtrType) ToExpressionEvaluationOptionsPtrOutput() ExpressionEvaluationOptionsPtrOutput {
	return i.ToExpressionEvaluationOptionsPtrOutputWithContext(context.Background())
}

func (i *expressionEvaluationOptionsPtrType) ToExpressionEvaluationOptionsPtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressionEvaluationOptionsPtrOutput)
}

type ExpressionEvaluationOptionsOutput struct{ *pulumi.OutputState }

func (ExpressionEvaluationOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionEvaluationOptions)(nil)).Elem()
}

func (o ExpressionEvaluationOptionsOutput) ToExpressionEvaluationOptionsOutput() ExpressionEvaluationOptionsOutput {
	return o
}

func (o ExpressionEvaluationOptionsOutput) ToExpressionEvaluationOptionsOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsOutput {
	return o
}

func (o ExpressionEvaluationOptionsOutput) ToExpressionEvaluationOptionsPtrOutput() ExpressionEvaluationOptionsPtrOutput {
	return o.ToExpressionEvaluationOptionsPtrOutputWithContext(context.Background())
}

func (o ExpressionEvaluationOptionsOutput) ToExpressionEvaluationOptionsPtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressionEvaluationOptions) *ExpressionEvaluationOptions {
		return &v
	}).(ExpressionEvaluationOptionsPtrOutput)
}

func (o ExpressionEvaluationOptionsOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressionEvaluationOptions) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

type ExpressionEvaluationOptionsPtrOutput struct{ *pulumi.OutputState }

func (ExpressionEvaluationOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressionEvaluationOptions)(nil)).Elem()
}

func (o ExpressionEvaluationOptionsPtrOutput) ToExpressionEvaluationOptionsPtrOutput() ExpressionEvaluationOptionsPtrOutput {
	return o
}

func (o ExpressionEvaluationOptionsPtrOutput) ToExpressionEvaluationOptionsPtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsPtrOutput {
	return o
}

func (o ExpressionEvaluationOptionsPtrOutput) Elem() ExpressionEvaluationOptionsOutput {
	return o.ApplyT(func(v *ExpressionEvaluationOptions) ExpressionEvaluationOptions {
		if v != nil {
			return *v
		}
		var ret ExpressionEvaluationOptions
		return ret
	}).(ExpressionEvaluationOptionsOutput)
}

func (o ExpressionEvaluationOptionsPtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressionEvaluationOptions) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (i ExtendedLocationArgs) ToExtendedLocationOutput() ExtendedLocationOutput {
	return i.ToExtendedLocationOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput)
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput).ToExtendedLocationPtrOutputWithContext(ctx)
}









type ExtendedLocationPtrInput interface {
	pulumi.Input

	ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput
	ToExtendedLocationPtrOutputWithContext(context.Context) ExtendedLocationPtrOutput
}

type extendedLocationPtrType ExtendedLocationArgs

func ExtendedLocationPtr(v *ExtendedLocationArgs) ExtendedLocationPtrInput {
	return (*extendedLocationPtrType)(v)
}

func (*extendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationPtrOutput)
}

type ExtendedLocationOutput struct{ *pulumi.OutputState }

func (ExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationOutput) ToExtendedLocationOutput() ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocation) *ExtendedLocation {
		return &v
	}).(ExtendedLocationPtrOutput)
}

func (o ExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) Elem() ExtendedLocationOutput {
	return o.ApplyT(func(v *ExtendedLocation) ExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ExtendedLocation
		return ret
	}).(ExtendedLocationOutput)
}

func (o ExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationResponseInput interface {
	pulumi.Input

	ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput
	ToExtendedLocationResponseOutputWithContext(context.Context) ExtendedLocationResponseOutput
}

type ExtendedLocationResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return i.ToExtendedLocationResponseOutputWithContext(context.Background())
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponseOutput)
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return i.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponseOutput).ToExtendedLocationResponsePtrOutputWithContext(ctx)
}









type ExtendedLocationResponsePtrInput interface {
	pulumi.Input

	ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput
	ToExtendedLocationResponsePtrOutputWithContext(context.Context) ExtendedLocationResponsePtrOutput
}

type extendedLocationResponsePtrType ExtendedLocationResponseArgs

func ExtendedLocationResponsePtr(v *ExtendedLocationResponseArgs) ExtendedLocationResponsePtrInput {
	return (*extendedLocationResponsePtrType)(v)
}

func (*extendedLocationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (i *extendedLocationResponsePtrType) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return i.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (i *extendedLocationResponsePtrType) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponsePtrOutput)
}

type ExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocationResponse) *ExtendedLocationResponse {
		return &v
	}).(ExtendedLocationResponsePtrOutput)
}

func (o ExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) Elem() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) ExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationResponse
		return ret
	}).(ExtendedLocationResponseOutput)
}

func (o ExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
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





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                             `pulumi:"principalId"`
	TenantId               pulumi.StringInput                             `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                          `pulumi:"type"`
	UserAssignedIdentities IdentityResponseUserAssignedIdentitiesMapInput `pulumi:"userAssignedIdentities"`
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





type IdentityResponseUserAssignedIdentitiesInput interface {
	pulumi.Input

	ToIdentityResponseUserAssignedIdentitiesOutput() IdentityResponseUserAssignedIdentitiesOutput
	ToIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Context) IdentityResponseUserAssignedIdentitiesOutput
}

type IdentityResponseUserAssignedIdentitiesArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (IdentityResponseUserAssignedIdentitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i IdentityResponseUserAssignedIdentitiesArgs) ToIdentityResponseUserAssignedIdentitiesOutput() IdentityResponseUserAssignedIdentitiesOutput {
	return i.ToIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Background())
}

func (i IdentityResponseUserAssignedIdentitiesArgs) ToIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseUserAssignedIdentitiesOutput)
}





type IdentityResponseUserAssignedIdentitiesMapInput interface {
	pulumi.Input

	ToIdentityResponseUserAssignedIdentitiesMapOutput() IdentityResponseUserAssignedIdentitiesMapOutput
	ToIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Context) IdentityResponseUserAssignedIdentitiesMapOutput
}

type IdentityResponseUserAssignedIdentitiesMap map[string]IdentityResponseUserAssignedIdentitiesInput

func (IdentityResponseUserAssignedIdentitiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i IdentityResponseUserAssignedIdentitiesMap) ToIdentityResponseUserAssignedIdentitiesMapOutput() IdentityResponseUserAssignedIdentitiesMapOutput {
	return i.ToIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Background())
}

func (i IdentityResponseUserAssignedIdentitiesMap) ToIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseUserAssignedIdentitiesMapOutput)
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





type OnErrorDeploymentExtendedResponseInput interface {
	pulumi.Input

	ToOnErrorDeploymentExtendedResponseOutput() OnErrorDeploymentExtendedResponseOutput
	ToOnErrorDeploymentExtendedResponseOutputWithContext(context.Context) OnErrorDeploymentExtendedResponseOutput
}

type OnErrorDeploymentExtendedResponseArgs struct {
	DeploymentName    pulumi.StringPtrInput `pulumi:"deploymentName"`
	ProvisioningState pulumi.StringInput    `pulumi:"provisioningState"`
	Type              pulumi.StringPtrInput `pulumi:"type"`
}

func (OnErrorDeploymentExtendedResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnErrorDeploymentExtendedResponse)(nil)).Elem()
}

func (i OnErrorDeploymentExtendedResponseArgs) ToOnErrorDeploymentExtendedResponseOutput() OnErrorDeploymentExtendedResponseOutput {
	return i.ToOnErrorDeploymentExtendedResponseOutputWithContext(context.Background())
}

func (i OnErrorDeploymentExtendedResponseArgs) ToOnErrorDeploymentExtendedResponseOutputWithContext(ctx context.Context) OnErrorDeploymentExtendedResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnErrorDeploymentExtendedResponseOutput)
}

func (i OnErrorDeploymentExtendedResponseArgs) ToOnErrorDeploymentExtendedResponsePtrOutput() OnErrorDeploymentExtendedResponsePtrOutput {
	return i.ToOnErrorDeploymentExtendedResponsePtrOutputWithContext(context.Background())
}

func (i OnErrorDeploymentExtendedResponseArgs) ToOnErrorDeploymentExtendedResponsePtrOutputWithContext(ctx context.Context) OnErrorDeploymentExtendedResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnErrorDeploymentExtendedResponseOutput).ToOnErrorDeploymentExtendedResponsePtrOutputWithContext(ctx)
}









type OnErrorDeploymentExtendedResponsePtrInput interface {
	pulumi.Input

	ToOnErrorDeploymentExtendedResponsePtrOutput() OnErrorDeploymentExtendedResponsePtrOutput
	ToOnErrorDeploymentExtendedResponsePtrOutputWithContext(context.Context) OnErrorDeploymentExtendedResponsePtrOutput
}

type onErrorDeploymentExtendedResponsePtrType OnErrorDeploymentExtendedResponseArgs

func OnErrorDeploymentExtendedResponsePtr(v *OnErrorDeploymentExtendedResponseArgs) OnErrorDeploymentExtendedResponsePtrInput {
	return (*onErrorDeploymentExtendedResponsePtrType)(v)
}

func (*onErrorDeploymentExtendedResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OnErrorDeploymentExtendedResponse)(nil)).Elem()
}

func (i *onErrorDeploymentExtendedResponsePtrType) ToOnErrorDeploymentExtendedResponsePtrOutput() OnErrorDeploymentExtendedResponsePtrOutput {
	return i.ToOnErrorDeploymentExtendedResponsePtrOutputWithContext(context.Background())
}

func (i *onErrorDeploymentExtendedResponsePtrType) ToOnErrorDeploymentExtendedResponsePtrOutputWithContext(ctx context.Context) OnErrorDeploymentExtendedResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnErrorDeploymentExtendedResponsePtrOutput)
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

func (o OnErrorDeploymentExtendedResponseOutput) ToOnErrorDeploymentExtendedResponsePtrOutput() OnErrorDeploymentExtendedResponsePtrOutput {
	return o.ToOnErrorDeploymentExtendedResponsePtrOutputWithContext(context.Background())
}

func (o OnErrorDeploymentExtendedResponseOutput) ToOnErrorDeploymentExtendedResponsePtrOutputWithContext(ctx context.Context) OnErrorDeploymentExtendedResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OnErrorDeploymentExtendedResponse) *OnErrorDeploymentExtendedResponse {
		return &v
	}).(OnErrorDeploymentExtendedResponsePtrOutput)
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
	Version       pulumi.StringPtrInput `pulumi:"version"`
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

type ProviderExtendedLocationResponse struct {
	ExtendedLocations []string `pulumi:"extendedLocations"`
	Location          *string  `pulumi:"location"`
	Type              *string  `pulumi:"type"`
}





type ProviderExtendedLocationResponseInput interface {
	pulumi.Input

	ToProviderExtendedLocationResponseOutput() ProviderExtendedLocationResponseOutput
	ToProviderExtendedLocationResponseOutputWithContext(context.Context) ProviderExtendedLocationResponseOutput
}

type ProviderExtendedLocationResponseArgs struct {
	ExtendedLocations pulumi.StringArrayInput `pulumi:"extendedLocations"`
	Location          pulumi.StringPtrInput   `pulumi:"location"`
	Type              pulumi.StringPtrInput   `pulumi:"type"`
}

func (ProviderExtendedLocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderExtendedLocationResponse)(nil)).Elem()
}

func (i ProviderExtendedLocationResponseArgs) ToProviderExtendedLocationResponseOutput() ProviderExtendedLocationResponseOutput {
	return i.ToProviderExtendedLocationResponseOutputWithContext(context.Background())
}

func (i ProviderExtendedLocationResponseArgs) ToProviderExtendedLocationResponseOutputWithContext(ctx context.Context) ProviderExtendedLocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderExtendedLocationResponseOutput)
}





type ProviderExtendedLocationResponseArrayInput interface {
	pulumi.Input

	ToProviderExtendedLocationResponseArrayOutput() ProviderExtendedLocationResponseArrayOutput
	ToProviderExtendedLocationResponseArrayOutputWithContext(context.Context) ProviderExtendedLocationResponseArrayOutput
}

type ProviderExtendedLocationResponseArray []ProviderExtendedLocationResponseInput

func (ProviderExtendedLocationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderExtendedLocationResponse)(nil)).Elem()
}

func (i ProviderExtendedLocationResponseArray) ToProviderExtendedLocationResponseArrayOutput() ProviderExtendedLocationResponseArrayOutput {
	return i.ToProviderExtendedLocationResponseArrayOutputWithContext(context.Background())
}

func (i ProviderExtendedLocationResponseArray) ToProviderExtendedLocationResponseArrayOutputWithContext(ctx context.Context) ProviderExtendedLocationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderExtendedLocationResponseArrayOutput)
}

type ProviderExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ProviderExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderExtendedLocationResponse)(nil)).Elem()
}

func (o ProviderExtendedLocationResponseOutput) ToProviderExtendedLocationResponseOutput() ProviderExtendedLocationResponseOutput {
	return o
}

func (o ProviderExtendedLocationResponseOutput) ToProviderExtendedLocationResponseOutputWithContext(ctx context.Context) ProviderExtendedLocationResponseOutput {
	return o
}

func (o ProviderExtendedLocationResponseOutput) ExtendedLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProviderExtendedLocationResponse) []string { return v.ExtendedLocations }).(pulumi.StringArrayOutput)
}

func (o ProviderExtendedLocationResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderExtendedLocationResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ProviderExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ProviderExtendedLocationResponseArrayOutput struct{ *pulumi.OutputState }

func (ProviderExtendedLocationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderExtendedLocationResponse)(nil)).Elem()
}

func (o ProviderExtendedLocationResponseArrayOutput) ToProviderExtendedLocationResponseArrayOutput() ProviderExtendedLocationResponseArrayOutput {
	return o
}

func (o ProviderExtendedLocationResponseArrayOutput) ToProviderExtendedLocationResponseArrayOutputWithContext(ctx context.Context) ProviderExtendedLocationResponseArrayOutput {
	return o
}

func (o ProviderExtendedLocationResponseArrayOutput) Index(i pulumi.IntInput) ProviderExtendedLocationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderExtendedLocationResponse {
		return vs[0].([]ProviderExtendedLocationResponse)[vs[1].(int)]
	}).(ProviderExtendedLocationResponseOutput)
}

type ProviderResourceTypeResponse struct {
	Aliases           []AliasResponse                    `pulumi:"aliases"`
	ApiProfiles       []ApiProfileResponse               `pulumi:"apiProfiles"`
	ApiVersions       []string                           `pulumi:"apiVersions"`
	Capabilities      *string                            `pulumi:"capabilities"`
	DefaultApiVersion string                             `pulumi:"defaultApiVersion"`
	LocationMappings  []ProviderExtendedLocationResponse `pulumi:"locationMappings"`
	Locations         []string                           `pulumi:"locations"`
	Properties        map[string]string                  `pulumi:"properties"`
	ResourceType      *string                            `pulumi:"resourceType"`
}





type ProviderResourceTypeResponseInput interface {
	pulumi.Input

	ToProviderResourceTypeResponseOutput() ProviderResourceTypeResponseOutput
	ToProviderResourceTypeResponseOutputWithContext(context.Context) ProviderResourceTypeResponseOutput
}

type ProviderResourceTypeResponseArgs struct {
	Aliases           AliasResponseArrayInput                    `pulumi:"aliases"`
	ApiProfiles       ApiProfileResponseArrayInput               `pulumi:"apiProfiles"`
	ApiVersions       pulumi.StringArrayInput                    `pulumi:"apiVersions"`
	Capabilities      pulumi.StringPtrInput                      `pulumi:"capabilities"`
	DefaultApiVersion pulumi.StringInput                         `pulumi:"defaultApiVersion"`
	LocationMappings  ProviderExtendedLocationResponseArrayInput `pulumi:"locationMappings"`
	Locations         pulumi.StringArrayInput                    `pulumi:"locations"`
	Properties        pulumi.StringMapInput                      `pulumi:"properties"`
	ResourceType      pulumi.StringPtrInput                      `pulumi:"resourceType"`
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

func (o ProviderResourceTypeResponseOutput) Aliases() AliasResponseArrayOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) []AliasResponse { return v.Aliases }).(AliasResponseArrayOutput)
}

func (o ProviderResourceTypeResponseOutput) ApiProfiles() ApiProfileResponseArrayOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) []ApiProfileResponse { return v.ApiProfiles }).(ApiProfileResponseArrayOutput)
}

func (o ProviderResourceTypeResponseOutput) ApiVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) []string { return v.ApiVersions }).(pulumi.StringArrayOutput)
}

func (o ProviderResourceTypeResponseOutput) Capabilities() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) *string { return v.Capabilities }).(pulumi.StringPtrOutput)
}

func (o ProviderResourceTypeResponseOutput) DefaultApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) string { return v.DefaultApiVersion }).(pulumi.StringOutput)
}

func (o ProviderResourceTypeResponseOutput) LocationMappings() ProviderExtendedLocationResponseArrayOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) []ProviderExtendedLocationResponse { return v.LocationMappings }).(ProviderExtendedLocationResponseArrayOutput)
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





type ProviderResponseInput interface {
	pulumi.Input

	ToProviderResponseOutput() ProviderResponseOutput
	ToProviderResponseOutputWithContext(context.Context) ProviderResponseOutput
}

type ProviderResponseArgs struct {
	Id                 pulumi.StringInput                     `pulumi:"id"`
	Namespace          pulumi.StringPtrInput                  `pulumi:"namespace"`
	RegistrationPolicy pulumi.StringInput                     `pulumi:"registrationPolicy"`
	RegistrationState  pulumi.StringInput                     `pulumi:"registrationState"`
	ResourceTypes      ProviderResourceTypeResponseArrayInput `pulumi:"resourceTypes"`
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

type ResourceReferenceResponse struct {
	Id string `pulumi:"id"`
}





type ResourceReferenceResponseInput interface {
	pulumi.Input

	ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput
	ToResourceReferenceResponseOutputWithContext(context.Context) ResourceReferenceResponseOutput
}

type ResourceReferenceResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (ResourceReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return i.ToResourceReferenceResponseOutputWithContext(context.Background())
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponseOutput)
}





type ResourceReferenceResponseArrayInput interface {
	pulumi.Input

	ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput
	ToResourceReferenceResponseArrayOutputWithContext(context.Context) ResourceReferenceResponseArrayOutput
}

type ResourceReferenceResponseArray []ResourceReferenceResponseInput

func (ResourceReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReferenceResponse)(nil)).Elem()
}

func (i ResourceReferenceResponseArray) ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput {
	return i.ToResourceReferenceResponseArrayOutputWithContext(context.Background())
}

func (i ResourceReferenceResponseArray) ToResourceReferenceResponseArrayOutputWithContext(ctx context.Context) ResourceReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponseArrayOutput)
}

type ResourceReferenceResponseOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ResourceReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutputWithContext(ctx context.Context) ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) Index(i pulumi.IntInput) ResourceReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceReferenceResponse {
		return vs[0].([]ResourceReferenceResponse)[vs[1].(int)]
	}).(ResourceReferenceResponseOutput)
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

type Tags struct {
	Tags map[string]string `pulumi:"tags"`
}





type TagsInput interface {
	pulumi.Input

	ToTagsOutput() TagsOutput
	ToTagsOutputWithContext(context.Context) TagsOutput
}

type TagsArgs struct {
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (TagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Tags)(nil)).Elem()
}

func (i TagsArgs) ToTagsOutput() TagsOutput {
	return i.ToTagsOutputWithContext(context.Background())
}

func (i TagsArgs) ToTagsOutputWithContext(ctx context.Context) TagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsOutput)
}

func (i TagsArgs) ToTagsPtrOutput() TagsPtrOutput {
	return i.ToTagsPtrOutputWithContext(context.Background())
}

func (i TagsArgs) ToTagsPtrOutputWithContext(ctx context.Context) TagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsOutput).ToTagsPtrOutputWithContext(ctx)
}









type TagsPtrInput interface {
	pulumi.Input

	ToTagsPtrOutput() TagsPtrOutput
	ToTagsPtrOutputWithContext(context.Context) TagsPtrOutput
}

type tagsPtrType TagsArgs

func TagsPtr(v *TagsArgs) TagsPtrInput {
	return (*tagsPtrType)(v)
}

func (*tagsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Tags)(nil)).Elem()
}

func (i *tagsPtrType) ToTagsPtrOutput() TagsPtrOutput {
	return i.ToTagsPtrOutputWithContext(context.Background())
}

func (i *tagsPtrType) ToTagsPtrOutputWithContext(ctx context.Context) TagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsPtrOutput)
}

type TagsOutput struct{ *pulumi.OutputState }

func (TagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Tags)(nil)).Elem()
}

func (o TagsOutput) ToTagsOutput() TagsOutput {
	return o
}

func (o TagsOutput) ToTagsOutputWithContext(ctx context.Context) TagsOutput {
	return o
}

func (o TagsOutput) ToTagsPtrOutput() TagsPtrOutput {
	return o.ToTagsPtrOutputWithContext(context.Background())
}

func (o TagsOutput) ToTagsPtrOutputWithContext(ctx context.Context) TagsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Tags) *Tags {
		return &v
	}).(TagsPtrOutput)
}

func (o TagsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v Tags) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type TagsPtrOutput struct{ *pulumi.OutputState }

func (TagsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tags)(nil)).Elem()
}

func (o TagsPtrOutput) ToTagsPtrOutput() TagsPtrOutput {
	return o
}

func (o TagsPtrOutput) ToTagsPtrOutputWithContext(ctx context.Context) TagsPtrOutput {
	return o
}

func (o TagsPtrOutput) Elem() TagsOutput {
	return o.ApplyT(func(v *Tags) Tags {
		if v != nil {
			return *v
		}
		var ret Tags
		return ret
	}).(TagsOutput)
}

func (o TagsPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Tags) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type TagsResponse struct {
	Tags map[string]string `pulumi:"tags"`
}





type TagsResponseInput interface {
	pulumi.Input

	ToTagsResponseOutput() TagsResponseOutput
	ToTagsResponseOutputWithContext(context.Context) TagsResponseOutput
}

type TagsResponseArgs struct {
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (TagsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagsResponse)(nil)).Elem()
}

func (i TagsResponseArgs) ToTagsResponseOutput() TagsResponseOutput {
	return i.ToTagsResponseOutputWithContext(context.Background())
}

func (i TagsResponseArgs) ToTagsResponseOutputWithContext(ctx context.Context) TagsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsResponseOutput)
}

func (i TagsResponseArgs) ToTagsResponsePtrOutput() TagsResponsePtrOutput {
	return i.ToTagsResponsePtrOutputWithContext(context.Background())
}

func (i TagsResponseArgs) ToTagsResponsePtrOutputWithContext(ctx context.Context) TagsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsResponseOutput).ToTagsResponsePtrOutputWithContext(ctx)
}









type TagsResponsePtrInput interface {
	pulumi.Input

	ToTagsResponsePtrOutput() TagsResponsePtrOutput
	ToTagsResponsePtrOutputWithContext(context.Context) TagsResponsePtrOutput
}

type tagsResponsePtrType TagsResponseArgs

func TagsResponsePtr(v *TagsResponseArgs) TagsResponsePtrInput {
	return (*tagsResponsePtrType)(v)
}

func (*tagsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagsResponse)(nil)).Elem()
}

func (i *tagsResponsePtrType) ToTagsResponsePtrOutput() TagsResponsePtrOutput {
	return i.ToTagsResponsePtrOutputWithContext(context.Background())
}

func (i *tagsResponsePtrType) ToTagsResponsePtrOutputWithContext(ctx context.Context) TagsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagsResponsePtrOutput)
}

type TagsResponseOutput struct{ *pulumi.OutputState }

func (TagsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagsResponse)(nil)).Elem()
}

func (o TagsResponseOutput) ToTagsResponseOutput() TagsResponseOutput {
	return o
}

func (o TagsResponseOutput) ToTagsResponseOutputWithContext(ctx context.Context) TagsResponseOutput {
	return o
}

func (o TagsResponseOutput) ToTagsResponsePtrOutput() TagsResponsePtrOutput {
	return o.ToTagsResponsePtrOutputWithContext(context.Background())
}

func (o TagsResponseOutput) ToTagsResponsePtrOutputWithContext(ctx context.Context) TagsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagsResponse) *TagsResponse {
		return &v
	}).(TagsResponsePtrOutput)
}

func (o TagsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v TagsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type TagsResponsePtrOutput struct{ *pulumi.OutputState }

func (TagsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagsResponse)(nil)).Elem()
}

func (o TagsResponsePtrOutput) ToTagsResponsePtrOutput() TagsResponsePtrOutput {
	return o
}

func (o TagsResponsePtrOutput) ToTagsResponsePtrOutputWithContext(ctx context.Context) TagsResponsePtrOutput {
	return o
}

func (o TagsResponsePtrOutput) Elem() TagsResponseOutput {
	return o.ApplyT(func(v *TagsResponse) TagsResponse {
		if v != nil {
			return *v
		}
		var ret TagsResponse
		return ret
	}).(TagsResponseOutput)
}

func (o TagsResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TagsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type TemplateLink struct {
	ContentVersion *string `pulumi:"contentVersion"`
	Id             *string `pulumi:"id"`
	QueryString    *string `pulumi:"queryString"`
	RelativePath   *string `pulumi:"relativePath"`
	Uri            *string `pulumi:"uri"`
}





type TemplateLinkInput interface {
	pulumi.Input

	ToTemplateLinkOutput() TemplateLinkOutput
	ToTemplateLinkOutputWithContext(context.Context) TemplateLinkOutput
}

type TemplateLinkArgs struct {
	ContentVersion pulumi.StringPtrInput `pulumi:"contentVersion"`
	Id             pulumi.StringPtrInput `pulumi:"id"`
	QueryString    pulumi.StringPtrInput `pulumi:"queryString"`
	RelativePath   pulumi.StringPtrInput `pulumi:"relativePath"`
	Uri            pulumi.StringPtrInput `pulumi:"uri"`
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

func (o TemplateLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o TemplateLinkOutput) QueryString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLink) *string { return v.QueryString }).(pulumi.StringPtrOutput)
}

func (o TemplateLinkOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLink) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func (o TemplateLinkOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLink) *string { return v.Uri }).(pulumi.StringPtrOutput)
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

func (o TemplateLinkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLink) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o TemplateLinkPtrOutput) QueryString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLink) *string {
		if v == nil {
			return nil
		}
		return v.QueryString
	}).(pulumi.StringPtrOutput)
}

func (o TemplateLinkPtrOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLink) *string {
		if v == nil {
			return nil
		}
		return v.RelativePath
	}).(pulumi.StringPtrOutput)
}

func (o TemplateLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLink) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type TemplateLinkResponse struct {
	ContentVersion *string `pulumi:"contentVersion"`
	Id             *string `pulumi:"id"`
	QueryString    *string `pulumi:"queryString"`
	RelativePath   *string `pulumi:"relativePath"`
	Uri            *string `pulumi:"uri"`
}





type TemplateLinkResponseInput interface {
	pulumi.Input

	ToTemplateLinkResponseOutput() TemplateLinkResponseOutput
	ToTemplateLinkResponseOutputWithContext(context.Context) TemplateLinkResponseOutput
}

type TemplateLinkResponseArgs struct {
	ContentVersion pulumi.StringPtrInput `pulumi:"contentVersion"`
	Id             pulumi.StringPtrInput `pulumi:"id"`
	QueryString    pulumi.StringPtrInput `pulumi:"queryString"`
	RelativePath   pulumi.StringPtrInput `pulumi:"relativePath"`
	Uri            pulumi.StringPtrInput `pulumi:"uri"`
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

func (o TemplateLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o TemplateLinkResponseOutput) QueryString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLinkResponse) *string { return v.QueryString }).(pulumi.StringPtrOutput)
}

func (o TemplateLinkResponseOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLinkResponse) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func (o TemplateLinkResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLinkResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
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

func (o TemplateLinkResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o TemplateLinkResponsePtrOutput) QueryString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.QueryString
	}).(pulumi.StringPtrOutput)
}

func (o TemplateLinkResponsePtrOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.RelativePath
	}).(pulumi.StringPtrOutput)
}

func (o TemplateLinkResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AliasPathMetadataResponseInput)(nil)).Elem(), AliasPathMetadataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasPathResponseInput)(nil)).Elem(), AliasPathResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasPathResponseArrayInput)(nil)).Elem(), AliasPathResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasPatternResponseInput)(nil)).Elem(), AliasPatternResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasPatternResponsePtrInput)(nil)).Elem(), AliasPatternResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasResponseInput)(nil)).Elem(), AliasResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasResponseArrayInput)(nil)).Elem(), AliasResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiProfileResponseInput)(nil)).Elem(), ApiProfileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiProfileResponseArrayInput)(nil)).Elem(), ApiProfileResponseArray{})
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
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorAdditionalInfoResponseInput)(nil)).Elem(), ErrorAdditionalInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorAdditionalInfoResponseArrayInput)(nil)).Elem(), ErrorAdditionalInfoResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorResponseResponseInput)(nil)).Elem(), ErrorResponseResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorResponseResponsePtrInput)(nil)).Elem(), ErrorResponseResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorResponseResponseArrayInput)(nil)).Elem(), ErrorResponseResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressionEvaluationOptionsInput)(nil)).Elem(), ExpressionEvaluationOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressionEvaluationOptionsPtrInput)(nil)).Elem(), ExpressionEvaluationOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationInput)(nil)).Elem(), ExtendedLocationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationPtrInput)(nil)).Elem(), ExtendedLocationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationResponseInput)(nil)).Elem(), ExtendedLocationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationResponsePtrInput)(nil)).Elem(), ExtendedLocationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPtrInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponseInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponsePtrInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponseUserAssignedIdentitiesInput)(nil)).Elem(), IdentityResponseUserAssignedIdentitiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponseUserAssignedIdentitiesMapInput)(nil)).Elem(), IdentityResponseUserAssignedIdentitiesMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnErrorDeploymentInput)(nil)).Elem(), OnErrorDeploymentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnErrorDeploymentPtrInput)(nil)).Elem(), OnErrorDeploymentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnErrorDeploymentExtendedResponseInput)(nil)).Elem(), OnErrorDeploymentExtendedResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OnErrorDeploymentExtendedResponsePtrInput)(nil)).Elem(), OnErrorDeploymentExtendedResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersLinkInput)(nil)).Elem(), ParametersLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersLinkPtrInput)(nil)).Elem(), ParametersLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersLinkResponseInput)(nil)).Elem(), ParametersLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersLinkResponsePtrInput)(nil)).Elem(), ParametersLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanInput)(nil)).Elem(), PlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanPtrInput)(nil)).Elem(), PlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanResponseInput)(nil)).Elem(), PlanResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanResponsePtrInput)(nil)).Elem(), PlanResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderExtendedLocationResponseInput)(nil)).Elem(), ProviderExtendedLocationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderExtendedLocationResponseArrayInput)(nil)).Elem(), ProviderExtendedLocationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderResourceTypeResponseInput)(nil)).Elem(), ProviderResourceTypeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderResourceTypeResponseArrayInput)(nil)).Elem(), ProviderResourceTypeResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderResponseInput)(nil)).Elem(), ProviderResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderResponseArrayInput)(nil)).Elem(), ProviderResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupPropertiesResponseInput)(nil)).Elem(), ResourceGroupPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupPropertiesResponsePtrInput)(nil)).Elem(), ResourceGroupPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceReferenceResponseInput)(nil)).Elem(), ResourceReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceReferenceResponseArrayInput)(nil)).Elem(), ResourceReferenceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagsInput)(nil)).Elem(), TagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagsPtrInput)(nil)).Elem(), TagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagsResponseInput)(nil)).Elem(), TagsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagsResponsePtrInput)(nil)).Elem(), TagsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateLinkInput)(nil)).Elem(), TemplateLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateLinkPtrInput)(nil)).Elem(), TemplateLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateLinkResponseInput)(nil)).Elem(), TemplateLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateLinkResponsePtrInput)(nil)).Elem(), TemplateLinkResponseArgs{})
	pulumi.RegisterOutputType(AliasPathMetadataResponseOutput{})
	pulumi.RegisterOutputType(AliasPathResponseOutput{})
	pulumi.RegisterOutputType(AliasPathResponseArrayOutput{})
	pulumi.RegisterOutputType(AliasPatternResponseOutput{})
	pulumi.RegisterOutputType(AliasPatternResponsePtrOutput{})
	pulumi.RegisterOutputType(AliasResponseOutput{})
	pulumi.RegisterOutputType(AliasResponseArrayOutput{})
	pulumi.RegisterOutputType(ApiProfileResponseOutput{})
	pulumi.RegisterOutputType(ApiProfileResponseArrayOutput{})
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
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseArrayOutput{})
	pulumi.RegisterOutputType(ExpressionEvaluationOptionsOutput{})
	pulumi.RegisterOutputType(ExpressionEvaluationOptionsPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ProviderExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ProviderExtendedLocationResponseArrayOutput{})
	pulumi.RegisterOutputType(ProviderResourceTypeResponseOutput{})
	pulumi.RegisterOutputType(ProviderResourceTypeResponseArrayOutput{})
	pulumi.RegisterOutputType(ProviderResponseOutput{})
	pulumi.RegisterOutputType(ProviderResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceGroupPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ResourceGroupPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(TagsOutput{})
	pulumi.RegisterOutputType(TagsPtrOutput{})
	pulumi.RegisterOutputType(TagsResponseOutput{})
	pulumi.RegisterOutputType(TagsResponsePtrOutput{})
	pulumi.RegisterOutputType(TemplateLinkOutput{})
	pulumi.RegisterOutputType(TemplateLinkPtrOutput{})
	pulumi.RegisterOutputType(TemplateLinkResponseOutput{})
	pulumi.RegisterOutputType(TemplateLinkResponsePtrOutput{})
}
