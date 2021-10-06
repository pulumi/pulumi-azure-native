


package v20210901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

type ErrorDetailResponse struct {
	AdditionalInfo []ErrorAdditionalInfoResponse `pulumi:"additionalInfo"`
	Code           string                        `pulumi:"code"`
	Details        []ErrorDetailResponse         `pulumi:"details"`
	Message        string                        `pulumi:"message"`
	Target         string                        `pulumi:"target"`
}





type ErrorDetailResponseInput interface {
	pulumi.Input

	ToErrorDetailResponseOutput() ErrorDetailResponseOutput
	ToErrorDetailResponseOutputWithContext(context.Context) ErrorDetailResponseOutput
}

type ErrorDetailResponseArgs struct {
	AdditionalInfo ErrorAdditionalInfoResponseArrayInput `pulumi:"additionalInfo"`
	Code           pulumi.StringInput                    `pulumi:"code"`
	Details        ErrorDetailResponseArrayInput         `pulumi:"details"`
	Message        pulumi.StringInput                    `pulumi:"message"`
	Target         pulumi.StringInput                    `pulumi:"target"`
}

func (ErrorDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return i.ToErrorDetailResponseOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput)
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return i.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput).ToErrorDetailResponsePtrOutputWithContext(ctx)
}









type ErrorDetailResponsePtrInput interface {
	pulumi.Input

	ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput
	ToErrorDetailResponsePtrOutputWithContext(context.Context) ErrorDetailResponsePtrOutput
}

type errorDetailResponsePtrType ErrorDetailResponseArgs

func ErrorDetailResponsePtr(v *ErrorDetailResponseArgs) ErrorDetailResponsePtrInput {
	return (*errorDetailResponsePtrType)(v)
}

func (*errorDetailResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDetailResponse)(nil)).Elem()
}

func (i *errorDetailResponsePtrType) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return i.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (i *errorDetailResponsePtrType) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponsePtrOutput)
}





type ErrorDetailResponseArrayInput interface {
	pulumi.Input

	ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput
	ToErrorDetailResponseArrayOutputWithContext(context.Context) ErrorDetailResponseArrayOutput
}

type ErrorDetailResponseArray []ErrorDetailResponseInput

func (ErrorDetailResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return i.ToErrorDetailResponseArrayOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseArrayOutput)
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return o.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ErrorDetailResponse) *ErrorDetailResponse {
		return &v
	}).(ErrorDetailResponsePtrOutput)
}

func (o ErrorDetailResponseOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorAdditionalInfoResponse { return v.AdditionalInfo }).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorDetailResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponsePtrOutput) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return o
}

func (o ErrorDetailResponsePtrOutput) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return o
}

func (o ErrorDetailResponsePtrOutput) Elem() ErrorDetailResponseOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) ErrorDetailResponse {
		if v != nil {
			return *v
		}
		var ret ErrorDetailResponse
		return ret
	}).(ErrorDetailResponseOutput)
}

func (o ErrorDetailResponsePtrOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) []ErrorAdditionalInfoResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalInfo
	}).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorDetailResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDetailResponsePtrOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) []ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDetailResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Target
	}).(pulumi.StringPtrOutput)
}

type ErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) ErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetailResponse {
		return vs[0].([]ErrorDetailResponse)[vs[1].(int)]
	}).(ErrorDetailResponseOutput)
}

type ExtensionAksAssignedIdentity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type ExtensionAksAssignedIdentityInput interface {
	pulumi.Input

	ToExtensionAksAssignedIdentityOutput() ExtensionAksAssignedIdentityOutput
	ToExtensionAksAssignedIdentityOutputWithContext(context.Context) ExtensionAksAssignedIdentityOutput
}

type ExtensionAksAssignedIdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (ExtensionAksAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionAksAssignedIdentity)(nil)).Elem()
}

