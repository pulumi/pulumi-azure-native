package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewLocalUser(ctx, "localUser", &storage.LocalUserArgs{
			AccountName:    pulumi.String("sto2527"),
			HasSshPassword: pulumi.Bool(true),
			HomeDirectory:  pulumi.String("homedirectory"),
			PermissionScopes: storage.PermissionScopeArray{
				&storage.PermissionScopeArgs{
					Permissions:  pulumi.String("rwd"),
					ResourceName: pulumi.String("share1"),
					Service:      pulumi.String("file"),
				},
				&storage.PermissionScopeArgs{
					Permissions:  pulumi.String("rw"),
					ResourceName: pulumi.String("share2"),
					Service:      pulumi.String("file"),
				},
			},
			ResourceGroupName: pulumi.String("res6977"),
			SshAuthorizedKeys: storage.SshPublicKeyArray{
				&storage.SshPublicKeyArgs{
					Description: pulumi.String("key name"),
					Key:         pulumi.String("ssh-rsa keykeykeykeykey="),
				},
			},
			Username: pulumi.String("user1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
