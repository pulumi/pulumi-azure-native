package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewPartnerDestination(ctx, "partnerDestination", &eventgrid.PartnerDestinationArgs{
			EndpointBaseUrl:                 pulumi.String("https://www.example/endpoint"),
			EndpointServiceContext:          pulumi.String("This is an example"),
			ExpirationTimeIfNotActivatedUtc: pulumi.String("2022-03-14T19:33:43.430Z"),
			Location:                        pulumi.String("westus2"),
			MessageForActivation:            pulumi.String("Sample Activation message"),
			PartnerDestinationName:          pulumi.String("examplePartnerDestinationName1"),
			PartnerRegistrationImmutableId:  pulumi.String("0bd70ee2-7d95-447e-ab1f-c4f320019404"),
			ResourceGroupName:               pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
