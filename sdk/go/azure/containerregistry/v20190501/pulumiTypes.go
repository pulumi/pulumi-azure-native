


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActorResponse struct {
	Name *string `pulumi:"name"`
}

type EventContentResponse struct {
	Action    *string          `pulumi:"action"`
	Actor     *ActorResponse   `pulumi:"actor"`
	Id        *string          `pulumi:"id"`
	Request   *RequestResponse `pulumi:"request"`
	Source    *SourceResponse  `pulumi:"source"`
	Target    *TargetResponse  `pulumi:"target"`
	Timestamp *string          `pulumi:"timestamp"`
}

type EventRequestMessageResponse struct {
	Content    *EventContentResponse `pulumi:"content"`
	Headers    map[string]string     `pulumi:"headers"`
	Method     *string               `pulumi:"method"`
	RequestUri *string               `pulumi:"requestUri"`
	Version    *string               `pulumi:"version"`
}

type EventResponse struct {
	EventRequestMessage  *EventRequestMessageResponse  `pulumi:"eventRequestMessage"`
	EventResponseMessage *EventResponseMessageResponse `pulumi:"eventResponseMessage"`
	Id                   *string                       `pulumi:"id"`
}

type EventResponseMessageResponse struct {
	Content      *string           `pulumi:"content"`
	Headers      map[string]string `pulumi:"headers"`
	ReasonPhrase *string           `pulumi:"reasonPhrase"`
	StatusCode   *string           `pulumi:"statusCode"`
	Version      *string           `pulumi:"version"`
}

type IPRule struct {
	Action           *string `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}


func (val *IPRule) Defaults() *IPRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type IPRuleInput interface {
	pulumi.Input

	ToIPRuleOutput() IPRuleOutput
	ToIPRuleOutputWithContext(context.Context) IPRuleOutput
}

type IPRuleArgs struct {
	Action           pulumi.StringPtrInput `pulumi:"action"`
	IPAddressOrRange pulumi.StringInput    `pulumi:"iPAddressOrRange"`
}

func (IPRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (i IPRuleArgs) ToIPRuleOutput() IPRuleOutput {
	return i.ToIPRuleOutputWithContext(context.Background())
}

func (i IPRuleArgs) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleOutput)
}





type IPRuleArrayInput interface {
	pulumi.Input

	ToIPRuleArrayOutput() IPRuleArrayOutput
	ToIPRuleArrayOutputWithContext(context.Context) IPRuleArrayOutput
}

type IPRuleArray []IPRuleInput

func (IPRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (i IPRuleArray) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return i.ToIPRuleArrayOutputWithContext(context.Background())
}

func (i IPRuleArray) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleArrayOutput)
}

type IPRuleOutput struct{ *pulumi.OutputState }

func (IPRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (o IPRuleOutput) ToIPRuleOutput() IPRuleOutput {
	return o
}

func (o IPRuleOutput) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return o
}

func (o IPRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRule) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IPRuleOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRule) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleArrayOutput struct{ *pulumi.OutputState }

func (IPRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) Index(i pulumi.IntInput) IPRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRule {
		return vs[0].([]IPRule)[vs[1].(int)]
	}).(IPRuleOutput)
}

type IPRuleResponse struct {
	Action           *string `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}


func (val *IPRuleResponse) Defaults() *IPRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type IPRuleResponseOutput struct{ *pulumi.OutputState }

func (IPRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutput() IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutputWithContext(ctx context.Context) IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IPRuleResponseOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRuleResponse) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IPRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutputWithContext(ctx context.Context) IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) Index(i pulumi.IntInput) IPRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRuleResponse {
		return vs[0].([]IPRuleResponse)[vs[1].(int)]
	}).(IPRuleResponseOutput)
}

type NetworkRuleSet struct {
	DefaultAction       string               `pulumi:"defaultAction"`
	IpRules             []IPRule             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}


func (val *NetworkRuleSet) Defaults() *NetworkRuleSet {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultAction) {
		tmp.DefaultAction = "Allow"
	}
	return &tmp
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	DefaultAction       pulumi.StringInput           `pulumi:"defaultAction"`
	IpRules             IPRuleArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSet) string { return v.DefaultAction }).(pulumi.StringOutput)
}

func (o NetworkRuleSetOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []IPRule { return v.IpRules }).(IPRuleArrayOutput)
}

