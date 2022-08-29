


package v20150201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContentHash struct {
	Algorithm *string `pulumi:"algorithm"`
	Value     *string `pulumi:"value"`
}





type ContentHashInput interface {
	pulumi.Input

	ToContentHashOutput() ContentHashOutput
	ToContentHashOutputWithContext(context.Context) ContentHashOutput
}

type ContentHashArgs struct {
	Algorithm pulumi.StringPtrInput `pulumi:"algorithm"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (ContentHashArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHash)(nil)).Elem()
}

func (i ContentHashArgs) ToContentHashOutput() ContentHashOutput {
	return i.ToContentHashOutputWithContext(context.Background())
}

func (i ContentHashArgs) ToContentHashOutputWithContext(ctx context.Context) ContentHashOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashOutput)
}

func (i ContentHashArgs) ToContentHashPtrOutput() ContentHashPtrOutput {
	return i.ToContentHashPtrOutputWithContext(context.Background())
}

func (i ContentHashArgs) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashOutput).ToContentHashPtrOutputWithContext(ctx)
}









type ContentHashPtrInput interface {
	pulumi.Input

	ToContentHashPtrOutput() ContentHashPtrOutput
	ToContentHashPtrOutputWithContext(context.Context) ContentHashPtrOutput
}

type contentHashPtrType ContentHashArgs

func ContentHashPtr(v *ContentHashArgs) ContentHashPtrInput {
	return (*contentHashPtrType)(v)
}

func (*contentHashPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHash)(nil)).Elem()
}

func (i *contentHashPtrType) ToContentHashPtrOutput() ContentHashPtrOutput {
	return i.ToContentHashPtrOutputWithContext(context.Background())
}

func (i *contentHashPtrType) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashPtrOutput)
}

type ContentHashOutput struct{ *pulumi.OutputState }

func (ContentHashOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHash)(nil)).Elem()
}

func (o ContentHashOutput) ToContentHashOutput() ContentHashOutput {
	return o
}

func (o ContentHashOutput) ToContentHashOutputWithContext(ctx context.Context) ContentHashOutput {
	return o
}

func (o ContentHashOutput) ToContentHashPtrOutput() ContentHashPtrOutput {
	return o.ToContentHashPtrOutputWithContext(context.Background())
}

func (o ContentHashOutput) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentHash) *ContentHash {
		return &v
	}).(ContentHashPtrOutput)
}

func (o ContentHashOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentHash) *string { return v.Algorithm }).(pulumi.StringPtrOutput)
}

func (o ContentHashOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentHash) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ContentHashPtrOutput struct{ *pulumi.OutputState }

func (ContentHashPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHash)(nil)).Elem()
}

func (o ContentHashPtrOutput) ToContentHashPtrOutput() ContentHashPtrOutput {
	return o
}

func (o ContentHashPtrOutput) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return o
}

func (o ContentHashPtrOutput) Elem() ContentHashOutput {
	return o.ApplyT(func(v *ContentHash) ContentHash {
		if v != nil {
			return *v
		}
		var ret ContentHash
		return ret
	}).(ContentHashOutput)
}

func (o ContentHashPtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHash) *string {
		if v == nil {
			return nil
		}
		return v.Algorithm
	}).(pulumi.StringPtrOutput)
}

func (o ContentHashPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHash) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type ContentHashResponse struct {
	Algorithm *string `pulumi:"algorithm"`
	Value     *string `pulumi:"value"`
}

type ContentHashResponseOutput struct{ *pulumi.OutputState }

func (ContentHashResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHashResponse)(nil)).Elem()
}

func (o ContentHashResponseOutput) ToContentHashResponseOutput() ContentHashResponseOutput {
	return o
}

func (o ContentHashResponseOutput) ToContentHashResponseOutputWithContext(ctx context.Context) ContentHashResponseOutput {
	return o
}

func (o ContentHashResponseOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentHashResponse) *string { return v.Algorithm }).(pulumi.StringPtrOutput)
}

func (o ContentHashResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentHashResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ContentHashResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentHashResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHashResponse)(nil)).Elem()
}

func (o ContentHashResponsePtrOutput) ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput {
	return o
}

func (o ContentHashResponsePtrOutput) ToContentHashResponsePtrOutputWithContext(ctx context.Context) ContentHashResponsePtrOutput {
	return o
}

func (o ContentHashResponsePtrOutput) Elem() ContentHashResponseOutput {
	return o.ApplyT(func(v *ContentHashResponse) ContentHashResponse {
		if v != nil {
			return *v
		}
		var ret ContentHashResponse
		return ret
	}).(ContentHashResponseOutput)
}

func (o ContentHashResponsePtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return v.Algorithm
	}).(pulumi.StringPtrOutput)
}

func (o ContentHashResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type ContentLink struct {
	ContentHash    *ContentHash `pulumi:"contentHash"`
	ContentSize    *float64     `pulumi:"contentSize"`
	ContentVersion *string      `pulumi:"contentVersion"`
	Metadata       interface{}  `pulumi:"metadata"`
	Uri            *string      `pulumi:"uri"`
}





type ContentLinkInput interface {
	pulumi.Input

	ToContentLinkOutput() ContentLinkOutput
	ToContentLinkOutputWithContext(context.Context) ContentLinkOutput
}

type ContentLinkArgs struct {
	ContentHash    ContentHashPtrInput    `pulumi:"contentHash"`
	ContentSize    pulumi.Float64PtrInput `pulumi:"contentSize"`
	ContentVersion pulumi.StringPtrInput  `pulumi:"contentVersion"`
	Metadata       pulumi.Input           `pulumi:"metadata"`
	Uri            pulumi.StringPtrInput  `pulumi:"uri"`
}

func (ContentLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLink)(nil)).Elem()
}

func (i ContentLinkArgs) ToContentLinkOutput() ContentLinkOutput {
	return i.ToContentLinkOutputWithContext(context.Background())
}

func (i ContentLinkArgs) ToContentLinkOutputWithContext(ctx context.Context) ContentLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkOutput)
}

func (i ContentLinkArgs) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return i.ToContentLinkPtrOutputWithContext(context.Background())
}

func (i ContentLinkArgs) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkOutput).ToContentLinkPtrOutputWithContext(ctx)
}









type ContentLinkPtrInput interface {
	pulumi.Input

	ToContentLinkPtrOutput() ContentLinkPtrOutput
	ToContentLinkPtrOutputWithContext(context.Context) ContentLinkPtrOutput
}

type contentLinkPtrType ContentLinkArgs

func ContentLinkPtr(v *ContentLinkArgs) ContentLinkPtrInput {
	return (*contentLinkPtrType)(v)
}

func (*contentLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLink)(nil)).Elem()
}

func (i *contentLinkPtrType) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return i.ToContentLinkPtrOutputWithContext(context.Background())
}

func (i *contentLinkPtrType) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkPtrOutput)
}

type ContentLinkOutput struct{ *pulumi.OutputState }

func (ContentLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLink)(nil)).Elem()
}

func (o ContentLinkOutput) ToContentLinkOutput() ContentLinkOutput {
	return o
}

func (o ContentLinkOutput) ToContentLinkOutputWithContext(ctx context.Context) ContentLinkOutput {
	return o
}

func (o ContentLinkOutput) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return o.ToContentLinkPtrOutputWithContext(context.Background())
}

func (o ContentLinkOutput) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentLink) *ContentLink {
		return &v
	}).(ContentLinkPtrOutput)
}

func (o ContentLinkOutput) ContentHash() ContentHashPtrOutput {
	return o.ApplyT(func(v ContentLink) *ContentHash { return v.ContentHash }).(ContentHashPtrOutput)
}

func (o ContentLinkOutput) ContentSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContentLink) *float64 { return v.ContentSize }).(pulumi.Float64PtrOutput)
}

func (o ContentLinkOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLink) *string { return v.ContentVersion }).(pulumi.StringPtrOutput)
}

func (o ContentLinkOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentLink) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o ContentLinkOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLink) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ContentLinkPtrOutput struct{ *pulumi.OutputState }

func (ContentLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLink)(nil)).Elem()
}

func (o ContentLinkPtrOutput) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return o
}

func (o ContentLinkPtrOutput) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return o
}

func (o ContentLinkPtrOutput) Elem() ContentLinkOutput {
	return o.ApplyT(func(v *ContentLink) ContentLink {
		if v != nil {
			return *v
		}
		var ret ContentLink
		return ret
	}).(ContentLinkOutput)
}

func (o ContentLinkPtrOutput) ContentHash() ContentHashPtrOutput {
	return o.ApplyT(func(v *ContentLink) *ContentHash {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(ContentHashPtrOutput)
}

func (o ContentLinkPtrOutput) ContentSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContentLink) *float64 {
		if v == nil {
			return nil
		}
		return v.ContentSize
	}).(pulumi.Float64PtrOutput)
}

func (o ContentLinkPtrOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLink) *string {
		if v == nil {
			return nil
		}
		return v.ContentVersion
	}).(pulumi.StringPtrOutput)
}

func (o ContentLinkPtrOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *ContentLink) interface{} {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.AnyOutput)
}

func (o ContentLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLink) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type ContentLinkResponse struct {
	ContentHash    *ContentHashResponse `pulumi:"contentHash"`
	ContentSize    *float64             `pulumi:"contentSize"`
	ContentVersion *string              `pulumi:"contentVersion"`
	Metadata       interface{}          `pulumi:"metadata"`
	Uri            *string              `pulumi:"uri"`
}

type ContentLinkResponseOutput struct{ *pulumi.OutputState }

func (ContentLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLinkResponse)(nil)).Elem()
}

func (o ContentLinkResponseOutput) ToContentLinkResponseOutput() ContentLinkResponseOutput {
	return o
}

func (o ContentLinkResponseOutput) ToContentLinkResponseOutputWithContext(ctx context.Context) ContentLinkResponseOutput {
	return o
}

func (o ContentLinkResponseOutput) ContentHash() ContentHashResponsePtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *ContentHashResponse { return v.ContentHash }).(ContentHashResponsePtrOutput)
}

func (o ContentLinkResponseOutput) ContentSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *float64 { return v.ContentSize }).(pulumi.Float64PtrOutput)
}

func (o ContentLinkResponseOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *string { return v.ContentVersion }).(pulumi.StringPtrOutput)
}

func (o ContentLinkResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentLinkResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o ContentLinkResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ContentLinkResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentLinkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLinkResponse)(nil)).Elem()
}

func (o ContentLinkResponsePtrOutput) ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput {
	return o
}

func (o ContentLinkResponsePtrOutput) ToContentLinkResponsePtrOutputWithContext(ctx context.Context) ContentLinkResponsePtrOutput {
	return o
}

func (o ContentLinkResponsePtrOutput) Elem() ContentLinkResponseOutput {
	return o.ApplyT(func(v *ContentLinkResponse) ContentLinkResponse {
		if v != nil {
			return *v
		}
		var ret ContentLinkResponse
		return ret
	}).(ContentLinkResponseOutput)
}

func (o ContentLinkResponsePtrOutput) ContentHash() ContentHashResponsePtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *ContentHashResponse {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(ContentHashResponsePtrOutput)
}

func (o ContentLinkResponsePtrOutput) ContentSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.ContentSize
	}).(pulumi.Float64PtrOutput)
}

func (o ContentLinkResponsePtrOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentVersion
	}).(pulumi.StringPtrOutput)
}

func (o ContentLinkResponsePtrOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *ContentLinkResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.AnyOutput)
}

func (o ContentLinkResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type ResourceReference struct {
	Id *string `pulumi:"id"`
}





type ResourceReferenceInput interface {
	pulumi.Input

	ToResourceReferenceOutput() ResourceReferenceOutput
	ToResourceReferenceOutputWithContext(context.Context) ResourceReferenceOutput
}

type ResourceReferenceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArgs) ToResourceReferenceOutput() ResourceReferenceOutput {
	return i.ToResourceReferenceOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput)
}

func (i ResourceReferenceArgs) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return i.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput).ToResourceReferencePtrOutputWithContext(ctx)
}









type ResourceReferencePtrInput interface {
	pulumi.Input

	ToResourceReferencePtrOutput() ResourceReferencePtrOutput
	ToResourceReferencePtrOutputWithContext(context.Context) ResourceReferencePtrOutput
}

type resourceReferencePtrType ResourceReferenceArgs

func ResourceReferencePtr(v *ResourceReferenceArgs) ResourceReferencePtrInput {
	return (*resourceReferencePtrType)(v)
}

func (*resourceReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReference)(nil)).Elem()
}

func (i *resourceReferencePtrType) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return i.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (i *resourceReferencePtrType) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferencePtrOutput)
}

type ResourceReferenceOutput struct{ *pulumi.OutputState }

func (ResourceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceOutput) ToResourceReferenceOutput() ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return o.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (o ResourceReferenceOutput) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceReference) *ResourceReference {
		return &v
	}).(ResourceReferencePtrOutput)
}

func (o ResourceReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferencePtrOutput struct{ *pulumi.OutputState }

func (ResourceReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReference)(nil)).Elem()
}

func (o ResourceReferencePtrOutput) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return o
}

func (o ResourceReferencePtrOutput) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return o
}

func (o ResourceReferencePtrOutput) Elem() ResourceReferenceOutput {
	return o.ApplyT(func(v *ResourceReference) ResourceReference {
		if v != nil {
			return *v
		}
		var ret ResourceReference
		return ret
	}).(ResourceReferenceOutput)
}

func (o ResourceReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourceReferenceResponse struct {
	Id   *string `pulumi:"id"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}

type ResourceReferenceResponseOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ResourceReferenceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceReferenceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ResourceReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) Elem() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) ResourceReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ResourceReferenceResponse
		return ret
	}).(ResourceReferenceResponseOutput)
}

func (o ResourceReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ResourceReferenceResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceReferenceResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name *SkuName           `pulumi:"name"`
	Plan *ResourceReference `pulumi:"plan"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name SkuNamePtrInput           `pulumi:"name"`
	Plan ResourceReferencePtrInput `pulumi:"plan"`
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

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
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

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v Sku) *SkuName { return v.Name }).(SkuNamePtrOutput)
}

func (o SkuOutput) Plan() ResourceReferencePtrOutput {
	return o.ApplyT(func(v Sku) *ResourceReference { return v.Plan }).(ResourceReferencePtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v *Sku) *SkuName {
		if v == nil {
			return nil
		}
		return v.Name
	}).(SkuNamePtrOutput)
}

func (o SkuPtrOutput) Plan() ResourceReferencePtrOutput {
	return o.ApplyT(func(v *Sku) *ResourceReference {
		if v == nil {
			return nil
		}
		return v.Plan
	}).(ResourceReferencePtrOutput)
}

type SkuResponse struct {
	Name *string                    `pulumi:"name"`
	Plan *ResourceReferenceResponse `pulumi:"plan"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Plan() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v SkuResponse) *ResourceReferenceResponse { return v.Plan }).(ResourceReferenceResponsePtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Plan() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *SkuResponse) *ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.Plan
	}).(ResourceReferenceResponsePtrOutput)
}

