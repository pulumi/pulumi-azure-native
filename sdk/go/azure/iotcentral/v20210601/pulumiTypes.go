


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppSkuInfo struct {
	Name string `pulumi:"name"`
}





type AppSkuInfoInput interface {
	pulumi.Input

	ToAppSkuInfoOutput() AppSkuInfoOutput
	ToAppSkuInfoOutputWithContext(context.Context) AppSkuInfoOutput
}

type AppSkuInfoArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (AppSkuInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfo)(nil)).Elem()
}

func (i AppSkuInfoArgs) ToAppSkuInfoOutput() AppSkuInfoOutput {
	return i.ToAppSkuInfoOutputWithContext(context.Background())
}

func (i AppSkuInfoArgs) ToAppSkuInfoOutputWithContext(ctx context.Context) AppSkuInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoOutput)
}

func (i AppSkuInfoArgs) ToAppSkuInfoPtrOutput() AppSkuInfoPtrOutput {
	return i.ToAppSkuInfoPtrOutputWithContext(context.Background())
}

func (i AppSkuInfoArgs) ToAppSkuInfoPtrOutputWithContext(ctx context.Context) AppSkuInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoOutput).ToAppSkuInfoPtrOutputWithContext(ctx)
}









type AppSkuInfoPtrInput interface {
	pulumi.Input

	ToAppSkuInfoPtrOutput() AppSkuInfoPtrOutput
	ToAppSkuInfoPtrOutputWithContext(context.Context) AppSkuInfoPtrOutput
}

type appSkuInfoPtrType AppSkuInfoArgs

func AppSkuInfoPtr(v *AppSkuInfoArgs) AppSkuInfoPtrInput {
	return (*appSkuInfoPtrType)(v)
}

func (*appSkuInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSkuInfo)(nil)).Elem()
}

func (i *appSkuInfoPtrType) ToAppSkuInfoPtrOutput() AppSkuInfoPtrOutput {
	return i.ToAppSkuInfoPtrOutputWithContext(context.Background())
}

func (i *appSkuInfoPtrType) ToAppSkuInfoPtrOutputWithContext(ctx context.Context) AppSkuInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoPtrOutput)
}

type AppSkuInfoOutput struct{ *pulumi.OutputState }

func (AppSkuInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfo)(nil)).Elem()
}

func (o AppSkuInfoOutput) ToAppSkuInfoOutput() AppSkuInfoOutput {
	return o
}

func (o AppSkuInfoOutput) ToAppSkuInfoOutputWithContext(ctx context.Context) AppSkuInfoOutput {
	return o
}

func (o AppSkuInfoOutput) ToAppSkuInfoPtrOutput() AppSkuInfoPtrOutput {
	return o.ToAppSkuInfoPtrOutputWithContext(context.Background())
}

func (o AppSkuInfoOutput) ToAppSkuInfoPtrOutputWithContext(ctx context.Context) AppSkuInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSkuInfo) *AppSkuInfo {
		return &v
	}).(AppSkuInfoPtrOutput)
}

func (o AppSkuInfoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AppSkuInfo) string { return v.Name }).(pulumi.StringOutput)
}

type AppSkuInfoPtrOutput struct{ *pulumi.OutputState }

func (AppSkuInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSkuInfo)(nil)).Elem()
}

func (o AppSkuInfoPtrOutput) ToAppSkuInfoPtrOutput() AppSkuInfoPtrOutput {
	return o
}

func (o AppSkuInfoPtrOutput) ToAppSkuInfoPtrOutputWithContext(ctx context.Context) AppSkuInfoPtrOutput {
	return o
}

func (o AppSkuInfoPtrOutput) Elem() AppSkuInfoOutput {
	return o.ApplyT(func(v *AppSkuInfo) AppSkuInfo {
		if v != nil {
			return *v
		}
		var ret AppSkuInfo
		return ret
	}).(AppSkuInfoOutput)
}

func (o AppSkuInfoPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppSkuInfo) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type AppSkuInfoResponse struct {
	Name string `pulumi:"name"`
}





type AppSkuInfoResponseInput interface {
	pulumi.Input

	ToAppSkuInfoResponseOutput() AppSkuInfoResponseOutput
	ToAppSkuInfoResponseOutputWithContext(context.Context) AppSkuInfoResponseOutput
}

type AppSkuInfoResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (AppSkuInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfoResponse)(nil)).Elem()
}

func (i AppSkuInfoResponseArgs) ToAppSkuInfoResponseOutput() AppSkuInfoResponseOutput {
	return i.ToAppSkuInfoResponseOutputWithContext(context.Background())
}

func (i AppSkuInfoResponseArgs) ToAppSkuInfoResponseOutputWithContext(ctx context.Context) AppSkuInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoResponseOutput)
}

func (i AppSkuInfoResponseArgs) ToAppSkuInfoResponsePtrOutput() AppSkuInfoResponsePtrOutput {
	return i.ToAppSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (i AppSkuInfoResponseArgs) ToAppSkuInfoResponsePtrOutputWithContext(ctx context.Context) AppSkuInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoResponseOutput).ToAppSkuInfoResponsePtrOutputWithContext(ctx)
}









type AppSkuInfoResponsePtrInput interface {
	pulumi.Input

	ToAppSkuInfoResponsePtrOutput() AppSkuInfoResponsePtrOutput
	ToAppSkuInfoResponsePtrOutputWithContext(context.Context) AppSkuInfoResponsePtrOutput
}

type appSkuInfoResponsePtrType AppSkuInfoResponseArgs

func AppSkuInfoResponsePtr(v *AppSkuInfoResponseArgs) AppSkuInfoResponsePtrInput {
	return (*appSkuInfoResponsePtrType)(v)
}

func (*appSkuInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSkuInfoResponse)(nil)).Elem()
}

func (i *appSkuInfoResponsePtrType) ToAppSkuInfoResponsePtrOutput() AppSkuInfoResponsePtrOutput {
	return i.ToAppSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (i *appSkuInfoResponsePtrType) ToAppSkuInfoResponsePtrOutputWithContext(ctx context.Context) AppSkuInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoResponsePtrOutput)
}

type AppSkuInfoResponseOutput struct{ *pulumi.OutputState }

func (AppSkuInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfoResponse)(nil)).Elem()
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponseOutput() AppSkuInfoResponseOutput {
	return o
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponseOutputWithContext(ctx context.Context) AppSkuInfoResponseOutput {
	return o
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponsePtrOutput() AppSkuInfoResponsePtrOutput {
	return o.ToAppSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponsePtrOutputWithContext(ctx context.Context) AppSkuInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSkuInfoResponse) *AppSkuInfoResponse {
		return &v
	}).(AppSkuInfoResponsePtrOutput)
}

func (o AppSkuInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AppSkuInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

type AppSkuInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AppSkuInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSkuInfoResponse)(nil)).Elem()
}

func (o AppSkuInfoResponsePtrOutput) ToAppSkuInfoResponsePtrOutput() AppSkuInfoResponsePtrOutput {
	return o
}

func (o AppSkuInfoResponsePtrOutput) ToAppSkuInfoResponsePtrOutputWithContext(ctx context.Context) AppSkuInfoResponsePtrOutput {
	return o
}

func (o AppSkuInfoResponsePtrOutput) Elem() AppSkuInfoResponseOutput {
	return o.ApplyT(func(v *AppSkuInfoResponse) AppSkuInfoResponse {
		if v != nil {
			return *v
		}
		var ret AppSkuInfoResponse
		return ret
	}).(AppSkuInfoResponseOutput)
}

func (o AppSkuInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppSkuInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SystemAssignedServiceIdentity struct {
	Type string `pulumi:"type"`
}





type SystemAssignedServiceIdentityInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput
	ToSystemAssignedServiceIdentityOutputWithContext(context.Context) SystemAssignedServiceIdentityOutput
}

type SystemAssignedServiceIdentityArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (SystemAssignedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentity)(nil)).Elem()
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput {
	return i.ToSystemAssignedServiceIdentityOutputWithContext(context.Background())
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityOutput)
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return i.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityOutput).ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx)
}









type SystemAssignedServiceIdentityPtrInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput
	ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Context) SystemAssignedServiceIdentityPtrOutput
}

type systemAssignedServiceIdentityPtrType SystemAssignedServiceIdentityArgs

func SystemAssignedServiceIdentityPtr(v *SystemAssignedServiceIdentityArgs) SystemAssignedServiceIdentityPtrInput {
	return (*systemAssignedServiceIdentityPtrType)(v)
}

