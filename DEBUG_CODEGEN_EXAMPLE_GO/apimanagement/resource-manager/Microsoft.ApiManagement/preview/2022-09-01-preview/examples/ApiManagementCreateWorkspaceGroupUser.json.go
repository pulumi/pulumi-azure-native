package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewWorkspaceGroupUser(ctx, "workspaceGroupUser", &apimanagement.WorkspaceGroupUserArgs{
			GroupId:           pulumi.String("tempgroup"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			UserId:            pulumi.String("59307d350af58404d8a26300"),
			WorkspaceId:       pulumi.String("wks1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