type WorkflowParameter struct {
	Metadata interface{}    `pulumi:"metadata"`
	Type     *ParameterType `pulumi:"type"`
	Value    interface{}    `pulumi:"value"`
}





type WorkflowParameterInput interface {
	pulumi.Input

	ToWorkflowParameterOutput() WorkflowParameterOutput
	ToWorkflowParameterOutputWithContext(context.Context) WorkflowParameterOutput
}

type WorkflowParameterArgs struct {
	Metadata pulumi.Input          `pulumi:"metadata"`
	Type     ParameterTypePtrInput `pulumi:"type"`
	Value    pulumi.Input          `pulumi:"value"`
}

func (WorkflowParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowParameter)(nil)).Elem()
}

func (i WorkflowParameterArgs) ToWorkflowParameterOutput() WorkflowParameterOutput {
	return i.ToWorkflowParameterOutputWithContext(context.Background())
}

func (i WorkflowParameterArgs) ToWorkflowParameterOutputWithContext(ctx context.Context) WorkflowParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowParameterOutput)
}





type WorkflowParameterMapInput interface {
	pulumi.Input

	ToWorkflowParameterMapOutput() WorkflowParameterMapOutput
	ToWorkflowParameterMapOutputWithContext(context.Context) WorkflowParameterMapOutput
}

