


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomRPActionRouteDefinition struct {
	Endpoint    string  `pulumi:"endpoint"`
	Name        string  `pulumi:"name"`
	RoutingType *string `pulumi:"routingType"`
}





type CustomRPActionRouteDefinitionInput interface {
	pulumi.Input

	ToCustomRPActionRouteDefinitionOutput() CustomRPActionRouteDefinitionOutput
	ToCustomRPActionRouteDefinitionOutputWithContext(context.Context) CustomRPActionRouteDefinitionOutput
}

type CustomRPActionRouteDefinitionArgs struct {
	Endpoint    pulumi.StringInput    `pulumi:"endpoint"`
	Name        pulumi.StringInput    `pulumi:"name"`
	RoutingType pulumi.StringPtrInput `pulumi:"routingType"`
}

func (CustomRPActionRouteDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRPActionRouteDefinition)(nil)).Elem()
}

func (i CustomRPActionRouteDefinitionArgs) ToCustomRPActionRouteDefinitionOutput() CustomRPActionRouteDefinitionOutput {
	return i.ToCustomRPActionRouteDefinitionOutputWithContext(context.Background())
}

func (i CustomRPActionRouteDefinitionArgs) ToCustomRPActionRouteDefinitionOutputWithContext(ctx context.Context) CustomRPActionRouteDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRPActionRouteDefinitionOutput)
}





type CustomRPActionRouteDefinitionArrayInput interface {
	pulumi.Input

	ToCustomRPActionRouteDefinitionArrayOutput() CustomRPActionRouteDefinitionArrayOutput
	ToCustomRPActionRouteDefinitionArrayOutputWithContext(context.Context) CustomRPActionRouteDefinitionArrayOutput
}

type CustomRPActionRouteDefinitionArray []CustomRPActionRouteDefinitionInput

func (CustomRPActionRouteDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRPActionRouteDefinition)(nil)).Elem()
}

func (i CustomRPActionRouteDefinitionArray) ToCustomRPActionRouteDefinitionArrayOutput() CustomRPActionRouteDefinitionArrayOutput {
	return i.ToCustomRPActionRouteDefinitionArrayOutputWithContext(context.Background())
}

func (i CustomRPActionRouteDefinitionArray) ToCustomRPActionRouteDefinitionArrayOutputWithContext(ctx context.Context) CustomRPActionRouteDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRPActionRouteDefinitionArrayOutput)
}

type CustomRPActionRouteDefinitionOutput struct{ *pulumi.OutputState }

func (CustomRPActionRouteDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRPActionRouteDefinition)(nil)).Elem()
}

func (o CustomRPActionRouteDefinitionOutput) ToCustomRPActionRouteDefinitionOutput() CustomRPActionRouteDefinitionOutput {
	return o
}

func (o CustomRPActionRouteDefinitionOutput) ToCustomRPActionRouteDefinitionOutputWithContext(ctx context.Context) CustomRPActionRouteDefinitionOutput {
	return o
}

func (o CustomRPActionRouteDefinitionOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRPActionRouteDefinition) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o CustomRPActionRouteDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRPActionRouteDefinition) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomRPActionRouteDefinitionOutput) RoutingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRPActionRouteDefinition) *string { return v.RoutingType }).(pulumi.StringPtrOutput)
}

type CustomRPActionRouteDefinitionArrayOutput struct{ *pulumi.OutputState }

func (CustomRPActionRouteDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRPActionRouteDefinition)(nil)).Elem()
}

func (o CustomRPActionRouteDefinitionArrayOutput) ToCustomRPActionRouteDefinitionArrayOutput() CustomRPActionRouteDefinitionArrayOutput {
	return o
}

func (o CustomRPActionRouteDefinitionArrayOutput) ToCustomRPActionRouteDefinitionArrayOutputWithContext(ctx context.Context) CustomRPActionRouteDefinitionArrayOutput {
	return o
}

