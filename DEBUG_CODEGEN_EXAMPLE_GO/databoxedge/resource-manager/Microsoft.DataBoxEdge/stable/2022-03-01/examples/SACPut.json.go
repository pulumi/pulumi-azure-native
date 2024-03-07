package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databoxedge/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databoxedge.NewStorageAccountCredential(ctx, "storageAccountCredential", &databoxedge.StorageAccountCredentialArgs{
			AccountKey: &databoxedge.AsymmetricEncryptedSecretArgs{
				EncryptionAlgorithm:      pulumi.String("AES256"),
				EncryptionCertThumbprint: pulumi.String("2A9D8D6BE51574B5461230AEF02F162C5F01AD31"),
				Value:                    pulumi.String("lAeZEYi6rNP1/EyNaVUYmTSZEYyaIaWmwUsGwek0+xiZj54GM9Ue9/UA2ed/ClC03wuSit2XzM/cLRU5eYiFBwks23rGwiQOr3sruEL2a74EjPD050xYjA6M1I2hu/w2yjVHhn5j+DbXS4Xzi+rHHNZK3DgfDO3PkbECjPck+PbpSBjy9+6Mrjcld5DIZhUAeMlMHrFlg+WKRKB14o/og56u5/xX6WKlrMLEQ+y6E18dUwvWs2elTNoVO8PBE8SM/CfooX4AMNvaNdSObNBPdP+F6Lzc556nFNWXrBLRt0vC7s9qTiVRO4x/qCNaK/B4y7IqXMllwQFf4Np9UQ2ECA=="),
			},
			AccountType:       pulumi.String("BlobStorage"),
			Alias:             pulumi.String("sac1"),
			DeviceName:        pulumi.String("testedgedevice"),
			Name:              pulumi.String("sac1"),
			ResourceGroupName: pulumi.String("GroupForEdgeAutomation"),
			SslStatus:         pulumi.String("Disabled"),
			UserName:          pulumi.String("cisbvt"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
