


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthorizationProvider struct {
	pulumi.CustomResourceState

	DisplayName      pulumi.StringPtrOutput                               `pulumi:"displayName"`
	IdentityProvider pulumi.StringPtrOutput                               `pulumi:"identityProvider"`
	Name             pulumi.StringOutput                                  `pulumi:"name"`
	Oauth2           AuthorizationProviderOAuth2SettingsResponsePtrOutput `pulumi:"oauth2"`
	Type             pulumi.StringOutput                                  `pulumi:"type"`
}


func NewAuthorizationProvider(ctx *pulumi.Context,
	name string, args *AuthorizationProviderArgs, opts ...pulumi.ResourceOption) (*AuthorizationProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource AuthorizationProvider
	err := ctx.RegisterResource("azure-native:apimanagement/v20220401preview:AuthorizationProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAuthorizationProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationProviderState, opts ...pulumi.ResourceOption) (*AuthorizationProvider, error) {
	var resource AuthorizationProvider
	err := ctx.ReadResource("azure-native:apimanagement/v20220401preview:AuthorizationProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type authorizationProviderState struct {
}

type AuthorizationProviderState struct {
}

func (AuthorizationProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationProviderState)(nil)).Elem()
}

type authorizationProviderArgs struct {
	AuthorizationProviderId *string                              `pulumi:"authorizationProviderId"`
	DisplayName             *string                              `pulumi:"displayName"`
	IdentityProvider        *string                              `pulumi:"identityProvider"`
	Oauth2                  *AuthorizationProviderOAuth2Settings `pulumi:"oauth2"`
	ResourceGroupName       string                               `pulumi:"resourceGroupName"`
	ServiceName             string                               `pulumi:"serviceName"`
}


type AuthorizationProviderArgs struct {
	AuthorizationProviderId pulumi.StringPtrInput
	DisplayName             pulumi.StringPtrInput
	IdentityProvider        pulumi.StringPtrInput
	Oauth2                  AuthorizationProviderOAuth2SettingsPtrInput
	ResourceGroupName       pulumi.StringInput
	ServiceName             pulumi.StringInput
}

func (AuthorizationProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationProviderArgs)(nil)).Elem()
}

type AuthorizationProviderInput interface {
	pulumi.Input

	ToAuthorizationProviderOutput() AuthorizationProviderOutput
	ToAuthorizationProviderOutputWithContext(ctx context.Context) AuthorizationProviderOutput
}

func (*AuthorizationProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationProvider)(nil)).Elem()
}

func (i *AuthorizationProvider) ToAuthorizationProviderOutput() AuthorizationProviderOutput {
	return i.ToAuthorizationProviderOutputWithContext(context.Background())
}

func (i *AuthorizationProvider) ToAuthorizationProviderOutputWithContext(ctx context.Context) AuthorizationProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationProviderOutput)
}

type AuthorizationProviderOutput struct{ *pulumi.OutputState }

func (AuthorizationProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationProvider)(nil)).Elem()
}

func (o AuthorizationProviderOutput) ToAuthorizationProviderOutput() AuthorizationProviderOutput {
	return o
}

func (o AuthorizationProviderOutput) ToAuthorizationProviderOutputWithContext(ctx context.Context) AuthorizationProviderOutput {
	return o
}

func (o AuthorizationProviderOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationProvider) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o AuthorizationProviderOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationProvider) pulumi.StringPtrOutput { return v.IdentityProvider }).(pulumi.StringPtrOutput)
}

func (o AuthorizationProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AuthorizationProviderOutput) Oauth2() AuthorizationProviderOAuth2SettingsResponsePtrOutput {
	return o.ApplyT(func(v *AuthorizationProvider) AuthorizationProviderOAuth2SettingsResponsePtrOutput { return v.Oauth2 }).(AuthorizationProviderOAuth2SettingsResponsePtrOutput)
}

func (o AuthorizationProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationProviderOutput{})
}
