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
			ProviderSettings: workloads.SapNetWeaverProviderInstanceProperties{
				ProviderType: "SapNetWeaver",
				SapClientId:  "111",
				SapHostFileEntries: []string{
					"127.0.0.1 name fqdn",
				},
				SapHostname:       "name",
				SapInstanceNr:     "00",
				SapPassword:       "****",
				SapPasswordUri:    "",
				SapPortNumber:     "1234",
				SapSid:            "SID",
				SapUsername:       "username",
				SslCertificateUri: "https://storageaccount.blob.core.windows.net/containername/filename",
				SslPreference:     "ServerCertificate",
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
