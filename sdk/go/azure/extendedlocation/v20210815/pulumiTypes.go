


package v20210815

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomLocationPropertiesAuthentication struct {
	Type  *string `pulumi:"type"`
	Value *string `pulumi:"value"`
}





type CustomLocationPropertiesAuthenticationInput interface {
	pulumi.Input

	ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput
	ToCustomLocationPropertiesAuthenticationOutputWithContext(context.Context) CustomLocationPropertiesAuthenticationOutput
}

type CustomLocationPropertiesAuthenticationArgs struct {
	Type  pulumi.StringPtrInput `pulumi:"type"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (CustomLocationPropertiesAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput {
	return i.ToCustomLocationPropertiesAuthenticationOutputWithContext(context.Background())
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationOutput)
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return i.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationOutput).ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx)
}









type CustomLocationPropertiesAuthenticationPtrInput interface {
	pulumi.Input

	ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput
	ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Context) CustomLocationPropertiesAuthenticationPtrOutput
}

type customLocationPropertiesAuthenticationPtrType CustomLocationPropertiesAuthenticationArgs

func CustomLocationPropertiesAuthenticationPtr(v *CustomLocationPropertiesAuthenticationArgs) CustomLocationPropertiesAuthenticationPtrInput {
	return (*customLocationPropertiesAuthenticationPtrType)(v)
}

func (*customLocationPropertiesAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (i *customLocationPropertiesAuthenticationPtrType) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return i.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (i *customLocationPropertiesAuthenticationPtrType) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationPtrOutput)
}

type CustomLocationPropertiesAuthenticationOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return o.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomLocationPropertiesAuthentication) *CustomLocationPropertiesAuthentication {
		return &v
	}).(CustomLocationPropertiesAuthenticationPtrOutput)
}

func (o CustomLocationPropertiesAuthenticationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesAuthentication) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o CustomLocationPropertiesAuthenticationOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesAuthentication) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type CustomLocationPropertiesAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) Elem() CustomLocationPropertiesAuthenticationOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) CustomLocationPropertiesAuthentication {
		if v != nil {
			return *v
		}
		var ret CustomLocationPropertiesAuthentication
		return ret
	}).(CustomLocationPropertiesAuthenticationOutput)
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type CustomLocationPropertiesResponseAuthentication struct {
	Type *string `pulumi:"type"`
}





type CustomLocationPropertiesResponseAuthenticationInput interface {
	pulumi.Input

	ToCustomLocationPropertiesResponseAuthenticationOutput() CustomLocationPropertiesResponseAuthenticationOutput
	ToCustomLocationPropertiesResponseAuthenticationOutputWithContext(context.Context) CustomLocationPropertiesResponseAuthenticationOutput
}

type CustomLocationPropertiesResponseAuthenticationArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (CustomLocationPropertiesResponseAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesResponseAuthentication)(nil)).Elem()
}

func (i CustomLocationPropertiesResponseAuthenticationArgs) ToCustomLocationPropertiesResponseAuthenticationOutput() CustomLocationPropertiesResponseAuthenticationOutput {
	return i.ToCustomLocationPropertiesResponseAuthenticationOutputWithContext(context.Background())
}

func (i CustomLocationPropertiesResponseAuthenticationArgs) ToCustomLocationPropertiesResponseAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesResponseAuthenticationOutput)
}

func (i CustomLocationPropertiesResponseAuthenticationArgs) ToCustomLocationPropertiesResponseAuthenticationPtrOutput() CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return i.ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(context.Background())
}

func (i CustomLocationPropertiesResponseAuthenticationArgs) ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesResponseAuthenticationOutput).ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(ctx)
}









type CustomLocationPropertiesResponseAuthenticationPtrInput interface {
	pulumi.Input

	ToCustomLocationPropertiesResponseAuthenticationPtrOutput() CustomLocationPropertiesResponseAuthenticationPtrOutput
	ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(context.Context) CustomLocationPropertiesResponseAuthenticationPtrOutput
}

type customLocationPropertiesResponseAuthenticationPtrType CustomLocationPropertiesResponseAuthenticationArgs

func CustomLocationPropertiesResponseAuthenticationPtr(v *CustomLocationPropertiesResponseAuthenticationArgs) CustomLocationPropertiesResponseAuthenticationPtrInput {
	return (*customLocationPropertiesResponseAuthenticationPtrType)(v)
}

func (*customLocationPropertiesResponseAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesResponseAuthentication)(nil)).Elem()
}

func (i *customLocationPropertiesResponseAuthenticationPtrType) ToCustomLocationPropertiesResponseAuthenticationPtrOutput() CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return i.ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(context.Background())
}

func (i *customLocationPropertiesResponseAuthenticationPtrType) ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesResponseAuthenticationPtrOutput)
}

type CustomLocationPropertiesResponseAuthenticationOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesResponseAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesResponseAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) ToCustomLocationPropertiesResponseAuthenticationOutput() CustomLocationPropertiesResponseAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) ToCustomLocationPropertiesResponseAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) ToCustomLocationPropertiesResponseAuthenticationPtrOutput() CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o.ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(context.Background())
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomLocationPropertiesResponseAuthentication) *CustomLocationPropertiesResponseAuthentication {
		return &v
	}).(CustomLocationPropertiesResponseAuthenticationPtrOutput)
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesResponseAuthentication) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type CustomLocationPropertiesResponseAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesResponseAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesResponseAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) ToCustomLocationPropertiesResponseAuthenticationPtrOutput() CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) Elem() CustomLocationPropertiesResponseAuthenticationOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesResponseAuthentication) CustomLocationPropertiesResponseAuthentication {
		if v != nil {
			return *v
		}
		var ret CustomLocationPropertiesResponseAuthentication
		return ret
	}).(CustomLocationPropertiesResponseAuthenticationOutput)
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesResponseAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type Identity struct {
	Type *string `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
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

func (o IdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.Type }).(pulumi.StringPtrOutput)
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

func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(CustomLocationPropertiesAuthenticationOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesResponseAuthenticationOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesResponseAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
