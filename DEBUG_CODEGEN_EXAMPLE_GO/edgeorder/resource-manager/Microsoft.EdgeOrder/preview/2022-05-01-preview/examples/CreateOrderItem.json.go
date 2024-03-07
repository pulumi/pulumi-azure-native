package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/edgeorder/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := edgeorder.NewOrderItem(ctx, "orderItem", &edgeorder.OrderItemArgs{
			AddressDetails: &edgeorder.AddressDetailsArgs{
				ForwardAddress: &edgeorder.AddressPropertiesArgs{
					ContactDetails: &edgeorder.ContactDetailsArgs{
						ContactName: pulumi.String("XXXX XXXX"),
						EmailList: pulumi.StringArray{
							pulumi.String("xxxx@xxxx.xxx"),
						},
						Phone:          pulumi.String("0000000000"),
						PhoneExtension: pulumi.String(""),
					},
					ShippingAddress: &edgeorder.ShippingAddressArgs{
						AddressType:     pulumi.String("None"),
						City:            pulumi.String("San Francisco"),
						CompanyName:     pulumi.String("Microsoft"),
						Country:         pulumi.String("US"),
						PostalCode:      pulumi.String("94107"),
						StateOrProvince: pulumi.String("CA"),
						StreetAddress1:  pulumi.String("16 TOWNSEND ST"),
						StreetAddress2:  pulumi.String("UNIT 1"),
					},
				},
			},
			Location: pulumi.String("eastus"),
			OrderId:  pulumi.String("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/locations/eastus/orders/TestOrderName2"),
			OrderItemDetails: &edgeorder.OrderItemDetailsArgs{
				OrderItemType: pulumi.String("Purchase"),
				Preferences: &edgeorder.PreferencesArgs{
					TransportPreferences: &edgeorder.TransportPreferencesArgs{
						PreferredShipmentType: pulumi.String("MicrosoftManaged"),
					},
				},
				ProductDetails: &edgeorder.ProductDetailsArgs{
					HierarchyInformation: &edgeorder.HierarchyInformationArgs{
						ConfigurationName: pulumi.String("edgep_base"),
						ProductFamilyName: pulumi.String("azurestackedge"),
						ProductLineName:   pulumi.String("azurestackedge"),
						ProductName:       pulumi.String("azurestackedgegpu"),
					},
				},
			},
			OrderItemName:     pulumi.String("TestOrderItemName2"),
			ResourceGroupName: pulumi.String("YourResourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
