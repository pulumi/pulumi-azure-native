package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewScopeMap(ctx, "scopeMap", &containerregistry.ScopeMapArgs{
			Actions: pulumi.StringArray{
				pulumi.String("repositories/myrepository/contentWrite"),
				pulumi.String("repositories/myrepository/delete"),
			},
			Description:       pulumi.String("Developer Scopes"),
			RegistryName:      pulumi.String("myRegistry"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ScopeMapName:      pulumi.String("myScopeMap"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
