package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewCacheRule(ctx, "cacheRule", &containerregistry.CacheRuleArgs{
			CacheRuleName:           pulumi.String("myCacheRule"),
			CredentialSetResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/credentialSets/myCredentialSet"),
			RegistryName:            pulumi.String("myRegistry"),
			ResourceGroupName:       pulumi.String("myResourceGroup"),
			SourceRepository:        pulumi.String("docker.io/library/hello-world"),
			TargetRepository:        pulumi.String("cached-docker-hub/hello-world"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
