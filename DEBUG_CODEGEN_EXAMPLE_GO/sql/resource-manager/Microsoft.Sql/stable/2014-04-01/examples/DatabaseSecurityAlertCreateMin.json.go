package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabaseThreatDetectionPolicy(ctx, "databaseThreatDetectionPolicy", &sql.DatabaseThreatDetectionPolicyArgs{
			DatabaseName:            pulumi.String("testdb"),
			ResourceGroupName:       pulumi.String("securityalert-4799"),
			SecurityAlertPolicyName: pulumi.String("default"),
			ServerName:              pulumi.String("securityalert-6440"),
			State:                   pulumi.String("Enabled"),
			StorageAccountAccessKey: pulumi.String("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
			StorageEndpoint:         pulumi.String("https://mystorage.blob.core.windows.net"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
