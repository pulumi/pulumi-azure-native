package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securitydevops/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securitydevops.NewGitHubConnector(ctx, "gitHubConnector", &securitydevops.GitHubConnectorArgs{
			GitHubConnectorName: pulumi.String("testconnector"),
			Location:            pulumi.String("West US"),
			Properties: &securitydevops.GitHubConnectorPropertiesArgs{
				Code: pulumi.String("00000000000000000000"),
			},
			ResourceGroupName: pulumi.String("westusrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
