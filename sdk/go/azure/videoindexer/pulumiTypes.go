


package videoindexer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedServiceIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type MediaServicesForPutRequest struct {
	ResourceId           *string `pulumi:"resourceId"`
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type MediaServicesForPutRequestInput interface {
	pulumi.Input

	ToMediaServicesForPutRequestOutput() MediaServicesForPutRequestOutput
	ToMediaServicesForPutRequestOutputWithContext(context.Context) MediaServicesForPutRequestOutput
}

type MediaServicesForPutRequestArgs struct {
	ResourceId           pulumi.StringPtrInput `pulumi:"resourceId"`
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (MediaServicesForPutRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServicesForPutRequest)(nil)).Elem()
}

func (i MediaServicesForPutRequestArgs) ToMediaServicesForPutRequestOutput() MediaServicesForPutRequestOutput {
	return i.ToMediaServicesForPutRequestOutputWithContext(context.Background())
}

func (i MediaServicesForPutRequestArgs) ToMediaServicesForPutRequestOutputWithContext(ctx context.Context) MediaServicesForPutRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServicesForPutRequestOutput)
}

func (i MediaServicesForPutRequestArgs) ToMediaServicesForPutRequestPtrOutput() MediaServicesForPutRequestPtrOutput {
	return i.ToMediaServicesForPutRequestPtrOutputWithContext(context.Background())
}

func (i MediaServicesForPutRequestArgs) ToMediaServicesForPutRequestPtrOutputWithContext(ctx context.Context) MediaServicesForPutRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServicesForPutRequestOutput).ToMediaServicesForPutRequestPtrOutputWithContext(ctx)
}









type MediaServicesForPutRequestPtrInput interface {
	pulumi.Input

	ToMediaServicesForPutRequestPtrOutput() MediaServicesForPutRequestPtrOutput
	ToMediaServicesForPutRequestPtrOutputWithContext(context.Context) MediaServicesForPutRequestPtrOutput
}

type mediaServicesForPutRequestPtrType MediaServicesForPutRequestArgs

func MediaServicesForPutRequestPtr(v *MediaServicesForPutRequestArgs) MediaServicesForPutRequestPtrInput {
	return (*mediaServicesForPutRequestPtrType)(v)
}

func (*mediaServicesForPutRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServicesForPutRequest)(nil)).Elem()
}

func (i *mediaServicesForPutRequestPtrType) ToMediaServicesForPutRequestPtrOutput() MediaServicesForPutRequestPtrOutput {
	return i.ToMediaServicesForPutRequestPtrOutputWithContext(context.Background())
}

func (i *mediaServicesForPutRequestPtrType) ToMediaServicesForPutRequestPtrOutputWithContext(ctx context.Context) MediaServicesForPutRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServicesForPutRequestPtrOutput)
}

type MediaServicesForPutRequestOutput struct{ *pulumi.OutputState }

func (MediaServicesForPutRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServicesForPutRequest)(nil)).Elem()
}

func (o MediaServicesForPutRequestOutput) ToMediaServicesForPutRequestOutput() MediaServicesForPutRequestOutput {
	return o
}

func (o MediaServicesForPutRequestOutput) ToMediaServicesForPutRequestOutputWithContext(ctx context.Context) MediaServicesForPutRequestOutput {
	return o
}

func (o MediaServicesForPutRequestOutput) ToMediaServicesForPutRequestPtrOutput() MediaServicesForPutRequestPtrOutput {
	return o.ToMediaServicesForPutRequestPtrOutputWithContext(context.Background())
}

func (o MediaServicesForPutRequestOutput) ToMediaServicesForPutRequestPtrOutputWithContext(ctx context.Context) MediaServicesForPutRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaServicesForPutRequest) *MediaServicesForPutRequest {
		return &v
	}).(MediaServicesForPutRequestPtrOutput)
}

func (o MediaServicesForPutRequestOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MediaServicesForPutRequest) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o MediaServicesForPutRequestOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MediaServicesForPutRequest) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type MediaServicesForPutRequestPtrOutput struct{ *pulumi.OutputState }

func (MediaServicesForPutRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServicesForPutRequest)(nil)).Elem()
}

func (o MediaServicesForPutRequestPtrOutput) ToMediaServicesForPutRequestPtrOutput() MediaServicesForPutRequestPtrOutput {
	return o
}

func (o MediaServicesForPutRequestPtrOutput) ToMediaServicesForPutRequestPtrOutputWithContext(ctx context.Context) MediaServicesForPutRequestPtrOutput {
	return o
}

func (o MediaServicesForPutRequestPtrOutput) Elem() MediaServicesForPutRequestOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequest) MediaServicesForPutRequest {
		if v != nil {
			return *v
		}
		var ret MediaServicesForPutRequest
		return ret
	}).(MediaServicesForPutRequestOutput)
}

func (o MediaServicesForPutRequestPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequest) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func (o MediaServicesForPutRequestPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequest) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type MediaServicesForPutRequestResponse struct {
	ResourceId           *string `pulumi:"resourceId"`
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}

type MediaServicesForPutRequestResponseOutput struct{ *pulumi.OutputState }

func (MediaServicesForPutRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServicesForPutRequestResponse)(nil)).Elem()
}

func (o MediaServicesForPutRequestResponseOutput) ToMediaServicesForPutRequestResponseOutput() MediaServicesForPutRequestResponseOutput {
	return o
}

func (o MediaServicesForPutRequestResponseOutput) ToMediaServicesForPutRequestResponseOutputWithContext(ctx context.Context) MediaServicesForPutRequestResponseOutput {
	return o
}

func (o MediaServicesForPutRequestResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MediaServicesForPutRequestResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o MediaServicesForPutRequestResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MediaServicesForPutRequestResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type MediaServicesForPutRequestResponsePtrOutput struct{ *pulumi.OutputState }

func (MediaServicesForPutRequestResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServicesForPutRequestResponse)(nil)).Elem()
}

func (o MediaServicesForPutRequestResponsePtrOutput) ToMediaServicesForPutRequestResponsePtrOutput() MediaServicesForPutRequestResponsePtrOutput {
	return o
}

func (o MediaServicesForPutRequestResponsePtrOutput) ToMediaServicesForPutRequestResponsePtrOutputWithContext(ctx context.Context) MediaServicesForPutRequestResponsePtrOutput {
	return o
}

func (o MediaServicesForPutRequestResponsePtrOutput) Elem() MediaServicesForPutRequestResponseOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequestResponse) MediaServicesForPutRequestResponse {
		if v != nil {
			return *v
		}
		var ret MediaServicesForPutRequestResponse
		return ret
	}).(MediaServicesForPutRequestResponseOutput)
}

func (o MediaServicesForPutRequestResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func (o MediaServicesForPutRequestResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
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

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(MediaServicesForPutRequestOutput{})
	pulumi.RegisterOutputType(MediaServicesForPutRequestPtrOutput{})
	pulumi.RegisterOutputType(MediaServicesForPutRequestResponseOutput{})
	pulumi.RegisterOutputType(MediaServicesForPutRequestResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}
