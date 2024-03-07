package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewServiceRegistry(ctx, "serviceRegistry", &appplatform.ServiceRegistryArgs{
			ResourceGroupName:   pulumi.String("myResourceGroup"),
			ServiceName:         pulumi.String("myservice"),
			ServiceRegistryName: pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
