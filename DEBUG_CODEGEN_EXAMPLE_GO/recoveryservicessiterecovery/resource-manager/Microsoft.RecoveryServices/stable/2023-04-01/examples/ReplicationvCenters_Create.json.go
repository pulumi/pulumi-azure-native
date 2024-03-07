package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewReplicationvCenter(ctx, "replicationvCenter", &recoveryservices.ReplicationvCenterArgs{
			FabricName: pulumi.String("MadhaviFabric"),
			Properties: &recoveryservices.AddVCenterRequestPropertiesArgs{
				FriendlyName:    pulumi.String("esx-78"),
				IpAddress:       pulumi.String("inmtest78"),
				Port:            pulumi.String("443"),
				ProcessServerId: pulumi.String("5A720CAB-39CB-F445-BD1662B0B33164B5"),
				RunAsAccountId:  pulumi.String("2"),
			},
			ResourceGroupName: pulumi.String("MadhaviVRG"),
			ResourceName:      pulumi.String("MadhaviVault"),
			VcenterName:       pulumi.String("esx-78"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
