package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewSAPDatabaseInstance(ctx, "sapDatabaseInstance", &workloads.SAPDatabaseInstanceArgs{
			DatabaseInstanceName:   pulumi.String("databaseServer"),
			Location:               pulumi.String("westcentralus"),
			ResourceGroupName:      pulumi.String("test-rg"),
			SapVirtualInstanceName: pulumi.String("X00"),
			Tags:                   nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
