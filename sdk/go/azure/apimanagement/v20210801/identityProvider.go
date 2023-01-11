


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IdentityProvider struct {
	pulumi.CustomResourceState

	AllowedTenants           pulumi.StringArrayOutput `pulumi:"allowedTenants"`
	Authority                pulumi.StringPtrOutput   `pulumi:"authority"`
	ClientId                 pulumi.StringOutput      `pulumi:"clientId"`
	ClientSecret             pulumi.StringPtrOutput   `pulumi:"clientSecret"`
	Name                     pulumi.StringOutput      `pulumi:"name"`
	PasswordResetPolicyName  pulumi.StringPtrOutput   `pulumi:"passwordResetPolicyName"`
	ProfileEditingPolicyName pulumi.StringPtrOutput   `pulumi:"profileEditingPolicyName"`
	SigninPolicyName         pulumi.StringPtrOutput   `pulumi:"signinPolicyName"`
	SigninTenant             pulumi.StringPtrOutput   `pulumi:"signinTenant"`
	SignupPolicyName         pulumi.StringPtrOutput   `pulumi:"signupPolicyName"`
	Type                     pulumi.StringOutput      `pulumi:"type"`
}


func NewIdentityProvider(ctx *pulumi.Context,
	name string, args *IdentityProviderArgs, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:IdentityProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:IdentityProvider"),
		},
	})
	opts = append(opts, aliases)
	var resource IdentityProvider
	err := ctx.RegisterResource("azure-native:apimanagement/v20210801:IdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderState, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	var resource IdentityProvider
	err := ctx.ReadResource("azure-native:apimanagement/v20210801:IdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type identityProviderState struct {
}

type IdentityProviderState struct {
}

func (IdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderState)(nil)).Elem()
}

type identityProviderArgs struct {
	AllowedTenants           []string `pulumi:"allowedTenants"`
	Authority                *string  `pulumi:"authority"`
	ClientId                 string   `pulumi:"clientId"`
	ClientSecret             string   `pulumi:"clientSecret"`
	IdentityProviderName     *string  `pulumi:"identityProviderName"`
	PasswordResetPolicyName  *string  `pulumi:"passwordResetPolicyName"`
	ProfileEditingPolicyName *string  `pulumi:"profileEditingPolicyName"`
	ResourceGroupName        string   `pulumi:"resourceGroupName"`
	ServiceName              string   `pulumi:"serviceName"`
	SigninPolicyName         *string  `pulumi:"signinPolicyName"`
	SigninTenant             *string  `pulumi:"signinTenant"`
	SignupPolicyName         *string  `pulumi:"signupPolicyName"`
	Type                     *string  `pulumi:"type"`
}


type IdentityProviderArgs struct {
	AllowedTenants           pulumi.StringArrayInput
	Authority                pulumi.StringPtrInput
	ClientId                 pulumi.StringInput
	ClientSecret             pulumi.StringInput
	IdentityProviderName     pulumi.StringPtrInput
	PasswordResetPolicyName  pulumi.StringPtrInput
	ProfileEditingPolicyName pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	ServiceName              pulumi.StringInput
	SigninPolicyName         pulumi.StringPtrInput
	SigninTenant             pulumi.StringPtrInput
	SignupPolicyName         pulumi.StringPtrInput
	Type                     pulumi.StringPtrInput
}

func (IdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderArgs)(nil)).Elem()
}

type IdentityProviderInput interface {
	pulumi.Input

	ToIdentityProviderOutput() IdentityProviderOutput
	ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput
}

func (*IdentityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil)).Elem()
}

func (i *IdentityProvider) ToIdentityProviderOutput() IdentityProviderOutput {
	return i.ToIdentityProviderOutputWithContext(context.Background())
}

func (i *IdentityProvider) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderOutput)
}

type IdentityProviderOutput struct{ *pulumi.OutputState }

func (IdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil)).Elem()
}

func (o IdentityProviderOutput) ToIdentityProviderOutput() IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) AllowedTenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringArrayOutput { return v.AllowedTenants }).(pulumi.StringArrayOutput)
}

func (o IdentityProviderOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

func (o IdentityProviderOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IdentityProviderOutput) PasswordResetPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.PasswordResetPolicyName }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderOutput) ProfileEditingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.ProfileEditingPolicyName }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderOutput) SigninPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.SigninPolicyName }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderOutput) SigninTenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.SigninTenant }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderOutput) SignupPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.SignupPolicyName }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityProviderOutput{})
}
