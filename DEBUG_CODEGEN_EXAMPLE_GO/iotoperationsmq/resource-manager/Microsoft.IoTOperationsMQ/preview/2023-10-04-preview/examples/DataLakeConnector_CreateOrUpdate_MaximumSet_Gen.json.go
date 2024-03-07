package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsmq/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsmq.NewDataLakeConnector(ctx, "dataLakeConnector", &iotoperationsmq.DataLakeConnectorArgs{
			DataLakeConnectorName: pulumi.String("87v4D"),
			DatabaseFormat:        pulumi.String("delta"),
			ExtendedLocation: &iotoperationsmq.ExtendedLocationPropertyArgs{
				Name: pulumi.String("an"),
				Type: pulumi.String("CustomLocation"),
			},
			Image: &iotoperationsmq.ContainerImageArgs{
				PullPolicy:  pulumi.String("imfuzvqxgbdwliqnn"),
				PullSecrets: pulumi.String("klnqimxqsrdwhcqldjvdtsrs"),
				Repository:  pulumi.String("m"),
				Tag:         pulumi.String("jygfdiamhhm"),
			},
			Instances: pulumi.Int(53467),
			LocalBrokerConnection: &iotoperationsmq.LocalBrokerConnectionSpecArgs{
				Authentication: &iotoperationsmq.LocalBrokerAuthenticationMethodsArgs{
					Kubernetes: &iotoperationsmq.LocalBrokerKubernetesAuthenticationArgs{
						SecretPath:              pulumi.String("soukzfkouir"),
						ServiceAccountTokenName: pulumi.String("suwetviuhmhorhvsidlznnufe"),
					},
				},
				Endpoint: pulumi.String("xc"),
				Tls: &iotoperationsmq.LocalBrokerConnectionTlsArgs{
					TlsEnabled:                    pulumi.Bool(true),
					TrustedCaCertificateConfigMap: pulumi.String("rinkomfeznsfedbmllxlbmmhc"),
				},
			},
			Location: pulumi.String("ewguwvlahlu"),
			LogLevel: pulumi.String("ikicyoalavfmqlodnupfjayxjti"),
			MqName:   pulumi.String("Ox--3e65kYN0731DJ1Qg"),
			NodeTolerations: &iotoperationsmq.NodeTolerationsArgs{
				Effect:   pulumi.String("eeswvciblqmmaeesjoflyvxqbz"),
				Key:      pulumi.String("wbrstdwxgm"),
				Operator: pulumi.String("lbegegneekwnyodtzraarivtwhmzep"),
				Value:    pulumi.String("sfafsjdcezdmkwibxeluukxgl"),
			},
			Protocol:          pulumi.String("v3"),
			ResourceGroupName: pulumi.String("rgiotoperationsmq"),
			Tags:              nil,
			Target: &iotoperationsmq.DataLakeTargetStorageArgs{
				DatalakeStorage: &iotoperationsmq.DataLakeServiceStorageArgs{
					Authentication: &iotoperationsmq.DataLakeServiceStorageAuthenticationArgs{
						AccessTokenSecretName: pulumi.String("cfaoxjbfbwdldqjbfczvovgooyqkl"),
						SystemAssignedManagedIdentity: &iotoperationsmq.ManagedIdentityAuthenticationArgs{
							Audience:      pulumi.String("kjderojhpehosgfcrxxbh"),
							ExtensionName: pulumi.String("cyckjqqzspleajbtkniwrfsqygjfhe"),
						},
					},
					Endpoint: pulumi.String("bddy"),
				},
				FabricOneLake: &iotoperationsmq.DataLakeFabricStorageArgs{
					Authentication: &iotoperationsmq.DataLakeFabricStorageAuthenticationArgs{
						SystemAssignedManagedIdentity: &iotoperationsmq.ManagedIdentityAuthenticationArgs{
							Audience:      pulumi.String("kjderojhpehosgfcrxxbh"),
							ExtensionName: pulumi.String("cyckjqqzspleajbtkniwrfsqygjfhe"),
						},
					},
					Endpoint:   pulumi.String("S.fabric.microsoft.com"),
					FabricPath: pulumi.String("files"),
					Guids: &iotoperationsmq.FabricGuidsArgs{
						LakehouseGuid: pulumi.String("grsapopwjnuzbmnxbjnawaae"),
						WorkspaceGuid: pulumi.String("iwqfplayvdkdxumpdc"),
					},
					Names: &iotoperationsmq.FabricNamesArgs{
						LakehouseName: pulumi.String("iqgqtk"),
						WorkspaceName: pulumi.String("fxvlfhfcmlhcbgpopyqfikqsryct"),
					},
				},
				LocalStorage: &iotoperationsmq.DataLakeLocalStorageArgs{
					VolumeName: pulumi.String("nmzsioldiwteljpplmftk"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
