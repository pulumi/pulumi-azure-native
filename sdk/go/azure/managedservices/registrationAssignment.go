


package managedservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegistrationAssignment struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                            `pulumi:"name"`
	Properties RegistrationAssignmentPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                            `pulumi:"type"`
}


func NewRegistrationAssignment(ctx *pulumi.Context,
	name string, args *RegistrationAssignmentArgs, opts ...pulumi.ResourceOption) (*RegistrationAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:managedservices:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20180601preview:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:managedservices/v20180601preview:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20190401preview:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:managedservices/v20190401preview:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20190601:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:managedservices/v20190601:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20190901:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:managedservices/v20190901:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:managedservices/v20200201preview:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azure-nextgen:managedservices/v20200201preview:RegistrationAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource RegistrationAssignment
	err := ctx.RegisterResource("azure-native:managedservices:RegistrationAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistrationAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistrationAssignmentState, opts ...pulumi.ResourceOption) (*RegistrationAssignment, error) {
	var resource RegistrationAssignment
	err := ctx.ReadResource("azure-native:managedservices:RegistrationAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registrationAssignmentState struct {
}

type RegistrationAssignmentState struct {
}

func (RegistrationAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationAssignmentState)(nil)).Elem()
}

type registrationAssignmentArgs struct {
	Properties               *RegistrationAssignmentProperties `pulumi:"properties"`
	RegistrationAssignmentId *string                           `pulumi:"registrationAssignmentId"`
	Scope                    string                            `pulumi:"scope"`
}


type RegistrationAssignmentArgs struct {
	Properties               RegistrationAssignmentPropertiesPtrInput
	RegistrationAssignmentId pulumi.StringPtrInput
	Scope                    pulumi.StringInput
}

func (RegistrationAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationAssignmentArgs)(nil)).Elem()
}

type RegistrationAssignmentInput interface {
	pulumi.Input

	ToRegistrationAssignmentOutput() RegistrationAssignmentOutput
	ToRegistrationAssignmentOutputWithContext(ctx context.Context) RegistrationAssignmentOutput
}

func (*RegistrationAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignment)(nil))
}

func (i *RegistrationAssignment) ToRegistrationAssignmentOutput() RegistrationAssignmentOutput {
	return i.ToRegistrationAssignmentOutputWithContext(context.Background())
}

func (i *RegistrationAssignment) ToRegistrationAssignmentOutputWithContext(ctx context.Context) RegistrationAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentOutput)
}

type RegistrationAssignmentOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignment)(nil))
}

func (o RegistrationAssignmentOutput) ToRegistrationAssignmentOutput() RegistrationAssignmentOutput {
	return o
}

func (o RegistrationAssignmentOutput) ToRegistrationAssignmentOutputWithContext(ctx context.Context) RegistrationAssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegistrationAssignmentOutput{})
}
