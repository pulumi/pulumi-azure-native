


package azure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Provider struct {
	pulumi.ProviderResourceState
}


func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if isZero(args.Environment) {
		args.Environment = pulumi.StringPtr("public")
	}
	if isZero(args.UseMsi) {
		args.UseMsi = pulumi.BoolPtr(false)
	}
	if args.ClientCertificatePassword != nil {
		args.ClientCertificatePassword = pulumi.ToSecret(args.ClientCertificatePassword).(pulumi.StringPtrOutput)
	}
	if args.ClientId != nil {
		args.ClientId = pulumi.ToSecret(args.ClientId).(pulumi.StringPtrOutput)
	}
	if args.ClientSecret != nil {
		args.ClientSecret = pulumi.ToSecret(args.ClientSecret).(pulumi.StringPtrOutput)
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:azure-native", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// Any additional Tenant IDs which should be used for authentication.
	AuxiliaryTenantIds []string `pulumi:"auxiliaryTenantIds"`
	// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate
	ClientCertificatePassword *string `pulumi:"clientCertificatePassword"`
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.
	ClientCertificatePath *string `pulumi:"clientCertificatePath"`
	// The Client ID which should be used.
	ClientId *string `pulumi:"clientId"`
	// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
	ClientSecret *string `pulumi:"clientSecret"`
	// This will disable the Pulumi Partner ID which is used if a custom `partnerId` isn't specified.
	DisablePulumiPartnerId *bool `pulumi:"disablePulumiPartnerId"`
	// The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to public.
	Environment *string `pulumi:"environment"`
	// The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically.
	MsiEndpoint *string `pulumi:"msiEndpoint"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
	PartnerId *string `pulumi:"partnerId"`
	// The Subscription ID which should be used.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The Tenant ID which should be used.
	TenantId *string `pulumi:"tenantId"`
	// Allowed Managed Service Identity be used for Authentication.
	UseMsi *bool `pulumi:"useMsi"`
}


type ProviderArgs struct {
	// Any additional Tenant IDs which should be used for authentication.
	AuxiliaryTenantIds pulumi.StringArrayInput
	// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate
	ClientCertificatePassword pulumi.StringPtrInput
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.
	ClientCertificatePath pulumi.StringPtrInput
	// The Client ID which should be used.
	ClientId pulumi.StringPtrInput
	// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
	ClientSecret pulumi.StringPtrInput
	// This will disable the Pulumi Partner ID which is used if a custom `partnerId` isn't specified.
	DisablePulumiPartnerId pulumi.BoolPtrInput
	// The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to public.
	Environment pulumi.StringPtrInput
	// The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically.
	MsiEndpoint pulumi.StringPtrInput
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
	PartnerId pulumi.StringPtrInput
	// The Subscription ID which should be used.
	SubscriptionId pulumi.StringPtrInput
	// The Tenant ID which should be used.
	TenantId pulumi.StringPtrInput
	// Allowed Managed Service Identity be used for Authentication.
	UseMsi pulumi.BoolPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
}
