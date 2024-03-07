package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewManagedInstance(ctx, "managedInstance", &sql.ManagedInstanceArgs{
			AdministratorLogin:         pulumi.String("dummylogin"),
			AdministratorLoginPassword: pulumi.String("PLACEHOLDER"),
			Administrators: &sql.ManagedInstanceExternalAdministratorArgs{
				AzureADOnlyAuthentication: pulumi.Bool(true),
				Login:                     pulumi.String("bob@contoso.com"),
				PrincipalType:             pulumi.String("User"),
				Sid:                       pulumi.String("00000011-1111-2222-2222-123456789111"),
				TenantId:                  pulumi.String("00000011-1111-2222-2222-123456789111"),
			},
			Collation:                        pulumi.String("SQL_Latin1_General_CP1_CI_AS"),
			DnsZonePartner:                   pulumi.String("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance"),
			InstancePoolId:                   pulumi.String("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/instancePools/pool1"),
			LicenseType:                      pulumi.String("LicenseIncluded"),
			Location:                         pulumi.String("Japan East"),
			MaintenanceConfigurationId:       pulumi.String("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1"),
			ManagedInstanceName:              pulumi.String("testinstance"),
			MinimalTlsVersion:                pulumi.String("1.2"),
			ProxyOverride:                    pulumi.String("Redirect"),
			PublicDataEndpointEnabled:        pulumi.Bool(false),
			RequestedBackupStorageRedundancy: pulumi.String("Geo"),
			ResourceGroupName:                pulumi.String("testrg"),
			ServicePrincipal: &sql.ServicePrincipalArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Sku: &sql.SkuArgs{
				Name: pulumi.String("GP_Gen5"),
				Tier: pulumi.String("GeneralPurpose"),
			},
			StorageSizeInGB: pulumi.Int(1024),
			SubnetId:        pulumi.String("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
			Tags: pulumi.StringMap{
				"tagKey1": pulumi.String("TagValue1"),
			},
			TimezoneId: pulumi.String("UTC"),
			VCores:     pulumi.Int(8),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
