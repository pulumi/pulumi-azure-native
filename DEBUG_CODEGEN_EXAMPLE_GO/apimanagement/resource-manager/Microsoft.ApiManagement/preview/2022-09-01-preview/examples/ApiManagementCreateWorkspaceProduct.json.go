package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewWorkspaceProduct(ctx, "workspaceProduct", &apimanagement.WorkspaceProductArgs{
			DisplayName:       pulumi.String("Test Template ProductName 4"),
			ProductId:         pulumi.String("testproduct"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			WorkspaceId:       pulumi.String("wks1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
