package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewPartnerConfiguration(ctx, "partnerConfiguration", &eventgrid.PartnerConfigurationArgs{
			PartnerAuthorization: eventgrid.PartnerAuthorizationResponse{
				AuthorizedPartnersList: eventgrid.PartnerArray{
					&eventgrid.PartnerArgs{
						AuthorizationExpirationTimeInUtc: pulumi.String("2022-01-28T01:20:55.142Z"),
						PartnerName:                      pulumi.String("Contoso.Finance"),
						PartnerRegistrationImmutableId:   pulumi.String("941892bc-f5d0-4d1c-8fb5-477570fc2b71"),
					},
					&eventgrid.PartnerArgs{
						AuthorizationExpirationTimeInUtc: pulumi.String("2022-02-20T01:00:00.142Z"),
						PartnerName:                      pulumi.String("fabrikam.HR"),
						PartnerRegistrationImmutableId:   pulumi.String("5362bdb6-ce3e-4d0d-9a5b-3eb92c8aab38"),
					},
				},
				DefaultMaximumExpirationTimeInDays: pulumi.Int(10),
			},
			ResourceGroupName: pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
