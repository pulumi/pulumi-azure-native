package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewEventHub(ctx, "eventHub", &eventhub.EventHubArgs{
			CaptureDescription: &eventhub.CaptureDescriptionArgs{
				Destination: &eventhub.DestinationArgs{
					ArchiveNameFormat:        pulumi.String("{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}"),
					BlobContainer:            pulumi.String("container"),
					Name:                     pulumi.String("EventHubArchive.AzureBlockBlob"),
					StorageAccountResourceId: pulumi.String("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage"),
				},
				Enabled:           pulumi.Bool(true),
				Encoding:          eventhub.EncodingCaptureDescriptionAvro,
				IntervalInSeconds: pulumi.Int(120),
				SizeLimitInBytes:  pulumi.Int(10485763),
			},
			EventHubName:           pulumi.String("sdk-EventHub-6547"),
			MessageRetentionInDays: pulumi.Float64(4),
			NamespaceName:          pulumi.String("sdk-Namespace-5357"),
			PartitionCount:         pulumi.Float64(4),
			ResourceGroupName:      pulumi.String("Default-NotificationHubs-AustraliaEast"),
			RetentionDescription: &eventhub.RetentionDescriptionArgs{
				CleanupPolicy:                 pulumi.String("Compact"),
				RetentionTimeInHours:          pulumi.Float64(96),
				TombstoneRetentionTimeInHours: pulumi.Int(1),
			},
			Status: eventhub.EntityStatusActive,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