func (o NetworkRuleSetOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []VirtualNetworkRule { return v.VirtualNetworkRules }).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPtrOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []IPRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleArrayOutput)
}

func (o NetworkRuleSetPtrOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []VirtualNetworkRule {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	DefaultAction       string                       `pulumi:"defaultAction"`
	IpRules             []IPRuleResponse             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}


func (val *NetworkRuleSetResponse) Defaults() *NetworkRuleSetResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultAction) {
		tmp.DefaultAction = "Allow"
	}
	return &tmp
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) string { return v.DefaultAction }).(pulumi.StringOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IPRuleResponse { return v.IpRules }).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponseOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

type NetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) Elem() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) NetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetResponse
		return ret
	}).(NetworkRuleSetResponseOutput)
}

func (o NetworkRuleSetResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []IPRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponsePtrOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []VirtualNetworkRuleResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleResponseArrayOutput)
}

type Policies struct {
	QuarantinePolicy *QuarantinePolicy `pulumi:"quarantinePolicy"`
	RetentionPolicy  *RetentionPolicy  `pulumi:"retentionPolicy"`
	TrustPolicy      *TrustPolicy      `pulumi:"trustPolicy"`
}


func (val *Policies) Defaults() *Policies {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.QuarantinePolicy = tmp.QuarantinePolicy.Defaults()

	tmp.RetentionPolicy = tmp.RetentionPolicy.Defaults()

	tmp.TrustPolicy = tmp.TrustPolicy.Defaults()

	return &tmp
}





type PoliciesInput interface {
	pulumi.Input

	ToPoliciesOutput() PoliciesOutput
	ToPoliciesOutputWithContext(context.Context) PoliciesOutput
}

type PoliciesArgs struct {
	QuarantinePolicy QuarantinePolicyPtrInput `pulumi:"quarantinePolicy"`
	RetentionPolicy  RetentionPolicyPtrInput  `pulumi:"retentionPolicy"`
	TrustPolicy      TrustPolicyPtrInput      `pulumi:"trustPolicy"`
}

func (PoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Policies)(nil)).Elem()
}

func (i PoliciesArgs) ToPoliciesOutput() PoliciesOutput {
	return i.ToPoliciesOutputWithContext(context.Background())
}

func (i PoliciesArgs) ToPoliciesOutputWithContext(ctx context.Context) PoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesOutput)
}

func (i PoliciesArgs) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return i.ToPoliciesPtrOutputWithContext(context.Background())
}

func (i PoliciesArgs) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesOutput).ToPoliciesPtrOutputWithContext(ctx)
}









type PoliciesPtrInput interface {
	pulumi.Input

	ToPoliciesPtrOutput() PoliciesPtrOutput
	ToPoliciesPtrOutputWithContext(context.Context) PoliciesPtrOutput
}

type policiesPtrType PoliciesArgs

func PoliciesPtr(v *PoliciesArgs) PoliciesPtrInput {
	return (*policiesPtrType)(v)
}

func (*policiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Policies)(nil)).Elem()
}

func (i *policiesPtrType) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return i.ToPoliciesPtrOutputWithContext(context.Background())
}

func (i *policiesPtrType) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesPtrOutput)
}

type PoliciesOutput struct{ *pulumi.OutputState }

func (PoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policies)(nil)).Elem()
}

func (o PoliciesOutput) ToPoliciesOutput() PoliciesOutput {
	return o
}

func (o PoliciesOutput) ToPoliciesOutputWithContext(ctx context.Context) PoliciesOutput {
	return o
}

func (o PoliciesOutput) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return o.ToPoliciesPtrOutputWithContext(context.Background())
}

func (o PoliciesOutput) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Policies) *Policies {
		return &v
	}).(PoliciesPtrOutput)
}

func (o PoliciesOutput) QuarantinePolicy() QuarantinePolicyPtrOutput {
	return o.ApplyT(func(v Policies) *QuarantinePolicy { return v.QuarantinePolicy }).(QuarantinePolicyPtrOutput)
}

func (o PoliciesOutput) RetentionPolicy() RetentionPolicyPtrOutput {
	return o.ApplyT(func(v Policies) *RetentionPolicy { return v.RetentionPolicy }).(RetentionPolicyPtrOutput)
}

func (o PoliciesOutput) TrustPolicy() TrustPolicyPtrOutput {
	return o.ApplyT(func(v Policies) *TrustPolicy { return v.TrustPolicy }).(TrustPolicyPtrOutput)
}

