package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkanalytics.NewDataType(ctx, "dataType", &networkanalytics.DataTypeArgs{
			DataProductName:        pulumi.String("dataproduct01"),
			DataTypeName:           pulumi.String("datatypename"),
			DatabaseCacheRetention: pulumi.Int(23),
			DatabaseRetention:      pulumi.Int(6),
			ResourceGroupName:      pulumi.String("aoiresourceGroupName"),
			State:                  pulumi.String("STARTED"),
			StorageOutputRetention: pulumi.Int(27),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
