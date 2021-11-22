


package v20200315preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Acl struct {
	InitiatorIqn string   `pulumi:"initiatorIqn"`
	MappedLuns   []string `pulumi:"mappedLuns"`
	Password     string   `pulumi:"password"`
	Username     string   `pulumi:"username"`
}





type AclInput interface {
	pulumi.Input

	ToAclOutput() AclOutput
	ToAclOutputWithContext(context.Context) AclOutput
}

type AclArgs struct {
	InitiatorIqn pulumi.StringInput      `pulumi:"initiatorIqn"`
	MappedLuns   pulumi.StringArrayInput `pulumi:"mappedLuns"`
	Password     pulumi.StringInput      `pulumi:"password"`
	Username     pulumi.StringInput      `pulumi:"username"`
}

func (AclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Acl)(nil)).Elem()
}

func (i AclArgs) ToAclOutput() AclOutput {
	return i.ToAclOutputWithContext(context.Background())
}

func (i AclArgs) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclOutput)
}





type AclArrayInput interface {
	pulumi.Input

	ToAclArrayOutput() AclArrayOutput
	ToAclArrayOutputWithContext(context.Context) AclArrayOutput
}

type AclArray []AclInput

func (AclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Acl)(nil)).Elem()
}

func (i AclArray) ToAclArrayOutput() AclArrayOutput {
	return i.ToAclArrayOutputWithContext(context.Background())
}

func (i AclArray) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclArrayOutput)
}

type AclOutput struct{ *pulumi.OutputState }

func (AclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Acl)(nil)).Elem()
}

func (o AclOutput) ToAclOutput() AclOutput {
	return o
}

func (o AclOutput) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return o
}

func (o AclOutput) InitiatorIqn() pulumi.StringOutput {
	return o.ApplyT(func(v Acl) string { return v.InitiatorIqn }).(pulumi.StringOutput)
}

func (o AclOutput) MappedLuns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Acl) []string { return v.MappedLuns }).(pulumi.StringArrayOutput)
}

func (o AclOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v Acl) string { return v.Password }).(pulumi.StringOutput)
}

func (o AclOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v Acl) string { return v.Username }).(pulumi.StringOutput)
}

type AclArrayOutput struct{ *pulumi.OutputState }

func (AclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Acl)(nil)).Elem()
}

func (o AclArrayOutput) ToAclArrayOutput() AclArrayOutput {
	return o
}

func (o AclArrayOutput) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return o
}

func (o AclArrayOutput) Index(i pulumi.IntInput) AclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Acl {
		return vs[0].([]Acl)[vs[1].(int)]
	}).(AclOutput)
}

type AclResponse struct {
	InitiatorIqn string   `pulumi:"initiatorIqn"`
	MappedLuns   []string `pulumi:"mappedLuns"`
	Password     string   `pulumi:"password"`
	Username     string   `pulumi:"username"`
}





type AclResponseInput interface {
	pulumi.Input

	ToAclResponseOutput() AclResponseOutput
	ToAclResponseOutputWithContext(context.Context) AclResponseOutput
}

type AclResponseArgs struct {
	InitiatorIqn pulumi.StringInput      `pulumi:"initiatorIqn"`
	MappedLuns   pulumi.StringArrayInput `pulumi:"mappedLuns"`
	Password     pulumi.StringInput      `pulumi:"password"`
	Username     pulumi.StringInput      `pulumi:"username"`
}

func (AclResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AclResponse)(nil)).Elem()
}

func (i AclResponseArgs) ToAclResponseOutput() AclResponseOutput {
	return i.ToAclResponseOutputWithContext(context.Background())
}

func (i AclResponseArgs) ToAclResponseOutputWithContext(ctx context.Context) AclResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclResponseOutput)
}





type AclResponseArrayInput interface {
	pulumi.Input

	ToAclResponseArrayOutput() AclResponseArrayOutput
	ToAclResponseArrayOutputWithContext(context.Context) AclResponseArrayOutput
}

type AclResponseArray []AclResponseInput

func (AclResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AclResponse)(nil)).Elem()
}

func (i AclResponseArray) ToAclResponseArrayOutput() AclResponseArrayOutput {
	return i.ToAclResponseArrayOutputWithContext(context.Background())
}

func (i AclResponseArray) ToAclResponseArrayOutputWithContext(ctx context.Context) AclResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclResponseArrayOutput)
}