type PoliciesPtrOutput struct{ *pulumi.OutputState }

func (PoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policies)(nil)).Elem()
}

func (o PoliciesPtrOutput) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return o
}

func (o PoliciesPtrOutput) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return o
}

func (o PoliciesPtrOutput) Elem() PoliciesOutput {
	return o.ApplyT(func(v *Policies) Policies {
		if v != nil {
			return *v
		}
		var ret Policies
		return ret
	}).(PoliciesOutput)
}

func (o PoliciesPtrOutput) QuarantinePolicy() QuarantinePolicyPtrOutput {
	return o.ApplyT(func(v *Policies) *QuarantinePolicy {
		if v == nil {
			return nil
		}
		return v.QuarantinePolicy
	}).(QuarantinePolicyPtrOutput)
}

func (o PoliciesPtrOutput) RetentionPolicy() RetentionPolicyPtrOutput {
	return o.ApplyT(func(v *Policies) *RetentionPolicy {
		if v == nil {
			return nil
		}
		return v.RetentionPolicy
	}).(RetentionPolicyPtrOutput)
}

func (o PoliciesPtrOutput) TrustPolicy() TrustPolicyPtrOutput {
	return o.ApplyT(func(v *Policies) *TrustPolicy {
		if v == nil {
			return nil
		}
		return v.TrustPolicy
	}).(TrustPolicyPtrOutput)
}

type PoliciesResponse struct {
	QuarantinePolicy *QuarantinePolicyResponse `pulumi:"quarantinePolicy"`
	RetentionPolicy  *RetentionPolicyResponse  `pulumi:"retentionPolicy"`
	TrustPolicy      *TrustPolicyResponse      `pulumi:"trustPolicy"`
}


func (val *PoliciesResponse) Defaults() *PoliciesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.QuarantinePolicy = tmp.QuarantinePolicy.Defaults()

	tmp.RetentionPolicy = tmp.RetentionPolicy.Defaults()

	tmp.TrustPolicy = tmp.TrustPolicy.Defaults()

	return &tmp
}

type PoliciesResponseOutput struct{ *pulumi.OutputState }

func (PoliciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PoliciesResponse)(nil)).Elem()
}

func (o PoliciesResponseOutput) ToPoliciesResponseOutput() PoliciesResponseOutput {
	return o
}

func (o PoliciesResponseOutput) ToPoliciesResponseOutputWithContext(ctx context.Context) PoliciesResponseOutput {
	return o
}

func (o PoliciesResponseOutput) QuarantinePolicy() QuarantinePolicyResponsePtrOutput {
	return o.ApplyT(func(v PoliciesResponse) *QuarantinePolicyResponse { return v.QuarantinePolicy }).(QuarantinePolicyResponsePtrOutput)
}

func (o PoliciesResponseOutput) RetentionPolicy() RetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v PoliciesResponse) *RetentionPolicyResponse { return v.RetentionPolicy }).(RetentionPolicyResponsePtrOutput)
}

func (o PoliciesResponseOutput) TrustPolicy() TrustPolicyResponsePtrOutput {
	return o.ApplyT(func(v PoliciesResponse) *TrustPolicyResponse { return v.TrustPolicy }).(TrustPolicyResponsePtrOutput)
}

type PoliciesResponsePtrOutput struct{ *pulumi.OutputState }

func (PoliciesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PoliciesResponse)(nil)).Elem()
}

func (o PoliciesResponsePtrOutput) ToPoliciesResponsePtrOutput() PoliciesResponsePtrOutput {
	return o
}

func (o PoliciesResponsePtrOutput) ToPoliciesResponsePtrOutputWithContext(ctx context.Context) PoliciesResponsePtrOutput {
	return o
}

func (o PoliciesResponsePtrOutput) Elem() PoliciesResponseOutput {
	return o.ApplyT(func(v *PoliciesResponse) PoliciesResponse {
		if v != nil {
			return *v
		}
		var ret PoliciesResponse
		return ret
	}).(PoliciesResponseOutput)
}

func (o PoliciesResponsePtrOutput) QuarantinePolicy() QuarantinePolicyResponsePtrOutput {
	return o.ApplyT(func(v *PoliciesResponse) *QuarantinePolicyResponse {
		if v == nil {
			return nil
		}
		return v.QuarantinePolicy
	}).(QuarantinePolicyResponsePtrOutput)
}

