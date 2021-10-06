


package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationPackage(ctx *pulumi.Context, args *LookupApplicationPackageArgs, opts ...pulumi.InvokeOption) (*LookupApplicationPackageResult, error) {
	var rv LookupApplicationPackageResult
	err := ctx.Invoke("azure-native:batch/v20200501:getApplicationPackage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationPackageArgs struct {
	AccountName       string `pulumi:"accountName"`
	ApplicationName   string `pulumi:"applicationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VersionName       string `pulumi:"versionName"`
}


type LookupApplicationPackageResult struct {
	Etag               string `pulumi:"etag"`
	Format             string `pulumi:"format"`
	Id                 string `pulumi:"id"`
	LastActivationTime string `pulumi:"lastActivationTime"`
	Name               string `pulumi:"name"`
	State              string `pulumi:"state"`
	StorageUrl         string `pulumi:"storageUrl"`
	StorageUrlExpiry   string `pulumi:"storageUrlExpiry"`
	Type               string `pulumi:"type"`
}
