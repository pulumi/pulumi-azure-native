


package v20200202preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSyncGroup(ctx *pulumi.Context, args *LookupSyncGroupArgs, opts ...pulumi.InvokeOption) (*LookupSyncGroupResult, error) {
	var rv LookupSyncGroupResult
	err := ctx.Invoke("azure-native:sql/v20200202preview:getSyncGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncGroupArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	SyncGroupName     string `pulumi:"syncGroupName"`
}


type LookupSyncGroupResult struct {
	ConflictLoggingRetentionInDays *int                     `pulumi:"conflictLoggingRetentionInDays"`
	ConflictResolutionPolicy       *string                  `pulumi:"conflictResolutionPolicy"`
	EnableConflictLogging          *bool                    `pulumi:"enableConflictLogging"`
	HubDatabaseUserName            *string                  `pulumi:"hubDatabaseUserName"`
	Id                             string                   `pulumi:"id"`
	Interval                       *int                     `pulumi:"interval"`
	LastSyncTime                   string                   `pulumi:"lastSyncTime"`
	Name                           string                   `pulumi:"name"`
	PrivateEndpointName            string                   `pulumi:"privateEndpointName"`
	Schema                         *SyncGroupSchemaResponse `pulumi:"schema"`
	Sku                            *SkuResponse             `pulumi:"sku"`
	SyncDatabaseId                 *string                  `pulumi:"syncDatabaseId"`
	SyncState                      string                   `pulumi:"syncState"`
	Type                           string                   `pulumi:"type"`
	UsePrivateLinkConnection       *bool                    `pulumi:"usePrivateLinkConnection"`
}
