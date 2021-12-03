


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TenantConfiguration struct {
	pulumi.CustomResourceState

	EnforcePrivateMarkdownStorage pulumi.BoolPtrOutput `pulumi:"enforcePrivateMarkdownStorage"`
	Name                          pulumi.StringOutput  `pulumi:"name"`
	Type                          pulumi.StringOutput  `pulumi:"type"`
}


func NewTenantConfiguration(ctx *pulumi.Context,
	name string, args *TenantConfigurationArgs, opts ...pulumi.ResourceOption) (*TenantConfiguration, error) {
	if args == nil {
		args = &TenantConfigurationArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:portal:TenantConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:portal/v20190101preview:TenantConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource TenantConfiguration
	err := ctx.RegisterResource("azure-native:portal/v20200901preview:TenantConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTenantConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantConfigurationState, opts ...pulumi.ResourceOption) (*TenantConfiguration, error) {
	var resource TenantConfiguration
	err := ctx.ReadResource("azure-native:portal/v20200901preview:TenantConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tenantConfigurationState struct {
}

type TenantConfigurationState struct {
}

func (TenantConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantConfigurationState)(nil)).Elem()
}

type tenantConfigurationArgs struct {
	ConfigurationName             *string `pulumi:"configurationName"`
	EnforcePrivateMarkdownStorage *bool   `pulumi:"enforcePrivateMarkdownStorage"`
}


type TenantConfigurationArgs struct {
	ConfigurationName             pulumi.StringPtrInput
	EnforcePrivateMarkdownStorage pulumi.BoolPtrInput
}

func (TenantConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantConfigurationArgs)(nil)).Elem()
}

type TenantConfigurationInput interface {
	pulumi.Input

	ToTenantConfigurationOutput() TenantConfigurationOutput
	ToTenantConfigurationOutputWithContext(ctx context.Context) TenantConfigurationOutput
}

func (*TenantConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*TenantConfiguration)(nil))
}

func (i *TenantConfiguration) ToTenantConfigurationOutput() TenantConfigurationOutput {
	return i.ToTenantConfigurationOutputWithContext(context.Background())
}

func (i *TenantConfiguration) ToTenantConfigurationOutputWithContext(ctx context.Context) TenantConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantConfigurationOutput)
}

type TenantConfigurationOutput struct{ *pulumi.OutputState }

func (TenantConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TenantConfiguration)(nil))
}

func (o TenantConfigurationOutput) ToTenantConfigurationOutput() TenantConfigurationOutput {
	return o
}

func (o TenantConfigurationOutput) ToTenantConfigurationOutputWithContext(ctx context.Context) TenantConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TenantConfigurationOutput{})
}
