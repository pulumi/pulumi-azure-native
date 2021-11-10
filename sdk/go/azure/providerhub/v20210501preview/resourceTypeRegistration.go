


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceTypeRegistration struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                              `pulumi:"name"`
	Properties ResourceTypeRegistrationResponsePropertiesOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                              `pulumi:"type"`
}


func NewResourceTypeRegistration(ctx *pulumi.Context,
	name string, args *ResourceTypeRegistrationArgs, opts ...pulumi.ResourceOption) (*ResourceTypeRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:providerhub:ResourceTypeRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:ResourceTypeRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210601preview:ResourceTypeRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210901preview:ResourceTypeRegistration"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourceTypeRegistration
	err := ctx.RegisterResource("azure-native:providerhub/v20210501preview:ResourceTypeRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetResourceTypeRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceTypeRegistrationState, opts ...pulumi.ResourceOption) (*ResourceTypeRegistration, error) {
	var resource ResourceTypeRegistration
	err := ctx.ReadResource("azure-native:providerhub/v20210501preview:ResourceTypeRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type resourceTypeRegistrationState struct {
}

type ResourceTypeRegistrationState struct {
}

func (ResourceTypeRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceTypeRegistrationState)(nil)).Elem()
}

type resourceTypeRegistrationArgs struct {
	Properties        *ResourceTypeRegistrationProperties `pulumi:"properties"`
	ProviderNamespace string                              `pulumi:"providerNamespace"`
	ResourceType      *string                             `pulumi:"resourceType"`
}


type ResourceTypeRegistrationArgs struct {
	Properties        ResourceTypeRegistrationPropertiesPtrInput
	ProviderNamespace pulumi.StringInput
	ResourceType      pulumi.StringPtrInput
}

func (ResourceTypeRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceTypeRegistrationArgs)(nil)).Elem()
}

type ResourceTypeRegistrationInput interface {
	pulumi.Input

	ToResourceTypeRegistrationOutput() ResourceTypeRegistrationOutput
	ToResourceTypeRegistrationOutputWithContext(ctx context.Context) ResourceTypeRegistrationOutput
}

func (*ResourceTypeRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceTypeRegistration)(nil))
}

func (i *ResourceTypeRegistration) ToResourceTypeRegistrationOutput() ResourceTypeRegistrationOutput {
	return i.ToResourceTypeRegistrationOutputWithContext(context.Background())
}

func (i *ResourceTypeRegistration) ToResourceTypeRegistrationOutputWithContext(ctx context.Context) ResourceTypeRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceTypeRegistrationOutput)
}

type ResourceTypeRegistrationOutput struct{ *pulumi.OutputState }

func (ResourceTypeRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceTypeRegistration)(nil))
}

func (o ResourceTypeRegistrationOutput) ToResourceTypeRegistrationOutput() ResourceTypeRegistrationOutput {
	return o
}

func (o ResourceTypeRegistrationOutput) ToResourceTypeRegistrationOutputWithContext(ctx context.Context) ResourceTypeRegistrationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceTypeRegistrationOutput{})
}
