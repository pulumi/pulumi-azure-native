


package customerinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetImageUploadUrlForEntityType(ctx *pulumi.Context, args *GetImageUploadUrlForEntityTypeArgs, opts ...pulumi.InvokeOption) (*GetImageUploadUrlForEntityTypeResult, error) {
	var rv GetImageUploadUrlForEntityTypeResult
	err := ctx.Invoke("azure-native:customerinsights:getImageUploadUrlForEntityType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetImageUploadUrlForEntityTypeArgs struct {
	EntityType        *string `pulumi:"entityType"`
	EntityTypeName    *string `pulumi:"entityTypeName"`
	HubName           string  `pulumi:"hubName"`
	RelativePath      *string `pulumi:"relativePath"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type GetImageUploadUrlForEntityTypeResult struct {
	ContentUrl   *string `pulumi:"contentUrl"`
	ImageExists  *bool   `pulumi:"imageExists"`
	RelativePath *string `pulumi:"relativePath"`
}
