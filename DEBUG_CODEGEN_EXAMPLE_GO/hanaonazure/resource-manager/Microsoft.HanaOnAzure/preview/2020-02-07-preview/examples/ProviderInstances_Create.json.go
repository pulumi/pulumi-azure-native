package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hanaonazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hanaonazure.NewProviderInstance(ctx, "providerInstance", &hanaonazure.ProviderInstanceArgs{
			Metadata:             pulumi.String("{\"key\":\"value\"}"),
			Properties:           pulumi.String("{\"hostname\":\"10.0.0.6\",\"dbName\":\"SYSTEMDB\",\"sqlPort\":30013,\"dbUsername\":\"SYSTEM\",\"dbPassword\":\"PASSWORD\"}"),
			ProviderInstanceName: pulumi.String("myProviderInstance"),
			ResourceGroupName:    pulumi.String("myResourceGroup"),
			SapMonitorName:       pulumi.String("mySapMonitor"),
			Type:                 pulumi.String("hana"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
