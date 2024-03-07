package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkfunction/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfunction.NewCollectorPolicy(ctx, "collectorPolicy", &networkfunction.CollectorPolicyArgs{
			AzureTrafficCollectorName: pulumi.String("atc"),
			CollectorPolicyName:       pulumi.String("cp1"),
			EmissionPolicies: networkfunction.EmissionPoliciesPropertiesFormatArray{
				&networkfunction.EmissionPoliciesPropertiesFormatArgs{
					EmissionDestinations: networkfunction.EmissionPolicyDestinationArray{
						&networkfunction.EmissionPolicyDestinationArgs{
							DestinationType: pulumi.String("AzureMonitor"),
						},
					},
					EmissionType: pulumi.String("IPFIX"),
				},
			},
			IngestionPolicy: &networkfunction.IngestionPolicyPropertiesFormatArgs{
				IngestionSources: networkfunction.IngestionSourcesPropertiesFormatArray{
					&networkfunction.IngestionSourcesPropertiesFormatArgs{
						ResourceId: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName"),
						SourceType: pulumi.String("Resource"),
					},
				},
				IngestionType: pulumi.String("IPFIX"),
			},
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
