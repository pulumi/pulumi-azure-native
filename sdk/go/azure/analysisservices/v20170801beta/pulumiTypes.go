


package v20170801beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewayDetails struct {
	GatewayResourceId *string `pulumi:"gatewayResourceId"`
}





type GatewayDetailsInput interface {
	pulumi.Input

	ToGatewayDetailsOutput() GatewayDetailsOutput
	ToGatewayDetailsOutputWithContext(context.Context) GatewayDetailsOutput
}

type GatewayDetailsArgs struct {
	GatewayResourceId pulumi.StringPtrInput `pulumi:"gatewayResourceId"`
}

func (GatewayDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDetails)(nil)).Elem()
}

func (i GatewayDetailsArgs) ToGatewayDetailsOutput() GatewayDetailsOutput {
	return i.ToGatewayDetailsOutputWithContext(context.Background())
}

func (i GatewayDetailsArgs) ToGatewayDetailsOutputWithContext(ctx context.Context) GatewayDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsOutput)
}

func (i GatewayDetailsArgs) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return i.ToGatewayDetailsPtrOutputWithContext(context.Background())
}

func (i GatewayDetailsArgs) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsOutput).ToGatewayDetailsPtrOutputWithContext(ctx)
}









type GatewayDetailsPtrInput interface {
	pulumi.Input

	ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput
	ToGatewayDetailsPtrOutputWithContext(context.Context) GatewayDetailsPtrOutput
}

type gatewayDetailsPtrType GatewayDetailsArgs

func GatewayDetailsPtr(v *GatewayDetailsArgs) GatewayDetailsPtrInput {
	return (*gatewayDetailsPtrType)(v)
}

func (*gatewayDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayDetails)(nil)).Elem()
}

func (i *gatewayDetailsPtrType) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return i.ToGatewayDetailsPtrOutputWithContext(context.Background())
}

func (i *gatewayDetailsPtrType) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsPtrOutput)
}

type GatewayDetailsOutput struct{ *pulumi.OutputState }

func (GatewayDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDetails)(nil)).Elem()
}

func (o GatewayDetailsOutput) ToGatewayDetailsOutput() GatewayDetailsOutput {
	return o
}

func (o GatewayDetailsOutput) ToGatewayDetailsOutputWithContext(ctx context.Context) GatewayDetailsOutput {
	return o
}

func (o GatewayDetailsOutput) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return o.ToGatewayDetailsPtrOutputWithContext(context.Background())
}

func (o GatewayDetailsOutput) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayDetails) *GatewayDetails {
		return &v
	}).(GatewayDetailsPtrOutput)
}

func (o GatewayDetailsOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayDetails) *string { return v.GatewayResourceId }).(pulumi.StringPtrOutput)
}

type GatewayDetailsPtrOutput struct{ *pulumi.OutputState }

func (GatewayDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayDetails)(nil)).Elem()
}

func (o GatewayDetailsPtrOutput) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return o
}

func (o GatewayDetailsPtrOutput) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return o
}

func (o GatewayDetailsPtrOutput) Elem() GatewayDetailsOutput {
	return o.ApplyT(func(v *GatewayDetails) GatewayDetails {
		if v != nil {
			return *v
		}
		var ret GatewayDetails
		return ret
	}).(GatewayDetailsOutput)
}

func (o GatewayDetailsPtrOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetails) *string {
		if v == nil {
			return nil
		}
		return v.GatewayResourceId
	}).(pulumi.StringPtrOutput)
}

type GatewayDetailsResponse struct {
	DmtsClusterUri    string  `pulumi:"dmtsClusterUri"`
	GatewayObjectId   string  `pulumi:"gatewayObjectId"`
	GatewayResourceId *string `pulumi:"gatewayResourceId"`
}





type GatewayDetailsResponseInput interface {
	pulumi.Input

	ToGatewayDetailsResponseOutput() GatewayDetailsResponseOutput
	ToGatewayDetailsResponseOutputWithContext(context.Context) GatewayDetailsResponseOutput
}

type GatewayDetailsResponseArgs struct {
	DmtsClusterUri    pulumi.StringInput    `pulumi:"dmtsClusterUri"`
	GatewayObjectId   pulumi.StringInput    `pulumi:"gatewayObjectId"`
	GatewayResourceId pulumi.StringPtrInput `pulumi:"gatewayResourceId"`
}

func (GatewayDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDetailsResponse)(nil)).Elem()
}

func (i GatewayDetailsResponseArgs) ToGatewayDetailsResponseOutput() GatewayDetailsResponseOutput {
	return i.ToGatewayDetailsResponseOutputWithContext(context.Background())
}

func (i GatewayDetailsResponseArgs) ToGatewayDetailsResponseOutputWithContext(ctx context.Context) GatewayDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsResponseOutput)
}

func (i GatewayDetailsResponseArgs) ToGatewayDetailsResponsePtrOutput() GatewayDetailsResponsePtrOutput {
	return i.ToGatewayDetailsResponsePtrOutputWithContext(context.Background())
}

