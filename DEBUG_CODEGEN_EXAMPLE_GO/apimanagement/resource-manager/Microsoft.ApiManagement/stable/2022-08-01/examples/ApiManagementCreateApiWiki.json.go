package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiWiki(ctx, "apiWiki", &apimanagement.ApiWikiArgs{
			ApiId: pulumi.String("57d1f7558aa04f15146d9d8a"),
			Documents: apimanagement.WikiDocumentationContractArray{
				&apimanagement.WikiDocumentationContractArgs{
					DocumentationId: pulumi.String("docId1"),
				},
				&apimanagement.WikiDocumentationContractArgs{
					DocumentationId: pulumi.String("docId2"),
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
