package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerinstance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerinstance.NewContainerGroup(ctx, "containerGroup", &containerinstance.ContainerGroupArgs{
			ContainerGroupName: pulumi.String("demo1"),
			Containers: containerinstance.ContainerArray{
				&containerinstance.ContainerArgs{
					Command: pulumi.StringArray{
						pulumi.String("/bin/sh"),
						pulumi.String("-c"),
						pulumi.String("sleep 10"),
					},
					Image: pulumi.String("alpine:latest"),
					Name:  pulumi.String("test-container-001"),
					Resources: &containerinstance.ResourceRequirementsArgs{
						Requests: &containerinstance.ResourceRequestsArgs{
							Cpu:        pulumi.Float64(1),
							MemoryInGB: pulumi.Float64(1),
						},
					},
				},
			},
			Location:          pulumi.String("eastus"),
			OsType:            pulumi.String("Linux"),
			Priority:          pulumi.String("Spot"),
			ResourceGroupName: pulumi.String("demo"),
			RestartPolicy:     pulumi.String("Never"),
			Sku:               pulumi.String("Standard"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
