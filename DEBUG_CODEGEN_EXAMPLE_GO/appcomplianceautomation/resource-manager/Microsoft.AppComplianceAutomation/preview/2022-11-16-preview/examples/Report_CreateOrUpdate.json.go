package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appcomplianceautomation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appcomplianceautomation.NewReport(ctx, "report", &appcomplianceautomation.ReportArgs{
			Properties: &appcomplianceautomation.ReportPropertiesArgs{
				OfferGuid: pulumi.String("0000"),
				Resources: appcomplianceautomation.ResourceMetadataArray{
					&appcomplianceautomation.ResourceMetadataArgs{
						ResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
						Tags: pulumi.StringMap{
							"key1": pulumi.String("value1"),
						},
					},
				},
				TimeZone:    pulumi.String("GMT Standard Time"),
				TriggerTime: pulumi.String("2022-03-04T05:11:56.197Z"),
			},
			ReportName: pulumi.String("testReportName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
