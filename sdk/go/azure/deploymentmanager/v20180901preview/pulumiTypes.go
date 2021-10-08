


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudErrorBodyResponse struct {
	Code    string                   `pulumi:"code"`
	Details []CloudErrorBodyResponse `pulumi:"details"`
	Message string                   `pulumi:"message"`
	Target  *string                  `pulumi:"target"`
}





type CloudErrorBodyResponseInput interface {
	pulumi.Input

	ToCloudErrorBodyResponseOutput() CloudErrorBodyResponseOutput
	ToCloudErrorBodyResponseOutputWithContext(context.Context) CloudErrorBodyResponseOutput
}

type CloudErrorBodyResponseArgs struct {
	Code    pulumi.StringInput               `pulumi:"code"`
	Details CloudErrorBodyResponseArrayInput `pulumi:"details"`
	Message pulumi.StringInput               `pulumi:"message"`
	Target  pulumi.StringPtrInput            `pulumi:"target"`
}

func (CloudErrorBodyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudErrorBodyResponse)(nil)).Elem()
}

func (i CloudErrorBodyResponseArgs) ToCloudErrorBodyResponseOutput() CloudErrorBodyResponseOutput {
	return i.ToCloudErrorBodyResponseOutputWithContext(context.Background())
}

func (i CloudErrorBodyResponseArgs) ToCloudErrorBodyResponseOutputWithContext(ctx context.Context) CloudErrorBodyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorBodyResponseOutput)
}

func (i CloudErrorBodyResponseArgs) ToCloudErrorBodyResponsePtrOutput() CloudErrorBodyResponsePtrOutput {
	return i.ToCloudErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (i CloudErrorBodyResponseArgs) ToCloudErrorBodyResponsePtrOutputWithContext(ctx context.Context) CloudErrorBodyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorBodyResponseOutput).ToCloudErrorBodyResponsePtrOutputWithContext(ctx)
}









type CloudErrorBodyResponsePtrInput interface {
	pulumi.Input

	ToCloudErrorBodyResponsePtrOutput() CloudErrorBodyResponsePtrOutput
	ToCloudErrorBodyResponsePtrOutputWithContext(context.Context) CloudErrorBodyResponsePtrOutput
}

type cloudErrorBodyResponsePtrType CloudErrorBodyResponseArgs

func CloudErrorBodyResponsePtr(v *CloudErrorBodyResponseArgs) CloudErrorBodyResponsePtrInput {
	return (*cloudErrorBodyResponsePtrType)(v)
}

func (*cloudErrorBodyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudErrorBodyResponse)(nil)).Elem()
}

func (i *cloudErrorBodyResponsePtrType) ToCloudErrorBodyResponsePtrOutput() CloudErrorBodyResponsePtrOutput {
	return i.ToCloudErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (i *cloudErrorBodyResponsePtrType) ToCloudErrorBodyResponsePtrOutputWithContext(ctx context.Context) CloudErrorBodyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorBodyResponsePtrOutput)
}





type CloudErrorBodyResponseArrayInput interface {
	pulumi.Input

	ToCloudErrorBodyResponseArrayOutput() CloudErrorBodyResponseArrayOutput
	ToCloudErrorBodyResponseArrayOutputWithContext(context.Context) CloudErrorBodyResponseArrayOutput
}

type CloudErrorBodyResponseArray []CloudErrorBodyResponseInput

func (CloudErrorBodyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudErrorBodyResponse)(nil)).Elem()
}

func (i CloudErrorBodyResponseArray) ToCloudErrorBodyResponseArrayOutput() CloudErrorBodyResponseArrayOutput {
	return i.ToCloudErrorBodyResponseArrayOutputWithContext(context.Background())
}

func (i CloudErrorBodyResponseArray) ToCloudErrorBodyResponseArrayOutputWithContext(ctx context.Context) CloudErrorBodyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorBodyResponseArrayOutput)
}

type CloudErrorBodyResponseOutput struct{ *pulumi.OutputState }

func (CloudErrorBodyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudErrorBodyResponse)(nil)).Elem()
}

func (o CloudErrorBodyResponseOutput) ToCloudErrorBodyResponseOutput() CloudErrorBodyResponseOutput {
	return o
}

func (o CloudErrorBodyResponseOutput) ToCloudErrorBodyResponseOutputWithContext(ctx context.Context) CloudErrorBodyResponseOutput {
	return o
}

func (o CloudErrorBodyResponseOutput) ToCloudErrorBodyResponsePtrOutput() CloudErrorBodyResponsePtrOutput {
	return o.ToCloudErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (o CloudErrorBodyResponseOutput) ToCloudErrorBodyResponsePtrOutputWithContext(ctx context.Context) CloudErrorBodyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudErrorBodyResponse) *CloudErrorBodyResponse {
		return &v
	}).(CloudErrorBodyResponsePtrOutput)
}

func (o CloudErrorBodyResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o CloudErrorBodyResponseOutput) Details() CloudErrorBodyResponseArrayOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) []CloudErrorBodyResponse { return v.Details }).(CloudErrorBodyResponseArrayOutput)
}

func (o CloudErrorBodyResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o CloudErrorBodyResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudErrorBodyResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type CloudErrorBodyResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudErrorBodyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudErrorBodyResponse)(nil)).Elem()
}

func (o CloudErrorBodyResponsePtrOutput) ToCloudErrorBodyResponsePtrOutput() CloudErrorBodyResponsePtrOutput {
	return o
}

func (o CloudErrorBodyResponsePtrOutput) ToCloudErrorBodyResponsePtrOutputWithContext(ctx context.Context) CloudErrorBodyResponsePtrOutput {
	return o
}

func (o CloudErrorBodyResponsePtrOutput) Elem() CloudErrorBodyResponseOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) CloudErrorBodyResponse {
		if v != nil {
			return *v
		}
		var ret CloudErrorBodyResponse
		return ret
	}).(CloudErrorBodyResponseOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Details() CloudErrorBodyResponseArrayOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) []CloudErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(CloudErrorBodyResponseArrayOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o CloudErrorBodyResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type CloudErrorBodyResponseArrayOutput struct{ *pulumi.OutputState }

func (CloudErrorBodyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudErrorBodyResponse)(nil)).Elem()
}

func (o CloudErrorBodyResponseArrayOutput) ToCloudErrorBodyResponseArrayOutput() CloudErrorBodyResponseArrayOutput {
	return o
}

func (o CloudErrorBodyResponseArrayOutput) ToCloudErrorBodyResponseArrayOutputWithContext(ctx context.Context) CloudErrorBodyResponseArrayOutput {
	return o
}

func (o CloudErrorBodyResponseArrayOutput) Index(i pulumi.IntInput) CloudErrorBodyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudErrorBodyResponse {
		return vs[0].([]CloudErrorBodyResponse)[vs[1].(int)]
	}).(CloudErrorBodyResponseOutput)
}

type Identity struct {
	IdentityIds []string `pulumi:"identityIds"`
	Type        string   `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	IdentityIds pulumi.StringArrayInput `pulumi:"identityIds"`
	Type        pulumi.StringInput      `pulumi:"type"`
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

func (o IdentityOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Identity) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

func (o IdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v Identity) string { return v.Type }).(pulumi.StringOutput)
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

func (o IdentityPtrOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Identity) []string {
		if v == nil {
			return nil
		}
		return v.IdentityIds
	}).(pulumi.StringArrayOutput)
}

func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityResponse struct {
	IdentityIds []string `pulumi:"identityIds"`
	Type        string   `pulumi:"type"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	IdentityIds pulumi.StringArrayInput `pulumi:"identityIds"`
	Type        pulumi.StringInput      `pulumi:"type"`
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

func (o IdentityResponseOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IdentityResponse) []string { return v.IdentityIds }).(pulumi.StringArrayOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
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

func (o IdentityResponsePtrOutput) IdentityIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdentityResponse) []string {
		if v == nil {
			return nil
		}
		return v.IdentityIds
	}).(pulumi.StringArrayOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type MessageResponse struct {
	Message   string `pulumi:"message"`
	TimeStamp string `pulumi:"timeStamp"`
}





type MessageResponseInput interface {
	pulumi.Input

	ToMessageResponseOutput() MessageResponseOutput
	ToMessageResponseOutputWithContext(context.Context) MessageResponseOutput
}

type MessageResponseArgs struct {
	Message   pulumi.StringInput `pulumi:"message"`
	TimeStamp pulumi.StringInput `pulumi:"timeStamp"`
}

func (MessageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageResponse)(nil)).Elem()
}

func (i MessageResponseArgs) ToMessageResponseOutput() MessageResponseOutput {
	return i.ToMessageResponseOutputWithContext(context.Background())
}

func (i MessageResponseArgs) ToMessageResponseOutputWithContext(ctx context.Context) MessageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessageResponseOutput)
}





type MessageResponseArrayInput interface {
	pulumi.Input

	ToMessageResponseArrayOutput() MessageResponseArrayOutput
	ToMessageResponseArrayOutputWithContext(context.Context) MessageResponseArrayOutput
}

type MessageResponseArray []MessageResponseInput

func (MessageResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MessageResponse)(nil)).Elem()
}

func (i MessageResponseArray) ToMessageResponseArrayOutput() MessageResponseArrayOutput {
	return i.ToMessageResponseArrayOutputWithContext(context.Background())
}

func (i MessageResponseArray) ToMessageResponseArrayOutputWithContext(ctx context.Context) MessageResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessageResponseArrayOutput)
}

type MessageResponseOutput struct{ *pulumi.OutputState }

func (MessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageResponse)(nil)).Elem()
}

func (o MessageResponseOutput) ToMessageResponseOutput() MessageResponseOutput {
	return o
}

func (o MessageResponseOutput) ToMessageResponseOutputWithContext(ctx context.Context) MessageResponseOutput {
	return o
}

func (o MessageResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v MessageResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o MessageResponseOutput) TimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v MessageResponse) string { return v.TimeStamp }).(pulumi.StringOutput)
}

type MessageResponseArrayOutput struct{ *pulumi.OutputState }

func (MessageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MessageResponse)(nil)).Elem()
}

func (o MessageResponseArrayOutput) ToMessageResponseArrayOutput() MessageResponseArrayOutput {
	return o
}

func (o MessageResponseArrayOutput) ToMessageResponseArrayOutputWithContext(ctx context.Context) MessageResponseArrayOutput {
	return o
}

func (o MessageResponseArrayOutput) Index(i pulumi.IntInput) MessageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MessageResponse {
		return vs[0].([]MessageResponse)[vs[1].(int)]
	}).(MessageResponseOutput)
}

type PrePostStep struct {
	StepId string `pulumi:"stepId"`
}





type PrePostStepInput interface {
	pulumi.Input

	ToPrePostStepOutput() PrePostStepOutput
	ToPrePostStepOutputWithContext(context.Context) PrePostStepOutput
}

type PrePostStepArgs struct {
	StepId pulumi.StringInput `pulumi:"stepId"`
}

func (PrePostStepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrePostStep)(nil)).Elem()
}

func (i PrePostStepArgs) ToPrePostStepOutput() PrePostStepOutput {
	return i.ToPrePostStepOutputWithContext(context.Background())
}

func (i PrePostStepArgs) ToPrePostStepOutputWithContext(ctx context.Context) PrePostStepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrePostStepOutput)
}





type PrePostStepArrayInput interface {
	pulumi.Input

	ToPrePostStepArrayOutput() PrePostStepArrayOutput
	ToPrePostStepArrayOutputWithContext(context.Context) PrePostStepArrayOutput
}