func (i ExtensionAksAssignedIdentityArgs) ToExtensionAksAssignedIdentityOutput() ExtensionAksAssignedIdentityOutput {
	return i.ToExtensionAksAssignedIdentityOutputWithContext(context.Background())
}

func (i ExtensionAksAssignedIdentityArgs) ToExtensionAksAssignedIdentityOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionAksAssignedIdentityOutput)
}

func (i ExtensionAksAssignedIdentityArgs) ToExtensionAksAssignedIdentityPtrOutput() ExtensionAksAssignedIdentityPtrOutput {
	return i.ToExtensionAksAssignedIdentityPtrOutputWithContext(context.Background())
}

func (i ExtensionAksAssignedIdentityArgs) ToExtensionAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionAksAssignedIdentityOutput).ToExtensionAksAssignedIdentityPtrOutputWithContext(ctx)
}









type ExtensionAksAssignedIdentityPtrInput interface {
	pulumi.Input

	ToExtensionAksAssignedIdentityPtrOutput() ExtensionAksAssignedIdentityPtrOutput
	ToExtensionAksAssignedIdentityPtrOutputWithContext(context.Context) ExtensionAksAssignedIdentityPtrOutput
}

type extensionAksAssignedIdentityPtrType ExtensionAksAssignedIdentityArgs

func ExtensionAksAssignedIdentityPtr(v *ExtensionAksAssignedIdentityArgs) ExtensionAksAssignedIdentityPtrInput {
	return (*extensionAksAssignedIdentityPtrType)(v)
}

func (*extensionAksAssignedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionAksAssignedIdentity)(nil)).Elem()
}

func (i *extensionAksAssignedIdentityPtrType) ToExtensionAksAssignedIdentityPtrOutput() ExtensionAksAssignedIdentityPtrOutput {
	return i.ToExtensionAksAssignedIdentityPtrOutputWithContext(context.Background())
}

func (i *extensionAksAssignedIdentityPtrType) ToExtensionAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionAksAssignedIdentityPtrOutput)
}

type ExtensionAksAssignedIdentityOutput struct{ *pulumi.OutputState }

func (ExtensionAksAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionAksAssignedIdentity)(nil)).Elem()
}

func (o ExtensionAksAssignedIdentityOutput) ToExtensionAksAssignedIdentityOutput() ExtensionAksAssignedIdentityOutput {
	return o
}

func (o ExtensionAksAssignedIdentityOutput) ToExtensionAksAssignedIdentityOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityOutput {
	return o
}

func (o ExtensionAksAssignedIdentityOutput) ToExtensionAksAssignedIdentityPtrOutput() ExtensionAksAssignedIdentityPtrOutput {
	return o.ToExtensionAksAssignedIdentityPtrOutputWithContext(context.Background())
}

func (o ExtensionAksAssignedIdentityOutput) ToExtensionAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtensionAksAssignedIdentity) *ExtensionAksAssignedIdentity {
		return &v
	}).(ExtensionAksAssignedIdentityPtrOutput)
}

func (o ExtensionAksAssignedIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ExtensionAksAssignedIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type ExtensionAksAssignedIdentityPtrOutput struct{ *pulumi.OutputState }

func (ExtensionAksAssignedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionAksAssignedIdentity)(nil)).Elem()
}

func (o ExtensionAksAssignedIdentityPtrOutput) ToExtensionAksAssignedIdentityPtrOutput() ExtensionAksAssignedIdentityPtrOutput {
	return o
}

func (o ExtensionAksAssignedIdentityPtrOutput) ToExtensionAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityPtrOutput {
	return o
}

func (o ExtensionAksAssignedIdentityPtrOutput) Elem() ExtensionAksAssignedIdentityOutput {
	return o.ApplyT(func(v *ExtensionAksAssignedIdentity) ExtensionAksAssignedIdentity {
		if v != nil {
			return *v
		}
		var ret ExtensionAksAssignedIdentity
		return ret
	}).(ExtensionAksAssignedIdentityOutput)
}

