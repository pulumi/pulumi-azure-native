package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
			ApiId:             pulumi.String("petstore"),
			Format:            pulumi.String("swagger-link-json"),
			Path:              pulumi.String("petstore"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Value:             pulumi.String("http://petstore.swagger.io/v2/swagger.json"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
