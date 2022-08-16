


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AFDDomainHttpsParameters struct {
	CertificateType   string                `pulumi:"certificateType"`
	MinimumTlsVersion *AfdMinimumTlsVersion `pulumi:"minimumTlsVersion"`
	Secret            *ResourceReference    `pulumi:"secret"`
}





type AFDDomainHttpsParametersInput interface {
	pulumi.Input

	ToAFDDomainHttpsParametersOutput() AFDDomainHttpsParametersOutput
	ToAFDDomainHttpsParametersOutputWithContext(context.Context) AFDDomainHttpsParametersOutput
}

type AFDDomainHttpsParametersArgs struct {
	CertificateType   pulumi.StringInput           `pulumi:"certificateType"`
	MinimumTlsVersion AfdMinimumTlsVersionPtrInput `pulumi:"minimumTlsVersion"`
	Secret            ResourceReferencePtrInput    `pulumi:"secret"`
}

func (AFDDomainHttpsParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AFDDomainHttpsParameters)(nil)).Elem()
}

func (i AFDDomainHttpsParametersArgs) ToAFDDomainHttpsParametersOutput() AFDDomainHttpsParametersOutput {
	return i.ToAFDDomainHttpsParametersOutputWithContext(context.Background())
}

func (i AFDDomainHttpsParametersArgs) ToAFDDomainHttpsParametersOutputWithContext(ctx context.Context) AFDDomainHttpsParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDDomainHttpsParametersOutput)
}

func (i AFDDomainHttpsParametersArgs) ToAFDDomainHttpsParametersPtrOutput() AFDDomainHttpsParametersPtrOutput {
	return i.ToAFDDomainHttpsParametersPtrOutputWithContext(context.Background())
}

func (i AFDDomainHttpsParametersArgs) ToAFDDomainHttpsParametersPtrOutputWithContext(ctx context.Context) AFDDomainHttpsParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDDomainHttpsParametersOutput).ToAFDDomainHttpsParametersPtrOutputWithContext(ctx)
}









type AFDDomainHttpsParametersPtrInput interface {
	pulumi.Input

	ToAFDDomainHttpsParametersPtrOutput() AFDDomainHttpsParametersPtrOutput
	ToAFDDomainHttpsParametersPtrOutputWithContext(context.Context) AFDDomainHttpsParametersPtrOutput
}

type afddomainHttpsParametersPtrType AFDDomainHttpsParametersArgs

func AFDDomainHttpsParametersPtr(v *AFDDomainHttpsParametersArgs) AFDDomainHttpsParametersPtrInput {
	return (*afddomainHttpsParametersPtrType)(v)
}

func (*afddomainHttpsParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDDomainHttpsParameters)(nil)).Elem()
}

func (i *afddomainHttpsParametersPtrType) ToAFDDomainHttpsParametersPtrOutput() AFDDomainHttpsParametersPtrOutput {
	return i.ToAFDDomainHttpsParametersPtrOutputWithContext(context.Background())
}

func (i *afddomainHttpsParametersPtrType) ToAFDDomainHttpsParametersPtrOutputWithContext(ctx context.Context) AFDDomainHttpsParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDDomainHttpsParametersPtrOutput)
}

type AFDDomainHttpsParametersOutput struct{ *pulumi.OutputState }

func (AFDDomainHttpsParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AFDDomainHttpsParameters)(nil)).Elem()
}

func (o AFDDomainHttpsParametersOutput) ToAFDDomainHttpsParametersOutput() AFDDomainHttpsParametersOutput {
	return o
}

func (o AFDDomainHttpsParametersOutput) ToAFDDomainHttpsParametersOutputWithContext(ctx context.Context) AFDDomainHttpsParametersOutput {
	return o
}

func (o AFDDomainHttpsParametersOutput) ToAFDDomainHttpsParametersPtrOutput() AFDDomainHttpsParametersPtrOutput {
	return o.ToAFDDomainHttpsParametersPtrOutputWithContext(context.Background())
}

func (o AFDDomainHttpsParametersOutput) ToAFDDomainHttpsParametersPtrOutputWithContext(ctx context.Context) AFDDomainHttpsParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AFDDomainHttpsParameters) *AFDDomainHttpsParameters {
		return &v
	}).(AFDDomainHttpsParametersPtrOutput)
}

func (o AFDDomainHttpsParametersOutput) CertificateType() pulumi.StringOutput {
	return o.ApplyT(func(v AFDDomainHttpsParameters) string { return v.CertificateType }).(pulumi.StringOutput)
}

func (o AFDDomainHttpsParametersOutput) MinimumTlsVersion() AfdMinimumTlsVersionPtrOutput {
	return o.ApplyT(func(v AFDDomainHttpsParameters) *AfdMinimumTlsVersion { return v.MinimumTlsVersion }).(AfdMinimumTlsVersionPtrOutput)
}

func (o AFDDomainHttpsParametersOutput) Secret() ResourceReferencePtrOutput {
	return o.ApplyT(func(v AFDDomainHttpsParameters) *ResourceReference { return v.Secret }).(ResourceReferencePtrOutput)
}

type AFDDomainHttpsParametersPtrOutput struct{ *pulumi.OutputState }

func (AFDDomainHttpsParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDDomainHttpsParameters)(nil)).Elem()
}

func (o AFDDomainHttpsParametersPtrOutput) ToAFDDomainHttpsParametersPtrOutput() AFDDomainHttpsParametersPtrOutput {
	return o
}

func (o AFDDomainHttpsParametersPtrOutput) ToAFDDomainHttpsParametersPtrOutputWithContext(ctx context.Context) AFDDomainHttpsParametersPtrOutput {
	return o
}

func (o AFDDomainHttpsParametersPtrOutput) Elem() AFDDomainHttpsParametersOutput {
	return o.ApplyT(func(v *AFDDomainHttpsParameters) AFDDomainHttpsParameters {
		if v != nil {
			return *v
		}
		var ret AFDDomainHttpsParameters
		return ret
	}).(AFDDomainHttpsParametersOutput)
}

func (o AFDDomainHttpsParametersPtrOutput) CertificateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AFDDomainHttpsParameters) *string {
		if v == nil {
			return nil
		}
		return &v.CertificateType
	}).(pulumi.StringPtrOutput)
}

func (o AFDDomainHttpsParametersPtrOutput) MinimumTlsVersion() AfdMinimumTlsVersionPtrOutput {
	return o.ApplyT(func(v *AFDDomainHttpsParameters) *AfdMinimumTlsVersion {
		if v == nil {
			return nil
		}
		return v.MinimumTlsVersion
	}).(AfdMinimumTlsVersionPtrOutput)
}

func (o AFDDomainHttpsParametersPtrOutput) Secret() ResourceReferencePtrOutput {
	return o.ApplyT(func(v *AFDDomainHttpsParameters) *ResourceReference {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(ResourceReferencePtrOutput)
}

type AFDDomainHttpsParametersResponse struct {
	CertificateType   string                     `pulumi:"certificateType"`
	MinimumTlsVersion *string                    `pulumi:"minimumTlsVersion"`
	Secret            *ResourceReferenceResponse `pulumi:"secret"`
}

type AFDDomainHttpsParametersResponseOutput struct{ *pulumi.OutputState }

func (AFDDomainHttpsParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AFDDomainHttpsParametersResponse)(nil)).Elem()
}

func (o AFDDomainHttpsParametersResponseOutput) ToAFDDomainHttpsParametersResponseOutput() AFDDomainHttpsParametersResponseOutput {
	return o
}

