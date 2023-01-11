


package hybridnetwork

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VendorSkuPreview struct {
	pulumi.CustomResourceState

	Name pulumi.StringOutput `pulumi:"name"`
	Type pulumi.StringOutput `pulumi:"type"`
}


func NewVendorSkuPreview(ctx *pulumi.Context,
	name string, args *VendorSkuPreviewArgs, opts ...pulumi.ResourceOption) (*VendorSkuPreview, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SkuName == nil {
		return nil, errors.New("invalid value for required argument 'SkuName'")
	}
	if args.VendorName == nil {
		return nil, errors.New("invalid value for required argument 'VendorName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20200101preview:VendorSkuPreview"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20210501:VendorSkuPreview"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20220101preview:VendorSkuPreview"),
		},
	})
	opts = append(opts, aliases)
	var resource VendorSkuPreview
	err := ctx.RegisterResource("azure-native:hybridnetwork:VendorSkuPreview", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVendorSkuPreview(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VendorSkuPreviewState, opts ...pulumi.ResourceOption) (*VendorSkuPreview, error) {
	var resource VendorSkuPreview
	err := ctx.ReadResource("azure-native:hybridnetwork:VendorSkuPreview", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vendorSkuPreviewState struct {
}

type VendorSkuPreviewState struct {
}

func (VendorSkuPreviewState) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkuPreviewState)(nil)).Elem()
}

type vendorSkuPreviewArgs struct {
	PreviewSubscription *string `pulumi:"previewSubscription"`
	SkuName             string  `pulumi:"skuName"`
	VendorName          string  `pulumi:"vendorName"`
}


type VendorSkuPreviewArgs struct {
	PreviewSubscription pulumi.StringPtrInput
	SkuName             pulumi.StringInput
	VendorName          pulumi.StringInput
}

func (VendorSkuPreviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkuPreviewArgs)(nil)).Elem()
}

type VendorSkuPreviewInput interface {
	pulumi.Input

	ToVendorSkuPreviewOutput() VendorSkuPreviewOutput
	ToVendorSkuPreviewOutputWithContext(ctx context.Context) VendorSkuPreviewOutput
}

func (*VendorSkuPreview) ElementType() reflect.Type {
	return reflect.TypeOf((**VendorSkuPreview)(nil)).Elem()
}

func (i *VendorSkuPreview) ToVendorSkuPreviewOutput() VendorSkuPreviewOutput {
	return i.ToVendorSkuPreviewOutputWithContext(context.Background())
}

func (i *VendorSkuPreview) ToVendorSkuPreviewOutputWithContext(ctx context.Context) VendorSkuPreviewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VendorSkuPreviewOutput)
}

type VendorSkuPreviewOutput struct{ *pulumi.OutputState }

func (VendorSkuPreviewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VendorSkuPreview)(nil)).Elem()
}

func (o VendorSkuPreviewOutput) ToVendorSkuPreviewOutput() VendorSkuPreviewOutput {
	return o
}

func (o VendorSkuPreviewOutput) ToVendorSkuPreviewOutputWithContext(ctx context.Context) VendorSkuPreviewOutput {
	return o
}

func (o VendorSkuPreviewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkuPreview) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VendorSkuPreviewOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkuPreview) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VendorSkuPreviewOutput{})
}
