package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewVirtualEndpoint(ctx, "virtualEndpoint", &dbforpostgresql.VirtualEndpointArgs{
			EndpointType: pulumi.String("ReadWrite"),
			Members: pulumi.StringArray{
				pulumi.String("testPrimary1"),
			},
			ResourceGroupName:   pulumi.String("testrg"),
			ServerName:          pulumi.String("pgtestsvc4"),
			VirtualEndpointName: pulumi.String("pgVirtualEndpoint1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
