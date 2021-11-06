


package v20200331

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





type CacheExpirationActionParametersResponseInput interface {
	pulumi.Input

	ToCacheExpirationActionParametersResponseOutput() CacheExpirationActionParametersResponseOutput
	ToCacheExpirationActionParametersResponseOutputWithContext(context.Context) CacheExpirationActionParametersResponseOutput
}

type CacheExpirationActionParametersResponseArgs struct {
	CacheBehavior pulumi.StringInput    `pulumi:"cacheBehavior"`
	CacheDuration pulumi.StringPtrInput `pulumi:"cacheDuration"`
	CacheType     pulumi.StringInput    `pulumi:"cacheType"`
	OdataType     pulumi.StringInput    `pulumi:"odataType"`
}

func (CacheExpirationActionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheExpirationActionParametersResponse)(nil)).Elem()
}

func (i CacheExpirationActionParametersResponseArgs) ToCacheExpirationActionParametersResponseOutput() CacheExpirationActionParametersResponseOutput {
	return i.ToCacheExpirationActionParametersResponseOutputWithContext(context.Background())
}

func (i CacheExpirationActionParametersResponseArgs) ToCacheExpirationActionParametersResponseOutputWithContext(ctx context.Context) CacheExpirationActionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheExpirationActionParametersResponseOutput)
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

type CacheKeyQueryStringActionParameters struct {
	OdataType           string  `pulumi:"odataType"`
	QueryParameters     *string `pulumi:"queryParameters"`
	QueryStringBehavior string  `pulumi:"queryStringBehavior"`
}





type CacheKeyQueryStringActionParametersInput interface {
	pulumi.Input

	ToCacheKeyQueryStringActionParametersOutput() CacheKeyQueryStringActionParametersOutput
	ToCacheKeyQueryStringActionParametersOutputWithContext(context.Context) CacheKeyQueryStringActionParametersOutput
}

type CacheKeyQueryStringActionParametersArgs struct {
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
	QueryParameters     pulumi.StringPtrInput `pulumi:"queryParameters"`
	QueryStringBehavior pulumi.StringInput    `pulumi:"queryStringBehavior"`
}

func (CacheKeyQueryStringActionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheKeyQueryStringActionParameters)(nil)).Elem()
}

func (i CacheKeyQueryStringActionParametersArgs) ToCacheKeyQueryStringActionParametersOutput() CacheKeyQueryStringActionParametersOutput {
	return i.ToCacheKeyQueryStringActionParametersOutputWithContext(context.Background())
}

func (i CacheKeyQueryStringActionParametersArgs) ToCacheKeyQueryStringActionParametersOutputWithContext(ctx context.Context) CacheKeyQueryStringActionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheKeyQueryStringActionParametersOutput)
}

type CacheKeyQueryStringActionParametersOutput struct{ *pulumi.OutputState }

func (CacheKeyQueryStringActionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheKeyQueryStringActionParameters)(nil)).Elem()
}

func (o CacheKeyQueryStringActionParametersOutput) ToCacheKeyQueryStringActionParametersOutput() CacheKeyQueryStringActionParametersOutput {
	return o
}

func (o CacheKeyQueryStringActionParametersOutput) ToCacheKeyQueryStringActionParametersOutputWithContext(ctx context.Context) CacheKeyQueryStringActionParametersOutput {
	return o
}

func (o CacheKeyQueryStringActionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v CacheKeyQueryStringActionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o CacheKeyQueryStringActionParametersOutput) QueryParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheKeyQueryStringActionParameters) *string { return v.QueryParameters }).(pulumi.StringPtrOutput)
}

func (o CacheKeyQueryStringActionParametersOutput) QueryStringBehavior() pulumi.StringOutput {
	return o.ApplyT(func(v CacheKeyQueryStringActionParameters) string { return v.QueryStringBehavior }).(pulumi.StringOutput)
}

type CacheKeyQueryStringActionParametersResponse struct {
	OdataType           string  `pulumi:"odataType"`
	QueryParameters     *string `pulumi:"queryParameters"`
	QueryStringBehavior string  `pulumi:"queryStringBehavior"`
}





type CacheKeyQueryStringActionParametersResponseInput interface {
	pulumi.Input

	ToCacheKeyQueryStringActionParametersResponseOutput() CacheKeyQueryStringActionParametersResponseOutput
	ToCacheKeyQueryStringActionParametersResponseOutputWithContext(context.Context) CacheKeyQueryStringActionParametersResponseOutput
}

type CacheKeyQueryStringActionParametersResponseArgs struct {
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
	QueryParameters     pulumi.StringPtrInput `pulumi:"queryParameters"`
	QueryStringBehavior pulumi.StringInput    `pulumi:"queryStringBehavior"`
}

func (CacheKeyQueryStringActionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheKeyQueryStringActionParametersResponse)(nil)).Elem()
}

func (i CacheKeyQueryStringActionParametersResponseArgs) ToCacheKeyQueryStringActionParametersResponseOutput() CacheKeyQueryStringActionParametersResponseOutput {
	return i.ToCacheKeyQueryStringActionParametersResponseOutputWithContext(context.Background())
}

func (i CacheKeyQueryStringActionParametersResponseArgs) ToCacheKeyQueryStringActionParametersResponseOutputWithContext(ctx context.Context) CacheKeyQueryStringActionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheKeyQueryStringActionParametersResponseOutput)
}

type CacheKeyQueryStringActionParametersResponseOutput struct{ *pulumi.OutputState }

func (CacheKeyQueryStringActionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheKeyQueryStringActionParametersResponse)(nil)).Elem()
}

func (o CacheKeyQueryStringActionParametersResponseOutput) ToCacheKeyQueryStringActionParametersResponseOutput() CacheKeyQueryStringActionParametersResponseOutput {
	return o
}

func (o CacheKeyQueryStringActionParametersResponseOutput) ToCacheKeyQueryStringActionParametersResponseOutputWithContext(ctx context.Context) CacheKeyQueryStringActionParametersResponseOutput {
	return o
}

func (o CacheKeyQueryStringActionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v CacheKeyQueryStringActionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o CacheKeyQueryStringActionParametersResponseOutput) QueryParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheKeyQueryStringActionParametersResponse) *string { return v.QueryParameters }).(pulumi.StringPtrOutput)
}

func (o CacheKeyQueryStringActionParametersResponseOutput) QueryStringBehavior() pulumi.StringOutput {
	return o.ApplyT(func(v CacheKeyQueryStringActionParametersResponse) string { return v.QueryStringBehavior }).(pulumi.StringOutput)
}

type CdnCertificateSourceParametersResponse struct {
	CertificateType string `pulumi:"certificateType"`
	OdataType       string `pulumi:"odataType"`
}





type CdnCertificateSourceParametersResponseInput interface {
	pulumi.Input

	ToCdnCertificateSourceParametersResponseOutput() CdnCertificateSourceParametersResponseOutput
	ToCdnCertificateSourceParametersResponseOutputWithContext(context.Context) CdnCertificateSourceParametersResponseOutput
}

type CdnCertificateSourceParametersResponseArgs struct {
	CertificateType pulumi.StringInput `pulumi:"certificateType"`
	OdataType       pulumi.StringInput `pulumi:"odataType"`
}

func (CdnCertificateSourceParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CdnCertificateSourceParametersResponse)(nil)).Elem()
}

func (i CdnCertificateSourceParametersResponseArgs) ToCdnCertificateSourceParametersResponseOutput() CdnCertificateSourceParametersResponseOutput {
	return i.ToCdnCertificateSourceParametersResponseOutputWithContext(context.Background())
}

func (i CdnCertificateSourceParametersResponseArgs) ToCdnCertificateSourceParametersResponseOutputWithContext(ctx context.Context) CdnCertificateSourceParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdnCertificateSourceParametersResponseOutput)
}

type CdnCertificateSourceParametersResponseOutput struct{ *pulumi.OutputState }

func (CdnCertificateSourceParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CdnCertificateSourceParametersResponse)(nil)).Elem()
}

func (o CdnCertificateSourceParametersResponseOutput) ToCdnCertificateSourceParametersResponseOutput() CdnCertificateSourceParametersResponseOutput {
	return o
}

func (o CdnCertificateSourceParametersResponseOutput) ToCdnCertificateSourceParametersResponseOutputWithContext(ctx context.Context) CdnCertificateSourceParametersResponseOutput {
	return o
}

func (o CdnCertificateSourceParametersResponseOutput) CertificateType() pulumi.StringOutput {
	return o.ApplyT(func(v CdnCertificateSourceParametersResponse) string { return v.CertificateType }).(pulumi.StringOutput)
}

func (o CdnCertificateSourceParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v CdnCertificateSourceParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type CdnEndpointResponse struct {
	Id *string `pulumi:"id"`
}





type CdnEndpointResponseInput interface {
	pulumi.Input

	ToCdnEndpointResponseOutput() CdnEndpointResponseOutput
	ToCdnEndpointResponseOutputWithContext(context.Context) CdnEndpointResponseOutput
}

type CdnEndpointResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (CdnEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CdnEndpointResponse)(nil)).Elem()
}

func (i CdnEndpointResponseArgs) ToCdnEndpointResponseOutput() CdnEndpointResponseOutput {
	return i.ToCdnEndpointResponseOutputWithContext(context.Background())
}

func (i CdnEndpointResponseArgs) ToCdnEndpointResponseOutputWithContext(ctx context.Context) CdnEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdnEndpointResponseOutput)
}





type CdnEndpointResponseArrayInput interface {
	pulumi.Input

	ToCdnEndpointResponseArrayOutput() CdnEndpointResponseArrayOutput
	ToCdnEndpointResponseArrayOutputWithContext(context.Context) CdnEndpointResponseArrayOutput
}

type CdnEndpointResponseArray []CdnEndpointResponseInput

func (CdnEndpointResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CdnEndpointResponse)(nil)).Elem()
}

func (i CdnEndpointResponseArray) ToCdnEndpointResponseArrayOutput() CdnEndpointResponseArrayOutput {
	return i.ToCdnEndpointResponseArrayOutputWithContext(context.Background())
}

func (i CdnEndpointResponseArray) ToCdnEndpointResponseArrayOutputWithContext(ctx context.Context) CdnEndpointResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdnEndpointResponseArrayOutput)
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





type CdnManagedHttpsParametersResponseInput interface {
	pulumi.Input

	ToCdnManagedHttpsParametersResponseOutput() CdnManagedHttpsParametersResponseOutput
	ToCdnManagedHttpsParametersResponseOutputWithContext(context.Context) CdnManagedHttpsParametersResponseOutput
}

type CdnManagedHttpsParametersResponseArgs struct {
	CertificateSource           pulumi.StringInput                          `pulumi:"certificateSource"`
	CertificateSourceParameters CdnCertificateSourceParametersResponseInput `pulumi:"certificateSourceParameters"`
	MinimumTlsVersion           pulumi.StringPtrInput                       `pulumi:"minimumTlsVersion"`
	ProtocolType                pulumi.StringInput                          `pulumi:"protocolType"`
}

func (CdnManagedHttpsParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CdnManagedHttpsParametersResponse)(nil)).Elem()
}

func (i CdnManagedHttpsParametersResponseArgs) ToCdnManagedHttpsParametersResponseOutput() CdnManagedHttpsParametersResponseOutput {
	return i.ToCdnManagedHttpsParametersResponseOutputWithContext(context.Background())
}

func (i CdnManagedHttpsParametersResponseArgs) ToCdnManagedHttpsParametersResponseOutputWithContext(ctx context.Context) CdnManagedHttpsParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CdnManagedHttpsParametersResponseOutput)
}

type CdnManagedHttpsParametersResponseOutput struct{ *pulumi.OutputState }

func (CdnManagedHttpsParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CdnManagedHttpsParametersResponse)(nil)).Elem()
}

func (o CdnManagedHttpsParametersResponseOutput) ToCdnManagedHttpsParametersResponseOutput() CdnManagedHttpsParametersResponseOutput {
	return o
}

func (o CdnManagedHttpsParametersResponseOutput) ToCdnManagedHttpsParametersResponseOutputWithContext(ctx context.Context) CdnManagedHttpsParametersResponseOutput {
	return o
}

func (o CdnManagedHttpsParametersResponseOutput) CertificateSource() pulumi.StringOutput {
	return o.ApplyT(func(v CdnManagedHttpsParametersResponse) string { return v.CertificateSource }).(pulumi.StringOutput)
}

func (o CdnManagedHttpsParametersResponseOutput) CertificateSourceParameters() CdnCertificateSourceParametersResponseOutput {
	return o.ApplyT(func(v CdnManagedHttpsParametersResponse) CdnCertificateSourceParametersResponse {
		return v.CertificateSourceParameters
	}).(CdnCertificateSourceParametersResponseOutput)
}

func (o CdnManagedHttpsParametersResponseOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CdnManagedHttpsParametersResponse) *string { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

func (o CdnManagedHttpsParametersResponseOutput) ProtocolType() pulumi.StringOutput {
	return o.ApplyT(func(v CdnManagedHttpsParametersResponse) string { return v.ProtocolType }).(pulumi.StringOutput)
}

type CookiesMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}





type CookiesMatchConditionParametersInput interface {
	pulumi.Input

	ToCookiesMatchConditionParametersOutput() CookiesMatchConditionParametersOutput
	ToCookiesMatchConditionParametersOutputWithContext(context.Context) CookiesMatchConditionParametersOutput
}

type CookiesMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Selector        pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (CookiesMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CookiesMatchConditionParameters)(nil)).Elem()
}

func (i CookiesMatchConditionParametersArgs) ToCookiesMatchConditionParametersOutput() CookiesMatchConditionParametersOutput {
	return i.ToCookiesMatchConditionParametersOutputWithContext(context.Background())
}

func (i CookiesMatchConditionParametersArgs) ToCookiesMatchConditionParametersOutputWithContext(ctx context.Context) CookiesMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CookiesMatchConditionParametersOutput)
}

type CookiesMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (CookiesMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CookiesMatchConditionParameters)(nil)).Elem()
}

func (o CookiesMatchConditionParametersOutput) ToCookiesMatchConditionParametersOutput() CookiesMatchConditionParametersOutput {
	return o
}

func (o CookiesMatchConditionParametersOutput) ToCookiesMatchConditionParametersOutputWithContext(ctx context.Context) CookiesMatchConditionParametersOutput {
	return o
}

func (o CookiesMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CookiesMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o CookiesMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CookiesMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o CookiesMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v CookiesMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o CookiesMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v CookiesMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

func (o CookiesMatchConditionParametersOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CookiesMatchConditionParameters) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o CookiesMatchConditionParametersOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CookiesMatchConditionParameters) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type CookiesMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}





type CookiesMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToCookiesMatchConditionParametersResponseOutput() CookiesMatchConditionParametersResponseOutput
	ToCookiesMatchConditionParametersResponseOutputWithContext(context.Context) CookiesMatchConditionParametersResponseOutput
}

type CookiesMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Selector        pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (CookiesMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CookiesMatchConditionParametersResponse)(nil)).Elem()
}

func (i CookiesMatchConditionParametersResponseArgs) ToCookiesMatchConditionParametersResponseOutput() CookiesMatchConditionParametersResponseOutput {
	return i.ToCookiesMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i CookiesMatchConditionParametersResponseArgs) ToCookiesMatchConditionParametersResponseOutputWithContext(ctx context.Context) CookiesMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CookiesMatchConditionParametersResponseOutput)
}

type CookiesMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (CookiesMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CookiesMatchConditionParametersResponse)(nil)).Elem()
}

func (o CookiesMatchConditionParametersResponseOutput) ToCookiesMatchConditionParametersResponseOutput() CookiesMatchConditionParametersResponseOutput {
	return o
}

func (o CookiesMatchConditionParametersResponseOutput) ToCookiesMatchConditionParametersResponseOutputWithContext(ctx context.Context) CookiesMatchConditionParametersResponseOutput {
	return o
}

func (o CookiesMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CookiesMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o CookiesMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CookiesMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o CookiesMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v CookiesMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o CookiesMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v CookiesMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o CookiesMatchConditionParametersResponseOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CookiesMatchConditionParametersResponse) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o CookiesMatchConditionParametersResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CookiesMatchConditionParametersResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
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





type CustomRuleListResponseInput interface {
	pulumi.Input

	ToCustomRuleListResponseOutput() CustomRuleListResponseOutput
	ToCustomRuleListResponseOutputWithContext(context.Context) CustomRuleListResponseOutput
}

type CustomRuleListResponseArgs struct {
	Rules CustomRuleResponseArrayInput `pulumi:"rules"`
}

func (CustomRuleListResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleListResponse)(nil)).Elem()
}

func (i CustomRuleListResponseArgs) ToCustomRuleListResponseOutput() CustomRuleListResponseOutput {
	return i.ToCustomRuleListResponseOutputWithContext(context.Background())
}

func (i CustomRuleListResponseArgs) ToCustomRuleListResponseOutputWithContext(ctx context.Context) CustomRuleListResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListResponseOutput)
}

func (i CustomRuleListResponseArgs) ToCustomRuleListResponsePtrOutput() CustomRuleListResponsePtrOutput {
	return i.ToCustomRuleListResponsePtrOutputWithContext(context.Background())
}

func (i CustomRuleListResponseArgs) ToCustomRuleListResponsePtrOutputWithContext(ctx context.Context) CustomRuleListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListResponseOutput).ToCustomRuleListResponsePtrOutputWithContext(ctx)
}









type CustomRuleListResponsePtrInput interface {
	pulumi.Input

	ToCustomRuleListResponsePtrOutput() CustomRuleListResponsePtrOutput
	ToCustomRuleListResponsePtrOutputWithContext(context.Context) CustomRuleListResponsePtrOutput
}

type customRuleListResponsePtrType CustomRuleListResponseArgs

func CustomRuleListResponsePtr(v *CustomRuleListResponseArgs) CustomRuleListResponsePtrInput {
	return (*customRuleListResponsePtrType)(v)
}

func (*customRuleListResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRuleListResponse)(nil)).Elem()
}

func (i *customRuleListResponsePtrType) ToCustomRuleListResponsePtrOutput() CustomRuleListResponsePtrOutput {
	return i.ToCustomRuleListResponsePtrOutputWithContext(context.Background())
}

func (i *customRuleListResponsePtrType) ToCustomRuleListResponsePtrOutputWithContext(ctx context.Context) CustomRuleListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListResponsePtrOutput)
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

func (o CustomRuleListResponseOutput) ToCustomRuleListResponsePtrOutput() CustomRuleListResponsePtrOutput {
	return o.ToCustomRuleListResponsePtrOutputWithContext(context.Background())
}

func (o CustomRuleListResponseOutput) ToCustomRuleListResponsePtrOutputWithContext(ctx context.Context) CustomRuleListResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomRuleListResponse) *CustomRuleListResponse {
		return &v
	}).(CustomRuleListResponsePtrOutput)
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





type CustomRuleResponseInput interface {
	pulumi.Input

	ToCustomRuleResponseOutput() CustomRuleResponseOutput
	ToCustomRuleResponseOutputWithContext(context.Context) CustomRuleResponseOutput
}

type CustomRuleResponseArgs struct {
	Action          pulumi.StringInput               `pulumi:"action"`
	EnabledState    pulumi.StringPtrInput            `pulumi:"enabledState"`
	MatchConditions MatchConditionResponseArrayInput `pulumi:"matchConditions"`
	Name            pulumi.StringInput               `pulumi:"name"`
	Priority        pulumi.IntInput                  `pulumi:"priority"`
}

func (CustomRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleResponse)(nil)).Elem()
}

func (i CustomRuleResponseArgs) ToCustomRuleResponseOutput() CustomRuleResponseOutput {
	return i.ToCustomRuleResponseOutputWithContext(context.Background())
}

func (i CustomRuleResponseArgs) ToCustomRuleResponseOutputWithContext(ctx context.Context) CustomRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleResponseOutput)
}





type CustomRuleResponseArrayInput interface {
	pulumi.Input

	ToCustomRuleResponseArrayOutput() CustomRuleResponseArrayOutput
	ToCustomRuleResponseArrayOutputWithContext(context.Context) CustomRuleResponseArrayOutput
}

type CustomRuleResponseArray []CustomRuleResponseInput

func (CustomRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRuleResponse)(nil)).Elem()
}

func (i CustomRuleResponseArray) ToCustomRuleResponseArrayOutput() CustomRuleResponseArrayOutput {
	return i.ToCustomRuleResponseArrayOutputWithContext(context.Background())
}

func (i CustomRuleResponseArray) ToCustomRuleResponseArrayOutputWithContext(ctx context.Context) CustomRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleResponseArrayOutput)
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





type DeepCreatedOriginGroupResponseInput interface {
	pulumi.Input

	ToDeepCreatedOriginGroupResponseOutput() DeepCreatedOriginGroupResponseOutput
	ToDeepCreatedOriginGroupResponseOutputWithContext(context.Context) DeepCreatedOriginGroupResponseOutput
}

type DeepCreatedOriginGroupResponseArgs struct {
	HealthProbeSettings                                   HealthProbeParametersResponsePtrInput                       `pulumi:"healthProbeSettings"`
	Name                                                  pulumi.StringInput                                          `pulumi:"name"`
	Origins                                               ResourceReferenceResponseArrayInput                         `pulumi:"origins"`
	ResponseBasedOriginErrorDetectionSettings             ResponseBasedOriginErrorDetectionParametersResponsePtrInput `pulumi:"responseBasedOriginErrorDetectionSettings"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes pulumi.IntPtrInput                                          `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
}

func (DeepCreatedOriginGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOriginGroupResponse)(nil)).Elem()
}

func (i DeepCreatedOriginGroupResponseArgs) ToDeepCreatedOriginGroupResponseOutput() DeepCreatedOriginGroupResponseOutput {
	return i.ToDeepCreatedOriginGroupResponseOutputWithContext(context.Background())
}

func (i DeepCreatedOriginGroupResponseArgs) ToDeepCreatedOriginGroupResponseOutputWithContext(ctx context.Context) DeepCreatedOriginGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginGroupResponseOutput)
}





type DeepCreatedOriginGroupResponseArrayInput interface {
	pulumi.Input

	ToDeepCreatedOriginGroupResponseArrayOutput() DeepCreatedOriginGroupResponseArrayOutput
	ToDeepCreatedOriginGroupResponseArrayOutputWithContext(context.Context) DeepCreatedOriginGroupResponseArrayOutput
}

type DeepCreatedOriginGroupResponseArray []DeepCreatedOriginGroupResponseInput

func (DeepCreatedOriginGroupResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOriginGroupResponse)(nil)).Elem()
}

func (i DeepCreatedOriginGroupResponseArray) ToDeepCreatedOriginGroupResponseArrayOutput() DeepCreatedOriginGroupResponseArrayOutput {
	return i.ToDeepCreatedOriginGroupResponseArrayOutputWithContext(context.Background())
}

func (i DeepCreatedOriginGroupResponseArray) ToDeepCreatedOriginGroupResponseArrayOutputWithContext(ctx context.Context) DeepCreatedOriginGroupResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginGroupResponseArrayOutput)
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





type DeepCreatedOriginResponseInput interface {
	pulumi.Input

	ToDeepCreatedOriginResponseOutput() DeepCreatedOriginResponseOutput
	ToDeepCreatedOriginResponseOutputWithContext(context.Context) DeepCreatedOriginResponseOutput
}

type DeepCreatedOriginResponseArgs struct {
	Enabled          pulumi.BoolPtrInput   `pulumi:"enabled"`
	HostName         pulumi.StringInput    `pulumi:"hostName"`
	HttpPort         pulumi.IntPtrInput    `pulumi:"httpPort"`
	HttpsPort        pulumi.IntPtrInput    `pulumi:"httpsPort"`
	Name             pulumi.StringInput    `pulumi:"name"`
	OriginHostHeader pulumi.StringPtrInput `pulumi:"originHostHeader"`
	Priority         pulumi.IntPtrInput    `pulumi:"priority"`
	Weight           pulumi.IntPtrInput    `pulumi:"weight"`
}

func (DeepCreatedOriginResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOriginResponse)(nil)).Elem()
}

func (i DeepCreatedOriginResponseArgs) ToDeepCreatedOriginResponseOutput() DeepCreatedOriginResponseOutput {
	return i.ToDeepCreatedOriginResponseOutputWithContext(context.Background())
}

func (i DeepCreatedOriginResponseArgs) ToDeepCreatedOriginResponseOutputWithContext(ctx context.Context) DeepCreatedOriginResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginResponseOutput)
}





type DeepCreatedOriginResponseArrayInput interface {
	pulumi.Input

	ToDeepCreatedOriginResponseArrayOutput() DeepCreatedOriginResponseArrayOutput
	ToDeepCreatedOriginResponseArrayOutputWithContext(context.Context) DeepCreatedOriginResponseArrayOutput
}

type DeepCreatedOriginResponseArray []DeepCreatedOriginResponseInput

func (DeepCreatedOriginResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOriginResponse)(nil)).Elem()
}

func (i DeepCreatedOriginResponseArray) ToDeepCreatedOriginResponseArrayOutput() DeepCreatedOriginResponseArrayOutput {
	return i.ToDeepCreatedOriginResponseArrayOutputWithContext(context.Background())
}

func (i DeepCreatedOriginResponseArray) ToDeepCreatedOriginResponseArrayOutputWithContext(ctx context.Context) DeepCreatedOriginResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginResponseArrayOutput)
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

type DeliveryRuleCacheExpirationActionResponse struct {
	Name       string                                  `pulumi:"name"`
	Parameters CacheExpirationActionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleCacheExpirationActionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleCacheExpirationActionResponseOutput() DeliveryRuleCacheExpirationActionResponseOutput
	ToDeliveryRuleCacheExpirationActionResponseOutputWithContext(context.Context) DeliveryRuleCacheExpirationActionResponseOutput
}

type DeliveryRuleCacheExpirationActionResponseArgs struct {
	Name       pulumi.StringInput                           `pulumi:"name"`
	Parameters CacheExpirationActionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleCacheExpirationActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCacheExpirationActionResponse)(nil)).Elem()
}

func (i DeliveryRuleCacheExpirationActionResponseArgs) ToDeliveryRuleCacheExpirationActionResponseOutput() DeliveryRuleCacheExpirationActionResponseOutput {
	return i.ToDeliveryRuleCacheExpirationActionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleCacheExpirationActionResponseArgs) ToDeliveryRuleCacheExpirationActionResponseOutputWithContext(ctx context.Context) DeliveryRuleCacheExpirationActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleCacheExpirationActionResponseOutput)
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

type DeliveryRuleCacheKeyQueryStringAction struct {
	Name       string                              `pulumi:"name"`
	Parameters CacheKeyQueryStringActionParameters `pulumi:"parameters"`
}





type DeliveryRuleCacheKeyQueryStringActionInput interface {
	pulumi.Input

	ToDeliveryRuleCacheKeyQueryStringActionOutput() DeliveryRuleCacheKeyQueryStringActionOutput
	ToDeliveryRuleCacheKeyQueryStringActionOutputWithContext(context.Context) DeliveryRuleCacheKeyQueryStringActionOutput
}

type DeliveryRuleCacheKeyQueryStringActionArgs struct {
	Name       pulumi.StringInput                       `pulumi:"name"`
	Parameters CacheKeyQueryStringActionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleCacheKeyQueryStringActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCacheKeyQueryStringAction)(nil)).Elem()
}

func (i DeliveryRuleCacheKeyQueryStringActionArgs) ToDeliveryRuleCacheKeyQueryStringActionOutput() DeliveryRuleCacheKeyQueryStringActionOutput {
	return i.ToDeliveryRuleCacheKeyQueryStringActionOutputWithContext(context.Background())
}

func (i DeliveryRuleCacheKeyQueryStringActionArgs) ToDeliveryRuleCacheKeyQueryStringActionOutputWithContext(ctx context.Context) DeliveryRuleCacheKeyQueryStringActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleCacheKeyQueryStringActionOutput)
}

type DeliveryRuleCacheKeyQueryStringActionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleCacheKeyQueryStringActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCacheKeyQueryStringAction)(nil)).Elem()
}

func (o DeliveryRuleCacheKeyQueryStringActionOutput) ToDeliveryRuleCacheKeyQueryStringActionOutput() DeliveryRuleCacheKeyQueryStringActionOutput {
	return o
}

func (o DeliveryRuleCacheKeyQueryStringActionOutput) ToDeliveryRuleCacheKeyQueryStringActionOutputWithContext(ctx context.Context) DeliveryRuleCacheKeyQueryStringActionOutput {
	return o
}

func (o DeliveryRuleCacheKeyQueryStringActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleCacheKeyQueryStringAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleCacheKeyQueryStringActionOutput) Parameters() CacheKeyQueryStringActionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleCacheKeyQueryStringAction) CacheKeyQueryStringActionParameters { return v.Parameters }).(CacheKeyQueryStringActionParametersOutput)
}

type DeliveryRuleCacheKeyQueryStringActionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters CacheKeyQueryStringActionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleCacheKeyQueryStringActionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleCacheKeyQueryStringActionResponseOutput() DeliveryRuleCacheKeyQueryStringActionResponseOutput
	ToDeliveryRuleCacheKeyQueryStringActionResponseOutputWithContext(context.Context) DeliveryRuleCacheKeyQueryStringActionResponseOutput
}

type DeliveryRuleCacheKeyQueryStringActionResponseArgs struct {
	Name       pulumi.StringInput                               `pulumi:"name"`
	Parameters CacheKeyQueryStringActionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleCacheKeyQueryStringActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCacheKeyQueryStringActionResponse)(nil)).Elem()
}

