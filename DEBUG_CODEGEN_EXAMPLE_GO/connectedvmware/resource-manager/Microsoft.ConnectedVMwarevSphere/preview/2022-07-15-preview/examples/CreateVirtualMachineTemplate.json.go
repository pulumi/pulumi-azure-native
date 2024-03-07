package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/connectedvmwarevsphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connectedvmwarevsphere.NewVirtualMachineTemplate(ctx, "virtualMachineTemplate", &connectedvmwarevsphere.VirtualMachineTemplateArgs{
			ExtendedLocation: &connectedvmwarevsphere.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
				Type: pulumi.String("customLocation"),
			},
			Location:                   pulumi.String("East US"),
			MoRefId:                    pulumi.String("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
			ResourceGroupName:          pulumi.String("testrg"),
			VCenterId:                  pulumi.String("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter"),
			VirtualMachineTemplateName: pulumi.String("WebFrontEndTemplate"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