func (o CustomRPActionRouteDefinitionArrayOutput) Index(i pulumi.IntInput) CustomRPActionRouteDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRPActionRouteDefinition {
		return vs[0].([]CustomRPActionRouteDefinition)[vs[1].(int)]
	}).(CustomRPActionRouteDefinitionOutput)
}

type CustomRPActionRouteDefinitionResponse struct {
	Endpoint    string  `pulumi:"endpoint"`
	Name        string  `pulumi:"name"`
	RoutingType *string `pulumi:"routingType"`
}

type CustomRPActionRouteDefinitionResponseOutput struct{ *pulumi.OutputState }

func (CustomRPActionRouteDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRPActionRouteDefinitionResponse)(nil)).Elem()
}

func (o CustomRPActionRouteDefinitionResponseOutput) ToCustomRPActionRouteDefinitionResponseOutput() CustomRPActionRouteDefinitionResponseOutput {
	return o
}

func (o CustomRPActionRouteDefinitionResponseOutput) ToCustomRPActionRouteDefinitionResponseOutputWithContext(ctx context.Context) CustomRPActionRouteDefinitionResponseOutput {
	return o
}

func (o CustomRPActionRouteDefinitionResponseOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRPActionRouteDefinitionResponse) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o CustomRPActionRouteDefinitionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRPActionRouteDefinitionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomRPActionRouteDefinitionResponseOutput) RoutingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRPActionRouteDefinitionResponse) *string { return v.RoutingType }).(pulumi.StringPtrOutput)
}

type CustomRPActionRouteDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (CustomRPActionRouteDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRPActionRouteDefinitionResponse)(nil)).Elem()
}

func (o CustomRPActionRouteDefinitionResponseArrayOutput) ToCustomRPActionRouteDefinitionResponseArrayOutput() CustomRPActionRouteDefinitionResponseArrayOutput {
	return o
}

func (o CustomRPActionRouteDefinitionResponseArrayOutput) ToCustomRPActionRouteDefinitionResponseArrayOutputWithContext(ctx context.Context) CustomRPActionRouteDefinitionResponseArrayOutput {
	return o
}

func (o CustomRPActionRouteDefinitionResponseArrayOutput) Index(i pulumi.IntInput) CustomRPActionRouteDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRPActionRouteDefinitionResponse {
		return vs[0].([]CustomRPActionRouteDefinitionResponse)[vs[1].(int)]
	}).(CustomRPActionRouteDefinitionResponseOutput)
}

type CustomRPResourceTypeRouteDefinition struct {
	Endpoint    string  `pulumi:"endpoint"`
	Name        string  `pulumi:"name"`
	RoutingType *string `pulumi:"routingType"`
}





type CustomRPResourceTypeRouteDefinitionInput interface {
	pulumi.Input

	ToCustomRPResourceTypeRouteDefinitionOutput() CustomRPResourceTypeRouteDefinitionOutput
	ToCustomRPResourceTypeRouteDefinitionOutputWithContext(context.Context) CustomRPResourceTypeRouteDefinitionOutput
}

type CustomRPResourceTypeRouteDefinitionArgs struct {
	Endpoint    pulumi.StringInput    `pulumi:"endpoint"`
	Name        pulumi.StringInput    `pulumi:"name"`
	RoutingType pulumi.StringPtrInput `pulumi:"routingType"`
}

func (CustomRPResourceTypeRouteDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRPResourceTypeRouteDefinition)(nil)).Elem()
}

func (i CustomRPResourceTypeRouteDefinitionArgs) ToCustomRPResourceTypeRouteDefinitionOutput() CustomRPResourceTypeRouteDefinitionOutput {
	return i.ToCustomRPResourceTypeRouteDefinitionOutputWithContext(context.Background())
}

func (i CustomRPResourceTypeRouteDefinitionArgs) ToCustomRPResourceTypeRouteDefinitionOutputWithContext(ctx context.Context) CustomRPResourceTypeRouteDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRPResourceTypeRouteDefinitionOutput)
}





