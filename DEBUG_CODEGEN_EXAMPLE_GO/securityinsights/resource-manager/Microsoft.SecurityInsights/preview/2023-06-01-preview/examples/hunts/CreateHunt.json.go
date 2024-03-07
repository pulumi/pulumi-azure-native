package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewHunt(ctx, "hunt", &securityinsights.HuntArgs{
			AttackTactics: pulumi.StringArray{
				pulumi.String("Reconnaissance"),
			},
			AttackTechniques: pulumi.StringArray{
				pulumi.String("T1595"),
			},
			Description:      pulumi.String("Log4J Hunt Description"),
			DisplayName:      pulumi.String("Log4J new hunt"),
			HuntId:           pulumi.String("163e7b2a-a2ec-4041-aaba-d878a38f265f"),
			HypothesisStatus: pulumi.String("Unknown"),
			Labels: pulumi.StringArray{
				pulumi.String("Label1"),
				pulumi.String("Label2"),
			},
			Owner: &securityinsights.HuntOwnerArgs{
				ObjectId: pulumi.String("873b5263-5d34-4149-b356-ad341b01e123"),
			},
			ResourceGroupName: pulumi.String("myRg"),
			Status:            pulumi.String("New"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
