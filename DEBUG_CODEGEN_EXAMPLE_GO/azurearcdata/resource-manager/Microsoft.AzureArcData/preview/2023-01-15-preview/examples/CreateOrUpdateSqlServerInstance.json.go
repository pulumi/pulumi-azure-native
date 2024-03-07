package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurearcdata/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurearcdata.NewSqlServerInstance(ctx, "sqlServerInstance", &azurearcdata.SqlServerInstanceArgs{
			Location: pulumi.String("northeurope"),
			Properties: &azurearcdata.SqlServerInstancePropertiesArgs{
				AzureDefenderStatus:            pulumi.String("Protected"),
				AzureDefenderStatusLastUpdated: pulumi.String("2020-01-02T17:18:19.1234567Z"),
				Collation:                      pulumi.String("collation"),
				ContainerResourceId:            pulumi.String("Resource id of hosting Arc Machine"),
				Cores:                          pulumi.String("4"),
				CurrentVersion:                 pulumi.String("2012"),
				Edition:                        pulumi.String("Developer"),
				HostType:                       pulumi.String("Physical Server"),
				InstanceName:                   pulumi.String("name of instance"),
				LicenseType:                    pulumi.String("Free"),
				PatchLevel:                     pulumi.String("patchlevel"),
				ProductId:                      pulumi.String("sql id"),
				Status:                         pulumi.String("Registered"),
				TcpDynamicPorts:                pulumi.String("1433"),
				TcpStaticPorts:                 pulumi.String("1433"),
				VCore:                          pulumi.String("4"),
				Version:                        pulumi.String("SQL Server 2012"),
			},
			ResourceGroupName:     pulumi.String("testrg"),
			SqlServerInstanceName: pulumi.String("testsqlServerInstance"),
			Tags: pulumi.StringMap{
				"mytag": pulumi.String("myval"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
