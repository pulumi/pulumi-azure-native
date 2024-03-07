package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewCustomApi(ctx, "customApi", &web.CustomApiArgs{
			ApiName: pulumi.String("testCustomApi"),
			Properties: web.CustomApiPropertiesDefinitionResponse{
				ApiDefinitions: &web.ApiResourceDefinitionsArgs{
					OriginalSwaggerUrl: pulumi.String("https://tempuri.org/swagger.json"),
				},
				ApiType:      pulumi.String("Rest"),
				Capabilities: pulumi.StringArray{},
				Description:  pulumi.String(""),
				DisplayName:  pulumi.String("testCustomApi"),
				IconUri:      pulumi.String("/testIcon.svg"),
			},
			ResourceGroupName: pulumi.String("testResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
