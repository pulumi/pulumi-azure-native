


package v20190401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ControllerConnectionDetailsResponse struct {
	OrchestratorSpecificConnectionDetails *KubernetesConnectionDetailsResponse `pulumi:"orchestratorSpecificConnectionDetails"`
}

type ControllerConnectionDetailsResponseOutput struct{ *pulumi.OutputState }

func (ControllerConnectionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ControllerConnectionDetailsResponse)(nil)).Elem()
}

func (o ControllerConnectionDetailsResponseOutput) ToControllerConnectionDetailsResponseOutput() ControllerConnectionDetailsResponseOutput {
	return o
}

func (o ControllerConnectionDetailsResponseOutput) ToControllerConnectionDetailsResponseOutputWithContext(ctx context.Context) ControllerConnectionDetailsResponseOutput {
	return o
}

func (o ControllerConnectionDetailsResponseOutput) OrchestratorSpecificConnectionDetails() KubernetesConnectionDetailsResponsePtrOutput {
	return o.ApplyT(func(v ControllerConnectionDetailsResponse) *KubernetesConnectionDetailsResponse {
		return v.OrchestratorSpecificConnectionDetails
	}).(KubernetesConnectionDetailsResponsePtrOutput)
}

type ControllerConnectionDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (ControllerConnectionDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ControllerConnectionDetailsResponse)(nil)).Elem()
}

func (o ControllerConnectionDetailsResponseArrayOutput) ToControllerConnectionDetailsResponseArrayOutput() ControllerConnectionDetailsResponseArrayOutput {
	return o
}

func (o ControllerConnectionDetailsResponseArrayOutput) ToControllerConnectionDetailsResponseArrayOutputWithContext(ctx context.Context) ControllerConnectionDetailsResponseArrayOutput {
	return o
}

func (o ControllerConnectionDetailsResponseArrayOutput) Index(i pulumi.IntInput) ControllerConnectionDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ControllerConnectionDetailsResponse {
		return vs[0].([]ControllerConnectionDetailsResponse)[vs[1].(int)]
	}).(ControllerConnectionDetailsResponseOutput)
}

type KubernetesConnectionDetailsResponse struct {
	InstanceType string  `pulumi:"instanceType"`
	KubeConfig   *string `pulumi:"kubeConfig"`
}

type KubernetesConnectionDetailsResponseOutput struct{ *pulumi.OutputState }

func (KubernetesConnectionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesConnectionDetailsResponse)(nil)).Elem()
}

func (o KubernetesConnectionDetailsResponseOutput) ToKubernetesConnectionDetailsResponseOutput() KubernetesConnectionDetailsResponseOutput {
	return o
}

func (o KubernetesConnectionDetailsResponseOutput) ToKubernetesConnectionDetailsResponseOutputWithContext(ctx context.Context) KubernetesConnectionDetailsResponseOutput {
	return o
}

func (o KubernetesConnectionDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesConnectionDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o KubernetesConnectionDetailsResponseOutput) KubeConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesConnectionDetailsResponse) *string { return v.KubeConfig }).(pulumi.StringPtrOutput)
}

type KubernetesConnectionDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (KubernetesConnectionDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesConnectionDetailsResponse)(nil)).Elem()
}

func (o KubernetesConnectionDetailsResponsePtrOutput) ToKubernetesConnectionDetailsResponsePtrOutput() KubernetesConnectionDetailsResponsePtrOutput {
	return o
}

func (o KubernetesConnectionDetailsResponsePtrOutput) ToKubernetesConnectionDetailsResponsePtrOutputWithContext(ctx context.Context) KubernetesConnectionDetailsResponsePtrOutput {
	return o
}

func (o KubernetesConnectionDetailsResponsePtrOutput) Elem() KubernetesConnectionDetailsResponseOutput {
	return o.ApplyT(func(v *KubernetesConnectionDetailsResponse) KubernetesConnectionDetailsResponse {
		if v != nil {
			return *v
		}
		var ret KubernetesConnectionDetailsResponse
		return ret
	}).(KubernetesConnectionDetailsResponseOutput)
}

func (o KubernetesConnectionDetailsResponsePtrOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesConnectionDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InstanceType
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesConnectionDetailsResponsePtrOutput) KubeConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesConnectionDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.KubeConfig
	}).(pulumi.StringPtrOutput)
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

type SkuResponse struct {
	Name string  `pulumi:"name"`
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

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ControllerConnectionDetailsResponseOutput{})
	pulumi.RegisterOutputType(ControllerConnectionDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(KubernetesConnectionDetailsResponseOutput{})
	pulumi.RegisterOutputType(KubernetesConnectionDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
}