func (i GatewayDetailsResponseArgs) ToGatewayDetailsResponsePtrOutputWithContext(ctx context.Context) GatewayDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsResponseOutput).ToGatewayDetailsResponsePtrOutputWithContext(ctx)
}









type GatewayDetailsResponsePtrInput interface {
	pulumi.Input

	ToGatewayDetailsResponsePtrOutput() GatewayDetailsResponsePtrOutput
	ToGatewayDetailsResponsePtrOutputWithContext(context.Context) GatewayDetailsResponsePtrOutput
}

type gatewayDetailsResponsePtrType GatewayDetailsResponseArgs

func GatewayDetailsResponsePtr(v *GatewayDetailsResponseArgs) GatewayDetailsResponsePtrInput {
	return (*gatewayDetailsResponsePtrType)(v)
}

func (*gatewayDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayDetailsResponse)(nil)).Elem()
}

func (i *gatewayDetailsResponsePtrType) ToGatewayDetailsResponsePtrOutput() GatewayDetailsResponsePtrOutput {
	return i.ToGatewayDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *gatewayDetailsResponsePtrType) ToGatewayDetailsResponsePtrOutputWithContext(ctx context.Context) GatewayDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsResponsePtrOutput)
}

type GatewayDetailsResponseOutput struct{ *pulumi.OutputState }

func (GatewayDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDetailsResponse)(nil)).Elem()
}

func (o GatewayDetailsResponseOutput) ToGatewayDetailsResponseOutput() GatewayDetailsResponseOutput {
	return o
}

func (o GatewayDetailsResponseOutput) ToGatewayDetailsResponseOutputWithContext(ctx context.Context) GatewayDetailsResponseOutput {
	return o
}

func (o GatewayDetailsResponseOutput) ToGatewayDetailsResponsePtrOutput() GatewayDetailsResponsePtrOutput {
	return o.ToGatewayDetailsResponsePtrOutputWithContext(context.Background())
}

func (o GatewayDetailsResponseOutput) ToGatewayDetailsResponsePtrOutputWithContext(ctx context.Context) GatewayDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayDetailsResponse) *GatewayDetailsResponse {
		return &v
	}).(GatewayDetailsResponsePtrOutput)
}

func (o GatewayDetailsResponseOutput) DmtsClusterUri() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDetailsResponse) string { return v.DmtsClusterUri }).(pulumi.StringOutput)
}

func (o GatewayDetailsResponseOutput) GatewayObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDetailsResponse) string { return v.GatewayObjectId }).(pulumi.StringOutput)
}

func (o GatewayDetailsResponseOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayDetailsResponse) *string { return v.GatewayResourceId }).(pulumi.StringPtrOutput)
}

type GatewayDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (GatewayDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayDetailsResponse)(nil)).Elem()
}

func (o GatewayDetailsResponsePtrOutput) ToGatewayDetailsResponsePtrOutput() GatewayDetailsResponsePtrOutput {
	return o
}

func (o GatewayDetailsResponsePtrOutput) ToGatewayDetailsResponsePtrOutputWithContext(ctx context.Context) GatewayDetailsResponsePtrOutput {
	return o
}

func (o GatewayDetailsResponsePtrOutput) Elem() GatewayDetailsResponseOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) GatewayDetailsResponse {
		if v != nil {
			return *v
		}
		var ret GatewayDetailsResponse
		return ret
	}).(GatewayDetailsResponseOutput)
}

func (o GatewayDetailsResponsePtrOutput) DmtsClusterUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DmtsClusterUri
	}).(pulumi.StringPtrOutput)
}

func (o GatewayDetailsResponsePtrOutput) GatewayObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.GatewayObjectId
	}).(pulumi.StringPtrOutput)
}

func (o GatewayDetailsResponsePtrOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GatewayResourceId
	}).(pulumi.StringPtrOutput)
}

type IPv4FirewallRule struct {
	FirewallRuleName *string `pulumi:"firewallRuleName"`
	RangeEnd         *string `pulumi:"rangeEnd"`
	RangeStart       *string `pulumi:"rangeStart"`
}





type IPv4FirewallRuleInput interface {
	pulumi.Input

	ToIPv4FirewallRuleOutput() IPv4FirewallRuleOutput
	ToIPv4FirewallRuleOutputWithContext(context.Context) IPv4FirewallRuleOutput
}

type IPv4FirewallRuleArgs struct {
	FirewallRuleName pulumi.StringPtrInput `pulumi:"firewallRuleName"`
	RangeEnd         pulumi.StringPtrInput `pulumi:"rangeEnd"`
	RangeStart       pulumi.StringPtrInput `pulumi:"rangeStart"`
}

func (IPv4FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallRule)(nil)).Elem()
}

func (i IPv4FirewallRuleArgs) ToIPv4FirewallRuleOutput() IPv4FirewallRuleOutput {
	return i.ToIPv4FirewallRuleOutputWithContext(context.Background())
}

func (i IPv4FirewallRuleArgs) ToIPv4FirewallRuleOutputWithContext(ctx context.Context) IPv4FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallRuleOutput)
}





