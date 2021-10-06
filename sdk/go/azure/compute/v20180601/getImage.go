


package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	var rv LookupImageResult
	err := ctx.Invoke("azure-native:compute/v20180601:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImageArgs struct {
	Expand            *string `pulumi:"expand"`
	ImageName         string  `pulumi:"imageName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupImageResult struct {
	Id                   string                       `pulumi:"id"`
	Location             string                       `pulumi:"location"`
	Name                 string                       `pulumi:"name"`
	ProvisioningState    string                       `pulumi:"provisioningState"`
	SourceVirtualMachine *SubResourceResponse         `pulumi:"sourceVirtualMachine"`
	StorageProfile       *ImageStorageProfileResponse `pulumi:"storageProfile"`
	Tags                 map[string]string            `pulumi:"tags"`
	Type                 string                       `pulumi:"type"`
}
