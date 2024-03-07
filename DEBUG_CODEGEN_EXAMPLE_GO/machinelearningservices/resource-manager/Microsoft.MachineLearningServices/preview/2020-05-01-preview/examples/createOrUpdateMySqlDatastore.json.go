package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewMachineLearningDatastore(ctx, "machineLearningDatastore", &machinelearningservices.MachineLearningDatastoreArgs{
			DataStoreType:     pulumi.String("mysqldb"),
			DatabaseName:      pulumi.String("dataset"),
			DatastoreName:     pulumi.String("mySqlDatastore"),
			Password:          pulumi.String("<password>"),
			ResourceGroupName: pulumi.String("acjain-mleastUS2"),
			ServerName:        pulumi.String("dataset-mysql-srv"),
			UserId:            pulumi.String("demo_user@dataset-mysql-srv"),
			WorkspaceName:     pulumi.String("acjain-mleastUS2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
