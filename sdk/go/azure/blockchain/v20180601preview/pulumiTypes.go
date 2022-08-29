


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiKeyResponse struct {
	KeyName *string `pulumi:"keyName"`
	Value   *string `pulumi:"value"`
}

type ApiKeyResponseOutput struct{ *pulumi.OutputState }

func (ApiKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiKeyResponse)(nil)).Elem()
}

func (o ApiKeyResponseOutput) ToApiKeyResponseOutput() ApiKeyResponseOutput {
	return o
}

func (o ApiKeyResponseOutput) ToApiKeyResponseOutputWithContext(ctx context.Context) ApiKeyResponseOutput {
	return o
}

func (o ApiKeyResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiKeyResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o ApiKeyResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiKeyResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ApiKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (ApiKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiKeyResponse)(nil)).Elem()
}

func (o ApiKeyResponseArrayOutput) ToApiKeyResponseArrayOutput() ApiKeyResponseArrayOutput {
	return o
}

func (o ApiKeyResponseArrayOutput) ToApiKeyResponseArrayOutputWithContext(ctx context.Context) ApiKeyResponseArrayOutput {
	return o
}

func (o ApiKeyResponseArrayOutput) Index(i pulumi.IntInput) ApiKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiKeyResponse {
		return vs[0].([]ApiKeyResponse)[vs[1].(int)]
	}).(ApiKeyResponseOutput)
}

type BlockchainMemberNodesSku struct {
	Capacity *int `pulumi:"capacity"`
}





type BlockchainMemberNodesSkuInput interface {
	pulumi.Input

	ToBlockchainMemberNodesSkuOutput() BlockchainMemberNodesSkuOutput
	ToBlockchainMemberNodesSkuOutputWithContext(context.Context) BlockchainMemberNodesSkuOutput
}

type BlockchainMemberNodesSkuArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
}

func (BlockchainMemberNodesSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlockchainMemberNodesSku)(nil)).Elem()
}

func (i BlockchainMemberNodesSkuArgs) ToBlockchainMemberNodesSkuOutput() BlockchainMemberNodesSkuOutput {
	return i.ToBlockchainMemberNodesSkuOutputWithContext(context.Background())
}

func (i BlockchainMemberNodesSkuArgs) ToBlockchainMemberNodesSkuOutputWithContext(ctx context.Context) BlockchainMemberNodesSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockchainMemberNodesSkuOutput)
}

func (i BlockchainMemberNodesSkuArgs) ToBlockchainMemberNodesSkuPtrOutput() BlockchainMemberNodesSkuPtrOutput {
	return i.ToBlockchainMemberNodesSkuPtrOutputWithContext(context.Background())
}

func (i BlockchainMemberNodesSkuArgs) ToBlockchainMemberNodesSkuPtrOutputWithContext(ctx context.Context) BlockchainMemberNodesSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockchainMemberNodesSkuOutput).ToBlockchainMemberNodesSkuPtrOutputWithContext(ctx)
}









type BlockchainMemberNodesSkuPtrInput interface {
	pulumi.Input

	ToBlockchainMemberNodesSkuPtrOutput() BlockchainMemberNodesSkuPtrOutput
	ToBlockchainMemberNodesSkuPtrOutputWithContext(context.Context) BlockchainMemberNodesSkuPtrOutput
}

type blockchainMemberNodesSkuPtrType BlockchainMemberNodesSkuArgs

func BlockchainMemberNodesSkuPtr(v *BlockchainMemberNodesSkuArgs) BlockchainMemberNodesSkuPtrInput {
	return (*blockchainMemberNodesSkuPtrType)(v)
}

func (*blockchainMemberNodesSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockchainMemberNodesSku)(nil)).Elem()
}

func (i *blockchainMemberNodesSkuPtrType) ToBlockchainMemberNodesSkuPtrOutput() BlockchainMemberNodesSkuPtrOutput {
	return i.ToBlockchainMemberNodesSkuPtrOutputWithContext(context.Background())
}

func (i *blockchainMemberNodesSkuPtrType) ToBlockchainMemberNodesSkuPtrOutputWithContext(ctx context.Context) BlockchainMemberNodesSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockchainMemberNodesSkuPtrOutput)
}

type BlockchainMemberNodesSkuOutput struct{ *pulumi.OutputState }

func (BlockchainMemberNodesSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlockchainMemberNodesSku)(nil)).Elem()
}

func (o BlockchainMemberNodesSkuOutput) ToBlockchainMemberNodesSkuOutput() BlockchainMemberNodesSkuOutput {
	return o
}

func (o BlockchainMemberNodesSkuOutput) ToBlockchainMemberNodesSkuOutputWithContext(ctx context.Context) BlockchainMemberNodesSkuOutput {
	return o
}

