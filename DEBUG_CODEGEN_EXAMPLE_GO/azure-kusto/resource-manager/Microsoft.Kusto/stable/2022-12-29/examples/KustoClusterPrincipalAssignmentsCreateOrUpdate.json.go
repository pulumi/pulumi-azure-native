package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kusto/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kusto.NewClusterPrincipalAssignment(ctx, "clusterPrincipalAssignment", &kusto.ClusterPrincipalAssignmentArgs{
			ClusterName:             pulumi.String("kustoCluster"),
			PrincipalAssignmentName: pulumi.String("kustoprincipal1"),
			PrincipalId:             pulumi.String("87654321-1234-1234-1234-123456789123"),
			PrincipalType:           pulumi.String("App"),
			ResourceGroupName:       pulumi.String("kustorptest"),
			Role:                    pulumi.String("AllDatabasesAdmin"),
			TenantId:                pulumi.String("12345678-1234-1234-1234-123456789123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
