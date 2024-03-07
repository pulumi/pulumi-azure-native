package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/providerhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := providerhub.NewOperationByProviderRegistration(ctx, "operationByProviderRegistration", &providerhub.OperationByProviderRegistrationArgs{
			ProviderNamespace: pulumi.String("Microsoft.Contoso"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
