package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewOutput(ctx, "output", &streamanalytics.OutputArgs{
			Datasource: streamanalytics.GatewayMessageBusOutputDataSource{
				Topic: "EdgeTopic1",
				Type:  "GatewayMessageBus",
			},
			JobName:           pulumi.String("sj2331"),
			OutputName:        pulumi.String("output3022"),
			ResourceGroupName: pulumi.String("sjrg7983"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
