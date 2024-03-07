package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkcloud.NewVirtualMachine(ctx, "virtualMachine", &networkcloud.VirtualMachineArgs{
			AdminUsername: pulumi.String("username"),
			BootMethod:    pulumi.String("UEFI"),
			CloudServicesNetworkAttachment: &networkcloud.NetworkAttachmentArgs{
				AttachedNetworkId:  pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName"),
				IpAllocationMethod: pulumi.String("Dynamic"),
			},
			CpuCores: pulumi.Float64(2),
			ExtendedLocation: &networkcloud.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
				Type: pulumi.String("CustomLocation"),
			},
			Location:     pulumi.String("location"),
			MemorySizeGB: pulumi.Float64(8),
			NetworkAttachments: networkcloud.NetworkAttachmentArray{
				&networkcloud.NetworkAttachmentArgs{
					AttachedNetworkId:     pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
					DefaultGateway:        pulumi.String("True"),
					IpAllocationMethod:    pulumi.String("Dynamic"),
					Ipv4Address:           pulumi.String("198.51.100.1"),
					Ipv6Address:           pulumi.String("2001:0db8:0000:0000:0000:0000:0000:0000"),
					NetworkAttachmentName: pulumi.String("netAttachName01"),
				},
			},
			NetworkData: pulumi.String("bmV0d29ya0RhdGVTYW1wbGU="),
			PlacementHints: networkcloud.VirtualMachinePlacementHintArray{
				&networkcloud.VirtualMachinePlacementHintArgs{
					HintType:            pulumi.String("Affinity"),
					ResourceId:          pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName"),
					SchedulingExecution: pulumi.String("Hard"),
					Scope:               pulumi.String(""),
				},
			},
			ResourceGroupName: pulumi.String("resourceGroupName"),
			SshPublicKeys: networkcloud.SshPublicKeyArray{
				&networkcloud.SshPublicKeyArgs{
					KeyData: pulumi.String("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
				},
			},
			StorageProfile: &networkcloud.StorageProfileArgs{
				OsDisk: &networkcloud.OsDiskArgs{
					CreateOption: pulumi.String("Ephemeral"),
					DeleteOption: pulumi.String("Delete"),
					DiskSizeGB:   pulumi.Float64(120),
				},
				VolumeAttachments: pulumi.StringArray{
					pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/volumes/volumeName"),
				},
			},
			Tags: pulumi.StringMap{
				"key1": pulumi.String("myvalue1"),
				"key2": pulumi.String("myvalue2"),
			},
			UserData:           pulumi.String("dXNlckRhdGVTYW1wbGU="),
			VirtualMachineName: pulumi.String("virtualMachineName"),
			VmDeviceModel:      pulumi.String("T2"),
			VmImage:            pulumi.String("myacr.azurecr.io/foobar:latest"),
			VmImageRepositoryCredentials: &networkcloud.ImageRepositoryCredentialsArgs{
				Password:    pulumi.String("{password}"),
				RegistryUrl: pulumi.String("myacr.azurecr.io"),
				Username:    pulumi.String("myuser"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
