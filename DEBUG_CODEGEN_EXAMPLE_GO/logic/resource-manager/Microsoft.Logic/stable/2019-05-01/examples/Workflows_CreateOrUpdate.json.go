package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logic.NewWorkflow(ctx, "workflow", &logic.WorkflowArgs{
			Definition: pulumi.Any{
				Schema: "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
				Actions: map[string]interface{}{
					"Find_pet_by_ID": map[string]interface{}{
						"inputs": map[string]interface{}{
							"host": map[string]interface{}{
								"connection": map[string]interface{}{
									"name": "@parameters('$connections')['test-custom-connector']['connectionId']",
								},
							},
							"method": "get",
							"path":   "/pet/@{encodeURIComponent('1')}",
						},
						"runAfter": nil,
						"type":     "ApiConnection",
					},
				},
				ContentVersion: "1.0.0.0",
				Outputs:        nil,
				Parameters: map[string]interface{}{
					"$connections": map[string]interface{}{
						"defaultValue": nil,
						"type":         "Object",
					},
				},
				Triggers: map[string]interface{}{
					"manual": map[string]interface{}{
						"inputs": map[string]interface{}{
							"schema": nil,
						},
						"kind": "Http",
						"type": "Request",
					},
				},
			},
			IntegrationAccount: &logic.ResourceReferenceArgs{
				Id: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Logic/integrationAccounts/test-integration-account"),
			},
			Location: pulumi.String("brazilsouth"),
			Parameters: logic.WorkflowParameterMap{
				"$connections": &logic.WorkflowParameterArgs{
					Value: pulumi.Any{
						TestCustomConnector: map[string]interface{}{
							"connectionId":   "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Web/connections/test-custom-connector",
							"connectionName": "test-custom-connector",
							"id":             "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/locations/brazilsouth/managedApis/test-custom-connector",
						},
					},
				},
			},
			ResourceGroupName: pulumi.String("test-resource-group"),
			Tags:              nil,
			WorkflowName:      pulumi.String("test-workflow"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