func (i DeliveryRuleCacheKeyQueryStringActionResponseArgs) ToDeliveryRuleCacheKeyQueryStringActionResponseOutput() DeliveryRuleCacheKeyQueryStringActionResponseOutput {
	return i.ToDeliveryRuleCacheKeyQueryStringActionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleCacheKeyQueryStringActionResponseArgs) ToDeliveryRuleCacheKeyQueryStringActionResponseOutputWithContext(ctx context.Context) DeliveryRuleCacheKeyQueryStringActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleCacheKeyQueryStringActionResponseOutput)
}

type DeliveryRuleCacheKeyQueryStringActionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleCacheKeyQueryStringActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCacheKeyQueryStringActionResponse)(nil)).Elem()
}

func (o DeliveryRuleCacheKeyQueryStringActionResponseOutput) ToDeliveryRuleCacheKeyQueryStringActionResponseOutput() DeliveryRuleCacheKeyQueryStringActionResponseOutput {
	return o
}

func (o DeliveryRuleCacheKeyQueryStringActionResponseOutput) ToDeliveryRuleCacheKeyQueryStringActionResponseOutputWithContext(ctx context.Context) DeliveryRuleCacheKeyQueryStringActionResponseOutput {
	return o
}

func (o DeliveryRuleCacheKeyQueryStringActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleCacheKeyQueryStringActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleCacheKeyQueryStringActionResponseOutput) Parameters() CacheKeyQueryStringActionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleCacheKeyQueryStringActionResponse) CacheKeyQueryStringActionParametersResponse {
		return v.Parameters
	}).(CacheKeyQueryStringActionParametersResponseOutput)
}

type DeliveryRuleCookiesCondition struct {
	Name       string                          `pulumi:"name"`
	Parameters CookiesMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleCookiesConditionInput interface {
	pulumi.Input

	ToDeliveryRuleCookiesConditionOutput() DeliveryRuleCookiesConditionOutput
	ToDeliveryRuleCookiesConditionOutputWithContext(context.Context) DeliveryRuleCookiesConditionOutput
}

type DeliveryRuleCookiesConditionArgs struct {
	Name       pulumi.StringInput                   `pulumi:"name"`
	Parameters CookiesMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleCookiesConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCookiesCondition)(nil)).Elem()
}

func (i DeliveryRuleCookiesConditionArgs) ToDeliveryRuleCookiesConditionOutput() DeliveryRuleCookiesConditionOutput {
	return i.ToDeliveryRuleCookiesConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleCookiesConditionArgs) ToDeliveryRuleCookiesConditionOutputWithContext(ctx context.Context) DeliveryRuleCookiesConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleCookiesConditionOutput)
}

type DeliveryRuleCookiesConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleCookiesConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCookiesCondition)(nil)).Elem()
}

func (o DeliveryRuleCookiesConditionOutput) ToDeliveryRuleCookiesConditionOutput() DeliveryRuleCookiesConditionOutput {
	return o
}

func (o DeliveryRuleCookiesConditionOutput) ToDeliveryRuleCookiesConditionOutputWithContext(ctx context.Context) DeliveryRuleCookiesConditionOutput {
	return o
}

func (o DeliveryRuleCookiesConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleCookiesCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleCookiesConditionOutput) Parameters() CookiesMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleCookiesCondition) CookiesMatchConditionParameters { return v.Parameters }).(CookiesMatchConditionParametersOutput)
}

type DeliveryRuleCookiesConditionResponse struct {
	Name       string                                  `pulumi:"name"`
	Parameters CookiesMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleCookiesConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleCookiesConditionResponseOutput() DeliveryRuleCookiesConditionResponseOutput
	ToDeliveryRuleCookiesConditionResponseOutputWithContext(context.Context) DeliveryRuleCookiesConditionResponseOutput
}

type DeliveryRuleCookiesConditionResponseArgs struct {
	Name       pulumi.StringInput                           `pulumi:"name"`
	Parameters CookiesMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleCookiesConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCookiesConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleCookiesConditionResponseArgs) ToDeliveryRuleCookiesConditionResponseOutput() DeliveryRuleCookiesConditionResponseOutput {
	return i.ToDeliveryRuleCookiesConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleCookiesConditionResponseArgs) ToDeliveryRuleCookiesConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleCookiesConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleCookiesConditionResponseOutput)
}

type DeliveryRuleCookiesConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleCookiesConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleCookiesConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleCookiesConditionResponseOutput) ToDeliveryRuleCookiesConditionResponseOutput() DeliveryRuleCookiesConditionResponseOutput {
	return o
}

func (o DeliveryRuleCookiesConditionResponseOutput) ToDeliveryRuleCookiesConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleCookiesConditionResponseOutput {
	return o
}

func (o DeliveryRuleCookiesConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleCookiesConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleCookiesConditionResponseOutput) Parameters() CookiesMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleCookiesConditionResponse) CookiesMatchConditionParametersResponse {
		return v.Parameters
	}).(CookiesMatchConditionParametersResponseOutput)
}

type DeliveryRuleHttpVersionCondition struct {
	Name       string                              `pulumi:"name"`
	Parameters HttpVersionMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleHttpVersionConditionInput interface {
	pulumi.Input

	ToDeliveryRuleHttpVersionConditionOutput() DeliveryRuleHttpVersionConditionOutput
	ToDeliveryRuleHttpVersionConditionOutputWithContext(context.Context) DeliveryRuleHttpVersionConditionOutput
}

type DeliveryRuleHttpVersionConditionArgs struct {
	Name       pulumi.StringInput                       `pulumi:"name"`
	Parameters HttpVersionMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleHttpVersionConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleHttpVersionCondition)(nil)).Elem()
}

func (i DeliveryRuleHttpVersionConditionArgs) ToDeliveryRuleHttpVersionConditionOutput() DeliveryRuleHttpVersionConditionOutput {
	return i.ToDeliveryRuleHttpVersionConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleHttpVersionConditionArgs) ToDeliveryRuleHttpVersionConditionOutputWithContext(ctx context.Context) DeliveryRuleHttpVersionConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleHttpVersionConditionOutput)
}

type DeliveryRuleHttpVersionConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleHttpVersionConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleHttpVersionCondition)(nil)).Elem()
}

func (o DeliveryRuleHttpVersionConditionOutput) ToDeliveryRuleHttpVersionConditionOutput() DeliveryRuleHttpVersionConditionOutput {
	return o
}

func (o DeliveryRuleHttpVersionConditionOutput) ToDeliveryRuleHttpVersionConditionOutputWithContext(ctx context.Context) DeliveryRuleHttpVersionConditionOutput {
	return o
}

func (o DeliveryRuleHttpVersionConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleHttpVersionCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleHttpVersionConditionOutput) Parameters() HttpVersionMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleHttpVersionCondition) HttpVersionMatchConditionParameters { return v.Parameters }).(HttpVersionMatchConditionParametersOutput)
}

type DeliveryRuleHttpVersionConditionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters HttpVersionMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleHttpVersionConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleHttpVersionConditionResponseOutput() DeliveryRuleHttpVersionConditionResponseOutput
	ToDeliveryRuleHttpVersionConditionResponseOutputWithContext(context.Context) DeliveryRuleHttpVersionConditionResponseOutput
}

type DeliveryRuleHttpVersionConditionResponseArgs struct {
	Name       pulumi.StringInput                               `pulumi:"name"`
	Parameters HttpVersionMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleHttpVersionConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleHttpVersionConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleHttpVersionConditionResponseArgs) ToDeliveryRuleHttpVersionConditionResponseOutput() DeliveryRuleHttpVersionConditionResponseOutput {
	return i.ToDeliveryRuleHttpVersionConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleHttpVersionConditionResponseArgs) ToDeliveryRuleHttpVersionConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleHttpVersionConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleHttpVersionConditionResponseOutput)
}

type DeliveryRuleHttpVersionConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleHttpVersionConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleHttpVersionConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleHttpVersionConditionResponseOutput) ToDeliveryRuleHttpVersionConditionResponseOutput() DeliveryRuleHttpVersionConditionResponseOutput {
	return o
}

func (o DeliveryRuleHttpVersionConditionResponseOutput) ToDeliveryRuleHttpVersionConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleHttpVersionConditionResponseOutput {
	return o
}

func (o DeliveryRuleHttpVersionConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleHttpVersionConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleHttpVersionConditionResponseOutput) Parameters() HttpVersionMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleHttpVersionConditionResponse) HttpVersionMatchConditionParametersResponse {
		return v.Parameters
	}).(HttpVersionMatchConditionParametersResponseOutput)
}

type DeliveryRuleIsDeviceCondition struct {
	Name       string                           `pulumi:"name"`
	Parameters IsDeviceMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleIsDeviceConditionInput interface {
	pulumi.Input

	ToDeliveryRuleIsDeviceConditionOutput() DeliveryRuleIsDeviceConditionOutput
	ToDeliveryRuleIsDeviceConditionOutputWithContext(context.Context) DeliveryRuleIsDeviceConditionOutput
}

type DeliveryRuleIsDeviceConditionArgs struct {
	Name       pulumi.StringInput                    `pulumi:"name"`
	Parameters IsDeviceMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleIsDeviceConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleIsDeviceCondition)(nil)).Elem()
}

func (i DeliveryRuleIsDeviceConditionArgs) ToDeliveryRuleIsDeviceConditionOutput() DeliveryRuleIsDeviceConditionOutput {
	return i.ToDeliveryRuleIsDeviceConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleIsDeviceConditionArgs) ToDeliveryRuleIsDeviceConditionOutputWithContext(ctx context.Context) DeliveryRuleIsDeviceConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleIsDeviceConditionOutput)
}

type DeliveryRuleIsDeviceConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleIsDeviceConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleIsDeviceCondition)(nil)).Elem()
}

func (o DeliveryRuleIsDeviceConditionOutput) ToDeliveryRuleIsDeviceConditionOutput() DeliveryRuleIsDeviceConditionOutput {
	return o
}

func (o DeliveryRuleIsDeviceConditionOutput) ToDeliveryRuleIsDeviceConditionOutputWithContext(ctx context.Context) DeliveryRuleIsDeviceConditionOutput {
	return o
}

func (o DeliveryRuleIsDeviceConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleIsDeviceCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleIsDeviceConditionOutput) Parameters() IsDeviceMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleIsDeviceCondition) IsDeviceMatchConditionParameters { return v.Parameters }).(IsDeviceMatchConditionParametersOutput)
}

type DeliveryRuleIsDeviceConditionResponse struct {
	Name       string                                   `pulumi:"name"`
	Parameters IsDeviceMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleIsDeviceConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleIsDeviceConditionResponseOutput() DeliveryRuleIsDeviceConditionResponseOutput
	ToDeliveryRuleIsDeviceConditionResponseOutputWithContext(context.Context) DeliveryRuleIsDeviceConditionResponseOutput
}

type DeliveryRuleIsDeviceConditionResponseArgs struct {
	Name       pulumi.StringInput                            `pulumi:"name"`
	Parameters IsDeviceMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleIsDeviceConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleIsDeviceConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleIsDeviceConditionResponseArgs) ToDeliveryRuleIsDeviceConditionResponseOutput() DeliveryRuleIsDeviceConditionResponseOutput {
	return i.ToDeliveryRuleIsDeviceConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleIsDeviceConditionResponseArgs) ToDeliveryRuleIsDeviceConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleIsDeviceConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleIsDeviceConditionResponseOutput)
}

type DeliveryRuleIsDeviceConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleIsDeviceConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleIsDeviceConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleIsDeviceConditionResponseOutput) ToDeliveryRuleIsDeviceConditionResponseOutput() DeliveryRuleIsDeviceConditionResponseOutput {
	return o
}

func (o DeliveryRuleIsDeviceConditionResponseOutput) ToDeliveryRuleIsDeviceConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleIsDeviceConditionResponseOutput {
	return o
}

func (o DeliveryRuleIsDeviceConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleIsDeviceConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleIsDeviceConditionResponseOutput) Parameters() IsDeviceMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleIsDeviceConditionResponse) IsDeviceMatchConditionParametersResponse {
		return v.Parameters
	}).(IsDeviceMatchConditionParametersResponseOutput)
}

type DeliveryRulePostArgsCondition struct {
	Name       string                           `pulumi:"name"`
	Parameters PostArgsMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRulePostArgsConditionInput interface {
	pulumi.Input

	ToDeliveryRulePostArgsConditionOutput() DeliveryRulePostArgsConditionOutput
	ToDeliveryRulePostArgsConditionOutputWithContext(context.Context) DeliveryRulePostArgsConditionOutput
}

type DeliveryRulePostArgsConditionArgs struct {
	Name       pulumi.StringInput                    `pulumi:"name"`
	Parameters PostArgsMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRulePostArgsConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRulePostArgsCondition)(nil)).Elem()
}

func (i DeliveryRulePostArgsConditionArgs) ToDeliveryRulePostArgsConditionOutput() DeliveryRulePostArgsConditionOutput {
	return i.ToDeliveryRulePostArgsConditionOutputWithContext(context.Background())
}

func (i DeliveryRulePostArgsConditionArgs) ToDeliveryRulePostArgsConditionOutputWithContext(ctx context.Context) DeliveryRulePostArgsConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRulePostArgsConditionOutput)
}

type DeliveryRulePostArgsConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRulePostArgsConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRulePostArgsCondition)(nil)).Elem()
}

func (o DeliveryRulePostArgsConditionOutput) ToDeliveryRulePostArgsConditionOutput() DeliveryRulePostArgsConditionOutput {
	return o
}

func (o DeliveryRulePostArgsConditionOutput) ToDeliveryRulePostArgsConditionOutputWithContext(ctx context.Context) DeliveryRulePostArgsConditionOutput {
	return o
}

func (o DeliveryRulePostArgsConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRulePostArgsCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRulePostArgsConditionOutput) Parameters() PostArgsMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRulePostArgsCondition) PostArgsMatchConditionParameters { return v.Parameters }).(PostArgsMatchConditionParametersOutput)
}

type DeliveryRulePostArgsConditionResponse struct {
	Name       string                                   `pulumi:"name"`
	Parameters PostArgsMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRulePostArgsConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRulePostArgsConditionResponseOutput() DeliveryRulePostArgsConditionResponseOutput
	ToDeliveryRulePostArgsConditionResponseOutputWithContext(context.Context) DeliveryRulePostArgsConditionResponseOutput
}

type DeliveryRulePostArgsConditionResponseArgs struct {
	Name       pulumi.StringInput                            `pulumi:"name"`
	Parameters PostArgsMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRulePostArgsConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRulePostArgsConditionResponse)(nil)).Elem()
}

func (i DeliveryRulePostArgsConditionResponseArgs) ToDeliveryRulePostArgsConditionResponseOutput() DeliveryRulePostArgsConditionResponseOutput {
	return i.ToDeliveryRulePostArgsConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRulePostArgsConditionResponseArgs) ToDeliveryRulePostArgsConditionResponseOutputWithContext(ctx context.Context) DeliveryRulePostArgsConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRulePostArgsConditionResponseOutput)
}

type DeliveryRulePostArgsConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRulePostArgsConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRulePostArgsConditionResponse)(nil)).Elem()
}

func (o DeliveryRulePostArgsConditionResponseOutput) ToDeliveryRulePostArgsConditionResponseOutput() DeliveryRulePostArgsConditionResponseOutput {
	return o
}

func (o DeliveryRulePostArgsConditionResponseOutput) ToDeliveryRulePostArgsConditionResponseOutputWithContext(ctx context.Context) DeliveryRulePostArgsConditionResponseOutput {
	return o
}

func (o DeliveryRulePostArgsConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRulePostArgsConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRulePostArgsConditionResponseOutput) Parameters() PostArgsMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRulePostArgsConditionResponse) PostArgsMatchConditionParametersResponse {
		return v.Parameters
	}).(PostArgsMatchConditionParametersResponseOutput)
}

type DeliveryRuleQueryStringCondition struct {
	Name       string                              `pulumi:"name"`
	Parameters QueryStringMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleQueryStringConditionInput interface {
	pulumi.Input

	ToDeliveryRuleQueryStringConditionOutput() DeliveryRuleQueryStringConditionOutput
	ToDeliveryRuleQueryStringConditionOutputWithContext(context.Context) DeliveryRuleQueryStringConditionOutput
}

type DeliveryRuleQueryStringConditionArgs struct {
	Name       pulumi.StringInput                       `pulumi:"name"`
	Parameters QueryStringMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleQueryStringConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleQueryStringCondition)(nil)).Elem()
}

func (i DeliveryRuleQueryStringConditionArgs) ToDeliveryRuleQueryStringConditionOutput() DeliveryRuleQueryStringConditionOutput {
	return i.ToDeliveryRuleQueryStringConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleQueryStringConditionArgs) ToDeliveryRuleQueryStringConditionOutputWithContext(ctx context.Context) DeliveryRuleQueryStringConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleQueryStringConditionOutput)
}

type DeliveryRuleQueryStringConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleQueryStringConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleQueryStringCondition)(nil)).Elem()
}

func (o DeliveryRuleQueryStringConditionOutput) ToDeliveryRuleQueryStringConditionOutput() DeliveryRuleQueryStringConditionOutput {
	return o
}

func (o DeliveryRuleQueryStringConditionOutput) ToDeliveryRuleQueryStringConditionOutputWithContext(ctx context.Context) DeliveryRuleQueryStringConditionOutput {
	return o
}

func (o DeliveryRuleQueryStringConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleQueryStringCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleQueryStringConditionOutput) Parameters() QueryStringMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleQueryStringCondition) QueryStringMatchConditionParameters { return v.Parameters }).(QueryStringMatchConditionParametersOutput)
}

type DeliveryRuleQueryStringConditionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters QueryStringMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleQueryStringConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleQueryStringConditionResponseOutput() DeliveryRuleQueryStringConditionResponseOutput
	ToDeliveryRuleQueryStringConditionResponseOutputWithContext(context.Context) DeliveryRuleQueryStringConditionResponseOutput
}

type DeliveryRuleQueryStringConditionResponseArgs struct {
	Name       pulumi.StringInput                               `pulumi:"name"`
	Parameters QueryStringMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleQueryStringConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleQueryStringConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleQueryStringConditionResponseArgs) ToDeliveryRuleQueryStringConditionResponseOutput() DeliveryRuleQueryStringConditionResponseOutput {
	return i.ToDeliveryRuleQueryStringConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleQueryStringConditionResponseArgs) ToDeliveryRuleQueryStringConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleQueryStringConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleQueryStringConditionResponseOutput)
}

type DeliveryRuleQueryStringConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleQueryStringConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleQueryStringConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleQueryStringConditionResponseOutput) ToDeliveryRuleQueryStringConditionResponseOutput() DeliveryRuleQueryStringConditionResponseOutput {
	return o
}

func (o DeliveryRuleQueryStringConditionResponseOutput) ToDeliveryRuleQueryStringConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleQueryStringConditionResponseOutput {
	return o
}

func (o DeliveryRuleQueryStringConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleQueryStringConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleQueryStringConditionResponseOutput) Parameters() QueryStringMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleQueryStringConditionResponse) QueryStringMatchConditionParametersResponse {
		return v.Parameters
	}).(QueryStringMatchConditionParametersResponseOutput)
}

type DeliveryRuleRemoteAddressCondition struct {
	Name       string                                `pulumi:"name"`
	Parameters RemoteAddressMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleRemoteAddressConditionInput interface {
	pulumi.Input

	ToDeliveryRuleRemoteAddressConditionOutput() DeliveryRuleRemoteAddressConditionOutput
	ToDeliveryRuleRemoteAddressConditionOutputWithContext(context.Context) DeliveryRuleRemoteAddressConditionOutput
}

type DeliveryRuleRemoteAddressConditionArgs struct {
	Name       pulumi.StringInput                         `pulumi:"name"`
	Parameters RemoteAddressMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleRemoteAddressConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRemoteAddressCondition)(nil)).Elem()
}

func (i DeliveryRuleRemoteAddressConditionArgs) ToDeliveryRuleRemoteAddressConditionOutput() DeliveryRuleRemoteAddressConditionOutput {
	return i.ToDeliveryRuleRemoteAddressConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleRemoteAddressConditionArgs) ToDeliveryRuleRemoteAddressConditionOutputWithContext(ctx context.Context) DeliveryRuleRemoteAddressConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRemoteAddressConditionOutput)
}

type DeliveryRuleRemoteAddressConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRemoteAddressConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRemoteAddressCondition)(nil)).Elem()
}

func (o DeliveryRuleRemoteAddressConditionOutput) ToDeliveryRuleRemoteAddressConditionOutput() DeliveryRuleRemoteAddressConditionOutput {
	return o
}

func (o DeliveryRuleRemoteAddressConditionOutput) ToDeliveryRuleRemoteAddressConditionOutputWithContext(ctx context.Context) DeliveryRuleRemoteAddressConditionOutput {
	return o
}

func (o DeliveryRuleRemoteAddressConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRemoteAddressCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRemoteAddressConditionOutput) Parameters() RemoteAddressMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleRemoteAddressCondition) RemoteAddressMatchConditionParameters { return v.Parameters }).(RemoteAddressMatchConditionParametersOutput)
}

type DeliveryRuleRemoteAddressConditionResponse struct {
	Name       string                                        `pulumi:"name"`
	Parameters RemoteAddressMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleRemoteAddressConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleRemoteAddressConditionResponseOutput() DeliveryRuleRemoteAddressConditionResponseOutput
	ToDeliveryRuleRemoteAddressConditionResponseOutputWithContext(context.Context) DeliveryRuleRemoteAddressConditionResponseOutput
}

type DeliveryRuleRemoteAddressConditionResponseArgs struct {
	Name       pulumi.StringInput                                 `pulumi:"name"`
	Parameters RemoteAddressMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleRemoteAddressConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRemoteAddressConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleRemoteAddressConditionResponseArgs) ToDeliveryRuleRemoteAddressConditionResponseOutput() DeliveryRuleRemoteAddressConditionResponseOutput {
	return i.ToDeliveryRuleRemoteAddressConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleRemoteAddressConditionResponseArgs) ToDeliveryRuleRemoteAddressConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRemoteAddressConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRemoteAddressConditionResponseOutput)
}

type DeliveryRuleRemoteAddressConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRemoteAddressConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRemoteAddressConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleRemoteAddressConditionResponseOutput) ToDeliveryRuleRemoteAddressConditionResponseOutput() DeliveryRuleRemoteAddressConditionResponseOutput {
	return o
}

func (o DeliveryRuleRemoteAddressConditionResponseOutput) ToDeliveryRuleRemoteAddressConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRemoteAddressConditionResponseOutput {
	return o
}

func (o DeliveryRuleRemoteAddressConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRemoteAddressConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRemoteAddressConditionResponseOutput) Parameters() RemoteAddressMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleRemoteAddressConditionResponse) RemoteAddressMatchConditionParametersResponse {
		return v.Parameters
	}).(RemoteAddressMatchConditionParametersResponseOutput)
}

type DeliveryRuleRequestBodyCondition struct {
	Name       string                              `pulumi:"name"`
	Parameters RequestBodyMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleRequestBodyConditionInput interface {
	pulumi.Input

	ToDeliveryRuleRequestBodyConditionOutput() DeliveryRuleRequestBodyConditionOutput
	ToDeliveryRuleRequestBodyConditionOutputWithContext(context.Context) DeliveryRuleRequestBodyConditionOutput
}

type DeliveryRuleRequestBodyConditionArgs struct {
	Name       pulumi.StringInput                       `pulumi:"name"`
	Parameters RequestBodyMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestBodyConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestBodyCondition)(nil)).Elem()
}

func (i DeliveryRuleRequestBodyConditionArgs) ToDeliveryRuleRequestBodyConditionOutput() DeliveryRuleRequestBodyConditionOutput {
	return i.ToDeliveryRuleRequestBodyConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestBodyConditionArgs) ToDeliveryRuleRequestBodyConditionOutputWithContext(ctx context.Context) DeliveryRuleRequestBodyConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestBodyConditionOutput)
}

type DeliveryRuleRequestBodyConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestBodyConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestBodyCondition)(nil)).Elem()
}

func (o DeliveryRuleRequestBodyConditionOutput) ToDeliveryRuleRequestBodyConditionOutput() DeliveryRuleRequestBodyConditionOutput {
	return o
}

func (o DeliveryRuleRequestBodyConditionOutput) ToDeliveryRuleRequestBodyConditionOutputWithContext(ctx context.Context) DeliveryRuleRequestBodyConditionOutput {
	return o
}

func (o DeliveryRuleRequestBodyConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestBodyCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestBodyConditionOutput) Parameters() RequestBodyMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleRequestBodyCondition) RequestBodyMatchConditionParameters { return v.Parameters }).(RequestBodyMatchConditionParametersOutput)
}

type DeliveryRuleRequestBodyConditionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters RequestBodyMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleRequestBodyConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleRequestBodyConditionResponseOutput() DeliveryRuleRequestBodyConditionResponseOutput
	ToDeliveryRuleRequestBodyConditionResponseOutputWithContext(context.Context) DeliveryRuleRequestBodyConditionResponseOutput
}

type DeliveryRuleRequestBodyConditionResponseArgs struct {
	Name       pulumi.StringInput                               `pulumi:"name"`
	Parameters RequestBodyMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestBodyConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestBodyConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleRequestBodyConditionResponseArgs) ToDeliveryRuleRequestBodyConditionResponseOutput() DeliveryRuleRequestBodyConditionResponseOutput {
	return i.ToDeliveryRuleRequestBodyConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestBodyConditionResponseArgs) ToDeliveryRuleRequestBodyConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestBodyConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestBodyConditionResponseOutput)
}

type DeliveryRuleRequestBodyConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestBodyConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestBodyConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleRequestBodyConditionResponseOutput) ToDeliveryRuleRequestBodyConditionResponseOutput() DeliveryRuleRequestBodyConditionResponseOutput {
	return o
}

func (o DeliveryRuleRequestBodyConditionResponseOutput) ToDeliveryRuleRequestBodyConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestBodyConditionResponseOutput {
	return o
}

func (o DeliveryRuleRequestBodyConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestBodyConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestBodyConditionResponseOutput) Parameters() RequestBodyMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleRequestBodyConditionResponse) RequestBodyMatchConditionParametersResponse {
		return v.Parameters
	}).(RequestBodyMatchConditionParametersResponseOutput)
}

type DeliveryRuleRequestHeaderAction struct {
	Name       string                 `pulumi:"name"`
	Parameters HeaderActionParameters `pulumi:"parameters"`
}





type DeliveryRuleRequestHeaderActionInput interface {
	pulumi.Input

	ToDeliveryRuleRequestHeaderActionOutput() DeliveryRuleRequestHeaderActionOutput
	ToDeliveryRuleRequestHeaderActionOutputWithContext(context.Context) DeliveryRuleRequestHeaderActionOutput
}

type DeliveryRuleRequestHeaderActionArgs struct {
	Name       pulumi.StringInput          `pulumi:"name"`
	Parameters HeaderActionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestHeaderActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestHeaderAction)(nil)).Elem()
}

func (i DeliveryRuleRequestHeaderActionArgs) ToDeliveryRuleRequestHeaderActionOutput() DeliveryRuleRequestHeaderActionOutput {
	return i.ToDeliveryRuleRequestHeaderActionOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestHeaderActionArgs) ToDeliveryRuleRequestHeaderActionOutputWithContext(ctx context.Context) DeliveryRuleRequestHeaderActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestHeaderActionOutput)
}

type DeliveryRuleRequestHeaderActionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestHeaderActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestHeaderAction)(nil)).Elem()
}

func (o DeliveryRuleRequestHeaderActionOutput) ToDeliveryRuleRequestHeaderActionOutput() DeliveryRuleRequestHeaderActionOutput {
	return o
}

func (o DeliveryRuleRequestHeaderActionOutput) ToDeliveryRuleRequestHeaderActionOutputWithContext(ctx context.Context) DeliveryRuleRequestHeaderActionOutput {
	return o
}

func (o DeliveryRuleRequestHeaderActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestHeaderAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestHeaderActionOutput) Parameters() HeaderActionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleRequestHeaderAction) HeaderActionParameters { return v.Parameters }).(HeaderActionParametersOutput)
}

type DeliveryRuleRequestHeaderActionResponse struct {
	Name       string                         `pulumi:"name"`
	Parameters HeaderActionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleRequestHeaderActionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleRequestHeaderActionResponseOutput() DeliveryRuleRequestHeaderActionResponseOutput
	ToDeliveryRuleRequestHeaderActionResponseOutputWithContext(context.Context) DeliveryRuleRequestHeaderActionResponseOutput
}

type DeliveryRuleRequestHeaderActionResponseArgs struct {
	Name       pulumi.StringInput                  `pulumi:"name"`
	Parameters HeaderActionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestHeaderActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestHeaderActionResponse)(nil)).Elem()
}

func (i DeliveryRuleRequestHeaderActionResponseArgs) ToDeliveryRuleRequestHeaderActionResponseOutput() DeliveryRuleRequestHeaderActionResponseOutput {
	return i.ToDeliveryRuleRequestHeaderActionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestHeaderActionResponseArgs) ToDeliveryRuleRequestHeaderActionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestHeaderActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestHeaderActionResponseOutput)
}

type DeliveryRuleRequestHeaderActionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestHeaderActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestHeaderActionResponse)(nil)).Elem()
}

func (o DeliveryRuleRequestHeaderActionResponseOutput) ToDeliveryRuleRequestHeaderActionResponseOutput() DeliveryRuleRequestHeaderActionResponseOutput {
	return o
}

func (o DeliveryRuleRequestHeaderActionResponseOutput) ToDeliveryRuleRequestHeaderActionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestHeaderActionResponseOutput {
	return o
}

func (o DeliveryRuleRequestHeaderActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestHeaderActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestHeaderActionResponseOutput) Parameters() HeaderActionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleRequestHeaderActionResponse) HeaderActionParametersResponse { return v.Parameters }).(HeaderActionParametersResponseOutput)
}

type DeliveryRuleRequestHeaderCondition struct {
	Name       string                                `pulumi:"name"`
	Parameters RequestHeaderMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleRequestHeaderConditionInput interface {
	pulumi.Input

	ToDeliveryRuleRequestHeaderConditionOutput() DeliveryRuleRequestHeaderConditionOutput
	ToDeliveryRuleRequestHeaderConditionOutputWithContext(context.Context) DeliveryRuleRequestHeaderConditionOutput
}

type DeliveryRuleRequestHeaderConditionArgs struct {
	Name       pulumi.StringInput                         `pulumi:"name"`
	Parameters RequestHeaderMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestHeaderConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestHeaderCondition)(nil)).Elem()
}

func (i DeliveryRuleRequestHeaderConditionArgs) ToDeliveryRuleRequestHeaderConditionOutput() DeliveryRuleRequestHeaderConditionOutput {
	return i.ToDeliveryRuleRequestHeaderConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestHeaderConditionArgs) ToDeliveryRuleRequestHeaderConditionOutputWithContext(ctx context.Context) DeliveryRuleRequestHeaderConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestHeaderConditionOutput)
}

type DeliveryRuleRequestHeaderConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestHeaderConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestHeaderCondition)(nil)).Elem()
}

func (o DeliveryRuleRequestHeaderConditionOutput) ToDeliveryRuleRequestHeaderConditionOutput() DeliveryRuleRequestHeaderConditionOutput {
	return o
}

func (o DeliveryRuleRequestHeaderConditionOutput) ToDeliveryRuleRequestHeaderConditionOutputWithContext(ctx context.Context) DeliveryRuleRequestHeaderConditionOutput {
	return o
}

func (o DeliveryRuleRequestHeaderConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestHeaderCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestHeaderConditionOutput) Parameters() RequestHeaderMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleRequestHeaderCondition) RequestHeaderMatchConditionParameters { return v.Parameters }).(RequestHeaderMatchConditionParametersOutput)
}

type DeliveryRuleRequestHeaderConditionResponse struct {
	Name       string                                        `pulumi:"name"`
	Parameters RequestHeaderMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleRequestHeaderConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleRequestHeaderConditionResponseOutput() DeliveryRuleRequestHeaderConditionResponseOutput
	ToDeliveryRuleRequestHeaderConditionResponseOutputWithContext(context.Context) DeliveryRuleRequestHeaderConditionResponseOutput
}

type DeliveryRuleRequestHeaderConditionResponseArgs struct {
	Name       pulumi.StringInput                                 `pulumi:"name"`
	Parameters RequestHeaderMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestHeaderConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestHeaderConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleRequestHeaderConditionResponseArgs) ToDeliveryRuleRequestHeaderConditionResponseOutput() DeliveryRuleRequestHeaderConditionResponseOutput {
	return i.ToDeliveryRuleRequestHeaderConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestHeaderConditionResponseArgs) ToDeliveryRuleRequestHeaderConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestHeaderConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestHeaderConditionResponseOutput)
}

type DeliveryRuleRequestHeaderConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestHeaderConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestHeaderConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleRequestHeaderConditionResponseOutput) ToDeliveryRuleRequestHeaderConditionResponseOutput() DeliveryRuleRequestHeaderConditionResponseOutput {
	return o
}

func (o DeliveryRuleRequestHeaderConditionResponseOutput) ToDeliveryRuleRequestHeaderConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestHeaderConditionResponseOutput {
	return o
}

func (o DeliveryRuleRequestHeaderConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestHeaderConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestHeaderConditionResponseOutput) Parameters() RequestHeaderMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleRequestHeaderConditionResponse) RequestHeaderMatchConditionParametersResponse {
		return v.Parameters
	}).(RequestHeaderMatchConditionParametersResponseOutput)
}

type DeliveryRuleRequestMethodCondition struct {
	Name       string                                `pulumi:"name"`
	Parameters RequestMethodMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleRequestMethodConditionInput interface {
	pulumi.Input

	ToDeliveryRuleRequestMethodConditionOutput() DeliveryRuleRequestMethodConditionOutput
	ToDeliveryRuleRequestMethodConditionOutputWithContext(context.Context) DeliveryRuleRequestMethodConditionOutput
}

type DeliveryRuleRequestMethodConditionArgs struct {
	Name       pulumi.StringInput                         `pulumi:"name"`
	Parameters RequestMethodMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestMethodConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestMethodCondition)(nil)).Elem()
}

func (i DeliveryRuleRequestMethodConditionArgs) ToDeliveryRuleRequestMethodConditionOutput() DeliveryRuleRequestMethodConditionOutput {
	return i.ToDeliveryRuleRequestMethodConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestMethodConditionArgs) ToDeliveryRuleRequestMethodConditionOutputWithContext(ctx context.Context) DeliveryRuleRequestMethodConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestMethodConditionOutput)
}

type DeliveryRuleRequestMethodConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestMethodConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestMethodCondition)(nil)).Elem()
}

func (o DeliveryRuleRequestMethodConditionOutput) ToDeliveryRuleRequestMethodConditionOutput() DeliveryRuleRequestMethodConditionOutput {
	return o
}

func (o DeliveryRuleRequestMethodConditionOutput) ToDeliveryRuleRequestMethodConditionOutputWithContext(ctx context.Context) DeliveryRuleRequestMethodConditionOutput {
	return o
}

func (o DeliveryRuleRequestMethodConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestMethodCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestMethodConditionOutput) Parameters() RequestMethodMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleRequestMethodCondition) RequestMethodMatchConditionParameters { return v.Parameters }).(RequestMethodMatchConditionParametersOutput)
}

type DeliveryRuleRequestMethodConditionResponse struct {
	Name       string                                        `pulumi:"name"`
	Parameters RequestMethodMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleRequestMethodConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleRequestMethodConditionResponseOutput() DeliveryRuleRequestMethodConditionResponseOutput
	ToDeliveryRuleRequestMethodConditionResponseOutputWithContext(context.Context) DeliveryRuleRequestMethodConditionResponseOutput
}

type DeliveryRuleRequestMethodConditionResponseArgs struct {
	Name       pulumi.StringInput                                 `pulumi:"name"`
	Parameters RequestMethodMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestMethodConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestMethodConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleRequestMethodConditionResponseArgs) ToDeliveryRuleRequestMethodConditionResponseOutput() DeliveryRuleRequestMethodConditionResponseOutput {
	return i.ToDeliveryRuleRequestMethodConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestMethodConditionResponseArgs) ToDeliveryRuleRequestMethodConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestMethodConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestMethodConditionResponseOutput)
}

type DeliveryRuleRequestMethodConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestMethodConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestMethodConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleRequestMethodConditionResponseOutput) ToDeliveryRuleRequestMethodConditionResponseOutput() DeliveryRuleRequestMethodConditionResponseOutput {
	return o
}

func (o DeliveryRuleRequestMethodConditionResponseOutput) ToDeliveryRuleRequestMethodConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestMethodConditionResponseOutput {
	return o
}

func (o DeliveryRuleRequestMethodConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestMethodConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestMethodConditionResponseOutput) Parameters() RequestMethodMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleRequestMethodConditionResponse) RequestMethodMatchConditionParametersResponse {
		return v.Parameters
	}).(RequestMethodMatchConditionParametersResponseOutput)
}

type DeliveryRuleRequestSchemeCondition struct {
	Name       string                                `pulumi:"name"`
	Parameters RequestSchemeMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleRequestSchemeConditionInput interface {
	pulumi.Input

	ToDeliveryRuleRequestSchemeConditionOutput() DeliveryRuleRequestSchemeConditionOutput
	ToDeliveryRuleRequestSchemeConditionOutputWithContext(context.Context) DeliveryRuleRequestSchemeConditionOutput
}

type DeliveryRuleRequestSchemeConditionArgs struct {
	Name       pulumi.StringInput                         `pulumi:"name"`
	Parameters RequestSchemeMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestSchemeConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestSchemeCondition)(nil)).Elem()
}

func (i DeliveryRuleRequestSchemeConditionArgs) ToDeliveryRuleRequestSchemeConditionOutput() DeliveryRuleRequestSchemeConditionOutput {
	return i.ToDeliveryRuleRequestSchemeConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestSchemeConditionArgs) ToDeliveryRuleRequestSchemeConditionOutputWithContext(ctx context.Context) DeliveryRuleRequestSchemeConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestSchemeConditionOutput)
}

type DeliveryRuleRequestSchemeConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestSchemeConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestSchemeCondition)(nil)).Elem()
}

func (o DeliveryRuleRequestSchemeConditionOutput) ToDeliveryRuleRequestSchemeConditionOutput() DeliveryRuleRequestSchemeConditionOutput {
	return o
}

func (o DeliveryRuleRequestSchemeConditionOutput) ToDeliveryRuleRequestSchemeConditionOutputWithContext(ctx context.Context) DeliveryRuleRequestSchemeConditionOutput {
	return o
}

func (o DeliveryRuleRequestSchemeConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestSchemeCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestSchemeConditionOutput) Parameters() RequestSchemeMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleRequestSchemeCondition) RequestSchemeMatchConditionParameters { return v.Parameters }).(RequestSchemeMatchConditionParametersOutput)
}

type DeliveryRuleRequestSchemeConditionResponse struct {
	Name       string                                        `pulumi:"name"`
	Parameters RequestSchemeMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleRequestSchemeConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleRequestSchemeConditionResponseOutput() DeliveryRuleRequestSchemeConditionResponseOutput
	ToDeliveryRuleRequestSchemeConditionResponseOutputWithContext(context.Context) DeliveryRuleRequestSchemeConditionResponseOutput
}

type DeliveryRuleRequestSchemeConditionResponseArgs struct {
	Name       pulumi.StringInput                                 `pulumi:"name"`
	Parameters RequestSchemeMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestSchemeConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestSchemeConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleRequestSchemeConditionResponseArgs) ToDeliveryRuleRequestSchemeConditionResponseOutput() DeliveryRuleRequestSchemeConditionResponseOutput {
	return i.ToDeliveryRuleRequestSchemeConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestSchemeConditionResponseArgs) ToDeliveryRuleRequestSchemeConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestSchemeConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestSchemeConditionResponseOutput)
}

type DeliveryRuleRequestSchemeConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestSchemeConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestSchemeConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleRequestSchemeConditionResponseOutput) ToDeliveryRuleRequestSchemeConditionResponseOutput() DeliveryRuleRequestSchemeConditionResponseOutput {
	return o
}

func (o DeliveryRuleRequestSchemeConditionResponseOutput) ToDeliveryRuleRequestSchemeConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestSchemeConditionResponseOutput {
	return o
}

func (o DeliveryRuleRequestSchemeConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestSchemeConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestSchemeConditionResponseOutput) Parameters() RequestSchemeMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleRequestSchemeConditionResponse) RequestSchemeMatchConditionParametersResponse {
		return v.Parameters
	}).(RequestSchemeMatchConditionParametersResponseOutput)
}

type DeliveryRuleRequestUriCondition struct {
	Name       string                             `pulumi:"name"`
	Parameters RequestUriMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleRequestUriConditionInput interface {
	pulumi.Input

	ToDeliveryRuleRequestUriConditionOutput() DeliveryRuleRequestUriConditionOutput
	ToDeliveryRuleRequestUriConditionOutputWithContext(context.Context) DeliveryRuleRequestUriConditionOutput
}

type DeliveryRuleRequestUriConditionArgs struct {
	Name       pulumi.StringInput                      `pulumi:"name"`
	Parameters RequestUriMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestUriConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestUriCondition)(nil)).Elem()
}

func (i DeliveryRuleRequestUriConditionArgs) ToDeliveryRuleRequestUriConditionOutput() DeliveryRuleRequestUriConditionOutput {
	return i.ToDeliveryRuleRequestUriConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestUriConditionArgs) ToDeliveryRuleRequestUriConditionOutputWithContext(ctx context.Context) DeliveryRuleRequestUriConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestUriConditionOutput)
}

type DeliveryRuleRequestUriConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestUriConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestUriCondition)(nil)).Elem()
}

func (o DeliveryRuleRequestUriConditionOutput) ToDeliveryRuleRequestUriConditionOutput() DeliveryRuleRequestUriConditionOutput {
	return o
}

func (o DeliveryRuleRequestUriConditionOutput) ToDeliveryRuleRequestUriConditionOutputWithContext(ctx context.Context) DeliveryRuleRequestUriConditionOutput {
	return o
}

func (o DeliveryRuleRequestUriConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestUriCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestUriConditionOutput) Parameters() RequestUriMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleRequestUriCondition) RequestUriMatchConditionParameters { return v.Parameters }).(RequestUriMatchConditionParametersOutput)
}

type DeliveryRuleRequestUriConditionResponse struct {
	Name       string                                     `pulumi:"name"`
	Parameters RequestUriMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleRequestUriConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleRequestUriConditionResponseOutput() DeliveryRuleRequestUriConditionResponseOutput
	ToDeliveryRuleRequestUriConditionResponseOutputWithContext(context.Context) DeliveryRuleRequestUriConditionResponseOutput
}

type DeliveryRuleRequestUriConditionResponseArgs struct {
	Name       pulumi.StringInput                              `pulumi:"name"`
	Parameters RequestUriMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleRequestUriConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestUriConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleRequestUriConditionResponseArgs) ToDeliveryRuleRequestUriConditionResponseOutput() DeliveryRuleRequestUriConditionResponseOutput {
	return i.ToDeliveryRuleRequestUriConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleRequestUriConditionResponseArgs) ToDeliveryRuleRequestUriConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestUriConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleRequestUriConditionResponseOutput)
}

type DeliveryRuleRequestUriConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleRequestUriConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleRequestUriConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleRequestUriConditionResponseOutput) ToDeliveryRuleRequestUriConditionResponseOutput() DeliveryRuleRequestUriConditionResponseOutput {
	return o
}

func (o DeliveryRuleRequestUriConditionResponseOutput) ToDeliveryRuleRequestUriConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleRequestUriConditionResponseOutput {
	return o
}

func (o DeliveryRuleRequestUriConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleRequestUriConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleRequestUriConditionResponseOutput) Parameters() RequestUriMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleRequestUriConditionResponse) RequestUriMatchConditionParametersResponse {
		return v.Parameters
	}).(RequestUriMatchConditionParametersResponseOutput)
}

type DeliveryRuleResponse struct {
	Actions    []interface{} `pulumi:"actions"`
	Conditions []interface{} `pulumi:"conditions"`
	Name       *string       `pulumi:"name"`
	Order      int           `pulumi:"order"`
}





type DeliveryRuleResponseInput interface {
	pulumi.Input

	ToDeliveryRuleResponseOutput() DeliveryRuleResponseOutput
	ToDeliveryRuleResponseOutputWithContext(context.Context) DeliveryRuleResponseOutput
}

