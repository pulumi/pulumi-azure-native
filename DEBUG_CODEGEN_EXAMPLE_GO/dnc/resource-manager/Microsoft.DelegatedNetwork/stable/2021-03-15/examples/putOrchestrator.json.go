package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/delegatednetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := delegatednetwork.NewOrchestratorInstanceServiceDetails(ctx, "orchestratorInstanceServiceDetails", &delegatednetwork.OrchestratorInstanceServiceDetailsArgs{
			ApiServerEndpoint: pulumi.String("https://testk8s.cloudapp.net"),
			ClusterRootCA:     pulumi.String("ddsadsad344mfdsfdl"),
			ControllerDetails: &delegatednetwork.ControllerDetailsTypeArgs{
				Id: pulumi.String("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/controller/testcontroller"),
			},
			Identity: &delegatednetwork.OrchestratorIdentityArgs{
				Type: delegatednetwork.ResourceIdentityTypeSystemAssigned,
			},
			Kind:                  pulumi.String("Kubernetes"),
			Location:              pulumi.String("West US"),
			OrchestratorAppId:     pulumi.String("546192d7-503f-477a-9cfe-4efc3ee2b6e1"),
			OrchestratorTenantId:  pulumi.String("da6192d7-503f-477a-9cfe-4efc3ee2b6c3"),
			PrivateLinkResourceId: pulumi.String("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.Network/privateLinkServices/plresource1"),
			ResourceGroupName:     pulumi.String("TestRG"),
			ResourceName:          pulumi.String("testk8s1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
