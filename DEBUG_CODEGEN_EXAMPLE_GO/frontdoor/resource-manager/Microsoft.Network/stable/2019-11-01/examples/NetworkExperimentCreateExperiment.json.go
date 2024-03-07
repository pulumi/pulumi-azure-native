package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewExperiment(ctx, "experiment", &network.ExperimentArgs{
			Description:  pulumi.String("this is my first experiment!"),
			EnabledState: pulumi.String("Enabled"),
			EndpointA: &network.ExperimentEndpointArgs{
				Endpoint: pulumi.String("endpointA.net"),
				Name:     pulumi.String("endpoint A"),
			},
			EndpointB: &network.ExperimentEndpointArgs{
				Endpoint: pulumi.String("endpointB.net"),
				Name:     pulumi.String("endpoint B"),
			},
			ExperimentName:    pulumi.String("MyExperiment"),
			ProfileName:       pulumi.String("MyProfile"),
			ResourceGroupName: pulumi.String("MyResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
