package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsmq/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsmq.NewKafkaConnector(ctx, "kafkaConnector", &iotoperationsmq.KafkaConnectorArgs{
			ClientIdPrefix: pulumi.String("yybbbeowkw"),
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
			Instances: pulumi.Int(55252),
			KafkaConnection: &iotoperationsmq.KafkaRemoteBrokerConnectionSpecArgs{
				Authentication: &iotoperationsmq.KafkaRemoteBrokerAuthenticationPropertiesArgs{
					AuthType: &iotoperationsmq.KafkaRemoteBrokerAuthenticationTypesArgs{
						Sasl: &iotoperationsmq.SaslRemoteBrokerBasicAuthenticationArgs{
							SaslType: pulumi.String("plain"),
							Token: &iotoperationsmq.SaslRemoteBrokerBasicAuthenticationTokenArgs{
								KeyVault: &iotoperationsmq.KafkaTokenKeyVaultPropertiesArgs{
									Username: pulumi.String("fb"),
									Vault: &iotoperationsmq.KeyVaultConnectionPropertiesArgs{
										Credentials: &iotoperationsmq.KeyVaultCredentialsPropertiesArgs{
											ServicePrincipalLocalSecretName: pulumi.String("wuimjwpbhoglbsxxa"),
										},
										DirectoryId: pulumi.String("eyjniptiykzcgbzok"),
										Name:        pulumi.String("lxmwfan"),
									},
									VaultSecret: &iotoperationsmq.KeyVaultSecretObjectArgs{
										Name:    pulumi.String("bmectskddmpjxnsogwooexj"),
										Version: pulumi.String("unjfbf"),
									},
								},
								SecretName: pulumi.String("hxmqokubwldgjdtjv"),
							},
						},
						SystemAssignedManagedIdentity: &iotoperationsmq.ManagedIdentityAuthenticationArgs{
							Audience:      pulumi.String("kjderojhpehosgfcrxxbh"),
							ExtensionName: pulumi.String("cyckjqqzspleajbtkniwrfsqygjfhe"),
						},
						X509: &iotoperationsmq.KafkaX509AuthenticationArgs{
							KeyVault: &iotoperationsmq.KeyVaultCertificatePropertiesArgs{
								Vault: &iotoperationsmq.KeyVaultConnectionPropertiesArgs{
									Credentials: &iotoperationsmq.KeyVaultCredentialsPropertiesArgs{
										ServicePrincipalLocalSecretName: pulumi.String("wuimjwpbhoglbsxxa"),
									},
									DirectoryId: pulumi.String("eyjniptiykzcgbzok"),
									Name:        pulumi.String("lxmwfan"),
								},
								VaultCaChainSecret: &iotoperationsmq.KeyVaultSecretObjectArgs{
									Name:    pulumi.String("bmectskddmpjxnsogwooexj"),
									Version: pulumi.String("unjfbf"),
								},
								VaultCert: &iotoperationsmq.KeyVaultSecretObjectArgs{
									Name:    pulumi.String("bmectskddmpjxnsogwooexj"),
									Version: pulumi.String("unjfbf"),
								},
							},
							SecretName: pulumi.String("jlrjvqyoygynlpsekfbvyrpu"),
						},
					},
					Enabled: pulumi.Bool(true),
				},
				Endpoint: pulumi.String("odxpssuhjkbonjmbhbebfjcah"),
				Tls: &iotoperationsmq.KafkaRemoteBrokerConnectionTlsArgs{
					TlsEnabled:                    pulumi.Bool(true),
					TrustedCaCertificateConfigMap: pulumi.String("kndjozglnxsgnzxq"),
				},
			},
			KafkaConnectorName: pulumi.String("V5--OL8-R"),
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
			Location: pulumi.String("mpbupgmqssnb"),
			LogLevel: pulumi.String("g"),
			MqName:   pulumi.String("s47Lj-S2S-Q-XY"),
			NodeTolerations: &iotoperationsmq.NodeTolerationsArgs{
				Effect:   pulumi.String("eeswvciblqmmaeesjoflyvxqbz"),
				Key:      pulumi.String("wbrstdwxgm"),
				Operator: pulumi.String("lbegegneekwnyodtzraarivtwhmzep"),
				Value:    pulumi.String("sfafsjdcezdmkwibxeluukxgl"),
			},
			ResourceGroupName: pulumi.String("rgiotoperationsmq"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
