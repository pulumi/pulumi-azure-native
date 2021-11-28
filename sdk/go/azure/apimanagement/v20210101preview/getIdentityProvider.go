


package v20210101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIdentityProvider(ctx *pulumi.Context, args *LookupIdentityProviderArgs, opts ...pulumi.InvokeOption) (*LookupIdentityProviderResult, error) {
	var rv LookupIdentityProviderResult
	err := ctx.Invoke("azure-native:apimanagement/v20210101preview:getIdentityProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIdentityProviderArgs struct {
	IdentityProviderName string `pulumi:"identityProviderName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	ServiceName          string `pulumi:"serviceName"`
}


type LookupIdentityProviderResult struct {
	AllowedTenants           []string `pulumi:"allowedTenants"`
	Authority                *string  `pulumi:"authority"`
	ClientId                 string   `pulumi:"clientId"`
	ClientSecret             *string  `pulumi:"clientSecret"`
	Id                       string   `pulumi:"id"`
	Name                     string   `pulumi:"name"`
	PasswordResetPolicyName  *string  `pulumi:"passwordResetPolicyName"`
	ProfileEditingPolicyName *string  `pulumi:"profileEditingPolicyName"`
	SigninPolicyName         *string  `pulumi:"signinPolicyName"`
	SigninTenant             *string  `pulumi:"signinTenant"`
	SignupPolicyName         *string  `pulumi:"signupPolicyName"`
	Type                     string   `pulumi:"type"`
}
