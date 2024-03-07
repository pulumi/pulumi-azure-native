package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewDatabaseAccount(ctx, "databaseAccount", &documentdb.DatabaseAccountArgs{
			AccountName: pulumi.String("ddb1"),
			ApiProperties: &documentdb.ApiPropertiesArgs{
				ServerVersion: pulumi.String("3.2"),
			},
			BackupPolicy: documentdb.ContinuousModeBackupPolicy{
				ContinuousModeProperties: documentdb.ContinuousModeProperties{
					Tier: "Continuous30Days",
				},
				Type: "Continuous",
			},
			ConsistencyPolicy: &documentdb.ConsistencyPolicyArgs{
				DefaultConsistencyLevel: documentdb.DefaultConsistencyLevelBoundedStaleness,
				MaxIntervalInSeconds:    pulumi.Int(10),
				MaxStalenessPrefix:      pulumi.Float64(200),
			},
			CreateMode:               pulumi.String("Restore"),
			DatabaseAccountOfferType: documentdb.DatabaseAccountOfferTypeStandard,
			EnableAnalyticalStorage:  pulumi.Bool(true),
			EnableFreeTier:           pulumi.Bool(false),
			KeyVaultKeyUri:           pulumi.String("https://myKeyVault.vault.azure.net"),
			Kind:                     pulumi.String("GlobalDocumentDB"),
			Location:                 pulumi.String("westus"),
			Locations: []documentdb.LocationArgs{
				{
					FailoverPriority: pulumi.Int(0),
					IsZoneRedundant:  pulumi.Bool(false),
					LocationName:     pulumi.String("southcentralus"),
				},
			},
			MinimalTlsVersion: pulumi.String("Tls"),
			ResourceGroupName: pulumi.String("rg1"),
			RestoreParameters: documentdb.RestoreParametersResponse{
				DatabasesToRestore: documentdb.DatabaseRestoreResourceArray{
					&documentdb.DatabaseRestoreResourceArgs{
						CollectionNames: pulumi.StringArray{
							pulumi.String("collection1"),
							pulumi.String("collection2"),
						},
						DatabaseName: pulumi.String("db1"),
					},
					&documentdb.DatabaseRestoreResourceArgs{
						CollectionNames: pulumi.StringArray{
							pulumi.String("collection3"),
							pulumi.String("collection4"),
						},
						DatabaseName: pulumi.String("db2"),
					},
				},
				RestoreMode:           pulumi.String("PointInTime"),
				RestoreSource:         pulumi.String("/subscriptions/subid/providers/Microsoft.DocumentDB/locations/westus/restorableDatabaseAccounts/1a97b4bb-f6a0-430e-ade1-638d781830cc"),
				RestoreTimestampInUtc: pulumi.String("2021-03-11T22:05:09Z"),
			},
			Tags: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
