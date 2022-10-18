


package v20190615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CacheExpirationActionParameters struct {
	CacheBehavior string  `pulumi:"cacheBehavior"`
	CacheDuration *string `pulumi:"cacheDuration"`
	CacheType     string  `pulumi:"cacheType"`
	OdataType     string  `pulumi:"odataType"`
}

type CacheExpirationActionParametersResponse struct {
	CacheBehavior string  `pulumi:"cacheBehavior"`
	CacheDuration *string `pulumi:"cacheDuration"`
	CacheType     string  `pulumi:"cacheType"`
	OdataType     string  `pulumi:"odataType"`
}

type CacheKeyQueryStringActionParameters struct {
	OdataType           string  `pulumi:"odataType"`
	QueryParameters     *string `pulumi:"queryParameters"`
	QueryStringBehavior string  `pulumi:"queryStringBehavior"`
}

type CacheKeyQueryStringActionParametersResponse struct {
	OdataType           string  `pulumi:"odataType"`
	QueryParameters     *string `pulumi:"queryParameters"`
	QueryStringBehavior string  `pulumi:"queryStringBehavior"`
}

type CdnCertificateSourceParametersResponse struct {
	CertificateType string `pulumi:"certificateType"`
	OdataType       string `pulumi:"odataType"`
}

type CdnEndpointResponse struct {
	Id *string `pulumi:"id"`
}

type CdnEndpointResponseOutput struct{ *pulumi.OutputState }

func (CdnEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CdnEndpointResponse)(nil)).Elem()
}

func (o CdnEndpointResponseOutput) ToCdnEndpointResponseOutput() CdnEndpointResponseOutput {
	return o
}

func (o CdnEndpointResponseOutput) ToCdnEndpointResponseOutputWithContext(ctx context.Context) CdnEndpointResponseOutput {
	return o
}

func (o CdnEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CdnEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type CdnEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (CdnEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CdnEndpointResponse)(nil)).Elem()
}

func (o CdnEndpointResponseArrayOutput) ToCdnEndpointResponseArrayOutput() CdnEndpointResponseArrayOutput {
	return o
}

func (o CdnEndpointResponseArrayOutput) ToCdnEndpointResponseArrayOutputWithContext(ctx context.Context) CdnEndpointResponseArrayOutput {
	return o
}

func (o CdnEndpointResponseArrayOutput) Index(i pulumi.IntInput) CdnEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CdnEndpointResponse {
		return vs[0].([]CdnEndpointResponse)[vs[1].(int)]
	}).(CdnEndpointResponseOutput)
}

type CdnManagedHttpsParametersResponse struct {
	CertificateSource           string                                 `pulumi:"certificateSource"`
	CertificateSourceParameters CdnCertificateSourceParametersResponse `pulumi:"certificateSourceParameters"`
	MinimumTlsVersion           *string                                `pulumi:"minimumTlsVersion"`
	ProtocolType                string                                 `pulumi:"protocolType"`
}

type CookiesMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}

type CookiesMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}

type CustomRule struct {
	Action          string           `pulumi:"action"`
	EnabledState    *string          `pulumi:"enabledState"`
	MatchConditions []MatchCondition `pulumi:"matchConditions"`
	Name            string           `pulumi:"name"`
	Priority        int              `pulumi:"priority"`
}





type CustomRuleInput interface {
	pulumi.Input

	ToCustomRuleOutput() CustomRuleOutput
	ToCustomRuleOutputWithContext(context.Context) CustomRuleOutput
}

type CustomRuleArgs struct {
	Action          pulumi.StringInput       `pulumi:"action"`
	EnabledState    pulumi.StringPtrInput    `pulumi:"enabledState"`
	MatchConditions MatchConditionArrayInput `pulumi:"matchConditions"`
	Name            pulumi.StringInput       `pulumi:"name"`
	Priority        pulumi.IntInput          `pulumi:"priority"`
}

func (CustomRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRule)(nil)).Elem()
}

func (i CustomRuleArgs) ToCustomRuleOutput() CustomRuleOutput {
	return i.ToCustomRuleOutputWithContext(context.Background())
}

func (i CustomRuleArgs) ToCustomRuleOutputWithContext(ctx context.Context) CustomRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleOutput)
}





type CustomRuleArrayInput interface {
	pulumi.Input

	ToCustomRuleArrayOutput() CustomRuleArrayOutput
	ToCustomRuleArrayOutputWithContext(context.Context) CustomRuleArrayOutput
}

type CustomRuleArray []CustomRuleInput

func (CustomRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRule)(nil)).Elem()
}

func (i CustomRuleArray) ToCustomRuleArrayOutput() CustomRuleArrayOutput {
	return i.ToCustomRuleArrayOutputWithContext(context.Background())
}

func (i CustomRuleArray) ToCustomRuleArrayOutputWithContext(ctx context.Context) CustomRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleArrayOutput)
}

type CustomRuleOutput struct{ *pulumi.OutputState }

func (CustomRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRule)(nil)).Elem()
}

func (o CustomRuleOutput) ToCustomRuleOutput() CustomRuleOutput {
	return o
}

func (o CustomRuleOutput) ToCustomRuleOutputWithContext(ctx context.Context) CustomRuleOutput {
	return o
}

func (o CustomRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRule) string { return v.Action }).(pulumi.StringOutput)
}

func (o CustomRuleOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRule) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o CustomRuleOutput) MatchConditions() MatchConditionArrayOutput {
	return o.ApplyT(func(v CustomRule) []MatchCondition { return v.MatchConditions }).(MatchConditionArrayOutput)
}

func (o CustomRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRule) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v CustomRule) int { return v.Priority }).(pulumi.IntOutput)
}

type CustomRuleArrayOutput struct{ *pulumi.OutputState }

func (CustomRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRule)(nil)).Elem()
}

func (o CustomRuleArrayOutput) ToCustomRuleArrayOutput() CustomRuleArrayOutput {
	return o
}

func (o CustomRuleArrayOutput) ToCustomRuleArrayOutputWithContext(ctx context.Context) CustomRuleArrayOutput {
	return o
}

func (o CustomRuleArrayOutput) Index(i pulumi.IntInput) CustomRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRule {
		return vs[0].([]CustomRule)[vs[1].(int)]
	}).(CustomRuleOutput)
}

type CustomRuleList struct {
	Rules []CustomRule `pulumi:"rules"`
}





type CustomRuleListInput interface {
	pulumi.Input

	ToCustomRuleListOutput() CustomRuleListOutput
	ToCustomRuleListOutputWithContext(context.Context) CustomRuleListOutput
}

type CustomRuleListArgs struct {
	Rules CustomRuleArrayInput `pulumi:"rules"`
}

func (CustomRuleListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleList)(nil)).Elem()
}

func (i CustomRuleListArgs) ToCustomRuleListOutput() CustomRuleListOutput {
	return i.ToCustomRuleListOutputWithContext(context.Background())
}

func (i CustomRuleListArgs) ToCustomRuleListOutputWithContext(ctx context.Context) CustomRuleListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListOutput)
}

func (i CustomRuleListArgs) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return i.ToCustomRuleListPtrOutputWithContext(context.Background())
}

func (i CustomRuleListArgs) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListOutput).ToCustomRuleListPtrOutputWithContext(ctx)
}









type CustomRuleListPtrInput interface {
	pulumi.Input

	ToCustomRuleListPtrOutput() CustomRuleListPtrOutput
	ToCustomRuleListPtrOutputWithContext(context.Context) CustomRuleListPtrOutput
}

type customRuleListPtrType CustomRuleListArgs

func CustomRuleListPtr(v *CustomRuleListArgs) CustomRuleListPtrInput {
	return (*customRuleListPtrType)(v)
}

func (*customRuleListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRuleList)(nil)).Elem()
}

func (i *customRuleListPtrType) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return i.ToCustomRuleListPtrOutputWithContext(context.Background())
}

func (i *customRuleListPtrType) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListPtrOutput)
}

type CustomRuleListOutput struct{ *pulumi.OutputState }

func (CustomRuleListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleList)(nil)).Elem()
}

func (o CustomRuleListOutput) ToCustomRuleListOutput() CustomRuleListOutput {
	return o
}

func (o CustomRuleListOutput) ToCustomRuleListOutputWithContext(ctx context.Context) CustomRuleListOutput {
	return o
}

func (o CustomRuleListOutput) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return o.ToCustomRuleListPtrOutputWithContext(context.Background())
}

func (o CustomRuleListOutput) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomRuleList) *CustomRuleList {
		return &v
	}).(CustomRuleListPtrOutput)
}

func (o CustomRuleListOutput) Rules() CustomRuleArrayOutput {
	return o.ApplyT(func(v CustomRuleList) []CustomRule { return v.Rules }).(CustomRuleArrayOutput)
}

type CustomRuleListPtrOutput struct{ *pulumi.OutputState }

func (CustomRuleListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRuleList)(nil)).Elem()
}

func (o CustomRuleListPtrOutput) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return o
}

func (o CustomRuleListPtrOutput) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return o
}

func (o CustomRuleListPtrOutput) Elem() CustomRuleListOutput {
	return o.ApplyT(func(v *CustomRuleList) CustomRuleList {
		if v != nil {
			return *v
		}
		var ret CustomRuleList
		return ret
	}).(CustomRuleListOutput)
}

func (o CustomRuleListPtrOutput) Rules() CustomRuleArrayOutput {
	return o.ApplyT(func(v *CustomRuleList) []CustomRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(CustomRuleArrayOutput)
}

type CustomRuleListResponse struct {
	Rules []CustomRuleResponse `pulumi:"rules"`
}

type CustomRuleListResponseOutput struct{ *pulumi.OutputState }

func (CustomRuleListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleListResponse)(nil)).Elem()
}

func (o CustomRuleListResponseOutput) ToCustomRuleListResponseOutput() CustomRuleListResponseOutput {
	return o
}

func (o CustomRuleListResponseOutput) ToCustomRuleListResponseOutputWithContext(ctx context.Context) CustomRuleListResponseOutput {
	return o
}

func (o CustomRuleListResponseOutput) Rules() CustomRuleResponseArrayOutput {
	return o.ApplyT(func(v CustomRuleListResponse) []CustomRuleResponse { return v.Rules }).(CustomRuleResponseArrayOutput)
}

type CustomRuleListResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomRuleListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRuleListResponse)(nil)).Elem()
}

func (o CustomRuleListResponsePtrOutput) ToCustomRuleListResponsePtrOutput() CustomRuleListResponsePtrOutput {
	return o
}

func (o CustomRuleListResponsePtrOutput) ToCustomRuleListResponsePtrOutputWithContext(ctx context.Context) CustomRuleListResponsePtrOutput {
	return o
}

func (o CustomRuleListResponsePtrOutput) Elem() CustomRuleListResponseOutput {
	return o.ApplyT(func(v *CustomRuleListResponse) CustomRuleListResponse {
		if v != nil {
			return *v
		}
		var ret CustomRuleListResponse
		return ret
	}).(CustomRuleListResponseOutput)
}

func (o CustomRuleListResponsePtrOutput) Rules() CustomRuleResponseArrayOutput {
	return o.ApplyT(func(v *CustomRuleListResponse) []CustomRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(CustomRuleResponseArrayOutput)
}

type CustomRuleResponse struct {
	Action          string                   `pulumi:"action"`
	EnabledState    *string                  `pulumi:"enabledState"`
	MatchConditions []MatchConditionResponse `pulumi:"matchConditions"`
	Name            string                   `pulumi:"name"`
	Priority        int                      `pulumi:"priority"`
}

type CustomRuleResponseOutput struct{ *pulumi.OutputState }

func (CustomRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleResponse)(nil)).Elem()
}

func (o CustomRuleResponseOutput) ToCustomRuleResponseOutput() CustomRuleResponseOutput {
	return o
}

func (o CustomRuleResponseOutput) ToCustomRuleResponseOutputWithContext(ctx context.Context) CustomRuleResponseOutput {
	return o
}

