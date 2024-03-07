package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azuredata/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azuredata.NewSqlServer(ctx, "sqlServer", &azuredata.SqlServerArgs{
			Cores:                     pulumi.Int(8),
			Edition:                   pulumi.String("Latin"),
			PropertyBag:               pulumi.String(""),
			RegistrationID:            pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureData/SqlServerRegistrations/testsqlregistration"),
			ResourceGroupName:         pulumi.String("testrg"),
			SqlServerName:             pulumi.String("testsqlserver"),
			SqlServerRegistrationName: pulumi.String("testsqlregistration"),
			Version:                   pulumi.String("2008"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
