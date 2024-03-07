package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsmq/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsmq.NewBrokerAuthentication(ctx, "brokerAuthentication", &iotoperationsmq.BrokerAuthenticationArgs{
			AuthenticationMethods: iotoperationsmq.BrokerAuthenticatorMethodsArray{
				&iotoperationsmq.BrokerAuthenticatorMethodsArgs{
					Custom: &iotoperationsmq.BrokerAuthenticatorMethodCustomArgs{
						Auth: &iotoperationsmq.BrokerAuthenticatorCustomAuthArgs{
							X509: &iotoperationsmq.BrokerAuthenticatorCustomAuthX509Args{
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
								SecretName: pulumi.String("dordbwjewnqkhfd"),
							},
						},
						CaCertConfigMap: pulumi.String("diufihyysdcosgy"),
						Endpoint:        pulumi.String("yy"),
						Headers:         nil,
					},
					Sat: &iotoperationsmq.BrokerAuthenticatorMethodSatArgs{
						Audiences: pulumi.StringArray{
							pulumi.String("fiyitxutbuuhwtltukyjacads"),
						},
					},
					Svid: &iotoperationsmq.BrokerAuthenticatorMethodSvidArgs{
						AgentSocketPath:     pulumi.String("gnyowebmeaj"),
						IdentityMaxRetry:    pulumi.Float64(4031184731),
						IdentityWaitRetryMs: pulumi.Float64(2243705844935085568),
					},
					UsernamePassword: &iotoperationsmq.BrokerAuthenticatorMethodUsernamePasswordArgs{
						KeyVault: &iotoperationsmq.KeyVaultSecretPropertiesArgs{
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
						SecretName: pulumi.String("blk"),
					},
					X509: &iotoperationsmq.BrokerAuthenticatorMethodX509Args{
						Attributes: &iotoperationsmq.BrokerAuthenticatorMethodX509AttributesArgs{
							KeyVault: &iotoperationsmq.KeyVaultSecretPropertiesArgs{
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
							SecretName: pulumi.String("ybcke"),
						},
						TrustedClientCaCertConfigMap: pulumi.String("udidafmnpt"),
					},
				},
			},
			AuthenticationName: pulumi.String("lUo-GQ3-95F-1O-"),
			BrokerName:         pulumi.String("87v1GC9557XuP-JLI4-"),
			ExtendedLocation: &iotoperationsmq.ExtendedLocationPropertyArgs{
				Name: pulumi.String("an"),
				Type: pulumi.String("CustomLocation"),
			},
			ListenerRef: pulumi.StringArray{
				pulumi.String("dhjpypfjzzmwm"),
			},
			Location:          pulumi.String("vtxegvaeqwyupplnm"),
			MqName:            pulumi.String("2S-A2-D9kC946K"),
			ResourceGroupName: pulumi.String("rgiotoperationsmq"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
