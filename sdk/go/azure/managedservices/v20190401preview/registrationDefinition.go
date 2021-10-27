


package v20190401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistrationDefinition struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                            `pulumi:"name"`
	Plan       PlanResponsePtrOutput                          `pulumi:"plan"`
	Properties RegistrationDefinitionPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                            `pulumi:"type"`
}


func NewRegistrationDefinition(ctx *pulumi.Context,
	name string, args *RegistrationDefinitionArgs, opts ...pulumi.ResourceOption) (*RegistrationDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:managedservices/v20190401preview:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:managedservices:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:managedservices:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20180601preview:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:managedservices/v20180601preview:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20190601:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:managedservices/v20190601:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20190901:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:managedservices/v20190901:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20200201preview:RegistrationDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:managedservices/v20200201preview:RegistrationDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource RegistrationDefinition
	err := ctx.RegisterResource("azure-native:managedservices/v20190401preview:RegistrationDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistrationDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistrationDefinitionState, opts ...pulumi.ResourceOption) (*RegistrationDefinition, error) {
	var resource RegistrationDefinition
	err := ctx.ReadResource("azure-native:managedservices/v20190401preview:RegistrationDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registrationDefinitionState struct {
}

type RegistrationDefinitionState struct {
}

func (RegistrationDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationDefinitionState)(nil)).Elem()
}

type registrationDefinitionArgs struct {
	Plan                     *Plan                             `pulumi:"plan"`
	Properties               *RegistrationDefinitionProperties `pulumi:"properties"`
	RegistrationDefinitionId *string                           `pulumi:"registrationDefinitionId"`
	Scope                    string                            `pulumi:"scope"`
}


type RegistrationDefinitionArgs struct {
	Plan                     PlanPtrInput
	Properties               RegistrationDefinitionPropertiesPtrInput
	RegistrationDefinitionId pulumi.StringPtrInput
	Scope                    pulumi.StringInput
}

func (RegistrationDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationDefinitionArgs)(nil)).Elem()
}

type RegistrationDefinitionInput interface {
	pulumi.Input

	ToRegistrationDefinitionOutput() RegistrationDefinitionOutput
	ToRegistrationDefinitionOutputWithContext(ctx context.Context) RegistrationDefinitionOutput
}

func (*RegistrationDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationDefinition)(nil))
}

func (i *RegistrationDefinition) ToRegistrationDefinitionOutput() RegistrationDefinitionOutput {
	return i.ToRegistrationDefinitionOutputWithContext(context.Background())
}

func (i *RegistrationDefinition) ToRegistrationDefinitionOutputWithContext(ctx context.Context) RegistrationDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationDefinitionOutput)
}

type RegistrationDefinitionOutput struct{ *pulumi.OutputState }

func (RegistrationDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationDefinition)(nil))
}

func (o RegistrationDefinitionOutput) ToRegistrationDefinitionOutput() RegistrationDefinitionOutput {
	return o
}

func (o RegistrationDefinitionOutput) ToRegistrationDefinitionOutputWithContext(ctx context.Context) RegistrationDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistrationDefinitionInput)(nil)).Elem(), &RegistrationDefinition{})
	pulumi.RegisterOutputType(RegistrationDefinitionOutput{})
}