func (o CustomRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o CustomRuleResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRuleResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o CustomRuleResponseOutput) MatchConditions() MatchConditionResponseArrayOutput {
	return o.ApplyT(func(v CustomRuleResponse) []MatchConditionResponse { return v.MatchConditions }).(MatchConditionResponseArrayOutput)
}

func (o CustomRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v CustomRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

type CustomRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (CustomRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRuleResponse)(nil)).Elem()
}

func (o CustomRuleResponseArrayOutput) ToCustomRuleResponseArrayOutput() CustomRuleResponseArrayOutput {
	return o
}

func (o CustomRuleResponseArrayOutput) ToCustomRuleResponseArrayOutputWithContext(ctx context.Context) CustomRuleResponseArrayOutput {
	return o
}

func (o CustomRuleResponseArrayOutput) Index(i pulumi.IntInput) CustomRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRuleResponse {
		return vs[0].([]CustomRuleResponse)[vs[1].(int)]
	}).(CustomRuleResponseOutput)
}

type DeepCreatedOrigin struct {
	HostName  string `pulumi:"hostName"`
	HttpPort  *int   `pulumi:"httpPort"`
	HttpsPort *int   `pulumi:"httpsPort"`
	Name      string `pulumi:"name"`
}





type DeepCreatedOriginInput interface {
	pulumi.Input

	ToDeepCreatedOriginOutput() DeepCreatedOriginOutput
	ToDeepCreatedOriginOutputWithContext(context.Context) DeepCreatedOriginOutput
}

type DeepCreatedOriginArgs struct {
	HostName  pulumi.StringInput `pulumi:"hostName"`
	HttpPort  pulumi.IntPtrInput `pulumi:"httpPort"`
	HttpsPort pulumi.IntPtrInput `pulumi:"httpsPort"`
	Name      pulumi.StringInput `pulumi:"name"`
}

func (DeepCreatedOriginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOrigin)(nil)).Elem()
}

func (i DeepCreatedOriginArgs) ToDeepCreatedOriginOutput() DeepCreatedOriginOutput {
	return i.ToDeepCreatedOriginOutputWithContext(context.Background())
}

func (i DeepCreatedOriginArgs) ToDeepCreatedOriginOutputWithContext(ctx context.Context) DeepCreatedOriginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginOutput)
}





type DeepCreatedOriginArrayInput interface {
	pulumi.Input

	ToDeepCreatedOriginArrayOutput() DeepCreatedOriginArrayOutput
	ToDeepCreatedOriginArrayOutputWithContext(context.Context) DeepCreatedOriginArrayOutput
}

type DeepCreatedOriginArray []DeepCreatedOriginInput

func (DeepCreatedOriginArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOrigin)(nil)).Elem()
}

func (i DeepCreatedOriginArray) ToDeepCreatedOriginArrayOutput() DeepCreatedOriginArrayOutput {
	return i.ToDeepCreatedOriginArrayOutputWithContext(context.Background())
}

func (i DeepCreatedOriginArray) ToDeepCreatedOriginArrayOutputWithContext(ctx context.Context) DeepCreatedOriginArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginArrayOutput)
}

type DeepCreatedOriginOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOrigin)(nil)).Elem()
}

func (o DeepCreatedOriginOutput) ToDeepCreatedOriginOutput() DeepCreatedOriginOutput {
	return o
}

func (o DeepCreatedOriginOutput) ToDeepCreatedOriginOutputWithContext(ctx context.Context) DeepCreatedOriginOutput {
	return o
}

func (o DeepCreatedOriginOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) string { return v.HostName }).(pulumi.StringOutput)
}

func (o DeepCreatedOriginOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *int { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o DeepCreatedOriginOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *int { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o DeepCreatedOriginOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) string { return v.Name }).(pulumi.StringOutput)
}

type DeepCreatedOriginArrayOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOrigin)(nil)).Elem()
}

func (o DeepCreatedOriginArrayOutput) ToDeepCreatedOriginArrayOutput() DeepCreatedOriginArrayOutput {
	return o
}

func (o DeepCreatedOriginArrayOutput) ToDeepCreatedOriginArrayOutputWithContext(ctx context.Context) DeepCreatedOriginArrayOutput {
	return o
}

func (o DeepCreatedOriginArrayOutput) Index(i pulumi.IntInput) DeepCreatedOriginOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeepCreatedOrigin {
		return vs[0].([]DeepCreatedOrigin)[vs[1].(int)]
	}).(DeepCreatedOriginOutput)
}

type DeepCreatedOriginResponse struct {
	HostName  string `pulumi:"hostName"`
	HttpPort  *int   `pulumi:"httpPort"`
	HttpsPort *int   `pulumi:"httpsPort"`
	Name      string `pulumi:"name"`
}

type DeepCreatedOriginResponseOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOriginResponse)(nil)).Elem()
}

func (o DeepCreatedOriginResponseOutput) ToDeepCreatedOriginResponseOutput() DeepCreatedOriginResponseOutput {
	return o
}

func (o DeepCreatedOriginResponseOutput) ToDeepCreatedOriginResponseOutputWithContext(ctx context.Context) DeepCreatedOriginResponseOutput {
	return o
}

func (o DeepCreatedOriginResponseOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) string { return v.HostName }).(pulumi.StringOutput)
}

func (o DeepCreatedOriginResponseOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *int { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o DeepCreatedOriginResponseOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *int { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o DeepCreatedOriginResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) string { return v.Name }).(pulumi.StringOutput)
}

type DeepCreatedOriginResponseArrayOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOriginResponse)(nil)).Elem()
}

func (o DeepCreatedOriginResponseArrayOutput) ToDeepCreatedOriginResponseArrayOutput() DeepCreatedOriginResponseArrayOutput {
	return o
}

func (o DeepCreatedOriginResponseArrayOutput) ToDeepCreatedOriginResponseArrayOutputWithContext(ctx context.Context) DeepCreatedOriginResponseArrayOutput {
	return o
}

func (o DeepCreatedOriginResponseArrayOutput) Index(i pulumi.IntInput) DeepCreatedOriginResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeepCreatedOriginResponse {
		return vs[0].([]DeepCreatedOriginResponse)[vs[1].(int)]
	}).(DeepCreatedOriginResponseOutput)
}

type DeliveryRule struct {
	Actions    []interface{} `pulumi:"actions"`
	Conditions []interface{} `pulumi:"conditions"`
	Name       *string       `pulumi:"name"`
	Order      int           `pulumi:"order"`
}





type DeliveryRuleInput interface {
	pulumi.Input

	ToDeliveryRuleOutput() DeliveryRuleOutput
	ToDeliveryRuleOutputWithContext(context.Context) DeliveryRuleOutput
}

