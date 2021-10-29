


package v20180701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobServiceProperties(ctx *pulumi.Context, args *LookupBlobServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupBlobServicePropertiesResult, error) {
	var rv LookupBlobServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20180701:getBlobServiceProperties", args, &rv, opts...)
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
	Cors                  *CorsRulesResponse             `pulumi:"cors"`
	DefaultServiceVersion *string                        `pulumi:"defaultServiceVersion"`
	DeleteRetentionPolicy *DeleteRetentionPolicyResponse `pulumi:"deleteRetentionPolicy"`
	Id                    string                         `pulumi:"id"`
	Name                  string                         `pulumi:"name"`
	Type                  string                         `pulumi:"type"`
}
