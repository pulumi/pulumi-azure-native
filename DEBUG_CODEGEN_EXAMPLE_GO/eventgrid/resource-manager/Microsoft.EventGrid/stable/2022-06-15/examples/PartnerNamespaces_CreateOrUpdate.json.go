package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventgrid/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventgrid.NewPartnerNamespace(ctx, "partnerNamespace", &eventgrid.PartnerNamespaceArgs{
			Location:                            pulumi.String("westus"),
			PartnerNamespaceName:                pulumi.String("examplePartnerNamespaceName1"),
			PartnerRegistrationFullyQualifiedId: pulumi.String("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerRegistrations/ContosoCorpAccount1"),
			ResourceGroupName:                   pulumi.String("examplerg"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
