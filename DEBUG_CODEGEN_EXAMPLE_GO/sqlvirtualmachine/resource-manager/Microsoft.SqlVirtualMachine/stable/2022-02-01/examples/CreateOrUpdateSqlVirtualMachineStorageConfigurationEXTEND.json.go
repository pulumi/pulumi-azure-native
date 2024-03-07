package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sqlvirtualmachine/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sqlvirtualmachine.NewSqlVirtualMachine(ctx, "sqlVirtualMachine", &sqlvirtualmachine.SqlVirtualMachineArgs{
			Location:              pulumi.String("northeurope"),
			ResourceGroupName:     pulumi.String("testrg"),
			SqlVirtualMachineName: pulumi.String("testvm"),
			StorageConfigurationSettings: &sqlvirtualmachine.StorageConfigurationSettingsArgs{
				DiskConfigurationType: pulumi.String("EXTEND"),
				SqlDataSettings: &sqlvirtualmachine.SQLStorageSettingsArgs{
					Luns: pulumi.IntArray{
						pulumi.Int(2),
					},
				},
			},
			VirtualMachineResourceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