type PrePostStepArray []PrePostStepInput

func (PrePostStepArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrePostStep)(nil)).Elem()
}

func (i PrePostStepArray) ToPrePostStepArrayOutput() PrePostStepArrayOutput {
	return i.ToPrePostStepArrayOutputWithContext(context.Background())
}

func (i PrePostStepArray) ToPrePostStepArrayOutputWithContext(ctx context.Context) PrePostStepArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrePostStepArrayOutput)
}

type PrePostStepOutput struct{ *pulumi.OutputState }

func (PrePostStepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrePostStep)(nil)).Elem()
}

func (o PrePostStepOutput) ToPrePostStepOutput() PrePostStepOutput {
	return o
}

func (o PrePostStepOutput) ToPrePostStepOutputWithContext(ctx context.Context) PrePostStepOutput {
	return o
}

func (o PrePostStepOutput) StepId() pulumi.StringOutput {
	return o.ApplyT(func(v PrePostStep) string { return v.StepId }).(pulumi.StringOutput)
}

type PrePostStepArrayOutput struct{ *pulumi.OutputState }

func (PrePostStepArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrePostStep)(nil)).Elem()
}

func (o PrePostStepArrayOutput) ToPrePostStepArrayOutput() PrePostStepArrayOutput {
	return o
}

func (o PrePostStepArrayOutput) ToPrePostStepArrayOutputWithContext(ctx context.Context) PrePostStepArrayOutput {
	return o
}

func (o PrePostStepArrayOutput) Index(i pulumi.IntInput) PrePostStepOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrePostStep {
		return vs[0].([]PrePostStep)[vs[1].(int)]
	}).(PrePostStepOutput)
}

type PrePostStepResponse struct {
	StepId string `pulumi:"stepId"`
}





type PrePostStepResponseInput interface {
	pulumi.Input

	ToPrePostStepResponseOutput() PrePostStepResponseOutput
	ToPrePostStepResponseOutputWithContext(context.Context) PrePostStepResponseOutput
}

type PrePostStepResponseArgs struct {
	StepId pulumi.StringInput `pulumi:"stepId"`
}

func (PrePostStepResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrePostStepResponse)(nil)).Elem()
}

func (i PrePostStepResponseArgs) ToPrePostStepResponseOutput() PrePostStepResponseOutput {
	return i.ToPrePostStepResponseOutputWithContext(context.Background())
}

func (i PrePostStepResponseArgs) ToPrePostStepResponseOutputWithContext(ctx context.Context) PrePostStepResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrePostStepResponseOutput)
}





type PrePostStepResponseArrayInput interface {
	pulumi.Input

	ToPrePostStepResponseArrayOutput() PrePostStepResponseArrayOutput
	ToPrePostStepResponseArrayOutputWithContext(context.Context) PrePostStepResponseArrayOutput
}

type PrePostStepResponseArray []PrePostStepResponseInput

func (PrePostStepResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrePostStepResponse)(nil)).Elem()
}

func (i PrePostStepResponseArray) ToPrePostStepResponseArrayOutput() PrePostStepResponseArrayOutput {
	return i.ToPrePostStepResponseArrayOutputWithContext(context.Background())
}

func (i PrePostStepResponseArray) ToPrePostStepResponseArrayOutputWithContext(ctx context.Context) PrePostStepResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrePostStepResponseArrayOutput)
}

type PrePostStepResponseOutput struct{ *pulumi.OutputState }

func (PrePostStepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrePostStepResponse)(nil)).Elem()
}

func (o PrePostStepResponseOutput) ToPrePostStepResponseOutput() PrePostStepResponseOutput {
	return o
}

func (o PrePostStepResponseOutput) ToPrePostStepResponseOutputWithContext(ctx context.Context) PrePostStepResponseOutput {
	return o
}

func (o PrePostStepResponseOutput) StepId() pulumi.StringOutput {
	return o.ApplyT(func(v PrePostStepResponse) string { return v.StepId }).(pulumi.StringOutput)
}

type PrePostStepResponseArrayOutput struct{ *pulumi.OutputState }

func (PrePostStepResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrePostStepResponse)(nil)).Elem()
}

func (o PrePostStepResponseArrayOutput) ToPrePostStepResponseArrayOutput() PrePostStepResponseArrayOutput {
	return o
}

func (o PrePostStepResponseArrayOutput) ToPrePostStepResponseArrayOutputWithContext(ctx context.Context) PrePostStepResponseArrayOutput {
	return o
}

func (o PrePostStepResponseArrayOutput) Index(i pulumi.IntInput) PrePostStepResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrePostStepResponse {
		return vs[0].([]PrePostStepResponse)[vs[1].(int)]
	}).(PrePostStepResponseOutput)
}

type ResourceOperationResponse struct {
	OperationId       string  `pulumi:"operationId"`
	ProvisioningState string  `pulumi:"provisioningState"`
	ResourceName      *string `pulumi:"resourceName"`
	ResourceType      *string `pulumi:"resourceType"`
	StatusCode        string  `pulumi:"statusCode"`
	StatusMessage     string  `pulumi:"statusMessage"`
}





type ResourceOperationResponseInput interface {
	pulumi.Input

	ToResourceOperationResponseOutput() ResourceOperationResponseOutput
	ToResourceOperationResponseOutputWithContext(context.Context) ResourceOperationResponseOutput
}

type ResourceOperationResponseArgs struct {
	OperationId       pulumi.StringInput    `pulumi:"operationId"`
	ProvisioningState pulumi.StringInput    `pulumi:"provisioningState"`
	ResourceName      pulumi.StringPtrInput `pulumi:"resourceName"`
	ResourceType      pulumi.StringPtrInput `pulumi:"resourceType"`
	StatusCode        pulumi.StringInput    `pulumi:"statusCode"`
	StatusMessage     pulumi.StringInput    `pulumi:"statusMessage"`
}

func (ResourceOperationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceOperationResponse)(nil)).Elem()
}

func (i ResourceOperationResponseArgs) ToResourceOperationResponseOutput() ResourceOperationResponseOutput {
	return i.ToResourceOperationResponseOutputWithContext(context.Background())
}

func (i ResourceOperationResponseArgs) ToResourceOperationResponseOutputWithContext(ctx context.Context) ResourceOperationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOperationResponseOutput)
}





type ResourceOperationResponseArrayInput interface {
	pulumi.Input

	ToResourceOperationResponseArrayOutput() ResourceOperationResponseArrayOutput
	ToResourceOperationResponseArrayOutputWithContext(context.Context) ResourceOperationResponseArrayOutput
}

type ResourceOperationResponseArray []ResourceOperationResponseInput

func (ResourceOperationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceOperationResponse)(nil)).Elem()
}

func (i ResourceOperationResponseArray) ToResourceOperationResponseArrayOutput() ResourceOperationResponseArrayOutput {
	return i.ToResourceOperationResponseArrayOutputWithContext(context.Background())
}

func (i ResourceOperationResponseArray) ToResourceOperationResponseArrayOutputWithContext(ctx context.Context) ResourceOperationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOperationResponseArrayOutput)
}

type ResourceOperationResponseOutput struct{ *pulumi.OutputState }

func (ResourceOperationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceOperationResponse)(nil)).Elem()
}

func (o ResourceOperationResponseOutput) ToResourceOperationResponseOutput() ResourceOperationResponseOutput {
	return o
}

func (o ResourceOperationResponseOutput) ToResourceOperationResponseOutputWithContext(ctx context.Context) ResourceOperationResponseOutput {
	return o
}

func (o ResourceOperationResponseOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceOperationResponse) string { return v.OperationId }).(pulumi.StringOutput)
}

func (o ResourceOperationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceOperationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ResourceOperationResponseOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceOperationResponse) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o ResourceOperationResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceOperationResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o ResourceOperationResponseOutput) StatusCode() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceOperationResponse) string { return v.StatusCode }).(pulumi.StringOutput)
}

func (o ResourceOperationResponseOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceOperationResponse) string { return v.StatusMessage }).(pulumi.StringOutput)
}

type ResourceOperationResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceOperationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceOperationResponse)(nil)).Elem()
}

func (o ResourceOperationResponseArrayOutput) ToResourceOperationResponseArrayOutput() ResourceOperationResponseArrayOutput {
	return o
}

func (o ResourceOperationResponseArrayOutput) ToResourceOperationResponseArrayOutputWithContext(ctx context.Context) ResourceOperationResponseArrayOutput {
	return o
}

func (o ResourceOperationResponseArrayOutput) Index(i pulumi.IntInput) ResourceOperationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceOperationResponse {
		return vs[0].([]ResourceOperationResponse)[vs[1].(int)]
	}).(ResourceOperationResponseOutput)
}

type RolloutOperationInfoResponse struct {
	EndTime              string                 `pulumi:"endTime"`
	Error                CloudErrorBodyResponse `pulumi:"error"`
	RetryAttempt         int                    `pulumi:"retryAttempt"`
	SkipSucceededOnRetry bool                   `pulumi:"skipSucceededOnRetry"`
	StartTime            string                 `pulumi:"startTime"`
}





type RolloutOperationInfoResponseInput interface {
	pulumi.Input

	ToRolloutOperationInfoResponseOutput() RolloutOperationInfoResponseOutput
	ToRolloutOperationInfoResponseOutputWithContext(context.Context) RolloutOperationInfoResponseOutput
}

type RolloutOperationInfoResponseArgs struct {
	EndTime              pulumi.StringInput          `pulumi:"endTime"`
	Error                CloudErrorBodyResponseInput `pulumi:"error"`
	RetryAttempt         pulumi.IntInput             `pulumi:"retryAttempt"`
	SkipSucceededOnRetry pulumi.BoolInput            `pulumi:"skipSucceededOnRetry"`
	StartTime            pulumi.StringInput          `pulumi:"startTime"`
}

func (RolloutOperationInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RolloutOperationInfoResponse)(nil)).Elem()
}

func (i RolloutOperationInfoResponseArgs) ToRolloutOperationInfoResponseOutput() RolloutOperationInfoResponseOutput {
	return i.ToRolloutOperationInfoResponseOutputWithContext(context.Background())
}

func (i RolloutOperationInfoResponseArgs) ToRolloutOperationInfoResponseOutputWithContext(ctx context.Context) RolloutOperationInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolloutOperationInfoResponseOutput)
}

type RolloutOperationInfoResponseOutput struct{ *pulumi.OutputState }

func (RolloutOperationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RolloutOperationInfoResponse)(nil)).Elem()
}

func (o RolloutOperationInfoResponseOutput) ToRolloutOperationInfoResponseOutput() RolloutOperationInfoResponseOutput {
	return o
}

func (o RolloutOperationInfoResponseOutput) ToRolloutOperationInfoResponseOutputWithContext(ctx context.Context) RolloutOperationInfoResponseOutput {
	return o
}

func (o RolloutOperationInfoResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v RolloutOperationInfoResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o RolloutOperationInfoResponseOutput) Error() CloudErrorBodyResponseOutput {
	return o.ApplyT(func(v RolloutOperationInfoResponse) CloudErrorBodyResponse { return v.Error }).(CloudErrorBodyResponseOutput)
}

func (o RolloutOperationInfoResponseOutput) RetryAttempt() pulumi.IntOutput {
	return o.ApplyT(func(v RolloutOperationInfoResponse) int { return v.RetryAttempt }).(pulumi.IntOutput)
}

func (o RolloutOperationInfoResponseOutput) SkipSucceededOnRetry() pulumi.BoolOutput {
	return o.ApplyT(func(v RolloutOperationInfoResponse) bool { return v.SkipSucceededOnRetry }).(pulumi.BoolOutput)
}

