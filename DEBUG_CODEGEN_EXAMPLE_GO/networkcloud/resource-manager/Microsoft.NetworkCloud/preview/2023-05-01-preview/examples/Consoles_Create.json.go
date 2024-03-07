package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkcloud.NewConsole(ctx, "console", &networkcloud.ConsoleArgs{
			ConsoleName: pulumi.String("default"),
			Enabled:     pulumi.String("True"),
			Expiration:  pulumi.String("2022-06-01T01:27:03.008Z"),
			ExtendedLocation: &networkcloud.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterManagerExtendedLocationName"),
				Type: pulumi.String("CustomLocation"),
			},
			Location:          pulumi.String("location"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			SshPublicKey: &networkcloud.SshPublicKeyArgs{
				KeyData: pulumi.String("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
			},
			Tags: pulumi.StringMap{
				"key1": pulumi.String("myvalue1"),
				"key2": pulumi.String("myvalue2"),
			},
			VirtualMachineName: pulumi.String("virtualMachineName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
