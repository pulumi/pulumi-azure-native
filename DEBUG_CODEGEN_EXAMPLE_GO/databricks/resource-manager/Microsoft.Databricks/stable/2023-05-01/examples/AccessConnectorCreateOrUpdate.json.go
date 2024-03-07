package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databricks/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databricks.NewAccessConnector(ctx, "accessConnector", &databricks.AccessConnectorArgs{
			ConnectorName:     pulumi.String("myAccessConnector"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
