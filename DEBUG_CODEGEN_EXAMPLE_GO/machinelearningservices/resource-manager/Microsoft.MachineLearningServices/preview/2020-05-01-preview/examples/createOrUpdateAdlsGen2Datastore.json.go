package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewMachineLearningDatastore(ctx, "machineLearningDatastore", &machinelearningservices.MachineLearningDatastoreArgs{
			AccountName:       pulumi.String("nicksadlsgen2storage"),
			ClientId:          pulumi.String("233d7008-b157-4354-88d1-ba191f06a900"),
			ClientSecret:      pulumi.String("vdegbvedgeg"),
			DataStoreType:     pulumi.String("adls-gen2"),
			DatastoreName:     pulumi.String("adlsgen2Datastore"),
			FileSystem:        pulumi.String("testfs1"),
			ResourceGroupName: pulumi.String("acjain-mleastUS2"),
			TenantId:          pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
			WorkspaceName:     pulumi.String("acjain-mleastUS2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
