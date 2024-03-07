package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewDisasterRecoveryConfig(ctx, "disasterRecoveryConfig", &eventhub.DisasterRecoveryConfigArgs{
			Alias:             pulumi.String("sdk-DisasterRecovery-3814"),
			NamespaceName:     pulumi.String("sdk-Namespace-8859"),
			PartnerNamespace:  pulumi.String("sdk-Namespace-37"),
			ResourceGroupName: pulumi.String("exampleResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
