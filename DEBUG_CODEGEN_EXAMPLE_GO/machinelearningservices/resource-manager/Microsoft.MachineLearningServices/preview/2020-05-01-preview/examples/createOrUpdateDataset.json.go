package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewMachineLearningDataset(ctx, "machineLearningDataset", &machinelearningservices.MachineLearningDatasetArgs{
			DatasetName: pulumi.String("datasetName123"),
			DatasetType: pulumi.String("file"),
			Parameters: &machinelearningservices.DatasetCreateRequestParametersArgs{
				Path: &machinelearningservices.DatasetCreateRequestPathArgs{
					DataPath: &machinelearningservices.DatasetCreateRequestDataPathArgs{
						DatastoreName: pulumi.String("testblobfromarm"),
						RelativePath:  pulumi.String("UI/03-26-2020_083359_UTC/latin1encoding.csv"),
					},
				},
			},
			Registration: &machinelearningservices.DatasetCreateRequestRegistrationArgs{
				Description: pulumi.String("test description"),
				Name:        pulumi.String("datasetName123"),
			},
			ResourceGroupName: pulumi.String("acjain-mleastUS2"),
			SkipValidation:    pulumi.Bool(false),
			WorkspaceName:     pulumi.String("acjain-mleastUS2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
