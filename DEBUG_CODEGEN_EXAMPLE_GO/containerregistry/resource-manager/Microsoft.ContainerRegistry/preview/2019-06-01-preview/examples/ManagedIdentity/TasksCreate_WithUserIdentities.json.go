package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewTask(ctx, "task", &containerregistry.TaskArgs{
			AgentConfiguration: &containerregistry.AgentPropertiesArgs{
				Cpu: pulumi.Int(2),
			},
			Identity: &containerregistry.IdentityPropertiesArgs{
				Type: containerregistry.ResourceIdentityTypeUserAssigned,
				UserAssignedIdentities: containerregistry.UserIdentityPropertiesMap{
					"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1":  nil,
					"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": nil,
				},
			},
			IsSystemTask: pulumi.Bool(false),
			Location:     pulumi.String("eastus"),
			Platform: &containerregistry.PlatformPropertiesArgs{
				Architecture: pulumi.String("amd64"),
				Os:           pulumi.String("Linux"),
			},
			RegistryName:      pulumi.String("myRegistry"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Status:            pulumi.String("Enabled"),
			Step: containerregistry.DockerBuildStep{
				Arguments: []containerregistry.Argument{
					{
						IsSecret: false,
						Name:     "mytestargument",
						Value:    "mytestvalue",
					},
					{
						IsSecret: true,
						Name:     "mysecrettestargument",
						Value:    "mysecrettestvalue",
					},
				},
				ContextPath:    "src",
				DockerFilePath: "src/DockerFile",
				ImageNames: []string{
					"azurerest:testtag",
				},
				IsPushEnabled: true,
				NoCache:       false,
				Type:          "Docker",
			},
			Tags: pulumi.StringMap{
				"testkey": pulumi.String("value"),
			},
			TaskName: pulumi.String("mytTask"),
			Trigger: &containerregistry.TriggerPropertiesArgs{
				BaseImageTrigger: &containerregistry.BaseImageTriggerArgs{
					BaseImageTriggerType:     pulumi.String("Runtime"),
					Name:                     pulumi.String("myBaseImageTrigger"),
					UpdateTriggerEndpoint:    pulumi.String("https://user:pass@mycicd.webhook.com?token=foo"),
					UpdateTriggerPayloadType: pulumi.String("Default"),
				},
				SourceTriggers: containerregistry.SourceTriggerArray{
					&containerregistry.SourceTriggerArgs{
						Name: pulumi.String("mySourceTrigger"),
						SourceRepository: &containerregistry.SourcePropertiesArgs{
							Branch:        pulumi.String("master"),
							RepositoryUrl: pulumi.String("https://github.com/Azure/azure-rest-api-specs"),
							SourceControlAuthProperties: &containerregistry.AuthInfoArgs{
								Token:     pulumi.String("xxxxx"),
								TokenType: pulumi.String("PAT"),
							},
							SourceControlType: pulumi.String("Github"),
						},
						SourceTriggerEvents: pulumi.StringArray{
							pulumi.String("commit"),
						},
					},
				},
				TimerTriggers: containerregistry.TimerTriggerArray{
					&containerregistry.TimerTriggerArgs{
						Name:     pulumi.String("myTimerTrigger"),
						Schedule: pulumi.String("30 9 * * 1-5"),
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
