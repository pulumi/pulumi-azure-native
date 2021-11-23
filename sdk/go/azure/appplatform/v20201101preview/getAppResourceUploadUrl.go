


package v20201101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAppResourceUploadUrl(ctx *pulumi.Context, args *GetAppResourceUploadUrlArgs, opts ...pulumi.InvokeOption) (*GetAppResourceUploadUrlResult, error) {
	var rv GetAppResourceUploadUrlResult
	err := ctx.Invoke("azure-native:appplatform/v20201101preview:getAppResourceUploadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAppResourceUploadUrlArgs struct {
	AppName           string `pulumi:"appName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type GetAppResourceUploadUrlResult struct {
	RelativePath *string `pulumi:"relativePath"`
	UploadUrl    *string `pulumi:"uploadUrl"`
}
