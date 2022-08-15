


package app

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContainerAppsAuthConfig(ctx *pulumi.Context, args *LookupContainerAppsAuthConfigArgs, opts ...pulumi.InvokeOption) (*LookupContainerAppsAuthConfigResult, error) {
	var rv LookupContainerAppsAuthConfigResult
	err := ctx.Invoke("azure-native:app:getContainerAppsAuthConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerAppsAuthConfigArgs struct {
	AuthConfigName    string `pulumi:"authConfigName"`
	ContainerAppName  string `pulumi:"containerAppName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupContainerAppsAuthConfigResult struct {
	GlobalValidation  *GlobalValidationResponse  `pulumi:"globalValidation"`
	HttpSettings      *HttpSettingsResponse      `pulumi:"httpSettings"`
	Id                string                     `pulumi:"id"`
	IdentityProviders *IdentityProvidersResponse `pulumi:"identityProviders"`
	Login             *LoginResponse             `pulumi:"login"`
	Name              string                     `pulumi:"name"`
	Platform          *AuthPlatformResponse      `pulumi:"platform"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	Type              string                     `pulumi:"type"`
}

func LookupContainerAppsAuthConfigOutput(ctx *pulumi.Context, args LookupContainerAppsAuthConfigOutputArgs, opts ...pulumi.InvokeOption) LookupContainerAppsAuthConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerAppsAuthConfigResult, error) {
			args := v.(LookupContainerAppsAuthConfigArgs)
			r, err := LookupContainerAppsAuthConfig(ctx, &args, opts...)
			var s LookupContainerAppsAuthConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerAppsAuthConfigResultOutput)
}

type LookupContainerAppsAuthConfigOutputArgs struct {
	AuthConfigName    pulumi.StringInput `pulumi:"authConfigName"`
	ContainerAppName  pulumi.StringInput `pulumi:"containerAppName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContainerAppsAuthConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppsAuthConfigArgs)(nil)).Elem()
}


type LookupContainerAppsAuthConfigResultOutput struct{ *pulumi.OutputState }

func (LookupContainerAppsAuthConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppsAuthConfigResult)(nil)).Elem()
}

func (o LookupContainerAppsAuthConfigResultOutput) ToLookupContainerAppsAuthConfigResultOutput() LookupContainerAppsAuthConfigResultOutput {
	return o
}

func (o LookupContainerAppsAuthConfigResultOutput) ToLookupContainerAppsAuthConfigResultOutputWithContext(ctx context.Context) LookupContainerAppsAuthConfigResultOutput {
	return o
}

func (o LookupContainerAppsAuthConfigResultOutput) GlobalValidation() GlobalValidationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) *GlobalValidationResponse { return v.GlobalValidation }).(GlobalValidationResponsePtrOutput)
}

func (o LookupContainerAppsAuthConfigResultOutput) HttpSettings() HttpSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) *HttpSettingsResponse { return v.HttpSettings }).(HttpSettingsResponsePtrOutput)
}

func (o LookupContainerAppsAuthConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContainerAppsAuthConfigResultOutput) IdentityProviders() IdentityProvidersResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) *IdentityProvidersResponse { return v.IdentityProviders }).(IdentityProvidersResponsePtrOutput)
}

func (o LookupContainerAppsAuthConfigResultOutput) Login() LoginResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) *LoginResponse { return v.Login }).(LoginResponsePtrOutput)
}

func (o LookupContainerAppsAuthConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContainerAppsAuthConfigResultOutput) Platform() AuthPlatformResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) *AuthPlatformResponse { return v.Platform }).(AuthPlatformResponsePtrOutput)
}

func (o LookupContainerAppsAuthConfigResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupContainerAppsAuthConfigResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerAppsAuthConfigResultOutput{})
}
