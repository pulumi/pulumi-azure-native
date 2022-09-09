


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApi(ctx *pulumi.Context, args *LookupApiArgs, opts ...pulumi.InvokeOption) (*LookupApiResult, error) {
	var rv LookupApiResult
	err := ctx.Invoke("azure-native:apimanagement/v20210401preview:getApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiArgs struct {
	ApiId             string `pulumi:"apiId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiResult struct {
	ApiRevision                   *string                                        `pulumi:"apiRevision"`
	ApiRevisionDescription        *string                                        `pulumi:"apiRevisionDescription"`
	ApiType                       *string                                        `pulumi:"apiType"`
	ApiVersion                    *string                                        `pulumi:"apiVersion"`
	ApiVersionDescription         *string                                        `pulumi:"apiVersionDescription"`
	ApiVersionSet                 *ApiVersionSetContractDetailsResponse          `pulumi:"apiVersionSet"`
	ApiVersionSetId               *string                                        `pulumi:"apiVersionSetId"`
	AuthenticationSettings        *AuthenticationSettingsContractResponse        `pulumi:"authenticationSettings"`
	Contact                       *ApiContactInformationResponse                 `pulumi:"contact"`
	Description                   *string                                        `pulumi:"description"`
	DisplayName                   *string                                        `pulumi:"displayName"`
	Id                            string                                         `pulumi:"id"`
	IsCurrent                     *bool                                          `pulumi:"isCurrent"`
	IsOnline                      bool                                           `pulumi:"isOnline"`
	License                       *ApiLicenseInformationResponse                 `pulumi:"license"`
	Name                          string                                         `pulumi:"name"`
	Path                          string                                         `pulumi:"path"`
	Protocols                     []string                                       `pulumi:"protocols"`
	ServiceUrl                    *string                                        `pulumi:"serviceUrl"`
	SourceApiId                   *string                                        `pulumi:"sourceApiId"`
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContractResponse `pulumi:"subscriptionKeyParameterNames"`
	SubscriptionRequired          *bool                                          `pulumi:"subscriptionRequired"`
	TermsOfServiceUrl             *string                                        `pulumi:"termsOfServiceUrl"`
	Type                          string                                         `pulumi:"type"`
}

func LookupApiOutput(ctx *pulumi.Context, args LookupApiOutputArgs, opts ...pulumi.InvokeOption) LookupApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiResult, error) {
			args := v.(LookupApiArgs)
			r, err := LookupApi(ctx, &args, opts...)
			var s LookupApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiResultOutput)
}

type LookupApiOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiArgs)(nil)).Elem()
}


type LookupApiResultOutput struct{ *pulumi.OutputState }

func (LookupApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiResult)(nil)).Elem()
}

func (o LookupApiResultOutput) ToLookupApiResultOutput() LookupApiResultOutput {
	return o
}

func (o LookupApiResultOutput) ToLookupApiResultOutputWithContext(ctx context.Context) LookupApiResultOutput {
	return o
}

func (o LookupApiResultOutput) ApiRevision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiRevision }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) ApiRevisionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiRevisionDescription }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiType }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) ApiVersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiVersionDescription }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) ApiVersionSet() ApiVersionSetContractDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *ApiVersionSetContractDetailsResponse { return v.ApiVersionSet }).(ApiVersionSetContractDetailsResponsePtrOutput)
}

func (o LookupApiResultOutput) ApiVersionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ApiVersionSetId }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) AuthenticationSettings() AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *AuthenticationSettingsContractResponse { return v.AuthenticationSettings }).(AuthenticationSettingsContractResponsePtrOutput)
}

func (o LookupApiResultOutput) Contact() ApiContactInformationResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *ApiContactInformationResponse { return v.Contact }).(ApiContactInformationResponsePtrOutput)
}

func (o LookupApiResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiResultOutput) IsCurrent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *bool { return v.IsCurrent }).(pulumi.BoolPtrOutput)
}

func (o LookupApiResultOutput) IsOnline() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApiResult) bool { return v.IsOnline }).(pulumi.BoolOutput)
}

func (o LookupApiResultOutput) License() ApiLicenseInformationResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *ApiLicenseInformationResponse { return v.License }).(ApiLicenseInformationResponsePtrOutput)
}

func (o LookupApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o LookupApiResultOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApiResult) []string { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o LookupApiResultOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) SourceApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.SourceApiId }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) SubscriptionKeyParameterNames() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o.ApplyT(func(v LookupApiResult) *SubscriptionKeyParameterNamesContractResponse {
		return v.SubscriptionKeyParameterNames
	}).(SubscriptionKeyParameterNamesContractResponsePtrOutput)
}

func (o LookupApiResultOutput) SubscriptionRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *bool { return v.SubscriptionRequired }).(pulumi.BoolPtrOutput)
}

func (o LookupApiResultOutput) TermsOfServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiResult) *string { return v.TermsOfServiceUrl }).(pulumi.StringPtrOutput)
}

func (o LookupApiResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiResultOutput{})
}
