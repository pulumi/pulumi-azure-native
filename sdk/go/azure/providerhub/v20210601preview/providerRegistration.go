


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProviderRegistration struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                          `pulumi:"name"`
	Properties ProviderRegistrationResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                     `pulumi:"systemData"`
	Type       pulumi.StringOutput                          `pulumi:"type"`
}


func NewProviderRegistration(ctx *pulumi.Context,
	name string, args *ProviderRegistrationArgs, opts ...pulumi.ResourceOption) (*ProviderRegistration, error) {
	if args == nil {
		args = &ProviderRegistrationArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20210601preview:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20201120:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20210501preview:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210901preview:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20210901preview:ProviderRegistration"),
		},
	})
	opts = append(opts, aliases)
	var resource ProviderRegistration
	err := ctx.RegisterResource("azure-native:providerhub/v20210601preview:ProviderRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProviderRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderRegistrationState, opts ...pulumi.ResourceOption) (*ProviderRegistration, error) {
	var resource ProviderRegistration
	err := ctx.ReadResource("azure-native:providerhub/v20210601preview:ProviderRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type providerRegistrationState struct {
}

type ProviderRegistrationState struct {
}

func (ProviderRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerRegistrationState)(nil)).Elem()
}

type providerRegistrationArgs struct {
	Properties        *ProviderRegistrationProperties `pulumi:"properties"`
	ProviderNamespace *string                         `pulumi:"providerNamespace"`
}


type ProviderRegistrationArgs struct {
	Properties        ProviderRegistrationPropertiesPtrInput
	ProviderNamespace pulumi.StringPtrInput
}

func (ProviderRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerRegistrationArgs)(nil)).Elem()
}

type ProviderRegistrationInput interface {
	pulumi.Input

	ToProviderRegistrationOutput() ProviderRegistrationOutput
	ToProviderRegistrationOutputWithContext(ctx context.Context) ProviderRegistrationOutput
}

func (*ProviderRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderRegistration)(nil))
}

func (i *ProviderRegistration) ToProviderRegistrationOutput() ProviderRegistrationOutput {
	return i.ToProviderRegistrationOutputWithContext(context.Background())
}

func (i *ProviderRegistration) ToProviderRegistrationOutputWithContext(ctx context.Context) ProviderRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderRegistrationOutput)
}

type ProviderRegistrationOutput struct{ *pulumi.OutputState }

func (ProviderRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderRegistration)(nil))
}

func (o ProviderRegistrationOutput) ToProviderRegistrationOutput() ProviderRegistrationOutput {
	return o
}

func (o ProviderRegistrationOutput) ToProviderRegistrationOutputWithContext(ctx context.Context) ProviderRegistrationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderRegistrationInput)(nil)).Elem(), &ProviderRegistration{})
	pulumi.RegisterOutputType(ProviderRegistrationOutput{})
}
