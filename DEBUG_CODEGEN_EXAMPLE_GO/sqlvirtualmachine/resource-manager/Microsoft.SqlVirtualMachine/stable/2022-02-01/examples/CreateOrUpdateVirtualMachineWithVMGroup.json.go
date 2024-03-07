package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sqlvirtualmachine/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sqlvirtualmachine.NewSqlVirtualMachine(ctx, "sqlVirtualMachine", &sqlvirtualmachine.SqlVirtualMachineArgs{
			Location:                         pulumi.String("northeurope"),
			ResourceGroupName:                pulumi.String("testrg"),
			SqlVirtualMachineGroupResourceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/testvmgroup"),
			SqlVirtualMachineName:            pulumi.String("testvm"),
			VirtualMachineResourceId:         pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm2"),
			WsfcDomainCredentials: &sqlvirtualmachine.WsfcDomainCredentialsArgs{
				ClusterBootstrapAccountPassword: pulumi.String("<Password>"),
				ClusterOperatorAccountPassword:  pulumi.String("<Password>"),
				SqlServiceAccountPassword:       pulumi.String("<Password>"),
			},
			WsfcStaticIp: pulumi.String("10.0.0.7"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
