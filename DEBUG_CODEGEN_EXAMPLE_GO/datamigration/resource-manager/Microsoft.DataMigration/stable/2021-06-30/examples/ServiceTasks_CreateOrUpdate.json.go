package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datamigration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datamigration.NewServiceTask(ctx, "serviceTask", &datamigration.ServiceTaskArgs{
			GroupName: pulumi.String("DmsSdkRg"),
			Properties: pulumi.Any{
				Input: &datamigration.MongoDbConnectionInfoArgs{
					ServerVersion: "NA",
				},
				TaskType: pulumi.String("Service.Check.OCI"),
			},
			ServiceName: pulumi.String("DmsSdkService"),
			TaskName:    pulumi.String("DmsSdkTask"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
