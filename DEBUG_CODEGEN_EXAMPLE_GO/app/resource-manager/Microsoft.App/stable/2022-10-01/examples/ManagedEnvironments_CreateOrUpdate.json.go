package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewManagedEnvironment(ctx, "managedEnvironment", &app.ManagedEnvironmentArgs{
			AppLogsConfiguration: app.AppLogsConfigurationResponse{
				LogAnalyticsConfiguration: &app.LogAnalyticsConfigurationArgs{
					CustomerId: pulumi.String("string"),
					SharedKey:  pulumi.String("string"),
				},
			},
			CustomDomainConfiguration: &app.CustomDomainConfigurationArgs{
				CertificatePassword: pulumi.String("private key password"),
				CertificateValue:    pulumi.String("Y2VydA=="),
				DnsSuffix:           pulumi.String("www.my-name.com"),
			},
			DaprAIConnectionString: pulumi.String("InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/"),
			EnvironmentName:        pulumi.String("testcontainerenv"),
			Kind:                   pulumi.String("serverless"),
			Location:               pulumi.String("East US"),
			ResourceGroupName:      pulumi.String("examplerg"),
			Sku: &app.EnvironmentSkuPropertiesArgs{
				Name: pulumi.String("Premium"),
			},
			VnetConfiguration: app.VnetConfigurationResponse{
				OutboundSettings: &app.ManagedEnvironmentOutboundSettingsArgs{
					OutBoundType:              pulumi.String("UserDefinedRouting"),
					VirtualNetworkApplianceIp: pulumi.String("192.168.1.20"),
				},
			},
			WorkloadProfiles: []app.WorkloadProfileArgs{
				{
					MaximumCount:        pulumi.Int(12),
					MinimumCount:        pulumi.Int(3),
					WorkloadProfileType: pulumi.String("GeneralPurpose"),
				},
				{
					MaximumCount:        pulumi.Int(6),
					MinimumCount:        pulumi.Int(3),
					WorkloadProfileType: pulumi.String("MemoryOptimized"),
				},
				{
					MaximumCount:        pulumi.Int(6),
					MinimumCount:        pulumi.Int(3),
					WorkloadProfileType: pulumi.String("ComputeOptimized"),
				},
			},
			ZoneRedundant: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
