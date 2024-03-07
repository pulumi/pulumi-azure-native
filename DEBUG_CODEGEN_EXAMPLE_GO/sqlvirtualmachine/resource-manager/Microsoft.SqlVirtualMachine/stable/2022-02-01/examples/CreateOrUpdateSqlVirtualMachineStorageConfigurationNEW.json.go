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
			StorageConfigurationSettings: sqlvirtualmachine.StorageConfigurationSettingsResponse{
				DiskConfigurationType: pulumi.String("NEW"),
				SqlDataSettings: &sqlvirtualmachine.SQLStorageSettingsArgs{
					DefaultFilePath: pulumi.String("F:\\folderpath\\"),
					Luns: pulumi.IntArray{
						pulumi.Int(0),
					},
				},
				SqlLogSettings: &sqlvirtualmachine.SQLStorageSettingsArgs{
					DefaultFilePath: pulumi.String("G:\\folderpath\\"),
					Luns: pulumi.IntArray{
						pulumi.Int(1),
					},
				},
				SqlSystemDbOnDataDisk: pulumi.Bool(true),
				SqlTempDbSettings: &sqlvirtualmachine.SQLTempDbSettingsArgs{
					DataFileCount:   pulumi.Int(8),
					DataFileSize:    pulumi.Int(256),
					DataGrowth:      pulumi.Int(512),
					DefaultFilePath: pulumi.String("D:\\TEMP"),
					LogFileSize:     pulumi.Int(256),
					LogGrowth:       pulumi.Int(512),
				},
				StorageWorkloadType: pulumi.String("OLTP"),
			},
			VirtualMachineResourceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