type IPv4FirewallRuleArrayInput interface {
	pulumi.Input

	ToIPv4FirewallRuleArrayOutput() IPv4FirewallRuleArrayOutput
	ToIPv4FirewallRuleArrayOutputWithContext(context.Context) IPv4FirewallRuleArrayOutput
}

type IPv4FirewallRuleArray []IPv4FirewallRuleInput

func (IPv4FirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPv4FirewallRule)(nil)).Elem()
}

func (i IPv4FirewallRuleArray) ToIPv4FirewallRuleArrayOutput() IPv4FirewallRuleArrayOutput {
	return i.ToIPv4FirewallRuleArrayOutputWithContext(context.Background())
}

func (i IPv4FirewallRuleArray) ToIPv4FirewallRuleArrayOutputWithContext(ctx context.Context) IPv4FirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallRuleArrayOutput)
}

type IPv4FirewallRuleOutput struct{ *pulumi.OutputState }

func (IPv4FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallRule)(nil)).Elem()
}

func (o IPv4FirewallRuleOutput) ToIPv4FirewallRuleOutput() IPv4FirewallRuleOutput {
	return o
}

func (o IPv4FirewallRuleOutput) ToIPv4FirewallRuleOutputWithContext(ctx context.Context) IPv4FirewallRuleOutput {
	return o
}

func (o IPv4FirewallRuleOutput) FirewallRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRule) *string { return v.FirewallRuleName }).(pulumi.StringPtrOutput)
}

func (o IPv4FirewallRuleOutput) RangeEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRule) *string { return v.RangeEnd }).(pulumi.StringPtrOutput)
}

func (o IPv4FirewallRuleOutput) RangeStart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRule) *string { return v.RangeStart }).(pulumi.StringPtrOutput)
}

type IPv4FirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (IPv4FirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPv4FirewallRule)(nil)).Elem()
}

func (o IPv4FirewallRuleArrayOutput) ToIPv4FirewallRuleArrayOutput() IPv4FirewallRuleArrayOutput {
	return o
}

func (o IPv4FirewallRuleArrayOutput) ToIPv4FirewallRuleArrayOutputWithContext(ctx context.Context) IPv4FirewallRuleArrayOutput {
	return o
}

func (o IPv4FirewallRuleArrayOutput) Index(i pulumi.IntInput) IPv4FirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPv4FirewallRule {
		return vs[0].([]IPv4FirewallRule)[vs[1].(int)]
	}).(IPv4FirewallRuleOutput)
}

type IPv4FirewallRuleResponse struct {
	FirewallRuleName *string `pulumi:"firewallRuleName"`
	RangeEnd         *string `pulumi:"rangeEnd"`
	RangeStart       *string `pulumi:"rangeStart"`
}





type IPv4FirewallRuleResponseInput interface {
	pulumi.Input

	ToIPv4FirewallRuleResponseOutput() IPv4FirewallRuleResponseOutput
	ToIPv4FirewallRuleResponseOutputWithContext(context.Context) IPv4FirewallRuleResponseOutput
}

type IPv4FirewallRuleResponseArgs struct {
	FirewallRuleName pulumi.StringPtrInput `pulumi:"firewallRuleName"`
	RangeEnd         pulumi.StringPtrInput `pulumi:"rangeEnd"`
	RangeStart       pulumi.StringPtrInput `pulumi:"rangeStart"`
}

func (IPv4FirewallRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallRuleResponse)(nil)).Elem()
}

func (i IPv4FirewallRuleResponseArgs) ToIPv4FirewallRuleResponseOutput() IPv4FirewallRuleResponseOutput {
	return i.ToIPv4FirewallRuleResponseOutputWithContext(context.Background())
}

func (i IPv4FirewallRuleResponseArgs) ToIPv4FirewallRuleResponseOutputWithContext(ctx context.Context) IPv4FirewallRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallRuleResponseOutput)
}





type IPv4FirewallRuleResponseArrayInput interface {
	pulumi.Input

	ToIPv4FirewallRuleResponseArrayOutput() IPv4FirewallRuleResponseArrayOutput
	ToIPv4FirewallRuleResponseArrayOutputWithContext(context.Context) IPv4FirewallRuleResponseArrayOutput
}

type IPv4FirewallRuleResponseArray []IPv4FirewallRuleResponseInput

func (IPv4FirewallRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPv4FirewallRuleResponse)(nil)).Elem()
}

func (i IPv4FirewallRuleResponseArray) ToIPv4FirewallRuleResponseArrayOutput() IPv4FirewallRuleResponseArrayOutput {
	return i.ToIPv4FirewallRuleResponseArrayOutputWithContext(context.Background())
}

func (i IPv4FirewallRuleResponseArray) ToIPv4FirewallRuleResponseArrayOutputWithContext(ctx context.Context) IPv4FirewallRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallRuleResponseArrayOutput)
}

type IPv4FirewallRuleResponseOutput struct{ *pulumi.OutputState }

func (IPv4FirewallRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallRuleResponse)(nil)).Elem()
}