type DeliveryRuleResponseArgs struct {
	Actions    pulumi.ArrayInput     `pulumi:"actions"`
	Conditions pulumi.ArrayInput     `pulumi:"conditions"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Order      pulumi.IntInput       `pulumi:"order"`
}

func (DeliveryRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleResponse)(nil)).Elem()
}

func (i DeliveryRuleResponseArgs) ToDeliveryRuleResponseOutput() DeliveryRuleResponseOutput {
	return i.ToDeliveryRuleResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleResponseArgs) ToDeliveryRuleResponseOutputWithContext(ctx context.Context) DeliveryRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleResponseOutput)
}





type DeliveryRuleResponseArrayInput interface {
	pulumi.Input

	ToDeliveryRuleResponseArrayOutput() DeliveryRuleResponseArrayOutput
	ToDeliveryRuleResponseArrayOutputWithContext(context.Context) DeliveryRuleResponseArrayOutput
}

type DeliveryRuleResponseArray []DeliveryRuleResponseInput

func (DeliveryRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeliveryRuleResponse)(nil)).Elem()
}

func (i DeliveryRuleResponseArray) ToDeliveryRuleResponseArrayOutput() DeliveryRuleResponseArrayOutput {
	return i.ToDeliveryRuleResponseArrayOutputWithContext(context.Background())
}

func (i DeliveryRuleResponseArray) ToDeliveryRuleResponseArrayOutputWithContext(ctx context.Context) DeliveryRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleResponseArrayOutput)
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





type DeliveryRuleResponseHeaderActionInput interface {
	pulumi.Input

	ToDeliveryRuleResponseHeaderActionOutput() DeliveryRuleResponseHeaderActionOutput
	ToDeliveryRuleResponseHeaderActionOutputWithContext(context.Context) DeliveryRuleResponseHeaderActionOutput
}

type DeliveryRuleResponseHeaderActionArgs struct {
	Name       pulumi.StringInput          `pulumi:"name"`
	Parameters HeaderActionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleResponseHeaderActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleResponseHeaderAction)(nil)).Elem()
}

func (i DeliveryRuleResponseHeaderActionArgs) ToDeliveryRuleResponseHeaderActionOutput() DeliveryRuleResponseHeaderActionOutput {
	return i.ToDeliveryRuleResponseHeaderActionOutputWithContext(context.Background())
}

func (i DeliveryRuleResponseHeaderActionArgs) ToDeliveryRuleResponseHeaderActionOutputWithContext(ctx context.Context) DeliveryRuleResponseHeaderActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleResponseHeaderActionOutput)
}

type DeliveryRuleResponseHeaderActionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleResponseHeaderActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleResponseHeaderAction)(nil)).Elem()
}

func (o DeliveryRuleResponseHeaderActionOutput) ToDeliveryRuleResponseHeaderActionOutput() DeliveryRuleResponseHeaderActionOutput {
	return o
}

func (o DeliveryRuleResponseHeaderActionOutput) ToDeliveryRuleResponseHeaderActionOutputWithContext(ctx context.Context) DeliveryRuleResponseHeaderActionOutput {
	return o
}

func (o DeliveryRuleResponseHeaderActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleResponseHeaderAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleResponseHeaderActionOutput) Parameters() HeaderActionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleResponseHeaderAction) HeaderActionParameters { return v.Parameters }).(HeaderActionParametersOutput)
}

type DeliveryRuleResponseHeaderActionResponse struct {
	Name       string                         `pulumi:"name"`
	Parameters HeaderActionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleResponseHeaderActionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleResponseHeaderActionResponseOutput() DeliveryRuleResponseHeaderActionResponseOutput
	ToDeliveryRuleResponseHeaderActionResponseOutputWithContext(context.Context) DeliveryRuleResponseHeaderActionResponseOutput
}

type DeliveryRuleResponseHeaderActionResponseArgs struct {
	Name       pulumi.StringInput                  `pulumi:"name"`
	Parameters HeaderActionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleResponseHeaderActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleResponseHeaderActionResponse)(nil)).Elem()
}

func (i DeliveryRuleResponseHeaderActionResponseArgs) ToDeliveryRuleResponseHeaderActionResponseOutput() DeliveryRuleResponseHeaderActionResponseOutput {
	return i.ToDeliveryRuleResponseHeaderActionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleResponseHeaderActionResponseArgs) ToDeliveryRuleResponseHeaderActionResponseOutputWithContext(ctx context.Context) DeliveryRuleResponseHeaderActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleResponseHeaderActionResponseOutput)
}

type DeliveryRuleResponseHeaderActionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleResponseHeaderActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleResponseHeaderActionResponse)(nil)).Elem()
}

func (o DeliveryRuleResponseHeaderActionResponseOutput) ToDeliveryRuleResponseHeaderActionResponseOutput() DeliveryRuleResponseHeaderActionResponseOutput {
	return o
}

func (o DeliveryRuleResponseHeaderActionResponseOutput) ToDeliveryRuleResponseHeaderActionResponseOutputWithContext(ctx context.Context) DeliveryRuleResponseHeaderActionResponseOutput {
	return o
}

func (o DeliveryRuleResponseHeaderActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleResponseHeaderActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleResponseHeaderActionResponseOutput) Parameters() HeaderActionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleResponseHeaderActionResponse) HeaderActionParametersResponse { return v.Parameters }).(HeaderActionParametersResponseOutput)
}

type DeliveryRuleUrlFileExtensionCondition struct {
	Name       string                                   `pulumi:"name"`
	Parameters UrlFileExtensionMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleUrlFileExtensionConditionInput interface {
	pulumi.Input

	ToDeliveryRuleUrlFileExtensionConditionOutput() DeliveryRuleUrlFileExtensionConditionOutput
	ToDeliveryRuleUrlFileExtensionConditionOutputWithContext(context.Context) DeliveryRuleUrlFileExtensionConditionOutput
}

type DeliveryRuleUrlFileExtensionConditionArgs struct {
	Name       pulumi.StringInput                            `pulumi:"name"`
	Parameters UrlFileExtensionMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleUrlFileExtensionConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlFileExtensionCondition)(nil)).Elem()
}

func (i DeliveryRuleUrlFileExtensionConditionArgs) ToDeliveryRuleUrlFileExtensionConditionOutput() DeliveryRuleUrlFileExtensionConditionOutput {
	return i.ToDeliveryRuleUrlFileExtensionConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleUrlFileExtensionConditionArgs) ToDeliveryRuleUrlFileExtensionConditionOutputWithContext(ctx context.Context) DeliveryRuleUrlFileExtensionConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleUrlFileExtensionConditionOutput)
}

type DeliveryRuleUrlFileExtensionConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleUrlFileExtensionConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlFileExtensionCondition)(nil)).Elem()
}

func (o DeliveryRuleUrlFileExtensionConditionOutput) ToDeliveryRuleUrlFileExtensionConditionOutput() DeliveryRuleUrlFileExtensionConditionOutput {
	return o
}

func (o DeliveryRuleUrlFileExtensionConditionOutput) ToDeliveryRuleUrlFileExtensionConditionOutputWithContext(ctx context.Context) DeliveryRuleUrlFileExtensionConditionOutput {
	return o
}

func (o DeliveryRuleUrlFileExtensionConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleUrlFileExtensionCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleUrlFileExtensionConditionOutput) Parameters() UrlFileExtensionMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleUrlFileExtensionCondition) UrlFileExtensionMatchConditionParameters {
		return v.Parameters
	}).(UrlFileExtensionMatchConditionParametersOutput)
}

type DeliveryRuleUrlFileExtensionConditionResponse struct {
	Name       string                                           `pulumi:"name"`
	Parameters UrlFileExtensionMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleUrlFileExtensionConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleUrlFileExtensionConditionResponseOutput() DeliveryRuleUrlFileExtensionConditionResponseOutput
	ToDeliveryRuleUrlFileExtensionConditionResponseOutputWithContext(context.Context) DeliveryRuleUrlFileExtensionConditionResponseOutput
}

type DeliveryRuleUrlFileExtensionConditionResponseArgs struct {
	Name       pulumi.StringInput                                    `pulumi:"name"`
	Parameters UrlFileExtensionMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleUrlFileExtensionConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlFileExtensionConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleUrlFileExtensionConditionResponseArgs) ToDeliveryRuleUrlFileExtensionConditionResponseOutput() DeliveryRuleUrlFileExtensionConditionResponseOutput {
	return i.ToDeliveryRuleUrlFileExtensionConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleUrlFileExtensionConditionResponseArgs) ToDeliveryRuleUrlFileExtensionConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleUrlFileExtensionConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleUrlFileExtensionConditionResponseOutput)
}

type DeliveryRuleUrlFileExtensionConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleUrlFileExtensionConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlFileExtensionConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleUrlFileExtensionConditionResponseOutput) ToDeliveryRuleUrlFileExtensionConditionResponseOutput() DeliveryRuleUrlFileExtensionConditionResponseOutput {
	return o
}

func (o DeliveryRuleUrlFileExtensionConditionResponseOutput) ToDeliveryRuleUrlFileExtensionConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleUrlFileExtensionConditionResponseOutput {
	return o
}

func (o DeliveryRuleUrlFileExtensionConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleUrlFileExtensionConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleUrlFileExtensionConditionResponseOutput) Parameters() UrlFileExtensionMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleUrlFileExtensionConditionResponse) UrlFileExtensionMatchConditionParametersResponse {
		return v.Parameters
	}).(UrlFileExtensionMatchConditionParametersResponseOutput)
}

type DeliveryRuleUrlFileNameCondition struct {
	Name       string                              `pulumi:"name"`
	Parameters UrlFileNameMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleUrlFileNameConditionInput interface {
	pulumi.Input

	ToDeliveryRuleUrlFileNameConditionOutput() DeliveryRuleUrlFileNameConditionOutput
	ToDeliveryRuleUrlFileNameConditionOutputWithContext(context.Context) DeliveryRuleUrlFileNameConditionOutput
}

type DeliveryRuleUrlFileNameConditionArgs struct {
	Name       pulumi.StringInput                       `pulumi:"name"`
	Parameters UrlFileNameMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleUrlFileNameConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlFileNameCondition)(nil)).Elem()
}

func (i DeliveryRuleUrlFileNameConditionArgs) ToDeliveryRuleUrlFileNameConditionOutput() DeliveryRuleUrlFileNameConditionOutput {
	return i.ToDeliveryRuleUrlFileNameConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleUrlFileNameConditionArgs) ToDeliveryRuleUrlFileNameConditionOutputWithContext(ctx context.Context) DeliveryRuleUrlFileNameConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleUrlFileNameConditionOutput)
}

type DeliveryRuleUrlFileNameConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleUrlFileNameConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlFileNameCondition)(nil)).Elem()
}

func (o DeliveryRuleUrlFileNameConditionOutput) ToDeliveryRuleUrlFileNameConditionOutput() DeliveryRuleUrlFileNameConditionOutput {
	return o
}

func (o DeliveryRuleUrlFileNameConditionOutput) ToDeliveryRuleUrlFileNameConditionOutputWithContext(ctx context.Context) DeliveryRuleUrlFileNameConditionOutput {
	return o
}

func (o DeliveryRuleUrlFileNameConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleUrlFileNameCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleUrlFileNameConditionOutput) Parameters() UrlFileNameMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleUrlFileNameCondition) UrlFileNameMatchConditionParameters { return v.Parameters }).(UrlFileNameMatchConditionParametersOutput)
}

type DeliveryRuleUrlFileNameConditionResponse struct {
	Name       string                                      `pulumi:"name"`
	Parameters UrlFileNameMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleUrlFileNameConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleUrlFileNameConditionResponseOutput() DeliveryRuleUrlFileNameConditionResponseOutput
	ToDeliveryRuleUrlFileNameConditionResponseOutputWithContext(context.Context) DeliveryRuleUrlFileNameConditionResponseOutput
}

type DeliveryRuleUrlFileNameConditionResponseArgs struct {
	Name       pulumi.StringInput                               `pulumi:"name"`
	Parameters UrlFileNameMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleUrlFileNameConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlFileNameConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleUrlFileNameConditionResponseArgs) ToDeliveryRuleUrlFileNameConditionResponseOutput() DeliveryRuleUrlFileNameConditionResponseOutput {
	return i.ToDeliveryRuleUrlFileNameConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleUrlFileNameConditionResponseArgs) ToDeliveryRuleUrlFileNameConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleUrlFileNameConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleUrlFileNameConditionResponseOutput)
}

type DeliveryRuleUrlFileNameConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleUrlFileNameConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlFileNameConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleUrlFileNameConditionResponseOutput) ToDeliveryRuleUrlFileNameConditionResponseOutput() DeliveryRuleUrlFileNameConditionResponseOutput {
	return o
}

func (o DeliveryRuleUrlFileNameConditionResponseOutput) ToDeliveryRuleUrlFileNameConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleUrlFileNameConditionResponseOutput {
	return o
}

func (o DeliveryRuleUrlFileNameConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleUrlFileNameConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleUrlFileNameConditionResponseOutput) Parameters() UrlFileNameMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleUrlFileNameConditionResponse) UrlFileNameMatchConditionParametersResponse {
		return v.Parameters
	}).(UrlFileNameMatchConditionParametersResponseOutput)
}

type DeliveryRuleUrlPathCondition struct {
	Name       string                          `pulumi:"name"`
	Parameters UrlPathMatchConditionParameters `pulumi:"parameters"`
}





type DeliveryRuleUrlPathConditionInput interface {
	pulumi.Input

	ToDeliveryRuleUrlPathConditionOutput() DeliveryRuleUrlPathConditionOutput
	ToDeliveryRuleUrlPathConditionOutputWithContext(context.Context) DeliveryRuleUrlPathConditionOutput
}

type DeliveryRuleUrlPathConditionArgs struct {
	Name       pulumi.StringInput                   `pulumi:"name"`
	Parameters UrlPathMatchConditionParametersInput `pulumi:"parameters"`
}

func (DeliveryRuleUrlPathConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlPathCondition)(nil)).Elem()
}

func (i DeliveryRuleUrlPathConditionArgs) ToDeliveryRuleUrlPathConditionOutput() DeliveryRuleUrlPathConditionOutput {
	return i.ToDeliveryRuleUrlPathConditionOutputWithContext(context.Background())
}

func (i DeliveryRuleUrlPathConditionArgs) ToDeliveryRuleUrlPathConditionOutputWithContext(ctx context.Context) DeliveryRuleUrlPathConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleUrlPathConditionOutput)
}

type DeliveryRuleUrlPathConditionOutput struct{ *pulumi.OutputState }

func (DeliveryRuleUrlPathConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlPathCondition)(nil)).Elem()
}

func (o DeliveryRuleUrlPathConditionOutput) ToDeliveryRuleUrlPathConditionOutput() DeliveryRuleUrlPathConditionOutput {
	return o
}

func (o DeliveryRuleUrlPathConditionOutput) ToDeliveryRuleUrlPathConditionOutputWithContext(ctx context.Context) DeliveryRuleUrlPathConditionOutput {
	return o
}

func (o DeliveryRuleUrlPathConditionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleUrlPathCondition) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleUrlPathConditionOutput) Parameters() UrlPathMatchConditionParametersOutput {
	return o.ApplyT(func(v DeliveryRuleUrlPathCondition) UrlPathMatchConditionParameters { return v.Parameters }).(UrlPathMatchConditionParametersOutput)
}

type DeliveryRuleUrlPathConditionResponse struct {
	Name       string                                  `pulumi:"name"`
	Parameters UrlPathMatchConditionParametersResponse `pulumi:"parameters"`
}





type DeliveryRuleUrlPathConditionResponseInput interface {
	pulumi.Input

	ToDeliveryRuleUrlPathConditionResponseOutput() DeliveryRuleUrlPathConditionResponseOutput
	ToDeliveryRuleUrlPathConditionResponseOutputWithContext(context.Context) DeliveryRuleUrlPathConditionResponseOutput
}

type DeliveryRuleUrlPathConditionResponseArgs struct {
	Name       pulumi.StringInput                           `pulumi:"name"`
	Parameters UrlPathMatchConditionParametersResponseInput `pulumi:"parameters"`
}

func (DeliveryRuleUrlPathConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlPathConditionResponse)(nil)).Elem()
}

func (i DeliveryRuleUrlPathConditionResponseArgs) ToDeliveryRuleUrlPathConditionResponseOutput() DeliveryRuleUrlPathConditionResponseOutput {
	return i.ToDeliveryRuleUrlPathConditionResponseOutputWithContext(context.Background())
}

func (i DeliveryRuleUrlPathConditionResponseArgs) ToDeliveryRuleUrlPathConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleUrlPathConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryRuleUrlPathConditionResponseOutput)
}

type DeliveryRuleUrlPathConditionResponseOutput struct{ *pulumi.OutputState }

func (DeliveryRuleUrlPathConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryRuleUrlPathConditionResponse)(nil)).Elem()
}

func (o DeliveryRuleUrlPathConditionResponseOutput) ToDeliveryRuleUrlPathConditionResponseOutput() DeliveryRuleUrlPathConditionResponseOutput {
	return o
}

func (o DeliveryRuleUrlPathConditionResponseOutput) ToDeliveryRuleUrlPathConditionResponseOutputWithContext(ctx context.Context) DeliveryRuleUrlPathConditionResponseOutput {
	return o
}

func (o DeliveryRuleUrlPathConditionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryRuleUrlPathConditionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeliveryRuleUrlPathConditionResponseOutput) Parameters() UrlPathMatchConditionParametersResponseOutput {
	return o.ApplyT(func(v DeliveryRuleUrlPathConditionResponse) UrlPathMatchConditionParametersResponse {
		return v.Parameters
	}).(UrlPathMatchConditionParametersResponseOutput)
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





type EndpointPropertiesUpdateParametersResponseDeliveryPolicyInput interface {
	pulumi.Input

	ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput() EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput
	ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyOutputWithContext(context.Context) EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput
}

type EndpointPropertiesUpdateParametersResponseDeliveryPolicyArgs struct {
	Description pulumi.StringPtrInput          `pulumi:"description"`
	Rules       DeliveryRuleResponseArrayInput `pulumi:"rules"`
}

func (EndpointPropertiesUpdateParametersResponseDeliveryPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesUpdateParametersResponseDeliveryPolicy)(nil)).Elem()
}

func (i EndpointPropertiesUpdateParametersResponseDeliveryPolicyArgs) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput() EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput {
	return i.ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyOutputWithContext(context.Background())
}

func (i EndpointPropertiesUpdateParametersResponseDeliveryPolicyArgs) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput)
}

func (i EndpointPropertiesUpdateParametersResponseDeliveryPolicyArgs) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput() EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
	return i.ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutputWithContext(context.Background())
}

func (i EndpointPropertiesUpdateParametersResponseDeliveryPolicyArgs) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput).ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutputWithContext(ctx)
}









type EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrInput interface {
	pulumi.Input

	ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput() EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput
	ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutputWithContext(context.Context) EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput
}

type endpointPropertiesUpdateParametersResponseDeliveryPolicyPtrType EndpointPropertiesUpdateParametersResponseDeliveryPolicyArgs

func EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtr(v *EndpointPropertiesUpdateParametersResponseDeliveryPolicyArgs) EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrInput {
	return (*endpointPropertiesUpdateParametersResponseDeliveryPolicyPtrType)(v)
}

func (*endpointPropertiesUpdateParametersResponseDeliveryPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointPropertiesUpdateParametersResponseDeliveryPolicy)(nil)).Elem()
}

func (i *endpointPropertiesUpdateParametersResponseDeliveryPolicyPtrType) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput() EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
	return i.ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutputWithContext(context.Background())
}

func (i *endpointPropertiesUpdateParametersResponseDeliveryPolicyPtrType) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput)
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

func (o EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput() EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
	return o.ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutputWithContext(context.Background())
}

func (o EndpointPropertiesUpdateParametersResponseDeliveryPolicyOutput) ToEndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointPropertiesUpdateParametersResponseDeliveryPolicy) *EndpointPropertiesUpdateParametersResponseDeliveryPolicy {
		return &v
	}).(EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput)
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





type EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkInput interface {
	pulumi.Input

	ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput
	ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutputWithContext(context.Context) EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput
}

type EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkArgs) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return i.ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutputWithContext(context.Background())
}

func (i EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkArgs) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput)
}

func (i EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkArgs) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkArgs) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput).ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx)
}









type EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrInput interface {
	pulumi.Input

	ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput
	ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Context) EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput
}

type endpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrType EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkArgs

func EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtr(v *EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkArgs) EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrInput {
	return (*endpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrType)(v)
}

func (*endpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i *endpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrType) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i *endpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrType) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput)
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

func (o EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (o EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToEndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink) *EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink {
		return &v
	}).(EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput)
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





type GeoFilterResponseInput interface {
	pulumi.Input

	ToGeoFilterResponseOutput() GeoFilterResponseOutput
	ToGeoFilterResponseOutputWithContext(context.Context) GeoFilterResponseOutput
}

type GeoFilterResponseArgs struct {
	Action       pulumi.StringInput      `pulumi:"action"`
	CountryCodes pulumi.StringArrayInput `pulumi:"countryCodes"`
	RelativePath pulumi.StringInput      `pulumi:"relativePath"`
}

func (GeoFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoFilterResponse)(nil)).Elem()
}

func (i GeoFilterResponseArgs) ToGeoFilterResponseOutput() GeoFilterResponseOutput {
	return i.ToGeoFilterResponseOutputWithContext(context.Background())
}

func (i GeoFilterResponseArgs) ToGeoFilterResponseOutputWithContext(ctx context.Context) GeoFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoFilterResponseOutput)
}





type GeoFilterResponseArrayInput interface {
	pulumi.Input

	ToGeoFilterResponseArrayOutput() GeoFilterResponseArrayOutput
	ToGeoFilterResponseArrayOutputWithContext(context.Context) GeoFilterResponseArrayOutput
}

type GeoFilterResponseArray []GeoFilterResponseInput

func (GeoFilterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GeoFilterResponse)(nil)).Elem()
}

func (i GeoFilterResponseArray) ToGeoFilterResponseArrayOutput() GeoFilterResponseArrayOutput {
	return i.ToGeoFilterResponseArrayOutputWithContext(context.Background())
}

func (i GeoFilterResponseArray) ToGeoFilterResponseArrayOutputWithContext(ctx context.Context) GeoFilterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoFilterResponseArrayOutput)
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





type HeaderActionParametersInput interface {
	pulumi.Input

	ToHeaderActionParametersOutput() HeaderActionParametersOutput
	ToHeaderActionParametersOutputWithContext(context.Context) HeaderActionParametersOutput
}

type HeaderActionParametersArgs struct {
	HeaderAction pulumi.StringInput    `pulumi:"headerAction"`
	HeaderName   pulumi.StringInput    `pulumi:"headerName"`
	OdataType    pulumi.StringInput    `pulumi:"odataType"`
	Value        pulumi.StringPtrInput `pulumi:"value"`
}

func (HeaderActionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderActionParameters)(nil)).Elem()
}

func (i HeaderActionParametersArgs) ToHeaderActionParametersOutput() HeaderActionParametersOutput {
	return i.ToHeaderActionParametersOutputWithContext(context.Background())
}

func (i HeaderActionParametersArgs) ToHeaderActionParametersOutputWithContext(ctx context.Context) HeaderActionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeaderActionParametersOutput)
}

type HeaderActionParametersOutput struct{ *pulumi.OutputState }

func (HeaderActionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderActionParameters)(nil)).Elem()
}

func (o HeaderActionParametersOutput) ToHeaderActionParametersOutput() HeaderActionParametersOutput {
	return o
}

func (o HeaderActionParametersOutput) ToHeaderActionParametersOutputWithContext(ctx context.Context) HeaderActionParametersOutput {
	return o
}

func (o HeaderActionParametersOutput) HeaderAction() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderActionParameters) string { return v.HeaderAction }).(pulumi.StringOutput)
}

func (o HeaderActionParametersOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderActionParameters) string { return v.HeaderName }).(pulumi.StringOutput)
}

func (o HeaderActionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderActionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o HeaderActionParametersOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HeaderActionParameters) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type HeaderActionParametersResponse struct {
	HeaderAction string  `pulumi:"headerAction"`
	HeaderName   string  `pulumi:"headerName"`
	OdataType    string  `pulumi:"odataType"`
	Value        *string `pulumi:"value"`
}





type HeaderActionParametersResponseInput interface {
	pulumi.Input

	ToHeaderActionParametersResponseOutput() HeaderActionParametersResponseOutput
	ToHeaderActionParametersResponseOutputWithContext(context.Context) HeaderActionParametersResponseOutput
}

type HeaderActionParametersResponseArgs struct {
	HeaderAction pulumi.StringInput    `pulumi:"headerAction"`
	HeaderName   pulumi.StringInput    `pulumi:"headerName"`
	OdataType    pulumi.StringInput    `pulumi:"odataType"`
	Value        pulumi.StringPtrInput `pulumi:"value"`
}

func (HeaderActionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderActionParametersResponse)(nil)).Elem()
}

func (i HeaderActionParametersResponseArgs) ToHeaderActionParametersResponseOutput() HeaderActionParametersResponseOutput {
	return i.ToHeaderActionParametersResponseOutputWithContext(context.Background())
}

func (i HeaderActionParametersResponseArgs) ToHeaderActionParametersResponseOutputWithContext(ctx context.Context) HeaderActionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeaderActionParametersResponseOutput)
}

type HeaderActionParametersResponseOutput struct{ *pulumi.OutputState }

func (HeaderActionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderActionParametersResponse)(nil)).Elem()
}

func (o HeaderActionParametersResponseOutput) ToHeaderActionParametersResponseOutput() HeaderActionParametersResponseOutput {
	return o
}

func (o HeaderActionParametersResponseOutput) ToHeaderActionParametersResponseOutputWithContext(ctx context.Context) HeaderActionParametersResponseOutput {
	return o
}

func (o HeaderActionParametersResponseOutput) HeaderAction() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderActionParametersResponse) string { return v.HeaderAction }).(pulumi.StringOutput)
}

func (o HeaderActionParametersResponseOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderActionParametersResponse) string { return v.HeaderName }).(pulumi.StringOutput)
}

func (o HeaderActionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderActionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o HeaderActionParametersResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HeaderActionParametersResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
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





type HealthProbeParametersResponseInput interface {
	pulumi.Input

	ToHealthProbeParametersResponseOutput() HealthProbeParametersResponseOutput
	ToHealthProbeParametersResponseOutputWithContext(context.Context) HealthProbeParametersResponseOutput
}

type HealthProbeParametersResponseArgs struct {
	ProbeIntervalInSeconds pulumi.IntPtrInput    `pulumi:"probeIntervalInSeconds"`
	ProbePath              pulumi.StringPtrInput `pulumi:"probePath"`
	ProbeProtocol          pulumi.StringPtrInput `pulumi:"probeProtocol"`
	ProbeRequestType       pulumi.StringPtrInput `pulumi:"probeRequestType"`
}

func (HealthProbeParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeParametersResponse)(nil)).Elem()
}

func (i HealthProbeParametersResponseArgs) ToHealthProbeParametersResponseOutput() HealthProbeParametersResponseOutput {
	return i.ToHealthProbeParametersResponseOutputWithContext(context.Background())
}

func (i HealthProbeParametersResponseArgs) ToHealthProbeParametersResponseOutputWithContext(ctx context.Context) HealthProbeParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthProbeParametersResponseOutput)
}

func (i HealthProbeParametersResponseArgs) ToHealthProbeParametersResponsePtrOutput() HealthProbeParametersResponsePtrOutput {
	return i.ToHealthProbeParametersResponsePtrOutputWithContext(context.Background())
}

func (i HealthProbeParametersResponseArgs) ToHealthProbeParametersResponsePtrOutputWithContext(ctx context.Context) HealthProbeParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthProbeParametersResponseOutput).ToHealthProbeParametersResponsePtrOutputWithContext(ctx)
}









type HealthProbeParametersResponsePtrInput interface {
	pulumi.Input

	ToHealthProbeParametersResponsePtrOutput() HealthProbeParametersResponsePtrOutput
	ToHealthProbeParametersResponsePtrOutputWithContext(context.Context) HealthProbeParametersResponsePtrOutput
}

type healthProbeParametersResponsePtrType HealthProbeParametersResponseArgs

func HealthProbeParametersResponsePtr(v *HealthProbeParametersResponseArgs) HealthProbeParametersResponsePtrInput {
	return (*healthProbeParametersResponsePtrType)(v)
}

func (*healthProbeParametersResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthProbeParametersResponse)(nil)).Elem()
}

func (i *healthProbeParametersResponsePtrType) ToHealthProbeParametersResponsePtrOutput() HealthProbeParametersResponsePtrOutput {
	return i.ToHealthProbeParametersResponsePtrOutputWithContext(context.Background())
}

func (i *healthProbeParametersResponsePtrType) ToHealthProbeParametersResponsePtrOutputWithContext(ctx context.Context) HealthProbeParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthProbeParametersResponsePtrOutput)
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

func (o HealthProbeParametersResponseOutput) ToHealthProbeParametersResponsePtrOutput() HealthProbeParametersResponsePtrOutput {
	return o.ToHealthProbeParametersResponsePtrOutputWithContext(context.Background())
}

func (o HealthProbeParametersResponseOutput) ToHealthProbeParametersResponsePtrOutputWithContext(ctx context.Context) HealthProbeParametersResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HealthProbeParametersResponse) *HealthProbeParametersResponse {
		return &v
	}).(HealthProbeParametersResponsePtrOutput)
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





type HttpErrorRangeParametersResponseInput interface {
	pulumi.Input

	ToHttpErrorRangeParametersResponseOutput() HttpErrorRangeParametersResponseOutput
	ToHttpErrorRangeParametersResponseOutputWithContext(context.Context) HttpErrorRangeParametersResponseOutput
}

type HttpErrorRangeParametersResponseArgs struct {
	Begin pulumi.IntPtrInput `pulumi:"begin"`
	End   pulumi.IntPtrInput `pulumi:"end"`
}

func (HttpErrorRangeParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpErrorRangeParametersResponse)(nil)).Elem()
}

func (i HttpErrorRangeParametersResponseArgs) ToHttpErrorRangeParametersResponseOutput() HttpErrorRangeParametersResponseOutput {
	return i.ToHttpErrorRangeParametersResponseOutputWithContext(context.Background())
}

func (i HttpErrorRangeParametersResponseArgs) ToHttpErrorRangeParametersResponseOutputWithContext(ctx context.Context) HttpErrorRangeParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpErrorRangeParametersResponseOutput)
}





type HttpErrorRangeParametersResponseArrayInput interface {
	pulumi.Input

	ToHttpErrorRangeParametersResponseArrayOutput() HttpErrorRangeParametersResponseArrayOutput
	ToHttpErrorRangeParametersResponseArrayOutputWithContext(context.Context) HttpErrorRangeParametersResponseArrayOutput
}

type HttpErrorRangeParametersResponseArray []HttpErrorRangeParametersResponseInput

func (HttpErrorRangeParametersResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpErrorRangeParametersResponse)(nil)).Elem()
}

func (i HttpErrorRangeParametersResponseArray) ToHttpErrorRangeParametersResponseArrayOutput() HttpErrorRangeParametersResponseArrayOutput {
	return i.ToHttpErrorRangeParametersResponseArrayOutputWithContext(context.Background())
}

func (i HttpErrorRangeParametersResponseArray) ToHttpErrorRangeParametersResponseArrayOutputWithContext(ctx context.Context) HttpErrorRangeParametersResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpErrorRangeParametersResponseArrayOutput)
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





type HttpVersionMatchConditionParametersInput interface {
	pulumi.Input

	ToHttpVersionMatchConditionParametersOutput() HttpVersionMatchConditionParametersOutput
	ToHttpVersionMatchConditionParametersOutputWithContext(context.Context) HttpVersionMatchConditionParametersOutput
}

type HttpVersionMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
}

func (HttpVersionMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpVersionMatchConditionParameters)(nil)).Elem()
}

func (i HttpVersionMatchConditionParametersArgs) ToHttpVersionMatchConditionParametersOutput() HttpVersionMatchConditionParametersOutput {
	return i.ToHttpVersionMatchConditionParametersOutputWithContext(context.Background())
}

func (i HttpVersionMatchConditionParametersArgs) ToHttpVersionMatchConditionParametersOutputWithContext(ctx context.Context) HttpVersionMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpVersionMatchConditionParametersOutput)
}

type HttpVersionMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (HttpVersionMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpVersionMatchConditionParameters)(nil)).Elem()
}

func (o HttpVersionMatchConditionParametersOutput) ToHttpVersionMatchConditionParametersOutput() HttpVersionMatchConditionParametersOutput {
	return o
}

func (o HttpVersionMatchConditionParametersOutput) ToHttpVersionMatchConditionParametersOutputWithContext(ctx context.Context) HttpVersionMatchConditionParametersOutput {
	return o
}

func (o HttpVersionMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HttpVersionMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o HttpVersionMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HttpVersionMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o HttpVersionMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v HttpVersionMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o HttpVersionMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v HttpVersionMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

type HttpVersionMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
}





type HttpVersionMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToHttpVersionMatchConditionParametersResponseOutput() HttpVersionMatchConditionParametersResponseOutput
	ToHttpVersionMatchConditionParametersResponseOutputWithContext(context.Context) HttpVersionMatchConditionParametersResponseOutput
}

type HttpVersionMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
}

func (HttpVersionMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpVersionMatchConditionParametersResponse)(nil)).Elem()
}

func (i HttpVersionMatchConditionParametersResponseArgs) ToHttpVersionMatchConditionParametersResponseOutput() HttpVersionMatchConditionParametersResponseOutput {
	return i.ToHttpVersionMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i HttpVersionMatchConditionParametersResponseArgs) ToHttpVersionMatchConditionParametersResponseOutputWithContext(ctx context.Context) HttpVersionMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpVersionMatchConditionParametersResponseOutput)
}

type HttpVersionMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (HttpVersionMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpVersionMatchConditionParametersResponse)(nil)).Elem()
}

func (o HttpVersionMatchConditionParametersResponseOutput) ToHttpVersionMatchConditionParametersResponseOutput() HttpVersionMatchConditionParametersResponseOutput {
	return o
}

func (o HttpVersionMatchConditionParametersResponseOutput) ToHttpVersionMatchConditionParametersResponseOutputWithContext(ctx context.Context) HttpVersionMatchConditionParametersResponseOutput {
	return o
}

func (o HttpVersionMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HttpVersionMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o HttpVersionMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HttpVersionMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o HttpVersionMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v HttpVersionMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o HttpVersionMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v HttpVersionMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

type IsDeviceMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type IsDeviceMatchConditionParametersInput interface {
	pulumi.Input

	ToIsDeviceMatchConditionParametersOutput() IsDeviceMatchConditionParametersOutput
	ToIsDeviceMatchConditionParametersOutputWithContext(context.Context) IsDeviceMatchConditionParametersOutput
}

type IsDeviceMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (IsDeviceMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IsDeviceMatchConditionParameters)(nil)).Elem()
}

func (i IsDeviceMatchConditionParametersArgs) ToIsDeviceMatchConditionParametersOutput() IsDeviceMatchConditionParametersOutput {
	return i.ToIsDeviceMatchConditionParametersOutputWithContext(context.Background())
}

func (i IsDeviceMatchConditionParametersArgs) ToIsDeviceMatchConditionParametersOutputWithContext(ctx context.Context) IsDeviceMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IsDeviceMatchConditionParametersOutput)
}

type IsDeviceMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (IsDeviceMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IsDeviceMatchConditionParameters)(nil)).Elem()
}

func (o IsDeviceMatchConditionParametersOutput) ToIsDeviceMatchConditionParametersOutput() IsDeviceMatchConditionParametersOutput {
	return o
}

func (o IsDeviceMatchConditionParametersOutput) ToIsDeviceMatchConditionParametersOutputWithContext(ctx context.Context) IsDeviceMatchConditionParametersOutput {
	return o
}

func (o IsDeviceMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IsDeviceMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o IsDeviceMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IsDeviceMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o IsDeviceMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v IsDeviceMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o IsDeviceMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v IsDeviceMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

func (o IsDeviceMatchConditionParametersOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IsDeviceMatchConditionParameters) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type IsDeviceMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type IsDeviceMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToIsDeviceMatchConditionParametersResponseOutput() IsDeviceMatchConditionParametersResponseOutput
	ToIsDeviceMatchConditionParametersResponseOutputWithContext(context.Context) IsDeviceMatchConditionParametersResponseOutput
}

type IsDeviceMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (IsDeviceMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IsDeviceMatchConditionParametersResponse)(nil)).Elem()
}

func (i IsDeviceMatchConditionParametersResponseArgs) ToIsDeviceMatchConditionParametersResponseOutput() IsDeviceMatchConditionParametersResponseOutput {
	return i.ToIsDeviceMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i IsDeviceMatchConditionParametersResponseArgs) ToIsDeviceMatchConditionParametersResponseOutputWithContext(ctx context.Context) IsDeviceMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IsDeviceMatchConditionParametersResponseOutput)
}

type IsDeviceMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (IsDeviceMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IsDeviceMatchConditionParametersResponse)(nil)).Elem()
}

func (o IsDeviceMatchConditionParametersResponseOutput) ToIsDeviceMatchConditionParametersResponseOutput() IsDeviceMatchConditionParametersResponseOutput {
	return o
}

func (o IsDeviceMatchConditionParametersResponseOutput) ToIsDeviceMatchConditionParametersResponseOutputWithContext(ctx context.Context) IsDeviceMatchConditionParametersResponseOutput {
	return o
}

func (o IsDeviceMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IsDeviceMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o IsDeviceMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IsDeviceMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o IsDeviceMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v IsDeviceMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o IsDeviceMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v IsDeviceMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o IsDeviceMatchConditionParametersResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IsDeviceMatchConditionParametersResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
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





type KeyVaultCertificateSourceParametersResponseInput interface {
	pulumi.Input

	ToKeyVaultCertificateSourceParametersResponseOutput() KeyVaultCertificateSourceParametersResponseOutput
	ToKeyVaultCertificateSourceParametersResponseOutputWithContext(context.Context) KeyVaultCertificateSourceParametersResponseOutput
}

type KeyVaultCertificateSourceParametersResponseArgs struct {
	DeleteRule        pulumi.StringInput    `pulumi:"deleteRule"`
	OdataType         pulumi.StringInput    `pulumi:"odataType"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	SecretName        pulumi.StringInput    `pulumi:"secretName"`
	SecretVersion     pulumi.StringPtrInput `pulumi:"secretVersion"`
	SubscriptionId    pulumi.StringInput    `pulumi:"subscriptionId"`
	UpdateRule        pulumi.StringInput    `pulumi:"updateRule"`
	VaultName         pulumi.StringInput    `pulumi:"vaultName"`
}

func (KeyVaultCertificateSourceParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultCertificateSourceParametersResponse)(nil)).Elem()
}

func (i KeyVaultCertificateSourceParametersResponseArgs) ToKeyVaultCertificateSourceParametersResponseOutput() KeyVaultCertificateSourceParametersResponseOutput {
	return i.ToKeyVaultCertificateSourceParametersResponseOutputWithContext(context.Background())
}

func (i KeyVaultCertificateSourceParametersResponseArgs) ToKeyVaultCertificateSourceParametersResponseOutputWithContext(ctx context.Context) KeyVaultCertificateSourceParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultCertificateSourceParametersResponseOutput)
}

type KeyVaultCertificateSourceParametersResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultCertificateSourceParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultCertificateSourceParametersResponse)(nil)).Elem()
}

func (o KeyVaultCertificateSourceParametersResponseOutput) ToKeyVaultCertificateSourceParametersResponseOutput() KeyVaultCertificateSourceParametersResponseOutput {
	return o
}

func (o KeyVaultCertificateSourceParametersResponseOutput) ToKeyVaultCertificateSourceParametersResponseOutputWithContext(ctx context.Context) KeyVaultCertificateSourceParametersResponseOutput {
	return o
}

func (o KeyVaultCertificateSourceParametersResponseOutput) DeleteRule() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultCertificateSourceParametersResponse) string { return v.DeleteRule }).(pulumi.StringOutput)
}

func (o KeyVaultCertificateSourceParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultCertificateSourceParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o KeyVaultCertificateSourceParametersResponseOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultCertificateSourceParametersResponse) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o KeyVaultCertificateSourceParametersResponseOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultCertificateSourceParametersResponse) string { return v.SecretName }).(pulumi.StringOutput)
}

func (o KeyVaultCertificateSourceParametersResponseOutput) SecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCertificateSourceParametersResponse) *string { return v.SecretVersion }).(pulumi.StringPtrOutput)
}

func (o KeyVaultCertificateSourceParametersResponseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultCertificateSourceParametersResponse) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o KeyVaultCertificateSourceParametersResponseOutput) UpdateRule() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultCertificateSourceParametersResponse) string { return v.UpdateRule }).(pulumi.StringOutput)
}

func (o KeyVaultCertificateSourceParametersResponseOutput) VaultName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultCertificateSourceParametersResponse) string { return v.VaultName }).(pulumi.StringOutput)
}

type KeyVaultSigningKeyParameters struct {
	OdataType         string `pulumi:"odataType"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SecretName        string `pulumi:"secretName"`
	SecretVersion     string `pulumi:"secretVersion"`
	SubscriptionId    string `pulumi:"subscriptionId"`
	VaultName         string `pulumi:"vaultName"`
}





type KeyVaultSigningKeyParametersInput interface {
	pulumi.Input

	ToKeyVaultSigningKeyParametersOutput() KeyVaultSigningKeyParametersOutput
	ToKeyVaultSigningKeyParametersOutputWithContext(context.Context) KeyVaultSigningKeyParametersOutput
}

type KeyVaultSigningKeyParametersArgs struct {
	OdataType         pulumi.StringInput `pulumi:"odataType"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SecretName        pulumi.StringInput `pulumi:"secretName"`
	SecretVersion     pulumi.StringInput `pulumi:"secretVersion"`
	SubscriptionId    pulumi.StringInput `pulumi:"subscriptionId"`
	VaultName         pulumi.StringInput `pulumi:"vaultName"`
}

func (KeyVaultSigningKeyParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSigningKeyParameters)(nil)).Elem()
}

func (i KeyVaultSigningKeyParametersArgs) ToKeyVaultSigningKeyParametersOutput() KeyVaultSigningKeyParametersOutput {
	return i.ToKeyVaultSigningKeyParametersOutputWithContext(context.Background())
}

func (i KeyVaultSigningKeyParametersArgs) ToKeyVaultSigningKeyParametersOutputWithContext(ctx context.Context) KeyVaultSigningKeyParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSigningKeyParametersOutput)
}

type KeyVaultSigningKeyParametersOutput struct{ *pulumi.OutputState }

func (KeyVaultSigningKeyParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSigningKeyParameters)(nil)).Elem()
}

func (o KeyVaultSigningKeyParametersOutput) ToKeyVaultSigningKeyParametersOutput() KeyVaultSigningKeyParametersOutput {
	return o
}

