package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerinstance/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerinstance.NewContainerGroup(ctx, "containerGroup", &containerinstance.ContainerGroupArgs{
			ConfidentialComputeProperties: &containerinstance.ConfidentialComputePropertiesArgs{
				CcePolicy: pulumi.String("eyJhbGxvd19hbGwiOiB0cnVlLCAiY29udGFpbmVycyI6IHsibGVuZ3RoIjogMCwgImVsZW1lbnRzIjogbnVsbH19"),
			},
			ContainerGroupName: pulumi.String("demo1"),
			Containers: containerinstance.ContainerArray{
				&containerinstance.ContainerArgs{
					Command:              pulumi.StringArray{},
					EnvironmentVariables: containerinstance.EnvironmentVariableArray{},
					Image:                pulumi.String("confiimage"),
					Name:                 pulumi.String("accdemo"),
					Ports: containerinstance.ContainerPortArray{
						&containerinstance.ContainerPortArgs{
							Port: pulumi.Int(8000),
						},
					},
					Resources: &containerinstance.ResourceRequirementsArgs{
						Requests: &containerinstance.ResourceRequestsArgs{
							Cpu:        pulumi.Float64(1),
							MemoryInGB: pulumi.Float64(1.5),
						},
					},
					SecurityContext: &containerinstance.SecurityContextDefinitionArgs{
						Capabilities: &containerinstance.SecurityContextCapabilitiesDefinitionArgs{
							Add: pulumi.StringArray{
								pulumi.String("CAP_NET_ADMIN"),
							},
						},
						Privileged: pulumi.Bool(false),
					},
				},
			},
			ImageRegistryCredentials: containerinstance.ImageRegistryCredentialArray{},
			IpAddress: &containerinstance.IpAddressArgs{
				Ports: containerinstance.PortArray{
					&containerinstance.PortArgs{
						Port:     pulumi.Int(8000),
						Protocol: pulumi.String("TCP"),
					},
				},
				Type: pulumi.String("Public"),
			},
			Location:          pulumi.String("westeurope"),
			OsType:            pulumi.String("Linux"),
			ResourceGroupName: pulumi.String("demo"),
			Sku:               pulumi.String("Confidential"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
