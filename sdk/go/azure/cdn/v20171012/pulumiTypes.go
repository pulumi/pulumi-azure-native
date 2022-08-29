


package v20171012

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





type CacheExpirationActionParametersInput interface {
	pulumi.Input

	ToCacheExpirationActionParametersOutput() CacheExpirationActionParametersOutput
	ToCacheExpirationActionParametersOutputWithContext(context.Context) CacheExpirationActionParametersOutput
}

type CacheExpirationActionParametersArgs struct {
	CacheBehavior pulumi.StringInput    `pulumi:"cacheBehavior"`
	CacheDuration pulumi.StringPtrInput `pulumi:"cacheDuration"`
	CacheType     pulumi.StringInput    `pulumi:"cacheType"`
	OdataType     pulumi.StringInput    `pulumi:"odataType"`
}

func (CacheExpirationActionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheExpirationActionParameters)(nil)).Elem()
}

func (i CacheExpirationActionParametersArgs) ToCacheExpirationActionParametersOutput() CacheExpirationActionParametersOutput {
	return i.ToCacheExpirationActionParametersOutputWithContext(context.Background())
}

func (i CacheExpirationActionParametersArgs) ToCacheExpirationActionParametersOutputWithContext(ctx context.Context) CacheExpirationActionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheExpirationActionParametersOutput)
}

type CacheExpirationActionParametersOutput struct{ *pulumi.OutputState }

func (CacheExpirationActionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheExpirationActionParameters)(nil)).Elem()
}

func (o CacheExpirationActionParametersOutput) ToCacheExpirationActionParametersOutput() CacheExpirationActionParametersOutput {
	return o
}

func (o CacheExpirationActionParametersOutput) ToCacheExpirationActionParametersOutputWithContext(ctx context.Context) CacheExpirationActionParametersOutput {
	return o
}

func (o CacheExpirationActionParametersOutput) CacheBehavior() pulumi.StringOutput {
	return o.ApplyT(func(v CacheExpirationActionParameters) string { return v.CacheBehavior }).(pulumi.StringOutput)
}

func (o CacheExpirationActionParametersOutput) CacheDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheExpirationActionParameters) *string { return v.CacheDuration }).(pulumi.StringPtrOutput)
}

func (o CacheExpirationActionParametersOutput) CacheType() pulumi.StringOutput {
	return o.ApplyT(func(v CacheExpirationActionParameters) string { return v.CacheType }).(pulumi.StringOutput)
}

func (o CacheExpirationActionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v CacheExpirationActionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

type CacheExpirationActionParametersResponse struct {
	CacheBehavior string  `pulumi:"cacheBehavior"`
	CacheDuration *string `pulumi:"cacheDuration"`
	CacheType     string  `pulumi:"cacheType"`
	OdataType     string  `pulumi:"odataType"`
}

type CacheExpirationActionParametersResponseOutput struct{ *pulumi.OutputState }

func (CacheExpirationActionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheExpirationActionParametersResponse)(nil)).Elem()
}

func (o CacheExpirationActionParametersResponseOutput) ToCacheExpirationActionParametersResponseOutput() CacheExpirationActionParametersResponseOutput {
	return o
}

func (o CacheExpirationActionParametersResponseOutput) ToCacheExpirationActionParametersResponseOutputWithContext(ctx context.Context) CacheExpirationActionParametersResponseOutput {
	return o
}

func (o CacheExpirationActionParametersResponseOutput) CacheBehavior() pulumi.StringOutput {
	return o.ApplyT(func(v CacheExpirationActionParametersResponse) string { return v.CacheBehavior }).(pulumi.StringOutput)
}

func (o CacheExpirationActionParametersResponseOutput) CacheDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheExpirationActionParametersResponse) *string { return v.CacheDuration }).(pulumi.StringPtrOutput)
}

func (o CacheExpirationActionParametersResponseOutput) CacheType() pulumi.StringOutput {
	return o.ApplyT(func(v CacheExpirationActionParametersResponse) string { return v.CacheType }).(pulumi.StringOutput)
}

func (o CacheExpirationActionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v CacheExpirationActionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
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
	Actions    []DeliveryRuleCacheExpirationAction `pulumi:"actions"`
	Conditions []interface{}                       `pulumi:"conditions"`
	Order      int                                 `pulumi:"order"`
}





type DeliveryRuleInput interface {
	pulumi.Input

	ToDeliveryRuleOutput() DeliveryRuleOutput
	ToDeliveryRuleOutputWithContext(context.Context) DeliveryRuleOutput
}

