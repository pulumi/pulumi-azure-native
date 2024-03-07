package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApiVersionSet(ctx, "apiVersionSet", &apimanagement.ApiVersionSetArgs{
			Description:       pulumi.String("Version configuration"),
			DisplayName:       pulumi.String("api set 1"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			VersionSetId:      pulumi.String("api1"),
			VersioningScheme:  pulumi.String("Segment"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
