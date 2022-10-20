


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink struct {
	Id *string `pulumi:"id"`
}

type RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput struct{ *pulumi.OutputState }

func (RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput() RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput struct{ *pulumi.OutputState }

func (RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ToRoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ToRoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) Elem() RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o.ApplyT(func(v *RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink) RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink {
		if v != nil {
			return *v
		}
		var ret RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink
		return ret
	}).(RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput)
}

func (o RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLink) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink struct {
	Id *string `pulumi:"id"`
}





type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkInput interface {
	pulumi.Input

	ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput
	ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput
}

type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return i.ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(context.Background())
}

func (i RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput)
}

func (i RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput).ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx)
}









type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrInput interface {
	pulumi.Input

	ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput
	ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput
}

type routingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrType RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs

func RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtr(v *RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkArgs) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrInput {
	return (*routingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrType)(v)
}

func (*routingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i *routingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i *routingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput struct{ *pulumi.OutputState }

func (RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink) *RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink {
		return &v
	}).(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput struct{ *pulumi.OutputState }

func (RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ToRoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) Elem() RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o.ApplyT(func(v *RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink) RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink {
		if v != nil {
			return *v
		}
		var ret RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink
		return ret
	}).(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput)
}

func (o RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type RulesEngineAction struct {
	RequestHeaderActions       []HeaderAction `pulumi:"requestHeaderActions"`
	ResponseHeaderActions      []HeaderAction `pulumi:"responseHeaderActions"`
	RouteConfigurationOverride interface{}    `pulumi:"routeConfigurationOverride"`
}





type RulesEngineActionInput interface {
	pulumi.Input

	ToRulesEngineActionOutput() RulesEngineActionOutput
	ToRulesEngineActionOutputWithContext(context.Context) RulesEngineActionOutput
}

type RulesEngineActionArgs struct {
	RequestHeaderActions       HeaderActionArrayInput `pulumi:"requestHeaderActions"`
	ResponseHeaderActions      HeaderActionArrayInput `pulumi:"responseHeaderActions"`
	RouteConfigurationOverride pulumi.Input           `pulumi:"routeConfigurationOverride"`
}

func (RulesEngineActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineAction)(nil)).Elem()
}

func (i RulesEngineActionArgs) ToRulesEngineActionOutput() RulesEngineActionOutput {
	return i.ToRulesEngineActionOutputWithContext(context.Background())
}

func (i RulesEngineActionArgs) ToRulesEngineActionOutputWithContext(ctx context.Context) RulesEngineActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineActionOutput)
}

type RulesEngineActionOutput struct{ *pulumi.OutputState }

func (RulesEngineActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineAction)(nil)).Elem()
}

func (o RulesEngineActionOutput) ToRulesEngineActionOutput() RulesEngineActionOutput {
	return o
}

func (o RulesEngineActionOutput) ToRulesEngineActionOutputWithContext(ctx context.Context) RulesEngineActionOutput {
	return o
}

func (o RulesEngineActionOutput) RequestHeaderActions() HeaderActionArrayOutput {
	return o.ApplyT(func(v RulesEngineAction) []HeaderAction { return v.RequestHeaderActions }).(HeaderActionArrayOutput)
}

func (o RulesEngineActionOutput) ResponseHeaderActions() HeaderActionArrayOutput {
	return o.ApplyT(func(v RulesEngineAction) []HeaderAction { return v.ResponseHeaderActions }).(HeaderActionArrayOutput)
}

func (o RulesEngineActionOutput) RouteConfigurationOverride() pulumi.AnyOutput {
	return o.ApplyT(func(v RulesEngineAction) interface{} { return v.RouteConfigurationOverride }).(pulumi.AnyOutput)
}

type RulesEngineActionResponse struct {
	RequestHeaderActions       []HeaderActionResponse `pulumi:"requestHeaderActions"`
	ResponseHeaderActions      []HeaderActionResponse `pulumi:"responseHeaderActions"`
	RouteConfigurationOverride interface{}            `pulumi:"routeConfigurationOverride"`
}

type RulesEngineActionResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineActionResponse)(nil)).Elem()
}

func (o RulesEngineActionResponseOutput) ToRulesEngineActionResponseOutput() RulesEngineActionResponseOutput {
	return o
}

func (o RulesEngineActionResponseOutput) ToRulesEngineActionResponseOutputWithContext(ctx context.Context) RulesEngineActionResponseOutput {
	return o
}

func (o RulesEngineActionResponseOutput) RequestHeaderActions() HeaderActionResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineActionResponse) []HeaderActionResponse { return v.RequestHeaderActions }).(HeaderActionResponseArrayOutput)
}

func (o RulesEngineActionResponseOutput) ResponseHeaderActions() HeaderActionResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineActionResponse) []HeaderActionResponse { return v.ResponseHeaderActions }).(HeaderActionResponseArrayOutput)
}

func (o RulesEngineActionResponseOutput) RouteConfigurationOverride() pulumi.AnyOutput {
	return o.ApplyT(func(v RulesEngineActionResponse) interface{} { return v.RouteConfigurationOverride }).(pulumi.AnyOutput)
}

type RulesEngineMatchCondition struct {
	NegateCondition          *bool    `pulumi:"negateCondition"`
	RulesEngineMatchValue    []string `pulumi:"rulesEngineMatchValue"`
	RulesEngineMatchVariable string   `pulumi:"rulesEngineMatchVariable"`
	RulesEngineOperator      string   `pulumi:"rulesEngineOperator"`
	Selector                 *string  `pulumi:"selector"`
	Transforms               []string `pulumi:"transforms"`
}





type RulesEngineMatchConditionInput interface {
	pulumi.Input

	ToRulesEngineMatchConditionOutput() RulesEngineMatchConditionOutput
	ToRulesEngineMatchConditionOutputWithContext(context.Context) RulesEngineMatchConditionOutput
}

type RulesEngineMatchConditionArgs struct {
	NegateCondition          pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	RulesEngineMatchValue    pulumi.StringArrayInput `pulumi:"rulesEngineMatchValue"`
	RulesEngineMatchVariable pulumi.StringInput      `pulumi:"rulesEngineMatchVariable"`
	RulesEngineOperator      pulumi.StringInput      `pulumi:"rulesEngineOperator"`
	Selector                 pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms               pulumi.StringArrayInput `pulumi:"transforms"`
}

func (RulesEngineMatchConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineMatchCondition)(nil)).Elem()
}

func (i RulesEngineMatchConditionArgs) ToRulesEngineMatchConditionOutput() RulesEngineMatchConditionOutput {
	return i.ToRulesEngineMatchConditionOutputWithContext(context.Background())
}

func (i RulesEngineMatchConditionArgs) ToRulesEngineMatchConditionOutputWithContext(ctx context.Context) RulesEngineMatchConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineMatchConditionOutput)
}





type RulesEngineMatchConditionArrayInput interface {
	pulumi.Input

	ToRulesEngineMatchConditionArrayOutput() RulesEngineMatchConditionArrayOutput
	ToRulesEngineMatchConditionArrayOutputWithContext(context.Context) RulesEngineMatchConditionArrayOutput
}

type RulesEngineMatchConditionArray []RulesEngineMatchConditionInput

func (RulesEngineMatchConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineMatchCondition)(nil)).Elem()
}

func (i RulesEngineMatchConditionArray) ToRulesEngineMatchConditionArrayOutput() RulesEngineMatchConditionArrayOutput {
	return i.ToRulesEngineMatchConditionArrayOutputWithContext(context.Background())
}

func (i RulesEngineMatchConditionArray) ToRulesEngineMatchConditionArrayOutputWithContext(ctx context.Context) RulesEngineMatchConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineMatchConditionArrayOutput)
}

type RulesEngineMatchConditionOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineMatchCondition)(nil)).Elem()
}

func (o RulesEngineMatchConditionOutput) ToRulesEngineMatchConditionOutput() RulesEngineMatchConditionOutput {
	return o
}

func (o RulesEngineMatchConditionOutput) ToRulesEngineMatchConditionOutputWithContext(ctx context.Context) RulesEngineMatchConditionOutput {
	return o
}

func (o RulesEngineMatchConditionOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RulesEngineMatchConditionOutput) RulesEngineMatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) []string { return v.RulesEngineMatchValue }).(pulumi.StringArrayOutput)
}

func (o RulesEngineMatchConditionOutput) RulesEngineMatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) string { return v.RulesEngineMatchVariable }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionOutput) RulesEngineOperator() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) string { return v.RulesEngineOperator }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o RulesEngineMatchConditionOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RulesEngineMatchConditionArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineMatchCondition)(nil)).Elem()
}

func (o RulesEngineMatchConditionArrayOutput) ToRulesEngineMatchConditionArrayOutput() RulesEngineMatchConditionArrayOutput {
	return o
}

func (o RulesEngineMatchConditionArrayOutput) ToRulesEngineMatchConditionArrayOutputWithContext(ctx context.Context) RulesEngineMatchConditionArrayOutput {
	return o
}

func (o RulesEngineMatchConditionArrayOutput) Index(i pulumi.IntInput) RulesEngineMatchConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineMatchCondition {
		return vs[0].([]RulesEngineMatchCondition)[vs[1].(int)]
	}).(RulesEngineMatchConditionOutput)
}

type RulesEngineMatchConditionResponse struct {
	NegateCondition          *bool    `pulumi:"negateCondition"`
	RulesEngineMatchValue    []string `pulumi:"rulesEngineMatchValue"`
	RulesEngineMatchVariable string   `pulumi:"rulesEngineMatchVariable"`
	RulesEngineOperator      string   `pulumi:"rulesEngineOperator"`
	Selector                 *string  `pulumi:"selector"`
	Transforms               []string `pulumi:"transforms"`
}

type RulesEngineMatchConditionResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineMatchConditionResponse)(nil)).Elem()
}

func (o RulesEngineMatchConditionResponseOutput) ToRulesEngineMatchConditionResponseOutput() RulesEngineMatchConditionResponseOutput {
	return o
}

func (o RulesEngineMatchConditionResponseOutput) ToRulesEngineMatchConditionResponseOutputWithContext(ctx context.Context) RulesEngineMatchConditionResponseOutput {
	return o
}

func (o RulesEngineMatchConditionResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RulesEngineMatchConditionResponseOutput) RulesEngineMatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) []string { return v.RulesEngineMatchValue }).(pulumi.StringArrayOutput)
}

func (o RulesEngineMatchConditionResponseOutput) RulesEngineMatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) string { return v.RulesEngineMatchVariable }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionResponseOutput) RulesEngineOperator() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) string { return v.RulesEngineOperator }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionResponseOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o RulesEngineMatchConditionResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RulesEngineMatchConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineMatchConditionResponse)(nil)).Elem()
}

func (o RulesEngineMatchConditionResponseArrayOutput) ToRulesEngineMatchConditionResponseArrayOutput() RulesEngineMatchConditionResponseArrayOutput {
	return o
}

func (o RulesEngineMatchConditionResponseArrayOutput) ToRulesEngineMatchConditionResponseArrayOutputWithContext(ctx context.Context) RulesEngineMatchConditionResponseArrayOutput {
	return o
}

func (o RulesEngineMatchConditionResponseArrayOutput) Index(i pulumi.IntInput) RulesEngineMatchConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineMatchConditionResponse {
		return vs[0].([]RulesEngineMatchConditionResponse)[vs[1].(int)]
	}).(RulesEngineMatchConditionResponseOutput)
}

type RulesEngineResponse struct {
	Id            string                    `pulumi:"id"`
	Name          string                    `pulumi:"name"`
	ResourceState string                    `pulumi:"resourceState"`
	Rules         []RulesEngineRuleResponse `pulumi:"rules"`
	Type          string                    `pulumi:"type"`
}

type RulesEngineResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineResponse)(nil)).Elem()
}

func (o RulesEngineResponseOutput) ToRulesEngineResponseOutput() RulesEngineResponseOutput {
	return o
}

func (o RulesEngineResponseOutput) ToRulesEngineResponseOutputWithContext(ctx context.Context) RulesEngineResponseOutput {
	return o
}

func (o RulesEngineResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o RulesEngineResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RulesEngineResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o RulesEngineResponseOutput) Rules() RulesEngineRuleResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineResponse) []RulesEngineRuleResponse { return v.Rules }).(RulesEngineRuleResponseArrayOutput)
}

func (o RulesEngineResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RulesEngineResponseArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineResponse)(nil)).Elem()
}

func (o RulesEngineResponseArrayOutput) ToRulesEngineResponseArrayOutput() RulesEngineResponseArrayOutput {
	return o
}

func (o RulesEngineResponseArrayOutput) ToRulesEngineResponseArrayOutputWithContext(ctx context.Context) RulesEngineResponseArrayOutput {
	return o
}

func (o RulesEngineResponseArrayOutput) Index(i pulumi.IntInput) RulesEngineResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineResponse {
		return vs[0].([]RulesEngineResponse)[vs[1].(int)]
	}).(RulesEngineResponseOutput)
}

type RulesEngineRule struct {
	Action                  RulesEngineAction           `pulumi:"action"`
	MatchConditions         []RulesEngineMatchCondition `pulumi:"matchConditions"`
	MatchProcessingBehavior *string                     `pulumi:"matchProcessingBehavior"`
	Name                    string                      `pulumi:"name"`
	Priority                int                         `pulumi:"priority"`
}





type RulesEngineRuleInput interface {
	pulumi.Input

	ToRulesEngineRuleOutput() RulesEngineRuleOutput
	ToRulesEngineRuleOutputWithContext(context.Context) RulesEngineRuleOutput
}

type RulesEngineRuleArgs struct {
	Action                  RulesEngineActionInput              `pulumi:"action"`
	MatchConditions         RulesEngineMatchConditionArrayInput `pulumi:"matchConditions"`
	MatchProcessingBehavior pulumi.StringPtrInput               `pulumi:"matchProcessingBehavior"`
	Name                    pulumi.StringInput                  `pulumi:"name"`
	Priority                pulumi.IntInput                     `pulumi:"priority"`
}

func (RulesEngineRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineRule)(nil)).Elem()
}

func (i RulesEngineRuleArgs) ToRulesEngineRuleOutput() RulesEngineRuleOutput {
	return i.ToRulesEngineRuleOutputWithContext(context.Background())
}

func (i RulesEngineRuleArgs) ToRulesEngineRuleOutputWithContext(ctx context.Context) RulesEngineRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineRuleOutput)
}





type RulesEngineRuleArrayInput interface {
	pulumi.Input

	ToRulesEngineRuleArrayOutput() RulesEngineRuleArrayOutput
	ToRulesEngineRuleArrayOutputWithContext(context.Context) RulesEngineRuleArrayOutput
}

type RulesEngineRuleArray []RulesEngineRuleInput

func (RulesEngineRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineRule)(nil)).Elem()
}

func (i RulesEngineRuleArray) ToRulesEngineRuleArrayOutput() RulesEngineRuleArrayOutput {
	return i.ToRulesEngineRuleArrayOutputWithContext(context.Background())
}

func (i RulesEngineRuleArray) ToRulesEngineRuleArrayOutputWithContext(ctx context.Context) RulesEngineRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineRuleArrayOutput)
}

type RulesEngineRuleOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineRule)(nil)).Elem()
}

func (o RulesEngineRuleOutput) ToRulesEngineRuleOutput() RulesEngineRuleOutput {
	return o
}

func (o RulesEngineRuleOutput) ToRulesEngineRuleOutputWithContext(ctx context.Context) RulesEngineRuleOutput {
	return o
}

func (o RulesEngineRuleOutput) Action() RulesEngineActionOutput {
	return o.ApplyT(func(v RulesEngineRule) RulesEngineAction { return v.Action }).(RulesEngineActionOutput)
}

func (o RulesEngineRuleOutput) MatchConditions() RulesEngineMatchConditionArrayOutput {
	return o.ApplyT(func(v RulesEngineRule) []RulesEngineMatchCondition { return v.MatchConditions }).(RulesEngineMatchConditionArrayOutput)
}

func (o RulesEngineRuleOutput) MatchProcessingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineRule) *string { return v.MatchProcessingBehavior }).(pulumi.StringPtrOutput)
}

func (o RulesEngineRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineRule) string { return v.Name }).(pulumi.StringOutput)
}

func (o RulesEngineRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v RulesEngineRule) int { return v.Priority }).(pulumi.IntOutput)
}

type RulesEngineRuleArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineRule)(nil)).Elem()
}

func (o RulesEngineRuleArrayOutput) ToRulesEngineRuleArrayOutput() RulesEngineRuleArrayOutput {
	return o
}

func (o RulesEngineRuleArrayOutput) ToRulesEngineRuleArrayOutputWithContext(ctx context.Context) RulesEngineRuleArrayOutput {
	return o
}

func (o RulesEngineRuleArrayOutput) Index(i pulumi.IntInput) RulesEngineRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineRule {
		return vs[0].([]RulesEngineRule)[vs[1].(int)]
	}).(RulesEngineRuleOutput)
}

type RulesEngineRuleResponse struct {
	Action                  RulesEngineActionResponse           `pulumi:"action"`
	MatchConditions         []RulesEngineMatchConditionResponse `pulumi:"matchConditions"`
	MatchProcessingBehavior *string                             `pulumi:"matchProcessingBehavior"`
	Name                    string                              `pulumi:"name"`
	Priority                int                                 `pulumi:"priority"`
}

type RulesEngineRuleResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineRuleResponse)(nil)).Elem()
}

func (o RulesEngineRuleResponseOutput) ToRulesEngineRuleResponseOutput() RulesEngineRuleResponseOutput {
	return o
}

func (o RulesEngineRuleResponseOutput) ToRulesEngineRuleResponseOutputWithContext(ctx context.Context) RulesEngineRuleResponseOutput {
	return o
}

func (o RulesEngineRuleResponseOutput) Action() RulesEngineActionResponseOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) RulesEngineActionResponse { return v.Action }).(RulesEngineActionResponseOutput)
}

func (o RulesEngineRuleResponseOutput) MatchConditions() RulesEngineMatchConditionResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) []RulesEngineMatchConditionResponse { return v.MatchConditions }).(RulesEngineMatchConditionResponseArrayOutput)
}

func (o RulesEngineRuleResponseOutput) MatchProcessingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) *string { return v.MatchProcessingBehavior }).(pulumi.StringPtrOutput)
}

func (o RulesEngineRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RulesEngineRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

type RulesEngineRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineRuleResponse)(nil)).Elem()
}

func (o RulesEngineRuleResponseArrayOutput) ToRulesEngineRuleResponseArrayOutput() RulesEngineRuleResponseArrayOutput {
	return o
}

func (o RulesEngineRuleResponseArrayOutput) ToRulesEngineRuleResponseArrayOutputWithContext(ctx context.Context) RulesEngineRuleResponseArrayOutput {
	return o
}

func (o RulesEngineRuleResponseArrayOutput) Index(i pulumi.IntInput) RulesEngineRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineRuleResponse {
		return vs[0].([]RulesEngineRuleResponse)[vs[1].(int)]
	}).(RulesEngineRuleResponseOutput)
}

type SecurityPolicyLinkResponse struct {
	Id *string `pulumi:"id"`
}

type SecurityPolicyLinkResponseOutput struct{ *pulumi.OutputState }

func (SecurityPolicyLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPolicyLinkResponse)(nil)).Elem()
}

func (o SecurityPolicyLinkResponseOutput) ToSecurityPolicyLinkResponseOutput() SecurityPolicyLinkResponseOutput {
	return o
}

func (o SecurityPolicyLinkResponseOutput) ToSecurityPolicyLinkResponseOutputWithContext(ctx context.Context) SecurityPolicyLinkResponseOutput {
	return o
}

func (o SecurityPolicyLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityPolicyLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SecurityPolicyLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (SecurityPolicyLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityPolicyLinkResponse)(nil)).Elem()
}

func (o SecurityPolicyLinkResponseArrayOutput) ToSecurityPolicyLinkResponseArrayOutput() SecurityPolicyLinkResponseArrayOutput {
	return o
}

func (o SecurityPolicyLinkResponseArrayOutput) ToSecurityPolicyLinkResponseArrayOutputWithContext(ctx context.Context) SecurityPolicyLinkResponseArrayOutput {
	return o
}

func (o SecurityPolicyLinkResponseArrayOutput) Index(i pulumi.IntInput) SecurityPolicyLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityPolicyLinkResponse {
		return vs[0].([]SecurityPolicyLinkResponse)[vs[1].(int)]
	}).(SecurityPolicyLinkResponseOutput)
}

type SecurityRuleType struct {
	Access                               string                         `pulumi:"access"`
	Description                          *string                        `pulumi:"description"`
	DestinationAddressPrefix             *string                        `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes           []string                       `pulumi:"destinationAddressPrefixes"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupType `pulumi:"destinationApplicationSecurityGroups"`
	DestinationPortRange                 *string                        `pulumi:"destinationPortRange"`
	DestinationPortRanges                []string                       `pulumi:"destinationPortRanges"`
	Direction                            string                         `pulumi:"direction"`
	Id                                   *string                        `pulumi:"id"`
	Name                                 *string                        `pulumi:"name"`
	Priority                             *int                           `pulumi:"priority"`
	Protocol                             string                         `pulumi:"protocol"`
	SourceAddressPrefix                  *string                        `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes                []string                       `pulumi:"sourceAddressPrefixes"`
	SourceApplicationSecurityGroups      []ApplicationSecurityGroupType `pulumi:"sourceApplicationSecurityGroups"`
	SourcePortRange                      *string                        `pulumi:"sourcePortRange"`
	SourcePortRanges                     []string                       `pulumi:"sourcePortRanges"`
	Type                                 *string                        `pulumi:"type"`
}





type SecurityRuleTypeInput interface {
	pulumi.Input

	ToSecurityRuleTypeOutput() SecurityRuleTypeOutput
	ToSecurityRuleTypeOutputWithContext(context.Context) SecurityRuleTypeOutput
}

type SecurityRuleTypeArgs struct {
	Access                               pulumi.StringInput                     `pulumi:"access"`
	Description                          pulumi.StringPtrInput                  `pulumi:"description"`
	DestinationAddressPrefix             pulumi.StringPtrInput                  `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes           pulumi.StringArrayInput                `pulumi:"destinationAddressPrefixes"`
	DestinationApplicationSecurityGroups ApplicationSecurityGroupTypeArrayInput `pulumi:"destinationApplicationSecurityGroups"`
	DestinationPortRange                 pulumi.StringPtrInput                  `pulumi:"destinationPortRange"`
	DestinationPortRanges                pulumi.StringArrayInput                `pulumi:"destinationPortRanges"`
	Direction                            pulumi.StringInput                     `pulumi:"direction"`
	Id                                   pulumi.StringPtrInput                  `pulumi:"id"`
	Name                                 pulumi.StringPtrInput                  `pulumi:"name"`
	Priority                             pulumi.IntPtrInput                     `pulumi:"priority"`
	Protocol                             pulumi.StringInput                     `pulumi:"protocol"`
	SourceAddressPrefix                  pulumi.StringPtrInput                  `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes                pulumi.StringArrayInput                `pulumi:"sourceAddressPrefixes"`
	SourceApplicationSecurityGroups      ApplicationSecurityGroupTypeArrayInput `pulumi:"sourceApplicationSecurityGroups"`
	SourcePortRange                      pulumi.StringPtrInput                  `pulumi:"sourcePortRange"`
	SourcePortRanges                     pulumi.StringArrayInput                `pulumi:"sourcePortRanges"`
	Type                                 pulumi.StringPtrInput                  `pulumi:"type"`
}

func (SecurityRuleTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleType)(nil)).Elem()
}

func (i SecurityRuleTypeArgs) ToSecurityRuleTypeOutput() SecurityRuleTypeOutput {
	return i.ToSecurityRuleTypeOutputWithContext(context.Background())
}

func (i SecurityRuleTypeArgs) ToSecurityRuleTypeOutputWithContext(ctx context.Context) SecurityRuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityRuleTypeOutput)
}





type SecurityRuleTypeArrayInput interface {
	pulumi.Input

	ToSecurityRuleTypeArrayOutput() SecurityRuleTypeArrayOutput
	ToSecurityRuleTypeArrayOutputWithContext(context.Context) SecurityRuleTypeArrayOutput
}

type SecurityRuleTypeArray []SecurityRuleTypeInput

func (SecurityRuleTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityRuleType)(nil)).Elem()
}

func (i SecurityRuleTypeArray) ToSecurityRuleTypeArrayOutput() SecurityRuleTypeArrayOutput {
	return i.ToSecurityRuleTypeArrayOutputWithContext(context.Background())
}

func (i SecurityRuleTypeArray) ToSecurityRuleTypeArrayOutputWithContext(ctx context.Context) SecurityRuleTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityRuleTypeArrayOutput)
}

type SecurityRuleTypeOutput struct{ *pulumi.OutputState }

func (SecurityRuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleType)(nil)).Elem()
}

func (o SecurityRuleTypeOutput) ToSecurityRuleTypeOutput() SecurityRuleTypeOutput {
	return o
}

func (o SecurityRuleTypeOutput) ToSecurityRuleTypeOutputWithContext(ctx context.Context) SecurityRuleTypeOutput {
	return o
}

func (o SecurityRuleTypeOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.Access }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) DestinationAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []string { return v.DestinationAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleTypeOutput) DestinationApplicationSecurityGroups() ApplicationSecurityGroupTypeArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []ApplicationSecurityGroupType { return v.DestinationApplicationSecurityGroups }).(ApplicationSecurityGroupTypeArrayOutput)
}

func (o SecurityRuleTypeOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleTypeOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.Direction }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SecurityRuleTypeOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) SourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []string { return v.SourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleTypeOutput) SourceApplicationSecurityGroups() ApplicationSecurityGroupTypeArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []ApplicationSecurityGroupType { return v.SourceApplicationSecurityGroups }).(ApplicationSecurityGroupTypeArrayOutput)
}

func (o SecurityRuleTypeOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleType) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type SecurityRuleTypeArrayOutput struct{ *pulumi.OutputState }

func (SecurityRuleTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityRuleType)(nil)).Elem()
}

func (o SecurityRuleTypeArrayOutput) ToSecurityRuleTypeArrayOutput() SecurityRuleTypeArrayOutput {
	return o
}

func (o SecurityRuleTypeArrayOutput) ToSecurityRuleTypeArrayOutputWithContext(ctx context.Context) SecurityRuleTypeArrayOutput {
	return o
}

func (o SecurityRuleTypeArrayOutput) Index(i pulumi.IntInput) SecurityRuleTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityRuleType {
		return vs[0].([]SecurityRuleType)[vs[1].(int)]
	}).(SecurityRuleTypeOutput)
}

type SecurityRuleResponse struct {
	Access                               string                             `pulumi:"access"`
	Description                          *string                            `pulumi:"description"`
	DestinationAddressPrefix             *string                            `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes           []string                           `pulumi:"destinationAddressPrefixes"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupResponse `pulumi:"destinationApplicationSecurityGroups"`
	DestinationPortRange                 *string                            `pulumi:"destinationPortRange"`
	DestinationPortRanges                []string                           `pulumi:"destinationPortRanges"`
	Direction                            string                             `pulumi:"direction"`
	Etag                                 string                             `pulumi:"etag"`
	Id                                   *string                            `pulumi:"id"`
	Name                                 *string                            `pulumi:"name"`
	Priority                             *int                               `pulumi:"priority"`
	Protocol                             string                             `pulumi:"protocol"`
	ProvisioningState                    string                             `pulumi:"provisioningState"`
	SourceAddressPrefix                  *string                            `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes                []string                           `pulumi:"sourceAddressPrefixes"`
	SourceApplicationSecurityGroups      []ApplicationSecurityGroupResponse `pulumi:"sourceApplicationSecurityGroups"`
	SourcePortRange                      *string                            `pulumi:"sourcePortRange"`
	SourcePortRanges                     []string                           `pulumi:"sourcePortRanges"`
	Type                                 *string                            `pulumi:"type"`
}

type SecurityRuleResponseOutput struct{ *pulumi.OutputState }

func (SecurityRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleResponse)(nil)).Elem()
}

func (o SecurityRuleResponseOutput) ToSecurityRuleResponseOutput() SecurityRuleResponseOutput {
	return o
}

func (o SecurityRuleResponseOutput) ToSecurityRuleResponseOutputWithContext(ctx context.Context) SecurityRuleResponseOutput {
	return o
}

func (o SecurityRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) DestinationAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []string { return v.DestinationAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleResponseOutput) DestinationApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []ApplicationSecurityGroupResponse {
		return v.DestinationApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

func (o SecurityRuleResponseOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SecurityRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) SourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []string { return v.SourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleResponseOutput) SourceApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []ApplicationSecurityGroupResponse {
		return v.SourceApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

func (o SecurityRuleResponseOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityRuleResponse) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func (o SecurityRuleResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type SecurityRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (SecurityRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityRuleResponse)(nil)).Elem()
}

func (o SecurityRuleResponseArrayOutput) ToSecurityRuleResponseArrayOutput() SecurityRuleResponseArrayOutput {
	return o
}

func (o SecurityRuleResponseArrayOutput) ToSecurityRuleResponseArrayOutputWithContext(ctx context.Context) SecurityRuleResponseArrayOutput {
	return o
}

func (o SecurityRuleResponseArrayOutput) Index(i pulumi.IntInput) SecurityRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityRuleResponse {
		return vs[0].([]SecurityRuleResponse)[vs[1].(int)]
	}).(SecurityRuleResponseOutput)
}

type ServiceAssociationLinkResponse struct {
	AllowDelete        *bool    `pulumi:"allowDelete"`
	Etag               string   `pulumi:"etag"`
	Id                 *string  `pulumi:"id"`
	Link               *string  `pulumi:"link"`
	LinkedResourceType *string  `pulumi:"linkedResourceType"`
	Locations          []string `pulumi:"locations"`
	Name               *string  `pulumi:"name"`
	ProvisioningState  string   `pulumi:"provisioningState"`
	Type               string   `pulumi:"type"`
}

type ServiceAssociationLinkResponseOutput struct{ *pulumi.OutputState }

func (ServiceAssociationLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAssociationLinkResponse)(nil)).Elem()
}

func (o ServiceAssociationLinkResponseOutput) ToServiceAssociationLinkResponseOutput() ServiceAssociationLinkResponseOutput {
	return o
}

func (o ServiceAssociationLinkResponseOutput) ToServiceAssociationLinkResponseOutputWithContext(ctx context.Context) ServiceAssociationLinkResponseOutput {
	return o
}

func (o ServiceAssociationLinkResponseOutput) AllowDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceAssociationLinkResponse) *bool { return v.AllowDelete }).(pulumi.BoolPtrOutput)
}

func (o ServiceAssociationLinkResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAssociationLinkResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ServiceAssociationLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAssociationLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ServiceAssociationLinkResponseOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAssociationLinkResponse) *string { return v.Link }).(pulumi.StringPtrOutput)
}

func (o ServiceAssociationLinkResponseOutput) LinkedResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAssociationLinkResponse) *string { return v.LinkedResourceType }).(pulumi.StringPtrOutput)
}

func (o ServiceAssociationLinkResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceAssociationLinkResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o ServiceAssociationLinkResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAssociationLinkResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceAssociationLinkResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAssociationLinkResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceAssociationLinkResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAssociationLinkResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceAssociationLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceAssociationLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAssociationLinkResponse)(nil)).Elem()
}

func (o ServiceAssociationLinkResponseArrayOutput) ToServiceAssociationLinkResponseArrayOutput() ServiceAssociationLinkResponseArrayOutput {
	return o
}

func (o ServiceAssociationLinkResponseArrayOutput) ToServiceAssociationLinkResponseArrayOutputWithContext(ctx context.Context) ServiceAssociationLinkResponseArrayOutput {
	return o
}

func (o ServiceAssociationLinkResponseArrayOutput) Index(i pulumi.IntInput) ServiceAssociationLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceAssociationLinkResponse {
		return vs[0].([]ServiceAssociationLinkResponse)[vs[1].(int)]
	}).(ServiceAssociationLinkResponseOutput)
}

type ServiceEndpointPolicyType struct {
	Id                               *string                               `pulumi:"id"`
	Location                         *string                               `pulumi:"location"`
	ServiceEndpointPolicyDefinitions []ServiceEndpointPolicyDefinitionType `pulumi:"serviceEndpointPolicyDefinitions"`
	Tags                             map[string]string                     `pulumi:"tags"`
}





type ServiceEndpointPolicyTypeInput interface {
	pulumi.Input

	ToServiceEndpointPolicyTypeOutput() ServiceEndpointPolicyTypeOutput
	ToServiceEndpointPolicyTypeOutputWithContext(context.Context) ServiceEndpointPolicyTypeOutput
}

type ServiceEndpointPolicyTypeArgs struct {
	Id                               pulumi.StringPtrInput                         `pulumi:"id"`
	Location                         pulumi.StringPtrInput                         `pulumi:"location"`
	ServiceEndpointPolicyDefinitions ServiceEndpointPolicyDefinitionTypeArrayInput `pulumi:"serviceEndpointPolicyDefinitions"`
	Tags                             pulumi.StringMapInput                         `pulumi:"tags"`
}

func (ServiceEndpointPolicyTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPolicyType)(nil)).Elem()
}

func (i ServiceEndpointPolicyTypeArgs) ToServiceEndpointPolicyTypeOutput() ServiceEndpointPolicyTypeOutput {
	return i.ToServiceEndpointPolicyTypeOutputWithContext(context.Background())
}

func (i ServiceEndpointPolicyTypeArgs) ToServiceEndpointPolicyTypeOutputWithContext(ctx context.Context) ServiceEndpointPolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPolicyTypeOutput)
}





type ServiceEndpointPolicyTypeArrayInput interface {
	pulumi.Input

	ToServiceEndpointPolicyTypeArrayOutput() ServiceEndpointPolicyTypeArrayOutput
	ToServiceEndpointPolicyTypeArrayOutputWithContext(context.Context) ServiceEndpointPolicyTypeArrayOutput
}

type ServiceEndpointPolicyTypeArray []ServiceEndpointPolicyTypeInput

func (ServiceEndpointPolicyTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPolicyType)(nil)).Elem()
}

func (i ServiceEndpointPolicyTypeArray) ToServiceEndpointPolicyTypeArrayOutput() ServiceEndpointPolicyTypeArrayOutput {
	return i.ToServiceEndpointPolicyTypeArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointPolicyTypeArray) ToServiceEndpointPolicyTypeArrayOutputWithContext(ctx context.Context) ServiceEndpointPolicyTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPolicyTypeArrayOutput)
}

type ServiceEndpointPolicyTypeOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPolicyType)(nil)).Elem()
}

func (o ServiceEndpointPolicyTypeOutput) ToServiceEndpointPolicyTypeOutput() ServiceEndpointPolicyTypeOutput {
	return o
}

func (o ServiceEndpointPolicyTypeOutput) ToServiceEndpointPolicyTypeOutputWithContext(ctx context.Context) ServiceEndpointPolicyTypeOutput {
	return o
}

func (o ServiceEndpointPolicyTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyTypeOutput) ServiceEndpointPolicyDefinitions() ServiceEndpointPolicyDefinitionTypeArrayOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyType) []ServiceEndpointPolicyDefinitionType {
		return v.ServiceEndpointPolicyDefinitions
	}).(ServiceEndpointPolicyDefinitionTypeArrayOutput)
}

func (o ServiceEndpointPolicyTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ServiceEndpointPolicyTypeArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPolicyType)(nil)).Elem()
}

func (o ServiceEndpointPolicyTypeArrayOutput) ToServiceEndpointPolicyTypeArrayOutput() ServiceEndpointPolicyTypeArrayOutput {
	return o
}

func (o ServiceEndpointPolicyTypeArrayOutput) ToServiceEndpointPolicyTypeArrayOutputWithContext(ctx context.Context) ServiceEndpointPolicyTypeArrayOutput {
	return o
}

func (o ServiceEndpointPolicyTypeArrayOutput) Index(i pulumi.IntInput) ServiceEndpointPolicyTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointPolicyType {
		return vs[0].([]ServiceEndpointPolicyType)[vs[1].(int)]
	}).(ServiceEndpointPolicyTypeOutput)
}

type ServiceEndpointPolicyDefinitionType struct {
	Description      *string  `pulumi:"description"`
	Id               *string  `pulumi:"id"`
	Name             *string  `pulumi:"name"`
	Service          *string  `pulumi:"service"`
	ServiceResources []string `pulumi:"serviceResources"`
}





type ServiceEndpointPolicyDefinitionTypeInput interface {
	pulumi.Input

	ToServiceEndpointPolicyDefinitionTypeOutput() ServiceEndpointPolicyDefinitionTypeOutput
	ToServiceEndpointPolicyDefinitionTypeOutputWithContext(context.Context) ServiceEndpointPolicyDefinitionTypeOutput
}

type ServiceEndpointPolicyDefinitionTypeArgs struct {
	Description      pulumi.StringPtrInput   `pulumi:"description"`
	Id               pulumi.StringPtrInput   `pulumi:"id"`
	Name             pulumi.StringPtrInput   `pulumi:"name"`
	Service          pulumi.StringPtrInput   `pulumi:"service"`
	ServiceResources pulumi.StringArrayInput `pulumi:"serviceResources"`
}

func (ServiceEndpointPolicyDefinitionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPolicyDefinitionType)(nil)).Elem()
}

func (i ServiceEndpointPolicyDefinitionTypeArgs) ToServiceEndpointPolicyDefinitionTypeOutput() ServiceEndpointPolicyDefinitionTypeOutput {
	return i.ToServiceEndpointPolicyDefinitionTypeOutputWithContext(context.Background())
}

func (i ServiceEndpointPolicyDefinitionTypeArgs) ToServiceEndpointPolicyDefinitionTypeOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPolicyDefinitionTypeOutput)
}





type ServiceEndpointPolicyDefinitionTypeArrayInput interface {
	pulumi.Input

	ToServiceEndpointPolicyDefinitionTypeArrayOutput() ServiceEndpointPolicyDefinitionTypeArrayOutput
	ToServiceEndpointPolicyDefinitionTypeArrayOutputWithContext(context.Context) ServiceEndpointPolicyDefinitionTypeArrayOutput
}

type ServiceEndpointPolicyDefinitionTypeArray []ServiceEndpointPolicyDefinitionTypeInput

func (ServiceEndpointPolicyDefinitionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPolicyDefinitionType)(nil)).Elem()
}

func (i ServiceEndpointPolicyDefinitionTypeArray) ToServiceEndpointPolicyDefinitionTypeArrayOutput() ServiceEndpointPolicyDefinitionTypeArrayOutput {
	return i.ToServiceEndpointPolicyDefinitionTypeArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointPolicyDefinitionTypeArray) ToServiceEndpointPolicyDefinitionTypeArrayOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPolicyDefinitionTypeArrayOutput)
}

type ServiceEndpointPolicyDefinitionTypeOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyDefinitionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPolicyDefinitionType)(nil)).Elem()
}

func (o ServiceEndpointPolicyDefinitionTypeOutput) ToServiceEndpointPolicyDefinitionTypeOutput() ServiceEndpointPolicyDefinitionTypeOutput {
	return o
}

func (o ServiceEndpointPolicyDefinitionTypeOutput) ToServiceEndpointPolicyDefinitionTypeOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionTypeOutput {
	return o
}

func (o ServiceEndpointPolicyDefinitionTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyDefinitionTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyDefinitionTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyDefinitionTypeOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionType) *string { return v.Service }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyDefinitionTypeOutput) ServiceResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionType) []string { return v.ServiceResources }).(pulumi.StringArrayOutput)
}

type ServiceEndpointPolicyDefinitionTypeArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyDefinitionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPolicyDefinitionType)(nil)).Elem()
}

func (o ServiceEndpointPolicyDefinitionTypeArrayOutput) ToServiceEndpointPolicyDefinitionTypeArrayOutput() ServiceEndpointPolicyDefinitionTypeArrayOutput {
	return o
}

func (o ServiceEndpointPolicyDefinitionTypeArrayOutput) ToServiceEndpointPolicyDefinitionTypeArrayOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionTypeArrayOutput {
	return o
}

func (o ServiceEndpointPolicyDefinitionTypeArrayOutput) Index(i pulumi.IntInput) ServiceEndpointPolicyDefinitionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointPolicyDefinitionType {
		return vs[0].([]ServiceEndpointPolicyDefinitionType)[vs[1].(int)]
	}).(ServiceEndpointPolicyDefinitionTypeOutput)
}

type ServiceEndpointPolicyDefinitionResponse struct {
	Description       *string  `pulumi:"description"`
	Etag              string   `pulumi:"etag"`
	Id                *string  `pulumi:"id"`
	Name              *string  `pulumi:"name"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Service           *string  `pulumi:"service"`
	ServiceResources  []string `pulumi:"serviceResources"`
}

type ServiceEndpointPolicyDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPolicyDefinitionResponse)(nil)).Elem()
}

func (o ServiceEndpointPolicyDefinitionResponseOutput) ToServiceEndpointPolicyDefinitionResponseOutput() ServiceEndpointPolicyDefinitionResponseOutput {
	return o
}

func (o ServiceEndpointPolicyDefinitionResponseOutput) ToServiceEndpointPolicyDefinitionResponseOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionResponseOutput {
	return o
}

func (o ServiceEndpointPolicyDefinitionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyDefinitionResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ServiceEndpointPolicyDefinitionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyDefinitionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyDefinitionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceEndpointPolicyDefinitionResponseOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionResponse) *string { return v.Service }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyDefinitionResponseOutput) ServiceResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyDefinitionResponse) []string { return v.ServiceResources }).(pulumi.StringArrayOutput)
}

type ServiceEndpointPolicyDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPolicyDefinitionResponse)(nil)).Elem()
}

func (o ServiceEndpointPolicyDefinitionResponseArrayOutput) ToServiceEndpointPolicyDefinitionResponseArrayOutput() ServiceEndpointPolicyDefinitionResponseArrayOutput {
	return o
}

func (o ServiceEndpointPolicyDefinitionResponseArrayOutput) ToServiceEndpointPolicyDefinitionResponseArrayOutputWithContext(ctx context.Context) ServiceEndpointPolicyDefinitionResponseArrayOutput {
	return o
}

func (o ServiceEndpointPolicyDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ServiceEndpointPolicyDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointPolicyDefinitionResponse {
		return vs[0].([]ServiceEndpointPolicyDefinitionResponse)[vs[1].(int)]
	}).(ServiceEndpointPolicyDefinitionResponseOutput)
}

type ServiceEndpointPolicyResponse struct {
	Etag                             string                                    `pulumi:"etag"`
	Id                               *string                                   `pulumi:"id"`
	Kind                             string                                    `pulumi:"kind"`
	Location                         *string                                   `pulumi:"location"`
	Name                             string                                    `pulumi:"name"`
	ProvisioningState                string                                    `pulumi:"provisioningState"`
	ResourceGuid                     string                                    `pulumi:"resourceGuid"`
	ServiceEndpointPolicyDefinitions []ServiceEndpointPolicyDefinitionResponse `pulumi:"serviceEndpointPolicyDefinitions"`
	Subnets                          []SubnetResponse                          `pulumi:"subnets"`
	Tags                             map[string]string                         `pulumi:"tags"`
	Type                             string                                    `pulumi:"type"`
}

type ServiceEndpointPolicyResponseOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPolicyResponse)(nil)).Elem()
}

func (o ServiceEndpointPolicyResponseOutput) ToServiceEndpointPolicyResponseOutput() ServiceEndpointPolicyResponseOutput {
	return o
}

func (o ServiceEndpointPolicyResponseOutput) ToServiceEndpointPolicyResponseOutputWithContext(ctx context.Context) ServiceEndpointPolicyResponseOutput {
	return o
}

func (o ServiceEndpointPolicyResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ServiceEndpointPolicyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o ServiceEndpointPolicyResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ServiceEndpointPolicyResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceEndpointPolicyResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceEndpointPolicyResponseOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyResponse) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o ServiceEndpointPolicyResponseOutput) ServiceEndpointPolicyDefinitions() ServiceEndpointPolicyDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyResponse) []ServiceEndpointPolicyDefinitionResponse {
		return v.ServiceEndpointPolicyDefinitions
	}).(ServiceEndpointPolicyDefinitionResponseArrayOutput)
}

func (o ServiceEndpointPolicyResponseOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyResponse) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

func (o ServiceEndpointPolicyResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServiceEndpointPolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointPolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceEndpointPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPolicyResponse)(nil)).Elem()
}

func (o ServiceEndpointPolicyResponseArrayOutput) ToServiceEndpointPolicyResponseArrayOutput() ServiceEndpointPolicyResponseArrayOutput {
	return o
}

func (o ServiceEndpointPolicyResponseArrayOutput) ToServiceEndpointPolicyResponseArrayOutputWithContext(ctx context.Context) ServiceEndpointPolicyResponseArrayOutput {
	return o
}

func (o ServiceEndpointPolicyResponseArrayOutput) Index(i pulumi.IntInput) ServiceEndpointPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointPolicyResponse {
		return vs[0].([]ServiceEndpointPolicyResponse)[vs[1].(int)]
	}).(ServiceEndpointPolicyResponseOutput)
}

type ServiceEndpointPropertiesFormat struct {
	Locations []string `pulumi:"locations"`
	Service   *string  `pulumi:"service"`
}





type ServiceEndpointPropertiesFormatInput interface {
	pulumi.Input

	ToServiceEndpointPropertiesFormatOutput() ServiceEndpointPropertiesFormatOutput
	ToServiceEndpointPropertiesFormatOutputWithContext(context.Context) ServiceEndpointPropertiesFormatOutput
}

type ServiceEndpointPropertiesFormatArgs struct {
	Locations pulumi.StringArrayInput `pulumi:"locations"`
	Service   pulumi.StringPtrInput   `pulumi:"service"`
}

func (ServiceEndpointPropertiesFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPropertiesFormat)(nil)).Elem()
}

func (i ServiceEndpointPropertiesFormatArgs) ToServiceEndpointPropertiesFormatOutput() ServiceEndpointPropertiesFormatOutput {
	return i.ToServiceEndpointPropertiesFormatOutputWithContext(context.Background())
}

func (i ServiceEndpointPropertiesFormatArgs) ToServiceEndpointPropertiesFormatOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPropertiesFormatOutput)
}





type ServiceEndpointPropertiesFormatArrayInput interface {
	pulumi.Input

	ToServiceEndpointPropertiesFormatArrayOutput() ServiceEndpointPropertiesFormatArrayOutput
	ToServiceEndpointPropertiesFormatArrayOutputWithContext(context.Context) ServiceEndpointPropertiesFormatArrayOutput
}

type ServiceEndpointPropertiesFormatArray []ServiceEndpointPropertiesFormatInput

func (ServiceEndpointPropertiesFormatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPropertiesFormat)(nil)).Elem()
}

func (i ServiceEndpointPropertiesFormatArray) ToServiceEndpointPropertiesFormatArrayOutput() ServiceEndpointPropertiesFormatArrayOutput {
	return i.ToServiceEndpointPropertiesFormatArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointPropertiesFormatArray) ToServiceEndpointPropertiesFormatArrayOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPropertiesFormatArrayOutput)
}

type ServiceEndpointPropertiesFormatOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPropertiesFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPropertiesFormat)(nil)).Elem()
}

func (o ServiceEndpointPropertiesFormatOutput) ToServiceEndpointPropertiesFormatOutput() ServiceEndpointPropertiesFormatOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatOutput) ToServiceEndpointPropertiesFormatOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceEndpointPropertiesFormat) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o ServiceEndpointPropertiesFormatOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPropertiesFormat) *string { return v.Service }).(pulumi.StringPtrOutput)
}

type ServiceEndpointPropertiesFormatArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPropertiesFormatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPropertiesFormat)(nil)).Elem()
}

func (o ServiceEndpointPropertiesFormatArrayOutput) ToServiceEndpointPropertiesFormatArrayOutput() ServiceEndpointPropertiesFormatArrayOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatArrayOutput) ToServiceEndpointPropertiesFormatArrayOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatArrayOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatArrayOutput) Index(i pulumi.IntInput) ServiceEndpointPropertiesFormatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointPropertiesFormat {
		return vs[0].([]ServiceEndpointPropertiesFormat)[vs[1].(int)]
	}).(ServiceEndpointPropertiesFormatOutput)
}

type ServiceEndpointPropertiesFormatResponse struct {
	Locations         []string `pulumi:"locations"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Service           *string  `pulumi:"service"`
}

type ServiceEndpointPropertiesFormatResponseOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPropertiesFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointPropertiesFormatResponse)(nil)).Elem()
}

func (o ServiceEndpointPropertiesFormatResponseOutput) ToServiceEndpointPropertiesFormatResponseOutput() ServiceEndpointPropertiesFormatResponseOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatResponseOutput) ToServiceEndpointPropertiesFormatResponseOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatResponseOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceEndpointPropertiesFormatResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o ServiceEndpointPropertiesFormatResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointPropertiesFormatResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceEndpointPropertiesFormatResponseOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceEndpointPropertiesFormatResponse) *string { return v.Service }).(pulumi.StringPtrOutput)
}

type ServiceEndpointPropertiesFormatResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPropertiesFormatResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointPropertiesFormatResponse)(nil)).Elem()
}

func (o ServiceEndpointPropertiesFormatResponseArrayOutput) ToServiceEndpointPropertiesFormatResponseArrayOutput() ServiceEndpointPropertiesFormatResponseArrayOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatResponseArrayOutput) ToServiceEndpointPropertiesFormatResponseArrayOutputWithContext(ctx context.Context) ServiceEndpointPropertiesFormatResponseArrayOutput {
	return o
}

func (o ServiceEndpointPropertiesFormatResponseArrayOutput) Index(i pulumi.IntInput) ServiceEndpointPropertiesFormatResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointPropertiesFormatResponse {
		return vs[0].([]ServiceEndpointPropertiesFormatResponse)[vs[1].(int)]
	}).(ServiceEndpointPropertiesFormatResponseOutput)
}

type SingleQueryResultResponse struct {
	Description               *string  `pulumi:"description"`
	DestinationPorts          []string `pulumi:"destinationPorts"`
	Direction                 *int     `pulumi:"direction"`
	Group                     *string  `pulumi:"group"`
	InheritedFromParentPolicy *bool    `pulumi:"inheritedFromParentPolicy"`
	LastUpdated               *string  `pulumi:"lastUpdated"`
	Mode                      *int     `pulumi:"mode"`
	Protocol                  *string  `pulumi:"protocol"`
	Severity                  *int     `pulumi:"severity"`
	SignatureId               *int     `pulumi:"signatureId"`
	SourcePorts               []string `pulumi:"sourcePorts"`
}

type SingleQueryResultResponseOutput struct{ *pulumi.OutputState }

func (SingleQueryResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SingleQueryResultResponse)(nil)).Elem()
}

func (o SingleQueryResultResponseOutput) ToSingleQueryResultResponseOutput() SingleQueryResultResponseOutput {
	return o
}

func (o SingleQueryResultResponseOutput) ToSingleQueryResultResponseOutputWithContext(ctx context.Context) SingleQueryResultResponseOutput {
	return o
}

func (o SingleQueryResultResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SingleQueryResultResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SingleQueryResultResponseOutput) DestinationPorts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SingleQueryResultResponse) []string { return v.DestinationPorts }).(pulumi.StringArrayOutput)
}

func (o SingleQueryResultResponseOutput) Direction() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SingleQueryResultResponse) *int { return v.Direction }).(pulumi.IntPtrOutput)
}

func (o SingleQueryResultResponseOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SingleQueryResultResponse) *string { return v.Group }).(pulumi.StringPtrOutput)
}

func (o SingleQueryResultResponseOutput) InheritedFromParentPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SingleQueryResultResponse) *bool { return v.InheritedFromParentPolicy }).(pulumi.BoolPtrOutput)
}

func (o SingleQueryResultResponseOutput) LastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SingleQueryResultResponse) *string { return v.LastUpdated }).(pulumi.StringPtrOutput)
}

func (o SingleQueryResultResponseOutput) Mode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SingleQueryResultResponse) *int { return v.Mode }).(pulumi.IntPtrOutput)
}

func (o SingleQueryResultResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SingleQueryResultResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o SingleQueryResultResponseOutput) Severity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SingleQueryResultResponse) *int { return v.Severity }).(pulumi.IntPtrOutput)
}

func (o SingleQueryResultResponseOutput) SignatureId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SingleQueryResultResponse) *int { return v.SignatureId }).(pulumi.IntPtrOutput)
}

func (o SingleQueryResultResponseOutput) SourcePorts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SingleQueryResultResponse) []string { return v.SourcePorts }).(pulumi.StringArrayOutput)
}

type SingleQueryResultResponseArrayOutput struct{ *pulumi.OutputState }

func (SingleQueryResultResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SingleQueryResultResponse)(nil)).Elem()
}

func (o SingleQueryResultResponseArrayOutput) ToSingleQueryResultResponseArrayOutput() SingleQueryResultResponseArrayOutput {
	return o
}

func (o SingleQueryResultResponseArrayOutput) ToSingleQueryResultResponseArrayOutputWithContext(ctx context.Context) SingleQueryResultResponseArrayOutput {
	return o
}

func (o SingleQueryResultResponseArrayOutput) Index(i pulumi.IntInput) SingleQueryResultResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SingleQueryResultResponse {
		return vs[0].([]SingleQueryResultResponse)[vs[1].(int)]
	}).(SingleQueryResultResponseOutput)
}

type Sku struct {
	Name *string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
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

type SkuResponse struct {
	Name *string `pulumi:"name"`
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

type SoaRecord struct {
	Email        *string  `pulumi:"email"`
	ExpireTime   *float64 `pulumi:"expireTime"`
	Host         *string  `pulumi:"host"`
	MinimumTtl   *float64 `pulumi:"minimumTtl"`
	RefreshTime  *float64 `pulumi:"refreshTime"`
	RetryTime    *float64 `pulumi:"retryTime"`
	SerialNumber *float64 `pulumi:"serialNumber"`
}





type SoaRecordInput interface {
	pulumi.Input

	ToSoaRecordOutput() SoaRecordOutput
	ToSoaRecordOutputWithContext(context.Context) SoaRecordOutput
}

type SoaRecordArgs struct {
	Email        pulumi.StringPtrInput  `pulumi:"email"`
	ExpireTime   pulumi.Float64PtrInput `pulumi:"expireTime"`
	Host         pulumi.StringPtrInput  `pulumi:"host"`
	MinimumTtl   pulumi.Float64PtrInput `pulumi:"minimumTtl"`
	RefreshTime  pulumi.Float64PtrInput `pulumi:"refreshTime"`
	RetryTime    pulumi.Float64PtrInput `pulumi:"retryTime"`
	SerialNumber pulumi.Float64PtrInput `pulumi:"serialNumber"`
}

func (SoaRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecord)(nil)).Elem()
}

func (i SoaRecordArgs) ToSoaRecordOutput() SoaRecordOutput {
	return i.ToSoaRecordOutputWithContext(context.Background())
}

func (i SoaRecordArgs) ToSoaRecordOutputWithContext(ctx context.Context) SoaRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordOutput)
}

func (i SoaRecordArgs) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return i.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (i SoaRecordArgs) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordOutput).ToSoaRecordPtrOutputWithContext(ctx)
}









type SoaRecordPtrInput interface {
	pulumi.Input

	ToSoaRecordPtrOutput() SoaRecordPtrOutput
	ToSoaRecordPtrOutputWithContext(context.Context) SoaRecordPtrOutput
}

type soaRecordPtrType SoaRecordArgs

func SoaRecordPtr(v *SoaRecordArgs) SoaRecordPtrInput {
	return (*soaRecordPtrType)(v)
}

func (*soaRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecord)(nil)).Elem()
}

func (i *soaRecordPtrType) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return i.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (i *soaRecordPtrType) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordPtrOutput)
}

type SoaRecordOutput struct{ *pulumi.OutputState }

func (SoaRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecord)(nil)).Elem()
}

func (o SoaRecordOutput) ToSoaRecordOutput() SoaRecordOutput {
	return o
}

func (o SoaRecordOutput) ToSoaRecordOutputWithContext(ctx context.Context) SoaRecordOutput {
	return o
}

func (o SoaRecordOutput) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return o.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (o SoaRecordOutput) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SoaRecord) *SoaRecord {
		return &v
	}).(SoaRecordPtrOutput)
}

func (o SoaRecordOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecord) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o SoaRecordOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.ExpireTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecord) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o SoaRecordOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.MinimumTtl }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.RefreshTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.RetryTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.SerialNumber }).(pulumi.Float64PtrOutput)
}

type SoaRecordPtrOutput struct{ *pulumi.OutputState }

func (SoaRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecord)(nil)).Elem()
}

func (o SoaRecordPtrOutput) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return o
}

func (o SoaRecordPtrOutput) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return o
}

func (o SoaRecordPtrOutput) Elem() SoaRecordOutput {
	return o.ApplyT(func(v *SoaRecord) SoaRecord {
		if v != nil {
			return *v
		}
		var ret SoaRecord
		return ret
	}).(SoaRecordOutput)
}

func (o SoaRecordPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecord) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordPtrOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.ExpireTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecord) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordPtrOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumTtl
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.Float64PtrOutput)
}

type SoaRecordResponse struct {
	Email        *string  `pulumi:"email"`
	ExpireTime   *float64 `pulumi:"expireTime"`
	Host         *string  `pulumi:"host"`
	MinimumTtl   *float64 `pulumi:"minimumTtl"`
	RefreshTime  *float64 `pulumi:"refreshTime"`
	RetryTime    *float64 `pulumi:"retryTime"`
	SerialNumber *float64 `pulumi:"serialNumber"`
}

type SoaRecordResponseOutput struct{ *pulumi.OutputState }

func (SoaRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecordResponse)(nil)).Elem()
}

func (o SoaRecordResponseOutput) ToSoaRecordResponseOutput() SoaRecordResponseOutput {
	return o
}

func (o SoaRecordResponseOutput) ToSoaRecordResponseOutputWithContext(ctx context.Context) SoaRecordResponseOutput {
	return o
}

func (o SoaRecordResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponseOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.ExpireTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponseOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.MinimumTtl }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.RefreshTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.RetryTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.SerialNumber }).(pulumi.Float64PtrOutput)
}

type SoaRecordResponsePtrOutput struct{ *pulumi.OutputState }

func (SoaRecordResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecordResponse)(nil)).Elem()
}

func (o SoaRecordResponsePtrOutput) ToSoaRecordResponsePtrOutput() SoaRecordResponsePtrOutput {
	return o
}

func (o SoaRecordResponsePtrOutput) ToSoaRecordResponsePtrOutputWithContext(ctx context.Context) SoaRecordResponsePtrOutput {
	return o
}

func (o SoaRecordResponsePtrOutput) Elem() SoaRecordResponseOutput {
	return o.ApplyT(func(v *SoaRecordResponse) SoaRecordResponse {
		if v != nil {
			return *v
		}
		var ret SoaRecordResponse
		return ret
	}).(SoaRecordResponseOutput)
}

func (o SoaRecordResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponsePtrOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.ExpireTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponsePtrOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumTtl
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.Float64PtrOutput)
}

type SrvRecord struct {
	Port     *int    `pulumi:"port"`
	Priority *int    `pulumi:"priority"`
	Target   *string `pulumi:"target"`
	Weight   *int    `pulumi:"weight"`
}





type SrvRecordInput interface {
	pulumi.Input

	ToSrvRecordOutput() SrvRecordOutput
	ToSrvRecordOutputWithContext(context.Context) SrvRecordOutput
}

type SrvRecordArgs struct {
	Port     pulumi.IntPtrInput    `pulumi:"port"`
	Priority pulumi.IntPtrInput    `pulumi:"priority"`
	Target   pulumi.StringPtrInput `pulumi:"target"`
	Weight   pulumi.IntPtrInput    `pulumi:"weight"`
}

func (SrvRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecord)(nil)).Elem()
}

func (i SrvRecordArgs) ToSrvRecordOutput() SrvRecordOutput {
	return i.ToSrvRecordOutputWithContext(context.Background())
}

func (i SrvRecordArgs) ToSrvRecordOutputWithContext(ctx context.Context) SrvRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordOutput)
}





type SrvRecordArrayInput interface {
	pulumi.Input

	ToSrvRecordArrayOutput() SrvRecordArrayOutput
	ToSrvRecordArrayOutputWithContext(context.Context) SrvRecordArrayOutput
}

type SrvRecordArray []SrvRecordInput

func (SrvRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecord)(nil)).Elem()
}

func (i SrvRecordArray) ToSrvRecordArrayOutput() SrvRecordArrayOutput {
	return i.ToSrvRecordArrayOutputWithContext(context.Background())
}

func (i SrvRecordArray) ToSrvRecordArrayOutputWithContext(ctx context.Context) SrvRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordArrayOutput)
}

type SrvRecordOutput struct{ *pulumi.OutputState }

func (SrvRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecord)(nil)).Elem()
}

func (o SrvRecordOutput) ToSrvRecordOutput() SrvRecordOutput {
	return o
}

func (o SrvRecordOutput) ToSrvRecordOutputWithContext(ctx context.Context) SrvRecordOutput {
	return o
}

func (o SrvRecordOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o SrvRecordOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SrvRecordOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SrvRecord) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o SrvRecordOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type SrvRecordArrayOutput struct{ *pulumi.OutputState }

func (SrvRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecord)(nil)).Elem()
}

func (o SrvRecordArrayOutput) ToSrvRecordArrayOutput() SrvRecordArrayOutput {
	return o
}

func (o SrvRecordArrayOutput) ToSrvRecordArrayOutputWithContext(ctx context.Context) SrvRecordArrayOutput {
	return o
}

func (o SrvRecordArrayOutput) Index(i pulumi.IntInput) SrvRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SrvRecord {
		return vs[0].([]SrvRecord)[vs[1].(int)]
	}).(SrvRecordOutput)
}

type SrvRecordResponse struct {
	Port     *int    `pulumi:"port"`
	Priority *int    `pulumi:"priority"`
	Target   *string `pulumi:"target"`
	Weight   *int    `pulumi:"weight"`
}

type SrvRecordResponseOutput struct{ *pulumi.OutputState }

func (SrvRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecordResponse)(nil)).Elem()
}

func (o SrvRecordResponseOutput) ToSrvRecordResponseOutput() SrvRecordResponseOutput {
	return o
}

func (o SrvRecordResponseOutput) ToSrvRecordResponseOutputWithContext(ctx context.Context) SrvRecordResponseOutput {
	return o
}

func (o SrvRecordResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o SrvRecordResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SrvRecordResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o SrvRecordResponseOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type SrvRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (SrvRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecordResponse)(nil)).Elem()
}

func (o SrvRecordResponseArrayOutput) ToSrvRecordResponseArrayOutput() SrvRecordResponseArrayOutput {
	return o
}

func (o SrvRecordResponseArrayOutput) ToSrvRecordResponseArrayOutputWithContext(ctx context.Context) SrvRecordResponseArrayOutput {
	return o
}

func (o SrvRecordResponseArrayOutput) Index(i pulumi.IntInput) SrvRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SrvRecordResponse {
		return vs[0].([]SrvRecordResponse)[vs[1].(int)]
	}).(SrvRecordResponseOutput)
}

type StaticRoute struct {
	AddressPrefixes  []string `pulumi:"addressPrefixes"`
	Name             *string  `pulumi:"name"`
	NextHopIpAddress *string  `pulumi:"nextHopIpAddress"`
}





type StaticRouteInput interface {
	pulumi.Input

	ToStaticRouteOutput() StaticRouteOutput
	ToStaticRouteOutputWithContext(context.Context) StaticRouteOutput
}

type StaticRouteArgs struct {
	AddressPrefixes  pulumi.StringArrayInput `pulumi:"addressPrefixes"`
	Name             pulumi.StringPtrInput   `pulumi:"name"`
	NextHopIpAddress pulumi.StringPtrInput   `pulumi:"nextHopIpAddress"`
}

func (StaticRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticRoute)(nil)).Elem()
}

func (i StaticRouteArgs) ToStaticRouteOutput() StaticRouteOutput {
	return i.ToStaticRouteOutputWithContext(context.Background())
}

func (i StaticRouteArgs) ToStaticRouteOutputWithContext(ctx context.Context) StaticRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticRouteOutput)
}





type StaticRouteArrayInput interface {
	pulumi.Input

	ToStaticRouteArrayOutput() StaticRouteArrayOutput
	ToStaticRouteArrayOutputWithContext(context.Context) StaticRouteArrayOutput
}

type StaticRouteArray []StaticRouteInput

func (StaticRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticRoute)(nil)).Elem()
}

func (i StaticRouteArray) ToStaticRouteArrayOutput() StaticRouteArrayOutput {
	return i.ToStaticRouteArrayOutputWithContext(context.Background())
}

func (i StaticRouteArray) ToStaticRouteArrayOutputWithContext(ctx context.Context) StaticRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticRouteArrayOutput)
}

type StaticRouteOutput struct{ *pulumi.OutputState }

func (StaticRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticRoute)(nil)).Elem()
}

func (o StaticRouteOutput) ToStaticRouteOutput() StaticRouteOutput {
	return o
}

func (o StaticRouteOutput) ToStaticRouteOutputWithContext(ctx context.Context) StaticRouteOutput {
	return o
}

func (o StaticRouteOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StaticRoute) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o StaticRouteOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticRoute) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StaticRouteOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticRoute) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

type StaticRouteArrayOutput struct{ *pulumi.OutputState }

func (StaticRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticRoute)(nil)).Elem()
}

func (o StaticRouteArrayOutput) ToStaticRouteArrayOutput() StaticRouteArrayOutput {
	return o
}

func (o StaticRouteArrayOutput) ToStaticRouteArrayOutputWithContext(ctx context.Context) StaticRouteArrayOutput {
	return o
}

func (o StaticRouteArrayOutput) Index(i pulumi.IntInput) StaticRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StaticRoute {
		return vs[0].([]StaticRoute)[vs[1].(int)]
	}).(StaticRouteOutput)
}

type StaticRouteResponse struct {
	AddressPrefixes  []string `pulumi:"addressPrefixes"`
	Name             *string  `pulumi:"name"`
	NextHopIpAddress *string  `pulumi:"nextHopIpAddress"`
}

type StaticRouteResponseOutput struct{ *pulumi.OutputState }

func (StaticRouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticRouteResponse)(nil)).Elem()
}

func (o StaticRouteResponseOutput) ToStaticRouteResponseOutput() StaticRouteResponseOutput {
	return o
}

func (o StaticRouteResponseOutput) ToStaticRouteResponseOutputWithContext(ctx context.Context) StaticRouteResponseOutput {
	return o
}

func (o StaticRouteResponseOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StaticRouteResponse) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o StaticRouteResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticRouteResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StaticRouteResponseOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticRouteResponse) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

type StaticRouteResponseArrayOutput struct{ *pulumi.OutputState }

func (StaticRouteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticRouteResponse)(nil)).Elem()
}

func (o StaticRouteResponseArrayOutput) ToStaticRouteResponseArrayOutput() StaticRouteResponseArrayOutput {
	return o
}

func (o StaticRouteResponseArrayOutput) ToStaticRouteResponseArrayOutputWithContext(ctx context.Context) StaticRouteResponseArrayOutput {
	return o
}

func (o StaticRouteResponseArrayOutput) Index(i pulumi.IntInput) StaticRouteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StaticRouteResponse {
		return vs[0].([]StaticRouteResponse)[vs[1].(int)]
	}).(StaticRouteResponseOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}









type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}





type SubResourceArrayInput interface {
	pulumi.Input

	ToSubResourceArrayOutput() SubResourceArrayOutput
	ToSubResourceArrayOutputWithContext(context.Context) SubResourceArrayOutput
}

type SubResourceArray []SubResourceInput

func (SubResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (i SubResourceArray) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return i.ToSubResourceArrayOutputWithContext(context.Background())
}

func (i SubResourceArray) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceArrayOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceArrayOutput struct{ *pulumi.OutputState }

func (SubResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) Index(i pulumi.IntInput) SubResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResource {
		return vs[0].([]SubResource)[vs[1].(int)]
	}).(SubResourceOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) Index(i pulumi.IntInput) SubResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceResponse {
		return vs[0].([]SubResourceResponse)[vs[1].(int)]
	}).(SubResourceResponseOutput)
}

type SubnetType struct {
	AddressPrefix                      *string                             `pulumi:"addressPrefix"`
	AddressPrefixes                    []string                            `pulumi:"addressPrefixes"`
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration `pulumi:"applicationGatewayIpConfigurations"`
	Delegations                        []Delegation                        `pulumi:"delegations"`
	Id                                 *string                             `pulumi:"id"`
	IpAllocations                      []SubResource                       `pulumi:"ipAllocations"`
	Name                               *string                             `pulumi:"name"`
	NatGateway                         *SubResource                        `pulumi:"natGateway"`
	NetworkSecurityGroup               *NetworkSecurityGroupType           `pulumi:"networkSecurityGroup"`
	PrivateEndpointNetworkPolicies     *string                             `pulumi:"privateEndpointNetworkPolicies"`
	PrivateLinkServiceNetworkPolicies  *string                             `pulumi:"privateLinkServiceNetworkPolicies"`
	RouteTable                         *RouteTableType                     `pulumi:"routeTable"`
	ServiceEndpointPolicies            []ServiceEndpointPolicyType         `pulumi:"serviceEndpointPolicies"`
	ServiceEndpoints                   []ServiceEndpointPropertiesFormat   `pulumi:"serviceEndpoints"`
	Type                               *string                             `pulumi:"type"`
}


func (val *SubnetType) Defaults() *SubnetType {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrivateEndpointNetworkPolicies) {
		privateEndpointNetworkPolicies_ := "Enabled"
		tmp.PrivateEndpointNetworkPolicies = &privateEndpointNetworkPolicies_
	}
	if isZero(tmp.PrivateLinkServiceNetworkPolicies) {
		privateLinkServiceNetworkPolicies_ := "Enabled"
		tmp.PrivateLinkServiceNetworkPolicies = &privateLinkServiceNetworkPolicies_
	}
	return &tmp
}





type SubnetTypeInput interface {
	pulumi.Input

	ToSubnetTypeOutput() SubnetTypeOutput
	ToSubnetTypeOutputWithContext(context.Context) SubnetTypeOutput
}

type SubnetTypeArgs struct {
	AddressPrefix                      pulumi.StringPtrInput                       `pulumi:"addressPrefix"`
	AddressPrefixes                    pulumi.StringArrayInput                     `pulumi:"addressPrefixes"`
	ApplicationGatewayIpConfigurations ApplicationGatewayIPConfigurationArrayInput `pulumi:"applicationGatewayIpConfigurations"`
	Delegations                        DelegationArrayInput                        `pulumi:"delegations"`
	Id                                 pulumi.StringPtrInput                       `pulumi:"id"`
	IpAllocations                      SubResourceArrayInput                       `pulumi:"ipAllocations"`
	Name                               pulumi.StringPtrInput                       `pulumi:"name"`
	NatGateway                         SubResourcePtrInput                         `pulumi:"natGateway"`
	NetworkSecurityGroup               NetworkSecurityGroupTypePtrInput            `pulumi:"networkSecurityGroup"`
	PrivateEndpointNetworkPolicies     pulumi.StringPtrInput                       `pulumi:"privateEndpointNetworkPolicies"`
	PrivateLinkServiceNetworkPolicies  pulumi.StringPtrInput                       `pulumi:"privateLinkServiceNetworkPolicies"`
	RouteTable                         RouteTableTypePtrInput                      `pulumi:"routeTable"`
	ServiceEndpointPolicies            ServiceEndpointPolicyTypeArrayInput         `pulumi:"serviceEndpointPolicies"`
	ServiceEndpoints                   ServiceEndpointPropertiesFormatArrayInput   `pulumi:"serviceEndpoints"`
	Type                               pulumi.StringPtrInput                       `pulumi:"type"`
}


func (val *SubnetTypeArgs) Defaults() *SubnetTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrivateEndpointNetworkPolicies) {
		tmp.PrivateEndpointNetworkPolicies = pulumi.StringPtr("Enabled")
	}
	if isZero(tmp.PrivateLinkServiceNetworkPolicies) {
		tmp.PrivateLinkServiceNetworkPolicies = pulumi.StringPtr("Enabled")
	}
	return &tmp
}
func (SubnetTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetType)(nil)).Elem()
}

func (i SubnetTypeArgs) ToSubnetTypeOutput() SubnetTypeOutput {
	return i.ToSubnetTypeOutputWithContext(context.Background())
}

func (i SubnetTypeArgs) ToSubnetTypeOutputWithContext(ctx context.Context) SubnetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypeOutput)
}

func (i SubnetTypeArgs) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return i.ToSubnetTypePtrOutputWithContext(context.Background())
}

func (i SubnetTypeArgs) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypeOutput).ToSubnetTypePtrOutputWithContext(ctx)
}









type SubnetTypePtrInput interface {
	pulumi.Input

	ToSubnetTypePtrOutput() SubnetTypePtrOutput
	ToSubnetTypePtrOutputWithContext(context.Context) SubnetTypePtrOutput
}

type subnetTypePtrType SubnetTypeArgs

func SubnetTypePtr(v *SubnetTypeArgs) SubnetTypePtrInput {
	return (*subnetTypePtrType)(v)
}

func (*subnetTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetType)(nil)).Elem()
}

func (i *subnetTypePtrType) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return i.ToSubnetTypePtrOutputWithContext(context.Background())
}

func (i *subnetTypePtrType) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypePtrOutput)
}





type SubnetTypeArrayInput interface {
	pulumi.Input

	ToSubnetTypeArrayOutput() SubnetTypeArrayOutput
	ToSubnetTypeArrayOutputWithContext(context.Context) SubnetTypeArrayOutput
}

type SubnetTypeArray []SubnetTypeInput

func (SubnetTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetType)(nil)).Elem()
}

func (i SubnetTypeArray) ToSubnetTypeArrayOutput() SubnetTypeArrayOutput {
	return i.ToSubnetTypeArrayOutputWithContext(context.Background())
}

func (i SubnetTypeArray) ToSubnetTypeArrayOutputWithContext(ctx context.Context) SubnetTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypeArrayOutput)
}

type SubnetTypeOutput struct{ *pulumi.OutputState }

func (SubnetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetType)(nil)).Elem()
}

func (o SubnetTypeOutput) ToSubnetTypeOutput() SubnetTypeOutput {
	return o
}

func (o SubnetTypeOutput) ToSubnetTypeOutputWithContext(ctx context.Context) SubnetTypeOutput {
	return o
}

func (o SubnetTypeOutput) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return o.ToSubnetTypePtrOutputWithContext(context.Background())
}

func (o SubnetTypeOutput) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetType) *SubnetType {
		return &v
	}).(SubnetTypePtrOutput)
}

func (o SubnetTypeOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SubnetType) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o SubnetTypeOutput) ApplicationGatewayIpConfigurations() ApplicationGatewayIPConfigurationArrayOutput {
	return o.ApplyT(func(v SubnetType) []ApplicationGatewayIPConfiguration { return v.ApplicationGatewayIpConfigurations }).(ApplicationGatewayIPConfigurationArrayOutput)
}

func (o SubnetTypeOutput) Delegations() DelegationArrayOutput {
	return o.ApplyT(func(v SubnetType) []Delegation { return v.Delegations }).(DelegationArrayOutput)
}

func (o SubnetTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) IpAllocations() SubResourceArrayOutput {
	return o.ApplyT(func(v SubnetType) []SubResource { return v.IpAllocations }).(SubResourceArrayOutput)
}

func (o SubnetTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) NatGateway() SubResourcePtrOutput {
	return o.ApplyT(func(v SubnetType) *SubResource { return v.NatGateway }).(SubResourcePtrOutput)
}

func (o SubnetTypeOutput) NetworkSecurityGroup() NetworkSecurityGroupTypePtrOutput {
	return o.ApplyT(func(v SubnetType) *NetworkSecurityGroupType { return v.NetworkSecurityGroup }).(NetworkSecurityGroupTypePtrOutput)
}

func (o SubnetTypeOutput) PrivateEndpointNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.PrivateEndpointNetworkPolicies }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) PrivateLinkServiceNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.PrivateLinkServiceNetworkPolicies }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) RouteTable() RouteTableTypePtrOutput {
	return o.ApplyT(func(v SubnetType) *RouteTableType { return v.RouteTable }).(RouteTableTypePtrOutput)
}

func (o SubnetTypeOutput) ServiceEndpointPolicies() ServiceEndpointPolicyTypeArrayOutput {
	return o.ApplyT(func(v SubnetType) []ServiceEndpointPolicyType { return v.ServiceEndpointPolicies }).(ServiceEndpointPolicyTypeArrayOutput)
}

func (o SubnetTypeOutput) ServiceEndpoints() ServiceEndpointPropertiesFormatArrayOutput {
	return o.ApplyT(func(v SubnetType) []ServiceEndpointPropertiesFormat { return v.ServiceEndpoints }).(ServiceEndpointPropertiesFormatArrayOutput)
}

func (o SubnetTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type SubnetTypePtrOutput struct{ *pulumi.OutputState }

func (SubnetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetType)(nil)).Elem()
}

func (o SubnetTypePtrOutput) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return o
}

func (o SubnetTypePtrOutput) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return o
}

func (o SubnetTypePtrOutput) Elem() SubnetTypeOutput {
	return o.ApplyT(func(v *SubnetType) SubnetType {
		if v != nil {
			return *v
		}
		var ret SubnetType
		return ret
	}).(SubnetTypeOutput)
}

func (o SubnetTypePtrOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.AddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubnetType) []string {
		if v == nil {
			return nil
		}
		return v.AddressPrefixes
	}).(pulumi.StringArrayOutput)
}

func (o SubnetTypePtrOutput) ApplicationGatewayIpConfigurations() ApplicationGatewayIPConfigurationArrayOutput {
	return o.ApplyT(func(v *SubnetType) []ApplicationGatewayIPConfiguration {
		if v == nil {
			return nil
		}
		return v.ApplicationGatewayIpConfigurations
	}).(ApplicationGatewayIPConfigurationArrayOutput)
}

func (o SubnetTypePtrOutput) Delegations() DelegationArrayOutput {
	return o.ApplyT(func(v *SubnetType) []Delegation {
		if v == nil {
			return nil
		}
		return v.Delegations
	}).(DelegationArrayOutput)
}

