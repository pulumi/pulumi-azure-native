package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azuresphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azuresphere.NewImage(ctx, "image", &azuresphere.ImageArgs{
			CatalogName:       pulumi.String("MyCatalog1"),
			Image:             pulumi.String("bXliYXNlNjRzdHJpbmc="),
			ImageName:         pulumi.String(".default"),
			ResourceGroupName: pulumi.String("MyResourceGroup1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
