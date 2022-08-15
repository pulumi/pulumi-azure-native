


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NspAssociation struct {
	pulumi.CustomResourceState

	AccessMode            pulumi.StringPtrOutput       `pulumi:"accessMode"`
	HasProvisioningIssues pulumi.StringOutput          `pulumi:"hasProvisioningIssues"`
	Location              pulumi.StringPtrOutput       `pulumi:"location"`
	Name                  pulumi.StringOutput          `pulumi:"name"`
	PrivateLinkResource   SubResourceResponsePtrOutput `pulumi:"privateLinkResource"`
	Profile               SubResourceResponsePtrOutput `pulumi:"profile"`
	ProvisioningState     pulumi.StringOutput          `pulumi:"provisioningState"`
	Tags                  pulumi.StringMapOutput       `pulumi:"tags"`
	Type                  pulumi.StringOutput          `pulumi:"type"`
}


func NewNspAssociation(ctx *pulumi.Context,
	name string, args *NspAssociationArgs, opts ...pulumi.ResourceOption) (*NspAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkSecurityPerimeterName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityPerimeterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210201preview:NspAssociation"),
		},
	})
	opts = append(opts, aliases)
	var resource NspAssociation
	err := ctx.RegisterResource("azure-native:network:NspAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNspAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NspAssociationState, opts ...pulumi.ResourceOption) (*NspAssociation, error) {
	var resource NspAssociation
	err := ctx.ReadResource("azure-native:network:NspAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type nspAssociationState struct {
}

type NspAssociationState struct {
}

func (NspAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*nspAssociationState)(nil)).Elem()
}

type nspAssociationArgs struct {
	AccessMode                   *string           `pulumi:"accessMode"`
	AssociationName              *string           `pulumi:"associationName"`
	Id                           *string           `pulumi:"id"`
	Location                     *string           `pulumi:"location"`
	Name                         *string           `pulumi:"name"`
	NetworkSecurityPerimeterName string            `pulumi:"networkSecurityPerimeterName"`
	PrivateLinkResource          *SubResource      `pulumi:"privateLinkResource"`
	Profile                      *SubResource      `pulumi:"profile"`
	ResourceGroupName            string            `pulumi:"resourceGroupName"`
	Tags                         map[string]string `pulumi:"tags"`
}


type NspAssociationArgs struct {
	AccessMode                   pulumi.StringPtrInput
	AssociationName              pulumi.StringPtrInput
	Id                           pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	NetworkSecurityPerimeterName pulumi.StringInput
	PrivateLinkResource          SubResourcePtrInput
	Profile                      SubResourcePtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
}

func (NspAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nspAssociationArgs)(nil)).Elem()
}

type NspAssociationInput interface {
	pulumi.Input

	ToNspAssociationOutput() NspAssociationOutput
	ToNspAssociationOutputWithContext(ctx context.Context) NspAssociationOutput
}

func (*NspAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**NspAssociation)(nil)).Elem()
}

func (i *NspAssociation) ToNspAssociationOutput() NspAssociationOutput {
	return i.ToNspAssociationOutputWithContext(context.Background())
}

func (i *NspAssociation) ToNspAssociationOutputWithContext(ctx context.Context) NspAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NspAssociationOutput)
}

type NspAssociationOutput struct{ *pulumi.OutputState }

func (NspAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NspAssociation)(nil)).Elem()
}

func (o NspAssociationOutput) ToNspAssociationOutput() NspAssociationOutput {
	return o
}

func (o NspAssociationOutput) ToNspAssociationOutputWithContext(ctx context.Context) NspAssociationOutput {
	return o
}

func (o NspAssociationOutput) AccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringPtrOutput { return v.AccessMode }).(pulumi.StringPtrOutput)
}

func (o NspAssociationOutput) HasProvisioningIssues() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringOutput { return v.HasProvisioningIssues }).(pulumi.StringOutput)
}

func (o NspAssociationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NspAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NspAssociationOutput) PrivateLinkResource() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NspAssociation) SubResourceResponsePtrOutput { return v.PrivateLinkResource }).(SubResourceResponsePtrOutput)
}

func (o NspAssociationOutput) Profile() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NspAssociation) SubResourceResponsePtrOutput { return v.Profile }).(SubResourceResponsePtrOutput)
}

func (o NspAssociationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NspAssociationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NspAssociationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NspAssociationOutput{})
}
