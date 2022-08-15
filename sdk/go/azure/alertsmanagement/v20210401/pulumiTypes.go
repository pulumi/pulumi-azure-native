


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionGroupsInformation struct {
	CustomEmailSubject   *string  `pulumi:"customEmailSubject"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	GroupIds             []string `pulumi:"groupIds"`
}





type ActionGroupsInformationInput interface {
	pulumi.Input

	ToActionGroupsInformationOutput() ActionGroupsInformationOutput
	ToActionGroupsInformationOutputWithContext(context.Context) ActionGroupsInformationOutput
}

type ActionGroupsInformationArgs struct {
	CustomEmailSubject   pulumi.StringPtrInput   `pulumi:"customEmailSubject"`
	CustomWebhookPayload pulumi.StringPtrInput   `pulumi:"customWebhookPayload"`
	GroupIds             pulumi.StringArrayInput `pulumi:"groupIds"`
}

func (ActionGroupsInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupsInformation)(nil)).Elem()
}

func (i ActionGroupsInformationArgs) ToActionGroupsInformationOutput() ActionGroupsInformationOutput {
	return i.ToActionGroupsInformationOutputWithContext(context.Background())
}

func (i ActionGroupsInformationArgs) ToActionGroupsInformationOutputWithContext(ctx context.Context) ActionGroupsInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationOutput)
}

type ActionGroupsInformationOutput struct{ *pulumi.OutputState }

func (ActionGroupsInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupsInformation)(nil)).Elem()
}

func (o ActionGroupsInformationOutput) ToActionGroupsInformationOutput() ActionGroupsInformationOutput {
	return o
}

func (o ActionGroupsInformationOutput) ToActionGroupsInformationOutputWithContext(ctx context.Context) ActionGroupsInformationOutput {
	return o
}

func (o ActionGroupsInformationOutput) CustomEmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformation) *string { return v.CustomEmailSubject }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformation) *string { return v.CustomWebhookPayload }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActionGroupsInformation) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

type ActionGroupsInformationResponse struct {
	CustomEmailSubject   *string  `pulumi:"customEmailSubject"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	GroupIds             []string `pulumi:"groupIds"`
}

type ActionGroupsInformationResponseOutput struct{ *pulumi.OutputState }

func (ActionGroupsInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupsInformationResponse)(nil)).Elem()
}

func (o ActionGroupsInformationResponseOutput) ToActionGroupsInformationResponseOutput() ActionGroupsInformationResponseOutput {
	return o
}

func (o ActionGroupsInformationResponseOutput) ToActionGroupsInformationResponseOutputWithContext(ctx context.Context) ActionGroupsInformationResponseOutput {
	return o
}

func (o ActionGroupsInformationResponseOutput) CustomEmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformationResponse) *string { return v.CustomEmailSubject }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationResponseOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformationResponse) *string { return v.CustomWebhookPayload }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActionGroupsInformationResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

type Detector struct {
	Id         string                 `pulumi:"id"`
	Parameters map[string]interface{} `pulumi:"parameters"`
}





type DetectorInput interface {
	pulumi.Input

	ToDetectorOutput() DetectorOutput
	ToDetectorOutputWithContext(context.Context) DetectorOutput
}

type DetectorArgs struct {
	Id         pulumi.StringInput `pulumi:"id"`
	Parameters pulumi.MapInput    `pulumi:"parameters"`
}

func (DetectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Detector)(nil)).Elem()
}

func (i DetectorArgs) ToDetectorOutput() DetectorOutput {
	return i.ToDetectorOutputWithContext(context.Background())
}

func (i DetectorArgs) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorOutput)
}

type DetectorOutput struct{ *pulumi.OutputState }

func (DetectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Detector)(nil)).Elem()
}

func (o DetectorOutput) ToDetectorOutput() DetectorOutput {
	return o
}

func (o DetectorOutput) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return o
}

func (o DetectorOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Detector) string { return v.Id }).(pulumi.StringOutput)
}

func (o DetectorOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v Detector) map[string]interface{} { return v.Parameters }).(pulumi.MapOutput)
}

type DetectorParameterDefinitionResponse struct {
	Description *string `pulumi:"description"`
	DisplayName *string `pulumi:"displayName"`
	IsMandatory *bool   `pulumi:"isMandatory"`
	Name        *string `pulumi:"name"`
	Type        *string `pulumi:"type"`
}

type DetectorParameterDefinitionResponseOutput struct{ *pulumi.OutputState }

