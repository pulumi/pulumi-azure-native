


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementLockOwner struct {
	ApplicationId *string `pulumi:"applicationId"`
}





type ManagementLockOwnerInput interface {
	pulumi.Input

	ToManagementLockOwnerOutput() ManagementLockOwnerOutput
	ToManagementLockOwnerOutputWithContext(context.Context) ManagementLockOwnerOutput
}

type ManagementLockOwnerArgs struct {
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
}

func (ManagementLockOwnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwner)(nil)).Elem()
}

func (i ManagementLockOwnerArgs) ToManagementLockOwnerOutput() ManagementLockOwnerOutput {
	return i.ToManagementLockOwnerOutputWithContext(context.Background())
}

func (i ManagementLockOwnerArgs) ToManagementLockOwnerOutputWithContext(ctx context.Context) ManagementLockOwnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerOutput)
}





type ManagementLockOwnerArrayInput interface {
	pulumi.Input

	ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput
	ToManagementLockOwnerArrayOutputWithContext(context.Context) ManagementLockOwnerArrayOutput
}

type ManagementLockOwnerArray []ManagementLockOwnerInput

func (ManagementLockOwnerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwner)(nil)).Elem()
}

func (i ManagementLockOwnerArray) ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput {
	return i.ToManagementLockOwnerArrayOutputWithContext(context.Background())
}

func (i ManagementLockOwnerArray) ToManagementLockOwnerArrayOutputWithContext(ctx context.Context) ManagementLockOwnerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerArrayOutput)
}

type ManagementLockOwnerOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwner)(nil)).Elem()
}

func (o ManagementLockOwnerOutput) ToManagementLockOwnerOutput() ManagementLockOwnerOutput {
	return o
}

func (o ManagementLockOwnerOutput) ToManagementLockOwnerOutputWithContext(ctx context.Context) ManagementLockOwnerOutput {
	return o
}

func (o ManagementLockOwnerOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementLockOwner) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

type ManagementLockOwnerArrayOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwner)(nil)).Elem()
}

func (o ManagementLockOwnerArrayOutput) ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput {
	return o
}

func (o ManagementLockOwnerArrayOutput) ToManagementLockOwnerArrayOutputWithContext(ctx context.Context) ManagementLockOwnerArrayOutput {
	return o
}

func (o ManagementLockOwnerArrayOutput) Index(i pulumi.IntInput) ManagementLockOwnerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementLockOwner {
		return vs[0].([]ManagementLockOwner)[vs[1].(int)]
	}).(ManagementLockOwnerOutput)
}

type ManagementLockOwnerResponse struct {
	ApplicationId *string `pulumi:"applicationId"`
}





type ManagementLockOwnerResponseInput interface {
	pulumi.Input

	ToManagementLockOwnerResponseOutput() ManagementLockOwnerResponseOutput
	ToManagementLockOwnerResponseOutputWithContext(context.Context) ManagementLockOwnerResponseOutput
}

type ManagementLockOwnerResponseArgs struct {
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
}

func (ManagementLockOwnerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwnerResponse)(nil)).Elem()
}

func (i ManagementLockOwnerResponseArgs) ToManagementLockOwnerResponseOutput() ManagementLockOwnerResponseOutput {
	return i.ToManagementLockOwnerResponseOutputWithContext(context.Background())
}

func (i ManagementLockOwnerResponseArgs) ToManagementLockOwnerResponseOutputWithContext(ctx context.Context) ManagementLockOwnerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerResponseOutput)
}





type ManagementLockOwnerResponseArrayInput interface {
	pulumi.Input

	ToManagementLockOwnerResponseArrayOutput() ManagementLockOwnerResponseArrayOutput
	ToManagementLockOwnerResponseArrayOutputWithContext(context.Context) ManagementLockOwnerResponseArrayOutput
}

type ManagementLockOwnerResponseArray []ManagementLockOwnerResponseInput

func (ManagementLockOwnerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwnerResponse)(nil)).Elem()
}

func (i ManagementLockOwnerResponseArray) ToManagementLockOwnerResponseArrayOutput() ManagementLockOwnerResponseArrayOutput {
	return i.ToManagementLockOwnerResponseArrayOutputWithContext(context.Background())
}

func (i ManagementLockOwnerResponseArray) ToManagementLockOwnerResponseArrayOutputWithContext(ctx context.Context) ManagementLockOwnerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerResponseArrayOutput)
}

type ManagementLockOwnerResponseOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwnerResponse)(nil)).Elem()
}

func (o ManagementLockOwnerResponseOutput) ToManagementLockOwnerResponseOutput() ManagementLockOwnerResponseOutput {
	return o
}

func (o ManagementLockOwnerResponseOutput) ToManagementLockOwnerResponseOutputWithContext(ctx context.Context) ManagementLockOwnerResponseOutput {
	return o
}

func (o ManagementLockOwnerResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementLockOwnerResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

type ManagementLockOwnerResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwnerResponse)(nil)).Elem()
}

