package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
			ApiId:             pulumi.String("apidocs"),
			Format:            pulumi.String("swagger-link"),
			Path:              pulumi.String("petstoreapi123"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			ServiceUrl:        pulumi.String("http://petstore.swagger.wordnik.com/api"),
			Value:             pulumi.String("http://apimpimportviaurl.azurewebsites.net/api/apidocs/"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