type AclResponseOutput struct{ *pulumi.OutputState }

func (AclResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AclResponse)(nil)).Elem()
}

func (o AclResponseOutput) ToAclResponseOutput() AclResponseOutput {
	return o
}

func (o AclResponseOutput) ToAclResponseOutputWithContext(ctx context.Context) AclResponseOutput {
	return o
}

func (o AclResponseOutput) InitiatorIqn() pulumi.StringOutput {
	return o.ApplyT(func(v AclResponse) string { return v.InitiatorIqn }).(pulumi.StringOutput)
}

func (o AclResponseOutput) MappedLuns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AclResponse) []string { return v.MappedLuns }).(pulumi.StringArrayOutput)
}

func (o AclResponseOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v AclResponse) string { return v.Password }).(pulumi.StringOutput)
}

func (o AclResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v AclResponse) string { return v.Username }).(pulumi.StringOutput)
}

type AclResponseArrayOutput struct{ *pulumi.OutputState }

func (AclResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AclResponse)(nil)).Elem()
}

func (o AclResponseArrayOutput) ToAclResponseArrayOutput() AclResponseArrayOutput {
	return o
}

func (o AclResponseArrayOutput) ToAclResponseArrayOutputWithContext(ctx context.Context) AclResponseArrayOutput {
	return o
}

func (o AclResponseArrayOutput) Index(i pulumi.IntInput) AclResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AclResponse {
		return vs[0].([]AclResponse)[vs[1].(int)]
	}).(AclResponseOutput)
}

type Attributes struct {
	Authentication       bool `pulumi:"authentication"`
	ProdModeWriteProtect bool `pulumi:"prodModeWriteProtect"`
}





type AttributesInput interface {
	pulumi.Input

	ToAttributesOutput() AttributesOutput
	ToAttributesOutputWithContext(context.Context) AttributesOutput
}

type AttributesArgs struct {
	Authentication       pulumi.BoolInput `pulumi:"authentication"`
	ProdModeWriteProtect pulumi.BoolInput `pulumi:"prodModeWriteProtect"`
}

func (AttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Attributes)(nil)).Elem()
}

func (i AttributesArgs) ToAttributesOutput() AttributesOutput {
	return i.ToAttributesOutputWithContext(context.Background())
}

func (i AttributesArgs) ToAttributesOutputWithContext(ctx context.Context) AttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributesOutput)
}

type AttributesOutput struct{ *pulumi.OutputState }

func (AttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Attributes)(nil)).Elem()
}

func (o AttributesOutput) ToAttributesOutput() AttributesOutput {
	return o
}

func (o AttributesOutput) ToAttributesOutputWithContext(ctx context.Context) AttributesOutput {
	return o
}

func (o AttributesOutput) Authentication() pulumi.BoolOutput {
	return o.ApplyT(func(v Attributes) bool { return v.Authentication }).(pulumi.BoolOutput)
}

func (o AttributesOutput) ProdModeWriteProtect() pulumi.BoolOutput {
	return o.ApplyT(func(v Attributes) bool { return v.ProdModeWriteProtect }).(pulumi.BoolOutput)
}

type AttributesResponse struct {
	Authentication       bool `pulumi:"authentication"`
	ProdModeWriteProtect bool `pulumi:"prodModeWriteProtect"`
}





type AttributesResponseInput interface {
	pulumi.Input

	ToAttributesResponseOutput() AttributesResponseOutput
	ToAttributesResponseOutputWithContext(context.Context) AttributesResponseOutput
}

type AttributesResponseArgs struct {
	Authentication       pulumi.BoolInput `pulumi:"authentication"`
	ProdModeWriteProtect pulumi.BoolInput `pulumi:"prodModeWriteProtect"`
}

func (AttributesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AttributesResponse)(nil)).Elem()
}

func (i AttributesResponseArgs) ToAttributesResponseOutput() AttributesResponseOutput {
	return i.ToAttributesResponseOutputWithContext(context.Background())
}

func (i AttributesResponseArgs) ToAttributesResponseOutputWithContext(ctx context.Context) AttributesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributesResponseOutput)
}

type AttributesResponseOutput struct{ *pulumi.OutputState }

func (AttributesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttributesResponse)(nil)).Elem()
}