type WorkflowParameterMap map[string]WorkflowParameterInput

func (WorkflowParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkflowParameter)(nil)).Elem()
}

func (i WorkflowParameterMap) ToWorkflowParameterMapOutput() WorkflowParameterMapOutput {
	return i.ToWorkflowParameterMapOutputWithContext(context.Background())
}

func (i WorkflowParameterMap) ToWorkflowParameterMapOutputWithContext(ctx context.Context) WorkflowParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowParameterMapOutput)
}

type WorkflowParameterOutput struct{ *pulumi.OutputState }

func (WorkflowParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowParameter)(nil)).Elem()
}

func (o WorkflowParameterOutput) ToWorkflowParameterOutput() WorkflowParameterOutput {
	return o
}

func (o WorkflowParameterOutput) ToWorkflowParameterOutputWithContext(ctx context.Context) WorkflowParameterOutput {
	return o
}

func (o WorkflowParameterOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameter) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o WorkflowParameterOutput) Type() ParameterTypePtrOutput {
	return o.ApplyT(func(v WorkflowParameter) *ParameterType { return v.Type }).(ParameterTypePtrOutput)
}

func (o WorkflowParameterOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameter) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type WorkflowParameterMapOutput struct{ *pulumi.OutputState }

func (WorkflowParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkflowParameter)(nil)).Elem()
}

