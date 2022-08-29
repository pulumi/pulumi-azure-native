


package v20191231

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

type DeepCreatedOrigin struct {
	Enabled          *bool   `pulumi:"enabled"`
	HostName         string  `pulumi:"hostName"`
	HttpPort         *int    `pulumi:"httpPort"`
	HttpsPort        *int    `pulumi:"httpsPort"`
	Name             string  `pulumi:"name"`
	OriginHostHeader *string `pulumi:"originHostHeader"`
	Priority         *int    `pulumi:"priority"`
	Weight           *int    `pulumi:"weight"`
}





type DeepCreatedOriginInput interface {
	pulumi.Input

	ToDeepCreatedOriginOutput() DeepCreatedOriginOutput
	ToDeepCreatedOriginOutputWithContext(context.Context) DeepCreatedOriginOutput
}

type DeepCreatedOriginArgs struct {
	Enabled          pulumi.BoolPtrInput   `pulumi:"enabled"`
	HostName         pulumi.StringInput    `pulumi:"hostName"`
	HttpPort         pulumi.IntPtrInput    `pulumi:"httpPort"`
	HttpsPort        pulumi.IntPtrInput    `pulumi:"httpsPort"`
	Name             pulumi.StringInput    `pulumi:"name"`
	OriginHostHeader pulumi.StringPtrInput `pulumi:"originHostHeader"`
	Priority         pulumi.IntPtrInput    `pulumi:"priority"`
	Weight           pulumi.IntPtrInput    `pulumi:"weight"`
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

func (o DeepCreatedOriginOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
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

func (o DeepCreatedOriginOutput) OriginHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *string { return v.OriginHostHeader }).(pulumi.StringPtrOutput)
}

func (o DeepCreatedOriginOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o DeepCreatedOriginOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *int { return v.Weight }).(pulumi.IntPtrOutput)
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

type DeepCreatedOriginGroup struct {
	HealthProbeSettings                                   *HealthProbeParameters                       `pulumi:"healthProbeSettings"`
	Name                                                  string                                       `pulumi:"name"`
	Origins                                               []ResourceReference                          `pulumi:"origins"`
	ResponseBasedOriginErrorDetectionSettings             *ResponseBasedOriginErrorDetectionParameters `pulumi:"responseBasedOriginErrorDetectionSettings"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                                         `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
}





type DeepCreatedOriginGroupInput interface {
	pulumi.Input

	ToDeepCreatedOriginGroupOutput() DeepCreatedOriginGroupOutput
	ToDeepCreatedOriginGroupOutputWithContext(context.Context) DeepCreatedOriginGroupOutput
}

type DeepCreatedOriginGroupArgs struct {
	HealthProbeSettings                                   HealthProbeParametersPtrInput                       `pulumi:"healthProbeSettings"`
	Name                                                  pulumi.StringInput                                  `pulumi:"name"`
	Origins                                               ResourceReferenceArrayInput                         `pulumi:"origins"`
	ResponseBasedOriginErrorDetectionSettings             ResponseBasedOriginErrorDetectionParametersPtrInput `pulumi:"responseBasedOriginErrorDetectionSettings"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes pulumi.IntPtrInput                                  `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
}

func (DeepCreatedOriginGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOriginGroup)(nil)).Elem()
}

func (i DeepCreatedOriginGroupArgs) ToDeepCreatedOriginGroupOutput() DeepCreatedOriginGroupOutput {
	return i.ToDeepCreatedOriginGroupOutputWithContext(context.Background())
}

func (i DeepCreatedOriginGroupArgs) ToDeepCreatedOriginGroupOutputWithContext(ctx context.Context) DeepCreatedOriginGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginGroupOutput)
}





type DeepCreatedOriginGroupArrayInput interface {
	pulumi.Input

	ToDeepCreatedOriginGroupArrayOutput() DeepCreatedOriginGroupArrayOutput
	ToDeepCreatedOriginGroupArrayOutputWithContext(context.Context) DeepCreatedOriginGroupArrayOutput
}

type DeepCreatedOriginGroupArray []DeepCreatedOriginGroupInput

func (DeepCreatedOriginGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOriginGroup)(nil)).Elem()
}

func (i DeepCreatedOriginGroupArray) ToDeepCreatedOriginGroupArrayOutput() DeepCreatedOriginGroupArrayOutput {
	return i.ToDeepCreatedOriginGroupArrayOutputWithContext(context.Background())
}

func (i DeepCreatedOriginGroupArray) ToDeepCreatedOriginGroupArrayOutputWithContext(ctx context.Context) DeepCreatedOriginGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginGroupArrayOutput)
}

type DeepCreatedOriginGroupOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOriginGroup)(nil)).Elem()
}

func (o DeepCreatedOriginGroupOutput) ToDeepCreatedOriginGroupOutput() DeepCreatedOriginGroupOutput {
	return o
}

func (o DeepCreatedOriginGroupOutput) ToDeepCreatedOriginGroupOutputWithContext(ctx context.Context) DeepCreatedOriginGroupOutput {
	return o
}

func (o DeepCreatedOriginGroupOutput) HealthProbeSettings() HealthProbeParametersPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginGroup) *HealthProbeParameters { return v.HealthProbeSettings }).(HealthProbeParametersPtrOutput)
}

func (o DeepCreatedOriginGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeepCreatedOriginGroup) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeepCreatedOriginGroupOutput) Origins() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v DeepCreatedOriginGroup) []ResourceReference { return v.Origins }).(ResourceReferenceArrayOutput)
}

