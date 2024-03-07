package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewMachineLearningDatastore(ctx, "machineLearningDatastore", &machinelearningservices.MachineLearningDatastoreArgs{
			DataStoreType:     pulumi.String("psqldb"),
			DatabaseName:      pulumi.String("dataset"),
			DatastoreName:     pulumi.String("postgreSqlDatastore"),
			Password:          pulumi.String("<password>"),
			ResourceGroupName: pulumi.String("acjain-mleastUS2"),
			ServerName:        pulumi.String("dataset-pgsql-srv"),
			UserId:            pulumi.String("demo_user@dataset-pgsql-srv"),
			WorkspaceName:     pulumi.String("acjain-mleastUS2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
