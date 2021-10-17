


package v20170501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:machinelearningexperimentation/v20170501preview:getAccount", args, &rv, opts...)
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
	AccountId         string                           `pulumi:"accountId"`
	CreationDate      string                           `pulumi:"creationDate"`
	Description       *string                          `pulumi:"description"`
	DiscoveryUri      string                           `pulumi:"discoveryUri"`
	FriendlyName      *string                          `pulumi:"friendlyName"`
	Id                string                           `pulumi:"id"`
	KeyVaultId        string                           `pulumi:"keyVaultId"`
	Location          string                           `pulumi:"location"`
	Name              string                           `pulumi:"name"`
	ProvisioningState string                           `pulumi:"provisioningState"`
	Seats             *string                          `pulumi:"seats"`
	StorageAccount    StorageAccountPropertiesResponse `pulumi:"storageAccount"`
	Tags              map[string]string                `pulumi:"tags"`
	Type              string                           `pulumi:"type"`
	VsoAccountId      string                           `pulumi:"vsoAccountId"`
}