func (o KeyVaultSigningKeyParametersOutput) ToKeyVaultSigningKeyParametersOutputWithContext(ctx context.Context) KeyVaultSigningKeyParametersOutput {
	return o
}

func (o KeyVaultSigningKeyParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o KeyVaultSigningKeyParametersOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParameters) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o KeyVaultSigningKeyParametersOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParameters) string { return v.SecretName }).(pulumi.StringOutput)
}

func (o KeyVaultSigningKeyParametersOutput) SecretVersion() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParameters) string { return v.SecretVersion }).(pulumi.StringOutput)
}

func (o KeyVaultSigningKeyParametersOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParameters) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o KeyVaultSigningKeyParametersOutput) VaultName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParameters) string { return v.VaultName }).(pulumi.StringOutput)
}

type KeyVaultSigningKeyParametersResponse struct {
	OdataType         string `pulumi:"odataType"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SecretName        string `pulumi:"secretName"`
	SecretVersion     string `pulumi:"secretVersion"`
	SubscriptionId    string `pulumi:"subscriptionId"`
	VaultName         string `pulumi:"vaultName"`
}





type KeyVaultSigningKeyParametersResponseInput interface {
	pulumi.Input

	ToKeyVaultSigningKeyParametersResponseOutput() KeyVaultSigningKeyParametersResponseOutput
	ToKeyVaultSigningKeyParametersResponseOutputWithContext(context.Context) KeyVaultSigningKeyParametersResponseOutput
}

type KeyVaultSigningKeyParametersResponseArgs struct {
	OdataType         pulumi.StringInput `pulumi:"odataType"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SecretName        pulumi.StringInput `pulumi:"secretName"`
	SecretVersion     pulumi.StringInput `pulumi:"secretVersion"`
	SubscriptionId    pulumi.StringInput `pulumi:"subscriptionId"`
	VaultName         pulumi.StringInput `pulumi:"vaultName"`
}

func (KeyVaultSigningKeyParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSigningKeyParametersResponse)(nil)).Elem()
}

func (i KeyVaultSigningKeyParametersResponseArgs) ToKeyVaultSigningKeyParametersResponseOutput() KeyVaultSigningKeyParametersResponseOutput {
	return i.ToKeyVaultSigningKeyParametersResponseOutputWithContext(context.Background())
}

func (i KeyVaultSigningKeyParametersResponseArgs) ToKeyVaultSigningKeyParametersResponseOutputWithContext(ctx context.Context) KeyVaultSigningKeyParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSigningKeyParametersResponseOutput)
}

type KeyVaultSigningKeyParametersResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultSigningKeyParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSigningKeyParametersResponse)(nil)).Elem()
}

func (o KeyVaultSigningKeyParametersResponseOutput) ToKeyVaultSigningKeyParametersResponseOutput() KeyVaultSigningKeyParametersResponseOutput {
	return o
}

func (o KeyVaultSigningKeyParametersResponseOutput) ToKeyVaultSigningKeyParametersResponseOutputWithContext(ctx context.Context) KeyVaultSigningKeyParametersResponseOutput {
	return o
}

func (o KeyVaultSigningKeyParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o KeyVaultSigningKeyParametersResponseOutput) ResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParametersResponse) string { return v.ResourceGroupName }).(pulumi.StringOutput)
}

func (o KeyVaultSigningKeyParametersResponseOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParametersResponse) string { return v.SecretName }).(pulumi.StringOutput)
}

func (o KeyVaultSigningKeyParametersResponseOutput) SecretVersion() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParametersResponse) string { return v.SecretVersion }).(pulumi.StringOutput)
}

func (o KeyVaultSigningKeyParametersResponseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParametersResponse) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o KeyVaultSigningKeyParametersResponseOutput) VaultName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSigningKeyParametersResponse) string { return v.VaultName }).(pulumi.StringOutput)
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





type ManagedRuleGroupOverrideResponseInput interface {
	pulumi.Input

	ToManagedRuleGroupOverrideResponseOutput() ManagedRuleGroupOverrideResponseOutput
	ToManagedRuleGroupOverrideResponseOutputWithContext(context.Context) ManagedRuleGroupOverrideResponseOutput
}

type ManagedRuleGroupOverrideResponseArgs struct {
	RuleGroupName pulumi.StringInput                    `pulumi:"ruleGroupName"`
	Rules         ManagedRuleOverrideResponseArrayInput `pulumi:"rules"`
}

func (ManagedRuleGroupOverrideResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleGroupOverrideResponse)(nil)).Elem()
}

func (i ManagedRuleGroupOverrideResponseArgs) ToManagedRuleGroupOverrideResponseOutput() ManagedRuleGroupOverrideResponseOutput {
	return i.ToManagedRuleGroupOverrideResponseOutputWithContext(context.Background())
}

func (i ManagedRuleGroupOverrideResponseArgs) ToManagedRuleGroupOverrideResponseOutputWithContext(ctx context.Context) ManagedRuleGroupOverrideResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleGroupOverrideResponseOutput)
}





type ManagedRuleGroupOverrideResponseArrayInput interface {
	pulumi.Input

	ToManagedRuleGroupOverrideResponseArrayOutput() ManagedRuleGroupOverrideResponseArrayOutput
	ToManagedRuleGroupOverrideResponseArrayOutputWithContext(context.Context) ManagedRuleGroupOverrideResponseArrayOutput
}

type ManagedRuleGroupOverrideResponseArray []ManagedRuleGroupOverrideResponseInput

func (ManagedRuleGroupOverrideResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleGroupOverrideResponse)(nil)).Elem()
}

func (i ManagedRuleGroupOverrideResponseArray) ToManagedRuleGroupOverrideResponseArrayOutput() ManagedRuleGroupOverrideResponseArrayOutput {
	return i.ToManagedRuleGroupOverrideResponseArrayOutputWithContext(context.Background())
}

func (i ManagedRuleGroupOverrideResponseArray) ToManagedRuleGroupOverrideResponseArrayOutputWithContext(ctx context.Context) ManagedRuleGroupOverrideResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleGroupOverrideResponseArrayOutput)
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





type ManagedRuleOverrideResponseInput interface {
	pulumi.Input

	ToManagedRuleOverrideResponseOutput() ManagedRuleOverrideResponseOutput
	ToManagedRuleOverrideResponseOutputWithContext(context.Context) ManagedRuleOverrideResponseOutput
}

type ManagedRuleOverrideResponseArgs struct {
	Action       pulumi.StringPtrInput `pulumi:"action"`
	EnabledState pulumi.StringPtrInput `pulumi:"enabledState"`
	RuleId       pulumi.StringInput    `pulumi:"ruleId"`
}

func (ManagedRuleOverrideResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleOverrideResponse)(nil)).Elem()
}

func (i ManagedRuleOverrideResponseArgs) ToManagedRuleOverrideResponseOutput() ManagedRuleOverrideResponseOutput {
	return i.ToManagedRuleOverrideResponseOutputWithContext(context.Background())
}

func (i ManagedRuleOverrideResponseArgs) ToManagedRuleOverrideResponseOutputWithContext(ctx context.Context) ManagedRuleOverrideResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleOverrideResponseOutput)
}





type ManagedRuleOverrideResponseArrayInput interface {
	pulumi.Input

	ToManagedRuleOverrideResponseArrayOutput() ManagedRuleOverrideResponseArrayOutput
	ToManagedRuleOverrideResponseArrayOutputWithContext(context.Context) ManagedRuleOverrideResponseArrayOutput
}

type ManagedRuleOverrideResponseArray []ManagedRuleOverrideResponseInput

func (ManagedRuleOverrideResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleOverrideResponse)(nil)).Elem()
}

func (i ManagedRuleOverrideResponseArray) ToManagedRuleOverrideResponseArrayOutput() ManagedRuleOverrideResponseArrayOutput {
	return i.ToManagedRuleOverrideResponseArrayOutputWithContext(context.Background())
}

func (i ManagedRuleOverrideResponseArray) ToManagedRuleOverrideResponseArrayOutputWithContext(ctx context.Context) ManagedRuleOverrideResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleOverrideResponseArrayOutput)
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





type ManagedRuleSetListResponseInput interface {
	pulumi.Input

	ToManagedRuleSetListResponseOutput() ManagedRuleSetListResponseOutput
	ToManagedRuleSetListResponseOutputWithContext(context.Context) ManagedRuleSetListResponseOutput
}

type ManagedRuleSetListResponseArgs struct {
	ManagedRuleSets ManagedRuleSetResponseArrayInput `pulumi:"managedRuleSets"`
}

func (ManagedRuleSetListResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetListResponse)(nil)).Elem()
}

func (i ManagedRuleSetListResponseArgs) ToManagedRuleSetListResponseOutput() ManagedRuleSetListResponseOutput {
	return i.ToManagedRuleSetListResponseOutputWithContext(context.Background())
}

func (i ManagedRuleSetListResponseArgs) ToManagedRuleSetListResponseOutputWithContext(ctx context.Context) ManagedRuleSetListResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListResponseOutput)
}

func (i ManagedRuleSetListResponseArgs) ToManagedRuleSetListResponsePtrOutput() ManagedRuleSetListResponsePtrOutput {
	return i.ToManagedRuleSetListResponsePtrOutputWithContext(context.Background())
}

func (i ManagedRuleSetListResponseArgs) ToManagedRuleSetListResponsePtrOutputWithContext(ctx context.Context) ManagedRuleSetListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListResponseOutput).ToManagedRuleSetListResponsePtrOutputWithContext(ctx)
}









type ManagedRuleSetListResponsePtrInput interface {
	pulumi.Input

	ToManagedRuleSetListResponsePtrOutput() ManagedRuleSetListResponsePtrOutput
	ToManagedRuleSetListResponsePtrOutputWithContext(context.Context) ManagedRuleSetListResponsePtrOutput
}

type managedRuleSetListResponsePtrType ManagedRuleSetListResponseArgs

func ManagedRuleSetListResponsePtr(v *ManagedRuleSetListResponseArgs) ManagedRuleSetListResponsePtrInput {
	return (*managedRuleSetListResponsePtrType)(v)
}

func (*managedRuleSetListResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRuleSetListResponse)(nil)).Elem()
}

func (i *managedRuleSetListResponsePtrType) ToManagedRuleSetListResponsePtrOutput() ManagedRuleSetListResponsePtrOutput {
	return i.ToManagedRuleSetListResponsePtrOutputWithContext(context.Background())
}

func (i *managedRuleSetListResponsePtrType) ToManagedRuleSetListResponsePtrOutputWithContext(ctx context.Context) ManagedRuleSetListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListResponsePtrOutput)
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

func (o ManagedRuleSetListResponseOutput) ToManagedRuleSetListResponsePtrOutput() ManagedRuleSetListResponsePtrOutput {
	return o.ToManagedRuleSetListResponsePtrOutputWithContext(context.Background())
}

func (o ManagedRuleSetListResponseOutput) ToManagedRuleSetListResponsePtrOutputWithContext(ctx context.Context) ManagedRuleSetListResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedRuleSetListResponse) *ManagedRuleSetListResponse {
		return &v
	}).(ManagedRuleSetListResponsePtrOutput)
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





type ManagedRuleSetResponseInput interface {
	pulumi.Input

	ToManagedRuleSetResponseOutput() ManagedRuleSetResponseOutput
	ToManagedRuleSetResponseOutputWithContext(context.Context) ManagedRuleSetResponseOutput
}

type ManagedRuleSetResponseArgs struct {
	AnomalyScore       pulumi.IntPtrInput                         `pulumi:"anomalyScore"`
	RuleGroupOverrides ManagedRuleGroupOverrideResponseArrayInput `pulumi:"ruleGroupOverrides"`
	RuleSetType        pulumi.StringInput                         `pulumi:"ruleSetType"`
	RuleSetVersion     pulumi.StringInput                         `pulumi:"ruleSetVersion"`
}

func (ManagedRuleSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetResponse)(nil)).Elem()
}

func (i ManagedRuleSetResponseArgs) ToManagedRuleSetResponseOutput() ManagedRuleSetResponseOutput {
	return i.ToManagedRuleSetResponseOutputWithContext(context.Background())
}

func (i ManagedRuleSetResponseArgs) ToManagedRuleSetResponseOutputWithContext(ctx context.Context) ManagedRuleSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetResponseOutput)
}





type ManagedRuleSetResponseArrayInput interface {
	pulumi.Input

	ToManagedRuleSetResponseArrayOutput() ManagedRuleSetResponseArrayOutput
	ToManagedRuleSetResponseArrayOutputWithContext(context.Context) ManagedRuleSetResponseArrayOutput
}

type ManagedRuleSetResponseArray []ManagedRuleSetResponseInput

func (ManagedRuleSetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleSetResponse)(nil)).Elem()
}

func (i ManagedRuleSetResponseArray) ToManagedRuleSetResponseArrayOutput() ManagedRuleSetResponseArrayOutput {
	return i.ToManagedRuleSetResponseArrayOutputWithContext(context.Background())
}

func (i ManagedRuleSetResponseArray) ToManagedRuleSetResponseArrayOutputWithContext(ctx context.Context) ManagedRuleSetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetResponseArrayOutput)
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





type MatchConditionResponseInput interface {
	pulumi.Input

	ToMatchConditionResponseOutput() MatchConditionResponseOutput
	ToMatchConditionResponseOutputWithContext(context.Context) MatchConditionResponseOutput
}

