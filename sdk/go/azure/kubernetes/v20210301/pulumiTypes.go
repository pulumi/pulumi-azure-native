


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectedClusterIdentity struct {
	Type ResourceIdentityType `pulumi:"type"`
}


func (val *ConnectedClusterIdentity) Defaults() *ConnectedClusterIdentity {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		tmp.Type = ResourceIdentityType("SystemAssigned")
	}
	return &tmp
}





type ConnectedClusterIdentityInput interface {
	pulumi.Input

	ToConnectedClusterIdentityOutput() ConnectedClusterIdentityOutput
	ToConnectedClusterIdentityOutputWithContext(context.Context) ConnectedClusterIdentityOutput
}

type ConnectedClusterIdentityArgs struct {
	Type ResourceIdentityTypeInput `pulumi:"type"`
}

func (ConnectedClusterIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentity)(nil)).Elem()
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityOutput() ConnectedClusterIdentityOutput {
	return i.ToConnectedClusterIdentityOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityOutputWithContext(ctx context.Context) ConnectedClusterIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityOutput)
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return i.ToConnectedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityArgs) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityOutput).ToConnectedClusterIdentityPtrOutputWithContext(ctx)
}









type ConnectedClusterIdentityPtrInput interface {
	pulumi.Input

	ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput
	ToConnectedClusterIdentityPtrOutputWithContext(context.Context) ConnectedClusterIdentityPtrOutput
}

type connectedClusterIdentityPtrType ConnectedClusterIdentityArgs

func ConnectedClusterIdentityPtr(v *ConnectedClusterIdentityArgs) ConnectedClusterIdentityPtrInput {
	return (*connectedClusterIdentityPtrType)(v)
}

func (*connectedClusterIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentity)(nil)).Elem()
}

func (i *connectedClusterIdentityPtrType) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return i.ToConnectedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i *connectedClusterIdentityPtrType) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityPtrOutput)
}

type ConnectedClusterIdentityOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentity)(nil)).Elem()
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityOutput() ConnectedClusterIdentityOutput {
	return o
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityOutputWithContext(ctx context.Context) ConnectedClusterIdentityOutput {
	return o
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return o.ToConnectedClusterIdentityPtrOutputWithContext(context.Background())
}

func (o ConnectedClusterIdentityOutput) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectedClusterIdentity) *ConnectedClusterIdentity {
		return &v
	}).(ConnectedClusterIdentityPtrOutput)
}

func (o ConnectedClusterIdentityOutput) Type() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v ConnectedClusterIdentity) ResourceIdentityType { return v.Type }).(ResourceIdentityTypeOutput)
}

type ConnectedClusterIdentityPtrOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentity)(nil)).Elem()
}

func (o ConnectedClusterIdentityPtrOutput) ToConnectedClusterIdentityPtrOutput() ConnectedClusterIdentityPtrOutput {
	return o
}

func (o ConnectedClusterIdentityPtrOutput) ToConnectedClusterIdentityPtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityPtrOutput {
	return o
}

func (o ConnectedClusterIdentityPtrOutput) Elem() ConnectedClusterIdentityOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentity) ConnectedClusterIdentity {
		if v != nil {
			return *v
		}
		var ret ConnectedClusterIdentity
		return ret
	}).(ConnectedClusterIdentityOutput)
}

func (o ConnectedClusterIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type ConnectedClusterIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}


func (val *ConnectedClusterIdentityResponse) Defaults() *ConnectedClusterIdentityResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		tmp.Type = "SystemAssigned"
	}
	return &tmp
}





type ConnectedClusterIdentityResponseInput interface {
	pulumi.Input

	ToConnectedClusterIdentityResponseOutput() ConnectedClusterIdentityResponseOutput
	ToConnectedClusterIdentityResponseOutputWithContext(context.Context) ConnectedClusterIdentityResponseOutput
}

type ConnectedClusterIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (ConnectedClusterIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponseOutput() ConnectedClusterIdentityResponseOutput {
	return i.ToConnectedClusterIdentityResponseOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponseOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityResponseOutput)
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return i.ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ConnectedClusterIdentityResponseArgs) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityResponseOutput).ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx)
}









type ConnectedClusterIdentityResponsePtrInput interface {
	pulumi.Input

	ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput
	ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Context) ConnectedClusterIdentityResponsePtrOutput
}

type connectedClusterIdentityResponsePtrType ConnectedClusterIdentityResponseArgs

func ConnectedClusterIdentityResponsePtr(v *ConnectedClusterIdentityResponseArgs) ConnectedClusterIdentityResponsePtrInput {
	return (*connectedClusterIdentityResponsePtrType)(v)
}

func (*connectedClusterIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (i *connectedClusterIdentityResponsePtrType) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return i.ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *connectedClusterIdentityResponsePtrType) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterIdentityResponsePtrOutput)
}

type ConnectedClusterIdentityResponseOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponseOutput() ConnectedClusterIdentityResponseOutput {
	return o
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponseOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponseOutput {
	return o
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return o.ToConnectedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ConnectedClusterIdentityResponseOutput) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectedClusterIdentityResponse) *ConnectedClusterIdentityResponse {
		return &v
	}).(ConnectedClusterIdentityResponsePtrOutput)
}

func (o ConnectedClusterIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ConnectedClusterIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ConnectedClusterIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ConnectedClusterIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectedClusterIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterIdentityResponse)(nil)).Elem()
}

func (o ConnectedClusterIdentityResponsePtrOutput) ToConnectedClusterIdentityResponsePtrOutput() ConnectedClusterIdentityResponsePtrOutput {
	return o
}

func (o ConnectedClusterIdentityResponsePtrOutput) ToConnectedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterIdentityResponsePtrOutput {
	return o
}

func (o ConnectedClusterIdentityResponsePtrOutput) Elem() ConnectedClusterIdentityResponseOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) ConnectedClusterIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ConnectedClusterIdentityResponse
		return ret
	}).(ConnectedClusterIdentityResponseOutput)
}

func (o ConnectedClusterIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ConnectedClusterIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ConnectedClusterIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterIdentityResponse) *string {
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
	pulumi.RegisterOutputType(ConnectedClusterIdentityOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityPtrOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityResponseOutput{})
	pulumi.RegisterOutputType(ConnectedClusterIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
