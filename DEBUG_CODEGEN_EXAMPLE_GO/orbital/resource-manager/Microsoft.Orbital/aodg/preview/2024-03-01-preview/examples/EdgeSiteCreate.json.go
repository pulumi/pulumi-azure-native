package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/orbital/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := orbital.NewEdgeSite(ctx, "edgeSite", &orbital.EdgeSiteArgs{
			EdgeSiteName: pulumi.String("es1"),
			GlobalCommunicationsSite: &orbital.EdgeSitesPropertiesGlobalCommunicationsSiteArgs{
				Id: pulumi.String("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/providers/Microsoft.Orbital/globalCommunicationsSites/contoso-Vernon"),
			},
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
