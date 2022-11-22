


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Address struct {
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


func NewAddress(ctx *pulumi.Context,
	name string, args *AddressArgs, opts ...pulumi.ResourceOption) (*Address, error) {
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
			Type: pulumi.String("azure-native:edgeorder:Address"),
		},
		{
			Type: pulumi.String("azure-native:edgeorder/v20201201preview:Address"),
		},
		{
			Type: pulumi.String("azure-native:edgeorder/v20211201:Address"),
		},
	})
	opts = append(opts, aliases)
	var resource Address
	err := ctx.RegisterResource("azure-native:edgeorder/v20220501preview:Address", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressState, opts ...pulumi.ResourceOption) (*Address, error) {
	var resource Address
	err := ctx.ReadResource("azure-native:edgeorder/v20220501preview:Address", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type addressState struct {
}

type AddressState struct {
}

func (AddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressState)(nil)).Elem()
}

type addressArgs struct {
	AddressName       *string           `pulumi:"addressName"`
	ContactDetails    ContactDetails    `pulumi:"contactDetails"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ShippingAddress   *ShippingAddress  `pulumi:"shippingAddress"`
	Tags              map[string]string `pulumi:"tags"`
}


type AddressArgs struct {
	AddressName       pulumi.StringPtrInput
	ContactDetails    ContactDetailsInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ShippingAddress   ShippingAddressPtrInput
	Tags              pulumi.StringMapInput
}

func (AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressArgs)(nil)).Elem()
}

type AddressInput interface {
	pulumi.Input

	ToAddressOutput() AddressOutput
	ToAddressOutputWithContext(ctx context.Context) AddressOutput
}

func (*Address) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (i *Address) ToAddressOutput() AddressOutput {
	return i.ToAddressOutputWithContext(context.Background())
}

func (i *Address) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput)
}

type AddressOutput struct{ *pulumi.OutputState }

func (AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (o AddressOutput) ToAddressOutput() AddressOutput {
	return o
}

func (o AddressOutput) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return o
}

func (o AddressOutput) AddressValidationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.AddressValidationStatus }).(pulumi.StringOutput)
}

func (o AddressOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v *Address) ContactDetailsResponseOutput { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o AddressOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AddressOutput) ShippingAddress() ShippingAddressResponsePtrOutput {
	return o.ApplyT(func(v *Address) ShippingAddressResponsePtrOutput { return v.ShippingAddress }).(ShippingAddressResponsePtrOutput)
}

func (o AddressOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Address) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AddressOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Address) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AddressOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AddressOutput{})
}
