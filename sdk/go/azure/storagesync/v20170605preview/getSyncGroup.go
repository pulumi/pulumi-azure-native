


package v20170605preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSyncGroup(ctx *pulumi.Context, args *LookupSyncGroupArgs, opts ...pulumi.InvokeOption) (*LookupSyncGroupResult, error) {
	var rv LookupSyncGroupResult
	err := ctx.Invoke("azure-native:storagesync/v20170605preview:getSyncGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncGroupArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	SyncGroupName          string `pulumi:"syncGroupName"`
}


type LookupSyncGroupResult struct {
	Id              string  `pulumi:"id"`
	Name            string  `pulumi:"name"`
	SyncGroupStatus string  `pulumi:"syncGroupStatus"`
	Type            string  `pulumi:"type"`
	UniqueId        *string `pulumi:"uniqueId"`
}
