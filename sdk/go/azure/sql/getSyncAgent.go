


package sql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSyncAgent(ctx *pulumi.Context, args *LookupSyncAgentArgs, opts ...pulumi.InvokeOption) (*LookupSyncAgentResult, error) {
	var rv LookupSyncAgentResult
	err := ctx.Invoke("azure-native:sql:getSyncAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSyncAgentArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	SyncAgentName     string `pulumi:"syncAgentName"`
}


type LookupSyncAgentResult struct {
	ExpiryTime     string  `pulumi:"expiryTime"`
	Id             string  `pulumi:"id"`
	IsUpToDate     bool    `pulumi:"isUpToDate"`
	LastAliveTime  string  `pulumi:"lastAliveTime"`
	Name           string  `pulumi:"name"`
	State          string  `pulumi:"state"`
	SyncDatabaseId *string `pulumi:"syncDatabaseId"`
	Type           string  `pulumi:"type"`
	Version        string  `pulumi:"version"`
}