type MatchConditionResponseArgs struct {
	MatchValue      pulumi.StringArrayInput `pulumi:"matchValue"`
	MatchVariable   pulumi.StringInput      `pulumi:"matchVariable"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Selector        pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (MatchConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchConditionResponse)(nil)).Elem()
}

func (i MatchConditionResponseArgs) ToMatchConditionResponseOutput() MatchConditionResponseOutput {
	return i.ToMatchConditionResponseOutputWithContext(context.Background())
}

func (i MatchConditionResponseArgs) ToMatchConditionResponseOutputWithContext(ctx context.Context) MatchConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MatchConditionResponseOutput)
}





type MatchConditionResponseArrayInput interface {
	pulumi.Input

	ToMatchConditionResponseArrayOutput() MatchConditionResponseArrayOutput
	ToMatchConditionResponseArrayOutputWithContext(context.Context) MatchConditionResponseArrayOutput
}

type MatchConditionResponseArray []MatchConditionResponseInput

func (MatchConditionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MatchConditionResponse)(nil)).Elem()
}

func (i MatchConditionResponseArray) ToMatchConditionResponseArrayOutput() MatchConditionResponseArrayOutput {
	return i.ToMatchConditionResponseArrayOutputWithContext(context.Background())
}

func (i MatchConditionResponseArray) ToMatchConditionResponseArrayOutputWithContext(ctx context.Context) MatchConditionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MatchConditionResponseArrayOutput)
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





type PolicySettingsResponseInput interface {
	pulumi.Input

	ToPolicySettingsResponseOutput() PolicySettingsResponseOutput
	ToPolicySettingsResponseOutputWithContext(context.Context) PolicySettingsResponseOutput
}

type PolicySettingsResponseArgs struct {
	DefaultCustomBlockResponseBody       pulumi.StringPtrInput `pulumi:"defaultCustomBlockResponseBody"`
	DefaultCustomBlockResponseStatusCode pulumi.IntPtrInput    `pulumi:"defaultCustomBlockResponseStatusCode"`
	DefaultRedirectUrl                   pulumi.StringPtrInput `pulumi:"defaultRedirectUrl"`
	EnabledState                         pulumi.StringPtrInput `pulumi:"enabledState"`
	Mode                                 pulumi.StringPtrInput `pulumi:"mode"`
}

func (PolicySettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySettingsResponse)(nil)).Elem()
}

func (i PolicySettingsResponseArgs) ToPolicySettingsResponseOutput() PolicySettingsResponseOutput {
	return i.ToPolicySettingsResponseOutputWithContext(context.Background())
}

func (i PolicySettingsResponseArgs) ToPolicySettingsResponseOutputWithContext(ctx context.Context) PolicySettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySettingsResponseOutput)
}

func (i PolicySettingsResponseArgs) ToPolicySettingsResponsePtrOutput() PolicySettingsResponsePtrOutput {
	return i.ToPolicySettingsResponsePtrOutputWithContext(context.Background())
}

func (i PolicySettingsResponseArgs) ToPolicySettingsResponsePtrOutputWithContext(ctx context.Context) PolicySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySettingsResponseOutput).ToPolicySettingsResponsePtrOutputWithContext(ctx)
}









type PolicySettingsResponsePtrInput interface {
	pulumi.Input

	ToPolicySettingsResponsePtrOutput() PolicySettingsResponsePtrOutput
	ToPolicySettingsResponsePtrOutputWithContext(context.Context) PolicySettingsResponsePtrOutput
}

type policySettingsResponsePtrType PolicySettingsResponseArgs

func PolicySettingsResponsePtr(v *PolicySettingsResponseArgs) PolicySettingsResponsePtrInput {
	return (*policySettingsResponsePtrType)(v)
}

func (*policySettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySettingsResponse)(nil)).Elem()
}

func (i *policySettingsResponsePtrType) ToPolicySettingsResponsePtrOutput() PolicySettingsResponsePtrOutput {
	return i.ToPolicySettingsResponsePtrOutputWithContext(context.Background())
}

func (i *policySettingsResponsePtrType) ToPolicySettingsResponsePtrOutputWithContext(ctx context.Context) PolicySettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySettingsResponsePtrOutput)
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

func (o PolicySettingsResponseOutput) ToPolicySettingsResponsePtrOutput() PolicySettingsResponsePtrOutput {
	return o.ToPolicySettingsResponsePtrOutputWithContext(context.Background())
}

func (o PolicySettingsResponseOutput) ToPolicySettingsResponsePtrOutputWithContext(ctx context.Context) PolicySettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicySettingsResponse) *PolicySettingsResponse {
		return &v
	}).(PolicySettingsResponsePtrOutput)
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





type PostArgsMatchConditionParametersInput interface {
	pulumi.Input

	ToPostArgsMatchConditionParametersOutput() PostArgsMatchConditionParametersOutput
	ToPostArgsMatchConditionParametersOutputWithContext(context.Context) PostArgsMatchConditionParametersOutput
}

type PostArgsMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Selector        pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (PostArgsMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PostArgsMatchConditionParameters)(nil)).Elem()
}

func (i PostArgsMatchConditionParametersArgs) ToPostArgsMatchConditionParametersOutput() PostArgsMatchConditionParametersOutput {
	return i.ToPostArgsMatchConditionParametersOutputWithContext(context.Background())
}

func (i PostArgsMatchConditionParametersArgs) ToPostArgsMatchConditionParametersOutputWithContext(ctx context.Context) PostArgsMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostArgsMatchConditionParametersOutput)
}

type PostArgsMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (PostArgsMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PostArgsMatchConditionParameters)(nil)).Elem()
}

func (o PostArgsMatchConditionParametersOutput) ToPostArgsMatchConditionParametersOutput() PostArgsMatchConditionParametersOutput {
	return o
}

func (o PostArgsMatchConditionParametersOutput) ToPostArgsMatchConditionParametersOutputWithContext(ctx context.Context) PostArgsMatchConditionParametersOutput {
	return o
}

func (o PostArgsMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o PostArgsMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o PostArgsMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o PostArgsMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

func (o PostArgsMatchConditionParametersOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParameters) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o PostArgsMatchConditionParametersOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParameters) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type PostArgsMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}





type PostArgsMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToPostArgsMatchConditionParametersResponseOutput() PostArgsMatchConditionParametersResponseOutput
	ToPostArgsMatchConditionParametersResponseOutputWithContext(context.Context) PostArgsMatchConditionParametersResponseOutput
}

type PostArgsMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Selector        pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (PostArgsMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PostArgsMatchConditionParametersResponse)(nil)).Elem()
}

func (i PostArgsMatchConditionParametersResponseArgs) ToPostArgsMatchConditionParametersResponseOutput() PostArgsMatchConditionParametersResponseOutput {
	return i.ToPostArgsMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i PostArgsMatchConditionParametersResponseArgs) ToPostArgsMatchConditionParametersResponseOutputWithContext(ctx context.Context) PostArgsMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostArgsMatchConditionParametersResponseOutput)
}

type PostArgsMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (PostArgsMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PostArgsMatchConditionParametersResponse)(nil)).Elem()
}

func (o PostArgsMatchConditionParametersResponseOutput) ToPostArgsMatchConditionParametersResponseOutput() PostArgsMatchConditionParametersResponseOutput {
	return o
}

func (o PostArgsMatchConditionParametersResponseOutput) ToPostArgsMatchConditionParametersResponseOutputWithContext(ctx context.Context) PostArgsMatchConditionParametersResponseOutput {
	return o
}

func (o PostArgsMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o PostArgsMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o PostArgsMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o PostArgsMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o PostArgsMatchConditionParametersResponseOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParametersResponse) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o PostArgsMatchConditionParametersResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PostArgsMatchConditionParametersResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type QueryStringMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type QueryStringMatchConditionParametersInput interface {
	pulumi.Input

	ToQueryStringMatchConditionParametersOutput() QueryStringMatchConditionParametersOutput
	ToQueryStringMatchConditionParametersOutputWithContext(context.Context) QueryStringMatchConditionParametersOutput
}

type QueryStringMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (QueryStringMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryStringMatchConditionParameters)(nil)).Elem()
}

func (i QueryStringMatchConditionParametersArgs) ToQueryStringMatchConditionParametersOutput() QueryStringMatchConditionParametersOutput {
	return i.ToQueryStringMatchConditionParametersOutputWithContext(context.Background())
}

func (i QueryStringMatchConditionParametersArgs) ToQueryStringMatchConditionParametersOutputWithContext(ctx context.Context) QueryStringMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryStringMatchConditionParametersOutput)
}

type QueryStringMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (QueryStringMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryStringMatchConditionParameters)(nil)).Elem()
}

func (o QueryStringMatchConditionParametersOutput) ToQueryStringMatchConditionParametersOutput() QueryStringMatchConditionParametersOutput {
	return o
}

func (o QueryStringMatchConditionParametersOutput) ToQueryStringMatchConditionParametersOutputWithContext(ctx context.Context) QueryStringMatchConditionParametersOutput {
	return o
}

func (o QueryStringMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryStringMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o QueryStringMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v QueryStringMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o QueryStringMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v QueryStringMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o QueryStringMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v QueryStringMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

func (o QueryStringMatchConditionParametersOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryStringMatchConditionParameters) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type QueryStringMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type QueryStringMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToQueryStringMatchConditionParametersResponseOutput() QueryStringMatchConditionParametersResponseOutput
	ToQueryStringMatchConditionParametersResponseOutputWithContext(context.Context) QueryStringMatchConditionParametersResponseOutput
}

type QueryStringMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (QueryStringMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryStringMatchConditionParametersResponse)(nil)).Elem()
}

func (i QueryStringMatchConditionParametersResponseArgs) ToQueryStringMatchConditionParametersResponseOutput() QueryStringMatchConditionParametersResponseOutput {
	return i.ToQueryStringMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i QueryStringMatchConditionParametersResponseArgs) ToQueryStringMatchConditionParametersResponseOutputWithContext(ctx context.Context) QueryStringMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryStringMatchConditionParametersResponseOutput)
}

type QueryStringMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (QueryStringMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryStringMatchConditionParametersResponse)(nil)).Elem()
}

func (o QueryStringMatchConditionParametersResponseOutput) ToQueryStringMatchConditionParametersResponseOutput() QueryStringMatchConditionParametersResponseOutput {
	return o
}

func (o QueryStringMatchConditionParametersResponseOutput) ToQueryStringMatchConditionParametersResponseOutputWithContext(ctx context.Context) QueryStringMatchConditionParametersResponseOutput {
	return o
}

func (o QueryStringMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryStringMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o QueryStringMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v QueryStringMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o QueryStringMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v QueryStringMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o QueryStringMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v QueryStringMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o QueryStringMatchConditionParametersResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryStringMatchConditionParametersResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
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





type RateLimitRuleListResponseInput interface {
	pulumi.Input

	ToRateLimitRuleListResponseOutput() RateLimitRuleListResponseOutput
	ToRateLimitRuleListResponseOutputWithContext(context.Context) RateLimitRuleListResponseOutput
}

type RateLimitRuleListResponseArgs struct {
	Rules RateLimitRuleResponseArrayInput `pulumi:"rules"`
}

func (RateLimitRuleListResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RateLimitRuleListResponse)(nil)).Elem()
}

func (i RateLimitRuleListResponseArgs) ToRateLimitRuleListResponseOutput() RateLimitRuleListResponseOutput {
	return i.ToRateLimitRuleListResponseOutputWithContext(context.Background())
}

func (i RateLimitRuleListResponseArgs) ToRateLimitRuleListResponseOutputWithContext(ctx context.Context) RateLimitRuleListResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateLimitRuleListResponseOutput)
}

func (i RateLimitRuleListResponseArgs) ToRateLimitRuleListResponsePtrOutput() RateLimitRuleListResponsePtrOutput {
	return i.ToRateLimitRuleListResponsePtrOutputWithContext(context.Background())
}

func (i RateLimitRuleListResponseArgs) ToRateLimitRuleListResponsePtrOutputWithContext(ctx context.Context) RateLimitRuleListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateLimitRuleListResponseOutput).ToRateLimitRuleListResponsePtrOutputWithContext(ctx)
}









type RateLimitRuleListResponsePtrInput interface {
	pulumi.Input

	ToRateLimitRuleListResponsePtrOutput() RateLimitRuleListResponsePtrOutput
	ToRateLimitRuleListResponsePtrOutputWithContext(context.Context) RateLimitRuleListResponsePtrOutput
}

type rateLimitRuleListResponsePtrType RateLimitRuleListResponseArgs

func RateLimitRuleListResponsePtr(v *RateLimitRuleListResponseArgs) RateLimitRuleListResponsePtrInput {
	return (*rateLimitRuleListResponsePtrType)(v)
}

func (*rateLimitRuleListResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RateLimitRuleListResponse)(nil)).Elem()
}

func (i *rateLimitRuleListResponsePtrType) ToRateLimitRuleListResponsePtrOutput() RateLimitRuleListResponsePtrOutput {
	return i.ToRateLimitRuleListResponsePtrOutputWithContext(context.Background())
}

func (i *rateLimitRuleListResponsePtrType) ToRateLimitRuleListResponsePtrOutputWithContext(ctx context.Context) RateLimitRuleListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateLimitRuleListResponsePtrOutput)
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

func (o RateLimitRuleListResponseOutput) ToRateLimitRuleListResponsePtrOutput() RateLimitRuleListResponsePtrOutput {
	return o.ToRateLimitRuleListResponsePtrOutputWithContext(context.Background())
}

func (o RateLimitRuleListResponseOutput) ToRateLimitRuleListResponsePtrOutputWithContext(ctx context.Context) RateLimitRuleListResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RateLimitRuleListResponse) *RateLimitRuleListResponse {
		return &v
	}).(RateLimitRuleListResponsePtrOutput)
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





type RateLimitRuleResponseInput interface {
	pulumi.Input

	ToRateLimitRuleResponseOutput() RateLimitRuleResponseOutput
	ToRateLimitRuleResponseOutputWithContext(context.Context) RateLimitRuleResponseOutput
}

type RateLimitRuleResponseArgs struct {
	Action                     pulumi.StringInput               `pulumi:"action"`
	EnabledState               pulumi.StringPtrInput            `pulumi:"enabledState"`
	MatchConditions            MatchConditionResponseArrayInput `pulumi:"matchConditions"`
	Name                       pulumi.StringInput               `pulumi:"name"`
	Priority                   pulumi.IntInput                  `pulumi:"priority"`
	RateLimitDurationInMinutes pulumi.IntInput                  `pulumi:"rateLimitDurationInMinutes"`
	RateLimitThreshold         pulumi.IntInput                  `pulumi:"rateLimitThreshold"`
}

func (RateLimitRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RateLimitRuleResponse)(nil)).Elem()
}

func (i RateLimitRuleResponseArgs) ToRateLimitRuleResponseOutput() RateLimitRuleResponseOutput {
	return i.ToRateLimitRuleResponseOutputWithContext(context.Background())
}

func (i RateLimitRuleResponseArgs) ToRateLimitRuleResponseOutputWithContext(ctx context.Context) RateLimitRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateLimitRuleResponseOutput)
}





type RateLimitRuleResponseArrayInput interface {
	pulumi.Input

	ToRateLimitRuleResponseArrayOutput() RateLimitRuleResponseArrayOutput
	ToRateLimitRuleResponseArrayOutputWithContext(context.Context) RateLimitRuleResponseArrayOutput
}

type RateLimitRuleResponseArray []RateLimitRuleResponseInput

func (RateLimitRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RateLimitRuleResponse)(nil)).Elem()
}

func (i RateLimitRuleResponseArray) ToRateLimitRuleResponseArrayOutput() RateLimitRuleResponseArrayOutput {
	return i.ToRateLimitRuleResponseArrayOutputWithContext(context.Background())
}

func (i RateLimitRuleResponseArray) ToRateLimitRuleResponseArrayOutputWithContext(ctx context.Context) RateLimitRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateLimitRuleResponseArrayOutput)
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





type RemoteAddressMatchConditionParametersInput interface {
	pulumi.Input

	ToRemoteAddressMatchConditionParametersOutput() RemoteAddressMatchConditionParametersOutput
	ToRemoteAddressMatchConditionParametersOutputWithContext(context.Context) RemoteAddressMatchConditionParametersOutput
}

type RemoteAddressMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (RemoteAddressMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteAddressMatchConditionParameters)(nil)).Elem()
}

func (i RemoteAddressMatchConditionParametersArgs) ToRemoteAddressMatchConditionParametersOutput() RemoteAddressMatchConditionParametersOutput {
	return i.ToRemoteAddressMatchConditionParametersOutputWithContext(context.Background())
}

func (i RemoteAddressMatchConditionParametersArgs) ToRemoteAddressMatchConditionParametersOutputWithContext(ctx context.Context) RemoteAddressMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteAddressMatchConditionParametersOutput)
}

type RemoteAddressMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (RemoteAddressMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteAddressMatchConditionParameters)(nil)).Elem()
}

func (o RemoteAddressMatchConditionParametersOutput) ToRemoteAddressMatchConditionParametersOutput() RemoteAddressMatchConditionParametersOutput {
	return o
}

func (o RemoteAddressMatchConditionParametersOutput) ToRemoteAddressMatchConditionParametersOutputWithContext(ctx context.Context) RemoteAddressMatchConditionParametersOutput {
	return o
}

func (o RemoteAddressMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemoteAddressMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RemoteAddressMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RemoteAddressMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RemoteAddressMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RemoteAddressMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RemoteAddressMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RemoteAddressMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

func (o RemoteAddressMatchConditionParametersOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemoteAddressMatchConditionParameters) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RemoteAddressMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type RemoteAddressMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToRemoteAddressMatchConditionParametersResponseOutput() RemoteAddressMatchConditionParametersResponseOutput
	ToRemoteAddressMatchConditionParametersResponseOutputWithContext(context.Context) RemoteAddressMatchConditionParametersResponseOutput
}

type RemoteAddressMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (RemoteAddressMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteAddressMatchConditionParametersResponse)(nil)).Elem()
}

func (i RemoteAddressMatchConditionParametersResponseArgs) ToRemoteAddressMatchConditionParametersResponseOutput() RemoteAddressMatchConditionParametersResponseOutput {
	return i.ToRemoteAddressMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i RemoteAddressMatchConditionParametersResponseArgs) ToRemoteAddressMatchConditionParametersResponseOutputWithContext(ctx context.Context) RemoteAddressMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteAddressMatchConditionParametersResponseOutput)
}

type RemoteAddressMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (RemoteAddressMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteAddressMatchConditionParametersResponse)(nil)).Elem()
}

func (o RemoteAddressMatchConditionParametersResponseOutput) ToRemoteAddressMatchConditionParametersResponseOutput() RemoteAddressMatchConditionParametersResponseOutput {
	return o
}

func (o RemoteAddressMatchConditionParametersResponseOutput) ToRemoteAddressMatchConditionParametersResponseOutputWithContext(ctx context.Context) RemoteAddressMatchConditionParametersResponseOutput {
	return o
}

func (o RemoteAddressMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemoteAddressMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RemoteAddressMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RemoteAddressMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RemoteAddressMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RemoteAddressMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RemoteAddressMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RemoteAddressMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o RemoteAddressMatchConditionParametersResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemoteAddressMatchConditionParametersResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RequestBodyMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type RequestBodyMatchConditionParametersInput interface {
	pulumi.Input

	ToRequestBodyMatchConditionParametersOutput() RequestBodyMatchConditionParametersOutput
	ToRequestBodyMatchConditionParametersOutputWithContext(context.Context) RequestBodyMatchConditionParametersOutput
}

type RequestBodyMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (RequestBodyMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestBodyMatchConditionParameters)(nil)).Elem()
}

func (i RequestBodyMatchConditionParametersArgs) ToRequestBodyMatchConditionParametersOutput() RequestBodyMatchConditionParametersOutput {
	return i.ToRequestBodyMatchConditionParametersOutputWithContext(context.Background())
}

func (i RequestBodyMatchConditionParametersArgs) ToRequestBodyMatchConditionParametersOutputWithContext(ctx context.Context) RequestBodyMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestBodyMatchConditionParametersOutput)
}

type RequestBodyMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (RequestBodyMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestBodyMatchConditionParameters)(nil)).Elem()
}

func (o RequestBodyMatchConditionParametersOutput) ToRequestBodyMatchConditionParametersOutput() RequestBodyMatchConditionParametersOutput {
	return o
}

func (o RequestBodyMatchConditionParametersOutput) ToRequestBodyMatchConditionParametersOutputWithContext(ctx context.Context) RequestBodyMatchConditionParametersOutput {
	return o
}

func (o RequestBodyMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestBodyMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RequestBodyMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequestBodyMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RequestBodyMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RequestBodyMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RequestBodyMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RequestBodyMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

func (o RequestBodyMatchConditionParametersOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestBodyMatchConditionParameters) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RequestBodyMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type RequestBodyMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToRequestBodyMatchConditionParametersResponseOutput() RequestBodyMatchConditionParametersResponseOutput
	ToRequestBodyMatchConditionParametersResponseOutputWithContext(context.Context) RequestBodyMatchConditionParametersResponseOutput
}

type RequestBodyMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (RequestBodyMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestBodyMatchConditionParametersResponse)(nil)).Elem()
}

func (i RequestBodyMatchConditionParametersResponseArgs) ToRequestBodyMatchConditionParametersResponseOutput() RequestBodyMatchConditionParametersResponseOutput {
	return i.ToRequestBodyMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i RequestBodyMatchConditionParametersResponseArgs) ToRequestBodyMatchConditionParametersResponseOutputWithContext(ctx context.Context) RequestBodyMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestBodyMatchConditionParametersResponseOutput)
}

type RequestBodyMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (RequestBodyMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestBodyMatchConditionParametersResponse)(nil)).Elem()
}

func (o RequestBodyMatchConditionParametersResponseOutput) ToRequestBodyMatchConditionParametersResponseOutput() RequestBodyMatchConditionParametersResponseOutput {
	return o
}

func (o RequestBodyMatchConditionParametersResponseOutput) ToRequestBodyMatchConditionParametersResponseOutputWithContext(ctx context.Context) RequestBodyMatchConditionParametersResponseOutput {
	return o
}

func (o RequestBodyMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestBodyMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RequestBodyMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequestBodyMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RequestBodyMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RequestBodyMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RequestBodyMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RequestBodyMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o RequestBodyMatchConditionParametersResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestBodyMatchConditionParametersResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RequestHeaderMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}





type RequestHeaderMatchConditionParametersInput interface {
	pulumi.Input

	ToRequestHeaderMatchConditionParametersOutput() RequestHeaderMatchConditionParametersOutput
	ToRequestHeaderMatchConditionParametersOutputWithContext(context.Context) RequestHeaderMatchConditionParametersOutput
}

type RequestHeaderMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Selector        pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (RequestHeaderMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestHeaderMatchConditionParameters)(nil)).Elem()
}

func (i RequestHeaderMatchConditionParametersArgs) ToRequestHeaderMatchConditionParametersOutput() RequestHeaderMatchConditionParametersOutput {
	return i.ToRequestHeaderMatchConditionParametersOutputWithContext(context.Background())
}

func (i RequestHeaderMatchConditionParametersArgs) ToRequestHeaderMatchConditionParametersOutputWithContext(ctx context.Context) RequestHeaderMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestHeaderMatchConditionParametersOutput)
}

type RequestHeaderMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (RequestHeaderMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestHeaderMatchConditionParameters)(nil)).Elem()
}

func (o RequestHeaderMatchConditionParametersOutput) ToRequestHeaderMatchConditionParametersOutput() RequestHeaderMatchConditionParametersOutput {
	return o
}

func (o RequestHeaderMatchConditionParametersOutput) ToRequestHeaderMatchConditionParametersOutputWithContext(ctx context.Context) RequestHeaderMatchConditionParametersOutput {
	return o
}

func (o RequestHeaderMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RequestHeaderMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RequestHeaderMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RequestHeaderMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

func (o RequestHeaderMatchConditionParametersOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParameters) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o RequestHeaderMatchConditionParametersOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParameters) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RequestHeaderMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}





type RequestHeaderMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToRequestHeaderMatchConditionParametersResponseOutput() RequestHeaderMatchConditionParametersResponseOutput
	ToRequestHeaderMatchConditionParametersResponseOutputWithContext(context.Context) RequestHeaderMatchConditionParametersResponseOutput
}

type RequestHeaderMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Selector        pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (RequestHeaderMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestHeaderMatchConditionParametersResponse)(nil)).Elem()
}

func (i RequestHeaderMatchConditionParametersResponseArgs) ToRequestHeaderMatchConditionParametersResponseOutput() RequestHeaderMatchConditionParametersResponseOutput {
	return i.ToRequestHeaderMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i RequestHeaderMatchConditionParametersResponseArgs) ToRequestHeaderMatchConditionParametersResponseOutputWithContext(ctx context.Context) RequestHeaderMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestHeaderMatchConditionParametersResponseOutput)
}

type RequestHeaderMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (RequestHeaderMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestHeaderMatchConditionParametersResponse)(nil)).Elem()
}

func (o RequestHeaderMatchConditionParametersResponseOutput) ToRequestHeaderMatchConditionParametersResponseOutput() RequestHeaderMatchConditionParametersResponseOutput {
	return o
}

func (o RequestHeaderMatchConditionParametersResponseOutput) ToRequestHeaderMatchConditionParametersResponseOutputWithContext(ctx context.Context) RequestHeaderMatchConditionParametersResponseOutput {
	return o
}

func (o RequestHeaderMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RequestHeaderMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RequestHeaderMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RequestHeaderMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o RequestHeaderMatchConditionParametersResponseOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParametersResponse) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o RequestHeaderMatchConditionParametersResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestHeaderMatchConditionParametersResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RequestMethodMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
}





type RequestMethodMatchConditionParametersInput interface {
	pulumi.Input

	ToRequestMethodMatchConditionParametersOutput() RequestMethodMatchConditionParametersOutput
	ToRequestMethodMatchConditionParametersOutputWithContext(context.Context) RequestMethodMatchConditionParametersOutput
}

type RequestMethodMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
}

func (RequestMethodMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestMethodMatchConditionParameters)(nil)).Elem()
}

func (i RequestMethodMatchConditionParametersArgs) ToRequestMethodMatchConditionParametersOutput() RequestMethodMatchConditionParametersOutput {
	return i.ToRequestMethodMatchConditionParametersOutputWithContext(context.Background())
}

func (i RequestMethodMatchConditionParametersArgs) ToRequestMethodMatchConditionParametersOutputWithContext(ctx context.Context) RequestMethodMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestMethodMatchConditionParametersOutput)
}

type RequestMethodMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (RequestMethodMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestMethodMatchConditionParameters)(nil)).Elem()
}

func (o RequestMethodMatchConditionParametersOutput) ToRequestMethodMatchConditionParametersOutput() RequestMethodMatchConditionParametersOutput {
	return o
}

func (o RequestMethodMatchConditionParametersOutput) ToRequestMethodMatchConditionParametersOutputWithContext(ctx context.Context) RequestMethodMatchConditionParametersOutput {
	return o
}

func (o RequestMethodMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestMethodMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RequestMethodMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequestMethodMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RequestMethodMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RequestMethodMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RequestMethodMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RequestMethodMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

type RequestMethodMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
}





type RequestMethodMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToRequestMethodMatchConditionParametersResponseOutput() RequestMethodMatchConditionParametersResponseOutput
	ToRequestMethodMatchConditionParametersResponseOutputWithContext(context.Context) RequestMethodMatchConditionParametersResponseOutput
}

type RequestMethodMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
}

func (RequestMethodMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestMethodMatchConditionParametersResponse)(nil)).Elem()
}

func (i RequestMethodMatchConditionParametersResponseArgs) ToRequestMethodMatchConditionParametersResponseOutput() RequestMethodMatchConditionParametersResponseOutput {
	return i.ToRequestMethodMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i RequestMethodMatchConditionParametersResponseArgs) ToRequestMethodMatchConditionParametersResponseOutputWithContext(ctx context.Context) RequestMethodMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestMethodMatchConditionParametersResponseOutput)
}

type RequestMethodMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (RequestMethodMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestMethodMatchConditionParametersResponse)(nil)).Elem()
}

func (o RequestMethodMatchConditionParametersResponseOutput) ToRequestMethodMatchConditionParametersResponseOutput() RequestMethodMatchConditionParametersResponseOutput {
	return o
}

func (o RequestMethodMatchConditionParametersResponseOutput) ToRequestMethodMatchConditionParametersResponseOutputWithContext(ctx context.Context) RequestMethodMatchConditionParametersResponseOutput {
	return o
}

func (o RequestMethodMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestMethodMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RequestMethodMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequestMethodMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RequestMethodMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RequestMethodMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RequestMethodMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RequestMethodMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

type RequestSchemeMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
}





type RequestSchemeMatchConditionParametersInput interface {
	pulumi.Input

	ToRequestSchemeMatchConditionParametersOutput() RequestSchemeMatchConditionParametersOutput
	ToRequestSchemeMatchConditionParametersOutputWithContext(context.Context) RequestSchemeMatchConditionParametersOutput
}

type RequestSchemeMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
}

func (RequestSchemeMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestSchemeMatchConditionParameters)(nil)).Elem()
}

func (i RequestSchemeMatchConditionParametersArgs) ToRequestSchemeMatchConditionParametersOutput() RequestSchemeMatchConditionParametersOutput {
	return i.ToRequestSchemeMatchConditionParametersOutputWithContext(context.Background())
}

func (i RequestSchemeMatchConditionParametersArgs) ToRequestSchemeMatchConditionParametersOutputWithContext(ctx context.Context) RequestSchemeMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestSchemeMatchConditionParametersOutput)
}

type RequestSchemeMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (RequestSchemeMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestSchemeMatchConditionParameters)(nil)).Elem()
}

func (o RequestSchemeMatchConditionParametersOutput) ToRequestSchemeMatchConditionParametersOutput() RequestSchemeMatchConditionParametersOutput {
	return o
}

func (o RequestSchemeMatchConditionParametersOutput) ToRequestSchemeMatchConditionParametersOutputWithContext(ctx context.Context) RequestSchemeMatchConditionParametersOutput {
	return o
}

func (o RequestSchemeMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestSchemeMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RequestSchemeMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequestSchemeMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RequestSchemeMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RequestSchemeMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RequestSchemeMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RequestSchemeMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

type RequestSchemeMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
}





type RequestSchemeMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToRequestSchemeMatchConditionParametersResponseOutput() RequestSchemeMatchConditionParametersResponseOutput
	ToRequestSchemeMatchConditionParametersResponseOutputWithContext(context.Context) RequestSchemeMatchConditionParametersResponseOutput
}

type RequestSchemeMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
}

func (RequestSchemeMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestSchemeMatchConditionParametersResponse)(nil)).Elem()
}

func (i RequestSchemeMatchConditionParametersResponseArgs) ToRequestSchemeMatchConditionParametersResponseOutput() RequestSchemeMatchConditionParametersResponseOutput {
	return i.ToRequestSchemeMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i RequestSchemeMatchConditionParametersResponseArgs) ToRequestSchemeMatchConditionParametersResponseOutputWithContext(ctx context.Context) RequestSchemeMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestSchemeMatchConditionParametersResponseOutput)
}

type RequestSchemeMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (RequestSchemeMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestSchemeMatchConditionParametersResponse)(nil)).Elem()
}

func (o RequestSchemeMatchConditionParametersResponseOutput) ToRequestSchemeMatchConditionParametersResponseOutput() RequestSchemeMatchConditionParametersResponseOutput {
	return o
}

func (o RequestSchemeMatchConditionParametersResponseOutput) ToRequestSchemeMatchConditionParametersResponseOutputWithContext(ctx context.Context) RequestSchemeMatchConditionParametersResponseOutput {
	return o
}

func (o RequestSchemeMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestSchemeMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RequestSchemeMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequestSchemeMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RequestSchemeMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RequestSchemeMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RequestSchemeMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RequestSchemeMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

type RequestUriMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type RequestUriMatchConditionParametersInput interface {
	pulumi.Input

	ToRequestUriMatchConditionParametersOutput() RequestUriMatchConditionParametersOutput
	ToRequestUriMatchConditionParametersOutputWithContext(context.Context) RequestUriMatchConditionParametersOutput
}

type RequestUriMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (RequestUriMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestUriMatchConditionParameters)(nil)).Elem()
}

func (i RequestUriMatchConditionParametersArgs) ToRequestUriMatchConditionParametersOutput() RequestUriMatchConditionParametersOutput {
	return i.ToRequestUriMatchConditionParametersOutputWithContext(context.Background())
}

func (i RequestUriMatchConditionParametersArgs) ToRequestUriMatchConditionParametersOutputWithContext(ctx context.Context) RequestUriMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestUriMatchConditionParametersOutput)
}

type RequestUriMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (RequestUriMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestUriMatchConditionParameters)(nil)).Elem()
}

func (o RequestUriMatchConditionParametersOutput) ToRequestUriMatchConditionParametersOutput() RequestUriMatchConditionParametersOutput {
	return o
}

func (o RequestUriMatchConditionParametersOutput) ToRequestUriMatchConditionParametersOutputWithContext(ctx context.Context) RequestUriMatchConditionParametersOutput {
	return o
}

func (o RequestUriMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestUriMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RequestUriMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequestUriMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RequestUriMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RequestUriMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RequestUriMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RequestUriMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

func (o RequestUriMatchConditionParametersOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestUriMatchConditionParameters) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RequestUriMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type RequestUriMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToRequestUriMatchConditionParametersResponseOutput() RequestUriMatchConditionParametersResponseOutput
	ToRequestUriMatchConditionParametersResponseOutputWithContext(context.Context) RequestUriMatchConditionParametersResponseOutput
}

type RequestUriMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (RequestUriMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestUriMatchConditionParametersResponse)(nil)).Elem()
}

func (i RequestUriMatchConditionParametersResponseArgs) ToRequestUriMatchConditionParametersResponseOutput() RequestUriMatchConditionParametersResponseOutput {
	return i.ToRequestUriMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i RequestUriMatchConditionParametersResponseArgs) ToRequestUriMatchConditionParametersResponseOutputWithContext(ctx context.Context) RequestUriMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestUriMatchConditionParametersResponseOutput)
}

type RequestUriMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (RequestUriMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestUriMatchConditionParametersResponse)(nil)).Elem()
}

func (o RequestUriMatchConditionParametersResponseOutput) ToRequestUriMatchConditionParametersResponseOutput() RequestUriMatchConditionParametersResponseOutput {
	return o
}

func (o RequestUriMatchConditionParametersResponseOutput) ToRequestUriMatchConditionParametersResponseOutputWithContext(ctx context.Context) RequestUriMatchConditionParametersResponseOutput {
	return o
}

func (o RequestUriMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestUriMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o RequestUriMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequestUriMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RequestUriMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RequestUriMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RequestUriMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v RequestUriMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o RequestUriMatchConditionParametersResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequestUriMatchConditionParametersResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
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





type ResourceReferenceResponseInput interface {
	pulumi.Input

	ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput
	ToResourceReferenceResponseOutputWithContext(context.Context) ResourceReferenceResponseOutput
}

type ResourceReferenceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return i.ToResourceReferenceResponseOutputWithContext(context.Background())
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponseOutput)
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return i.ToResourceReferenceResponsePtrOutputWithContext(context.Background())
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponseOutput).ToResourceReferenceResponsePtrOutputWithContext(ctx)
}









type ResourceReferenceResponsePtrInput interface {
	pulumi.Input

	ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput
	ToResourceReferenceResponsePtrOutputWithContext(context.Context) ResourceReferenceResponsePtrOutput
}

type resourceReferenceResponsePtrType ResourceReferenceResponseArgs

func ResourceReferenceResponsePtr(v *ResourceReferenceResponseArgs) ResourceReferenceResponsePtrInput {
	return (*resourceReferenceResponsePtrType)(v)
}

func (*resourceReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReferenceResponse)(nil)).Elem()
}

func (i *resourceReferenceResponsePtrType) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return i.ToResourceReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *resourceReferenceResponsePtrType) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponsePtrOutput)
}





type ResourceReferenceResponseArrayInput interface {
	pulumi.Input

	ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput
	ToResourceReferenceResponseArrayOutputWithContext(context.Context) ResourceReferenceResponseArrayOutput
}

type ResourceReferenceResponseArray []ResourceReferenceResponseInput

func (ResourceReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReferenceResponse)(nil)).Elem()
}

func (i ResourceReferenceResponseArray) ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput {
	return i.ToResourceReferenceResponseArrayOutputWithContext(context.Background())
}

func (i ResourceReferenceResponseArray) ToResourceReferenceResponseArrayOutputWithContext(ctx context.Context) ResourceReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponseArrayOutput)
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

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return o.ToResourceReferenceResponsePtrOutputWithContext(context.Background())
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceReferenceResponse) *ResourceReferenceResponse {
		return &v
	}).(ResourceReferenceResponsePtrOutput)
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





type ResponseBasedOriginErrorDetectionParametersResponseInput interface {
	pulumi.Input

	ToResponseBasedOriginErrorDetectionParametersResponseOutput() ResponseBasedOriginErrorDetectionParametersResponseOutput
	ToResponseBasedOriginErrorDetectionParametersResponseOutputWithContext(context.Context) ResponseBasedOriginErrorDetectionParametersResponseOutput
}

type ResponseBasedOriginErrorDetectionParametersResponseArgs struct {
	HttpErrorRanges                          HttpErrorRangeParametersResponseArrayInput `pulumi:"httpErrorRanges"`
	ResponseBasedDetectedErrorTypes          pulumi.StringPtrInput                      `pulumi:"responseBasedDetectedErrorTypes"`
	ResponseBasedFailoverThresholdPercentage pulumi.IntPtrInput                         `pulumi:"responseBasedFailoverThresholdPercentage"`
}

func (ResponseBasedOriginErrorDetectionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseBasedOriginErrorDetectionParametersResponse)(nil)).Elem()
}

func (i ResponseBasedOriginErrorDetectionParametersResponseArgs) ToResponseBasedOriginErrorDetectionParametersResponseOutput() ResponseBasedOriginErrorDetectionParametersResponseOutput {
	return i.ToResponseBasedOriginErrorDetectionParametersResponseOutputWithContext(context.Background())
}

func (i ResponseBasedOriginErrorDetectionParametersResponseArgs) ToResponseBasedOriginErrorDetectionParametersResponseOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseBasedOriginErrorDetectionParametersResponseOutput)
}

func (i ResponseBasedOriginErrorDetectionParametersResponseArgs) ToResponseBasedOriginErrorDetectionParametersResponsePtrOutput() ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return i.ToResponseBasedOriginErrorDetectionParametersResponsePtrOutputWithContext(context.Background())
}

func (i ResponseBasedOriginErrorDetectionParametersResponseArgs) ToResponseBasedOriginErrorDetectionParametersResponsePtrOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseBasedOriginErrorDetectionParametersResponseOutput).ToResponseBasedOriginErrorDetectionParametersResponsePtrOutputWithContext(ctx)
}









type ResponseBasedOriginErrorDetectionParametersResponsePtrInput interface {
	pulumi.Input

	ToResponseBasedOriginErrorDetectionParametersResponsePtrOutput() ResponseBasedOriginErrorDetectionParametersResponsePtrOutput
	ToResponseBasedOriginErrorDetectionParametersResponsePtrOutputWithContext(context.Context) ResponseBasedOriginErrorDetectionParametersResponsePtrOutput
}

type responseBasedOriginErrorDetectionParametersResponsePtrType ResponseBasedOriginErrorDetectionParametersResponseArgs

func ResponseBasedOriginErrorDetectionParametersResponsePtr(v *ResponseBasedOriginErrorDetectionParametersResponseArgs) ResponseBasedOriginErrorDetectionParametersResponsePtrInput {
	return (*responseBasedOriginErrorDetectionParametersResponsePtrType)(v)
}

func (*responseBasedOriginErrorDetectionParametersResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseBasedOriginErrorDetectionParametersResponse)(nil)).Elem()
}

func (i *responseBasedOriginErrorDetectionParametersResponsePtrType) ToResponseBasedOriginErrorDetectionParametersResponsePtrOutput() ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return i.ToResponseBasedOriginErrorDetectionParametersResponsePtrOutputWithContext(context.Background())
}

func (i *responseBasedOriginErrorDetectionParametersResponsePtrType) ToResponseBasedOriginErrorDetectionParametersResponsePtrOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseBasedOriginErrorDetectionParametersResponsePtrOutput)
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

func (o ResponseBasedOriginErrorDetectionParametersResponseOutput) ToResponseBasedOriginErrorDetectionParametersResponsePtrOutput() ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return o.ToResponseBasedOriginErrorDetectionParametersResponsePtrOutputWithContext(context.Background())
}

func (o ResponseBasedOriginErrorDetectionParametersResponseOutput) ToResponseBasedOriginErrorDetectionParametersResponsePtrOutputWithContext(ctx context.Context) ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResponseBasedOriginErrorDetectionParametersResponse) *ResponseBasedOriginErrorDetectionParametersResponse {
		return &v
	}).(ResponseBasedOriginErrorDetectionParametersResponsePtrOutput)
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





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
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

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
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

type UrlFileExtensionMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type UrlFileExtensionMatchConditionParametersInput interface {
	pulumi.Input

	ToUrlFileExtensionMatchConditionParametersOutput() UrlFileExtensionMatchConditionParametersOutput
	ToUrlFileExtensionMatchConditionParametersOutputWithContext(context.Context) UrlFileExtensionMatchConditionParametersOutput
}

type UrlFileExtensionMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (UrlFileExtensionMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlFileExtensionMatchConditionParameters)(nil)).Elem()
}

func (i UrlFileExtensionMatchConditionParametersArgs) ToUrlFileExtensionMatchConditionParametersOutput() UrlFileExtensionMatchConditionParametersOutput {
	return i.ToUrlFileExtensionMatchConditionParametersOutputWithContext(context.Background())
}

func (i UrlFileExtensionMatchConditionParametersArgs) ToUrlFileExtensionMatchConditionParametersOutputWithContext(ctx context.Context) UrlFileExtensionMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlFileExtensionMatchConditionParametersOutput)
}

type UrlFileExtensionMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (UrlFileExtensionMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlFileExtensionMatchConditionParameters)(nil)).Elem()
}

func (o UrlFileExtensionMatchConditionParametersOutput) ToUrlFileExtensionMatchConditionParametersOutput() UrlFileExtensionMatchConditionParametersOutput {
	return o
}

func (o UrlFileExtensionMatchConditionParametersOutput) ToUrlFileExtensionMatchConditionParametersOutputWithContext(ctx context.Context) UrlFileExtensionMatchConditionParametersOutput {
	return o
}

func (o UrlFileExtensionMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlFileExtensionMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o UrlFileExtensionMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UrlFileExtensionMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o UrlFileExtensionMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlFileExtensionMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UrlFileExtensionMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v UrlFileExtensionMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

func (o UrlFileExtensionMatchConditionParametersOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlFileExtensionMatchConditionParameters) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type UrlFileExtensionMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type UrlFileExtensionMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToUrlFileExtensionMatchConditionParametersResponseOutput() UrlFileExtensionMatchConditionParametersResponseOutput
	ToUrlFileExtensionMatchConditionParametersResponseOutputWithContext(context.Context) UrlFileExtensionMatchConditionParametersResponseOutput
}

type UrlFileExtensionMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (UrlFileExtensionMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlFileExtensionMatchConditionParametersResponse)(nil)).Elem()
}

func (i UrlFileExtensionMatchConditionParametersResponseArgs) ToUrlFileExtensionMatchConditionParametersResponseOutput() UrlFileExtensionMatchConditionParametersResponseOutput {
	return i.ToUrlFileExtensionMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i UrlFileExtensionMatchConditionParametersResponseArgs) ToUrlFileExtensionMatchConditionParametersResponseOutputWithContext(ctx context.Context) UrlFileExtensionMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlFileExtensionMatchConditionParametersResponseOutput)
}

type UrlFileExtensionMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (UrlFileExtensionMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlFileExtensionMatchConditionParametersResponse)(nil)).Elem()
}

func (o UrlFileExtensionMatchConditionParametersResponseOutput) ToUrlFileExtensionMatchConditionParametersResponseOutput() UrlFileExtensionMatchConditionParametersResponseOutput {
	return o
}

func (o UrlFileExtensionMatchConditionParametersResponseOutput) ToUrlFileExtensionMatchConditionParametersResponseOutputWithContext(ctx context.Context) UrlFileExtensionMatchConditionParametersResponseOutput {
	return o
}

func (o UrlFileExtensionMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlFileExtensionMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o UrlFileExtensionMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UrlFileExtensionMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o UrlFileExtensionMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlFileExtensionMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UrlFileExtensionMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v UrlFileExtensionMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o UrlFileExtensionMatchConditionParametersResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlFileExtensionMatchConditionParametersResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type UrlFileNameMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type UrlFileNameMatchConditionParametersInput interface {
	pulumi.Input

	ToUrlFileNameMatchConditionParametersOutput() UrlFileNameMatchConditionParametersOutput
	ToUrlFileNameMatchConditionParametersOutputWithContext(context.Context) UrlFileNameMatchConditionParametersOutput
}

type UrlFileNameMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (UrlFileNameMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlFileNameMatchConditionParameters)(nil)).Elem()
}

func (i UrlFileNameMatchConditionParametersArgs) ToUrlFileNameMatchConditionParametersOutput() UrlFileNameMatchConditionParametersOutput {
	return i.ToUrlFileNameMatchConditionParametersOutputWithContext(context.Background())
}

func (i UrlFileNameMatchConditionParametersArgs) ToUrlFileNameMatchConditionParametersOutputWithContext(ctx context.Context) UrlFileNameMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlFileNameMatchConditionParametersOutput)
}

type UrlFileNameMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (UrlFileNameMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlFileNameMatchConditionParameters)(nil)).Elem()
}

func (o UrlFileNameMatchConditionParametersOutput) ToUrlFileNameMatchConditionParametersOutput() UrlFileNameMatchConditionParametersOutput {
	return o
}

func (o UrlFileNameMatchConditionParametersOutput) ToUrlFileNameMatchConditionParametersOutputWithContext(ctx context.Context) UrlFileNameMatchConditionParametersOutput {
	return o
}

func (o UrlFileNameMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlFileNameMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o UrlFileNameMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UrlFileNameMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o UrlFileNameMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlFileNameMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UrlFileNameMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v UrlFileNameMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

func (o UrlFileNameMatchConditionParametersOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlFileNameMatchConditionParameters) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type UrlFileNameMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type UrlFileNameMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToUrlFileNameMatchConditionParametersResponseOutput() UrlFileNameMatchConditionParametersResponseOutput
	ToUrlFileNameMatchConditionParametersResponseOutputWithContext(context.Context) UrlFileNameMatchConditionParametersResponseOutput
}

type UrlFileNameMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (UrlFileNameMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlFileNameMatchConditionParametersResponse)(nil)).Elem()
}

func (i UrlFileNameMatchConditionParametersResponseArgs) ToUrlFileNameMatchConditionParametersResponseOutput() UrlFileNameMatchConditionParametersResponseOutput {
	return i.ToUrlFileNameMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i UrlFileNameMatchConditionParametersResponseArgs) ToUrlFileNameMatchConditionParametersResponseOutputWithContext(ctx context.Context) UrlFileNameMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlFileNameMatchConditionParametersResponseOutput)
}

type UrlFileNameMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (UrlFileNameMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlFileNameMatchConditionParametersResponse)(nil)).Elem()
}

func (o UrlFileNameMatchConditionParametersResponseOutput) ToUrlFileNameMatchConditionParametersResponseOutput() UrlFileNameMatchConditionParametersResponseOutput {
	return o
}

func (o UrlFileNameMatchConditionParametersResponseOutput) ToUrlFileNameMatchConditionParametersResponseOutputWithContext(ctx context.Context) UrlFileNameMatchConditionParametersResponseOutput {
	return o
}

func (o UrlFileNameMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlFileNameMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o UrlFileNameMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UrlFileNameMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o UrlFileNameMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlFileNameMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UrlFileNameMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v UrlFileNameMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o UrlFileNameMatchConditionParametersResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlFileNameMatchConditionParametersResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type UrlPathMatchConditionParameters struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type UrlPathMatchConditionParametersInput interface {
	pulumi.Input

	ToUrlPathMatchConditionParametersOutput() UrlPathMatchConditionParametersOutput
	ToUrlPathMatchConditionParametersOutputWithContext(context.Context) UrlPathMatchConditionParametersOutput
}

type UrlPathMatchConditionParametersArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (UrlPathMatchConditionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlPathMatchConditionParameters)(nil)).Elem()
}

func (i UrlPathMatchConditionParametersArgs) ToUrlPathMatchConditionParametersOutput() UrlPathMatchConditionParametersOutput {
	return i.ToUrlPathMatchConditionParametersOutputWithContext(context.Background())
}

func (i UrlPathMatchConditionParametersArgs) ToUrlPathMatchConditionParametersOutputWithContext(ctx context.Context) UrlPathMatchConditionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlPathMatchConditionParametersOutput)
}

type UrlPathMatchConditionParametersOutput struct{ *pulumi.OutputState }

func (UrlPathMatchConditionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlPathMatchConditionParameters)(nil)).Elem()
}

func (o UrlPathMatchConditionParametersOutput) ToUrlPathMatchConditionParametersOutput() UrlPathMatchConditionParametersOutput {
	return o
}

func (o UrlPathMatchConditionParametersOutput) ToUrlPathMatchConditionParametersOutputWithContext(ctx context.Context) UrlPathMatchConditionParametersOutput {
	return o
}

func (o UrlPathMatchConditionParametersOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlPathMatchConditionParameters) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o UrlPathMatchConditionParametersOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UrlPathMatchConditionParameters) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o UrlPathMatchConditionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlPathMatchConditionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UrlPathMatchConditionParametersOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v UrlPathMatchConditionParameters) string { return v.Operator }).(pulumi.StringOutput)
}

func (o UrlPathMatchConditionParametersOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlPathMatchConditionParameters) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type UrlPathMatchConditionParametersResponse struct {
	MatchValues     []string `pulumi:"matchValues"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	OdataType       string   `pulumi:"odataType"`
	Operator        string   `pulumi:"operator"`
	Transforms      []string `pulumi:"transforms"`
}