func (o DeepCreatedOriginGroupOutput) ResponseBasedOriginErrorDetectionSettings() ResponseBasedOriginErrorDetectionParametersPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginGroup) *ResponseBasedOriginErrorDetectionParameters {
		return v.ResponseBasedOriginErrorDetectionSettings
	}).(ResponseBasedOriginErrorDetectionParametersPtrOutput)
}

func (o DeepCreatedOriginGroupOutput) TrafficRestorationTimeToHealedOrNewEndpointsInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginGroup) *int { return v.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes }).(pulumi.IntPtrOutput)
}

type DeepCreatedOriginGroupArrayOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOriginGroup)(nil)).Elem()
}

func (o DeepCreatedOriginGroupArrayOutput) ToDeepCreatedOriginGroupArrayOutput() DeepCreatedOriginGroupArrayOutput {
	return o
}

func (o DeepCreatedOriginGroupArrayOutput) ToDeepCreatedOriginGroupArrayOutputWithContext(ctx context.Context) DeepCreatedOriginGroupArrayOutput {
	return o
}

func (o DeepCreatedOriginGroupArrayOutput) Index(i pulumi.IntInput) DeepCreatedOriginGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeepCreatedOriginGroup {
		return vs[0].([]DeepCreatedOriginGroup)[vs[1].(int)]
	}).(DeepCreatedOriginGroupOutput)
}

type DeepCreatedOriginGroupResponse struct {
	HealthProbeSettings                                   *HealthProbeParametersResponse                       `pulumi:"healthProbeSettings"`
	Name                                                  string                                               `pulumi:"name"`
	Origins                                               []ResourceReferenceResponse                          `pulumi:"origins"`
	ResponseBasedOriginErrorDetectionSettings             *ResponseBasedOriginErrorDetectionParametersResponse `pulumi:"responseBasedOriginErrorDetectionSettings"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                                                 `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
}

type DeepCreatedOriginGroupResponseOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOriginGroupResponse)(nil)).Elem()
}

func (o DeepCreatedOriginGroupResponseOutput) ToDeepCreatedOriginGroupResponseOutput() DeepCreatedOriginGroupResponseOutput {
	return o
}

func (o DeepCreatedOriginGroupResponseOutput) ToDeepCreatedOriginGroupResponseOutputWithContext(ctx context.Context) DeepCreatedOriginGroupResponseOutput {
	return o
}

func (o DeepCreatedOriginGroupResponseOutput) HealthProbeSettings() HealthProbeParametersResponsePtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginGroupResponse) *HealthProbeParametersResponse { return v.HealthProbeSettings }).(HealthProbeParametersResponsePtrOutput)
}

func (o DeepCreatedOriginGroupResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeepCreatedOriginGroupResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeepCreatedOriginGroupResponseOutput) Origins() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v DeepCreatedOriginGroupResponse) []ResourceReferenceResponse { return v.Origins }).(ResourceReferenceResponseArrayOutput)
}

func (o DeepCreatedOriginGroupResponseOutput) ResponseBasedOriginErrorDetectionSettings() ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginGroupResponse) *ResponseBasedOriginErrorDetectionParametersResponse {
		return v.ResponseBasedOriginErrorDetectionSettings
	}).(ResponseBasedOriginErrorDetectionParametersResponsePtrOutput)
}

func (o DeepCreatedOriginGroupResponseOutput) TrafficRestorationTimeToHealedOrNewEndpointsInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginGroupResponse) *int {
		return v.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes
	}).(pulumi.IntPtrOutput)
}

type DeepCreatedOriginGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOriginGroupResponse)(nil)).Elem()
}

func (o DeepCreatedOriginGroupResponseArrayOutput) ToDeepCreatedOriginGroupResponseArrayOutput() DeepCreatedOriginGroupResponseArrayOutput {
	return o
}

func (o DeepCreatedOriginGroupResponseArrayOutput) ToDeepCreatedOriginGroupResponseArrayOutputWithContext(ctx context.Context) DeepCreatedOriginGroupResponseArrayOutput {
	return o
}

func (o DeepCreatedOriginGroupResponseArrayOutput) Index(i pulumi.IntInput) DeepCreatedOriginGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeepCreatedOriginGroupResponse {
		return vs[0].([]DeepCreatedOriginGroupResponse)[vs[1].(int)]
	}).(DeepCreatedOriginGroupResponseOutput)
}

type DeepCreatedOriginResponse struct {
	Enabled          *bool   `pulumi:"enabled"`
	HostName         string  `pulumi:"hostName"`
	HttpPort         *int    `pulumi:"httpPort"`
	HttpsPort        *int    `pulumi:"httpsPort"`
	Name             string  `pulumi:"name"`
	OriginHostHeader *string `pulumi:"originHostHeader"`
	Priority         *int    `pulumi:"priority"`
	Weight           *int    `pulumi:"weight"`
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

func (o DeepCreatedOriginResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
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

func (o DeepCreatedOriginResponseOutput) OriginHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *string { return v.OriginHostHeader }).(pulumi.StringPtrOutput)
}

func (o DeepCreatedOriginResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o DeepCreatedOriginResponseOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *int { return v.Weight }).(pulumi.IntPtrOutput)
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

type HealthProbeParameters struct {
	ProbeIntervalInSeconds *int                    `pulumi:"probeIntervalInSeconds"`
	ProbePath              *string                 `pulumi:"probePath"`
	ProbeProtocol          *ProbeProtocol          `pulumi:"probeProtocol"`
	ProbeRequestType       *HealthProbeRequestType `pulumi:"probeRequestType"`
}





