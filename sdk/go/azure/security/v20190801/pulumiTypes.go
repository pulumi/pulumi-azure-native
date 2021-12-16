


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdditionalWorkspacesProperties struct {
	DataTypes []string `pulumi:"dataTypes"`
	Type      *string  `pulumi:"type"`
	Workspace *string  `pulumi:"workspace"`
}


func (val *AdditionalWorkspacesProperties) Defaults() *AdditionalWorkspacesProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "Sentinel"
		tmp.Type = &type_
	}
	return &tmp
}





type AdditionalWorkspacesPropertiesInput interface {
	pulumi.Input

	ToAdditionalWorkspacesPropertiesOutput() AdditionalWorkspacesPropertiesOutput
	ToAdditionalWorkspacesPropertiesOutputWithContext(context.Context) AdditionalWorkspacesPropertiesOutput
}

type AdditionalWorkspacesPropertiesArgs struct {
	DataTypes pulumi.StringArrayInput `pulumi:"dataTypes"`
	Type      pulumi.StringPtrInput   `pulumi:"type"`
	Workspace pulumi.StringPtrInput   `pulumi:"workspace"`
}

func (AdditionalWorkspacesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspacesProperties)(nil)).Elem()
}

func (i AdditionalWorkspacesPropertiesArgs) ToAdditionalWorkspacesPropertiesOutput() AdditionalWorkspacesPropertiesOutput {
	return i.ToAdditionalWorkspacesPropertiesOutputWithContext(context.Background())
}

func (i AdditionalWorkspacesPropertiesArgs) ToAdditionalWorkspacesPropertiesOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalWorkspacesPropertiesOutput)
}





type AdditionalWorkspacesPropertiesArrayInput interface {
	pulumi.Input

	ToAdditionalWorkspacesPropertiesArrayOutput() AdditionalWorkspacesPropertiesArrayOutput
	ToAdditionalWorkspacesPropertiesArrayOutputWithContext(context.Context) AdditionalWorkspacesPropertiesArrayOutput
}

type AdditionalWorkspacesPropertiesArray []AdditionalWorkspacesPropertiesInput

func (AdditionalWorkspacesPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalWorkspacesProperties)(nil)).Elem()
}

func (i AdditionalWorkspacesPropertiesArray) ToAdditionalWorkspacesPropertiesArrayOutput() AdditionalWorkspacesPropertiesArrayOutput {
	return i.ToAdditionalWorkspacesPropertiesArrayOutputWithContext(context.Background())
}

func (i AdditionalWorkspacesPropertiesArray) ToAdditionalWorkspacesPropertiesArrayOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalWorkspacesPropertiesArrayOutput)
}

type AdditionalWorkspacesPropertiesOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspacesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspacesProperties)(nil)).Elem()
}

func (o AdditionalWorkspacesPropertiesOutput) ToAdditionalWorkspacesPropertiesOutput() AdditionalWorkspacesPropertiesOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesOutput) ToAdditionalWorkspacesPropertiesOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesOutput) DataTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdditionalWorkspacesProperties) []string { return v.DataTypes }).(pulumi.StringArrayOutput)
}

func (o AdditionalWorkspacesPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalWorkspacesProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o AdditionalWorkspacesPropertiesOutput) Workspace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalWorkspacesProperties) *string { return v.Workspace }).(pulumi.StringPtrOutput)
}

type AdditionalWorkspacesPropertiesArrayOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspacesPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalWorkspacesProperties)(nil)).Elem()
}

func (o AdditionalWorkspacesPropertiesArrayOutput) ToAdditionalWorkspacesPropertiesArrayOutput() AdditionalWorkspacesPropertiesArrayOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesArrayOutput) ToAdditionalWorkspacesPropertiesArrayOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesArrayOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesArrayOutput) Index(i pulumi.IntInput) AdditionalWorkspacesPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalWorkspacesProperties {
		return vs[0].([]AdditionalWorkspacesProperties)[vs[1].(int)]
	}).(AdditionalWorkspacesPropertiesOutput)
}

type AdditionalWorkspacesPropertiesResponse struct {
	DataTypes []string `pulumi:"dataTypes"`
	Type      *string  `pulumi:"type"`
	Workspace *string  `pulumi:"workspace"`
}


func (val *AdditionalWorkspacesPropertiesResponse) Defaults() *AdditionalWorkspacesPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "Sentinel"
		tmp.Type = &type_
	}
	return &tmp
}





type AdditionalWorkspacesPropertiesResponseInput interface {
	pulumi.Input

	ToAdditionalWorkspacesPropertiesResponseOutput() AdditionalWorkspacesPropertiesResponseOutput
	ToAdditionalWorkspacesPropertiesResponseOutputWithContext(context.Context) AdditionalWorkspacesPropertiesResponseOutput
}

type AdditionalWorkspacesPropertiesResponseArgs struct {
	DataTypes pulumi.StringArrayInput `pulumi:"dataTypes"`
	Type      pulumi.StringPtrInput   `pulumi:"type"`
	Workspace pulumi.StringPtrInput   `pulumi:"workspace"`
}

func (AdditionalWorkspacesPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspacesPropertiesResponse)(nil)).Elem()
}

func (i AdditionalWorkspacesPropertiesResponseArgs) ToAdditionalWorkspacesPropertiesResponseOutput() AdditionalWorkspacesPropertiesResponseOutput {
	return i.ToAdditionalWorkspacesPropertiesResponseOutputWithContext(context.Background())
}

func (i AdditionalWorkspacesPropertiesResponseArgs) ToAdditionalWorkspacesPropertiesResponseOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalWorkspacesPropertiesResponseOutput)
}





type AdditionalWorkspacesPropertiesResponseArrayInput interface {
	pulumi.Input

	ToAdditionalWorkspacesPropertiesResponseArrayOutput() AdditionalWorkspacesPropertiesResponseArrayOutput
	ToAdditionalWorkspacesPropertiesResponseArrayOutputWithContext(context.Context) AdditionalWorkspacesPropertiesResponseArrayOutput
}

type AdditionalWorkspacesPropertiesResponseArray []AdditionalWorkspacesPropertiesResponseInput

func (AdditionalWorkspacesPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalWorkspacesPropertiesResponse)(nil)).Elem()
}

func (i AdditionalWorkspacesPropertiesResponseArray) ToAdditionalWorkspacesPropertiesResponseArrayOutput() AdditionalWorkspacesPropertiesResponseArrayOutput {
	return i.ToAdditionalWorkspacesPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i AdditionalWorkspacesPropertiesResponseArray) ToAdditionalWorkspacesPropertiesResponseArrayOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalWorkspacesPropertiesResponseArrayOutput)
}

type AdditionalWorkspacesPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspacesPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspacesPropertiesResponse)(nil)).Elem()
}

func (o AdditionalWorkspacesPropertiesResponseOutput) ToAdditionalWorkspacesPropertiesResponseOutput() AdditionalWorkspacesPropertiesResponseOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesResponseOutput) ToAdditionalWorkspacesPropertiesResponseOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesResponseOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesResponseOutput) DataTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdditionalWorkspacesPropertiesResponse) []string { return v.DataTypes }).(pulumi.StringArrayOutput)
}

func (o AdditionalWorkspacesPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalWorkspacesPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o AdditionalWorkspacesPropertiesResponseOutput) Workspace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalWorkspacesPropertiesResponse) *string { return v.Workspace }).(pulumi.StringPtrOutput)
}

type AdditionalWorkspacesPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspacesPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalWorkspacesPropertiesResponse)(nil)).Elem()
}

func (o AdditionalWorkspacesPropertiesResponseArrayOutput) ToAdditionalWorkspacesPropertiesResponseArrayOutput() AdditionalWorkspacesPropertiesResponseArrayOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesResponseArrayOutput) ToAdditionalWorkspacesPropertiesResponseArrayOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesResponseArrayOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesResponseArrayOutput) Index(i pulumi.IntInput) AdditionalWorkspacesPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalWorkspacesPropertiesResponse {
		return vs[0].([]AdditionalWorkspacesPropertiesResponse)[vs[1].(int)]
	}).(AdditionalWorkspacesPropertiesResponseOutput)
}

