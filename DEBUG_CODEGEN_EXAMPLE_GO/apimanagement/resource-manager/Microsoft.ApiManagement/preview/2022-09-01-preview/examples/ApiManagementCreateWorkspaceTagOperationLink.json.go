package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewWorkspaceTagOperationLink(ctx, "workspaceTagOperationLink", &apimanagement.WorkspaceTagOperationLinkArgs{
			OperationId:       pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/echo-api/operations/op1"),
			OperationLinkId:   pulumi.String("link1"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			TagId:             pulumi.String("tag1"),
			WorkspaceId:       pulumi.String("wks1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
