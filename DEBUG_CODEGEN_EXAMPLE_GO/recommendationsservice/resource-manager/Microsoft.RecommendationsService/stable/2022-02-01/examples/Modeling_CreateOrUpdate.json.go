package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recommendationsservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recommendationsservice.NewModeling(ctx, "modeling", &recommendationsservice.ModelingArgs{
			AccountName:  pulumi.String("sampleAccount"),
			Location:     pulumi.String("West US"),
			ModelingName: pulumi.String("c1"),
			Properties: recommendationsservice.ModelingResourceResponseProperties{
				Features:  pulumi.String("Standard"),
				Frequency: pulumi.String("High"),
				InputData: &recommendationsservice.ModelingInputDataArgs{
					ConnectionString: pulumi.String("https://storageAccount.blob.core.windows.net/container/root"),
				},
				Size: pulumi.String("Medium"),
			},
			ResourceGroupName: pulumi.String("rg"),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Prod"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
