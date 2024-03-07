package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logic.NewIntegrationAccountAssembly(ctx, "integrationAccountAssembly", &logic.IntegrationAccountAssemblyArgs{
			AssemblyArtifactName:   pulumi.String("testAssembly"),
			IntegrationAccountName: pulumi.String("testIntegrationAccount"),
			Location:               pulumi.String("westus"),
			Properties: &logic.AssemblyPropertiesArgs{
				AssemblyName: pulumi.String("System.IdentityModel.Tokens.Jwt"),
				Content:      pulumi.Any("Base64 encoded Assembly Content"),
				Metadata:     nil,
			},
			ResourceGroupName: pulumi.String("testResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