func (o IPv4FirewallRuleResponseOutput) ToIPv4FirewallRuleResponseOutput() IPv4FirewallRuleResponseOutput {
	return o
}

func (o IPv4FirewallRuleResponseOutput) ToIPv4FirewallRuleResponseOutputWithContext(ctx context.Context) IPv4FirewallRuleResponseOutput {
	return o
}

func (o IPv4FirewallRuleResponseOutput) FirewallRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRuleResponse) *string { return v.FirewallRuleName }).(pulumi.StringPtrOutput)
}

func (o IPv4FirewallRuleResponseOutput) RangeEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRuleResponse) *string { return v.RangeEnd }).(pulumi.StringPtrOutput)
}

func (o IPv4FirewallRuleResponseOutput) RangeStart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRuleResponse) *string { return v.RangeStart }).(pulumi.StringPtrOutput)
}

type IPv4FirewallRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IPv4FirewallRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPv4FirewallRuleResponse)(nil)).Elem()
}

func (o IPv4FirewallRuleResponseArrayOutput) ToIPv4FirewallRuleResponseArrayOutput() IPv4FirewallRuleResponseArrayOutput {
	return o
}

func (o IPv4FirewallRuleResponseArrayOutput) ToIPv4FirewallRuleResponseArrayOutputWithContext(ctx context.Context) IPv4FirewallRuleResponseArrayOutput {
	return o
}

func (o IPv4FirewallRuleResponseArrayOutput) Index(i pulumi.IntInput) IPv4FirewallRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPv4FirewallRuleResponse {
		return vs[0].([]IPv4FirewallRuleResponse)[vs[1].(int)]
	}).(IPv4FirewallRuleResponseOutput)
}

type IPv4FirewallSettings struct {
	EnablePowerBIService *bool              `pulumi:"enablePowerBIService"`
	FirewallRules        []IPv4FirewallRule `pulumi:"firewallRules"`
}





type IPv4FirewallSettingsInput interface {
	pulumi.Input

	ToIPv4FirewallSettingsOutput() IPv4FirewallSettingsOutput
	ToIPv4FirewallSettingsOutputWithContext(context.Context) IPv4FirewallSettingsOutput
}

type IPv4FirewallSettingsArgs struct {
	EnablePowerBIService pulumi.BoolPtrInput        `pulumi:"enablePowerBIService"`
	FirewallRules        IPv4FirewallRuleArrayInput `pulumi:"firewallRules"`
}

func (IPv4FirewallSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallSettings)(nil)).Elem()
}

func (i IPv4FirewallSettingsArgs) ToIPv4FirewallSettingsOutput() IPv4FirewallSettingsOutput {
	return i.ToIPv4FirewallSettingsOutputWithContext(context.Background())
}

func (i IPv4FirewallSettingsArgs) ToIPv4FirewallSettingsOutputWithContext(ctx context.Context) IPv4FirewallSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallSettingsOutput)
}

func (i IPv4FirewallSettingsArgs) ToIPv4FirewallSettingsPtrOutput() IPv4FirewallSettingsPtrOutput {
	return i.ToIPv4FirewallSettingsPtrOutputWithContext(context.Background())
}

func (i IPv4FirewallSettingsArgs) ToIPv4FirewallSettingsPtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallSettingsOutput).ToIPv4FirewallSettingsPtrOutputWithContext(ctx)
}









type IPv4FirewallSettingsPtrInput interface {
	pulumi.Input

	ToIPv4FirewallSettingsPtrOutput() IPv4FirewallSettingsPtrOutput
	ToIPv4FirewallSettingsPtrOutputWithContext(context.Context) IPv4FirewallSettingsPtrOutput
}

type ipv4FirewallSettingsPtrType IPv4FirewallSettingsArgs

func IPv4FirewallSettingsPtr(v *IPv4FirewallSettingsArgs) IPv4FirewallSettingsPtrInput {
	return (*ipv4FirewallSettingsPtrType)(v)
}

func (*ipv4FirewallSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IPv4FirewallSettings)(nil)).Elem()
}

func (i *ipv4FirewallSettingsPtrType) ToIPv4FirewallSettingsPtrOutput() IPv4FirewallSettingsPtrOutput {
	return i.ToIPv4FirewallSettingsPtrOutputWithContext(context.Background())
}

func (i *ipv4FirewallSettingsPtrType) ToIPv4FirewallSettingsPtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallSettingsPtrOutput)
}

type IPv4FirewallSettingsOutput struct{ *pulumi.OutputState }

func (IPv4FirewallSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallSettings)(nil)).Elem()
}

func (o IPv4FirewallSettingsOutput) ToIPv4FirewallSettingsOutput() IPv4FirewallSettingsOutput {
	return o
}

func (o IPv4FirewallSettingsOutput) ToIPv4FirewallSettingsOutputWithContext(ctx context.Context) IPv4FirewallSettingsOutput {
	return o
}

func (o IPv4FirewallSettingsOutput) ToIPv4FirewallSettingsPtrOutput() IPv4FirewallSettingsPtrOutput {
	return o.ToIPv4FirewallSettingsPtrOutputWithContext(context.Background())
}