func (o AFDDomainHttpsParametersResponseOutput) ToAFDDomainHttpsParametersResponseOutputWithContext(ctx context.Context) AFDDomainHttpsParametersResponseOutput {
	return o
}

func (o AFDDomainHttpsParametersResponseOutput) CertificateType() pulumi.StringOutput {
	return o.ApplyT(func(v AFDDomainHttpsParametersResponse) string { return v.CertificateType }).(pulumi.StringOutput)
}

func (o AFDDomainHttpsParametersResponseOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AFDDomainHttpsParametersResponse) *string { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

func (o AFDDomainHttpsParametersResponseOutput) Secret() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v AFDDomainHttpsParametersResponse) *ResourceReferenceResponse { return v.Secret }).(ResourceReferenceResponsePtrOutput)
}

type AFDDomainHttpsParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (AFDDomainHttpsParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDDomainHttpsParametersResponse)(nil)).Elem()
}

func (o AFDDomainHttpsParametersResponsePtrOutput) ToAFDDomainHttpsParametersResponsePtrOutput() AFDDomainHttpsParametersResponsePtrOutput {
	return o
}

func (o AFDDomainHttpsParametersResponsePtrOutput) ToAFDDomainHttpsParametersResponsePtrOutputWithContext(ctx context.Context) AFDDomainHttpsParametersResponsePtrOutput {
	return o
}

func (o AFDDomainHttpsParametersResponsePtrOutput) Elem() AFDDomainHttpsParametersResponseOutput {
	return o.ApplyT(func(v *AFDDomainHttpsParametersResponse) AFDDomainHttpsParametersResponse {
		if v != nil {
			return *v
		}
		var ret AFDDomainHttpsParametersResponse
		return ret
	}).(AFDDomainHttpsParametersResponseOutput)
}

func (o AFDDomainHttpsParametersResponsePtrOutput) CertificateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AFDDomainHttpsParametersResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CertificateType
	}).(pulumi.StringPtrOutput)
}

func (o AFDDomainHttpsParametersResponsePtrOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AFDDomainHttpsParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.MinimumTlsVersion
	}).(pulumi.StringPtrOutput)
}

func (o AFDDomainHttpsParametersResponsePtrOutput) Secret() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *AFDDomainHttpsParametersResponse) *ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(ResourceReferenceResponsePtrOutput)
}

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

type CompressionSettings struct {
	ContentTypesToCompress []string `pulumi:"contentTypesToCompress"`
	IsCompressionEnabled   *bool    `pulumi:"isCompressionEnabled"`
}





type CompressionSettingsInput interface {
	pulumi.Input

	ToCompressionSettingsOutput() CompressionSettingsOutput
	ToCompressionSettingsOutputWithContext(context.Context) CompressionSettingsOutput
}

type CompressionSettingsArgs struct {
	ContentTypesToCompress pulumi.StringArrayInput `pulumi:"contentTypesToCompress"`
	IsCompressionEnabled   pulumi.BoolPtrInput     `pulumi:"isCompressionEnabled"`
}

func (CompressionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CompressionSettings)(nil)).Elem()
}

func (i CompressionSettingsArgs) ToCompressionSettingsOutput() CompressionSettingsOutput {
	return i.ToCompressionSettingsOutputWithContext(context.Background())
}

func (i CompressionSettingsArgs) ToCompressionSettingsOutputWithContext(ctx context.Context) CompressionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompressionSettingsOutput)
}

func (i CompressionSettingsArgs) ToCompressionSettingsPtrOutput() CompressionSettingsPtrOutput {
	return i.ToCompressionSettingsPtrOutputWithContext(context.Background())
}

func (i CompressionSettingsArgs) ToCompressionSettingsPtrOutputWithContext(ctx context.Context) CompressionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompressionSettingsOutput).ToCompressionSettingsPtrOutputWithContext(ctx)
}









type CompressionSettingsPtrInput interface {
	pulumi.Input

	ToCompressionSettingsPtrOutput() CompressionSettingsPtrOutput
	ToCompressionSettingsPtrOutputWithContext(context.Context) CompressionSettingsPtrOutput
}

type compressionSettingsPtrType CompressionSettingsArgs

func CompressionSettingsPtr(v *CompressionSettingsArgs) CompressionSettingsPtrInput {
	return (*compressionSettingsPtrType)(v)
}

func (*compressionSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CompressionSettings)(nil)).Elem()
}

func (i *compressionSettingsPtrType) ToCompressionSettingsPtrOutput() CompressionSettingsPtrOutput {
	return i.ToCompressionSettingsPtrOutputWithContext(context.Background())
}

func (i *compressionSettingsPtrType) ToCompressionSettingsPtrOutputWithContext(ctx context.Context) CompressionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompressionSettingsPtrOutput)
}

type CompressionSettingsOutput struct{ *pulumi.OutputState }

func (CompressionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompressionSettings)(nil)).Elem()
}

func (o CompressionSettingsOutput) ToCompressionSettingsOutput() CompressionSettingsOutput {
	return o
}

func (o CompressionSettingsOutput) ToCompressionSettingsOutputWithContext(ctx context.Context) CompressionSettingsOutput {
	return o
}

func (o CompressionSettingsOutput) ToCompressionSettingsPtrOutput() CompressionSettingsPtrOutput {
	return o.ToCompressionSettingsPtrOutputWithContext(context.Background())
}

func (o CompressionSettingsOutput) ToCompressionSettingsPtrOutputWithContext(ctx context.Context) CompressionSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CompressionSettings) *CompressionSettings {
		return &v
	}).(CompressionSettingsPtrOutput)
}

func (o CompressionSettingsOutput) ContentTypesToCompress() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CompressionSettings) []string { return v.ContentTypesToCompress }).(pulumi.StringArrayOutput)
}

func (o CompressionSettingsOutput) IsCompressionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CompressionSettings) *bool { return v.IsCompressionEnabled }).(pulumi.BoolPtrOutput)
}

type CompressionSettingsPtrOutput struct{ *pulumi.OutputState }

func (CompressionSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CompressionSettings)(nil)).Elem()
}

func (o CompressionSettingsPtrOutput) ToCompressionSettingsPtrOutput() CompressionSettingsPtrOutput {
	return o
}

func (o CompressionSettingsPtrOutput) ToCompressionSettingsPtrOutputWithContext(ctx context.Context) CompressionSettingsPtrOutput {
	return o
}

func (o CompressionSettingsPtrOutput) Elem() CompressionSettingsOutput {
	return o.ApplyT(func(v *CompressionSettings) CompressionSettings {
		if v != nil {
			return *v
		}
		var ret CompressionSettings
		return ret
	}).(CompressionSettingsOutput)
}

func (o CompressionSettingsPtrOutput) ContentTypesToCompress() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CompressionSettings) []string {
		if v == nil {
			return nil
		}
		return v.ContentTypesToCompress
	}).(pulumi.StringArrayOutput)
}

func (o CompressionSettingsPtrOutput) IsCompressionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CompressionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.IsCompressionEnabled
	}).(pulumi.BoolPtrOutput)
}

type CompressionSettingsResponse struct {
	ContentTypesToCompress []string `pulumi:"contentTypesToCompress"`
	IsCompressionEnabled   *bool    `pulumi:"isCompressionEnabled"`
}

type CompressionSettingsResponseOutput struct{ *pulumi.OutputState }

func (CompressionSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompressionSettingsResponse)(nil)).Elem()
}

func (o CompressionSettingsResponseOutput) ToCompressionSettingsResponseOutput() CompressionSettingsResponseOutput {
	return o
}