type DeliveryRuleArgs struct {
	Actions    pulumi.ArrayInput     `pulumi:"actions"`
	Conditions pulumi.ArrayInput     `pulumi:"conditions"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Order      pulumi.IntInput       `pulumi:"order"`
}

func (DeliveryRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRule)(nil)).Elem()
}

func (i DeliveryRuleArgs) ToDeliveryRuleOutput() DeliveryRuleOutput {
	return i.ToDeliveryRuleOutputWithContext(context.Background())
}

func (i DeliveryRuleArgs) ToDeliveryRuleOutputWithContext(ctx context.Context) DeliveryRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleOutput)
}





type DeliveryRuleArrayInput interface {
	pulumi.Input

	ToDeliveryRuleArrayOutput() DeliveryRuleArrayOutput
	ToDeliveryRuleArrayOutputWithContext(context.Context) DeliveryRuleArrayOutput
}

type DeliveryRuleArray []DeliveryRuleInput

func (DeliveryRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeliveryRule)(nil)).Elem()
}

func (i DeliveryRuleArray) ToDeliveryRuleArrayOutput() DeliveryRuleArrayOutput {
	return i.ToDeliveryRuleArrayOutputWithContext(context.Background())
}

func (i DeliveryRuleArray) ToDeliveryRuleArrayOutputWithContext(ctx context.Context) DeliveryRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleArrayOutput)
}

type DeliveryRuleOutput struct{ *pulumi.OutputState }

func (DeliveryRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRule)(nil)).Elem()
}

func (o DeliveryRuleOutput) ToDeliveryRuleOutput() DeliveryRuleOutput {
	return o
}

func (o DeliveryRuleOutput) ToDeliveryRuleOutputWithContext(ctx context.Context) DeliveryRuleOutput {
	return o
}

func (o DeliveryRuleOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v DeliveryRule) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o DeliveryRuleOutput) Conditions() pulumi.ArrayOutput {
	return o.ApplyT(func(v DeliveryRule) []interface{} { return v.Conditions }).(pulumi.ArrayOutput)
}

func (o DeliveryRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeliveryRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DeliveryRuleOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v DeliveryRule) int { return v.Order }).(pulumi.IntOutput)
}

type DeliveryRuleArrayOutput struct{ *pulumi.OutputState }

func (DeliveryRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeliveryRule)(nil)).Elem()
}

func (o DeliveryRuleArrayOutput) ToDeliveryRuleArrayOutput() DeliveryRuleArrayOutput {
	return o
}

func (o DeliveryRuleArrayOutput) ToDeliveryRuleArrayOutputWithContext(ctx context.Context) DeliveryRuleArrayOutput {
	return o
}

func (o DeliveryRuleArrayOutput) Index(i pulumi.IntInput) DeliveryRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeliveryRule {
		return vs[0].([]DeliveryRule)[vs[1].(int)]
	}).(DeliveryRuleOutput)
}

type DeliveryRuleCacheExpirationAction struct {
	Name       string                          `pulumi:"name"`
	Parameters CacheExpirationActionParameters `pulumi:"parameters"`
}

type DeliveryRuleCacheExpirationActionResponse struct {
	Name       string                                  `pulumi:"name"`
	Parameters CacheExpirationActionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleCacheKeyQueryStringAction struct {
	Name       string                              `pulumi:"name"`
	Parameters CacheKeyQueryStringActionParameters `pulumi:"parameters"`
}

type DeliveryRuleCacheKeyQueryStringActionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters CacheKeyQueryStringActionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleCookiesCondition struct {
	Name       string                          `pulumi:"name"`
	Parameters CookiesMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleCookiesConditionResponse struct {
	Name       string                                  `pulumi:"name"`
	Parameters CookiesMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleHttpVersionCondition struct {
	Name       string                              `pulumi:"name"`
	Parameters HttpVersionMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleHttpVersionConditionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters HttpVersionMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleIsDeviceCondition struct {
	Name       string                           `pulumi:"name"`
	Parameters IsDeviceMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleIsDeviceConditionResponse struct {
	Name       string                                   `pulumi:"name"`
	Parameters IsDeviceMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRulePostArgsCondition struct {
	Name       string                           `pulumi:"name"`
	Parameters PostArgsMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRulePostArgsConditionResponse struct {
	Name       string                                   `pulumi:"name"`
	Parameters PostArgsMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleQueryStringCondition struct {
	Name       string                              `pulumi:"name"`
	Parameters QueryStringMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleQueryStringConditionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters QueryStringMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleRemoteAddressCondition struct {
	Name       string                                `pulumi:"name"`
	Parameters RemoteAddressMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleRemoteAddressConditionResponse struct {
	Name       string                                        `pulumi:"name"`
	Parameters RemoteAddressMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleRequestBodyCondition struct {
	Name       string                              `pulumi:"name"`
	Parameters RequestBodyMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleRequestBodyConditionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters RequestBodyMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleRequestHeaderAction struct {
	Name       string                 `pulumi:"name"`
	Parameters HeaderActionParameters `pulumi:"parameters"`
}

type DeliveryRuleRequestHeaderActionResponse struct {
	Name       string                         `pulumi:"name"`
	Parameters HeaderActionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleRequestHeaderCondition struct {
	Name       string                                `pulumi:"name"`
	Parameters RequestHeaderMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleRequestHeaderConditionResponse struct {
	Name       string                                        `pulumi:"name"`
	Parameters RequestHeaderMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleRequestMethodCondition struct {
	Name       string                                `pulumi:"name"`
	Parameters RequestMethodMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleRequestMethodConditionResponse struct {
	Name       string                                        `pulumi:"name"`
	Parameters RequestMethodMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleRequestSchemeCondition struct {
	Name       string                                `pulumi:"name"`
	Parameters RequestSchemeMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleRequestSchemeConditionResponse struct {
	Name       string                                        `pulumi:"name"`
	Parameters RequestSchemeMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleRequestUriCondition struct {
	Name       string                             `pulumi:"name"`
	Parameters RequestUriMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleRequestUriConditionResponse struct {
	Name       string                                     `pulumi:"name"`
	Parameters RequestUriMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleResponse struct {
	Actions    []interface{} `pulumi:"actions"`
	Conditions []interface{} `pulumi:"conditions"`
	Name       *string       `pulumi:"name"`
	Order      int           `pulumi:"order"`
}

type DeliveryRuleResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleResponse)(nil)).Elem()
}

func (o DeliveryRuleResponseOutput) ToDeliveryRuleResponseOutput() DeliveryRuleResponseOutput {
	return o
}

func (o DeliveryRuleResponseOutput) ToDeliveryRuleResponseOutputWithContext(ctx context.Context) DeliveryRuleResponseOutput {
	return o
}

func (o DeliveryRuleResponseOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v DeliveryRuleResponse) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o DeliveryRuleResponseOutput) Conditions() pulumi.ArrayOutput {
	return o.ApplyT(func(v DeliveryRuleResponse) []interface{} { return v.Conditions }).(pulumi.ArrayOutput)
}

func (o DeliveryRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeliveryRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DeliveryRuleResponseOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v DeliveryRuleResponse) int { return v.Order }).(pulumi.IntOutput)
}

type DeliveryRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (DeliveryRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeliveryRuleResponse)(nil)).Elem()
}

func (o DeliveryRuleResponseArrayOutput) ToDeliveryRuleResponseArrayOutput() DeliveryRuleResponseArrayOutput {
	return o
}

func (o DeliveryRuleResponseArrayOutput) ToDeliveryRuleResponseArrayOutputWithContext(ctx context.Context) DeliveryRuleResponseArrayOutput {
	return o
}

func (o DeliveryRuleResponseArrayOutput) Index(i pulumi.IntInput) DeliveryRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeliveryRuleResponse {
		return vs[0].([]DeliveryRuleResponse)[vs[1].(int)]
	}).(DeliveryRuleResponseOutput)
}

type DeliveryRuleResponseHeaderAction struct {
	Name       string                 `pulumi:"name"`
	Parameters HeaderActionParameters `pulumi:"parameters"`
}

type DeliveryRuleResponseHeaderActionResponse struct {
	Name       string                         `pulumi:"name"`
	Parameters HeaderActionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleUrlFileExtensionCondition struct {
	Name       string                                   `pulumi:"name"`
	Parameters UrlFileExtensionMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleUrlFileExtensionConditionResponse struct {
	Name       string                                           `pulumi:"name"`
	Parameters UrlFileExtensionMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleUrlFileNameCondition struct {
	Name       string                              `pulumi:"name"`
	Parameters UrlFileNameMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleUrlFileNameConditionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters UrlFileNameMatchConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleUrlPathCondition struct {
	Name       string                          `pulumi:"name"`
	Parameters UrlPathMatchConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleUrlPathConditionResponse struct {
	Name       string                                  `pulumi:"name"`
	Parameters UrlPathMatchConditionParametersResponse `pulumi:"parameters"`
}

type EndpointPropertiesUpdateParametersDeliveryPolicy struct {
	Description *string        `pulumi:"description"`
	Rules       []DeliveryRule `pulumi:"rules"`
}





type EndpointPropertiesUpdateParametersDeliveryPolicyInput interface {
	pulumi.Input

	ToEndpointPropertiesUpdateParametersDeliveryPolicyOutput() EndpointPropertiesUpdateParametersDeliveryPolicyOutput
	ToEndpointPropertiesUpdateParametersDeliveryPolicyOutputWithContext(context.Context) EndpointPropertiesUpdateParametersDeliveryPolicyOutput
}

type EndpointPropertiesUpdateParametersDeliveryPolicyArgs struct {
	Description pulumi.StringPtrInput  `pulumi:"description"`
	Rules       DeliveryRuleArrayInput `pulumi:"rules"`
}

func (EndpointPropertiesUpdateParametersDeliveryPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesUpdateParametersDeliveryPolicy)(nil)).Elem()
}

func (i EndpointPropertiesUpdateParametersDeliveryPolicyArgs) ToEndpointPropertiesUpdateParametersDeliveryPolicyOutput() EndpointPropertiesUpdateParametersDeliveryPolicyOutput {
	return i.ToEndpointPropertiesUpdateParametersDeliveryPolicyOutputWithContext(context.Background())
}

func (i EndpointPropertiesUpdateParametersDeliveryPolicyArgs) ToEndpointPropertiesUpdateParametersDeliveryPolicyOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersDeliveryPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersDeliveryPolicyOutput)
}

func (i EndpointPropertiesUpdateParametersDeliveryPolicyArgs) ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput() EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput {
	return i.ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutputWithContext(context.Background())
}

func (i EndpointPropertiesUpdateParametersDeliveryPolicyArgs) ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersDeliveryPolicyOutput).ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutputWithContext(ctx)
}









type EndpointPropertiesUpdateParametersDeliveryPolicyPtrInput interface {
	pulumi.Input

	ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput() EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput
	ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutputWithContext(context.Context) EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput
}

type endpointPropertiesUpdateParametersDeliveryPolicyPtrType EndpointPropertiesUpdateParametersDeliveryPolicyArgs

func EndpointPropertiesUpdateParametersDeliveryPolicyPtr(v *EndpointPropertiesUpdateParametersDeliveryPolicyArgs) EndpointPropertiesUpdateParametersDeliveryPolicyPtrInput {
	return (*endpointPropertiesUpdateParametersDeliveryPolicyPtrType)(v)
}

func (*endpointPropertiesUpdateParametersDeliveryPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointPropertiesUpdateParametersDeliveryPolicy)(nil)).Elem()
}

func (i *endpointPropertiesUpdateParametersDeliveryPolicyPtrType) ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput() EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput {
	return i.ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutputWithContext(context.Background())
}

func (i *endpointPropertiesUpdateParametersDeliveryPolicyPtrType) ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput)
}

type EndpointPropertiesUpdateParametersDeliveryPolicyOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesUpdateParametersDeliveryPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesUpdateParametersDeliveryPolicy)(nil)).Elem()
}

func (o EndpointPropertiesUpdateParametersDeliveryPolicyOutput) ToEndpointPropertiesUpdateParametersDeliveryPolicyOutput() EndpointPropertiesUpdateParametersDeliveryPolicyOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersDeliveryPolicyOutput) ToEndpointPropertiesUpdateParametersDeliveryPolicyOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersDeliveryPolicyOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersDeliveryPolicyOutput) ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput() EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput {
	return o.ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutputWithContext(context.Background())
}

func (o EndpointPropertiesUpdateParametersDeliveryPolicyOutput) ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointPropertiesUpdateParametersDeliveryPolicy) *EndpointPropertiesUpdateParametersDeliveryPolicy {
		return &v
	}).(EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput)
}

func (o EndpointPropertiesUpdateParametersDeliveryPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesUpdateParametersDeliveryPolicy) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesUpdateParametersDeliveryPolicyOutput) Rules() DeliveryRuleArrayOutput {
	return o.ApplyT(func(v EndpointPropertiesUpdateParametersDeliveryPolicy) []DeliveryRule { return v.Rules }).(DeliveryRuleArrayOutput)
}

type EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointPropertiesUpdateParametersDeliveryPolicy)(nil)).Elem()
}

func (o EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput) ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput() EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput) ToEndpointPropertiesUpdateParametersDeliveryPolicyPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput) Elem() EndpointPropertiesUpdateParametersDeliveryPolicyOutput {
	return o.ApplyT(func(v *EndpointPropertiesUpdateParametersDeliveryPolicy) EndpointPropertiesUpdateParametersDeliveryPolicy {
		if v != nil {
			return *v
		}
		var ret EndpointPropertiesUpdateParametersDeliveryPolicy
		return ret
	}).(EndpointPropertiesUpdateParametersDeliveryPolicyOutput)
}

func (o EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointPropertiesUpdateParametersDeliveryPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput) Rules() DeliveryRuleArrayOutput {
	return o.ApplyT(func(v *EndpointPropertiesUpdateParametersDeliveryPolicy) []DeliveryRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(DeliveryRuleArrayOutput)
}

type EndpointPropertiesUpdateParametersResponseDeliveryPolicy struct {
	Description *string                `pulumi:"description"`
	Rules       []DeliveryRuleResponse `pulumi:"rules"`
}

type EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesUpdateParametersResponseDeliveryPolicy)(nil)).Elem()
}

func (o EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput() EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesUpdateParametersResponseDeliveryPolicy) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput) Rules() DeliveryRuleResponseArrayOutput {
	return o.ApplyT(func(v EndpointPropertiesUpdateParametersResponseDeliveryPolicy) []DeliveryRuleResponse {
		return v.Rules
	}).(DeliveryRuleResponseArrayOutput)
}

type EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointPropertiesUpdateParametersResponseDeliveryPolicy)(nil)).Elem()
}

func (o EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput() EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput) Elem() EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput {
	return o.ApplyT(func(v *EndpointPropertiesUpdateParametersResponseDeliveryPolicy) EndpointPropertiesUpdateParametersResponseDeliveryPolicy {
		if v != nil {
			return *v
		}
		var ret EndpointPropertiesUpdateParametersResponseDeliveryPolicy
		return ret
	}).(EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput)
}

func (o EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointPropertiesUpdateParametersResponseDeliveryPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput) Rules() DeliveryRuleResponseArrayOutput {
	return o.ApplyT(func(v *EndpointPropertiesUpdateParametersResponseDeliveryPolicy) []DeliveryRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(DeliveryRuleResponseArrayOutput)
}

type EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink struct {
	Id *string `pulumi:"id"`
}

type EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink) *string {
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) Elem() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o.ApplyT(func(v *EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink) EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink {
		if v != nil {
			return *v
		}
		var ret EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink
		return ret
	}).(EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput)
}

func (o EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink struct {
	Id *string `pulumi:"id"`
}





type EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkInput interface {
	pulumi.Input

	ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput() EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput
	ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(context.Context) EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput
}

type EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput() EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return i.ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(context.Background())
}

func (i EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput)
}

func (i EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput).ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx)
}









type EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrInput interface {
	pulumi.Input

	ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput
	ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Context) EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput
}

type endpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrType EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkArgs

func EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtr(v *EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkArgs) EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrInput {
	return (*endpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrType)(v)
}

func (*endpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i *endpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i *endpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

type EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput() EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (o EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink) *EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink {
		return &v
	}).(EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

func (o EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ToEndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) Elem() EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o.ApplyT(func(v *EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink) EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink {
		if v != nil {
			return *v
		}
		var ret EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink
		return ret
	}).(EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput)
}

func (o EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type GeoFilter struct {
	Action       GeoFilterActions `pulumi:"action"`
	CountryCodes []string         `pulumi:"countryCodes"`
	RelativePath string           `pulumi:"relativePath"`
}





type GeoFilterInput interface {
	pulumi.Input

	ToGeoFilterOutput() GeoFilterOutput
	ToGeoFilterOutputWithContext(context.Context) GeoFilterOutput
}

type GeoFilterArgs struct {
	Action       GeoFilterActionsInput   `pulumi:"action"`
	CountryCodes pulumi.StringArrayInput `pulumi:"countryCodes"`
	RelativePath pulumi.StringInput      `pulumi:"relativePath"`
}

func (GeoFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoFilter)(nil)).Elem()
}

func (i GeoFilterArgs) ToGeoFilterOutput() GeoFilterOutput {
	return i.ToGeoFilterOutputWithContext(context.Background())
}

func (i GeoFilterArgs) ToGeoFilterOutputWithContext(ctx context.Context) GeoFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoFilterOutput)
}





type GeoFilterArrayInput interface {
	pulumi.Input

	ToGeoFilterArrayOutput() GeoFilterArrayOutput
	ToGeoFilterArrayOutputWithContext(context.Context) GeoFilterArrayOutput
}

type GeoFilterArray []GeoFilterInput

func (GeoFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GeoFilter)(nil)).Elem()
}

func (i GeoFilterArray) ToGeoFilterArrayOutput() GeoFilterArrayOutput {
	return i.ToGeoFilterArrayOutputWithContext(context.Background())
}

func (i GeoFilterArray) ToGeoFilterArrayOutputWithContext(ctx context.Context) GeoFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoFilterArrayOutput)
}

type GeoFilterOutput struct{ *pulumi.OutputState }

func (GeoFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoFilter)(nil)).Elem()
}

func (o GeoFilterOutput) ToGeoFilterOutput() GeoFilterOutput {
	return o
}

func (o GeoFilterOutput) ToGeoFilterOutputWithContext(ctx context.Context) GeoFilterOutput {
	return o
}

func (o GeoFilterOutput) Action() GeoFilterActionsOutput {
	return o.ApplyT(func(v GeoFilter) GeoFilterActions { return v.Action }).(GeoFilterActionsOutput)
}

func (o GeoFilterOutput) CountryCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GeoFilter) []string { return v.CountryCodes }).(pulumi.StringArrayOutput)
}

func (o GeoFilterOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v GeoFilter) string { return v.RelativePath }).(pulumi.StringOutput)
}

type GeoFilterArrayOutput struct{ *pulumi.OutputState }

func (GeoFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GeoFilter)(nil)).Elem()
}

func (o GeoFilterArrayOutput) ToGeoFilterArrayOutput() GeoFilterArrayOutput {
	return o
}

func (o GeoFilterArrayOutput) ToGeoFilterArrayOutputWithContext(ctx context.Context) GeoFilterArrayOutput {
	return o
}

func (o GeoFilterArrayOutput) Index(i pulumi.IntInput) GeoFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GeoFilter {
		return vs[0].([]GeoFilter)[vs[1].(int)]
	}).(GeoFilterOutput)
}

type GeoFilterResponse struct {
	Action       string   `pulumi:"action"`
	CountryCodes []string `pulumi:"countryCodes"`
	RelativePath string   `pulumi:"relativePath"`
}

type GeoFilterResponseOutput struct{ *pulumi.OutputState }

func (GeoFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoFilterResponse)(nil)).Elem()
}

func (o GeoFilterResponseOutput) ToGeoFilterResponseOutput() GeoFilterResponseOutput {
	return o
}

func (o GeoFilterResponseOutput) ToGeoFilterResponseOutputWithContext(ctx context.Context) GeoFilterResponseOutput {
	return o
}

func (o GeoFilterResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v GeoFilterResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o GeoFilterResponseOutput) CountryCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GeoFilterResponse) []string { return v.CountryCodes }).(pulumi.StringArrayOutput)
}

func (o GeoFilterResponseOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v GeoFilterResponse) string { return v.RelativePath }).(pulumi.StringOutput)
}

type GeoFilterResponseArrayOutput struct{ *pulumi.OutputState }

func (GeoFilterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GeoFilterResponse)(nil)).Elem()
}

func (o GeoFilterResponseArrayOutput) ToGeoFilterResponseArrayOutput() GeoFilterResponseArrayOutput {
	return o
}

func (o GeoFilterResponseArrayOutput) ToGeoFilterResponseArrayOutputWithContext(ctx context.Context) GeoFilterResponseArrayOutput {
	return o
}

func (o GeoFilterResponseArrayOutput) Index(i pulumi.IntInput) GeoFilterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GeoFilterResponse {
		return vs[0].([]GeoFilterResponse)[vs[1].(int)]
	}).(GeoFilterResponseOutput)
}

type HeaderActionParameters struct {
	HeaderAction string  `pulumi:"headerAction"`
	HeaderName   string  `pulumi:"headerName"`
	OdataType    string  `pulumi:"odataType"`
	Value        *string `pulumi:"value"`
}

type HeaderActionParametersResponse struct {
	HeaderAction string  `pulumi:"headerAction"`
	HeaderName   string  `pulumi:"headerName"`
	OdataType    string  `pulumi:"odataType"`
	Value        *string `pulumi:"value"`
}

type HttpVersionMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
}

type HttpVersionMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
}

type IsDeviceMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type IsDeviceMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type KeyVaultCertificateSourceParametersResponse struct {
	DeleteRule        string  `pulumi:"deleteRule"`
	OdataType         string  `pulumi:"odataType"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SecretName        string  `pulumi:"secretName"`
	SecretVersion     *string `pulumi:"secretVersion"`
	SubscriptionId    string  `pulumi:"subscriptionId"`
	UpdateRule        string  `pulumi:"updateRule"`
	VaultName         string  `pulumi:"vaultName"`
}