type UrlPathMatchConditionParametersResponseInput interface {
	pulumi.Input

	ToUrlPathMatchConditionParametersResponseOutput() UrlPathMatchConditionParametersResponseOutput
	ToUrlPathMatchConditionParametersResponseOutputWithContext(context.Context) UrlPathMatchConditionParametersResponseOutput
}

type UrlPathMatchConditionParametersResponseArgs struct {
	MatchValues     pulumi.StringArrayInput `pulumi:"matchValues"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (UrlPathMatchConditionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlPathMatchConditionParametersResponse)(nil)).Elem()
}

func (i UrlPathMatchConditionParametersResponseArgs) ToUrlPathMatchConditionParametersResponseOutput() UrlPathMatchConditionParametersResponseOutput {
	return i.ToUrlPathMatchConditionParametersResponseOutputWithContext(context.Background())
}

func (i UrlPathMatchConditionParametersResponseArgs) ToUrlPathMatchConditionParametersResponseOutputWithContext(ctx context.Context) UrlPathMatchConditionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlPathMatchConditionParametersResponseOutput)
}

type UrlPathMatchConditionParametersResponseOutput struct{ *pulumi.OutputState }

func (UrlPathMatchConditionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlPathMatchConditionParametersResponse)(nil)).Elem()
}

func (o UrlPathMatchConditionParametersResponseOutput) ToUrlPathMatchConditionParametersResponseOutput() UrlPathMatchConditionParametersResponseOutput {
	return o
}

func (o UrlPathMatchConditionParametersResponseOutput) ToUrlPathMatchConditionParametersResponseOutputWithContext(ctx context.Context) UrlPathMatchConditionParametersResponseOutput {
	return o
}

func (o UrlPathMatchConditionParametersResponseOutput) MatchValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlPathMatchConditionParametersResponse) []string { return v.MatchValues }).(pulumi.StringArrayOutput)
}

func (o UrlPathMatchConditionParametersResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UrlPathMatchConditionParametersResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o UrlPathMatchConditionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlPathMatchConditionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UrlPathMatchConditionParametersResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v UrlPathMatchConditionParametersResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o UrlPathMatchConditionParametersResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlPathMatchConditionParametersResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type UrlRedirectAction struct {
	Name       string                      `pulumi:"name"`
	Parameters UrlRedirectActionParameters `pulumi:"parameters"`
}





type UrlRedirectActionInput interface {
	pulumi.Input

	ToUrlRedirectActionOutput() UrlRedirectActionOutput
	ToUrlRedirectActionOutputWithContext(context.Context) UrlRedirectActionOutput
}

type UrlRedirectActionArgs struct {
	Name       pulumi.StringInput               `pulumi:"name"`
	Parameters UrlRedirectActionParametersInput `pulumi:"parameters"`
}

func (UrlRedirectActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRedirectAction)(nil)).Elem()
}

func (i UrlRedirectActionArgs) ToUrlRedirectActionOutput() UrlRedirectActionOutput {
	return i.ToUrlRedirectActionOutputWithContext(context.Background())
}

func (i UrlRedirectActionArgs) ToUrlRedirectActionOutputWithContext(ctx context.Context) UrlRedirectActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlRedirectActionOutput)
}

type UrlRedirectActionOutput struct{ *pulumi.OutputState }

func (UrlRedirectActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRedirectAction)(nil)).Elem()
}

func (o UrlRedirectActionOutput) ToUrlRedirectActionOutput() UrlRedirectActionOutput {
	return o
}

func (o UrlRedirectActionOutput) ToUrlRedirectActionOutputWithContext(ctx context.Context) UrlRedirectActionOutput {
	return o
}

func (o UrlRedirectActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRedirectAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o UrlRedirectActionOutput) Parameters() UrlRedirectActionParametersOutput {
	return o.ApplyT(func(v UrlRedirectAction) UrlRedirectActionParameters { return v.Parameters }).(UrlRedirectActionParametersOutput)
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





type UrlRedirectActionParametersInput interface {
	pulumi.Input

	ToUrlRedirectActionParametersOutput() UrlRedirectActionParametersOutput
	ToUrlRedirectActionParametersOutputWithContext(context.Context) UrlRedirectActionParametersOutput
}

type UrlRedirectActionParametersArgs struct {
	CustomFragment      pulumi.StringPtrInput `pulumi:"customFragment"`
	CustomHostname      pulumi.StringPtrInput `pulumi:"customHostname"`
	CustomPath          pulumi.StringPtrInput `pulumi:"customPath"`
	CustomQueryString   pulumi.StringPtrInput `pulumi:"customQueryString"`
	DestinationProtocol pulumi.StringPtrInput `pulumi:"destinationProtocol"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
	RedirectType        pulumi.StringInput    `pulumi:"redirectType"`
}

func (UrlRedirectActionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRedirectActionParameters)(nil)).Elem()
}

func (i UrlRedirectActionParametersArgs) ToUrlRedirectActionParametersOutput() UrlRedirectActionParametersOutput {
	return i.ToUrlRedirectActionParametersOutputWithContext(context.Background())
}

func (i UrlRedirectActionParametersArgs) ToUrlRedirectActionParametersOutputWithContext(ctx context.Context) UrlRedirectActionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlRedirectActionParametersOutput)
}

type UrlRedirectActionParametersOutput struct{ *pulumi.OutputState }

func (UrlRedirectActionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRedirectActionParameters)(nil)).Elem()
}

func (o UrlRedirectActionParametersOutput) ToUrlRedirectActionParametersOutput() UrlRedirectActionParametersOutput {
	return o
}

func (o UrlRedirectActionParametersOutput) ToUrlRedirectActionParametersOutputWithContext(ctx context.Context) UrlRedirectActionParametersOutput {
	return o
}

func (o UrlRedirectActionParametersOutput) CustomFragment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlRedirectActionParameters) *string { return v.CustomFragment }).(pulumi.StringPtrOutput)
}

func (o UrlRedirectActionParametersOutput) CustomHostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlRedirectActionParameters) *string { return v.CustomHostname }).(pulumi.StringPtrOutput)
}

func (o UrlRedirectActionParametersOutput) CustomPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlRedirectActionParameters) *string { return v.CustomPath }).(pulumi.StringPtrOutput)
}

func (o UrlRedirectActionParametersOutput) CustomQueryString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlRedirectActionParameters) *string { return v.CustomQueryString }).(pulumi.StringPtrOutput)
}

func (o UrlRedirectActionParametersOutput) DestinationProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlRedirectActionParameters) *string { return v.DestinationProtocol }).(pulumi.StringPtrOutput)
}

func (o UrlRedirectActionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRedirectActionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UrlRedirectActionParametersOutput) RedirectType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRedirectActionParameters) string { return v.RedirectType }).(pulumi.StringOutput)
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





type UrlRedirectActionParametersResponseInput interface {
	pulumi.Input

	ToUrlRedirectActionParametersResponseOutput() UrlRedirectActionParametersResponseOutput
	ToUrlRedirectActionParametersResponseOutputWithContext(context.Context) UrlRedirectActionParametersResponseOutput
}

