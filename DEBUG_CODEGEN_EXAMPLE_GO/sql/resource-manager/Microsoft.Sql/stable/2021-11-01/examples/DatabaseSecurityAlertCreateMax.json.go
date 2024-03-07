package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabaseSecurityAlertPolicy(ctx, "databaseSecurityAlertPolicy", &sql.DatabaseSecurityAlertPolicyArgs{
			DatabaseName: pulumi.String("testdb"),
			DisabledAlerts: pulumi.StringArray{
				pulumi.String("Sql_Injection"),
				pulumi.String("Usage_Anomaly"),
			},
			EmailAccountAdmins: pulumi.Bool(true),
			EmailAddresses: pulumi.StringArray{
				pulumi.String("test@microsoft.com"),
				pulumi.String("user@microsoft.com"),
			},
			ResourceGroupName:       pulumi.String("securityalert-4799"),
			RetentionDays:           pulumi.Int(6),
			SecurityAlertPolicyName: pulumi.String("default"),
			ServerName:              pulumi.String("securityalert-6440"),
			State:                   sql.SecurityAlertsPolicyStateEnabled,
			StorageAccountAccessKey: pulumi.String("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
			StorageEndpoint:         pulumi.String("https://mystorage.blob.core.windows.net"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
