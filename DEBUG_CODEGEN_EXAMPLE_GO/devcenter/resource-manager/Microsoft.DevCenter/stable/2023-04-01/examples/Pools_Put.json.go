package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devcenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devcenter.NewPool(ctx, "pool", &devcenter.PoolArgs{
			DevBoxDefinitionName:  pulumi.String("WebDevBox"),
			LicenseType:           pulumi.String("Windows_Client"),
			LocalAdministrator:    pulumi.String("Enabled"),
			Location:              pulumi.String("centralus"),
			NetworkConnectionName: pulumi.String("Network1-westus2"),
			PoolName:              pulumi.String("DevPool"),
			ProjectName:           pulumi.String("DevProject"),
			ResourceGroupName:     pulumi.String("rg1"),
			StopOnDisconnect: &devcenter.StopOnDisconnectConfigurationArgs{
				GracePeriodMinutes: pulumi.Int(60),
				Status:             pulumi.String("Enabled"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