func (o RolloutOperationInfoResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v RolloutOperationInfoResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

type RolloutStepResponse struct {
	Messages           []MessageResponse           `pulumi:"messages"`
	Name               string                      `pulumi:"name"`
	OperationInfo      StepOperationInfoResponse   `pulumi:"operationInfo"`
	ResourceOperations []ResourceOperationResponse `pulumi:"resourceOperations"`
	Status             string                      `pulumi:"status"`
	StepGroup          *string                     `pulumi:"stepGroup"`
}





type RolloutStepResponseInput interface {
	pulumi.Input

	ToRolloutStepResponseOutput() RolloutStepResponseOutput
	ToRolloutStepResponseOutputWithContext(context.Context) RolloutStepResponseOutput
}

type RolloutStepResponseArgs struct {
	Messages           MessageResponseArrayInput           `pulumi:"messages"`
	Name               pulumi.StringInput                  `pulumi:"name"`
	OperationInfo      StepOperationInfoResponseInput      `pulumi:"operationInfo"`
	ResourceOperations ResourceOperationResponseArrayInput `pulumi:"resourceOperations"`
	Status             pulumi.StringInput                  `pulumi:"status"`
	StepGroup          pulumi.StringPtrInput               `pulumi:"stepGroup"`
}

func (RolloutStepResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RolloutStepResponse)(nil)).Elem()
}

func (i RolloutStepResponseArgs) ToRolloutStepResponseOutput() RolloutStepResponseOutput {
	return i.ToRolloutStepResponseOutputWithContext(context.Background())
}

func (i RolloutStepResponseArgs) ToRolloutStepResponseOutputWithContext(ctx context.Context) RolloutStepResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolloutStepResponseOutput)
}





type RolloutStepResponseArrayInput interface {
	pulumi.Input

	ToRolloutStepResponseArrayOutput() RolloutStepResponseArrayOutput
	ToRolloutStepResponseArrayOutputWithContext(context.Context) RolloutStepResponseArrayOutput
}

type RolloutStepResponseArray []RolloutStepResponseInput

func (RolloutStepResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RolloutStepResponse)(nil)).Elem()
}

func (i RolloutStepResponseArray) ToRolloutStepResponseArrayOutput() RolloutStepResponseArrayOutput {
	return i.ToRolloutStepResponseArrayOutputWithContext(context.Background())
}

func (i RolloutStepResponseArray) ToRolloutStepResponseArrayOutputWithContext(ctx context.Context) RolloutStepResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolloutStepResponseArrayOutput)
}

type RolloutStepResponseOutput struct{ *pulumi.OutputState }

func (RolloutStepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RolloutStepResponse)(nil)).Elem()
}

func (o RolloutStepResponseOutput) ToRolloutStepResponseOutput() RolloutStepResponseOutput {
	return o
}

func (o RolloutStepResponseOutput) ToRolloutStepResponseOutputWithContext(ctx context.Context) RolloutStepResponseOutput {
	return o
}

func (o RolloutStepResponseOutput) Messages() MessageResponseArrayOutput {
	return o.ApplyT(func(v RolloutStepResponse) []MessageResponse { return v.Messages }).(MessageResponseArrayOutput)
}

func (o RolloutStepResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RolloutStepResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RolloutStepResponseOutput) OperationInfo() StepOperationInfoResponseOutput {
	return o.ApplyT(func(v RolloutStepResponse) StepOperationInfoResponse { return v.OperationInfo }).(StepOperationInfoResponseOutput)
}

func (o RolloutStepResponseOutput) ResourceOperations() ResourceOperationResponseArrayOutput {
	return o.ApplyT(func(v RolloutStepResponse) []ResourceOperationResponse { return v.ResourceOperations }).(ResourceOperationResponseArrayOutput)
}

func (o RolloutStepResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v RolloutStepResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o RolloutStepResponseOutput) StepGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RolloutStepResponse) *string { return v.StepGroup }).(pulumi.StringPtrOutput)
}

type RolloutStepResponseArrayOutput struct{ *pulumi.OutputState }

func (RolloutStepResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RolloutStepResponse)(nil)).Elem()
}

func (o RolloutStepResponseArrayOutput) ToRolloutStepResponseArrayOutput() RolloutStepResponseArrayOutput {
	return o
}

func (o RolloutStepResponseArrayOutput) ToRolloutStepResponseArrayOutputWithContext(ctx context.Context) RolloutStepResponseArrayOutput {
	return o
}

func (o RolloutStepResponseArrayOutput) Index(i pulumi.IntInput) RolloutStepResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RolloutStepResponse {
		return vs[0].([]RolloutStepResponse)[vs[1].(int)]
	}).(RolloutStepResponseOutput)
}

type SasAuthentication struct {
	SasUri string `pulumi:"sasUri"`
	Type   string `pulumi:"type"`
}





type SasAuthenticationInput interface {
	pulumi.Input

	ToSasAuthenticationOutput() SasAuthenticationOutput
	ToSasAuthenticationOutputWithContext(context.Context) SasAuthenticationOutput
}

type SasAuthenticationArgs struct {
	SasUri pulumi.StringInput `pulumi:"sasUri"`
	Type   pulumi.StringInput `pulumi:"type"`
}

func (SasAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SasAuthentication)(nil)).Elem()
}

func (i SasAuthenticationArgs) ToSasAuthenticationOutput() SasAuthenticationOutput {
	return i.ToSasAuthenticationOutputWithContext(context.Background())
}

func (i SasAuthenticationArgs) ToSasAuthenticationOutputWithContext(ctx context.Context) SasAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SasAuthenticationOutput)
}

func (i SasAuthenticationArgs) ToSasAuthenticationPtrOutput() SasAuthenticationPtrOutput {
	return i.ToSasAuthenticationPtrOutputWithContext(context.Background())
}

func (i SasAuthenticationArgs) ToSasAuthenticationPtrOutputWithContext(ctx context.Context) SasAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SasAuthenticationOutput).ToSasAuthenticationPtrOutputWithContext(ctx)
}









type SasAuthenticationPtrInput interface {
	pulumi.Input

	ToSasAuthenticationPtrOutput() SasAuthenticationPtrOutput
	ToSasAuthenticationPtrOutputWithContext(context.Context) SasAuthenticationPtrOutput
}

type sasAuthenticationPtrType SasAuthenticationArgs

func SasAuthenticationPtr(v *SasAuthenticationArgs) SasAuthenticationPtrInput {
	return (*sasAuthenticationPtrType)(v)
}

func (*sasAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SasAuthentication)(nil)).Elem()
}

func (i *sasAuthenticationPtrType) ToSasAuthenticationPtrOutput() SasAuthenticationPtrOutput {
	return i.ToSasAuthenticationPtrOutputWithContext(context.Background())
}

func (i *sasAuthenticationPtrType) ToSasAuthenticationPtrOutputWithContext(ctx context.Context) SasAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SasAuthenticationPtrOutput)
}

type SasAuthenticationOutput struct{ *pulumi.OutputState }

func (SasAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SasAuthentication)(nil)).Elem()
}

func (o SasAuthenticationOutput) ToSasAuthenticationOutput() SasAuthenticationOutput {
	return o
}

func (o SasAuthenticationOutput) ToSasAuthenticationOutputWithContext(ctx context.Context) SasAuthenticationOutput {
	return o
}

func (o SasAuthenticationOutput) ToSasAuthenticationPtrOutput() SasAuthenticationPtrOutput {
	return o.ToSasAuthenticationPtrOutputWithContext(context.Background())
}

func (o SasAuthenticationOutput) ToSasAuthenticationPtrOutputWithContext(ctx context.Context) SasAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SasAuthentication) *SasAuthentication {
		return &v
	}).(SasAuthenticationPtrOutput)
}

func (o SasAuthenticationOutput) SasUri() pulumi.StringOutput {
	return o.ApplyT(func(v SasAuthentication) string { return v.SasUri }).(pulumi.StringOutput)
}

func (o SasAuthenticationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SasAuthentication) string { return v.Type }).(pulumi.StringOutput)
}

type SasAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (SasAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SasAuthentication)(nil)).Elem()
}

func (o SasAuthenticationPtrOutput) ToSasAuthenticationPtrOutput() SasAuthenticationPtrOutput {
	return o
}

func (o SasAuthenticationPtrOutput) ToSasAuthenticationPtrOutputWithContext(ctx context.Context) SasAuthenticationPtrOutput {
	return o
}

func (o SasAuthenticationPtrOutput) Elem() SasAuthenticationOutput {
	return o.ApplyT(func(v *SasAuthentication) SasAuthentication {
		if v != nil {
			return *v
		}
		var ret SasAuthentication
		return ret
	}).(SasAuthenticationOutput)
}

func (o SasAuthenticationPtrOutput) SasUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SasAuthentication) *string {
		if v == nil {
			return nil
		}
		return &v.SasUri
	}).(pulumi.StringPtrOutput)
}

func (o SasAuthenticationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SasAuthentication) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type SasAuthenticationResponse struct {
	SasUri string `pulumi:"sasUri"`
	Type   string `pulumi:"type"`
}





type SasAuthenticationResponseInput interface {
	pulumi.Input

	ToSasAuthenticationResponseOutput() SasAuthenticationResponseOutput
	ToSasAuthenticationResponseOutputWithContext(context.Context) SasAuthenticationResponseOutput
}

type SasAuthenticationResponseArgs struct {
	SasUri pulumi.StringInput `pulumi:"sasUri"`
	Type   pulumi.StringInput `pulumi:"type"`
}

func (SasAuthenticationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SasAuthenticationResponse)(nil)).Elem()
}

func (i SasAuthenticationResponseArgs) ToSasAuthenticationResponseOutput() SasAuthenticationResponseOutput {
	return i.ToSasAuthenticationResponseOutputWithContext(context.Background())
}

func (i SasAuthenticationResponseArgs) ToSasAuthenticationResponseOutputWithContext(ctx context.Context) SasAuthenticationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SasAuthenticationResponseOutput)
}

func (i SasAuthenticationResponseArgs) ToSasAuthenticationResponsePtrOutput() SasAuthenticationResponsePtrOutput {
	return i.ToSasAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i SasAuthenticationResponseArgs) ToSasAuthenticationResponsePtrOutputWithContext(ctx context.Context) SasAuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SasAuthenticationResponseOutput).ToSasAuthenticationResponsePtrOutputWithContext(ctx)
}









type SasAuthenticationResponsePtrInput interface {
	pulumi.Input

	ToSasAuthenticationResponsePtrOutput() SasAuthenticationResponsePtrOutput
	ToSasAuthenticationResponsePtrOutputWithContext(context.Context) SasAuthenticationResponsePtrOutput
}

type sasAuthenticationResponsePtrType SasAuthenticationResponseArgs

func SasAuthenticationResponsePtr(v *SasAuthenticationResponseArgs) SasAuthenticationResponsePtrInput {
	return (*sasAuthenticationResponsePtrType)(v)
}

func (*sasAuthenticationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SasAuthenticationResponse)(nil)).Elem()
}

func (i *sasAuthenticationResponsePtrType) ToSasAuthenticationResponsePtrOutput() SasAuthenticationResponsePtrOutput {
	return i.ToSasAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i *sasAuthenticationResponsePtrType) ToSasAuthenticationResponsePtrOutputWithContext(ctx context.Context) SasAuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SasAuthenticationResponsePtrOutput)
}

type SasAuthenticationResponseOutput struct{ *pulumi.OutputState }

