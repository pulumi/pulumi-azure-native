package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
			ApiId:       pulumi.String("tempgroup"),
			ApiType:     pulumi.String("graphql"),
			Description: pulumi.String("apidescription5200"),
			DisplayName: pulumi.String("apiname1463"),
			Path:        pulumi.String("graphql-api"),
			Protocols: pulumi.StringArray{
				pulumi.String("http"),
				pulumi.String("https"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			ServiceUrl:        pulumi.String("https://api.spacex.land/graphql"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
