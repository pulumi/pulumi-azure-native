package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurearcdata/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurearcdata.NewDataController(ctx, "dataController", &azurearcdata.DataControllerArgs{
			DataControllerName: pulumi.String("testdataController"),
			ExtendedLocation: &azurearcdata.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
				Type: pulumi.String("CustomLocation"),
			},
			Location: pulumi.String("northeurope"),
			Properties: &azurearcdata.DataControllerPropertiesArgs{
				BasicLoginInformation: &azurearcdata.BasicLoginInformationArgs{
					Password: pulumi.String("********"),
					Username: pulumi.String("username"),
				},
				ClusterId:      pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
				ExtensionId:    pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
				Infrastructure: azurearcdata.InfrastructureOnpremises,
				LogAnalyticsWorkspaceConfig: &azurearcdata.LogAnalyticsWorkspaceConfigArgs{
					PrimaryKey:  pulumi.String("********"),
					WorkspaceId: pulumi.String("00000000-1111-2222-3333-444444444444"),
				},
				LogsDashboardCredential: &azurearcdata.BasicLoginInformationArgs{
					Password: pulumi.String("********"),
					Username: pulumi.String("username"),
				},
				MetricsDashboardCredential: &azurearcdata.BasicLoginInformationArgs{
					Password: pulumi.String("********"),
					Username: pulumi.String("username"),
				},
				OnPremiseProperty: &azurearcdata.OnPremisePropertyArgs{
					Id:               pulumi.String("12345678-1234-1234-ab12-1a2b3c4d5e6f"),
					PublicSigningKey: pulumi.String("publicOnPremSigningKey"),
				},
				UploadServicePrincipal: &azurearcdata.UploadServicePrincipalArgs{
					Authority:    pulumi.String("https://login.microsoftonline.com/"),
					ClientId:     pulumi.String("00000000-1111-2222-3333-444444444444"),
					ClientSecret: pulumi.String("********"),
					TenantId:     pulumi.String("00000000-1111-2222-3333-444444444444"),
				},
				UploadWatermark: &azurearcdata.UploadWatermarkArgs{
					Logs:    pulumi.String("2020-01-01T17:18:19.1234567Z"),
					Metrics: pulumi.String("2020-01-01T17:18:19.1234567Z"),
					Usages:  pulumi.String("2020-01-01T17:18:19.1234567Z"),
				},
			},
			ResourceGroupName: pulumi.String("testrg"),
			Tags: pulumi.StringMap{
				"mytag": pulumi.String("myval"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
