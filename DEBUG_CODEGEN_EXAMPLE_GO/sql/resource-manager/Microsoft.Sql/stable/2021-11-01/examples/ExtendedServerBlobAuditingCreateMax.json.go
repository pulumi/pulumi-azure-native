package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewExtendedServerBlobAuditingPolicy(ctx, "extendedServerBlobAuditingPolicy", &sql.ExtendedServerBlobAuditingPolicyArgs{
			AuditActionsAndGroups: pulumi.StringArray{
				pulumi.String("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP"),
				pulumi.String("FAILED_DATABASE_AUTHENTICATION_GROUP"),
				pulumi.String("BATCH_COMPLETED_GROUP"),
			},
			BlobAuditingPolicyName:       pulumi.String("default"),
			IsAzureMonitorTargetEnabled:  pulumi.Bool(true),
			IsStorageSecondaryKeyInUse:   pulumi.Bool(false),
			PredicateExpression:          pulumi.String("object_name = 'SensitiveData'"),
			QueueDelayMs:                 pulumi.Int(4000),
			ResourceGroupName:            pulumi.String("blobauditingtest-4799"),
			RetentionDays:                pulumi.Int(6),
			ServerName:                   pulumi.String("blobauditingtest-6440"),
			State:                        sql.BlobAuditingPolicyStateEnabled,
			StorageAccountAccessKey:      pulumi.String("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
			StorageAccountSubscriptionId: pulumi.String("00000000-1234-0000-5678-000000000000"),
			StorageEndpoint:              pulumi.String("https://mystorage.blob.core.windows.net"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
