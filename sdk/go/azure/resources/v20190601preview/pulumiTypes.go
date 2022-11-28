


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

type TemplateSpecTemplateArtifact struct {
	Kind     string      `pulumi:"kind"`
	Path     string      `pulumi:"path"`
	Template interface{} `pulumi:"template"`
}





type TemplateSpecTemplateArtifactInput interface {
	pulumi.Input

	ToTemplateSpecTemplateArtifactOutput() TemplateSpecTemplateArtifactOutput
	ToTemplateSpecTemplateArtifactOutputWithContext(context.Context) TemplateSpecTemplateArtifactOutput
}

type TemplateSpecTemplateArtifactArgs struct {
	Kind     pulumi.StringInput `pulumi:"kind"`
	Path     pulumi.StringInput `pulumi:"path"`
	Template pulumi.Input       `pulumi:"template"`
}

func (TemplateSpecTemplateArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecTemplateArtifact)(nil)).Elem()
}

func (i TemplateSpecTemplateArtifactArgs) ToTemplateSpecTemplateArtifactOutput() TemplateSpecTemplateArtifactOutput {
	return i.ToTemplateSpecTemplateArtifactOutputWithContext(context.Background())
}

func (i TemplateSpecTemplateArtifactArgs) ToTemplateSpecTemplateArtifactOutputWithContext(ctx context.Context) TemplateSpecTemplateArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecTemplateArtifactOutput)
}





type TemplateSpecTemplateArtifactArrayInput interface {
	pulumi.Input

	ToTemplateSpecTemplateArtifactArrayOutput() TemplateSpecTemplateArtifactArrayOutput
	ToTemplateSpecTemplateArtifactArrayOutputWithContext(context.Context) TemplateSpecTemplateArtifactArrayOutput
}

type TemplateSpecTemplateArtifactArray []TemplateSpecTemplateArtifactInput

func (TemplateSpecTemplateArtifactArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TemplateSpecTemplateArtifact)(nil)).Elem()
}

func (i TemplateSpecTemplateArtifactArray) ToTemplateSpecTemplateArtifactArrayOutput() TemplateSpecTemplateArtifactArrayOutput {
	return i.ToTemplateSpecTemplateArtifactArrayOutputWithContext(context.Background())
}

func (i TemplateSpecTemplateArtifactArray) ToTemplateSpecTemplateArtifactArrayOutputWithContext(ctx context.Context) TemplateSpecTemplateArtifactArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecTemplateArtifactArrayOutput)
}

type TemplateSpecTemplateArtifactOutput struct{ *pulumi.OutputState }

func (TemplateSpecTemplateArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecTemplateArtifact)(nil)).Elem()
}

func (o TemplateSpecTemplateArtifactOutput) ToTemplateSpecTemplateArtifactOutput() TemplateSpecTemplateArtifactOutput {
	return o
}

func (o TemplateSpecTemplateArtifactOutput) ToTemplateSpecTemplateArtifactOutputWithContext(ctx context.Context) TemplateSpecTemplateArtifactOutput {
	return o
}

func (o TemplateSpecTemplateArtifactOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecTemplateArtifact) string { return v.Kind }).(pulumi.StringOutput)
}

func (o TemplateSpecTemplateArtifactOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecTemplateArtifact) string { return v.Path }).(pulumi.StringOutput)
}

func (o TemplateSpecTemplateArtifactOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v TemplateSpecTemplateArtifact) interface{} { return v.Template }).(pulumi.AnyOutput)
}

type TemplateSpecTemplateArtifactArrayOutput struct{ *pulumi.OutputState }

func (TemplateSpecTemplateArtifactArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TemplateSpecTemplateArtifact)(nil)).Elem()
}

func (o TemplateSpecTemplateArtifactArrayOutput) ToTemplateSpecTemplateArtifactArrayOutput() TemplateSpecTemplateArtifactArrayOutput {
	return o
}

func (o TemplateSpecTemplateArtifactArrayOutput) ToTemplateSpecTemplateArtifactArrayOutputWithContext(ctx context.Context) TemplateSpecTemplateArtifactArrayOutput {
	return o
}

func (o TemplateSpecTemplateArtifactArrayOutput) Index(i pulumi.IntInput) TemplateSpecTemplateArtifactOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TemplateSpecTemplateArtifact {
		return vs[0].([]TemplateSpecTemplateArtifact)[vs[1].(int)]
	}).(TemplateSpecTemplateArtifactOutput)
}