func (o IPv4FirewallSettingsOutput) ToIPv4FirewallSettingsPtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPv4FirewallSettings) *IPv4FirewallSettings {
		return &v
	}).(IPv4FirewallSettingsPtrOutput)
}

func (o IPv4FirewallSettingsOutput) EnablePowerBIService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IPv4FirewallSettings) *bool { return v.EnablePowerBIService }).(pulumi.BoolPtrOutput)
}

func (o IPv4FirewallSettingsOutput) FirewallRules() IPv4FirewallRuleArrayOutput {
	return o.ApplyT(func(v IPv4FirewallSettings) []IPv4FirewallRule { return v.FirewallRules }).(IPv4FirewallRuleArrayOutput)
}

type IPv4FirewallSettingsPtrOutput struct{ *pulumi.OutputState }

func (IPv4FirewallSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPv4FirewallSettings)(nil)).Elem()
}

func (o IPv4FirewallSettingsPtrOutput) ToIPv4FirewallSettingsPtrOutput() IPv4FirewallSettingsPtrOutput {
	return o
}

func (o IPv4FirewallSettingsPtrOutput) ToIPv4FirewallSettingsPtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsPtrOutput {
	return o
}

func (o IPv4FirewallSettingsPtrOutput) Elem() IPv4FirewallSettingsOutput {
	return o.ApplyT(func(v *IPv4FirewallSettings) IPv4FirewallSettings {
		if v != nil {
			return *v
		}
		var ret IPv4FirewallSettings
		return ret
	}).(IPv4FirewallSettingsOutput)
}

func (o IPv4FirewallSettingsPtrOutput) EnablePowerBIService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IPv4FirewallSettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePowerBIService
	}).(pulumi.BoolPtrOutput)
}

func (o IPv4FirewallSettingsPtrOutput) FirewallRules() IPv4FirewallRuleArrayOutput {
	return o.ApplyT(func(v *IPv4FirewallSettings) []IPv4FirewallRule {
		if v == nil {
			return nil
		}
		return v.FirewallRules
	}).(IPv4FirewallRuleArrayOutput)
}

type IPv4FirewallSettingsResponse struct {
	EnablePowerBIService *bool                      `pulumi:"enablePowerBIService"`
	FirewallRules        []IPv4FirewallRuleResponse `pulumi:"firewallRules"`
}





type IPv4FirewallSettingsResponseInput interface {
	pulumi.Input

	ToIPv4FirewallSettingsResponseOutput() IPv4FirewallSettingsResponseOutput
	ToIPv4FirewallSettingsResponseOutputWithContext(context.Context) IPv4FirewallSettingsResponseOutput
}

type IPv4FirewallSettingsResponseArgs struct {
	EnablePowerBIService pulumi.BoolPtrInput                `pulumi:"enablePowerBIService"`
	FirewallRules        IPv4FirewallRuleResponseArrayInput `pulumi:"firewallRules"`
}

func (IPv4FirewallSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallSettingsResponse)(nil)).Elem()
}

func (i IPv4FirewallSettingsResponseArgs) ToIPv4FirewallSettingsResponseOutput() IPv4FirewallSettingsResponseOutput {
	return i.ToIPv4FirewallSettingsResponseOutputWithContext(context.Background())
}

func (i IPv4FirewallSettingsResponseArgs) ToIPv4FirewallSettingsResponseOutputWithContext(ctx context.Context) IPv4FirewallSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallSettingsResponseOutput)
}

func (i IPv4FirewallSettingsResponseArgs) ToIPv4FirewallSettingsResponsePtrOutput() IPv4FirewallSettingsResponsePtrOutput {
	return i.ToIPv4FirewallSettingsResponsePtrOutputWithContext(context.Background())
}

func (i IPv4FirewallSettingsResponseArgs) ToIPv4FirewallSettingsResponsePtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallSettingsResponseOutput).ToIPv4FirewallSettingsResponsePtrOutputWithContext(ctx)
}









type IPv4FirewallSettingsResponsePtrInput interface {
	pulumi.Input

	ToIPv4FirewallSettingsResponsePtrOutput() IPv4FirewallSettingsResponsePtrOutput
	ToIPv4FirewallSettingsResponsePtrOutputWithContext(context.Context) IPv4FirewallSettingsResponsePtrOutput
}

type ipv4FirewallSettingsResponsePtrType IPv4FirewallSettingsResponseArgs

func IPv4FirewallSettingsResponsePtr(v *IPv4FirewallSettingsResponseArgs) IPv4FirewallSettingsResponsePtrInput {
	return (*ipv4FirewallSettingsResponsePtrType)(v)
}

func (*ipv4FirewallSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IPv4FirewallSettingsResponse)(nil)).Elem()
}

func (i *ipv4FirewallSettingsResponsePtrType) ToIPv4FirewallSettingsResponsePtrOutput() IPv4FirewallSettingsResponsePtrOutput {
	return i.ToIPv4FirewallSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *ipv4FirewallSettingsResponsePtrType) ToIPv4FirewallSettingsResponsePtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallSettingsResponsePtrOutput)
}

