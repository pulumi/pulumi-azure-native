


package v20180115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContentHash struct {
	Algorithm string `pulumi:"algorithm"`
	Value     string `pulumi:"value"`
}





type ContentHashInput interface {
	pulumi.Input

	ToContentHashOutput() ContentHashOutput
	ToContentHashOutputWithContext(context.Context) ContentHashOutput
}

type ContentHashArgs struct {
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	Value     pulumi.StringInput `pulumi:"value"`
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

func (o ContentHashOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHash) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o ContentHashOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHash) string { return v.Value }).(pulumi.StringOutput)
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
		return &v.Algorithm
	}).(pulumi.StringPtrOutput)
}

func (o ContentHashPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHash) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type ContentSource struct {
	Hash    *ContentHash `pulumi:"hash"`
	Type    *string      `pulumi:"type"`
	Value   *string      `pulumi:"value"`
	Version *string      `pulumi:"version"`
}





type ContentSourceInput interface {
	pulumi.Input

	ToContentSourceOutput() ContentSourceOutput
	ToContentSourceOutputWithContext(context.Context) ContentSourceOutput
}

type ContentSourceArgs struct {
	Hash    ContentHashPtrInput   `pulumi:"hash"`
	Type    pulumi.StringPtrInput `pulumi:"type"`
	Value   pulumi.StringPtrInput `pulumi:"value"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (ContentSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSource)(nil)).Elem()
}

func (i ContentSourceArgs) ToContentSourceOutput() ContentSourceOutput {
	return i.ToContentSourceOutputWithContext(context.Background())
}

func (i ContentSourceArgs) ToContentSourceOutputWithContext(ctx context.Context) ContentSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentSourceOutput)
}

type ContentSourceOutput struct{ *pulumi.OutputState }

func (ContentSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSource)(nil)).Elem()
}

func (o ContentSourceOutput) ToContentSourceOutput() ContentSourceOutput {
	return o
}

func (o ContentSourceOutput) ToContentSourceOutputWithContext(ctx context.Context) ContentSourceOutput {
	return o
}

func (o ContentSourceOutput) Hash() ContentHashPtrOutput {
	return o.ApplyT(func(v ContentSource) *ContentHash { return v.Hash }).(ContentHashPtrOutput)
}

func (o ContentSourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSource) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ContentSourceOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSource) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func (o ContentSourceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSource) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type DscConfigurationAssociationProperty struct {
	Name *string `pulumi:"name"`
}





type DscConfigurationAssociationPropertyInput interface {
	pulumi.Input

	ToDscConfigurationAssociationPropertyOutput() DscConfigurationAssociationPropertyOutput
	ToDscConfigurationAssociationPropertyOutputWithContext(context.Context) DscConfigurationAssociationPropertyOutput
}

type DscConfigurationAssociationPropertyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DscConfigurationAssociationPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationAssociationProperty)(nil)).Elem()
}

func (i DscConfigurationAssociationPropertyArgs) ToDscConfigurationAssociationPropertyOutput() DscConfigurationAssociationPropertyOutput {
	return i.ToDscConfigurationAssociationPropertyOutputWithContext(context.Background())
}

func (i DscConfigurationAssociationPropertyArgs) ToDscConfigurationAssociationPropertyOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyOutput)
}

type DscConfigurationAssociationPropertyOutput struct{ *pulumi.OutputState }

func (DscConfigurationAssociationPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationAssociationProperty)(nil)).Elem()
}

func (o DscConfigurationAssociationPropertyOutput) ToDscConfigurationAssociationPropertyOutput() DscConfigurationAssociationPropertyOutput {
	return o
}

func (o DscConfigurationAssociationPropertyOutput) ToDscConfigurationAssociationPropertyOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyOutput {
	return o
}

func (o DscConfigurationAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscConfigurationAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DscConfigurationAssociationPropertyResponse struct {
	Name *string `pulumi:"name"`
}

type DscConfigurationAssociationPropertyResponseOutput struct{ *pulumi.OutputState }

func (DscConfigurationAssociationPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationAssociationPropertyResponse)(nil)).Elem()
}

func (o DscConfigurationAssociationPropertyResponseOutput) ToDscConfigurationAssociationPropertyResponseOutput() DscConfigurationAssociationPropertyResponseOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponseOutput) ToDscConfigurationAssociationPropertyResponseOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponseOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscConfigurationAssociationPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DscConfigurationAssociationPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (DscConfigurationAssociationPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfigurationAssociationPropertyResponse)(nil)).Elem()
}

func (o DscConfigurationAssociationPropertyResponsePtrOutput) ToDscConfigurationAssociationPropertyResponsePtrOutput() DscConfigurationAssociationPropertyResponsePtrOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponsePtrOutput) ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponsePtrOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponsePtrOutput) Elem() DscConfigurationAssociationPropertyResponseOutput {
	return o.ApplyT(func(v *DscConfigurationAssociationPropertyResponse) DscConfigurationAssociationPropertyResponse {
		if v != nil {
			return *v
		}
		var ret DscConfigurationAssociationPropertyResponse
		return ret
	}).(DscConfigurationAssociationPropertyResponseOutput)
}

func (o DscConfigurationAssociationPropertyResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfigurationAssociationPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContentHashOutput{})
	pulumi.RegisterOutputType(ContentHashPtrOutput{})
	pulumi.RegisterOutputType(ContentSourceOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyResponseOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyResponsePtrOutput{})
}