func (o ManagementLockOwnerResponseArrayOutput) ToManagementLockOwnerResponseArrayOutput() ManagementLockOwnerResponseArrayOutput {
	return o
}

func (o ManagementLockOwnerResponseArrayOutput) ToManagementLockOwnerResponseArrayOutputWithContext(ctx context.Context) ManagementLockOwnerResponseArrayOutput {
	return o
}

func (o ManagementLockOwnerResponseArrayOutput) Index(i pulumi.IntInput) ManagementLockOwnerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementLockOwnerResponse {
		return vs[0].([]ManagementLockOwnerResponse)[vs[1].(int)]
	}).(ManagementLockOwnerResponseOutput)
}

type ResourceManagementPrivateLinkEndpointConnectionsResponse struct {
	PrivateEndpointConnections []string `pulumi:"privateEndpointConnections"`
}





type ResourceManagementPrivateLinkEndpointConnectionsResponseInput interface {
	pulumi.Input

	ToResourceManagementPrivateLinkEndpointConnectionsResponseOutput() ResourceManagementPrivateLinkEndpointConnectionsResponseOutput
	ToResourceManagementPrivateLinkEndpointConnectionsResponseOutputWithContext(context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponseOutput
}

type ResourceManagementPrivateLinkEndpointConnectionsResponseArgs struct {
	PrivateEndpointConnections pulumi.StringArrayInput `pulumi:"privateEndpointConnections"`
}

func (ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceManagementPrivateLinkEndpointConnectionsResponse)(nil)).Elem()
}

func (i ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ToResourceManagementPrivateLinkEndpointConnectionsResponseOutput() ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return i.ToResourceManagementPrivateLinkEndpointConnectionsResponseOutputWithContext(context.Background())
}

func (i ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ToResourceManagementPrivateLinkEndpointConnectionsResponseOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput)
}

func (i ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput() ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return i.ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(context.Background())
}

func (i ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput).ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(ctx)
}









type ResourceManagementPrivateLinkEndpointConnectionsResponsePtrInput interface {
	pulumi.Input

	ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput() ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput
	ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput
}

type resourceManagementPrivateLinkEndpointConnectionsResponsePtrType ResourceManagementPrivateLinkEndpointConnectionsResponseArgs

func ResourceManagementPrivateLinkEndpointConnectionsResponsePtr(v *ResourceManagementPrivateLinkEndpointConnectionsResponseArgs) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrInput {
	return (*resourceManagementPrivateLinkEndpointConnectionsResponsePtrType)(v)
}

func (*resourceManagementPrivateLinkEndpointConnectionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceManagementPrivateLinkEndpointConnectionsResponse)(nil)).Elem()
}

func (i *resourceManagementPrivateLinkEndpointConnectionsResponsePtrType) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput() ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return i.ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(context.Background())
}

func (i *resourceManagementPrivateLinkEndpointConnectionsResponsePtrType) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput)
}

type ResourceManagementPrivateLinkEndpointConnectionsResponseOutput struct{ *pulumi.OutputState }

func (ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceManagementPrivateLinkEndpointConnectionsResponse)(nil)).Elem()
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponseOutput() ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return o
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponseOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return o
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput() ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return o.ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(context.Background())
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceManagementPrivateLinkEndpointConnectionsResponse) *ResourceManagementPrivateLinkEndpointConnectionsResponse {
		return &v
	}).(ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput)
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) PrivateEndpointConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceManagementPrivateLinkEndpointConnectionsResponse) []string {
		return v.PrivateEndpointConnections
	}).(pulumi.StringArrayOutput)
}

type ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceManagementPrivateLinkEndpointConnectionsResponse)(nil)).Elem()
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput() ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return o
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput) ToResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput {
	return o
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput) Elem() ResourceManagementPrivateLinkEndpointConnectionsResponseOutput {
	return o.ApplyT(func(v *ResourceManagementPrivateLinkEndpointConnectionsResponse) ResourceManagementPrivateLinkEndpointConnectionsResponse {
		if v != nil {
			return *v
		}
		var ret ResourceManagementPrivateLinkEndpointConnectionsResponse
		return ret
	}).(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput)
}

func (o ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput) PrivateEndpointConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceManagementPrivateLinkEndpointConnectionsResponse) []string {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(pulumi.StringArrayOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementLockOwnerInput)(nil)).Elem(), ManagementLockOwnerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementLockOwnerArrayInput)(nil)).Elem(), ManagementLockOwnerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementLockOwnerResponseInput)(nil)).Elem(), ManagementLockOwnerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementLockOwnerResponseArrayInput)(nil)).Elem(), ManagementLockOwnerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceManagementPrivateLinkEndpointConnectionsResponseInput)(nil)).Elem(), ResourceManagementPrivateLinkEndpointConnectionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceManagementPrivateLinkEndpointConnectionsResponsePtrInput)(nil)).Elem(), ResourceManagementPrivateLinkEndpointConnectionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterOutputType(ManagementLockOwnerOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerArrayOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerResponseOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput{})
	pulumi.RegisterOutputType(ResourceManagementPrivateLinkEndpointConnectionsResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