type HealthProbeParametersInput interface {
	pulumi.Input

	ToHealthProbeParametersOutput() HealthProbeParametersOutput
	ToHealthProbeParametersOutputWithContext(context.Context) HealthProbeParametersOutput
}

type HealthProbeParametersArgs struct {
	ProbeIntervalInSeconds pulumi.IntPtrInput             `pulumi:"probeIntervalInSeconds"`
	ProbePath              pulumi.StringPtrInput          `pulumi:"probePath"`
	ProbeProtocol          ProbeProtocolPtrInput          `pulumi:"probeProtocol"`
	ProbeRequestType       HealthProbeRequestTypePtrInput `pulumi:"probeRequestType"`
}

func (HealthProbeParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeParameters)(nil)).Elem()
}

func (i HealthProbeParametersArgs) ToHealthProbeParametersOutput() HealthProbeParametersOutput {
	return i.ToHealthProbeParametersOutputWithContext(context.Background())
}

func (i HealthProbeParametersArgs) ToHealthProbeParametersOutputWithContext(ctx context.Context) HealthProbeParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthProbeParametersOutput)
}

func (i HealthProbeParametersArgs) ToHealthProbeParametersPtrOutput() HealthProbeParametersPtrOutput {
	return i.ToHealthProbeParametersPtrOutputWithContext(context.Background())
}

func (i HealthProbeParametersArgs) ToHealthProbeParametersPtrOutputWithContext(ctx context.Context) HealthProbeParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthProbeParametersOutput).ToHealthProbeParametersPtrOutputWithContext(ctx)
}









type HealthProbeParametersPtrInput interface {
	pulumi.Input

	ToHealthProbeParametersPtrOutput() HealthProbeParametersPtrOutput
	ToHealthProbeParametersPtrOutputWithContext(context.Context) HealthProbeParametersPtrOutput
}

type healthProbeParametersPtrType HealthProbeParametersArgs

func HealthProbeParametersPtr(v *HealthProbeParametersArgs) HealthProbeParametersPtrInput {
	return (*healthProbeParametersPtrType)(v)
}

func (*healthProbeParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthProbeParameters)(nil)).Elem()
}

func (i *healthProbeParametersPtrType) ToHealthProbeParametersPtrOutput() HealthProbeParametersPtrOutput {
	return i.ToHealthProbeParametersPtrOutputWithContext(context.Background())
}

func (i *healthProbeParametersPtrType) ToHealthProbeParametersPtrOutputWithContext(ctx context.Context) HealthProbeParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthProbeParametersPtrOutput)
}

type HealthProbeParametersOutput struct{ *pulumi.OutputState }

func (HealthProbeParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeParameters)(nil)).Elem()
}

func (o HealthProbeParametersOutput) ToHealthProbeParametersOutput() HealthProbeParametersOutput {
	return o
}

func (o HealthProbeParametersOutput) ToHealthProbeParametersOutputWithContext(ctx context.Context) HealthProbeParametersOutput {
	return o
}

func (o HealthProbeParametersOutput) ToHealthProbeParametersPtrOutput() HealthProbeParametersPtrOutput {
	return o.ToHealthProbeParametersPtrOutputWithContext(context.Background())
}

func (o HealthProbeParametersOutput) ToHealthProbeParametersPtrOutputWithContext(ctx context.Context) HealthProbeParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HealthProbeParameters) *HealthProbeParameters {
		return &v
	}).(HealthProbeParametersPtrOutput)
}

func (o HealthProbeParametersOutput) ProbeIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HealthProbeParameters) *int { return v.ProbeIntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o HealthProbeParametersOutput) ProbePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeParameters) *string { return v.ProbePath }).(pulumi.StringPtrOutput)
}

func (o HealthProbeParametersOutput) ProbeProtocol() ProbeProtocolPtrOutput {
	return o.ApplyT(func(v HealthProbeParameters) *ProbeProtocol { return v.ProbeProtocol }).(ProbeProtocolPtrOutput)
}

func (o HealthProbeParametersOutput) ProbeRequestType() HealthProbeRequestTypePtrOutput {
	return o.ApplyT(func(v HealthProbeParameters) *HealthProbeRequestType { return v.ProbeRequestType }).(HealthProbeRequestTypePtrOutput)
}

type HealthProbeParametersPtrOutput struct{ *pulumi.OutputState }

func (HealthProbeParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthProbeParameters)(nil)).Elem()
}

func (o HealthProbeParametersPtrOutput) ToHealthProbeParametersPtrOutput() HealthProbeParametersPtrOutput {
	return o
}

func (o HealthProbeParametersPtrOutput) ToHealthProbeParametersPtrOutputWithContext(ctx context.Context) HealthProbeParametersPtrOutput {
	return o
}

func (o HealthProbeParametersPtrOutput) Elem() HealthProbeParametersOutput {
	return o.ApplyT(func(v *HealthProbeParameters) HealthProbeParameters {
		if v != nil {
			return *v
		}
		var ret HealthProbeParameters
		return ret
	}).(HealthProbeParametersOutput)
}

func (o HealthProbeParametersPtrOutput) ProbeIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HealthProbeParameters) *int {
		if v == nil {
			return nil
		}
		return v.ProbeIntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o HealthProbeParametersPtrOutput) ProbePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthProbeParameters) *string {
		if v == nil {
			return nil
		}
		return v.ProbePath
	}).(pulumi.StringPtrOutput)
}

