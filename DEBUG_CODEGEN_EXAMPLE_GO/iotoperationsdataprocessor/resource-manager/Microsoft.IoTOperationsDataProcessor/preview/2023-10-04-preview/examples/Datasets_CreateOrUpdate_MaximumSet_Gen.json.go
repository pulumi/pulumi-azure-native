package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsdataprocessor/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsdataprocessor.NewDataset(ctx, "dataset", &iotoperationsdataprocessor.DatasetArgs{
			DatasetName: pulumi.String("709v7-95-5-t-52oc5--s-5-5876j45wp6mf6--n-8bh--l55-r477"),
			Description: pulumi.String("pakdvhh"),
			ExtendedLocation: &iotoperationsdataprocessor.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/e0aaa3df-e9a4-456a-9824-3c3b5c438110/resourceGroups/IoTOperationsDataProcessor-rg/providers/Microsoft.ExtendedLocation/customLocations/dev-space"),
				Type: pulumi.String("CustomLocation"),
			},
			InstanceName:      pulumi.String("xh--6h732-2-6-21-4513-2-597q-5412971q"),
			Keys:              nil,
			Location:          pulumi.String("westus"),
			Payload:           pulumi.String(".value"),
			ResourceGroupName: pulumi.String("rgopenapi"),
			Tags:              nil,
			Timestamp:         pulumi.String(".timestamp"),
			Ttl:               pulumi.String("72h"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