type IPv4FirewallSettingsResponseOutput struct{ *pulumi.OutputState }

func (IPv4FirewallSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallSettingsResponse)(nil)).Elem()
}

func (o IPv4FirewallSettingsResponseOutput) ToIPv4FirewallSettingsResponseOutput() IPv4FirewallSettingsResponseOutput {
	return o
}

func (o IPv4FirewallSettingsResponseOutput) ToIPv4FirewallSettingsResponseOutputWithContext(ctx context.Context) IPv4FirewallSettingsResponseOutput {
	return o
}

func (o IPv4FirewallSettingsResponseOutput) ToIPv4FirewallSettingsResponsePtrOutput() IPv4FirewallSettingsResponsePtrOutput {
	return o.ToIPv4FirewallSettingsResponsePtrOutputWithContext(context.Background())
}

func (o IPv4FirewallSettingsResponseOutput) ToIPv4FirewallSettingsResponsePtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPv4FirewallSettingsResponse) *IPv4FirewallSettingsResponse {
		return &v
	}).(IPv4FirewallSettingsResponsePtrOutput)
}

func (o IPv4FirewallSettingsResponseOutput) EnablePowerBIService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IPv4FirewallSettingsResponse) *bool { return v.EnablePowerBIService }).(pulumi.BoolPtrOutput)
}

func (o IPv4FirewallSettingsResponseOutput) FirewallRules() IPv4FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v IPv4FirewallSettingsResponse) []IPv4FirewallRuleResponse { return v.FirewallRules }).(IPv4FirewallRuleResponseArrayOutput)
}

type IPv4FirewallSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (IPv4FirewallSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPv4FirewallSettingsResponse)(nil)).Elem()
}

func (o IPv4FirewallSettingsResponsePtrOutput) ToIPv4FirewallSettingsResponsePtrOutput() IPv4FirewallSettingsResponsePtrOutput {
	return o
}

func (o IPv4FirewallSettingsResponsePtrOutput) ToIPv4FirewallSettingsResponsePtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsResponsePtrOutput {
	return o
}

func (o IPv4FirewallSettingsResponsePtrOutput) Elem() IPv4FirewallSettingsResponseOutput {
	return o.ApplyT(func(v *IPv4FirewallSettingsResponse) IPv4FirewallSettingsResponse {
		if v != nil {
			return *v
		}
		var ret IPv4FirewallSettingsResponse
		return ret
	}).(IPv4FirewallSettingsResponseOutput)
}

func (o IPv4FirewallSettingsResponsePtrOutput) EnablePowerBIService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IPv4FirewallSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePowerBIService
	}).(pulumi.BoolPtrOutput)
}

func (o IPv4FirewallSettingsResponsePtrOutput) FirewallRules() IPv4FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v *IPv4FirewallSettingsResponse) []IPv4FirewallRuleResponse {
		if v == nil {
			return nil
		}
		return v.FirewallRules
	}).(IPv4FirewallRuleResponseArrayOutput)
}

type ResourceSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (i ResourceSkuArgs) ToResourceSkuOutput() ResourceSkuOutput {
	return i.ToResourceSkuOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput)
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput).ToResourceSkuPtrOutputWithContext(ctx)
}









type ResourceSkuPtrInput interface {
	pulumi.Input

	ToResourceSkuPtrOutput() ResourceSkuPtrOutput
	ToResourceSkuPtrOutputWithContext(context.Context) ResourceSkuPtrOutput
}

type resourceSkuPtrType ResourceSkuArgs

func ResourceSkuPtr(v *ResourceSkuArgs) ResourceSkuPtrInput {
	return (*resourceSkuPtrType)(v)
}

func (*resourceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuPtrOutput)
}

type ResourceSkuOutput struct{ *pulumi.OutputState }

func (ResourceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (o ResourceSkuOutput) ToResourceSkuOutput() ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSku) *ResourceSku {
		return &v
	}).(ResourceSkuPtrOutput)
}

func (o ResourceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) Elem() ResourceSkuOutput {
	return o.ApplyT(func(v *ResourceSku) ResourceSku {
		if v != nil {
			return *v
		}
		var ret ResourceSku
		return ret
	}).(ResourceSkuOutput)
}

func (o ResourceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type ResourceSkuResponseInput interface {
	pulumi.Input

	ToResourceSkuResponseOutput() ResourceSkuResponseOutput
	ToResourceSkuResponseOutputWithContext(context.Context) ResourceSkuResponseOutput
}

type ResourceSkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return i.ToResourceSkuResponseOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput)
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput).ToResourceSkuResponsePtrOutputWithContext(ctx)
}









type ResourceSkuResponsePtrInput interface {
	pulumi.Input

	ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput
	ToResourceSkuResponsePtrOutputWithContext(context.Context) ResourceSkuResponsePtrOutput
}

type resourceSkuResponsePtrType ResourceSkuResponseArgs

