package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewLocalUser(ctx, "localUser", &storage.LocalUserArgs{
			AccountName:       pulumi.String("sto2527"),
			HasSharedKey:      pulumi.Bool(false),
			HasSshKey:         pulumi.Bool(false),
			HasSshPassword:    pulumi.Bool(false),
			HomeDirectory:     pulumi.String("homedirectory2"),
			ResourceGroupName: pulumi.String("res6977"),
			Username:          pulumi.String("user1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
