


package authorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkAssociation struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                                    `pulumi:"name"`
	Properties PrivateLinkAssociationPropertiesExpandedResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                                    `pulumi:"type"`
}


func NewPrivateLinkAssociation(ctx *pulumi.Context,
	name string, args *PrivateLinkAssociationArgs, opts ...pulumi.ResourceOption) (*PrivateLinkAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20200501:PrivateLinkAssociation"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkAssociation
	err := ctx.RegisterResource("azure-native:authorization:PrivateLinkAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkAssociationState, opts ...pulumi.ResourceOption) (*PrivateLinkAssociation, error) {
	var resource PrivateLinkAssociation
	err := ctx.ReadResource("azure-native:authorization:PrivateLinkAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkAssociationState struct {
}

type PrivateLinkAssociationState struct {
}

func (PrivateLinkAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkAssociationState)(nil)).Elem()
}

type privateLinkAssociationArgs struct {
	GroupId    string                            `pulumi:"groupId"`
	PlaId      *string                           `pulumi:"plaId"`
	Properties *PrivateLinkAssociationProperties `pulumi:"properties"`
}


type PrivateLinkAssociationArgs struct {
	GroupId    pulumi.StringInput
	PlaId      pulumi.StringPtrInput
	Properties PrivateLinkAssociationPropertiesPtrInput
}

func (PrivateLinkAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkAssociationArgs)(nil)).Elem()
}

type PrivateLinkAssociationInput interface {
	pulumi.Input

	ToPrivateLinkAssociationOutput() PrivateLinkAssociationOutput
	ToPrivateLinkAssociationOutputWithContext(ctx context.Context) PrivateLinkAssociationOutput
}

func (*PrivateLinkAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkAssociation)(nil)).Elem()
}

func (i *PrivateLinkAssociation) ToPrivateLinkAssociationOutput() PrivateLinkAssociationOutput {
	return i.ToPrivateLinkAssociationOutputWithContext(context.Background())
}

func (i *PrivateLinkAssociation) ToPrivateLinkAssociationOutputWithContext(ctx context.Context) PrivateLinkAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkAssociationOutput)
}

type PrivateLinkAssociationOutput struct{ *pulumi.OutputState }

func (PrivateLinkAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkAssociation)(nil)).Elem()
}

func (o PrivateLinkAssociationOutput) ToPrivateLinkAssociationOutput() PrivateLinkAssociationOutput {
	return o
}

func (o PrivateLinkAssociationOutput) ToPrivateLinkAssociationOutputWithContext(ctx context.Context) PrivateLinkAssociationOutput {
	return o
}

func (o PrivateLinkAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateLinkAssociationOutput) Properties() PrivateLinkAssociationPropertiesExpandedResponseOutput {
	return o.ApplyT(func(v *PrivateLinkAssociation) PrivateLinkAssociationPropertiesExpandedResponseOutput {
		return v.Properties
	}).(PrivateLinkAssociationPropertiesExpandedResponseOutput)
}

func (o PrivateLinkAssociationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkAssociation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkAssociationOutput{})
}
