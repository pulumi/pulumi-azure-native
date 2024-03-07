package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewServer(ctx, "server", &dbforpostgresql.ServerArgs{
			AdministratorLogin:         pulumi.String("cloudsa"),
			AdministratorLoginPassword: pulumi.String("password"),
			AvailabilityZone:           pulumi.String("1"),
			Backup: &dbforpostgresql.BackupArgs{
				BackupRetentionDays: pulumi.Int(7),
				GeoRedundantBackup:  pulumi.String("Disabled"),
			},
			CreateMode: pulumi.String("Create"),
			DataEncryption: &dbforpostgresql.DataEncryptionArgs{
				PrimaryKeyURI:                 pulumi.String("https://test-kv.vault.azure.net/keys/test-key1/77f57315bab34b0189daa113fbc78787"),
				PrimaryUserAssignedIdentityId: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity"),
				Type:                          pulumi.String("AzureKeyVault"),
			},
			HighAvailability: &dbforpostgresql.HighAvailabilityArgs{
				Mode: pulumi.String("ZoneRedundant"),
			},
			Identity: &dbforpostgresql.UserAssignedIdentityArgs{
				Type: pulumi.String("UserAssigned"),
				UserAssignedIdentities: dbforpostgresql.UserIdentityMap{
					"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testresourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-usermanagedidentity": nil,
				},
			},
			Location: pulumi.String("westus"),
			Network: &dbforpostgresql.NetworkArgs{
				DelegatedSubnetResourceId:   pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-vnet-subnet"),
				PrivateDnsZoneArmResourceId: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/testrg/providers/Microsoft.Network/privateDnsZones/test-private-dns-zone.postgres.database.azure.com"),
			},
			ResourceGroupName: pulumi.String("testrg"),
			ServerName:        pulumi.String("pgtestsvc4"),
			Sku: &dbforpostgresql.SkuArgs{
				Name: pulumi.String("Standard_D4s_v3"),
				Tier: pulumi.String("GeneralPurpose"),
			},
			Storage: &dbforpostgresql.StorageArgs{
				StorageSizeGB: pulumi.Int(512),
			},
			Tags: pulumi.StringMap{
				"ElasticServer": pulumi.String("1"),
			},
			Version: pulumi.String("12"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