func (o AttributesResponseOutput) ToAttributesResponseOutput() AttributesResponseOutput {
	return o
}

func (o AttributesResponseOutput) ToAttributesResponseOutputWithContext(ctx context.Context) AttributesResponseOutput {
	return o
}

func (o AttributesResponseOutput) Authentication() pulumi.BoolOutput {
	return o.ApplyT(func(v AttributesResponse) bool { return v.Authentication }).(pulumi.BoolOutput)
}

func (o AttributesResponseOutput) ProdModeWriteProtect() pulumi.BoolOutput {
	return o.ApplyT(func(v AttributesResponse) bool { return v.ProdModeWriteProtect }).(pulumi.BoolOutput)
}

type Disk struct {
	Id string `pulumi:"id"`
}





type DiskInput interface {
	pulumi.Input

	ToDiskOutput() DiskOutput
	ToDiskOutputWithContext(context.Context) DiskOutput
}

type DiskArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Disk)(nil)).Elem()
}

func (i DiskArgs) ToDiskOutput() DiskOutput {
	return i.ToDiskOutputWithContext(context.Background())
}

func (i DiskArgs) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskOutput)
}





type DiskArrayInput interface {
	pulumi.Input

	ToDiskArrayOutput() DiskArrayOutput
	ToDiskArrayOutputWithContext(context.Context) DiskArrayOutput
}

type DiskArray []DiskInput

func (DiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Disk)(nil)).Elem()
}

func (i DiskArray) ToDiskArrayOutput() DiskArrayOutput {
	return i.ToDiskArrayOutputWithContext(context.Background())
}

func (i DiskArray) ToDiskArrayOutputWithContext(ctx context.Context) DiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskArrayOutput)
}

type DiskOutput struct{ *pulumi.OutputState }

func (DiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Disk)(nil)).Elem()
}

func (o DiskOutput) ToDiskOutput() DiskOutput {
	return o
}

func (o DiskOutput) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return o
}

func (o DiskOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Disk) string { return v.Id }).(pulumi.StringOutput)
}

type DiskArrayOutput struct{ *pulumi.OutputState }

func (DiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Disk)(nil)).Elem()
}

func (o DiskArrayOutput) ToDiskArrayOutput() DiskArrayOutput {
	return o
}

func (o DiskArrayOutput) ToDiskArrayOutputWithContext(ctx context.Context) DiskArrayOutput {
	return o
}

func (o DiskArrayOutput) Index(i pulumi.IntInput) DiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Disk {
		return vs[0].([]Disk)[vs[1].(int)]
	}).(DiskOutput)
}

type DiskResponse struct {
	Id string `pulumi:"id"`
}





type DiskResponseInput interface {
	pulumi.Input

	ToDiskResponseOutput() DiskResponseOutput
	ToDiskResponseOutputWithContext(context.Context) DiskResponseOutput
}

type DiskResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (DiskResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskResponse)(nil)).Elem()
}

func (i DiskResponseArgs) ToDiskResponseOutput() DiskResponseOutput {
	return i.ToDiskResponseOutputWithContext(context.Background())
}

func (i DiskResponseArgs) ToDiskResponseOutputWithContext(ctx context.Context) DiskResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskResponseOutput)
}





type DiskResponseArrayInput interface {
	pulumi.Input

	ToDiskResponseArrayOutput() DiskResponseArrayOutput
	ToDiskResponseArrayOutputWithContext(context.Context) DiskResponseArrayOutput
}

type DiskResponseArray []DiskResponseInput

func (DiskResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskResponse)(nil)).Elem()
}

func (i DiskResponseArray) ToDiskResponseArrayOutput() DiskResponseArrayOutput {
	return i.ToDiskResponseArrayOutputWithContext(context.Background())
}

func (i DiskResponseArray) ToDiskResponseArrayOutputWithContext(ctx context.Context) DiskResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskResponseArrayOutput)
}

type DiskResponseOutput struct{ *pulumi.OutputState }

func (DiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskResponse)(nil)).Elem()
}

func (o DiskResponseOutput) ToDiskResponseOutput() DiskResponseOutput {
	return o
}

func (o DiskResponseOutput) ToDiskResponseOutputWithContext(ctx context.Context) DiskResponseOutput {
	return o
}

func (o DiskResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DiskResponse) string { return v.Id }).(pulumi.StringOutput)
}

type DiskResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskResponse)(nil)).Elem()
}

func (o DiskResponseArrayOutput) ToDiskResponseArrayOutput() DiskResponseArrayOutput {
	return o
}

func (o DiskResponseArrayOutput) ToDiskResponseArrayOutputWithContext(ctx context.Context) DiskResponseArrayOutput {
	return o
}

func (o DiskResponseArrayOutput) Index(i pulumi.IntInput) DiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskResponse {
		return vs[0].([]DiskResponse)[vs[1].(int)]
	}).(DiskResponseOutput)
}

type IscsiLun struct {
	ManagedDiskAzureResourceId string `pulumi:"managedDiskAzureResourceId"`
	Name                       string `pulumi:"name"`
}





type IscsiLunInput interface {
	pulumi.Input

	ToIscsiLunOutput() IscsiLunOutput
	ToIscsiLunOutputWithContext(context.Context) IscsiLunOutput
}

type IscsiLunArgs struct {
	ManagedDiskAzureResourceId pulumi.StringInput `pulumi:"managedDiskAzureResourceId"`
	Name                       pulumi.StringInput `pulumi:"name"`
}

func (IscsiLunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiLun)(nil)).Elem()
}

func (i IscsiLunArgs) ToIscsiLunOutput() IscsiLunOutput {
	return i.ToIscsiLunOutputWithContext(context.Background())
}

func (i IscsiLunArgs) ToIscsiLunOutputWithContext(ctx context.Context) IscsiLunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiLunOutput)
}





type IscsiLunArrayInput interface {
	pulumi.Input

	ToIscsiLunArrayOutput() IscsiLunArrayOutput
	ToIscsiLunArrayOutputWithContext(context.Context) IscsiLunArrayOutput
}

type IscsiLunArray []IscsiLunInput

func (IscsiLunArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IscsiLun)(nil)).Elem()
}

func (i IscsiLunArray) ToIscsiLunArrayOutput() IscsiLunArrayOutput {
	return i.ToIscsiLunArrayOutputWithContext(context.Background())
}

func (i IscsiLunArray) ToIscsiLunArrayOutputWithContext(ctx context.Context) IscsiLunArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiLunArrayOutput)
}

type IscsiLunOutput struct{ *pulumi.OutputState }

func (IscsiLunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiLun)(nil)).Elem()
}

func (o IscsiLunOutput) ToIscsiLunOutput() IscsiLunOutput {
	return o
}

func (o IscsiLunOutput) ToIscsiLunOutputWithContext(ctx context.Context) IscsiLunOutput {
	return o
}

func (o IscsiLunOutput) ManagedDiskAzureResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v IscsiLun) string { return v.ManagedDiskAzureResourceId }).(pulumi.StringOutput)
}

func (o IscsiLunOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IscsiLun) string { return v.Name }).(pulumi.StringOutput)
}

type IscsiLunArrayOutput struct{ *pulumi.OutputState }

func (IscsiLunArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IscsiLun)(nil)).Elem()
}

func (o IscsiLunArrayOutput) ToIscsiLunArrayOutput() IscsiLunArrayOutput {
	return o
}

func (o IscsiLunArrayOutput) ToIscsiLunArrayOutputWithContext(ctx context.Context) IscsiLunArrayOutput {
	return o
}

func (o IscsiLunArrayOutput) Index(i pulumi.IntInput) IscsiLunOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IscsiLun {
		return vs[0].([]IscsiLun)[vs[1].(int)]
	}).(IscsiLunOutput)
}

type IscsiLunResponse struct {
	ManagedDiskAzureResourceId string `pulumi:"managedDiskAzureResourceId"`
	Name                       string `pulumi:"name"`
}





type IscsiLunResponseInput interface {
	pulumi.Input

	ToIscsiLunResponseOutput() IscsiLunResponseOutput
	ToIscsiLunResponseOutputWithContext(context.Context) IscsiLunResponseOutput
}

type IscsiLunResponseArgs struct {
	ManagedDiskAzureResourceId pulumi.StringInput `pulumi:"managedDiskAzureResourceId"`
	Name                       pulumi.StringInput `pulumi:"name"`
}

func (IscsiLunResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiLunResponse)(nil)).Elem()
}

func (i IscsiLunResponseArgs) ToIscsiLunResponseOutput() IscsiLunResponseOutput {
	return i.ToIscsiLunResponseOutputWithContext(context.Background())
}