type UrlRedirectActionParametersResponseArgs struct {
	CustomFragment      pulumi.StringPtrInput `pulumi:"customFragment"`
	CustomHostname      pulumi.StringPtrInput `pulumi:"customHostname"`
	CustomPath          pulumi.StringPtrInput `pulumi:"customPath"`
	CustomQueryString   pulumi.StringPtrInput `pulumi:"customQueryString"`
	DestinationProtocol pulumi.StringPtrInput `pulumi:"destinationProtocol"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
	RedirectType        pulumi.StringInput    `pulumi:"redirectType"`
}

func (UrlRedirectActionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRedirectActionParametersResponse)(nil)).Elem()
}

func (i UrlRedirectActionParametersResponseArgs) ToUrlRedirectActionParametersResponseOutput() UrlRedirectActionParametersResponseOutput {
	return i.ToUrlRedirectActionParametersResponseOutputWithContext(context.Background())
}

func (i UrlRedirectActionParametersResponseArgs) ToUrlRedirectActionParametersResponseOutputWithContext(ctx context.Context) UrlRedirectActionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlRedirectActionParametersResponseOutput)
}

type UrlRedirectActionParametersResponseOutput struct{ *pulumi.OutputState }

func (UrlRedirectActionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRedirectActionParametersResponse)(nil)).Elem()
}

func (o UrlRedirectActionParametersResponseOutput) ToUrlRedirectActionParametersResponseOutput() UrlRedirectActionParametersResponseOutput {
	return o
}

func (o UrlRedirectActionParametersResponseOutput) ToUrlRedirectActionParametersResponseOutputWithContext(ctx context.Context) UrlRedirectActionParametersResponseOutput {
	return o
}

func (o UrlRedirectActionParametersResponseOutput) CustomFragment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlRedirectActionParametersResponse) *string { return v.CustomFragment }).(pulumi.StringPtrOutput)
}

func (o UrlRedirectActionParametersResponseOutput) CustomHostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlRedirectActionParametersResponse) *string { return v.CustomHostname }).(pulumi.StringPtrOutput)
}

func (o UrlRedirectActionParametersResponseOutput) CustomPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlRedirectActionParametersResponse) *string { return v.CustomPath }).(pulumi.StringPtrOutput)
}

func (o UrlRedirectActionParametersResponseOutput) CustomQueryString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlRedirectActionParametersResponse) *string { return v.CustomQueryString }).(pulumi.StringPtrOutput)
}

func (o UrlRedirectActionParametersResponseOutput) DestinationProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlRedirectActionParametersResponse) *string { return v.DestinationProtocol }).(pulumi.StringPtrOutput)
}

func (o UrlRedirectActionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRedirectActionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UrlRedirectActionParametersResponseOutput) RedirectType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRedirectActionParametersResponse) string { return v.RedirectType }).(pulumi.StringOutput)
}

type UrlRedirectActionResponse struct {
	Name       string                              `pulumi:"name"`
	Parameters UrlRedirectActionParametersResponse `pulumi:"parameters"`
}





type UrlRedirectActionResponseInput interface {
	pulumi.Input

	ToUrlRedirectActionResponseOutput() UrlRedirectActionResponseOutput
	ToUrlRedirectActionResponseOutputWithContext(context.Context) UrlRedirectActionResponseOutput
}

type UrlRedirectActionResponseArgs struct {
	Name       pulumi.StringInput                       `pulumi:"name"`
	Parameters UrlRedirectActionParametersResponseInput `pulumi:"parameters"`
}

func (UrlRedirectActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRedirectActionResponse)(nil)).Elem()
}

func (i UrlRedirectActionResponseArgs) ToUrlRedirectActionResponseOutput() UrlRedirectActionResponseOutput {
	return i.ToUrlRedirectActionResponseOutputWithContext(context.Background())
}

func (i UrlRedirectActionResponseArgs) ToUrlRedirectActionResponseOutputWithContext(ctx context.Context) UrlRedirectActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlRedirectActionResponseOutput)
}

type UrlRedirectActionResponseOutput struct{ *pulumi.OutputState }

func (UrlRedirectActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRedirectActionResponse)(nil)).Elem()
}

func (o UrlRedirectActionResponseOutput) ToUrlRedirectActionResponseOutput() UrlRedirectActionResponseOutput {
	return o
}

func (o UrlRedirectActionResponseOutput) ToUrlRedirectActionResponseOutputWithContext(ctx context.Context) UrlRedirectActionResponseOutput {
	return o
}

func (o UrlRedirectActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRedirectActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o UrlRedirectActionResponseOutput) Parameters() UrlRedirectActionParametersResponseOutput {
	return o.ApplyT(func(v UrlRedirectActionResponse) UrlRedirectActionParametersResponse { return v.Parameters }).(UrlRedirectActionParametersResponseOutput)
}

type UrlRewriteAction struct {
	Name       string                     `pulumi:"name"`
	Parameters UrlRewriteActionParameters `pulumi:"parameters"`
}





type UrlRewriteActionInput interface {
	pulumi.Input

	ToUrlRewriteActionOutput() UrlRewriteActionOutput
	ToUrlRewriteActionOutputWithContext(context.Context) UrlRewriteActionOutput
}

type UrlRewriteActionArgs struct {
	Name       pulumi.StringInput              `pulumi:"name"`
	Parameters UrlRewriteActionParametersInput `pulumi:"parameters"`
}

func (UrlRewriteActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRewriteAction)(nil)).Elem()
}

func (i UrlRewriteActionArgs) ToUrlRewriteActionOutput() UrlRewriteActionOutput {
	return i.ToUrlRewriteActionOutputWithContext(context.Background())
}

func (i UrlRewriteActionArgs) ToUrlRewriteActionOutputWithContext(ctx context.Context) UrlRewriteActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlRewriteActionOutput)
}

type UrlRewriteActionOutput struct{ *pulumi.OutputState }

func (UrlRewriteActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRewriteAction)(nil)).Elem()
}

func (o UrlRewriteActionOutput) ToUrlRewriteActionOutput() UrlRewriteActionOutput {
	return o
}

func (o UrlRewriteActionOutput) ToUrlRewriteActionOutputWithContext(ctx context.Context) UrlRewriteActionOutput {
	return o
}

func (o UrlRewriteActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRewriteAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o UrlRewriteActionOutput) Parameters() UrlRewriteActionParametersOutput {
	return o.ApplyT(func(v UrlRewriteAction) UrlRewriteActionParameters { return v.Parameters }).(UrlRewriteActionParametersOutput)
}

type UrlRewriteActionParameters struct {
	Destination           string `pulumi:"destination"`
	OdataType             string `pulumi:"odataType"`
	PreserveUnmatchedPath *bool  `pulumi:"preserveUnmatchedPath"`
	SourcePattern         string `pulumi:"sourcePattern"`
}





type UrlRewriteActionParametersInput interface {
	pulumi.Input

	ToUrlRewriteActionParametersOutput() UrlRewriteActionParametersOutput
	ToUrlRewriteActionParametersOutputWithContext(context.Context) UrlRewriteActionParametersOutput
}

type UrlRewriteActionParametersArgs struct {
	Destination           pulumi.StringInput  `pulumi:"destination"`
	OdataType             pulumi.StringInput  `pulumi:"odataType"`
	PreserveUnmatchedPath pulumi.BoolPtrInput `pulumi:"preserveUnmatchedPath"`
	SourcePattern         pulumi.StringInput  `pulumi:"sourcePattern"`
}

func (UrlRewriteActionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRewriteActionParameters)(nil)).Elem()
}

func (i UrlRewriteActionParametersArgs) ToUrlRewriteActionParametersOutput() UrlRewriteActionParametersOutput {
	return i.ToUrlRewriteActionParametersOutputWithContext(context.Background())
}

func (i UrlRewriteActionParametersArgs) ToUrlRewriteActionParametersOutputWithContext(ctx context.Context) UrlRewriteActionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlRewriteActionParametersOutput)
}

type UrlRewriteActionParametersOutput struct{ *pulumi.OutputState }

func (UrlRewriteActionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRewriteActionParameters)(nil)).Elem()
}

func (o UrlRewriteActionParametersOutput) ToUrlRewriteActionParametersOutput() UrlRewriteActionParametersOutput {
	return o
}

func (o UrlRewriteActionParametersOutput) ToUrlRewriteActionParametersOutputWithContext(ctx context.Context) UrlRewriteActionParametersOutput {
	return o
}

func (o UrlRewriteActionParametersOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRewriteActionParameters) string { return v.Destination }).(pulumi.StringOutput)
}

func (o UrlRewriteActionParametersOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRewriteActionParameters) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UrlRewriteActionParametersOutput) PreserveUnmatchedPath() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UrlRewriteActionParameters) *bool { return v.PreserveUnmatchedPath }).(pulumi.BoolPtrOutput)
}

func (o UrlRewriteActionParametersOutput) SourcePattern() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRewriteActionParameters) string { return v.SourcePattern }).(pulumi.StringOutput)
}

type UrlRewriteActionParametersResponse struct {
	Destination           string `pulumi:"destination"`
	OdataType             string `pulumi:"odataType"`
	PreserveUnmatchedPath *bool  `pulumi:"preserveUnmatchedPath"`
	SourcePattern         string `pulumi:"sourcePattern"`
}





type UrlRewriteActionParametersResponseInput interface {
	pulumi.Input

	ToUrlRewriteActionParametersResponseOutput() UrlRewriteActionParametersResponseOutput
	ToUrlRewriteActionParametersResponseOutputWithContext(context.Context) UrlRewriteActionParametersResponseOutput
}

type UrlRewriteActionParametersResponseArgs struct {
	Destination           pulumi.StringInput  `pulumi:"destination"`
	OdataType             pulumi.StringInput  `pulumi:"odataType"`
	PreserveUnmatchedPath pulumi.BoolPtrInput `pulumi:"preserveUnmatchedPath"`
	SourcePattern         pulumi.StringInput  `pulumi:"sourcePattern"`
}

func (UrlRewriteActionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRewriteActionParametersResponse)(nil)).Elem()
}

func (i UrlRewriteActionParametersResponseArgs) ToUrlRewriteActionParametersResponseOutput() UrlRewriteActionParametersResponseOutput {
	return i.ToUrlRewriteActionParametersResponseOutputWithContext(context.Background())
}

func (i UrlRewriteActionParametersResponseArgs) ToUrlRewriteActionParametersResponseOutputWithContext(ctx context.Context) UrlRewriteActionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlRewriteActionParametersResponseOutput)
}

type UrlRewriteActionParametersResponseOutput struct{ *pulumi.OutputState }

func (UrlRewriteActionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRewriteActionParametersResponse)(nil)).Elem()
}

func (o UrlRewriteActionParametersResponseOutput) ToUrlRewriteActionParametersResponseOutput() UrlRewriteActionParametersResponseOutput {
	return o
}

func (o UrlRewriteActionParametersResponseOutput) ToUrlRewriteActionParametersResponseOutputWithContext(ctx context.Context) UrlRewriteActionParametersResponseOutput {
	return o
}

func (o UrlRewriteActionParametersResponseOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRewriteActionParametersResponse) string { return v.Destination }).(pulumi.StringOutput)
}

func (o UrlRewriteActionParametersResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRewriteActionParametersResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UrlRewriteActionParametersResponseOutput) PreserveUnmatchedPath() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UrlRewriteActionParametersResponse) *bool { return v.PreserveUnmatchedPath }).(pulumi.BoolPtrOutput)
}

func (o UrlRewriteActionParametersResponseOutput) SourcePattern() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRewriteActionParametersResponse) string { return v.SourcePattern }).(pulumi.StringOutput)
}

type UrlRewriteActionResponse struct {
	Name       string                             `pulumi:"name"`
	Parameters UrlRewriteActionParametersResponse `pulumi:"parameters"`
}





type UrlRewriteActionResponseInput interface {
	pulumi.Input

	ToUrlRewriteActionResponseOutput() UrlRewriteActionResponseOutput
	ToUrlRewriteActionResponseOutputWithContext(context.Context) UrlRewriteActionResponseOutput
}

type UrlRewriteActionResponseArgs struct {
	Name       pulumi.StringInput                      `pulumi:"name"`
	Parameters UrlRewriteActionParametersResponseInput `pulumi:"parameters"`
}

func (UrlRewriteActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRewriteActionResponse)(nil)).Elem()
}

func (i UrlRewriteActionResponseArgs) ToUrlRewriteActionResponseOutput() UrlRewriteActionResponseOutput {
	return i.ToUrlRewriteActionResponseOutputWithContext(context.Background())
}

func (i UrlRewriteActionResponseArgs) ToUrlRewriteActionResponseOutputWithContext(ctx context.Context) UrlRewriteActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlRewriteActionResponseOutput)
}

type UrlRewriteActionResponseOutput struct{ *pulumi.OutputState }

func (UrlRewriteActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlRewriteActionResponse)(nil)).Elem()
}

func (o UrlRewriteActionResponseOutput) ToUrlRewriteActionResponseOutput() UrlRewriteActionResponseOutput {
	return o
}

func (o UrlRewriteActionResponseOutput) ToUrlRewriteActionResponseOutputWithContext(ctx context.Context) UrlRewriteActionResponseOutput {
	return o
}

func (o UrlRewriteActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v UrlRewriteActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o UrlRewriteActionResponseOutput) Parameters() UrlRewriteActionParametersResponseOutput {
	return o.ApplyT(func(v UrlRewriteActionResponse) UrlRewriteActionParametersResponse { return v.Parameters }).(UrlRewriteActionParametersResponseOutput)
}

type UrlSigningAction struct {
	Name       string                     `pulumi:"name"`
	Parameters UrlSigningActionParameters `pulumi:"parameters"`
}





type UrlSigningActionInput interface {
	pulumi.Input

	ToUrlSigningActionOutput() UrlSigningActionOutput
	ToUrlSigningActionOutputWithContext(context.Context) UrlSigningActionOutput
}

type UrlSigningActionArgs struct {
	Name       pulumi.StringInput              `pulumi:"name"`
	Parameters UrlSigningActionParametersInput `pulumi:"parameters"`
}

func (UrlSigningActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningAction)(nil)).Elem()
}

func (i UrlSigningActionArgs) ToUrlSigningActionOutput() UrlSigningActionOutput {
	return i.ToUrlSigningActionOutputWithContext(context.Background())
}

func (i UrlSigningActionArgs) ToUrlSigningActionOutputWithContext(ctx context.Context) UrlSigningActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningActionOutput)
}

type UrlSigningActionOutput struct{ *pulumi.OutputState }

func (UrlSigningActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningAction)(nil)).Elem()
}

func (o UrlSigningActionOutput) ToUrlSigningActionOutput() UrlSigningActionOutput {
	return o
}

func (o UrlSigningActionOutput) ToUrlSigningActionOutputWithContext(ctx context.Context) UrlSigningActionOutput {
	return o
}

func (o UrlSigningActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v UrlSigningAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o UrlSigningActionOutput) Parameters() UrlSigningActionParametersOutput {
	return o.ApplyT(func(v UrlSigningAction) UrlSigningActionParameters { return v.Parameters }).(UrlSigningActionParametersOutput)
}

type UrlSigningActionParameters struct {
	Algorithm             *string                     `pulumi:"algorithm"`
	IpSubnets             []string                    `pulumi:"ipSubnets"`
	KeyId                 string                      `pulumi:"keyId"`
	OdataType             *string                     `pulumi:"odataType"`
	ParameterNameOverride []UrlSigningParamIdentifier `pulumi:"parameterNameOverride"`
}





type UrlSigningActionParametersInput interface {
	pulumi.Input

	ToUrlSigningActionParametersOutput() UrlSigningActionParametersOutput
	ToUrlSigningActionParametersOutputWithContext(context.Context) UrlSigningActionParametersOutput
}

type UrlSigningActionParametersArgs struct {
	Algorithm             pulumi.StringPtrInput               `pulumi:"algorithm"`
	IpSubnets             pulumi.StringArrayInput             `pulumi:"ipSubnets"`
	KeyId                 pulumi.StringInput                  `pulumi:"keyId"`
	OdataType             pulumi.StringPtrInput               `pulumi:"odataType"`
	ParameterNameOverride UrlSigningParamIdentifierArrayInput `pulumi:"parameterNameOverride"`
}

func (UrlSigningActionParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningActionParameters)(nil)).Elem()
}

func (i UrlSigningActionParametersArgs) ToUrlSigningActionParametersOutput() UrlSigningActionParametersOutput {
	return i.ToUrlSigningActionParametersOutputWithContext(context.Background())
}

func (i UrlSigningActionParametersArgs) ToUrlSigningActionParametersOutputWithContext(ctx context.Context) UrlSigningActionParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningActionParametersOutput)
}

type UrlSigningActionParametersOutput struct{ *pulumi.OutputState }

func (UrlSigningActionParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningActionParameters)(nil)).Elem()
}

func (o UrlSigningActionParametersOutput) ToUrlSigningActionParametersOutput() UrlSigningActionParametersOutput {
	return o
}

func (o UrlSigningActionParametersOutput) ToUrlSigningActionParametersOutputWithContext(ctx context.Context) UrlSigningActionParametersOutput {
	return o
}

func (o UrlSigningActionParametersOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlSigningActionParameters) *string { return v.Algorithm }).(pulumi.StringPtrOutput)
}

func (o UrlSigningActionParametersOutput) IpSubnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlSigningActionParameters) []string { return v.IpSubnets }).(pulumi.StringArrayOutput)
}

func (o UrlSigningActionParametersOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v UrlSigningActionParameters) string { return v.KeyId }).(pulumi.StringOutput)
}

func (o UrlSigningActionParametersOutput) OdataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlSigningActionParameters) *string { return v.OdataType }).(pulumi.StringPtrOutput)
}

func (o UrlSigningActionParametersOutput) ParameterNameOverride() UrlSigningParamIdentifierArrayOutput {
	return o.ApplyT(func(v UrlSigningActionParameters) []UrlSigningParamIdentifier { return v.ParameterNameOverride }).(UrlSigningParamIdentifierArrayOutput)
}

type UrlSigningActionParametersResponse struct {
	Algorithm             *string                             `pulumi:"algorithm"`
	IpSubnets             []string                            `pulumi:"ipSubnets"`
	KeyId                 string                              `pulumi:"keyId"`
	OdataType             *string                             `pulumi:"odataType"`
	ParameterNameOverride []UrlSigningParamIdentifierResponse `pulumi:"parameterNameOverride"`
}





type UrlSigningActionParametersResponseInput interface {
	pulumi.Input

	ToUrlSigningActionParametersResponseOutput() UrlSigningActionParametersResponseOutput
	ToUrlSigningActionParametersResponseOutputWithContext(context.Context) UrlSigningActionParametersResponseOutput
}

type UrlSigningActionParametersResponseArgs struct {
	Algorithm             pulumi.StringPtrInput                       `pulumi:"algorithm"`
	IpSubnets             pulumi.StringArrayInput                     `pulumi:"ipSubnets"`
	KeyId                 pulumi.StringInput                          `pulumi:"keyId"`
	OdataType             pulumi.StringPtrInput                       `pulumi:"odataType"`
	ParameterNameOverride UrlSigningParamIdentifierResponseArrayInput `pulumi:"parameterNameOverride"`
}

func (UrlSigningActionParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningActionParametersResponse)(nil)).Elem()
}

func (i UrlSigningActionParametersResponseArgs) ToUrlSigningActionParametersResponseOutput() UrlSigningActionParametersResponseOutput {
	return i.ToUrlSigningActionParametersResponseOutputWithContext(context.Background())
}

func (i UrlSigningActionParametersResponseArgs) ToUrlSigningActionParametersResponseOutputWithContext(ctx context.Context) UrlSigningActionParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningActionParametersResponseOutput)
}

type UrlSigningActionParametersResponseOutput struct{ *pulumi.OutputState }

func (UrlSigningActionParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningActionParametersResponse)(nil)).Elem()
}

func (o UrlSigningActionParametersResponseOutput) ToUrlSigningActionParametersResponseOutput() UrlSigningActionParametersResponseOutput {
	return o
}

func (o UrlSigningActionParametersResponseOutput) ToUrlSigningActionParametersResponseOutputWithContext(ctx context.Context) UrlSigningActionParametersResponseOutput {
	return o
}

func (o UrlSigningActionParametersResponseOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlSigningActionParametersResponse) *string { return v.Algorithm }).(pulumi.StringPtrOutput)
}

func (o UrlSigningActionParametersResponseOutput) IpSubnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UrlSigningActionParametersResponse) []string { return v.IpSubnets }).(pulumi.StringArrayOutput)
}

func (o UrlSigningActionParametersResponseOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v UrlSigningActionParametersResponse) string { return v.KeyId }).(pulumi.StringOutput)
}

func (o UrlSigningActionParametersResponseOutput) OdataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UrlSigningActionParametersResponse) *string { return v.OdataType }).(pulumi.StringPtrOutput)
}

func (o UrlSigningActionParametersResponseOutput) ParameterNameOverride() UrlSigningParamIdentifierResponseArrayOutput {
	return o.ApplyT(func(v UrlSigningActionParametersResponse) []UrlSigningParamIdentifierResponse {
		return v.ParameterNameOverride
	}).(UrlSigningParamIdentifierResponseArrayOutput)
}

type UrlSigningActionResponse struct {
	Name       string                             `pulumi:"name"`
	Parameters UrlSigningActionParametersResponse `pulumi:"parameters"`
}





type UrlSigningActionResponseInput interface {
	pulumi.Input

	ToUrlSigningActionResponseOutput() UrlSigningActionResponseOutput
	ToUrlSigningActionResponseOutputWithContext(context.Context) UrlSigningActionResponseOutput
}

type UrlSigningActionResponseArgs struct {
	Name       pulumi.StringInput                      `pulumi:"name"`
	Parameters UrlSigningActionParametersResponseInput `pulumi:"parameters"`
}

func (UrlSigningActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningActionResponse)(nil)).Elem()
}

func (i UrlSigningActionResponseArgs) ToUrlSigningActionResponseOutput() UrlSigningActionResponseOutput {
	return i.ToUrlSigningActionResponseOutputWithContext(context.Background())
}

func (i UrlSigningActionResponseArgs) ToUrlSigningActionResponseOutputWithContext(ctx context.Context) UrlSigningActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningActionResponseOutput)
}

type UrlSigningActionResponseOutput struct{ *pulumi.OutputState }

func (UrlSigningActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningActionResponse)(nil)).Elem()
}

func (o UrlSigningActionResponseOutput) ToUrlSigningActionResponseOutput() UrlSigningActionResponseOutput {
	return o
}

func (o UrlSigningActionResponseOutput) ToUrlSigningActionResponseOutputWithContext(ctx context.Context) UrlSigningActionResponseOutput {
	return o
}

func (o UrlSigningActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v UrlSigningActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o UrlSigningActionResponseOutput) Parameters() UrlSigningActionParametersResponseOutput {
	return o.ApplyT(func(v UrlSigningActionResponse) UrlSigningActionParametersResponse { return v.Parameters }).(UrlSigningActionParametersResponseOutput)
}

type UrlSigningKey struct {
	KeyId               string                       `pulumi:"keyId"`
	KeySourceParameters KeyVaultSigningKeyParameters `pulumi:"keySourceParameters"`
}





type UrlSigningKeyInput interface {
	pulumi.Input

	ToUrlSigningKeyOutput() UrlSigningKeyOutput
	ToUrlSigningKeyOutputWithContext(context.Context) UrlSigningKeyOutput
}

type UrlSigningKeyArgs struct {
	KeyId               pulumi.StringInput                `pulumi:"keyId"`
	KeySourceParameters KeyVaultSigningKeyParametersInput `pulumi:"keySourceParameters"`
}

func (UrlSigningKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningKey)(nil)).Elem()
}

func (i UrlSigningKeyArgs) ToUrlSigningKeyOutput() UrlSigningKeyOutput {
	return i.ToUrlSigningKeyOutputWithContext(context.Background())
}

func (i UrlSigningKeyArgs) ToUrlSigningKeyOutputWithContext(ctx context.Context) UrlSigningKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningKeyOutput)
}





type UrlSigningKeyArrayInput interface {
	pulumi.Input

	ToUrlSigningKeyArrayOutput() UrlSigningKeyArrayOutput
	ToUrlSigningKeyArrayOutputWithContext(context.Context) UrlSigningKeyArrayOutput
}

type UrlSigningKeyArray []UrlSigningKeyInput

func (UrlSigningKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UrlSigningKey)(nil)).Elem()
}

func (i UrlSigningKeyArray) ToUrlSigningKeyArrayOutput() UrlSigningKeyArrayOutput {
	return i.ToUrlSigningKeyArrayOutputWithContext(context.Background())
}

func (i UrlSigningKeyArray) ToUrlSigningKeyArrayOutputWithContext(ctx context.Context) UrlSigningKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningKeyArrayOutput)
}

type UrlSigningKeyOutput struct{ *pulumi.OutputState }

func (UrlSigningKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningKey)(nil)).Elem()
}

func (o UrlSigningKeyOutput) ToUrlSigningKeyOutput() UrlSigningKeyOutput {
	return o
}

func (o UrlSigningKeyOutput) ToUrlSigningKeyOutputWithContext(ctx context.Context) UrlSigningKeyOutput {
	return o
}

func (o UrlSigningKeyOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v UrlSigningKey) string { return v.KeyId }).(pulumi.StringOutput)
}

func (o UrlSigningKeyOutput) KeySourceParameters() KeyVaultSigningKeyParametersOutput {
	return o.ApplyT(func(v UrlSigningKey) KeyVaultSigningKeyParameters { return v.KeySourceParameters }).(KeyVaultSigningKeyParametersOutput)
}

type UrlSigningKeyArrayOutput struct{ *pulumi.OutputState }

func (UrlSigningKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UrlSigningKey)(nil)).Elem()
}

func (o UrlSigningKeyArrayOutput) ToUrlSigningKeyArrayOutput() UrlSigningKeyArrayOutput {
	return o
}

func (o UrlSigningKeyArrayOutput) ToUrlSigningKeyArrayOutputWithContext(ctx context.Context) UrlSigningKeyArrayOutput {
	return o
}

func (o UrlSigningKeyArrayOutput) Index(i pulumi.IntInput) UrlSigningKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UrlSigningKey {
		return vs[0].([]UrlSigningKey)[vs[1].(int)]
	}).(UrlSigningKeyOutput)
}

type UrlSigningKeyResponse struct {
	KeyId               string                               `pulumi:"keyId"`
	KeySourceParameters KeyVaultSigningKeyParametersResponse `pulumi:"keySourceParameters"`
}





type UrlSigningKeyResponseInput interface {
	pulumi.Input

	ToUrlSigningKeyResponseOutput() UrlSigningKeyResponseOutput
	ToUrlSigningKeyResponseOutputWithContext(context.Context) UrlSigningKeyResponseOutput
}

type UrlSigningKeyResponseArgs struct {
	KeyId               pulumi.StringInput                        `pulumi:"keyId"`
	KeySourceParameters KeyVaultSigningKeyParametersResponseInput `pulumi:"keySourceParameters"`
}

func (UrlSigningKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningKeyResponse)(nil)).Elem()
}

func (i UrlSigningKeyResponseArgs) ToUrlSigningKeyResponseOutput() UrlSigningKeyResponseOutput {
	return i.ToUrlSigningKeyResponseOutputWithContext(context.Background())
}

func (i UrlSigningKeyResponseArgs) ToUrlSigningKeyResponseOutputWithContext(ctx context.Context) UrlSigningKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningKeyResponseOutput)
}





type UrlSigningKeyResponseArrayInput interface {
	pulumi.Input

	ToUrlSigningKeyResponseArrayOutput() UrlSigningKeyResponseArrayOutput
	ToUrlSigningKeyResponseArrayOutputWithContext(context.Context) UrlSigningKeyResponseArrayOutput
}

type UrlSigningKeyResponseArray []UrlSigningKeyResponseInput

func (UrlSigningKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UrlSigningKeyResponse)(nil)).Elem()
}

func (i UrlSigningKeyResponseArray) ToUrlSigningKeyResponseArrayOutput() UrlSigningKeyResponseArrayOutput {
	return i.ToUrlSigningKeyResponseArrayOutputWithContext(context.Background())
}

func (i UrlSigningKeyResponseArray) ToUrlSigningKeyResponseArrayOutputWithContext(ctx context.Context) UrlSigningKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningKeyResponseArrayOutput)
}

type UrlSigningKeyResponseOutput struct{ *pulumi.OutputState }

func (UrlSigningKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningKeyResponse)(nil)).Elem()
}

func (o UrlSigningKeyResponseOutput) ToUrlSigningKeyResponseOutput() UrlSigningKeyResponseOutput {
	return o
}

func (o UrlSigningKeyResponseOutput) ToUrlSigningKeyResponseOutputWithContext(ctx context.Context) UrlSigningKeyResponseOutput {
	return o
}

func (o UrlSigningKeyResponseOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v UrlSigningKeyResponse) string { return v.KeyId }).(pulumi.StringOutput)
}

func (o UrlSigningKeyResponseOutput) KeySourceParameters() KeyVaultSigningKeyParametersResponseOutput {
	return o.ApplyT(func(v UrlSigningKeyResponse) KeyVaultSigningKeyParametersResponse { return v.KeySourceParameters }).(KeyVaultSigningKeyParametersResponseOutput)
}

type UrlSigningKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (UrlSigningKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UrlSigningKeyResponse)(nil)).Elem()
}

func (o UrlSigningKeyResponseArrayOutput) ToUrlSigningKeyResponseArrayOutput() UrlSigningKeyResponseArrayOutput {
	return o
}

func (o UrlSigningKeyResponseArrayOutput) ToUrlSigningKeyResponseArrayOutputWithContext(ctx context.Context) UrlSigningKeyResponseArrayOutput {
	return o
}

func (o UrlSigningKeyResponseArrayOutput) Index(i pulumi.IntInput) UrlSigningKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UrlSigningKeyResponse {
		return vs[0].([]UrlSigningKeyResponse)[vs[1].(int)]
	}).(UrlSigningKeyResponseOutput)
}

type UrlSigningParamIdentifier struct {
	ParamIndicator string `pulumi:"paramIndicator"`
	ParamName      string `pulumi:"paramName"`
}





type UrlSigningParamIdentifierInput interface {
	pulumi.Input

	ToUrlSigningParamIdentifierOutput() UrlSigningParamIdentifierOutput
	ToUrlSigningParamIdentifierOutputWithContext(context.Context) UrlSigningParamIdentifierOutput
}

type UrlSigningParamIdentifierArgs struct {
	ParamIndicator pulumi.StringInput `pulumi:"paramIndicator"`
	ParamName      pulumi.StringInput `pulumi:"paramName"`
}

func (UrlSigningParamIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningParamIdentifier)(nil)).Elem()
}

func (i UrlSigningParamIdentifierArgs) ToUrlSigningParamIdentifierOutput() UrlSigningParamIdentifierOutput {
	return i.ToUrlSigningParamIdentifierOutputWithContext(context.Background())
}

func (i UrlSigningParamIdentifierArgs) ToUrlSigningParamIdentifierOutputWithContext(ctx context.Context) UrlSigningParamIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningParamIdentifierOutput)
}





type UrlSigningParamIdentifierArrayInput interface {
	pulumi.Input

	ToUrlSigningParamIdentifierArrayOutput() UrlSigningParamIdentifierArrayOutput
	ToUrlSigningParamIdentifierArrayOutputWithContext(context.Context) UrlSigningParamIdentifierArrayOutput
}

type UrlSigningParamIdentifierArray []UrlSigningParamIdentifierInput

func (UrlSigningParamIdentifierArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UrlSigningParamIdentifier)(nil)).Elem()
}

func (i UrlSigningParamIdentifierArray) ToUrlSigningParamIdentifierArrayOutput() UrlSigningParamIdentifierArrayOutput {
	return i.ToUrlSigningParamIdentifierArrayOutputWithContext(context.Background())
}

func (i UrlSigningParamIdentifierArray) ToUrlSigningParamIdentifierArrayOutputWithContext(ctx context.Context) UrlSigningParamIdentifierArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningParamIdentifierArrayOutput)
}

type UrlSigningParamIdentifierOutput struct{ *pulumi.OutputState }

func (UrlSigningParamIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningParamIdentifier)(nil)).Elem()
}

func (o UrlSigningParamIdentifierOutput) ToUrlSigningParamIdentifierOutput() UrlSigningParamIdentifierOutput {
	return o
}

func (o UrlSigningParamIdentifierOutput) ToUrlSigningParamIdentifierOutputWithContext(ctx context.Context) UrlSigningParamIdentifierOutput {
	return o
}

func (o UrlSigningParamIdentifierOutput) ParamIndicator() pulumi.StringOutput {
	return o.ApplyT(func(v UrlSigningParamIdentifier) string { return v.ParamIndicator }).(pulumi.StringOutput)
}

func (o UrlSigningParamIdentifierOutput) ParamName() pulumi.StringOutput {
	return o.ApplyT(func(v UrlSigningParamIdentifier) string { return v.ParamName }).(pulumi.StringOutput)
}

type UrlSigningParamIdentifierArrayOutput struct{ *pulumi.OutputState }

func (UrlSigningParamIdentifierArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UrlSigningParamIdentifier)(nil)).Elem()
}

func (o UrlSigningParamIdentifierArrayOutput) ToUrlSigningParamIdentifierArrayOutput() UrlSigningParamIdentifierArrayOutput {
	return o
}

func (o UrlSigningParamIdentifierArrayOutput) ToUrlSigningParamIdentifierArrayOutputWithContext(ctx context.Context) UrlSigningParamIdentifierArrayOutput {
	return o
}

func (o UrlSigningParamIdentifierArrayOutput) Index(i pulumi.IntInput) UrlSigningParamIdentifierOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UrlSigningParamIdentifier {
		return vs[0].([]UrlSigningParamIdentifier)[vs[1].(int)]
	}).(UrlSigningParamIdentifierOutput)
}

type UrlSigningParamIdentifierResponse struct {
	ParamIndicator string `pulumi:"paramIndicator"`
	ParamName      string `pulumi:"paramName"`
}





type UrlSigningParamIdentifierResponseInput interface {
	pulumi.Input

	ToUrlSigningParamIdentifierResponseOutput() UrlSigningParamIdentifierResponseOutput
	ToUrlSigningParamIdentifierResponseOutputWithContext(context.Context) UrlSigningParamIdentifierResponseOutput
}

type UrlSigningParamIdentifierResponseArgs struct {
	ParamIndicator pulumi.StringInput `pulumi:"paramIndicator"`
	ParamName      pulumi.StringInput `pulumi:"paramName"`
}

func (UrlSigningParamIdentifierResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningParamIdentifierResponse)(nil)).Elem()
}

func (i UrlSigningParamIdentifierResponseArgs) ToUrlSigningParamIdentifierResponseOutput() UrlSigningParamIdentifierResponseOutput {
	return i.ToUrlSigningParamIdentifierResponseOutputWithContext(context.Background())
}

func (i UrlSigningParamIdentifierResponseArgs) ToUrlSigningParamIdentifierResponseOutputWithContext(ctx context.Context) UrlSigningParamIdentifierResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningParamIdentifierResponseOutput)
}





type UrlSigningParamIdentifierResponseArrayInput interface {
	pulumi.Input

	ToUrlSigningParamIdentifierResponseArrayOutput() UrlSigningParamIdentifierResponseArrayOutput
	ToUrlSigningParamIdentifierResponseArrayOutputWithContext(context.Context) UrlSigningParamIdentifierResponseArrayOutput
}

type UrlSigningParamIdentifierResponseArray []UrlSigningParamIdentifierResponseInput

func (UrlSigningParamIdentifierResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UrlSigningParamIdentifierResponse)(nil)).Elem()
}

func (i UrlSigningParamIdentifierResponseArray) ToUrlSigningParamIdentifierResponseArrayOutput() UrlSigningParamIdentifierResponseArrayOutput {
	return i.ToUrlSigningParamIdentifierResponseArrayOutputWithContext(context.Background())
}

func (i UrlSigningParamIdentifierResponseArray) ToUrlSigningParamIdentifierResponseArrayOutputWithContext(ctx context.Context) UrlSigningParamIdentifierResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlSigningParamIdentifierResponseArrayOutput)
}

type UrlSigningParamIdentifierResponseOutput struct{ *pulumi.OutputState }

func (UrlSigningParamIdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlSigningParamIdentifierResponse)(nil)).Elem()
}

func (o UrlSigningParamIdentifierResponseOutput) ToUrlSigningParamIdentifierResponseOutput() UrlSigningParamIdentifierResponseOutput {
	return o
}

func (o UrlSigningParamIdentifierResponseOutput) ToUrlSigningParamIdentifierResponseOutputWithContext(ctx context.Context) UrlSigningParamIdentifierResponseOutput {
	return o
}

func (o UrlSigningParamIdentifierResponseOutput) ParamIndicator() pulumi.StringOutput {
	return o.ApplyT(func(v UrlSigningParamIdentifierResponse) string { return v.ParamIndicator }).(pulumi.StringOutput)
}

func (o UrlSigningParamIdentifierResponseOutput) ParamName() pulumi.StringOutput {
	return o.ApplyT(func(v UrlSigningParamIdentifierResponse) string { return v.ParamName }).(pulumi.StringOutput)
}

type UrlSigningParamIdentifierResponseArrayOutput struct{ *pulumi.OutputState }

func (UrlSigningParamIdentifierResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UrlSigningParamIdentifierResponse)(nil)).Elem()
}

func (o UrlSigningParamIdentifierResponseArrayOutput) ToUrlSigningParamIdentifierResponseArrayOutput() UrlSigningParamIdentifierResponseArrayOutput {
	return o
}

func (o UrlSigningParamIdentifierResponseArrayOutput) ToUrlSigningParamIdentifierResponseArrayOutputWithContext(ctx context.Context) UrlSigningParamIdentifierResponseArrayOutput {
	return o
}

func (o UrlSigningParamIdentifierResponseArrayOutput) Index(i pulumi.IntInput) UrlSigningParamIdentifierResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UrlSigningParamIdentifierResponse {
		return vs[0].([]UrlSigningParamIdentifierResponse)[vs[1].(int)]
	}).(UrlSigningParamIdentifierResponseOutput)
}

type UserManagedHttpsParametersResponse struct {
	CertificateSource           string                                      `pulumi:"certificateSource"`
	CertificateSourceParameters KeyVaultCertificateSourceParametersResponse `pulumi:"certificateSourceParameters"`
	MinimumTlsVersion           *string                                     `pulumi:"minimumTlsVersion"`
	ProtocolType                string                                      `pulumi:"protocolType"`
}





type UserManagedHttpsParametersResponseInput interface {
	pulumi.Input

	ToUserManagedHttpsParametersResponseOutput() UserManagedHttpsParametersResponseOutput
	ToUserManagedHttpsParametersResponseOutputWithContext(context.Context) UserManagedHttpsParametersResponseOutput
}

type UserManagedHttpsParametersResponseArgs struct {
	CertificateSource           pulumi.StringInput                               `pulumi:"certificateSource"`
	CertificateSourceParameters KeyVaultCertificateSourceParametersResponseInput `pulumi:"certificateSourceParameters"`
	MinimumTlsVersion           pulumi.StringPtrInput                            `pulumi:"minimumTlsVersion"`
	ProtocolType                pulumi.StringInput                               `pulumi:"protocolType"`
}

func (UserManagedHttpsParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserManagedHttpsParametersResponse)(nil)).Elem()
}

func (i UserManagedHttpsParametersResponseArgs) ToUserManagedHttpsParametersResponseOutput() UserManagedHttpsParametersResponseOutput {
	return i.ToUserManagedHttpsParametersResponseOutputWithContext(context.Background())
}

func (i UserManagedHttpsParametersResponseArgs) ToUserManagedHttpsParametersResponseOutputWithContext(ctx context.Context) UserManagedHttpsParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserManagedHttpsParametersResponseOutput)
}

type UserManagedHttpsParametersResponseOutput struct{ *pulumi.OutputState }

func (UserManagedHttpsParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserManagedHttpsParametersResponse)(nil)).Elem()
}

func (o UserManagedHttpsParametersResponseOutput) ToUserManagedHttpsParametersResponseOutput() UserManagedHttpsParametersResponseOutput {
	return o
}

func (o UserManagedHttpsParametersResponseOutput) ToUserManagedHttpsParametersResponseOutputWithContext(ctx context.Context) UserManagedHttpsParametersResponseOutput {
	return o
}

func (o UserManagedHttpsParametersResponseOutput) CertificateSource() pulumi.StringOutput {
	return o.ApplyT(func(v UserManagedHttpsParametersResponse) string { return v.CertificateSource }).(pulumi.StringOutput)
}

func (o UserManagedHttpsParametersResponseOutput) CertificateSourceParameters() KeyVaultCertificateSourceParametersResponseOutput {
	return o.ApplyT(func(v UserManagedHttpsParametersResponse) KeyVaultCertificateSourceParametersResponse {
		return v.CertificateSourceParameters
	}).(KeyVaultCertificateSourceParametersResponseOutput)
}

func (o UserManagedHttpsParametersResponseOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserManagedHttpsParametersResponse) *string { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

func (o UserManagedHttpsParametersResponseOutput) ProtocolType() pulumi.StringOutput {
	return o.ApplyT(func(v UserManagedHttpsParametersResponse) string { return v.ProtocolType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CacheExpirationActionParametersOutput{})
	pulumi.RegisterOutputType(CacheExpirationActionParametersResponseOutput{})
	pulumi.RegisterOutputType(CacheKeyQueryStringActionParametersOutput{})
	pulumi.RegisterOutputType(CacheKeyQueryStringActionParametersResponseOutput{})
	pulumi.RegisterOutputType(CdnCertificateSourceParametersResponseOutput{})
	pulumi.RegisterOutputType(CdnEndpointResponseOutput{})
	pulumi.RegisterOutputType(CdnEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(CdnManagedHttpsParametersResponseOutput{})
	pulumi.RegisterOutputType(CookiesMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(CookiesMatchConditionParametersResponseOutput{})
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
	pulumi.RegisterOutputType(DeepCreatedOriginGroupOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginGroupArrayOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginGroupResponseOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginResponseOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginResponseArrayOutput{})
	pulumi.RegisterOutputType(DeliveryRuleOutput{})
	pulumi.RegisterOutputType(DeliveryRuleArrayOutput{})
	pulumi.RegisterOutputType(DeliveryRuleCacheExpirationActionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleCacheExpirationActionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleCacheKeyQueryStringActionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleCacheKeyQueryStringActionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleCookiesConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleCookiesConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleHttpVersionConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleHttpVersionConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleIsDeviceConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleIsDeviceConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRulePostArgsConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRulePostArgsConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleQueryStringConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleQueryStringConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRemoteAddressConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRemoteAddressConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestBodyConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestBodyConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestHeaderActionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestHeaderActionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestHeaderConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestHeaderConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestMethodConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestMethodConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestSchemeConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestSchemeConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestUriConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleRequestUriConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(DeliveryRuleResponseHeaderActionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleResponseHeaderActionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleUrlFileExtensionConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleUrlFileExtensionConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleUrlFileNameConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleUrlFileNameConditionResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleUrlPathConditionOutput{})
	pulumi.RegisterOutputType(DeliveryRuleUrlPathConditionResponseOutput{})
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
	pulumi.RegisterOutputType(HeaderActionParametersOutput{})
	pulumi.RegisterOutputType(HeaderActionParametersResponseOutput{})
	pulumi.RegisterOutputType(HealthProbeParametersOutput{})
	pulumi.RegisterOutputType(HealthProbeParametersPtrOutput{})
	pulumi.RegisterOutputType(HealthProbeParametersResponseOutput{})
	pulumi.RegisterOutputType(HealthProbeParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersArrayOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersResponseOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersResponseArrayOutput{})
	pulumi.RegisterOutputType(HttpVersionMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(HttpVersionMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(IsDeviceMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(IsDeviceMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultCertificateSourceParametersResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultSigningKeyParametersOutput{})
	pulumi.RegisterOutputType(KeyVaultSigningKeyParametersResponseOutput{})
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
	pulumi.RegisterOutputType(PostArgsMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(PostArgsMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(QueryStringMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(QueryStringMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(RateLimitRuleOutput{})
	pulumi.RegisterOutputType(RateLimitRuleArrayOutput{})
	pulumi.RegisterOutputType(RateLimitRuleListOutput{})
	pulumi.RegisterOutputType(RateLimitRuleListPtrOutput{})
	pulumi.RegisterOutputType(RateLimitRuleListResponseOutput{})
	pulumi.RegisterOutputType(RateLimitRuleListResponsePtrOutput{})
	pulumi.RegisterOutputType(RateLimitRuleResponseOutput{})
	pulumi.RegisterOutputType(RateLimitRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(RemoteAddressMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(RemoteAddressMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(RequestBodyMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(RequestBodyMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(RequestHeaderMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(RequestHeaderMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(RequestMethodMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(RequestMethodMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(RequestSchemeMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(RequestSchemeMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(RequestUriMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(RequestUriMatchConditionParametersResponseOutput{})
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
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(UrlFileExtensionMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(UrlFileExtensionMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(UrlFileNameMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(UrlFileNameMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(UrlPathMatchConditionParametersOutput{})
	pulumi.RegisterOutputType(UrlPathMatchConditionParametersResponseOutput{})
	pulumi.RegisterOutputType(UrlRedirectActionOutput{})
	pulumi.RegisterOutputType(UrlRedirectActionParametersOutput{})
	pulumi.RegisterOutputType(UrlRedirectActionParametersResponseOutput{})
	pulumi.RegisterOutputType(UrlRedirectActionResponseOutput{})
	pulumi.RegisterOutputType(UrlRewriteActionOutput{})
	pulumi.RegisterOutputType(UrlRewriteActionParametersOutput{})
	pulumi.RegisterOutputType(UrlRewriteActionParametersResponseOutput{})
	pulumi.RegisterOutputType(UrlRewriteActionResponseOutput{})
	pulumi.RegisterOutputType(UrlSigningActionOutput{})
	pulumi.RegisterOutputType(UrlSigningActionParametersOutput{})
	pulumi.RegisterOutputType(UrlSigningActionParametersResponseOutput{})
	pulumi.RegisterOutputType(UrlSigningActionResponseOutput{})
	pulumi.RegisterOutputType(UrlSigningKeyOutput{})
	pulumi.RegisterOutputType(UrlSigningKeyArrayOutput{})
	pulumi.RegisterOutputType(UrlSigningKeyResponseOutput{})
	pulumi.RegisterOutputType(UrlSigningKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(UrlSigningParamIdentifierOutput{})
	pulumi.RegisterOutputType(UrlSigningParamIdentifierArrayOutput{})
	pulumi.RegisterOutputType(UrlSigningParamIdentifierResponseOutput{})
	pulumi.RegisterOutputType(UrlSigningParamIdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(UserManagedHttpsParametersResponseOutput{})
}