func (o WorkflowParameterMapOutput) ToWorkflowParameterMapOutput() WorkflowParameterMapOutput {
	return o
}

func (o WorkflowParameterMapOutput) ToWorkflowParameterMapOutputWithContext(ctx context.Context) WorkflowParameterMapOutput {
	return o
}

func (o WorkflowParameterMapOutput) MapIndex(k pulumi.StringInput) WorkflowParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkflowParameter {
		return vs[0].(map[string]WorkflowParameter)[vs[1].(string)]
	}).(WorkflowParameterOutput)
}

type WorkflowParameterResponse struct {
	Metadata interface{} `pulumi:"metadata"`
	Type     *string     `pulumi:"type"`
	Value    interface{} `pulumi:"value"`
}

type WorkflowParameterResponseOutput struct{ *pulumi.OutputState }

func (WorkflowParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowParameterResponse)(nil)).Elem()
}

func (o WorkflowParameterResponseOutput) ToWorkflowParameterResponseOutput() WorkflowParameterResponseOutput {
	return o
}

func (o WorkflowParameterResponseOutput) ToWorkflowParameterResponseOutputWithContext(ctx context.Context) WorkflowParameterResponseOutput {
	return o
}

func (o WorkflowParameterResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameterResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o WorkflowParameterResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowParameterResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o WorkflowParameterResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameterResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type WorkflowParameterResponseMapOutput struct{ *pulumi.OutputState }

func (WorkflowParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkflowParameterResponse)(nil)).Elem()
}

func (o WorkflowParameterResponseMapOutput) ToWorkflowParameterResponseMapOutput() WorkflowParameterResponseMapOutput {
	return o
}

func (o WorkflowParameterResponseMapOutput) ToWorkflowParameterResponseMapOutputWithContext(ctx context.Context) WorkflowParameterResponseMapOutput {
	return o
}

func (o WorkflowParameterResponseMapOutput) MapIndex(k pulumi.StringInput) WorkflowParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkflowParameterResponse {
		return vs[0].(map[string]WorkflowParameterResponse)[vs[1].(string)]
	}).(WorkflowParameterResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ContentHashOutput{})
	pulumi.RegisterOutputType(ContentHashPtrOutput{})
	pulumi.RegisterOutputType(ContentHashResponseOutput{})
	pulumi.RegisterOutputType(ContentHashResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentLinkOutput{})
	pulumi.RegisterOutputType(ContentLinkPtrOutput{})
	pulumi.RegisterOutputType(ContentLinkResponseOutput{})
	pulumi.RegisterOutputType(ContentLinkResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceOutput{})
	pulumi.RegisterOutputType(ResourceReferencePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkflowParameterOutput{})
	pulumi.RegisterOutputType(WorkflowParameterMapOutput{})
	pulumi.RegisterOutputType(WorkflowParameterResponseOutput{})
	pulumi.RegisterOutputType(WorkflowParameterResponseMapOutput{})
}