type AllowlistCustomAlertRule struct {
	AllowlistValues []string `pulumi:"allowlistValues"`
	IsEnabled       bool     `pulumi:"isEnabled"`
	RuleType        string   `pulumi:"ruleType"`
}





type AllowlistCustomAlertRuleInput interface {
	pulumi.Input

	ToAllowlistCustomAlertRuleOutput() AllowlistCustomAlertRuleOutput
	ToAllowlistCustomAlertRuleOutputWithContext(context.Context) AllowlistCustomAlertRuleOutput
}

type AllowlistCustomAlertRuleArgs struct {
	AllowlistValues pulumi.StringArrayInput `pulumi:"allowlistValues"`
	IsEnabled       pulumi.BoolInput        `pulumi:"isEnabled"`
	RuleType        pulumi.StringInput      `pulumi:"ruleType"`
}

func (AllowlistCustomAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowlistCustomAlertRule)(nil)).Elem()
}

func (i AllowlistCustomAlertRuleArgs) ToAllowlistCustomAlertRuleOutput() AllowlistCustomAlertRuleOutput {
	return i.ToAllowlistCustomAlertRuleOutputWithContext(context.Background())
}

func (i AllowlistCustomAlertRuleArgs) ToAllowlistCustomAlertRuleOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowlistCustomAlertRuleOutput)
}





type AllowlistCustomAlertRuleArrayInput interface {
	pulumi.Input

	ToAllowlistCustomAlertRuleArrayOutput() AllowlistCustomAlertRuleArrayOutput
	ToAllowlistCustomAlertRuleArrayOutputWithContext(context.Context) AllowlistCustomAlertRuleArrayOutput
}

type AllowlistCustomAlertRuleArray []AllowlistCustomAlertRuleInput

func (AllowlistCustomAlertRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AllowlistCustomAlertRule)(nil)).Elem()
}

func (i AllowlistCustomAlertRuleArray) ToAllowlistCustomAlertRuleArrayOutput() AllowlistCustomAlertRuleArrayOutput {
	return i.ToAllowlistCustomAlertRuleArrayOutputWithContext(context.Background())
}

func (i AllowlistCustomAlertRuleArray) ToAllowlistCustomAlertRuleArrayOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowlistCustomAlertRuleArrayOutput)
}

type AllowlistCustomAlertRuleOutput struct{ *pulumi.OutputState }

func (AllowlistCustomAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowlistCustomAlertRule)(nil)).Elem()
}

func (o AllowlistCustomAlertRuleOutput) ToAllowlistCustomAlertRuleOutput() AllowlistCustomAlertRuleOutput {
	return o
}

func (o AllowlistCustomAlertRuleOutput) ToAllowlistCustomAlertRuleOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleOutput {
	return o
}

func (o AllowlistCustomAlertRuleOutput) AllowlistValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRule) []string { return v.AllowlistValues }).(pulumi.StringArrayOutput)
}

func (o AllowlistCustomAlertRuleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRule) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o AllowlistCustomAlertRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type AllowlistCustomAlertRuleArrayOutput struct{ *pulumi.OutputState }

func (AllowlistCustomAlertRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AllowlistCustomAlertRule)(nil)).Elem()
}

func (o AllowlistCustomAlertRuleArrayOutput) ToAllowlistCustomAlertRuleArrayOutput() AllowlistCustomAlertRuleArrayOutput {
	return o
}

func (o AllowlistCustomAlertRuleArrayOutput) ToAllowlistCustomAlertRuleArrayOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleArrayOutput {
	return o
}

func (o AllowlistCustomAlertRuleArrayOutput) Index(i pulumi.IntInput) AllowlistCustomAlertRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AllowlistCustomAlertRule {
		return vs[0].([]AllowlistCustomAlertRule)[vs[1].(int)]
	}).(AllowlistCustomAlertRuleOutput)
}

type AllowlistCustomAlertRuleResponse struct {
	AllowlistValues []string `pulumi:"allowlistValues"`
	Description     string   `pulumi:"description"`
	DisplayName     string   `pulumi:"displayName"`
	IsEnabled       bool     `pulumi:"isEnabled"`
	RuleType        string   `pulumi:"ruleType"`
	ValueType       string   `pulumi:"valueType"`
}





type AllowlistCustomAlertRuleResponseInput interface {
	pulumi.Input

	ToAllowlistCustomAlertRuleResponseOutput() AllowlistCustomAlertRuleResponseOutput
	ToAllowlistCustomAlertRuleResponseOutputWithContext(context.Context) AllowlistCustomAlertRuleResponseOutput
}

type AllowlistCustomAlertRuleResponseArgs struct {
	AllowlistValues pulumi.StringArrayInput `pulumi:"allowlistValues"`
	Description     pulumi.StringInput      `pulumi:"description"`
	DisplayName     pulumi.StringInput      `pulumi:"displayName"`
	IsEnabled       pulumi.BoolInput        `pulumi:"isEnabled"`
	RuleType        pulumi.StringInput      `pulumi:"ruleType"`
	ValueType       pulumi.StringInput      `pulumi:"valueType"`
}

func (AllowlistCustomAlertRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowlistCustomAlertRuleResponse)(nil)).Elem()
}

func (i AllowlistCustomAlertRuleResponseArgs) ToAllowlistCustomAlertRuleResponseOutput() AllowlistCustomAlertRuleResponseOutput {
	return i.ToAllowlistCustomAlertRuleResponseOutputWithContext(context.Background())
}

func (i AllowlistCustomAlertRuleResponseArgs) ToAllowlistCustomAlertRuleResponseOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowlistCustomAlertRuleResponseOutput)
}





type AllowlistCustomAlertRuleResponseArrayInput interface {
	pulumi.Input

	ToAllowlistCustomAlertRuleResponseArrayOutput() AllowlistCustomAlertRuleResponseArrayOutput
	ToAllowlistCustomAlertRuleResponseArrayOutputWithContext(context.Context) AllowlistCustomAlertRuleResponseArrayOutput
}

type AllowlistCustomAlertRuleResponseArray []AllowlistCustomAlertRuleResponseInput

func (AllowlistCustomAlertRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AllowlistCustomAlertRuleResponse)(nil)).Elem()
}

func (i AllowlistCustomAlertRuleResponseArray) ToAllowlistCustomAlertRuleResponseArrayOutput() AllowlistCustomAlertRuleResponseArrayOutput {
	return i.ToAllowlistCustomAlertRuleResponseArrayOutputWithContext(context.Background())
}

func (i AllowlistCustomAlertRuleResponseArray) ToAllowlistCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowlistCustomAlertRuleResponseArrayOutput)
}

type AllowlistCustomAlertRuleResponseOutput struct{ *pulumi.OutputState }

func (AllowlistCustomAlertRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowlistCustomAlertRuleResponse)(nil)).Elem()
}

func (o AllowlistCustomAlertRuleResponseOutput) ToAllowlistCustomAlertRuleResponseOutput() AllowlistCustomAlertRuleResponseOutput {
	return o
}

func (o AllowlistCustomAlertRuleResponseOutput) ToAllowlistCustomAlertRuleResponseOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleResponseOutput {
	return o
}

func (o AllowlistCustomAlertRuleResponseOutput) AllowlistValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) []string { return v.AllowlistValues }).(pulumi.StringArrayOutput)
}

func (o AllowlistCustomAlertRuleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o AllowlistCustomAlertRuleResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o AllowlistCustomAlertRuleResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o AllowlistCustomAlertRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o AllowlistCustomAlertRuleResponseOutput) ValueType() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) string { return v.ValueType }).(pulumi.StringOutput)
}

type AllowlistCustomAlertRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (AllowlistCustomAlertRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AllowlistCustomAlertRuleResponse)(nil)).Elem()
}

