package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewContentType(ctx, "contentType", &apimanagement.ContentTypeArgs{
			ContentTypeId:     pulumi.String("page"),
			Description:       pulumi.String("A regular page"),
			Name:              pulumi.String("Page"),
			ResourceGroupName: pulumi.String("rg1"),
			Schema: pulumi.Any{
				AdditionalProperties: false,
				Properties: map[string]interface{}{
					"en_us": map[string]interface{}{
						"additionalProperties": false,
						"properties": map[string]interface{}{
							"description": map[string]interface{}{
								"description": "Page description. This property gets included in SEO attributes.",
								"indexed":     true,
								"title":       "Description",
								"type":        "string",
							},
							"documentId": map[string]interface{}{
								"description": "Reference to page content document.",
								"title":       "Document ID",
								"type":        "string",
							},
							"keywords": map[string]interface{}{
								"description": "Page keywords. This property gets included in SEO attributes.",
								"indexed":     true,
								"title":       "Keywords",
								"type":        "string",
							},
							"permalink": map[string]interface{}{
								"description": "Page permalink, e.g. '/about'.",
								"indexed":     true,
								"title":       "Permalink",
								"type":        "string",
							},
							"title": map[string]interface{}{
								"description": "Page title. This property gets included in SEO attributes.",
								"indexed":     true,
								"title":       "Title",
								"type":        "string",
							},
						},
						"required": []string{
							"title",
							"permalink",
							"documentId",
						},
						"type": "object",
					},
				},
			},
			ServiceName: pulumi.String("apimService1"),
			Version:     pulumi.String("1.0.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
