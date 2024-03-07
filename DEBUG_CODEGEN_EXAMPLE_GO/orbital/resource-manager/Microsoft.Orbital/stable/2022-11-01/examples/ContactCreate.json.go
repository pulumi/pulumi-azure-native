package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/orbital/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := orbital.NewContact(ctx, "contact", &orbital.ContactArgs{
			ContactName: pulumi.String("contact1"),
			ContactProfile: &orbital.ContactsPropertiesContactProfileArgs{
				Id: pulumi.String("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/contactProfiles/CONTOSO-CP"),
			},
			GroundStationName:    pulumi.String("EASTUS2_0"),
			ReservationEndTime:   pulumi.String("2023-02-22T11:10:45Z"),
			ReservationStartTime: pulumi.String("2023-02-22T10:58:30Z"),
			ResourceGroupName:    pulumi.String("contoso-Rgp"),
			SpacecraftName:       pulumi.String("CONTOSO_SAT"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