func (o AllowlistCustomAlertRuleResponseArrayOutput) ToAllowlistCustomAlertRuleResponseArrayOutput() AllowlistCustomAlertRuleResponseArrayOutput {
	return o
}

func (o AllowlistCustomAlertRuleResponseArrayOutput) ToAllowlistCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleResponseArrayOutput {
	return o
}

func (o AllowlistCustomAlertRuleResponseArrayOutput) Index(i pulumi.IntInput) AllowlistCustomAlertRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AllowlistCustomAlertRuleResponse {
		return vs[0].([]AllowlistCustomAlertRuleResponse)[vs[1].(int)]
	}).(AllowlistCustomAlertRuleResponseOutput)
}

type DenylistCustomAlertRule struct {
	DenylistValues []string `pulumi:"denylistValues"`
	IsEnabled      bool     `pulumi:"isEnabled"`
	RuleType       string   `pulumi:"ruleType"`
}





type DenylistCustomAlertRuleInput interface {
	pulumi.Input

	ToDenylistCustomAlertRuleOutput() DenylistCustomAlertRuleOutput
	ToDenylistCustomAlertRuleOutputWithContext(context.Context) DenylistCustomAlertRuleOutput
}

type DenylistCustomAlertRuleArgs struct {
	DenylistValues pulumi.StringArrayInput `pulumi:"denylistValues"`
	IsEnabled      pulumi.BoolInput        `pulumi:"isEnabled"`
	RuleType       pulumi.StringInput      `pulumi:"ruleType"`
}

func (DenylistCustomAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DenylistCustomAlertRule)(nil)).Elem()
}

func (i DenylistCustomAlertRuleArgs) ToDenylistCustomAlertRuleOutput() DenylistCustomAlertRuleOutput {
	return i.ToDenylistCustomAlertRuleOutputWithContext(context.Background())
}

func (i DenylistCustomAlertRuleArgs) ToDenylistCustomAlertRuleOutputWithContext(ctx context.Context) DenylistCustomAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DenylistCustomAlertRuleOutput)
}





type DenylistCustomAlertRuleArrayInput interface {
	pulumi.Input

	ToDenylistCustomAlertRuleArrayOutput() DenylistCustomAlertRuleArrayOutput
	ToDenylistCustomAlertRuleArrayOutputWithContext(context.Context) DenylistCustomAlertRuleArrayOutput
}

type DenylistCustomAlertRuleArray []DenylistCustomAlertRuleInput

func (DenylistCustomAlertRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DenylistCustomAlertRule)(nil)).Elem()
}

func (i DenylistCustomAlertRuleArray) ToDenylistCustomAlertRuleArrayOutput() DenylistCustomAlertRuleArrayOutput {
	return i.ToDenylistCustomAlertRuleArrayOutputWithContext(context.Background())
}

func (i DenylistCustomAlertRuleArray) ToDenylistCustomAlertRuleArrayOutputWithContext(ctx context.Context) DenylistCustomAlertRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DenylistCustomAlertRuleArrayOutput)
}

type DenylistCustomAlertRuleOutput struct{ *pulumi.OutputState }

func (DenylistCustomAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DenylistCustomAlertRule)(nil)).Elem()
}

func (o DenylistCustomAlertRuleOutput) ToDenylistCustomAlertRuleOutput() DenylistCustomAlertRuleOutput {
	return o
}

func (o DenylistCustomAlertRuleOutput) ToDenylistCustomAlertRuleOutputWithContext(ctx context.Context) DenylistCustomAlertRuleOutput {
	return o
}

func (o DenylistCustomAlertRuleOutput) DenylistValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DenylistCustomAlertRule) []string { return v.DenylistValues }).(pulumi.StringArrayOutput)
}

func (o DenylistCustomAlertRuleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v DenylistCustomAlertRule) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o DenylistCustomAlertRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v DenylistCustomAlertRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type DenylistCustomAlertRuleArrayOutput struct{ *pulumi.OutputState }

func (DenylistCustomAlertRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DenylistCustomAlertRule)(nil)).Elem()
}

func (o DenylistCustomAlertRuleArrayOutput) ToDenylistCustomAlertRuleArrayOutput() DenylistCustomAlertRuleArrayOutput {
	return o
}

func (o DenylistCustomAlertRuleArrayOutput) ToDenylistCustomAlertRuleArrayOutputWithContext(ctx context.Context) DenylistCustomAlertRuleArrayOutput {
	return o
}

func (o DenylistCustomAlertRuleArrayOutput) Index(i pulumi.IntInput) DenylistCustomAlertRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DenylistCustomAlertRule {
		return vs[0].([]DenylistCustomAlertRule)[vs[1].(int)]
	}).(DenylistCustomAlertRuleOutput)
}

type DenylistCustomAlertRuleResponse struct {
	DenylistValues []string `pulumi:"denylistValues"`
	Description    string   `pulumi:"description"`
	DisplayName    string   `pulumi:"displayName"`
	IsEnabled      bool     `pulumi:"isEnabled"`
	RuleType       string   `pulumi:"ruleType"`
	ValueType      string   `pulumi:"valueType"`
}





type DenylistCustomAlertRuleResponseInput interface {
	pulumi.Input

	ToDenylistCustomAlertRuleResponseOutput() DenylistCustomAlertRuleResponseOutput
	ToDenylistCustomAlertRuleResponseOutputWithContext(context.Context) DenylistCustomAlertRuleResponseOutput
}

type DenylistCustomAlertRuleResponseArgs struct {
	DenylistValues pulumi.StringArrayInput `pulumi:"denylistValues"`
	Description    pulumi.StringInput      `pulumi:"description"`
	DisplayName    pulumi.StringInput      `pulumi:"displayName"`
	IsEnabled      pulumi.BoolInput        `pulumi:"isEnabled"`
	RuleType       pulumi.StringInput      `pulumi:"ruleType"`
	ValueType      pulumi.StringInput      `pulumi:"valueType"`
}

func (DenylistCustomAlertRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DenylistCustomAlertRuleResponse)(nil)).Elem()
}

func (i DenylistCustomAlertRuleResponseArgs) ToDenylistCustomAlertRuleResponseOutput() DenylistCustomAlertRuleResponseOutput {
	return i.ToDenylistCustomAlertRuleResponseOutputWithContext(context.Background())
}

func (i DenylistCustomAlertRuleResponseArgs) ToDenylistCustomAlertRuleResponseOutputWithContext(ctx context.Context) DenylistCustomAlertRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DenylistCustomAlertRuleResponseOutput)
}





type DenylistCustomAlertRuleResponseArrayInput interface {
	pulumi.Input

	ToDenylistCustomAlertRuleResponseArrayOutput() DenylistCustomAlertRuleResponseArrayOutput
	ToDenylistCustomAlertRuleResponseArrayOutputWithContext(context.Context) DenylistCustomAlertRuleResponseArrayOutput
}

type DenylistCustomAlertRuleResponseArray []DenylistCustomAlertRuleResponseInput

func (DenylistCustomAlertRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DenylistCustomAlertRuleResponse)(nil)).Elem()
}

func (i DenylistCustomAlertRuleResponseArray) ToDenylistCustomAlertRuleResponseArrayOutput() DenylistCustomAlertRuleResponseArrayOutput {
	return i.ToDenylistCustomAlertRuleResponseArrayOutputWithContext(context.Background())
}

func (i DenylistCustomAlertRuleResponseArray) ToDenylistCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) DenylistCustomAlertRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DenylistCustomAlertRuleResponseArrayOutput)
}

type DenylistCustomAlertRuleResponseOutput struct{ *pulumi.OutputState }

func (DenylistCustomAlertRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DenylistCustomAlertRuleResponse)(nil)).Elem()
}

func (o DenylistCustomAlertRuleResponseOutput) ToDenylistCustomAlertRuleResponseOutput() DenylistCustomAlertRuleResponseOutput {
	return o
}

