package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
			ApiId:                  pulumi.String("echo-api;rev=3"),
			ApiRevisionDescription: pulumi.String("Creating a Revision of an existing API"),
			Path:                   pulumi.String("echo"),
			ResourceGroupName:      pulumi.String("rg1"),
			ServiceName:            pulumi.String("apimService1"),
			ServiceUrl:             pulumi.String("http://echoapi.cloudapp.net/apiv3"),
			SourceApiId:            pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echo-api"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
