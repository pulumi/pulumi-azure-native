package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewCluster(ctx, "cluster", &dbforpostgresql.ClusterArgs{
			AdministratorLoginPassword:      pulumi.String("password"),
			CitusVersion:                    pulumi.String("11.3"),
			ClusterName:                     pulumi.String("testcluster-burstablev2"),
			CoordinatorEnablePublicIpAccess: pulumi.Bool(true),
			CoordinatorServerEdition:        pulumi.String("BurstableGeneralPurpose"),
			CoordinatorStorageQuotaInMb:     pulumi.Int(131072),
			CoordinatorVCores:               pulumi.Int(2),
			EnableHa:                        pulumi.Bool(false),
			EnableShardsOnCoordinator:       pulumi.Bool(true),
			Location:                        pulumi.String("westus"),
			NodeCount:                       pulumi.Int(0),
			PostgresqlVersion:               pulumi.String("15"),
			PreferredPrimaryZone:            pulumi.String("1"),
			ResourceGroupName:               pulumi.String("TestGroup"),
			Tags: pulumi.StringMap{
				"owner": pulumi.String("JohnDoe"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
