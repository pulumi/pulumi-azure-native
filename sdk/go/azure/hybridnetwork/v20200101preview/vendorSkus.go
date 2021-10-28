


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VendorSkus struct {
	pulumi.CustomResourceState

	DeploymentMode               pulumi.StringPtrOutput                   `pulumi:"deploymentMode"`
	ManagedApplicationParameters pulumi.AnyOutput                         `pulumi:"managedApplicationParameters"`
	ManagedApplicationTemplate   pulumi.AnyOutput                         `pulumi:"managedApplicationTemplate"`
	Name                         pulumi.StringOutput                      `pulumi:"name"`
	NetworkFunctionTemplate      NetworkFunctionTemplateResponsePtrOutput `pulumi:"networkFunctionTemplate"`
	Preview                      pulumi.BoolPtrOutput                     `pulumi:"preview"`
	ProvisioningState            pulumi.StringOutput                      `pulumi:"provisioningState"`
	SkuType                      pulumi.StringPtrOutput                   `pulumi:"skuType"`
	Type                         pulumi.StringOutput                      `pulumi:"type"`
}


func NewVendorSkus(ctx *pulumi.Context,
	name string, args *VendorSkusArgs, opts ...pulumi.ResourceOption) (*VendorSkus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VendorName == nil {
		return nil, errors.New("invalid value for required argument 'VendorName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:hybridnetwork/v20200101preview:VendorSkus"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork:VendorSkus"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridnetwork:VendorSkus"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20210501:VendorSkus"),
		},
		{
			Type: pulumi.String("azure-nextgen:hybridnetwork/v20210501:VendorSkus"),
		},
	})
	opts = append(opts, aliases)
	var resource VendorSkus
	err := ctx.RegisterResource("azure-native:hybridnetwork/v20200101preview:VendorSkus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVendorSkus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VendorSkusState, opts ...pulumi.ResourceOption) (*VendorSkus, error) {
	var resource VendorSkus
	err := ctx.ReadResource("azure-native:hybridnetwork/v20200101preview:VendorSkus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vendorSkusState struct {
}

type VendorSkusState struct {
}

func (VendorSkusState) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkusState)(nil)).Elem()
}

type vendorSkusArgs struct {
	DeploymentMode               *string                  `pulumi:"deploymentMode"`
	ManagedApplicationParameters interface{}              `pulumi:"managedApplicationParameters"`
	ManagedApplicationTemplate   interface{}              `pulumi:"managedApplicationTemplate"`
	NetworkFunctionTemplate      *NetworkFunctionTemplate `pulumi:"networkFunctionTemplate"`
	Preview                      *bool                    `pulumi:"preview"`
	SkuName                      *string                  `pulumi:"skuName"`
	SkuType                      *string                  `pulumi:"skuType"`
	VendorName                   string                   `pulumi:"vendorName"`
}


type VendorSkusArgs struct {
	DeploymentMode               pulumi.StringPtrInput
	ManagedApplicationParameters pulumi.Input
	ManagedApplicationTemplate   pulumi.Input
	NetworkFunctionTemplate      NetworkFunctionTemplatePtrInput
	Preview                      pulumi.BoolPtrInput
	SkuName                      pulumi.StringPtrInput
	SkuType                      pulumi.StringPtrInput
	VendorName                   pulumi.StringInput
}

func (VendorSkusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkusArgs)(nil)).Elem()
}

type VendorSkusInput interface {
	pulumi.Input

	ToVendorSkusOutput() VendorSkusOutput
	ToVendorSkusOutputWithContext(ctx context.Context) VendorSkusOutput
}

func (*VendorSkus) ElementType() reflect.Type {
	return reflect.TypeOf((*VendorSkus)(nil))
}

func (i *VendorSkus) ToVendorSkusOutput() VendorSkusOutput {
	return i.ToVendorSkusOutputWithContext(context.Background())
}

func (i *VendorSkus) ToVendorSkusOutputWithContext(ctx context.Context) VendorSkusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VendorSkusOutput)
}

type VendorSkusOutput struct{ *pulumi.OutputState }

func (VendorSkusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VendorSkus)(nil))
}

func (o VendorSkusOutput) ToVendorSkusOutput() VendorSkusOutput {
	return o
}

func (o VendorSkusOutput) ToVendorSkusOutputWithContext(ctx context.Context) VendorSkusOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VendorSkusInput)(nil)).Elem(), &VendorSkus{})
	pulumi.RegisterOutputType(VendorSkusOutput{})
}
