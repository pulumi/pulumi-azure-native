package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewConnector(ctx, "connector", &security.ConnectorArgs{
			AuthenticationDetails: security.AwsCredsAuthenticationDetailsProperties{
				AuthenticationType: "awsCreds",
				AwsAccessKeyId:     "AKIARPZCNODDNAEQFSOE",
				AwsSecretAccessKey: "<awsSecretAccessKey>",
			},
			ConnectorName: pulumi.String("aws_dev1"),
			HybridComputeSettings: &security.HybridComputeSettingsPropertiesArgs{
				AutoProvision: pulumi.String("On"),
				ProxyServer: &security.ProxyServerPropertiesArgs{
					Ip:   pulumi.String("167.220.197.140"),
					Port: pulumi.String("34"),
				},
				Region:            pulumi.String("West US 2"),
				ResourceGroupName: pulumi.String("AwsConnectorRG"),
				ServicePrincipal: &security.ServicePrincipalPropertiesArgs{
					ApplicationId: pulumi.String("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
					Secret:        pulumi.String("<secret>"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
