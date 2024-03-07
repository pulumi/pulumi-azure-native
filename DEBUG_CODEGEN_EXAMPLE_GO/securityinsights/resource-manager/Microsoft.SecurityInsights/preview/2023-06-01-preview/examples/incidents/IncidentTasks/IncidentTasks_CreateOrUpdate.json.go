package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewIncidentTask(ctx, "incidentTask", &securityinsights.IncidentTaskArgs{
			Description:       pulumi.String("Task description"),
			IncidentId:        pulumi.String("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
			IncidentTaskId:    pulumi.String("4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014"),
			ResourceGroupName: pulumi.String("myRg"),
			Status:            pulumi.String("New"),
			Title:             pulumi.String("Task title"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
