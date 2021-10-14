


package v20180402

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerEndpoint(ctx *pulumi.Context, args *LookupServerEndpointArgs, opts ...pulumi.InvokeOption) (*LookupServerEndpointResult, error) {
	var rv LookupServerEndpointResult
	err := ctx.Invoke("azure-native:storagesync/v20180402:getServerEndpoint", args, &rv, opts...)
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
	CloudTiering           *string     `pulumi:"cloudTiering"`
	FriendlyName           *string     `pulumi:"friendlyName"`
	Id                     string      `pulumi:"id"`
	LastOperationName      *string     `pulumi:"lastOperationName"`
	LastWorkflowId         *string     `pulumi:"lastWorkflowId"`
	Name                   string      `pulumi:"name"`
	ProvisioningState      *string     `pulumi:"provisioningState"`
	ServerLocalPath        *string     `pulumi:"serverLocalPath"`
	ServerResourceId       *string     `pulumi:"serverResourceId"`
	SyncStatus             interface{} `pulumi:"syncStatus"`
	Type                   string      `pulumi:"type"`
	VolumeFreeSpacePercent *int        `pulumi:"volumeFreeSpacePercent"`
}