func (o HealthProbeParametersPtrOutput) ProbeProtocol() ProbeProtocolPtrOutput {
	return o.ApplyT(func(v *HealthProbeParameters) *ProbeProtocol {
		if v == nil {
			return nil
		}
		return v.ProbeProtocol
	}).(ProbeProtocolPtrOutput)
}

func (o HealthProbeParametersPtrOutput) ProbeRequestType() HealthProbeRequestTypePtrOutput {
	return o.ApplyT(func(v *HealthProbeParameters) *HealthProbeRequestType {
		if v == nil {
			return nil
		}
		return v.ProbeRequestType
	}).(HealthProbeRequestTypePtrOutput)
}

type HealthProbeParametersResponse struct {
	ProbeIntervalInSeconds *int    `pulumi:"probeIntervalInSeconds"`
	ProbePath              *string `pulumi:"probePath"`
	ProbeProtocol          *string `pulumi:"probeProtocol"`
	ProbeRequestType       *string `pulumi:"probeRequestType"`
}

type HealthProbeParametersResponseOutput struct{ *pulumi.OutputState }

func (HealthProbeParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeParametersResponse)(nil)).Elem()
}

func (o HealthProbeParametersResponseOutput) ToHealthProbeParametersResponseOutput() HealthProbeParametersResponseOutput {
	return o
}

func (o HealthProbeParametersResponseOutput) ToHealthProbeParametersResponseOutputWithContext(ctx context.Context) HealthProbeParametersResponseOutput {
	return o
}

func (o HealthProbeParametersResponseOutput) ProbeIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HealthProbeParametersResponse) *int { return v.ProbeIntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o HealthProbeParametersResponseOutput) ProbePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeParametersResponse) *string { return v.ProbePath }).(pulumi.StringPtrOutput)
}

func (o HealthProbeParametersResponseOutput) ProbeProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeParametersResponse) *string { return v.ProbeProtocol }).(pulumi.StringPtrOutput)
}

func (o HealthProbeParametersResponseOutput) ProbeRequestType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeParametersResponse) *string { return v.ProbeRequestType }).(pulumi.StringPtrOutput)
}

type HealthProbeParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (HealthProbeParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthProbeParametersResponse)(nil)).Elem()
}

func (o HealthProbeParametersResponsePtrOutput) ToHealthProbeParametersResponsePtrOutput() HealthProbeParametersResponsePtrOutput {
	return o
}

func (o HealthProbeParametersResponsePtrOutput) ToHealthProbeParametersResponsePtrOutputWithContext(ctx context.Context) HealthProbeParametersResponsePtrOutput {
	return o
}

func (o HealthProbeParametersResponsePtrOutput) Elem() HealthProbeParametersResponseOutput {
	return o.ApplyT(func(v *HealthProbeParametersResponse) HealthProbeParametersResponse {
		if v != nil {
			return *v
		}
		var ret HealthProbeParametersResponse
		return ret
	}).(HealthProbeParametersResponseOutput)
}

func (o HealthProbeParametersResponsePtrOutput) ProbeIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HealthProbeParametersResponse) *int {
		if v == nil {
			return nil
		}
		return v.ProbeIntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o HealthProbeParametersResponsePtrOutput) ProbePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthProbeParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProbePath
	}).(pulumi.StringPtrOutput)
}

func (o HealthProbeParametersResponsePtrOutput) ProbeProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthProbeParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProbeProtocol
	}).(pulumi.StringPtrOutput)
}

func (o HealthProbeParametersResponsePtrOutput) ProbeRequestType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthProbeParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProbeRequestType
	}).(pulumi.StringPtrOutput)
}

type HttpErrorRangeParameters struct {
	Begin *int `pulumi:"begin"`
	End   *int `pulumi:"end"`
}





type HttpErrorRangeParametersInput interface {
	pulumi.Input

	ToHttpErrorRangeParametersOutput() HttpErrorRangeParametersOutput
	ToHttpErrorRangeParametersOutputWithContext(context.Context) HttpErrorRangeParametersOutput
}

type HttpErrorRangeParametersArgs struct {
	Begin pulumi.IntPtrInput `pulumi:"begin"`
	End   pulumi.IntPtrInput `pulumi:"end"`
}

func (HttpErrorRangeParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpErrorRangeParameters)(nil)).Elem()
}

func (i HttpErrorRangeParametersArgs) ToHttpErrorRangeParametersOutput() HttpErrorRangeParametersOutput {
	return i.ToHttpErrorRangeParametersOutputWithContext(context.Background())
}

func (i HttpErrorRangeParametersArgs) ToHttpErrorRangeParametersOutputWithContext(ctx context.Context) HttpErrorRangeParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpErrorRangeParametersOutput)
}





type HttpErrorRangeParametersArrayInput interface {
	pulumi.Input

	ToHttpErrorRangeParametersArrayOutput() HttpErrorRangeParametersArrayOutput
	ToHttpErrorRangeParametersArrayOutputWithContext(context.Context) HttpErrorRangeParametersArrayOutput
}

type HttpErrorRangeParametersArray []HttpErrorRangeParametersInput

func (HttpErrorRangeParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpErrorRangeParameters)(nil)).Elem()
}

func (i HttpErrorRangeParametersArray) ToHttpErrorRangeParametersArrayOutput() HttpErrorRangeParametersArrayOutput {
	return i.ToHttpErrorRangeParametersArrayOutputWithContext(context.Background())
}

func (i HttpErrorRangeParametersArray) ToHttpErrorRangeParametersArrayOutputWithContext(ctx context.Context) HttpErrorRangeParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpErrorRangeParametersArrayOutput)
}