func (o DenylistCustomAlertRuleResponseOutput) ToDenylistCustomAlertRuleResponseOutputWithContext(ctx context.Context) DenylistCustomAlertRuleResponseOutput {
	return o
}

func (o DenylistCustomAlertRuleResponseOutput) DenylistValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) []string { return v.DenylistValues }).(pulumi.StringArrayOutput)
}

func (o DenylistCustomAlertRuleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o DenylistCustomAlertRuleResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o DenylistCustomAlertRuleResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o DenylistCustomAlertRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o DenylistCustomAlertRuleResponseOutput) ValueType() pulumi.StringOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) string { return v.ValueType }).(pulumi.StringOutput)
}

type DenylistCustomAlertRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (DenylistCustomAlertRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DenylistCustomAlertRuleResponse)(nil)).Elem()
}

func (o DenylistCustomAlertRuleResponseArrayOutput) ToDenylistCustomAlertRuleResponseArrayOutput() DenylistCustomAlertRuleResponseArrayOutput {
	return o
}

func (o DenylistCustomAlertRuleResponseArrayOutput) ToDenylistCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) DenylistCustomAlertRuleResponseArrayOutput {
	return o
}

func (o DenylistCustomAlertRuleResponseArrayOutput) Index(i pulumi.IntInput) DenylistCustomAlertRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DenylistCustomAlertRuleResponse {
		return vs[0].([]DenylistCustomAlertRuleResponse)[vs[1].(int)]
	}).(DenylistCustomAlertRuleResponseOutput)
}

type RecommendationConfigurationProperties struct {
	RecommendationType string `pulumi:"recommendationType"`
	Status             string `pulumi:"status"`
}


func (val *RecommendationConfigurationProperties) Defaults() *RecommendationConfigurationProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		tmp.Status = "Enabled"
	}
	return &tmp
}





type RecommendationConfigurationPropertiesInput interface {
	pulumi.Input

	ToRecommendationConfigurationPropertiesOutput() RecommendationConfigurationPropertiesOutput
	ToRecommendationConfigurationPropertiesOutputWithContext(context.Context) RecommendationConfigurationPropertiesOutput
}

type RecommendationConfigurationPropertiesArgs struct {
	RecommendationType pulumi.StringInput `pulumi:"recommendationType"`
	Status             pulumi.StringInput `pulumi:"status"`
}

func (RecommendationConfigurationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationConfigurationProperties)(nil)).Elem()
}

func (i RecommendationConfigurationPropertiesArgs) ToRecommendationConfigurationPropertiesOutput() RecommendationConfigurationPropertiesOutput {
	return i.ToRecommendationConfigurationPropertiesOutputWithContext(context.Background())
}

func (i RecommendationConfigurationPropertiesArgs) ToRecommendationConfigurationPropertiesOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendationConfigurationPropertiesOutput)
}





type RecommendationConfigurationPropertiesArrayInput interface {
	pulumi.Input

	ToRecommendationConfigurationPropertiesArrayOutput() RecommendationConfigurationPropertiesArrayOutput
	ToRecommendationConfigurationPropertiesArrayOutputWithContext(context.Context) RecommendationConfigurationPropertiesArrayOutput
}

type RecommendationConfigurationPropertiesArray []RecommendationConfigurationPropertiesInput

func (RecommendationConfigurationPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendationConfigurationProperties)(nil)).Elem()
}

func (i RecommendationConfigurationPropertiesArray) ToRecommendationConfigurationPropertiesArrayOutput() RecommendationConfigurationPropertiesArrayOutput {
	return i.ToRecommendationConfigurationPropertiesArrayOutputWithContext(context.Background())
}

func (i RecommendationConfigurationPropertiesArray) ToRecommendationConfigurationPropertiesArrayOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendationConfigurationPropertiesArrayOutput)
}

type RecommendationConfigurationPropertiesOutput struct{ *pulumi.OutputState }

func (RecommendationConfigurationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationConfigurationProperties)(nil)).Elem()
}

func (o RecommendationConfigurationPropertiesOutput) ToRecommendationConfigurationPropertiesOutput() RecommendationConfigurationPropertiesOutput {
	return o
}

func (o RecommendationConfigurationPropertiesOutput) ToRecommendationConfigurationPropertiesOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesOutput {
	return o
}

func (o RecommendationConfigurationPropertiesOutput) RecommendationType() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendationConfigurationProperties) string { return v.RecommendationType }).(pulumi.StringOutput)
}

func (o RecommendationConfigurationPropertiesOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendationConfigurationProperties) string { return v.Status }).(pulumi.StringOutput)
}

type RecommendationConfigurationPropertiesArrayOutput struct{ *pulumi.OutputState }

func (RecommendationConfigurationPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendationConfigurationProperties)(nil)).Elem()
}

func (o RecommendationConfigurationPropertiesArrayOutput) ToRecommendationConfigurationPropertiesArrayOutput() RecommendationConfigurationPropertiesArrayOutput {
	return o
}

func (o RecommendationConfigurationPropertiesArrayOutput) ToRecommendationConfigurationPropertiesArrayOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesArrayOutput {
	return o
}

func (o RecommendationConfigurationPropertiesArrayOutput) Index(i pulumi.IntInput) RecommendationConfigurationPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecommendationConfigurationProperties {
		return vs[0].([]RecommendationConfigurationProperties)[vs[1].(int)]
	}).(RecommendationConfigurationPropertiesOutput)
}

type RecommendationConfigurationPropertiesResponse struct {
	Name               string `pulumi:"name"`
	RecommendationType string `pulumi:"recommendationType"`
	Status             string `pulumi:"status"`
}


func (val *RecommendationConfigurationPropertiesResponse) Defaults() *RecommendationConfigurationPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		tmp.Status = "Enabled"
	}
	return &tmp
}





type RecommendationConfigurationPropertiesResponseInput interface {
	pulumi.Input

	ToRecommendationConfigurationPropertiesResponseOutput() RecommendationConfigurationPropertiesResponseOutput
	ToRecommendationConfigurationPropertiesResponseOutputWithContext(context.Context) RecommendationConfigurationPropertiesResponseOutput
}

type RecommendationConfigurationPropertiesResponseArgs struct {
	Name               pulumi.StringInput `pulumi:"name"`
	RecommendationType pulumi.StringInput `pulumi:"recommendationType"`
	Status             pulumi.StringInput `pulumi:"status"`
}

func (RecommendationConfigurationPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationConfigurationPropertiesResponse)(nil)).Elem()
}

func (i RecommendationConfigurationPropertiesResponseArgs) ToRecommendationConfigurationPropertiesResponseOutput() RecommendationConfigurationPropertiesResponseOutput {
	return i.ToRecommendationConfigurationPropertiesResponseOutputWithContext(context.Background())
}

func (i RecommendationConfigurationPropertiesResponseArgs) ToRecommendationConfigurationPropertiesResponseOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendationConfigurationPropertiesResponseOutput)
}





type RecommendationConfigurationPropertiesResponseArrayInput interface {
	pulumi.Input

	ToRecommendationConfigurationPropertiesResponseArrayOutput() RecommendationConfigurationPropertiesResponseArrayOutput
	ToRecommendationConfigurationPropertiesResponseArrayOutputWithContext(context.Context) RecommendationConfigurationPropertiesResponseArrayOutput
}

type RecommendationConfigurationPropertiesResponseArray []RecommendationConfigurationPropertiesResponseInput

func (RecommendationConfigurationPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendationConfigurationPropertiesResponse)(nil)).Elem()
}

func (i RecommendationConfigurationPropertiesResponseArray) ToRecommendationConfigurationPropertiesResponseArrayOutput() RecommendationConfigurationPropertiesResponseArrayOutput {
	return i.ToRecommendationConfigurationPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i RecommendationConfigurationPropertiesResponseArray) ToRecommendationConfigurationPropertiesResponseArrayOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendationConfigurationPropertiesResponseArrayOutput)
}

type RecommendationConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RecommendationConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationConfigurationPropertiesResponse)(nil)).Elem()
}

