


package v20201101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSyncMember(ctx *pulumi.Context, args *LookupSyncMemberArgs, opts ...pulumi.InvokeOption) (*LookupSyncMemberResult, error) {
	var rv LookupSyncMemberResult
	err := ctx.Invoke("azure-native:sql/v20201101preview:getSyncMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncMemberArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	SyncGroupName     string `pulumi:"syncGroupName"`
	SyncMemberName    string `pulumi:"syncMemberName"`
}


type LookupSyncMemberResult struct {
	DatabaseName                      *string `pulumi:"databaseName"`
	DatabaseType                      *string `pulumi:"databaseType"`
	Id                                string  `pulumi:"id"`
	Name                              string  `pulumi:"name"`
	PrivateEndpointName               string  `pulumi:"privateEndpointName"`
	ServerName                        *string `pulumi:"serverName"`
	SqlServerDatabaseId               *string `pulumi:"sqlServerDatabaseId"`
	SyncAgentId                       *string `pulumi:"syncAgentId"`
	SyncDirection                     *string `pulumi:"syncDirection"`
	SyncMemberAzureDatabaseResourceId *string `pulumi:"syncMemberAzureDatabaseResourceId"`
	SyncState                         string  `pulumi:"syncState"`
	Type                              string  `pulumi:"type"`
	UsePrivateLinkConnection          *bool   `pulumi:"usePrivateLinkConnection"`
	UserName                          *string `pulumi:"userName"`
}
