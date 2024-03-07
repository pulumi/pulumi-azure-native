package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/integrationspaces/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := integrationspaces.NewBusinessProcess(ctx, "businessProcess", &integrationspaces.BusinessProcessArgs{
			ApplicationName: pulumi.String("Application1"),
			BusinessProcessMapping: integrationspaces.BusinessProcessMappingItemMap{
				"Completed": &integrationspaces.BusinessProcessMappingItemArgs{
					LogicAppResourceId: pulumi.String("subscriptions/sub1/resourcegroups/group1/providers/Microsoft.Web/sites/logicApp1"),
					OperationName:      pulumi.String("CompletedPO"),
					OperationType:      pulumi.String("Action"),
					WorkflowName:       pulumi.String("Fulfillment"),
				},
				"Denied": &integrationspaces.BusinessProcessMappingItemArgs{
					LogicAppResourceId: pulumi.String("subscriptions/sub1/resourcegroups/group1/providers/Microsoft.Web/sites/logicApp1"),
					OperationName:      pulumi.String("DeniedPO"),
					OperationType:      pulumi.String("Action"),
					WorkflowName:       pulumi.String("Fulfillment"),
				},
				"Processing": &integrationspaces.BusinessProcessMappingItemArgs{
					LogicAppResourceId: pulumi.String("subscriptions/sub1/resourcegroups/group1/providers/Microsoft.Web/sites/logicApp1"),
					OperationName:      pulumi.String("ApprovedPO"),
					OperationType:      pulumi.String("Action"),
					WorkflowName:       pulumi.String("PurchaseOrder"),
				},
				"Received": &integrationspaces.BusinessProcessMappingItemArgs{
					LogicAppResourceId: pulumi.String("subscriptions/sub1/resourcegroups/group1/providers/Microsoft.Web/sites/logicApp1"),
					OperationName:      pulumi.String("manual"),
					OperationType:      pulumi.String("Trigger"),
					WorkflowName:       pulumi.String("PurchaseOrder"),
				},
				"Shipped": &integrationspaces.BusinessProcessMappingItemArgs{
					LogicAppResourceId: pulumi.String("subscriptions/sub1/resourcegroups/group1/providers/Microsoft.Web/sites/logicApp1"),
					OperationName:      pulumi.String("ShippedPO"),
					OperationType:      pulumi.String("Action"),
					WorkflowName:       pulumi.String("Fulfillment"),
				},
			},
			BusinessProcessName: pulumi.String("BusinessProcess1"),
			BusinessProcessStages: integrationspaces.BusinessProcessStageMap{
				"Completed": &integrationspaces.BusinessProcessStageArgs{
					Description: pulumi.String("Completed"),
					StagesBefore: pulumi.StringArray{
						pulumi.String("Shipped"),
					},
				},
				"Denied": &integrationspaces.BusinessProcessStageArgs{
					Description: pulumi.String("Denied"),
					StagesBefore: pulumi.StringArray{
						pulumi.String("Processing"),
					},
				},
				"Processing": &integrationspaces.BusinessProcessStageArgs{
					Description: pulumi.String("Processing"),
					Properties: pulumi.StringMap{
						"ApprovalState": pulumi.String("String"),
						"ApproverName":  pulumi.String("String"),
						"POAmount":      pulumi.String("Integer"),
					},
					StagesBefore: pulumi.StringArray{
						pulumi.String("Received"),
					},
				},
				"Received": &integrationspaces.BusinessProcessStageArgs{
					Description: pulumi.String("received"),
					Properties: pulumi.StringMap{
						"City":     pulumi.String("String"),
						"Product":  pulumi.String("String"),
						"Quantity": pulumi.String("Integer"),
						"State":    pulumi.String("String"),
					},
				},
				"Shipped": &integrationspaces.BusinessProcessStageArgs{
					Description: pulumi.String("Shipped"),
					Properties: pulumi.StringMap{
						"ShipPriority": pulumi.String("Integer"),
						"TrackingID":   pulumi.String("Integer"),
					},
					StagesBefore: pulumi.StringArray{
						pulumi.String("Denied"),
					},
				},
			},
			Description: pulumi.String("First Business Process"),
			Identifier: &integrationspaces.BusinessProcessIdentifierArgs{
				PropertyName: pulumi.String("businessIdentifier-1"),
				PropertyType: pulumi.String("String"),
			},
			ResourceGroupName:              pulumi.String("testrg"),
			SpaceName:                      pulumi.String("Space1"),
			TableName:                      pulumi.String("table1"),
			TrackingDataStoreReferenceName: pulumi.String("trackingDataStoreReferenceName1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
