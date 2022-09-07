


package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddressByName struct {
	pulumi.CustomResourceState

	AddressValidationStatus pulumi.StringOutput              `pulumi:"addressValidationStatus"`
	ContactDetails          ContactDetailsResponseOutput     `pulumi:"contactDetails"`
	Location                pulumi.StringOutput              `pulumi:"location"`
	Name                    pulumi.StringOutput              `pulumi:"name"`
	ShippingAddress         ShippingAddressResponsePtrOutput `pulumi:"shippingAddress"`
	SystemData              SystemDataResponseOutput         `pulumi:"systemData"`
	Tags                    pulumi.StringMapOutput           `pulumi:"tags"`
	Type                    pulumi.StringOutput              `pulumi:"type"`
}


func NewAddressByName(ctx *pulumi.Context,
	name string, args *AddressByNameArgs, opts ...pulumi.ResourceOption) (*AddressByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContactDetails == nil {
		return nil, errors.New("invalid value for required argument 'ContactDetails'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:edgeorder:AddressByName"),
		},
		{
			Type: pulumi.String("azure-native:edgeorder/v20201201preview:AddressByName"),
		},
		{
			Type: pulumi.String("azure-native:edgeorder/v20220501preview:AddressByName"),
		},
	})
	opts = append(opts, aliases)
	var resource AddressByName
	err := ctx.RegisterResource("azure-native:edgeorder/v20211201:AddressByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAddressByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressByNameState, opts ...pulumi.ResourceOption) (*AddressByName, error) {
	var resource AddressByName
	err := ctx.ReadResource("azure-native:edgeorder/v20211201:AddressByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type addressByNameState struct {
}

type AddressByNameState struct {
}

func (AddressByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressByNameState)(nil)).Elem()
}

type addressByNameArgs struct {
	AddressName       *string           `pulumi:"addressName"`
	ContactDetails    ContactDetails    `pulumi:"contactDetails"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ShippingAddress   *ShippingAddress  `pulumi:"shippingAddress"`
	Tags              map[string]string `pulumi:"tags"`
}


type AddressByNameArgs struct {
	AddressName       pulumi.StringPtrInput
	ContactDetails    ContactDetailsInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ShippingAddress   ShippingAddressPtrInput
	Tags              pulumi.StringMapInput
}

func (AddressByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressByNameArgs)(nil)).Elem()
}

type AddressByNameInput interface {
	pulumi.Input

	ToAddressByNameOutput() AddressByNameOutput
	ToAddressByNameOutputWithContext(ctx context.Context) AddressByNameOutput
}

func (*AddressByName) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressByName)(nil)).Elem()
}

func (i *AddressByName) ToAddressByNameOutput() AddressByNameOutput {
	return i.ToAddressByNameOutputWithContext(context.Background())
}

func (i *AddressByName) ToAddressByNameOutputWithContext(ctx context.Context) AddressByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressByNameOutput)
}

type AddressByNameOutput struct{ *pulumi.OutputState }

func (AddressByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressByName)(nil)).Elem()
}

func (o AddressByNameOutput) ToAddressByNameOutput() AddressByNameOutput {
	return o
}

func (o AddressByNameOutput) ToAddressByNameOutputWithContext(ctx context.Context) AddressByNameOutput {
	return o
}

func (o AddressByNameOutput) AddressValidationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressByName) pulumi.StringOutput { return v.AddressValidationStatus }).(pulumi.StringOutput)
}

func (o AddressByNameOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v *AddressByName) ContactDetailsResponseOutput { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o AddressByNameOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressByName) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AddressByNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressByName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AddressByNameOutput) ShippingAddress() ShippingAddressResponsePtrOutput {
	return o.ApplyT(func(v *AddressByName) ShippingAddressResponsePtrOutput { return v.ShippingAddress }).(ShippingAddressResponsePtrOutput)
}

func (o AddressByNameOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AddressByName) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AddressByNameOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AddressByName) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AddressByNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressByName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AddressByNameOutput{})
}