type TemplateSpecTemplateArtifactResponse struct {
	Kind     string      `pulumi:"kind"`
	Path     string      `pulumi:"path"`
	Template interface{} `pulumi:"template"`
}

type TemplateSpecTemplateArtifactResponseOutput struct{ *pulumi.OutputState }

func (TemplateSpecTemplateArtifactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecTemplateArtifactResponse)(nil)).Elem()
}

func (o TemplateSpecTemplateArtifactResponseOutput) ToTemplateSpecTemplateArtifactResponseOutput() TemplateSpecTemplateArtifactResponseOutput {
	return o
}

func (o TemplateSpecTemplateArtifactResponseOutput) ToTemplateSpecTemplateArtifactResponseOutputWithContext(ctx context.Context) TemplateSpecTemplateArtifactResponseOutput {
	return o
}

func (o TemplateSpecTemplateArtifactResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecTemplateArtifactResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o TemplateSpecTemplateArtifactResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecTemplateArtifactResponse) string { return v.Path }).(pulumi.StringOutput)
}

func (o TemplateSpecTemplateArtifactResponseOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v TemplateSpecTemplateArtifactResponse) interface{} { return v.Template }).(pulumi.AnyOutput)
}

type TemplateSpecTemplateArtifactResponseArrayOutput struct{ *pulumi.OutputState }

func (TemplateSpecTemplateArtifactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TemplateSpecTemplateArtifactResponse)(nil)).Elem()
}

func (o TemplateSpecTemplateArtifactResponseArrayOutput) ToTemplateSpecTemplateArtifactResponseArrayOutput() TemplateSpecTemplateArtifactResponseArrayOutput {
	return o
}

func (o TemplateSpecTemplateArtifactResponseArrayOutput) ToTemplateSpecTemplateArtifactResponseArrayOutputWithContext(ctx context.Context) TemplateSpecTemplateArtifactResponseArrayOutput {
	return o
}

func (o TemplateSpecTemplateArtifactResponseArrayOutput) Index(i pulumi.IntInput) TemplateSpecTemplateArtifactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TemplateSpecTemplateArtifactResponse {
		return vs[0].([]TemplateSpecTemplateArtifactResponse)[vs[1].(int)]
	}).(TemplateSpecTemplateArtifactResponseOutput)
}

type TemplateSpecVersionInfoResponse struct {
	Description  string `pulumi:"description"`
	TimeCreated  string `pulumi:"timeCreated"`
	TimeModified string `pulumi:"timeModified"`
}

type TemplateSpecVersionInfoResponseOutput struct{ *pulumi.OutputState }

func (TemplateSpecVersionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecVersionInfoResponse)(nil)).Elem()
}

func (o TemplateSpecVersionInfoResponseOutput) ToTemplateSpecVersionInfoResponseOutput() TemplateSpecVersionInfoResponseOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseOutput) ToTemplateSpecVersionInfoResponseOutputWithContext(ctx context.Context) TemplateSpecVersionInfoResponseOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o TemplateSpecVersionInfoResponseOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o TemplateSpecVersionInfoResponseOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.TimeModified }).(pulumi.StringOutput)
}

type TemplateSpecVersionInfoResponseMapOutput struct{ *pulumi.OutputState }

func (TemplateSpecVersionInfoResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TemplateSpecVersionInfoResponse)(nil)).Elem()
}

func (o TemplateSpecVersionInfoResponseMapOutput) ToTemplateSpecVersionInfoResponseMapOutput() TemplateSpecVersionInfoResponseMapOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseMapOutput) ToTemplateSpecVersionInfoResponseMapOutputWithContext(ctx context.Context) TemplateSpecVersionInfoResponseMapOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseMapOutput) MapIndex(k pulumi.StringInput) TemplateSpecVersionInfoResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TemplateSpecVersionInfoResponse {
		return vs[0].(map[string]TemplateSpecVersionInfoResponse)[vs[1].(string)]
	}).(TemplateSpecVersionInfoResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TemplateSpecTemplateArtifactOutput{})
	pulumi.RegisterOutputType(TemplateSpecTemplateArtifactArrayOutput{})
	pulumi.RegisterOutputType(TemplateSpecTemplateArtifactResponseOutput{})
	pulumi.RegisterOutputType(TemplateSpecTemplateArtifactResponseArrayOutput{})
	pulumi.RegisterOutputType(TemplateSpecVersionInfoResponseOutput{})
	pulumi.RegisterOutputType(TemplateSpecVersionInfoResponseMapOutput{})
}