func (SasAuthenticationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SasAuthenticationResponse)(nil)).Elem()
}

func (o SasAuthenticationResponseOutput) ToSasAuthenticationResponseOutput() SasAuthenticationResponseOutput {
	return o
}

func (o SasAuthenticationResponseOutput) ToSasAuthenticationResponseOutputWithContext(ctx context.Context) SasAuthenticationResponseOutput {
	return o
}

func (o SasAuthenticationResponseOutput) ToSasAuthenticationResponsePtrOutput() SasAuthenticationResponsePtrOutput {
	return o.ToSasAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (o SasAuthenticationResponseOutput) ToSasAuthenticationResponsePtrOutputWithContext(ctx context.Context) SasAuthenticationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SasAuthenticationResponse) *SasAuthenticationResponse {
		return &v
	}).(SasAuthenticationResponsePtrOutput)
}

func (o SasAuthenticationResponseOutput) SasUri() pulumi.StringOutput {
	return o.ApplyT(func(v SasAuthenticationResponse) string { return v.SasUri }).(pulumi.StringOutput)
}

func (o SasAuthenticationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SasAuthenticationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SasAuthenticationResponsePtrOutput struct{ *pulumi.OutputState }

func (SasAuthenticationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SasAuthenticationResponse)(nil)).Elem()
}

func (o SasAuthenticationResponsePtrOutput) ToSasAuthenticationResponsePtrOutput() SasAuthenticationResponsePtrOutput {
	return o
}

func (o SasAuthenticationResponsePtrOutput) ToSasAuthenticationResponsePtrOutputWithContext(ctx context.Context) SasAuthenticationResponsePtrOutput {
	return o
}

func (o SasAuthenticationResponsePtrOutput) Elem() SasAuthenticationResponseOutput {
	return o.ApplyT(func(v *SasAuthenticationResponse) SasAuthenticationResponse {
		if v != nil {
			return *v
		}
		var ret SasAuthenticationResponse
		return ret
	}).(SasAuthenticationResponseOutput)
}

func (o SasAuthenticationResponsePtrOutput) SasUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SasAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SasUri
	}).(pulumi.StringPtrOutput)
}

func (o SasAuthenticationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SasAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ServiceResponse struct {
	Name                 *string               `pulumi:"name"`
	ServiceUnits         []ServiceUnitResponse `pulumi:"serviceUnits"`
	TargetLocation       string                `pulumi:"targetLocation"`
	TargetSubscriptionId string                `pulumi:"targetSubscriptionId"`
}





type ServiceResponseInput interface {
	pulumi.Input

	ToServiceResponseOutput() ServiceResponseOutput
	ToServiceResponseOutputWithContext(context.Context) ServiceResponseOutput
}

type ServiceResponseArgs struct {
	Name                 pulumi.StringPtrInput         `pulumi:"name"`
	ServiceUnits         ServiceUnitResponseArrayInput `pulumi:"serviceUnits"`
	TargetLocation       pulumi.StringInput            `pulumi:"targetLocation"`
	TargetSubscriptionId pulumi.StringInput            `pulumi:"targetSubscriptionId"`
}

func (ServiceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResponse)(nil)).Elem()
}

func (i ServiceResponseArgs) ToServiceResponseOutput() ServiceResponseOutput {
	return i.ToServiceResponseOutputWithContext(context.Background())
}

func (i ServiceResponseArgs) ToServiceResponseOutputWithContext(ctx context.Context) ServiceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceResponseOutput)
}





type ServiceResponseArrayInput interface {
	pulumi.Input

	ToServiceResponseArrayOutput() ServiceResponseArrayOutput
	ToServiceResponseArrayOutputWithContext(context.Context) ServiceResponseArrayOutput
}

type ServiceResponseArray []ServiceResponseInput

func (ServiceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResponse)(nil)).Elem()
}

func (i ServiceResponseArray) ToServiceResponseArrayOutput() ServiceResponseArrayOutput {
	return i.ToServiceResponseArrayOutputWithContext(context.Background())
}

func (i ServiceResponseArray) ToServiceResponseArrayOutputWithContext(ctx context.Context) ServiceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceResponseArrayOutput)
}

type ServiceResponseOutput struct{ *pulumi.OutputState }

func (ServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResponse)(nil)).Elem()
}

func (o ServiceResponseOutput) ToServiceResponseOutput() ServiceResponseOutput {
	return o
}

func (o ServiceResponseOutput) ToServiceResponseOutputWithContext(ctx context.Context) ServiceResponseOutput {
	return o
}

func (o ServiceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceResponseOutput) ServiceUnits() ServiceUnitResponseArrayOutput {
	return o.ApplyT(func(v ServiceResponse) []ServiceUnitResponse { return v.ServiceUnits }).(ServiceUnitResponseArrayOutput)
}

func (o ServiceResponseOutput) TargetLocation() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResponse) string { return v.TargetLocation }).(pulumi.StringOutput)
}

func (o ServiceResponseOutput) TargetSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResponse) string { return v.TargetSubscriptionId }).(pulumi.StringOutput)
}

type ServiceResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResponse)(nil)).Elem()
}

func (o ServiceResponseArrayOutput) ToServiceResponseArrayOutput() ServiceResponseArrayOutput {
	return o
}

func (o ServiceResponseArrayOutput) ToServiceResponseArrayOutputWithContext(ctx context.Context) ServiceResponseArrayOutput {
	return o
}

func (o ServiceResponseArrayOutput) Index(i pulumi.IntInput) ServiceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceResponse {
		return vs[0].([]ServiceResponse)[vs[1].(int)]
	}).(ServiceResponseOutput)
}

type ServiceUnitArtifacts struct {
	ParametersArtifactSourceRelativePath *string `pulumi:"parametersArtifactSourceRelativePath"`
	ParametersUri                        *string `pulumi:"parametersUri"`
	TemplateArtifactSourceRelativePath   *string `pulumi:"templateArtifactSourceRelativePath"`
	TemplateUri                          *string `pulumi:"templateUri"`
}





type ServiceUnitArtifactsInput interface {
	pulumi.Input

	ToServiceUnitArtifactsOutput() ServiceUnitArtifactsOutput
	ToServiceUnitArtifactsOutputWithContext(context.Context) ServiceUnitArtifactsOutput
}

type ServiceUnitArtifactsArgs struct {
	ParametersArtifactSourceRelativePath pulumi.StringPtrInput `pulumi:"parametersArtifactSourceRelativePath"`
	ParametersUri                        pulumi.StringPtrInput `pulumi:"parametersUri"`
	TemplateArtifactSourceRelativePath   pulumi.StringPtrInput `pulumi:"templateArtifactSourceRelativePath"`
	TemplateUri                          pulumi.StringPtrInput `pulumi:"templateUri"`
}

func (ServiceUnitArtifactsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceUnitArtifacts)(nil)).Elem()
}

func (i ServiceUnitArtifactsArgs) ToServiceUnitArtifactsOutput() ServiceUnitArtifactsOutput {
	return i.ToServiceUnitArtifactsOutputWithContext(context.Background())
}

func (i ServiceUnitArtifactsArgs) ToServiceUnitArtifactsOutputWithContext(ctx context.Context) ServiceUnitArtifactsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitArtifactsOutput)
}

func (i ServiceUnitArtifactsArgs) ToServiceUnitArtifactsPtrOutput() ServiceUnitArtifactsPtrOutput {
	return i.ToServiceUnitArtifactsPtrOutputWithContext(context.Background())
}

func (i ServiceUnitArtifactsArgs) ToServiceUnitArtifactsPtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitArtifactsOutput).ToServiceUnitArtifactsPtrOutputWithContext(ctx)
}









type ServiceUnitArtifactsPtrInput interface {
	pulumi.Input

	ToServiceUnitArtifactsPtrOutput() ServiceUnitArtifactsPtrOutput
	ToServiceUnitArtifactsPtrOutputWithContext(context.Context) ServiceUnitArtifactsPtrOutput
}

type serviceUnitArtifactsPtrType ServiceUnitArtifactsArgs

func ServiceUnitArtifactsPtr(v *ServiceUnitArtifactsArgs) ServiceUnitArtifactsPtrInput {
	return (*serviceUnitArtifactsPtrType)(v)
}

func (*serviceUnitArtifactsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceUnitArtifacts)(nil)).Elem()
}

func (i *serviceUnitArtifactsPtrType) ToServiceUnitArtifactsPtrOutput() ServiceUnitArtifactsPtrOutput {
	return i.ToServiceUnitArtifactsPtrOutputWithContext(context.Background())
}

func (i *serviceUnitArtifactsPtrType) ToServiceUnitArtifactsPtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitArtifactsPtrOutput)
}

type ServiceUnitArtifactsOutput struct{ *pulumi.OutputState }

func (ServiceUnitArtifactsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceUnitArtifacts)(nil)).Elem()
}

func (o ServiceUnitArtifactsOutput) ToServiceUnitArtifactsOutput() ServiceUnitArtifactsOutput {
	return o
}

func (o ServiceUnitArtifactsOutput) ToServiceUnitArtifactsOutputWithContext(ctx context.Context) ServiceUnitArtifactsOutput {
	return o
}

func (o ServiceUnitArtifactsOutput) ToServiceUnitArtifactsPtrOutput() ServiceUnitArtifactsPtrOutput {
	return o.ToServiceUnitArtifactsPtrOutputWithContext(context.Background())
}

func (o ServiceUnitArtifactsOutput) ToServiceUnitArtifactsPtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceUnitArtifacts) *ServiceUnitArtifacts {
		return &v
	}).(ServiceUnitArtifactsPtrOutput)
}

func (o ServiceUnitArtifactsOutput) ParametersArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifacts) *string { return v.ParametersArtifactSourceRelativePath }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsOutput) ParametersUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifacts) *string { return v.ParametersUri }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsOutput) TemplateArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifacts) *string { return v.TemplateArtifactSourceRelativePath }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsOutput) TemplateUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifacts) *string { return v.TemplateUri }).(pulumi.StringPtrOutput)
}

type ServiceUnitArtifactsPtrOutput struct{ *pulumi.OutputState }

func (ServiceUnitArtifactsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceUnitArtifacts)(nil)).Elem()
}

func (o ServiceUnitArtifactsPtrOutput) ToServiceUnitArtifactsPtrOutput() ServiceUnitArtifactsPtrOutput {
	return o
}

func (o ServiceUnitArtifactsPtrOutput) ToServiceUnitArtifactsPtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsPtrOutput {
	return o
}

func (o ServiceUnitArtifactsPtrOutput) Elem() ServiceUnitArtifactsOutput {
	return o.ApplyT(func(v *ServiceUnitArtifacts) ServiceUnitArtifacts {
		if v != nil {
			return *v
		}
		var ret ServiceUnitArtifacts
		return ret
	}).(ServiceUnitArtifactsOutput)
}

func (o ServiceUnitArtifactsPtrOutput) ParametersArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifacts) *string {
		if v == nil {
			return nil
		}
		return v.ParametersArtifactSourceRelativePath
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsPtrOutput) ParametersUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifacts) *string {
		if v == nil {
			return nil
		}
		return v.ParametersUri
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsPtrOutput) TemplateArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifacts) *string {
		if v == nil {
			return nil
		}
		return v.TemplateArtifactSourceRelativePath
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsPtrOutput) TemplateUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifacts) *string {
		if v == nil {
			return nil
		}
		return v.TemplateUri
	}).(pulumi.StringPtrOutput)
}

