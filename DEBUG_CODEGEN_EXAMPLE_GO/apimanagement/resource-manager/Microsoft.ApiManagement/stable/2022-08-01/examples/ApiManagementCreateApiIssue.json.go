package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiIssue(ctx, "apiIssue", &apimanagement.ApiIssueArgs{
			ApiId:             pulumi.String("57d1f7558aa04f15146d9d8a"),
			CreatedDate:       pulumi.String("2018-02-01T22:21:20.467Z"),
			Description:       pulumi.String("New API issue description"),
			IssueId:           pulumi.String("57d2ef278aa04f0ad01d6cdc"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			State:             pulumi.String("open"),
			Title:             pulumi.String("New API issue"),
			UserId:            pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
