


package v20160101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListProductDetails(ctx *pulumi.Context, args *ListProductDetailsArgs, opts ...pulumi.InvokeOption) (*ListProductDetailsResult, error) {
	var rv ListProductDetailsResult
	err := ctx.Invoke("azure-native:azurestack/v20160101:listProductDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProductDetailsArgs struct {
	ProductName      string `pulumi:"productName"`
	RegistrationName string `pulumi:"registrationName"`
	ResourceGroup    string `pulumi:"resourceGroup"`
}


type ListProductDetailsResult struct {
	ComputeRole               string                  `pulumi:"computeRole"`
	DataDiskImages            []DataDiskImageResponse `pulumi:"dataDiskImages"`
	GalleryPackageBlobSasUri  string                  `pulumi:"galleryPackageBlobSasUri"`
	IsSystemExtension         bool                    `pulumi:"isSystemExtension"`
	OsDiskImage               OsDiskImageResponse     `pulumi:"osDiskImage"`
	ProductKind               string                  `pulumi:"productKind"`
	SupportMultipleExtensions bool                    `pulumi:"supportMultipleExtensions"`
	Uri                       string                  `pulumi:"uri"`
	Version                   string                  `pulumi:"version"`
	VmOsType                  string                  `pulumi:"vmOsType"`
	VmScaleSetEnabled         bool                    `pulumi:"vmScaleSetEnabled"`
}

func ListProductDetailsOutput(ctx *pulumi.Context, args ListProductDetailsOutputArgs, opts ...pulumi.InvokeOption) ListProductDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListProductDetailsResult, error) {
			args := v.(ListProductDetailsArgs)
			r, err := ListProductDetails(ctx, &args, opts...)
			var s ListProductDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListProductDetailsResultOutput)
}

type ListProductDetailsOutputArgs struct {
	ProductName      pulumi.StringInput `pulumi:"productName"`
	RegistrationName pulumi.StringInput `pulumi:"registrationName"`
	ResourceGroup    pulumi.StringInput `pulumi:"resourceGroup"`
}

func (ListProductDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductDetailsArgs)(nil)).Elem()
}


type ListProductDetailsResultOutput struct{ *pulumi.OutputState }

func (ListProductDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductDetailsResult)(nil)).Elem()
}

func (o ListProductDetailsResultOutput) ToListProductDetailsResultOutput() ListProductDetailsResultOutput {
	return o
}

func (o ListProductDetailsResultOutput) ToListProductDetailsResultOutputWithContext(ctx context.Context) ListProductDetailsResultOutput {
	return o
}

func (o ListProductDetailsResultOutput) ComputeRole() pulumi.StringOutput {
	return o.ApplyT(func(v ListProductDetailsResult) string { return v.ComputeRole }).(pulumi.StringOutput)
}

func (o ListProductDetailsResultOutput) DataDiskImages() DataDiskImageResponseArrayOutput {
	return o.ApplyT(func(v ListProductDetailsResult) []DataDiskImageResponse { return v.DataDiskImages }).(DataDiskImageResponseArrayOutput)
}

func (o ListProductDetailsResultOutput) GalleryPackageBlobSasUri() pulumi.StringOutput {
	return o.ApplyT(func(v ListProductDetailsResult) string { return v.GalleryPackageBlobSasUri }).(pulumi.StringOutput)
}

func (o ListProductDetailsResultOutput) IsSystemExtension() pulumi.BoolOutput {
	return o.ApplyT(func(v ListProductDetailsResult) bool { return v.IsSystemExtension }).(pulumi.BoolOutput)
}

func (o ListProductDetailsResultOutput) OsDiskImage() OsDiskImageResponseOutput {
	return o.ApplyT(func(v ListProductDetailsResult) OsDiskImageResponse { return v.OsDiskImage }).(OsDiskImageResponseOutput)
}

func (o ListProductDetailsResultOutput) ProductKind() pulumi.StringOutput {
	return o.ApplyT(func(v ListProductDetailsResult) string { return v.ProductKind }).(pulumi.StringOutput)
}

func (o ListProductDetailsResultOutput) SupportMultipleExtensions() pulumi.BoolOutput {
	return o.ApplyT(func(v ListProductDetailsResult) bool { return v.SupportMultipleExtensions }).(pulumi.BoolOutput)
}

func (o ListProductDetailsResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ListProductDetailsResult) string { return v.Uri }).(pulumi.StringOutput)
}

func (o ListProductDetailsResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ListProductDetailsResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o ListProductDetailsResultOutput) VmOsType() pulumi.StringOutput {
	return o.ApplyT(func(v ListProductDetailsResult) string { return v.VmOsType }).(pulumi.StringOutput)
}

func (o ListProductDetailsResultOutput) VmScaleSetEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ListProductDetailsResult) bool { return v.VmScaleSetEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProductDetailsResultOutput{})
}
