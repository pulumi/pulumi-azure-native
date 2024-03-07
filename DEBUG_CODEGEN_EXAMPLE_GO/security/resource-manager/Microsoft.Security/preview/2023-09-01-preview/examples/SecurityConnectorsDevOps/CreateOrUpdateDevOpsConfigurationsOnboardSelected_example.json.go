package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewDevOpsConfiguration(ctx, "devOpsConfiguration", &security.DevOpsConfigurationArgs{
			Properties: &security.DevOpsConfigurationPropertiesArgs{
				Authorization: &security.AuthorizationArgs{
					Code: pulumi.String("00000000000000000000"),
				},
				AutoDiscovery: pulumi.String("Disabled"),
				TopLevelInventoryList: pulumi.StringArray{
					pulumi.String("org1"),
					pulumi.String("org2"),
				},
			},
			ResourceGroupName:     pulumi.String("myRg"),
			SecurityConnectorName: pulumi.String("mySecurityConnectorName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
