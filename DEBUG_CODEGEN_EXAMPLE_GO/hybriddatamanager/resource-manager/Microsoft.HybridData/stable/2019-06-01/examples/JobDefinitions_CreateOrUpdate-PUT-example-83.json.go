package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybriddata/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybriddata.NewJobDefinition(ctx, "jobDefinition", &hybriddata.JobDefinitionArgs{
			DataManagerName: pulumi.String("TestAzureSDKOperations"),
			DataServiceInput: pulumi.Any{
				AzureStorageType: "Blob",
				BackupChoice:     "UseExistingLatest",
				ContainerName:    "containerfromtest",
				DeviceName:       "8600-SHG0997877L71FC",
				FileNameFilter:   "*",
				IsDirectoryMode:  false,
				RootDirectories: []string{
					"\\",
				},
				VolumeNames: []string{
					"TestAutomation",
				},
			},
			DataServiceName:   pulumi.String("DataTransformation"),
			DataSinkId:        pulumi.String("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStores/TestAzureStorage1"),
			DataSourceId:      pulumi.String("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStores/TestStorSimpleSource1"),
			JobDefinitionName: pulumi.String("jobdeffromtestcode1"),
			ResourceGroupName: pulumi.String("ResourceGroupForSDKTest"),
			RunLocation:       hybriddata.RunLocationWestus,
			State:             hybriddata.StateEnabled,
			UserConfirmation:  hybriddata.UserConfirmationRequired,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
