package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datafactory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datafactory.NewChangeDataCapture(ctx, "changeDataCapture", &datafactory.ChangeDataCaptureArgs{
			AllowVNetOverride:     pulumi.Bool(false),
			ChangeDataCaptureName: pulumi.String("exampleChangeDataCapture"),
			Description:           pulumi.String("Sample demo change data capture to transfer data from delimited (csv) to Azure SQL Database with automapped and non-automapped mappings."),
			FactoryName:           pulumi.String("exampleFactoryName"),
			ResourceGroupName:     pulumi.String("exampleResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
