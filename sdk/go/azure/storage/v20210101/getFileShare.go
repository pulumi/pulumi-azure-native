


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFileShare(ctx *pulumi.Context, args *LookupFileShareArgs, opts ...pulumi.InvokeOption) (*LookupFileShareResult, error) {
	var rv LookupFileShareResult
	err := ctx.Invoke("azure-native:storage/v20210101:getFileShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileShareArgs struct {
	AccountName       string  `pulumi:"accountName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ShareName         string  `pulumi:"shareName"`
}


type LookupFileShareResult struct {
	AccessTier             *string           `pulumi:"accessTier"`
	AccessTierChangeTime   string            `pulumi:"accessTierChangeTime"`
	AccessTierStatus       string            `pulumi:"accessTierStatus"`
	Deleted                bool              `pulumi:"deleted"`
	DeletedTime            string            `pulumi:"deletedTime"`
	EnabledProtocols       *string           `pulumi:"enabledProtocols"`
	Etag                   string            `pulumi:"etag"`
	Id                     string            `pulumi:"id"`
	LastModifiedTime       string            `pulumi:"lastModifiedTime"`
	Metadata               map[string]string `pulumi:"metadata"`
	Name                   string            `pulumi:"name"`
	RemainingRetentionDays int               `pulumi:"remainingRetentionDays"`
	RootSquash             *string           `pulumi:"rootSquash"`
	ShareQuota             *int              `pulumi:"shareQuota"`
	ShareUsageBytes        float64           `pulumi:"shareUsageBytes"`
	SnapshotTime           string            `pulumi:"snapshotTime"`
	Type                   string            `pulumi:"type"`
	Version                string            `pulumi:"version"`
}