type HttpErrorRangeParametersOutput struct{ *pulumi.OutputState }

func (HttpErrorRangeParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpErrorRangeParameters)(nil)).Elem()
}

func (o HttpErrorRangeParametersOutput) ToHttpErrorRangeParametersOutput() HttpErrorRangeParametersOutput {
	return o
}

func (o HttpErrorRangeParametersOutput) ToHttpErrorRangeParametersOutputWithContext(ctx context.Context) HttpErrorRangeParametersOutput {
	return o
}

func (o HttpErrorRangeParametersOutput) Begin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HttpErrorRangeParameters) *int { return v.Begin }).(pulumi.IntPtrOutput)
}

func (o HttpErrorRangeParametersOutput) End() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HttpErrorRangeParameters) *int { return v.End }).(pulumi.IntPtrOutput)
}

type HttpErrorRangeParametersArrayOutput struct{ *pulumi.OutputState }

func (HttpErrorRangeParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpErrorRangeParameters)(nil)).Elem()
}

func (o HttpErrorRangeParametersArrayOutput) ToHttpErrorRangeParametersArrayOutput() HttpErrorRangeParametersArrayOutput {
	return o
}

func (o HttpErrorRangeParametersArrayOutput) ToHttpErrorRangeParametersArrayOutputWithContext(ctx context.Context) HttpErrorRangeParametersArrayOutput {
	return o
}

func (o HttpErrorRangeParametersArrayOutput) Index(i pulumi.IntInput) HttpErrorRangeParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpErrorRangeParameters {
		return vs[0].([]HttpErrorRangeParameters)[vs[1].(int)]
	}).(HttpErrorRangeParametersOutput)
}

type HttpErrorRangeParametersResponse struct {
	Begin *int `pulumi:"begin"`
	End   *int `pulumi:"end"`
}

type HttpErrorRangeParametersResponseOutput struct{ *pulumi.OutputState }

func (HttpErrorRangeParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpErrorRangeParametersResponse)(nil)).Elem()
}

func (o HttpErrorRangeParametersResponseOutput) ToHttpErrorRangeParametersResponseOutput() HttpErrorRangeParametersResponseOutput {
	return o
}

func (o HttpErrorRangeParametersResponseOutput) ToHttpErrorRangeParametersResponseOutputWithContext(ctx context.Context) HttpErrorRangeParametersResponseOutput {
	return o
}

func (o HttpErrorRangeParametersResponseOutput) Begin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HttpErrorRangeParametersResponse) *int { return v.Begin }).(pulumi.IntPtrOutput)
}

func (o HttpErrorRangeParametersResponseOutput) End() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HttpErrorRangeParametersResponse) *int { return v.End }).(pulumi.IntPtrOutput)
}

type HttpErrorRangeParametersResponseArrayOutput struct{ *pulumi.OutputState }

func (HttpErrorRangeParametersResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpErrorRangeParametersResponse)(nil)).Elem()
}

func (o HttpErrorRangeParametersResponseArrayOutput) ToHttpErrorRangeParametersResponseArrayOutput() HttpErrorRangeParametersResponseArrayOutput {
	return o
}

func (o HttpErrorRangeParametersResponseArrayOutput) ToHttpErrorRangeParametersResponseArrayOutputWithContext(ctx context.Context) HttpErrorRangeParametersResponseArrayOutput {
	return o
}

func (o HttpErrorRangeParametersResponseArrayOutput) Index(i pulumi.IntInput) HttpErrorRangeParametersResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpErrorRangeParametersResponse {
		return vs[0].([]HttpErrorRangeParametersResponse)[vs[1].(int)]
	}).(HttpErrorRangeParametersResponseOutput)
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

type OriginGroupOverrideAction struct {
	Name       string                              `pulumi:"name"`
	Parameters OriginGroupOverrideActionParameters `pulumi:"parameters"`
}

type OriginGroupOverrideActionParameters struct {
	OdataType   string            `pulumi:"odataType"`
	OriginGroup ResourceReference `pulumi:"originGroup"`
}

type OriginGroupOverrideActionParametersResponse struct {
	OdataType   string                    `pulumi:"odataType"`
	OriginGroup ResourceReferenceResponse `pulumi:"originGroup"`
}

type OriginGroupOverrideActionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters OriginGroupOverrideActionParametersResponse `pulumi:"parameters"`
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

type ResourceReference struct {
	Id *string `pulumi:"id"`
}





type ResourceReferenceInput interface {
	pulumi.Input

	ToResourceReferenceOutput() ResourceReferenceOutput
	ToResourceReferenceOutputWithContext(context.Context) ResourceReferenceOutput
}

type ResourceReferenceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArgs) ToResourceReferenceOutput() ResourceReferenceOutput {
	return i.ToResourceReferenceOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput)
}

func (i ResourceReferenceArgs) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return i.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput).ToResourceReferencePtrOutputWithContext(ctx)
}









type ResourceReferencePtrInput interface {
	pulumi.Input

	ToResourceReferencePtrOutput() ResourceReferencePtrOutput
	ToResourceReferencePtrOutputWithContext(context.Context) ResourceReferencePtrOutput
}

type resourceReferencePtrType ResourceReferenceArgs

func ResourceReferencePtr(v *ResourceReferenceArgs) ResourceReferencePtrInput {
	return (*resourceReferencePtrType)(v)
}

func (*resourceReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReference)(nil)).Elem()
}

func (i *resourceReferencePtrType) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return i.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (i *resourceReferencePtrType) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferencePtrOutput)
}





