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
			Properties: appplatform.DeploymentResourcePropertiesResponse{
				DeploymentSettings: appplatform.DeploymentSettingsResponse{
					AddonConfigs: pulumi.Map{
						"ApplicationConfigurationService": pulumi.Any{
							Patterns: []string{
								"mypattern",
							},
						},
					},
					Apms: appplatform.ApmReferenceArray{
						&appplatform.ApmReferenceArgs{
							ResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apms/myappinsights"),
						},
					},
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
				Source: appplatform.SourceUploadedUserSourceInfo{
					ArtifactSelector: "sub-module-1",
					RelativePath:     "resources/a172cedcae47474b615c54d510a5d84a8dea3032e958587430b413538be3f333-2019082605-e3095339-1723-44b7-8b5e-31b1003978bc",
					Type:             "Source",
					Version:          "1.0",
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
			Sku: &appplatform.SkuArgs{
				Capacity: pulumi.Int(1),
				Name:     pulumi.String("S0"),
				Tier:     pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
