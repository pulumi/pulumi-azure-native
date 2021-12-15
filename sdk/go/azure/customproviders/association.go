


package customproviders

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Association struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	TargetResourceId  pulumi.StringPtrOutput `pulumi:"targetResourceId"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewAssociation(ctx *pulumi.Context,
	name string, args *AssociationArgs, opts ...pulumi.ResourceOption) (*Association, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customproviders/v20180901preview:Association"),
		},
	})
	opts = append(opts, aliases)
	var resource Association
	err := ctx.RegisterResource("azure-native:customproviders:Association", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssociationState, opts ...pulumi.ResourceOption) (*Association, error) {
	var resource Association
	err := ctx.ReadResource("azure-native:customproviders:Association", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type associationState struct {
}

type AssociationState struct {
}

func (AssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*associationState)(nil)).Elem()
}

type associationArgs struct {
	AssociationName  *string `pulumi:"associationName"`
	Scope            string  `pulumi:"scope"`
	TargetResourceId *string `pulumi:"targetResourceId"`
}


type AssociationArgs struct {
	AssociationName  pulumi.StringPtrInput
	Scope            pulumi.StringInput
	TargetResourceId pulumi.StringPtrInput
}

func (AssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*associationArgs)(nil)).Elem()
}

type AssociationInput interface {
	pulumi.Input

	ToAssociationOutput() AssociationOutput
	ToAssociationOutputWithContext(ctx context.Context) AssociationOutput
}

func (*Association) ElementType() reflect.Type {
	return reflect.TypeOf((**Association)(nil)).Elem()
}

func (i *Association) ToAssociationOutput() AssociationOutput {
	return i.ToAssociationOutputWithContext(context.Background())
}

func (i *Association) ToAssociationOutputWithContext(ctx context.Context) AssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationOutput)
}

type AssociationOutput struct{ *pulumi.OutputState }

func (AssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Association)(nil)).Elem()
}

func (o AssociationOutput) ToAssociationOutput() AssociationOutput {
	return o
}

func (o AssociationOutput) ToAssociationOutputWithContext(ctx context.Context) AssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AssociationOutput{})
}
