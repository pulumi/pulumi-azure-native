package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewDeployment(ctx, "deployment", &appplatform.DeploymentArgs{
			AppName:        pulumi.String("myapp"),
			DeploymentName: pulumi.String("mydeployment"),
			Properties: &appplatform.DeploymentResourcePropertiesArgs{
				DeploymentSettings: &appplatform.DeploymentSettingsArgs{
					EnvironmentVariables: pulumi.StringMap{
						"env": pulumi.String("test"),
					},
					LivenessProbe: &appplatform.ProbeArgs{
						DisableProbe:        pulumi.Bool(false),
						FailureThreshold:    pulumi.Int(3),
						InitialDelaySeconds: pulumi.Int(30),
						PeriodSeconds:       pulumi.Int(10),
						ProbeAction: appplatform.HTTPGetAction{
							Path:   "/health",
							Scheme: "HTTP",
							Type:   "HTTPGetAction",
						},
					},
					ReadinessProbe: &appplatform.ProbeArgs{
						DisableProbe:        pulumi.Bool(false),
						FailureThreshold:    pulumi.Int(3),
						InitialDelaySeconds: pulumi.Int(30),
						PeriodSeconds:       pulumi.Int(10),
						ProbeAction: appplatform.HTTPGetAction{
							Path:   "/health",
							Scheme: "HTTP",
							Type:   "HTTPGetAction",
						},
					},
					ResourceRequests: &appplatform.ResourceRequestsArgs{
						Cpu:    pulumi.String("1000m"),
						Memory: pulumi.String("3Gi"),
					},
					TerminationGracePeriodSeconds: pulumi.Int(30),
				},
				Source: appplatform.CustomContainerUserSourceInfo{
					CustomContainer: appplatform.CustomContainer{
						Args: []string{
							"-c",
							"while true; do echo hello; sleep 10;done",
						},
						Command: []string{
							"/bin/sh",
						},
						ContainerImage: "myContainerImage:v1",
						ImageRegistryCredential: appplatform.ImageRegistryCredential{
							Password: "myPassword",
							Username: "myUsername",
						},
						LanguageFramework: "springboot",
						Server:            "myacr.azurecr.io",
					},
					Type: "Container",
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
