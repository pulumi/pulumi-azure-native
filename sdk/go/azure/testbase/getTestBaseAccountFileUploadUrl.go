


package testbase

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTestBaseAccountFileUploadUrl(ctx *pulumi.Context, args *GetTestBaseAccountFileUploadUrlArgs, opts ...pulumi.InvokeOption) (*GetTestBaseAccountFileUploadUrlResult, error) {
	var rv GetTestBaseAccountFileUploadUrlResult
	err := ctx.Invoke("azure-native:testbase:getTestBaseAccountFileUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTestBaseAccountFileUploadUrlArgs struct {
	BlobName            *string `pulumi:"blobName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	TestBaseAccountName string  `pulumi:"testBaseAccountName"`
}


type GetTestBaseAccountFileUploadUrlResult struct {
	BlobPath  string `pulumi:"blobPath"`
	UploadUrl string `pulumi:"uploadUrl"`
}