func (i IscsiLunResponseArgs) ToIscsiLunResponseOutputWithContext(ctx context.Context) IscsiLunResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiLunResponseOutput)
}





type IscsiLunResponseArrayInput interface {
	pulumi.Input

	ToIscsiLunResponseArrayOutput() IscsiLunResponseArrayOutput
	ToIscsiLunResponseArrayOutputWithContext(context.Context) IscsiLunResponseArrayOutput
}

type IscsiLunResponseArray []IscsiLunResponseInput

func (IscsiLunResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IscsiLunResponse)(nil)).Elem()
}

func (i IscsiLunResponseArray) ToIscsiLunResponseArrayOutput() IscsiLunResponseArrayOutput {
	return i.ToIscsiLunResponseArrayOutputWithContext(context.Background())
}

func (i IscsiLunResponseArray) ToIscsiLunResponseArrayOutputWithContext(ctx context.Context) IscsiLunResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiLunResponseArrayOutput)
}

type IscsiLunResponseOutput struct{ *pulumi.OutputState }

func (IscsiLunResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiLunResponse)(nil)).Elem()
}

func (o IscsiLunResponseOutput) ToIscsiLunResponseOutput() IscsiLunResponseOutput {
	return o
}

func (o IscsiLunResponseOutput) ToIscsiLunResponseOutputWithContext(ctx context.Context) IscsiLunResponseOutput {
	return o
}

func (o IscsiLunResponseOutput) ManagedDiskAzureResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v IscsiLunResponse) string { return v.ManagedDiskAzureResourceId }).(pulumi.StringOutput)
}

func (o IscsiLunResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IscsiLunResponse) string { return v.Name }).(pulumi.StringOutput)
}

type IscsiLunResponseArrayOutput struct{ *pulumi.OutputState }

func (IscsiLunResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IscsiLunResponse)(nil)).Elem()
}

func (o IscsiLunResponseArrayOutput) ToIscsiLunResponseArrayOutput() IscsiLunResponseArrayOutput {
	return o
}

func (o IscsiLunResponseArrayOutput) ToIscsiLunResponseArrayOutputWithContext(ctx context.Context) IscsiLunResponseArrayOutput {
	return o
}

func (o IscsiLunResponseArrayOutput) Index(i pulumi.IntInput) IscsiLunResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IscsiLunResponse {
		return vs[0].([]IscsiLunResponse)[vs[1].(int)]
	}).(IscsiLunResponseOutput)
}

type SystemMetadataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemMetadataResponseInput interface {
	pulumi.Input

	ToSystemMetadataResponseOutput() SystemMetadataResponseOutput
	ToSystemMetadataResponseOutputWithContext(context.Context) SystemMetadataResponseOutput
}

type SystemMetadataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemMetadataResponse)(nil)).Elem()
}

func (i SystemMetadataResponseArgs) ToSystemMetadataResponseOutput() SystemMetadataResponseOutput {
	return i.ToSystemMetadataResponseOutputWithContext(context.Background())
}

func (i SystemMetadataResponseArgs) ToSystemMetadataResponseOutputWithContext(ctx context.Context) SystemMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemMetadataResponseOutput)
}

func (i SystemMetadataResponseArgs) ToSystemMetadataResponsePtrOutput() SystemMetadataResponsePtrOutput {
	return i.ToSystemMetadataResponsePtrOutputWithContext(context.Background())
}

func (i SystemMetadataResponseArgs) ToSystemMetadataResponsePtrOutputWithContext(ctx context.Context) SystemMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemMetadataResponseOutput).ToSystemMetadataResponsePtrOutputWithContext(ctx)
}









type SystemMetadataResponsePtrInput interface {
	pulumi.Input

	ToSystemMetadataResponsePtrOutput() SystemMetadataResponsePtrOutput
	ToSystemMetadataResponsePtrOutputWithContext(context.Context) SystemMetadataResponsePtrOutput
}

type systemMetadataResponsePtrType SystemMetadataResponseArgs

func SystemMetadataResponsePtr(v *SystemMetadataResponseArgs) SystemMetadataResponsePtrInput {
	return (*systemMetadataResponsePtrType)(v)
}

func (*systemMetadataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemMetadataResponse)(nil)).Elem()
}

