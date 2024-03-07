package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewServiceRunner(ctx, "serviceRunner", &devtestlab.ServiceRunnerArgs{
			Identity: &devtestlab.IdentityPropertiesArgs{
				ClientSecretUrl: pulumi.String("{identityClientSecretUrl}"),
				PrincipalId:     pulumi.String("{identityPrincipalId}"),
				TenantId:        pulumi.String("{identityTenantId}"),
				Type:            pulumi.String("{identityType}"),
			},
			LabName:           pulumi.String("{devtestlabName}"),
			Location:          pulumi.String("{location}"),
			Name:              pulumi.String("{servicerunnerName}"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"tagName1": pulumi.String("tagValue1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
