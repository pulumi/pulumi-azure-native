package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewSecret(ctx, "secret", &devtestlab.SecretArgs{
			LabName:           pulumi.String("{labName}"),
			Name:              pulumi.String("{secretName}"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			UserName:          pulumi.String("{userName}"),
			Value:             pulumi.String("{secret}"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
