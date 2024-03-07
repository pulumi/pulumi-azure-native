package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewTask(ctx, "task", &containerregistry.TaskArgs{
			IsSystemTask:      pulumi.Bool(true),
			Location:          pulumi.String("eastus"),
			LogTemplate:       pulumi.String("acr/tasks:{{.Run.OS}}"),
			RegistryName:      pulumi.String("myRegistry"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Status:            pulumi.String("Enabled"),
			Tags: pulumi.StringMap{
				"testkey": pulumi.String("value"),
			},
			TaskName: pulumi.String("quicktask"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
