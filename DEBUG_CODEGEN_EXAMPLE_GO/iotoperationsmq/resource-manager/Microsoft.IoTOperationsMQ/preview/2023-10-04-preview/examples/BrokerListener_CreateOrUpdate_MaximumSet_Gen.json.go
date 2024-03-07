package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsmq/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsmq.NewBrokerListener(ctx, "brokerListener", &iotoperationsmq.BrokerListenerArgs{
			AuthenticationEnabled: pulumi.Bool(true),
			AuthorizationEnabled:  pulumi.Bool(true),
			BrokerName:            pulumi.String("HGF6WIy6oHv756MjW0JRLILF"),
			BrokerRef:             pulumi.String("ikuszpfycikq"),
			ExtendedLocation: &iotoperationsmq.ExtendedLocationPropertyArgs{
				Name: pulumi.String("an"),
				Type: pulumi.String("CustomLocation"),
			},
			ListenerName:      pulumi.String("XGEP"),
			Location:          pulumi.String("dppbdcpstouifyko"),
			MqName:            pulumi.String("Z1-2BCdNY4JO--84"),
			NodePort:          pulumi.Int(34375),
			Port:              pulumi.Int(19791),
			ResourceGroupName: pulumi.String("rgiotoperationsmq"),
			ServiceName:       pulumi.String("euxa"),
			ServiceType:       pulumi.String("clusterIp"),
			Tags:              nil,
			Tls: &iotoperationsmq.TlsCertMethodArgs{
				Automatic: &iotoperationsmq.AutomaticCertMethodArgs{
					Duration: pulumi.String("rv"),
					IssuerRef: &iotoperationsmq.CertManagerIssuerRefArgs{
						Group: pulumi.String("wxydv"),
						Kind:  pulumi.String("birgjwuxfjcvyqe"),
						Name:  pulumi.String("krmdlovyynymtvgffaveker"),
					},
					PrivateKey: &iotoperationsmq.CertManagerPrivateKeyArgs{
						Algorithm:      pulumi.String("wwewfsddymjefuhxzqybwvay"),
						RotationPolicy: pulumi.String("jxmpyvfneckopjiakjtous"),
						Size:           pulumi.Int(63427),
					},
					RenewBefore: pulumi.String("dexxoqqkgyofhkbk"),
					San: &iotoperationsmq.SanForCertArgs{
						Dns: pulumi.StringArray{
							pulumi.String("nknzptgqgjvbkgzv"),
						},
						Ip: pulumi.StringArray{
							pulumi.String("jpdkemham"),
						},
					},
					SecretName:      pulumi.String("hquvygbuueerkspqqktviya"),
					SecretNamespace: pulumi.String("aevwndhcnfxitdjykp"),
				},
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
				Manual: &iotoperationsmq.ManualCertMethodArgs{
					SecretName:      pulumi.String("fezcl"),
					SecretNamespace: pulumi.String("ozhayajoooingoczovfusqyilin"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
