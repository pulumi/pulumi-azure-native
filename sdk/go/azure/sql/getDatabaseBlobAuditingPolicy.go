


package sql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseBlobAuditingPolicy(ctx *pulumi.Context, args *LookupDatabaseBlobAuditingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseBlobAuditingPolicyResult, error) {
	var rv LookupDatabaseBlobAuditingPolicyResult
	err := ctx.Invoke("azure-native:sql:getDatabaseBlobAuditingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseBlobAuditingPolicyArgs struct {
	BlobAuditingPolicyName string `pulumi:"blobAuditingPolicyName"`
	DatabaseName           string `pulumi:"databaseName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerName             string `pulumi:"serverName"`
}


type LookupDatabaseBlobAuditingPolicyResult struct {
	AuditActionsAndGroups        []string `pulumi:"auditActionsAndGroups"`
	Id                           string   `pulumi:"id"`
	IsAzureMonitorTargetEnabled  *bool    `pulumi:"isAzureMonitorTargetEnabled"`
	IsStorageSecondaryKeyInUse   *bool    `pulumi:"isStorageSecondaryKeyInUse"`
	Kind                         string   `pulumi:"kind"`
	Name                         string   `pulumi:"name"`
	QueueDelayMs                 *int     `pulumi:"queueDelayMs"`
	RetentionDays                *int     `pulumi:"retentionDays"`
	State                        string   `pulumi:"state"`
	StorageAccountSubscriptionId *string  `pulumi:"storageAccountSubscriptionId"`
	StorageEndpoint              *string  `pulumi:"storageEndpoint"`
	Type                         string   `pulumi:"type"`
}
