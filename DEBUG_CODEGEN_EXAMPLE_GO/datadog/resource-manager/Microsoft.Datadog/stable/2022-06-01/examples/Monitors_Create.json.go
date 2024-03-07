package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datadog/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datadog.NewMonitor(ctx, "monitor", &datadog.MonitorArgs{
			Location:    pulumi.String("West US"),
			MonitorName: pulumi.String("myMonitor"),
			Properties: datadog.MonitorPropertiesResponse{
				DatadogOrganizationProperties: &datadog.DatadogOrganizationPropertiesArgs{
					EnterpriseAppId: pulumi.String("00000000-0000-0000-0000-000000000000"),
					Id:              pulumi.String("myOrg123"),
					LinkingAuthCode: pulumi.String("someAuthCode"),
					LinkingClientId: pulumi.String("00000000-0000-0000-0000-000000000000"),
					Name:            pulumi.String("myOrg"),
				},
				MonitoringStatus: pulumi.String("Enabled"),
				UserInfo: &datadog.UserInfoArgs{
					EmailAddress: pulumi.String("alice@microsoft.com"),
					Name:         pulumi.String("Alice"),
					PhoneNumber:  pulumi.String("123-456-7890"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &datadog.ResourceSkuArgs{
				Name: pulumi.String("free_Monthly"),
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Dev"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
