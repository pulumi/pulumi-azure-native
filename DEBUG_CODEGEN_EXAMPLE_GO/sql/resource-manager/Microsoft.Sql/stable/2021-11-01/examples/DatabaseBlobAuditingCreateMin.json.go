package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabaseBlobAuditingPolicy(ctx, "databaseBlobAuditingPolicy", &sql.DatabaseBlobAuditingPolicyArgs{
			BlobAuditingPolicyName:  pulumi.String("default"),
			DatabaseName:            pulumi.String("testdb"),
			ResourceGroupName:       pulumi.String("blobauditingtest-4799"),
			ServerName:              pulumi.String("blobauditingtest-6440"),
			State:                   sql.BlobAuditingPolicyStateEnabled,
			StorageAccountAccessKey: pulumi.String("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
			StorageEndpoint:         pulumi.String("https://mystorage.blob.core.windows.net"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
