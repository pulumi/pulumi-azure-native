package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewSshPublicKey(ctx, "sshPublicKey", &compute.SshPublicKeyArgs{
			Location:          pulumi.String("westus"),
			PublicKey:         pulumi.String("{ssh-rsa public key}"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SshPublicKeyName:  pulumi.String("mySshPublicKeyName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