type DeliveryRuleArgs struct {
	Actions    DeliveryRuleCacheExpirationActionArrayInput `pulumi:"actions"`
	Conditions pulumi.ArrayInput                           `pulumi:"conditions"`
	Order      pulumi.IntInput                             `pulumi:"order"`
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

func (o DeliveryRuleOutput) Actions() DeliveryRuleCacheExpirationActionArrayOutput {
	return o.ApplyT(func(v DeliveryRule) []DeliveryRuleCacheExpirationAction { return v.Actions }).(DeliveryRuleCacheExpirationActionArrayOutput)
}

func (o DeliveryRuleOutput) Conditions() pulumi.ArrayOutput {
	return o.ApplyT(func(v DeliveryRule) []interface{} { return v.Conditions }).(pulumi.ArrayOutput)
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





type DeliveryRuleCacheExpirationActionInput interface {
	pulumi.Input

	ToDeliveryRuleCacheExpirationActionOutput() DeliveryRuleCacheExpirationActionOutput
	ToDeliveryRuleCacheExpirationActionOutputWithContext(context.Context) DeliveryRuleCacheExpirationActionOutput
}

type DeliveryRuleCacheExpirationActionArgs struct {
	Name       pulumi.StringInput                   `pulumi:"name"`
	Parameters CacheExpirationActionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleCacheExpirationActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCacheExpirationAction)(nil)).Elem()
}

func (i DeliveryRuleCacheExpirationActionArgs) ToDeliveryRuleCacheExpirationActionOutput() DeliveryRuleCacheExpirationActionOutput {
	return i.ToDeliveryRuleCacheExpirationActionOutputWithContext(context.Background())
}

func (i DeliveryRuleCacheExpirationActionArgs) ToDeliveryRuleCacheExpirationActionOutputWithContext(ctx context.Context) DeliveryRuleCacheExpirationActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleCacheExpirationActionOutput)
}





type DeliveryRuleCacheExpirationActionArrayInput interface {
	pulumi.Input

	ToDeliveryRuleCacheExpirationActionArrayOutput() DeliveryRuleCacheExpirationActionArrayOutput
	ToDeliveryRuleCacheExpirationActionArrayOutputWithContext(context.Context) DeliveryRuleCacheExpirationActionArrayOutput
}

type DeliveryRuleCacheExpirationActionArray []DeliveryRuleCacheExpirationActionInput

func (DeliveryRuleCacheExpirationActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeliveryRuleCacheExpirationAction)(nil)).Elem()
}

func (i DeliveryRuleCacheExpirationActionArray) ToDeliveryRuleCacheExpirationActionArrayOutput() DeliveryRuleCacheExpirationActionArrayOutput {
	return i.ToDeliveryRuleCacheExpirationActionArrayOutputWithContext(context.Background())
}

func (i DeliveryRuleCacheExpirationActionArray) ToDeliveryRuleCacheExpirationActionArrayOutputWithContext(ctx context.Context) DeliveryRuleCacheExpirationActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleCacheExpirationActionArrayOutput)
}

type DeliveryRuleCacheExpirationActionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleCacheExpirationActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCacheExpirationAction)(nil)).Elem()
}

func (o DeliveryRuleCacheExpirationActionOutput) ToDeliveryRuleCacheExpirationActionOutput() DeliveryRuleCacheExpirationActionOutput {
	return o
}

func (o DeliveryRuleCacheExpirationActionOutput) ToDeliveryRuleCacheExpirationActionOutputWithContext(ctx context.Context) DeliveryRuleCacheExpirationActionOutput {
	return o
}

func (o DeliveryRuleCacheExpirationActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleCacheExpirationAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleCacheExpirationActionOutput) Parameters() CacheExpirationActionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleCacheExpirationAction) CacheExpirationActionParameters { return v.Parameters }).(CacheExpirationActionParametersOutput)
}

type DeliveryRuleCacheExpirationActionArrayOutput struct{ *pulumi.OutputState }

func (DeliveryRuleCacheExpirationActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeliveryRuleCacheExpirationAction)(nil)).Elem()
}

func (o DeliveryRuleCacheExpirationActionArrayOutput) ToDeliveryRuleCacheExpirationActionArrayOutput() DeliveryRuleCacheExpirationActionArrayOutput {
	return o
}

func (o DeliveryRuleCacheExpirationActionArrayOutput) ToDeliveryRuleCacheExpirationActionArrayOutputWithContext(ctx context.Context) DeliveryRuleCacheExpirationActionArrayOutput {
	return o
}

func (o DeliveryRuleCacheExpirationActionArrayOutput) Index(i pulumi.IntInput) DeliveryRuleCacheExpirationActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeliveryRuleCacheExpirationAction {
		return vs[0].([]DeliveryRuleCacheExpirationAction)[vs[1].(int)]
	}).(DeliveryRuleCacheExpirationActionOutput)
}

type DeliveryRuleCacheExpirationActionResponse struct {
	Name       string                                  `pulumi:"name"`
	Parameters CacheExpirationActionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleCacheExpirationActionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleCacheExpirationActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCacheExpirationActionResponse)(nil)).Elem()
}

