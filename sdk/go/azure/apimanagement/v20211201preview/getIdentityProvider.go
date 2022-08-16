


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIdentityProvider(ctx *pulumi.Context, args *LookupIdentityProviderArgs, opts ...pulumi.InvokeOption) (*LookupIdentityProviderResult, error) {
	var rv LookupIdentityProviderResult
	err := ctx.Invoke("azure-native:apimanagement/v20211201preview:getIdentityProvider", args, &rv, opts...)
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
	ClientLibrary            *string  `pulumi:"clientLibrary"`
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

func LookupIdentityProviderOutput(ctx *pulumi.Context, args LookupIdentityProviderOutputArgs, opts ...pulumi.InvokeOption) LookupIdentityProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIdentityProviderResult, error) {
			args := v.(LookupIdentityProviderArgs)
			r, err := LookupIdentityProvider(ctx, &args, opts...)
			var s LookupIdentityProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIdentityProviderResultOutput)
}

type LookupIdentityProviderOutputArgs struct {
	IdentityProviderName pulumi.StringInput `pulumi:"identityProviderName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName          pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupIdentityProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityProviderArgs)(nil)).Elem()
}


type LookupIdentityProviderResultOutput struct{ *pulumi.OutputState }

func (LookupIdentityProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityProviderResult)(nil)).Elem()
}

func (o LookupIdentityProviderResultOutput) ToLookupIdentityProviderResultOutput() LookupIdentityProviderResultOutput {
	return o
}

func (o LookupIdentityProviderResultOutput) ToLookupIdentityProviderResultOutputWithContext(ctx context.Context) LookupIdentityProviderResultOutput {
	return o
}

func (o LookupIdentityProviderResultOutput) AllowedTenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) []string { return v.AllowedTenants }).(pulumi.StringArrayOutput)
}

func (o LookupIdentityProviderResultOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityProviderResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o LookupIdentityProviderResultOutput) ClientLibrary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) *string { return v.ClientLibrary }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityProviderResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIdentityProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIdentityProviderResultOutput) PasswordResetPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) *string { return v.PasswordResetPolicyName }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityProviderResultOutput) ProfileEditingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) *string { return v.ProfileEditingPolicyName }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityProviderResultOutput) SigninPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) *string { return v.SigninPolicyName }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityProviderResultOutput) SigninTenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) *string { return v.SigninTenant }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityProviderResultOutput) SignupPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) *string { return v.SignupPolicyName }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdentityProviderResultOutput{})
}