type ResourceReferenceArrayInput interface {
	pulumi.Input

	ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput
	ToResourceReferenceArrayOutputWithContext(context.Context) ResourceReferenceArrayOutput
}

type ResourceReferenceArray []ResourceReferenceInput

func (ResourceReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArray) ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput {
	return i.ToResourceReferenceArrayOutputWithContext(context.Background())
}

func (i ResourceReferenceArray) ToResourceReferenceArrayOutputWithContext(ctx context.Context) ResourceReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceArrayOutput)
}

type ResourceReferenceOutput struct{ *pulumi.OutputState }

func (ResourceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceOutput) ToResourceReferenceOutput() ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return o.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (o ResourceReferenceOutput) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceReference) *ResourceReference {
		return &v
	}).(ResourceReferencePtrOutput)
}

func (o ResourceReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferencePtrOutput struct{ *pulumi.OutputState }

func (ResourceReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReference)(nil)).Elem()
}

func (o ResourceReferencePtrOutput) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return o
}

func (o ResourceReferencePtrOutput) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return o
}

func (o ResourceReferencePtrOutput) Elem() ResourceReferenceOutput {
	return o.ApplyT(func(v *ResourceReference) ResourceReference {
		if v != nil {
			return *v
		}
		var ret ResourceReference
		return ret
	}).(ResourceReferenceOutput)
}

func (o ResourceReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourceReferenceArrayOutput struct{ *pulumi.OutputState }

func (ResourceReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceArrayOutput) ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput {
	return o
}

func (o ResourceReferenceArrayOutput) ToResourceReferenceArrayOutputWithContext(ctx context.Context) ResourceReferenceArrayOutput {
	return o
}

func (o ResourceReferenceArrayOutput) Index(i pulumi.IntInput) ResourceReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceReference {
		return vs[0].([]ResourceReference)[vs[1].(int)]
	}).(ResourceReferenceOutput)
}

type ResourceReferenceResponse struct {
	Id *string `pulumi:"id"`
}

type ResourceReferenceResponseOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) Elem() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) ResourceReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ResourceReferenceResponse
		return ret
	}).(ResourceReferenceResponseOutput)
}

func (o ResourceReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourceReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutputWithContext(ctx context.Context) ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) Index(i pulumi.IntInput) ResourceReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceReferenceResponse {
		return vs[0].([]ResourceReferenceResponse)[vs[1].(int)]
	}).(ResourceReferenceResponseOutput)
}

type ResponseBasedOriginErrorDetectionParameters struct {
	HttpErrorRanges                          []HttpErrorRangeParameters       `pulumi:"httpErrorRanges"`
	ResponseBasedDetectedErrorTypes          *ResponseBasedDetectedErrorTypes `pulumi:"responseBasedDetectedErrorTypes"`
	ResponseBasedFailoverThresholdPercentage *int                             `pulumi:"responseBasedFailoverThresholdPercentage"`
}





type ResponseBasedOriginErrorDetectionParametersInput interface {
	pulumi.Input

	ToResponseBasedOriginErrorDetectionParametersOutput() ResponseBasedOriginErrorDetectionParametersOutput
	ToResponseBasedOriginErrorDetectionParametersOutputWithContext(context.Context) ResponseBasedOriginErrorDetectionParametersOutput
}

type ResponseBasedOriginErrorDetectionParametersArgs struct {
	HttpErrorRanges                          HttpErrorRangeParametersArrayInput      `pulumi:"httpErrorRanges"`
	ResponseBasedDetectedErrorTypes          ResponseBasedDetectedErrorTypesPtrInput `pulumi:"responseBasedDetectedErrorTypes"`
	ResponseBasedFailoverThresholdPercentage pulumi.IntPtrInput                      `pulumi:"responseBasedFailoverThresholdPercentage"`
}

func (ResponseBasedOriginErrorDetectionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseBasedOriginErrorDetectionParameters)(nil)).Elem()
}

func (i ResponseBasedOriginErrorDetectionParametersArgs) ToResponseBasedOriginErrorDetectionParametersOutput() ResponseBasedOriginErrorDetectionParametersOutput {
	return i.ToResponseBasedOriginErrorDetectionParametersOutputWithContext(context.Background())
}

func (i ResponseBasedOriginErrorDetectionParametersArgs) ToResponseBasedOriginErrorDetectionParametersOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseBasedOriginErrorDetectionParametersOutput)
}

func (i ResponseBasedOriginErrorDetectionParametersArgs) ToResponseBasedOriginErrorDetectionParametersPtrOutput() ResponseBasedOriginErrorDetectionParametersPtrOutput {
	return i.ToResponseBasedOriginErrorDetectionParametersPtrOutputWithContext(context.Background())
}

func (i ResponseBasedOriginErrorDetectionParametersArgs) ToResponseBasedOriginErrorDetectionParametersPtrOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseBasedOriginErrorDetectionParametersOutput).ToResponseBasedOriginErrorDetectionParametersPtrOutputWithContext(ctx)
}









type ResponseBasedOriginErrorDetectionParametersPtrInput interface {
	pulumi.Input

	ToResponseBasedOriginErrorDetectionParametersPtrOutput() ResponseBasedOriginErrorDetectionParametersPtrOutput
	ToResponseBasedOriginErrorDetectionParametersPtrOutputWithContext(context.Context) ResponseBasedOriginErrorDetectionParametersPtrOutput
}

type responseBasedOriginErrorDetectionParametersPtrType ResponseBasedOriginErrorDetectionParametersArgs

