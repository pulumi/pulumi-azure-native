package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewIotHubDataConnection(ctx, "iotHubDataConnection", &synapse.IotHubDataConnectionArgs{
			DataConnectionName: pulumi.String("DataConnections8"),
			DatabaseName:       pulumi.String("KustoDatabase8"),
			KustoPoolName:      pulumi.String("kustoclusterrptest4"),
			ResourceGroupName:  pulumi.String("kustorptest"),
			WorkspaceName:      pulumi.String("synapseWorkspaceName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
