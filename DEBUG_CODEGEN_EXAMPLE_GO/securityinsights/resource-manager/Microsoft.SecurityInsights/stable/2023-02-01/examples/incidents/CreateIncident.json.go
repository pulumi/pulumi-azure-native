package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewIncident(ctx, "incident", &securityinsights.IncidentArgs{
			Classification:        pulumi.String("FalsePositive"),
			ClassificationComment: pulumi.String("Not a malicious activity"),
			ClassificationReason:  pulumi.String("IncorrectAlertLogic"),
			Description:           pulumi.String("This is a demo incident"),
			FirstActivityTimeUtc:  pulumi.String("2019-01-01T13:00:30Z"),
			IncidentId:            pulumi.String("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
			LastActivityTimeUtc:   pulumi.String("2019-01-01T13:05:30Z"),
			Owner: &securityinsights.IncidentOwnerInfoArgs{
				ObjectId: pulumi.String("2046feea-040d-4a46-9e2b-91c2941bfa70"),
			},
			ResourceGroupName: pulumi.String("myRg"),
			Severity:          pulumi.String("High"),
			Status:            pulumi.String("Closed"),
			Title:             pulumi.String("My incident"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
