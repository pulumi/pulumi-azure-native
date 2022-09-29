


package hybridnetwork

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
			Type: pulumi.String("azure-native:hybridnetwork/v20200101preview:VendorSkus"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20210501:VendorSkus"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20220101preview:VendorSkus"),
		},
	})
	opts = append(opts, aliases)
	var resource VendorSkus
	err := ctx.RegisterResource("azure-native:hybridnetwork:VendorSkus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVendorSkus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VendorSkusState, opts ...pulumi.ResourceOption) (*VendorSkus, error) {
	var resource VendorSkus
	err := ctx.ReadResource("azure-native:hybridnetwork:VendorSkus", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**VendorSkus)(nil)).Elem()
}

func (i *VendorSkus) ToVendorSkusOutput() VendorSkusOutput {
	return i.ToVendorSkusOutputWithContext(context.Background())
}

func (i *VendorSkus) ToVendorSkusOutputWithContext(ctx context.Context) VendorSkusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VendorSkusOutput)
}

type VendorSkusOutput struct{ *pulumi.OutputState }

func (VendorSkusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VendorSkus)(nil)).Elem()
}

func (o VendorSkusOutput) ToVendorSkusOutput() VendorSkusOutput {
	return o
}

func (o VendorSkusOutput) ToVendorSkusOutputWithContext(ctx context.Context) VendorSkusOutput {
	return o
}

func (o VendorSkusOutput) DeploymentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.StringPtrOutput { return v.DeploymentMode }).(pulumi.StringPtrOutput)
}

func (o VendorSkusOutput) ManagedApplicationParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.AnyOutput { return v.ManagedApplicationParameters }).(pulumi.AnyOutput)
}

func (o VendorSkusOutput) ManagedApplicationTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.AnyOutput { return v.ManagedApplicationTemplate }).(pulumi.AnyOutput)
}

func (o VendorSkusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VendorSkusOutput) NetworkFunctionTemplate() NetworkFunctionTemplateResponsePtrOutput {
	return o.ApplyT(func(v *VendorSkus) NetworkFunctionTemplateResponsePtrOutput { return v.NetworkFunctionTemplate }).(NetworkFunctionTemplateResponsePtrOutput)
}

func (o VendorSkusOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.BoolPtrOutput { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o VendorSkusOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VendorSkusOutput) SkuType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.StringPtrOutput { return v.SkuType }).(pulumi.StringPtrOutput)
}

func (o VendorSkusOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VendorSkusOutput{})
}
