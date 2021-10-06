


package v20191001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerEndpoint(ctx *pulumi.Context, args *LookupServerEndpointArgs, opts ...pulumi.InvokeOption) (*LookupServerEndpointResult, error) {
	var rv LookupServerEndpointResult
	err := ctx.Invoke("azure-native:storagesync/v20191001:getServerEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerEndpointArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerEndpointName     string `pulumi:"serverEndpointName"`
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	SyncGroupName          string `pulumi:"syncGroupName"`
}


type LookupServerEndpointResult struct {
	CloudTiering                                *string                                  `pulumi:"cloudTiering"`
	CloudTieringStatus                          ServerEndpointCloudTieringStatusResponse `pulumi:"cloudTieringStatus"`
	FriendlyName                                *string                                  `pulumi:"friendlyName"`
	Id                                          string                                   `pulumi:"id"`
	LastOperationName                           string                                   `pulumi:"lastOperationName"`
	LastWorkflowId                              string                                   `pulumi:"lastWorkflowId"`
	Name                                        string                                   `pulumi:"name"`
	OfflineDataTransfer                         *string                                  `pulumi:"offlineDataTransfer"`
	OfflineDataTransferShareName                *string                                  `pulumi:"offlineDataTransferShareName"`
	OfflineDataTransferStorageAccountResourceId string                                   `pulumi:"offlineDataTransferStorageAccountResourceId"`
	OfflineDataTransferStorageAccountTenantId   string                                   `pulumi:"offlineDataTransferStorageAccountTenantId"`
	ProvisioningState                           string                                   `pulumi:"provisioningState"`
	RecallStatus                                ServerEndpointRecallStatusResponse       `pulumi:"recallStatus"`
	ServerLocalPath                             *string                                  `pulumi:"serverLocalPath"`
	ServerResourceId                            *string                                  `pulumi:"serverResourceId"`
	SyncStatus                                  ServerEndpointSyncStatusResponse         `pulumi:"syncStatus"`
	TierFilesOlderThanDays                      *int                                     `pulumi:"tierFilesOlderThanDays"`
	Type                                        string                                   `pulumi:"type"`
	VolumeFreeSpacePercent                      *int                                     `pulumi:"volumeFreeSpacePercent"`
}