func (o PoliciesResponsePtrOutput) RetentionPolicy() RetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *PoliciesResponse) *RetentionPolicyResponse {
		if v == nil {
			return nil
		}
		return v.RetentionPolicy
	}).(RetentionPolicyResponsePtrOutput)
}

func (o PoliciesResponsePtrOutput) TrustPolicy() TrustPolicyResponsePtrOutput {
	return o.ApplyT(func(v *PoliciesResponse) *TrustPolicyResponse {
		if v == nil {
			return nil
		}
		return v.TrustPolicy
	}).(TrustPolicyResponsePtrOutput)
}

type QuarantinePolicy struct {
	Status *string `pulumi:"status"`
}


func (val *QuarantinePolicy) Defaults() *QuarantinePolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}





type QuarantinePolicyInput interface {
	pulumi.Input

	ToQuarantinePolicyOutput() QuarantinePolicyOutput
	ToQuarantinePolicyOutputWithContext(context.Context) QuarantinePolicyOutput
}

type QuarantinePolicyArgs struct {
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (QuarantinePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QuarantinePolicy)(nil)).Elem()
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyOutput() QuarantinePolicyOutput {
	return i.ToQuarantinePolicyOutputWithContext(context.Background())
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyOutputWithContext(ctx context.Context) QuarantinePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyOutput)
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return i.ToQuarantinePolicyPtrOutputWithContext(context.Background())
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyOutput).ToQuarantinePolicyPtrOutputWithContext(ctx)
}









type QuarantinePolicyPtrInput interface {
	pulumi.Input

	ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput
	ToQuarantinePolicyPtrOutputWithContext(context.Context) QuarantinePolicyPtrOutput
}

type quarantinePolicyPtrType QuarantinePolicyArgs

func QuarantinePolicyPtr(v *QuarantinePolicyArgs) QuarantinePolicyPtrInput {
	return (*quarantinePolicyPtrType)(v)
}

func (*quarantinePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QuarantinePolicy)(nil)).Elem()
}

func (i *quarantinePolicyPtrType) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return i.ToQuarantinePolicyPtrOutputWithContext(context.Background())
}

func (i *quarantinePolicyPtrType) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyPtrOutput)
}

type QuarantinePolicyOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuarantinePolicy)(nil)).Elem()
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyOutput() QuarantinePolicyOutput {
	return o
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyOutputWithContext(ctx context.Context) QuarantinePolicyOutput {
	return o
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return o.ToQuarantinePolicyPtrOutputWithContext(context.Background())
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QuarantinePolicy) *QuarantinePolicy {
		return &v
	}).(QuarantinePolicyPtrOutput)
}

func (o QuarantinePolicyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuarantinePolicy) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type QuarantinePolicyPtrOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuarantinePolicy)(nil)).Elem()
}

func (o QuarantinePolicyPtrOutput) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return o
}

func (o QuarantinePolicyPtrOutput) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return o
}

func (o QuarantinePolicyPtrOutput) Elem() QuarantinePolicyOutput {
	return o.ApplyT(func(v *QuarantinePolicy) QuarantinePolicy {
		if v != nil {
			return *v
		}
		var ret QuarantinePolicy
		return ret
	}).(QuarantinePolicyOutput)
}

func (o QuarantinePolicyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuarantinePolicy) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type QuarantinePolicyResponse struct {
	Status *string `pulumi:"status"`
}


func (val *QuarantinePolicyResponse) Defaults() *QuarantinePolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}

type QuarantinePolicyResponseOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuarantinePolicyResponse)(nil)).Elem()
}

func (o QuarantinePolicyResponseOutput) ToQuarantinePolicyResponseOutput() QuarantinePolicyResponseOutput {
	return o
}

func (o QuarantinePolicyResponseOutput) ToQuarantinePolicyResponseOutputWithContext(ctx context.Context) QuarantinePolicyResponseOutput {
	return o
}

func (o QuarantinePolicyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuarantinePolicyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type QuarantinePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuarantinePolicyResponse)(nil)).Elem()
}

func (o QuarantinePolicyResponsePtrOutput) ToQuarantinePolicyResponsePtrOutput() QuarantinePolicyResponsePtrOutput {
	return o
}

func (o QuarantinePolicyResponsePtrOutput) ToQuarantinePolicyResponsePtrOutputWithContext(ctx context.Context) QuarantinePolicyResponsePtrOutput {
	return o
}

