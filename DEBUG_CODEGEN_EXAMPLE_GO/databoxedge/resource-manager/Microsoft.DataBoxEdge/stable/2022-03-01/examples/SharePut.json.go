package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databoxedge/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databoxedge.NewShare(ctx, "share", &databoxedge.ShareArgs{
			AccessProtocol: pulumi.String("SMB"),
			AzureContainerInfo: &databoxedge.AzureContainerInfoArgs{
				ContainerName:              pulumi.String("testContainerSMB"),
				DataFormat:                 pulumi.String("BlockBlob"),
				StorageAccountCredentialId: pulumi.String("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/sac1"),
			},
			DataPolicy:        pulumi.String("Cloud"),
			Description:       pulumi.String(""),
			DeviceName:        pulumi.String("testedgedevice"),
			MonitoringStatus:  pulumi.String("Enabled"),
			Name:              pulumi.String("smbshare"),
			ResourceGroupName: pulumi.String("GroupForEdgeAutomation"),
			ShareStatus:       pulumi.String("Online"),
			UserAccessRights: []databoxedge.UserAccessRightArgs{
				{
					AccessType: pulumi.String("Change"),
					UserId:     pulumi.String("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/users/user2"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