func ResponseBasedOriginErrorDetectionParametersPtr(v *ResponseBasedOriginErrorDetectionParametersArgs) ResponseBasedOriginErrorDetectionParametersPtrInput {
	return (*responseBasedOriginErrorDetectionParametersPtrType)(v)
}

func (*responseBasedOriginErrorDetectionParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseBasedOriginErrorDetectionParameters)(nil)).Elem()
}

func (i *responseBasedOriginErrorDetectionParametersPtrType) ToResponseBasedOriginErrorDetectionParametersPtrOutput() ResponseBasedOriginErrorDetectionParametersPtrOutput {
	return i.ToResponseBasedOriginErrorDetectionParametersPtrOutputWithContext(context.Background())
}

func (i *responseBasedOriginErrorDetectionParametersPtrType) ToResponseBasedOriginErrorDetectionParametersPtrOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseBasedOriginErrorDetectionParametersPtrOutput)
}

type ResponseBasedOriginErrorDetectionParametersOutput struct{ *pulumi.OutputState }

func (ResponseBasedOriginErrorDetectionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseBasedOriginErrorDetectionParameters)(nil)).Elem()
}

func (o ResponseBasedOriginErrorDetectionParametersOutput) ToResponseBasedOriginErrorDetectionParametersOutput() ResponseBasedOriginErrorDetectionParametersOutput {
	return o
}

func (o ResponseBasedOriginErrorDetectionParametersOutput) ToResponseBasedOriginErrorDetectionParametersOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersOutput {
	return o
}

func (o ResponseBasedOriginErrorDetectionParametersOutput) ToResponseBasedOriginErrorDetectionParametersPtrOutput() ResponseBasedOriginErrorDetectionParametersPtrOutput {
	return o.ToResponseBasedOriginErrorDetectionParametersPtrOutputWithContext(context.Background())
}

func (o ResponseBasedOriginErrorDetectionParametersOutput) ToResponseBasedOriginErrorDetectionParametersPtrOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResponseBasedOriginErrorDetectionParameters) *ResponseBasedOriginErrorDetectionParameters {
		return &v
	}).(ResponseBasedOriginErrorDetectionParametersPtrOutput)
}

func (o ResponseBasedOriginErrorDetectionParametersOutput) HttpErrorRanges() HttpErrorRangeParametersArrayOutput {
	return o.ApplyT(func(v ResponseBasedOriginErrorDetectionParameters) []HttpErrorRangeParameters {
		return v.HttpErrorRanges
	}).(HttpErrorRangeParametersArrayOutput)
}

func (o ResponseBasedOriginErrorDetectionParametersOutput) ResponseBasedDetectedErrorTypes() ResponseBasedDetectedErrorTypesPtrOutput {
	return o.ApplyT(func(v ResponseBasedOriginErrorDetectionParameters) *ResponseBasedDetectedErrorTypes {
		return v.ResponseBasedDetectedErrorTypes
	}).(ResponseBasedDetectedErrorTypesPtrOutput)
}

func (o ResponseBasedOriginErrorDetectionParametersOutput) ResponseBasedFailoverThresholdPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResponseBasedOriginErrorDetectionParameters) *int {
		return v.ResponseBasedFailoverThresholdPercentage
	}).(pulumi.IntPtrOutput)
}

type ResponseBasedOriginErrorDetectionParametersPtrOutput struct{ *pulumi.OutputState }

func (ResponseBasedOriginErrorDetectionParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseBasedOriginErrorDetectionParameters)(nil)).Elem()
}

func (o ResponseBasedOriginErrorDetectionParametersPtrOutput) ToResponseBasedOriginErrorDetectionParametersPtrOutput() ResponseBasedOriginErrorDetectionParametersPtrOutput {
	return o
}

func (o ResponseBasedOriginErrorDetectionParametersPtrOutput) ToResponseBasedOriginErrorDetectionParametersPtrOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersPtrOutput {
	return o
}

func (o ResponseBasedOriginErrorDetectionParametersPtrOutput) Elem() ResponseBasedOriginErrorDetectionParametersOutput {
	return o.ApplyT(func(v *ResponseBasedOriginErrorDetectionParameters) ResponseBasedOriginErrorDetectionParameters {
		if v != nil {
			return *v
		}
		var ret ResponseBasedOriginErrorDetectionParameters
		return ret
	}).(ResponseBasedOriginErrorDetectionParametersOutput)
}

func (o ResponseBasedOriginErrorDetectionParametersPtrOutput) HttpErrorRanges() HttpErrorRangeParametersArrayOutput {
	return o.ApplyT(func(v *ResponseBasedOriginErrorDetectionParameters) []HttpErrorRangeParameters {
		if v == nil {
			return nil
		}
		return v.HttpErrorRanges
	}).(HttpErrorRangeParametersArrayOutput)
}

func (o ResponseBasedOriginErrorDetectionParametersPtrOutput) ResponseBasedDetectedErrorTypes() ResponseBasedDetectedErrorTypesPtrOutput {
	return o.ApplyT(func(v *ResponseBasedOriginErrorDetectionParameters) *ResponseBasedDetectedErrorTypes {
		if v == nil {
			return nil
		}
		return v.ResponseBasedDetectedErrorTypes
	}).(ResponseBasedDetectedErrorTypesPtrOutput)
}

func (o ResponseBasedOriginErrorDetectionParametersPtrOutput) ResponseBasedFailoverThresholdPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResponseBasedOriginErrorDetectionParameters) *int {
		if v == nil {
			return nil
		}
		return v.ResponseBasedFailoverThresholdPercentage
	}).(pulumi.IntPtrOutput)
}

