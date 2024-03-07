package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewContentItem(ctx, "contentItem", &apimanagement.ContentItemArgs{
			ContentItemId: pulumi.String("4e3cf6a5-574a-ba08-1f23-2e7a38faa6d8"),
			ContentTypeId: pulumi.String("page"),
			Properties: pulumi.Any{
				En_us: map[string]interface{}{
					"description": "Short story about the company.",
					"documentId":  "contentTypes/document/contentItems/4e3cf6a5-574a-ba08-1f23-2e7a38faa6d8",
					"keywords":    "company, about",
					"permalink":   "/about",
					"title":       "About",
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
