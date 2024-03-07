package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resources.NewDeploymentStackAtSubscription(ctx, "deploymentStackAtSubscription", &resources.DeploymentStackAtSubscriptionArgs{
			ActionOnUnmanage: &resources.DeploymentStackPropertiesActionOnUnmanageArgs{
				ManagementGroups: pulumi.String("detach"),
				ResourceGroups:   pulumi.String("delete"),
				Resources:        pulumi.String("delete"),
			},
			DenySettings: &resources.DenySettingsArgs{
				ApplyToChildScopes: pulumi.Bool(false),
				ExcludedActions: pulumi.StringArray{
					pulumi.String("action"),
				},
				ExcludedPrincipals: pulumi.StringArray{
					pulumi.String("principal"),
				},
				Mode: pulumi.String("denyDelete"),
			},
			DeploymentStackName: pulumi.String("simpleDeploymentStack"),
			Location:            pulumi.String("eastus"),
			Parameters: pulumi.Any{
				Parameter1: map[string]interface{}{
					"value": "a string",
				},
			},
			Tags: pulumi.StringMap{
				"tagkey": pulumi.String("tagVal"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
