package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datamigration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datamigration.NewDatabaseMigrationsMongoToCosmosDbvCoreMongo(ctx, "databaseMigrationsMongoToCosmosDbvCoreMongo", &datamigration.DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs{
			CollectionList: datamigration.MongoMigrationCollectionArray{
				&datamigration.MongoMigrationCollectionArgs{
					SourceCollection: pulumi.String("sourceCol1"),
					SourceDatabase:   pulumi.String("sourceDb1"),
					TargetCollection: pulumi.String("targetCol1"),
					TargetDatabase:   pulumi.String("targetDb1"),
				},
				&datamigration.MongoMigrationCollectionArgs{
					SourceCollection: pulumi.String("sourceCol2"),
					SourceDatabase:   pulumi.String("sourceDb2"),
				},
			},
			Kind:              pulumi.String("MongoToCosmosDbMongo"),
			MigrationName:     pulumi.String("migrationRequest"),
			MigrationService:  pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/MigrationServices/testMigrationService"),
			ResourceGroupName: pulumi.String("testrg"),
			Scope:             pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName"),
			SourceMongoConnection: &datamigration.MongoConnectionInformationArgs{
				Host:     pulumi.String("abc.mongodb.com"),
				Password: pulumi.String("placeholder"),
				Port:     pulumi.Int(88),
				UseSsl:   pulumi.Bool(true),
				UserName: pulumi.String("abc"),
			},
			TargetMongoConnection: &datamigration.MongoConnectionInformationArgs{
				ConnectionString: pulumi.String("placeholder"),
			},
			TargetResourceName: pulumi.String("targetCosmosDbClusterName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
