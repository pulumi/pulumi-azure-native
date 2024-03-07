package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewTaskRun(ctx, "taskRun", &containerregistry.TaskRunArgs{
			ForceUpdateTag:    pulumi.String("test"),
			RegistryName:      pulumi.String("myRegistry"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			RunRequest: containerregistry.EncodedTaskRunRequest{
				Credentials:          nil,
				EncodedTaskContent:   "c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K",
				EncodedValuesContent: "Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg==",
				Platform: containerregistry.PlatformProperties{
					Architecture: "amd64",
					Os:           "Linux",
				},
				Type:   "EncodedTaskRunRequest",
				Values: []containerregistry.SetValue{},
			},
			TaskRunName: pulumi.String("myRun"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
