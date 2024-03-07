package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
			ApiId:       pulumi.String("echo-api2"),
			Description: pulumi.String("Copy of Existing Echo Api including Operations."),
			DisplayName: pulumi.String("Echo API2"),
			IsCurrent:   pulumi.Bool(true),
			Path:        pulumi.String("echo2"),
			Protocols: pulumi.StringArray{
				pulumi.String("http"),
				pulumi.String("https"),
			},
			ResourceGroupName:    pulumi.String("rg1"),
			ServiceName:          pulumi.String("apimService1"),
			ServiceUrl:           pulumi.String("http://echoapi.cloudapp.net/api"),
			SourceApiId:          pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/58a4aeac497000007d040001"),
			SubscriptionRequired: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