type ManagedRuleGroupOverride struct {
	RuleGroupName string                `pulumi:"ruleGroupName"`
	Rules         []ManagedRuleOverride `pulumi:"rules"`
}





type ManagedRuleGroupOverrideInput interface {
	pulumi.Input

	ToManagedRuleGroupOverrideOutput() ManagedRuleGroupOverrideOutput
	ToManagedRuleGroupOverrideOutputWithContext(context.Context) ManagedRuleGroupOverrideOutput
}

type ManagedRuleGroupOverrideArgs struct {
	RuleGroupName pulumi.StringInput            `pulumi:"ruleGroupName"`
	Rules         ManagedRuleOverrideArrayInput `pulumi:"rules"`
}

func (ManagedRuleGroupOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleGroupOverride)(nil)).Elem()
}

func (i ManagedRuleGroupOverrideArgs) ToManagedRuleGroupOverrideOutput() ManagedRuleGroupOverrideOutput {
	return i.ToManagedRuleGroupOverrideOutputWithContext(context.Background())
}

func (i ManagedRuleGroupOverrideArgs) ToManagedRuleGroupOverrideOutputWithContext(ctx context.Context) ManagedRuleGroupOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleGroupOverrideOutput)
}





type ManagedRuleGroupOverrideArrayInput interface {
	pulumi.Input

	ToManagedRuleGroupOverrideArrayOutput() ManagedRuleGroupOverrideArrayOutput
	ToManagedRuleGroupOverrideArrayOutputWithContext(context.Context) ManagedRuleGroupOverrideArrayOutput
}

type ManagedRuleGroupOverrideArray []ManagedRuleGroupOverrideInput

func (ManagedRuleGroupOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleGroupOverride)(nil)).Elem()
}

func (i ManagedRuleGroupOverrideArray) ToManagedRuleGroupOverrideArrayOutput() ManagedRuleGroupOverrideArrayOutput {
	return i.ToManagedRuleGroupOverrideArrayOutputWithContext(context.Background())
}

func (i ManagedRuleGroupOverrideArray) ToManagedRuleGroupOverrideArrayOutputWithContext(ctx context.Context) ManagedRuleGroupOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleGroupOverrideArrayOutput)
}

type ManagedRuleGroupOverrideOutput struct{ *pulumi.OutputState }

func (ManagedRuleGroupOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleGroupOverride)(nil)).Elem()
}

func (o ManagedRuleGroupOverrideOutput) ToManagedRuleGroupOverrideOutput() ManagedRuleGroupOverrideOutput {
	return o
}

func (o ManagedRuleGroupOverrideOutput) ToManagedRuleGroupOverrideOutputWithContext(ctx context.Context) ManagedRuleGroupOverrideOutput {
	return o
}

func (o ManagedRuleGroupOverrideOutput) RuleGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleGroupOverride) string { return v.RuleGroupName }).(pulumi.StringOutput)
}

func (o ManagedRuleGroupOverrideOutput) Rules() ManagedRuleOverrideArrayOutput {
	return o.ApplyT(func(v ManagedRuleGroupOverride) []ManagedRuleOverride { return v.Rules }).(ManagedRuleOverrideArrayOutput)
}

type ManagedRuleGroupOverrideArrayOutput struct{ *pulumi.OutputState }

func (ManagedRuleGroupOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleGroupOverride)(nil)).Elem()
}

func (o ManagedRuleGroupOverrideArrayOutput) ToManagedRuleGroupOverrideArrayOutput() ManagedRuleGroupOverrideArrayOutput {
	return o
}

func (o ManagedRuleGroupOverrideArrayOutput) ToManagedRuleGroupOverrideArrayOutputWithContext(ctx context.Context) ManagedRuleGroupOverrideArrayOutput {
	return o
}

func (o ManagedRuleGroupOverrideArrayOutput) Index(i pulumi.IntInput) ManagedRuleGroupOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedRuleGroupOverride {
		return vs[0].([]ManagedRuleGroupOverride)[vs[1].(int)]
	}).(ManagedRuleGroupOverrideOutput)
}

type ManagedRuleGroupOverrideResponse struct {
	RuleGroupName string                        `pulumi:"ruleGroupName"`
	Rules         []ManagedRuleOverrideResponse `pulumi:"rules"`
}

type ManagedRuleGroupOverrideResponseOutput struct{ *pulumi.OutputState }

func (ManagedRuleGroupOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleGroupOverrideResponse)(nil)).Elem()
}

func (o ManagedRuleGroupOverrideResponseOutput) ToManagedRuleGroupOverrideResponseOutput() ManagedRuleGroupOverrideResponseOutput {
	return o
}

func (o ManagedRuleGroupOverrideResponseOutput) ToManagedRuleGroupOverrideResponseOutputWithContext(ctx context.Context) ManagedRuleGroupOverrideResponseOutput {
	return o
}

func (o ManagedRuleGroupOverrideResponseOutput) RuleGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleGroupOverrideResponse) string { return v.RuleGroupName }).(pulumi.StringOutput)
}

func (o ManagedRuleGroupOverrideResponseOutput) Rules() ManagedRuleOverrideResponseArrayOutput {
	return o.ApplyT(func(v ManagedRuleGroupOverrideResponse) []ManagedRuleOverrideResponse { return v.Rules }).(ManagedRuleOverrideResponseArrayOutput)
}

type ManagedRuleGroupOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedRuleGroupOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleGroupOverrideResponse)(nil)).Elem()
}

func (o ManagedRuleGroupOverrideResponseArrayOutput) ToManagedRuleGroupOverrideResponseArrayOutput() ManagedRuleGroupOverrideResponseArrayOutput {
	return o
}

func (o ManagedRuleGroupOverrideResponseArrayOutput) ToManagedRuleGroupOverrideResponseArrayOutputWithContext(ctx context.Context) ManagedRuleGroupOverrideResponseArrayOutput {
	return o
}

func (o ManagedRuleGroupOverrideResponseArrayOutput) Index(i pulumi.IntInput) ManagedRuleGroupOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedRuleGroupOverrideResponse {
		return vs[0].([]ManagedRuleGroupOverrideResponse)[vs[1].(int)]
	}).(ManagedRuleGroupOverrideResponseOutput)
}

type ManagedRuleOverride struct {
	Action       *string `pulumi:"action"`
	EnabledState *string `pulumi:"enabledState"`
	RuleId       string  `pulumi:"ruleId"`
}





type ManagedRuleOverrideInput interface {
	pulumi.Input

	ToManagedRuleOverrideOutput() ManagedRuleOverrideOutput
	ToManagedRuleOverrideOutputWithContext(context.Context) ManagedRuleOverrideOutput
}

type ManagedRuleOverrideArgs struct {
	Action       pulumi.StringPtrInput `pulumi:"action"`
	EnabledState pulumi.StringPtrInput `pulumi:"enabledState"`
	RuleId       pulumi.StringInput    `pulumi:"ruleId"`
}

func (ManagedRuleOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleOverride)(nil)).Elem()
}

func (i ManagedRuleOverrideArgs) ToManagedRuleOverrideOutput() ManagedRuleOverrideOutput {
	return i.ToManagedRuleOverrideOutputWithContext(context.Background())
}

func (i ManagedRuleOverrideArgs) ToManagedRuleOverrideOutputWithContext(ctx context.Context) ManagedRuleOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleOverrideOutput)
}





