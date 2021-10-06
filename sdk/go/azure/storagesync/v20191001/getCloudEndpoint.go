


package v20191001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudEndpoint(ctx *pulumi.Context, args *LookupCloudEndpointArgs, opts ...pulumi.InvokeOption) (*LookupCloudEndpointResult, error) {
	var rv LookupCloudEndpointResult
	err := ctx.Invoke("azure-native:storagesync/v20191001:getCloudEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudEndpointArgs struct {
	CloudEndpointName      string `pulumi:"cloudEndpointName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	SyncGroupName          string `pulumi:"syncGroupName"`
}


type LookupCloudEndpointResult struct {
	AzureFileShareName       *string `pulumi:"azureFileShareName"`
	BackupEnabled            string  `pulumi:"backupEnabled"`
	FriendlyName             *string `pulumi:"friendlyName"`
	Id                       string  `pulumi:"id"`
	LastOperationName        *string `pulumi:"lastOperationName"`
	LastWorkflowId           *string `pulumi:"lastWorkflowId"`
	Name                     string  `pulumi:"name"`
	PartnershipId            *string `pulumi:"partnershipId"`
	ProvisioningState        *string `pulumi:"provisioningState"`
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
	StorageAccountTenantId   *string `pulumi:"storageAccountTenantId"`
	Type                     string  `pulumi:"type"`
}
