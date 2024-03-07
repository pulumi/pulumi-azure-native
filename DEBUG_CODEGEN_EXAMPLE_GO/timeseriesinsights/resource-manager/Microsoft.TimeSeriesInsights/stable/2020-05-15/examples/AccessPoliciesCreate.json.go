package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/timeseriesinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := timeseriesinsights.NewAccessPolicy(ctx, "accessPolicy", &timeseriesinsights.AccessPolicyArgs{
			AccessPolicyName:  pulumi.String("ap1"),
			Description:       pulumi.String("some description"),
			EnvironmentName:   pulumi.String("env1"),
			PrincipalObjectId: pulumi.String("aGuid"),
			ResourceGroupName: pulumi.String("rg1"),
			Roles: pulumi.StringArray{
				pulumi.String("Reader"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
