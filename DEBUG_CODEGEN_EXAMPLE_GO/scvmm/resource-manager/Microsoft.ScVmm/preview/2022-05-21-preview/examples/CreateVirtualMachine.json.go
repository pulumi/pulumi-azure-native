package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/scvmm/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := scvmm.NewVirtualMachine(ctx, "virtualMachine", &scvmm.VirtualMachineArgs{
			CloudId: pulumi.String("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/Clouds/HRCloud"),
			ExtendedLocation: &scvmm.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
				Type: pulumi.String("customLocation"),
			},
			HardwareProfile: &scvmm.HardwareProfileArgs{
				CpuCount: pulumi.Int(4),
				MemoryMB: pulumi.Int(4096),
			},
			Location:           pulumi.String("East US"),
			ResourceGroupName:  pulumi.String("testrg"),
			TemplateId:         pulumi.String("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VirtualMachineTemplates/HRVirtualMachineTemplate"),
			VirtualMachineName: pulumi.String("DemoVM"),
			VmmServerId:        pulumi.String("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
