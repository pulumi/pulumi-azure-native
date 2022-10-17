


package v20170301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthorizationServer struct {
	pulumi.CustomResourceState

	AuthorizationEndpoint      pulumi.StringOutput                           `pulumi:"authorizationEndpoint"`
	AuthorizationMethods       pulumi.StringArrayOutput                      `pulumi:"authorizationMethods"`
	BearerTokenSendingMethods  pulumi.StringArrayOutput                      `pulumi:"bearerTokenSendingMethods"`
	ClientAuthenticationMethod pulumi.StringArrayOutput                      `pulumi:"clientAuthenticationMethod"`
	ClientId                   pulumi.StringOutput                           `pulumi:"clientId"`
	ClientRegistrationEndpoint pulumi.StringOutput                           `pulumi:"clientRegistrationEndpoint"`
	ClientSecret               pulumi.StringPtrOutput                        `pulumi:"clientSecret"`
	DefaultScope               pulumi.StringPtrOutput                        `pulumi:"defaultScope"`
	Description                pulumi.StringPtrOutput                        `pulumi:"description"`
	DisplayName                pulumi.StringOutput                           `pulumi:"displayName"`
	GrantTypes                 pulumi.StringArrayOutput                      `pulumi:"grantTypes"`
	Name                       pulumi.StringOutput                           `pulumi:"name"`
	ResourceOwnerPassword      pulumi.StringPtrOutput                        `pulumi:"resourceOwnerPassword"`
	ResourceOwnerUsername      pulumi.StringPtrOutput                        `pulumi:"resourceOwnerUsername"`
	SupportState               pulumi.BoolPtrOutput                          `pulumi:"supportState"`
	TokenBodyParameters        TokenBodyParameterContractResponseArrayOutput `pulumi:"tokenBodyParameters"`
	TokenEndpoint              pulumi.StringPtrOutput                        `pulumi:"tokenEndpoint"`
	Type                       pulumi.StringOutput                           `pulumi:"type"`
}


func NewAuthorizationServer(ctx *pulumi.Context,
	name string, args *AuthorizationServerArgs, opts ...pulumi.ResourceOption) (*AuthorizationServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizationEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationEndpoint'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientRegistrationEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ClientRegistrationEndpoint'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.GrantTypes == nil {
		return nil, errors.New("invalid value for required argument 'GrantTypes'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:AuthorizationServer"),
		},
	})
	opts = append(opts, aliases)
	var resource AuthorizationServer
	err := ctx.RegisterResource("azure-native:apimanagement/v20170301:AuthorizationServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAuthorizationServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationServerState, opts ...pulumi.ResourceOption) (*AuthorizationServer, error) {
	var resource AuthorizationServer
	err := ctx.ReadResource("azure-native:apimanagement/v20170301:AuthorizationServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type authorizationServerState struct {
}

type AuthorizationServerState struct {
}

func (AuthorizationServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationServerState)(nil)).Elem()
}

type authorizationServerArgs struct {
	AuthorizationEndpoint      string                       `pulumi:"authorizationEndpoint"`
	AuthorizationMethods       []AuthorizationMethod        `pulumi:"authorizationMethods"`
	Authsid                    *string                      `pulumi:"authsid"`
	BearerTokenSendingMethods  []string                     `pulumi:"bearerTokenSendingMethods"`
	ClientAuthenticationMethod []string                     `pulumi:"clientAuthenticationMethod"`
	ClientId                   string                       `pulumi:"clientId"`
	ClientRegistrationEndpoint string                       `pulumi:"clientRegistrationEndpoint"`
	ClientSecret               *string                      `pulumi:"clientSecret"`
	DefaultScope               *string                      `pulumi:"defaultScope"`
	Description                *string                      `pulumi:"description"`
	DisplayName                string                       `pulumi:"displayName"`
	GrantTypes                 []string                     `pulumi:"grantTypes"`
	ResourceGroupName          string                       `pulumi:"resourceGroupName"`
	ResourceOwnerPassword      *string                      `pulumi:"resourceOwnerPassword"`
	ResourceOwnerUsername      *string                      `pulumi:"resourceOwnerUsername"`
	ServiceName                string                       `pulumi:"serviceName"`
	SupportState               *bool                        `pulumi:"supportState"`
	TokenBodyParameters        []TokenBodyParameterContract `pulumi:"tokenBodyParameters"`
	TokenEndpoint              *string                      `pulumi:"tokenEndpoint"`
}


