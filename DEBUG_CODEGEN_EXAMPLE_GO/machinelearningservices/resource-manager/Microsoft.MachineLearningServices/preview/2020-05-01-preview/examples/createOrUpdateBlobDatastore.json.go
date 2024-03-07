package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewMachineLearningDatastore(ctx, "machineLearningDatastore", &machinelearningservices.MachineLearningDatastoreArgs{
			AccountKey:        pulumi.String("wddrfewfewsgewgrrwegwreg"),
			AccountName:       pulumi.String("acjainmleastus9484093746"),
			ContainerName:     pulumi.String("azureml-blobstore-5da947c5-53aa-41a5-bb2b-074074e73b7"),
			DataStoreType:     pulumi.String("blob"),
			DatastoreName:     pulumi.String("blobDatastore"),
			ResourceGroupName: pulumi.String("acjain-mleastUS2"),
			WorkspaceName:     pulumi.String("acjain-mleastUS2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