func ResourceSkuResponsePtr(v *ResourceSkuResponseArgs) ResourceSkuResponsePtrInput {
	return (*resourceSkuResponsePtrType)(v)
}

func (*resourceSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponsePtrOutput)
}

type ResourceSkuResponseOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSkuResponse) *ResourceSkuResponse {
		return &v
	}).(ResourceSkuResponsePtrOutput)
}

func (o ResourceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) Elem() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) ResourceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSkuResponse
		return ret
	}).(ResourceSkuResponseOutput)
}

func (o ResourceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ServerAdministrators struct {
	Members []string `pulumi:"members"`
}





type ServerAdministratorsInput interface {
	pulumi.Input

	ToServerAdministratorsOutput() ServerAdministratorsOutput
	ToServerAdministratorsOutputWithContext(context.Context) ServerAdministratorsOutput
}

type ServerAdministratorsArgs struct {
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (ServerAdministratorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministrators)(nil)).Elem()
}

func (i ServerAdministratorsArgs) ToServerAdministratorsOutput() ServerAdministratorsOutput {
	return i.ToServerAdministratorsOutputWithContext(context.Background())
}

func (i ServerAdministratorsArgs) ToServerAdministratorsOutputWithContext(ctx context.Context) ServerAdministratorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsOutput)
}

func (i ServerAdministratorsArgs) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return i.ToServerAdministratorsPtrOutputWithContext(context.Background())
}

func (i ServerAdministratorsArgs) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsOutput).ToServerAdministratorsPtrOutputWithContext(ctx)
}









type ServerAdministratorsPtrInput interface {
	pulumi.Input

	ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput
	ToServerAdministratorsPtrOutputWithContext(context.Context) ServerAdministratorsPtrOutput
}

type serverAdministratorsPtrType ServerAdministratorsArgs

func ServerAdministratorsPtr(v *ServerAdministratorsArgs) ServerAdministratorsPtrInput {
	return (*serverAdministratorsPtrType)(v)
}

func (*serverAdministratorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdministrators)(nil)).Elem()
}

func (i *serverAdministratorsPtrType) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return i.ToServerAdministratorsPtrOutputWithContext(context.Background())
}

func (i *serverAdministratorsPtrType) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsPtrOutput)
}

type ServerAdministratorsOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministrators)(nil)).Elem()
}

func (o ServerAdministratorsOutput) ToServerAdministratorsOutput() ServerAdministratorsOutput {
	return o
}

func (o ServerAdministratorsOutput) ToServerAdministratorsOutputWithContext(ctx context.Context) ServerAdministratorsOutput {
	return o
}

func (o ServerAdministratorsOutput) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return o.ToServerAdministratorsPtrOutputWithContext(context.Background())
}

func (o ServerAdministratorsOutput) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerAdministrators) *ServerAdministrators {
		return &v
	}).(ServerAdministratorsPtrOutput)
}

func (o ServerAdministratorsOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServerAdministrators) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type ServerAdministratorsPtrOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdministrators)(nil)).Elem()
}

func (o ServerAdministratorsPtrOutput) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return o
}

func (o ServerAdministratorsPtrOutput) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return o
}

func (o ServerAdministratorsPtrOutput) Elem() ServerAdministratorsOutput {
	return o.ApplyT(func(v *ServerAdministrators) ServerAdministrators {
		if v != nil {
			return *v
		}
		var ret ServerAdministrators
		return ret
	}).(ServerAdministratorsOutput)
}

func (o ServerAdministratorsPtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerAdministrators) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

type ServerAdministratorsResponse struct {
	Members []string `pulumi:"members"`
}





type ServerAdministratorsResponseInput interface {
	pulumi.Input

	ToServerAdministratorsResponseOutput() ServerAdministratorsResponseOutput
	ToServerAdministratorsResponseOutputWithContext(context.Context) ServerAdministratorsResponseOutput
}

type ServerAdministratorsResponseArgs struct {
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (ServerAdministratorsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministratorsResponse)(nil)).Elem()
}

func (i ServerAdministratorsResponseArgs) ToServerAdministratorsResponseOutput() ServerAdministratorsResponseOutput {
	return i.ToServerAdministratorsResponseOutputWithContext(context.Background())
}

func (i ServerAdministratorsResponseArgs) ToServerAdministratorsResponseOutputWithContext(ctx context.Context) ServerAdministratorsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsResponseOutput)
}

func (i ServerAdministratorsResponseArgs) ToServerAdministratorsResponsePtrOutput() ServerAdministratorsResponsePtrOutput {
	return i.ToServerAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (i ServerAdministratorsResponseArgs) ToServerAdministratorsResponsePtrOutputWithContext(ctx context.Context) ServerAdministratorsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsResponseOutput).ToServerAdministratorsResponsePtrOutputWithContext(ctx)
}









type ServerAdministratorsResponsePtrInput interface {
	pulumi.Input

	ToServerAdministratorsResponsePtrOutput() ServerAdministratorsResponsePtrOutput
	ToServerAdministratorsResponsePtrOutputWithContext(context.Context) ServerAdministratorsResponsePtrOutput
}

