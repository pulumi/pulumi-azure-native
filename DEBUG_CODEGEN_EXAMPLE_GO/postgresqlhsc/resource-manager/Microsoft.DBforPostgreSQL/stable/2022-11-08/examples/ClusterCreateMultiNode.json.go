package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewCluster(ctx, "cluster", &dbforpostgresql.ClusterArgs{
			AdministratorLoginPassword:      pulumi.String("password"),
			CitusVersion:                    pulumi.String("11.1"),
			ClusterName:                     pulumi.String("testcluster-multinode"),
			CoordinatorEnablePublicIpAccess: pulumi.Bool(true),
			CoordinatorServerEdition:        pulumi.String("GeneralPurpose"),
			CoordinatorStorageQuotaInMb:     pulumi.Int(524288),
			CoordinatorVCores:               pulumi.Int(4),
			EnableHa:                        pulumi.Bool(true),
			EnableShardsOnCoordinator:       pulumi.Bool(false),
			Location:                        pulumi.String("westus"),
			NodeCount:                       pulumi.Int(3),
			NodeEnablePublicIpAccess:        pulumi.Bool(false),
			NodeServerEdition:               pulumi.String("MemoryOptimized"),
			NodeStorageQuotaInMb:            pulumi.Int(524288),
			NodeVCores:                      pulumi.Int(8),
			PostgresqlVersion:               pulumi.String("15"),
			PreferredPrimaryZone:            pulumi.String("1"),
			ResourceGroupName:               pulumi.String("TestGroup"),
			Tags:                            nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
