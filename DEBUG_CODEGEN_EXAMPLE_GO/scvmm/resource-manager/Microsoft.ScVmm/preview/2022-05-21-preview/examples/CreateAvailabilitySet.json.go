package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/scvmm/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := scvmm.NewAvailabilitySet(ctx, "availabilitySet", &scvmm.AvailabilitySetArgs{
			AvailabilitySetName: pulumi.String("hr-avset"),
			ExtendedLocation: &scvmm.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
				Type: pulumi.String("customLocation"),
			},
			Location:          pulumi.String("East US"),
			ResourceGroupName: pulumi.String("testrg"),
			VmmServerId:       pulumi.String("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/VMMServers/ContosoVMMServer"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