type ServiceUnitArtifactsResponse struct {
	ParametersArtifactSourceRelativePath *string `pulumi:"parametersArtifactSourceRelativePath"`
	ParametersUri                        *string `pulumi:"parametersUri"`
	TemplateArtifactSourceRelativePath   *string `pulumi:"templateArtifactSourceRelativePath"`
	TemplateUri                          *string `pulumi:"templateUri"`
}





type ServiceUnitArtifactsResponseInput interface {
	pulumi.Input

	ToServiceUnitArtifactsResponseOutput() ServiceUnitArtifactsResponseOutput
	ToServiceUnitArtifactsResponseOutputWithContext(context.Context) ServiceUnitArtifactsResponseOutput
}

type ServiceUnitArtifactsResponseArgs struct {
	ParametersArtifactSourceRelativePath pulumi.StringPtrInput `pulumi:"parametersArtifactSourceRelativePath"`
	ParametersUri                        pulumi.StringPtrInput `pulumi:"parametersUri"`
	TemplateArtifactSourceRelativePath   pulumi.StringPtrInput `pulumi:"templateArtifactSourceRelativePath"`
	TemplateUri                          pulumi.StringPtrInput `pulumi:"templateUri"`
}

func (ServiceUnitArtifactsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceUnitArtifactsResponse)(nil)).Elem()
}

func (i ServiceUnitArtifactsResponseArgs) ToServiceUnitArtifactsResponseOutput() ServiceUnitArtifactsResponseOutput {
	return i.ToServiceUnitArtifactsResponseOutputWithContext(context.Background())
}

func (i ServiceUnitArtifactsResponseArgs) ToServiceUnitArtifactsResponseOutputWithContext(ctx context.Context) ServiceUnitArtifactsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitArtifactsResponseOutput)
}

func (i ServiceUnitArtifactsResponseArgs) ToServiceUnitArtifactsResponsePtrOutput() ServiceUnitArtifactsResponsePtrOutput {
	return i.ToServiceUnitArtifactsResponsePtrOutputWithContext(context.Background())
}

func (i ServiceUnitArtifactsResponseArgs) ToServiceUnitArtifactsResponsePtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitArtifactsResponseOutput).ToServiceUnitArtifactsResponsePtrOutputWithContext(ctx)
}









type ServiceUnitArtifactsResponsePtrInput interface {
	pulumi.Input

	ToServiceUnitArtifactsResponsePtrOutput() ServiceUnitArtifactsResponsePtrOutput
	ToServiceUnitArtifactsResponsePtrOutputWithContext(context.Context) ServiceUnitArtifactsResponsePtrOutput
}

type serviceUnitArtifactsResponsePtrType ServiceUnitArtifactsResponseArgs

func ServiceUnitArtifactsResponsePtr(v *ServiceUnitArtifactsResponseArgs) ServiceUnitArtifactsResponsePtrInput {
	return (*serviceUnitArtifactsResponsePtrType)(v)
}

func (*serviceUnitArtifactsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceUnitArtifactsResponse)(nil)).Elem()
}

func (i *serviceUnitArtifactsResponsePtrType) ToServiceUnitArtifactsResponsePtrOutput() ServiceUnitArtifactsResponsePtrOutput {
	return i.ToServiceUnitArtifactsResponsePtrOutputWithContext(context.Background())
}

func (i *serviceUnitArtifactsResponsePtrType) ToServiceUnitArtifactsResponsePtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitArtifactsResponsePtrOutput)
}

type ServiceUnitArtifactsResponseOutput struct{ *pulumi.OutputState }

func (ServiceUnitArtifactsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceUnitArtifactsResponse)(nil)).Elem()
}

func (o ServiceUnitArtifactsResponseOutput) ToServiceUnitArtifactsResponseOutput() ServiceUnitArtifactsResponseOutput {
	return o
}

func (o ServiceUnitArtifactsResponseOutput) ToServiceUnitArtifactsResponseOutputWithContext(ctx context.Context) ServiceUnitArtifactsResponseOutput {
	return o
}

func (o ServiceUnitArtifactsResponseOutput) ToServiceUnitArtifactsResponsePtrOutput() ServiceUnitArtifactsResponsePtrOutput {
	return o.ToServiceUnitArtifactsResponsePtrOutputWithContext(context.Background())
}

func (o ServiceUnitArtifactsResponseOutput) ToServiceUnitArtifactsResponsePtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceUnitArtifactsResponse) *ServiceUnitArtifactsResponse {
		return &v
	}).(ServiceUnitArtifactsResponsePtrOutput)
}

func (o ServiceUnitArtifactsResponseOutput) ParametersArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifactsResponse) *string { return v.ParametersArtifactSourceRelativePath }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponseOutput) ParametersUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifactsResponse) *string { return v.ParametersUri }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponseOutput) TemplateArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifactsResponse) *string { return v.TemplateArtifactSourceRelativePath }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponseOutput) TemplateUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitArtifactsResponse) *string { return v.TemplateUri }).(pulumi.StringPtrOutput)
}

type ServiceUnitArtifactsResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceUnitArtifactsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceUnitArtifactsResponse)(nil)).Elem()
}

func (o ServiceUnitArtifactsResponsePtrOutput) ToServiceUnitArtifactsResponsePtrOutput() ServiceUnitArtifactsResponsePtrOutput {
	return o
}

func (o ServiceUnitArtifactsResponsePtrOutput) ToServiceUnitArtifactsResponsePtrOutputWithContext(ctx context.Context) ServiceUnitArtifactsResponsePtrOutput {
	return o
}

func (o ServiceUnitArtifactsResponsePtrOutput) Elem() ServiceUnitArtifactsResponseOutput {
	return o.ApplyT(func(v *ServiceUnitArtifactsResponse) ServiceUnitArtifactsResponse {
		if v != nil {
			return *v
		}
		var ret ServiceUnitArtifactsResponse
		return ret
	}).(ServiceUnitArtifactsResponseOutput)
}

func (o ServiceUnitArtifactsResponsePtrOutput) ParametersArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifactsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ParametersArtifactSourceRelativePath
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponsePtrOutput) ParametersUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifactsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ParametersUri
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponsePtrOutput) TemplateArtifactSourceRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifactsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TemplateArtifactSourceRelativePath
	}).(pulumi.StringPtrOutput)
}

func (o ServiceUnitArtifactsResponsePtrOutput) TemplateUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceUnitArtifactsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TemplateUri
	}).(pulumi.StringPtrOutput)
}

type ServiceUnitResponse struct {
	Artifacts           *ServiceUnitArtifactsResponse `pulumi:"artifacts"`
	DeploymentMode      string                        `pulumi:"deploymentMode"`
	Name                *string                       `pulumi:"name"`
	Steps               []RolloutStepResponse         `pulumi:"steps"`
	TargetResourceGroup string                        `pulumi:"targetResourceGroup"`
}





type ServiceUnitResponseInput interface {
	pulumi.Input

	ToServiceUnitResponseOutput() ServiceUnitResponseOutput
	ToServiceUnitResponseOutputWithContext(context.Context) ServiceUnitResponseOutput
}

type ServiceUnitResponseArgs struct {
	Artifacts           ServiceUnitArtifactsResponsePtrInput `pulumi:"artifacts"`
	DeploymentMode      pulumi.StringInput                   `pulumi:"deploymentMode"`
	Name                pulumi.StringPtrInput                `pulumi:"name"`
	Steps               RolloutStepResponseArrayInput        `pulumi:"steps"`
	TargetResourceGroup pulumi.StringInput                   `pulumi:"targetResourceGroup"`
}

func (ServiceUnitResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceUnitResponse)(nil)).Elem()
}

func (i ServiceUnitResponseArgs) ToServiceUnitResponseOutput() ServiceUnitResponseOutput {
	return i.ToServiceUnitResponseOutputWithContext(context.Background())
}

func (i ServiceUnitResponseArgs) ToServiceUnitResponseOutputWithContext(ctx context.Context) ServiceUnitResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitResponseOutput)
}





type ServiceUnitResponseArrayInput interface {
	pulumi.Input

	ToServiceUnitResponseArrayOutput() ServiceUnitResponseArrayOutput
	ToServiceUnitResponseArrayOutputWithContext(context.Context) ServiceUnitResponseArrayOutput
}

type ServiceUnitResponseArray []ServiceUnitResponseInput

func (ServiceUnitResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceUnitResponse)(nil)).Elem()
}

func (i ServiceUnitResponseArray) ToServiceUnitResponseArrayOutput() ServiceUnitResponseArrayOutput {
	return i.ToServiceUnitResponseArrayOutputWithContext(context.Background())
}

func (i ServiceUnitResponseArray) ToServiceUnitResponseArrayOutputWithContext(ctx context.Context) ServiceUnitResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceUnitResponseArrayOutput)
}

type ServiceUnitResponseOutput struct{ *pulumi.OutputState }

func (ServiceUnitResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceUnitResponse)(nil)).Elem()
}

func (o ServiceUnitResponseOutput) ToServiceUnitResponseOutput() ServiceUnitResponseOutput {
	return o
}

func (o ServiceUnitResponseOutput) ToServiceUnitResponseOutputWithContext(ctx context.Context) ServiceUnitResponseOutput {
	return o
}

func (o ServiceUnitResponseOutput) Artifacts() ServiceUnitArtifactsResponsePtrOutput {
	return o.ApplyT(func(v ServiceUnitResponse) *ServiceUnitArtifactsResponse { return v.Artifacts }).(ServiceUnitArtifactsResponsePtrOutput)
}

func (o ServiceUnitResponseOutput) DeploymentMode() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceUnitResponse) string { return v.DeploymentMode }).(pulumi.StringOutput)
}

func (o ServiceUnitResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceUnitResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceUnitResponseOutput) Steps() RolloutStepResponseArrayOutput {
	return o.ApplyT(func(v ServiceUnitResponse) []RolloutStepResponse { return v.Steps }).(RolloutStepResponseArrayOutput)
}

func (o ServiceUnitResponseOutput) TargetResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceUnitResponse) string { return v.TargetResourceGroup }).(pulumi.StringOutput)
}

type ServiceUnitResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceUnitResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceUnitResponse)(nil)).Elem()
}

func (o ServiceUnitResponseArrayOutput) ToServiceUnitResponseArrayOutput() ServiceUnitResponseArrayOutput {
	return o
}

func (o ServiceUnitResponseArrayOutput) ToServiceUnitResponseArrayOutputWithContext(ctx context.Context) ServiceUnitResponseArrayOutput {
	return o
}

func (o ServiceUnitResponseArrayOutput) Index(i pulumi.IntInput) ServiceUnitResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceUnitResponse {
		return vs[0].([]ServiceUnitResponse)[vs[1].(int)]
	}).(ServiceUnitResponseOutput)
}

type StepType struct {
	DependsOnStepGroups []string      `pulumi:"dependsOnStepGroups"`
	DeploymentTargetId  string        `pulumi:"deploymentTargetId"`
	Name                string        `pulumi:"name"`
	PostDeploymentSteps []PrePostStep `pulumi:"postDeploymentSteps"`
	PreDeploymentSteps  []PrePostStep `pulumi:"preDeploymentSteps"`
}





type StepTypeInput interface {
	pulumi.Input

	ToStepTypeOutput() StepTypeOutput
	ToStepTypeOutputWithContext(context.Context) StepTypeOutput
}

