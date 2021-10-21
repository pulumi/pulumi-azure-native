


package v20190601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:netapp/v20190601:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	ActiveDirectories []ActiveDirectoryResponse `pulumi:"activeDirectories"`
	Id                string                    `pulumi:"id"`
	Location          string                    `pulumi:"location"`
	Name              string                    `pulumi:"name"`
	ProvisioningState string                    `pulumi:"provisioningState"`
	Tags              interface{}               `pulumi:"tags"`
	Type              string                    `pulumi:"type"`
}
