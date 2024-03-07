package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
			ApiId:                                   pulumi.String("petstore"),
			Format:                                  pulumi.String("openapi-link"),
			Path:                                    pulumi.String("petstore"),
			ResourceGroupName:                       pulumi.String("rg1"),
			ServiceName:                             pulumi.String("apimService1"),
			TranslateRequiredQueryParametersConduct: pulumi.String("template"),
			Value:                                   pulumi.String("https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/examples/v3.0/petstore.yaml"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