type AuthorizationServerArgs struct {
	AuthorizationEndpoint      pulumi.StringInput
	AuthorizationMethods       AuthorizationMethodArrayInput
	Authsid                    pulumi.StringPtrInput
	BearerTokenSendingMethods  pulumi.StringArrayInput
	ClientAuthenticationMethod pulumi.StringArrayInput
	ClientId                   pulumi.StringInput
	ClientRegistrationEndpoint pulumi.StringInput
	ClientSecret               pulumi.StringPtrInput
	DefaultScope               pulumi.StringPtrInput
	Description                pulumi.StringPtrInput
	DisplayName                pulumi.StringInput
	GrantTypes                 pulumi.StringArrayInput
	ResourceGroupName          pulumi.StringInput
	ResourceOwnerPassword      pulumi.StringPtrInput
	ResourceOwnerUsername      pulumi.StringPtrInput
	ServiceName                pulumi.StringInput
	SupportState               pulumi.BoolPtrInput
	TokenBodyParameters        TokenBodyParameterContractArrayInput
	TokenEndpoint              pulumi.StringPtrInput
}

func (AuthorizationServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationServerArgs)(nil)).Elem()
}

type AuthorizationServerInput interface {
	pulumi.Input

	ToAuthorizationServerOutput() AuthorizationServerOutput
	ToAuthorizationServerOutputWithContext(ctx context.Context) AuthorizationServerOutput
}

func (*AuthorizationServer) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationServer)(nil)).Elem()
}

func (i *AuthorizationServer) ToAuthorizationServerOutput() AuthorizationServerOutput {
	return i.ToAuthorizationServerOutputWithContext(context.Background())
}

func (i *AuthorizationServer) ToAuthorizationServerOutputWithContext(ctx context.Context) AuthorizationServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationServerOutput)
}

type AuthorizationServerOutput struct{ *pulumi.OutputState }

func (AuthorizationServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationServer)(nil)).Elem()
}

func (o AuthorizationServerOutput) ToAuthorizationServerOutput() AuthorizationServerOutput {
	return o
}

func (o AuthorizationServerOutput) ToAuthorizationServerOutputWithContext(ctx context.Context) AuthorizationServerOutput {
	return o
}

func (o AuthorizationServerOutput) AuthorizationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.AuthorizationEndpoint }).(pulumi.StringOutput)
}

func (o AuthorizationServerOutput) AuthorizationMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringArrayOutput { return v.AuthorizationMethods }).(pulumi.StringArrayOutput)
}

func (o AuthorizationServerOutput) BearerTokenSendingMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringArrayOutput { return v.BearerTokenSendingMethods }).(pulumi.StringArrayOutput)
}

func (o AuthorizationServerOutput) ClientAuthenticationMethod() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringArrayOutput { return v.ClientAuthenticationMethod }).(pulumi.StringArrayOutput)
}

func (o AuthorizationServerOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

func (o AuthorizationServerOutput) ClientRegistrationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.ClientRegistrationEndpoint }).(pulumi.StringOutput)
}

func (o AuthorizationServerOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o AuthorizationServerOutput) DefaultScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.DefaultScope }).(pulumi.StringPtrOutput)
}

func (o AuthorizationServerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AuthorizationServerOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o AuthorizationServerOutput) GrantTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringArrayOutput { return v.GrantTypes }).(pulumi.StringArrayOutput)
}

func (o AuthorizationServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AuthorizationServerOutput) ResourceOwnerPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.ResourceOwnerPassword }).(pulumi.StringPtrOutput)
}

func (o AuthorizationServerOutput) ResourceOwnerUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.ResourceOwnerUsername }).(pulumi.StringPtrOutput)
}

func (o AuthorizationServerOutput) SupportState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.BoolPtrOutput { return v.SupportState }).(pulumi.BoolPtrOutput)
}

func (o AuthorizationServerOutput) TokenBodyParameters() TokenBodyParameterContractResponseArrayOutput {
	return o.ApplyT(func(v *AuthorizationServer) TokenBodyParameterContractResponseArrayOutput {
		return v.TokenBodyParameters
	}).(TokenBodyParameterContractResponseArrayOutput)
}

func (o AuthorizationServerOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.TokenEndpoint }).(pulumi.StringPtrOutput)
}

func (o AuthorizationServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationServerOutput{})
}
