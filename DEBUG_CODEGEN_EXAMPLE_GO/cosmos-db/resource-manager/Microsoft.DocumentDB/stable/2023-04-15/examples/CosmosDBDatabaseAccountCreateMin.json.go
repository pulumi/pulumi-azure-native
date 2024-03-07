package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewDatabaseAccount(ctx, "databaseAccount", &documentdb.DatabaseAccountArgs{
			AccountName:              pulumi.String("ddb1"),
			CreateMode:               pulumi.String("Default"),
			DatabaseAccountOfferType: documentdb.DatabaseAccountOfferTypeStandard,
			Location:                 pulumi.String("westus"),
			Locations: documentdb.LocationArray{
				&documentdb.LocationArgs{
					FailoverPriority: pulumi.Int(0),
					IsZoneRedundant:  pulumi.Bool(false),
					LocationName:     pulumi.String("southcentralus"),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
