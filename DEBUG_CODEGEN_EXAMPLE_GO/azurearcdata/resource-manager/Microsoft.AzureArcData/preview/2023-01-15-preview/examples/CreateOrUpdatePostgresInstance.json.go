package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurearcdata/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurearcdata.NewPostgresInstance(ctx, "postgresInstance", &azurearcdata.PostgresInstanceArgs{
			ExtendedLocation: &azurearcdata.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
				Type: pulumi.String("CustomLocation"),
			},
			Location:             pulumi.String("eastus"),
			PostgresInstanceName: pulumi.String("testpostgresInstance"),
			Properties: azurearcdata.PostgresInstancePropertiesResponse{
				Admin: pulumi.String("admin"),
				BasicLoginInformation: &azurearcdata.BasicLoginInformationArgs{
					Password: pulumi.String("********"),
					Username: pulumi.String("username"),
				},
				DataControllerId: pulumi.String("dataControllerId"),
				K8sRaw: pulumi.Any{
					ApiVersion: "apiVersion",
					Kind:       "postgresql-12",
					Metadata: map[string]interface{}{
						"creationTimestamp": "2020-08-25T14:55:10Z",
						"generation":        1,
						"name":              "pg1",
						"namespace":         "test",
						"resourceVersion":   "527780",
						"selfLink":          "/apis/arcdata.microsoft.com/v1alpha1/namespaces/test/postgresql-12s/pg1",
						"uid":               "1111aaaa-ffff-ffff-ffff-99999aaaaaaa",
					},
					Spec: map[string]interface{}{
						"backups": map[string]interface{}{
							"deltaMinutes": 3,
							"fullMinutes":  10,
							"tiers": []map[string]interface{}{
								map[string]interface{}{
									"retention": map[string]interface{}{
										"maximums": []string{
											"6",
											"512MB",
										},
										"minimums": []string{
											"3",
										},
									},
									"storage": map[string]interface{}{
										"volumeSize": "1Gi",
									},
								},
							},
						},
						"engine": map[string]interface{}{
							"extensions": []map[string]interface{}{
								map[string]interface{}{
									"name": "citus",
								},
							},
						},
						"scale": map[string]interface{}{
							"shards": 3,
						},
						"scheduling": map[string]interface{}{
							"default": map[string]interface{}{
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"memory": "256Mi",
									},
								},
							},
						},
						"service": map[string]interface{}{
							"type": "NodePort",
						},
						"storage": map[string]interface{}{
							"data": map[string]interface{}{
								"className": "local-storage",
								"size":      "5Gi",
							},
							"logs": map[string]interface{}{
								"className": "local-storage",
								"size":      "5Gi",
							},
						},
					},
					Status: map[string]interface{}{
						"externalEndpoint": nil,
						"readyPods":        "4/4",
						"state":            "Ready",
					},
				},
			},
			ResourceGroupName: pulumi.String("testrg"),
			Sku: &azurearcdata.PostgresInstanceSkuArgs{
				Dev:  pulumi.Bool(true),
				Name: pulumi.String("default"),
				Tier: azurearcdata.PostgresInstanceSkuTierHyperscale,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