type CustomRPResourceTypeRouteDefinitionArrayInput interface {
	pulumi.Input

	ToCustomRPResourceTypeRouteDefinitionArrayOutput() CustomRPResourceTypeRouteDefinitionArrayOutput
	ToCustomRPResourceTypeRouteDefinitionArrayOutputWithContext(context.Context) CustomRPResourceTypeRouteDefinitionArrayOutput
}

type CustomRPResourceTypeRouteDefinitionArray []CustomRPResourceTypeRouteDefinitionInput

func (CustomRPResourceTypeRouteDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRPResourceTypeRouteDefinition)(nil)).Elem()
}

func (i CustomRPResourceTypeRouteDefinitionArray) ToCustomRPResourceTypeRouteDefinitionArrayOutput() CustomRPResourceTypeRouteDefinitionArrayOutput {
	return i.ToCustomRPResourceTypeRouteDefinitionArrayOutputWithContext(context.Background())
}

func (i CustomRPResourceTypeRouteDefinitionArray) ToCustomRPResourceTypeRouteDefinitionArrayOutputWithContext(ctx context.Context) CustomRPResourceTypeRouteDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRPResourceTypeRouteDefinitionArrayOutput)
}

type CustomRPResourceTypeRouteDefinitionOutput struct{ *pulumi.OutputState }

func (CustomRPResourceTypeRouteDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRPResourceTypeRouteDefinition)(nil)).Elem()
}

func (o CustomRPResourceTypeRouteDefinitionOutput) ToCustomRPResourceTypeRouteDefinitionOutput() CustomRPResourceTypeRouteDefinitionOutput {
	return o
}

func (o CustomRPResourceTypeRouteDefinitionOutput) ToCustomRPResourceTypeRouteDefinitionOutputWithContext(ctx context.Context) CustomRPResourceTypeRouteDefinitionOutput {
	return o
}

func (o CustomRPResourceTypeRouteDefinitionOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRPResourceTypeRouteDefinition) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o CustomRPResourceTypeRouteDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRPResourceTypeRouteDefinition) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomRPResourceTypeRouteDefinitionOutput) RoutingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRPResourceTypeRouteDefinition) *string { return v.RoutingType }).(pulumi.StringPtrOutput)
}

type CustomRPResourceTypeRouteDefinitionArrayOutput struct{ *pulumi.OutputState }

func (CustomRPResourceTypeRouteDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRPResourceTypeRouteDefinition)(nil)).Elem()
}

func (o CustomRPResourceTypeRouteDefinitionArrayOutput) ToCustomRPResourceTypeRouteDefinitionArrayOutput() CustomRPResourceTypeRouteDefinitionArrayOutput {
	return o
}

func (o CustomRPResourceTypeRouteDefinitionArrayOutput) ToCustomRPResourceTypeRouteDefinitionArrayOutputWithContext(ctx context.Context) CustomRPResourceTypeRouteDefinitionArrayOutput {
	return o
}

func (o CustomRPResourceTypeRouteDefinitionArrayOutput) Index(i pulumi.IntInput) CustomRPResourceTypeRouteDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRPResourceTypeRouteDefinition {
		return vs[0].([]CustomRPResourceTypeRouteDefinition)[vs[1].(int)]
	}).(CustomRPResourceTypeRouteDefinitionOutput)
}

type CustomRPResourceTypeRouteDefinitionResponse struct {
	Endpoint    string  `pulumi:"endpoint"`
	Name        string  `pulumi:"name"`
	RoutingType *string `pulumi:"routingType"`
}

type CustomRPResourceTypeRouteDefinitionResponseOutput struct{ *pulumi.OutputState }

func (CustomRPResourceTypeRouteDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRPResourceTypeRouteDefinitionResponse)(nil)).Elem()
}

func (o CustomRPResourceTypeRouteDefinitionResponseOutput) ToCustomRPResourceTypeRouteDefinitionResponseOutput() CustomRPResourceTypeRouteDefinitionResponseOutput {
	return o
}

