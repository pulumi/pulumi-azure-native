package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewAssessmentProjectsOperation(ctx, "assessmentProjectsOperation", &migrate.AssessmentProjectsOperationArgs{
			AssessmentSolutionId:        pulumi.String("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sakanwar/providers/Microsoft.Storage/storageAccounts/sakanwar1204usa"),
			CustomerStorageAccountArmId: pulumi.String("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sakanwar/providers/Microsoft.Storage/storageAccounts/sakanwar1204usa"),
			Location:                    pulumi.String("southeastasia"),
			ProjectName:                 pulumi.String("sakanwar1204project"),
			ProjectStatus:               pulumi.String("Active"),
			ProvisioningState:           pulumi.String("Succeeded"),
			PublicNetworkAccess:         pulumi.String("Disabled"),
			ResourceGroupName:           pulumi.String("sakanwar"),
			Tags: pulumi.StringMap{
				"Migrate Project": pulumi.String("sakanwar-PE-SEA"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
