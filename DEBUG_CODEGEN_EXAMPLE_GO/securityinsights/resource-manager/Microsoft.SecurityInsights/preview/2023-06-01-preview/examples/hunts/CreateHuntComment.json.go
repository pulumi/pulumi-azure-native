package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewHuntComment(ctx, "huntComment", &securityinsights.HuntCommentArgs{
			HuntCommentId:     pulumi.String("2216d0e1-91e3-4902-89fd-d2df8c535096"),
			HuntId:            pulumi.String("163e7b2a-a2ec-4041-aaba-d878a38f265f"),
			Message:           pulumi.String("This is a test comment."),
			ResourceGroupName: pulumi.String("myRg"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
