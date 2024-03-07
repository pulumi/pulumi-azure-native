package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabaseSecurityAlertPolicy(ctx, "databaseSecurityAlertPolicy", &sql.DatabaseSecurityAlertPolicyArgs{
			DatabaseName:            pulumi.String("testdb"),
			ResourceGroupName:       pulumi.String("securityalert-4799"),
			SecurityAlertPolicyName: pulumi.String("default"),
			ServerName:              pulumi.String("securityalert-6440"),
			State:                   sql.SecurityAlertsPolicyStateEnabled,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
