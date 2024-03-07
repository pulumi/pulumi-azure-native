package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabricmesh/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabricmesh.NewGateway(ctx, "gateway", &servicefabricmesh.GatewayArgs{
			Description: pulumi.String("Service Fabric Mesh sample gateway."),
			DestinationNetwork: &servicefabricmesh.NetworkRefArgs{
				Name: pulumi.String("helloWorldNetwork"),
			},
			GatewayResourceName: pulumi.String("sampleGateway"),
			Http: []servicefabricmesh.HttpConfigArgs{
				{
					Hosts: servicefabricmesh.HttpHostConfigArray{
						{
							Name: pulumi.String("contoso.com"),
							Routes: servicefabricmesh.HttpRouteConfigArray{
								{
									Destination: {
										ApplicationName: pulumi.String("httpHelloWorldApp"),
										EndpointName:    pulumi.String("indexHttpEndpoint"),
										ServiceName:     pulumi.String("indexService"),
									},
									Match: {
										Headers: servicefabricmesh.HttpRouteMatchHeaderArray{
											{
												Name:  pulumi.String("accept"),
												Type:  pulumi.String("exact"),
												Value: pulumi.String("application/json"),
											},
										},
										Path: {
											Rewrite: pulumi.String("/"),
											Type:    pulumi.String("prefix"),
											Value:   pulumi.String("/index"),
										},
									},
									Name: pulumi.String("index"),
								},
							},
						},
					},
					Name: pulumi.String("contosoWebsite"),
					Port: pulumi.Int(8081),
				},
			},
			Location:          pulumi.String("EastUS"),
			ResourceGroupName: pulumi.String("sbz_demo"),
			SourceNetwork: &servicefabricmesh.NetworkRefArgs{
				Name: pulumi.String("Open"),
			},
			Tags: nil,
			Tcp: []servicefabricmesh.TcpConfigArgs{
				{
					Destination: {
						ApplicationName: pulumi.String("helloWorldApp"),
						EndpointName:    pulumi.String("helloWorldListener"),
						ServiceName:     pulumi.String("helloWorldService"),
					},
					Name: pulumi.String("web"),
					Port: pulumi.Int(80),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
