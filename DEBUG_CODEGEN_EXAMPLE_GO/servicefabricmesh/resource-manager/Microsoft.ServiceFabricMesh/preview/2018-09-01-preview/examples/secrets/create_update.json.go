package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabricmesh/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabricmesh.NewSecret(ctx, "secret", &servicefabricmesh.SecretArgs{
			Location:           pulumi.String("EastUS"),
			Properties:         nil,
			ResourceGroupName:  pulumi.String("sbz_demo"),
			SecretResourceName: pulumi.String("dbConnectionString"),
			Tags:               nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
