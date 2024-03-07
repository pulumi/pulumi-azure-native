package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/solutions/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := solutions.NewApplicationDefinition(ctx, "applicationDefinition", &solutions.ApplicationDefinitionArgs{
			ApplicationDefinitionName: pulumi.String("myManagedApplicationDef"),
			Authorizations: []solutions.ApplicationAuthorizationArgs{
				{
					PrincipalId:      pulumi.String("validprincipalguid"),
					RoleDefinitionId: pulumi.String("validroleguid"),
				},
			},
			Description:       pulumi.String("myManagedApplicationDef description"),
			DisplayName:       pulumi.String("myManagedApplicationDef"),
			LockLevel:         solutions.ApplicationLockLevelNone,
			PackageFileUri:    pulumi.String("https://path/to/packagezipfile"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