func (o RecommendationConfigurationPropertiesResponseOutput) ToRecommendationConfigurationPropertiesResponseOutput() RecommendationConfigurationPropertiesResponseOutput {
	return o
}

func (o RecommendationConfigurationPropertiesResponseOutput) ToRecommendationConfigurationPropertiesResponseOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesResponseOutput {
	return o
}

func (o RecommendationConfigurationPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendationConfigurationPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RecommendationConfigurationPropertiesResponseOutput) RecommendationType() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendationConfigurationPropertiesResponse) string { return v.RecommendationType }).(pulumi.StringOutput)
}

func (o RecommendationConfigurationPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendationConfigurationPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

type RecommendationConfigurationPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RecommendationConfigurationPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendationConfigurationPropertiesResponse)(nil)).Elem()
}

func (o RecommendationConfigurationPropertiesResponseArrayOutput) ToRecommendationConfigurationPropertiesResponseArrayOutput() RecommendationConfigurationPropertiesResponseArrayOutput {
	return o
}

func (o RecommendationConfigurationPropertiesResponseArrayOutput) ToRecommendationConfigurationPropertiesResponseArrayOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesResponseArrayOutput {
	return o
}

func (o RecommendationConfigurationPropertiesResponseArrayOutput) Index(i pulumi.IntInput) RecommendationConfigurationPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecommendationConfigurationPropertiesResponse {
		return vs[0].([]RecommendationConfigurationPropertiesResponse)[vs[1].(int)]
	}).(RecommendationConfigurationPropertiesResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type ThresholdCustomAlertRule struct {
	IsEnabled    bool   `pulumi:"isEnabled"`
	MaxThreshold int    `pulumi:"maxThreshold"`
	MinThreshold int    `pulumi:"minThreshold"`
	RuleType     string `pulumi:"ruleType"`
}





type ThresholdCustomAlertRuleInput interface {
	pulumi.Input

	ToThresholdCustomAlertRuleOutput() ThresholdCustomAlertRuleOutput
	ToThresholdCustomAlertRuleOutputWithContext(context.Context) ThresholdCustomAlertRuleOutput
}

type ThresholdCustomAlertRuleArgs struct {
	IsEnabled    pulumi.BoolInput   `pulumi:"isEnabled"`
	MaxThreshold pulumi.IntInput    `pulumi:"maxThreshold"`
	MinThreshold pulumi.IntInput    `pulumi:"minThreshold"`
	RuleType     pulumi.StringInput `pulumi:"ruleType"`
}

func (ThresholdCustomAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdCustomAlertRule)(nil)).Elem()
}

func (i ThresholdCustomAlertRuleArgs) ToThresholdCustomAlertRuleOutput() ThresholdCustomAlertRuleOutput {
	return i.ToThresholdCustomAlertRuleOutputWithContext(context.Background())
}

func (i ThresholdCustomAlertRuleArgs) ToThresholdCustomAlertRuleOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdCustomAlertRuleOutput)
}





type ThresholdCustomAlertRuleArrayInput interface {
	pulumi.Input

	ToThresholdCustomAlertRuleArrayOutput() ThresholdCustomAlertRuleArrayOutput
	ToThresholdCustomAlertRuleArrayOutputWithContext(context.Context) ThresholdCustomAlertRuleArrayOutput
}

type ThresholdCustomAlertRuleArray []ThresholdCustomAlertRuleInput

func (ThresholdCustomAlertRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThresholdCustomAlertRule)(nil)).Elem()
}

func (i ThresholdCustomAlertRuleArray) ToThresholdCustomAlertRuleArrayOutput() ThresholdCustomAlertRuleArrayOutput {
	return i.ToThresholdCustomAlertRuleArrayOutputWithContext(context.Background())
}

func (i ThresholdCustomAlertRuleArray) ToThresholdCustomAlertRuleArrayOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdCustomAlertRuleArrayOutput)
}

type ThresholdCustomAlertRuleOutput struct{ *pulumi.OutputState }

func (ThresholdCustomAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdCustomAlertRule)(nil)).Elem()
}

func (o ThresholdCustomAlertRuleOutput) ToThresholdCustomAlertRuleOutput() ThresholdCustomAlertRuleOutput {
	return o
}

func (o ThresholdCustomAlertRuleOutput) ToThresholdCustomAlertRuleOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleOutput {
	return o
}

func (o ThresholdCustomAlertRuleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRule) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o ThresholdCustomAlertRuleOutput) MaxThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRule) int { return v.MaxThreshold }).(pulumi.IntOutput)
}

func (o ThresholdCustomAlertRuleOutput) MinThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRule) int { return v.MinThreshold }).(pulumi.IntOutput)
}

func (o ThresholdCustomAlertRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type ThresholdCustomAlertRuleArrayOutput struct{ *pulumi.OutputState }

func (ThresholdCustomAlertRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThresholdCustomAlertRule)(nil)).Elem()
}

func (o ThresholdCustomAlertRuleArrayOutput) ToThresholdCustomAlertRuleArrayOutput() ThresholdCustomAlertRuleArrayOutput {
	return o
}

func (o ThresholdCustomAlertRuleArrayOutput) ToThresholdCustomAlertRuleArrayOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleArrayOutput {
	return o
}

func (o ThresholdCustomAlertRuleArrayOutput) Index(i pulumi.IntInput) ThresholdCustomAlertRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThresholdCustomAlertRule {
		return vs[0].([]ThresholdCustomAlertRule)[vs[1].(int)]
	}).(ThresholdCustomAlertRuleOutput)
}

type ThresholdCustomAlertRuleResponse struct {
	Description  string `pulumi:"description"`
	DisplayName  string `pulumi:"displayName"`
	IsEnabled    bool   `pulumi:"isEnabled"`
	MaxThreshold int    `pulumi:"maxThreshold"`
	MinThreshold int    `pulumi:"minThreshold"`
	RuleType     string `pulumi:"ruleType"`
}





type ThresholdCustomAlertRuleResponseInput interface {
	pulumi.Input

	ToThresholdCustomAlertRuleResponseOutput() ThresholdCustomAlertRuleResponseOutput
	ToThresholdCustomAlertRuleResponseOutputWithContext(context.Context) ThresholdCustomAlertRuleResponseOutput
}

type ThresholdCustomAlertRuleResponseArgs struct {
	Description  pulumi.StringInput `pulumi:"description"`
	DisplayName  pulumi.StringInput `pulumi:"displayName"`
	IsEnabled    pulumi.BoolInput   `pulumi:"isEnabled"`
	MaxThreshold pulumi.IntInput    `pulumi:"maxThreshold"`
	MinThreshold pulumi.IntInput    `pulumi:"minThreshold"`
	RuleType     pulumi.StringInput `pulumi:"ruleType"`
}

func (ThresholdCustomAlertRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdCustomAlertRuleResponse)(nil)).Elem()
}

func (i ThresholdCustomAlertRuleResponseArgs) ToThresholdCustomAlertRuleResponseOutput() ThresholdCustomAlertRuleResponseOutput {
	return i.ToThresholdCustomAlertRuleResponseOutputWithContext(context.Background())
}

func (i ThresholdCustomAlertRuleResponseArgs) ToThresholdCustomAlertRuleResponseOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdCustomAlertRuleResponseOutput)
}





type ThresholdCustomAlertRuleResponseArrayInput interface {
	pulumi.Input

	ToThresholdCustomAlertRuleResponseArrayOutput() ThresholdCustomAlertRuleResponseArrayOutput
	ToThresholdCustomAlertRuleResponseArrayOutputWithContext(context.Context) ThresholdCustomAlertRuleResponseArrayOutput
}

type ThresholdCustomAlertRuleResponseArray []ThresholdCustomAlertRuleResponseInput

func (ThresholdCustomAlertRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThresholdCustomAlertRuleResponse)(nil)).Elem()
}

