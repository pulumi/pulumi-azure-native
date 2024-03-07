package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsdataprocessor/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsdataprocessor.NewInstance(ctx, "instance", &iotoperationsdataprocessor.InstanceArgs{
			Description: pulumi.String("ytazzjdwgnnwsmexqasgpyabrtkgtf"),
			ExtendedLocation: &iotoperationsdataprocessor.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/e0aaa3df-e9a4-456a-9824-3c3b5c438110/resourceGroups/IoTOperationsDataProcessor-rg/providers/Microsoft.ExtendedLocation/customLocations/dev-space"),
				Type: pulumi.String("CustomLocation"),
			},
			InstanceName:      pulumi.String("15wp-47-e60s18w"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("rgopenapi"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