func (o QuarantinePolicyResponsePtrOutput) Elem() QuarantinePolicyResponseOutput {
	return o.ApplyT(func(v *QuarantinePolicyResponse) QuarantinePolicyResponse {
		if v != nil {
			return *v
		}
		var ret QuarantinePolicyResponse
		return ret
	}).(QuarantinePolicyResponseOutput)
}

func (o QuarantinePolicyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuarantinePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type RegistryPasswordResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type RequestResponse struct {
	Addr      *string `pulumi:"addr"`
	Host      *string `pulumi:"host"`
	Id        *string `pulumi:"id"`
	Method    *string `pulumi:"method"`
	Useragent *string `pulumi:"useragent"`
}

type RetentionPolicy struct {
	Days   *int    `pulumi:"days"`
	Status *string `pulumi:"status"`
}


func (val *RetentionPolicy) Defaults() *RetentionPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Days) {
		days_ := 7
		tmp.Days = &days_
	}
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}





type RetentionPolicyInput interface {
	pulumi.Input

	ToRetentionPolicyOutput() RetentionPolicyOutput
	ToRetentionPolicyOutputWithContext(context.Context) RetentionPolicyOutput
}

type RetentionPolicyArgs struct {
	Days   pulumi.IntPtrInput    `pulumi:"days"`
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (RetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return i.ToRetentionPolicyOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput)
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput).ToRetentionPolicyPtrOutputWithContext(ctx)
}









type RetentionPolicyPtrInput interface {
	pulumi.Input

	ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput
	ToRetentionPolicyPtrOutputWithContext(context.Context) RetentionPolicyPtrOutput
}

type retentionPolicyPtrType RetentionPolicyArgs

func RetentionPolicyPtr(v *RetentionPolicyArgs) RetentionPolicyPtrInput {
	return (*retentionPolicyPtrType)(v)
}

func (*retentionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyPtrOutput)
}

type RetentionPolicyOutput struct{ *pulumi.OutputState }

func (RetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionPolicy) *RetentionPolicy {
		return &v
	}).(RetentionPolicyPtrOutput)
}

func (o RetentionPolicyOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionPolicy) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionPolicy) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type RetentionPolicyPtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) Elem() RetentionPolicyOutput {
	return o.ApplyT(func(v *RetentionPolicy) RetentionPolicy {
		if v != nil {
			return *v
		}
		var ret RetentionPolicy
		return ret
	}).(RetentionPolicyOutput)
}

func (o RetentionPolicyPtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type RetentionPolicyResponse struct {
	Days            *int    `pulumi:"days"`
	LastUpdatedTime string  `pulumi:"lastUpdatedTime"`
	Status          *string `pulumi:"status"`
}


func (val *RetentionPolicyResponse) Defaults() *RetentionPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Days) {
		days_ := 7
		tmp.Days = &days_
	}
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}

type RetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutputWithContext(ctx context.Context) RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyResponseOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) string { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o RetentionPolicyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type RetentionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) Elem() RetentionPolicyResponseOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) RetentionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RetentionPolicyResponse
		return ret
	}).(RetentionPolicyResponseOutput)
}

func (o RetentionPolicyResponsePtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyResponsePtrOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTime
	}).(pulumi.StringPtrOutput)
}

func (o RetentionPolicyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
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

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
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

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SourceResponse struct {
	Addr       *string `pulumi:"addr"`
	InstanceID *string `pulumi:"instanceID"`
}

type StatusResponse struct {
	DisplayStatus string `pulumi:"displayStatus"`
	Message       string `pulumi:"message"`
	Timestamp     string `pulumi:"timestamp"`
}

type StatusResponseOutput struct{ *pulumi.OutputState }

func (StatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusResponse)(nil)).Elem()
}

func (o StatusResponseOutput) ToStatusResponseOutput() StatusResponseOutput {
	return o
}

func (o StatusResponseOutput) ToStatusResponseOutputWithContext(ctx context.Context) StatusResponseOutput {
	return o
}

func (o StatusResponseOutput) DisplayStatus() pulumi.StringOutput {
	return o.ApplyT(func(v StatusResponse) string { return v.DisplayStatus }).(pulumi.StringOutput)
}

