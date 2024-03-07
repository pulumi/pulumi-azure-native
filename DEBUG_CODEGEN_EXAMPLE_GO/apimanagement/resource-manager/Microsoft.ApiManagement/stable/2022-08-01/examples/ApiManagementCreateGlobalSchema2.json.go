package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewGlobalSchema(ctx, "globalSchema", &apimanagement.GlobalSchemaArgs{
			Description:       pulumi.String("sample schema description"),
			ResourceGroupName: pulumi.String("rg1"),
			SchemaId:          pulumi.String("schema1"),
			SchemaType:        pulumi.String("json"),
			ServiceName:       pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
