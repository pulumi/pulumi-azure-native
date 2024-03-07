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
			ProviderSettings: workloads.HanaDbProviderInstanceProperties{
				DbName:                   "db",
				DbPassword:               "****",
				DbPasswordUri:            "",
				DbUsername:               "user",
				Hostname:                 "name",
				InstanceNumber:           "00",
				ProviderType:             "SapHana",
				SapSid:                   "SID",
				SqlPort:                  "0000",
				SslCertificateUri:        "https://storageaccount.blob.core.windows.net/containername/filename",
				SslHostNameInCertificate: "xyz.domain.com",
				SslPreference:            "ServerCertificate",
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