func (o StatusResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v StatusResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o StatusResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v StatusResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

type StorageAccountProperties struct {
	Id string `pulumi:"id"`
}





type StorageAccountPropertiesInput interface {
	pulumi.Input

	ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput
	ToStorageAccountPropertiesOutputWithContext(context.Context) StorageAccountPropertiesOutput
}

type StorageAccountPropertiesArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (StorageAccountPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return i.ToStorageAccountPropertiesOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesOutput)
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return i.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesOutput).ToStorageAccountPropertiesPtrOutputWithContext(ctx)
}









type StorageAccountPropertiesPtrInput interface {
	pulumi.Input

	ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput
	ToStorageAccountPropertiesPtrOutputWithContext(context.Context) StorageAccountPropertiesPtrOutput
}

type storageAccountPropertiesPtrType StorageAccountPropertiesArgs

func StorageAccountPropertiesPtr(v *StorageAccountPropertiesArgs) StorageAccountPropertiesPtrInput {
	return (*storageAccountPropertiesPtrType)(v)
}

func (*storageAccountPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountProperties)(nil)).Elem()
}

func (i *storageAccountPropertiesPtrType) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return i.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i *storageAccountPropertiesPtrType) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesPtrOutput)
}

type StorageAccountPropertiesOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return o.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountProperties) *StorageAccountProperties {
		return &v
	}).(StorageAccountPropertiesPtrOutput)
}

func (o StorageAccountPropertiesOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountProperties) string { return v.Id }).(pulumi.StringOutput)
}

type StorageAccountPropertiesPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountProperties)(nil)).Elem()
}

func (o StorageAccountPropertiesPtrOutput) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return o
}

func (o StorageAccountPropertiesPtrOutput) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return o
}

func (o StorageAccountPropertiesPtrOutput) Elem() StorageAccountPropertiesOutput {
	return o.ApplyT(func(v *StorageAccountProperties) StorageAccountProperties {
		if v != nil {
			return *v
		}
		var ret StorageAccountProperties
		return ret
	}).(StorageAccountPropertiesOutput)
}

func (o StorageAccountPropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type StorageAccountPropertiesResponse struct {
	Id string `pulumi:"id"`
}

type StorageAccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutputWithContext(ctx context.Context) StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

type StorageAccountPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) Elem() StorageAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) StorageAccountPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountPropertiesResponse
		return ret
	}).(StorageAccountPropertiesResponseOutput)
}

func (o StorageAccountPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type TargetResponse struct {
	Digest     *string  `pulumi:"digest"`
	Length     *float64 `pulumi:"length"`
	MediaType  *string  `pulumi:"mediaType"`
	Name       *string  `pulumi:"name"`
	Repository *string  `pulumi:"repository"`
	Size       *float64 `pulumi:"size"`
	Tag        *string  `pulumi:"tag"`
	Url        *string  `pulumi:"url"`
	Version    *string  `pulumi:"version"`
}

type TrustPolicy struct {
	Status *string `pulumi:"status"`
	Type   *string `pulumi:"type"`
}


func (val *TrustPolicy) Defaults() *TrustPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	if isZero(tmp.Type) {
		type_ := "Notary"
		tmp.Type = &type_
	}
	return &tmp
}





type TrustPolicyInput interface {
	pulumi.Input

	ToTrustPolicyOutput() TrustPolicyOutput
	ToTrustPolicyOutputWithContext(context.Context) TrustPolicyOutput
}

type TrustPolicyArgs struct {
	Status pulumi.StringPtrInput `pulumi:"status"`
	Type   pulumi.StringPtrInput `pulumi:"type"`
}

func (TrustPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicy)(nil)).Elem()
}

func (i TrustPolicyArgs) ToTrustPolicyOutput() TrustPolicyOutput {
	return i.ToTrustPolicyOutputWithContext(context.Background())
}

func (i TrustPolicyArgs) ToTrustPolicyOutputWithContext(ctx context.Context) TrustPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyOutput)
}

func (i TrustPolicyArgs) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return i.ToTrustPolicyPtrOutputWithContext(context.Background())
}

func (i TrustPolicyArgs) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyOutput).ToTrustPolicyPtrOutputWithContext(ctx)
}









type TrustPolicyPtrInput interface {
	pulumi.Input

	ToTrustPolicyPtrOutput() TrustPolicyPtrOutput
	ToTrustPolicyPtrOutputWithContext(context.Context) TrustPolicyPtrOutput
}

type trustPolicyPtrType TrustPolicyArgs

func TrustPolicyPtr(v *TrustPolicyArgs) TrustPolicyPtrInput {
	return (*trustPolicyPtrType)(v)
}

