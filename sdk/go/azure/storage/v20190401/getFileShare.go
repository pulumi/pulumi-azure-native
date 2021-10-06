


package v20190401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFileShare(ctx *pulumi.Context, args *LookupFileShareArgs, opts ...pulumi.InvokeOption) (*LookupFileShareResult, error) {
	var rv LookupFileShareResult
	err := ctx.Invoke("azure-native:storage/v20190401:getFileShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileShareArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupFileShareResult struct {
	Etag             string            `pulumi:"etag"`
	Id               string            `pulumi:"id"`
	LastModifiedTime string            `pulumi:"lastModifiedTime"`
	Metadata         map[string]string `pulumi:"metadata"`
	Name             string            `pulumi:"name"`
	ShareQuota       *int              `pulumi:"shareQuota"`
	Type             string            `pulumi:"type"`
}
