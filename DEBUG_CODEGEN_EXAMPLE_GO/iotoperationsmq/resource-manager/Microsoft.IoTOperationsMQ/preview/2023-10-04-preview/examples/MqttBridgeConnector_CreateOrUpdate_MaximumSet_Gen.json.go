package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsmq/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := iotoperationsmq.NewMqttBridgeConnector(ctx, "mqttBridgeConnector", &iotoperationsmq.MqttBridgeConnectorArgs{
BridgeInstances: pulumi.Int(4528),
ClientIdPrefix: pulumi.String("yqipejvabahsexbnttiegjnh"),
ExtendedLocation: &iotoperationsmq.ExtendedLocationPropertyArgs{
Name: pulumi.String("an"),
Type: pulumi.String("CustomLocation"),
},
Image: &iotoperationsmq.ContainerImageArgs{
PullPolicy: pulumi.String("imfuzvqxgbdwliqnn"),
PullSecrets: pulumi.String("klnqimxqsrdwhcqldjvdtsrs"),
Repository: pulumi.String("m"),
Tag: pulumi.String("jygfdiamhhm"),
},
LocalBrokerConnection: iotoperationsmq.LocalBrokerConnectionSpecResponse{
Authentication: interface{}{
Kubernetes: &iotoperationsmq.LocalBrokerKubernetesAuthenticationArgs{
SecretPath: pulumi.String("soukzfkouir"),
ServiceAccountTokenName: pulumi.String("suwetviuhmhorhvsidlznnufe"),
},
},
Endpoint: pulumi.String("xc"),
Tls: &iotoperationsmq.LocalBrokerConnectionTlsArgs{
TlsEnabled: pulumi.Bool(true),
TrustedCaCertificateConfigMap: pulumi.String("rinkomfeznsfedbmllxlbmmhc"),
},
},
Location: pulumi.String("frztvxzhskx"),
LogLevel: pulumi.String("gpgijsotipdtvvkpnckuziqqv"),
MqName: pulumi.String("R8-6x-Y-L-F-21RP5-XVv"),
MqttBridgeConnectorName: pulumi.String("k1v-U4P2440C1b7T8y-G"),
NodeTolerations: &iotoperationsmq.NodeTolerationsArgs{
Effect: pulumi.String("eeswvciblqmmaeesjoflyvxqbz"),
Key: pulumi.String("wbrstdwxgm"),
Operator: pulumi.String("lbegegneekwnyodtzraarivtwhmzep"),
Value: pulumi.String("sfafsjdcezdmkwibxeluukxgl"),
},
Protocol: pulumi.String("v3"),
RemoteBrokerConnection: iotoperationsmq.MqttBridgeRemoteBrokerConnectionSpecResponse{
Authentication: interface{}{
SystemAssignedManagedIdentity: &iotoperationsmq.ManagedIdentityAuthenticationArgs{
Audience: pulumi.String("kjderojhpehosgfcrxxbh"),
ExtensionName: pulumi.String("cyckjqqzspleajbtkniwrfsqygjfhe"),
},
X509: interface{}{
KeyVault: interface{}{
Vault: interface{}{
Credentials: &iotoperationsmq.KeyVaultCredentialsPropertiesArgs{
ServicePrincipalLocalSecretName: pulumi.String("wuimjwpbhoglbsxxa"),
},
DirectoryId: pulumi.String("eyjniptiykzcgbzok"),
Name: pulumi.String("lxmwfan"),
},
VaultCaChainSecret: &iotoperationsmq.KeyVaultSecretObjectArgs{
Name: pulumi.String("bmectskddmpjxnsogwooexj"),
Version: pulumi.String("unjfbf"),
},
VaultCert: &iotoperationsmq.KeyVaultSecretObjectArgs{
Name: pulumi.String("bmectskddmpjxnsogwooexj"),
Version: pulumi.String("unjfbf"),
},
},
SecretName: pulumi.String("x"),
},
},
Endpoint: pulumi.String("bshzrukafmxjgnrlhzlxbbzjdbqh"),
Protocol: pulumi.String("mqtt"),
Tls: &iotoperationsmq.MqttBridgeRemoteBrokerConnectionTlsArgs{
TlsEnabled: pulumi.Bool(true),
TrustedCaCertificateConfigMap: pulumi.String("ivtebqmclgfjx"),
},
},
ResourceGroupName: pulumi.String("rgiotoperationsmq"),
Tags: nil,
})
if err != nil {
return err
}
return nil
})
}
