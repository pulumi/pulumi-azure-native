


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAuthorizationProvider(ctx *pulumi.Context, args *LookupAuthorizationProviderArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationProviderResult, error) {
	var rv LookupAuthorizationProviderResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:getAuthorizationProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationProviderArgs struct {
	AuthorizationProviderId string `pulumi:"authorizationProviderId"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	ServiceName             string `pulumi:"serviceName"`
}


type LookupAuthorizationProviderResult struct {
	DisplayName      *string                                      `pulumi:"displayName"`
	Id               string                                       `pulumi:"id"`
	IdentityProvider *string                                      `pulumi:"identityProvider"`
	Name             string                                       `pulumi:"name"`
	Oauth2           *AuthorizationProviderOAuth2SettingsResponse `pulumi:"oauth2"`
	Type             string                                       `pulumi:"type"`
}

func LookupAuthorizationProviderOutput(ctx *pulumi.Context, args LookupAuthorizationProviderOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationProviderResult, error) {
			args := v.(LookupAuthorizationProviderArgs)
			r, err := LookupAuthorizationProvider(ctx, &args, opts...)
			var s LookupAuthorizationProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizationProviderResultOutput)
}

type LookupAuthorizationProviderOutputArgs struct {
	AuthorizationProviderId pulumi.StringInput `pulumi:"authorizationProviderId"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName             pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupAuthorizationProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationProviderArgs)(nil)).Elem()
}


type LookupAuthorizationProviderResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationProviderResult)(nil)).Elem()
}

func (o LookupAuthorizationProviderResultOutput) ToLookupAuthorizationProviderResultOutput() LookupAuthorizationProviderResultOutput {
	return o
}

func (o LookupAuthorizationProviderResultOutput) ToLookupAuthorizationProviderResultOutputWithContext(ctx context.Context) LookupAuthorizationProviderResultOutput {
	return o
}

func (o LookupAuthorizationProviderResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAuthorizationProviderResultOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) *string { return v.IdentityProvider }).(pulumi.StringPtrOutput)
}

func (o LookupAuthorizationProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAuthorizationProviderResultOutput) Oauth2() AuthorizationProviderOAuth2SettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) *AuthorizationProviderOAuth2SettingsResponse {
		return v.Oauth2
	}).(AuthorizationProviderOAuth2SettingsResponsePtrOutput)
}

func (o LookupAuthorizationProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationProviderResultOutput{})
}
