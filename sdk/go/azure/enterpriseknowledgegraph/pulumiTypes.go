


package enterpriseknowledgegraph

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnterpriseKnowledgeGraphProperties struct {
	Description       *string     `pulumi:"description"`
	Metadata          interface{} `pulumi:"metadata"`
	ProvisioningState *string     `pulumi:"provisioningState"`
}





type EnterpriseKnowledgeGraphPropertiesInput interface {
	pulumi.Input

	ToEnterpriseKnowledgeGraphPropertiesOutput() EnterpriseKnowledgeGraphPropertiesOutput
	ToEnterpriseKnowledgeGraphPropertiesOutputWithContext(context.Context) EnterpriseKnowledgeGraphPropertiesOutput
}

type EnterpriseKnowledgeGraphPropertiesArgs struct {
	Description       pulumi.StringPtrInput `pulumi:"description"`
	Metadata          pulumi.Input          `pulumi:"metadata"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (EnterpriseKnowledgeGraphPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseKnowledgeGraphProperties)(nil)).Elem()
}

func (i EnterpriseKnowledgeGraphPropertiesArgs) ToEnterpriseKnowledgeGraphPropertiesOutput() EnterpriseKnowledgeGraphPropertiesOutput {
	return i.ToEnterpriseKnowledgeGraphPropertiesOutputWithContext(context.Background())
}

func (i EnterpriseKnowledgeGraphPropertiesArgs) ToEnterpriseKnowledgeGraphPropertiesOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseKnowledgeGraphPropertiesOutput)
}

func (i EnterpriseKnowledgeGraphPropertiesArgs) ToEnterpriseKnowledgeGraphPropertiesPtrOutput() EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return i.ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(context.Background())
}

func (i EnterpriseKnowledgeGraphPropertiesArgs) ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseKnowledgeGraphPropertiesOutput).ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(ctx)
}









type EnterpriseKnowledgeGraphPropertiesPtrInput interface {
	pulumi.Input

	ToEnterpriseKnowledgeGraphPropertiesPtrOutput() EnterpriseKnowledgeGraphPropertiesPtrOutput
	ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(context.Context) EnterpriseKnowledgeGraphPropertiesPtrOutput
}

type enterpriseKnowledgeGraphPropertiesPtrType EnterpriseKnowledgeGraphPropertiesArgs

func EnterpriseKnowledgeGraphPropertiesPtr(v *EnterpriseKnowledgeGraphPropertiesArgs) EnterpriseKnowledgeGraphPropertiesPtrInput {
	return (*enterpriseKnowledgeGraphPropertiesPtrType)(v)
}

func (*enterpriseKnowledgeGraphPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseKnowledgeGraphProperties)(nil)).Elem()
}

func (i *enterpriseKnowledgeGraphPropertiesPtrType) ToEnterpriseKnowledgeGraphPropertiesPtrOutput() EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return i.ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(context.Background())
}

func (i *enterpriseKnowledgeGraphPropertiesPtrType) ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseKnowledgeGraphPropertiesPtrOutput)
}

type EnterpriseKnowledgeGraphPropertiesOutput struct{ *pulumi.OutputState }

func (EnterpriseKnowledgeGraphPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseKnowledgeGraphProperties)(nil)).Elem()
}

func (o EnterpriseKnowledgeGraphPropertiesOutput) ToEnterpriseKnowledgeGraphPropertiesOutput() EnterpriseKnowledgeGraphPropertiesOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesOutput) ToEnterpriseKnowledgeGraphPropertiesOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesOutput) ToEnterpriseKnowledgeGraphPropertiesPtrOutput() EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return o.ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(context.Background())
}

func (o EnterpriseKnowledgeGraphPropertiesOutput) ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnterpriseKnowledgeGraphProperties) *EnterpriseKnowledgeGraphProperties {
		return &v
	}).(EnterpriseKnowledgeGraphPropertiesPtrOutput)
}

func (o EnterpriseKnowledgeGraphPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnterpriseKnowledgeGraphPropertiesOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphProperties) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o EnterpriseKnowledgeGraphPropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type EnterpriseKnowledgeGraphPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EnterpriseKnowledgeGraphPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseKnowledgeGraphProperties)(nil)).Elem()
}

func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) ToEnterpriseKnowledgeGraphPropertiesPtrOutput() EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) Elem() EnterpriseKnowledgeGraphPropertiesOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphProperties) EnterpriseKnowledgeGraphProperties {
		if v != nil {
			return *v
		}
		var ret EnterpriseKnowledgeGraphProperties
		return ret
	}).(EnterpriseKnowledgeGraphPropertiesOutput)
}

func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.AnyOutput)
}

func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type EnterpriseKnowledgeGraphPropertiesResponse struct {
	Description       *string     `pulumi:"description"`
	Metadata          interface{} `pulumi:"metadata"`
	ProvisioningState *string     `pulumi:"provisioningState"`
}

type EnterpriseKnowledgeGraphPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EnterpriseKnowledgeGraphPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseKnowledgeGraphPropertiesResponse)(nil)).Elem()
}

func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) ToEnterpriseKnowledgeGraphPropertiesResponseOutput() EnterpriseKnowledgeGraphPropertiesResponseOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) ToEnterpriseKnowledgeGraphPropertiesResponseOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesResponseOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphPropertiesResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
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

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
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

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
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

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
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
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EnterpriseKnowledgeGraphPropertiesOutput{})
	pulumi.RegisterOutputType(EnterpriseKnowledgeGraphPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EnterpriseKnowledgeGraphPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