func (DetectorParameterDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorParameterDefinitionResponse)(nil)).Elem()
}

func (o DetectorParameterDefinitionResponseOutput) ToDetectorParameterDefinitionResponseOutput() DetectorParameterDefinitionResponseOutput {
	return o
}

func (o DetectorParameterDefinitionResponseOutput) ToDetectorParameterDefinitionResponseOutputWithContext(ctx context.Context) DetectorParameterDefinitionResponseOutput {
	return o
}

func (o DetectorParameterDefinitionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DetectorParameterDefinitionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DetectorParameterDefinitionResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DetectorParameterDefinitionResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o DetectorParameterDefinitionResponseOutput) IsMandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DetectorParameterDefinitionResponse) *bool { return v.IsMandatory }).(pulumi.BoolPtrOutput)
}

func (o DetectorParameterDefinitionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DetectorParameterDefinitionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DetectorParameterDefinitionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DetectorParameterDefinitionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type DetectorParameterDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (DetectorParameterDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DetectorParameterDefinitionResponse)(nil)).Elem()
}

func (o DetectorParameterDefinitionResponseArrayOutput) ToDetectorParameterDefinitionResponseArrayOutput() DetectorParameterDefinitionResponseArrayOutput {
	return o
}

func (o DetectorParameterDefinitionResponseArrayOutput) ToDetectorParameterDefinitionResponseArrayOutputWithContext(ctx context.Context) DetectorParameterDefinitionResponseArrayOutput {
	return o
}

func (o DetectorParameterDefinitionResponseArrayOutput) Index(i pulumi.IntInput) DetectorParameterDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DetectorParameterDefinitionResponse {
		return vs[0].([]DetectorParameterDefinitionResponse)[vs[1].(int)]
	}).(DetectorParameterDefinitionResponseOutput)
}

type DetectorResponse struct {
	Description            string                                `pulumi:"description"`
	Id                     string                                `pulumi:"id"`
	ImagePaths             []string                              `pulumi:"imagePaths"`
	Name                   string                                `pulumi:"name"`
	ParameterDefinitions   []DetectorParameterDefinitionResponse `pulumi:"parameterDefinitions"`
	Parameters             map[string]interface{}                `pulumi:"parameters"`
	SupportedCadences      []int                                 `pulumi:"supportedCadences"`
	SupportedResourceTypes []string                              `pulumi:"supportedResourceTypes"`
}

type DetectorResponseOutput struct{ *pulumi.OutputState }

func (DetectorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorResponse)(nil)).Elem()
}

func (o DetectorResponseOutput) ToDetectorResponseOutput() DetectorResponseOutput {
	return o
}

func (o DetectorResponseOutput) ToDetectorResponseOutputWithContext(ctx context.Context) DetectorResponseOutput {
	return o
}

func (o DetectorResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v DetectorResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o DetectorResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DetectorResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o DetectorResponseOutput) ImagePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DetectorResponse) []string { return v.ImagePaths }).(pulumi.StringArrayOutput)
}

func (o DetectorResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DetectorResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DetectorResponseOutput) ParameterDefinitions() DetectorParameterDefinitionResponseArrayOutput {
	return o.ApplyT(func(v DetectorResponse) []DetectorParameterDefinitionResponse { return v.ParameterDefinitions }).(DetectorParameterDefinitionResponseArrayOutput)
}

func (o DetectorResponseOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v DetectorResponse) map[string]interface{} { return v.Parameters }).(pulumi.MapOutput)
}

func (o DetectorResponseOutput) SupportedCadences() pulumi.IntArrayOutput {
	return o.ApplyT(func(v DetectorResponse) []int { return v.SupportedCadences }).(pulumi.IntArrayOutput)
}

func (o DetectorResponseOutput) SupportedResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DetectorResponse) []string { return v.SupportedResourceTypes }).(pulumi.StringArrayOutput)
}

type ThrottlingInformation struct {
	Duration *string `pulumi:"duration"`
}





type ThrottlingInformationInput interface {
	pulumi.Input

	ToThrottlingInformationOutput() ThrottlingInformationOutput
	ToThrottlingInformationOutputWithContext(context.Context) ThrottlingInformationOutput
}

type ThrottlingInformationArgs struct {
	Duration pulumi.StringPtrInput `pulumi:"duration"`
}

func (ThrottlingInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingInformation)(nil)).Elem()
}

func (i ThrottlingInformationArgs) ToThrottlingInformationOutput() ThrottlingInformationOutput {
	return i.ToThrottlingInformationOutputWithContext(context.Background())
}

