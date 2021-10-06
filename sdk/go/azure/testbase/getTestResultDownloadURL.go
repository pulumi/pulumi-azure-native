


package testbase

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTestResultDownloadURL(ctx *pulumi.Context, args *GetTestResultDownloadURLArgs, opts ...pulumi.InvokeOption) (*GetTestResultDownloadURLResult, error) {
	var rv GetTestResultDownloadURLResult
	err := ctx.Invoke("azure-native:testbase:getTestResultDownloadURL", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTestResultDownloadURLArgs struct {
	PackageName         string `pulumi:"packageName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
	TestResultName      string `pulumi:"testResultName"`
}


type GetTestResultDownloadURLResult struct {
	DownloadUrl    string `pulumi:"downloadUrl"`
	ExpirationTime string `pulumi:"expirationTime"`
}