func (o ExtensionAksAssignedIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ExtensionAksAssignedIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type ExtensionResponseAksAssignedIdentity struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type ExtensionResponseAksAssignedIdentityInput interface {
	pulumi.Input

	ToExtensionResponseAksAssignedIdentityOutput() ExtensionResponseAksAssignedIdentityOutput
	ToExtensionResponseAksAssignedIdentityOutputWithContext(context.Context) ExtensionResponseAksAssignedIdentityOutput
}

type ExtensionResponseAksAssignedIdentityArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtensionResponseAksAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionResponseAksAssignedIdentity)(nil)).Elem()
}

func (i ExtensionResponseAksAssignedIdentityArgs) ToExtensionResponseAksAssignedIdentityOutput() ExtensionResponseAksAssignedIdentityOutput {
	return i.ToExtensionResponseAksAssignedIdentityOutputWithContext(context.Background())
}

func (i ExtensionResponseAksAssignedIdentityArgs) ToExtensionResponseAksAssignedIdentityOutputWithContext(ctx context.Context) ExtensionResponseAksAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionResponseAksAssignedIdentityOutput)
}

func (i ExtensionResponseAksAssignedIdentityArgs) ToExtensionResponseAksAssignedIdentityPtrOutput() ExtensionResponseAksAssignedIdentityPtrOutput {
	return i.ToExtensionResponseAksAssignedIdentityPtrOutputWithContext(context.Background())
}

func (i ExtensionResponseAksAssignedIdentityArgs) ToExtensionResponseAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionResponseAksAssignedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionResponseAksAssignedIdentityOutput).ToExtensionResponseAksAssignedIdentityPtrOutputWithContext(ctx)
}









type ExtensionResponseAksAssignedIdentityPtrInput interface {
	pulumi.Input

	ToExtensionResponseAksAssignedIdentityPtrOutput() ExtensionResponseAksAssignedIdentityPtrOutput
	ToExtensionResponseAksAssignedIdentityPtrOutputWithContext(context.Context) ExtensionResponseAksAssignedIdentityPtrOutput
}

type extensionResponseAksAssignedIdentityPtrType ExtensionResponseAksAssignedIdentityArgs

func ExtensionResponseAksAssignedIdentityPtr(v *ExtensionResponseAksAssignedIdentityArgs) ExtensionResponseAksAssignedIdentityPtrInput {
	return (*extensionResponseAksAssignedIdentityPtrType)(v)
}

func (*extensionResponseAksAssignedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionResponseAksAssignedIdentity)(nil)).Elem()
}

func (i *extensionResponseAksAssignedIdentityPtrType) ToExtensionResponseAksAssignedIdentityPtrOutput() ExtensionResponseAksAssignedIdentityPtrOutput {
	return i.ToExtensionResponseAksAssignedIdentityPtrOutputWithContext(context.Background())
}

func (i *extensionResponseAksAssignedIdentityPtrType) ToExtensionResponseAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionResponseAksAssignedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionResponseAksAssignedIdentityPtrOutput)
}

type ExtensionResponseAksAssignedIdentityOutput struct{ *pulumi.OutputState }

func (ExtensionResponseAksAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionResponseAksAssignedIdentity)(nil)).Elem()
}

func (o ExtensionResponseAksAssignedIdentityOutput) ToExtensionResponseAksAssignedIdentityOutput() ExtensionResponseAksAssignedIdentityOutput {
	return o
}

func (o ExtensionResponseAksAssignedIdentityOutput) ToExtensionResponseAksAssignedIdentityOutputWithContext(ctx context.Context) ExtensionResponseAksAssignedIdentityOutput {
	return o
}

func (o ExtensionResponseAksAssignedIdentityOutput) ToExtensionResponseAksAssignedIdentityPtrOutput() ExtensionResponseAksAssignedIdentityPtrOutput {
	return o.ToExtensionResponseAksAssignedIdentityPtrOutputWithContext(context.Background())
}

