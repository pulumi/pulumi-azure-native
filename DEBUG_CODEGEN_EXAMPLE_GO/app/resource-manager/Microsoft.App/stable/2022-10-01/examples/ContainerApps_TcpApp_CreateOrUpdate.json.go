package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewContainerApp(ctx, "containerApp", &app.ContainerAppArgs{
			Configuration: &app.ConfigurationArgs{
				Ingress: &app.IngressArgs{
					ExposedPort: pulumi.Int(4000),
					External:    pulumi.Bool(true),
					TargetPort:  pulumi.Int(3000),
					Traffic: app.TrafficWeightArray{
						&app.TrafficWeightArgs{
							RevisionName: pulumi.String("testcontainerapptcp-ab1234"),
							Weight:       pulumi.Int(100),
						},
					},
					Transport: pulumi.String("tcp"),
				},
			},
			ContainerAppName:  pulumi.String("testcontainerapptcp"),
			EnvironmentId:     pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
			Location:          pulumi.String("East US"),
			ResourceGroupName: pulumi.String("rg"),
			Template: &app.TemplateArgs{
				Containers: app.ContainerArray{
					&app.ContainerArgs{
						Image: pulumi.String("repo/testcontainerapptcp:v1"),
						Name:  pulumi.String("testcontainerapptcp"),
						Probes: app.ContainerAppProbeArray{
							&app.ContainerAppProbeArgs{
								InitialDelaySeconds: pulumi.Int(3),
								PeriodSeconds:       pulumi.Int(3),
								TcpSocket: &app.ContainerAppProbeTcpSocketArgs{
									Port: pulumi.Int(8080),
								},
								Type: pulumi.String("Liveness"),
							},
						},
					},
				},
				Scale: &app.ScaleArgs{
					MaxReplicas: pulumi.Int(5),
					MinReplicas: pulumi.Int(1),
					Rules: app.ScaleRuleArray{
						&app.ScaleRuleArgs{
							Name: pulumi.String("tcpscalingrule"),
							Tcp: &app.TcpScaleRuleArgs{
								Metadata: pulumi.StringMap{
									"concurrentConnections": pulumi.String("50"),
								},
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
