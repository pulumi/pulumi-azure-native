


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
