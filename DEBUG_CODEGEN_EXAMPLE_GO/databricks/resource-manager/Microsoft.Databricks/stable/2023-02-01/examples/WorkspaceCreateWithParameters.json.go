package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databricks/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databricks.NewWorkspace(ctx, "workspace", &databricks.WorkspaceArgs{
			Location:               pulumi.String("westus"),
			ManagedResourceGroupId: pulumi.String("/subscriptions/subid/resourceGroups/myManagedRG"),
			Parameters: databricks.WorkspaceCustomParametersResponse{
				CustomPrivateSubnetName: &databricks.WorkspaceCustomStringParameterArgs{
					Value: pulumi.String("myPrivateSubnet"),
				},
				CustomPublicSubnetName: &databricks.WorkspaceCustomStringParameterArgs{
					Value: pulumi.String("myPublicSubnet"),
				},
				CustomVirtualNetworkId: &databricks.WorkspaceCustomStringParameterArgs{
					Value: pulumi.String("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/myNetwork"),
				},
			},
			ResourceGroupName: pulumi.String("rg"),
			WorkspaceName:     pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