func (o DeliveryRuleCacheExpirationActionResponseOutput) ToDeliveryRuleCacheExpirationActionResponseOutput() DeliveryRuleCacheExpirationActionResponseOutput {
	return o
}

func (o DeliveryRuleCacheExpirationActionResponseOutput) ToDeliveryRuleCacheExpirationActionResponseOutputWithContext(ctx context.Context) DeliveryRuleCacheExpirationActionResponseOutput {
	return o
}

func (o DeliveryRuleCacheExpirationActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleCacheExpirationActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleCacheExpirationActionResponseOutput) Parameters() CacheExpirationActionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleCacheExpirationActionResponse) CacheExpirationActionParametersResponse {
		return v.Parameters
	}).(CacheExpirationActionParametersResponseOutput)
}

type DeliveryRuleCacheExpirationActionResponseArrayOutput struct{ *pulumi.OutputState }

func (DeliveryRuleCacheExpirationActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeliveryRuleCacheExpirationActionResponse)(nil)).Elem()
}

func (o DeliveryRuleCacheExpirationActionResponseArrayOutput) ToDeliveryRuleCacheExpirationActionResponseArrayOutput() DeliveryRuleCacheExpirationActionResponseArrayOutput {
	return o
}

func (o DeliveryRuleCacheExpirationActionResponseArrayOutput) ToDeliveryRuleCacheExpirationActionResponseArrayOutputWithContext(ctx context.Context) DeliveryRuleCacheExpirationActionResponseArrayOutput {
	return o
}

func (o DeliveryRuleCacheExpirationActionResponseArrayOutput) Index(i pulumi.IntInput) DeliveryRuleCacheExpirationActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeliveryRuleCacheExpirationActionResponse {
		return vs[0].([]DeliveryRuleCacheExpirationActionResponse)[vs[1].(int)]
	}).(DeliveryRuleCacheExpirationActionResponseOutput)
}

type DeliveryRuleResponse struct {
	Actions    []DeliveryRuleCacheExpirationActionResponse `pulumi:"actions"`
	Conditions []interface{}                               `pulumi:"conditions"`
	Order      int                                         `pulumi:"order"`
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

func (o DeliveryRuleResponseOutput) Actions() DeliveryRuleCacheExpirationActionResponseArrayOutput {
	return o.ApplyT(func(v DeliveryRuleResponse) []DeliveryRuleCacheExpirationActionResponse { return v.Actions }).(DeliveryRuleCacheExpirationActionResponseArrayOutput)
}

func (o DeliveryRuleResponseOutput) Conditions() pulumi.ArrayOutput {
	return o.ApplyT(func(v DeliveryRuleResponse) []interface{} { return v.Conditions }).(pulumi.ArrayOutput)
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

type DeliveryRuleUrlFileExtensionCondition struct {
	Name       string                              `pulumi:"name"`
	Parameters UrlFileExtensionConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleUrlFileExtensionConditionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters UrlFileExtensionConditionParametersResponse `pulumi:"parameters"`
}

type DeliveryRuleUrlPathCondition struct {
	Name       string                     `pulumi:"name"`
	Parameters UrlPathConditionParameters `pulumi:"parameters"`
}

type DeliveryRuleUrlPathConditionResponse struct {
	Name       string                             `pulumi:"name"`
	Parameters UrlPathConditionParametersResponse `pulumi:"parameters"`
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

type UrlFileExtensionConditionParameters struct {
	Extensions []string `pulumi:"extensions"`
	OdataType  string   `pulumi:"odataType"`
}

type UrlFileExtensionConditionParametersResponse struct {
	Extensions []string `pulumi:"extensions"`
	OdataType  string   `pulumi:"odataType"`
}

type UrlPathConditionParameters struct {
	MatchType string `pulumi:"matchType"`
	OdataType string `pulumi:"odataType"`
	Path      string `pulumi:"path"`
}

type UrlPathConditionParametersResponse struct {
	MatchType string `pulumi:"matchType"`
	OdataType string `pulumi:"odataType"`
	Path      string `pulumi:"path"`
}

func init() {
	pulumi.RegisterOutputType(CacheExpirationActionParametersOutput{})
	pulumi.RegisterOutputType(CacheExpirationActionParametersResponseOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginArrayOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginResponseOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginResponseArrayOutput{})
	pulumi.RegisterOutputType(DeliveryRuleOutput{})
	pulumi.RegisterOutputType(DeliveryRuleArrayOutput{})
	pulumi.RegisterOutputType(DeliveryRuleCacheExpirationActionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleCacheExpirationActionArrayOutput{})
	pulumi.RegisterOutputType(DeliveryRuleCacheExpirationActionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleCacheExpirationActionResponseArrayOutput{})
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
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
}
