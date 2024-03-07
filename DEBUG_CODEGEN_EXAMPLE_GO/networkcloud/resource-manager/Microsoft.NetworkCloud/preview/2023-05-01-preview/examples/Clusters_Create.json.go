package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkcloud.NewCluster(ctx, "cluster", &networkcloud.ClusterArgs{
			AggregatorOrSingleRackDefinition: &networkcloud.RackDefinitionArgs{
				BareMetalMachineConfigurationData: networkcloud.BareMetalMachineConfigurationDataArray{
					&networkcloud.BareMetalMachineConfigurationDataArgs{
						BmcCredentials: &networkcloud.AdministrativeCredentialsArgs{
							Password: pulumi.String("{password}"),
							Username: pulumi.String("username"),
						},
						BmcMacAddress:  pulumi.String("AA:BB:CC:DD:EE:FF"),
						BootMacAddress: pulumi.String("00:BB:CC:DD:EE:FF"),
						MachineDetails: pulumi.String("extraDetails"),
						MachineName:    pulumi.String("bmmName1"),
						RackSlot:       pulumi.Float64(1),
						SerialNumber:   pulumi.String("BM1219XXX"),
					},
					&networkcloud.BareMetalMachineConfigurationDataArgs{
						BmcCredentials: &networkcloud.AdministrativeCredentialsArgs{
							Password: pulumi.String("{password}"),
							Username: pulumi.String("username"),
						},
						BmcMacAddress:  pulumi.String("AA:BB:CC:DD:EE:00"),
						BootMacAddress: pulumi.String("00:BB:CC:DD:EE:00"),
						MachineDetails: pulumi.String("extraDetails"),
						MachineName:    pulumi.String("bmmName2"),
						RackSlot:       pulumi.Float64(2),
						SerialNumber:   pulumi.String("BM1219YYY"),
					},
				},
				NetworkRackId:    pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName"),
				RackLocation:     pulumi.String("Foo Datacenter, Floor 3, Aisle 9, Rack 2"),
				RackSerialNumber: pulumi.String("AA1234"),
				RackSkuId:        pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"),
				StorageApplianceConfigurationData: networkcloud.StorageApplianceConfigurationDataArray{
					&networkcloud.StorageApplianceConfigurationDataArgs{
						AdminCredentials: &networkcloud.AdministrativeCredentialsArgs{
							Password: pulumi.String("{password}"),
							Username: pulumi.String("username"),
						},
						RackSlot:             pulumi.Float64(1),
						SerialNumber:         pulumi.String("BM1219XXX"),
						StorageApplianceName: pulumi.String("vmName"),
					},
				},
			},
			AnalyticsWorkspaceId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/microsoft.operationalInsights/workspaces/logAnalyticsWorkspaceName"),
			ClusterLocation:      pulumi.String("Foo Street, 3rd Floor, row 9"),
			ClusterName:          pulumi.String("clusterName"),
			ClusterServicePrincipal: &networkcloud.ServicePrincipalInformationArgs{
				ApplicationId: pulumi.String("12345678-1234-1234-1234-123456789012"),
				Password:      pulumi.String("{password}"),
				PrincipalId:   pulumi.String("00000008-0004-0004-0004-000000000012"),
				TenantId:      pulumi.String("80000000-4000-4000-4000-120000000000"),
			},
			ClusterType:    pulumi.String("SingleRack"),
			ClusterVersion: pulumi.String("1.0.0"),
			ComputeDeploymentThreshold: &networkcloud.ValidationThresholdArgs{
				Grouping: pulumi.String("PerCluster"),
				Type:     pulumi.String("PercentSuccess"),
				Value:    pulumi.Float64(90),
			},
			ComputeRackDefinitions: networkcloud.RackDefinitionArray{
				&networkcloud.RackDefinitionArgs{
					BareMetalMachineConfigurationData: networkcloud.BareMetalMachineConfigurationDataArray{
						&networkcloud.BareMetalMachineConfigurationDataArgs{
							BmcCredentials: &networkcloud.AdministrativeCredentialsArgs{
								Password: pulumi.String("{password}"),
								Username: pulumi.String("username"),
							},
							BmcMacAddress:  pulumi.String("AA:BB:CC:DD:EE:FF"),
							BootMacAddress: pulumi.String("00:BB:CC:DD:EE:FF"),
							MachineDetails: pulumi.String("extraDetails"),
							MachineName:    pulumi.String("bmmName1"),
							RackSlot:       pulumi.Float64(1),
							SerialNumber:   pulumi.String("BM1219XXX"),
						},
						&networkcloud.BareMetalMachineConfigurationDataArgs{
							BmcCredentials: &networkcloud.AdministrativeCredentialsArgs{
								Password: pulumi.String("{password}"),
								Username: pulumi.String("username"),
							},
							BmcMacAddress:  pulumi.String("AA:BB:CC:DD:EE:00"),
							BootMacAddress: pulumi.String("00:BB:CC:DD:EE:00"),
							MachineDetails: pulumi.String("extraDetails"),
							MachineName:    pulumi.String("bmmName2"),
							RackSlot:       pulumi.Float64(2),
							SerialNumber:   pulumi.String("BM1219YYY"),
						},
					},
					NetworkRackId:    pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName"),
					RackLocation:     pulumi.String("Foo Datacenter, Floor 3, Aisle 9, Rack 2"),
					RackSerialNumber: pulumi.String("AA1234"),
					RackSkuId:        pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"),
					StorageApplianceConfigurationData: networkcloud.StorageApplianceConfigurationDataArray{
						&networkcloud.StorageApplianceConfigurationDataArgs{
							AdminCredentials: &networkcloud.AdministrativeCredentialsArgs{
								Password: pulumi.String("{password}"),
								Username: pulumi.String("username"),
							},
							RackSlot:             pulumi.Float64(1),
							SerialNumber:         pulumi.String("BM1219XXX"),
							StorageApplianceName: pulumi.String("vmName"),
						},
					},
				},
			},
			ExtendedLocation: &networkcloud.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterManagerExtendedLocationName"),
				Type: pulumi.String("CustomLocation"),
			},
			Location: pulumi.String("location"),
			ManagedResourceGroupConfiguration: &networkcloud.ManagedResourceGroupConfigurationArgs{
				Location: pulumi.String("East US"),
				Name:     pulumi.String("my-managed-rg"),
			},
			NetworkFabricId:   pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabrics/fabricName"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("myvalue1"),
				"key2": pulumi.String("myvalue2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