func (i *systemMetadataResponsePtrType) ToSystemMetadataResponsePtrOutput() SystemMetadataResponsePtrOutput {
	return i.ToSystemMetadataResponsePtrOutputWithContext(context.Background())
}

func (i *systemMetadataResponsePtrType) ToSystemMetadataResponsePtrOutputWithContext(ctx context.Context) SystemMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemMetadataResponsePtrOutput)
}

type SystemMetadataResponseOutput struct{ *pulumi.OutputState }

func (SystemMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemMetadataResponse)(nil)).Elem()
}

func (o SystemMetadataResponseOutput) ToSystemMetadataResponseOutput() SystemMetadataResponseOutput {
	return o
}

func (o SystemMetadataResponseOutput) ToSystemMetadataResponseOutputWithContext(ctx context.Context) SystemMetadataResponseOutput {
	return o
}

func (o SystemMetadataResponseOutput) ToSystemMetadataResponsePtrOutput() SystemMetadataResponsePtrOutput {
	return o.ToSystemMetadataResponsePtrOutputWithContext(context.Background())
}

func (o SystemMetadataResponseOutput) ToSystemMetadataResponsePtrOutputWithContext(ctx context.Context) SystemMetadataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemMetadataResponse) *SystemMetadataResponse {
		return &v
	}).(SystemMetadataResponsePtrOutput)
}

func (o SystemMetadataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemMetadataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemMetadataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemMetadataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemMetadataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemMetadataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemMetadataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemMetadataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemMetadataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemMetadataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemMetadataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemMetadataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemMetadataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemMetadataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemMetadataResponse)(nil)).Elem()
}

func (o SystemMetadataResponsePtrOutput) ToSystemMetadataResponsePtrOutput() SystemMetadataResponsePtrOutput {
	return o
}

func (o SystemMetadataResponsePtrOutput) ToSystemMetadataResponsePtrOutputWithContext(ctx context.Context) SystemMetadataResponsePtrOutput {
	return o
}

func (o SystemMetadataResponsePtrOutput) Elem() SystemMetadataResponseOutput {
	return o.ApplyT(func(v *SystemMetadataResponse) SystemMetadataResponse {
		if v != nil {
			return *v
		}
		var ret SystemMetadataResponse
		return ret
	}).(SystemMetadataResponseOutput)
}

func (o SystemMetadataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemMetadataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemMetadataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemMetadataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemMetadataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemMetadataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type TargetPortalGroupCreate struct {
	Acls       []Acl      `pulumi:"acls"`
	Attributes Attributes `pulumi:"attributes"`
	Luns       []IscsiLun `pulumi:"luns"`
}





type TargetPortalGroupCreateInput interface {
	pulumi.Input

	ToTargetPortalGroupCreateOutput() TargetPortalGroupCreateOutput
	ToTargetPortalGroupCreateOutputWithContext(context.Context) TargetPortalGroupCreateOutput
}

type TargetPortalGroupCreateArgs struct {
	Acls       AclArrayInput      `pulumi:"acls"`
	Attributes AttributesInput    `pulumi:"attributes"`
	Luns       IscsiLunArrayInput `pulumi:"luns"`
}

func (TargetPortalGroupCreateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetPortalGroupCreate)(nil)).Elem()
}

func (i TargetPortalGroupCreateArgs) ToTargetPortalGroupCreateOutput() TargetPortalGroupCreateOutput {
	return i.ToTargetPortalGroupCreateOutputWithContext(context.Background())
}

func (i TargetPortalGroupCreateArgs) ToTargetPortalGroupCreateOutputWithContext(ctx context.Context) TargetPortalGroupCreateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPortalGroupCreateOutput)
}





type TargetPortalGroupCreateArrayInput interface {
	pulumi.Input

	ToTargetPortalGroupCreateArrayOutput() TargetPortalGroupCreateArrayOutput
	ToTargetPortalGroupCreateArrayOutputWithContext(context.Context) TargetPortalGroupCreateArrayOutput
}

type TargetPortalGroupCreateArray []TargetPortalGroupCreateInput

func (TargetPortalGroupCreateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetPortalGroupCreate)(nil)).Elem()
}

func (i TargetPortalGroupCreateArray) ToTargetPortalGroupCreateArrayOutput() TargetPortalGroupCreateArrayOutput {
	return i.ToTargetPortalGroupCreateArrayOutputWithContext(context.Background())
}

