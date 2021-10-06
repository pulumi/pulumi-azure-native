


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobInventoryPolicy(ctx *pulumi.Context, args *LookupBlobInventoryPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBlobInventoryPolicyResult, error) {
	var rv LookupBlobInventoryPolicyResult
	err := ctx.Invoke("azure-native:storage/v20210201:getBlobInventoryPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobInventoryPolicyArgs struct {
	AccountName             string `pulumi:"accountName"`
	BlobInventoryPolicyName string `pulumi:"blobInventoryPolicyName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupBlobInventoryPolicyResult struct {
	Id               string                            `pulumi:"id"`
	LastModifiedTime string                            `pulumi:"lastModifiedTime"`
	Name             string                            `pulumi:"name"`
	Policy           BlobInventoryPolicySchemaResponse `pulumi:"policy"`
	SystemData       SystemDataResponse                `pulumi:"systemData"`
	Type             string                            `pulumi:"type"`
}
