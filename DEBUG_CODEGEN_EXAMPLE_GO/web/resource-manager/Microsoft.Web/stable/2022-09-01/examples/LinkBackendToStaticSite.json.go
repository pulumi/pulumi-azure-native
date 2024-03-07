package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewStaticSiteLinkedBackend(ctx, "staticSiteLinkedBackend", &web.StaticSiteLinkedBackendArgs{
			BackendResourceId: pulumi.String("/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/backendRg/providers/Microsoft.Web/sites/testBackend"),
			LinkedBackendName: pulumi.String("testBackend"),
			Name:              pulumi.String("testStaticSite0"),
			Region:            pulumi.String("West US 2"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