func (i ThresholdCustomAlertRuleResponseArray) ToThresholdCustomAlertRuleResponseArrayOutput() ThresholdCustomAlertRuleResponseArrayOutput {
	return i.ToThresholdCustomAlertRuleResponseArrayOutputWithContext(context.Background())
}

func (i ThresholdCustomAlertRuleResponseArray) ToThresholdCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdCustomAlertRuleResponseArrayOutput)
}

type ThresholdCustomAlertRuleResponseOutput struct{ *pulumi.OutputState }

func (ThresholdCustomAlertRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdCustomAlertRuleResponse)(nil)).Elem()
}

func (o ThresholdCustomAlertRuleResponseOutput) ToThresholdCustomAlertRuleResponseOutput() ThresholdCustomAlertRuleResponseOutput {
	return o
}

func (o ThresholdCustomAlertRuleResponseOutput) ToThresholdCustomAlertRuleResponseOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleResponseOutput {
	return o
}

func (o ThresholdCustomAlertRuleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o ThresholdCustomAlertRuleResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ThresholdCustomAlertRuleResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o ThresholdCustomAlertRuleResponseOutput) MaxThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) int { return v.MaxThreshold }).(pulumi.IntOutput)
}

func (o ThresholdCustomAlertRuleResponseOutput) MinThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) int { return v.MinThreshold }).(pulumi.IntOutput)
}

func (o ThresholdCustomAlertRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

type ThresholdCustomAlertRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ThresholdCustomAlertRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThresholdCustomAlertRuleResponse)(nil)).Elem()
}

func (o ThresholdCustomAlertRuleResponseArrayOutput) ToThresholdCustomAlertRuleResponseArrayOutput() ThresholdCustomAlertRuleResponseArrayOutput {
	return o
}

func (o ThresholdCustomAlertRuleResponseArrayOutput) ToThresholdCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleResponseArrayOutput {
	return o
}

func (o ThresholdCustomAlertRuleResponseArrayOutput) Index(i pulumi.IntInput) ThresholdCustomAlertRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThresholdCustomAlertRuleResponse {
		return vs[0].([]ThresholdCustomAlertRuleResponse)[vs[1].(int)]
	}).(ThresholdCustomAlertRuleResponseOutput)
}

type TimeWindowCustomAlertRule struct {
	IsEnabled      bool   `pulumi:"isEnabled"`
	MaxThreshold   int    `pulumi:"maxThreshold"`
	MinThreshold   int    `pulumi:"minThreshold"`
	RuleType       string `pulumi:"ruleType"`
	TimeWindowSize string `pulumi:"timeWindowSize"`
}





type TimeWindowCustomAlertRuleInput interface {
	pulumi.Input

	ToTimeWindowCustomAlertRuleOutput() TimeWindowCustomAlertRuleOutput
	ToTimeWindowCustomAlertRuleOutputWithContext(context.Context) TimeWindowCustomAlertRuleOutput
}

type TimeWindowCustomAlertRuleArgs struct {
	IsEnabled      pulumi.BoolInput   `pulumi:"isEnabled"`
	MaxThreshold   pulumi.IntInput    `pulumi:"maxThreshold"`
	MinThreshold   pulumi.IntInput    `pulumi:"minThreshold"`
	RuleType       pulumi.StringInput `pulumi:"ruleType"`
	TimeWindowSize pulumi.StringInput `pulumi:"timeWindowSize"`
}

func (TimeWindowCustomAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindowCustomAlertRule)(nil)).Elem()
}

func (i TimeWindowCustomAlertRuleArgs) ToTimeWindowCustomAlertRuleOutput() TimeWindowCustomAlertRuleOutput {
	return i.ToTimeWindowCustomAlertRuleOutputWithContext(context.Background())
}

func (i TimeWindowCustomAlertRuleArgs) ToTimeWindowCustomAlertRuleOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowCustomAlertRuleOutput)
}





type TimeWindowCustomAlertRuleArrayInput interface {
	pulumi.Input

	ToTimeWindowCustomAlertRuleArrayOutput() TimeWindowCustomAlertRuleArrayOutput
	ToTimeWindowCustomAlertRuleArrayOutputWithContext(context.Context) TimeWindowCustomAlertRuleArrayOutput
}

type TimeWindowCustomAlertRuleArray []TimeWindowCustomAlertRuleInput

func (TimeWindowCustomAlertRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeWindowCustomAlertRule)(nil)).Elem()
}

func (i TimeWindowCustomAlertRuleArray) ToTimeWindowCustomAlertRuleArrayOutput() TimeWindowCustomAlertRuleArrayOutput {
	return i.ToTimeWindowCustomAlertRuleArrayOutputWithContext(context.Background())
}

func (i TimeWindowCustomAlertRuleArray) ToTimeWindowCustomAlertRuleArrayOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowCustomAlertRuleArrayOutput)
}

type TimeWindowCustomAlertRuleOutput struct{ *pulumi.OutputState }

func (TimeWindowCustomAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindowCustomAlertRule)(nil)).Elem()
}

func (o TimeWindowCustomAlertRuleOutput) ToTimeWindowCustomAlertRuleOutput() TimeWindowCustomAlertRuleOutput {
	return o
}

func (o TimeWindowCustomAlertRuleOutput) ToTimeWindowCustomAlertRuleOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleOutput {
	return o
}

func (o TimeWindowCustomAlertRuleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRule) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o TimeWindowCustomAlertRuleOutput) MaxThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRule) int { return v.MaxThreshold }).(pulumi.IntOutput)
}

func (o TimeWindowCustomAlertRuleOutput) MinThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRule) int { return v.MinThreshold }).(pulumi.IntOutput)
}

func (o TimeWindowCustomAlertRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRule) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o TimeWindowCustomAlertRuleOutput) TimeWindowSize() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRule) string { return v.TimeWindowSize }).(pulumi.StringOutput)
}

type TimeWindowCustomAlertRuleArrayOutput struct{ *pulumi.OutputState }

func (TimeWindowCustomAlertRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeWindowCustomAlertRule)(nil)).Elem()
}

func (o TimeWindowCustomAlertRuleArrayOutput) ToTimeWindowCustomAlertRuleArrayOutput() TimeWindowCustomAlertRuleArrayOutput {
	return o
}

func (o TimeWindowCustomAlertRuleArrayOutput) ToTimeWindowCustomAlertRuleArrayOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleArrayOutput {
	return o
}

func (o TimeWindowCustomAlertRuleArrayOutput) Index(i pulumi.IntInput) TimeWindowCustomAlertRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeWindowCustomAlertRule {
		return vs[0].([]TimeWindowCustomAlertRule)[vs[1].(int)]
	}).(TimeWindowCustomAlertRuleOutput)
}

type TimeWindowCustomAlertRuleResponse struct {
	Description    string `pulumi:"description"`
	DisplayName    string `pulumi:"displayName"`
	IsEnabled      bool   `pulumi:"isEnabled"`
	MaxThreshold   int    `pulumi:"maxThreshold"`
	MinThreshold   int    `pulumi:"minThreshold"`
	RuleType       string `pulumi:"ruleType"`
	TimeWindowSize string `pulumi:"timeWindowSize"`
}





type TimeWindowCustomAlertRuleResponseInput interface {
	pulumi.Input

	ToTimeWindowCustomAlertRuleResponseOutput() TimeWindowCustomAlertRuleResponseOutput
	ToTimeWindowCustomAlertRuleResponseOutputWithContext(context.Context) TimeWindowCustomAlertRuleResponseOutput
}

type TimeWindowCustomAlertRuleResponseArgs struct {
	Description    pulumi.StringInput `pulumi:"description"`
	DisplayName    pulumi.StringInput `pulumi:"displayName"`
	IsEnabled      pulumi.BoolInput   `pulumi:"isEnabled"`
	MaxThreshold   pulumi.IntInput    `pulumi:"maxThreshold"`
	MinThreshold   pulumi.IntInput    `pulumi:"minThreshold"`
	RuleType       pulumi.StringInput `pulumi:"ruleType"`
	TimeWindowSize pulumi.StringInput `pulumi:"timeWindowSize"`
}

func (TimeWindowCustomAlertRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindowCustomAlertRuleResponse)(nil)).Elem()
}

func (i TimeWindowCustomAlertRuleResponseArgs) ToTimeWindowCustomAlertRuleResponseOutput() TimeWindowCustomAlertRuleResponseOutput {
	return i.ToTimeWindowCustomAlertRuleResponseOutputWithContext(context.Background())
}

func (i TimeWindowCustomAlertRuleResponseArgs) ToTimeWindowCustomAlertRuleResponseOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowCustomAlertRuleResponseOutput)
}





type TimeWindowCustomAlertRuleResponseArrayInput interface {
	pulumi.Input

	ToTimeWindowCustomAlertRuleResponseArrayOutput() TimeWindowCustomAlertRuleResponseArrayOutput
	ToTimeWindowCustomAlertRuleResponseArrayOutputWithContext(context.Context) TimeWindowCustomAlertRuleResponseArrayOutput
}

type TimeWindowCustomAlertRuleResponseArray []TimeWindowCustomAlertRuleResponseInput

func (TimeWindowCustomAlertRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeWindowCustomAlertRuleResponse)(nil)).Elem()
}

func (i TimeWindowCustomAlertRuleResponseArray) ToTimeWindowCustomAlertRuleResponseArrayOutput() TimeWindowCustomAlertRuleResponseArrayOutput {
	return i.ToTimeWindowCustomAlertRuleResponseArrayOutputWithContext(context.Background())
}

func (i TimeWindowCustomAlertRuleResponseArray) ToTimeWindowCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowCustomAlertRuleResponseArrayOutput)
}

type TimeWindowCustomAlertRuleResponseOutput struct{ *pulumi.OutputState }

func (TimeWindowCustomAlertRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindowCustomAlertRuleResponse)(nil)).Elem()
}

func (o TimeWindowCustomAlertRuleResponseOutput) ToTimeWindowCustomAlertRuleResponseOutput() TimeWindowCustomAlertRuleResponseOutput {
	return o
}

func (o TimeWindowCustomAlertRuleResponseOutput) ToTimeWindowCustomAlertRuleResponseOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleResponseOutput {
	return o
}

func (o TimeWindowCustomAlertRuleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) MaxThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) int { return v.MaxThreshold }).(pulumi.IntOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) MinThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) int { return v.MinThreshold }).(pulumi.IntOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) TimeWindowSize() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) string { return v.TimeWindowSize }).(pulumi.StringOutput)
}

type TimeWindowCustomAlertRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (TimeWindowCustomAlertRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeWindowCustomAlertRuleResponse)(nil)).Elem()
}

func (o TimeWindowCustomAlertRuleResponseArrayOutput) ToTimeWindowCustomAlertRuleResponseArrayOutput() TimeWindowCustomAlertRuleResponseArrayOutput {
	return o
}

func (o TimeWindowCustomAlertRuleResponseArrayOutput) ToTimeWindowCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleResponseArrayOutput {
	return o
}

func (o TimeWindowCustomAlertRuleResponseArrayOutput) Index(i pulumi.IntInput) TimeWindowCustomAlertRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeWindowCustomAlertRuleResponse {
		return vs[0].([]TimeWindowCustomAlertRuleResponse)[vs[1].(int)]
	}).(TimeWindowCustomAlertRuleResponseOutput)
}

type UserDefinedResourcesProperties struct {
	Query              string   `pulumi:"query"`
	QuerySubscriptions []string `pulumi:"querySubscriptions"`
}





type UserDefinedResourcesPropertiesInput interface {
	pulumi.Input

	ToUserDefinedResourcesPropertiesOutput() UserDefinedResourcesPropertiesOutput
	ToUserDefinedResourcesPropertiesOutputWithContext(context.Context) UserDefinedResourcesPropertiesOutput
}

type UserDefinedResourcesPropertiesArgs struct {
	Query              pulumi.StringInput      `pulumi:"query"`
	QuerySubscriptions pulumi.StringArrayInput `pulumi:"querySubscriptions"`
}

func (UserDefinedResourcesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDefinedResourcesProperties)(nil)).Elem()
}

func (i UserDefinedResourcesPropertiesArgs) ToUserDefinedResourcesPropertiesOutput() UserDefinedResourcesPropertiesOutput {
	return i.ToUserDefinedResourcesPropertiesOutputWithContext(context.Background())
}

func (i UserDefinedResourcesPropertiesArgs) ToUserDefinedResourcesPropertiesOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesOutput)
}

func (i UserDefinedResourcesPropertiesArgs) ToUserDefinedResourcesPropertiesPtrOutput() UserDefinedResourcesPropertiesPtrOutput {
	return i.ToUserDefinedResourcesPropertiesPtrOutputWithContext(context.Background())
}

func (i UserDefinedResourcesPropertiesArgs) ToUserDefinedResourcesPropertiesPtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesOutput).ToUserDefinedResourcesPropertiesPtrOutputWithContext(ctx)
}









type UserDefinedResourcesPropertiesPtrInput interface {
	pulumi.Input

	ToUserDefinedResourcesPropertiesPtrOutput() UserDefinedResourcesPropertiesPtrOutput
	ToUserDefinedResourcesPropertiesPtrOutputWithContext(context.Context) UserDefinedResourcesPropertiesPtrOutput
}

type userDefinedResourcesPropertiesPtrType UserDefinedResourcesPropertiesArgs

func UserDefinedResourcesPropertiesPtr(v *UserDefinedResourcesPropertiesArgs) UserDefinedResourcesPropertiesPtrInput {
	return (*userDefinedResourcesPropertiesPtrType)(v)
}

func (*userDefinedResourcesPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDefinedResourcesProperties)(nil)).Elem()
}

func (i *userDefinedResourcesPropertiesPtrType) ToUserDefinedResourcesPropertiesPtrOutput() UserDefinedResourcesPropertiesPtrOutput {
	return i.ToUserDefinedResourcesPropertiesPtrOutputWithContext(context.Background())
}

func (i *userDefinedResourcesPropertiesPtrType) ToUserDefinedResourcesPropertiesPtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesPtrOutput)
}

type UserDefinedResourcesPropertiesOutput struct{ *pulumi.OutputState }

func (UserDefinedResourcesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDefinedResourcesProperties)(nil)).Elem()
}

func (o UserDefinedResourcesPropertiesOutput) ToUserDefinedResourcesPropertiesOutput() UserDefinedResourcesPropertiesOutput {
	return o
}

func (o UserDefinedResourcesPropertiesOutput) ToUserDefinedResourcesPropertiesOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesOutput {
	return o
}

func (o UserDefinedResourcesPropertiesOutput) ToUserDefinedResourcesPropertiesPtrOutput() UserDefinedResourcesPropertiesPtrOutput {
	return o.ToUserDefinedResourcesPropertiesPtrOutputWithContext(context.Background())
}

func (o UserDefinedResourcesPropertiesOutput) ToUserDefinedResourcesPropertiesPtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserDefinedResourcesProperties) *UserDefinedResourcesProperties {
		return &v
	}).(UserDefinedResourcesPropertiesPtrOutput)
}

func (o UserDefinedResourcesPropertiesOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v UserDefinedResourcesProperties) string { return v.Query }).(pulumi.StringOutput)
}

func (o UserDefinedResourcesPropertiesOutput) QuerySubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UserDefinedResourcesProperties) []string { return v.QuerySubscriptions }).(pulumi.StringArrayOutput)
}

type UserDefinedResourcesPropertiesPtrOutput struct{ *pulumi.OutputState }

func (UserDefinedResourcesPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDefinedResourcesProperties)(nil)).Elem()
}

func (o UserDefinedResourcesPropertiesPtrOutput) ToUserDefinedResourcesPropertiesPtrOutput() UserDefinedResourcesPropertiesPtrOutput {
	return o
}

