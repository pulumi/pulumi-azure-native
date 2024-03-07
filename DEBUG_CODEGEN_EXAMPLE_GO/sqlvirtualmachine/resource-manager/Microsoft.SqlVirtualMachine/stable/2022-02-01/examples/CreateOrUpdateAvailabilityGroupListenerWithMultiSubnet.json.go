package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sqlvirtualmachine/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sqlvirtualmachine.NewAvailabilityGroupListener(ctx, "availabilityGroupListener", &sqlvirtualmachine.AvailabilityGroupListenerArgs{
			AvailabilityGroupListenerName: pulumi.String("agl-test"),
			AvailabilityGroupName:         pulumi.String("ag-test"),
			MultiSubnetIpConfigurations: []sqlvirtualmachine.MultiSubnetIpConfigurationArgs{
				{
					PrivateIpAddress: {
						IpAddress:        pulumi.String("10.0.0.112"),
						SubnetResourceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
					},
					SqlVirtualMachineInstance: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2"),
				},
				{
					PrivateIpAddress: {
						IpAddress:        pulumi.String("10.0.1.112"),
						SubnetResourceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/alternate"),
					},
					SqlVirtualMachineInstance: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm1"),
				},
			},
			Port:                       pulumi.Int(1433),
			ResourceGroupName:          pulumi.String("testrg"),
			SqlVirtualMachineGroupName: pulumi.String("testvmgroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