func (*trustPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustPolicy)(nil)).Elem()
}

func (i *trustPolicyPtrType) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return i.ToTrustPolicyPtrOutputWithContext(context.Background())
}

func (i *trustPolicyPtrType) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyPtrOutput)
}

type TrustPolicyOutput struct{ *pulumi.OutputState }

func (TrustPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicy)(nil)).Elem()
}

func (o TrustPolicyOutput) ToTrustPolicyOutput() TrustPolicyOutput {
	return o
}

func (o TrustPolicyOutput) ToTrustPolicyOutputWithContext(ctx context.Context) TrustPolicyOutput {
	return o
}

func (o TrustPolicyOutput) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return o.ToTrustPolicyPtrOutputWithContext(context.Background())
}

func (o TrustPolicyOutput) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrustPolicy) *TrustPolicy {
		return &v
	}).(TrustPolicyPtrOutput)
}

func (o TrustPolicyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicy) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o TrustPolicyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicy) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type TrustPolicyPtrOutput struct{ *pulumi.OutputState }

func (TrustPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustPolicy)(nil)).Elem()
}

func (o TrustPolicyPtrOutput) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return o
}

func (o TrustPolicyPtrOutput) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return o
}

func (o TrustPolicyPtrOutput) Elem() TrustPolicyOutput {
	return o.ApplyT(func(v *TrustPolicy) TrustPolicy {
		if v != nil {
			return *v
		}
		var ret TrustPolicy
		return ret
	}).(TrustPolicyOutput)
}

func (o TrustPolicyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o TrustPolicyPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type TrustPolicyResponse struct {
	Status *string `pulumi:"status"`
	Type   *string `pulumi:"type"`
}


func (val *TrustPolicyResponse) Defaults() *TrustPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	if isZero(tmp.Type) {
		type_ := "Notary"
		tmp.Type = &type_
	}
	return &tmp
}

type TrustPolicyResponseOutput struct{ *pulumi.OutputState }

func (TrustPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicyResponse)(nil)).Elem()
}

func (o TrustPolicyResponseOutput) ToTrustPolicyResponseOutput() TrustPolicyResponseOutput {
	return o
}

func (o TrustPolicyResponseOutput) ToTrustPolicyResponseOutputWithContext(ctx context.Context) TrustPolicyResponseOutput {
	return o
}

func (o TrustPolicyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o TrustPolicyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type TrustPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (TrustPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustPolicyResponse)(nil)).Elem()
}

func (o TrustPolicyResponsePtrOutput) ToTrustPolicyResponsePtrOutput() TrustPolicyResponsePtrOutput {
	return o
}

func (o TrustPolicyResponsePtrOutput) ToTrustPolicyResponsePtrOutputWithContext(ctx context.Context) TrustPolicyResponsePtrOutput {
	return o
}

func (o TrustPolicyResponsePtrOutput) Elem() TrustPolicyResponseOutput {
	return o.ApplyT(func(v *TrustPolicyResponse) TrustPolicyResponse {
		if v != nil {
			return *v
		}
		var ret TrustPolicyResponse
		return ret
	}).(TrustPolicyResponseOutput)
}

func (o TrustPolicyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o TrustPolicyResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkRule struct {
	Action                   *string `pulumi:"action"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRule) Defaults() *VirtualNetworkRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Action                   pulumi.StringPtrInput `pulumi:"action"`
	VirtualNetworkResourceId pulumi.StringInput    `pulumi:"virtualNetworkResourceId"`
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRule) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Action                   *string `pulumi:"action"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRuleResponse) Defaults() *VirtualNetworkRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(IPRuleOutput{})
	pulumi.RegisterOutputType(IPRuleArrayOutput{})
	pulumi.RegisterOutputType(IPRuleResponseOutput{})
	pulumi.RegisterOutputType(IPRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(PoliciesOutput{})
	pulumi.RegisterOutputType(PoliciesPtrOutput{})
	pulumi.RegisterOutputType(PoliciesResponseOutput{})
	pulumi.RegisterOutputType(PoliciesResponsePtrOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyPtrOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyResponseOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyOutput{})
	pulumi.RegisterOutputType(RetentionPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(StatusResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TrustPolicyOutput{})
	pulumi.RegisterOutputType(TrustPolicyPtrOutput{})
	pulumi.RegisterOutputType(TrustPolicyResponseOutput{})
	pulumi.RegisterOutputType(TrustPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
