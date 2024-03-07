package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewCodeContainer(ctx, "codeContainer", &machinelearningservices.CodeContainerArgs{
			CodeContainerProperties: &machinelearningservices.CodeContainerTypeArgs{
				Description: pulumi.String("string"),
				Tags: pulumi.StringMap{
					"tag1": pulumi.String("value1"),
					"tag2": pulumi.String("value2"),
				},
			},
			Name:              pulumi.String("testContainer"),
			ResourceGroupName: pulumi.String("testrg123"),
			WorkspaceName:     pulumi.String("testworkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
