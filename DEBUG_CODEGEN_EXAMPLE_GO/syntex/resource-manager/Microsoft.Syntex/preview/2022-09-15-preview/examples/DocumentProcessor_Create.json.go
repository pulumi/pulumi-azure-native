package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/syntex/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := syntex.NewDocumentProcessor(ctx, "documentProcessor", &syntex.DocumentProcessorArgs{
			Location:      pulumi.String("westus"),
			ProcessorName: pulumi.String("myprocessor"),
			Properties: &syntex.DocumentProcessorPropertiesArgs{
				SpoTenantId:  pulumi.String("e9bb744b-9558-4dc6-9e50-a3297e3332fa"),
				SpoTenantUrl: pulumi.String("https://test123.sharepoint.com"),
			},
			ResourceGroupName: pulumi.String("mygroup"),
			Tags: pulumi.StringMap{
				"additionalProp1": pulumi.String("string1"),
				"additionalProp2": pulumi.String("string2"),
				"additionalProp3": pulumi.String("string3"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
