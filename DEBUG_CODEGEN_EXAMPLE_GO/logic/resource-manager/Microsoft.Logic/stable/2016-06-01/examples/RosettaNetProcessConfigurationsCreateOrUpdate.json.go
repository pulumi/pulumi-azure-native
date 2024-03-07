package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logic.NewRosettaNetProcessConfiguration(ctx, "rosettaNetProcessConfiguration", &logic.RosettaNetProcessConfigurationArgs{
			ActivitySettings: &logic.RosettaNetPipActivitySettingsArgs{
				AcknowledgmentOfReceiptSettings: &logic.RosettaNetPipAcknowledgmentOfReceiptSettingsArgs{
					IsNonRepudiationRequired:   pulumi.Bool(false),
					TimeToAcknowledgeInSeconds: pulumi.Int(600),
				},
				ActivityBehavior: &logic.RosettaNetPipActivityBehaviorArgs{
					ActionType:                       logic.RosettaNetActionTypeDoubleAction,
					IsAuthorizationRequired:          pulumi.Bool(false),
					IsSecuredTransportRequired:       pulumi.Bool(false),
					NonRepudiationOfOriginAndContent: pulumi.Bool(false),
					PersistentConfidentialityScope:   logic.RosettaNetPipConfidentialityScopeNone,
					ResponseType:                     logic.RosettaNetResponseTypeAsync,
					RetryCount:                       pulumi.Int(2),
					TimeToPerformInSeconds:           pulumi.Int(7200),
				},
				ActivityType: logic.RosettaNetPipActivityTypeRequestResponse,
			},
			Description: pulumi.String("Test description"),
			InitiatorRoleSettings: &logic.RosettaNetPipRoleSettingsArgs{
				Action: pulumi.String("Purchase Order Request"),
				BusinessDocument: &logic.RosettaNetPipBusinessDocumentArgs{
					Description: pulumi.String("A request to accept a purchase order for fulfillment.."),
					Name:        pulumi.String("Purchase Order Request"),
					Version:     pulumi.String("V02.02.00"),
				},
				Description:           pulumi.String("This partner role creates a demand for a product or service."),
				Role:                  pulumi.String("Buyer"),
				RoleType:              logic.RosettaNetPipRoleTypeFunctional,
				Service:               pulumi.String("Buyer Service"),
				ServiceClassification: pulumi.String("Business Service"),
			},
			IntegrationAccountName: pulumi.String("testia123"),
			ProcessCode:            pulumi.String("3A4"),
			ProcessName:            pulumi.String("Request Purchase Order"),
			ProcessVersion:         pulumi.String("V02.02.00"),
			ResourceGroupName:      pulumi.String("testrg123"),
			ResponderRoleSettings: &logic.RosettaNetPipRoleSettingsArgs{
				Action: pulumi.String("Purchase Order Confirmation Action"),
				BusinessDocument: &logic.RosettaNetPipBusinessDocumentArgs{
					Description: pulumi.String("Formally confirms the status of line item(s) in a Purchase Order. A Purchase Order line item may have one of the following states: accepted, rejected, or pending."),
					Name:        pulumi.String("Purchase Order Confirmation"),
					Version:     pulumi.String("V02.02.00"),
				},
				Description:           pulumi.String("An organization that sells products to partners in the supply chain."),
				Role:                  pulumi.String("Seller"),
				RoleType:              logic.RosettaNetPipRoleTypeOrganizational,
				Service:               pulumi.String("Seller Service"),
				ServiceClassification: pulumi.String("Business Service"),
			},
			RosettaNetProcessConfigurationName: pulumi.String("3A4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
