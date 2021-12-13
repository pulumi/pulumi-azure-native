


package features

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubscriptionFeatureRegistration struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                                     `pulumi:"name"`
	Properties SubscriptionFeatureRegistrationResponsePropertiesOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                                     `pulumi:"type"`
}


func NewSubscriptionFeatureRegistration(ctx *pulumi.Context,
	name string, args *SubscriptionFeatureRegistrationArgs, opts ...pulumi.ResourceOption) (*SubscriptionFeatureRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToSubscriptionFeatureRegistrationPropertiesPtrOutput().ApplyT(func(v *SubscriptionFeatureRegistrationProperties) *SubscriptionFeatureRegistrationProperties {
			return v.Defaults()
		}).(SubscriptionFeatureRegistrationPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:features/v20210701:SubscriptionFeatureRegistration"),
		},
	})
	opts = append(opts, aliases)
	var resource SubscriptionFeatureRegistration
	err := ctx.RegisterResource("azure-native:features:SubscriptionFeatureRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubscriptionFeatureRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionFeatureRegistrationState, opts ...pulumi.ResourceOption) (*SubscriptionFeatureRegistration, error) {
	var resource SubscriptionFeatureRegistration
	err := ctx.ReadResource("azure-native:features:SubscriptionFeatureRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type subscriptionFeatureRegistrationState struct {
}

type SubscriptionFeatureRegistrationState struct {
}

func (SubscriptionFeatureRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionFeatureRegistrationState)(nil)).Elem()
}

type subscriptionFeatureRegistrationArgs struct {
	FeatureName       *string                                    `pulumi:"featureName"`
	Properties        *SubscriptionFeatureRegistrationProperties `pulumi:"properties"`
	ProviderNamespace string                                     `pulumi:"providerNamespace"`
}


type SubscriptionFeatureRegistrationArgs struct {
	FeatureName       pulumi.StringPtrInput
	Properties        SubscriptionFeatureRegistrationPropertiesPtrInput
	ProviderNamespace pulumi.StringInput
}

func (SubscriptionFeatureRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionFeatureRegistrationArgs)(nil)).Elem()
}

type SubscriptionFeatureRegistrationInput interface {
	pulumi.Input

	ToSubscriptionFeatureRegistrationOutput() SubscriptionFeatureRegistrationOutput
	ToSubscriptionFeatureRegistrationOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationOutput
}

func (*SubscriptionFeatureRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionFeatureRegistration)(nil)).Elem()
}

func (i *SubscriptionFeatureRegistration) ToSubscriptionFeatureRegistrationOutput() SubscriptionFeatureRegistrationOutput {
	return i.ToSubscriptionFeatureRegistrationOutputWithContext(context.Background())
}

func (i *SubscriptionFeatureRegistration) ToSubscriptionFeatureRegistrationOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionFeatureRegistrationOutput)
}

type SubscriptionFeatureRegistrationOutput struct{ *pulumi.OutputState }

func (SubscriptionFeatureRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionFeatureRegistration)(nil)).Elem()
}

func (o SubscriptionFeatureRegistrationOutput) ToSubscriptionFeatureRegistrationOutput() SubscriptionFeatureRegistrationOutput {
	return o
}

func (o SubscriptionFeatureRegistrationOutput) ToSubscriptionFeatureRegistrationOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubscriptionFeatureRegistrationOutput{})
}