func (o SubnetTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) IpAllocations() SubResourceArrayOutput {
	return o.ApplyT(func(v *SubnetType) []SubResource {
		if v == nil {
			return nil
		}
		return v.IpAllocations
	}).(SubResourceArrayOutput)
}

func (o SubnetTypePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) NatGateway() SubResourcePtrOutput {
	return o.ApplyT(func(v *SubnetType) *SubResource {
		if v == nil {
			return nil
		}
		return v.NatGateway
	}).(SubResourcePtrOutput)
}

func (o SubnetTypePtrOutput) NetworkSecurityGroup() NetworkSecurityGroupTypePtrOutput {
	return o.ApplyT(func(v *SubnetType) *NetworkSecurityGroupType {
		if v == nil {
			return nil
		}
		return v.NetworkSecurityGroup
	}).(NetworkSecurityGroupTypePtrOutput)
}

func (o SubnetTypePtrOutput) PrivateEndpointNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointNetworkPolicies
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) PrivateLinkServiceNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceNetworkPolicies
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) RouteTable() RouteTableTypePtrOutput {
	return o.ApplyT(func(v *SubnetType) *RouteTableType {
		if v == nil {
			return nil
		}
		return v.RouteTable
	}).(RouteTableTypePtrOutput)
}

func (o SubnetTypePtrOutput) ServiceEndpointPolicies() ServiceEndpointPolicyTypeArrayOutput {
	return o.ApplyT(func(v *SubnetType) []ServiceEndpointPolicyType {
		if v == nil {
			return nil
		}
		return v.ServiceEndpointPolicies
	}).(ServiceEndpointPolicyTypeArrayOutput)
}

func (o SubnetTypePtrOutput) ServiceEndpoints() ServiceEndpointPropertiesFormatArrayOutput {
	return o.ApplyT(func(v *SubnetType) []ServiceEndpointPropertiesFormat {
		if v == nil {
			return nil
		}
		return v.ServiceEndpoints
	}).(ServiceEndpointPropertiesFormatArrayOutput)
}

func (o SubnetTypePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type SubnetTypeArrayOutput struct{ *pulumi.OutputState }

func (SubnetTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetType)(nil)).Elem()
}

func (o SubnetTypeArrayOutput) ToSubnetTypeArrayOutput() SubnetTypeArrayOutput {
	return o
}

func (o SubnetTypeArrayOutput) ToSubnetTypeArrayOutputWithContext(ctx context.Context) SubnetTypeArrayOutput {
	return o
}

func (o SubnetTypeArrayOutput) Index(i pulumi.IntInput) SubnetTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetType {
		return vs[0].([]SubnetType)[vs[1].(int)]
	}).(SubnetTypeOutput)
}

type SubnetResponse struct {
	AddressPrefix                      *string                                     `pulumi:"addressPrefix"`
	AddressPrefixes                    []string                                    `pulumi:"addressPrefixes"`
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfigurationResponse `pulumi:"applicationGatewayIpConfigurations"`
	Delegations                        []DelegationResponse                        `pulumi:"delegations"`
	Etag                               string                                      `pulumi:"etag"`
	Id                                 *string                                     `pulumi:"id"`
	IpAllocations                      []SubResourceResponse                       `pulumi:"ipAllocations"`
	IpConfigurationProfiles            []IPConfigurationProfileResponse            `pulumi:"ipConfigurationProfiles"`
	IpConfigurations                   []IPConfigurationResponse                   `pulumi:"ipConfigurations"`
	Name                               *string                                     `pulumi:"name"`
	NatGateway                         *SubResourceResponse                        `pulumi:"natGateway"`
	NetworkSecurityGroup               *NetworkSecurityGroupResponse               `pulumi:"networkSecurityGroup"`
	PrivateEndpointNetworkPolicies     *string                                     `pulumi:"privateEndpointNetworkPolicies"`
	PrivateEndpoints                   []PrivateEndpointResponse                   `pulumi:"privateEndpoints"`
	PrivateLinkServiceNetworkPolicies  *string                                     `pulumi:"privateLinkServiceNetworkPolicies"`
	ProvisioningState                  string                                      `pulumi:"provisioningState"`
	Purpose                            string                                      `pulumi:"purpose"`
	ResourceNavigationLinks            []ResourceNavigationLinkResponse            `pulumi:"resourceNavigationLinks"`
	RouteTable                         *RouteTableResponse                         `pulumi:"routeTable"`
	ServiceAssociationLinks            []ServiceAssociationLinkResponse            `pulumi:"serviceAssociationLinks"`
	ServiceEndpointPolicies            []ServiceEndpointPolicyResponse             `pulumi:"serviceEndpointPolicies"`
	ServiceEndpoints                   []ServiceEndpointPropertiesFormatResponse   `pulumi:"serviceEndpoints"`
	Type                               *string                                     `pulumi:"type"`
}


func (val *SubnetResponse) Defaults() *SubnetResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrivateEndpointNetworkPolicies) {
		privateEndpointNetworkPolicies_ := "Enabled"
		tmp.PrivateEndpointNetworkPolicies = &privateEndpointNetworkPolicies_
	}
	if isZero(tmp.PrivateLinkServiceNetworkPolicies) {
		privateLinkServiceNetworkPolicies_ := "Enabled"
		tmp.PrivateLinkServiceNetworkPolicies = &privateLinkServiceNetworkPolicies_
	}
	return &tmp
}

type SubnetResponseOutput struct{ *pulumi.OutputState }

func (SubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseOutput) ToSubnetResponseOutput() SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o SubnetResponseOutput) ApplicationGatewayIpConfigurations() ApplicationGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []ApplicationGatewayIPConfigurationResponse {
		return v.ApplicationGatewayIpConfigurations
	}).(ApplicationGatewayIPConfigurationResponseArrayOutput)
}

func (o SubnetResponseOutput) Delegations() DelegationResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []DelegationResponse { return v.Delegations }).(DelegationResponseArrayOutput)
}

func (o SubnetResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SubnetResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) IpAllocations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []SubResourceResponse { return v.IpAllocations }).(SubResourceResponseArrayOutput)
}

func (o SubnetResponseOutput) IpConfigurationProfiles() IPConfigurationProfileResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []IPConfigurationProfileResponse { return v.IpConfigurationProfiles }).(IPConfigurationProfileResponseArrayOutput)
}

func (o SubnetResponseOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []IPConfigurationResponse { return v.IpConfigurations }).(IPConfigurationResponseArrayOutput)
}

func (o SubnetResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) NatGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v SubnetResponse) *SubResourceResponse { return v.NatGateway }).(SubResourceResponsePtrOutput)
}

func (o SubnetResponseOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v SubnetResponse) *NetworkSecurityGroupResponse { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

func (o SubnetResponseOutput) PrivateEndpointNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.PrivateEndpointNetworkPolicies }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) PrivateEndpoints() PrivateEndpointResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []PrivateEndpointResponse { return v.PrivateEndpoints }).(PrivateEndpointResponseArrayOutput)
}

func (o SubnetResponseOutput) PrivateLinkServiceNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.PrivateLinkServiceNetworkPolicies }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SubnetResponseOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetResponse) string { return v.Purpose }).(pulumi.StringOutput)
}

func (o SubnetResponseOutput) ResourceNavigationLinks() ResourceNavigationLinkResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []ResourceNavigationLinkResponse { return v.ResourceNavigationLinks }).(ResourceNavigationLinkResponseArrayOutput)
}

func (o SubnetResponseOutput) RouteTable() RouteTableResponsePtrOutput {
	return o.ApplyT(func(v SubnetResponse) *RouteTableResponse { return v.RouteTable }).(RouteTableResponsePtrOutput)
}

func (o SubnetResponseOutput) ServiceAssociationLinks() ServiceAssociationLinkResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []ServiceAssociationLinkResponse { return v.ServiceAssociationLinks }).(ServiceAssociationLinkResponseArrayOutput)
}

func (o SubnetResponseOutput) ServiceEndpointPolicies() ServiceEndpointPolicyResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []ServiceEndpointPolicyResponse { return v.ServiceEndpointPolicies }).(ServiceEndpointPolicyResponseArrayOutput)
}

func (o SubnetResponseOutput) ServiceEndpoints() ServiceEndpointPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []ServiceEndpointPropertiesFormatResponse { return v.ServiceEndpoints }).(ServiceEndpointPropertiesFormatResponseArrayOutput)
}

func (o SubnetResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type SubnetResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetResponse)(nil)).Elem()
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutput() SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutputWithContext(ctx context.Context) SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) Elem() SubnetResponseOutput {
	return o.ApplyT(func(v *SubnetResponse) SubnetResponse {
		if v != nil {
			return *v
		}
		var ret SubnetResponse
		return ret
	}).(SubnetResponseOutput)
}

func (o SubnetResponsePtrOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.AddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []string {
		if v == nil {
			return nil
		}
		return v.AddressPrefixes
	}).(pulumi.StringArrayOutput)
}

func (o SubnetResponsePtrOutput) ApplicationGatewayIpConfigurations() ApplicationGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []ApplicationGatewayIPConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.ApplicationGatewayIpConfigurations
	}).(ApplicationGatewayIPConfigurationResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) Delegations() DelegationResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []DelegationResponse {
		if v == nil {
			return nil
		}
		return v.Delegations
	}).(DelegationResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) IpAllocations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.IpAllocations
	}).(SubResourceResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) IpConfigurationProfiles() IPConfigurationProfileResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []IPConfigurationProfileResponse {
		if v == nil {
			return nil
		}
		return v.IpConfigurationProfiles
	}).(IPConfigurationProfileResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []IPConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.IpConfigurations
	}).(IPConfigurationResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) NatGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.NatGateway
	}).(SubResourceResponsePtrOutput)
}

func (o SubnetResponsePtrOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *NetworkSecurityGroupResponse {
		if v == nil {
			return nil
		}
		return v.NetworkSecurityGroup
	}).(NetworkSecurityGroupResponsePtrOutput)
}

func (o SubnetResponsePtrOutput) PrivateEndpointNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointNetworkPolicies
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) PrivateEndpoints() PrivateEndpointResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []PrivateEndpointResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoints
	}).(PrivateEndpointResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) PrivateLinkServiceNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceNetworkPolicies
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Purpose
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) ResourceNavigationLinks() ResourceNavigationLinkResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []ResourceNavigationLinkResponse {
		if v == nil {
			return nil
		}
		return v.ResourceNavigationLinks
	}).(ResourceNavigationLinkResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) RouteTable() RouteTableResponsePtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *RouteTableResponse {
		if v == nil {
			return nil
		}
		return v.RouteTable
	}).(RouteTableResponsePtrOutput)
}

func (o SubnetResponsePtrOutput) ServiceAssociationLinks() ServiceAssociationLinkResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []ServiceAssociationLinkResponse {
		if v == nil {
			return nil
		}
		return v.ServiceAssociationLinks
	}).(ServiceAssociationLinkResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) ServiceEndpointPolicies() ServiceEndpointPolicyResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []ServiceEndpointPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ServiceEndpointPolicies
	}).(ServiceEndpointPolicyResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) ServiceEndpoints() ServiceEndpointPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []ServiceEndpointPropertiesFormatResponse {
		if v == nil {
			return nil
		}
		return v.ServiceEndpoints
	}).(ServiceEndpointPropertiesFormatResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type SubnetResponseArrayOutput struct{ *pulumi.OutputState }

func (SubnetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutput() SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutputWithContext(ctx context.Context) SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) Index(i pulumi.IntInput) SubnetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetResponse {
		return vs[0].([]SubnetResponse)[vs[1].(int)]
	}).(SubnetResponseOutput)
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

type TargetDnsServer struct {
	IpAddress *string `pulumi:"ipAddress"`
	Port      *int    `pulumi:"port"`
}


func (val *TargetDnsServer) Defaults() *TargetDnsServer {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Port) {
		port_ := 53
		tmp.Port = &port_
	}
	return &tmp
}





type TargetDnsServerInput interface {
	pulumi.Input

	ToTargetDnsServerOutput() TargetDnsServerOutput
	ToTargetDnsServerOutputWithContext(context.Context) TargetDnsServerOutput
}

type TargetDnsServerArgs struct {
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	Port      pulumi.IntPtrInput    `pulumi:"port"`
}


func (val *TargetDnsServerArgs) Defaults() *TargetDnsServerArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Port) {
		tmp.Port = pulumi.IntPtr(53)
	}
	return &tmp
}
func (TargetDnsServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetDnsServer)(nil)).Elem()
}

func (i TargetDnsServerArgs) ToTargetDnsServerOutput() TargetDnsServerOutput {
	return i.ToTargetDnsServerOutputWithContext(context.Background())
}

func (i TargetDnsServerArgs) ToTargetDnsServerOutputWithContext(ctx context.Context) TargetDnsServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetDnsServerOutput)
}





type TargetDnsServerArrayInput interface {
	pulumi.Input

	ToTargetDnsServerArrayOutput() TargetDnsServerArrayOutput
	ToTargetDnsServerArrayOutputWithContext(context.Context) TargetDnsServerArrayOutput
}

type TargetDnsServerArray []TargetDnsServerInput

func (TargetDnsServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetDnsServer)(nil)).Elem()
}

func (i TargetDnsServerArray) ToTargetDnsServerArrayOutput() TargetDnsServerArrayOutput {
	return i.ToTargetDnsServerArrayOutputWithContext(context.Background())
}

func (i TargetDnsServerArray) ToTargetDnsServerArrayOutputWithContext(ctx context.Context) TargetDnsServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetDnsServerArrayOutput)
}

type TargetDnsServerOutput struct{ *pulumi.OutputState }

func (TargetDnsServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetDnsServer)(nil)).Elem()
}

func (o TargetDnsServerOutput) ToTargetDnsServerOutput() TargetDnsServerOutput {
	return o
}

func (o TargetDnsServerOutput) ToTargetDnsServerOutputWithContext(ctx context.Context) TargetDnsServerOutput {
	return o
}

func (o TargetDnsServerOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetDnsServer) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o TargetDnsServerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TargetDnsServer) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type TargetDnsServerArrayOutput struct{ *pulumi.OutputState }

func (TargetDnsServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetDnsServer)(nil)).Elem()
}

func (o TargetDnsServerArrayOutput) ToTargetDnsServerArrayOutput() TargetDnsServerArrayOutput {
	return o
}

func (o TargetDnsServerArrayOutput) ToTargetDnsServerArrayOutputWithContext(ctx context.Context) TargetDnsServerArrayOutput {
	return o
}

func (o TargetDnsServerArrayOutput) Index(i pulumi.IntInput) TargetDnsServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetDnsServer {
		return vs[0].([]TargetDnsServer)[vs[1].(int)]
	}).(TargetDnsServerOutput)
}

type TargetDnsServerResponse struct {
	IpAddress *string `pulumi:"ipAddress"`
	Port      *int    `pulumi:"port"`
}


func (val *TargetDnsServerResponse) Defaults() *TargetDnsServerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Port) {
		port_ := 53
		tmp.Port = &port_
	}
	return &tmp
}

type TargetDnsServerResponseOutput struct{ *pulumi.OutputState }

func (TargetDnsServerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetDnsServerResponse)(nil)).Elem()
}

func (o TargetDnsServerResponseOutput) ToTargetDnsServerResponseOutput() TargetDnsServerResponseOutput {
	return o
}

func (o TargetDnsServerResponseOutput) ToTargetDnsServerResponseOutputWithContext(ctx context.Context) TargetDnsServerResponseOutput {
	return o
}

func (o TargetDnsServerResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetDnsServerResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o TargetDnsServerResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TargetDnsServerResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type TargetDnsServerResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetDnsServerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetDnsServerResponse)(nil)).Elem()
}

func (o TargetDnsServerResponseArrayOutput) ToTargetDnsServerResponseArrayOutput() TargetDnsServerResponseArrayOutput {
	return o
}

func (o TargetDnsServerResponseArrayOutput) ToTargetDnsServerResponseArrayOutputWithContext(ctx context.Context) TargetDnsServerResponseArrayOutput {
	return o
}

func (o TargetDnsServerResponseArrayOutput) Index(i pulumi.IntInput) TargetDnsServerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetDnsServerResponse {
		return vs[0].([]TargetDnsServerResponse)[vs[1].(int)]
	}).(TargetDnsServerResponseOutput)
}

type TrafficAnalyticsConfigurationProperties struct {
	Enabled                  *bool   `pulumi:"enabled"`
	TrafficAnalyticsInterval *int    `pulumi:"trafficAnalyticsInterval"`
	WorkspaceId              *string `pulumi:"workspaceId"`
	WorkspaceRegion          *string `pulumi:"workspaceRegion"`
	WorkspaceResourceId      *string `pulumi:"workspaceResourceId"`
}





type TrafficAnalyticsConfigurationPropertiesInput interface {
	pulumi.Input

	ToTrafficAnalyticsConfigurationPropertiesOutput() TrafficAnalyticsConfigurationPropertiesOutput
	ToTrafficAnalyticsConfigurationPropertiesOutputWithContext(context.Context) TrafficAnalyticsConfigurationPropertiesOutput
}

type TrafficAnalyticsConfigurationPropertiesArgs struct {
	Enabled                  pulumi.BoolPtrInput   `pulumi:"enabled"`
	TrafficAnalyticsInterval pulumi.IntPtrInput    `pulumi:"trafficAnalyticsInterval"`
	WorkspaceId              pulumi.StringPtrInput `pulumi:"workspaceId"`
	WorkspaceRegion          pulumi.StringPtrInput `pulumi:"workspaceRegion"`
	WorkspaceResourceId      pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (TrafficAnalyticsConfigurationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficAnalyticsConfigurationProperties)(nil)).Elem()
}

func (i TrafficAnalyticsConfigurationPropertiesArgs) ToTrafficAnalyticsConfigurationPropertiesOutput() TrafficAnalyticsConfigurationPropertiesOutput {
	return i.ToTrafficAnalyticsConfigurationPropertiesOutputWithContext(context.Background())
}

func (i TrafficAnalyticsConfigurationPropertiesArgs) ToTrafficAnalyticsConfigurationPropertiesOutputWithContext(ctx context.Context) TrafficAnalyticsConfigurationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficAnalyticsConfigurationPropertiesOutput)
}

func (i TrafficAnalyticsConfigurationPropertiesArgs) ToTrafficAnalyticsConfigurationPropertiesPtrOutput() TrafficAnalyticsConfigurationPropertiesPtrOutput {
	return i.ToTrafficAnalyticsConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (i TrafficAnalyticsConfigurationPropertiesArgs) ToTrafficAnalyticsConfigurationPropertiesPtrOutputWithContext(ctx context.Context) TrafficAnalyticsConfigurationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficAnalyticsConfigurationPropertiesOutput).ToTrafficAnalyticsConfigurationPropertiesPtrOutputWithContext(ctx)
}









type TrafficAnalyticsConfigurationPropertiesPtrInput interface {
	pulumi.Input

	ToTrafficAnalyticsConfigurationPropertiesPtrOutput() TrafficAnalyticsConfigurationPropertiesPtrOutput
	ToTrafficAnalyticsConfigurationPropertiesPtrOutputWithContext(context.Context) TrafficAnalyticsConfigurationPropertiesPtrOutput
}

type trafficAnalyticsConfigurationPropertiesPtrType TrafficAnalyticsConfigurationPropertiesArgs

func TrafficAnalyticsConfigurationPropertiesPtr(v *TrafficAnalyticsConfigurationPropertiesArgs) TrafficAnalyticsConfigurationPropertiesPtrInput {
	return (*trafficAnalyticsConfigurationPropertiesPtrType)(v)
}

func (*trafficAnalyticsConfigurationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficAnalyticsConfigurationProperties)(nil)).Elem()
}

func (i *trafficAnalyticsConfigurationPropertiesPtrType) ToTrafficAnalyticsConfigurationPropertiesPtrOutput() TrafficAnalyticsConfigurationPropertiesPtrOutput {
	return i.ToTrafficAnalyticsConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (i *trafficAnalyticsConfigurationPropertiesPtrType) ToTrafficAnalyticsConfigurationPropertiesPtrOutputWithContext(ctx context.Context) TrafficAnalyticsConfigurationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficAnalyticsConfigurationPropertiesPtrOutput)
}

type TrafficAnalyticsConfigurationPropertiesOutput struct{ *pulumi.OutputState }

func (TrafficAnalyticsConfigurationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficAnalyticsConfigurationProperties)(nil)).Elem()
}

func (o TrafficAnalyticsConfigurationPropertiesOutput) ToTrafficAnalyticsConfigurationPropertiesOutput() TrafficAnalyticsConfigurationPropertiesOutput {
	return o
}

func (o TrafficAnalyticsConfigurationPropertiesOutput) ToTrafficAnalyticsConfigurationPropertiesOutputWithContext(ctx context.Context) TrafficAnalyticsConfigurationPropertiesOutput {
	return o
}

func (o TrafficAnalyticsConfigurationPropertiesOutput) ToTrafficAnalyticsConfigurationPropertiesPtrOutput() TrafficAnalyticsConfigurationPropertiesPtrOutput {
	return o.ToTrafficAnalyticsConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (o TrafficAnalyticsConfigurationPropertiesOutput) ToTrafficAnalyticsConfigurationPropertiesPtrOutputWithContext(ctx context.Context) TrafficAnalyticsConfigurationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrafficAnalyticsConfigurationProperties) *TrafficAnalyticsConfigurationProperties {
		return &v
	}).(TrafficAnalyticsConfigurationPropertiesPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsConfigurationProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesOutput) TrafficAnalyticsInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsConfigurationProperties) *int { return v.TrafficAnalyticsInterval }).(pulumi.IntPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsConfigurationProperties) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesOutput) WorkspaceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsConfigurationProperties) *string { return v.WorkspaceRegion }).(pulumi.StringPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsConfigurationProperties) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type TrafficAnalyticsConfigurationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TrafficAnalyticsConfigurationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficAnalyticsConfigurationProperties)(nil)).Elem()
}

func (o TrafficAnalyticsConfigurationPropertiesPtrOutput) ToTrafficAnalyticsConfigurationPropertiesPtrOutput() TrafficAnalyticsConfigurationPropertiesPtrOutput {
	return o
}

func (o TrafficAnalyticsConfigurationPropertiesPtrOutput) ToTrafficAnalyticsConfigurationPropertiesPtrOutputWithContext(ctx context.Context) TrafficAnalyticsConfigurationPropertiesPtrOutput {
	return o
}

func (o TrafficAnalyticsConfigurationPropertiesPtrOutput) Elem() TrafficAnalyticsConfigurationPropertiesOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationProperties) TrafficAnalyticsConfigurationProperties {
		if v != nil {
			return *v
		}
		var ret TrafficAnalyticsConfigurationProperties
		return ret
	}).(TrafficAnalyticsConfigurationPropertiesOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesPtrOutput) TrafficAnalyticsInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationProperties) *int {
		if v == nil {
			return nil
		}
		return v.TrafficAnalyticsInterval
	}).(pulumi.IntPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesPtrOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceId
	}).(pulumi.StringPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesPtrOutput) WorkspaceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceRegion
	}).(pulumi.StringPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesPtrOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceResourceId
	}).(pulumi.StringPtrOutput)
}

type TrafficAnalyticsConfigurationPropertiesResponse struct {
	Enabled                  *bool   `pulumi:"enabled"`
	TrafficAnalyticsInterval *int    `pulumi:"trafficAnalyticsInterval"`
	WorkspaceId              *string `pulumi:"workspaceId"`
	WorkspaceRegion          *string `pulumi:"workspaceRegion"`
	WorkspaceResourceId      *string `pulumi:"workspaceResourceId"`
}

type TrafficAnalyticsConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TrafficAnalyticsConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficAnalyticsConfigurationPropertiesResponse)(nil)).Elem()
}

func (o TrafficAnalyticsConfigurationPropertiesResponseOutput) ToTrafficAnalyticsConfigurationPropertiesResponseOutput() TrafficAnalyticsConfigurationPropertiesResponseOutput {
	return o
}

func (o TrafficAnalyticsConfigurationPropertiesResponseOutput) ToTrafficAnalyticsConfigurationPropertiesResponseOutputWithContext(ctx context.Context) TrafficAnalyticsConfigurationPropertiesResponseOutput {
	return o
}

func (o TrafficAnalyticsConfigurationPropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsConfigurationPropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesResponseOutput) TrafficAnalyticsInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsConfigurationPropertiesResponse) *int { return v.TrafficAnalyticsInterval }).(pulumi.IntPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesResponseOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsConfigurationPropertiesResponse) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesResponseOutput) WorkspaceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsConfigurationPropertiesResponse) *string { return v.WorkspaceRegion }).(pulumi.StringPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesResponseOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsConfigurationPropertiesResponse) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type TrafficAnalyticsConfigurationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TrafficAnalyticsConfigurationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficAnalyticsConfigurationPropertiesResponse)(nil)).Elem()
}

func (o TrafficAnalyticsConfigurationPropertiesResponsePtrOutput) ToTrafficAnalyticsConfigurationPropertiesResponsePtrOutput() TrafficAnalyticsConfigurationPropertiesResponsePtrOutput {
	return o
}

func (o TrafficAnalyticsConfigurationPropertiesResponsePtrOutput) ToTrafficAnalyticsConfigurationPropertiesResponsePtrOutputWithContext(ctx context.Context) TrafficAnalyticsConfigurationPropertiesResponsePtrOutput {
	return o
}

func (o TrafficAnalyticsConfigurationPropertiesResponsePtrOutput) Elem() TrafficAnalyticsConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationPropertiesResponse) TrafficAnalyticsConfigurationPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TrafficAnalyticsConfigurationPropertiesResponse
		return ret
	}).(TrafficAnalyticsConfigurationPropertiesResponseOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesResponsePtrOutput) TrafficAnalyticsInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.TrafficAnalyticsInterval
	}).(pulumi.IntPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesResponsePtrOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceId
	}).(pulumi.StringPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesResponsePtrOutput) WorkspaceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceRegion
	}).(pulumi.StringPtrOutput)
}

func (o TrafficAnalyticsConfigurationPropertiesResponsePtrOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsConfigurationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceResourceId
	}).(pulumi.StringPtrOutput)
}

type TrafficAnalyticsProperties struct {
	NetworkWatcherFlowAnalyticsConfiguration *TrafficAnalyticsConfigurationProperties `pulumi:"networkWatcherFlowAnalyticsConfiguration"`
}





type TrafficAnalyticsPropertiesInput interface {
	pulumi.Input

	ToTrafficAnalyticsPropertiesOutput() TrafficAnalyticsPropertiesOutput
	ToTrafficAnalyticsPropertiesOutputWithContext(context.Context) TrafficAnalyticsPropertiesOutput
}

type TrafficAnalyticsPropertiesArgs struct {
	NetworkWatcherFlowAnalyticsConfiguration TrafficAnalyticsConfigurationPropertiesPtrInput `pulumi:"networkWatcherFlowAnalyticsConfiguration"`
}

func (TrafficAnalyticsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficAnalyticsProperties)(nil)).Elem()
}

func (i TrafficAnalyticsPropertiesArgs) ToTrafficAnalyticsPropertiesOutput() TrafficAnalyticsPropertiesOutput {
	return i.ToTrafficAnalyticsPropertiesOutputWithContext(context.Background())
}

func (i TrafficAnalyticsPropertiesArgs) ToTrafficAnalyticsPropertiesOutputWithContext(ctx context.Context) TrafficAnalyticsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficAnalyticsPropertiesOutput)
}

func (i TrafficAnalyticsPropertiesArgs) ToTrafficAnalyticsPropertiesPtrOutput() TrafficAnalyticsPropertiesPtrOutput {
	return i.ToTrafficAnalyticsPropertiesPtrOutputWithContext(context.Background())
}

func (i TrafficAnalyticsPropertiesArgs) ToTrafficAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) TrafficAnalyticsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficAnalyticsPropertiesOutput).ToTrafficAnalyticsPropertiesPtrOutputWithContext(ctx)
}









type TrafficAnalyticsPropertiesPtrInput interface {
	pulumi.Input

	ToTrafficAnalyticsPropertiesPtrOutput() TrafficAnalyticsPropertiesPtrOutput
	ToTrafficAnalyticsPropertiesPtrOutputWithContext(context.Context) TrafficAnalyticsPropertiesPtrOutput
}

type trafficAnalyticsPropertiesPtrType TrafficAnalyticsPropertiesArgs

func TrafficAnalyticsPropertiesPtr(v *TrafficAnalyticsPropertiesArgs) TrafficAnalyticsPropertiesPtrInput {
	return (*trafficAnalyticsPropertiesPtrType)(v)
}

func (*trafficAnalyticsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficAnalyticsProperties)(nil)).Elem()
}

func (i *trafficAnalyticsPropertiesPtrType) ToTrafficAnalyticsPropertiesPtrOutput() TrafficAnalyticsPropertiesPtrOutput {
	return i.ToTrafficAnalyticsPropertiesPtrOutputWithContext(context.Background())
}

func (i *trafficAnalyticsPropertiesPtrType) ToTrafficAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) TrafficAnalyticsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficAnalyticsPropertiesPtrOutput)
}

type TrafficAnalyticsPropertiesOutput struct{ *pulumi.OutputState }

func (TrafficAnalyticsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficAnalyticsProperties)(nil)).Elem()
}

func (o TrafficAnalyticsPropertiesOutput) ToTrafficAnalyticsPropertiesOutput() TrafficAnalyticsPropertiesOutput {
	return o
}

func (o TrafficAnalyticsPropertiesOutput) ToTrafficAnalyticsPropertiesOutputWithContext(ctx context.Context) TrafficAnalyticsPropertiesOutput {
	return o
}

func (o TrafficAnalyticsPropertiesOutput) ToTrafficAnalyticsPropertiesPtrOutput() TrafficAnalyticsPropertiesPtrOutput {
	return o.ToTrafficAnalyticsPropertiesPtrOutputWithContext(context.Background())
}

func (o TrafficAnalyticsPropertiesOutput) ToTrafficAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) TrafficAnalyticsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrafficAnalyticsProperties) *TrafficAnalyticsProperties {
		return &v
	}).(TrafficAnalyticsPropertiesPtrOutput)
}

func (o TrafficAnalyticsPropertiesOutput) NetworkWatcherFlowAnalyticsConfiguration() TrafficAnalyticsConfigurationPropertiesPtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsProperties) *TrafficAnalyticsConfigurationProperties {
		return v.NetworkWatcherFlowAnalyticsConfiguration
	}).(TrafficAnalyticsConfigurationPropertiesPtrOutput)
}

type TrafficAnalyticsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TrafficAnalyticsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficAnalyticsProperties)(nil)).Elem()
}

func (o TrafficAnalyticsPropertiesPtrOutput) ToTrafficAnalyticsPropertiesPtrOutput() TrafficAnalyticsPropertiesPtrOutput {
	return o
}

func (o TrafficAnalyticsPropertiesPtrOutput) ToTrafficAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) TrafficAnalyticsPropertiesPtrOutput {
	return o
}

func (o TrafficAnalyticsPropertiesPtrOutput) Elem() TrafficAnalyticsPropertiesOutput {
	return o.ApplyT(func(v *TrafficAnalyticsProperties) TrafficAnalyticsProperties {
		if v != nil {
			return *v
		}
		var ret TrafficAnalyticsProperties
		return ret
	}).(TrafficAnalyticsPropertiesOutput)
}

func (o TrafficAnalyticsPropertiesPtrOutput) NetworkWatcherFlowAnalyticsConfiguration() TrafficAnalyticsConfigurationPropertiesPtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsProperties) *TrafficAnalyticsConfigurationProperties {
		if v == nil {
			return nil
		}
		return v.NetworkWatcherFlowAnalyticsConfiguration
	}).(TrafficAnalyticsConfigurationPropertiesPtrOutput)
}

type TrafficAnalyticsPropertiesResponse struct {
	NetworkWatcherFlowAnalyticsConfiguration *TrafficAnalyticsConfigurationPropertiesResponse `pulumi:"networkWatcherFlowAnalyticsConfiguration"`
}

type TrafficAnalyticsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TrafficAnalyticsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficAnalyticsPropertiesResponse)(nil)).Elem()
}

func (o TrafficAnalyticsPropertiesResponseOutput) ToTrafficAnalyticsPropertiesResponseOutput() TrafficAnalyticsPropertiesResponseOutput {
	return o
}

func (o TrafficAnalyticsPropertiesResponseOutput) ToTrafficAnalyticsPropertiesResponseOutputWithContext(ctx context.Context) TrafficAnalyticsPropertiesResponseOutput {
	return o
}

func (o TrafficAnalyticsPropertiesResponseOutput) NetworkWatcherFlowAnalyticsConfiguration() TrafficAnalyticsConfigurationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v TrafficAnalyticsPropertiesResponse) *TrafficAnalyticsConfigurationPropertiesResponse {
		return v.NetworkWatcherFlowAnalyticsConfiguration
	}).(TrafficAnalyticsConfigurationPropertiesResponsePtrOutput)
}

type TrafficAnalyticsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TrafficAnalyticsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficAnalyticsPropertiesResponse)(nil)).Elem()
}

func (o TrafficAnalyticsPropertiesResponsePtrOutput) ToTrafficAnalyticsPropertiesResponsePtrOutput() TrafficAnalyticsPropertiesResponsePtrOutput {
	return o
}

func (o TrafficAnalyticsPropertiesResponsePtrOutput) ToTrafficAnalyticsPropertiesResponsePtrOutputWithContext(ctx context.Context) TrafficAnalyticsPropertiesResponsePtrOutput {
	return o
}

func (o TrafficAnalyticsPropertiesResponsePtrOutput) Elem() TrafficAnalyticsPropertiesResponseOutput {
	return o.ApplyT(func(v *TrafficAnalyticsPropertiesResponse) TrafficAnalyticsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TrafficAnalyticsPropertiesResponse
		return ret
	}).(TrafficAnalyticsPropertiesResponseOutput)
}

func (o TrafficAnalyticsPropertiesResponsePtrOutput) NetworkWatcherFlowAnalyticsConfiguration() TrafficAnalyticsConfigurationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *TrafficAnalyticsPropertiesResponse) *TrafficAnalyticsConfigurationPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.NetworkWatcherFlowAnalyticsConfiguration
	}).(TrafficAnalyticsConfigurationPropertiesResponsePtrOutput)
}

type TrafficSelectorPolicy struct {
	LocalAddressRanges  []string `pulumi:"localAddressRanges"`
	RemoteAddressRanges []string `pulumi:"remoteAddressRanges"`
}





type TrafficSelectorPolicyInput interface {
	pulumi.Input

	ToTrafficSelectorPolicyOutput() TrafficSelectorPolicyOutput
	ToTrafficSelectorPolicyOutputWithContext(context.Context) TrafficSelectorPolicyOutput
}

type TrafficSelectorPolicyArgs struct {
	LocalAddressRanges  pulumi.StringArrayInput `pulumi:"localAddressRanges"`
	RemoteAddressRanges pulumi.StringArrayInput `pulumi:"remoteAddressRanges"`
}

func (TrafficSelectorPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficSelectorPolicy)(nil)).Elem()
}

func (i TrafficSelectorPolicyArgs) ToTrafficSelectorPolicyOutput() TrafficSelectorPolicyOutput {
	return i.ToTrafficSelectorPolicyOutputWithContext(context.Background())
}

func (i TrafficSelectorPolicyArgs) ToTrafficSelectorPolicyOutputWithContext(ctx context.Context) TrafficSelectorPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficSelectorPolicyOutput)
}





type TrafficSelectorPolicyArrayInput interface {
	pulumi.Input

	ToTrafficSelectorPolicyArrayOutput() TrafficSelectorPolicyArrayOutput
	ToTrafficSelectorPolicyArrayOutputWithContext(context.Context) TrafficSelectorPolicyArrayOutput
}

type TrafficSelectorPolicyArray []TrafficSelectorPolicyInput

func (TrafficSelectorPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrafficSelectorPolicy)(nil)).Elem()
}

func (i TrafficSelectorPolicyArray) ToTrafficSelectorPolicyArrayOutput() TrafficSelectorPolicyArrayOutput {
	return i.ToTrafficSelectorPolicyArrayOutputWithContext(context.Background())
}

func (i TrafficSelectorPolicyArray) ToTrafficSelectorPolicyArrayOutputWithContext(ctx context.Context) TrafficSelectorPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficSelectorPolicyArrayOutput)
}

type TrafficSelectorPolicyOutput struct{ *pulumi.OutputState }

func (TrafficSelectorPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficSelectorPolicy)(nil)).Elem()
}

func (o TrafficSelectorPolicyOutput) ToTrafficSelectorPolicyOutput() TrafficSelectorPolicyOutput {
	return o
}

func (o TrafficSelectorPolicyOutput) ToTrafficSelectorPolicyOutputWithContext(ctx context.Context) TrafficSelectorPolicyOutput {
	return o
}

func (o TrafficSelectorPolicyOutput) LocalAddressRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TrafficSelectorPolicy) []string { return v.LocalAddressRanges }).(pulumi.StringArrayOutput)
}

func (o TrafficSelectorPolicyOutput) RemoteAddressRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TrafficSelectorPolicy) []string { return v.RemoteAddressRanges }).(pulumi.StringArrayOutput)
}

type TrafficSelectorPolicyArrayOutput struct{ *pulumi.OutputState }

func (TrafficSelectorPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrafficSelectorPolicy)(nil)).Elem()
}

func (o TrafficSelectorPolicyArrayOutput) ToTrafficSelectorPolicyArrayOutput() TrafficSelectorPolicyArrayOutput {
	return o
}

func (o TrafficSelectorPolicyArrayOutput) ToTrafficSelectorPolicyArrayOutputWithContext(ctx context.Context) TrafficSelectorPolicyArrayOutput {
	return o
}

func (o TrafficSelectorPolicyArrayOutput) Index(i pulumi.IntInput) TrafficSelectorPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrafficSelectorPolicy {
		return vs[0].([]TrafficSelectorPolicy)[vs[1].(int)]
	}).(TrafficSelectorPolicyOutput)
}

type TrafficSelectorPolicyResponse struct {
	LocalAddressRanges  []string `pulumi:"localAddressRanges"`
	RemoteAddressRanges []string `pulumi:"remoteAddressRanges"`
}

type TrafficSelectorPolicyResponseOutput struct{ *pulumi.OutputState }

func (TrafficSelectorPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficSelectorPolicyResponse)(nil)).Elem()
}

func (o TrafficSelectorPolicyResponseOutput) ToTrafficSelectorPolicyResponseOutput() TrafficSelectorPolicyResponseOutput {
	return o
}

func (o TrafficSelectorPolicyResponseOutput) ToTrafficSelectorPolicyResponseOutputWithContext(ctx context.Context) TrafficSelectorPolicyResponseOutput {
	return o
}

func (o TrafficSelectorPolicyResponseOutput) LocalAddressRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TrafficSelectorPolicyResponse) []string { return v.LocalAddressRanges }).(pulumi.StringArrayOutput)
}

func (o TrafficSelectorPolicyResponseOutput) RemoteAddressRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TrafficSelectorPolicyResponse) []string { return v.RemoteAddressRanges }).(pulumi.StringArrayOutput)
}

type TrafficSelectorPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (TrafficSelectorPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrafficSelectorPolicyResponse)(nil)).Elem()
}

