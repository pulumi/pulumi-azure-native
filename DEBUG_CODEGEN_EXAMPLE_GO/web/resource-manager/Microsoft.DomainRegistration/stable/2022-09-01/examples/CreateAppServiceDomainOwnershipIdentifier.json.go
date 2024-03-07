package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/domainregistration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := domainregistration.NewDomainOwnershipIdentifier(ctx, "domainOwnershipIdentifier", &domainregistration.DomainOwnershipIdentifierArgs{
			DomainName:        pulumi.String("example.com"),
			Name:              pulumi.String("SampleOwnershipId"),
			OwnershipId:       pulumi.String("SampleOwnershipId"),
			ResourceGroupName: pulumi.String("testrg123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