type StepTypeArgs struct {
	DependsOnStepGroups pulumi.StringArrayInput `pulumi:"dependsOnStepGroups"`
	DeploymentTargetId  pulumi.StringInput      `pulumi:"deploymentTargetId"`
	Name                pulumi.StringInput      `pulumi:"name"`
	PostDeploymentSteps PrePostStepArrayInput   `pulumi:"postDeploymentSteps"`
	PreDeploymentSteps  PrePostStepArrayInput   `pulumi:"preDeploymentSteps"`
}

func (StepTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StepType)(nil)).Elem()
}

func (i StepTypeArgs) ToStepTypeOutput() StepTypeOutput {
	return i.ToStepTypeOutputWithContext(context.Background())
}

func (i StepTypeArgs) ToStepTypeOutputWithContext(ctx context.Context) StepTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepTypeOutput)
}





type StepTypeArrayInput interface {
	pulumi.Input

	ToStepTypeArrayOutput() StepTypeArrayOutput
	ToStepTypeArrayOutputWithContext(context.Context) StepTypeArrayOutput
}

type StepTypeArray []StepTypeInput

func (StepTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepType)(nil)).Elem()
}

func (i StepTypeArray) ToStepTypeArrayOutput() StepTypeArrayOutput {
	return i.ToStepTypeArrayOutputWithContext(context.Background())
}

func (i StepTypeArray) ToStepTypeArrayOutputWithContext(ctx context.Context) StepTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepTypeArrayOutput)
}

type StepTypeOutput struct{ *pulumi.OutputState }

func (StepTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StepType)(nil)).Elem()
}

func (o StepTypeOutput) ToStepTypeOutput() StepTypeOutput {
	return o
}

func (o StepTypeOutput) ToStepTypeOutputWithContext(ctx context.Context) StepTypeOutput {
	return o
}

func (o StepTypeOutput) DependsOnStepGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StepType) []string { return v.DependsOnStepGroups }).(pulumi.StringArrayOutput)
}

func (o StepTypeOutput) DeploymentTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v StepType) string { return v.DeploymentTargetId }).(pulumi.StringOutput)
}

func (o StepTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StepType) string { return v.Name }).(pulumi.StringOutput)
}

func (o StepTypeOutput) PostDeploymentSteps() PrePostStepArrayOutput {
	return o.ApplyT(func(v StepType) []PrePostStep { return v.PostDeploymentSteps }).(PrePostStepArrayOutput)
}

func (o StepTypeOutput) PreDeploymentSteps() PrePostStepArrayOutput {
	return o.ApplyT(func(v StepType) []PrePostStep { return v.PreDeploymentSteps }).(PrePostStepArrayOutput)
}

type StepTypeArrayOutput struct{ *pulumi.OutputState }

func (StepTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepType)(nil)).Elem()
}

func (o StepTypeArrayOutput) ToStepTypeArrayOutput() StepTypeArrayOutput {
	return o
}

func (o StepTypeArrayOutput) ToStepTypeArrayOutputWithContext(ctx context.Context) StepTypeArrayOutput {
	return o
}

func (o StepTypeArrayOutput) Index(i pulumi.IntInput) StepTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StepType {
		return vs[0].([]StepType)[vs[1].(int)]
	}).(StepTypeOutput)
}

type StepOperationInfoResponse struct {
	CorrelationId   string                  `pulumi:"correlationId"`
	DeploymentName  string                  `pulumi:"deploymentName"`
	EndTime         string                  `pulumi:"endTime"`
	Error           *CloudErrorBodyResponse `pulumi:"error"`
	LastUpdatedTime string                  `pulumi:"lastUpdatedTime"`
	StartTime       string                  `pulumi:"startTime"`
}





type StepOperationInfoResponseInput interface {
	pulumi.Input

	ToStepOperationInfoResponseOutput() StepOperationInfoResponseOutput
	ToStepOperationInfoResponseOutputWithContext(context.Context) StepOperationInfoResponseOutput
}

type StepOperationInfoResponseArgs struct {
	CorrelationId   pulumi.StringInput             `pulumi:"correlationId"`
	DeploymentName  pulumi.StringInput             `pulumi:"deploymentName"`
	EndTime         pulumi.StringInput             `pulumi:"endTime"`
	Error           CloudErrorBodyResponsePtrInput `pulumi:"error"`
	LastUpdatedTime pulumi.StringInput             `pulumi:"lastUpdatedTime"`
	StartTime       pulumi.StringInput             `pulumi:"startTime"`
}

func (StepOperationInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StepOperationInfoResponse)(nil)).Elem()
}

func (i StepOperationInfoResponseArgs) ToStepOperationInfoResponseOutput() StepOperationInfoResponseOutput {
	return i.ToStepOperationInfoResponseOutputWithContext(context.Background())
}

func (i StepOperationInfoResponseArgs) ToStepOperationInfoResponseOutputWithContext(ctx context.Context) StepOperationInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepOperationInfoResponseOutput)
}

type StepOperationInfoResponseOutput struct{ *pulumi.OutputState }

func (StepOperationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StepOperationInfoResponse)(nil)).Elem()
}

func (o StepOperationInfoResponseOutput) ToStepOperationInfoResponseOutput() StepOperationInfoResponseOutput {
	return o
}

func (o StepOperationInfoResponseOutput) ToStepOperationInfoResponseOutputWithContext(ctx context.Context) StepOperationInfoResponseOutput {
	return o
}

func (o StepOperationInfoResponseOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o StepOperationInfoResponseOutput) DeploymentName() pulumi.StringOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) string { return v.DeploymentName }).(pulumi.StringOutput)
}

func (o StepOperationInfoResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o StepOperationInfoResponseOutput) Error() CloudErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) *CloudErrorBodyResponse { return v.Error }).(CloudErrorBodyResponsePtrOutput)
}

func (o StepOperationInfoResponseOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) string { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o StepOperationInfoResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v StepOperationInfoResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

type StepResponse struct {
	DependsOnStepGroups []string              `pulumi:"dependsOnStepGroups"`
	DeploymentTargetId  string                `pulumi:"deploymentTargetId"`
	Name                string                `pulumi:"name"`
	PostDeploymentSteps []PrePostStepResponse `pulumi:"postDeploymentSteps"`
	PreDeploymentSteps  []PrePostStepResponse `pulumi:"preDeploymentSteps"`
}





type StepResponseInput interface {
	pulumi.Input

	ToStepResponseOutput() StepResponseOutput
	ToStepResponseOutputWithContext(context.Context) StepResponseOutput
}

type StepResponseArgs struct {
	DependsOnStepGroups pulumi.StringArrayInput       `pulumi:"dependsOnStepGroups"`
	DeploymentTargetId  pulumi.StringInput            `pulumi:"deploymentTargetId"`
	Name                pulumi.StringInput            `pulumi:"name"`
	PostDeploymentSteps PrePostStepResponseArrayInput `pulumi:"postDeploymentSteps"`
	PreDeploymentSteps  PrePostStepResponseArrayInput `pulumi:"preDeploymentSteps"`
}

func (StepResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StepResponse)(nil)).Elem()
}

func (i StepResponseArgs) ToStepResponseOutput() StepResponseOutput {
	return i.ToStepResponseOutputWithContext(context.Background())
}

func (i StepResponseArgs) ToStepResponseOutputWithContext(ctx context.Context) StepResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepResponseOutput)
}





type StepResponseArrayInput interface {
	pulumi.Input

	ToStepResponseArrayOutput() StepResponseArrayOutput
	ToStepResponseArrayOutputWithContext(context.Context) StepResponseArrayOutput
}

type StepResponseArray []StepResponseInput

func (StepResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepResponse)(nil)).Elem()
}

func (i StepResponseArray) ToStepResponseArrayOutput() StepResponseArrayOutput {
	return i.ToStepResponseArrayOutputWithContext(context.Background())
}

func (i StepResponseArray) ToStepResponseArrayOutputWithContext(ctx context.Context) StepResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepResponseArrayOutput)
}

type StepResponseOutput struct{ *pulumi.OutputState }

func (StepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StepResponse)(nil)).Elem()
}

func (o StepResponseOutput) ToStepResponseOutput() StepResponseOutput {
	return o
}

func (o StepResponseOutput) ToStepResponseOutputWithContext(ctx context.Context) StepResponseOutput {
	return o
}

func (o StepResponseOutput) DependsOnStepGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StepResponse) []string { return v.DependsOnStepGroups }).(pulumi.StringArrayOutput)
}

func (o StepResponseOutput) DeploymentTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v StepResponse) string { return v.DeploymentTargetId }).(pulumi.StringOutput)
}

func (o StepResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StepResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o StepResponseOutput) PostDeploymentSteps() PrePostStepResponseArrayOutput {
	return o.ApplyT(func(v StepResponse) []PrePostStepResponse { return v.PostDeploymentSteps }).(PrePostStepResponseArrayOutput)
}

func (o StepResponseOutput) PreDeploymentSteps() PrePostStepResponseArrayOutput {
	return o.ApplyT(func(v StepResponse) []PrePostStepResponse { return v.PreDeploymentSteps }).(PrePostStepResponseArrayOutput)
}

type StepResponseArrayOutput struct{ *pulumi.OutputState }

func (StepResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepResponse)(nil)).Elem()
}

func (o StepResponseArrayOutput) ToStepResponseArrayOutput() StepResponseArrayOutput {
	return o
}

func (o StepResponseArrayOutput) ToStepResponseArrayOutputWithContext(ctx context.Context) StepResponseArrayOutput {
	return o
}

func (o StepResponseArrayOutput) Index(i pulumi.IntInput) StepResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StepResponse {
		return vs[0].([]StepResponse)[vs[1].(int)]
	}).(StepResponseOutput)
}

type WaitStepAttributes struct {
	Duration string `pulumi:"duration"`
}





type WaitStepAttributesInput interface {
	pulumi.Input

	ToWaitStepAttributesOutput() WaitStepAttributesOutput
	ToWaitStepAttributesOutputWithContext(context.Context) WaitStepAttributesOutput
}

type WaitStepAttributesArgs struct {
	Duration pulumi.StringInput `pulumi:"duration"`
}

func (WaitStepAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepAttributes)(nil)).Elem()
}

func (i WaitStepAttributesArgs) ToWaitStepAttributesOutput() WaitStepAttributesOutput {
	return i.ToWaitStepAttributesOutputWithContext(context.Background())
}

func (i WaitStepAttributesArgs) ToWaitStepAttributesOutputWithContext(ctx context.Context) WaitStepAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepAttributesOutput)
}

func (i WaitStepAttributesArgs) ToWaitStepAttributesPtrOutput() WaitStepAttributesPtrOutput {
	return i.ToWaitStepAttributesPtrOutputWithContext(context.Background())
}

func (i WaitStepAttributesArgs) ToWaitStepAttributesPtrOutputWithContext(ctx context.Context) WaitStepAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepAttributesOutput).ToWaitStepAttributesPtrOutputWithContext(ctx)
}









type WaitStepAttributesPtrInput interface {
	pulumi.Input

	ToWaitStepAttributesPtrOutput() WaitStepAttributesPtrOutput
	ToWaitStepAttributesPtrOutputWithContext(context.Context) WaitStepAttributesPtrOutput
}

type waitStepAttributesPtrType WaitStepAttributesArgs

func WaitStepAttributesPtr(v *WaitStepAttributesArgs) WaitStepAttributesPtrInput {
	return (*waitStepAttributesPtrType)(v)
}

func (*waitStepAttributesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WaitStepAttributes)(nil)).Elem()
}

func (i *waitStepAttributesPtrType) ToWaitStepAttributesPtrOutput() WaitStepAttributesPtrOutput {
	return i.ToWaitStepAttributesPtrOutputWithContext(context.Background())
}

