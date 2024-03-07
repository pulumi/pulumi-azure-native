package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewModernizeProject(ctx, "modernizeProject", &migrate.ModernizeProjectArgs{
			Identity: migrate.ResourceIdentityResponse{
				PrincipalId: pulumi.String("ins"),
				TenantId:    pulumi.String("fjnu"),
				Type:        pulumi.String("None"),
				UserAssignedIdentities: migrate.UserAssignedIdentityMap{
					"key6848": &migrate.UserAssignedIdentityArgs{
						ClientId:    pulumi.String("lvlngepacjdjryqmxuvfdxwtkc"),
						PrincipalId: pulumi.String("lumkynazsspljxiiwvz"),
					},
				},
			},
			Location:             pulumi.String("nbqyuxrgrlhx"),
			ModernizeProjectName: pulumi.String("b"),
			Properties: migrate.ModernizeProjectModelPropertiesResponse{
				MigrationConfiguration: &migrate.MigrationConfigurationArgs{
					KeyVaultResourceId:          pulumi.String("vekhittkyogvwnqmggknv"),
					MigrationSolutionResourceId: pulumi.String("bglfkwtzvqmhwpddwpvtdzaleaioxo"),
					StorageAccountResourceId:    pulumi.String("dgcoticysafrpynyoxkgrspooiia"),
				},
			},
			ResourceGroupName: pulumi.String("rgmigrateEngine"),
			Tags: pulumi.StringMap{
				"key8644": pulumi.String("wfyi"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
