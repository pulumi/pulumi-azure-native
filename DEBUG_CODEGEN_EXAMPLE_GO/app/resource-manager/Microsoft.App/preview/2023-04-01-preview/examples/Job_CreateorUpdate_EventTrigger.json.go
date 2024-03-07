package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewJob(ctx, "job", &app.JobArgs{
			Configuration: &app.JobConfigurationArgs{
				EventTriggerConfig: &app.JobConfigurationEventTriggerConfigArgs{
					Parallelism:            pulumi.Int(4),
					ReplicaCompletionCount: pulumi.Int(1),
					Scale: &app.JobScaleArgs{
						MaxExecutions:   pulumi.Int(5),
						MinExecutions:   pulumi.Int(1),
						PollingInterval: pulumi.Int(40),
						Rules: app.JobScaleRuleArray{
							&app.JobScaleRuleArgs{
								Metadata: pulumi.Any{
									TopicName: "my-topic",
								},
								Name: pulumi.String("servicebuscalingrule"),
								Type: pulumi.String("azure-servicebus"),
							},
						},
					},
				},
				ReplicaRetryLimit: pulumi.Int(10),
				ReplicaTimeout:    pulumi.Int(10),
				TriggerType:       pulumi.String("Event"),
			},
			EnvironmentId:     pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
			JobName:           pulumi.String("testcontainerappsjob0"),
			Location:          pulumi.String("East US"),
			ResourceGroupName: pulumi.String("rg"),
			Template: &app.JobTemplateArgs{
				Containers: app.ContainerArray{
					&app.ContainerArgs{
						Image: pulumi.String("repo/testcontainerappsjob0:v1"),
						Name:  pulumi.String("testcontainerappsjob0"),
					},
				},
				InitContainers: app.InitContainerArray{
					&app.InitContainerArgs{
						Args: pulumi.StringArray{
							pulumi.String("-c"),
							pulumi.String("while true; do echo hello; sleep 10;done"),
						},
						Command: pulumi.StringArray{
							pulumi.String("/bin/sh"),
						},
						Image: pulumi.String("repo/testcontainerappsjob0:v4"),
						Name:  pulumi.String("testinitcontainerAppsJob0"),
						Resources: &app.ContainerResourcesArgs{
							Cpu:    pulumi.Float64(0.5),
							Memory: pulumi.String("1Gi"),
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
