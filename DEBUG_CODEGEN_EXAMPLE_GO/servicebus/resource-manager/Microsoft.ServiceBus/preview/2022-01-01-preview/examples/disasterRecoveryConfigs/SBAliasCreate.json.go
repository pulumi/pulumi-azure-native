package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicebus.NewDisasterRecoveryConfig(ctx, "disasterRecoveryConfig", &servicebus.DisasterRecoveryConfigArgs{
			Alias:             pulumi.String("sdk-Namespace-8860"),
			AlternateName:     pulumi.String("alternameforAlias-Namespace-8860"),
			NamespaceName:     pulumi.String("sdk-Namespace-8860"),
			PartnerNamespace:  pulumi.String("sdk-Namespace-37"),
			ResourceGroupName: pulumi.String("ardsouzatestRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
