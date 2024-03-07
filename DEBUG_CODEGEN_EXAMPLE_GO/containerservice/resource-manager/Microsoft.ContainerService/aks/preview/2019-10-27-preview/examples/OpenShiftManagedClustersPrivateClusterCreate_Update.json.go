package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewOpenShiftManagedCluster(ctx, "openShiftManagedCluster", &containerservice.OpenShiftManagedClusterArgs{
			AgentPoolProfiles: []containerservice.OpenShiftManagedClusterAgentPoolProfileArgs{
				{
					Count:      pulumi.Int(3),
					Name:       pulumi.String("infra"),
					OsType:     pulumi.String("Linux"),
					Role:       pulumi.String("infra"),
					SubnetCidr: pulumi.String("10.0.0.0/24"),
					VmSize:     pulumi.String("Standard_D4s_v3"),
				},
				{
					Count:      pulumi.Int(4),
					Name:       pulumi.String("compute"),
					OsType:     pulumi.String("Linux"),
					Role:       pulumi.String("compute"),
					SubnetCidr: pulumi.String("10.0.0.0/24"),
					VmSize:     pulumi.String("Standard_D4s_v3"),
				},
			},
			AuthProfile: containerservice.OpenShiftManagedClusterAuthProfileResponse{
				IdentityProviders: []containerservice.OpenShiftManagedClusterIdentityProviderArgs{
					{
						Name: pulumi.String("Azure AD"),
						Provider: {
							ClientId:             pulumi.String("clientId"),
							CustomerAdminGroupId: pulumi.String("customerAdminGroupId"),
							Kind:                 pulumi.String("AADIdentityProvider"),
							Secret:               pulumi.String("secret"),
							TenantId:             pulumi.String("tenantId"),
						},
					},
				},
			},
			Location: pulumi.String("location1"),
			MasterPoolProfile: containerservice.OpenShiftManagedClusterMasterPoolProfileResponse{
				ApiProperties: &containerservice.OpenShiftAPIPropertiesArgs{
					PrivateApiServer: pulumi.Bool(true),
				},
				Count:      pulumi.Int(3),
				SubnetCidr: pulumi.String("10.0.0.0/24"),
				VmSize:     pulumi.String("Standard_D4s_v3"),
			},
			MonitorProfile: &containerservice.OpenShiftManagedClusterMonitorProfileArgs{
				Enabled:             pulumi.Bool(true),
				WorkspaceResourceID: pulumi.String("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.OperationalInsights/workspaces/workspacename1"),
			},
			NetworkProfile: &containerservice.NetworkProfileArgs{
				ManagementSubnetCidr: pulumi.String("10.0.1.0/24"),
				VnetCidr:             pulumi.String("10.0.0.0/8"),
			},
			OpenShiftVersion:  pulumi.String("v3.11"),
			RefreshCluster:    pulumi.Bool(true),
			ResourceGroupName: pulumi.String("rg1"),
			ResourceName:      pulumi.String("privateclustername1"),
			RouterProfiles: []containerservice.OpenShiftRouterProfileArgs{
				{
					Name: pulumi.String("default"),
				},
			},
			Tags: pulumi.StringMap{
				"archv2": pulumi.String(""),
				"tier":   pulumi.String("production"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