type ManagedRuleOverrideArrayInput interface {
	pulumi.Input

	ToManagedRuleOverrideArrayOutput() ManagedRuleOverrideArrayOutput
	ToManagedRuleOverrideArrayOutputWithContext(context.Context) ManagedRuleOverrideArrayOutput
}

type ManagedRuleOverrideArray []ManagedRuleOverrideInput

func (ManagedRuleOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleOverride)(nil)).Elem()
}

func (i ManagedRuleOverrideArray) ToManagedRuleOverrideArrayOutput() ManagedRuleOverrideArrayOutput {
	return i.ToManagedRuleOverrideArrayOutputWithContext(context.Background())
}

func (i ManagedRuleOverrideArray) ToManagedRuleOverrideArrayOutputWithContext(ctx context.Context) ManagedRuleOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleOverrideArrayOutput)
}

type ManagedRuleOverrideOutput struct{ *pulumi.OutputState }

func (ManagedRuleOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleOverride)(nil)).Elem()
}

func (o ManagedRuleOverrideOutput) ToManagedRuleOverrideOutput() ManagedRuleOverrideOutput {
	return o
}

func (o ManagedRuleOverrideOutput) ToManagedRuleOverrideOutputWithContext(ctx context.Context) ManagedRuleOverrideOutput {
	return o
}

func (o ManagedRuleOverrideOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedRuleOverride) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o ManagedRuleOverrideOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedRuleOverride) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o ManagedRuleOverrideOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleOverride) string { return v.RuleId }).(pulumi.StringOutput)
}

type ManagedRuleOverrideArrayOutput struct{ *pulumi.OutputState }

func (ManagedRuleOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleOverride)(nil)).Elem()
}

func (o ManagedRuleOverrideArrayOutput) ToManagedRuleOverrideArrayOutput() ManagedRuleOverrideArrayOutput {
	return o
}

func (o ManagedRuleOverrideArrayOutput) ToManagedRuleOverrideArrayOutputWithContext(ctx context.Context) ManagedRuleOverrideArrayOutput {
	return o
}

func (o ManagedRuleOverrideArrayOutput) Index(i pulumi.IntInput) ManagedRuleOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedRuleOverride {
		return vs[0].([]ManagedRuleOverride)[vs[1].(int)]
	}).(ManagedRuleOverrideOutput)
}

type ManagedRuleOverrideResponse struct {
	Action       *string `pulumi:"action"`
	EnabledState *string `pulumi:"enabledState"`
	RuleId       string  `pulumi:"ruleId"`
}

type ManagedRuleOverrideResponseOutput struct{ *pulumi.OutputState }

func (ManagedRuleOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleOverrideResponse)(nil)).Elem()
}

func (o ManagedRuleOverrideResponseOutput) ToManagedRuleOverrideResponseOutput() ManagedRuleOverrideResponseOutput {
	return o
}

func (o ManagedRuleOverrideResponseOutput) ToManagedRuleOverrideResponseOutputWithContext(ctx context.Context) ManagedRuleOverrideResponseOutput {
	return o
}

func (o ManagedRuleOverrideResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedRuleOverrideResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o ManagedRuleOverrideResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedRuleOverrideResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o ManagedRuleOverrideResponseOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleOverrideResponse) string { return v.RuleId }).(pulumi.StringOutput)
}

type ManagedRuleOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedRuleOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleOverrideResponse)(nil)).Elem()
}

func (o ManagedRuleOverrideResponseArrayOutput) ToManagedRuleOverrideResponseArrayOutput() ManagedRuleOverrideResponseArrayOutput {
	return o
}

func (o ManagedRuleOverrideResponseArrayOutput) ToManagedRuleOverrideResponseArrayOutputWithContext(ctx context.Context) ManagedRuleOverrideResponseArrayOutput {
	return o
}

func (o ManagedRuleOverrideResponseArrayOutput) Index(i pulumi.IntInput) ManagedRuleOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedRuleOverrideResponse {
		return vs[0].([]ManagedRuleOverrideResponse)[vs[1].(int)]
	}).(ManagedRuleOverrideResponseOutput)
}

type ManagedRuleSet struct {
	AnomalyScore       *int                       `pulumi:"anomalyScore"`
	RuleGroupOverrides []ManagedRuleGroupOverride `pulumi:"ruleGroupOverrides"`
	RuleSetType        string                     `pulumi:"ruleSetType"`
	RuleSetVersion     string                     `pulumi:"ruleSetVersion"`
}





type ManagedRuleSetInput interface {
	pulumi.Input

	ToManagedRuleSetOutput() ManagedRuleSetOutput
	ToManagedRuleSetOutputWithContext(context.Context) ManagedRuleSetOutput
}

type ManagedRuleSetArgs struct {
	AnomalyScore       pulumi.IntPtrInput                 `pulumi:"anomalyScore"`
	RuleGroupOverrides ManagedRuleGroupOverrideArrayInput `pulumi:"ruleGroupOverrides"`
	RuleSetType        pulumi.StringInput                 `pulumi:"ruleSetType"`
	RuleSetVersion     pulumi.StringInput                 `pulumi:"ruleSetVersion"`
}

func (ManagedRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSet)(nil)).Elem()
}

func (i ManagedRuleSetArgs) ToManagedRuleSetOutput() ManagedRuleSetOutput {
	return i.ToManagedRuleSetOutputWithContext(context.Background())
}

func (i ManagedRuleSetArgs) ToManagedRuleSetOutputWithContext(ctx context.Context) ManagedRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetOutput)
}





type ManagedRuleSetArrayInput interface {
	pulumi.Input

	ToManagedRuleSetArrayOutput() ManagedRuleSetArrayOutput
	ToManagedRuleSetArrayOutputWithContext(context.Context) ManagedRuleSetArrayOutput
}

type ManagedRuleSetArray []ManagedRuleSetInput

func (ManagedRuleSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleSet)(nil)).Elem()
}

func (i ManagedRuleSetArray) ToManagedRuleSetArrayOutput() ManagedRuleSetArrayOutput {
	return i.ToManagedRuleSetArrayOutputWithContext(context.Background())
}

func (i ManagedRuleSetArray) ToManagedRuleSetArrayOutputWithContext(ctx context.Context) ManagedRuleSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetArrayOutput)
}

type ManagedRuleSetOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSet)(nil)).Elem()
}

func (o ManagedRuleSetOutput) ToManagedRuleSetOutput() ManagedRuleSetOutput {
	return o
}

func (o ManagedRuleSetOutput) ToManagedRuleSetOutputWithContext(ctx context.Context) ManagedRuleSetOutput {
	return o
}

func (o ManagedRuleSetOutput) AnomalyScore() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedRuleSet) *int { return v.AnomalyScore }).(pulumi.IntPtrOutput)
}

func (o ManagedRuleSetOutput) RuleGroupOverrides() ManagedRuleGroupOverrideArrayOutput {
	return o.ApplyT(func(v ManagedRuleSet) []ManagedRuleGroupOverride { return v.RuleGroupOverrides }).(ManagedRuleGroupOverrideArrayOutput)
}

func (o ManagedRuleSetOutput) RuleSetType() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleSet) string { return v.RuleSetType }).(pulumi.StringOutput)
}

func (o ManagedRuleSetOutput) RuleSetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleSet) string { return v.RuleSetVersion }).(pulumi.StringOutput)
}

type ManagedRuleSetArrayOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleSet)(nil)).Elem()
}

func (o ManagedRuleSetArrayOutput) ToManagedRuleSetArrayOutput() ManagedRuleSetArrayOutput {
	return o
}

func (o ManagedRuleSetArrayOutput) ToManagedRuleSetArrayOutputWithContext(ctx context.Context) ManagedRuleSetArrayOutput {
	return o
}

func (o ManagedRuleSetArrayOutput) Index(i pulumi.IntInput) ManagedRuleSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedRuleSet {
		return vs[0].([]ManagedRuleSet)[vs[1].(int)]
	}).(ManagedRuleSetOutput)
}

type ManagedRuleSetList struct {
	ManagedRuleSets []ManagedRuleSet `pulumi:"managedRuleSets"`
}





type ManagedRuleSetListInput interface {
	pulumi.Input

	ToManagedRuleSetListOutput() ManagedRuleSetListOutput
	ToManagedRuleSetListOutputWithContext(context.Context) ManagedRuleSetListOutput
}

type ManagedRuleSetListArgs struct {
	ManagedRuleSets ManagedRuleSetArrayInput `pulumi:"managedRuleSets"`
}

func (ManagedRuleSetListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetList)(nil)).Elem()
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListOutput() ManagedRuleSetListOutput {
	return i.ToManagedRuleSetListOutputWithContext(context.Background())
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListOutputWithContext(ctx context.Context) ManagedRuleSetListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListOutput)
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return i.ToManagedRuleSetListPtrOutputWithContext(context.Background())
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListOutput).ToManagedRuleSetListPtrOutputWithContext(ctx)
}









type ManagedRuleSetListPtrInput interface {
	pulumi.Input

	ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput
	ToManagedRuleSetListPtrOutputWithContext(context.Context) ManagedRuleSetListPtrOutput
}

type managedRuleSetListPtrType ManagedRuleSetListArgs

func ManagedRuleSetListPtr(v *ManagedRuleSetListArgs) ManagedRuleSetListPtrInput {
	return (*managedRuleSetListPtrType)(v)
}

func (*managedRuleSetListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRuleSetList)(nil)).Elem()
}

func (i *managedRuleSetListPtrType) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return i.ToManagedRuleSetListPtrOutputWithContext(context.Background())
}

func (i *managedRuleSetListPtrType) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListPtrOutput)
}

type ManagedRuleSetListOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetList)(nil)).Elem()
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListOutput() ManagedRuleSetListOutput {
	return o
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListOutputWithContext(ctx context.Context) ManagedRuleSetListOutput {
	return o
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return o.ToManagedRuleSetListPtrOutputWithContext(context.Background())
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedRuleSetList) *ManagedRuleSetList {
		return &v
	}).(ManagedRuleSetListPtrOutput)
}

func (o ManagedRuleSetListOutput) ManagedRuleSets() ManagedRuleSetArrayOutput {
	return o.ApplyT(func(v ManagedRuleSetList) []ManagedRuleSet { return v.ManagedRuleSets }).(ManagedRuleSetArrayOutput)
}

type ManagedRuleSetListPtrOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRuleSetList)(nil)).Elem()
}

func (o ManagedRuleSetListPtrOutput) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return o
}

func (o ManagedRuleSetListPtrOutput) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return o
}

func (o ManagedRuleSetListPtrOutput) Elem() ManagedRuleSetListOutput {
	return o.ApplyT(func(v *ManagedRuleSetList) ManagedRuleSetList {
		if v != nil {
			return *v
		}
		var ret ManagedRuleSetList
		return ret
	}).(ManagedRuleSetListOutput)
}

func (o ManagedRuleSetListPtrOutput) ManagedRuleSets() ManagedRuleSetArrayOutput {
	return o.ApplyT(func(v *ManagedRuleSetList) []ManagedRuleSet {
		if v == nil {
			return nil
		}
		return v.ManagedRuleSets
	}).(ManagedRuleSetArrayOutput)
}

type ManagedRuleSetListResponse struct {
	ManagedRuleSets []ManagedRuleSetResponse `pulumi:"managedRuleSets"`
}

type ManagedRuleSetListResponseOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetListResponse)(nil)).Elem()
}

func (o ManagedRuleSetListResponseOutput) ToManagedRuleSetListResponseOutput() ManagedRuleSetListResponseOutput {
	return o
}

func (o ManagedRuleSetListResponseOutput) ToManagedRuleSetListResponseOutputWithContext(ctx context.Context) ManagedRuleSetListResponseOutput {
	return o
}

func (o ManagedRuleSetListResponseOutput) ManagedRuleSets() ManagedRuleSetResponseArrayOutput {
	return o.ApplyT(func(v ManagedRuleSetListResponse) []ManagedRuleSetResponse { return v.ManagedRuleSets }).(ManagedRuleSetResponseArrayOutput)
}

type ManagedRuleSetListResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRuleSetListResponse)(nil)).Elem()
}

