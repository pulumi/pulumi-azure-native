package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/peering/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := peering.NewConnectionMonitorTest(ctx, "connectionMonitorTest", &peering.ConnectionMonitorTestArgs{
			ConnectionMonitorTestName: pulumi.String("connectionMonitorTestName"),
			Destination:               pulumi.String("Example Destination"),
			DestinationPort:           pulumi.Int(443),
			PeeringServiceName:        pulumi.String("peeringServiceName"),
			ResourceGroupName:         pulumi.String("rgName"),
			SourceAgent:               pulumi.String("Example Source Agent"),
			TestFrequencyInSec:        pulumi.Int(30),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
