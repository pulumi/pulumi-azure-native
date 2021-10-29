


package v20190301

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

func (i ActionGroupsInformationArgs) ToActionGroupsInformationPtrOutput() ActionGroupsInformationPtrOutput {
	return i.ToActionGroupsInformationPtrOutputWithContext(context.Background())
}

func (i ActionGroupsInformationArgs) ToActionGroupsInformationPtrOutputWithContext(ctx context.Context) ActionGroupsInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationOutput).ToActionGroupsInformationPtrOutputWithContext(ctx)
}









type ActionGroupsInformationPtrInput interface {
	pulumi.Input

	ToActionGroupsInformationPtrOutput() ActionGroupsInformationPtrOutput
	ToActionGroupsInformationPtrOutputWithContext(context.Context) ActionGroupsInformationPtrOutput
}

type actionGroupsInformationPtrType ActionGroupsInformationArgs

func ActionGroupsInformationPtr(v *ActionGroupsInformationArgs) ActionGroupsInformationPtrInput {
	return (*actionGroupsInformationPtrType)(v)
}

func (*actionGroupsInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroupsInformation)(nil)).Elem()
}

func (i *actionGroupsInformationPtrType) ToActionGroupsInformationPtrOutput() ActionGroupsInformationPtrOutput {
	return i.ToActionGroupsInformationPtrOutputWithContext(context.Background())
}

func (i *actionGroupsInformationPtrType) ToActionGroupsInformationPtrOutputWithContext(ctx context.Context) ActionGroupsInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationPtrOutput)
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

func (o ActionGroupsInformationOutput) ToActionGroupsInformationPtrOutput() ActionGroupsInformationPtrOutput {
	return o.ToActionGroupsInformationPtrOutputWithContext(context.Background())
}

func (o ActionGroupsInformationOutput) ToActionGroupsInformationPtrOutputWithContext(ctx context.Context) ActionGroupsInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionGroupsInformation) *ActionGroupsInformation {
		return &v
	}).(ActionGroupsInformationPtrOutput)
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

type ActionGroupsInformationPtrOutput struct{ *pulumi.OutputState }

func (ActionGroupsInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroupsInformation)(nil)).Elem()
}

func (o ActionGroupsInformationPtrOutput) ToActionGroupsInformationPtrOutput() ActionGroupsInformationPtrOutput {
	return o
}

func (o ActionGroupsInformationPtrOutput) ToActionGroupsInformationPtrOutputWithContext(ctx context.Context) ActionGroupsInformationPtrOutput {
	return o
}

func (o ActionGroupsInformationPtrOutput) Elem() ActionGroupsInformationOutput {
	return o.ApplyT(func(v *ActionGroupsInformation) ActionGroupsInformation {
		if v != nil {
			return *v
		}
		var ret ActionGroupsInformation
		return ret
	}).(ActionGroupsInformationOutput)
}

func (o ActionGroupsInformationPtrOutput) CustomEmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionGroupsInformation) *string {
		if v == nil {
			return nil
		}
		return v.CustomEmailSubject
	}).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationPtrOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionGroupsInformation) *string {
		if v == nil {
			return nil
		}
		return v.CustomWebhookPayload
	}).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationPtrOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ActionGroupsInformation) []string {
		if v == nil {
			return nil
		}
		return v.GroupIds
	}).(pulumi.StringArrayOutput)
}

