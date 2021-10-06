


package v20210401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobServiceProperties(ctx *pulumi.Context, args *LookupBlobServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupBlobServicePropertiesResult, error) {
	var rv LookupBlobServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20210401:getBlobServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobServicePropertiesArgs struct {
	AccountName       string `pulumi:"accountName"`
	BlobServicesName  string `pulumi:"blobServicesName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBlobServicePropertiesResult struct {
	AutomaticSnapshotPolicyEnabled *bool                                 `pulumi:"automaticSnapshotPolicyEnabled"`
	ChangeFeed                     *ChangeFeedResponse                   `pulumi:"changeFeed"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicyResponse        `pulumi:"containerDeleteRetentionPolicy"`
	Cors                           *CorsRulesResponse                    `pulumi:"cors"`
	DefaultServiceVersion          *string                               `pulumi:"defaultServiceVersion"`
	DeleteRetentionPolicy          *DeleteRetentionPolicyResponse        `pulumi:"deleteRetentionPolicy"`
	Id                             string                                `pulumi:"id"`
	IsVersioningEnabled            *bool                                 `pulumi:"isVersioningEnabled"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicyResponse `pulumi:"lastAccessTimeTrackingPolicy"`
	Name                           string                                `pulumi:"name"`
	RestorePolicy                  *RestorePolicyPropertiesResponse      `pulumi:"restorePolicy"`
	Sku                            SkuResponse                           `pulumi:"sku"`
	Type                           string                                `pulumi:"type"`
}
