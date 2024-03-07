package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewProviderInstance(ctx, "providerInstance", &workloads.ProviderInstanceArgs{
			MonitorName:          pulumi.String("mySapMonitor"),
			ProviderInstanceName: pulumi.String("myProviderInstance"),
			ProviderSettings: workloads.DB2ProviderInstanceProperties{
				DbName:        "dbName",
				DbPassword:    "password",
				DbPasswordUri: "",
				DbPort:        "dbPort",
				DbUsername:    "username",
				Hostname:      "hostname",
				ProviderType:  "Db2",
				SapSid:        "SID",
				SslPreference: "RootCertificate",
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
