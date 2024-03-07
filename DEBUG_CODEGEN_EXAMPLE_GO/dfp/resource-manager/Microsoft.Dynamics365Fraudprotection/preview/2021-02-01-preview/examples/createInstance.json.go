package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dynamics365fraudprotection/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dynamics365fraudprotection.NewInstanceDetails(ctx, "instanceDetails", &dynamics365fraudprotection.InstanceDetailsArgs{
			Administration: &dynamics365fraudprotection.DFPInstanceAdministratorsArgs{
				Members: pulumi.StringArray{
					pulumi.String("azsdktest@microsoft.com"),
					pulumi.String("azsdktest2@microsoft.com"),
				},
			},
			InstanceName:      pulumi.String("azsdktest"),
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("TestRG"),
			Tags: pulumi.StringMap{
				"testKey": pulumi.String("testValue"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
