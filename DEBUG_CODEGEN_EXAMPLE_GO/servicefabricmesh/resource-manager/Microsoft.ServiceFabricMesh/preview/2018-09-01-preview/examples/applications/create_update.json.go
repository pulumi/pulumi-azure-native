package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabricmesh/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabricmesh.NewApplication(ctx, "application", &servicefabricmesh.ApplicationArgs{
			ApplicationResourceName: pulumi.String("sampleApplication"),
			Description:             pulumi.String("Service Fabric Mesh sample application."),
			Location:                pulumi.String("EastUS"),
			ResourceGroupName:       pulumi.String("sbz_demo"),
			Services: []servicefabricmesh.ServiceResourceDescriptionArgs{
				{
					CodePackages: []servicefabricmesh.ContainerCodePackagePropertiesArgs{
						{
							Endpoints: servicefabricmesh.EndpointPropertiesArray{
								{
									Name: pulumi.String("helloWorldListener"),
									Port: pulumi.Int(80),
								},
							},
							Image: pulumi.String("seabreeze/sbz-helloworld:1.0-alpine"),
							Name:  pulumi.String("helloWorldCode"),
							Resources: {
								Requests: {
									Cpu:        pulumi.Float64(1),
									MemoryInGB: pulumi.Float64(1),
								},
							},
						},
					},
					Description: pulumi.String("SeaBreeze Hello World Service."),
					Name:        pulumi.String("helloWorldService"),
					NetworkRefs: servicefabricmesh.NetworkRefArray{
						{
							EndpointRefs: servicefabricmesh.EndpointRefArray{
								{
									Name: pulumi.String("helloWorldListener"),
								},
							},
							Name: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/sbz_demo/providers/Microsoft.ServiceFabricMesh/networks/sampleNetwork"),
						},
					},
					OsType:       pulumi.String("Linux"),
					ReplicaCount: pulumi.Int(1),
				},
			},
			Tags: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