func (o CustomRPResourceTypeRouteDefinitionResponseOutput) ToCustomRPResourceTypeRouteDefinitionResponseOutputWithContext(ctx context.Context) CustomRPResourceTypeRouteDefinitionResponseOutput {
	return o
}

func (o CustomRPResourceTypeRouteDefinitionResponseOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRPResourceTypeRouteDefinitionResponse) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o CustomRPResourceTypeRouteDefinitionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRPResourceTypeRouteDefinitionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomRPResourceTypeRouteDefinitionResponseOutput) RoutingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRPResourceTypeRouteDefinitionResponse) *string { return v.RoutingType }).(pulumi.StringPtrOutput)
}

type CustomRPResourceTypeRouteDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (CustomRPResourceTypeRouteDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRPResourceTypeRouteDefinitionResponse)(nil)).Elem()
}

func (o CustomRPResourceTypeRouteDefinitionResponseArrayOutput) ToCustomRPResourceTypeRouteDefinitionResponseArrayOutput() CustomRPResourceTypeRouteDefinitionResponseArrayOutput {
	return o
}

func (o CustomRPResourceTypeRouteDefinitionResponseArrayOutput) ToCustomRPResourceTypeRouteDefinitionResponseArrayOutputWithContext(ctx context.Context) CustomRPResourceTypeRouteDefinitionResponseArrayOutput {
	return o
}

func (o CustomRPResourceTypeRouteDefinitionResponseArrayOutput) Index(i pulumi.IntInput) CustomRPResourceTypeRouteDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRPResourceTypeRouteDefinitionResponse {
		return vs[0].([]CustomRPResourceTypeRouteDefinitionResponse)[vs[1].(int)]
	}).(CustomRPResourceTypeRouteDefinitionResponseOutput)
}

type CustomRPValidations struct {
	Specification  string  `pulumi:"specification"`
	ValidationType *string `pulumi:"validationType"`
}





type CustomRPValidationsInput interface {
	pulumi.Input

	ToCustomRPValidationsOutput() CustomRPValidationsOutput
	ToCustomRPValidationsOutputWithContext(context.Context) CustomRPValidationsOutput
}

type CustomRPValidationsArgs struct {
	Specification  pulumi.StringInput    `pulumi:"specification"`
	ValidationType pulumi.StringPtrInput `pulumi:"validationType"`
}

func (CustomRPValidationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRPValidations)(nil)).Elem()
}

func (i CustomRPValidationsArgs) ToCustomRPValidationsOutput() CustomRPValidationsOutput {
	return i.ToCustomRPValidationsOutputWithContext(context.Background())
}

func (i CustomRPValidationsArgs) ToCustomRPValidationsOutputWithContext(ctx context.Context) CustomRPValidationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRPValidationsOutput)
}





type CustomRPValidationsArrayInput interface {
	pulumi.Input

	ToCustomRPValidationsArrayOutput() CustomRPValidationsArrayOutput
	ToCustomRPValidationsArrayOutputWithContext(context.Context) CustomRPValidationsArrayOutput
}

type CustomRPValidationsArray []CustomRPValidationsInput

func (CustomRPValidationsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRPValidations)(nil)).Elem()
}

func (i CustomRPValidationsArray) ToCustomRPValidationsArrayOutput() CustomRPValidationsArrayOutput {
	return i.ToCustomRPValidationsArrayOutputWithContext(context.Background())
}

func (i CustomRPValidationsArray) ToCustomRPValidationsArrayOutputWithContext(ctx context.Context) CustomRPValidationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRPValidationsArrayOutput)
}

type CustomRPValidationsOutput struct{ *pulumi.OutputState }

func (CustomRPValidationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRPValidations)(nil)).Elem()
}

func (o CustomRPValidationsOutput) ToCustomRPValidationsOutput() CustomRPValidationsOutput {
	return o
}

