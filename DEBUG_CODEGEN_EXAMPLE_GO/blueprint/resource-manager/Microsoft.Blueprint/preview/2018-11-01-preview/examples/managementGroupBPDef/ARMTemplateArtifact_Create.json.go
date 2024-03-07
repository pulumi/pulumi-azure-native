package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/blueprint/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := blueprint.NewTemplateArtifact(ctx, "templateArtifact", &blueprint.TemplateArtifactArgs{
			ArtifactName:  pulumi.String("storageTemplate"),
			BlueprintName: pulumi.String("simpleBlueprint"),
			Kind:          pulumi.String("template"),
			Parameters: blueprint.ParameterValueMap{
				"storageAccountType": &blueprint.ParameterValueArgs{
					Value: pulumi.Any("[parameters('storageAccountType')]"),
				},
			},
			ResourceGroup: pulumi.String("storageRG"),
			ResourceScope: pulumi.String("providers/Microsoft.Management/managementGroups/ContosoOnlineGroup"),
			Template: pulumi.Any{
				ContentVersion: "1.0.0.0",
				Outputs: map[string]interface{}{
					"storageAccountName": map[string]interface{}{
						"type":  "string",
						"value": "[variables('storageAccountName')]",
					},
				},
				Parameters: map[string]interface{}{
					"storageAccountType": map[string]interface{}{
						"allowedValues": []string{
							"Standard_LRS",
							"Standard_GRS",
							"Standard_ZRS",
							"Premium_LRS",
						},
						"defaultValue": "Standard_LRS",
						"metadata": map[string]interface{}{
							"description": "Storage Account type",
						},
						"type": "string",
					},
				},
				Resources: []map[string]interface{}{
					map[string]interface{}{
						"apiVersion": "2016-01-01",
						"kind":       "Storage",
						"location":   "[resourceGroup().location]",
						"name":       "[variables('storageAccountName')]",
						"properties": nil,
						"sku": map[string]interface{}{
							"name": "[parameters('storageAccountType')]",
						},
						"type": "Microsoft.Storage/storageAccounts",
					},
				},
				Variables: map[string]interface{}{
					"storageAccountName": "[concat(uniquestring(resourceGroup().id), 'standardsa')]",
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