func (o CompressionSettingsResponseOutput) ToCompressionSettingsResponseOutputWithContext(ctx context.Context) CompressionSettingsResponseOutput {
	return o
}

func (o CompressionSettingsResponseOutput) ContentTypesToCompress() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CompressionSettingsResponse) []string { return v.ContentTypesToCompress }).(pulumi.StringArrayOutput)
}

func (o CompressionSettingsResponseOutput) IsCompressionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CompressionSettingsResponse) *bool { return v.IsCompressionEnabled }).(pulumi.BoolPtrOutput)
}

type CompressionSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CompressionSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CompressionSettingsResponse)(nil)).Elem()
}

func (o CompressionSettingsResponsePtrOutput) ToCompressionSettingsResponsePtrOutput() CompressionSettingsResponsePtrOutput {
	return o
}

func (o CompressionSettingsResponsePtrOutput) ToCompressionSettingsResponsePtrOutputWithContext(ctx context.Context) CompressionSettingsResponsePtrOutput {
	return o
}

func (o CompressionSettingsResponsePtrOutput) Elem() CompressionSettingsResponseOutput {
	return o.ApplyT(func(v *CompressionSettingsResponse) CompressionSettingsResponse {
		if v != nil {
			return *v
		}
		var ret CompressionSettingsResponse
		return ret
	}).(CompressionSettingsResponseOutput)
}

func (o CompressionSettingsResponsePtrOutput) ContentTypesToCompress() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CompressionSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.ContentTypesToCompress
	}).(pulumi.StringArrayOutput)
}

func (o CompressionSettingsResponsePtrOutput) IsCompressionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CompressionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCompressionEnabled
	}).(pulumi.BoolPtrOutput)
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

type CustomerCertificateParameters struct {
	CertificateAuthority    *string           `pulumi:"certificateAuthority"`
	SecretSource            ResourceReference `pulumi:"secretSource"`
	SecretVersion           *string           `pulumi:"secretVersion"`
	SubjectAlternativeNames []string          `pulumi:"subjectAlternativeNames"`
	Type                    string            `pulumi:"type"`
	UseLatestVersion        *bool             `pulumi:"useLatestVersion"`
}

type CustomerCertificateParametersResponse struct {
	CertificateAuthority    *string                   `pulumi:"certificateAuthority"`
	SecretSource            ResourceReferenceResponse `pulumi:"secretSource"`
	SecretVersion           *string                   `pulumi:"secretVersion"`
	SubjectAlternativeNames []string                  `pulumi:"subjectAlternativeNames"`
	Type                    string                    `pulumi:"type"`
	UseLatestVersion        *bool                     `pulumi:"useLatestVersion"`
}

type DeepCreatedOrigin struct {
	Enabled                    *bool   `pulumi:"enabled"`
	HostName                   string  `pulumi:"hostName"`
	HttpPort                   *int    `pulumi:"httpPort"`
	HttpsPort                  *int    `pulumi:"httpsPort"`
	Name                       string  `pulumi:"name"`
	OriginHostHeader           *string `pulumi:"originHostHeader"`
	Priority                   *int    `pulumi:"priority"`
	PrivateLinkAlias           *string `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage *string `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        *string `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      *string `pulumi:"privateLinkResourceId"`
	Weight                     *int    `pulumi:"weight"`
}





type DeepCreatedOriginInput interface {
	pulumi.Input

	ToDeepCreatedOriginOutput() DeepCreatedOriginOutput
	ToDeepCreatedOriginOutputWithContext(context.Context) DeepCreatedOriginOutput
}

type DeepCreatedOriginArgs struct {
	Enabled                    pulumi.BoolPtrInput   `pulumi:"enabled"`
	HostName                   pulumi.StringInput    `pulumi:"hostName"`
	HttpPort                   pulumi.IntPtrInput    `pulumi:"httpPort"`
	HttpsPort                  pulumi.IntPtrInput    `pulumi:"httpsPort"`
	Name                       pulumi.StringInput    `pulumi:"name"`
	OriginHostHeader           pulumi.StringPtrInput `pulumi:"originHostHeader"`
	Priority                   pulumi.IntPtrInput    `pulumi:"priority"`
	PrivateLinkAlias           pulumi.StringPtrInput `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage pulumi.StringPtrInput `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        pulumi.StringPtrInput `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      pulumi.StringPtrInput `pulumi:"privateLinkResourceId"`
	Weight                     pulumi.IntPtrInput    `pulumi:"weight"`
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

func (o DeepCreatedOriginOutput) PrivateLinkAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *string { return v.PrivateLinkAlias }).(pulumi.StringPtrOutput)
}

func (o DeepCreatedOriginOutput) PrivateLinkApprovalMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *string { return v.PrivateLinkApprovalMessage }).(pulumi.StringPtrOutput)
}

func (o DeepCreatedOriginOutput) PrivateLinkLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *string { return v.PrivateLinkLocation }).(pulumi.StringPtrOutput)
}

