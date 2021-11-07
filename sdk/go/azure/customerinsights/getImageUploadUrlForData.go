


package customerinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetImageUploadUrlForData(ctx *pulumi.Context, args *GetImageUploadUrlForDataArgs, opts ...pulumi.InvokeOption) (*GetImageUploadUrlForDataResult, error) {
	var rv GetImageUploadUrlForDataResult
	err := ctx.Invoke("azure-native:customerinsights:getImageUploadUrlForData", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetImageUploadUrlForDataArgs struct {
	EntityType        *string `pulumi:"entityType"`
	EntityTypeName    *string `pulumi:"entityTypeName"`
	HubName           string  `pulumi:"hubName"`
	RelativePath      *string `pulumi:"relativePath"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type GetImageUploadUrlForDataResult struct {
	ContentUrl   *string `pulumi:"contentUrl"`
	ImageExists  *bool   `pulumi:"imageExists"`
	RelativePath *string `pulumi:"relativePath"`
}