func (o TrafficSelectorPolicyResponseArrayOutput) ToTrafficSelectorPolicyResponseArrayOutput() TrafficSelectorPolicyResponseArrayOutput {
	return o
}

func (o TrafficSelectorPolicyResponseArrayOutput) ToTrafficSelectorPolicyResponseArrayOutputWithContext(ctx context.Context) TrafficSelectorPolicyResponseArrayOutput {
	return o
}

func (o TrafficSelectorPolicyResponseArrayOutput) Index(i pulumi.IntInput) TrafficSelectorPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrafficSelectorPolicyResponse {
		return vs[0].([]TrafficSelectorPolicyResponse)[vs[1].(int)]
	}).(TrafficSelectorPolicyResponseOutput)
}

type TunnelConnectionHealthResponse struct {
	ConnectionStatus                 string  `pulumi:"connectionStatus"`
	EgressBytesTransferred           float64 `pulumi:"egressBytesTransferred"`
	IngressBytesTransferred          float64 `pulumi:"ingressBytesTransferred"`
	LastConnectionEstablishedUtcTime string  `pulumi:"lastConnectionEstablishedUtcTime"`
	Tunnel                           string  `pulumi:"tunnel"`
}

type TunnelConnectionHealthResponseOutput struct{ *pulumi.OutputState }

func (TunnelConnectionHealthResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TunnelConnectionHealthResponse)(nil)).Elem()
}

func (o TunnelConnectionHealthResponseOutput) ToTunnelConnectionHealthResponseOutput() TunnelConnectionHealthResponseOutput {
	return o
}

func (o TunnelConnectionHealthResponseOutput) ToTunnelConnectionHealthResponseOutputWithContext(ctx context.Context) TunnelConnectionHealthResponseOutput {
	return o
}

func (o TunnelConnectionHealthResponseOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o TunnelConnectionHealthResponseOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o TunnelConnectionHealthResponseOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o TunnelConnectionHealthResponseOutput) LastConnectionEstablishedUtcTime() pulumi.StringOutput {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) string { return v.LastConnectionEstablishedUtcTime }).(pulumi.StringOutput)
}

func (o TunnelConnectionHealthResponseOutput) Tunnel() pulumi.StringOutput {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) string { return v.Tunnel }).(pulumi.StringOutput)
}

type TunnelConnectionHealthResponseArrayOutput struct{ *pulumi.OutputState }

func (TunnelConnectionHealthResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TunnelConnectionHealthResponse)(nil)).Elem()
}

func (o TunnelConnectionHealthResponseArrayOutput) ToTunnelConnectionHealthResponseArrayOutput() TunnelConnectionHealthResponseArrayOutput {
	return o
}

func (o TunnelConnectionHealthResponseArrayOutput) ToTunnelConnectionHealthResponseArrayOutputWithContext(ctx context.Context) TunnelConnectionHealthResponseArrayOutput {
	return o
}

func (o TunnelConnectionHealthResponseArrayOutput) Index(i pulumi.IntInput) TunnelConnectionHealthResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TunnelConnectionHealthResponse {
		return vs[0].([]TunnelConnectionHealthResponse)[vs[1].(int)]
	}).(TunnelConnectionHealthResponseOutput)
}

type TxtRecord struct {
	Value []string `pulumi:"value"`
}





type TxtRecordInput interface {
	pulumi.Input

	ToTxtRecordOutput() TxtRecordOutput
	ToTxtRecordOutputWithContext(context.Context) TxtRecordOutput
}

type TxtRecordArgs struct {
	Value pulumi.StringArrayInput `pulumi:"value"`
}

func (TxtRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecord)(nil)).Elem()
}

func (i TxtRecordArgs) ToTxtRecordOutput() TxtRecordOutput {
	return i.ToTxtRecordOutputWithContext(context.Background())
}

func (i TxtRecordArgs) ToTxtRecordOutputWithContext(ctx context.Context) TxtRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordOutput)
}





type TxtRecordArrayInput interface {
	pulumi.Input

	ToTxtRecordArrayOutput() TxtRecordArrayOutput
	ToTxtRecordArrayOutputWithContext(context.Context) TxtRecordArrayOutput
}

type TxtRecordArray []TxtRecordInput

func (TxtRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecord)(nil)).Elem()
}

func (i TxtRecordArray) ToTxtRecordArrayOutput() TxtRecordArrayOutput {
	return i.ToTxtRecordArrayOutputWithContext(context.Background())
}

func (i TxtRecordArray) ToTxtRecordArrayOutputWithContext(ctx context.Context) TxtRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordArrayOutput)
}

type TxtRecordOutput struct{ *pulumi.OutputState }

func (TxtRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecord)(nil)).Elem()
}

func (o TxtRecordOutput) ToTxtRecordOutput() TxtRecordOutput {
	return o
}

func (o TxtRecordOutput) ToTxtRecordOutputWithContext(ctx context.Context) TxtRecordOutput {
	return o
}

func (o TxtRecordOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TxtRecord) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type TxtRecordArrayOutput struct{ *pulumi.OutputState }

func (TxtRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecord)(nil)).Elem()
}

func (o TxtRecordArrayOutput) ToTxtRecordArrayOutput() TxtRecordArrayOutput {
	return o
}

func (o TxtRecordArrayOutput) ToTxtRecordArrayOutputWithContext(ctx context.Context) TxtRecordArrayOutput {
	return o
}

func (o TxtRecordArrayOutput) Index(i pulumi.IntInput) TxtRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TxtRecord {
		return vs[0].([]TxtRecord)[vs[1].(int)]
	}).(TxtRecordOutput)
}

type TxtRecordResponse struct {
	Value []string `pulumi:"value"`
}

type TxtRecordResponseOutput struct{ *pulumi.OutputState }

func (TxtRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecordResponse)(nil)).Elem()
}

func (o TxtRecordResponseOutput) ToTxtRecordResponseOutput() TxtRecordResponseOutput {
	return o
}

func (o TxtRecordResponseOutput) ToTxtRecordResponseOutputWithContext(ctx context.Context) TxtRecordResponseOutput {
	return o
}

func (o TxtRecordResponseOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TxtRecordResponse) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type TxtRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (TxtRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecordResponse)(nil)).Elem()
}

func (o TxtRecordResponseArrayOutput) ToTxtRecordResponseArrayOutput() TxtRecordResponseArrayOutput {
	return o
}

func (o TxtRecordResponseArrayOutput) ToTxtRecordResponseArrayOutputWithContext(ctx context.Context) TxtRecordResponseArrayOutput {
	return o
}

func (o TxtRecordResponseArrayOutput) Index(i pulumi.IntInput) TxtRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TxtRecordResponse {
		return vs[0].([]TxtRecordResponse)[vs[1].(int)]
	}).(TxtRecordResponseOutput)
}

type VM struct {
	Id       *string           `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Tags     map[string]string `pulumi:"tags"`
}





type VMInput interface {
	pulumi.Input

	ToVMOutput() VMOutput
	ToVMOutputWithContext(context.Context) VMOutput
}

type VMArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringPtrInput `pulumi:"location"`
	Tags     pulumi.StringMapInput `pulumi:"tags"`
}

func (VMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VM)(nil)).Elem()
}

func (i VMArgs) ToVMOutput() VMOutput {
	return i.ToVMOutputWithContext(context.Background())
}

func (i VMArgs) ToVMOutputWithContext(ctx context.Context) VMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMOutput)
}

type VMOutput struct{ *pulumi.OutputState }

func (VMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VM)(nil)).Elem()
}

func (o VMOutput) ToVMOutput() VMOutput {
	return o
}

func (o VMOutput) ToVMOutputWithContext(ctx context.Context) VMOutput {
	return o
}

func (o VMOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VM) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VMOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VM) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VMOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VM) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type VMResponse struct {
	Id       *string           `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}

type VMResponseOutput struct{ *pulumi.OutputState }

func (VMResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMResponse)(nil)).Elem()
}

func (o VMResponseOutput) ToVMResponseOutput() VMResponseOutput {
	return o
}

func (o VMResponseOutput) ToVMResponseOutputWithContext(ctx context.Context) VMResponseOutput {
	return o
}

func (o VMResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VMResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VMResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VMResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VMResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VMResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VMResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VMResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VirtualApplianceNicPropertiesResponse struct {
	Name             string `pulumi:"name"`
	PrivateIpAddress string `pulumi:"privateIpAddress"`
	PublicIpAddress  string `pulumi:"publicIpAddress"`
}

type VirtualApplianceNicPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VirtualApplianceNicPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplianceNicPropertiesResponse)(nil)).Elem()
}

func (o VirtualApplianceNicPropertiesResponseOutput) ToVirtualApplianceNicPropertiesResponseOutput() VirtualApplianceNicPropertiesResponseOutput {
	return o
}

func (o VirtualApplianceNicPropertiesResponseOutput) ToVirtualApplianceNicPropertiesResponseOutputWithContext(ctx context.Context) VirtualApplianceNicPropertiesResponseOutput {
	return o
}

func (o VirtualApplianceNicPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualApplianceNicPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualApplianceNicPropertiesResponseOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualApplianceNicPropertiesResponse) string { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

func (o VirtualApplianceNicPropertiesResponseOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualApplianceNicPropertiesResponse) string { return v.PublicIpAddress }).(pulumi.StringOutput)
}

type VirtualApplianceNicPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualApplianceNicPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualApplianceNicPropertiesResponse)(nil)).Elem()
}

func (o VirtualApplianceNicPropertiesResponseArrayOutput) ToVirtualApplianceNicPropertiesResponseArrayOutput() VirtualApplianceNicPropertiesResponseArrayOutput {
	return o
}

func (o VirtualApplianceNicPropertiesResponseArrayOutput) ToVirtualApplianceNicPropertiesResponseArrayOutputWithContext(ctx context.Context) VirtualApplianceNicPropertiesResponseArrayOutput {
	return o
}

func (o VirtualApplianceNicPropertiesResponseArrayOutput) Index(i pulumi.IntInput) VirtualApplianceNicPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualApplianceNicPropertiesResponse {
		return vs[0].([]VirtualApplianceNicPropertiesResponse)[vs[1].(int)]
	}).(VirtualApplianceNicPropertiesResponseOutput)
}

type VirtualApplianceSkuProperties struct {
	BundledScaleUnit   *string `pulumi:"bundledScaleUnit"`
	MarketPlaceVersion *string `pulumi:"marketPlaceVersion"`
	Vendor             *string `pulumi:"vendor"`
}





type VirtualApplianceSkuPropertiesInput interface {
	pulumi.Input

	ToVirtualApplianceSkuPropertiesOutput() VirtualApplianceSkuPropertiesOutput
	ToVirtualApplianceSkuPropertiesOutputWithContext(context.Context) VirtualApplianceSkuPropertiesOutput
}

type VirtualApplianceSkuPropertiesArgs struct {
	BundledScaleUnit   pulumi.StringPtrInput `pulumi:"bundledScaleUnit"`
	MarketPlaceVersion pulumi.StringPtrInput `pulumi:"marketPlaceVersion"`
	Vendor             pulumi.StringPtrInput `pulumi:"vendor"`
}

func (VirtualApplianceSkuPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplianceSkuProperties)(nil)).Elem()
}

func (i VirtualApplianceSkuPropertiesArgs) ToVirtualApplianceSkuPropertiesOutput() VirtualApplianceSkuPropertiesOutput {
	return i.ToVirtualApplianceSkuPropertiesOutputWithContext(context.Background())
}

func (i VirtualApplianceSkuPropertiesArgs) ToVirtualApplianceSkuPropertiesOutputWithContext(ctx context.Context) VirtualApplianceSkuPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplianceSkuPropertiesOutput)
}

func (i VirtualApplianceSkuPropertiesArgs) ToVirtualApplianceSkuPropertiesPtrOutput() VirtualApplianceSkuPropertiesPtrOutput {
	return i.ToVirtualApplianceSkuPropertiesPtrOutputWithContext(context.Background())
}

func (i VirtualApplianceSkuPropertiesArgs) ToVirtualApplianceSkuPropertiesPtrOutputWithContext(ctx context.Context) VirtualApplianceSkuPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplianceSkuPropertiesOutput).ToVirtualApplianceSkuPropertiesPtrOutputWithContext(ctx)
}









type VirtualApplianceSkuPropertiesPtrInput interface {
	pulumi.Input

	ToVirtualApplianceSkuPropertiesPtrOutput() VirtualApplianceSkuPropertiesPtrOutput
	ToVirtualApplianceSkuPropertiesPtrOutputWithContext(context.Context) VirtualApplianceSkuPropertiesPtrOutput
}

type virtualApplianceSkuPropertiesPtrType VirtualApplianceSkuPropertiesArgs

func VirtualApplianceSkuPropertiesPtr(v *VirtualApplianceSkuPropertiesArgs) VirtualApplianceSkuPropertiesPtrInput {
	return (*virtualApplianceSkuPropertiesPtrType)(v)
}

func (*virtualApplianceSkuPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualApplianceSkuProperties)(nil)).Elem()
}

func (i *virtualApplianceSkuPropertiesPtrType) ToVirtualApplianceSkuPropertiesPtrOutput() VirtualApplianceSkuPropertiesPtrOutput {
	return i.ToVirtualApplianceSkuPropertiesPtrOutputWithContext(context.Background())
}

func (i *virtualApplianceSkuPropertiesPtrType) ToVirtualApplianceSkuPropertiesPtrOutputWithContext(ctx context.Context) VirtualApplianceSkuPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplianceSkuPropertiesPtrOutput)
}

type VirtualApplianceSkuPropertiesOutput struct{ *pulumi.OutputState }

func (VirtualApplianceSkuPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplianceSkuProperties)(nil)).Elem()
}

func (o VirtualApplianceSkuPropertiesOutput) ToVirtualApplianceSkuPropertiesOutput() VirtualApplianceSkuPropertiesOutput {
	return o
}

func (o VirtualApplianceSkuPropertiesOutput) ToVirtualApplianceSkuPropertiesOutputWithContext(ctx context.Context) VirtualApplianceSkuPropertiesOutput {
	return o
}

func (o VirtualApplianceSkuPropertiesOutput) ToVirtualApplianceSkuPropertiesPtrOutput() VirtualApplianceSkuPropertiesPtrOutput {
	return o.ToVirtualApplianceSkuPropertiesPtrOutputWithContext(context.Background())
}

func (o VirtualApplianceSkuPropertiesOutput) ToVirtualApplianceSkuPropertiesPtrOutputWithContext(ctx context.Context) VirtualApplianceSkuPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualApplianceSkuProperties) *VirtualApplianceSkuProperties {
		return &v
	}).(VirtualApplianceSkuPropertiesPtrOutput)
}

func (o VirtualApplianceSkuPropertiesOutput) BundledScaleUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplianceSkuProperties) *string { return v.BundledScaleUnit }).(pulumi.StringPtrOutput)
}

func (o VirtualApplianceSkuPropertiesOutput) MarketPlaceVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplianceSkuProperties) *string { return v.MarketPlaceVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualApplianceSkuPropertiesOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplianceSkuProperties) *string { return v.Vendor }).(pulumi.StringPtrOutput)
}

type VirtualApplianceSkuPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VirtualApplianceSkuPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualApplianceSkuProperties)(nil)).Elem()
}

func (o VirtualApplianceSkuPropertiesPtrOutput) ToVirtualApplianceSkuPropertiesPtrOutput() VirtualApplianceSkuPropertiesPtrOutput {
	return o
}

func (o VirtualApplianceSkuPropertiesPtrOutput) ToVirtualApplianceSkuPropertiesPtrOutputWithContext(ctx context.Context) VirtualApplianceSkuPropertiesPtrOutput {
	return o
}

func (o VirtualApplianceSkuPropertiesPtrOutput) Elem() VirtualApplianceSkuPropertiesOutput {
	return o.ApplyT(func(v *VirtualApplianceSkuProperties) VirtualApplianceSkuProperties {
		if v != nil {
			return *v
		}
		var ret VirtualApplianceSkuProperties
		return ret
	}).(VirtualApplianceSkuPropertiesOutput)
}

func (o VirtualApplianceSkuPropertiesPtrOutput) BundledScaleUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualApplianceSkuProperties) *string {
		if v == nil {
			return nil
		}
		return v.BundledScaleUnit
	}).(pulumi.StringPtrOutput)
}

func (o VirtualApplianceSkuPropertiesPtrOutput) MarketPlaceVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualApplianceSkuProperties) *string {
		if v == nil {
			return nil
		}
		return v.MarketPlaceVersion
	}).(pulumi.StringPtrOutput)
}

func (o VirtualApplianceSkuPropertiesPtrOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualApplianceSkuProperties) *string {
		if v == nil {
			return nil
		}
		return v.Vendor
	}).(pulumi.StringPtrOutput)
}

type VirtualApplianceSkuPropertiesResponse struct {
	BundledScaleUnit   *string `pulumi:"bundledScaleUnit"`
	MarketPlaceVersion *string `pulumi:"marketPlaceVersion"`
	Vendor             *string `pulumi:"vendor"`
}

type VirtualApplianceSkuPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VirtualApplianceSkuPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplianceSkuPropertiesResponse)(nil)).Elem()
}

func (o VirtualApplianceSkuPropertiesResponseOutput) ToVirtualApplianceSkuPropertiesResponseOutput() VirtualApplianceSkuPropertiesResponseOutput {
	return o
}

func (o VirtualApplianceSkuPropertiesResponseOutput) ToVirtualApplianceSkuPropertiesResponseOutputWithContext(ctx context.Context) VirtualApplianceSkuPropertiesResponseOutput {
	return o
}

func (o VirtualApplianceSkuPropertiesResponseOutput) BundledScaleUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplianceSkuPropertiesResponse) *string { return v.BundledScaleUnit }).(pulumi.StringPtrOutput)
}

func (o VirtualApplianceSkuPropertiesResponseOutput) MarketPlaceVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplianceSkuPropertiesResponse) *string { return v.MarketPlaceVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualApplianceSkuPropertiesResponseOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplianceSkuPropertiesResponse) *string { return v.Vendor }).(pulumi.StringPtrOutput)
}

type VirtualApplianceSkuPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualApplianceSkuPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualApplianceSkuPropertiesResponse)(nil)).Elem()
}

func (o VirtualApplianceSkuPropertiesResponsePtrOutput) ToVirtualApplianceSkuPropertiesResponsePtrOutput() VirtualApplianceSkuPropertiesResponsePtrOutput {
	return o
}

func (o VirtualApplianceSkuPropertiesResponsePtrOutput) ToVirtualApplianceSkuPropertiesResponsePtrOutputWithContext(ctx context.Context) VirtualApplianceSkuPropertiesResponsePtrOutput {
	return o
}

func (o VirtualApplianceSkuPropertiesResponsePtrOutput) Elem() VirtualApplianceSkuPropertiesResponseOutput {
	return o.ApplyT(func(v *VirtualApplianceSkuPropertiesResponse) VirtualApplianceSkuPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VirtualApplianceSkuPropertiesResponse
		return ret
	}).(VirtualApplianceSkuPropertiesResponseOutput)
}

func (o VirtualApplianceSkuPropertiesResponsePtrOutput) BundledScaleUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualApplianceSkuPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.BundledScaleUnit
	}).(pulumi.StringPtrOutput)
}

func (o VirtualApplianceSkuPropertiesResponsePtrOutput) MarketPlaceVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualApplianceSkuPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.MarketPlaceVersion
	}).(pulumi.StringPtrOutput)
}

func (o VirtualApplianceSkuPropertiesResponsePtrOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualApplianceSkuPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Vendor
	}).(pulumi.StringPtrOutput)
}

type VirtualHubId struct {
	Id *string `pulumi:"id"`
}





type VirtualHubIdInput interface {
	pulumi.Input

	ToVirtualHubIdOutput() VirtualHubIdOutput
	ToVirtualHubIdOutputWithContext(context.Context) VirtualHubIdOutput
}

type VirtualHubIdArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (VirtualHubIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubId)(nil)).Elem()
}

func (i VirtualHubIdArgs) ToVirtualHubIdOutput() VirtualHubIdOutput {
	return i.ToVirtualHubIdOutputWithContext(context.Background())
}

func (i VirtualHubIdArgs) ToVirtualHubIdOutputWithContext(ctx context.Context) VirtualHubIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubIdOutput)
}

type VirtualHubIdOutput struct{ *pulumi.OutputState }

func (VirtualHubIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubId)(nil)).Elem()
}

func (o VirtualHubIdOutput) ToVirtualHubIdOutput() VirtualHubIdOutput {
	return o
}

func (o VirtualHubIdOutput) ToVirtualHubIdOutputWithContext(ctx context.Context) VirtualHubIdOutput {
	return o
}

func (o VirtualHubIdOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubId) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualHubIdResponse struct {
	Id *string `pulumi:"id"`
}

type VirtualHubIdResponseOutput struct{ *pulumi.OutputState }

func (VirtualHubIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubIdResponse)(nil)).Elem()
}

func (o VirtualHubIdResponseOutput) ToVirtualHubIdResponseOutput() VirtualHubIdResponseOutput {
	return o
}

func (o VirtualHubIdResponseOutput) ToVirtualHubIdResponseOutputWithContext(ctx context.Context) VirtualHubIdResponseOutput {
	return o
}

func (o VirtualHubIdResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubIdResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualHubRoute struct {
	AddressPrefixes  []string `pulumi:"addressPrefixes"`
	NextHopIpAddress *string  `pulumi:"nextHopIpAddress"`
}





type VirtualHubRouteInput interface {
	pulumi.Input

	ToVirtualHubRouteOutput() VirtualHubRouteOutput
	ToVirtualHubRouteOutputWithContext(context.Context) VirtualHubRouteOutput
}

type VirtualHubRouteArgs struct {
	AddressPrefixes  pulumi.StringArrayInput `pulumi:"addressPrefixes"`
	NextHopIpAddress pulumi.StringPtrInput   `pulumi:"nextHopIpAddress"`
}

func (VirtualHubRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRoute)(nil)).Elem()
}

func (i VirtualHubRouteArgs) ToVirtualHubRouteOutput() VirtualHubRouteOutput {
	return i.ToVirtualHubRouteOutputWithContext(context.Background())
}

func (i VirtualHubRouteArgs) ToVirtualHubRouteOutputWithContext(ctx context.Context) VirtualHubRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteOutput)
}





type VirtualHubRouteArrayInput interface {
	pulumi.Input

	ToVirtualHubRouteArrayOutput() VirtualHubRouteArrayOutput
	ToVirtualHubRouteArrayOutputWithContext(context.Context) VirtualHubRouteArrayOutput
}

type VirtualHubRouteArray []VirtualHubRouteInput

func (VirtualHubRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHubRoute)(nil)).Elem()
}

func (i VirtualHubRouteArray) ToVirtualHubRouteArrayOutput() VirtualHubRouteArrayOutput {
	return i.ToVirtualHubRouteArrayOutputWithContext(context.Background())
}

func (i VirtualHubRouteArray) ToVirtualHubRouteArrayOutputWithContext(ctx context.Context) VirtualHubRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteArrayOutput)
}

type VirtualHubRouteOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRoute)(nil)).Elem()
}

func (o VirtualHubRouteOutput) ToVirtualHubRouteOutput() VirtualHubRouteOutput {
	return o
}

func (o VirtualHubRouteOutput) ToVirtualHubRouteOutputWithContext(ctx context.Context) VirtualHubRouteOutput {
	return o
}

func (o VirtualHubRouteOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualHubRoute) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o VirtualHubRouteOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubRoute) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

type VirtualHubRouteArrayOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHubRoute)(nil)).Elem()
}

func (o VirtualHubRouteArrayOutput) ToVirtualHubRouteArrayOutput() VirtualHubRouteArrayOutput {
	return o
}

func (o VirtualHubRouteArrayOutput) ToVirtualHubRouteArrayOutputWithContext(ctx context.Context) VirtualHubRouteArrayOutput {
	return o
}

func (o VirtualHubRouteArrayOutput) Index(i pulumi.IntInput) VirtualHubRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualHubRoute {
		return vs[0].([]VirtualHubRoute)[vs[1].(int)]
	}).(VirtualHubRouteOutput)
}

type VirtualHubRouteResponse struct {
	AddressPrefixes  []string `pulumi:"addressPrefixes"`
	NextHopIpAddress *string  `pulumi:"nextHopIpAddress"`
}

type VirtualHubRouteResponseOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteResponse)(nil)).Elem()
}

func (o VirtualHubRouteResponseOutput) ToVirtualHubRouteResponseOutput() VirtualHubRouteResponseOutput {
	return o
}

func (o VirtualHubRouteResponseOutput) ToVirtualHubRouteResponseOutputWithContext(ctx context.Context) VirtualHubRouteResponseOutput {
	return o
}

func (o VirtualHubRouteResponseOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualHubRouteResponse) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o VirtualHubRouteResponseOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubRouteResponse) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

type VirtualHubRouteResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHubRouteResponse)(nil)).Elem()
}

func (o VirtualHubRouteResponseArrayOutput) ToVirtualHubRouteResponseArrayOutput() VirtualHubRouteResponseArrayOutput {
	return o
}

func (o VirtualHubRouteResponseArrayOutput) ToVirtualHubRouteResponseArrayOutputWithContext(ctx context.Context) VirtualHubRouteResponseArrayOutput {
	return o
}

func (o VirtualHubRouteResponseArrayOutput) Index(i pulumi.IntInput) VirtualHubRouteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualHubRouteResponse {
		return vs[0].([]VirtualHubRouteResponse)[vs[1].(int)]
	}).(VirtualHubRouteResponseOutput)
}

type VirtualHubRouteTable struct {
	Routes []VirtualHubRoute `pulumi:"routes"`
}





type VirtualHubRouteTableInput interface {
	pulumi.Input

	ToVirtualHubRouteTableOutput() VirtualHubRouteTableOutput
	ToVirtualHubRouteTableOutputWithContext(context.Context) VirtualHubRouteTableOutput
}

type VirtualHubRouteTableArgs struct {
	Routes VirtualHubRouteArrayInput `pulumi:"routes"`
}

func (VirtualHubRouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteTable)(nil)).Elem()
}

func (i VirtualHubRouteTableArgs) ToVirtualHubRouteTableOutput() VirtualHubRouteTableOutput {
	return i.ToVirtualHubRouteTableOutputWithContext(context.Background())
}

func (i VirtualHubRouteTableArgs) ToVirtualHubRouteTableOutputWithContext(ctx context.Context) VirtualHubRouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteTableOutput)
}

func (i VirtualHubRouteTableArgs) ToVirtualHubRouteTablePtrOutput() VirtualHubRouteTablePtrOutput {
	return i.ToVirtualHubRouteTablePtrOutputWithContext(context.Background())
}

func (i VirtualHubRouteTableArgs) ToVirtualHubRouteTablePtrOutputWithContext(ctx context.Context) VirtualHubRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteTableOutput).ToVirtualHubRouteTablePtrOutputWithContext(ctx)
}









type VirtualHubRouteTablePtrInput interface {
	pulumi.Input

	ToVirtualHubRouteTablePtrOutput() VirtualHubRouteTablePtrOutput
	ToVirtualHubRouteTablePtrOutputWithContext(context.Context) VirtualHubRouteTablePtrOutput
}

type virtualHubRouteTablePtrType VirtualHubRouteTableArgs

func VirtualHubRouteTablePtr(v *VirtualHubRouteTableArgs) VirtualHubRouteTablePtrInput {
	return (*virtualHubRouteTablePtrType)(v)
}

func (*virtualHubRouteTablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubRouteTable)(nil)).Elem()
}

func (i *virtualHubRouteTablePtrType) ToVirtualHubRouteTablePtrOutput() VirtualHubRouteTablePtrOutput {
	return i.ToVirtualHubRouteTablePtrOutputWithContext(context.Background())
}

func (i *virtualHubRouteTablePtrType) ToVirtualHubRouteTablePtrOutputWithContext(ctx context.Context) VirtualHubRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteTablePtrOutput)
}

type VirtualHubRouteTableOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteTable)(nil)).Elem()
}

func (o VirtualHubRouteTableOutput) ToVirtualHubRouteTableOutput() VirtualHubRouteTableOutput {
	return o
}

func (o VirtualHubRouteTableOutput) ToVirtualHubRouteTableOutputWithContext(ctx context.Context) VirtualHubRouteTableOutput {
	return o
}

func (o VirtualHubRouteTableOutput) ToVirtualHubRouteTablePtrOutput() VirtualHubRouteTablePtrOutput {
	return o.ToVirtualHubRouteTablePtrOutputWithContext(context.Background())
}

func (o VirtualHubRouteTableOutput) ToVirtualHubRouteTablePtrOutputWithContext(ctx context.Context) VirtualHubRouteTablePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualHubRouteTable) *VirtualHubRouteTable {
		return &v
	}).(VirtualHubRouteTablePtrOutput)
}

func (o VirtualHubRouteTableOutput) Routes() VirtualHubRouteArrayOutput {
	return o.ApplyT(func(v VirtualHubRouteTable) []VirtualHubRoute { return v.Routes }).(VirtualHubRouteArrayOutput)
}

type VirtualHubRouteTablePtrOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubRouteTable)(nil)).Elem()
}

func (o VirtualHubRouteTablePtrOutput) ToVirtualHubRouteTablePtrOutput() VirtualHubRouteTablePtrOutput {
	return o
}

func (o VirtualHubRouteTablePtrOutput) ToVirtualHubRouteTablePtrOutputWithContext(ctx context.Context) VirtualHubRouteTablePtrOutput {
	return o
}

func (o VirtualHubRouteTablePtrOutput) Elem() VirtualHubRouteTableOutput {
	return o.ApplyT(func(v *VirtualHubRouteTable) VirtualHubRouteTable {
		if v != nil {
			return *v
		}
		var ret VirtualHubRouteTable
		return ret
	}).(VirtualHubRouteTableOutput)
}

func (o VirtualHubRouteTablePtrOutput) Routes() VirtualHubRouteArrayOutput {
	return o.ApplyT(func(v *VirtualHubRouteTable) []VirtualHubRoute {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(VirtualHubRouteArrayOutput)
}

type VirtualHubRouteTableResponse struct {
	Routes []VirtualHubRouteResponse `pulumi:"routes"`
}

type VirtualHubRouteTableResponseOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteTableResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteTableResponse)(nil)).Elem()
}

func (o VirtualHubRouteTableResponseOutput) ToVirtualHubRouteTableResponseOutput() VirtualHubRouteTableResponseOutput {
	return o
}

func (o VirtualHubRouteTableResponseOutput) ToVirtualHubRouteTableResponseOutputWithContext(ctx context.Context) VirtualHubRouteTableResponseOutput {
	return o
}

func (o VirtualHubRouteTableResponseOutput) Routes() VirtualHubRouteResponseArrayOutput {
	return o.ApplyT(func(v VirtualHubRouteTableResponse) []VirtualHubRouteResponse { return v.Routes }).(VirtualHubRouteResponseArrayOutput)
}

type VirtualHubRouteTableResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteTableResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubRouteTableResponse)(nil)).Elem()
}

func (o VirtualHubRouteTableResponsePtrOutput) ToVirtualHubRouteTableResponsePtrOutput() VirtualHubRouteTableResponsePtrOutput {
	return o
}

func (o VirtualHubRouteTableResponsePtrOutput) ToVirtualHubRouteTableResponsePtrOutputWithContext(ctx context.Context) VirtualHubRouteTableResponsePtrOutput {
	return o
}

func (o VirtualHubRouteTableResponsePtrOutput) Elem() VirtualHubRouteTableResponseOutput {
	return o.ApplyT(func(v *VirtualHubRouteTableResponse) VirtualHubRouteTableResponse {
		if v != nil {
			return *v
		}
		var ret VirtualHubRouteTableResponse
		return ret
	}).(VirtualHubRouteTableResponseOutput)
}

func (o VirtualHubRouteTableResponsePtrOutput) Routes() VirtualHubRouteResponseArrayOutput {
	return o.ApplyT(func(v *VirtualHubRouteTableResponse) []VirtualHubRouteResponse {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(VirtualHubRouteResponseArrayOutput)
}

type VirtualHubRouteTableV2Type struct {
	AttachedConnections []string            `pulumi:"attachedConnections"`
	Id                  *string             `pulumi:"id"`
	Name                *string             `pulumi:"name"`
	Routes              []VirtualHubRouteV2 `pulumi:"routes"`
}





type VirtualHubRouteTableV2TypeInput interface {
	pulumi.Input

	ToVirtualHubRouteTableV2TypeOutput() VirtualHubRouteTableV2TypeOutput
	ToVirtualHubRouteTableV2TypeOutputWithContext(context.Context) VirtualHubRouteTableV2TypeOutput
}

type VirtualHubRouteTableV2TypeArgs struct {
	AttachedConnections pulumi.StringArrayInput     `pulumi:"attachedConnections"`
	Id                  pulumi.StringPtrInput       `pulumi:"id"`
	Name                pulumi.StringPtrInput       `pulumi:"name"`
	Routes              VirtualHubRouteV2ArrayInput `pulumi:"routes"`
}

func (VirtualHubRouteTableV2TypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteTableV2Type)(nil)).Elem()
}

func (i VirtualHubRouteTableV2TypeArgs) ToVirtualHubRouteTableV2TypeOutput() VirtualHubRouteTableV2TypeOutput {
	return i.ToVirtualHubRouteTableV2TypeOutputWithContext(context.Background())
}

func (i VirtualHubRouteTableV2TypeArgs) ToVirtualHubRouteTableV2TypeOutputWithContext(ctx context.Context) VirtualHubRouteTableV2TypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteTableV2TypeOutput)
}





type VirtualHubRouteTableV2TypeArrayInput interface {
	pulumi.Input

	ToVirtualHubRouteTableV2TypeArrayOutput() VirtualHubRouteTableV2TypeArrayOutput
	ToVirtualHubRouteTableV2TypeArrayOutputWithContext(context.Context) VirtualHubRouteTableV2TypeArrayOutput
}

type VirtualHubRouteTableV2TypeArray []VirtualHubRouteTableV2TypeInput

func (VirtualHubRouteTableV2TypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHubRouteTableV2Type)(nil)).Elem()
}

func (i VirtualHubRouteTableV2TypeArray) ToVirtualHubRouteTableV2TypeArrayOutput() VirtualHubRouteTableV2TypeArrayOutput {
	return i.ToVirtualHubRouteTableV2TypeArrayOutputWithContext(context.Background())
}

func (i VirtualHubRouteTableV2TypeArray) ToVirtualHubRouteTableV2TypeArrayOutputWithContext(ctx context.Context) VirtualHubRouteTableV2TypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteTableV2TypeArrayOutput)
}

type VirtualHubRouteTableV2TypeOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteTableV2TypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteTableV2Type)(nil)).Elem()
}

func (o VirtualHubRouteTableV2TypeOutput) ToVirtualHubRouteTableV2TypeOutput() VirtualHubRouteTableV2TypeOutput {
	return o
}

func (o VirtualHubRouteTableV2TypeOutput) ToVirtualHubRouteTableV2TypeOutputWithContext(ctx context.Context) VirtualHubRouteTableV2TypeOutput {
	return o
}

func (o VirtualHubRouteTableV2TypeOutput) AttachedConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualHubRouteTableV2Type) []string { return v.AttachedConnections }).(pulumi.StringArrayOutput)
}

func (o VirtualHubRouteTableV2TypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubRouteTableV2Type) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualHubRouteTableV2TypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubRouteTableV2Type) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualHubRouteTableV2TypeOutput) Routes() VirtualHubRouteV2ArrayOutput {
	return o.ApplyT(func(v VirtualHubRouteTableV2Type) []VirtualHubRouteV2 { return v.Routes }).(VirtualHubRouteV2ArrayOutput)
}

type VirtualHubRouteTableV2TypeArrayOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteTableV2TypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHubRouteTableV2Type)(nil)).Elem()
}

func (o VirtualHubRouteTableV2TypeArrayOutput) ToVirtualHubRouteTableV2TypeArrayOutput() VirtualHubRouteTableV2TypeArrayOutput {
	return o
}

func (o VirtualHubRouteTableV2TypeArrayOutput) ToVirtualHubRouteTableV2TypeArrayOutputWithContext(ctx context.Context) VirtualHubRouteTableV2TypeArrayOutput {
	return o
}

func (o VirtualHubRouteTableV2TypeArrayOutput) Index(i pulumi.IntInput) VirtualHubRouteTableV2TypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualHubRouteTableV2Type {
		return vs[0].([]VirtualHubRouteTableV2Type)[vs[1].(int)]
	}).(VirtualHubRouteTableV2TypeOutput)
}

type VirtualHubRouteTableV2Response struct {
	AttachedConnections []string                    `pulumi:"attachedConnections"`
	Etag                string                      `pulumi:"etag"`
	Id                  *string                     `pulumi:"id"`
	Name                *string                     `pulumi:"name"`
	ProvisioningState   string                      `pulumi:"provisioningState"`
	Routes              []VirtualHubRouteV2Response `pulumi:"routes"`
}

type VirtualHubRouteTableV2ResponseOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteTableV2ResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteTableV2Response)(nil)).Elem()
}

func (o VirtualHubRouteTableV2ResponseOutput) ToVirtualHubRouteTableV2ResponseOutput() VirtualHubRouteTableV2ResponseOutput {
	return o
}

func (o VirtualHubRouteTableV2ResponseOutput) ToVirtualHubRouteTableV2ResponseOutputWithContext(ctx context.Context) VirtualHubRouteTableV2ResponseOutput {
	return o
}

func (o VirtualHubRouteTableV2ResponseOutput) AttachedConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualHubRouteTableV2Response) []string { return v.AttachedConnections }).(pulumi.StringArrayOutput)
}

func (o VirtualHubRouteTableV2ResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualHubRouteTableV2Response) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualHubRouteTableV2ResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubRouteTableV2Response) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualHubRouteTableV2ResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubRouteTableV2Response) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualHubRouteTableV2ResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualHubRouteTableV2Response) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualHubRouteTableV2ResponseOutput) Routes() VirtualHubRouteV2ResponseArrayOutput {
	return o.ApplyT(func(v VirtualHubRouteTableV2Response) []VirtualHubRouteV2Response { return v.Routes }).(VirtualHubRouteV2ResponseArrayOutput)
}

type VirtualHubRouteTableV2ResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteTableV2ResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHubRouteTableV2Response)(nil)).Elem()
}

func (o VirtualHubRouteTableV2ResponseArrayOutput) ToVirtualHubRouteTableV2ResponseArrayOutput() VirtualHubRouteTableV2ResponseArrayOutput {
	return o
}

func (o VirtualHubRouteTableV2ResponseArrayOutput) ToVirtualHubRouteTableV2ResponseArrayOutputWithContext(ctx context.Context) VirtualHubRouteTableV2ResponseArrayOutput {
	return o
}

func (o VirtualHubRouteTableV2ResponseArrayOutput) Index(i pulumi.IntInput) VirtualHubRouteTableV2ResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualHubRouteTableV2Response {
		return vs[0].([]VirtualHubRouteTableV2Response)[vs[1].(int)]
	}).(VirtualHubRouteTableV2ResponseOutput)
}

type VirtualHubRouteV2 struct {
	DestinationType *string  `pulumi:"destinationType"`
	Destinations    []string `pulumi:"destinations"`
	NextHopType     *string  `pulumi:"nextHopType"`
	NextHops        []string `pulumi:"nextHops"`
}





