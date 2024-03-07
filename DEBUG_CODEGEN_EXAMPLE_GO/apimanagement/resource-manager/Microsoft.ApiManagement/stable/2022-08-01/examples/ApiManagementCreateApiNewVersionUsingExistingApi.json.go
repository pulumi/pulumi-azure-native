package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
			ApiId:           pulumi.String("echoapiv3"),
			ApiVersion:      pulumi.String("v4"),
			ApiVersionSetId: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apiVersionSets/aa9c59e6-c0cd-4258-9356-9ca7d2f0b458"),
			Description:     pulumi.String("Create Echo API into a new Version using Existing Version Set and Copy all Operations."),
			DisplayName:     pulumi.String("Echo API2"),
			IsCurrent:       pulumi.Bool(true),
			Path:            pulumi.String("echo2"),
			Protocols: pulumi.StringArray{
				pulumi.String("http"),
				pulumi.String("https"),
			},
			ResourceGroupName:    pulumi.String("rg1"),
			ServiceName:          pulumi.String("apimService1"),
			ServiceUrl:           pulumi.String("http://echoapi.cloudapp.net/api"),
			SourceApiId:          pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echoPath"),
			SubscriptionRequired: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
