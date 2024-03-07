package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
			ApiId:             pulumi.String("petstore"),
			Format:            pulumi.String("wadl-link-json"),
			Path:              pulumi.String("collector"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Value:             pulumi.String("https://developer.cisco.com/media/wae-release-6-2-api-reference/wae-collector-rest-api/application.wadl"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
