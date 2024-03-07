package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkcloud.NewBareMetalMachineKeySet(ctx, "bareMetalMachineKeySet", &networkcloud.BareMetalMachineKeySetArgs{
			AzureGroupId:               pulumi.String("f110271b-XXXX-4163-9b99-214d91660f0e"),
			BareMetalMachineKeySetName: pulumi.String("bareMetalMachineKeySetName"),
			ClusterName:                pulumi.String("clusterName"),
			Expiration:                 pulumi.String("2022-12-31T23:59:59.008Z"),
			ExtendedLocation: &networkcloud.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
				Type: pulumi.String("CustomLocation"),
			},
			JumpHostsAllowed: pulumi.StringArray{
				pulumi.String("192.0.2.1"),
				pulumi.String("192.0.2.5"),
			},
			Location:          pulumi.String("location"),
			OsGroupName:       pulumi.String("standardAccessGroup"),
			PrivilegeLevel:    pulumi.String("Standard"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("myvalue1"),
				"key2": pulumi.String("myvalue2"),
			},
			UserList: []networkcloud.KeySetUserArgs{
				{
					AzureUserName: pulumi.String("userABC"),
					Description:   pulumi.String("Needs access for troubleshooting as a part of the support team"),
					SshPublicKey: {
						KeyData: pulumi.String("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
					},
				},
				{
					AzureUserName: pulumi.String("userXYZ"),
					Description:   pulumi.String("Needs access for troubleshooting as a part of the support team"),
					SshPublicKey: {
						KeyData: pulumi.String("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