func (o BlockchainMemberNodesSkuOutput) ToBlockchainMemberNodesSkuPtrOutput() BlockchainMemberNodesSkuPtrOutput {
	return o.ToBlockchainMemberNodesSkuPtrOutputWithContext(context.Background())
}

func (o BlockchainMemberNodesSkuOutput) ToBlockchainMemberNodesSkuPtrOutputWithContext(ctx context.Context) BlockchainMemberNodesSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlockchainMemberNodesSku) *BlockchainMemberNodesSku {
		return &v
	}).(BlockchainMemberNodesSkuPtrOutput)
}

func (o BlockchainMemberNodesSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BlockchainMemberNodesSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

type BlockchainMemberNodesSkuPtrOutput struct{ *pulumi.OutputState }

func (BlockchainMemberNodesSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockchainMemberNodesSku)(nil)).Elem()
}

func (o BlockchainMemberNodesSkuPtrOutput) ToBlockchainMemberNodesSkuPtrOutput() BlockchainMemberNodesSkuPtrOutput {
	return o
}

func (o BlockchainMemberNodesSkuPtrOutput) ToBlockchainMemberNodesSkuPtrOutputWithContext(ctx context.Context) BlockchainMemberNodesSkuPtrOutput {
	return o
}

func (o BlockchainMemberNodesSkuPtrOutput) Elem() BlockchainMemberNodesSkuOutput {
	return o.ApplyT(func(v *BlockchainMemberNodesSku) BlockchainMemberNodesSku {
		if v != nil {
			return *v
		}
		var ret BlockchainMemberNodesSku
		return ret
	}).(BlockchainMemberNodesSkuOutput)
}

func (o BlockchainMemberNodesSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BlockchainMemberNodesSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

type BlockchainMemberNodesSkuResponse struct {
	Capacity *int `pulumi:"capacity"`
}

type BlockchainMemberNodesSkuResponseOutput struct{ *pulumi.OutputState }

func (BlockchainMemberNodesSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlockchainMemberNodesSkuResponse)(nil)).Elem()
}

func (o BlockchainMemberNodesSkuResponseOutput) ToBlockchainMemberNodesSkuResponseOutput() BlockchainMemberNodesSkuResponseOutput {
	return o
}

func (o BlockchainMemberNodesSkuResponseOutput) ToBlockchainMemberNodesSkuResponseOutputWithContext(ctx context.Context) BlockchainMemberNodesSkuResponseOutput {
	return o
}

func (o BlockchainMemberNodesSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BlockchainMemberNodesSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

type BlockchainMemberNodesSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (BlockchainMemberNodesSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockchainMemberNodesSkuResponse)(nil)).Elem()
}

func (o BlockchainMemberNodesSkuResponsePtrOutput) ToBlockchainMemberNodesSkuResponsePtrOutput() BlockchainMemberNodesSkuResponsePtrOutput {
	return o
}

func (o BlockchainMemberNodesSkuResponsePtrOutput) ToBlockchainMemberNodesSkuResponsePtrOutputWithContext(ctx context.Context) BlockchainMemberNodesSkuResponsePtrOutput {
	return o
}

func (o BlockchainMemberNodesSkuResponsePtrOutput) Elem() BlockchainMemberNodesSkuResponseOutput {
	return o.ApplyT(func(v *BlockchainMemberNodesSkuResponse) BlockchainMemberNodesSkuResponse {
		if v != nil {
			return *v
		}
		var ret BlockchainMemberNodesSkuResponse
		return ret
	}).(BlockchainMemberNodesSkuResponseOutput)
}

func (o BlockchainMemberNodesSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BlockchainMemberNodesSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

type ConsortiumResponse struct {
	Name     *string `pulumi:"name"`
	Protocol *string `pulumi:"protocol"`
}

type ConsortiumResponseOutput struct{ *pulumi.OutputState }

func (ConsortiumResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsortiumResponse)(nil)).Elem()
}

func (o ConsortiumResponseOutput) ToConsortiumResponseOutput() ConsortiumResponseOutput {
	return o
}

func (o ConsortiumResponseOutput) ToConsortiumResponseOutputWithContext(ctx context.Context) ConsortiumResponseOutput {
	return o
}

func (o ConsortiumResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsortiumResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConsortiumResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsortiumResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type ConsortiumResponseArrayOutput struct{ *pulumi.OutputState }

func (ConsortiumResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConsortiumResponse)(nil)).Elem()
}

func (o ConsortiumResponseArrayOutput) ToConsortiumResponseArrayOutput() ConsortiumResponseArrayOutput {
	return o
}

func (o ConsortiumResponseArrayOutput) ToConsortiumResponseArrayOutputWithContext(ctx context.Context) ConsortiumResponseArrayOutput {
	return o
}