type serverAdministratorsResponsePtrType ServerAdministratorsResponseArgs

func ServerAdministratorsResponsePtr(v *ServerAdministratorsResponseArgs) ServerAdministratorsResponsePtrInput {
	return (*serverAdministratorsResponsePtrType)(v)
}

func (*serverAdministratorsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdministratorsResponse)(nil)).Elem()
}

func (i *serverAdministratorsResponsePtrType) ToServerAdministratorsResponsePtrOutput() ServerAdministratorsResponsePtrOutput {
	return i.ToServerAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (i *serverAdministratorsResponsePtrType) ToServerAdministratorsResponsePtrOutputWithContext(ctx context.Context) ServerAdministratorsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsResponsePtrOutput)
}

type ServerAdministratorsResponseOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministratorsResponse)(nil)).Elem()
}

func (o ServerAdministratorsResponseOutput) ToServerAdministratorsResponseOutput() ServerAdministratorsResponseOutput {
	return o
}

func (o ServerAdministratorsResponseOutput) ToServerAdministratorsResponseOutputWithContext(ctx context.Context) ServerAdministratorsResponseOutput {
	return o
}

func (o ServerAdministratorsResponseOutput) ToServerAdministratorsResponsePtrOutput() ServerAdministratorsResponsePtrOutput {
	return o.ToServerAdministratorsResponsePtrOutputWithContext(context.Background())
}

func (o ServerAdministratorsResponseOutput) ToServerAdministratorsResponsePtrOutputWithContext(ctx context.Context) ServerAdministratorsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerAdministratorsResponse) *ServerAdministratorsResponse {
		return &v
	}).(ServerAdministratorsResponsePtrOutput)
}

func (o ServerAdministratorsResponseOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServerAdministratorsResponse) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type ServerAdministratorsResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdministratorsResponse)(nil)).Elem()
}

func (o ServerAdministratorsResponsePtrOutput) ToServerAdministratorsResponsePtrOutput() ServerAdministratorsResponsePtrOutput {
	return o
}

func (o ServerAdministratorsResponsePtrOutput) ToServerAdministratorsResponsePtrOutputWithContext(ctx context.Context) ServerAdministratorsResponsePtrOutput {
	return o
}

func (o ServerAdministratorsResponsePtrOutput) Elem() ServerAdministratorsResponseOutput {
	return o.ApplyT(func(v *ServerAdministratorsResponse) ServerAdministratorsResponse {
		if v != nil {
			return *v
		}
		var ret ServerAdministratorsResponse
		return ret
	}).(ServerAdministratorsResponseOutput)
}

func (o ServerAdministratorsResponsePtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerAdministratorsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayDetailsInput)(nil)).Elem(), GatewayDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayDetailsPtrInput)(nil)).Elem(), GatewayDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayDetailsResponseInput)(nil)).Elem(), GatewayDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayDetailsResponsePtrInput)(nil)).Elem(), GatewayDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPv4FirewallRuleInput)(nil)).Elem(), IPv4FirewallRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPv4FirewallRuleArrayInput)(nil)).Elem(), IPv4FirewallRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPv4FirewallRuleResponseInput)(nil)).Elem(), IPv4FirewallRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPv4FirewallRuleResponseArrayInput)(nil)).Elem(), IPv4FirewallRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPv4FirewallSettingsInput)(nil)).Elem(), IPv4FirewallSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPv4FirewallSettingsPtrInput)(nil)).Elem(), IPv4FirewallSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPv4FirewallSettingsResponseInput)(nil)).Elem(), IPv4FirewallSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPv4FirewallSettingsResponsePtrInput)(nil)).Elem(), IPv4FirewallSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuInput)(nil)).Elem(), ResourceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuPtrInput)(nil)).Elem(), ResourceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuResponseInput)(nil)).Elem(), ResourceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSkuResponsePtrInput)(nil)).Elem(), ResourceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerAdministratorsInput)(nil)).Elem(), ServerAdministratorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerAdministratorsPtrInput)(nil)).Elem(), ServerAdministratorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerAdministratorsResponseInput)(nil)).Elem(), ServerAdministratorsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerAdministratorsResponsePtrInput)(nil)).Elem(), ServerAdministratorsResponseArgs{})
	pulumi.RegisterOutputType(GatewayDetailsOutput{})
	pulumi.RegisterOutputType(GatewayDetailsPtrOutput{})
	pulumi.RegisterOutputType(GatewayDetailsResponseOutput{})
	pulumi.RegisterOutputType(GatewayDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(IPv4FirewallRuleOutput{})
	pulumi.RegisterOutputType(IPv4FirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(IPv4FirewallRuleResponseOutput{})
	pulumi.RegisterOutputType(IPv4FirewallRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(IPv4FirewallSettingsOutput{})
	pulumi.RegisterOutputType(IPv4FirewallSettingsPtrOutput{})
	pulumi.RegisterOutputType(IPv4FirewallSettingsResponseOutput{})
	pulumi.RegisterOutputType(IPv4FirewallSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsResponsePtrOutput{})
}
