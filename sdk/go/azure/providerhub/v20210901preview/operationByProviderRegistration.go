


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OperationByProviderRegistration struct {
	pulumi.CustomResourceState

	Name pulumi.StringOutput `pulumi:"name"`
	Type pulumi.StringOutput `pulumi:"type"`
}


func NewOperationByProviderRegistration(ctx *pulumi.Context,
	name string, args *OperationByProviderRegistrationArgs, opts ...pulumi.ResourceOption) (*OperationByProviderRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:providerhub:OperationByProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:OperationByProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:OperationByProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210601preview:OperationByProviderRegistration"),
		},
	})
	opts = append(opts, aliases)
	var resource OperationByProviderRegistration
	err := ctx.RegisterResource("azure-native:providerhub/v20210901preview:OperationByProviderRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOperationByProviderRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OperationByProviderRegistrationState, opts ...pulumi.ResourceOption) (*OperationByProviderRegistration, error) {
	var resource OperationByProviderRegistration
	err := ctx.ReadResource("azure-native:providerhub/v20210901preview:OperationByProviderRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type operationByProviderRegistrationState struct {
}

type OperationByProviderRegistrationState struct {
}

func (OperationByProviderRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*operationByProviderRegistrationState)(nil)).Elem()
}

type operationByProviderRegistrationArgs struct {
	ProviderNamespace string `pulumi:"providerNamespace"`
}


type OperationByProviderRegistrationArgs struct {
	ProviderNamespace pulumi.StringInput
}

func (OperationByProviderRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*operationByProviderRegistrationArgs)(nil)).Elem()
}

type OperationByProviderRegistrationInput interface {
	pulumi.Input

	ToOperationByProviderRegistrationOutput() OperationByProviderRegistrationOutput
	ToOperationByProviderRegistrationOutputWithContext(ctx context.Context) OperationByProviderRegistrationOutput
}

func (*OperationByProviderRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationByProviderRegistration)(nil)).Elem()
}

func (i *OperationByProviderRegistration) ToOperationByProviderRegistrationOutput() OperationByProviderRegistrationOutput {
	return i.ToOperationByProviderRegistrationOutputWithContext(context.Background())
}

func (i *OperationByProviderRegistration) ToOperationByProviderRegistrationOutputWithContext(ctx context.Context) OperationByProviderRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperationByProviderRegistrationOutput)
}

type OperationByProviderRegistrationOutput struct{ *pulumi.OutputState }

func (OperationByProviderRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationByProviderRegistration)(nil)).Elem()
}

func (o OperationByProviderRegistrationOutput) ToOperationByProviderRegistrationOutput() OperationByProviderRegistrationOutput {
	return o
}

func (o OperationByProviderRegistrationOutput) ToOperationByProviderRegistrationOutputWithContext(ctx context.Context) OperationByProviderRegistrationOutput {
	return o
}

func (o OperationByProviderRegistrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationByProviderRegistration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OperationByProviderRegistrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationByProviderRegistration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OperationByProviderRegistrationOutput{})
}