type ResponseBasedOriginErrorDetectionParametersResponse struct {
	HttpErrorRanges                          []HttpErrorRangeParametersResponse `pulumi:"httpErrorRanges"`
	ResponseBasedDetectedErrorTypes          *string                            `pulumi:"responseBasedDetectedErrorTypes"`
	ResponseBasedFailoverThresholdPercentage *int                               `pulumi:"responseBasedFailoverThresholdPercentage"`
}

type ResponseBasedOriginErrorDetectionParametersResponseOutput struct{ *pulumi.OutputState }

func (ResponseBasedOriginErrorDetectionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseBasedOriginErrorDetectionParametersResponse)(nil)).Elem()
}

func (o ResponseBasedOriginErrorDetectionParametersResponseOutput) ToResponseBasedOriginErrorDetectionParametersResponseOutput() ResponseBasedOriginErrorDetectionParametersResponseOutput {
	return o
}

func (o ResponseBasedOriginErrorDetectionParametersResponseOutput) ToResponseBasedOriginErrorDetectionParametersResponseOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersResponseOutput {
	return o
}

func (o ResponseBasedOriginErrorDetectionParametersResponseOutput) HttpErrorRanges() HttpErrorRangeParametersResponseArrayOutput {
	return o.ApplyT(func(v ResponseBasedOriginErrorDetectionParametersResponse) []HttpErrorRangeParametersResponse {
		return v.HttpErrorRanges
	}).(HttpErrorRangeParametersResponseArrayOutput)
}

func (o ResponseBasedOriginErrorDetectionParametersResponseOutput) ResponseBasedDetectedErrorTypes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseBasedOriginErrorDetectionParametersResponse) *string {
		return v.ResponseBasedDetectedErrorTypes
	}).(pulumi.StringPtrOutput)
}

func (o ResponseBasedOriginErrorDetectionParametersResponseOutput) ResponseBasedFailoverThresholdPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResponseBasedOriginErrorDetectionParametersResponse) *int {
		return v.ResponseBasedFailoverThresholdPercentage
	}).(pulumi.IntPtrOutput)
}

type ResponseBasedOriginErrorDetectionParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (ResponseBasedOriginErrorDetectionParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseBasedOriginErrorDetectionParametersResponse)(nil)).Elem()
}

func (o ResponseBasedOriginErrorDetectionParametersResponsePtrOutput) ToResponseBasedOriginErrorDetectionParametersResponsePtrOutput() ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return o
}

func (o ResponseBasedOriginErrorDetectionParametersResponsePtrOutput) ToResponseBasedOriginErrorDetectionParametersResponsePtrOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return o
}

func (o ResponseBasedOriginErrorDetectionParametersResponsePtrOutput) Elem() ResponseBasedOriginErrorDetectionParametersResponseOutput {
	return o.ApplyT(func(v *ResponseBasedOriginErrorDetectionParametersResponse) ResponseBasedOriginErrorDetectionParametersResponse {
		if v != nil {
			return *v
		}
		var ret ResponseBasedOriginErrorDetectionParametersResponse
		return ret
	}).(ResponseBasedOriginErrorDetectionParametersResponseOutput)
}

func (o ResponseBasedOriginErrorDetectionParametersResponsePtrOutput) HttpErrorRanges() HttpErrorRangeParametersResponseArrayOutput {
	return o.ApplyT(func(v *ResponseBasedOriginErrorDetectionParametersResponse) []HttpErrorRangeParametersResponse {
		if v == nil {
			return nil
		}
		return v.HttpErrorRanges
	}).(HttpErrorRangeParametersResponseArrayOutput)
}

func (o ResponseBasedOriginErrorDetectionParametersResponsePtrOutput) ResponseBasedDetectedErrorTypes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResponseBasedOriginErrorDetectionParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResponseBasedDetectedErrorTypes
	}).(pulumi.StringPtrOutput)
}

func (o ResponseBasedOriginErrorDetectionParametersResponsePtrOutput) ResponseBasedFailoverThresholdPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResponseBasedOriginErrorDetectionParametersResponse) *int {
		if v == nil {
			return nil
		}
		return v.ResponseBasedFailoverThresholdPercentage
	}).(pulumi.IntPtrOutput)
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
	pulumi.RegisterOutputType(DeepCreatedOriginOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginArrayOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginGroupOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginGroupArrayOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginGroupResponseOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginGroupResponseArrayOutput{})
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
	pulumi.RegisterOutputType(GeoFilterOutput{})
	pulumi.RegisterOutputType(GeoFilterArrayOutput{})
	pulumi.RegisterOutputType(GeoFilterResponseOutput{})
	pulumi.RegisterOutputType(GeoFilterResponseArrayOutput{})
	pulumi.RegisterOutputType(HealthProbeParametersOutput{})
	pulumi.RegisterOutputType(HealthProbeParametersPtrOutput{})
	pulumi.RegisterOutputType(HealthProbeParametersResponseOutput{})
	pulumi.RegisterOutputType(HealthProbeParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersArrayOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersResponseOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceReferenceOutput{})
	pulumi.RegisterOutputType(ResourceReferencePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceArrayOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(ResponseBasedOriginErrorDetectionParametersOutput{})
	pulumi.RegisterOutputType(ResponseBasedOriginErrorDetectionParametersPtrOutput{})
	pulumi.RegisterOutputType(ResponseBasedOriginErrorDetectionParametersResponseOutput{})
	pulumi.RegisterOutputType(ResponseBasedOriginErrorDetectionParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
}
