


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExtendedServerBlobAuditingPolicy(ctx *pulumi.Context, args *LookupExtendedServerBlobAuditingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupExtendedServerBlobAuditingPolicyResult, error) {
	var rv LookupExtendedServerBlobAuditingPolicyResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getExtendedServerBlobAuditingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtendedServerBlobAuditingPolicyArgs struct {
	BlobAuditingPolicyName string `pulumi:"blobAuditingPolicyName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerName             string `pulumi:"serverName"`
}


type LookupExtendedServerBlobAuditingPolicyResult struct {
	AuditActionsAndGroups        []string `pulumi:"auditActionsAndGroups"`
	Id                           string   `pulumi:"id"`
	IsAzureMonitorTargetEnabled  *bool    `pulumi:"isAzureMonitorTargetEnabled"`
	IsDevopsAuditEnabled         *bool    `pulumi:"isDevopsAuditEnabled"`
	IsStorageSecondaryKeyInUse   *bool    `pulumi:"isStorageSecondaryKeyInUse"`
	Name                         string   `pulumi:"name"`
	PredicateExpression          *string  `pulumi:"predicateExpression"`
	QueueDelayMs                 *int     `pulumi:"queueDelayMs"`
	RetentionDays                *int     `pulumi:"retentionDays"`
	State                        string   `pulumi:"state"`
	StorageAccountSubscriptionId *string  `pulumi:"storageAccountSubscriptionId"`
	StorageEndpoint              *string  `pulumi:"storageEndpoint"`
	Type                         string   `pulumi:"type"`
}
