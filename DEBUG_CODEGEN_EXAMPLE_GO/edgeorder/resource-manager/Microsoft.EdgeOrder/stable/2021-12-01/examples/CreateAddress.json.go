package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/edgeorder/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := edgeorder.NewAddressByName(ctx, "addressByName", &edgeorder.AddressByNameArgs{
			AddressName: pulumi.String("TestAddressName2"),
			ContactDetails: &edgeorder.ContactDetailsArgs{
				ContactName: pulumi.String("XXXX XXXX"),
				EmailList: pulumi.StringArray{
					pulumi.String("xxxx@xxxx.xxx"),
				},
				Phone:          pulumi.String("0000000000"),
				PhoneExtension: pulumi.String(""),
			},
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("YourResourceGroupName"),
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
		})
		if err != nil {
			return err
		}
		return nil
	})
}
