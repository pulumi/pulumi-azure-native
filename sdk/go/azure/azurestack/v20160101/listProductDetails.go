


package v20160101

import (
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