type ActionGroupsInformationResponse struct {
	CustomEmailSubject   *string  `pulumi:"customEmailSubject"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	GroupIds             []string `pulumi:"groupIds"`
}





type ActionGroupsInformationResponseInput interface {
	pulumi.Input

	ToActionGroupsInformationResponseOutput() ActionGroupsInformationResponseOutput
	ToActionGroupsInformationResponseOutputWithContext(context.Context) ActionGroupsInformationResponseOutput
}

type ActionGroupsInformationResponseArgs struct {
	CustomEmailSubject   pulumi.StringPtrInput   `pulumi:"customEmailSubject"`
	CustomWebhookPayload pulumi.StringPtrInput   `pulumi:"customWebhookPayload"`
	GroupIds             pulumi.StringArrayInput `pulumi:"groupIds"`
}

func (ActionGroupsInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupsInformationResponse)(nil)).Elem()
}

func (i ActionGroupsInformationResponseArgs) ToActionGroupsInformationResponseOutput() ActionGroupsInformationResponseOutput {
	return i.ToActionGroupsInformationResponseOutputWithContext(context.Background())
}

func (i ActionGroupsInformationResponseArgs) ToActionGroupsInformationResponseOutputWithContext(ctx context.Context) ActionGroupsInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationResponseOutput)
}

func (i ActionGroupsInformationResponseArgs) ToActionGroupsInformationResponsePtrOutput() ActionGroupsInformationResponsePtrOutput {
	return i.ToActionGroupsInformationResponsePtrOutputWithContext(context.Background())
}

func (i ActionGroupsInformationResponseArgs) ToActionGroupsInformationResponsePtrOutputWithContext(ctx context.Context) ActionGroupsInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationResponseOutput).ToActionGroupsInformationResponsePtrOutputWithContext(ctx)
}









type ActionGroupsInformationResponsePtrInput interface {
	pulumi.Input

	ToActionGroupsInformationResponsePtrOutput() ActionGroupsInformationResponsePtrOutput
	ToActionGroupsInformationResponsePtrOutputWithContext(context.Context) ActionGroupsInformationResponsePtrOutput
}

type actionGroupsInformationResponsePtrType ActionGroupsInformationResponseArgs

func ActionGroupsInformationResponsePtr(v *ActionGroupsInformationResponseArgs) ActionGroupsInformationResponsePtrInput {
	return (*actionGroupsInformationResponsePtrType)(v)
}

func (*actionGroupsInformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroupsInformationResponse)(nil)).Elem()
}

func (i *actionGroupsInformationResponsePtrType) ToActionGroupsInformationResponsePtrOutput() ActionGroupsInformationResponsePtrOutput {
	return i.ToActionGroupsInformationResponsePtrOutputWithContext(context.Background())
}

func (i *actionGroupsInformationResponsePtrType) ToActionGroupsInformationResponsePtrOutputWithContext(ctx context.Context) ActionGroupsInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationResponsePtrOutput)
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

func (o ActionGroupsInformationResponseOutput) ToActionGroupsInformationResponsePtrOutput() ActionGroupsInformationResponsePtrOutput {
	return o.ToActionGroupsInformationResponsePtrOutputWithContext(context.Background())
}

func (o ActionGroupsInformationResponseOutput) ToActionGroupsInformationResponsePtrOutputWithContext(ctx context.Context) ActionGroupsInformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionGroupsInformationResponse) *ActionGroupsInformationResponse {
		return &v
	}).(ActionGroupsInformationResponsePtrOutput)
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

type ActionGroupsInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (ActionGroupsInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroupsInformationResponse)(nil)).Elem()
}

func (o ActionGroupsInformationResponsePtrOutput) ToActionGroupsInformationResponsePtrOutput() ActionGroupsInformationResponsePtrOutput {
	return o
}

func (o ActionGroupsInformationResponsePtrOutput) ToActionGroupsInformationResponsePtrOutputWithContext(ctx context.Context) ActionGroupsInformationResponsePtrOutput {
	return o
}

func (o ActionGroupsInformationResponsePtrOutput) Elem() ActionGroupsInformationResponseOutput {
	return o.ApplyT(func(v *ActionGroupsInformationResponse) ActionGroupsInformationResponse {
		if v != nil {
			return *v
		}
		var ret ActionGroupsInformationResponse
		return ret
	}).(ActionGroupsInformationResponseOutput)
}

func (o ActionGroupsInformationResponsePtrOutput) CustomEmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionGroupsInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomEmailSubject
	}).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationResponsePtrOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionGroupsInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomWebhookPayload
	}).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationResponsePtrOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ActionGroupsInformationResponse) []string {
		if v == nil {
			return nil
		}
		return v.GroupIds
	}).(pulumi.StringArrayOutput)
}

type Detector struct {
	Description            *string                `pulumi:"description"`
	Id                     string                 `pulumi:"id"`
	ImagePaths             []string               `pulumi:"imagePaths"`
	Name                   *string                `pulumi:"name"`
	Parameters             map[string]interface{} `pulumi:"parameters"`
	SupportedResourceTypes []string               `pulumi:"supportedResourceTypes"`
}





type DetectorInput interface {
	pulumi.Input

	ToDetectorOutput() DetectorOutput
	ToDetectorOutputWithContext(context.Context) DetectorOutput
}

type DetectorArgs struct {
	Description            pulumi.StringPtrInput   `pulumi:"description"`
	Id                     pulumi.StringInput      `pulumi:"id"`
	ImagePaths             pulumi.StringArrayInput `pulumi:"imagePaths"`
	Name                   pulumi.StringPtrInput   `pulumi:"name"`
	Parameters             pulumi.MapInput         `pulumi:"parameters"`
	SupportedResourceTypes pulumi.StringArrayInput `pulumi:"supportedResourceTypes"`
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

func (i DetectorArgs) ToDetectorPtrOutput() DetectorPtrOutput {
	return i.ToDetectorPtrOutputWithContext(context.Background())
}

func (i DetectorArgs) ToDetectorPtrOutputWithContext(ctx context.Context) DetectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorOutput).ToDetectorPtrOutputWithContext(ctx)
}









type DetectorPtrInput interface {
	pulumi.Input

	ToDetectorPtrOutput() DetectorPtrOutput
	ToDetectorPtrOutputWithContext(context.Context) DetectorPtrOutput
}

type detectorPtrType DetectorArgs

func DetectorPtr(v *DetectorArgs) DetectorPtrInput {
	return (*detectorPtrType)(v)
}

func (*detectorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Detector)(nil)).Elem()
}

func (i *detectorPtrType) ToDetectorPtrOutput() DetectorPtrOutput {
	return i.ToDetectorPtrOutputWithContext(context.Background())
}

func (i *detectorPtrType) ToDetectorPtrOutputWithContext(ctx context.Context) DetectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorPtrOutput)
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

func (o DetectorOutput) ToDetectorPtrOutput() DetectorPtrOutput {
	return o.ToDetectorPtrOutputWithContext(context.Background())
}

func (o DetectorOutput) ToDetectorPtrOutputWithContext(ctx context.Context) DetectorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Detector) *Detector {
		return &v
	}).(DetectorPtrOutput)
}

func (o DetectorOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Detector) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DetectorOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Detector) string { return v.Id }).(pulumi.StringOutput)
}

func (o DetectorOutput) ImagePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Detector) []string { return v.ImagePaths }).(pulumi.StringArrayOutput)
}

func (o DetectorOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Detector) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DetectorOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v Detector) map[string]interface{} { return v.Parameters }).(pulumi.MapOutput)
}

func (o DetectorOutput) SupportedResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Detector) []string { return v.SupportedResourceTypes }).(pulumi.StringArrayOutput)
}

type DetectorPtrOutput struct{ *pulumi.OutputState }

func (DetectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Detector)(nil)).Elem()
}

func (o DetectorPtrOutput) ToDetectorPtrOutput() DetectorPtrOutput {
	return o
}

func (o DetectorPtrOutput) ToDetectorPtrOutputWithContext(ctx context.Context) DetectorPtrOutput {
	return o
}

func (o DetectorPtrOutput) Elem() DetectorOutput {
	return o.ApplyT(func(v *Detector) Detector {
		if v != nil {
			return *v
		}
		var ret Detector
		return ret
	}).(DetectorOutput)
}

func (o DetectorPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Detector) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o DetectorPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Detector) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o DetectorPtrOutput) ImagePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Detector) []string {
		if v == nil {
			return nil
		}
		return v.ImagePaths
	}).(pulumi.StringArrayOutput)
}

func (o DetectorPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Detector) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DetectorPtrOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v *Detector) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.MapOutput)
}

func (o DetectorPtrOutput) SupportedResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Detector) []string {
		if v == nil {
			return nil
		}
		return v.SupportedResourceTypes
	}).(pulumi.StringArrayOutput)
}

type DetectorResponse struct {
	Description            *string                `pulumi:"description"`
	Id                     string                 `pulumi:"id"`
	ImagePaths             []string               `pulumi:"imagePaths"`
	Name                   *string                `pulumi:"name"`
	Parameters             map[string]interface{} `pulumi:"parameters"`
	SupportedResourceTypes []string               `pulumi:"supportedResourceTypes"`
}





type DetectorResponseInput interface {
	pulumi.Input

	ToDetectorResponseOutput() DetectorResponseOutput
	ToDetectorResponseOutputWithContext(context.Context) DetectorResponseOutput
}

type DetectorResponseArgs struct {
	Description            pulumi.StringPtrInput   `pulumi:"description"`
	Id                     pulumi.StringInput      `pulumi:"id"`
	ImagePaths             pulumi.StringArrayInput `pulumi:"imagePaths"`
	Name                   pulumi.StringPtrInput   `pulumi:"name"`
	Parameters             pulumi.MapInput         `pulumi:"parameters"`
	SupportedResourceTypes pulumi.StringArrayInput `pulumi:"supportedResourceTypes"`
}

func (DetectorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorResponse)(nil)).Elem()
}

func (i DetectorResponseArgs) ToDetectorResponseOutput() DetectorResponseOutput {
	return i.ToDetectorResponseOutputWithContext(context.Background())
}

func (i DetectorResponseArgs) ToDetectorResponseOutputWithContext(ctx context.Context) DetectorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorResponseOutput)
}

func (i DetectorResponseArgs) ToDetectorResponsePtrOutput() DetectorResponsePtrOutput {
	return i.ToDetectorResponsePtrOutputWithContext(context.Background())
}

func (i DetectorResponseArgs) ToDetectorResponsePtrOutputWithContext(ctx context.Context) DetectorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorResponseOutput).ToDetectorResponsePtrOutputWithContext(ctx)
}









type DetectorResponsePtrInput interface {
	pulumi.Input

	ToDetectorResponsePtrOutput() DetectorResponsePtrOutput
	ToDetectorResponsePtrOutputWithContext(context.Context) DetectorResponsePtrOutput
}

type detectorResponsePtrType DetectorResponseArgs

func DetectorResponsePtr(v *DetectorResponseArgs) DetectorResponsePtrInput {
	return (*detectorResponsePtrType)(v)
}

func (*detectorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DetectorResponse)(nil)).Elem()
}

func (i *detectorResponsePtrType) ToDetectorResponsePtrOutput() DetectorResponsePtrOutput {
	return i.ToDetectorResponsePtrOutputWithContext(context.Background())
}

func (i *detectorResponsePtrType) ToDetectorResponsePtrOutputWithContext(ctx context.Context) DetectorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorResponsePtrOutput)
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

func (o DetectorResponseOutput) ToDetectorResponsePtrOutput() DetectorResponsePtrOutput {
	return o.ToDetectorResponsePtrOutputWithContext(context.Background())
}

func (o DetectorResponseOutput) ToDetectorResponsePtrOutputWithContext(ctx context.Context) DetectorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DetectorResponse) *DetectorResponse {
		return &v
	}).(DetectorResponsePtrOutput)
}

func (o DetectorResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DetectorResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DetectorResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DetectorResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o DetectorResponseOutput) ImagePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DetectorResponse) []string { return v.ImagePaths }).(pulumi.StringArrayOutput)
}

func (o DetectorResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DetectorResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DetectorResponseOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v DetectorResponse) map[string]interface{} { return v.Parameters }).(pulumi.MapOutput)
}

func (o DetectorResponseOutput) SupportedResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DetectorResponse) []string { return v.SupportedResourceTypes }).(pulumi.StringArrayOutput)
}

type DetectorResponsePtrOutput struct{ *pulumi.OutputState }

func (DetectorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DetectorResponse)(nil)).Elem()
}

func (o DetectorResponsePtrOutput) ToDetectorResponsePtrOutput() DetectorResponsePtrOutput {
	return o
}

func (o DetectorResponsePtrOutput) ToDetectorResponsePtrOutputWithContext(ctx context.Context) DetectorResponsePtrOutput {
	return o
}

func (o DetectorResponsePtrOutput) Elem() DetectorResponseOutput {
	return o.ApplyT(func(v *DetectorResponse) DetectorResponse {
		if v != nil {
			return *v
		}
		var ret DetectorResponse
		return ret
	}).(DetectorResponseOutput)
}

func (o DetectorResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DetectorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o DetectorResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DetectorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o DetectorResponsePtrOutput) ImagePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DetectorResponse) []string {
		if v == nil {
			return nil
		}
		return v.ImagePaths
	}).(pulumi.StringArrayOutput)
}

func (o DetectorResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DetectorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DetectorResponsePtrOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v *DetectorResponse) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.MapOutput)
}

func (o DetectorResponsePtrOutput) SupportedResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DetectorResponse) []string {
		if v == nil {
			return nil
		}
		return v.SupportedResourceTypes
	}).(pulumi.StringArrayOutput)
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





type ThrottlingInformationResponseInput interface {
	pulumi.Input

	ToThrottlingInformationResponseOutput() ThrottlingInformationResponseOutput
	ToThrottlingInformationResponseOutputWithContext(context.Context) ThrottlingInformationResponseOutput
}

type ThrottlingInformationResponseArgs struct {
	Duration pulumi.StringPtrInput `pulumi:"duration"`
}

func (ThrottlingInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingInformationResponse)(nil)).Elem()
}

func (i ThrottlingInformationResponseArgs) ToThrottlingInformationResponseOutput() ThrottlingInformationResponseOutput {
	return i.ToThrottlingInformationResponseOutputWithContext(context.Background())
}

func (i ThrottlingInformationResponseArgs) ToThrottlingInformationResponseOutputWithContext(ctx context.Context) ThrottlingInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationResponseOutput)
}

func (i ThrottlingInformationResponseArgs) ToThrottlingInformationResponsePtrOutput() ThrottlingInformationResponsePtrOutput {
	return i.ToThrottlingInformationResponsePtrOutputWithContext(context.Background())
}

func (i ThrottlingInformationResponseArgs) ToThrottlingInformationResponsePtrOutputWithContext(ctx context.Context) ThrottlingInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationResponseOutput).ToThrottlingInformationResponsePtrOutputWithContext(ctx)
}









type ThrottlingInformationResponsePtrInput interface {
	pulumi.Input

	ToThrottlingInformationResponsePtrOutput() ThrottlingInformationResponsePtrOutput
	ToThrottlingInformationResponsePtrOutputWithContext(context.Context) ThrottlingInformationResponsePtrOutput
}

type throttlingInformationResponsePtrType ThrottlingInformationResponseArgs

func ThrottlingInformationResponsePtr(v *ThrottlingInformationResponseArgs) ThrottlingInformationResponsePtrInput {
	return (*throttlingInformationResponsePtrType)(v)
}

func (*throttlingInformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ThrottlingInformationResponse)(nil)).Elem()
}

func (i *throttlingInformationResponsePtrType) ToThrottlingInformationResponsePtrOutput() ThrottlingInformationResponsePtrOutput {
	return i.ToThrottlingInformationResponsePtrOutputWithContext(context.Background())
}

func (i *throttlingInformationResponsePtrType) ToThrottlingInformationResponsePtrOutputWithContext(ctx context.Context) ThrottlingInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationResponsePtrOutput)
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

func (o ThrottlingInformationResponseOutput) ToThrottlingInformationResponsePtrOutput() ThrottlingInformationResponsePtrOutput {
	return o.ToThrottlingInformationResponsePtrOutputWithContext(context.Background())
}

func (o ThrottlingInformationResponseOutput) ToThrottlingInformationResponsePtrOutputWithContext(ctx context.Context) ThrottlingInformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ThrottlingInformationResponse) *ThrottlingInformationResponse {
		return &v
	}).(ThrottlingInformationResponsePtrOutput)
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
	pulumi.RegisterOutputType(ActionGroupsInformationPtrOutput{})
	pulumi.RegisterOutputType(ActionGroupsInformationResponseOutput{})
	pulumi.RegisterOutputType(ActionGroupsInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(DetectorOutput{})
	pulumi.RegisterOutputType(DetectorPtrOutput{})
	pulumi.RegisterOutputType(DetectorResponseOutput{})
	pulumi.RegisterOutputType(DetectorResponsePtrOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationPtrOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationResponseOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationResponsePtrOutput{})
}
