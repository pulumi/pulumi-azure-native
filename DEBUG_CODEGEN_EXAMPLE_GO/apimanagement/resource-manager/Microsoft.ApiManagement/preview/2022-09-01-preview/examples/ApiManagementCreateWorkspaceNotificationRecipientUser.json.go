package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewWorkspaceNotificationRecipientUser(ctx, "workspaceNotificationRecipientUser", &apimanagement.WorkspaceNotificationRecipientUserArgs{
			NotificationName:  pulumi.String("RequestPublisherNotificationMessage"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			UserId:            pulumi.String("576823d0a40f7e74ec07d642"),
			WorkspaceId:       pulumi.String("wks1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
