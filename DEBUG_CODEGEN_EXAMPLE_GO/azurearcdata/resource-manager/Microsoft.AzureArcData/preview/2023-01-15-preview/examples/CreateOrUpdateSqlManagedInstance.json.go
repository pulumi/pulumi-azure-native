package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurearcdata/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurearcdata.NewSqlManagedInstance(ctx, "sqlManagedInstance", &azurearcdata.SqlManagedInstanceArgs{
			ExtendedLocation: &azurearcdata.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
				Type: pulumi.String("CustomLocation"),
			},
			Location: pulumi.String("northeurope"),
			Properties: &azurearcdata.SqlManagedInstancePropertiesArgs{
				ActiveDirectoryInformation: &azurearcdata.ActiveDirectoryInformationArgs{
					KeytabInformation: &azurearcdata.KeytabInformationArgs{
						Keytab: pulumi.String("********"),
					},
				},
				Admin: pulumi.String("Admin user"),
				BasicLoginInformation: &azurearcdata.BasicLoginInformationArgs{
					Password: pulumi.String("********"),
					Username: pulumi.String("username"),
				},
				ClusterId:   pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
				EndTime:     pulumi.String("Instance end time"),
				ExtensionId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
				K8sRaw: &azurearcdata.SqlManagedInstanceK8sRawArgs{
					Spec: &azurearcdata.SqlManagedInstanceK8sSpecArgs{
						Replicas: pulumi.Int(1),
						Scheduling: &azurearcdata.K8sSchedulingArgs{
							Default: &azurearcdata.K8sSchedulingOptionsArgs{
								Resources: &azurearcdata.K8sResourceRequirementsArgs{
									Limits: pulumi.StringMap{
										"additionalProperty": pulumi.String("additionalValue"),
										"cpu":                pulumi.String("1"),
										"memory":             pulumi.String("8Gi"),
									},
									Requests: pulumi.StringMap{
										"additionalProperty": pulumi.String("additionalValue"),
										"cpu":                pulumi.String("1"),
										"memory":             pulumi.String("8Gi"),
									},
								},
							},
						},
						Security: &azurearcdata.K8sSecurityArgs{
							ActiveDirectory: &azurearcdata.K8sActiveDirectoryArgs{
								AccountName: pulumi.String("Account name"),
								Connector: &azurearcdata.K8sActiveDirectoryConnectorArgs{
									Name:      pulumi.String("Name of connector"),
									Namespace: pulumi.String("Namespace of connector"),
								},
								EncryptionTypes: pulumi.StringArray{
									pulumi.String("Encryption type item1, Encryption type item2,..."),
								},
								KeytabSecret: pulumi.String("Key tab secret of account"),
							},
							AdminLoginSecret:         pulumi.String("test-sql-login-secret"),
							ServiceCertificateSecret: pulumi.String("Service Certificate Secret"),
							TransparentDataEncryption: &azurearcdata.K8stransparentDataEncryptionArgs{
								Mode: pulumi.String("SystemManaged"),
							},
						},
						Settings: &azurearcdata.K8sSettingsArgs{
							Network: &azurearcdata.K8sNetworkSettingsArgs{
								Forceencryption: pulumi.Int(0),
								Tlsciphers:      pulumi.String("ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES128-SHA256:ECDHE-RSA-AES256-SHA384"),
								Tlsprotocols:    pulumi.String("1.2"),
							},
						},
					},
				},
				LicenseType: pulumi.String("LicenseIncluded"),
				StartTime:   pulumi.String("Instance start time"),
			},
			ResourceGroupName: pulumi.String("testrg"),
			Sku: &azurearcdata.SqlManagedInstanceSkuArgs{
				Dev:  pulumi.Bool(true),
				Name: azurearcdata.SqlManagedInstanceSkuNameVCore,
				Tier: azurearcdata.SqlManagedInstanceSkuTierGeneralPurpose,
			},
			SqlManagedInstanceName: pulumi.String("testsqlManagedInstance"),
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
