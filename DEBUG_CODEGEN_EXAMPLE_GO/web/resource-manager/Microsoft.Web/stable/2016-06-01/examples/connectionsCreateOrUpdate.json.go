package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewConnection(ctx, "connection", &web.ConnectionArgs{
			ConnectionName: pulumi.String("testManagedApi"),
			Properties: &web.ApiConnectionDefinitionPropertiesArgs{
				Api: &web.ApiReferenceArgs{
					Id: pulumi.String("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/providers/Microsoft.Web/locations/centralus/managedApis/testManagedApi"),
				},
				CustomParameterValues: nil,
				DisplayName:           pulumi.String("testManagedApi"),
				ParameterValues:       nil,
			},
			ResourceGroupName: pulumi.String("testResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
