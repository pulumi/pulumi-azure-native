


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

type PrivateLinkAssociationProperties struct {
	PrivateLink         *string `pulumi:"privateLink"`
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}





type PrivateLinkAssociationPropertiesInput interface {
	pulumi.Input

	ToPrivateLinkAssociationPropertiesOutput() PrivateLinkAssociationPropertiesOutput
	ToPrivateLinkAssociationPropertiesOutputWithContext(context.Context) PrivateLinkAssociationPropertiesOutput
}

type PrivateLinkAssociationPropertiesArgs struct {
	PrivateLink         pulumi.StringPtrInput `pulumi:"privateLink"`
	PublicNetworkAccess pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (PrivateLinkAssociationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkAssociationProperties)(nil)).Elem()
}

func (i PrivateLinkAssociationPropertiesArgs) ToPrivateLinkAssociationPropertiesOutput() PrivateLinkAssociationPropertiesOutput {
	return i.ToPrivateLinkAssociationPropertiesOutputWithContext(context.Background())
}

func (i PrivateLinkAssociationPropertiesArgs) ToPrivateLinkAssociationPropertiesOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkAssociationPropertiesOutput)
}

func (i PrivateLinkAssociationPropertiesArgs) ToPrivateLinkAssociationPropertiesPtrOutput() PrivateLinkAssociationPropertiesPtrOutput {
	return i.ToPrivateLinkAssociationPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateLinkAssociationPropertiesArgs) ToPrivateLinkAssociationPropertiesPtrOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkAssociationPropertiesOutput).ToPrivateLinkAssociationPropertiesPtrOutputWithContext(ctx)
}









type PrivateLinkAssociationPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateLinkAssociationPropertiesPtrOutput() PrivateLinkAssociationPropertiesPtrOutput
	ToPrivateLinkAssociationPropertiesPtrOutputWithContext(context.Context) PrivateLinkAssociationPropertiesPtrOutput
}

type privateLinkAssociationPropertiesPtrType PrivateLinkAssociationPropertiesArgs

func PrivateLinkAssociationPropertiesPtr(v *PrivateLinkAssociationPropertiesArgs) PrivateLinkAssociationPropertiesPtrInput {
	return (*privateLinkAssociationPropertiesPtrType)(v)
}

func (*privateLinkAssociationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkAssociationProperties)(nil)).Elem()
}

func (i *privateLinkAssociationPropertiesPtrType) ToPrivateLinkAssociationPropertiesPtrOutput() PrivateLinkAssociationPropertiesPtrOutput {
	return i.ToPrivateLinkAssociationPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateLinkAssociationPropertiesPtrType) ToPrivateLinkAssociationPropertiesPtrOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkAssociationPropertiesPtrOutput)
}

type PrivateLinkAssociationPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateLinkAssociationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkAssociationProperties)(nil)).Elem()
}

func (o PrivateLinkAssociationPropertiesOutput) ToPrivateLinkAssociationPropertiesOutput() PrivateLinkAssociationPropertiesOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesOutput) ToPrivateLinkAssociationPropertiesOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesOutput) ToPrivateLinkAssociationPropertiesPtrOutput() PrivateLinkAssociationPropertiesPtrOutput {
	return o.ToPrivateLinkAssociationPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateLinkAssociationPropertiesOutput) ToPrivateLinkAssociationPropertiesPtrOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkAssociationProperties) *PrivateLinkAssociationProperties {
		return &v
	}).(PrivateLinkAssociationPropertiesPtrOutput)
}

func (o PrivateLinkAssociationPropertiesOutput) PrivateLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationProperties) *string { return v.PrivateLink }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkAssociationPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type PrivateLinkAssociationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkAssociationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkAssociationProperties)(nil)).Elem()
}

func (o PrivateLinkAssociationPropertiesPtrOutput) ToPrivateLinkAssociationPropertiesPtrOutput() PrivateLinkAssociationPropertiesPtrOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesPtrOutput) ToPrivateLinkAssociationPropertiesPtrOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesPtrOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesPtrOutput) Elem() PrivateLinkAssociationPropertiesOutput {
	return o.ApplyT(func(v *PrivateLinkAssociationProperties) PrivateLinkAssociationProperties {
		if v != nil {
			return *v
		}
		var ret PrivateLinkAssociationProperties
		return ret
	}).(PrivateLinkAssociationPropertiesOutput)
}

func (o PrivateLinkAssociationPropertiesPtrOutput) PrivateLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkAssociationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLink
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkAssociationPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkAssociationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkAssociationPropertiesExpandedResponse struct {
	PrivateLink         *string `pulumi:"privateLink"`
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	Scope               *string `pulumi:"scope"`
	TenantID            *string `pulumi:"tenantID"`
}

type PrivateLinkAssociationPropertiesExpandedResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkAssociationPropertiesExpandedResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkAssociationPropertiesExpandedResponse)(nil)).Elem()
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) ToPrivateLinkAssociationPropertiesExpandedResponseOutput() PrivateLinkAssociationPropertiesExpandedResponseOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) ToPrivateLinkAssociationPropertiesExpandedResponseOutputWithContext(ctx context.Context) PrivateLinkAssociationPropertiesExpandedResponseOutput {
	return o
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) PrivateLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationPropertiesExpandedResponse) *string { return v.PrivateLink }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationPropertiesExpandedResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationPropertiesExpandedResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkAssociationPropertiesExpandedResponseOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkAssociationPropertiesExpandedResponse) *string { return v.TenantID }).(pulumi.StringPtrOutput)
}

type ResourceManagementPrivateLinkEndpointConnectionsResponse struct {
	PrivateEndpointConnections []string `pulumi:"privateEndpointConnections"`
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

func (o ResourceManagementPrivateLinkEndpointConnectionsResponseOutput) PrivateEndpointConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceManagementPrivateLinkEndpointConnectionsResponse) []string {
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

func init() {
	pulumi.RegisterOutputType(ManagementLockOwnerOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerArrayOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerResponseOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkAssociationPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateLinkAssociationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkAssociationPropertiesExpandedResponseOutput{})
	pulumi.RegisterOutputType(ResourceManagementPrivateLinkEndpointConnectionsResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
