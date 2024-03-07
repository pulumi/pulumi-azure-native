package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/healthcareapis/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := healthcareapis.NewDicomService(ctx, "dicomService", &healthcareapis.DicomServiceArgs{
			DicomServiceName:  pulumi.String("blue"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("testRG"),
			WorkspaceName:     pulumi.String("workspace1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