func (o ManagedRuleSetListResponsePtrOutput) ToManagedRuleSetListResponsePtrOutput() ManagedRuleSetListResponsePtrOutput {
	return o
}

func (o ManagedRuleSetListResponsePtrOutput) ToManagedRuleSetListResponsePtrOutputWithContext(ctx context.Context) ManagedRuleSetListResponsePtrOutput {
	return o
}

func (o ManagedRuleSetListResponsePtrOutput) Elem() ManagedRuleSetListResponseOutput {
	return o.ApplyT(func(v *ManagedRuleSetListResponse) ManagedRuleSetListResponse {
		if v != nil {
			return *v
		}
		var ret ManagedRuleSetListResponse
		return ret
	}).(ManagedRuleSetListResponseOutput)
}

func (o ManagedRuleSetListResponsePtrOutput) ManagedRuleSets() ManagedRuleSetResponseArrayOutput {
	return o.ApplyT(func(v *ManagedRuleSetListResponse) []ManagedRuleSetResponse {
		if v == nil {
			return nil
		}
		return v.ManagedRuleSets
	}).(ManagedRuleSetResponseArrayOutput)
}

type ManagedRuleSetResponse struct {
	AnomalyScore       *int                               `pulumi:"anomalyScore"`
	RuleGroupOverrides []ManagedRuleGroupOverrideResponse `pulumi:"ruleGroupOverrides"`
	RuleSetType        string                             `pulumi:"ruleSetType"`
	RuleSetVersion     string                             `pulumi:"ruleSetVersion"`
}

type ManagedRuleSetResponseOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetResponse)(nil)).Elem()
}

func (o ManagedRuleSetResponseOutput) ToManagedRuleSetResponseOutput() ManagedRuleSetResponseOutput {
	return o
}

func (o ManagedRuleSetResponseOutput) ToManagedRuleSetResponseOutputWithContext(ctx context.Context) ManagedRuleSetResponseOutput {
	return o
}

func (o ManagedRuleSetResponseOutput) AnomalyScore() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedRuleSetResponse) *int { return v.AnomalyScore }).(pulumi.IntPtrOutput)
}

func (o ManagedRuleSetResponseOutput) RuleGroupOverrides() ManagedRuleGroupOverrideResponseArrayOutput {
	return o.ApplyT(func(v ManagedRuleSetResponse) []ManagedRuleGroupOverrideResponse { return v.RuleGroupOverrides }).(ManagedRuleGroupOverrideResponseArrayOutput)
}

func (o ManagedRuleSetResponseOutput) RuleSetType() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleSetResponse) string { return v.RuleSetType }).(pulumi.StringOutput)
}

func (o ManagedRuleSetResponseOutput) RuleSetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleSetResponse) string { return v.RuleSetVersion }).(pulumi.StringOutput)
}

type ManagedRuleSetResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleSetResponse)(nil)).Elem()
}

func (o ManagedRuleSetResponseArrayOutput) ToManagedRuleSetResponseArrayOutput() ManagedRuleSetResponseArrayOutput {
	return o
}

func (o ManagedRuleSetResponseArrayOutput) ToManagedRuleSetResponseArrayOutputWithContext(ctx context.Context) ManagedRuleSetResponseArrayOutput {
	return o
}

func (o ManagedRuleSetResponseArrayOutput) Index(i pulumi.IntInput) ManagedRuleSetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedRuleSetResponse {
		return vs[0].([]ManagedRuleSetResponse)[vs[1].(int)]
	}).(ManagedRuleSetResponseOutput)
}

type MatchCondition struct {
	MatchValue      []string `pulumi:"matchValue"`
	MatchVariable   string   `pulumi:"matchVariable"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}





type MatchConditionInput interface {
	pulumi.Input

	ToMatchConditionOutput() MatchConditionOutput
	ToMatchConditionOutputWithContext(context.Context) MatchConditionOutput
}

type MatchConditionArgs struct {
	MatchValue      pulumi.StringArrayInput `pulumi:"matchValue"`
	MatchVariable   pulumi.StringInput      `pulumi:"matchVariable"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Selector        pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (MatchConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchCondition)(nil)).Elem()
}

func (i MatchConditionArgs) ToMatchConditionOutput() MatchConditionOutput {
	return i.ToMatchConditionOutputWithContext(context.Background())
}

func (i MatchConditionArgs) ToMatchConditionOutputWithContext(ctx context.Context) MatchConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MatchConditionOutput)
}





type MatchConditionArrayInput interface {
	pulumi.Input

	ToMatchConditionArrayOutput() MatchConditionArrayOutput
	ToMatchConditionArrayOutputWithContext(context.Context) MatchConditionArrayOutput
}

type MatchConditionArray []MatchConditionInput

func (MatchConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MatchCondition)(nil)).Elem()
}

func (i MatchConditionArray) ToMatchConditionArrayOutput() MatchConditionArrayOutput {
	return i.ToMatchConditionArrayOutputWithContext(context.Background())
}

func (i MatchConditionArray) ToMatchConditionArrayOutputWithContext(ctx context.Context) MatchConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MatchConditionArrayOutput)
}

type MatchConditionOutput struct{ *pulumi.OutputState }

func (MatchConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchCondition)(nil)).Elem()
}

func (o MatchConditionOutput) ToMatchConditionOutput() MatchConditionOutput {
	return o
}

func (o MatchConditionOutput) ToMatchConditionOutputWithContext(ctx context.Context) MatchConditionOutput {
	return o
}

func (o MatchConditionOutput) MatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MatchCondition) []string { return v.MatchValue }).(pulumi.StringArrayOutput)
}

func (o MatchConditionOutput) MatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v MatchCondition) string { return v.MatchVariable }).(pulumi.StringOutput)
}

func (o MatchConditionOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MatchCondition) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o MatchConditionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v MatchCondition) string { return v.Operator }).(pulumi.StringOutput)
}

func (o MatchConditionOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MatchCondition) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o MatchConditionOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MatchCondition) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type MatchConditionArrayOutput struct{ *pulumi.OutputState }

func (MatchConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MatchCondition)(nil)).Elem()
}

func (o MatchConditionArrayOutput) ToMatchConditionArrayOutput() MatchConditionArrayOutput {
	return o
}

func (o MatchConditionArrayOutput) ToMatchConditionArrayOutputWithContext(ctx context.Context) MatchConditionArrayOutput {
	return o
}

func (o MatchConditionArrayOutput) Index(i pulumi.IntInput) MatchConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MatchCondition {
		return vs[0].([]MatchCondition)[vs[1].(int)]
	}).(MatchConditionOutput)
}

