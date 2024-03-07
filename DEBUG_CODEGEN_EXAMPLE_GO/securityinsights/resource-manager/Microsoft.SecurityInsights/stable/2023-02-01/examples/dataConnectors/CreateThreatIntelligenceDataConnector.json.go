package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewTIDataConnector(ctx, "tiDataConnector", &securityinsights.TIDataConnectorArgs{
			DataConnectorId: pulumi.String("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
			DataTypes: &securityinsights.TIDataConnectorDataTypesArgs{
				Indicators: &securityinsights.TIDataConnectorDataTypesIndicatorsArgs{
					State: pulumi.String("Enabled"),
				},
			},
			Kind:              pulumi.String("ThreatIntelligence"),
			ResourceGroupName: pulumi.String("myRg"),
			TenantId:          pulumi.String("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
			TipLookbackPeriod: pulumi.String("2020-01-01T13:00:30.123Z"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
