


package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAuthorizationServer(ctx *pulumi.Context, args *LookupAuthorizationServerArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationServerResult, error) {
	var rv LookupAuthorizationServerResult
	err := ctx.Invoke("azure-native:apimanagement/v20170301:getAuthorizationServer", args, &rv, opts...)
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
