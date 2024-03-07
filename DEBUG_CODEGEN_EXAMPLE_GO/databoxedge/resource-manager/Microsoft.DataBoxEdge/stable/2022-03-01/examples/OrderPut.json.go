package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databoxedge/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databoxedge.NewOrder(ctx, "order", &databoxedge.OrderArgs{
			ContactInformation: &databoxedge.ContactDetailsArgs{
				CompanyName:   pulumi.String("Microsoft"),
				ContactPerson: pulumi.String("John Mcclane"),
				EmailList: pulumi.StringArray{
					pulumi.String("john@microsoft.com"),
				},
				Phone: pulumi.String("(800) 426-9400"),
			},
			DeviceName:        pulumi.String("testedgedevice"),
			ResourceGroupName: pulumi.String("GroupForEdgeAutomation"),
			ShippingAddress: &databoxedge.AddressArgs{
				AddressLine1: pulumi.String("Microsoft Corporation"),
				AddressLine2: pulumi.String("One Microsoft Way"),
				AddressLine3: pulumi.String("Redmond"),
				City:         pulumi.String("WA"),
				Country:      pulumi.String("USA"),
				PostalCode:   pulumi.String("98052"),
				State:        pulumi.String("WA"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
