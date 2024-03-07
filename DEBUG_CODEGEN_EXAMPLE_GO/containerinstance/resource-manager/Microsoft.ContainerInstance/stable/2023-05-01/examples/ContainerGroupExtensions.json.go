package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerinstance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerinstance.NewContainerGroup(ctx, "containerGroup", &containerinstance.ContainerGroupArgs{
			ContainerGroupName: pulumi.String("demo1"),
			Containers: []containerinstance.ContainerArgs{
				{
					Command:              pulumi.StringArray{},
					EnvironmentVariables: containerinstance.EnvironmentVariableArray{},
					Image:                pulumi.String("nginx"),
					Name:                 pulumi.String("demo1"),
					Ports: containerinstance.ContainerPortArray{
						{
							Port: pulumi.Int(80),
						},
					},
					Resources: {
						Requests: {
							Cpu:        pulumi.Float64(1),
							MemoryInGB: pulumi.Float64(1.5),
						},
					},
				},
			},
			Extensions: []containerinstance.DeploymentExtensionSpecArgs{
				{
					ExtensionType: pulumi.String("kube-proxy"),
					Name:          pulumi.String("kube-proxy"),
					ProtectedSettings: {
						KubeConfig: "<kubeconfig encoded string>",
					},
					Settings: {
						ClusterCidr: "10.240.0.0/16",
						KubeVersion: "v1.9.10",
					},
					Version: pulumi.String("1.0"),
				},
				{
					ExtensionType: pulumi.String("realtime-metrics"),
					Name:          pulumi.String("vk-realtime-metrics"),
					Version:       pulumi.String("1.0"),
				},
			},
			ImageRegistryCredentials: containerinstance.ImageRegistryCredentialArray{},
			IpAddress: containerinstance.IpAddressResponse{
				Ports: containerinstance.PortArray{
					&containerinstance.PortArgs{
						Port:     pulumi.Int(80),
						Protocol: pulumi.String("TCP"),
					},
				},
				Type: pulumi.String("Private"),
			},
			Location:          pulumi.String("eastus2"),
			OsType:            pulumi.String("Linux"),
			ResourceGroupName: pulumi.String("demo"),
			SubnetIds: []containerinstance.ContainerGroupSubnetIdArgs{
				{
					Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-rg-vnet/subnets/test-subnet"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