func (o ExtensionResponseAksAssignedIdentityOutput) ToExtensionResponseAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionResponseAksAssignedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtensionResponseAksAssignedIdentity) *ExtensionResponseAksAssignedIdentity {
		return &v
	}).(ExtensionResponseAksAssignedIdentityPtrOutput)
}

func (o ExtensionResponseAksAssignedIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ExtensionResponseAksAssignedIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ExtensionResponseAksAssignedIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ExtensionResponseAksAssignedIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ExtensionResponseAksAssignedIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResponseAksAssignedIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtensionResponseAksAssignedIdentityPtrOutput struct{ *pulumi.OutputState }

func (ExtensionResponseAksAssignedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionResponseAksAssignedIdentity)(nil)).Elem()
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) ToExtensionResponseAksAssignedIdentityPtrOutput() ExtensionResponseAksAssignedIdentityPtrOutput {
	return o
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) ToExtensionResponseAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionResponseAksAssignedIdentityPtrOutput {
	return o
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) Elem() ExtensionResponseAksAssignedIdentityOutput {
	return o.ApplyT(func(v *ExtensionResponseAksAssignedIdentity) ExtensionResponseAksAssignedIdentity {
		if v != nil {
			return *v
		}
		var ret ExtensionResponseAksAssignedIdentity
		return ret
	}).(ExtensionResponseAksAssignedIdentityOutput)
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResponseAksAssignedIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResponseAksAssignedIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResponseAksAssignedIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtensionStatus struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
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





type ExtensionStatusResponseInput interface {
	pulumi.Input

	ToExtensionStatusResponseOutput() ExtensionStatusResponseOutput
	ToExtensionStatusResponseOutputWithContext(context.Context) ExtensionStatusResponseOutput
}

type ExtensionStatusResponseArgs struct {
	Code          pulumi.StringPtrInput `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput `pulumi:"displayStatus"`
	Level         pulumi.StringPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput `pulumi:"message"`
	Time          pulumi.StringPtrInput `pulumi:"time"`
}

func (ExtensionStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionStatusResponse)(nil)).Elem()
}

func (i ExtensionStatusResponseArgs) ToExtensionStatusResponseOutput() ExtensionStatusResponseOutput {
	return i.ToExtensionStatusResponseOutputWithContext(context.Background())
}

func (i ExtensionStatusResponseArgs) ToExtensionStatusResponseOutputWithContext(ctx context.Context) ExtensionStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionStatusResponseOutput)
}





type ExtensionStatusResponseArrayInput interface {
	pulumi.Input

	ToExtensionStatusResponseArrayOutput() ExtensionStatusResponseArrayOutput
	ToExtensionStatusResponseArrayOutputWithContext(context.Context) ExtensionStatusResponseArrayOutput
}

type ExtensionStatusResponseArray []ExtensionStatusResponseInput

func (ExtensionStatusResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionStatusResponse)(nil)).Elem()
}

func (i ExtensionStatusResponseArray) ToExtensionStatusResponseArrayOutput() ExtensionStatusResponseArrayOutput {
	return i.ToExtensionStatusResponseArrayOutputWithContext(context.Background())
}

func (i ExtensionStatusResponseArray) ToExtensionStatusResponseArrayOutputWithContext(ctx context.Context) ExtensionStatusResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionStatusResponseArrayOutput)
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





type ScopeClusterResponseInput interface {
	pulumi.Input

	ToScopeClusterResponseOutput() ScopeClusterResponseOutput
	ToScopeClusterResponseOutputWithContext(context.Context) ScopeClusterResponseOutput
}

type ScopeClusterResponseArgs struct {
	ReleaseNamespace pulumi.StringPtrInput `pulumi:"releaseNamespace"`
}

func (ScopeClusterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeClusterResponse)(nil)).Elem()
}

func (i ScopeClusterResponseArgs) ToScopeClusterResponseOutput() ScopeClusterResponseOutput {
	return i.ToScopeClusterResponseOutputWithContext(context.Background())
}

func (i ScopeClusterResponseArgs) ToScopeClusterResponseOutputWithContext(ctx context.Context) ScopeClusterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterResponseOutput)
}

func (i ScopeClusterResponseArgs) ToScopeClusterResponsePtrOutput() ScopeClusterResponsePtrOutput {
	return i.ToScopeClusterResponsePtrOutputWithContext(context.Background())
}

func (i ScopeClusterResponseArgs) ToScopeClusterResponsePtrOutputWithContext(ctx context.Context) ScopeClusterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterResponseOutput).ToScopeClusterResponsePtrOutputWithContext(ctx)
}









type ScopeClusterResponsePtrInput interface {
	pulumi.Input

	ToScopeClusterResponsePtrOutput() ScopeClusterResponsePtrOutput
	ToScopeClusterResponsePtrOutputWithContext(context.Context) ScopeClusterResponsePtrOutput
}

type scopeClusterResponsePtrType ScopeClusterResponseArgs

func ScopeClusterResponsePtr(v *ScopeClusterResponseArgs) ScopeClusterResponsePtrInput {
	return (*scopeClusterResponsePtrType)(v)
}

func (*scopeClusterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeClusterResponse)(nil)).Elem()
}

func (i *scopeClusterResponsePtrType) ToScopeClusterResponsePtrOutput() ScopeClusterResponsePtrOutput {
	return i.ToScopeClusterResponsePtrOutputWithContext(context.Background())
}

func (i *scopeClusterResponsePtrType) ToScopeClusterResponsePtrOutputWithContext(ctx context.Context) ScopeClusterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterResponsePtrOutput)
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

func (o ScopeClusterResponseOutput) ToScopeClusterResponsePtrOutput() ScopeClusterResponsePtrOutput {
	return o.ToScopeClusterResponsePtrOutputWithContext(context.Background())
}

func (o ScopeClusterResponseOutput) ToScopeClusterResponsePtrOutputWithContext(ctx context.Context) ScopeClusterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeClusterResponse) *ScopeClusterResponse {
		return &v
	}).(ScopeClusterResponsePtrOutput)
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





type ScopeNamespaceResponseInput interface {
	pulumi.Input

	ToScopeNamespaceResponseOutput() ScopeNamespaceResponseOutput
	ToScopeNamespaceResponseOutputWithContext(context.Context) ScopeNamespaceResponseOutput
}

type ScopeNamespaceResponseArgs struct {
	TargetNamespace pulumi.StringPtrInput `pulumi:"targetNamespace"`
}

func (ScopeNamespaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeNamespaceResponse)(nil)).Elem()
}

func (i ScopeNamespaceResponseArgs) ToScopeNamespaceResponseOutput() ScopeNamespaceResponseOutput {
	return i.ToScopeNamespaceResponseOutputWithContext(context.Background())
}

func (i ScopeNamespaceResponseArgs) ToScopeNamespaceResponseOutputWithContext(ctx context.Context) ScopeNamespaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceResponseOutput)
}

func (i ScopeNamespaceResponseArgs) ToScopeNamespaceResponsePtrOutput() ScopeNamespaceResponsePtrOutput {
	return i.ToScopeNamespaceResponsePtrOutputWithContext(context.Background())
}

func (i ScopeNamespaceResponseArgs) ToScopeNamespaceResponsePtrOutputWithContext(ctx context.Context) ScopeNamespaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceResponseOutput).ToScopeNamespaceResponsePtrOutputWithContext(ctx)
}









type ScopeNamespaceResponsePtrInput interface {
	pulumi.Input

	ToScopeNamespaceResponsePtrOutput() ScopeNamespaceResponsePtrOutput
	ToScopeNamespaceResponsePtrOutputWithContext(context.Context) ScopeNamespaceResponsePtrOutput
}

type scopeNamespaceResponsePtrType ScopeNamespaceResponseArgs

func ScopeNamespaceResponsePtr(v *ScopeNamespaceResponseArgs) ScopeNamespaceResponsePtrInput {
	return (*scopeNamespaceResponsePtrType)(v)
}

func (*scopeNamespaceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeNamespaceResponse)(nil)).Elem()
}

func (i *scopeNamespaceResponsePtrType) ToScopeNamespaceResponsePtrOutput() ScopeNamespaceResponsePtrOutput {
	return i.ToScopeNamespaceResponsePtrOutputWithContext(context.Background())
}

func (i *scopeNamespaceResponsePtrType) ToScopeNamespaceResponsePtrOutputWithContext(ctx context.Context) ScopeNamespaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceResponsePtrOutput)
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

func (o ScopeNamespaceResponseOutput) ToScopeNamespaceResponsePtrOutput() ScopeNamespaceResponsePtrOutput {
	return o.ToScopeNamespaceResponsePtrOutputWithContext(context.Background())
}

func (o ScopeNamespaceResponseOutput) ToScopeNamespaceResponsePtrOutputWithContext(ctx context.Context) ScopeNamespaceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeNamespaceResponse) *ScopeNamespaceResponse {
		return &v
	}).(ScopeNamespaceResponsePtrOutput)
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





type ScopeResponseInput interface {
	pulumi.Input

	ToScopeResponseOutput() ScopeResponseOutput
	ToScopeResponseOutputWithContext(context.Context) ScopeResponseOutput
}

type ScopeResponseArgs struct {
	Cluster   ScopeClusterResponsePtrInput   `pulumi:"cluster"`
	Namespace ScopeNamespaceResponsePtrInput `pulumi:"namespace"`
}

func (ScopeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeResponse)(nil)).Elem()
}

func (i ScopeResponseArgs) ToScopeResponseOutput() ScopeResponseOutput {
	return i.ToScopeResponseOutputWithContext(context.Background())
}

func (i ScopeResponseArgs) ToScopeResponseOutputWithContext(ctx context.Context) ScopeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeResponseOutput)
}

func (i ScopeResponseArgs) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return i.ToScopeResponsePtrOutputWithContext(context.Background())
}

func (i ScopeResponseArgs) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeResponseOutput).ToScopeResponsePtrOutputWithContext(ctx)
}









type ScopeResponsePtrInput interface {
	pulumi.Input

	ToScopeResponsePtrOutput() ScopeResponsePtrOutput
	ToScopeResponsePtrOutputWithContext(context.Context) ScopeResponsePtrOutput
}

type scopeResponsePtrType ScopeResponseArgs

func ScopeResponsePtr(v *ScopeResponseArgs) ScopeResponsePtrInput {
	return (*scopeResponsePtrType)(v)
}

func (*scopeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeResponse)(nil)).Elem()
}

func (i *scopeResponsePtrType) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return i.ToScopeResponsePtrOutputWithContext(context.Background())
}

func (i *scopeResponsePtrType) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeResponsePtrOutput)
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

func (o ScopeResponseOutput) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return o.ToScopeResponsePtrOutputWithContext(context.Background())
}

func (o ScopeResponseOutput) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeResponse) *ScopeResponse {
		return &v
	}).(ScopeResponsePtrOutput)
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





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ExtensionAksAssignedIdentityOutput{})
	pulumi.RegisterOutputType(ExtensionAksAssignedIdentityPtrOutput{})
	pulumi.RegisterOutputType(ExtensionResponseAksAssignedIdentityOutput{})
	pulumi.RegisterOutputType(ExtensionResponseAksAssignedIdentityPtrOutput{})
	pulumi.RegisterOutputType(ExtensionStatusOutput{})
	pulumi.RegisterOutputType(ExtensionStatusArrayOutput{})
	pulumi.RegisterOutputType(ExtensionStatusResponseOutput{})
	pulumi.RegisterOutputType(ExtensionStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
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