func (o CustomRPValidationsOutput) ToCustomRPValidationsOutputWithContext(ctx context.Context) CustomRPValidationsOutput {
	return o
}

func (o CustomRPValidationsOutput) Specification() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRPValidations) string { return v.Specification }).(pulumi.StringOutput)
}

func (o CustomRPValidationsOutput) ValidationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRPValidations) *string { return v.ValidationType }).(pulumi.StringPtrOutput)
}

type CustomRPValidationsArrayOutput struct{ *pulumi.OutputState }

func (CustomRPValidationsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRPValidations)(nil)).Elem()
}

func (o CustomRPValidationsArrayOutput) ToCustomRPValidationsArrayOutput() CustomRPValidationsArrayOutput {
	return o
}

func (o CustomRPValidationsArrayOutput) ToCustomRPValidationsArrayOutputWithContext(ctx context.Context) CustomRPValidationsArrayOutput {
	return o
}

func (o CustomRPValidationsArrayOutput) Index(i pulumi.IntInput) CustomRPValidationsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRPValidations {
		return vs[0].([]CustomRPValidations)[vs[1].(int)]
	}).(CustomRPValidationsOutput)
}

type CustomRPValidationsResponse struct {
	Specification  string  `pulumi:"specification"`
	ValidationType *string `pulumi:"validationType"`
}

type CustomRPValidationsResponseOutput struct{ *pulumi.OutputState }

func (CustomRPValidationsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRPValidationsResponse)(nil)).Elem()
}

func (o CustomRPValidationsResponseOutput) ToCustomRPValidationsResponseOutput() CustomRPValidationsResponseOutput {
	return o
}

func (o CustomRPValidationsResponseOutput) ToCustomRPValidationsResponseOutputWithContext(ctx context.Context) CustomRPValidationsResponseOutput {
	return o
}

func (o CustomRPValidationsResponseOutput) Specification() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRPValidationsResponse) string { return v.Specification }).(pulumi.StringOutput)
}

func (o CustomRPValidationsResponseOutput) ValidationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRPValidationsResponse) *string { return v.ValidationType }).(pulumi.StringPtrOutput)
}

type CustomRPValidationsResponseArrayOutput struct{ *pulumi.OutputState }

func (CustomRPValidationsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRPValidationsResponse)(nil)).Elem()
}

func (o CustomRPValidationsResponseArrayOutput) ToCustomRPValidationsResponseArrayOutput() CustomRPValidationsResponseArrayOutput {
	return o
}

func (o CustomRPValidationsResponseArrayOutput) ToCustomRPValidationsResponseArrayOutputWithContext(ctx context.Context) CustomRPValidationsResponseArrayOutput {
	return o
}

func (o CustomRPValidationsResponseArrayOutput) Index(i pulumi.IntInput) CustomRPValidationsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRPValidationsResponse {
		return vs[0].([]CustomRPValidationsResponse)[vs[1].(int)]
	}).(CustomRPValidationsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomRPActionRouteDefinitionOutput{})
	pulumi.RegisterOutputType(CustomRPActionRouteDefinitionArrayOutput{})
	pulumi.RegisterOutputType(CustomRPActionRouteDefinitionResponseOutput{})
	pulumi.RegisterOutputType(CustomRPActionRouteDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(CustomRPResourceTypeRouteDefinitionOutput{})
	pulumi.RegisterOutputType(CustomRPResourceTypeRouteDefinitionArrayOutput{})
	pulumi.RegisterOutputType(CustomRPResourceTypeRouteDefinitionResponseOutput{})
	pulumi.RegisterOutputType(CustomRPResourceTypeRouteDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(CustomRPValidationsOutput{})
	pulumi.RegisterOutputType(CustomRPValidationsArrayOutput{})
	pulumi.RegisterOutputType(CustomRPValidationsResponseOutput{})
	pulumi.RegisterOutputType(CustomRPValidationsResponseArrayOutput{})
}
