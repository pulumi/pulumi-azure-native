package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewServiceFabric(ctx, "serviceFabric", &devtestlab.ServiceFabricArgs{
			EnvironmentId:           pulumi.String("{environmentId}"),
			ExternalServiceFabricId: pulumi.String("{serviceFabricId}"),
			LabName:                 pulumi.String("{labName}"),
			Location:                pulumi.String("{location}"),
			Name:                    pulumi.String("{serviceFabricName}"),
			ResourceGroupName:       pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"tagName1": pulumi.String("tagValue1"),
			},
			UserName: pulumi.String("{userName}"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