func (i TargetPortalGroupCreateArray) ToTargetPortalGroupCreateArrayOutputWithContext(ctx context.Context) TargetPortalGroupCreateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPortalGroupCreateArrayOutput)
}

type TargetPortalGroupCreateOutput struct{ *pulumi.OutputState }

func (TargetPortalGroupCreateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetPortalGroupCreate)(nil)).Elem()
}

func (o TargetPortalGroupCreateOutput) ToTargetPortalGroupCreateOutput() TargetPortalGroupCreateOutput {
	return o
}

func (o TargetPortalGroupCreateOutput) ToTargetPortalGroupCreateOutputWithContext(ctx context.Context) TargetPortalGroupCreateOutput {
	return o
}

func (o TargetPortalGroupCreateOutput) Acls() AclArrayOutput {
	return o.ApplyT(func(v TargetPortalGroupCreate) []Acl { return v.Acls }).(AclArrayOutput)
}

func (o TargetPortalGroupCreateOutput) Attributes() AttributesOutput {
	return o.ApplyT(func(v TargetPortalGroupCreate) Attributes { return v.Attributes }).(AttributesOutput)
}

func (o TargetPortalGroupCreateOutput) Luns() IscsiLunArrayOutput {
	return o.ApplyT(func(v TargetPortalGroupCreate) []IscsiLun { return v.Luns }).(IscsiLunArrayOutput)
}

type TargetPortalGroupCreateArrayOutput struct{ *pulumi.OutputState }

func (TargetPortalGroupCreateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetPortalGroupCreate)(nil)).Elem()
}

func (o TargetPortalGroupCreateArrayOutput) ToTargetPortalGroupCreateArrayOutput() TargetPortalGroupCreateArrayOutput {
	return o
}

func (o TargetPortalGroupCreateArrayOutput) ToTargetPortalGroupCreateArrayOutputWithContext(ctx context.Context) TargetPortalGroupCreateArrayOutput {
	return o
}

func (o TargetPortalGroupCreateArrayOutput) Index(i pulumi.IntInput) TargetPortalGroupCreateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetPortalGroupCreate {
		return vs[0].([]TargetPortalGroupCreate)[vs[1].(int)]
	}).(TargetPortalGroupCreateOutput)
}

type TargetPortalGroupResponse struct {
	Acls       []AclResponse      `pulumi:"acls"`
	Attributes AttributesResponse `pulumi:"attributes"`
	Endpoints  []string           `pulumi:"endpoints"`
	Luns       []IscsiLunResponse `pulumi:"luns"`
	Port       int                `pulumi:"port"`
	Tag        int                `pulumi:"tag"`
}





type TargetPortalGroupResponseInput interface {
	pulumi.Input

	ToTargetPortalGroupResponseOutput() TargetPortalGroupResponseOutput
	ToTargetPortalGroupResponseOutputWithContext(context.Context) TargetPortalGroupResponseOutput
}

type TargetPortalGroupResponseArgs struct {
	Acls       AclResponseArrayInput      `pulumi:"acls"`
	Attributes AttributesResponseInput    `pulumi:"attributes"`
	Endpoints  pulumi.StringArrayInput    `pulumi:"endpoints"`
	Luns       IscsiLunResponseArrayInput `pulumi:"luns"`
	Port       pulumi.IntInput            `pulumi:"port"`
	Tag        pulumi.IntInput            `pulumi:"tag"`
}

func (TargetPortalGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetPortalGroupResponse)(nil)).Elem()
}

func (i TargetPortalGroupResponseArgs) ToTargetPortalGroupResponseOutput() TargetPortalGroupResponseOutput {
	return i.ToTargetPortalGroupResponseOutputWithContext(context.Background())
}

func (i TargetPortalGroupResponseArgs) ToTargetPortalGroupResponseOutputWithContext(ctx context.Context) TargetPortalGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPortalGroupResponseOutput)
}





type TargetPortalGroupResponseArrayInput interface {
	pulumi.Input

	ToTargetPortalGroupResponseArrayOutput() TargetPortalGroupResponseArrayOutput
	ToTargetPortalGroupResponseArrayOutputWithContext(context.Context) TargetPortalGroupResponseArrayOutput
}

type TargetPortalGroupResponseArray []TargetPortalGroupResponseInput

func (TargetPortalGroupResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetPortalGroupResponse)(nil)).Elem()
}

