package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewOpenIdConnectProvider(ctx, "openIdConnectProvider", &apimanagement.OpenIdConnectProviderArgs{
			ClientId:              pulumi.String("oidprovidertemplate3"),
			ClientSecret:          pulumi.String("x"),
			DisplayName:           pulumi.String("templateoidprovider3"),
			MetadataEndpoint:      pulumi.String("https://oidprovider-template3.net"),
			Opid:                  pulumi.String("templateOpenIdConnect3"),
			ResourceGroupName:     pulumi.String("rg1"),
			ServiceName:           pulumi.String("apimService1"),
			UseInApiDocumentation: pulumi.Bool(true),
			UseInTestConsole:      pulumi.Bool(false),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