type VirtualHubRouteV2Input interface {
	pulumi.Input

	ToVirtualHubRouteV2Output() VirtualHubRouteV2Output
	ToVirtualHubRouteV2OutputWithContext(context.Context) VirtualHubRouteV2Output
}

type VirtualHubRouteV2Args struct {
	DestinationType pulumi.StringPtrInput   `pulumi:"destinationType"`
	Destinations    pulumi.StringArrayInput `pulumi:"destinations"`
	NextHopType     pulumi.StringPtrInput   `pulumi:"nextHopType"`
	NextHops        pulumi.StringArrayInput `pulumi:"nextHops"`
}

func (VirtualHubRouteV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteV2)(nil)).Elem()
}

func (i VirtualHubRouteV2Args) ToVirtualHubRouteV2Output() VirtualHubRouteV2Output {
	return i.ToVirtualHubRouteV2OutputWithContext(context.Background())
}

func (i VirtualHubRouteV2Args) ToVirtualHubRouteV2OutputWithContext(ctx context.Context) VirtualHubRouteV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteV2Output)
}





type VirtualHubRouteV2ArrayInput interface {
	pulumi.Input

	ToVirtualHubRouteV2ArrayOutput() VirtualHubRouteV2ArrayOutput
	ToVirtualHubRouteV2ArrayOutputWithContext(context.Context) VirtualHubRouteV2ArrayOutput
}

type VirtualHubRouteV2Array []VirtualHubRouteV2Input

func (VirtualHubRouteV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHubRouteV2)(nil)).Elem()
}

func (i VirtualHubRouteV2Array) ToVirtualHubRouteV2ArrayOutput() VirtualHubRouteV2ArrayOutput {
	return i.ToVirtualHubRouteV2ArrayOutputWithContext(context.Background())
}

func (i VirtualHubRouteV2Array) ToVirtualHubRouteV2ArrayOutputWithContext(ctx context.Context) VirtualHubRouteV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubRouteV2ArrayOutput)
}

type VirtualHubRouteV2Output struct{ *pulumi.OutputState }

func (VirtualHubRouteV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteV2)(nil)).Elem()
}

func (o VirtualHubRouteV2Output) ToVirtualHubRouteV2Output() VirtualHubRouteV2Output {
	return o
}

func (o VirtualHubRouteV2Output) ToVirtualHubRouteV2OutputWithContext(ctx context.Context) VirtualHubRouteV2Output {
	return o
}

func (o VirtualHubRouteV2Output) DestinationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubRouteV2) *string { return v.DestinationType }).(pulumi.StringPtrOutput)
}

func (o VirtualHubRouteV2Output) Destinations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualHubRouteV2) []string { return v.Destinations }).(pulumi.StringArrayOutput)
}

func (o VirtualHubRouteV2Output) NextHopType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubRouteV2) *string { return v.NextHopType }).(pulumi.StringPtrOutput)
}

func (o VirtualHubRouteV2Output) NextHops() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualHubRouteV2) []string { return v.NextHops }).(pulumi.StringArrayOutput)
}

type VirtualHubRouteV2ArrayOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHubRouteV2)(nil)).Elem()
}

func (o VirtualHubRouteV2ArrayOutput) ToVirtualHubRouteV2ArrayOutput() VirtualHubRouteV2ArrayOutput {
	return o
}

func (o VirtualHubRouteV2ArrayOutput) ToVirtualHubRouteV2ArrayOutputWithContext(ctx context.Context) VirtualHubRouteV2ArrayOutput {
	return o
}

func (o VirtualHubRouteV2ArrayOutput) Index(i pulumi.IntInput) VirtualHubRouteV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualHubRouteV2 {
		return vs[0].([]VirtualHubRouteV2)[vs[1].(int)]
	}).(VirtualHubRouteV2Output)
}

type VirtualHubRouteV2Response struct {
	DestinationType *string  `pulumi:"destinationType"`
	Destinations    []string `pulumi:"destinations"`
	NextHopType     *string  `pulumi:"nextHopType"`
	NextHops        []string `pulumi:"nextHops"`
}

type VirtualHubRouteV2ResponseOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteV2ResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubRouteV2Response)(nil)).Elem()
}

func (o VirtualHubRouteV2ResponseOutput) ToVirtualHubRouteV2ResponseOutput() VirtualHubRouteV2ResponseOutput {
	return o
}

func (o VirtualHubRouteV2ResponseOutput) ToVirtualHubRouteV2ResponseOutputWithContext(ctx context.Context) VirtualHubRouteV2ResponseOutput {
	return o
}

func (o VirtualHubRouteV2ResponseOutput) DestinationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubRouteV2Response) *string { return v.DestinationType }).(pulumi.StringPtrOutput)
}

func (o VirtualHubRouteV2ResponseOutput) Destinations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualHubRouteV2Response) []string { return v.Destinations }).(pulumi.StringArrayOutput)
}

func (o VirtualHubRouteV2ResponseOutput) NextHopType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHubRouteV2Response) *string { return v.NextHopType }).(pulumi.StringPtrOutput)
}

func (o VirtualHubRouteV2ResponseOutput) NextHops() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualHubRouteV2Response) []string { return v.NextHops }).(pulumi.StringArrayOutput)
}

type VirtualHubRouteV2ResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualHubRouteV2ResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualHubRouteV2Response)(nil)).Elem()
}

func (o VirtualHubRouteV2ResponseArrayOutput) ToVirtualHubRouteV2ResponseArrayOutput() VirtualHubRouteV2ResponseArrayOutput {
	return o
}

func (o VirtualHubRouteV2ResponseArrayOutput) ToVirtualHubRouteV2ResponseArrayOutputWithContext(ctx context.Context) VirtualHubRouteV2ResponseArrayOutput {
	return o
}

func (o VirtualHubRouteV2ResponseArrayOutput) Index(i pulumi.IntInput) VirtualHubRouteV2ResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualHubRouteV2Response {
		return vs[0].([]VirtualHubRouteV2Response)[vs[1].(int)]
	}).(VirtualHubRouteV2ResponseOutput)
}

type VirtualNetworkBgpCommunities struct {
	VirtualNetworkCommunity string `pulumi:"virtualNetworkCommunity"`
}





type VirtualNetworkBgpCommunitiesInput interface {
	pulumi.Input

	ToVirtualNetworkBgpCommunitiesOutput() VirtualNetworkBgpCommunitiesOutput
	ToVirtualNetworkBgpCommunitiesOutputWithContext(context.Context) VirtualNetworkBgpCommunitiesOutput
}

type VirtualNetworkBgpCommunitiesArgs struct {
	VirtualNetworkCommunity pulumi.StringInput `pulumi:"virtualNetworkCommunity"`
}

func (VirtualNetworkBgpCommunitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkBgpCommunities)(nil)).Elem()
}

func (i VirtualNetworkBgpCommunitiesArgs) ToVirtualNetworkBgpCommunitiesOutput() VirtualNetworkBgpCommunitiesOutput {
	return i.ToVirtualNetworkBgpCommunitiesOutputWithContext(context.Background())
}

func (i VirtualNetworkBgpCommunitiesArgs) ToVirtualNetworkBgpCommunitiesOutputWithContext(ctx context.Context) VirtualNetworkBgpCommunitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkBgpCommunitiesOutput)
}

func (i VirtualNetworkBgpCommunitiesArgs) ToVirtualNetworkBgpCommunitiesPtrOutput() VirtualNetworkBgpCommunitiesPtrOutput {
	return i.ToVirtualNetworkBgpCommunitiesPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkBgpCommunitiesArgs) ToVirtualNetworkBgpCommunitiesPtrOutputWithContext(ctx context.Context) VirtualNetworkBgpCommunitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkBgpCommunitiesOutput).ToVirtualNetworkBgpCommunitiesPtrOutputWithContext(ctx)
}









type VirtualNetworkBgpCommunitiesPtrInput interface {
	pulumi.Input

	ToVirtualNetworkBgpCommunitiesPtrOutput() VirtualNetworkBgpCommunitiesPtrOutput
	ToVirtualNetworkBgpCommunitiesPtrOutputWithContext(context.Context) VirtualNetworkBgpCommunitiesPtrOutput
}

type virtualNetworkBgpCommunitiesPtrType VirtualNetworkBgpCommunitiesArgs

func VirtualNetworkBgpCommunitiesPtr(v *VirtualNetworkBgpCommunitiesArgs) VirtualNetworkBgpCommunitiesPtrInput {
	return (*virtualNetworkBgpCommunitiesPtrType)(v)
}

func (*virtualNetworkBgpCommunitiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkBgpCommunities)(nil)).Elem()
}

func (i *virtualNetworkBgpCommunitiesPtrType) ToVirtualNetworkBgpCommunitiesPtrOutput() VirtualNetworkBgpCommunitiesPtrOutput {
	return i.ToVirtualNetworkBgpCommunitiesPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkBgpCommunitiesPtrType) ToVirtualNetworkBgpCommunitiesPtrOutputWithContext(ctx context.Context) VirtualNetworkBgpCommunitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkBgpCommunitiesPtrOutput)
}

type VirtualNetworkBgpCommunitiesOutput struct{ *pulumi.OutputState }

func (VirtualNetworkBgpCommunitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkBgpCommunities)(nil)).Elem()
}

func (o VirtualNetworkBgpCommunitiesOutput) ToVirtualNetworkBgpCommunitiesOutput() VirtualNetworkBgpCommunitiesOutput {
	return o
}

func (o VirtualNetworkBgpCommunitiesOutput) ToVirtualNetworkBgpCommunitiesOutputWithContext(ctx context.Context) VirtualNetworkBgpCommunitiesOutput {
	return o
}

func (o VirtualNetworkBgpCommunitiesOutput) ToVirtualNetworkBgpCommunitiesPtrOutput() VirtualNetworkBgpCommunitiesPtrOutput {
	return o.ToVirtualNetworkBgpCommunitiesPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkBgpCommunitiesOutput) ToVirtualNetworkBgpCommunitiesPtrOutputWithContext(ctx context.Context) VirtualNetworkBgpCommunitiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkBgpCommunities) *VirtualNetworkBgpCommunities {
		return &v
	}).(VirtualNetworkBgpCommunitiesPtrOutput)
}

func (o VirtualNetworkBgpCommunitiesOutput) VirtualNetworkCommunity() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkBgpCommunities) string { return v.VirtualNetworkCommunity }).(pulumi.StringOutput)
}

type VirtualNetworkBgpCommunitiesPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkBgpCommunitiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkBgpCommunities)(nil)).Elem()
}

func (o VirtualNetworkBgpCommunitiesPtrOutput) ToVirtualNetworkBgpCommunitiesPtrOutput() VirtualNetworkBgpCommunitiesPtrOutput {
	return o
}

func (o VirtualNetworkBgpCommunitiesPtrOutput) ToVirtualNetworkBgpCommunitiesPtrOutputWithContext(ctx context.Context) VirtualNetworkBgpCommunitiesPtrOutput {
	return o
}

func (o VirtualNetworkBgpCommunitiesPtrOutput) Elem() VirtualNetworkBgpCommunitiesOutput {
	return o.ApplyT(func(v *VirtualNetworkBgpCommunities) VirtualNetworkBgpCommunities {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkBgpCommunities
		return ret
	}).(VirtualNetworkBgpCommunitiesOutput)
}

func (o VirtualNetworkBgpCommunitiesPtrOutput) VirtualNetworkCommunity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkBgpCommunities) *string {
		if v == nil {
			return nil
		}
		return &v.VirtualNetworkCommunity
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkBgpCommunitiesResponse struct {
	RegionalCommunity       string `pulumi:"regionalCommunity"`
	VirtualNetworkCommunity string `pulumi:"virtualNetworkCommunity"`
}

type VirtualNetworkBgpCommunitiesResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkBgpCommunitiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkBgpCommunitiesResponse)(nil)).Elem()
}

func (o VirtualNetworkBgpCommunitiesResponseOutput) ToVirtualNetworkBgpCommunitiesResponseOutput() VirtualNetworkBgpCommunitiesResponseOutput {
	return o
}

func (o VirtualNetworkBgpCommunitiesResponseOutput) ToVirtualNetworkBgpCommunitiesResponseOutputWithContext(ctx context.Context) VirtualNetworkBgpCommunitiesResponseOutput {
	return o
}

func (o VirtualNetworkBgpCommunitiesResponseOutput) RegionalCommunity() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkBgpCommunitiesResponse) string { return v.RegionalCommunity }).(pulumi.StringOutput)
}

func (o VirtualNetworkBgpCommunitiesResponseOutput) VirtualNetworkCommunity() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkBgpCommunitiesResponse) string { return v.VirtualNetworkCommunity }).(pulumi.StringOutput)
}

type VirtualNetworkBgpCommunitiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkBgpCommunitiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkBgpCommunitiesResponse)(nil)).Elem()
}

func (o VirtualNetworkBgpCommunitiesResponsePtrOutput) ToVirtualNetworkBgpCommunitiesResponsePtrOutput() VirtualNetworkBgpCommunitiesResponsePtrOutput {
	return o
}

func (o VirtualNetworkBgpCommunitiesResponsePtrOutput) ToVirtualNetworkBgpCommunitiesResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkBgpCommunitiesResponsePtrOutput {
	return o
}

func (o VirtualNetworkBgpCommunitiesResponsePtrOutput) Elem() VirtualNetworkBgpCommunitiesResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkBgpCommunitiesResponse) VirtualNetworkBgpCommunitiesResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkBgpCommunitiesResponse
		return ret
	}).(VirtualNetworkBgpCommunitiesResponseOutput)
}

func (o VirtualNetworkBgpCommunitiesResponsePtrOutput) RegionalCommunity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkBgpCommunitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RegionalCommunity
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkBgpCommunitiesResponsePtrOutput) VirtualNetworkCommunity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkBgpCommunitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VirtualNetworkCommunity
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkDnsForwardingRulesetResponse struct {
	Id                 *string              `pulumi:"id"`
	VirtualNetworkLink *SubResourceResponse `pulumi:"virtualNetworkLink"`
}

type VirtualNetworkDnsForwardingRulesetResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkDnsForwardingRulesetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkDnsForwardingRulesetResponse)(nil)).Elem()
}

func (o VirtualNetworkDnsForwardingRulesetResponseOutput) ToVirtualNetworkDnsForwardingRulesetResponseOutput() VirtualNetworkDnsForwardingRulesetResponseOutput {
	return o
}

func (o VirtualNetworkDnsForwardingRulesetResponseOutput) ToVirtualNetworkDnsForwardingRulesetResponseOutputWithContext(ctx context.Context) VirtualNetworkDnsForwardingRulesetResponseOutput {
	return o
}

func (o VirtualNetworkDnsForwardingRulesetResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkDnsForwardingRulesetResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkDnsForwardingRulesetResponseOutput) VirtualNetworkLink() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkDnsForwardingRulesetResponse) *SubResourceResponse { return v.VirtualNetworkLink }).(SubResourceResponsePtrOutput)
}

type VirtualNetworkDnsForwardingRulesetResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkDnsForwardingRulesetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkDnsForwardingRulesetResponse)(nil)).Elem()
}

func (o VirtualNetworkDnsForwardingRulesetResponseArrayOutput) ToVirtualNetworkDnsForwardingRulesetResponseArrayOutput() VirtualNetworkDnsForwardingRulesetResponseArrayOutput {
	return o
}

func (o VirtualNetworkDnsForwardingRulesetResponseArrayOutput) ToVirtualNetworkDnsForwardingRulesetResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkDnsForwardingRulesetResponseArrayOutput {
	return o
}

func (o VirtualNetworkDnsForwardingRulesetResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkDnsForwardingRulesetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkDnsForwardingRulesetResponse {
		return vs[0].([]VirtualNetworkDnsForwardingRulesetResponse)[vs[1].(int)]
	}).(VirtualNetworkDnsForwardingRulesetResponseOutput)
}

type VirtualNetworkGatewayType struct {
	ActiveActive                   *bool                                  `pulumi:"activeActive"`
	BgpSettings                    *BgpSettings                           `pulumi:"bgpSettings"`
	CustomRoutes                   *AddressSpace                          `pulumi:"customRoutes"`
	EnableBgp                      *bool                                  `pulumi:"enableBgp"`
	EnableDnsForwarding            *bool                                  `pulumi:"enableDnsForwarding"`
	EnablePrivateIpAddress         *bool                                  `pulumi:"enablePrivateIpAddress"`
	ExtendedLocation               *ExtendedLocation                      `pulumi:"extendedLocation"`
	GatewayDefaultSite             *SubResource                           `pulumi:"gatewayDefaultSite"`
	GatewayType                    *string                                `pulumi:"gatewayType"`
	Id                             *string                                `pulumi:"id"`
	IpConfigurations               []VirtualNetworkGatewayIPConfiguration `pulumi:"ipConfigurations"`
	Location                       *string                                `pulumi:"location"`
	Sku                            *VirtualNetworkGatewaySku              `pulumi:"sku"`
	Tags                           map[string]string                      `pulumi:"tags"`
	VNetExtendedLocationResourceId *string                                `pulumi:"vNetExtendedLocationResourceId"`
	VpnClientConfiguration         *VpnClientConfiguration                `pulumi:"vpnClientConfiguration"`
	VpnGatewayGeneration           *string                                `pulumi:"vpnGatewayGeneration"`
	VpnType                        *string                                `pulumi:"vpnType"`
}





type VirtualNetworkGatewayTypeInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayTypeOutput() VirtualNetworkGatewayTypeOutput
	ToVirtualNetworkGatewayTypeOutputWithContext(context.Context) VirtualNetworkGatewayTypeOutput
}

type VirtualNetworkGatewayTypeArgs struct {
	ActiveActive                   pulumi.BoolPtrInput                            `pulumi:"activeActive"`
	BgpSettings                    BgpSettingsPtrInput                            `pulumi:"bgpSettings"`
	CustomRoutes                   AddressSpacePtrInput                           `pulumi:"customRoutes"`
	EnableBgp                      pulumi.BoolPtrInput                            `pulumi:"enableBgp"`
	EnableDnsForwarding            pulumi.BoolPtrInput                            `pulumi:"enableDnsForwarding"`
	EnablePrivateIpAddress         pulumi.BoolPtrInput                            `pulumi:"enablePrivateIpAddress"`
	ExtendedLocation               ExtendedLocationPtrInput                       `pulumi:"extendedLocation"`
	GatewayDefaultSite             SubResourcePtrInput                            `pulumi:"gatewayDefaultSite"`
	GatewayType                    pulumi.StringPtrInput                          `pulumi:"gatewayType"`
	Id                             pulumi.StringPtrInput                          `pulumi:"id"`
	IpConfigurations               VirtualNetworkGatewayIPConfigurationArrayInput `pulumi:"ipConfigurations"`
	Location                       pulumi.StringPtrInput                          `pulumi:"location"`
	Sku                            VirtualNetworkGatewaySkuPtrInput               `pulumi:"sku"`
	Tags                           pulumi.StringMapInput                          `pulumi:"tags"`
	VNetExtendedLocationResourceId pulumi.StringPtrInput                          `pulumi:"vNetExtendedLocationResourceId"`
	VpnClientConfiguration         VpnClientConfigurationPtrInput                 `pulumi:"vpnClientConfiguration"`
	VpnGatewayGeneration           pulumi.StringPtrInput                          `pulumi:"vpnGatewayGeneration"`
	VpnType                        pulumi.StringPtrInput                          `pulumi:"vpnType"`
}

func (VirtualNetworkGatewayTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayType)(nil)).Elem()
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypeOutput() VirtualNetworkGatewayTypeOutput {
	return i.ToVirtualNetworkGatewayTypeOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypeOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayTypeOutput)
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return i.ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayTypeOutput).ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx)
}









type VirtualNetworkGatewayTypePtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput
	ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Context) VirtualNetworkGatewayTypePtrOutput
}

type virtualNetworkGatewayTypePtrType VirtualNetworkGatewayTypeArgs

func VirtualNetworkGatewayTypePtr(v *VirtualNetworkGatewayTypeArgs) VirtualNetworkGatewayTypePtrInput {
	return (*virtualNetworkGatewayTypePtrType)(v)
}

func (*virtualNetworkGatewayTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayType)(nil)).Elem()
}

func (i *virtualNetworkGatewayTypePtrType) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return i.ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkGatewayTypePtrType) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayTypePtrOutput)
}

type VirtualNetworkGatewayTypeOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayType)(nil)).Elem()
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypeOutput() VirtualNetworkGatewayTypeOutput {
	return o
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypeOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeOutput {
	return o
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return o.ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewayType) *VirtualNetworkGatewayType {
		return &v
	}).(VirtualNetworkGatewayTypePtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *bool { return v.ActiveActive }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) BgpSettings() BgpSettingsPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *BgpSettings { return v.BgpSettings }).(BgpSettingsPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) CustomRoutes() AddressSpacePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *AddressSpace { return v.CustomRoutes }).(AddressSpacePtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) EnableDnsForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *bool { return v.EnableDnsForwarding }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) EnablePrivateIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *bool { return v.EnablePrivateIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) ExtendedLocation() ExtendedLocationPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *ExtendedLocation { return v.ExtendedLocation }).(ExtendedLocationPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) GatewayDefaultSite() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *SubResource { return v.GatewayDefaultSite }).(SubResourcePtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.GatewayType }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) []VirtualNetworkGatewayIPConfiguration { return v.IpConfigurations }).(VirtualNetworkGatewayIPConfigurationArrayOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Sku() VirtualNetworkGatewaySkuPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *VirtualNetworkGatewaySku { return v.Sku }).(VirtualNetworkGatewaySkuPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayTypeOutput) VNetExtendedLocationResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.VNetExtendedLocationResourceId }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) VpnClientConfiguration() VpnClientConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *VpnClientConfiguration { return v.VpnClientConfiguration }).(VpnClientConfigurationPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) VpnGatewayGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.VpnGatewayGeneration }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.VpnType }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayTypePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayType)(nil)).Elem()
}

func (o VirtualNetworkGatewayTypePtrOutput) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return o
}

func (o VirtualNetworkGatewayTypePtrOutput) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return o
}

func (o VirtualNetworkGatewayTypePtrOutput) Elem() VirtualNetworkGatewayTypeOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) VirtualNetworkGatewayType {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewayType
		return ret
	}).(VirtualNetworkGatewayTypeOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *bool {
		if v == nil {
			return nil
		}
		return v.ActiveActive
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) BgpSettings() BgpSettingsPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *BgpSettings {
		if v == nil {
			return nil
		}
		return v.BgpSettings
	}).(BgpSettingsPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) CustomRoutes() AddressSpacePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *AddressSpace {
		if v == nil {
			return nil
		}
		return v.CustomRoutes
	}).(AddressSpacePtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *bool {
		if v == nil {
			return nil
		}
		return v.EnableBgp
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) EnableDnsForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDnsForwarding
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) EnablePrivateIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePrivateIpAddress
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) ExtendedLocation() ExtendedLocationPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *ExtendedLocation {
		if v == nil {
			return nil
		}
		return v.ExtendedLocation
	}).(ExtendedLocationPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) GatewayDefaultSite() SubResourcePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *SubResource {
		if v == nil {
			return nil
		}
		return v.GatewayDefaultSite
	}).(SubResourcePtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.GatewayType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) []VirtualNetworkGatewayIPConfiguration {
		if v == nil {
			return nil
		}
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationArrayOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Sku() VirtualNetworkGatewaySkuPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *VirtualNetworkGatewaySku {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(VirtualNetworkGatewaySkuPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) VNetExtendedLocationResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.VNetExtendedLocationResourceId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) VpnClientConfiguration() VpnClientConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *VpnClientConfiguration {
		if v == nil {
			return nil
		}
		return v.VpnClientConfiguration
	}).(VpnClientConfigurationPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) VpnGatewayGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.VpnGatewayGeneration
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.VpnType
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayIPConfiguration struct {
	Id                        *string      `pulumi:"id"`
	Name                      *string      `pulumi:"name"`
	PrivateIPAllocationMethod *string      `pulumi:"privateIPAllocationMethod"`
	PublicIPAddress           *SubResource `pulumi:"publicIPAddress"`
	Subnet                    *SubResource `pulumi:"subnet"`
}





type VirtualNetworkGatewayIPConfigurationInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayIPConfigurationOutput() VirtualNetworkGatewayIPConfigurationOutput
	ToVirtualNetworkGatewayIPConfigurationOutputWithContext(context.Context) VirtualNetworkGatewayIPConfigurationOutput
}

type VirtualNetworkGatewayIPConfigurationArgs struct {
	Id                        pulumi.StringPtrInput `pulumi:"id"`
	Name                      pulumi.StringPtrInput `pulumi:"name"`
	PrivateIPAllocationMethod pulumi.StringPtrInput `pulumi:"privateIPAllocationMethod"`
	PublicIPAddress           SubResourcePtrInput   `pulumi:"publicIPAddress"`
	Subnet                    SubResourcePtrInput   `pulumi:"subnet"`
}

func (VirtualNetworkGatewayIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (i VirtualNetworkGatewayIPConfigurationArgs) ToVirtualNetworkGatewayIPConfigurationOutput() VirtualNetworkGatewayIPConfigurationOutput {
	return i.ToVirtualNetworkGatewayIPConfigurationOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayIPConfigurationArgs) ToVirtualNetworkGatewayIPConfigurationOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayIPConfigurationOutput)
}





type VirtualNetworkGatewayIPConfigurationArrayInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayIPConfigurationArrayOutput() VirtualNetworkGatewayIPConfigurationArrayOutput
	ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(context.Context) VirtualNetworkGatewayIPConfigurationArrayOutput
}

type VirtualNetworkGatewayIPConfigurationArray []VirtualNetworkGatewayIPConfigurationInput

func (VirtualNetworkGatewayIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (i VirtualNetworkGatewayIPConfigurationArray) ToVirtualNetworkGatewayIPConfigurationArrayOutput() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return i.ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayIPConfigurationArray) ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayIPConfigurationArrayOutput)
}

type VirtualNetworkGatewayIPConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationOutput) ToVirtualNetworkGatewayIPConfigurationOutput() VirtualNetworkGatewayIPConfigurationOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationOutput) ToVirtualNetworkGatewayIPConfigurationOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) PublicIPAddress() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *SubResource { return v.PublicIPAddress }).(SubResourcePtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type VirtualNetworkGatewayIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationArrayOutput) ToVirtualNetworkGatewayIPConfigurationArrayOutput() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationArrayOutput) ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationArrayOutput) Index(i pulumi.IntInput) VirtualNetworkGatewayIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkGatewayIPConfiguration {
		return vs[0].([]VirtualNetworkGatewayIPConfiguration)[vs[1].(int)]
	}).(VirtualNetworkGatewayIPConfigurationOutput)
}

type VirtualNetworkGatewayIPConfigurationResponse struct {
	Etag                      string               `pulumi:"etag"`
	Id                        *string              `pulumi:"id"`
	Name                      *string              `pulumi:"name"`
	PrivateIPAddress          string               `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string              `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         string               `pulumi:"provisioningState"`
	PublicIPAddress           *SubResourceResponse `pulumi:"publicIPAddress"`
	Subnet                    *SubResourceResponse `pulumi:"subnet"`
}

type VirtualNetworkGatewayIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) ToVirtualNetworkGatewayIPConfigurationResponseOutput() VirtualNetworkGatewayIPConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) ToVirtualNetworkGatewayIPConfigurationResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) PrivateIPAddress() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) string { return v.PrivateIPAddress }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) PublicIPAddress() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *SubResourceResponse { return v.PublicIPAddress }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type VirtualNetworkGatewayIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkGatewayIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationResponseArrayOutput) ToVirtualNetworkGatewayIPConfigurationResponseArrayOutput() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseArrayOutput) ToVirtualNetworkGatewayIPConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkGatewayIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkGatewayIPConfigurationResponse {
		return vs[0].([]VirtualNetworkGatewayIPConfigurationResponse)[vs[1].(int)]
	}).(VirtualNetworkGatewayIPConfigurationResponseOutput)
}

type VirtualNetworkGatewayResponse struct {
	ActiveActive                   *bool                                          `pulumi:"activeActive"`
	BgpSettings                    *BgpSettingsResponse                           `pulumi:"bgpSettings"`
	CustomRoutes                   *AddressSpaceResponse                          `pulumi:"customRoutes"`
	EnableBgp                      *bool                                          `pulumi:"enableBgp"`
	EnableDnsForwarding            *bool                                          `pulumi:"enableDnsForwarding"`
	EnablePrivateIpAddress         *bool                                          `pulumi:"enablePrivateIpAddress"`
	Etag                           string                                         `pulumi:"etag"`
	ExtendedLocation               *ExtendedLocationResponse                      `pulumi:"extendedLocation"`
	GatewayDefaultSite             *SubResourceResponse                           `pulumi:"gatewayDefaultSite"`
	GatewayType                    *string                                        `pulumi:"gatewayType"`
	Id                             *string                                        `pulumi:"id"`
	InboundDnsForwardingEndpoint   string                                         `pulumi:"inboundDnsForwardingEndpoint"`
	IpConfigurations               []VirtualNetworkGatewayIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location                       *string                                        `pulumi:"location"`
	Name                           string                                         `pulumi:"name"`
	ProvisioningState              string                                         `pulumi:"provisioningState"`
	ResourceGuid                   string                                         `pulumi:"resourceGuid"`
	Sku                            *VirtualNetworkGatewaySkuResponse              `pulumi:"sku"`
	Tags                           map[string]string                              `pulumi:"tags"`
	Type                           string                                         `pulumi:"type"`
	VNetExtendedLocationResourceId *string                                        `pulumi:"vNetExtendedLocationResourceId"`
	VpnClientConfiguration         *VpnClientConfigurationResponse                `pulumi:"vpnClientConfiguration"`
	VpnGatewayGeneration           *string                                        `pulumi:"vpnGatewayGeneration"`
	VpnType                        *string                                        `pulumi:"vpnType"`
}

type VirtualNetworkGatewayResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayResponseOutput) ToVirtualNetworkGatewayResponseOutput() VirtualNetworkGatewayResponseOutput {
	return o
}

func (o VirtualNetworkGatewayResponseOutput) ToVirtualNetworkGatewayResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewayResponseOutput {
	return o
}

func (o VirtualNetworkGatewayResponseOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *bool { return v.ActiveActive }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *BgpSettingsResponse { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) CustomRoutes() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *AddressSpaceResponse { return v.CustomRoutes }).(AddressSpaceResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) EnableDnsForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *bool { return v.EnableDnsForwarding }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) EnablePrivateIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *bool { return v.EnablePrivateIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) GatewayDefaultSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *SubResourceResponse { return v.GatewayDefaultSite }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.GatewayType }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) InboundDnsForwardingEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.InboundDnsForwardingEndpoint }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) []VirtualNetworkGatewayIPConfigurationResponse {
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationResponseArrayOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Sku() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *VirtualNetworkGatewaySkuResponse { return v.Sku }).(VirtualNetworkGatewaySkuResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) VNetExtendedLocationResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.VNetExtendedLocationResourceId }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) VpnClientConfiguration() VpnClientConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *VpnClientConfigurationResponse { return v.VpnClientConfiguration }).(VpnClientConfigurationResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) VpnGatewayGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.VpnGatewayGeneration }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.VpnType }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayResponsePtrOutput) ToVirtualNetworkGatewayResponsePtrOutput() VirtualNetworkGatewayResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewayResponsePtrOutput) ToVirtualNetworkGatewayResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewayResponsePtrOutput) Elem() VirtualNetworkGatewayResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) VirtualNetworkGatewayResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewayResponse
		return ret
	}).(VirtualNetworkGatewayResponseOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ActiveActive
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *BgpSettingsResponse {
		if v == nil {
			return nil
		}
		return v.BgpSettings
	}).(BgpSettingsResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) CustomRoutes() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *AddressSpaceResponse {
		if v == nil {
			return nil
		}
		return v.CustomRoutes
	}).(AddressSpaceResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableBgp
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) EnableDnsForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDnsForwarding
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) EnablePrivateIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePrivateIpAddress
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *ExtendedLocationResponse {
		if v == nil {
			return nil
		}
		return v.ExtendedLocation
	}).(ExtendedLocationResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) GatewayDefaultSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.GatewayDefaultSite
	}).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.GatewayType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) InboundDnsForwardingEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InboundDnsForwardingEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) []VirtualNetworkGatewayIPConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationResponseArrayOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Sku() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *VirtualNetworkGatewaySkuResponse {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(VirtualNetworkGatewaySkuResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) VNetExtendedLocationResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.VNetExtendedLocationResourceId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) VpnClientConfiguration() VpnClientConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *VpnClientConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientConfiguration
	}).(VpnClientConfigurationResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) VpnGatewayGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.VpnGatewayGeneration
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.VpnType
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySku struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type VirtualNetworkGatewaySkuInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuOutput() VirtualNetworkGatewaySkuOutput
	ToVirtualNetworkGatewaySkuOutputWithContext(context.Context) VirtualNetworkGatewaySkuOutput
}

type VirtualNetworkGatewaySkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (VirtualNetworkGatewaySkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySku)(nil)).Elem()
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuOutput() VirtualNetworkGatewaySkuOutput {
	return i.ToVirtualNetworkGatewaySkuOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuOutput)
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return i.ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuOutput).ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx)
}









type VirtualNetworkGatewaySkuPtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput
	ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Context) VirtualNetworkGatewaySkuPtrOutput
}

type virtualNetworkGatewaySkuPtrType VirtualNetworkGatewaySkuArgs

func VirtualNetworkGatewaySkuPtr(v *VirtualNetworkGatewaySkuArgs) VirtualNetworkGatewaySkuPtrInput {
	return (*virtualNetworkGatewaySkuPtrType)(v)
}

func (*virtualNetworkGatewaySkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySku)(nil)).Elem()
}

func (i *virtualNetworkGatewaySkuPtrType) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return i.ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkGatewaySkuPtrType) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuPtrOutput)
}

type VirtualNetworkGatewaySkuOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySku)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuOutput() VirtualNetworkGatewaySkuOutput {
	return o
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuOutput {
	return o
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return o.ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewaySku) *VirtualNetworkGatewaySku {
		return &v
	}).(VirtualNetworkGatewaySkuPtrOutput)
}

func (o VirtualNetworkGatewaySkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySkuPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySku)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuPtrOutput) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuPtrOutput) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuPtrOutput) Elem() VirtualNetworkGatewaySkuOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySku) VirtualNetworkGatewaySku {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewaySku
		return ret
	}).(VirtualNetworkGatewaySkuOutput)
}

func (o VirtualNetworkGatewaySkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySkuResponse struct {
	Capacity int     `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}

type VirtualNetworkGatewaySkuResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySkuResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuResponseOutput) ToVirtualNetworkGatewaySkuResponseOutput() VirtualNetworkGatewaySkuResponseOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponseOutput) ToVirtualNetworkGatewaySkuResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuResponseOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponseOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySkuResponse) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o VirtualNetworkGatewaySkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySkuResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySkuResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) ToVirtualNetworkGatewaySkuResponsePtrOutput() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) ToVirtualNetworkGatewaySkuResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Elem() VirtualNetworkGatewaySkuResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) VirtualNetworkGatewaySkuResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewaySkuResponse
		return ret
	}).(VirtualNetworkGatewaySkuResponseOutput)
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringType struct {
	AllowForwardedTraffic     *bool                         `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       *bool                         `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess *bool                         `pulumi:"allowVirtualNetworkAccess"`
	DoNotVerifyRemoteGateways *bool                         `pulumi:"doNotVerifyRemoteGateways"`
	Id                        *string                       `pulumi:"id"`
	Name                      *string                       `pulumi:"name"`
	PeeringState              *string                       `pulumi:"peeringState"`
	RemoteAddressSpace        *AddressSpace                 `pulumi:"remoteAddressSpace"`
	RemoteBgpCommunities      *VirtualNetworkBgpCommunities `pulumi:"remoteBgpCommunities"`
	RemoteVirtualNetwork      *SubResource                  `pulumi:"remoteVirtualNetwork"`
	Type                      *string                       `pulumi:"type"`
	UseRemoteGateways         *bool                         `pulumi:"useRemoteGateways"`
}





type VirtualNetworkPeeringTypeInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringTypeOutput() VirtualNetworkPeeringTypeOutput
	ToVirtualNetworkPeeringTypeOutputWithContext(context.Context) VirtualNetworkPeeringTypeOutput
}

type VirtualNetworkPeeringTypeArgs struct {
	AllowForwardedTraffic     pulumi.BoolPtrInput                  `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       pulumi.BoolPtrInput                  `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess pulumi.BoolPtrInput                  `pulumi:"allowVirtualNetworkAccess"`
	DoNotVerifyRemoteGateways pulumi.BoolPtrInput                  `pulumi:"doNotVerifyRemoteGateways"`
	Id                        pulumi.StringPtrInput                `pulumi:"id"`
	Name                      pulumi.StringPtrInput                `pulumi:"name"`
	PeeringState              pulumi.StringPtrInput                `pulumi:"peeringState"`
	RemoteAddressSpace        AddressSpacePtrInput                 `pulumi:"remoteAddressSpace"`
	RemoteBgpCommunities      VirtualNetworkBgpCommunitiesPtrInput `pulumi:"remoteBgpCommunities"`
	RemoteVirtualNetwork      SubResourcePtrInput                  `pulumi:"remoteVirtualNetwork"`
	Type                      pulumi.StringPtrInput                `pulumi:"type"`
	UseRemoteGateways         pulumi.BoolPtrInput                  `pulumi:"useRemoteGateways"`
}

func (VirtualNetworkPeeringTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringType)(nil)).Elem()
}

func (i VirtualNetworkPeeringTypeArgs) ToVirtualNetworkPeeringTypeOutput() VirtualNetworkPeeringTypeOutput {
	return i.ToVirtualNetworkPeeringTypeOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringTypeArgs) ToVirtualNetworkPeeringTypeOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringTypeOutput)
}





type VirtualNetworkPeeringTypeArrayInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringTypeArrayOutput() VirtualNetworkPeeringTypeArrayOutput
	ToVirtualNetworkPeeringTypeArrayOutputWithContext(context.Context) VirtualNetworkPeeringTypeArrayOutput
}

type VirtualNetworkPeeringTypeArray []VirtualNetworkPeeringTypeInput

func (VirtualNetworkPeeringTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPeeringType)(nil)).Elem()
}

func (i VirtualNetworkPeeringTypeArray) ToVirtualNetworkPeeringTypeArrayOutput() VirtualNetworkPeeringTypeArrayOutput {
	return i.ToVirtualNetworkPeeringTypeArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringTypeArray) ToVirtualNetworkPeeringTypeArrayOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringTypeArrayOutput)
}

type VirtualNetworkPeeringTypeOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringType)(nil)).Elem()
}

func (o VirtualNetworkPeeringTypeOutput) ToVirtualNetworkPeeringTypeOutput() VirtualNetworkPeeringTypeOutput {
	return o
}

func (o VirtualNetworkPeeringTypeOutput) ToVirtualNetworkPeeringTypeOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeOutput {
	return o
}

func (o VirtualNetworkPeeringTypeOutput) AllowForwardedTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.AllowForwardedTraffic }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) AllowGatewayTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.AllowGatewayTransit }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) AllowVirtualNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.AllowVirtualNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) DoNotVerifyRemoteGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.DoNotVerifyRemoteGateways }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) PeeringState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.PeeringState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) RemoteAddressSpace() AddressSpacePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *AddressSpace { return v.RemoteAddressSpace }).(AddressSpacePtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) RemoteBgpCommunities() VirtualNetworkBgpCommunitiesPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *VirtualNetworkBgpCommunities { return v.RemoteBgpCommunities }).(VirtualNetworkBgpCommunitiesPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) RemoteVirtualNetwork() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *SubResource { return v.RemoteVirtualNetwork }).(SubResourcePtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) UseRemoteGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.UseRemoteGateways }).(pulumi.BoolPtrOutput)
}

type VirtualNetworkPeeringTypeArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPeeringType)(nil)).Elem()
}

func (o VirtualNetworkPeeringTypeArrayOutput) ToVirtualNetworkPeeringTypeArrayOutput() VirtualNetworkPeeringTypeArrayOutput {
	return o
}

func (o VirtualNetworkPeeringTypeArrayOutput) ToVirtualNetworkPeeringTypeArrayOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeArrayOutput {
	return o
}

func (o VirtualNetworkPeeringTypeArrayOutput) Index(i pulumi.IntInput) VirtualNetworkPeeringTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkPeeringType {
		return vs[0].([]VirtualNetworkPeeringType)[vs[1].(int)]
	}).(VirtualNetworkPeeringTypeOutput)
}

type VirtualNetworkPeeringResponse struct {
	AllowForwardedTraffic     *bool                                 `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       *bool                                 `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess *bool                                 `pulumi:"allowVirtualNetworkAccess"`
	DoNotVerifyRemoteGateways *bool                                 `pulumi:"doNotVerifyRemoteGateways"`
	Etag                      string                                `pulumi:"etag"`
	Id                        *string                               `pulumi:"id"`
	Name                      *string                               `pulumi:"name"`
	PeeringState              *string                               `pulumi:"peeringState"`
	ProvisioningState         string                                `pulumi:"provisioningState"`
	RemoteAddressSpace        *AddressSpaceResponse                 `pulumi:"remoteAddressSpace"`
	RemoteBgpCommunities      *VirtualNetworkBgpCommunitiesResponse `pulumi:"remoteBgpCommunities"`
	RemoteVirtualNetwork      *SubResourceResponse                  `pulumi:"remoteVirtualNetwork"`
	ResourceGuid              string                                `pulumi:"resourceGuid"`
	Type                      *string                               `pulumi:"type"`
	UseRemoteGateways         *bool                                 `pulumi:"useRemoteGateways"`
}

type VirtualNetworkPeeringResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringResponse)(nil)).Elem()
}

func (o VirtualNetworkPeeringResponseOutput) ToVirtualNetworkPeeringResponseOutput() VirtualNetworkPeeringResponseOutput {
	return o
}

func (o VirtualNetworkPeeringResponseOutput) ToVirtualNetworkPeeringResponseOutputWithContext(ctx context.Context) VirtualNetworkPeeringResponseOutput {
	return o
}

func (o VirtualNetworkPeeringResponseOutput) AllowForwardedTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.AllowForwardedTraffic }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) AllowGatewayTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.AllowGatewayTransit }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) AllowVirtualNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.AllowVirtualNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) DoNotVerifyRemoteGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.DoNotVerifyRemoteGateways }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualNetworkPeeringResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) PeeringState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.PeeringState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkPeeringResponseOutput) RemoteAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *AddressSpaceResponse { return v.RemoteAddressSpace }).(AddressSpaceResponsePtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) RemoteBgpCommunities() VirtualNetworkBgpCommunitiesResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *VirtualNetworkBgpCommunitiesResponse {
		return v.RemoteBgpCommunities
	}).(VirtualNetworkBgpCommunitiesResponsePtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) RemoteVirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *SubResourceResponse { return v.RemoteVirtualNetwork }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o VirtualNetworkPeeringResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) UseRemoteGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.UseRemoteGateways }).(pulumi.BoolPtrOutput)
}

type VirtualNetworkPeeringResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPeeringResponse)(nil)).Elem()
}

func (o VirtualNetworkPeeringResponseArrayOutput) ToVirtualNetworkPeeringResponseArrayOutput() VirtualNetworkPeeringResponseArrayOutput {
	return o
}

func (o VirtualNetworkPeeringResponseArrayOutput) ToVirtualNetworkPeeringResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkPeeringResponseArrayOutput {
	return o
}

func (o VirtualNetworkPeeringResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkPeeringResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkPeeringResponse {
		return vs[0].([]VirtualNetworkPeeringResponse)[vs[1].(int)]
	}).(VirtualNetworkPeeringResponseOutput)
}

type VirtualNetworkTapType struct {
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfiguration         `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	DestinationNetworkInterfaceIPConfiguration     *NetworkInterfaceIPConfiguration `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	DestinationPort                                *int                             `pulumi:"destinationPort"`
	Id                                             *string                          `pulumi:"id"`
	Location                                       *string                          `pulumi:"location"`
	Tags                                           map[string]string                `pulumi:"tags"`
}


func (val *VirtualNetworkTapType) Defaults() *VirtualNetworkTapType {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DestinationLoadBalancerFrontEndIPConfiguration = tmp.DestinationLoadBalancerFrontEndIPConfiguration.Defaults()

	tmp.DestinationNetworkInterfaceIPConfiguration = tmp.DestinationNetworkInterfaceIPConfiguration.Defaults()

	return &tmp
}





type VirtualNetworkTapTypeInput interface {
	pulumi.Input

	ToVirtualNetworkTapTypeOutput() VirtualNetworkTapTypeOutput
	ToVirtualNetworkTapTypeOutputWithContext(context.Context) VirtualNetworkTapTypeOutput
}

type VirtualNetworkTapTypeArgs struct {
	DestinationLoadBalancerFrontEndIPConfiguration FrontendIPConfigurationPtrInput         `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	DestinationNetworkInterfaceIPConfiguration     NetworkInterfaceIPConfigurationPtrInput `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	DestinationPort                                pulumi.IntPtrInput                      `pulumi:"destinationPort"`
	Id                                             pulumi.StringPtrInput                   `pulumi:"id"`
	Location                                       pulumi.StringPtrInput                   `pulumi:"location"`
	Tags                                           pulumi.StringMapInput                   `pulumi:"tags"`
}


func (val *VirtualNetworkTapTypeArgs) Defaults() *VirtualNetworkTapTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (VirtualNetworkTapTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkTapType)(nil)).Elem()
}

func (i VirtualNetworkTapTypeArgs) ToVirtualNetworkTapTypeOutput() VirtualNetworkTapTypeOutput {
	return i.ToVirtualNetworkTapTypeOutputWithContext(context.Background())
}

func (i VirtualNetworkTapTypeArgs) ToVirtualNetworkTapTypeOutputWithContext(ctx context.Context) VirtualNetworkTapTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkTapTypeOutput)
}

func (i VirtualNetworkTapTypeArgs) ToVirtualNetworkTapTypePtrOutput() VirtualNetworkTapTypePtrOutput {
	return i.ToVirtualNetworkTapTypePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkTapTypeArgs) ToVirtualNetworkTapTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTapTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkTapTypeOutput).ToVirtualNetworkTapTypePtrOutputWithContext(ctx)
}









type VirtualNetworkTapTypePtrInput interface {
	pulumi.Input

	ToVirtualNetworkTapTypePtrOutput() VirtualNetworkTapTypePtrOutput
	ToVirtualNetworkTapTypePtrOutputWithContext(context.Context) VirtualNetworkTapTypePtrOutput
}

type virtualNetworkTapTypePtrType VirtualNetworkTapTypeArgs

func VirtualNetworkTapTypePtr(v *VirtualNetworkTapTypeArgs) VirtualNetworkTapTypePtrInput {
	return (*virtualNetworkTapTypePtrType)(v)
}

func (*virtualNetworkTapTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkTapType)(nil)).Elem()
}

func (i *virtualNetworkTapTypePtrType) ToVirtualNetworkTapTypePtrOutput() VirtualNetworkTapTypePtrOutput {
	return i.ToVirtualNetworkTapTypePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkTapTypePtrType) ToVirtualNetworkTapTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTapTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkTapTypePtrOutput)
}





type VirtualNetworkTapTypeArrayInput interface {
	pulumi.Input

	ToVirtualNetworkTapTypeArrayOutput() VirtualNetworkTapTypeArrayOutput
	ToVirtualNetworkTapTypeArrayOutputWithContext(context.Context) VirtualNetworkTapTypeArrayOutput
}

type VirtualNetworkTapTypeArray []VirtualNetworkTapTypeInput

func (VirtualNetworkTapTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkTapType)(nil)).Elem()
}

func (i VirtualNetworkTapTypeArray) ToVirtualNetworkTapTypeArrayOutput() VirtualNetworkTapTypeArrayOutput {
	return i.ToVirtualNetworkTapTypeArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkTapTypeArray) ToVirtualNetworkTapTypeArrayOutputWithContext(ctx context.Context) VirtualNetworkTapTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkTapTypeArrayOutput)
}

type VirtualNetworkTapTypeOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkTapType)(nil)).Elem()
}

func (o VirtualNetworkTapTypeOutput) ToVirtualNetworkTapTypeOutput() VirtualNetworkTapTypeOutput {
	return o
}

func (o VirtualNetworkTapTypeOutput) ToVirtualNetworkTapTypeOutputWithContext(ctx context.Context) VirtualNetworkTapTypeOutput {
	return o
}

func (o VirtualNetworkTapTypeOutput) ToVirtualNetworkTapTypePtrOutput() VirtualNetworkTapTypePtrOutput {
	return o.ToVirtualNetworkTapTypePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkTapTypeOutput) ToVirtualNetworkTapTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTapTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkTapType) *VirtualNetworkTapType {
		return &v
	}).(VirtualNetworkTapTypePtrOutput)
}

func (o VirtualNetworkTapTypeOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) *FrontendIPConfiguration {
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationPtrOutput)
}

func (o VirtualNetworkTapTypeOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) *NetworkInterfaceIPConfiguration {
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationPtrOutput)
}

func (o VirtualNetworkTapTypeOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) *int { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkTapTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type VirtualNetworkTapTypePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkTapType)(nil)).Elem()
}

func (o VirtualNetworkTapTypePtrOutput) ToVirtualNetworkTapTypePtrOutput() VirtualNetworkTapTypePtrOutput {
	return o
}

func (o VirtualNetworkTapTypePtrOutput) ToVirtualNetworkTapTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTapTypePtrOutput {
	return o
}

func (o VirtualNetworkTapTypePtrOutput) Elem() VirtualNetworkTapTypeOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) VirtualNetworkTapType {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkTapType
		return ret
	}).(VirtualNetworkTapTypeOutput)
}

func (o VirtualNetworkTapTypePtrOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) *FrontendIPConfiguration {
		if v == nil {
			return nil
		}
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationPtrOutput)
}

func (o VirtualNetworkTapTypePtrOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) *NetworkInterfaceIPConfiguration {
		if v == nil {
			return nil
		}
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationPtrOutput)
}

func (o VirtualNetworkTapTypePtrOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) *int {
		if v == nil {
			return nil
		}
		return v.DestinationPort
	}).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkTapTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type VirtualNetworkTapTypeArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkTapType)(nil)).Elem()
}

func (o VirtualNetworkTapTypeArrayOutput) ToVirtualNetworkTapTypeArrayOutput() VirtualNetworkTapTypeArrayOutput {
	return o
}

func (o VirtualNetworkTapTypeArrayOutput) ToVirtualNetworkTapTypeArrayOutputWithContext(ctx context.Context) VirtualNetworkTapTypeArrayOutput {
	return o
}

func (o VirtualNetworkTapTypeArrayOutput) Index(i pulumi.IntInput) VirtualNetworkTapTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkTapType {
		return vs[0].([]VirtualNetworkTapType)[vs[1].(int)]
	}).(VirtualNetworkTapTypeOutput)
}

type VirtualNetworkTapResponse struct {
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfigurationResponse           `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	DestinationNetworkInterfaceIPConfiguration     *NetworkInterfaceIPConfigurationResponse   `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	DestinationPort                                *int                                       `pulumi:"destinationPort"`
	Etag                                           string                                     `pulumi:"etag"`
	Id                                             *string                                    `pulumi:"id"`
	Location                                       *string                                    `pulumi:"location"`
	Name                                           string                                     `pulumi:"name"`
	NetworkInterfaceTapConfigurations              []NetworkInterfaceTapConfigurationResponse `pulumi:"networkInterfaceTapConfigurations"`
	ProvisioningState                              string                                     `pulumi:"provisioningState"`
	ResourceGuid                                   string                                     `pulumi:"resourceGuid"`
	Tags                                           map[string]string                          `pulumi:"tags"`
	Type                                           string                                     `pulumi:"type"`
}


func (val *VirtualNetworkTapResponse) Defaults() *VirtualNetworkTapResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DestinationLoadBalancerFrontEndIPConfiguration = tmp.DestinationLoadBalancerFrontEndIPConfiguration.Defaults()

	tmp.DestinationNetworkInterfaceIPConfiguration = tmp.DestinationNetworkInterfaceIPConfiguration.Defaults()

	return &tmp
}

type VirtualNetworkTapResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkTapResponse)(nil)).Elem()
}

func (o VirtualNetworkTapResponseOutput) ToVirtualNetworkTapResponseOutput() VirtualNetworkTapResponseOutput {
	return o
}

func (o VirtualNetworkTapResponseOutput) ToVirtualNetworkTapResponseOutputWithContext(ctx context.Context) VirtualNetworkTapResponseOutput {
	return o
}

func (o VirtualNetworkTapResponseOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) *FrontendIPConfigurationResponse {
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationResponsePtrOutput)
}

func (o VirtualNetworkTapResponseOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) *NetworkInterfaceIPConfigurationResponse {
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponsePtrOutput)
}

func (o VirtualNetworkTapResponseOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) *int { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkTapResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapResponseOutput) NetworkInterfaceTapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) []NetworkInterfaceTapConfigurationResponse {
		return v.NetworkInterfaceTapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

func (o VirtualNetworkTapResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapResponseOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkTapResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VirtualNetworkTapResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkTapResponse)(nil)).Elem()
}

func (o VirtualNetworkTapResponsePtrOutput) ToVirtualNetworkTapResponsePtrOutput() VirtualNetworkTapResponsePtrOutput {
	return o
}

func (o VirtualNetworkTapResponsePtrOutput) ToVirtualNetworkTapResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkTapResponsePtrOutput {
	return o
}

func (o VirtualNetworkTapResponsePtrOutput) Elem() VirtualNetworkTapResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) VirtualNetworkTapResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkTapResponse
		return ret
	}).(VirtualNetworkTapResponseOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *FrontendIPConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationResponsePtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *NetworkInterfaceIPConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponsePtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *int {
		if v == nil {
			return nil
		}
		return v.DestinationPort
	}).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) NetworkInterfaceTapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) []NetworkInterfaceTapConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceTapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkTapResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkTapResponse)(nil)).Elem()
}

func (o VirtualNetworkTapResponseArrayOutput) ToVirtualNetworkTapResponseArrayOutput() VirtualNetworkTapResponseArrayOutput {
	return o
}

func (o VirtualNetworkTapResponseArrayOutput) ToVirtualNetworkTapResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkTapResponseArrayOutput {
	return o
}

func (o VirtualNetworkTapResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkTapResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkTapResponse {
		return vs[0].([]VirtualNetworkTapResponse)[vs[1].(int)]
	}).(VirtualNetworkTapResponseOutput)
}

type VnetRoute struct {
	StaticRoutes []StaticRoute `pulumi:"staticRoutes"`
}





type VnetRouteInput interface {
	pulumi.Input

	ToVnetRouteOutput() VnetRouteOutput
	ToVnetRouteOutputWithContext(context.Context) VnetRouteOutput
}

type VnetRouteArgs struct {
	StaticRoutes StaticRouteArrayInput `pulumi:"staticRoutes"`
}

func (VnetRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRoute)(nil)).Elem()
}

func (i VnetRouteArgs) ToVnetRouteOutput() VnetRouteOutput {
	return i.ToVnetRouteOutputWithContext(context.Background())
}

func (i VnetRouteArgs) ToVnetRouteOutputWithContext(ctx context.Context) VnetRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetRouteOutput)
}

func (i VnetRouteArgs) ToVnetRoutePtrOutput() VnetRoutePtrOutput {
	return i.ToVnetRoutePtrOutputWithContext(context.Background())
}

func (i VnetRouteArgs) ToVnetRoutePtrOutputWithContext(ctx context.Context) VnetRoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetRouteOutput).ToVnetRoutePtrOutputWithContext(ctx)
}









type VnetRoutePtrInput interface {
	pulumi.Input

	ToVnetRoutePtrOutput() VnetRoutePtrOutput
	ToVnetRoutePtrOutputWithContext(context.Context) VnetRoutePtrOutput
}

type vnetRoutePtrType VnetRouteArgs

func VnetRoutePtr(v *VnetRouteArgs) VnetRoutePtrInput {
	return (*vnetRoutePtrType)(v)
}

func (*vnetRoutePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VnetRoute)(nil)).Elem()
}

func (i *vnetRoutePtrType) ToVnetRoutePtrOutput() VnetRoutePtrOutput {
	return i.ToVnetRoutePtrOutputWithContext(context.Background())
}

func (i *vnetRoutePtrType) ToVnetRoutePtrOutputWithContext(ctx context.Context) VnetRoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetRoutePtrOutput)
}

type VnetRouteOutput struct{ *pulumi.OutputState }

func (VnetRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRoute)(nil)).Elem()
}

func (o VnetRouteOutput) ToVnetRouteOutput() VnetRouteOutput {
	return o
}

func (o VnetRouteOutput) ToVnetRouteOutputWithContext(ctx context.Context) VnetRouteOutput {
	return o
}

func (o VnetRouteOutput) ToVnetRoutePtrOutput() VnetRoutePtrOutput {
	return o.ToVnetRoutePtrOutputWithContext(context.Background())
}

func (o VnetRouteOutput) ToVnetRoutePtrOutputWithContext(ctx context.Context) VnetRoutePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VnetRoute) *VnetRoute {
		return &v
	}).(VnetRoutePtrOutput)
}

func (o VnetRouteOutput) StaticRoutes() StaticRouteArrayOutput {
	return o.ApplyT(func(v VnetRoute) []StaticRoute { return v.StaticRoutes }).(StaticRouteArrayOutput)
}

type VnetRoutePtrOutput struct{ *pulumi.OutputState }

func (VnetRoutePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VnetRoute)(nil)).Elem()
}

func (o VnetRoutePtrOutput) ToVnetRoutePtrOutput() VnetRoutePtrOutput {
	return o
}

func (o VnetRoutePtrOutput) ToVnetRoutePtrOutputWithContext(ctx context.Context) VnetRoutePtrOutput {
	return o
}

func (o VnetRoutePtrOutput) Elem() VnetRouteOutput {
	return o.ApplyT(func(v *VnetRoute) VnetRoute {
		if v != nil {
			return *v
		}
		var ret VnetRoute
		return ret
	}).(VnetRouteOutput)
}

func (o VnetRoutePtrOutput) StaticRoutes() StaticRouteArrayOutput {
	return o.ApplyT(func(v *VnetRoute) []StaticRoute {
		if v == nil {
			return nil
		}
		return v.StaticRoutes
	}).(StaticRouteArrayOutput)
}

type VnetRouteResponse struct {
	StaticRoutes []StaticRouteResponse `pulumi:"staticRoutes"`
}

type VnetRouteResponseOutput struct{ *pulumi.OutputState }

func (VnetRouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRouteResponse)(nil)).Elem()
}

func (o VnetRouteResponseOutput) ToVnetRouteResponseOutput() VnetRouteResponseOutput {
	return o
}

func (o VnetRouteResponseOutput) ToVnetRouteResponseOutputWithContext(ctx context.Context) VnetRouteResponseOutput {
	return o
}

func (o VnetRouteResponseOutput) StaticRoutes() StaticRouteResponseArrayOutput {
	return o.ApplyT(func(v VnetRouteResponse) []StaticRouteResponse { return v.StaticRoutes }).(StaticRouteResponseArrayOutput)
}

type VnetRouteResponsePtrOutput struct{ *pulumi.OutputState }

func (VnetRouteResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VnetRouteResponse)(nil)).Elem()
}

func (o VnetRouteResponsePtrOutput) ToVnetRouteResponsePtrOutput() VnetRouteResponsePtrOutput {
	return o
}

func (o VnetRouteResponsePtrOutput) ToVnetRouteResponsePtrOutputWithContext(ctx context.Context) VnetRouteResponsePtrOutput {
	return o
}

func (o VnetRouteResponsePtrOutput) Elem() VnetRouteResponseOutput {
	return o.ApplyT(func(v *VnetRouteResponse) VnetRouteResponse {
		if v != nil {
			return *v
		}
		var ret VnetRouteResponse
		return ret
	}).(VnetRouteResponseOutput)
}

func (o VnetRouteResponsePtrOutput) StaticRoutes() StaticRouteResponseArrayOutput {
	return o.ApplyT(func(v *VnetRouteResponse) []StaticRouteResponse {
		if v == nil {
			return nil
		}
		return v.StaticRoutes
	}).(StaticRouteResponseArrayOutput)
}

type VpnClientConfiguration struct {
	AadAudience                  *string                       `pulumi:"aadAudience"`
	AadIssuer                    *string                       `pulumi:"aadIssuer"`
	AadTenant                    *string                       `pulumi:"aadTenant"`
	RadiusServerAddress          *string                       `pulumi:"radiusServerAddress"`
	RadiusServerSecret           *string                       `pulumi:"radiusServerSecret"`
	RadiusServers                []RadiusServer                `pulumi:"radiusServers"`
	VpnAuthenticationTypes       []string                      `pulumi:"vpnAuthenticationTypes"`
	VpnClientAddressPool         *AddressSpace                 `pulumi:"vpnClientAddressPool"`
	VpnClientIpsecPolicies       []IpsecPolicy                 `pulumi:"vpnClientIpsecPolicies"`
	VpnClientProtocols           []string                      `pulumi:"vpnClientProtocols"`
	VpnClientRevokedCertificates []VpnClientRevokedCertificate `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    []VpnClientRootCertificate    `pulumi:"vpnClientRootCertificates"`
}





type VpnClientConfigurationInput interface {
	pulumi.Input

	ToVpnClientConfigurationOutput() VpnClientConfigurationOutput
	ToVpnClientConfigurationOutputWithContext(context.Context) VpnClientConfigurationOutput
}

type VpnClientConfigurationArgs struct {
	AadAudience                  pulumi.StringPtrInput                 `pulumi:"aadAudience"`
	AadIssuer                    pulumi.StringPtrInput                 `pulumi:"aadIssuer"`
	AadTenant                    pulumi.StringPtrInput                 `pulumi:"aadTenant"`
	RadiusServerAddress          pulumi.StringPtrInput                 `pulumi:"radiusServerAddress"`
	RadiusServerSecret           pulumi.StringPtrInput                 `pulumi:"radiusServerSecret"`
	RadiusServers                RadiusServerArrayInput                `pulumi:"radiusServers"`
	VpnAuthenticationTypes       pulumi.StringArrayInput               `pulumi:"vpnAuthenticationTypes"`
	VpnClientAddressPool         AddressSpacePtrInput                  `pulumi:"vpnClientAddressPool"`
	VpnClientIpsecPolicies       IpsecPolicyArrayInput                 `pulumi:"vpnClientIpsecPolicies"`
	VpnClientProtocols           pulumi.StringArrayInput               `pulumi:"vpnClientProtocols"`
	VpnClientRevokedCertificates VpnClientRevokedCertificateArrayInput `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    VpnClientRootCertificateArrayInput    `pulumi:"vpnClientRootCertificates"`
}

func (VpnClientConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfiguration)(nil)).Elem()
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationOutput() VpnClientConfigurationOutput {
	return i.ToVpnClientConfigurationOutputWithContext(context.Background())
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationOutputWithContext(ctx context.Context) VpnClientConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationOutput)
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return i.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationOutput).ToVpnClientConfigurationPtrOutputWithContext(ctx)
}









type VpnClientConfigurationPtrInput interface {
	pulumi.Input

	ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput
	ToVpnClientConfigurationPtrOutputWithContext(context.Context) VpnClientConfigurationPtrOutput
}

type vpnClientConfigurationPtrType VpnClientConfigurationArgs

func VpnClientConfigurationPtr(v *VpnClientConfigurationArgs) VpnClientConfigurationPtrInput {
	return (*vpnClientConfigurationPtrType)(v)
}

func (*vpnClientConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfiguration)(nil)).Elem()
}

func (i *vpnClientConfigurationPtrType) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return i.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (i *vpnClientConfigurationPtrType) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationPtrOutput)
}

type VpnClientConfigurationOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfiguration)(nil)).Elem()
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationOutput() VpnClientConfigurationOutput {
	return o
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationOutputWithContext(ctx context.Context) VpnClientConfigurationOutput {
	return o
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return o.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnClientConfiguration) *VpnClientConfiguration {
		return &v
	}).(VpnClientConfigurationPtrOutput)
}

func (o VpnClientConfigurationOutput) AadAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.AadAudience }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) AadIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.AadIssuer }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) AadTenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.AadTenant }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) RadiusServers() RadiusServerArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []RadiusServer { return v.RadiusServers }).(RadiusServerArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnAuthenticationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []string { return v.VpnAuthenticationTypes }).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnClientAddressPool() AddressSpacePtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *AddressSpace { return v.VpnClientAddressPool }).(AddressSpacePtrOutput)
}

func (o VpnClientConfigurationOutput) VpnClientIpsecPolicies() IpsecPolicyArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []IpsecPolicy { return v.VpnClientIpsecPolicies }).(IpsecPolicyArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []string { return v.VpnClientProtocols }).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []VpnClientRevokedCertificate { return v.VpnClientRevokedCertificates }).(VpnClientRevokedCertificateArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnClientRootCertificates() VpnClientRootCertificateArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []VpnClientRootCertificate { return v.VpnClientRootCertificates }).(VpnClientRootCertificateArrayOutput)
}

type VpnClientConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfiguration)(nil)).Elem()
}

func (o VpnClientConfigurationPtrOutput) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return o
}

func (o VpnClientConfigurationPtrOutput) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return o
}

func (o VpnClientConfigurationPtrOutput) Elem() VpnClientConfigurationOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) VpnClientConfiguration {
		if v != nil {
			return *v
		}
		var ret VpnClientConfiguration
		return ret
	}).(VpnClientConfigurationOutput)
}

func (o VpnClientConfigurationPtrOutput) AadAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AadAudience
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) AadIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AadIssuer
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) AadTenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AadTenant
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerAddress
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerSecret
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) RadiusServers() RadiusServerArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []RadiusServer {
		if v == nil {
			return nil
		}
		return v.RadiusServers
	}).(RadiusServerArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnAuthenticationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.VpnAuthenticationTypes
	}).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientAddressPool() AddressSpacePtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *AddressSpace {
		if v == nil {
			return nil
		}
		return v.VpnClientAddressPool
	}).(AddressSpacePtrOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientIpsecPolicies() IpsecPolicyArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []IpsecPolicy {
		if v == nil {
			return nil
		}
		return v.VpnClientIpsecPolicies
	}).(IpsecPolicyArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.VpnClientProtocols
	}).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []VpnClientRevokedCertificate {
		if v == nil {
			return nil
		}
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientRootCertificates() VpnClientRootCertificateArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []VpnClientRootCertificate {
		if v == nil {
			return nil
		}
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateArrayOutput)
}

type VpnClientConfigurationResponse struct {
	AadAudience                  *string                               `pulumi:"aadAudience"`
	AadIssuer                    *string                               `pulumi:"aadIssuer"`
	AadTenant                    *string                               `pulumi:"aadTenant"`
	RadiusServerAddress          *string                               `pulumi:"radiusServerAddress"`
	RadiusServerSecret           *string                               `pulumi:"radiusServerSecret"`
	RadiusServers                []RadiusServerResponse                `pulumi:"radiusServers"`
	VpnAuthenticationTypes       []string                              `pulumi:"vpnAuthenticationTypes"`
	VpnClientAddressPool         *AddressSpaceResponse                 `pulumi:"vpnClientAddressPool"`
	VpnClientIpsecPolicies       []IpsecPolicyResponse                 `pulumi:"vpnClientIpsecPolicies"`
	VpnClientProtocols           []string                              `pulumi:"vpnClientProtocols"`
	VpnClientRevokedCertificates []VpnClientRevokedCertificateResponse `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    []VpnClientRootCertificateResponse    `pulumi:"vpnClientRootCertificates"`
}

type VpnClientConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfigurationResponse)(nil)).Elem()
}

func (o VpnClientConfigurationResponseOutput) ToVpnClientConfigurationResponseOutput() VpnClientConfigurationResponseOutput {
	return o
}

func (o VpnClientConfigurationResponseOutput) ToVpnClientConfigurationResponseOutputWithContext(ctx context.Context) VpnClientConfigurationResponseOutput {
	return o
}

func (o VpnClientConfigurationResponseOutput) AadAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.AadAudience }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) AadIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.AadIssuer }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) AadTenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.AadTenant }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) RadiusServers() RadiusServerResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []RadiusServerResponse { return v.RadiusServers }).(RadiusServerResponseArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnAuthenticationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []string { return v.VpnAuthenticationTypes }).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientAddressPool() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *AddressSpaceResponse { return v.VpnClientAddressPool }).(AddressSpaceResponsePtrOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientIpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []IpsecPolicyResponse { return v.VpnClientIpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []string { return v.VpnClientProtocols }).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []VpnClientRevokedCertificateResponse {
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateResponseArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientRootCertificates() VpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []VpnClientRootCertificateResponse {
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateResponseArrayOutput)
}

type VpnClientConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfigurationResponse)(nil)).Elem()
}

func (o VpnClientConfigurationResponsePtrOutput) ToVpnClientConfigurationResponsePtrOutput() VpnClientConfigurationResponsePtrOutput {
	return o
}

func (o VpnClientConfigurationResponsePtrOutput) ToVpnClientConfigurationResponsePtrOutputWithContext(ctx context.Context) VpnClientConfigurationResponsePtrOutput {
	return o
}

func (o VpnClientConfigurationResponsePtrOutput) Elem() VpnClientConfigurationResponseOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) VpnClientConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VpnClientConfigurationResponse
		return ret
	}).(VpnClientConfigurationResponseOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) AadAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadAudience
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) AadIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadIssuer
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) AadTenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadTenant
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerAddress
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerSecret
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) RadiusServers() RadiusServerResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []RadiusServerResponse {
		if v == nil {
			return nil
		}
		return v.RadiusServers
	}).(RadiusServerResponseArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnAuthenticationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.VpnAuthenticationTypes
	}).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientAddressPool() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *AddressSpaceResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientAddressPool
	}).(AddressSpaceResponsePtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientIpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []IpsecPolicyResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientIpsecPolicies
	}).(IpsecPolicyResponseArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.VpnClientProtocols
	}).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []VpnClientRevokedCertificateResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateResponseArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientRootCertificates() VpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []VpnClientRootCertificateResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateResponseArrayOutput)
}

type VpnClientConnectionHealthDetailResponse struct {
	EgressBytesTransferred    float64 `pulumi:"egressBytesTransferred"`
	EgressPacketsTransferred  float64 `pulumi:"egressPacketsTransferred"`
	IngressBytesTransferred   float64 `pulumi:"ingressBytesTransferred"`
	IngressPacketsTransferred float64 `pulumi:"ingressPacketsTransferred"`
	MaxBandwidth              float64 `pulumi:"maxBandwidth"`
	MaxPacketsPerSecond       float64 `pulumi:"maxPacketsPerSecond"`
	PrivateIpAddress          string  `pulumi:"privateIpAddress"`
	PublicIpAddress           string  `pulumi:"publicIpAddress"`
	VpnConnectionDuration     float64 `pulumi:"vpnConnectionDuration"`
	VpnConnectionId           string  `pulumi:"vpnConnectionId"`
	VpnConnectionTime         string  `pulumi:"vpnConnectionTime"`
	VpnUserName               string  `pulumi:"vpnUserName"`
}

type VpnClientConnectionHealthDetailResponseOutput struct{ *pulumi.OutputState }

func (VpnClientConnectionHealthDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConnectionHealthDetailResponse)(nil)).Elem()
}

func (o VpnClientConnectionHealthDetailResponseOutput) ToVpnClientConnectionHealthDetailResponseOutput() VpnClientConnectionHealthDetailResponseOutput {
	return o
}

func (o VpnClientConnectionHealthDetailResponseOutput) ToVpnClientConnectionHealthDetailResponseOutputWithContext(ctx context.Context) VpnClientConnectionHealthDetailResponseOutput {
	return o
}

func (o VpnClientConnectionHealthDetailResponseOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) EgressPacketsTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.EgressPacketsTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) IngressPacketsTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.IngressPacketsTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) MaxBandwidth() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.MaxBandwidth }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) MaxPacketsPerSecond() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.MaxPacketsPerSecond }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) string { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

func (o VpnClientConnectionHealthDetailResponseOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) string { return v.PublicIpAddress }).(pulumi.StringOutput)
}

func (o VpnClientConnectionHealthDetailResponseOutput) VpnConnectionDuration() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.VpnConnectionDuration }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) VpnConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) string { return v.VpnConnectionId }).(pulumi.StringOutput)
}

func (o VpnClientConnectionHealthDetailResponseOutput) VpnConnectionTime() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) string { return v.VpnConnectionTime }).(pulumi.StringOutput)
}

func (o VpnClientConnectionHealthDetailResponseOutput) VpnUserName() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) string { return v.VpnUserName }).(pulumi.StringOutput)
}

type VpnClientConnectionHealthDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnClientConnectionHealthDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientConnectionHealthDetailResponse)(nil)).Elem()
}

func (o VpnClientConnectionHealthDetailResponseArrayOutput) ToVpnClientConnectionHealthDetailResponseArrayOutput() VpnClientConnectionHealthDetailResponseArrayOutput {
	return o
}

func (o VpnClientConnectionHealthDetailResponseArrayOutput) ToVpnClientConnectionHealthDetailResponseArrayOutputWithContext(ctx context.Context) VpnClientConnectionHealthDetailResponseArrayOutput {
	return o
}

func (o VpnClientConnectionHealthDetailResponseArrayOutput) Index(i pulumi.IntInput) VpnClientConnectionHealthDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientConnectionHealthDetailResponse {
		return vs[0].([]VpnClientConnectionHealthDetailResponse)[vs[1].(int)]
	}).(VpnClientConnectionHealthDetailResponseOutput)
}

type VpnClientConnectionHealthResponse struct {
	AllocatedIpAddresses         []string `pulumi:"allocatedIpAddresses"`
	TotalEgressBytesTransferred  float64  `pulumi:"totalEgressBytesTransferred"`
	TotalIngressBytesTransferred float64  `pulumi:"totalIngressBytesTransferred"`
	VpnClientConnectionsCount    *int     `pulumi:"vpnClientConnectionsCount"`
}

type VpnClientConnectionHealthResponseOutput struct{ *pulumi.OutputState }

func (VpnClientConnectionHealthResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConnectionHealthResponse)(nil)).Elem()
}

func (o VpnClientConnectionHealthResponseOutput) ToVpnClientConnectionHealthResponseOutput() VpnClientConnectionHealthResponseOutput {
	return o
}

func (o VpnClientConnectionHealthResponseOutput) ToVpnClientConnectionHealthResponseOutputWithContext(ctx context.Context) VpnClientConnectionHealthResponseOutput {
	return o
}

func (o VpnClientConnectionHealthResponseOutput) AllocatedIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthResponse) []string { return v.AllocatedIpAddresses }).(pulumi.StringArrayOutput)
}

func (o VpnClientConnectionHealthResponseOutput) TotalEgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthResponse) float64 { return v.TotalEgressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthResponseOutput) TotalIngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthResponse) float64 { return v.TotalIngressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthResponseOutput) VpnClientConnectionsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthResponse) *int { return v.VpnClientConnectionsCount }).(pulumi.IntPtrOutput)
}

type VpnClientRevokedCertificate struct {
	Id         *string `pulumi:"id"`
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type VpnClientRevokedCertificateInput interface {
	pulumi.Input

	ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput
	ToVpnClientRevokedCertificateOutputWithContext(context.Context) VpnClientRevokedCertificateOutput
}

type VpnClientRevokedCertificateArgs struct {
	Id         pulumi.StringPtrInput `pulumi:"id"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (VpnClientRevokedCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnClientRevokedCertificateArgs) ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput {
	return i.ToVpnClientRevokedCertificateOutputWithContext(context.Background())
}

func (i VpnClientRevokedCertificateArgs) ToVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnClientRevokedCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRevokedCertificateOutput)
}





type VpnClientRevokedCertificateArrayInput interface {
	pulumi.Input

	ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput
	ToVpnClientRevokedCertificateArrayOutputWithContext(context.Context) VpnClientRevokedCertificateArrayOutput
}

type VpnClientRevokedCertificateArray []VpnClientRevokedCertificateInput

func (VpnClientRevokedCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnClientRevokedCertificateArray) ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput {
	return i.ToVpnClientRevokedCertificateArrayOutputWithContext(context.Background())
}

func (i VpnClientRevokedCertificateArray) ToVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRevokedCertificateArrayOutput)
}

type VpnClientRevokedCertificateOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnClientRevokedCertificateOutput) ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput {
	return o
}

func (o VpnClientRevokedCertificateOutput) ToVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnClientRevokedCertificateOutput {
	return o
}

func (o VpnClientRevokedCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnClientRevokedCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnClientRevokedCertificateArrayOutput) ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateArrayOutput) ToVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateArrayOutput) Index(i pulumi.IntInput) VpnClientRevokedCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRevokedCertificate {
		return vs[0].([]VpnClientRevokedCertificate)[vs[1].(int)]
	}).(VpnClientRevokedCertificateOutput)
}

type VpnClientRevokedCertificateResponse struct {
	Etag              string  `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Thumbprint        *string `pulumi:"thumbprint"`
}

type VpnClientRevokedCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnClientRevokedCertificateResponseOutput) ToVpnClientRevokedCertificateResponseOutput() VpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseOutput) ToVpnClientRevokedCertificateResponseOutputWithContext(ctx context.Context) VpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnClientRevokedCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnClientRevokedCertificateResponseArrayOutput) ToVpnClientRevokedCertificateResponseArrayOutput() VpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseArrayOutput) ToVpnClientRevokedCertificateResponseArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnClientRevokedCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRevokedCertificateResponse {
		return vs[0].([]VpnClientRevokedCertificateResponse)[vs[1].(int)]
	}).(VpnClientRevokedCertificateResponseOutput)
}

type VpnClientRootCertificate struct {
	Id             *string `pulumi:"id"`
	Name           *string `pulumi:"name"`
	PublicCertData string  `pulumi:"publicCertData"`
}





type VpnClientRootCertificateInput interface {
	pulumi.Input

	ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput
	ToVpnClientRootCertificateOutputWithContext(context.Context) VpnClientRootCertificateOutput
}

