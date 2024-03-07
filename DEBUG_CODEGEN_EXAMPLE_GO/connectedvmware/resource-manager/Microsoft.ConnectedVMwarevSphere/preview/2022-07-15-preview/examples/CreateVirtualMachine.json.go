package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/connectedvmwarevsphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connectedvmwarevsphere.NewVirtualMachine(ctx, "virtualMachine", &connectedvmwarevsphere.VirtualMachineArgs{
			ExtendedLocation: &connectedvmwarevsphere.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
				Type: pulumi.String("customLocation"),
			},
			HardwareProfile: &connectedvmwarevsphere.HardwareProfileArgs{
				MemorySizeMB: pulumi.Int(4196),
				NumCPUs:      pulumi.Int(4),
			},
			Location:           pulumi.String("East US"),
			ResourceGroupName:  pulumi.String("testrg"),
			ResourcePoolId:     pulumi.String("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool"),
			TemplateId:         pulumi.String("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachineTemplates/WebFrontEndTemplate"),
			VCenterId:          pulumi.String("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter"),
			VirtualMachineName: pulumi.String("DemoVM"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