func (*systemAssignedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentity)(nil)).Elem()
}

func (i *systemAssignedServiceIdentityPtrType) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return i.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *systemAssignedServiceIdentityPtrType) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityPtrOutput)
}

type SystemAssignedServiceIdentityOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentity)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput {
	return o
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityOutput {
	return o
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return o.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemAssignedServiceIdentity) *SystemAssignedServiceIdentity {
		return &v
	}).(SystemAssignedServiceIdentityPtrOutput)
}

func (o SystemAssignedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type SystemAssignedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentity)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityPtrOutput) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityPtrOutput) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityPtrOutput) Elem() SystemAssignedServiceIdentityOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentity) SystemAssignedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret SystemAssignedServiceIdentity
		return ret
	}).(SystemAssignedServiceIdentityOutput)
}

func (o SystemAssignedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type SystemAssignedServiceIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}





type SystemAssignedServiceIdentityResponseInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityResponseOutput() SystemAssignedServiceIdentityResponseOutput
	ToSystemAssignedServiceIdentityResponseOutputWithContext(context.Context) SystemAssignedServiceIdentityResponseOutput
}

type SystemAssignedServiceIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (SystemAssignedServiceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentityResponse)(nil)).Elem()
}

func (i SystemAssignedServiceIdentityResponseArgs) ToSystemAssignedServiceIdentityResponseOutput() SystemAssignedServiceIdentityResponseOutput {
	return i.ToSystemAssignedServiceIdentityResponseOutputWithContext(context.Background())
}

func (i SystemAssignedServiceIdentityResponseArgs) ToSystemAssignedServiceIdentityResponseOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityResponseOutput)
}

func (i SystemAssignedServiceIdentityResponseArgs) ToSystemAssignedServiceIdentityResponsePtrOutput() SystemAssignedServiceIdentityResponsePtrOutput {
	return i.ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i SystemAssignedServiceIdentityResponseArgs) ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityResponseOutput).ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(ctx)
}









type SystemAssignedServiceIdentityResponsePtrInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityResponsePtrOutput() SystemAssignedServiceIdentityResponsePtrOutput
	ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(context.Context) SystemAssignedServiceIdentityResponsePtrOutput
}

type systemAssignedServiceIdentityResponsePtrType SystemAssignedServiceIdentityResponseArgs

func SystemAssignedServiceIdentityResponsePtr(v *SystemAssignedServiceIdentityResponseArgs) SystemAssignedServiceIdentityResponsePtrInput {
	return (*systemAssignedServiceIdentityResponsePtrType)(v)
}

func (*systemAssignedServiceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentityResponse)(nil)).Elem()
}

func (i *systemAssignedServiceIdentityResponsePtrType) ToSystemAssignedServiceIdentityResponsePtrOutput() SystemAssignedServiceIdentityResponsePtrOutput {
	return i.ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *systemAssignedServiceIdentityResponsePtrType) ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityResponsePtrOutput)
}

type SystemAssignedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentityResponse)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityResponseOutput) ToSystemAssignedServiceIdentityResponseOutput() SystemAssignedServiceIdentityResponseOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponseOutput) ToSystemAssignedServiceIdentityResponseOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponseOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponseOutput) ToSystemAssignedServiceIdentityResponsePtrOutput() SystemAssignedServiceIdentityResponsePtrOutput {
	return o.ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityResponseOutput) ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemAssignedServiceIdentityResponse) *SystemAssignedServiceIdentityResponse {
		return &v
	}).(SystemAssignedServiceIdentityResponsePtrOutput)
}

func (o SystemAssignedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o SystemAssignedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o SystemAssignedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SystemAssignedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentityResponse)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) ToSystemAssignedServiceIdentityResponsePtrOutput() SystemAssignedServiceIdentityResponsePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponsePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) Elem() SystemAssignedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) SystemAssignedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret SystemAssignedServiceIdentityResponse
		return ret
	}).(SystemAssignedServiceIdentityResponseOutput)
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSkuInfoOutput{})
	pulumi.RegisterOutputType(AppSkuInfoPtrOutput{})
	pulumi.RegisterOutputType(AppSkuInfoResponseOutput{})
	pulumi.RegisterOutputType(AppSkuInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityResponsePtrOutput{})
}