func (o ConsortiumResponseArrayOutput) Index(i pulumi.IntInput) ConsortiumResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConsortiumResponse {
		return vs[0].([]ConsortiumResponse)[vs[1].(int)]
	}).(ConsortiumResponseOutput)
}

type FirewallRule struct {
	EndIpAddress   *string `pulumi:"endIpAddress"`
	RuleName       *string `pulumi:"ruleName"`
	StartIpAddress *string `pulumi:"startIpAddress"`
}





type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(context.Context) FirewallRuleOutput
}

type FirewallRuleArgs struct {
	EndIpAddress   pulumi.StringPtrInput `pulumi:"endIpAddress"`
	RuleName       pulumi.StringPtrInput `pulumi:"ruleName"`
	StartIpAddress pulumi.StringPtrInput `pulumi:"startIpAddress"`
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleArgs) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i FirewallRuleArgs) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}





type FirewallRuleArrayInput interface {
	pulumi.Input

	ToFirewallRuleArrayOutput() FirewallRuleArrayOutput
	ToFirewallRuleArrayOutputWithContext(context.Context) FirewallRuleArrayOutput
}

type FirewallRuleArray []FirewallRuleInput

func (FirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallRule)(nil)).Elem()
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return i.ToFirewallRuleArrayOutputWithContext(context.Background())
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleArrayOutput)
}

type FirewallRuleOutput struct{ *pulumi.OutputState }

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) EndIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FirewallRule) *string { return v.EndIpAddress }).(pulumi.StringPtrOutput)
}

func (o FirewallRuleOutput) RuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FirewallRule) *string { return v.RuleName }).(pulumi.StringPtrOutput)
}

func (o FirewallRuleOutput) StartIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FirewallRule) *string { return v.StartIpAddress }).(pulumi.StringPtrOutput)
}

type FirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (FirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallRule)(nil)).Elem()
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) Index(i pulumi.IntInput) FirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallRule {
		return vs[0].([]FirewallRule)[vs[1].(int)]
	}).(FirewallRuleOutput)
}

type FirewallRuleResponse struct {
	EndIpAddress   *string `pulumi:"endIpAddress"`
	RuleName       *string `pulumi:"ruleName"`
	StartIpAddress *string `pulumi:"startIpAddress"`
}

type FirewallRuleResponseOutput struct{ *pulumi.OutputState }

func (FirewallRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRuleResponse)(nil)).Elem()
}

func (o FirewallRuleResponseOutput) ToFirewallRuleResponseOutput() FirewallRuleResponseOutput {
	return o
}

func (o FirewallRuleResponseOutput) ToFirewallRuleResponseOutputWithContext(ctx context.Context) FirewallRuleResponseOutput {
	return o
}

func (o FirewallRuleResponseOutput) EndIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FirewallRuleResponse) *string { return v.EndIpAddress }).(pulumi.StringPtrOutput)
}

func (o FirewallRuleResponseOutput) RuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FirewallRuleResponse) *string { return v.RuleName }).(pulumi.StringPtrOutput)
}

func (o FirewallRuleResponseOutput) StartIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FirewallRuleResponse) *string { return v.StartIpAddress }).(pulumi.StringPtrOutput)
}

type FirewallRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (FirewallRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallRuleResponse)(nil)).Elem()
}

func (o FirewallRuleResponseArrayOutput) ToFirewallRuleResponseArrayOutput() FirewallRuleResponseArrayOutput {
	return o
}

func (o FirewallRuleResponseArrayOutput) ToFirewallRuleResponseArrayOutputWithContext(ctx context.Context) FirewallRuleResponseArrayOutput {
	return o
}

func (o FirewallRuleResponseArrayOutput) Index(i pulumi.IntInput) FirewallRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallRuleResponse {
		return vs[0].([]FirewallRuleResponse)[vs[1].(int)]
	}).(FirewallRuleResponseOutput)
}

type Sku struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
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

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
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
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
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

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiKeyResponseOutput{})
	pulumi.RegisterOutputType(ApiKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(BlockchainMemberNodesSkuOutput{})
	pulumi.RegisterOutputType(BlockchainMemberNodesSkuPtrOutput{})
	pulumi.RegisterOutputType(BlockchainMemberNodesSkuResponseOutput{})
	pulumi.RegisterOutputType(BlockchainMemberNodesSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ConsortiumResponseOutput{})
	pulumi.RegisterOutputType(ConsortiumResponseArrayOutput{})
	pulumi.RegisterOutputType(FirewallRuleOutput{})
	pulumi.RegisterOutputType(FirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(FirewallRuleResponseOutput{})
	pulumi.RegisterOutputType(FirewallRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