func (o DeepCreatedOriginOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
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
	Enabled                    *bool   `pulumi:"enabled"`
	HostName                   string  `pulumi:"hostName"`
	HttpPort                   *int    `pulumi:"httpPort"`
	HttpsPort                  *int    `pulumi:"httpsPort"`
	Name                       string  `pulumi:"name"`
	OriginHostHeader           *string `pulumi:"originHostHeader"`
	Priority                   *int    `pulumi:"priority"`
	PrivateLinkAlias           *string `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage *string `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        *string `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      *string `pulumi:"privateLinkResourceId"`
	Weight                     *int    `pulumi:"weight"`
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

func (o DeepCreatedOriginResponseOutput) PrivateLinkAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *string { return v.PrivateLinkAlias }).(pulumi.StringPtrOutput)
}

func (o DeepCreatedOriginResponseOutput) PrivateLinkApprovalMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *string { return v.PrivateLinkApprovalMessage }).(pulumi.StringPtrOutput)
}

func (o DeepCreatedOriginResponseOutput) PrivateLinkLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *string { return v.PrivateLinkLocation }).(pulumi.StringPtrOutput)
}

func (o DeepCreatedOriginResponseOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
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

type DomainValidationPropertiesResponse struct {
	ExpirationDate  string `pulumi:"expirationDate"`
	ValidationToken string `pulumi:"validationToken"`
}

type DomainValidationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DomainValidationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainValidationPropertiesResponse)(nil)).Elem()
}

func (o DomainValidationPropertiesResponseOutput) ToDomainValidationPropertiesResponseOutput() DomainValidationPropertiesResponseOutput {
	return o
}

func (o DomainValidationPropertiesResponseOutput) ToDomainValidationPropertiesResponseOutputWithContext(ctx context.Context) DomainValidationPropertiesResponseOutput {
	return o
}

func (o DomainValidationPropertiesResponseOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v DomainValidationPropertiesResponse) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o DomainValidationPropertiesResponseOutput) ValidationToken() pulumi.StringOutput {
	return o.ApplyT(func(v DomainValidationPropertiesResponse) string { return v.ValidationToken }).(pulumi.StringOutput)
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

type LoadBalancingSettingsParameters struct {
	AdditionalLatencyInMilliseconds *int `pulumi:"additionalLatencyInMilliseconds"`
	SampleSize                      *int `pulumi:"sampleSize"`
	SuccessfulSamplesRequired       *int `pulumi:"successfulSamplesRequired"`
}





type LoadBalancingSettingsParametersInput interface {
	pulumi.Input

	ToLoadBalancingSettingsParametersOutput() LoadBalancingSettingsParametersOutput
	ToLoadBalancingSettingsParametersOutputWithContext(context.Context) LoadBalancingSettingsParametersOutput
}

type LoadBalancingSettingsParametersArgs struct {
	AdditionalLatencyInMilliseconds pulumi.IntPtrInput `pulumi:"additionalLatencyInMilliseconds"`
	SampleSize                      pulumi.IntPtrInput `pulumi:"sampleSize"`
	SuccessfulSamplesRequired       pulumi.IntPtrInput `pulumi:"successfulSamplesRequired"`
}

func (LoadBalancingSettingsParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingSettingsParameters)(nil)).Elem()
}

func (i LoadBalancingSettingsParametersArgs) ToLoadBalancingSettingsParametersOutput() LoadBalancingSettingsParametersOutput {
	return i.ToLoadBalancingSettingsParametersOutputWithContext(context.Background())
}

func (i LoadBalancingSettingsParametersArgs) ToLoadBalancingSettingsParametersOutputWithContext(ctx context.Context) LoadBalancingSettingsParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingSettingsParametersOutput)
}

func (i LoadBalancingSettingsParametersArgs) ToLoadBalancingSettingsParametersPtrOutput() LoadBalancingSettingsParametersPtrOutput {
	return i.ToLoadBalancingSettingsParametersPtrOutputWithContext(context.Background())
}

func (i LoadBalancingSettingsParametersArgs) ToLoadBalancingSettingsParametersPtrOutputWithContext(ctx context.Context) LoadBalancingSettingsParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingSettingsParametersOutput).ToLoadBalancingSettingsParametersPtrOutputWithContext(ctx)
}









type LoadBalancingSettingsParametersPtrInput interface {
	pulumi.Input

	ToLoadBalancingSettingsParametersPtrOutput() LoadBalancingSettingsParametersPtrOutput
	ToLoadBalancingSettingsParametersPtrOutputWithContext(context.Context) LoadBalancingSettingsParametersPtrOutput
}

type loadBalancingSettingsParametersPtrType LoadBalancingSettingsParametersArgs

func LoadBalancingSettingsParametersPtr(v *LoadBalancingSettingsParametersArgs) LoadBalancingSettingsParametersPtrInput {
	return (*loadBalancingSettingsParametersPtrType)(v)
}

func (*loadBalancingSettingsParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancingSettingsParameters)(nil)).Elem()
}

func (i *loadBalancingSettingsParametersPtrType) ToLoadBalancingSettingsParametersPtrOutput() LoadBalancingSettingsParametersPtrOutput {
	return i.ToLoadBalancingSettingsParametersPtrOutputWithContext(context.Background())
}

func (i *loadBalancingSettingsParametersPtrType) ToLoadBalancingSettingsParametersPtrOutputWithContext(ctx context.Context) LoadBalancingSettingsParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingSettingsParametersPtrOutput)
}

type LoadBalancingSettingsParametersOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingSettingsParameters)(nil)).Elem()
}

func (o LoadBalancingSettingsParametersOutput) ToLoadBalancingSettingsParametersOutput() LoadBalancingSettingsParametersOutput {
	return o
}

func (o LoadBalancingSettingsParametersOutput) ToLoadBalancingSettingsParametersOutputWithContext(ctx context.Context) LoadBalancingSettingsParametersOutput {
	return o
}

func (o LoadBalancingSettingsParametersOutput) ToLoadBalancingSettingsParametersPtrOutput() LoadBalancingSettingsParametersPtrOutput {
	return o.ToLoadBalancingSettingsParametersPtrOutputWithContext(context.Background())
}

func (o LoadBalancingSettingsParametersOutput) ToLoadBalancingSettingsParametersPtrOutputWithContext(ctx context.Context) LoadBalancingSettingsParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancingSettingsParameters) *LoadBalancingSettingsParameters {
		return &v
	}).(LoadBalancingSettingsParametersPtrOutput)
}

func (o LoadBalancingSettingsParametersOutput) AdditionalLatencyInMilliseconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsParameters) *int { return v.AdditionalLatencyInMilliseconds }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsParametersOutput) SampleSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsParameters) *int { return v.SampleSize }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsParametersOutput) SuccessfulSamplesRequired() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsParameters) *int { return v.SuccessfulSamplesRequired }).(pulumi.IntPtrOutput)
}

type LoadBalancingSettingsParametersPtrOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancingSettingsParameters)(nil)).Elem()
}

func (o LoadBalancingSettingsParametersPtrOutput) ToLoadBalancingSettingsParametersPtrOutput() LoadBalancingSettingsParametersPtrOutput {
	return o
}

func (o LoadBalancingSettingsParametersPtrOutput) ToLoadBalancingSettingsParametersPtrOutputWithContext(ctx context.Context) LoadBalancingSettingsParametersPtrOutput {
	return o
}

func (o LoadBalancingSettingsParametersPtrOutput) Elem() LoadBalancingSettingsParametersOutput {
	return o.ApplyT(func(v *LoadBalancingSettingsParameters) LoadBalancingSettingsParameters {
		if v != nil {
			return *v
		}
		var ret LoadBalancingSettingsParameters
		return ret
	}).(LoadBalancingSettingsParametersOutput)
}

func (o LoadBalancingSettingsParametersPtrOutput) AdditionalLatencyInMilliseconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancingSettingsParameters) *int {
		if v == nil {
			return nil
		}
		return v.AdditionalLatencyInMilliseconds
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsParametersPtrOutput) SampleSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancingSettingsParameters) *int {
		if v == nil {
			return nil
		}
		return v.SampleSize
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsParametersPtrOutput) SuccessfulSamplesRequired() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancingSettingsParameters) *int {
		if v == nil {
			return nil
		}
		return v.SuccessfulSamplesRequired
	}).(pulumi.IntPtrOutput)
}

type LoadBalancingSettingsParametersResponse struct {
	AdditionalLatencyInMilliseconds *int `pulumi:"additionalLatencyInMilliseconds"`
	SampleSize                      *int `pulumi:"sampleSize"`
	SuccessfulSamplesRequired       *int `pulumi:"successfulSamplesRequired"`
}

type LoadBalancingSettingsParametersResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingSettingsParametersResponse)(nil)).Elem()
}

func (o LoadBalancingSettingsParametersResponseOutput) ToLoadBalancingSettingsParametersResponseOutput() LoadBalancingSettingsParametersResponseOutput {
	return o
}

func (o LoadBalancingSettingsParametersResponseOutput) ToLoadBalancingSettingsParametersResponseOutputWithContext(ctx context.Context) LoadBalancingSettingsParametersResponseOutput {
	return o
}

func (o LoadBalancingSettingsParametersResponseOutput) AdditionalLatencyInMilliseconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsParametersResponse) *int { return v.AdditionalLatencyInMilliseconds }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsParametersResponseOutput) SampleSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsParametersResponse) *int { return v.SampleSize }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsParametersResponseOutput) SuccessfulSamplesRequired() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsParametersResponse) *int { return v.SuccessfulSamplesRequired }).(pulumi.IntPtrOutput)
}

type LoadBalancingSettingsParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancingSettingsParametersResponse)(nil)).Elem()
}

func (o LoadBalancingSettingsParametersResponsePtrOutput) ToLoadBalancingSettingsParametersResponsePtrOutput() LoadBalancingSettingsParametersResponsePtrOutput {
	return o
}

func (o LoadBalancingSettingsParametersResponsePtrOutput) ToLoadBalancingSettingsParametersResponsePtrOutputWithContext(ctx context.Context) LoadBalancingSettingsParametersResponsePtrOutput {
	return o
}

func (o LoadBalancingSettingsParametersResponsePtrOutput) Elem() LoadBalancingSettingsParametersResponseOutput {
	return o.ApplyT(func(v *LoadBalancingSettingsParametersResponse) LoadBalancingSettingsParametersResponse {
		if v != nil {
			return *v
		}
		var ret LoadBalancingSettingsParametersResponse
		return ret
	}).(LoadBalancingSettingsParametersResponseOutput)
}

func (o LoadBalancingSettingsParametersResponsePtrOutput) AdditionalLatencyInMilliseconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancingSettingsParametersResponse) *int {
		if v == nil {
			return nil
		}
		return v.AdditionalLatencyInMilliseconds
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsParametersResponsePtrOutput) SampleSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancingSettingsParametersResponse) *int {
		if v == nil {
			return nil
		}
		return v.SampleSize
	}).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsParametersResponsePtrOutput) SuccessfulSamplesRequired() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoadBalancingSettingsParametersResponse) *int {
		if v == nil {
			return nil
		}
		return v.SuccessfulSamplesRequired
	}).(pulumi.IntPtrOutput)
}

type ManagedCertificateParameters struct {
	Type string `pulumi:"type"`
}

type ManagedCertificateParametersResponse struct {
	Type string `pulumi:"type"`
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

type SecurityPolicyWebApplicationFirewallAssociation struct {
	Domains         []ResourceReference `pulumi:"domains"`
	PatternsToMatch []string            `pulumi:"patternsToMatch"`
}





type SecurityPolicyWebApplicationFirewallAssociationInput interface {
	pulumi.Input

	ToSecurityPolicyWebApplicationFirewallAssociationOutput() SecurityPolicyWebApplicationFirewallAssociationOutput
	ToSecurityPolicyWebApplicationFirewallAssociationOutputWithContext(context.Context) SecurityPolicyWebApplicationFirewallAssociationOutput
}

type SecurityPolicyWebApplicationFirewallAssociationArgs struct {
	Domains         ResourceReferenceArrayInput `pulumi:"domains"`
	PatternsToMatch pulumi.StringArrayInput     `pulumi:"patternsToMatch"`
}

func (SecurityPolicyWebApplicationFirewallAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPolicyWebApplicationFirewallAssociation)(nil)).Elem()
}

func (i SecurityPolicyWebApplicationFirewallAssociationArgs) ToSecurityPolicyWebApplicationFirewallAssociationOutput() SecurityPolicyWebApplicationFirewallAssociationOutput {
	return i.ToSecurityPolicyWebApplicationFirewallAssociationOutputWithContext(context.Background())
}

func (i SecurityPolicyWebApplicationFirewallAssociationArgs) ToSecurityPolicyWebApplicationFirewallAssociationOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyWebApplicationFirewallAssociationOutput)
}





type SecurityPolicyWebApplicationFirewallAssociationArrayInput interface {
	pulumi.Input

	ToSecurityPolicyWebApplicationFirewallAssociationArrayOutput() SecurityPolicyWebApplicationFirewallAssociationArrayOutput
	ToSecurityPolicyWebApplicationFirewallAssociationArrayOutputWithContext(context.Context) SecurityPolicyWebApplicationFirewallAssociationArrayOutput
}

type SecurityPolicyWebApplicationFirewallAssociationArray []SecurityPolicyWebApplicationFirewallAssociationInput

func (SecurityPolicyWebApplicationFirewallAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityPolicyWebApplicationFirewallAssociation)(nil)).Elem()
}

func (i SecurityPolicyWebApplicationFirewallAssociationArray) ToSecurityPolicyWebApplicationFirewallAssociationArrayOutput() SecurityPolicyWebApplicationFirewallAssociationArrayOutput {
	return i.ToSecurityPolicyWebApplicationFirewallAssociationArrayOutputWithContext(context.Background())
}

func (i SecurityPolicyWebApplicationFirewallAssociationArray) ToSecurityPolicyWebApplicationFirewallAssociationArrayOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyWebApplicationFirewallAssociationArrayOutput)
}

type SecurityPolicyWebApplicationFirewallAssociationOutput struct{ *pulumi.OutputState }

func (SecurityPolicyWebApplicationFirewallAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPolicyWebApplicationFirewallAssociation)(nil)).Elem()
}

func (o SecurityPolicyWebApplicationFirewallAssociationOutput) ToSecurityPolicyWebApplicationFirewallAssociationOutput() SecurityPolicyWebApplicationFirewallAssociationOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallAssociationOutput) ToSecurityPolicyWebApplicationFirewallAssociationOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallAssociationOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallAssociationOutput) Domains() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v SecurityPolicyWebApplicationFirewallAssociation) []ResourceReference { return v.Domains }).(ResourceReferenceArrayOutput)
}

func (o SecurityPolicyWebApplicationFirewallAssociationOutput) PatternsToMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityPolicyWebApplicationFirewallAssociation) []string { return v.PatternsToMatch }).(pulumi.StringArrayOutput)
}

type SecurityPolicyWebApplicationFirewallAssociationArrayOutput struct{ *pulumi.OutputState }

func (SecurityPolicyWebApplicationFirewallAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityPolicyWebApplicationFirewallAssociation)(nil)).Elem()
}

func (o SecurityPolicyWebApplicationFirewallAssociationArrayOutput) ToSecurityPolicyWebApplicationFirewallAssociationArrayOutput() SecurityPolicyWebApplicationFirewallAssociationArrayOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallAssociationArrayOutput) ToSecurityPolicyWebApplicationFirewallAssociationArrayOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallAssociationArrayOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallAssociationArrayOutput) Index(i pulumi.IntInput) SecurityPolicyWebApplicationFirewallAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityPolicyWebApplicationFirewallAssociation {
		return vs[0].([]SecurityPolicyWebApplicationFirewallAssociation)[vs[1].(int)]
	}).(SecurityPolicyWebApplicationFirewallAssociationOutput)
}

type SecurityPolicyWebApplicationFirewallAssociationResponse struct {
	Domains         []ResourceReferenceResponse `pulumi:"domains"`
	PatternsToMatch []string                    `pulumi:"patternsToMatch"`
}

type SecurityPolicyWebApplicationFirewallAssociationResponseOutput struct{ *pulumi.OutputState }

func (SecurityPolicyWebApplicationFirewallAssociationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPolicyWebApplicationFirewallAssociationResponse)(nil)).Elem()
}

func (o SecurityPolicyWebApplicationFirewallAssociationResponseOutput) ToSecurityPolicyWebApplicationFirewallAssociationResponseOutput() SecurityPolicyWebApplicationFirewallAssociationResponseOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallAssociationResponseOutput) ToSecurityPolicyWebApplicationFirewallAssociationResponseOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallAssociationResponseOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallAssociationResponseOutput) Domains() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v SecurityPolicyWebApplicationFirewallAssociationResponse) []ResourceReferenceResponse {
		return v.Domains
	}).(ResourceReferenceResponseArrayOutput)
}

func (o SecurityPolicyWebApplicationFirewallAssociationResponseOutput) PatternsToMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityPolicyWebApplicationFirewallAssociationResponse) []string { return v.PatternsToMatch }).(pulumi.StringArrayOutput)
}

type SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput struct{ *pulumi.OutputState }

func (SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityPolicyWebApplicationFirewallAssociationResponse)(nil)).Elem()
}

func (o SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput) ToSecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput() SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput) ToSecurityPolicyWebApplicationFirewallAssociationResponseArrayOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput) Index(i pulumi.IntInput) SecurityPolicyWebApplicationFirewallAssociationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityPolicyWebApplicationFirewallAssociationResponse {
		return vs[0].([]SecurityPolicyWebApplicationFirewallAssociationResponse)[vs[1].(int)]
	}).(SecurityPolicyWebApplicationFirewallAssociationResponseOutput)
}

type SecurityPolicyWebApplicationFirewallParameters struct {
	Associations []SecurityPolicyWebApplicationFirewallAssociation `pulumi:"associations"`
	Type         string                                            `pulumi:"type"`
	WafPolicy    *ResourceReference                                `pulumi:"wafPolicy"`
}





type SecurityPolicyWebApplicationFirewallParametersInput interface {
	pulumi.Input

	ToSecurityPolicyWebApplicationFirewallParametersOutput() SecurityPolicyWebApplicationFirewallParametersOutput
	ToSecurityPolicyWebApplicationFirewallParametersOutputWithContext(context.Context) SecurityPolicyWebApplicationFirewallParametersOutput
}

type SecurityPolicyWebApplicationFirewallParametersArgs struct {
	Associations SecurityPolicyWebApplicationFirewallAssociationArrayInput `pulumi:"associations"`
	Type         pulumi.StringInput                                        `pulumi:"type"`
	WafPolicy    ResourceReferencePtrInput                                 `pulumi:"wafPolicy"`
}

func (SecurityPolicyWebApplicationFirewallParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPolicyWebApplicationFirewallParameters)(nil)).Elem()
}

func (i SecurityPolicyWebApplicationFirewallParametersArgs) ToSecurityPolicyWebApplicationFirewallParametersOutput() SecurityPolicyWebApplicationFirewallParametersOutput {
	return i.ToSecurityPolicyWebApplicationFirewallParametersOutputWithContext(context.Background())
}

func (i SecurityPolicyWebApplicationFirewallParametersArgs) ToSecurityPolicyWebApplicationFirewallParametersOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyWebApplicationFirewallParametersOutput)
}

func (i SecurityPolicyWebApplicationFirewallParametersArgs) ToSecurityPolicyWebApplicationFirewallParametersPtrOutput() SecurityPolicyWebApplicationFirewallParametersPtrOutput {
	return i.ToSecurityPolicyWebApplicationFirewallParametersPtrOutputWithContext(context.Background())
}

func (i SecurityPolicyWebApplicationFirewallParametersArgs) ToSecurityPolicyWebApplicationFirewallParametersPtrOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyWebApplicationFirewallParametersOutput).ToSecurityPolicyWebApplicationFirewallParametersPtrOutputWithContext(ctx)
}









type SecurityPolicyWebApplicationFirewallParametersPtrInput interface {
	pulumi.Input

	ToSecurityPolicyWebApplicationFirewallParametersPtrOutput() SecurityPolicyWebApplicationFirewallParametersPtrOutput
	ToSecurityPolicyWebApplicationFirewallParametersPtrOutputWithContext(context.Context) SecurityPolicyWebApplicationFirewallParametersPtrOutput
}

type securityPolicyWebApplicationFirewallParametersPtrType SecurityPolicyWebApplicationFirewallParametersArgs

func SecurityPolicyWebApplicationFirewallParametersPtr(v *SecurityPolicyWebApplicationFirewallParametersArgs) SecurityPolicyWebApplicationFirewallParametersPtrInput {
	return (*securityPolicyWebApplicationFirewallParametersPtrType)(v)
}

func (*securityPolicyWebApplicationFirewallParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPolicyWebApplicationFirewallParameters)(nil)).Elem()
}

func (i *securityPolicyWebApplicationFirewallParametersPtrType) ToSecurityPolicyWebApplicationFirewallParametersPtrOutput() SecurityPolicyWebApplicationFirewallParametersPtrOutput {
	return i.ToSecurityPolicyWebApplicationFirewallParametersPtrOutputWithContext(context.Background())
}

func (i *securityPolicyWebApplicationFirewallParametersPtrType) ToSecurityPolicyWebApplicationFirewallParametersPtrOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyWebApplicationFirewallParametersPtrOutput)
}

type SecurityPolicyWebApplicationFirewallParametersOutput struct{ *pulumi.OutputState }

func (SecurityPolicyWebApplicationFirewallParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPolicyWebApplicationFirewallParameters)(nil)).Elem()
}

func (o SecurityPolicyWebApplicationFirewallParametersOutput) ToSecurityPolicyWebApplicationFirewallParametersOutput() SecurityPolicyWebApplicationFirewallParametersOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallParametersOutput) ToSecurityPolicyWebApplicationFirewallParametersOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallParametersOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallParametersOutput) ToSecurityPolicyWebApplicationFirewallParametersPtrOutput() SecurityPolicyWebApplicationFirewallParametersPtrOutput {
	return o.ToSecurityPolicyWebApplicationFirewallParametersPtrOutputWithContext(context.Background())
}

func (o SecurityPolicyWebApplicationFirewallParametersOutput) ToSecurityPolicyWebApplicationFirewallParametersPtrOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityPolicyWebApplicationFirewallParameters) *SecurityPolicyWebApplicationFirewallParameters {
		return &v
	}).(SecurityPolicyWebApplicationFirewallParametersPtrOutput)
}

func (o SecurityPolicyWebApplicationFirewallParametersOutput) Associations() SecurityPolicyWebApplicationFirewallAssociationArrayOutput {
	return o.ApplyT(func(v SecurityPolicyWebApplicationFirewallParameters) []SecurityPolicyWebApplicationFirewallAssociation {
		return v.Associations
	}).(SecurityPolicyWebApplicationFirewallAssociationArrayOutput)
}

func (o SecurityPolicyWebApplicationFirewallParametersOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityPolicyWebApplicationFirewallParameters) string { return v.Type }).(pulumi.StringOutput)
}

func (o SecurityPolicyWebApplicationFirewallParametersOutput) WafPolicy() ResourceReferencePtrOutput {
	return o.ApplyT(func(v SecurityPolicyWebApplicationFirewallParameters) *ResourceReference { return v.WafPolicy }).(ResourceReferencePtrOutput)
}

type SecurityPolicyWebApplicationFirewallParametersPtrOutput struct{ *pulumi.OutputState }

func (SecurityPolicyWebApplicationFirewallParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPolicyWebApplicationFirewallParameters)(nil)).Elem()
}

func (o SecurityPolicyWebApplicationFirewallParametersPtrOutput) ToSecurityPolicyWebApplicationFirewallParametersPtrOutput() SecurityPolicyWebApplicationFirewallParametersPtrOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallParametersPtrOutput) ToSecurityPolicyWebApplicationFirewallParametersPtrOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallParametersPtrOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallParametersPtrOutput) Elem() SecurityPolicyWebApplicationFirewallParametersOutput {
	return o.ApplyT(func(v *SecurityPolicyWebApplicationFirewallParameters) SecurityPolicyWebApplicationFirewallParameters {
		if v != nil {
			return *v
		}
		var ret SecurityPolicyWebApplicationFirewallParameters
		return ret
	}).(SecurityPolicyWebApplicationFirewallParametersOutput)
}

func (o SecurityPolicyWebApplicationFirewallParametersPtrOutput) Associations() SecurityPolicyWebApplicationFirewallAssociationArrayOutput {
	return o.ApplyT(func(v *SecurityPolicyWebApplicationFirewallParameters) []SecurityPolicyWebApplicationFirewallAssociation {
		if v == nil {
			return nil
		}
		return v.Associations
	}).(SecurityPolicyWebApplicationFirewallAssociationArrayOutput)
}

func (o SecurityPolicyWebApplicationFirewallParametersPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityPolicyWebApplicationFirewallParameters) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o SecurityPolicyWebApplicationFirewallParametersPtrOutput) WafPolicy() ResourceReferencePtrOutput {
	return o.ApplyT(func(v *SecurityPolicyWebApplicationFirewallParameters) *ResourceReference {
		if v == nil {
			return nil
		}
		return v.WafPolicy
	}).(ResourceReferencePtrOutput)
}

type SecurityPolicyWebApplicationFirewallParametersResponse struct {
	Associations []SecurityPolicyWebApplicationFirewallAssociationResponse `pulumi:"associations"`
	Type         string                                                    `pulumi:"type"`
	WafPolicy    *ResourceReferenceResponse                                `pulumi:"wafPolicy"`
}

type SecurityPolicyWebApplicationFirewallParametersResponseOutput struct{ *pulumi.OutputState }

func (SecurityPolicyWebApplicationFirewallParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPolicyWebApplicationFirewallParametersResponse)(nil)).Elem()
}

func (o SecurityPolicyWebApplicationFirewallParametersResponseOutput) ToSecurityPolicyWebApplicationFirewallParametersResponseOutput() SecurityPolicyWebApplicationFirewallParametersResponseOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallParametersResponseOutput) ToSecurityPolicyWebApplicationFirewallParametersResponseOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallParametersResponseOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallParametersResponseOutput) Associations() SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput {
	return o.ApplyT(func(v SecurityPolicyWebApplicationFirewallParametersResponse) []SecurityPolicyWebApplicationFirewallAssociationResponse {
		return v.Associations
	}).(SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput)
}

func (o SecurityPolicyWebApplicationFirewallParametersResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityPolicyWebApplicationFirewallParametersResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o SecurityPolicyWebApplicationFirewallParametersResponseOutput) WafPolicy() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v SecurityPolicyWebApplicationFirewallParametersResponse) *ResourceReferenceResponse {
		return v.WafPolicy
	}).(ResourceReferenceResponsePtrOutput)
}

type SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPolicyWebApplicationFirewallParametersResponse)(nil)).Elem()
}

func (o SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput) ToSecurityPolicyWebApplicationFirewallParametersResponsePtrOutput() SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput) ToSecurityPolicyWebApplicationFirewallParametersResponsePtrOutputWithContext(ctx context.Context) SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput {
	return o
}

func (o SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput) Elem() SecurityPolicyWebApplicationFirewallParametersResponseOutput {
	return o.ApplyT(func(v *SecurityPolicyWebApplicationFirewallParametersResponse) SecurityPolicyWebApplicationFirewallParametersResponse {
		if v != nil {
			return *v
		}
		var ret SecurityPolicyWebApplicationFirewallParametersResponse
		return ret
	}).(SecurityPolicyWebApplicationFirewallParametersResponseOutput)
}

func (o SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput) Associations() SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput {
	return o.ApplyT(func(v *SecurityPolicyWebApplicationFirewallParametersResponse) []SecurityPolicyWebApplicationFirewallAssociationResponse {
		if v == nil {
			return nil
		}
		return v.Associations
	}).(SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput)
}

func (o SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityPolicyWebApplicationFirewallParametersResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput) WafPolicy() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *SecurityPolicyWebApplicationFirewallParametersResponse) *ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.WafPolicy
	}).(ResourceReferenceResponsePtrOutput)
}

type SharedPrivateLinkResourceProperties struct {
	GroupId             *string                          `pulumi:"groupId"`
	PrivateLink         *ResourceReference               `pulumi:"privateLink"`
	PrivateLinkLocation *string                          `pulumi:"privateLinkLocation"`
	RequestMessage      *string                          `pulumi:"requestMessage"`
	Status              *SharedPrivateLinkResourceStatus `pulumi:"status"`
}





type SharedPrivateLinkResourcePropertiesInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourcePropertiesOutput() SharedPrivateLinkResourcePropertiesOutput
	ToSharedPrivateLinkResourcePropertiesOutputWithContext(context.Context) SharedPrivateLinkResourcePropertiesOutput
}

type SharedPrivateLinkResourcePropertiesArgs struct {
	GroupId             pulumi.StringPtrInput                   `pulumi:"groupId"`
	PrivateLink         ResourceReferencePtrInput               `pulumi:"privateLink"`
	PrivateLinkLocation pulumi.StringPtrInput                   `pulumi:"privateLinkLocation"`
	RequestMessage      pulumi.StringPtrInput                   `pulumi:"requestMessage"`
	Status              SharedPrivateLinkResourceStatusPtrInput `pulumi:"status"`
}

func (SharedPrivateLinkResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceProperties)(nil)).Elem()
}

func (i SharedPrivateLinkResourcePropertiesArgs) ToSharedPrivateLinkResourcePropertiesOutput() SharedPrivateLinkResourcePropertiesOutput {
	return i.ToSharedPrivateLinkResourcePropertiesOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourcePropertiesArgs) ToSharedPrivateLinkResourcePropertiesOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourcePropertiesOutput)
}

func (i SharedPrivateLinkResourcePropertiesArgs) ToSharedPrivateLinkResourcePropertiesPtrOutput() SharedPrivateLinkResourcePropertiesPtrOutput {
	return i.ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourcePropertiesArgs) ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourcePropertiesOutput).ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(ctx)
}









type SharedPrivateLinkResourcePropertiesPtrInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourcePropertiesPtrOutput() SharedPrivateLinkResourcePropertiesPtrOutput
	ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(context.Context) SharedPrivateLinkResourcePropertiesPtrOutput
}

type sharedPrivateLinkResourcePropertiesPtrType SharedPrivateLinkResourcePropertiesArgs

func SharedPrivateLinkResourcePropertiesPtr(v *SharedPrivateLinkResourcePropertiesArgs) SharedPrivateLinkResourcePropertiesPtrInput {
	return (*sharedPrivateLinkResourcePropertiesPtrType)(v)
}

func (*sharedPrivateLinkResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPrivateLinkResourceProperties)(nil)).Elem()
}

func (i *sharedPrivateLinkResourcePropertiesPtrType) ToSharedPrivateLinkResourcePropertiesPtrOutput() SharedPrivateLinkResourcePropertiesPtrOutput {
	return i.ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *sharedPrivateLinkResourcePropertiesPtrType) ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourcePropertiesPtrOutput)
}

type SharedPrivateLinkResourcePropertiesOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceProperties)(nil)).Elem()
}

func (o SharedPrivateLinkResourcePropertiesOutput) ToSharedPrivateLinkResourcePropertiesOutput() SharedPrivateLinkResourcePropertiesOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesOutput) ToSharedPrivateLinkResourcePropertiesOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesOutput) ToSharedPrivateLinkResourcePropertiesPtrOutput() SharedPrivateLinkResourcePropertiesPtrOutput {
	return o.ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o SharedPrivateLinkResourcePropertiesOutput) ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SharedPrivateLinkResourceProperties) *SharedPrivateLinkResourceProperties {
		return &v
	}).(SharedPrivateLinkResourcePropertiesPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceProperties) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesOutput) PrivateLink() ResourceReferencePtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceProperties) *ResourceReference { return v.PrivateLink }).(ResourceReferencePtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesOutput) PrivateLinkLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceProperties) *string { return v.PrivateLinkLocation }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceProperties) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesOutput) Status() SharedPrivateLinkResourceStatusPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceProperties) *SharedPrivateLinkResourceStatus { return v.Status }).(SharedPrivateLinkResourceStatusPtrOutput)
}

type SharedPrivateLinkResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPrivateLinkResourceProperties)(nil)).Elem()
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) ToSharedPrivateLinkResourcePropertiesPtrOutput() SharedPrivateLinkResourcePropertiesPtrOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesPtrOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) Elem() SharedPrivateLinkResourcePropertiesOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) SharedPrivateLinkResourceProperties {
		if v != nil {
			return *v
		}
		var ret SharedPrivateLinkResourceProperties
		return ret
	}).(SharedPrivateLinkResourcePropertiesOutput)
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.GroupId
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) PrivateLink() ResourceReferencePtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) *ResourceReference {
		if v == nil {
			return nil
		}
		return v.PrivateLink
	}).(ResourceReferencePtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) PrivateLinkLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLinkLocation
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.RequestMessage
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) Status() SharedPrivateLinkResourceStatusPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) *SharedPrivateLinkResourceStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(SharedPrivateLinkResourceStatusPtrOutput)
}

type SharedPrivateLinkResourcePropertiesResponse struct {
	GroupId             *string                    `pulumi:"groupId"`
	PrivateLink         *ResourceReferenceResponse `pulumi:"privateLink"`
	PrivateLinkLocation *string                    `pulumi:"privateLinkLocation"`
	RequestMessage      *string                    `pulumi:"requestMessage"`
	Status              *string                    `pulumi:"status"`
}

type SharedPrivateLinkResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourcePropertiesResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) ToSharedPrivateLinkResourcePropertiesResponseOutput() SharedPrivateLinkResourcePropertiesResponseOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) ToSharedPrivateLinkResourcePropertiesResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesResponseOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourcePropertiesResponse) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) PrivateLink() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourcePropertiesResponse) *ResourceReferenceResponse { return v.PrivateLink }).(ResourceReferenceResponsePtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) PrivateLinkLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourcePropertiesResponse) *string { return v.PrivateLinkLocation }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourcePropertiesResponse) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourcePropertiesResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPrivateLinkResourcePropertiesResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) ToSharedPrivateLinkResourcePropertiesResponsePtrOutput() SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) ToSharedPrivateLinkResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) Elem() SharedPrivateLinkResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) SharedPrivateLinkResourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SharedPrivateLinkResourcePropertiesResponse
		return ret
	}).(SharedPrivateLinkResourcePropertiesResponseOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupId
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) PrivateLink() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) *ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.PrivateLink
	}).(ResourceReferenceResponsePtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) PrivateLinkLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLinkLocation
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RequestMessage
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
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

type UrlSigningAction struct {
	Name       string                     `pulumi:"name"`
	Parameters UrlSigningActionParameters `pulumi:"parameters"`
}

type UrlSigningActionParameters struct {
	Algorithm             *string                     `pulumi:"algorithm"`
	OdataType             string                      `pulumi:"odataType"`
	ParameterNameOverride []UrlSigningParamIdentifier `pulumi:"parameterNameOverride"`
}

type UrlSigningActionParametersResponse struct {
	Algorithm             *string                             `pulumi:"algorithm"`
	OdataType             string                              `pulumi:"odataType"`
	ParameterNameOverride []UrlSigningParamIdentifierResponse `pulumi:"parameterNameOverride"`
}

type UrlSigningActionResponse struct {
	Name       string                             `pulumi:"name"`
	Parameters UrlSigningActionParametersResponse `pulumi:"parameters"`
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

type UrlSigningKeyParameters struct {
	KeyId         string            `pulumi:"keyId"`
	SecretSource  ResourceReference `pulumi:"secretSource"`
	SecretVersion *string           `pulumi:"secretVersion"`
	Type          string            `pulumi:"type"`
}

type UrlSigningKeyParametersResponse struct {
	KeyId         string                    `pulumi:"keyId"`
	SecretSource  ResourceReferenceResponse `pulumi:"secretSource"`
	SecretVersion *string                   `pulumi:"secretVersion"`
	Type          string                    `pulumi:"type"`
}

type UrlSigningKeyResponse struct {
	KeyId               string                               `pulumi:"keyId"`
	KeySourceParameters KeyVaultSigningKeyParametersResponse `pulumi:"keySourceParameters"`
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

type UrlSigningParamIdentifierResponse struct {
	ParamIndicator string `pulumi:"paramIndicator"`
	ParamName      string `pulumi:"paramName"`
}

type UserManagedHttpsParametersResponse struct {
	CertificateSource           string                                      `pulumi:"certificateSource"`
	CertificateSourceParameters KeyVaultCertificateSourceParametersResponse `pulumi:"certificateSourceParameters"`
	MinimumTlsVersion           *string                                     `pulumi:"minimumTlsVersion"`
	ProtocolType                string                                      `pulumi:"protocolType"`
}

func init() {
	pulumi.RegisterOutputType(AFDDomainHttpsParametersOutput{})
	pulumi.RegisterOutputType(AFDDomainHttpsParametersPtrOutput{})
	pulumi.RegisterOutputType(AFDDomainHttpsParametersResponseOutput{})
	pulumi.RegisterOutputType(AFDDomainHttpsParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(CdnEndpointResponseOutput{})
	pulumi.RegisterOutputType(CdnEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(CompressionSettingsOutput{})
	pulumi.RegisterOutputType(CompressionSettingsPtrOutput{})
	pulumi.RegisterOutputType(CompressionSettingsResponseOutput{})
	pulumi.RegisterOutputType(CompressionSettingsResponsePtrOutput{})
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
	pulumi.RegisterOutputType(DeliveryRuleResponseOutput{})
	pulumi.RegisterOutputType(DeliveryRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(DomainValidationPropertiesResponseOutput{})
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
	pulumi.RegisterOutputType(HealthProbeParametersOutput{})
	pulumi.RegisterOutputType(HealthProbeParametersPtrOutput{})
	pulumi.RegisterOutputType(HealthProbeParametersResponseOutput{})
	pulumi.RegisterOutputType(HealthProbeParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersArrayOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersResponseOutput{})
	pulumi.RegisterOutputType(HttpErrorRangeParametersResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultSigningKeyParametersOutput{})
	pulumi.RegisterOutputType(KeyVaultSigningKeyParametersResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsParametersOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsParametersPtrOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsParametersResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsParametersResponsePtrOutput{})
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
	pulumi.RegisterOutputType(SecurityPolicyWebApplicationFirewallAssociationOutput{})
	pulumi.RegisterOutputType(SecurityPolicyWebApplicationFirewallAssociationArrayOutput{})
	pulumi.RegisterOutputType(SecurityPolicyWebApplicationFirewallAssociationResponseOutput{})
	pulumi.RegisterOutputType(SecurityPolicyWebApplicationFirewallAssociationResponseArrayOutput{})
	pulumi.RegisterOutputType(SecurityPolicyWebApplicationFirewallParametersOutput{})
	pulumi.RegisterOutputType(SecurityPolicyWebApplicationFirewallParametersPtrOutput{})
	pulumi.RegisterOutputType(SecurityPolicyWebApplicationFirewallParametersResponseOutput{})
	pulumi.RegisterOutputType(SecurityPolicyWebApplicationFirewallParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourcePropertiesOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UrlSigningKeyOutput{})
	pulumi.RegisterOutputType(UrlSigningKeyArrayOutput{})
	pulumi.RegisterOutputType(UrlSigningKeyResponseOutput{})
	pulumi.RegisterOutputType(UrlSigningKeyResponseArrayOutput{})
}
