package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sqlvirtualmachine/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sqlvirtualmachine.NewSqlVirtualMachineGroup(ctx, "sqlVirtualMachineGroup", &sqlvirtualmachine.SqlVirtualMachineGroupArgs{
			Location:                   pulumi.String("northeurope"),
			ResourceGroupName:          pulumi.String("testrg"),
			SqlImageOffer:              pulumi.String("SQL2016-WS2016"),
			SqlImageSku:                pulumi.String("Enterprise"),
			SqlVirtualMachineGroupName: pulumi.String("testvmgroup"),
			Tags: pulumi.StringMap{
				"mytag": pulumi.String("myval"),
			},
			WsfcDomainProfile: &sqlvirtualmachine.WsfcDomainProfileArgs{
				ClusterBootstrapAccount:  pulumi.String("testrpadmin"),
				ClusterOperatorAccount:   pulumi.String("testrp@testdomain.com"),
				ClusterSubnetType:        pulumi.String("MultiSubnet"),
				DomainFqdn:               pulumi.String("testdomain.com"),
				OuPath:                   pulumi.String("OU=WSCluster,DC=testdomain,DC=com"),
				SqlServiceAccount:        pulumi.String("sqlservice@testdomain.com"),
				StorageAccountPrimaryKey: pulumi.String("<primary storage access key>"),
				StorageAccountUrl:        pulumi.String("https://storgact.blob.core.windows.net/"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
