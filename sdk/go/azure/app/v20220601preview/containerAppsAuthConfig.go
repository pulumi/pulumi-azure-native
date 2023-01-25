


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerAppsAuthConfig struct {
	pulumi.CustomResourceState

	GlobalValidation  GlobalValidationResponsePtrOutput  `pulumi:"globalValidation"`
	HttpSettings      HttpSettingsResponsePtrOutput      `pulumi:"httpSettings"`
	IdentityProviders IdentityProvidersResponsePtrOutput `pulumi:"identityProviders"`
	Login             LoginResponsePtrOutput             `pulumi:"login"`
	Name              pulumi.StringOutput                `pulumi:"name"`
	Platform          AuthPlatformResponsePtrOutput      `pulumi:"platform"`
	SystemData        SystemDataResponseOutput           `pulumi:"systemData"`
	Type              pulumi.StringOutput                `pulumi:"type"`
}


func NewContainerAppsAuthConfig(ctx *pulumi.Context,
	name string, args *ContainerAppsAuthConfigArgs, opts ...pulumi.ResourceOption) (*ContainerAppsAuthConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerAppName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerAppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:ContainerAppsAuthConfig"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220101preview:ContainerAppsAuthConfig"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220301:ContainerAppsAuthConfig"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221001:ContainerAppsAuthConfig"),
		},
	})
	opts = append(opts, aliases)
	var resource ContainerAppsAuthConfig
	err := ctx.RegisterResource("azure-native:app/v20220601preview:ContainerAppsAuthConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContainerAppsAuthConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerAppsAuthConfigState, opts ...pulumi.ResourceOption) (*ContainerAppsAuthConfig, error) {
	var resource ContainerAppsAuthConfig
	err := ctx.ReadResource("azure-native:app/v20220601preview:ContainerAppsAuthConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type containerAppsAuthConfigState struct {
}

type ContainerAppsAuthConfigState struct {
}

func (ContainerAppsAuthConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppsAuthConfigState)(nil)).Elem()
}

type containerAppsAuthConfigArgs struct {
	AuthConfigName    *string            `pulumi:"authConfigName"`
	ContainerAppName  string             `pulumi:"containerAppName"`
	GlobalValidation  *GlobalValidation  `pulumi:"globalValidation"`
	HttpSettings      *HttpSettings      `pulumi:"httpSettings"`
	IdentityProviders *IdentityProviders `pulumi:"identityProviders"`
	Login             *Login             `pulumi:"login"`
	Platform          *AuthPlatform      `pulumi:"platform"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
}


type ContainerAppsAuthConfigArgs struct {
	AuthConfigName    pulumi.StringPtrInput
	ContainerAppName  pulumi.StringInput
	GlobalValidation  GlobalValidationPtrInput
	HttpSettings      HttpSettingsPtrInput
	IdentityProviders IdentityProvidersPtrInput
	Login             LoginPtrInput
	Platform          AuthPlatformPtrInput
	ResourceGroupName pulumi.StringInput
}

func (ContainerAppsAuthConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppsAuthConfigArgs)(nil)).Elem()
}

type ContainerAppsAuthConfigInput interface {
	pulumi.Input

	ToContainerAppsAuthConfigOutput() ContainerAppsAuthConfigOutput
	ToContainerAppsAuthConfigOutputWithContext(ctx context.Context) ContainerAppsAuthConfigOutput
}

func (*ContainerAppsAuthConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppsAuthConfig)(nil)).Elem()
}

func (i *ContainerAppsAuthConfig) ToContainerAppsAuthConfigOutput() ContainerAppsAuthConfigOutput {
	return i.ToContainerAppsAuthConfigOutputWithContext(context.Background())
}

func (i *ContainerAppsAuthConfig) ToContainerAppsAuthConfigOutputWithContext(ctx context.Context) ContainerAppsAuthConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppsAuthConfigOutput)
}

type ContainerAppsAuthConfigOutput struct{ *pulumi.OutputState }

func (ContainerAppsAuthConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppsAuthConfig)(nil)).Elem()
}

func (o ContainerAppsAuthConfigOutput) ToContainerAppsAuthConfigOutput() ContainerAppsAuthConfigOutput {
	return o
}

func (o ContainerAppsAuthConfigOutput) ToContainerAppsAuthConfigOutputWithContext(ctx context.Context) ContainerAppsAuthConfigOutput {
	return o
}

func (o ContainerAppsAuthConfigOutput) GlobalValidation() GlobalValidationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) GlobalValidationResponsePtrOutput { return v.GlobalValidation }).(GlobalValidationResponsePtrOutput)
}

func (o ContainerAppsAuthConfigOutput) HttpSettings() HttpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) HttpSettingsResponsePtrOutput { return v.HttpSettings }).(HttpSettingsResponsePtrOutput)
}

func (o ContainerAppsAuthConfigOutput) IdentityProviders() IdentityProvidersResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) IdentityProvidersResponsePtrOutput { return v.IdentityProviders }).(IdentityProvidersResponsePtrOutput)
}

func (o ContainerAppsAuthConfigOutput) Login() LoginResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) LoginResponsePtrOutput { return v.Login }).(LoginResponsePtrOutput)
}

func (o ContainerAppsAuthConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerAppsAuthConfigOutput) Platform() AuthPlatformResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) AuthPlatformResponsePtrOutput { return v.Platform }).(AuthPlatformResponsePtrOutput)
}

func (o ContainerAppsAuthConfigOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ContainerAppsAuthConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAppsAuthConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerAppsAuthConfigOutput{})
}
