package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewPolicyDefinitionVersion(ctx, "policyDefinitionVersion", &authorization.PolicyDefinitionVersionArgs{
			Description: pulumi.String("Force resource names to begin with given 'prefix' and/or end with given 'suffix'"),
			DisplayName: pulumi.String("Enforce resource naming convention"),
			Metadata: pulumi.Any{
				Category: "Naming",
			},
			Mode: pulumi.String("All"),
			Parameters: authorization.ParameterDefinitionsValueMap{
				"prefix": &authorization.ParameterDefinitionsValueArgs{
					Metadata: &authorization.ParameterDefinitionsValueMetadataArgs{
						Description: pulumi.String("Resource name prefix"),
						DisplayName: pulumi.String("Prefix"),
					},
					Type: pulumi.String("String"),
				},
				"suffix": &authorization.ParameterDefinitionsValueArgs{
					Metadata: &authorization.ParameterDefinitionsValueMetadataArgs{
						Description: pulumi.String("Resource name suffix"),
						DisplayName: pulumi.String("Suffix"),
					},
					Type: pulumi.String("String"),
				},
			},
			PolicyDefinitionName:    pulumi.String("ResourceNaming"),
			PolicyDefinitionVersion: pulumi.String("1.2.1"),
			PolicyRule: pulumi.Any{
				If: map[string]interface{}{
					"not": map[string]interface{}{
						"field": "name",
						"like":  "[concat(parameters('prefix'), '*', parameters('suffix'))]",
					},
				},
				Then: map[string]interface{}{
					"effect": "deny",
				},
			},
			Version: pulumi.String("1.2.1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