func (i ThrottlingInformationArgs) ToThrottlingInformationOutputWithContext(ctx context.Context) ThrottlingInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationOutput)
}

func (i ThrottlingInformationArgs) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return i.ToThrottlingInformationPtrOutputWithContext(context.Background())
}

func (i ThrottlingInformationArgs) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationOutput).ToThrottlingInformationPtrOutputWithContext(ctx)
}









type ThrottlingInformationPtrInput interface {
	pulumi.Input

	ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput
	ToThrottlingInformationPtrOutputWithContext(context.Context) ThrottlingInformationPtrOutput
}

type throttlingInformationPtrType ThrottlingInformationArgs

func ThrottlingInformationPtr(v *ThrottlingInformationArgs) ThrottlingInformationPtrInput {
	return (*throttlingInformationPtrType)(v)
}

func (*throttlingInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ThrottlingInformation)(nil)).Elem()
}

func (i *throttlingInformationPtrType) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return i.ToThrottlingInformationPtrOutputWithContext(context.Background())
}

func (i *throttlingInformationPtrType) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationPtrOutput)
}

type ThrottlingInformationOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingInformation)(nil)).Elem()
}

func (o ThrottlingInformationOutput) ToThrottlingInformationOutput() ThrottlingInformationOutput {
	return o
}

func (o ThrottlingInformationOutput) ToThrottlingInformationOutputWithContext(ctx context.Context) ThrottlingInformationOutput {
	return o
}

func (o ThrottlingInformationOutput) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return o.ToThrottlingInformationPtrOutputWithContext(context.Background())
}

func (o ThrottlingInformationOutput) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ThrottlingInformation) *ThrottlingInformation {
		return &v
	}).(ThrottlingInformationPtrOutput)
}

func (o ThrottlingInformationOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThrottlingInformation) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

type ThrottlingInformationPtrOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThrottlingInformation)(nil)).Elem()
}

func (o ThrottlingInformationPtrOutput) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return o
}

func (o ThrottlingInformationPtrOutput) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return o
}

func (o ThrottlingInformationPtrOutput) Elem() ThrottlingInformationOutput {
	return o.ApplyT(func(v *ThrottlingInformation) ThrottlingInformation {
		if v != nil {
			return *v
		}
		var ret ThrottlingInformation
		return ret
	}).(ThrottlingInformationOutput)
}

func (o ThrottlingInformationPtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThrottlingInformation) *string {
		if v == nil {
			return nil
		}
		return v.Duration
	}).(pulumi.StringPtrOutput)
}

type ThrottlingInformationResponse struct {
	Duration *string `pulumi:"duration"`
}

type ThrottlingInformationResponseOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingInformationResponse)(nil)).Elem()
}

func (o ThrottlingInformationResponseOutput) ToThrottlingInformationResponseOutput() ThrottlingInformationResponseOutput {
	return o
}

func (o ThrottlingInformationResponseOutput) ToThrottlingInformationResponseOutputWithContext(ctx context.Context) ThrottlingInformationResponseOutput {
	return o
}

func (o ThrottlingInformationResponseOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThrottlingInformationResponse) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

type ThrottlingInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThrottlingInformationResponse)(nil)).Elem()
}

func (o ThrottlingInformationResponsePtrOutput) ToThrottlingInformationResponsePtrOutput() ThrottlingInformationResponsePtrOutput {
	return o
}

func (o ThrottlingInformationResponsePtrOutput) ToThrottlingInformationResponsePtrOutputWithContext(ctx context.Context) ThrottlingInformationResponsePtrOutput {
	return o
}

func (o ThrottlingInformationResponsePtrOutput) Elem() ThrottlingInformationResponseOutput {
	return o.ApplyT(func(v *ThrottlingInformationResponse) ThrottlingInformationResponse {
		if v != nil {
			return *v
		}
		var ret ThrottlingInformationResponse
		return ret
	}).(ThrottlingInformationResponseOutput)
}

func (o ThrottlingInformationResponsePtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThrottlingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Duration
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionGroupsInformationOutput{})
	pulumi.RegisterOutputType(ActionGroupsInformationResponseOutput{})
	pulumi.RegisterOutputType(DetectorOutput{})
	pulumi.RegisterOutputType(DetectorParameterDefinitionResponseOutput{})
	pulumi.RegisterOutputType(DetectorParameterDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(DetectorResponseOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationPtrOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationResponseOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationResponsePtrOutput{})
}
