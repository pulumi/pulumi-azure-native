package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databoxedge/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databoxedge.NewUser(ctx, "user", &databoxedge.UserArgs{
			DeviceName: pulumi.String("testedgedevice"),
			EncryptedPassword: &databoxedge.AsymmetricEncryptedSecretArgs{
				EncryptionAlgorithm:      pulumi.String("None"),
				EncryptionCertThumbprint: pulumi.String("blah"),
				Value:                    pulumi.String("<value>"),
			},
			Name:              pulumi.String("user1"),
			ResourceGroupName: pulumi.String("GroupForEdgeAutomation"),
			UserType:          pulumi.String("Share"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