func (o UserDefinedResourcesPropertiesPtrOutput) ToUserDefinedResourcesPropertiesPtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesPtrOutput {
	return o
}

func (o UserDefinedResourcesPropertiesPtrOutput) Elem() UserDefinedResourcesPropertiesOutput {
	return o.ApplyT(func(v *UserDefinedResourcesProperties) UserDefinedResourcesProperties {
		if v != nil {
			return *v
		}
		var ret UserDefinedResourcesProperties
		return ret
	}).(UserDefinedResourcesPropertiesOutput)
}

func (o UserDefinedResourcesPropertiesPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDefinedResourcesProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Query
	}).(pulumi.StringPtrOutput)
}

func (o UserDefinedResourcesPropertiesPtrOutput) QuerySubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserDefinedResourcesProperties) []string {
		if v == nil {
			return nil
		}
		return v.QuerySubscriptions
	}).(pulumi.StringArrayOutput)
}

type UserDefinedResourcesPropertiesResponse struct {
	Query              string   `pulumi:"query"`
	QuerySubscriptions []string `pulumi:"querySubscriptions"`
}





type UserDefinedResourcesPropertiesResponseInput interface {
	pulumi.Input

	ToUserDefinedResourcesPropertiesResponseOutput() UserDefinedResourcesPropertiesResponseOutput
	ToUserDefinedResourcesPropertiesResponseOutputWithContext(context.Context) UserDefinedResourcesPropertiesResponseOutput
}

type UserDefinedResourcesPropertiesResponseArgs struct {
	Query              pulumi.StringInput      `pulumi:"query"`
	QuerySubscriptions pulumi.StringArrayInput `pulumi:"querySubscriptions"`
}

func (UserDefinedResourcesPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDefinedResourcesPropertiesResponse)(nil)).Elem()
}

func (i UserDefinedResourcesPropertiesResponseArgs) ToUserDefinedResourcesPropertiesResponseOutput() UserDefinedResourcesPropertiesResponseOutput {
	return i.ToUserDefinedResourcesPropertiesResponseOutputWithContext(context.Background())
}

func (i UserDefinedResourcesPropertiesResponseArgs) ToUserDefinedResourcesPropertiesResponseOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesResponseOutput)
}

func (i UserDefinedResourcesPropertiesResponseArgs) ToUserDefinedResourcesPropertiesResponsePtrOutput() UserDefinedResourcesPropertiesResponsePtrOutput {
	return i.ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i UserDefinedResourcesPropertiesResponseArgs) ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesResponseOutput).ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(ctx)
}









type UserDefinedResourcesPropertiesResponsePtrInput interface {
	pulumi.Input

	ToUserDefinedResourcesPropertiesResponsePtrOutput() UserDefinedResourcesPropertiesResponsePtrOutput
	ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(context.Context) UserDefinedResourcesPropertiesResponsePtrOutput
}

type userDefinedResourcesPropertiesResponsePtrType UserDefinedResourcesPropertiesResponseArgs

func UserDefinedResourcesPropertiesResponsePtr(v *UserDefinedResourcesPropertiesResponseArgs) UserDefinedResourcesPropertiesResponsePtrInput {
	return (*userDefinedResourcesPropertiesResponsePtrType)(v)
}

func (*userDefinedResourcesPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDefinedResourcesPropertiesResponse)(nil)).Elem()
}

func (i *userDefinedResourcesPropertiesResponsePtrType) ToUserDefinedResourcesPropertiesResponsePtrOutput() UserDefinedResourcesPropertiesResponsePtrOutput {
	return i.ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *userDefinedResourcesPropertiesResponsePtrType) ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesResponsePtrOutput)
}

type UserDefinedResourcesPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserDefinedResourcesPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDefinedResourcesPropertiesResponse)(nil)).Elem()
}

func (o UserDefinedResourcesPropertiesResponseOutput) ToUserDefinedResourcesPropertiesResponseOutput() UserDefinedResourcesPropertiesResponseOutput {
	return o
}

func (o UserDefinedResourcesPropertiesResponseOutput) ToUserDefinedResourcesPropertiesResponseOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponseOutput {
	return o
}

func (o UserDefinedResourcesPropertiesResponseOutput) ToUserDefinedResourcesPropertiesResponsePtrOutput() UserDefinedResourcesPropertiesResponsePtrOutput {
	return o.ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o UserDefinedResourcesPropertiesResponseOutput) ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserDefinedResourcesPropertiesResponse) *UserDefinedResourcesPropertiesResponse {
		return &v
	}).(UserDefinedResourcesPropertiesResponsePtrOutput)
}

func (o UserDefinedResourcesPropertiesResponseOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v UserDefinedResourcesPropertiesResponse) string { return v.Query }).(pulumi.StringOutput)
}

func (o UserDefinedResourcesPropertiesResponseOutput) QuerySubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UserDefinedResourcesPropertiesResponse) []string { return v.QuerySubscriptions }).(pulumi.StringArrayOutput)
}

type UserDefinedResourcesPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (UserDefinedResourcesPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDefinedResourcesPropertiesResponse)(nil)).Elem()
}

func (o UserDefinedResourcesPropertiesResponsePtrOutput) ToUserDefinedResourcesPropertiesResponsePtrOutput() UserDefinedResourcesPropertiesResponsePtrOutput {
	return o
}

func (o UserDefinedResourcesPropertiesResponsePtrOutput) ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponsePtrOutput {
	return o
}

func (o UserDefinedResourcesPropertiesResponsePtrOutput) Elem() UserDefinedResourcesPropertiesResponseOutput {
	return o.ApplyT(func(v *UserDefinedResourcesPropertiesResponse) UserDefinedResourcesPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret UserDefinedResourcesPropertiesResponse
		return ret
	}).(UserDefinedResourcesPropertiesResponseOutput)
}

func (o UserDefinedResourcesPropertiesResponsePtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDefinedResourcesPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Query
	}).(pulumi.StringPtrOutput)
}

func (o UserDefinedResourcesPropertiesResponsePtrOutput) QuerySubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserDefinedResourcesPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.QuerySubscriptions
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AdditionalWorkspacesPropertiesOutput{})
	pulumi.RegisterOutputType(AdditionalWorkspacesPropertiesArrayOutput{})
	pulumi.RegisterOutputType(AdditionalWorkspacesPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AdditionalWorkspacesPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(AllowlistCustomAlertRuleOutput{})
	pulumi.RegisterOutputType(AllowlistCustomAlertRuleArrayOutput{})
	pulumi.RegisterOutputType(AllowlistCustomAlertRuleResponseOutput{})
	pulumi.RegisterOutputType(AllowlistCustomAlertRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(DenylistCustomAlertRuleOutput{})
	pulumi.RegisterOutputType(DenylistCustomAlertRuleArrayOutput{})
	pulumi.RegisterOutputType(DenylistCustomAlertRuleResponseOutput{})
	pulumi.RegisterOutputType(DenylistCustomAlertRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(RecommendationConfigurationPropertiesOutput{})
	pulumi.RegisterOutputType(RecommendationConfigurationPropertiesArrayOutput{})
	pulumi.RegisterOutputType(RecommendationConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RecommendationConfigurationPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(ThresholdCustomAlertRuleOutput{})
	pulumi.RegisterOutputType(ThresholdCustomAlertRuleArrayOutput{})
	pulumi.RegisterOutputType(ThresholdCustomAlertRuleResponseOutput{})
	pulumi.RegisterOutputType(ThresholdCustomAlertRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(TimeWindowCustomAlertRuleOutput{})
	pulumi.RegisterOutputType(TimeWindowCustomAlertRuleArrayOutput{})
	pulumi.RegisterOutputType(TimeWindowCustomAlertRuleResponseOutput{})
	pulumi.RegisterOutputType(TimeWindowCustomAlertRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(UserDefinedResourcesPropertiesOutput{})
	pulumi.RegisterOutputType(UserDefinedResourcesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(UserDefinedResourcesPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserDefinedResourcesPropertiesResponsePtrOutput{})
}