func (i *waitStepAttributesPtrType) ToWaitStepAttributesPtrOutputWithContext(ctx context.Context) WaitStepAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepAttributesPtrOutput)
}

type WaitStepAttributesOutput struct{ *pulumi.OutputState }

func (WaitStepAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepAttributes)(nil)).Elem()
}

func (o WaitStepAttributesOutput) ToWaitStepAttributesOutput() WaitStepAttributesOutput {
	return o
}

func (o WaitStepAttributesOutput) ToWaitStepAttributesOutputWithContext(ctx context.Context) WaitStepAttributesOutput {
	return o
}

func (o WaitStepAttributesOutput) ToWaitStepAttributesPtrOutput() WaitStepAttributesPtrOutput {
	return o.ToWaitStepAttributesPtrOutputWithContext(context.Background())
}

func (o WaitStepAttributesOutput) ToWaitStepAttributesPtrOutputWithContext(ctx context.Context) WaitStepAttributesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WaitStepAttributes) *WaitStepAttributes {
		return &v
	}).(WaitStepAttributesPtrOutput)
}

func (o WaitStepAttributesOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v WaitStepAttributes) string { return v.Duration }).(pulumi.StringOutput)
}

type WaitStepAttributesPtrOutput struct{ *pulumi.OutputState }

func (WaitStepAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WaitStepAttributes)(nil)).Elem()
}

func (o WaitStepAttributesPtrOutput) ToWaitStepAttributesPtrOutput() WaitStepAttributesPtrOutput {
	return o
}

func (o WaitStepAttributesPtrOutput) ToWaitStepAttributesPtrOutputWithContext(ctx context.Context) WaitStepAttributesPtrOutput {
	return o
}

func (o WaitStepAttributesPtrOutput) Elem() WaitStepAttributesOutput {
	return o.ApplyT(func(v *WaitStepAttributes) WaitStepAttributes {
		if v != nil {
			return *v
		}
		var ret WaitStepAttributes
		return ret
	}).(WaitStepAttributesOutput)
}

func (o WaitStepAttributesPtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WaitStepAttributes) *string {
		if v == nil {
			return nil
		}
		return &v.Duration
	}).(pulumi.StringPtrOutput)
}

type WaitStepAttributesResponse struct {
	Duration string `pulumi:"duration"`
}





type WaitStepAttributesResponseInput interface {
	pulumi.Input

	ToWaitStepAttributesResponseOutput() WaitStepAttributesResponseOutput
	ToWaitStepAttributesResponseOutputWithContext(context.Context) WaitStepAttributesResponseOutput
}

type WaitStepAttributesResponseArgs struct {
	Duration pulumi.StringInput `pulumi:"duration"`
}

func (WaitStepAttributesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepAttributesResponse)(nil)).Elem()
}

func (i WaitStepAttributesResponseArgs) ToWaitStepAttributesResponseOutput() WaitStepAttributesResponseOutput {
	return i.ToWaitStepAttributesResponseOutputWithContext(context.Background())
}

func (i WaitStepAttributesResponseArgs) ToWaitStepAttributesResponseOutputWithContext(ctx context.Context) WaitStepAttributesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepAttributesResponseOutput)
}

func (i WaitStepAttributesResponseArgs) ToWaitStepAttributesResponsePtrOutput() WaitStepAttributesResponsePtrOutput {
	return i.ToWaitStepAttributesResponsePtrOutputWithContext(context.Background())
}

func (i WaitStepAttributesResponseArgs) ToWaitStepAttributesResponsePtrOutputWithContext(ctx context.Context) WaitStepAttributesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepAttributesResponseOutput).ToWaitStepAttributesResponsePtrOutputWithContext(ctx)
}









type WaitStepAttributesResponsePtrInput interface {
	pulumi.Input

	ToWaitStepAttributesResponsePtrOutput() WaitStepAttributesResponsePtrOutput
	ToWaitStepAttributesResponsePtrOutputWithContext(context.Context) WaitStepAttributesResponsePtrOutput
}

type waitStepAttributesResponsePtrType WaitStepAttributesResponseArgs

func WaitStepAttributesResponsePtr(v *WaitStepAttributesResponseArgs) WaitStepAttributesResponsePtrInput {
	return (*waitStepAttributesResponsePtrType)(v)
}

func (*waitStepAttributesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WaitStepAttributesResponse)(nil)).Elem()
}

func (i *waitStepAttributesResponsePtrType) ToWaitStepAttributesResponsePtrOutput() WaitStepAttributesResponsePtrOutput {
	return i.ToWaitStepAttributesResponsePtrOutputWithContext(context.Background())
}

func (i *waitStepAttributesResponsePtrType) ToWaitStepAttributesResponsePtrOutputWithContext(ctx context.Context) WaitStepAttributesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepAttributesResponsePtrOutput)
}

type WaitStepAttributesResponseOutput struct{ *pulumi.OutputState }

func (WaitStepAttributesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepAttributesResponse)(nil)).Elem()
}

func (o WaitStepAttributesResponseOutput) ToWaitStepAttributesResponseOutput() WaitStepAttributesResponseOutput {
	return o
}

func (o WaitStepAttributesResponseOutput) ToWaitStepAttributesResponseOutputWithContext(ctx context.Context) WaitStepAttributesResponseOutput {
	return o
}

func (o WaitStepAttributesResponseOutput) ToWaitStepAttributesResponsePtrOutput() WaitStepAttributesResponsePtrOutput {
	return o.ToWaitStepAttributesResponsePtrOutputWithContext(context.Background())
}

func (o WaitStepAttributesResponseOutput) ToWaitStepAttributesResponsePtrOutputWithContext(ctx context.Context) WaitStepAttributesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WaitStepAttributesResponse) *WaitStepAttributesResponse {
		return &v
	}).(WaitStepAttributesResponsePtrOutput)
}

func (o WaitStepAttributesResponseOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v WaitStepAttributesResponse) string { return v.Duration }).(pulumi.StringOutput)
}

type WaitStepAttributesResponsePtrOutput struct{ *pulumi.OutputState }

func (WaitStepAttributesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WaitStepAttributesResponse)(nil)).Elem()
}

func (o WaitStepAttributesResponsePtrOutput) ToWaitStepAttributesResponsePtrOutput() WaitStepAttributesResponsePtrOutput {
	return o
}

func (o WaitStepAttributesResponsePtrOutput) ToWaitStepAttributesResponsePtrOutputWithContext(ctx context.Context) WaitStepAttributesResponsePtrOutput {
	return o
}

func (o WaitStepAttributesResponsePtrOutput) Elem() WaitStepAttributesResponseOutput {
	return o.ApplyT(func(v *WaitStepAttributesResponse) WaitStepAttributesResponse {
		if v != nil {
			return *v
		}
		var ret WaitStepAttributesResponse
		return ret
	}).(WaitStepAttributesResponseOutput)
}

func (o WaitStepAttributesResponsePtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WaitStepAttributesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Duration
	}).(pulumi.StringPtrOutput)
}

type WaitStepProperties struct {
	Attributes *WaitStepAttributes `pulumi:"attributes"`
	StepType   string              `pulumi:"stepType"`
}





type WaitStepPropertiesInput interface {
	pulumi.Input

	ToWaitStepPropertiesOutput() WaitStepPropertiesOutput
	ToWaitStepPropertiesOutputWithContext(context.Context) WaitStepPropertiesOutput
}

type WaitStepPropertiesArgs struct {
	Attributes WaitStepAttributesPtrInput `pulumi:"attributes"`
	StepType   pulumi.StringInput         `pulumi:"stepType"`
}

func (WaitStepPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepProperties)(nil)).Elem()
}

func (i WaitStepPropertiesArgs) ToWaitStepPropertiesOutput() WaitStepPropertiesOutput {
	return i.ToWaitStepPropertiesOutputWithContext(context.Background())
}

func (i WaitStepPropertiesArgs) ToWaitStepPropertiesOutputWithContext(ctx context.Context) WaitStepPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepPropertiesOutput)
}

func (i WaitStepPropertiesArgs) ToWaitStepPropertiesPtrOutput() WaitStepPropertiesPtrOutput {
	return i.ToWaitStepPropertiesPtrOutputWithContext(context.Background())
}

func (i WaitStepPropertiesArgs) ToWaitStepPropertiesPtrOutputWithContext(ctx context.Context) WaitStepPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepPropertiesOutput).ToWaitStepPropertiesPtrOutputWithContext(ctx)
}









type WaitStepPropertiesPtrInput interface {
	pulumi.Input

	ToWaitStepPropertiesPtrOutput() WaitStepPropertiesPtrOutput
	ToWaitStepPropertiesPtrOutputWithContext(context.Context) WaitStepPropertiesPtrOutput
}

type waitStepPropertiesPtrType WaitStepPropertiesArgs

func WaitStepPropertiesPtr(v *WaitStepPropertiesArgs) WaitStepPropertiesPtrInput {
	return (*waitStepPropertiesPtrType)(v)
}

func (*waitStepPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WaitStepProperties)(nil)).Elem()
}

func (i *waitStepPropertiesPtrType) ToWaitStepPropertiesPtrOutput() WaitStepPropertiesPtrOutput {
	return i.ToWaitStepPropertiesPtrOutputWithContext(context.Background())
}

func (i *waitStepPropertiesPtrType) ToWaitStepPropertiesPtrOutputWithContext(ctx context.Context) WaitStepPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepPropertiesPtrOutput)
}

type WaitStepPropertiesOutput struct{ *pulumi.OutputState }

func (WaitStepPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepProperties)(nil)).Elem()
}

func (o WaitStepPropertiesOutput) ToWaitStepPropertiesOutput() WaitStepPropertiesOutput {
	return o
}

func (o WaitStepPropertiesOutput) ToWaitStepPropertiesOutputWithContext(ctx context.Context) WaitStepPropertiesOutput {
	return o
}

func (o WaitStepPropertiesOutput) ToWaitStepPropertiesPtrOutput() WaitStepPropertiesPtrOutput {
	return o.ToWaitStepPropertiesPtrOutputWithContext(context.Background())
}

func (o WaitStepPropertiesOutput) ToWaitStepPropertiesPtrOutputWithContext(ctx context.Context) WaitStepPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WaitStepProperties) *WaitStepProperties {
		return &v
	}).(WaitStepPropertiesPtrOutput)
}

func (o WaitStepPropertiesOutput) Attributes() WaitStepAttributesPtrOutput {
	return o.ApplyT(func(v WaitStepProperties) *WaitStepAttributes { return v.Attributes }).(WaitStepAttributesPtrOutput)
}

func (o WaitStepPropertiesOutput) StepType() pulumi.StringOutput {
	return o.ApplyT(func(v WaitStepProperties) string { return v.StepType }).(pulumi.StringOutput)
}

type WaitStepPropertiesPtrOutput struct{ *pulumi.OutputState }

func (WaitStepPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WaitStepProperties)(nil)).Elem()
}

func (o WaitStepPropertiesPtrOutput) ToWaitStepPropertiesPtrOutput() WaitStepPropertiesPtrOutput {
	return o
}

func (o WaitStepPropertiesPtrOutput) ToWaitStepPropertiesPtrOutputWithContext(ctx context.Context) WaitStepPropertiesPtrOutput {
	return o
}

func (o WaitStepPropertiesPtrOutput) Elem() WaitStepPropertiesOutput {
	return o.ApplyT(func(v *WaitStepProperties) WaitStepProperties {
		if v != nil {
			return *v
		}
		var ret WaitStepProperties
		return ret
	}).(WaitStepPropertiesOutput)
}