type VpnClientRootCertificateArgs struct {
	Id             pulumi.StringPtrInput `pulumi:"id"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
	PublicCertData pulumi.StringInput    `pulumi:"publicCertData"`
}

func (VpnClientRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificate)(nil)).Elem()
}

func (i VpnClientRootCertificateArgs) ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput {
	return i.ToVpnClientRootCertificateOutputWithContext(context.Background())
}

func (i VpnClientRootCertificateArgs) ToVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnClientRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRootCertificateOutput)
}





type VpnClientRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput
	ToVpnClientRootCertificateArrayOutputWithContext(context.Context) VpnClientRootCertificateArrayOutput
}

type VpnClientRootCertificateArray []VpnClientRootCertificateInput

func (VpnClientRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificate)(nil)).Elem()
}

func (i VpnClientRootCertificateArray) ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput {
	return i.ToVpnClientRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnClientRootCertificateArray) ToVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRootCertificateArrayOutput)
}

type VpnClientRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificate)(nil)).Elem()
}

func (o VpnClientRootCertificateOutput) ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput {
	return o
}

func (o VpnClientRootCertificateOutput) ToVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnClientRootCertificateOutput {
	return o
}

func (o VpnClientRootCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateOutput) PublicCertData() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) string { return v.PublicCertData }).(pulumi.StringOutput)
}

type VpnClientRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificate)(nil)).Elem()
}

func (o VpnClientRootCertificateArrayOutput) ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnClientRootCertificateArrayOutput) ToVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnClientRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnClientRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRootCertificate {
		return vs[0].([]VpnClientRootCertificate)[vs[1].(int)]
	}).(VpnClientRootCertificateOutput)
}

type VpnClientRootCertificateResponse struct {
	Etag              string  `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	PublicCertData    string  `pulumi:"publicCertData"`
}

type VpnClientRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnClientRootCertificateResponseOutput) ToVpnClientRootCertificateResponseOutput() VpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnClientRootCertificateResponseOutput) ToVpnClientRootCertificateResponseOutputWithContext(ctx context.Context) VpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnClientRootCertificateResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnClientRootCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnClientRootCertificateResponseOutput) PublicCertData() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) string { return v.PublicCertData }).(pulumi.StringOutput)
}

type VpnClientRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnClientRootCertificateResponseArrayOutput) ToVpnClientRootCertificateResponseArrayOutput() VpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRootCertificateResponseArrayOutput) ToVpnClientRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnClientRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRootCertificateResponse {
		return vs[0].([]VpnClientRootCertificateResponse)[vs[1].(int)]
	}).(VpnClientRootCertificateResponseOutput)
}

type VpnConnectionType struct {
	ConnectionBandwidth            *int                    `pulumi:"connectionBandwidth"`
	DpdTimeoutSeconds              *int                    `pulumi:"dpdTimeoutSeconds"`
	EnableBgp                      *bool                   `pulumi:"enableBgp"`
	EnableInternetSecurity         *bool                   `pulumi:"enableInternetSecurity"`
	EnableRateLimiting             *bool                   `pulumi:"enableRateLimiting"`
	Id                             *string                 `pulumi:"id"`
	IpsecPolicies                  []IpsecPolicy           `pulumi:"ipsecPolicies"`
	Name                           *string                 `pulumi:"name"`
	RemoteVpnSite                  *SubResource            `pulumi:"remoteVpnSite"`
	RoutingConfiguration           *RoutingConfiguration   `pulumi:"routingConfiguration"`
	RoutingWeight                  *int                    `pulumi:"routingWeight"`
	SharedKey                      *string                 `pulumi:"sharedKey"`
	TrafficSelectorPolicies        []TrafficSelectorPolicy `pulumi:"trafficSelectorPolicies"`
	UseLocalAzureIpAddress         *bool                   `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                   `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      *string                 `pulumi:"vpnConnectionProtocolType"`
	VpnLinkConnections             []VpnSiteLinkConnection `pulumi:"vpnLinkConnections"`
}





type VpnConnectionTypeInput interface {
	pulumi.Input

	ToVpnConnectionTypeOutput() VpnConnectionTypeOutput
	ToVpnConnectionTypeOutputWithContext(context.Context) VpnConnectionTypeOutput
}

type VpnConnectionTypeArgs struct {
	ConnectionBandwidth            pulumi.IntPtrInput              `pulumi:"connectionBandwidth"`
	DpdTimeoutSeconds              pulumi.IntPtrInput              `pulumi:"dpdTimeoutSeconds"`
	EnableBgp                      pulumi.BoolPtrInput             `pulumi:"enableBgp"`
	EnableInternetSecurity         pulumi.BoolPtrInput             `pulumi:"enableInternetSecurity"`
	EnableRateLimiting             pulumi.BoolPtrInput             `pulumi:"enableRateLimiting"`
	Id                             pulumi.StringPtrInput           `pulumi:"id"`
	IpsecPolicies                  IpsecPolicyArrayInput           `pulumi:"ipsecPolicies"`
	Name                           pulumi.StringPtrInput           `pulumi:"name"`
	RemoteVpnSite                  SubResourcePtrInput             `pulumi:"remoteVpnSite"`
	RoutingConfiguration           RoutingConfigurationPtrInput    `pulumi:"routingConfiguration"`
	RoutingWeight                  pulumi.IntPtrInput              `pulumi:"routingWeight"`
	SharedKey                      pulumi.StringPtrInput           `pulumi:"sharedKey"`
	TrafficSelectorPolicies        TrafficSelectorPolicyArrayInput `pulumi:"trafficSelectorPolicies"`
	UseLocalAzureIpAddress         pulumi.BoolPtrInput             `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors pulumi.BoolPtrInput             `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      pulumi.StringPtrInput           `pulumi:"vpnConnectionProtocolType"`
	VpnLinkConnections             VpnSiteLinkConnectionArrayInput `pulumi:"vpnLinkConnections"`
}

func (VpnConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnConnectionType)(nil)).Elem()
}

func (i VpnConnectionTypeArgs) ToVpnConnectionTypeOutput() VpnConnectionTypeOutput {
	return i.ToVpnConnectionTypeOutputWithContext(context.Background())
}

func (i VpnConnectionTypeArgs) ToVpnConnectionTypeOutputWithContext(ctx context.Context) VpnConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionTypeOutput)
}





type VpnConnectionTypeArrayInput interface {
	pulumi.Input

	ToVpnConnectionTypeArrayOutput() VpnConnectionTypeArrayOutput
	ToVpnConnectionTypeArrayOutputWithContext(context.Context) VpnConnectionTypeArrayOutput
}

type VpnConnectionTypeArray []VpnConnectionTypeInput

func (VpnConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnConnectionType)(nil)).Elem()
}

func (i VpnConnectionTypeArray) ToVpnConnectionTypeArrayOutput() VpnConnectionTypeArrayOutput {
	return i.ToVpnConnectionTypeArrayOutputWithContext(context.Background())
}

func (i VpnConnectionTypeArray) ToVpnConnectionTypeArrayOutputWithContext(ctx context.Context) VpnConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionTypeArrayOutput)
}

type VpnConnectionTypeOutput struct{ *pulumi.OutputState }

func (VpnConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnConnectionType)(nil)).Elem()
}

func (o VpnConnectionTypeOutput) ToVpnConnectionTypeOutput() VpnConnectionTypeOutput {
	return o
}

func (o VpnConnectionTypeOutput) ToVpnConnectionTypeOutputWithContext(ctx context.Context) VpnConnectionTypeOutput {
	return o
}

func (o VpnConnectionTypeOutput) ConnectionBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *int { return v.ConnectionBandwidth }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionTypeOutput) DpdTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *int { return v.DpdTimeoutSeconds }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionTypeOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionTypeOutput) EnableInternetSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *bool { return v.EnableInternetSecurity }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionTypeOutput) EnableRateLimiting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *bool { return v.EnableRateLimiting }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionTypeOutput) IpsecPolicies() IpsecPolicyArrayOutput {
	return o.ApplyT(func(v VpnConnectionType) []IpsecPolicy { return v.IpsecPolicies }).(IpsecPolicyArrayOutput)
}

func (o VpnConnectionTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionTypeOutput) RemoteVpnSite() SubResourcePtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *SubResource { return v.RemoteVpnSite }).(SubResourcePtrOutput)
}

func (o VpnConnectionTypeOutput) RoutingConfiguration() RoutingConfigurationPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *RoutingConfiguration { return v.RoutingConfiguration }).(RoutingConfigurationPtrOutput)
}

func (o VpnConnectionTypeOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionTypeOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionTypeOutput) TrafficSelectorPolicies() TrafficSelectorPolicyArrayOutput {
	return o.ApplyT(func(v VpnConnectionType) []TrafficSelectorPolicy { return v.TrafficSelectorPolicies }).(TrafficSelectorPolicyArrayOutput)
}

func (o VpnConnectionTypeOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionTypeOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionTypeOutput) VpnConnectionProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *string { return v.VpnConnectionProtocolType }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionTypeOutput) VpnLinkConnections() VpnSiteLinkConnectionArrayOutput {
	return o.ApplyT(func(v VpnConnectionType) []VpnSiteLinkConnection { return v.VpnLinkConnections }).(VpnSiteLinkConnectionArrayOutput)
}

type VpnConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (VpnConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnConnectionType)(nil)).Elem()
}

func (o VpnConnectionTypeArrayOutput) ToVpnConnectionTypeArrayOutput() VpnConnectionTypeArrayOutput {
	return o
}

func (o VpnConnectionTypeArrayOutput) ToVpnConnectionTypeArrayOutputWithContext(ctx context.Context) VpnConnectionTypeArrayOutput {
	return o
}

func (o VpnConnectionTypeArrayOutput) Index(i pulumi.IntInput) VpnConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnConnectionType {
		return vs[0].([]VpnConnectionType)[vs[1].(int)]
	}).(VpnConnectionTypeOutput)
}

type VpnConnectionResponse struct {
	ConnectionBandwidth            *int                            `pulumi:"connectionBandwidth"`
	ConnectionStatus               string                          `pulumi:"connectionStatus"`
	DpdTimeoutSeconds              *int                            `pulumi:"dpdTimeoutSeconds"`
	EgressBytesTransferred         float64                         `pulumi:"egressBytesTransferred"`
	EnableBgp                      *bool                           `pulumi:"enableBgp"`
	EnableInternetSecurity         *bool                           `pulumi:"enableInternetSecurity"`
	EnableRateLimiting             *bool                           `pulumi:"enableRateLimiting"`
	Etag                           string                          `pulumi:"etag"`
	Id                             *string                         `pulumi:"id"`
	IngressBytesTransferred        float64                         `pulumi:"ingressBytesTransferred"`
	IpsecPolicies                  []IpsecPolicyResponse           `pulumi:"ipsecPolicies"`
	Name                           *string                         `pulumi:"name"`
	ProvisioningState              string                          `pulumi:"provisioningState"`
	RemoteVpnSite                  *SubResourceResponse            `pulumi:"remoteVpnSite"`
	RoutingConfiguration           *RoutingConfigurationResponse   `pulumi:"routingConfiguration"`
	RoutingWeight                  *int                            `pulumi:"routingWeight"`
	SharedKey                      *string                         `pulumi:"sharedKey"`
	TrafficSelectorPolicies        []TrafficSelectorPolicyResponse `pulumi:"trafficSelectorPolicies"`
	UseLocalAzureIpAddress         *bool                           `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                           `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      *string                         `pulumi:"vpnConnectionProtocolType"`
	VpnLinkConnections             []VpnSiteLinkConnectionResponse `pulumi:"vpnLinkConnections"`
}

type VpnConnectionResponseOutput struct{ *pulumi.OutputState }

func (VpnConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnConnectionResponse)(nil)).Elem()
}

func (o VpnConnectionResponseOutput) ToVpnConnectionResponseOutput() VpnConnectionResponseOutput {
	return o
}

func (o VpnConnectionResponseOutput) ToVpnConnectionResponseOutputWithContext(ctx context.Context) VpnConnectionResponseOutput {
	return o
}

func (o VpnConnectionResponseOutput) ConnectionBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *int { return v.ConnectionBandwidth }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionResponseOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v VpnConnectionResponse) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o VpnConnectionResponseOutput) DpdTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *int { return v.DpdTimeoutSeconds }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionResponseOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnConnectionResponse) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnConnectionResponseOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionResponseOutput) EnableInternetSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *bool { return v.EnableInternetSecurity }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionResponseOutput) EnableRateLimiting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *bool { return v.EnableRateLimiting }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnConnectionResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnConnectionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionResponseOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnConnectionResponse) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnConnectionResponseOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v VpnConnectionResponse) []IpsecPolicyResponse { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o VpnConnectionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnConnectionResponseOutput) RemoteVpnSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *SubResourceResponse { return v.RemoteVpnSite }).(SubResourceResponsePtrOutput)
}

func (o VpnConnectionResponseOutput) RoutingConfiguration() RoutingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *RoutingConfigurationResponse { return v.RoutingConfiguration }).(RoutingConfigurationResponsePtrOutput)
}

func (o VpnConnectionResponseOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionResponseOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionResponseOutput) TrafficSelectorPolicies() TrafficSelectorPolicyResponseArrayOutput {
	return o.ApplyT(func(v VpnConnectionResponse) []TrafficSelectorPolicyResponse { return v.TrafficSelectorPolicies }).(TrafficSelectorPolicyResponseArrayOutput)
}

func (o VpnConnectionResponseOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionResponseOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionResponseOutput) VpnConnectionProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *string { return v.VpnConnectionProtocolType }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionResponseOutput) VpnLinkConnections() VpnSiteLinkConnectionResponseArrayOutput {
	return o.ApplyT(func(v VpnConnectionResponse) []VpnSiteLinkConnectionResponse { return v.VpnLinkConnections }).(VpnSiteLinkConnectionResponseArrayOutput)
}

type VpnConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnConnectionResponse)(nil)).Elem()
}

func (o VpnConnectionResponseArrayOutput) ToVpnConnectionResponseArrayOutput() VpnConnectionResponseArrayOutput {
	return o
}

func (o VpnConnectionResponseArrayOutput) ToVpnConnectionResponseArrayOutputWithContext(ctx context.Context) VpnConnectionResponseArrayOutput {
	return o
}

func (o VpnConnectionResponseArrayOutput) Index(i pulumi.IntInput) VpnConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnConnectionResponse {
		return vs[0].([]VpnConnectionResponse)[vs[1].(int)]
	}).(VpnConnectionResponseOutput)
}

type VpnGatewayIpConfigurationResponse struct {
	Id               *string `pulumi:"id"`
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	PublicIpAddress  *string `pulumi:"publicIpAddress"`
}

type VpnGatewayIpConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VpnGatewayIpConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayIpConfigurationResponse)(nil)).Elem()
}

func (o VpnGatewayIpConfigurationResponseOutput) ToVpnGatewayIpConfigurationResponseOutput() VpnGatewayIpConfigurationResponseOutput {
	return o
}

func (o VpnGatewayIpConfigurationResponseOutput) ToVpnGatewayIpConfigurationResponseOutputWithContext(ctx context.Context) VpnGatewayIpConfigurationResponseOutput {
	return o
}

func (o VpnGatewayIpConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayIpConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayIpConfigurationResponseOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayIpConfigurationResponse) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayIpConfigurationResponseOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayIpConfigurationResponse) *string { return v.PublicIpAddress }).(pulumi.StringPtrOutput)
}

type VpnGatewayIpConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnGatewayIpConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnGatewayIpConfigurationResponse)(nil)).Elem()
}

func (o VpnGatewayIpConfigurationResponseArrayOutput) ToVpnGatewayIpConfigurationResponseArrayOutput() VpnGatewayIpConfigurationResponseArrayOutput {
	return o
}

func (o VpnGatewayIpConfigurationResponseArrayOutput) ToVpnGatewayIpConfigurationResponseArrayOutputWithContext(ctx context.Context) VpnGatewayIpConfigurationResponseArrayOutput {
	return o
}

func (o VpnGatewayIpConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VpnGatewayIpConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnGatewayIpConfigurationResponse {
		return vs[0].([]VpnGatewayIpConfigurationResponse)[vs[1].(int)]
	}).(VpnGatewayIpConfigurationResponseOutput)
}

type VpnGatewayNatRule struct {
	ExternalMappings  []VpnNatRuleMapping `pulumi:"externalMappings"`
	Id                *string             `pulumi:"id"`
	InternalMappings  []VpnNatRuleMapping `pulumi:"internalMappings"`
	IpConfigurationId *string             `pulumi:"ipConfigurationId"`
	Mode              *string             `pulumi:"mode"`
	Name              *string             `pulumi:"name"`
	Type              *string             `pulumi:"type"`
}





type VpnGatewayNatRuleInput interface {
	pulumi.Input

	ToVpnGatewayNatRuleOutput() VpnGatewayNatRuleOutput
	ToVpnGatewayNatRuleOutputWithContext(context.Context) VpnGatewayNatRuleOutput
}

type VpnGatewayNatRuleArgs struct {
	ExternalMappings  VpnNatRuleMappingArrayInput `pulumi:"externalMappings"`
	Id                pulumi.StringPtrInput       `pulumi:"id"`
	InternalMappings  VpnNatRuleMappingArrayInput `pulumi:"internalMappings"`
	IpConfigurationId pulumi.StringPtrInput       `pulumi:"ipConfigurationId"`
	Mode              pulumi.StringPtrInput       `pulumi:"mode"`
	Name              pulumi.StringPtrInput       `pulumi:"name"`
	Type              pulumi.StringPtrInput       `pulumi:"type"`
}

func (VpnGatewayNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayNatRule)(nil)).Elem()
}

func (i VpnGatewayNatRuleArgs) ToVpnGatewayNatRuleOutput() VpnGatewayNatRuleOutput {
	return i.ToVpnGatewayNatRuleOutputWithContext(context.Background())
}

func (i VpnGatewayNatRuleArgs) ToVpnGatewayNatRuleOutputWithContext(ctx context.Context) VpnGatewayNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayNatRuleOutput)
}





type VpnGatewayNatRuleArrayInput interface {
	pulumi.Input

	ToVpnGatewayNatRuleArrayOutput() VpnGatewayNatRuleArrayOutput
	ToVpnGatewayNatRuleArrayOutputWithContext(context.Context) VpnGatewayNatRuleArrayOutput
}

type VpnGatewayNatRuleArray []VpnGatewayNatRuleInput

func (VpnGatewayNatRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnGatewayNatRule)(nil)).Elem()
}

func (i VpnGatewayNatRuleArray) ToVpnGatewayNatRuleArrayOutput() VpnGatewayNatRuleArrayOutput {
	return i.ToVpnGatewayNatRuleArrayOutputWithContext(context.Background())
}

func (i VpnGatewayNatRuleArray) ToVpnGatewayNatRuleArrayOutputWithContext(ctx context.Context) VpnGatewayNatRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayNatRuleArrayOutput)
}

type VpnGatewayNatRuleOutput struct{ *pulumi.OutputState }

func (VpnGatewayNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayNatRule)(nil)).Elem()
}

func (o VpnGatewayNatRuleOutput) ToVpnGatewayNatRuleOutput() VpnGatewayNatRuleOutput {
	return o
}

func (o VpnGatewayNatRuleOutput) ToVpnGatewayNatRuleOutputWithContext(ctx context.Context) VpnGatewayNatRuleOutput {
	return o
}

func (o VpnGatewayNatRuleOutput) ExternalMappings() VpnNatRuleMappingArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) []VpnNatRuleMapping { return v.ExternalMappings }).(VpnNatRuleMappingArrayOutput)
}

func (o VpnGatewayNatRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleOutput) InternalMappings() VpnNatRuleMappingArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) []VpnNatRuleMapping { return v.InternalMappings }).(VpnNatRuleMappingArrayOutput)
}

func (o VpnGatewayNatRuleOutput) IpConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) *string { return v.IpConfigurationId }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VpnGatewayNatRuleArrayOutput struct{ *pulumi.OutputState }

func (VpnGatewayNatRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnGatewayNatRule)(nil)).Elem()
}

func (o VpnGatewayNatRuleArrayOutput) ToVpnGatewayNatRuleArrayOutput() VpnGatewayNatRuleArrayOutput {
	return o
}

func (o VpnGatewayNatRuleArrayOutput) ToVpnGatewayNatRuleArrayOutputWithContext(ctx context.Context) VpnGatewayNatRuleArrayOutput {
	return o
}

func (o VpnGatewayNatRuleArrayOutput) Index(i pulumi.IntInput) VpnGatewayNatRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnGatewayNatRule {
		return vs[0].([]VpnGatewayNatRule)[vs[1].(int)]
	}).(VpnGatewayNatRuleOutput)
}

type VpnGatewayNatRuleResponse struct {
	EgressVpnSiteLinkConnections  []SubResourceResponse       `pulumi:"egressVpnSiteLinkConnections"`
	Etag                          string                      `pulumi:"etag"`
	ExternalMappings              []VpnNatRuleMappingResponse `pulumi:"externalMappings"`
	Id                            *string                     `pulumi:"id"`
	IngressVpnSiteLinkConnections []SubResourceResponse       `pulumi:"ingressVpnSiteLinkConnections"`
	InternalMappings              []VpnNatRuleMappingResponse `pulumi:"internalMappings"`
	IpConfigurationId             *string                     `pulumi:"ipConfigurationId"`
	Mode                          *string                     `pulumi:"mode"`
	Name                          *string                     `pulumi:"name"`
	ProvisioningState             string                      `pulumi:"provisioningState"`
	Type                          string                      `pulumi:"type"`
}

type VpnGatewayNatRuleResponseOutput struct{ *pulumi.OutputState }

func (VpnGatewayNatRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayNatRuleResponse)(nil)).Elem()
}

func (o VpnGatewayNatRuleResponseOutput) ToVpnGatewayNatRuleResponseOutput() VpnGatewayNatRuleResponseOutput {
	return o
}

func (o VpnGatewayNatRuleResponseOutput) ToVpnGatewayNatRuleResponseOutputWithContext(ctx context.Context) VpnGatewayNatRuleResponseOutput {
	return o
}

func (o VpnGatewayNatRuleResponseOutput) EgressVpnSiteLinkConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) []SubResourceResponse { return v.EgressVpnSiteLinkConnections }).(SubResourceResponseArrayOutput)
}

func (o VpnGatewayNatRuleResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnGatewayNatRuleResponseOutput) ExternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) []VpnNatRuleMappingResponse { return v.ExternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

func (o VpnGatewayNatRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleResponseOutput) IngressVpnSiteLinkConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) []SubResourceResponse { return v.IngressVpnSiteLinkConnections }).(SubResourceResponseArrayOutput)
}

func (o VpnGatewayNatRuleResponseOutput) InternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) []VpnNatRuleMappingResponse { return v.InternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

func (o VpnGatewayNatRuleResponseOutput) IpConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) *string { return v.IpConfigurationId }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnGatewayNatRuleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VpnGatewayNatRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnGatewayNatRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnGatewayNatRuleResponse)(nil)).Elem()
}

func (o VpnGatewayNatRuleResponseArrayOutput) ToVpnGatewayNatRuleResponseArrayOutput() VpnGatewayNatRuleResponseArrayOutput {
	return o
}

func (o VpnGatewayNatRuleResponseArrayOutput) ToVpnGatewayNatRuleResponseArrayOutputWithContext(ctx context.Context) VpnGatewayNatRuleResponseArrayOutput {
	return o
}

func (o VpnGatewayNatRuleResponseArrayOutput) Index(i pulumi.IntInput) VpnGatewayNatRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnGatewayNatRuleResponse {
		return vs[0].([]VpnGatewayNatRuleResponse)[vs[1].(int)]
	}).(VpnGatewayNatRuleResponseOutput)
}

type VpnLinkBgpSettings struct {
	Asn               *float64 `pulumi:"asn"`
	BgpPeeringAddress *string  `pulumi:"bgpPeeringAddress"`
}





type VpnLinkBgpSettingsInput interface {
	pulumi.Input

	ToVpnLinkBgpSettingsOutput() VpnLinkBgpSettingsOutput
	ToVpnLinkBgpSettingsOutputWithContext(context.Context) VpnLinkBgpSettingsOutput
}

type VpnLinkBgpSettingsArgs struct {
	Asn               pulumi.Float64PtrInput `pulumi:"asn"`
	BgpPeeringAddress pulumi.StringPtrInput  `pulumi:"bgpPeeringAddress"`
}

func (VpnLinkBgpSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkBgpSettings)(nil)).Elem()
}

func (i VpnLinkBgpSettingsArgs) ToVpnLinkBgpSettingsOutput() VpnLinkBgpSettingsOutput {
	return i.ToVpnLinkBgpSettingsOutputWithContext(context.Background())
}

func (i VpnLinkBgpSettingsArgs) ToVpnLinkBgpSettingsOutputWithContext(ctx context.Context) VpnLinkBgpSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkBgpSettingsOutput)
}

func (i VpnLinkBgpSettingsArgs) ToVpnLinkBgpSettingsPtrOutput() VpnLinkBgpSettingsPtrOutput {
	return i.ToVpnLinkBgpSettingsPtrOutputWithContext(context.Background())
}

func (i VpnLinkBgpSettingsArgs) ToVpnLinkBgpSettingsPtrOutputWithContext(ctx context.Context) VpnLinkBgpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkBgpSettingsOutput).ToVpnLinkBgpSettingsPtrOutputWithContext(ctx)
}









type VpnLinkBgpSettingsPtrInput interface {
	pulumi.Input

	ToVpnLinkBgpSettingsPtrOutput() VpnLinkBgpSettingsPtrOutput
	ToVpnLinkBgpSettingsPtrOutputWithContext(context.Context) VpnLinkBgpSettingsPtrOutput
}

type vpnLinkBgpSettingsPtrType VpnLinkBgpSettingsArgs

func VpnLinkBgpSettingsPtr(v *VpnLinkBgpSettingsArgs) VpnLinkBgpSettingsPtrInput {
	return (*vpnLinkBgpSettingsPtrType)(v)
}

func (*vpnLinkBgpSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkBgpSettings)(nil)).Elem()
}

func (i *vpnLinkBgpSettingsPtrType) ToVpnLinkBgpSettingsPtrOutput() VpnLinkBgpSettingsPtrOutput {
	return i.ToVpnLinkBgpSettingsPtrOutputWithContext(context.Background())
}

func (i *vpnLinkBgpSettingsPtrType) ToVpnLinkBgpSettingsPtrOutputWithContext(ctx context.Context) VpnLinkBgpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkBgpSettingsPtrOutput)
}

type VpnLinkBgpSettingsOutput struct{ *pulumi.OutputState }

func (VpnLinkBgpSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkBgpSettings)(nil)).Elem()
}

func (o VpnLinkBgpSettingsOutput) ToVpnLinkBgpSettingsOutput() VpnLinkBgpSettingsOutput {
	return o
}

func (o VpnLinkBgpSettingsOutput) ToVpnLinkBgpSettingsOutputWithContext(ctx context.Context) VpnLinkBgpSettingsOutput {
	return o
}

func (o VpnLinkBgpSettingsOutput) ToVpnLinkBgpSettingsPtrOutput() VpnLinkBgpSettingsPtrOutput {
	return o.ToVpnLinkBgpSettingsPtrOutputWithContext(context.Background())
}

func (o VpnLinkBgpSettingsOutput) ToVpnLinkBgpSettingsPtrOutputWithContext(ctx context.Context) VpnLinkBgpSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnLinkBgpSettings) *VpnLinkBgpSettings {
		return &v
	}).(VpnLinkBgpSettingsPtrOutput)
}

func (o VpnLinkBgpSettingsOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VpnLinkBgpSettings) *float64 { return v.Asn }).(pulumi.Float64PtrOutput)
}

func (o VpnLinkBgpSettingsOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnLinkBgpSettings) *string { return v.BgpPeeringAddress }).(pulumi.StringPtrOutput)
}

type VpnLinkBgpSettingsPtrOutput struct{ *pulumi.OutputState }

func (VpnLinkBgpSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkBgpSettings)(nil)).Elem()
}

func (o VpnLinkBgpSettingsPtrOutput) ToVpnLinkBgpSettingsPtrOutput() VpnLinkBgpSettingsPtrOutput {
	return o
}

func (o VpnLinkBgpSettingsPtrOutput) ToVpnLinkBgpSettingsPtrOutputWithContext(ctx context.Context) VpnLinkBgpSettingsPtrOutput {
	return o
}

func (o VpnLinkBgpSettingsPtrOutput) Elem() VpnLinkBgpSettingsOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettings) VpnLinkBgpSettings {
		if v != nil {
			return *v
		}
		var ret VpnLinkBgpSettings
		return ret
	}).(VpnLinkBgpSettingsOutput)
}

func (o VpnLinkBgpSettingsPtrOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.Asn
	}).(pulumi.Float64PtrOutput)
}

func (o VpnLinkBgpSettingsPtrOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettings) *string {
		if v == nil {
			return nil
		}
		return v.BgpPeeringAddress
	}).(pulumi.StringPtrOutput)
}

type VpnLinkBgpSettingsResponse struct {
	Asn               *float64 `pulumi:"asn"`
	BgpPeeringAddress *string  `pulumi:"bgpPeeringAddress"`
}

type VpnLinkBgpSettingsResponseOutput struct{ *pulumi.OutputState }

func (VpnLinkBgpSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkBgpSettingsResponse)(nil)).Elem()
}

func (o VpnLinkBgpSettingsResponseOutput) ToVpnLinkBgpSettingsResponseOutput() VpnLinkBgpSettingsResponseOutput {
	return o
}

func (o VpnLinkBgpSettingsResponseOutput) ToVpnLinkBgpSettingsResponseOutputWithContext(ctx context.Context) VpnLinkBgpSettingsResponseOutput {
	return o
}

func (o VpnLinkBgpSettingsResponseOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VpnLinkBgpSettingsResponse) *float64 { return v.Asn }).(pulumi.Float64PtrOutput)
}

func (o VpnLinkBgpSettingsResponseOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnLinkBgpSettingsResponse) *string { return v.BgpPeeringAddress }).(pulumi.StringPtrOutput)
}

type VpnLinkBgpSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (VpnLinkBgpSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkBgpSettingsResponse)(nil)).Elem()
}

func (o VpnLinkBgpSettingsResponsePtrOutput) ToVpnLinkBgpSettingsResponsePtrOutput() VpnLinkBgpSettingsResponsePtrOutput {
	return o
}

func (o VpnLinkBgpSettingsResponsePtrOutput) ToVpnLinkBgpSettingsResponsePtrOutputWithContext(ctx context.Context) VpnLinkBgpSettingsResponsePtrOutput {
	return o
}

func (o VpnLinkBgpSettingsResponsePtrOutput) Elem() VpnLinkBgpSettingsResponseOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettingsResponse) VpnLinkBgpSettingsResponse {
		if v != nil {
			return *v
		}
		var ret VpnLinkBgpSettingsResponse
		return ret
	}).(VpnLinkBgpSettingsResponseOutput)
}

func (o VpnLinkBgpSettingsResponsePtrOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Asn
	}).(pulumi.Float64PtrOutput)
}

func (o VpnLinkBgpSettingsResponsePtrOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.BgpPeeringAddress
	}).(pulumi.StringPtrOutput)
}

type VpnLinkProviderProperties struct {
	LinkProviderName *string `pulumi:"linkProviderName"`
	LinkSpeedInMbps  *int    `pulumi:"linkSpeedInMbps"`
}





type VpnLinkProviderPropertiesInput interface {
	pulumi.Input

	ToVpnLinkProviderPropertiesOutput() VpnLinkProviderPropertiesOutput
	ToVpnLinkProviderPropertiesOutputWithContext(context.Context) VpnLinkProviderPropertiesOutput
}

type VpnLinkProviderPropertiesArgs struct {
	LinkProviderName pulumi.StringPtrInput `pulumi:"linkProviderName"`
	LinkSpeedInMbps  pulumi.IntPtrInput    `pulumi:"linkSpeedInMbps"`
}

func (VpnLinkProviderPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkProviderProperties)(nil)).Elem()
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesOutput() VpnLinkProviderPropertiesOutput {
	return i.ToVpnLinkProviderPropertiesOutputWithContext(context.Background())
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkProviderPropertiesOutput)
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return i.ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Background())
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkProviderPropertiesOutput).ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx)
}









type VpnLinkProviderPropertiesPtrInput interface {
	pulumi.Input

	ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput
	ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Context) VpnLinkProviderPropertiesPtrOutput
}

type vpnLinkProviderPropertiesPtrType VpnLinkProviderPropertiesArgs

func VpnLinkProviderPropertiesPtr(v *VpnLinkProviderPropertiesArgs) VpnLinkProviderPropertiesPtrInput {
	return (*vpnLinkProviderPropertiesPtrType)(v)
}

func (*vpnLinkProviderPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkProviderProperties)(nil)).Elem()
}

func (i *vpnLinkProviderPropertiesPtrType) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return i.ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Background())
}

func (i *vpnLinkProviderPropertiesPtrType) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkProviderPropertiesPtrOutput)
}

type VpnLinkProviderPropertiesOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkProviderProperties)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesOutput() VpnLinkProviderPropertiesOutput {
	return o
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesOutput {
	return o
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return o.ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Background())
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnLinkProviderProperties) *VpnLinkProviderProperties {
		return &v
	}).(VpnLinkProviderPropertiesPtrOutput)
}

func (o VpnLinkProviderPropertiesOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderProperties) *string { return v.LinkProviderName }).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderProperties) *int { return v.LinkSpeedInMbps }).(pulumi.IntPtrOutput)
}

type VpnLinkProviderPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkProviderProperties)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesPtrOutput) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesPtrOutput) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesPtrOutput) Elem() VpnLinkProviderPropertiesOutput {
	return o.ApplyT(func(v *VpnLinkProviderProperties) VpnLinkProviderProperties {
		if v != nil {
			return *v
		}
		var ret VpnLinkProviderProperties
		return ret
	}).(VpnLinkProviderPropertiesOutput)
}

func (o VpnLinkProviderPropertiesPtrOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderProperties) *string {
		if v == nil {
			return nil
		}
		return v.LinkProviderName
	}).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesPtrOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderProperties) *int {
		if v == nil {
			return nil
		}
		return v.LinkSpeedInMbps
	}).(pulumi.IntPtrOutput)
}

type VpnLinkProviderPropertiesResponse struct {
	LinkProviderName *string `pulumi:"linkProviderName"`
	LinkSpeedInMbps  *int    `pulumi:"linkSpeedInMbps"`
}

type VpnLinkProviderPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkProviderPropertiesResponse)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesResponseOutput) ToVpnLinkProviderPropertiesResponseOutput() VpnLinkProviderPropertiesResponseOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponseOutput) ToVpnLinkProviderPropertiesResponseOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesResponseOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponseOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderPropertiesResponse) *string { return v.LinkProviderName }).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesResponseOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderPropertiesResponse) *int { return v.LinkSpeedInMbps }).(pulumi.IntPtrOutput)
}

type VpnLinkProviderPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkProviderPropertiesResponse)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) ToVpnLinkProviderPropertiesResponsePtrOutput() VpnLinkProviderPropertiesResponsePtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) ToVpnLinkProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesResponsePtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) Elem() VpnLinkProviderPropertiesResponseOutput {
	return o.ApplyT(func(v *VpnLinkProviderPropertiesResponse) VpnLinkProviderPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VpnLinkProviderPropertiesResponse
		return ret
	}).(VpnLinkProviderPropertiesResponseOutput)
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LinkProviderName
	}).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.LinkSpeedInMbps
	}).(pulumi.IntPtrOutput)
}

type VpnNatRuleMapping struct {
	AddressSpace *string `pulumi:"addressSpace"`
}





type VpnNatRuleMappingInput interface {
	pulumi.Input

	ToVpnNatRuleMappingOutput() VpnNatRuleMappingOutput
	ToVpnNatRuleMappingOutputWithContext(context.Context) VpnNatRuleMappingOutput
}

type VpnNatRuleMappingArgs struct {
	AddressSpace pulumi.StringPtrInput `pulumi:"addressSpace"`
}

func (VpnNatRuleMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleMapping)(nil)).Elem()
}

func (i VpnNatRuleMappingArgs) ToVpnNatRuleMappingOutput() VpnNatRuleMappingOutput {
	return i.ToVpnNatRuleMappingOutputWithContext(context.Background())
}

func (i VpnNatRuleMappingArgs) ToVpnNatRuleMappingOutputWithContext(ctx context.Context) VpnNatRuleMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnNatRuleMappingOutput)
}





type VpnNatRuleMappingArrayInput interface {
	pulumi.Input

	ToVpnNatRuleMappingArrayOutput() VpnNatRuleMappingArrayOutput
	ToVpnNatRuleMappingArrayOutputWithContext(context.Context) VpnNatRuleMappingArrayOutput
}

type VpnNatRuleMappingArray []VpnNatRuleMappingInput

func (VpnNatRuleMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnNatRuleMapping)(nil)).Elem()
}

func (i VpnNatRuleMappingArray) ToVpnNatRuleMappingArrayOutput() VpnNatRuleMappingArrayOutput {
	return i.ToVpnNatRuleMappingArrayOutputWithContext(context.Background())
}

func (i VpnNatRuleMappingArray) ToVpnNatRuleMappingArrayOutputWithContext(ctx context.Context) VpnNatRuleMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnNatRuleMappingArrayOutput)
}

type VpnNatRuleMappingOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleMapping)(nil)).Elem()
}

func (o VpnNatRuleMappingOutput) ToVpnNatRuleMappingOutput() VpnNatRuleMappingOutput {
	return o
}

func (o VpnNatRuleMappingOutput) ToVpnNatRuleMappingOutputWithContext(ctx context.Context) VpnNatRuleMappingOutput {
	return o
}

func (o VpnNatRuleMappingOutput) AddressSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnNatRuleMapping) *string { return v.AddressSpace }).(pulumi.StringPtrOutput)
}

type VpnNatRuleMappingArrayOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnNatRuleMapping)(nil)).Elem()
}

func (o VpnNatRuleMappingArrayOutput) ToVpnNatRuleMappingArrayOutput() VpnNatRuleMappingArrayOutput {
	return o
}

func (o VpnNatRuleMappingArrayOutput) ToVpnNatRuleMappingArrayOutputWithContext(ctx context.Context) VpnNatRuleMappingArrayOutput {
	return o
}

func (o VpnNatRuleMappingArrayOutput) Index(i pulumi.IntInput) VpnNatRuleMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnNatRuleMapping {
		return vs[0].([]VpnNatRuleMapping)[vs[1].(int)]
	}).(VpnNatRuleMappingOutput)
}

type VpnNatRuleMappingResponse struct {
	AddressSpace *string `pulumi:"addressSpace"`
}

type VpnNatRuleMappingResponseOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleMappingResponse)(nil)).Elem()
}

func (o VpnNatRuleMappingResponseOutput) ToVpnNatRuleMappingResponseOutput() VpnNatRuleMappingResponseOutput {
	return o
}

func (o VpnNatRuleMappingResponseOutput) ToVpnNatRuleMappingResponseOutputWithContext(ctx context.Context) VpnNatRuleMappingResponseOutput {
	return o
}

func (o VpnNatRuleMappingResponseOutput) AddressSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnNatRuleMappingResponse) *string { return v.AddressSpace }).(pulumi.StringPtrOutput)
}

type VpnNatRuleMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnNatRuleMappingResponse)(nil)).Elem()
}

func (o VpnNatRuleMappingResponseArrayOutput) ToVpnNatRuleMappingResponseArrayOutput() VpnNatRuleMappingResponseArrayOutput {
	return o
}

func (o VpnNatRuleMappingResponseArrayOutput) ToVpnNatRuleMappingResponseArrayOutputWithContext(ctx context.Context) VpnNatRuleMappingResponseArrayOutput {
	return o
}

func (o VpnNatRuleMappingResponseArrayOutput) Index(i pulumi.IntInput) VpnNatRuleMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnNatRuleMappingResponse {
		return vs[0].([]VpnNatRuleMappingResponse)[vs[1].(int)]
	}).(VpnNatRuleMappingResponseOutput)
}

type VpnServerConfigRadiusClientRootCertificate struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type VpnServerConfigRadiusClientRootCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusClientRootCertificateOutput() VpnServerConfigRadiusClientRootCertificateOutput
	ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(context.Context) VpnServerConfigRadiusClientRootCertificateOutput
}

type VpnServerConfigRadiusClientRootCertificateArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (VpnServerConfigRadiusClientRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusClientRootCertificateArgs) ToVpnServerConfigRadiusClientRootCertificateOutput() VpnServerConfigRadiusClientRootCertificateOutput {
	return i.ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusClientRootCertificateArgs) ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusClientRootCertificateOutput)
}





type VpnServerConfigRadiusClientRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusClientRootCertificateArrayOutput() VpnServerConfigRadiusClientRootCertificateArrayOutput
	ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(context.Context) VpnServerConfigRadiusClientRootCertificateArrayOutput
}

type VpnServerConfigRadiusClientRootCertificateArray []VpnServerConfigRadiusClientRootCertificateInput

func (VpnServerConfigRadiusClientRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusClientRootCertificateArray) ToVpnServerConfigRadiusClientRootCertificateArrayOutput() VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return i.ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusClientRootCertificateArray) ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusClientRootCertificateArrayOutput)
}

type VpnServerConfigRadiusClientRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) ToVpnServerConfigRadiusClientRootCertificateOutput() VpnServerConfigRadiusClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusClientRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateArrayOutput) ToVpnServerConfigRadiusClientRootCertificateArrayOutput() VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateArrayOutput) ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusClientRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusClientRootCertificate {
		return vs[0].([]VpnServerConfigRadiusClientRootCertificate)[vs[1].(int)]
	}).(VpnServerConfigRadiusClientRootCertificateOutput)
}

type VpnServerConfigRadiusClientRootCertificateResponse struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}

type VpnServerConfigRadiusClientRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) ToVpnServerConfigRadiusClientRootCertificateResponseOutput() VpnServerConfigRadiusClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) ToVpnServerConfigRadiusClientRootCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusClientRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusClientRootCertificateResponseArrayOutput() VpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusClientRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusClientRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusClientRootCertificateResponse {
		return vs[0].([]VpnServerConfigRadiusClientRootCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigRadiusClientRootCertificateResponseOutput)
}

type VpnServerConfigRadiusServerRootCertificate struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}





type VpnServerConfigRadiusServerRootCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusServerRootCertificateOutput() VpnServerConfigRadiusServerRootCertificateOutput
	ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(context.Context) VpnServerConfigRadiusServerRootCertificateOutput
}

type VpnServerConfigRadiusServerRootCertificateArgs struct {
	Name           pulumi.StringPtrInput `pulumi:"name"`
	PublicCertData pulumi.StringPtrInput `pulumi:"publicCertData"`
}

func (VpnServerConfigRadiusServerRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusServerRootCertificateArgs) ToVpnServerConfigRadiusServerRootCertificateOutput() VpnServerConfigRadiusServerRootCertificateOutput {
	return i.ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusServerRootCertificateArgs) ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusServerRootCertificateOutput)
}





type VpnServerConfigRadiusServerRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusServerRootCertificateArrayOutput() VpnServerConfigRadiusServerRootCertificateArrayOutput
	ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(context.Context) VpnServerConfigRadiusServerRootCertificateArrayOutput
}

type VpnServerConfigRadiusServerRootCertificateArray []VpnServerConfigRadiusServerRootCertificateInput

func (VpnServerConfigRadiusServerRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusServerRootCertificateArray) ToVpnServerConfigRadiusServerRootCertificateArrayOutput() VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return i.ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusServerRootCertificateArray) ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusServerRootCertificateArrayOutput)
}

type VpnServerConfigRadiusServerRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) ToVpnServerConfigRadiusServerRootCertificateOutput() VpnServerConfigRadiusServerRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificate) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusServerRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateArrayOutput) ToVpnServerConfigRadiusServerRootCertificateArrayOutput() VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateArrayOutput) ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusServerRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusServerRootCertificate {
		return vs[0].([]VpnServerConfigRadiusServerRootCertificate)[vs[1].(int)]
	}).(VpnServerConfigRadiusServerRootCertificateOutput)
}

type VpnServerConfigRadiusServerRootCertificateResponse struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}

type VpnServerConfigRadiusServerRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusServerRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) ToVpnServerConfigRadiusServerRootCertificateResponseOutput() VpnServerConfigRadiusServerRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) ToVpnServerConfigRadiusServerRootCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificateResponse) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusServerRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusServerRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusServerRootCertificateResponseArrayOutput() VpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusServerRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusServerRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusServerRootCertificateResponse {
		return vs[0].([]VpnServerConfigRadiusServerRootCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigRadiusServerRootCertificateResponseOutput)
}

type VpnServerConfigVpnClientRevokedCertificate struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type VpnServerConfigVpnClientRevokedCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRevokedCertificateOutput() VpnServerConfigVpnClientRevokedCertificateOutput
	ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(context.Context) VpnServerConfigVpnClientRevokedCertificateOutput
}

type VpnServerConfigVpnClientRevokedCertificateArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (VpnServerConfigVpnClientRevokedCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRevokedCertificateArgs) ToVpnServerConfigVpnClientRevokedCertificateOutput() VpnServerConfigVpnClientRevokedCertificateOutput {
	return i.ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRevokedCertificateArgs) ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRevokedCertificateOutput)
}





type VpnServerConfigVpnClientRevokedCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRevokedCertificateArrayOutput() VpnServerConfigVpnClientRevokedCertificateArrayOutput
	ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(context.Context) VpnServerConfigVpnClientRevokedCertificateArrayOutput
}

type VpnServerConfigVpnClientRevokedCertificateArray []VpnServerConfigVpnClientRevokedCertificateInput

func (VpnServerConfigVpnClientRevokedCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRevokedCertificateArray) ToVpnServerConfigVpnClientRevokedCertificateArrayOutput() VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return i.ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRevokedCertificateArray) ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRevokedCertificateArrayOutput)
}

type VpnServerConfigVpnClientRevokedCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) ToVpnServerConfigVpnClientRevokedCertificateOutput() VpnServerConfigVpnClientRevokedCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRevokedCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateArrayOutput() VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRevokedCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRevokedCertificate {
		return vs[0].([]VpnServerConfigVpnClientRevokedCertificate)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRevokedCertificateOutput)
}

type VpnServerConfigVpnClientRevokedCertificateResponse struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}

type VpnServerConfigVpnClientRevokedCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseOutput() VpnServerConfigVpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput() VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRevokedCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRevokedCertificateResponse {
		return vs[0].([]VpnServerConfigVpnClientRevokedCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRevokedCertificateResponseOutput)
}

type VpnServerConfigVpnClientRootCertificate struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}





type VpnServerConfigVpnClientRootCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRootCertificateOutput() VpnServerConfigVpnClientRootCertificateOutput
	ToVpnServerConfigVpnClientRootCertificateOutputWithContext(context.Context) VpnServerConfigVpnClientRootCertificateOutput
}

type VpnServerConfigVpnClientRootCertificateArgs struct {
	Name           pulumi.StringPtrInput `pulumi:"name"`
	PublicCertData pulumi.StringPtrInput `pulumi:"publicCertData"`
}

func (VpnServerConfigVpnClientRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRootCertificateArgs) ToVpnServerConfigVpnClientRootCertificateOutput() VpnServerConfigVpnClientRootCertificateOutput {
	return i.ToVpnServerConfigVpnClientRootCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRootCertificateArgs) ToVpnServerConfigVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRootCertificateOutput)
}





type VpnServerConfigVpnClientRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRootCertificateArrayOutput() VpnServerConfigVpnClientRootCertificateArrayOutput
	ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(context.Context) VpnServerConfigVpnClientRootCertificateArrayOutput
}

type VpnServerConfigVpnClientRootCertificateArray []VpnServerConfigVpnClientRootCertificateInput

func (VpnServerConfigVpnClientRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRootCertificateArray) ToVpnServerConfigVpnClientRootCertificateArrayOutput() VpnServerConfigVpnClientRootCertificateArrayOutput {
	return i.ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRootCertificateArray) ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRootCertificateArrayOutput)
}

type VpnServerConfigVpnClientRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateOutput) ToVpnServerConfigVpnClientRootCertificateOutput() VpnServerConfigVpnClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateOutput) ToVpnServerConfigVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRootCertificateOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificate) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateArrayOutput) ToVpnServerConfigVpnClientRootCertificateArrayOutput() VpnServerConfigVpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateArrayOutput) ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRootCertificate {
		return vs[0].([]VpnServerConfigVpnClientRootCertificate)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRootCertificateOutput)
}

type VpnServerConfigVpnClientRootCertificateResponse struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}

type VpnServerConfigVpnClientRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) ToVpnServerConfigVpnClientRootCertificateResponseOutput() VpnServerConfigVpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) ToVpnServerConfigVpnClientRootCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificateResponse) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRootCertificateResponseArrayOutput() VpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRootCertificateResponse {
		return vs[0].([]VpnServerConfigVpnClientRootCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRootCertificateResponseOutput)
}

type VpnServerConfigurationPolicyGroupMember struct {
	AttributeType  *string `pulumi:"attributeType"`
	AttributeValue *string `pulumi:"attributeValue"`
	Name           *string `pulumi:"name"`
}





type VpnServerConfigurationPolicyGroupMemberInput interface {
	pulumi.Input

	ToVpnServerConfigurationPolicyGroupMemberOutput() VpnServerConfigurationPolicyGroupMemberOutput
	ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(context.Context) VpnServerConfigurationPolicyGroupMemberOutput
}

type VpnServerConfigurationPolicyGroupMemberArgs struct {
	AttributeType  pulumi.StringPtrInput `pulumi:"attributeType"`
	AttributeValue pulumi.StringPtrInput `pulumi:"attributeValue"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
}

func (VpnServerConfigurationPolicyGroupMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (i VpnServerConfigurationPolicyGroupMemberArgs) ToVpnServerConfigurationPolicyGroupMemberOutput() VpnServerConfigurationPolicyGroupMemberOutput {
	return i.ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(context.Background())
}

func (i VpnServerConfigurationPolicyGroupMemberArgs) ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationPolicyGroupMemberOutput)
}





type VpnServerConfigurationPolicyGroupMemberArrayInput interface {
	pulumi.Input

	ToVpnServerConfigurationPolicyGroupMemberArrayOutput() VpnServerConfigurationPolicyGroupMemberArrayOutput
	ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(context.Context) VpnServerConfigurationPolicyGroupMemberArrayOutput
}

type VpnServerConfigurationPolicyGroupMemberArray []VpnServerConfigurationPolicyGroupMemberInput

func (VpnServerConfigurationPolicyGroupMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (i VpnServerConfigurationPolicyGroupMemberArray) ToVpnServerConfigurationPolicyGroupMemberArrayOutput() VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return i.ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigurationPolicyGroupMemberArray) ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationPolicyGroupMemberArrayOutput)
}

type VpnServerConfigurationPolicyGroupMemberOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) ToVpnServerConfigurationPolicyGroupMemberOutput() VpnServerConfigurationPolicyGroupMemberOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) AttributeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMember) *string { return v.AttributeType }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) AttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMember) *string { return v.AttributeValue }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMember) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VpnServerConfigurationPolicyGroupMemberArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberArrayOutput) ToVpnServerConfigurationPolicyGroupMemberArrayOutput() VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberArrayOutput) ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberArrayOutput) Index(i pulumi.IntInput) VpnServerConfigurationPolicyGroupMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigurationPolicyGroupMember {
		return vs[0].([]VpnServerConfigurationPolicyGroupMember)[vs[1].(int)]
	}).(VpnServerConfigurationPolicyGroupMemberOutput)
}

type VpnServerConfigurationPolicyGroupMemberResponse struct {
	AttributeType  *string `pulumi:"attributeType"`
	AttributeValue *string `pulumi:"attributeValue"`
	Name           *string `pulumi:"name"`
}

type VpnServerConfigurationPolicyGroupMemberResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroupMemberResponse)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) ToVpnServerConfigurationPolicyGroupMemberResponseOutput() VpnServerConfigurationPolicyGroupMemberResponseOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) ToVpnServerConfigurationPolicyGroupMemberResponseOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberResponseOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) AttributeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMemberResponse) *string { return v.AttributeType }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) AttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMemberResponse) *string { return v.AttributeValue }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMemberResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VpnServerConfigurationPolicyGroupMemberResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroupMemberResponse)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) ToVpnServerConfigurationPolicyGroupMemberResponseArrayOutput() VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) ToVpnServerConfigurationPolicyGroupMemberResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigurationPolicyGroupMemberResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigurationPolicyGroupMemberResponse {
		return vs[0].([]VpnServerConfigurationPolicyGroupMemberResponse)[vs[1].(int)]
	}).(VpnServerConfigurationPolicyGroupMemberResponseOutput)
}

type VpnSiteLink struct {
	BgpProperties  *VpnLinkBgpSettings        `pulumi:"bgpProperties"`
	Fqdn           *string                    `pulumi:"fqdn"`
	Id             *string                    `pulumi:"id"`
	IpAddress      *string                    `pulumi:"ipAddress"`
	LinkProperties *VpnLinkProviderProperties `pulumi:"linkProperties"`
	Name           *string                    `pulumi:"name"`
}





type VpnSiteLinkInput interface {
	pulumi.Input

	ToVpnSiteLinkOutput() VpnSiteLinkOutput
	ToVpnSiteLinkOutputWithContext(context.Context) VpnSiteLinkOutput
}

type VpnSiteLinkArgs struct {
	BgpProperties  VpnLinkBgpSettingsPtrInput        `pulumi:"bgpProperties"`
	Fqdn           pulumi.StringPtrInput             `pulumi:"fqdn"`
	Id             pulumi.StringPtrInput             `pulumi:"id"`
	IpAddress      pulumi.StringPtrInput             `pulumi:"ipAddress"`
	LinkProperties VpnLinkProviderPropertiesPtrInput `pulumi:"linkProperties"`
	Name           pulumi.StringPtrInput             `pulumi:"name"`
}

func (VpnSiteLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLink)(nil)).Elem()
}

func (i VpnSiteLinkArgs) ToVpnSiteLinkOutput() VpnSiteLinkOutput {
	return i.ToVpnSiteLinkOutputWithContext(context.Background())
}

func (i VpnSiteLinkArgs) ToVpnSiteLinkOutputWithContext(ctx context.Context) VpnSiteLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkOutput)
}





type VpnSiteLinkArrayInput interface {
	pulumi.Input

	ToVpnSiteLinkArrayOutput() VpnSiteLinkArrayOutput
	ToVpnSiteLinkArrayOutputWithContext(context.Context) VpnSiteLinkArrayOutput
}

type VpnSiteLinkArray []VpnSiteLinkInput

func (VpnSiteLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLink)(nil)).Elem()
}

func (i VpnSiteLinkArray) ToVpnSiteLinkArrayOutput() VpnSiteLinkArrayOutput {
	return i.ToVpnSiteLinkArrayOutputWithContext(context.Background())
}

func (i VpnSiteLinkArray) ToVpnSiteLinkArrayOutputWithContext(ctx context.Context) VpnSiteLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkArrayOutput)
}

type VpnSiteLinkOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLink)(nil)).Elem()
}

func (o VpnSiteLinkOutput) ToVpnSiteLinkOutput() VpnSiteLinkOutput {
	return o
}

func (o VpnSiteLinkOutput) ToVpnSiteLinkOutputWithContext(ctx context.Context) VpnSiteLinkOutput {
	return o
}

func (o VpnSiteLinkOutput) BgpProperties() VpnLinkBgpSettingsPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *VpnLinkBgpSettings { return v.BgpProperties }).(VpnLinkBgpSettingsPtrOutput)
}

func (o VpnSiteLinkOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkOutput) LinkProperties() VpnLinkProviderPropertiesPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *VpnLinkProviderProperties { return v.LinkProperties }).(VpnLinkProviderPropertiesPtrOutput)
}

func (o VpnSiteLinkOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VpnSiteLinkArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLink)(nil)).Elem()
}

func (o VpnSiteLinkArrayOutput) ToVpnSiteLinkArrayOutput() VpnSiteLinkArrayOutput {
	return o
}

func (o VpnSiteLinkArrayOutput) ToVpnSiteLinkArrayOutputWithContext(ctx context.Context) VpnSiteLinkArrayOutput {
	return o
}

func (o VpnSiteLinkArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLink {
		return vs[0].([]VpnSiteLink)[vs[1].(int)]
	}).(VpnSiteLinkOutput)
}

type VpnSiteLinkConnection struct {
	ConnectionBandwidth            *int          `pulumi:"connectionBandwidth"`
	EgressNatRules                 []SubResource `pulumi:"egressNatRules"`
	EnableBgp                      *bool         `pulumi:"enableBgp"`
	EnableRateLimiting             *bool         `pulumi:"enableRateLimiting"`
	Id                             *string       `pulumi:"id"`
	IngressNatRules                []SubResource `pulumi:"ingressNatRules"`
	IpsecPolicies                  []IpsecPolicy `pulumi:"ipsecPolicies"`
	Name                           *string       `pulumi:"name"`
	RoutingWeight                  *int          `pulumi:"routingWeight"`
	SharedKey                      *string       `pulumi:"sharedKey"`
	UseLocalAzureIpAddress         *bool         `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool         `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      *string       `pulumi:"vpnConnectionProtocolType"`
	VpnLinkConnectionMode          *string       `pulumi:"vpnLinkConnectionMode"`
	VpnSiteLink                    *SubResource  `pulumi:"vpnSiteLink"`
}





type VpnSiteLinkConnectionInput interface {
	pulumi.Input

	ToVpnSiteLinkConnectionOutput() VpnSiteLinkConnectionOutput
	ToVpnSiteLinkConnectionOutputWithContext(context.Context) VpnSiteLinkConnectionOutput
}

type VpnSiteLinkConnectionArgs struct {
	ConnectionBandwidth            pulumi.IntPtrInput    `pulumi:"connectionBandwidth"`
	EgressNatRules                 SubResourceArrayInput `pulumi:"egressNatRules"`
	EnableBgp                      pulumi.BoolPtrInput   `pulumi:"enableBgp"`
	EnableRateLimiting             pulumi.BoolPtrInput   `pulumi:"enableRateLimiting"`
	Id                             pulumi.StringPtrInput `pulumi:"id"`
	IngressNatRules                SubResourceArrayInput `pulumi:"ingressNatRules"`
	IpsecPolicies                  IpsecPolicyArrayInput `pulumi:"ipsecPolicies"`
	Name                           pulumi.StringPtrInput `pulumi:"name"`
	RoutingWeight                  pulumi.IntPtrInput    `pulumi:"routingWeight"`
	SharedKey                      pulumi.StringPtrInput `pulumi:"sharedKey"`
	UseLocalAzureIpAddress         pulumi.BoolPtrInput   `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors pulumi.BoolPtrInput   `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      pulumi.StringPtrInput `pulumi:"vpnConnectionProtocolType"`
	VpnLinkConnectionMode          pulumi.StringPtrInput `pulumi:"vpnLinkConnectionMode"`
	VpnSiteLink                    SubResourcePtrInput   `pulumi:"vpnSiteLink"`
}

func (VpnSiteLinkConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkConnection)(nil)).Elem()
}

func (i VpnSiteLinkConnectionArgs) ToVpnSiteLinkConnectionOutput() VpnSiteLinkConnectionOutput {
	return i.ToVpnSiteLinkConnectionOutputWithContext(context.Background())
}

func (i VpnSiteLinkConnectionArgs) ToVpnSiteLinkConnectionOutputWithContext(ctx context.Context) VpnSiteLinkConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkConnectionOutput)
}





type VpnSiteLinkConnectionArrayInput interface {
	pulumi.Input

	ToVpnSiteLinkConnectionArrayOutput() VpnSiteLinkConnectionArrayOutput
	ToVpnSiteLinkConnectionArrayOutputWithContext(context.Context) VpnSiteLinkConnectionArrayOutput
}

type VpnSiteLinkConnectionArray []VpnSiteLinkConnectionInput

func (VpnSiteLinkConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkConnection)(nil)).Elem()
}

func (i VpnSiteLinkConnectionArray) ToVpnSiteLinkConnectionArrayOutput() VpnSiteLinkConnectionArrayOutput {
	return i.ToVpnSiteLinkConnectionArrayOutputWithContext(context.Background())
}

func (i VpnSiteLinkConnectionArray) ToVpnSiteLinkConnectionArrayOutputWithContext(ctx context.Context) VpnSiteLinkConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkConnectionArrayOutput)
}

type VpnSiteLinkConnectionOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkConnection)(nil)).Elem()
}

func (o VpnSiteLinkConnectionOutput) ToVpnSiteLinkConnectionOutput() VpnSiteLinkConnectionOutput {
	return o
}

func (o VpnSiteLinkConnectionOutput) ToVpnSiteLinkConnectionOutputWithContext(ctx context.Context) VpnSiteLinkConnectionOutput {
	return o
}

func (o VpnSiteLinkConnectionOutput) ConnectionBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *int { return v.ConnectionBandwidth }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) EgressNatRules() SubResourceArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) []SubResource { return v.EgressNatRules }).(SubResourceArrayOutput)
}

func (o VpnSiteLinkConnectionOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) EnableRateLimiting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.EnableRateLimiting }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) IngressNatRules() SubResourceArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) []SubResource { return v.IngressNatRules }).(SubResourceArrayOutput)
}

func (o VpnSiteLinkConnectionOutput) IpsecPolicies() IpsecPolicyArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) []IpsecPolicy { return v.IpsecPolicies }).(IpsecPolicyArrayOutput)
}

func (o VpnSiteLinkConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) VpnConnectionProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.VpnConnectionProtocolType }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) VpnLinkConnectionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.VpnLinkConnectionMode }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) VpnSiteLink() SubResourcePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *SubResource { return v.VpnSiteLink }).(SubResourcePtrOutput)
}

type VpnSiteLinkConnectionArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkConnection)(nil)).Elem()
}

func (o VpnSiteLinkConnectionArrayOutput) ToVpnSiteLinkConnectionArrayOutput() VpnSiteLinkConnectionArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionArrayOutput) ToVpnSiteLinkConnectionArrayOutputWithContext(ctx context.Context) VpnSiteLinkConnectionArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLinkConnection {
		return vs[0].([]VpnSiteLinkConnection)[vs[1].(int)]
	}).(VpnSiteLinkConnectionOutput)
}

type VpnSiteLinkConnectionResponse struct {
	ConnectionBandwidth            *int                  `pulumi:"connectionBandwidth"`
	ConnectionStatus               string                `pulumi:"connectionStatus"`
	EgressBytesTransferred         float64               `pulumi:"egressBytesTransferred"`
	EgressNatRules                 []SubResourceResponse `pulumi:"egressNatRules"`
	EnableBgp                      *bool                 `pulumi:"enableBgp"`
	EnableRateLimiting             *bool                 `pulumi:"enableRateLimiting"`
	Etag                           string                `pulumi:"etag"`
	Id                             *string               `pulumi:"id"`
	IngressBytesTransferred        float64               `pulumi:"ingressBytesTransferred"`
	IngressNatRules                []SubResourceResponse `pulumi:"ingressNatRules"`
	IpsecPolicies                  []IpsecPolicyResponse `pulumi:"ipsecPolicies"`
	Name                           *string               `pulumi:"name"`
	ProvisioningState              string                `pulumi:"provisioningState"`
	RoutingWeight                  *int                  `pulumi:"routingWeight"`
	SharedKey                      *string               `pulumi:"sharedKey"`
	Type                           string                `pulumi:"type"`
	UseLocalAzureIpAddress         *bool                 `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                 `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      *string               `pulumi:"vpnConnectionProtocolType"`
	VpnLinkConnectionMode          *string               `pulumi:"vpnLinkConnectionMode"`
	VpnSiteLink                    *SubResourceResponse  `pulumi:"vpnSiteLink"`
}

type VpnSiteLinkConnectionResponseOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkConnectionResponse)(nil)).Elem()
}

func (o VpnSiteLinkConnectionResponseOutput) ToVpnSiteLinkConnectionResponseOutput() VpnSiteLinkConnectionResponseOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseOutput) ToVpnSiteLinkConnectionResponseOutputWithContext(ctx context.Context) VpnSiteLinkConnectionResponseOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseOutput) ConnectionBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *int { return v.ConnectionBandwidth }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnSiteLinkConnectionResponseOutput) EgressNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) []SubResourceResponse { return v.EgressNatRules }).(SubResourceResponseArrayOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) EnableRateLimiting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.EnableRateLimiting }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnSiteLinkConnectionResponseOutput) IngressNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) []SubResourceResponse { return v.IngressNatRules }).(SubResourceResponseArrayOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) []IpsecPolicyResponse { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) VpnConnectionProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.VpnConnectionProtocolType }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) VpnLinkConnectionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.VpnLinkConnectionMode }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) VpnSiteLink() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *SubResourceResponse { return v.VpnSiteLink }).(SubResourceResponsePtrOutput)
}

type VpnSiteLinkConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkConnectionResponse)(nil)).Elem()
}

func (o VpnSiteLinkConnectionResponseArrayOutput) ToVpnSiteLinkConnectionResponseArrayOutput() VpnSiteLinkConnectionResponseArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseArrayOutput) ToVpnSiteLinkConnectionResponseArrayOutputWithContext(ctx context.Context) VpnSiteLinkConnectionResponseArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLinkConnectionResponse {
		return vs[0].([]VpnSiteLinkConnectionResponse)[vs[1].(int)]
	}).(VpnSiteLinkConnectionResponseOutput)
}

type VpnSiteLinkResponse struct {
	BgpProperties     *VpnLinkBgpSettingsResponse        `pulumi:"bgpProperties"`
	Etag              string                             `pulumi:"etag"`
	Fqdn              *string                            `pulumi:"fqdn"`
	Id                *string                            `pulumi:"id"`
	IpAddress         *string                            `pulumi:"ipAddress"`
	LinkProperties    *VpnLinkProviderPropertiesResponse `pulumi:"linkProperties"`
	Name              *string                            `pulumi:"name"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	Type              string                             `pulumi:"type"`
}

type VpnSiteLinkResponseOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkResponse)(nil)).Elem()
}

func (o VpnSiteLinkResponseOutput) ToVpnSiteLinkResponseOutput() VpnSiteLinkResponseOutput {
	return o
}

func (o VpnSiteLinkResponseOutput) ToVpnSiteLinkResponseOutputWithContext(ctx context.Context) VpnSiteLinkResponseOutput {
	return o
}

func (o VpnSiteLinkResponseOutput) BgpProperties() VpnLinkBgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *VpnLinkBgpSettingsResponse { return v.BgpProperties }).(VpnLinkBgpSettingsResponsePtrOutput)
}

func (o VpnSiteLinkResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnSiteLinkResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) LinkProperties() VpnLinkProviderPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *VpnLinkProviderPropertiesResponse { return v.LinkProperties }).(VpnLinkProviderPropertiesResponsePtrOutput)
}

func (o VpnSiteLinkResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnSiteLinkResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VpnSiteLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkResponse)(nil)).Elem()
}

func (o VpnSiteLinkResponseArrayOutput) ToVpnSiteLinkResponseArrayOutput() VpnSiteLinkResponseArrayOutput {
	return o
}

func (o VpnSiteLinkResponseArrayOutput) ToVpnSiteLinkResponseArrayOutputWithContext(ctx context.Context) VpnSiteLinkResponseArrayOutput {
	return o
}

func (o VpnSiteLinkResponseArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLinkResponse {
		return vs[0].([]VpnSiteLinkResponse)[vs[1].(int)]
	}).(VpnSiteLinkResponseOutput)
}

type WebApplicationFirewallCustomRule struct {
	Action          string           `pulumi:"action"`
	MatchConditions []MatchCondition `pulumi:"matchConditions"`
	Name            *string          `pulumi:"name"`
	Priority        int              `pulumi:"priority"`
	RuleType        string           `pulumi:"ruleType"`
}





type WebApplicationFirewallCustomRuleInput interface {
	pulumi.Input

	ToWebApplicationFirewallCustomRuleOutput() WebApplicationFirewallCustomRuleOutput
	ToWebApplicationFirewallCustomRuleOutputWithContext(context.Context) WebApplicationFirewallCustomRuleOutput
}

type WebApplicationFirewallCustomRuleArgs struct {
	Action          pulumi.StringInput       `pulumi:"action"`
	MatchConditions MatchConditionArrayInput `pulumi:"matchConditions"`
	Name            pulumi.StringPtrInput    `pulumi:"name"`
	Priority        pulumi.IntInput          `pulumi:"priority"`
	RuleType        pulumi.StringInput       `pulumi:"ruleType"`
}

func (WebApplicationFirewallCustomRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (i WebApplicationFirewallCustomRuleArgs) ToWebApplicationFirewallCustomRuleOutput() WebApplicationFirewallCustomRuleOutput {
	return i.ToWebApplicationFirewallCustomRuleOutputWithContext(context.Background())
}

func (i WebApplicationFirewallCustomRuleArgs) ToWebApplicationFirewallCustomRuleOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebApplicationFirewallCustomRuleOutput)
}





type WebApplicationFirewallCustomRuleArrayInput interface {
	pulumi.Input

	ToWebApplicationFirewallCustomRuleArrayOutput() WebApplicationFirewallCustomRuleArrayOutput
	ToWebApplicationFirewallCustomRuleArrayOutputWithContext(context.Context) WebApplicationFirewallCustomRuleArrayOutput
}

type WebApplicationFirewallCustomRuleArray []WebApplicationFirewallCustomRuleInput

func (WebApplicationFirewallCustomRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (i WebApplicationFirewallCustomRuleArray) ToWebApplicationFirewallCustomRuleArrayOutput() WebApplicationFirewallCustomRuleArrayOutput {
	return i.ToWebApplicationFirewallCustomRuleArrayOutputWithContext(context.Background())
}

func (i WebApplicationFirewallCustomRuleArray) ToWebApplicationFirewallCustomRuleArrayOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebApplicationFirewallCustomRuleArrayOutput)
}

type WebApplicationFirewallCustomRuleOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleOutput) ToWebApplicationFirewallCustomRuleOutput() WebApplicationFirewallCustomRuleOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleOutput) ToWebApplicationFirewallCustomRuleOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) string { return v.Action }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) MatchConditions() MatchConditionArrayOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) []MatchCondition { return v.MatchConditions }).(MatchConditionArrayOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) int { return v.Priority }).(pulumi.IntOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type WebApplicationFirewallCustomRuleArrayOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleArrayOutput) ToWebApplicationFirewallCustomRuleArrayOutput() WebApplicationFirewallCustomRuleArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleArrayOutput) ToWebApplicationFirewallCustomRuleArrayOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleArrayOutput) Index(i pulumi.IntInput) WebApplicationFirewallCustomRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebApplicationFirewallCustomRule {
		return vs[0].([]WebApplicationFirewallCustomRule)[vs[1].(int)]
	}).(WebApplicationFirewallCustomRuleOutput)
}

type WebApplicationFirewallCustomRuleResponse struct {
	Action          string                   `pulumi:"action"`
	Etag            string                   `pulumi:"etag"`
	MatchConditions []MatchConditionResponse `pulumi:"matchConditions"`
	Name            *string                  `pulumi:"name"`
	Priority        int                      `pulumi:"priority"`
	RuleType        string                   `pulumi:"ruleType"`
}

type WebApplicationFirewallCustomRuleResponseOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallCustomRuleResponse)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleResponseOutput) ToWebApplicationFirewallCustomRuleResponseOutput() WebApplicationFirewallCustomRuleResponseOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseOutput) ToWebApplicationFirewallCustomRuleResponseOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleResponseOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) MatchConditions() MatchConditionResponseArrayOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) []MatchConditionResponse { return v.MatchConditions }).(MatchConditionResponseArrayOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

type WebApplicationFirewallCustomRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebApplicationFirewallCustomRuleResponse)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleResponseArrayOutput) ToWebApplicationFirewallCustomRuleResponseArrayOutput() WebApplicationFirewallCustomRuleResponseArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseArrayOutput) ToWebApplicationFirewallCustomRuleResponseArrayOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleResponseArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseArrayOutput) Index(i pulumi.IntInput) WebApplicationFirewallCustomRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebApplicationFirewallCustomRuleResponse {
		return vs[0].([]WebApplicationFirewallCustomRuleResponse)[vs[1].(int)]
	}).(WebApplicationFirewallCustomRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(RoutingRuleUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput{})
	pulumi.RegisterOutputType(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(RoutingRuleUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput{})
	pulumi.RegisterOutputType(RulesEngineActionOutput{})
	pulumi.RegisterOutputType(RulesEngineActionResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineResponseArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SecurityPolicyLinkResponseOutput{})
	pulumi.RegisterOutputType(SecurityPolicyLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(SecurityRuleTypeOutput{})
	pulumi.RegisterOutputType(SecurityRuleTypeArrayOutput{})
	pulumi.RegisterOutputType(SecurityRuleResponseOutput{})
	pulumi.RegisterOutputType(SecurityRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceAssociationLinkResponseOutput{})
	pulumi.RegisterOutputType(ServiceAssociationLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPolicyTypeOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPolicyTypeArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPolicyDefinitionTypeOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPolicyDefinitionTypeArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPolicyDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPolicyDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPolicyResponseOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPropertiesFormatOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPropertiesFormatArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPropertiesFormatResponseOutput{})
	pulumi.RegisterOutputType(ServiceEndpointPropertiesFormatResponseArrayOutput{})
	pulumi.RegisterOutputType(SingleQueryResultResponseOutput{})
	pulumi.RegisterOutputType(SingleQueryResultResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SoaRecordOutput{})
	pulumi.RegisterOutputType(SoaRecordPtrOutput{})
	pulumi.RegisterOutputType(SoaRecordResponseOutput{})
	pulumi.RegisterOutputType(SoaRecordResponsePtrOutput{})
	pulumi.RegisterOutputType(SrvRecordOutput{})
	pulumi.RegisterOutputType(SrvRecordArrayOutput{})
	pulumi.RegisterOutputType(SrvRecordResponseOutput{})
	pulumi.RegisterOutputType(SrvRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(StaticRouteOutput{})
	pulumi.RegisterOutputType(StaticRouteArrayOutput{})
	pulumi.RegisterOutputType(StaticRouteResponseOutput{})
	pulumi.RegisterOutputType(StaticRouteResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SubnetTypeOutput{})
	pulumi.RegisterOutputType(SubnetTypePtrOutput{})
	pulumi.RegisterOutputType(SubnetTypeArrayOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TargetDnsServerOutput{})
	pulumi.RegisterOutputType(TargetDnsServerArrayOutput{})
	pulumi.RegisterOutputType(TargetDnsServerResponseOutput{})
	pulumi.RegisterOutputType(TargetDnsServerResponseArrayOutput{})
	pulumi.RegisterOutputType(TrafficAnalyticsConfigurationPropertiesOutput{})
	pulumi.RegisterOutputType(TrafficAnalyticsConfigurationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TrafficAnalyticsConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TrafficAnalyticsConfigurationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TrafficAnalyticsPropertiesOutput{})
	pulumi.RegisterOutputType(TrafficAnalyticsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TrafficAnalyticsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TrafficAnalyticsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TrafficSelectorPolicyOutput{})
	pulumi.RegisterOutputType(TrafficSelectorPolicyArrayOutput{})
	pulumi.RegisterOutputType(TrafficSelectorPolicyResponseOutput{})
	pulumi.RegisterOutputType(TrafficSelectorPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(TunnelConnectionHealthResponseOutput{})
	pulumi.RegisterOutputType(TunnelConnectionHealthResponseArrayOutput{})
	pulumi.RegisterOutputType(TxtRecordOutput{})
	pulumi.RegisterOutputType(TxtRecordArrayOutput{})
	pulumi.RegisterOutputType(TxtRecordResponseOutput{})
	pulumi.RegisterOutputType(TxtRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(VMOutput{})
	pulumi.RegisterOutputType(VMResponseOutput{})
	pulumi.RegisterOutputType(VirtualApplianceNicPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VirtualApplianceNicPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualApplianceSkuPropertiesOutput{})
	pulumi.RegisterOutputType(VirtualApplianceSkuPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VirtualApplianceSkuPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VirtualApplianceSkuPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualHubIdOutput{})
	pulumi.RegisterOutputType(VirtualHubIdResponseOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteArrayOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteResponseOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteTableOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteTablePtrOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteTableResponseOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteTableResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteTableV2TypeOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteTableV2TypeArrayOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteTableV2ResponseOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteTableV2ResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteV2Output{})
	pulumi.RegisterOutputType(VirtualHubRouteV2ArrayOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteV2ResponseOutput{})
	pulumi.RegisterOutputType(VirtualHubRouteV2ResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkBgpCommunitiesOutput{})
	pulumi.RegisterOutputType(VirtualNetworkBgpCommunitiesPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkBgpCommunitiesResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkBgpCommunitiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkDnsForwardingRulesetResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkDnsForwardingRulesetResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayTypeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayTypePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringTypeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringTypeArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTapTypeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTapTypePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTapTypeArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTapResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTapResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTapResponseArrayOutput{})
	pulumi.RegisterOutputType(VnetRouteOutput{})
	pulumi.RegisterOutputType(VnetRoutePtrOutput{})
	pulumi.RegisterOutputType(VnetRouteResponseOutput{})
	pulumi.RegisterOutputType(VnetRouteResponsePtrOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VpnClientConnectionHealthDetailResponseOutput{})
	pulumi.RegisterOutputType(VpnClientConnectionHealthDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnClientConnectionHealthResponseOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnConnectionTypeOutput{})
	pulumi.RegisterOutputType(VpnConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(VpnConnectionResponseOutput{})
	pulumi.RegisterOutputType(VpnConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnGatewayIpConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VpnGatewayIpConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnGatewayNatRuleOutput{})
	pulumi.RegisterOutputType(VpnGatewayNatRuleArrayOutput{})
	pulumi.RegisterOutputType(VpnGatewayNatRuleResponseOutput{})
	pulumi.RegisterOutputType(VpnGatewayNatRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnLinkBgpSettingsOutput{})
	pulumi.RegisterOutputType(VpnLinkBgpSettingsPtrOutput{})
	pulumi.RegisterOutputType(VpnLinkBgpSettingsResponseOutput{})
	pulumi.RegisterOutputType(VpnLinkBgpSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingArrayOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingResponseOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionResponseOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkResponseOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleArrayOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleResponseOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleResponseArrayOutput{})
}
