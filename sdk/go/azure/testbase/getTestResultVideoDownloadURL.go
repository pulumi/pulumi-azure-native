


package testbase

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTestResultVideoDownloadURL(ctx *pulumi.Context, args *GetTestResultVideoDownloadURLArgs, opts ...pulumi.InvokeOption) (*GetTestResultVideoDownloadURLResult, error) {
	var rv GetTestResultVideoDownloadURLResult
	err := ctx.Invoke("azure-native:testbase:getTestResultVideoDownloadURL", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTestResultVideoDownloadURLArgs struct {
	PackageName         string `pulumi:"packageName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
	TestResultName      string `pulumi:"testResultName"`
}


type GetTestResultVideoDownloadURLResult struct {
	DownloadUrl    string `pulumi:"downloadUrl"`
	ExpirationTime string `pulumi:"expirationTime"`
}