func (o WaitStepPropertiesPtrOutput) Attributes() WaitStepAttributesPtrOutput {
	return o.ApplyT(func(v *WaitStepProperties) *WaitStepAttributes {
		if v == nil {
			return nil
		}
		return v.Attributes
	}).(WaitStepAttributesPtrOutput)
}

func (o WaitStepPropertiesPtrOutput) StepType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WaitStepProperties) *string {
		if v == nil {
			return nil
		}
		return &v.StepType
	}).(pulumi.StringPtrOutput)
}

type WaitStepPropertiesResponse struct {
	Attributes *WaitStepAttributesResponse `pulumi:"attributes"`
	StepType   string                      `pulumi:"stepType"`
}





type WaitStepPropertiesResponseInput interface {
	pulumi.Input

	ToWaitStepPropertiesResponseOutput() WaitStepPropertiesResponseOutput
	ToWaitStepPropertiesResponseOutputWithContext(context.Context) WaitStepPropertiesResponseOutput
}

type WaitStepPropertiesResponseArgs struct {
	Attributes WaitStepAttributesResponsePtrInput `pulumi:"attributes"`
	StepType   pulumi.StringInput                 `pulumi:"stepType"`
}

func (WaitStepPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepPropertiesResponse)(nil)).Elem()
}

func (i WaitStepPropertiesResponseArgs) ToWaitStepPropertiesResponseOutput() WaitStepPropertiesResponseOutput {
	return i.ToWaitStepPropertiesResponseOutputWithContext(context.Background())
}

func (i WaitStepPropertiesResponseArgs) ToWaitStepPropertiesResponseOutputWithContext(ctx context.Context) WaitStepPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepPropertiesResponseOutput)
}

func (i WaitStepPropertiesResponseArgs) ToWaitStepPropertiesResponsePtrOutput() WaitStepPropertiesResponsePtrOutput {
	return i.ToWaitStepPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i WaitStepPropertiesResponseArgs) ToWaitStepPropertiesResponsePtrOutputWithContext(ctx context.Context) WaitStepPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepPropertiesResponseOutput).ToWaitStepPropertiesResponsePtrOutputWithContext(ctx)
}









type WaitStepPropertiesResponsePtrInput interface {
	pulumi.Input

	ToWaitStepPropertiesResponsePtrOutput() WaitStepPropertiesResponsePtrOutput
	ToWaitStepPropertiesResponsePtrOutputWithContext(context.Context) WaitStepPropertiesResponsePtrOutput
}

type waitStepPropertiesResponsePtrType WaitStepPropertiesResponseArgs

func WaitStepPropertiesResponsePtr(v *WaitStepPropertiesResponseArgs) WaitStepPropertiesResponsePtrInput {
	return (*waitStepPropertiesResponsePtrType)(v)
}

func (*waitStepPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WaitStepPropertiesResponse)(nil)).Elem()
}

func (i *waitStepPropertiesResponsePtrType) ToWaitStepPropertiesResponsePtrOutput() WaitStepPropertiesResponsePtrOutput {
	return i.ToWaitStepPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *waitStepPropertiesResponsePtrType) ToWaitStepPropertiesResponsePtrOutputWithContext(ctx context.Context) WaitStepPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitStepPropertiesResponsePtrOutput)
}

type WaitStepPropertiesResponseOutput struct{ *pulumi.OutputState }

func (WaitStepPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitStepPropertiesResponse)(nil)).Elem()
}

func (o WaitStepPropertiesResponseOutput) ToWaitStepPropertiesResponseOutput() WaitStepPropertiesResponseOutput {
	return o
}

func (o WaitStepPropertiesResponseOutput) ToWaitStepPropertiesResponseOutputWithContext(ctx context.Context) WaitStepPropertiesResponseOutput {
	return o
}

func (o WaitStepPropertiesResponseOutput) ToWaitStepPropertiesResponsePtrOutput() WaitStepPropertiesResponsePtrOutput {
	return o.ToWaitStepPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o WaitStepPropertiesResponseOutput) ToWaitStepPropertiesResponsePtrOutputWithContext(ctx context.Context) WaitStepPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WaitStepPropertiesResponse) *WaitStepPropertiesResponse {
		return &v
	}).(WaitStepPropertiesResponsePtrOutput)
}

func (o WaitStepPropertiesResponseOutput) Attributes() WaitStepAttributesResponsePtrOutput {
	return o.ApplyT(func(v WaitStepPropertiesResponse) *WaitStepAttributesResponse { return v.Attributes }).(WaitStepAttributesResponsePtrOutput)
}

func (o WaitStepPropertiesResponseOutput) StepType() pulumi.StringOutput {
	return o.ApplyT(func(v WaitStepPropertiesResponse) string { return v.StepType }).(pulumi.StringOutput)
}

type WaitStepPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (WaitStepPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WaitStepPropertiesResponse)(nil)).Elem()
}

func (o WaitStepPropertiesResponsePtrOutput) ToWaitStepPropertiesResponsePtrOutput() WaitStepPropertiesResponsePtrOutput {
	return o
}

func (o WaitStepPropertiesResponsePtrOutput) ToWaitStepPropertiesResponsePtrOutputWithContext(ctx context.Context) WaitStepPropertiesResponsePtrOutput {
	return o
}

func (o WaitStepPropertiesResponsePtrOutput) Elem() WaitStepPropertiesResponseOutput {
	return o.ApplyT(func(v *WaitStepPropertiesResponse) WaitStepPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret WaitStepPropertiesResponse
		return ret
	}).(WaitStepPropertiesResponseOutput)
}

func (o WaitStepPropertiesResponsePtrOutput) Attributes() WaitStepAttributesResponsePtrOutput {
	return o.ApplyT(func(v *WaitStepPropertiesResponse) *WaitStepAttributesResponse {
		if v == nil {
			return nil
		}
		return v.Attributes
	}).(WaitStepAttributesResponsePtrOutput)
}

func (o WaitStepPropertiesResponsePtrOutput) StepType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WaitStepPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StepType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudErrorBodyResponseInput)(nil)).Elem(), CloudErrorBodyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudErrorBodyResponsePtrInput)(nil)).Elem(), CloudErrorBodyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudErrorBodyResponseArrayInput)(nil)).Elem(), CloudErrorBodyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPtrInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponseInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponsePtrInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MessageResponseInput)(nil)).Elem(), MessageResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MessageResponseArrayInput)(nil)).Elem(), MessageResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrePostStepInput)(nil)).Elem(), PrePostStepArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrePostStepArrayInput)(nil)).Elem(), PrePostStepArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrePostStepResponseInput)(nil)).Elem(), PrePostStepResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrePostStepResponseArrayInput)(nil)).Elem(), PrePostStepResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceOperationResponseInput)(nil)).Elem(), ResourceOperationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceOperationResponseArrayInput)(nil)).Elem(), ResourceOperationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RolloutOperationInfoResponseInput)(nil)).Elem(), RolloutOperationInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RolloutStepResponseInput)(nil)).Elem(), RolloutStepResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RolloutStepResponseArrayInput)(nil)).Elem(), RolloutStepResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SasAuthenticationInput)(nil)).Elem(), SasAuthenticationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SasAuthenticationPtrInput)(nil)).Elem(), SasAuthenticationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SasAuthenticationResponseInput)(nil)).Elem(), SasAuthenticationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SasAuthenticationResponsePtrInput)(nil)).Elem(), SasAuthenticationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceResponseInput)(nil)).Elem(), ServiceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceResponseArrayInput)(nil)).Elem(), ServiceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceUnitArtifactsInput)(nil)).Elem(), ServiceUnitArtifactsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceUnitArtifactsPtrInput)(nil)).Elem(), ServiceUnitArtifactsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceUnitArtifactsResponseInput)(nil)).Elem(), ServiceUnitArtifactsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceUnitArtifactsResponsePtrInput)(nil)).Elem(), ServiceUnitArtifactsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceUnitResponseInput)(nil)).Elem(), ServiceUnitResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceUnitResponseArrayInput)(nil)).Elem(), ServiceUnitResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StepTypeInput)(nil)).Elem(), StepTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StepTypeArrayInput)(nil)).Elem(), StepTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StepOperationInfoResponseInput)(nil)).Elem(), StepOperationInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StepResponseInput)(nil)).Elem(), StepResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StepResponseArrayInput)(nil)).Elem(), StepResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WaitStepAttributesInput)(nil)).Elem(), WaitStepAttributesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WaitStepAttributesPtrInput)(nil)).Elem(), WaitStepAttributesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WaitStepAttributesResponseInput)(nil)).Elem(), WaitStepAttributesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WaitStepAttributesResponsePtrInput)(nil)).Elem(), WaitStepAttributesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WaitStepPropertiesInput)(nil)).Elem(), WaitStepPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WaitStepPropertiesPtrInput)(nil)).Elem(), WaitStepPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WaitStepPropertiesResponseInput)(nil)).Elem(), WaitStepPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WaitStepPropertiesResponsePtrInput)(nil)).Elem(), WaitStepPropertiesResponseArgs{})
	pulumi.RegisterOutputType(CloudErrorBodyResponseOutput{})
	pulumi.RegisterOutputType(CloudErrorBodyResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudErrorBodyResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(MessageResponseOutput{})
	pulumi.RegisterOutputType(MessageResponseArrayOutput{})
	pulumi.RegisterOutputType(PrePostStepOutput{})
	pulumi.RegisterOutputType(PrePostStepArrayOutput{})
	pulumi.RegisterOutputType(PrePostStepResponseOutput{})
	pulumi.RegisterOutputType(PrePostStepResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceOperationResponseOutput{})
	pulumi.RegisterOutputType(ResourceOperationResponseArrayOutput{})
	pulumi.RegisterOutputType(RolloutOperationInfoResponseOutput{})
	pulumi.RegisterOutputType(RolloutStepResponseOutput{})
	pulumi.RegisterOutputType(RolloutStepResponseArrayOutput{})
	pulumi.RegisterOutputType(SasAuthenticationOutput{})
	pulumi.RegisterOutputType(SasAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(SasAuthenticationResponseOutput{})
	pulumi.RegisterOutputType(SasAuthenticationResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceResponseOutput{})
	pulumi.RegisterOutputType(ServiceResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsPtrOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsResponseOutput{})
	pulumi.RegisterOutputType(ServiceUnitArtifactsResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceUnitResponseOutput{})
	pulumi.RegisterOutputType(ServiceUnitResponseArrayOutput{})
	pulumi.RegisterOutputType(StepTypeOutput{})
	pulumi.RegisterOutputType(StepTypeArrayOutput{})
	pulumi.RegisterOutputType(StepOperationInfoResponseOutput{})
	pulumi.RegisterOutputType(StepResponseOutput{})
	pulumi.RegisterOutputType(StepResponseArrayOutput{})
	pulumi.RegisterOutputType(WaitStepAttributesOutput{})
	pulumi.RegisterOutputType(WaitStepAttributesPtrOutput{})
	pulumi.RegisterOutputType(WaitStepAttributesResponseOutput{})
	pulumi.RegisterOutputType(WaitStepAttributesResponsePtrOutput{})
	pulumi.RegisterOutputType(WaitStepPropertiesOutput{})
	pulumi.RegisterOutputType(WaitStepPropertiesPtrOutput{})
	pulumi.RegisterOutputType(WaitStepPropertiesResponseOutput{})
	pulumi.RegisterOutputType(WaitStepPropertiesResponsePtrOutput{})
}
