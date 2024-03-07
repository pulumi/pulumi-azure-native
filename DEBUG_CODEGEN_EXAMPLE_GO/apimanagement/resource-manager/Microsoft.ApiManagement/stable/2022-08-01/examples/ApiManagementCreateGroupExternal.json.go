package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewGroup(ctx, "group", &apimanagement.GroupArgs{
			Description:       pulumi.String("new group to test"),
			DisplayName:       pulumi.String("NewGroup (samiraad.onmicrosoft.com)"),
			ExternalId:        pulumi.String("aad://samiraad.onmicrosoft.com/groups/83cf2753-5831-4675-bc0e-2f8dc067c58d"),
			GroupId:           pulumi.String("aadGroup"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Type:              apimanagement.GroupTypeExternal,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
