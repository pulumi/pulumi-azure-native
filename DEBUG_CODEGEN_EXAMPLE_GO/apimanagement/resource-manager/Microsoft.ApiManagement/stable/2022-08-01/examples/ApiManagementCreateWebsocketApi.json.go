package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
			ApiId:       pulumi.String("tempgroup"),
			ApiType:     pulumi.String("websocket"),
			Description: pulumi.String("apidescription5200"),
			DisplayName: pulumi.String("apiname1463"),
			Path:        pulumi.String("newapiPath"),
			Protocols: pulumi.StringArray{
				pulumi.String("wss"),
				pulumi.String("ws"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			ServiceUrl:        pulumi.String("wss://echo.websocket.org"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
