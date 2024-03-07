package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customerinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerinsights.NewConnectorMapping(ctx, "connectorMapping", &customerinsights.ConnectorMappingArgs{
			ConnectorName:  pulumi.String("testConnector8858"),
			Description:    pulumi.String("Test mapping"),
			DisplayName:    pulumi.String("testMapping12491"),
			EntityType:     customerinsights.EntityTypesInteraction,
			EntityTypeName: pulumi.String("TestInteractionType2967"),
			HubName:        pulumi.String("sdkTestHub"),
			MappingName:    pulumi.String("testMapping12491"),
			MappingProperties: &customerinsights.ConnectorMappingPropertiesArgs{
				Availability: &customerinsights.ConnectorMappingAvailabilityArgs{
					Frequency: customerinsights.FrequencyTypesHour,
					Interval:  pulumi.Int(5),
				},
				CompleteOperation: &customerinsights.ConnectorMappingCompleteOperationArgs{
					CompletionOperationType: customerinsights.CompletionOperationTypesDeleteFile,
					DestinationFolder:       pulumi.String("fakePath"),
				},
				ErrorManagement: &customerinsights.ConnectorMappingErrorManagementArgs{
					ErrorLimit:          pulumi.Int(10),
					ErrorManagementType: customerinsights.ErrorManagementTypesStopImport,
				},
				FileFilter: pulumi.String("unknown"),
				FolderPath: pulumi.String("http://sample.dne/file"),
				Format: &customerinsights.ConnectorMappingFormatArgs{
					ColumnDelimiter: pulumi.String("|"),
					FormatType:      customerinsights.FormatTypesTextFormat,
				},
				HasHeader: pulumi.Bool(false),
				Structure: customerinsights.ConnectorMappingStructureArray{
					&customerinsights.ConnectorMappingStructureArgs{
						ColumnName:   pulumi.String("unknown1"),
						IsEncrypted:  pulumi.Bool(false),
						PropertyName: pulumi.String("unknwon1"),
					},
					&customerinsights.ConnectorMappingStructureArgs{
						ColumnName:   pulumi.String("unknown2"),
						IsEncrypted:  pulumi.Bool(true),
						PropertyName: pulumi.String("unknwon2"),
					},
				},
			},
			ResourceGroupName: pulumi.String("TestHubRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
