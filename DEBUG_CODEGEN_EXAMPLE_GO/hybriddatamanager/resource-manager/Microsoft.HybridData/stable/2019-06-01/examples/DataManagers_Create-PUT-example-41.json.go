package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybriddata/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybriddata.NewDataManager(ctx, "dataManager", &hybriddata.DataManagerArgs{
			DataManagerName:   pulumi.String("TestAzureSDKOperations"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("ResourceGroupForSDKTest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