type MatchConditionResponse struct {
	MatchValue      []string `pulumi:"matchValue"`
	MatchVariable   string   `pulumi:"matchVariable"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}

type MatchConditionResponseOutput struct{ *pulumi.OutputState }

func (MatchConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchConditionResponse)(nil)).Elem()
}

func (o MatchConditionResponseOutput) ToMatchConditionResponseOutput() MatchConditionResponseOutput {
	return o
}

func (o MatchConditionResponseOutput) ToMatchConditionResponseOutputWithContext(ctx context.Context) MatchConditionResponseOutput {
	return o
}

func (o MatchConditionResponseOutput) MatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MatchConditionResponse) []string { return v.MatchValue }).(pulumi.StringArrayOutput)
}

func (o MatchConditionResponseOutput) MatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v MatchConditionResponse) string { return v.MatchVariable }).(pulumi.StringOutput)
}

func (o MatchConditionResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MatchConditionResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o MatchConditionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v MatchConditionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o MatchConditionResponseOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MatchConditionResponse) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o MatchConditionResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MatchConditionResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type MatchConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (MatchConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MatchConditionResponse)(nil)).Elem()
}

func (o MatchConditionResponseArrayOutput) ToMatchConditionResponseArrayOutput() MatchConditionResponseArrayOutput {
	return o
}

func (o MatchConditionResponseArrayOutput) ToMatchConditionResponseArrayOutputWithContext(ctx context.Context) MatchConditionResponseArrayOutput {
	return o
}

func (o MatchConditionResponseArrayOutput) Index(i pulumi.IntInput) MatchConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MatchConditionResponse {
		return vs[0].([]MatchConditionResponse)[vs[1].(int)]
	}).(MatchConditionResponseOutput)
}

type PolicySettings struct {
	DefaultCustomBlockResponseBody       *string `pulumi:"defaultCustomBlockResponseBody"`
	DefaultCustomBlockResponseStatusCode *int    `pulumi:"defaultCustomBlockResponseStatusCode"`
	DefaultRedirectUrl                   *string `pulumi:"defaultRedirectUrl"`
	EnabledState                         *string `pulumi:"enabledState"`
	Mode                                 *string `pulumi:"mode"`
}





type PolicySettingsInput interface {
	pulumi.Input

	ToPolicySettingsOutput() PolicySettingsOutput
	ToPolicySettingsOutputWithContext(context.Context) PolicySettingsOutput
}

type PolicySettingsArgs struct {
	DefaultCustomBlockResponseBody       pulumi.StringPtrInput `pulumi:"defaultCustomBlockResponseBody"`
	DefaultCustomBlockResponseStatusCode pulumi.IntPtrInput    `pulumi:"defaultCustomBlockResponseStatusCode"`
	DefaultRedirectUrl                   pulumi.StringPtrInput `pulumi:"defaultRedirectUrl"`
	EnabledState                         pulumi.StringPtrInput `pulumi:"enabledState"`
	Mode                                 pulumi.StringPtrInput `pulumi:"mode"`
}

func (PolicySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySettings)(nil)).Elem()
}

func (i PolicySettingsArgs) ToPolicySettingsOutput() PolicySettingsOutput {
	return i.ToPolicySettingsOutputWithContext(context.Background())
}

func (i PolicySettingsArgs) ToPolicySettingsOutputWithContext(ctx context.Context) PolicySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySettingsOutput)
}

func (i PolicySettingsArgs) ToPolicySettingsPtrOutput() PolicySettingsPtrOutput {
	return i.ToPolicySettingsPtrOutputWithContext(context.Background())
}

func (i PolicySettingsArgs) ToPolicySettingsPtrOutputWithContext(ctx context.Context) PolicySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySettingsOutput).ToPolicySettingsPtrOutputWithContext(ctx)
}









type PolicySettingsPtrInput interface {
	pulumi.Input

	ToPolicySettingsPtrOutput() PolicySettingsPtrOutput
	ToPolicySettingsPtrOutputWithContext(context.Context) PolicySettingsPtrOutput
}

type policySettingsPtrType PolicySettingsArgs

func PolicySettingsPtr(v *PolicySettingsArgs) PolicySettingsPtrInput {
	return (*policySettingsPtrType)(v)
}

func (*policySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySettings)(nil)).Elem()
}

func (i *policySettingsPtrType) ToPolicySettingsPtrOutput() PolicySettingsPtrOutput {
	return i.ToPolicySettingsPtrOutputWithContext(context.Background())
}

func (i *policySettingsPtrType) ToPolicySettingsPtrOutputWithContext(ctx context.Context) PolicySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySettingsPtrOutput)
}

type PolicySettingsOutput struct{ *pulumi.OutputState }

func (PolicySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySettings)(nil)).Elem()
}

func (o PolicySettingsOutput) ToPolicySettingsOutput() PolicySettingsOutput {
	return o
}

func (o PolicySettingsOutput) ToPolicySettingsOutputWithContext(ctx context.Context) PolicySettingsOutput {
	return o
}

func (o PolicySettingsOutput) ToPolicySettingsPtrOutput() PolicySettingsPtrOutput {
	return o.ToPolicySettingsPtrOutputWithContext(context.Background())
}

func (o PolicySettingsOutput) ToPolicySettingsPtrOutputWithContext(ctx context.Context) PolicySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicySettings) *PolicySettings {
		return &v
	}).(PolicySettingsPtrOutput)
}

func (o PolicySettingsOutput) DefaultCustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySettings) *string { return v.DefaultCustomBlockResponseBody }).(pulumi.StringPtrOutput)
}

func (o PolicySettingsOutput) DefaultCustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PolicySettings) *int { return v.DefaultCustomBlockResponseStatusCode }).(pulumi.IntPtrOutput)
}

func (o PolicySettingsOutput) DefaultRedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySettings) *string { return v.DefaultRedirectUrl }).(pulumi.StringPtrOutput)
}

func (o PolicySettingsOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySettings) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o PolicySettingsOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySettings) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type PolicySettingsPtrOutput struct{ *pulumi.OutputState }

func (PolicySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySettings)(nil)).Elem()
}

func (o PolicySettingsPtrOutput) ToPolicySettingsPtrOutput() PolicySettingsPtrOutput {
	return o
}

func (o PolicySettingsPtrOutput) ToPolicySettingsPtrOutputWithContext(ctx context.Context) PolicySettingsPtrOutput {
	return o
}

func (o PolicySettingsPtrOutput) Elem() PolicySettingsOutput {
	return o.ApplyT(func(v *PolicySettings) PolicySettings {
		if v != nil {
			return *v
		}
		var ret PolicySettings
		return ret
	}).(PolicySettingsOutput)
}

func (o PolicySettingsPtrOutput) DefaultCustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.DefaultCustomBlockResponseBody
	}).(pulumi.StringPtrOutput)
}

func (o PolicySettingsPtrOutput) DefaultCustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicySettings) *int {
		if v == nil {
			return nil
		}
		return v.DefaultCustomBlockResponseStatusCode
	}).(pulumi.IntPtrOutput)
}

func (o PolicySettingsPtrOutput) DefaultRedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.DefaultRedirectUrl
	}).(pulumi.StringPtrOutput)
}

func (o PolicySettingsPtrOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.EnabledState
	}).(pulumi.StringPtrOutput)
}

func (o PolicySettingsPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type PolicySettingsResponse struct {
	DefaultCustomBlockResponseBody       *string `pulumi:"defaultCustomBlockResponseBody"`
	DefaultCustomBlockResponseStatusCode *int    `pulumi:"defaultCustomBlockResponseStatusCode"`
	DefaultRedirectUrl                   *string `pulumi:"defaultRedirectUrl"`
	EnabledState                         *string `pulumi:"enabledState"`
	Mode                                 *string `pulumi:"mode"`
}

type PolicySettingsResponseOutput struct{ *pulumi.OutputState }

func (PolicySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySettingsResponse)(nil)).Elem()
}

func (o PolicySettingsResponseOutput) ToPolicySettingsResponseOutput() PolicySettingsResponseOutput {
	return o
}

func (o PolicySettingsResponseOutput) ToPolicySettingsResponseOutputWithContext(ctx context.Context) PolicySettingsResponseOutput {
	return o
}

func (o PolicySettingsResponseOutput) DefaultCustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySettingsResponse) *string { return v.DefaultCustomBlockResponseBody }).(pulumi.StringPtrOutput)
}

func (o PolicySettingsResponseOutput) DefaultCustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PolicySettingsResponse) *int { return v.DefaultCustomBlockResponseStatusCode }).(pulumi.IntPtrOutput)
}

func (o PolicySettingsResponseOutput) DefaultRedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySettingsResponse) *string { return v.DefaultRedirectUrl }).(pulumi.StringPtrOutput)
}

func (o PolicySettingsResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySettingsResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o PolicySettingsResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySettingsResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type PolicySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (PolicySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySettingsResponse)(nil)).Elem()
}

func (o PolicySettingsResponsePtrOutput) ToPolicySettingsResponsePtrOutput() PolicySettingsResponsePtrOutput {
	return o
}

func (o PolicySettingsResponsePtrOutput) ToPolicySettingsResponsePtrOutputWithContext(ctx context.Context) PolicySettingsResponsePtrOutput {
	return o
}

func (o PolicySettingsResponsePtrOutput) Elem() PolicySettingsResponseOutput {
	return o.ApplyT(func(v *PolicySettingsResponse) PolicySettingsResponse {
		if v != nil {
			return *v
		}
		var ret PolicySettingsResponse
		return ret
	}).(PolicySettingsResponseOutput)
}

func (o PolicySettingsResponsePtrOutput) DefaultCustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultCustomBlockResponseBody
	}).(pulumi.StringPtrOutput)
}

func (o PolicySettingsResponsePtrOutput) DefaultCustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicySettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.DefaultCustomBlockResponseStatusCode
	}).(pulumi.IntPtrOutput)
}

func (o PolicySettingsResponsePtrOutput) DefaultRedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultRedirectUrl
	}).(pulumi.StringPtrOutput)
}

func (o PolicySettingsResponsePtrOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EnabledState
	}).(pulumi.StringPtrOutput)
}

func (o PolicySettingsResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type PostArgsMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}

type PostArgsMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}

type QueryStringMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type QueryStringMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type RateLimitRule struct {
	Action                     string           `pulumi:"action"`
	EnabledState               *string          `pulumi:"enabledState"`
	MatchConditions            []MatchCondition `pulumi:"matchConditions"`
	Name                       string           `pulumi:"name"`
	Priority                   int              `pulumi:"priority"`
	RateLimitDurationInMinutes int              `pulumi:"rateLimitDurationInMinutes"`
	RateLimitThreshold         int              `pulumi:"rateLimitThreshold"`
}





type RateLimitRuleInput interface {
	pulumi.Input

	ToRateLimitRuleOutput() RateLimitRuleOutput
	ToRateLimitRuleOutputWithContext(context.Context) RateLimitRuleOutput
}

type RateLimitRuleArgs struct {
	Action                     pulumi.StringInput       `pulumi:"action"`
	EnabledState               pulumi.StringPtrInput    `pulumi:"enabledState"`
	MatchConditions            MatchConditionArrayInput `pulumi:"matchConditions"`
	Name                       pulumi.StringInput       `pulumi:"name"`
	Priority                   pulumi.IntInput          `pulumi:"priority"`
	RateLimitDurationInMinutes pulumi.IntInput          `pulumi:"rateLimitDurationInMinutes"`
	RateLimitThreshold         pulumi.IntInput          `pulumi:"rateLimitThreshold"`
}

func (RateLimitRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RateLimitRule)(nil)).Elem()
}

func (i RateLimitRuleArgs) ToRateLimitRuleOutput() RateLimitRuleOutput {
	return i.ToRateLimitRuleOutputWithContext(context.Background())
}

func (i RateLimitRuleArgs) ToRateLimitRuleOutputWithContext(ctx context.Context) RateLimitRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateLimitRuleOutput)
}





type RateLimitRuleArrayInput interface {
	pulumi.Input

	ToRateLimitRuleArrayOutput() RateLimitRuleArrayOutput
	ToRateLimitRuleArrayOutputWithContext(context.Context) RateLimitRuleArrayOutput
}

type RateLimitRuleArray []RateLimitRuleInput

func (RateLimitRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RateLimitRule)(nil)).Elem()
}

func (i RateLimitRuleArray) ToRateLimitRuleArrayOutput() RateLimitRuleArrayOutput {
	return i.ToRateLimitRuleArrayOutputWithContext(context.Background())
}

func (i RateLimitRuleArray) ToRateLimitRuleArrayOutputWithContext(ctx context.Context) RateLimitRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateLimitRuleArrayOutput)
}

type RateLimitRuleOutput struct{ *pulumi.OutputState }

func (RateLimitRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RateLimitRule)(nil)).Elem()
}

func (o RateLimitRuleOutput) ToRateLimitRuleOutput() RateLimitRuleOutput {
	return o
}

func (o RateLimitRuleOutput) ToRateLimitRuleOutputWithContext(ctx context.Context) RateLimitRuleOutput {
	return o
}

func (o RateLimitRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v RateLimitRule) string { return v.Action }).(pulumi.StringOutput)
}

func (o RateLimitRuleOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RateLimitRule) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o RateLimitRuleOutput) MatchConditions() MatchConditionArrayOutput {
	return o.ApplyT(func(v RateLimitRule) []MatchCondition { return v.MatchConditions }).(MatchConditionArrayOutput)
}

func (o RateLimitRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RateLimitRule) string { return v.Name }).(pulumi.StringOutput)
}

func (o RateLimitRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v RateLimitRule) int { return v.Priority }).(pulumi.IntOutput)
}

func (o RateLimitRuleOutput) RateLimitDurationInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v RateLimitRule) int { return v.RateLimitDurationInMinutes }).(pulumi.IntOutput)
}

func (o RateLimitRuleOutput) RateLimitThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v RateLimitRule) int { return v.RateLimitThreshold }).(pulumi.IntOutput)
}

type RateLimitRuleArrayOutput struct{ *pulumi.OutputState }

func (RateLimitRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RateLimitRule)(nil)).Elem()
}

func (o RateLimitRuleArrayOutput) ToRateLimitRuleArrayOutput() RateLimitRuleArrayOutput {
	return o
}

func (o RateLimitRuleArrayOutput) ToRateLimitRuleArrayOutputWithContext(ctx context.Context) RateLimitRuleArrayOutput {
	return o
}

func (o RateLimitRuleArrayOutput) Index(i pulumi.IntInput) RateLimitRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RateLimitRule {
		return vs[0].([]RateLimitRule)[vs[1].(int)]
	}).(RateLimitRuleOutput)
}

type RateLimitRuleList struct {
	Rules []RateLimitRule `pulumi:"rules"`
}





type RateLimitRuleListInput interface {
	pulumi.Input

	ToRateLimitRuleListOutput() RateLimitRuleListOutput
	ToRateLimitRuleListOutputWithContext(context.Context) RateLimitRuleListOutput
}

type RateLimitRuleListArgs struct {
	Rules RateLimitRuleArrayInput `pulumi:"rules"`
}

func (RateLimitRuleListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RateLimitRuleList)(nil)).Elem()
}

func (i RateLimitRuleListArgs) ToRateLimitRuleListOutput() RateLimitRuleListOutput {
	return i.ToRateLimitRuleListOutputWithContext(context.Background())
}

func (i RateLimitRuleListArgs) ToRateLimitRuleListOutputWithContext(ctx context.Context) RateLimitRuleListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateLimitRuleListOutput)
}

func (i RateLimitRuleListArgs) ToRateLimitRuleListPtrOutput() RateLimitRuleListPtrOutput {
	return i.ToRateLimitRuleListPtrOutputWithContext(context.Background())
}

func (i RateLimitRuleListArgs) ToRateLimitRuleListPtrOutputWithContext(ctx context.Context) RateLimitRuleListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateLimitRuleListOutput).ToRateLimitRuleListPtrOutputWithContext(ctx)
}









type RateLimitRuleListPtrInput interface {
	pulumi.Input

	ToRateLimitRuleListPtrOutput() RateLimitRuleListPtrOutput
	ToRateLimitRuleListPtrOutputWithContext(context.Context) RateLimitRuleListPtrOutput
}

type rateLimitRuleListPtrType RateLimitRuleListArgs

func RateLimitRuleListPtr(v *RateLimitRuleListArgs) RateLimitRuleListPtrInput {
	return (*rateLimitRuleListPtrType)(v)
}

func (*rateLimitRuleListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RateLimitRuleList)(nil)).Elem()
}

func (i *rateLimitRuleListPtrType) ToRateLimitRuleListPtrOutput() RateLimitRuleListPtrOutput {
	return i.ToRateLimitRuleListPtrOutputWithContext(context.Background())
}

func (i *rateLimitRuleListPtrType) ToRateLimitRuleListPtrOutputWithContext(ctx context.Context) RateLimitRuleListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateLimitRuleListPtrOutput)
}

type RateLimitRuleListOutput struct{ *pulumi.OutputState }

func (RateLimitRuleListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RateLimitRuleList)(nil)).Elem()
}

func (o RateLimitRuleListOutput) ToRateLimitRuleListOutput() RateLimitRuleListOutput {
	return o
}

func (o RateLimitRuleListOutput) ToRateLimitRuleListOutputWithContext(ctx context.Context) RateLimitRuleListOutput {
	return o
}

func (o RateLimitRuleListOutput) ToRateLimitRuleListPtrOutput() RateLimitRuleListPtrOutput {
	return o.ToRateLimitRuleListPtrOutputWithContext(context.Background())
}

func (o RateLimitRuleListOutput) ToRateLimitRuleListPtrOutputWithContext(ctx context.Context) RateLimitRuleListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RateLimitRuleList) *RateLimitRuleList {
		return &v
	}).(RateLimitRuleListPtrOutput)
}

func (o RateLimitRuleListOutput) Rules() RateLimitRuleArrayOutput {
	return o.ApplyT(func(v RateLimitRuleList) []RateLimitRule { return v.Rules }).(RateLimitRuleArrayOutput)
}

type RateLimitRuleListPtrOutput struct{ *pulumi.OutputState }

func (RateLimitRuleListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RateLimitRuleList)(nil)).Elem()
}

func (o RateLimitRuleListPtrOutput) ToRateLimitRuleListPtrOutput() RateLimitRuleListPtrOutput {
	return o
}

func (o RateLimitRuleListPtrOutput) ToRateLimitRuleListPtrOutputWithContext(ctx context.Context) RateLimitRuleListPtrOutput {
	return o
}

func (o RateLimitRuleListPtrOutput) Elem() RateLimitRuleListOutput {
	return o.ApplyT(func(v *RateLimitRuleList) RateLimitRuleList {
		if v != nil {
			return *v
		}
		var ret RateLimitRuleList
		return ret
	}).(RateLimitRuleListOutput)
}

func (o RateLimitRuleListPtrOutput) Rules() RateLimitRuleArrayOutput {
	return o.ApplyT(func(v *RateLimitRuleList) []RateLimitRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(RateLimitRuleArrayOutput)
}

type RateLimitRuleListResponse struct {
	Rules []RateLimitRuleResponse `pulumi:"rules"`
}

type RateLimitRuleListResponseOutput struct{ *pulumi.OutputState }

func (RateLimitRuleListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RateLimitRuleListResponse)(nil)).Elem()
}

func (o RateLimitRuleListResponseOutput) ToRateLimitRuleListResponseOutput() RateLimitRuleListResponseOutput {
	return o
}

func (o RateLimitRuleListResponseOutput) ToRateLimitRuleListResponseOutputWithContext(ctx context.Context) RateLimitRuleListResponseOutput {
	return o
}

func (o RateLimitRuleListResponseOutput) Rules() RateLimitRuleResponseArrayOutput {
	return o.ApplyT(func(v RateLimitRuleListResponse) []RateLimitRuleResponse { return v.Rules }).(RateLimitRuleResponseArrayOutput)
}

type RateLimitRuleListResponsePtrOutput struct{ *pulumi.OutputState }

func (RateLimitRuleListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RateLimitRuleListResponse)(nil)).Elem()
}

func (o RateLimitRuleListResponsePtrOutput) ToRateLimitRuleListResponsePtrOutput() RateLimitRuleListResponsePtrOutput {
	return o
}

func (o RateLimitRuleListResponsePtrOutput) ToRateLimitRuleListResponsePtrOutputWithContext(ctx context.Context) RateLimitRuleListResponsePtrOutput {
	return o
}

func (o RateLimitRuleListResponsePtrOutput) Elem() RateLimitRuleListResponseOutput {
	return o.ApplyT(func(v *RateLimitRuleListResponse) RateLimitRuleListResponse {
		if v != nil {
			return *v
		}
		var ret RateLimitRuleListResponse
		return ret
	}).(RateLimitRuleListResponseOutput)
}

func (o RateLimitRuleListResponsePtrOutput) Rules() RateLimitRuleResponseArrayOutput {
	return o.ApplyT(func(v *RateLimitRuleListResponse) []RateLimitRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(RateLimitRuleResponseArrayOutput)
}

type RateLimitRuleResponse struct {
	Action                     string                   `pulumi:"action"`
	EnabledState               *string                  `pulumi:"enabledState"`
	MatchConditions            []MatchConditionResponse `pulumi:"matchConditions"`
	Name                       string                   `pulumi:"name"`
	Priority                   int                      `pulumi:"priority"`
	RateLimitDurationInMinutes int                      `pulumi:"rateLimitDurationInMinutes"`
	RateLimitThreshold         int                      `pulumi:"rateLimitThreshold"`
}

type RateLimitRuleResponseOutput struct{ *pulumi.OutputState }

func (RateLimitRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RateLimitRuleResponse)(nil)).Elem()
}

func (o RateLimitRuleResponseOutput) ToRateLimitRuleResponseOutput() RateLimitRuleResponseOutput {
	return o
}

func (o RateLimitRuleResponseOutput) ToRateLimitRuleResponseOutputWithContext(ctx context.Context) RateLimitRuleResponseOutput {
	return o
}

func (o RateLimitRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v RateLimitRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o RateLimitRuleResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RateLimitRuleResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o RateLimitRuleResponseOutput) MatchConditions() MatchConditionResponseArrayOutput {
	return o.ApplyT(func(v RateLimitRuleResponse) []MatchConditionResponse { return v.MatchConditions }).(MatchConditionResponseArrayOutput)
}

func (o RateLimitRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RateLimitRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RateLimitRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v RateLimitRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

func (o RateLimitRuleResponseOutput) RateLimitDurationInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v RateLimitRuleResponse) int { return v.RateLimitDurationInMinutes }).(pulumi.IntOutput)
}

func (o RateLimitRuleResponseOutput) RateLimitThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v RateLimitRuleResponse) int { return v.RateLimitThreshold }).(pulumi.IntOutput)
}

type RateLimitRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RateLimitRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RateLimitRuleResponse)(nil)).Elem()
}

func (o RateLimitRuleResponseArrayOutput) ToRateLimitRuleResponseArrayOutput() RateLimitRuleResponseArrayOutput {
	return o
}

func (o RateLimitRuleResponseArrayOutput) ToRateLimitRuleResponseArrayOutputWithContext(ctx context.Context) RateLimitRuleResponseArrayOutput {
	return o
}

func (o RateLimitRuleResponseArrayOutput) Index(i pulumi.IntInput) RateLimitRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RateLimitRuleResponse {
		return vs[0].([]RateLimitRuleResponse)[vs[1].(int)]
	}).(RateLimitRuleResponseOutput)
}

type RemoteAddressMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type RemoteAddressMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type RequestBodyMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type RequestBodyMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type RequestHeaderMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}

type RequestHeaderMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}

type RequestMethodMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
}

type RequestMethodMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
}

type RequestSchemeMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
}

type RequestSchemeMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
}

type RequestUriMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type RequestUriMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
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

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
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

type UrlFileExtensionMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type UrlFileExtensionMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type UrlFileNameMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type UrlFileNameMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type UrlPathMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type UrlPathMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}

type UrlRedirectAction struct {
	Name       string                      `pulumi:"name"`
	Parameters UrlRedirectActionParameters `pulumi:"parameters"`
}

type UrlRedirectActionParameters struct {
	CustomFragment      *string `pulumi:"customFragment"`
	CustomHostname      *string `pulumi:"customHostname"`
	CustomPath          *string `pulumi:"customPath"`
	CustomQueryString   *string `pulumi:"customQueryString"`
	DestinationProtocol *string `pulumi:"destinationProtocol"`
	OdataType           string  `pulumi:"odataType"`
	RedirectType        string  `pulumi:"redirectType"`
}

type UrlRedirectActionParametersResponse struct {
	CustomFragment      *string `pulumi:"customFragment"`
	CustomHostname      *string `pulumi:"customHostname"`
	CustomPath          *string `pulumi:"customPath"`
	CustomQueryString   *string `pulumi:"customQueryString"`
	DestinationProtocol *string `pulumi:"destinationProtocol"`
	OdataType           string  `pulumi:"odataType"`
	RedirectType        string  `pulumi:"redirectType"`
}

type UrlRedirectActionResponse struct {
	Name       string                              `pulumi:"name"`
	Parameters UrlRedirectActionParametersResponse `pulumi:"parameters"`
}

type UrlRewriteAction struct {
	Name       string                     `pulumi:"name"`
	Parameters UrlRewriteActionParameters `pulumi:"parameters"`
}

type UrlRewriteActionParameters struct {
	Destination           string `pulumi:"destination"`
	OdataType             string `pulumi:"odataType"`
	PreserveUnmatchedPath *bool  `pulumi:"preserveUnmatchedPath"`
	SourcePattern         string `pulumi:"sourcePattern"`
}

type UrlRewriteActionParametersResponse struct {
	Destination           string `pulumi:"destination"`
	OdataType             string `pulumi:"odataType"`
	PreserveUnmatchedPath *bool  `pulumi:"preserveUnmatchedPath"`
	SourcePattern         string `pulumi:"sourcePattern"`
}

type UrlRewriteActionResponse struct {
	Name       string                             `pulumi:"name"`
	Parameters UrlRewriteActionParametersResponse `pulumi:"parameters"`
}

type UserManagedHttpsParametersResponse struct {
	CertificateSource           string                                      `pulumi:"certificateSource"`
	CertificateSourceParameters KeyVaultCertificateSourceParametersResponse `pulumi:"certificateSourceParameters"`
	MinimumTlsVersion           *string                                     `pulumi:"minimumTlsVersion"`
	ProtocolType                string                                      `pulumi:"protocolType"`
}

func init() {
	pulumi.RegisterOutputType(CdnEndpointResponseOutput{})
	pulumi.RegisterOutputType(CdnEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(CustomRuleOutput{})
	pulumi.RegisterOutputType(CustomRuleArrayOutput{})
	pulumi.RegisterOutputType(CustomRuleListOutput{})
	pulumi.RegisterOutputType(CustomRuleListPtrOutput{})
	pulumi.RegisterOutputType(CustomRuleListResponseOutput{})
	pulumi.RegisterOutputType(CustomRuleListResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomRuleResponseOutput{})
	pulumi.RegisterOutputType(CustomRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginArrayOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginResponseOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginResponseArrayOutput{})
	pulumi.RegisterOutputType(DeliveryRuleOutput{})
	pulumi.RegisterOutputType(DeliveryRuleArrayOutput{})
	pulumi.RegisterOutputType(DeliveryRuleResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesUpdateParametersDeliveryPolicyOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesUpdateParametersDeliveryPolicyPtrOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput{})
	pulumi.RegisterOutputType(GeoFilterOutput{})
	pulumi.RegisterOutputType(GeoFilterArrayOutput{})
	pulumi.RegisterOutputType(GeoFilterResponseOutput{})
	pulumi.RegisterOutputType(GeoFilterResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleGroupOverrideOutput{})
	pulumi.RegisterOutputType(ManagedRuleGroupOverrideArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleGroupOverrideResponseOutput{})
	pulumi.RegisterOutputType(ManagedRuleGroupOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleOverrideOutput{})
	pulumi.RegisterOutputType(ManagedRuleOverrideArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleOverrideResponseOutput{})
	pulumi.RegisterOutputType(ManagedRuleOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListPtrOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListResponseOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetResponseOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetResponseArrayOutput{})
	pulumi.RegisterOutputType(MatchConditionOutput{})
	pulumi.RegisterOutputType(MatchConditionArrayOutput{})
	pulumi.RegisterOutputType(MatchConditionResponseOutput{})
	pulumi.RegisterOutputType(MatchConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(PolicySettingsOutput{})
	pulumi.RegisterOutputType(PolicySettingsPtrOutput{})
	pulumi.RegisterOutputType(PolicySettingsResponseOutput{})
	pulumi.RegisterOutputType(PolicySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(RateLimitRuleOutput{})
	pulumi.RegisterOutputType(RateLimitRuleArrayOutput{})
	pulumi.RegisterOutputType(RateLimitRuleListOutput{})
	pulumi.RegisterOutputType(RateLimitRuleListPtrOutput{})
	pulumi.RegisterOutputType(RateLimitRuleListResponseOutput{})
	pulumi.RegisterOutputType(RateLimitRuleListResponsePtrOutput{})
	pulumi.RegisterOutputType(RateLimitRuleResponseOutput{})
	pulumi.RegisterOutputType(RateLimitRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
}
