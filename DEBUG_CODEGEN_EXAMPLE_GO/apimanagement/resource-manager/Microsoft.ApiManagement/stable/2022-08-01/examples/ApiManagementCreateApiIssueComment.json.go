package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiIssueComment(ctx, "apiIssueComment", &apimanagement.ApiIssueCommentArgs{
			ApiId:             pulumi.String("57d1f7558aa04f15146d9d8a"),
			CommentId:         pulumi.String("599e29ab193c3c0bd0b3e2fb"),
			CreatedDate:       pulumi.String("2018-02-01T22:21:20.467Z"),
			IssueId:           pulumi.String("57d2ef278aa04f0ad01d6cdc"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Text:              pulumi.String("Issue comment."),
			UserId:            pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
