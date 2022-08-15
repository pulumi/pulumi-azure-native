


package web

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppServiceEnvironmentAseCustomDnsSuffixConfiguration struct {
	pulumi.CustomResourceState

	CertificateUrl            pulumi.StringPtrOutput `pulumi:"certificateUrl"`
	DnsSuffix                 pulumi.StringPtrOutput `pulumi:"dnsSuffix"`
	KeyVaultReferenceIdentity pulumi.StringPtrOutput `pulumi:"keyVaultReferenceIdentity"`
	Kind                      pulumi.StringPtrOutput `pulumi:"kind"`
	Name                      pulumi.StringOutput    `pulumi:"name"`
	ProvisioningDetails       pulumi.StringOutput    `pulumi:"provisioningDetails"`
	ProvisioningState         pulumi.StringOutput    `pulumi:"provisioningState"`
	Type                      pulumi.StringOutput    `pulumi:"type"`
}


func NewAppServiceEnvironmentAseCustomDnsSuffixConfiguration(ctx *pulumi.Context,
	name string, args *AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs, opts ...pulumi.ResourceOption) (*AppServiceEnvironmentAseCustomDnsSuffixConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20220301:AppServiceEnvironmentAseCustomDnsSuffixConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServiceEnvironmentAseCustomDnsSuffixConfiguration
	err := ctx.RegisterResource("azure-native:web:AppServiceEnvironmentAseCustomDnsSuffixConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppServiceEnvironmentAseCustomDnsSuffixConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceEnvironmentAseCustomDnsSuffixConfigurationState, opts ...pulumi.ResourceOption) (*AppServiceEnvironmentAseCustomDnsSuffixConfiguration, error) {
	var resource AppServiceEnvironmentAseCustomDnsSuffixConfiguration
	err := ctx.ReadResource("azure-native:web:AppServiceEnvironmentAseCustomDnsSuffixConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appServiceEnvironmentAseCustomDnsSuffixConfigurationState struct {
}

type AppServiceEnvironmentAseCustomDnsSuffixConfigurationState struct {
}

func (AppServiceEnvironmentAseCustomDnsSuffixConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentAseCustomDnsSuffixConfigurationState)(nil)).Elem()
}

type appServiceEnvironmentAseCustomDnsSuffixConfigurationArgs struct {
	CertificateUrl            *string `pulumi:"certificateUrl"`
	DnsSuffix                 *string `pulumi:"dnsSuffix"`
	KeyVaultReferenceIdentity *string `pulumi:"keyVaultReferenceIdentity"`
	Kind                      *string `pulumi:"kind"`
	Name                      string  `pulumi:"name"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
}


type AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs struct {
	CertificateUrl            pulumi.StringPtrInput
	DnsSuffix                 pulumi.StringPtrInput
	KeyVaultReferenceIdentity pulumi.StringPtrInput
	Kind                      pulumi.StringPtrInput
	Name                      pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
}

func (AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentAseCustomDnsSuffixConfigurationArgs)(nil)).Elem()
}

type AppServiceEnvironmentAseCustomDnsSuffixConfigurationInput interface {
	pulumi.Input

	ToAppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput() AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput
	ToAppServiceEnvironmentAseCustomDnsSuffixConfigurationOutputWithContext(ctx context.Context) AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput
}

func (*AppServiceEnvironmentAseCustomDnsSuffixConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceEnvironmentAseCustomDnsSuffixConfiguration)(nil)).Elem()
}

func (i *AppServiceEnvironmentAseCustomDnsSuffixConfiguration) ToAppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput() AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput {
	return i.ToAppServiceEnvironmentAseCustomDnsSuffixConfigurationOutputWithContext(context.Background())
}

func (i *AppServiceEnvironmentAseCustomDnsSuffixConfiguration) ToAppServiceEnvironmentAseCustomDnsSuffixConfigurationOutputWithContext(ctx context.Context) AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput)
}

type AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput struct{ *pulumi.OutputState }

func (AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceEnvironmentAseCustomDnsSuffixConfiguration)(nil)).Elem()
}

func (o AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput) ToAppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput() AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput {
	return o
}

func (o AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput) ToAppServiceEnvironmentAseCustomDnsSuffixConfigurationOutputWithContext(ctx context.Context) AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput {
	return o
}

func (o AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentAseCustomDnsSuffixConfiguration) pulumi.StringPtrOutput {
		return v.CertificateUrl
	}).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentAseCustomDnsSuffixConfiguration) pulumi.StringPtrOutput {
		return v.DnsSuffix
	}).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentAseCustomDnsSuffixConfiguration) pulumi.StringPtrOutput {
		return v.KeyVaultReferenceIdentity
	}).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentAseCustomDnsSuffixConfiguration) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentAseCustomDnsSuffixConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput) ProvisioningDetails() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentAseCustomDnsSuffixConfiguration) pulumi.StringOutput {
		return v.ProvisioningDetails
	}).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentAseCustomDnsSuffixConfiguration) pulumi.StringOutput {
		return v.ProvisioningState
	}).(pulumi.StringOutput)
}

func (o AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentAseCustomDnsSuffixConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AppServiceEnvironmentAseCustomDnsSuffixConfigurationOutput{})
}
