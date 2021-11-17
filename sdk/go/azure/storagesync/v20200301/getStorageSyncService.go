


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageSyncService(ctx *pulumi.Context, args *LookupStorageSyncServiceArgs, opts ...pulumi.InvokeOption) (*LookupStorageSyncServiceResult, error) {
	var rv LookupStorageSyncServiceResult
	err := ctx.Invoke("azure-native:storagesync/v20200301:getStorageSyncService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageSyncServiceArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
}


type LookupStorageSyncServiceResult struct {
	Id                         string                              `pulumi:"id"`
	IncomingTrafficPolicy      *string                             `pulumi:"incomingTrafficPolicy"`
	LastOperationName          string                              `pulumi:"lastOperationName"`
	LastWorkflowId             string                              `pulumi:"lastWorkflowId"`
	Location                   string                              `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	StorageSyncServiceStatus   int                                 `pulumi:"storageSyncServiceStatus"`
	StorageSyncServiceUid      string                              `pulumi:"storageSyncServiceUid"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
}
