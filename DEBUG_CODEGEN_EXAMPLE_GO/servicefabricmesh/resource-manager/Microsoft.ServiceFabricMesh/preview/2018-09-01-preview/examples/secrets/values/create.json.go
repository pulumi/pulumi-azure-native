package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabricmesh/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabricmesh.NewSecretValue(ctx, "secretValue", &servicefabricmesh.SecretValueArgs{
			Location:                pulumi.String("West US"),
			ResourceGroupName:       pulumi.String("sbz_demo"),
			SecretResourceName:      pulumi.String("dbConnectionString"),
			SecretValueResourceName: pulumi.String("v1"),
			Value:                   pulumi.String("mongodb://contoso123:0Fc3IolnL12312asdfawejunASDF@asdfYXX2t8a97kghVcUzcDv98hawelufhawefafnoQRGwNj2nMPL1Y9qsIr9Srdw==@contoso123.documents.azure.com:10255/mydatabase?ssl=true"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
