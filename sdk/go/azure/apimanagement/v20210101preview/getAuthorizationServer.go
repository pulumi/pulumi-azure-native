


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAuthorizationServer(ctx *pulumi.Context, args *LookupAuthorizationServerArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationServerResult, error) {
	var rv LookupAuthorizationServerResult
	err := ctx.Invoke("azure-native:apimanagement/v20210101preview:getAuthorizationServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationServerArgs struct {
	Authsid           string `pulumi:"authsid"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupAuthorizationServerResult struct {
	AuthorizationEndpoint      string                               `pulumi:"authorizationEndpoint"`
	AuthorizationMethods       []string                             `pulumi:"authorizationMethods"`
	BearerTokenSendingMethods  []string                             `pulumi:"bearerTokenSendingMethods"`
	ClientAuthenticationMethod []string                             `pulumi:"clientAuthenticationMethod"`
	ClientId                   string                               `pulumi:"clientId"`
	ClientRegistrationEndpoint string                               `pulumi:"clientRegistrationEndpoint"`
	ClientSecret               *string                              `pulumi:"clientSecret"`
	DefaultScope               *string                              `pulumi:"defaultScope"`
	Description                *string                              `pulumi:"description"`
	DisplayName                string                               `pulumi:"displayName"`
	GrantTypes                 []string                             `pulumi:"grantTypes"`
	Id                         string                               `pulumi:"id"`
	Name                       string                               `pulumi:"name"`
	ResourceOwnerPassword      *string                              `pulumi:"resourceOwnerPassword"`
	ResourceOwnerUsername      *string                              `pulumi:"resourceOwnerUsername"`
	SupportState               *bool                                `pulumi:"supportState"`
	TokenBodyParameters        []TokenBodyParameterContractResponse `pulumi:"tokenBodyParameters"`
	TokenEndpoint              *string                              `pulumi:"tokenEndpoint"`
	Type                       string                               `pulumi:"type"`
}

func LookupAuthorizationServerOutput(ctx *pulumi.Context, args LookupAuthorizationServerOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationServerResult, error) {
			args := v.(LookupAuthorizationServerArgs)
			r, err := LookupAuthorizationServer(ctx, &args, opts...)
			var s LookupAuthorizationServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizationServerResultOutput)
}

type LookupAuthorizationServerOutputArgs struct {
	Authsid           pulumi.StringInput `pulumi:"authsid"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupAuthorizationServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationServerArgs)(nil)).Elem()
}


type LookupAuthorizationServerResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationServerResult)(nil)).Elem()
}

func (o LookupAuthorizationServerResultOutput) ToLookupAuthorizationServerResultOutput() LookupAuthorizationServerResultOutput {
	return o
}

func (o LookupAuthorizationServerResultOutput) ToLookupAuthorizationServerResultOutputWithContext(ctx context.Context) LookupAuthorizationServerResultOutput {
	return o
}

func (o LookupAuthorizationServerResultOutput) AuthorizationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.AuthorizationEndpoint }).(pulumi.StringOutput)
}

func (o LookupAuthorizationServerResultOutput) AuthorizationMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) []string { return v.AuthorizationMethods }).(pulumi.StringArrayOutput)
}

func (o LookupAuthorizationServerResultOutput) BearerTokenSendingMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) []string { return v.BearerTokenSendingMethods }).(pulumi.StringArrayOutput)
}

func (o LookupAuthorizationServerResultOutput) ClientAuthenticationMethod() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) []string { return v.ClientAuthenticationMethod }).(pulumi.StringArrayOutput)
}

func (o LookupAuthorizationServerResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o LookupAuthorizationServerResultOutput) ClientRegistrationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.ClientRegistrationEndpoint }).(pulumi.StringOutput)
}

func (o LookupAuthorizationServerResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationServerResultOutput) DefaultScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.DefaultScope }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationServerResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationServerResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupAuthorizationServerResultOutput) GrantTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) []string { return v.GrantTypes }).(pulumi.StringArrayOutput)
}

func (o LookupAuthorizationServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAuthorizationServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAuthorizationServerResultOutput) ResourceOwnerPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.ResourceOwnerPassword }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationServerResultOutput) ResourceOwnerUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.ResourceOwnerUsername }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationServerResultOutput) SupportState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *bool { return v.SupportState }).(pulumi.BoolPtrOutput)
}

func (o LookupAuthorizationServerResultOutput) TokenBodyParameters() TokenBodyParameterContractResponseArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) []TokenBodyParameterContractResponse {
		return v.TokenBodyParameters
	}).(TokenBodyParameterContractResponseArrayOutput)
}

func (o LookupAuthorizationServerResultOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.TokenEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationServerResultOutput{})
}
