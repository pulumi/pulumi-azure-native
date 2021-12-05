


package v20150521preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLabResource(ctx *pulumi.Context, args *LookupLabResourceArgs, opts ...pulumi.InvokeOption) (*LookupLabResourceResult, error) {
	var rv LookupLabResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getLabResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabResourceArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLabResourceResult struct {
	ArtifactsStorageAccount *string           `pulumi:"artifactsStorageAccount"`
	CreatedDate             *string           `pulumi:"createdDate"`
	DefaultStorageAccount   *string           `pulumi:"defaultStorageAccount"`
	DefaultVirtualNetworkId *string           `pulumi:"defaultVirtualNetworkId"`
	Id                      *string           `pulumi:"id"`
	LabStorageType          *string           `pulumi:"labStorageType"`
	Location                *string           `pulumi:"location"`
	Name                    *string           `pulumi:"name"`
	ProvisioningState       *string           `pulumi:"provisioningState"`
	StorageAccounts         []string          `pulumi:"storageAccounts"`
	Tags                    map[string]string `pulumi:"tags"`
	Type                    *string           `pulumi:"type"`
	VaultName               *string           `pulumi:"vaultName"`
}
