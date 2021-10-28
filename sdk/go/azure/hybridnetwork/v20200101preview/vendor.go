


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Vendor struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput            `pulumi:"name"`
	ProvisioningState pulumi.StringOutput            `pulumi:"provisioningState"`
	Skus              SubResourceResponseArrayOutput `pulumi:"skus"`
	Type              pulumi.StringOutput            `pulumi:"type"`
}


func NewVendor(ctx *pulumi.Context,
	name string, args *VendorArgs, opts ...pulumi.ResourceOption) (*Vendor, error) {
	if args == nil {
		args = &VendorArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:hybridnetwork/v20200101preview:Vendor"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork:Vendor"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridnetwork:Vendor"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20210501:Vendor"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridnetwork/v20210501:Vendor"),
		},
	})
	opts = append(opts, aliases)
	var resource Vendor
	err := ctx.RegisterResource("azure-native:hybridnetwork/v20200101preview:Vendor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVendor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VendorState, opts ...pulumi.ResourceOption) (*Vendor, error) {
	var resource Vendor
	err := ctx.ReadResource("azure-native:hybridnetwork/v20200101preview:Vendor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vendorState struct {
}

type VendorState struct {
}

func (VendorState) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorState)(nil)).Elem()
}

type vendorArgs struct {
	VendorName *string `pulumi:"vendorName"`
}


type VendorArgs struct {
	VendorName pulumi.StringPtrInput
}

func (VendorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorArgs)(nil)).Elem()
}

type VendorInput interface {
	pulumi.Input

	ToVendorOutput() VendorOutput
	ToVendorOutputWithContext(ctx context.Context) VendorOutput
}

func (*Vendor) ElementType() reflect.Type {
	return reflect.TypeOf((*Vendor)(nil))
}

func (i *Vendor) ToVendorOutput() VendorOutput {
	return i.ToVendorOutputWithContext(context.Background())
}

func (i *Vendor) ToVendorOutputWithContext(ctx context.Context) VendorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VendorOutput)
}

type VendorOutput struct{ *pulumi.OutputState }

func (VendorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Vendor)(nil))
}

func (o VendorOutput) ToVendorOutput() VendorOutput {
	return o
}

func (o VendorOutput) ToVendorOutputWithContext(ctx context.Context) VendorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VendorOutput{})
}
