package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azuredata/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azuredata.NewSqlServerRegistration(ctx, "sqlServerRegistration", &azuredata.SqlServerRegistrationArgs{
			Location:                  pulumi.String("northeurope"),
			ResourceGroupName:         pulumi.String("testrg"),
			SqlServerRegistrationName: pulumi.String("testsqlregistration"),
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
