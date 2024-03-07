package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewArchife(ctx, "archife", &containerregistry.ArchifeArgs{
			ArchiveName: pulumi.String("myArchiveName"),
			PackageSource: &containerregistry.ArchivePackageSourcePropertiesArgs{
				Type: pulumi.String("remote"),
				Url:  pulumi.String("string"),
			},
			PackageType:              pulumi.String("rpm"),
			PublishedVersion:         pulumi.String("string"),
			RegistryName:             pulumi.String("myRegistry"),
			RepositoryEndpointPrefix: pulumi.String("string"),
			ResourceGroupName:        pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