func (i TargetPortalGroupResponseArray) ToTargetPortalGroupResponseArrayOutput() TargetPortalGroupResponseArrayOutput {
	return i.ToTargetPortalGroupResponseArrayOutputWithContext(context.Background())
}

func (i TargetPortalGroupResponseArray) ToTargetPortalGroupResponseArrayOutputWithContext(ctx context.Context) TargetPortalGroupResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPortalGroupResponseArrayOutput)
}

type TargetPortalGroupResponseOutput struct{ *pulumi.OutputState }

func (TargetPortalGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetPortalGroupResponse)(nil)).Elem()
}

func (o TargetPortalGroupResponseOutput) ToTargetPortalGroupResponseOutput() TargetPortalGroupResponseOutput {
	return o
}

func (o TargetPortalGroupResponseOutput) ToTargetPortalGroupResponseOutputWithContext(ctx context.Context) TargetPortalGroupResponseOutput {
	return o
}

func (o TargetPortalGroupResponseOutput) Acls() AclResponseArrayOutput {
	return o.ApplyT(func(v TargetPortalGroupResponse) []AclResponse { return v.Acls }).(AclResponseArrayOutput)
}

func (o TargetPortalGroupResponseOutput) Attributes() AttributesResponseOutput {
	return o.ApplyT(func(v TargetPortalGroupResponse) AttributesResponse { return v.Attributes }).(AttributesResponseOutput)
}

func (o TargetPortalGroupResponseOutput) Endpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TargetPortalGroupResponse) []string { return v.Endpoints }).(pulumi.StringArrayOutput)
}

func (o TargetPortalGroupResponseOutput) Luns() IscsiLunResponseArrayOutput {
	return o.ApplyT(func(v TargetPortalGroupResponse) []IscsiLunResponse { return v.Luns }).(IscsiLunResponseArrayOutput)
}

func (o TargetPortalGroupResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v TargetPortalGroupResponse) int { return v.Port }).(pulumi.IntOutput)
}

func (o TargetPortalGroupResponseOutput) Tag() pulumi.IntOutput {
	return o.ApplyT(func(v TargetPortalGroupResponse) int { return v.Tag }).(pulumi.IntOutput)
}

type TargetPortalGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetPortalGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetPortalGroupResponse)(nil)).Elem()
}

func (o TargetPortalGroupResponseArrayOutput) ToTargetPortalGroupResponseArrayOutput() TargetPortalGroupResponseArrayOutput {
	return o
}

func (o TargetPortalGroupResponseArrayOutput) ToTargetPortalGroupResponseArrayOutputWithContext(ctx context.Context) TargetPortalGroupResponseArrayOutput {
	return o
}

func (o TargetPortalGroupResponseArrayOutput) Index(i pulumi.IntInput) TargetPortalGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetPortalGroupResponse {
		return vs[0].([]TargetPortalGroupResponse)[vs[1].(int)]
	}).(TargetPortalGroupResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AclOutput{})
	pulumi.RegisterOutputType(AclArrayOutput{})
	pulumi.RegisterOutputType(AclResponseOutput{})
	pulumi.RegisterOutputType(AclResponseArrayOutput{})
	pulumi.RegisterOutputType(AttributesOutput{})
	pulumi.RegisterOutputType(AttributesResponseOutput{})
	pulumi.RegisterOutputType(DiskOutput{})
	pulumi.RegisterOutputType(DiskArrayOutput{})
	pulumi.RegisterOutputType(DiskResponseOutput{})
	pulumi.RegisterOutputType(DiskResponseArrayOutput{})
	pulumi.RegisterOutputType(IscsiLunOutput{})
	pulumi.RegisterOutputType(IscsiLunArrayOutput{})
	pulumi.RegisterOutputType(IscsiLunResponseOutput{})
	pulumi.RegisterOutputType(IscsiLunResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemMetadataResponseOutput{})
	pulumi.RegisterOutputType(SystemMetadataResponsePtrOutput{})
	pulumi.RegisterOutputType(TargetPortalGroupCreateOutput{})
	pulumi.RegisterOutputType(TargetPortalGroupCreateArrayOutput{})
	pulumi.RegisterOutputType(TargetPortalGroupResponseOutput{})
	pulumi.RegisterOutputType(TargetPortalGroupResponseArrayOutput{})
}
