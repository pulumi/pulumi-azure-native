package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kusto/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kusto.NewIotHubDataConnection(ctx, "iotHubDataConnection", &kusto.IotHubDataConnectionArgs{
			ClusterName:        pulumi.String("kustoCluster"),
			DataConnectionName: pulumi.String("dataConnectionTest"),
			DatabaseName:       pulumi.String("KustoDatabase1"),
			ResourceGroupName:  pulumi.String("kustorptest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
