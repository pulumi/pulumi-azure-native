


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Acl struct {
	InitiatorIqn string   `pulumi:"initiatorIqn"`
	MappedLuns   []string `pulumi:"mappedLuns"`
}





type AclInput interface {
	pulumi.Input

	ToAclOutput() AclOutput
	ToAclOutputWithContext(context.Context) AclOutput
}

type AclArgs struct {
	InitiatorIqn pulumi.StringInput      `pulumi:"initiatorIqn"`
	MappedLuns   pulumi.StringArrayInput `pulumi:"mappedLuns"`
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
	Lun                        int    `pulumi:"lun"`
	ManagedDiskAzureResourceId string `pulumi:"managedDiskAzureResourceId"`
	Name                       string `pulumi:"name"`
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

func (o IscsiLunResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v IscsiLunResponse) int { return v.Lun }).(pulumi.IntOutput)
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

type Sku struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SystemMetadataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
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

func init() {
	pulumi.RegisterOutputType(AclOutput{})
	pulumi.RegisterOutputType(AclArrayOutput{})
	pulumi.RegisterOutputType(AclResponseOutput{})
	pulumi.RegisterOutputType(AclResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskOutput{})
	pulumi.RegisterOutputType(DiskArrayOutput{})
	pulumi.RegisterOutputType(DiskResponseOutput{})
	pulumi.RegisterOutputType(DiskResponseArrayOutput{})
	pulumi.RegisterOutputType(IscsiLunOutput{})
	pulumi.RegisterOutputType(IscsiLunArrayOutput{})
	pulumi.RegisterOutputType(IscsiLunResponseOutput{})
	pulumi.RegisterOutputType(IscsiLunResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SystemMetadataResponseOutput{})
}
