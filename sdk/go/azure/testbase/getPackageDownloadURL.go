


package testbase

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPackageDownloadURL(ctx *pulumi.Context, args *GetPackageDownloadURLArgs, opts ...pulumi.InvokeOption) (*GetPackageDownloadURLResult, error) {
	var rv GetPackageDownloadURLResult
	err := ctx.Invoke("azure-native:testbase:getPackageDownloadURL", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetPackageDownloadURLArgs struct {
	PackageName         string `pulumi:"packageName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}


type GetPackageDownloadURLResult struct {
	DownloadUrl    string `pulumi:"downloadUrl"`
	ExpirationTime string `pulumi:"expirationTime"`
}
