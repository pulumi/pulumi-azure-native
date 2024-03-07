package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewArchiveVersion(ctx, "archiveVersion", &containerregistry.ArchiveVersionArgs{
			ArchiveName:        pulumi.String("myArchiveName"),
			ArchiveVersionName: pulumi.String("myArchiveVersionName"),
			PackageType:        pulumi.String("rpm"),
			RegistryName:       pulumi.String("myRegistry"),
			ResourceGroupName:  pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
