package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsmq/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := iotoperationsmq.NewBroker(ctx, "broker", &iotoperationsmq.BrokerArgs{
AuthImage: &iotoperationsmq.ContainerImageArgs{
PullPolicy: pulumi.String("imfuzvqxgbdwliqnn"),
PullSecrets: pulumi.String("klnqimxqsrdwhcqldjvdtsrs"),
Repository: pulumi.String("m"),
Tag: pulumi.String("jygfdiamhhm"),
},
BrokerImage: &iotoperationsmq.ContainerImageArgs{
PullPolicy: pulumi.String("imfuzvqxgbdwliqnn"),
PullSecrets: pulumi.String("klnqimxqsrdwhcqldjvdtsrs"),
Repository: pulumi.String("m"),
Tag: pulumi.String("jygfdiamhhm"),
},
BrokerName: pulumi.String("29tAwt4A2-aH6nP"),
BrokerNodeTolerations: &iotoperationsmq.NodeTolerationsArgs{
Effect: pulumi.String("eeswvciblqmmaeesjoflyvxqbz"),
Key: pulumi.String("wbrstdwxgm"),
Operator: pulumi.String("lbegegneekwnyodtzraarivtwhmzep"),
Value: pulumi.String("sfafsjdcezdmkwibxeluukxgl"),
},
Cardinality: iotoperationsmq.CardinalityResponse{
BackendChain: interface{}{
Partitions: pulumi.Int(34721),
RedundancyFactor: pulumi.Int(468),
TemporaryDiskTransferEnabled: pulumi.Bool(true),
TemporaryDiskTransferHighWatermarkPercent: pulumi.Int(79),
TemporaryDiskTransferLowWatermarkPercent: pulumi.Int(94),
TemporaryMaxBackendMemUsagePercent: pulumi.Int(54),
TemporaryResourceLimits: &iotoperationsmq.TemporaryResourceLimitsConfigArgs{
MaxInflightMessages: pulumi.Int(33208),
MaxInflightPatches: pulumi.Int(3410),
MaxInflightPatchesPerClient: pulumi.Int(58933),
MaxMessageExpirySecs: pulumi.Float64(2036532516),
MaxQueuedMessages: pulumi.Float64(8083241696687839232),
MaxQueuedQos0Messages: pulumi.Float64(6545343433569253376),
MaxSessionExpirySecs: pulumi.Float64(2526293894),
},
Workers: pulumi.Int(15754),
},
Frontend: interface{}{
Replicas: pulumi.Int(38165),
TemporaryResourceLimits: &iotoperationsmq.TemporaryResourceLimitsConfigArgs{
MaxInflightMessages: pulumi.Int(33208),
MaxInflightPatches: pulumi.Int(3410),
MaxInflightPatchesPerClient: pulumi.Int(58933),
MaxMessageExpirySecs: pulumi.Float64(2036532516),
MaxQueuedMessages: pulumi.Float64(8083241696687839232),
MaxQueuedQos0Messages: pulumi.Float64(6545343433569253376),
MaxSessionExpirySecs: pulumi.Float64(2526293894),
},
Workers: pulumi.Int(38),
},
},
Diagnostics: &iotoperationsmq.BrokerDiagnosticsArgs{
DiagnosticServiceEndpoint: pulumi.String("cdvelitwasofaaqhdb"),
EnableMetrics: pulumi.Bool(true),
EnableSelfCheck: pulumi.Bool(true),
EnableSelfTracing: pulumi.Bool(true),
EnableTracing: pulumi.Bool(true),
LogFormat: pulumi.String("tcivnlakxcajynypbz"),
LogLevel: pulumi.String("zdjh"),
MaxCellMapLifetime: pulumi.Float64(997099872515057664),
MetricUpdateFrequencySeconds: pulumi.Float64(6156703238506293248),
ProbeImage: pulumi.String("uzizubdxsgcpjwly"),
SelfCheckFrequencySeconds: pulumi.Float64(579622483050303872),
SelfCheckTimeoutSeconds: pulumi.Float64(7847246333600883712),
SelfTraceFrequencySeconds: pulumi.Float64(6527612490765174784),
SpanChannelCapacity: pulumi.Float64(5533451650716961792),
},
DiskBackedMessageBufferSettings: iotoperationsmq.DiskBackedMessageBufferSettingsResponse{
EphemeralVolumeClaimSpec: interface{}{
AccessModes: pulumi.StringArray{
pulumi.String("cly"),
},
DataSource: &iotoperationsmq.VolumeClaimDataSourceArgs{
ApiGroup: pulumi.String("v"),
Kind: pulumi.String("pvzbnjebkoslzzucpaem"),
Name: pulumi.String("bgzdfwfpdrubbbnfwzyr"),
},
DataSourceRef: &iotoperationsmq.VolumeClaimDataSourceRefArgs{
ApiGroup: pulumi.String("e"),
Kind: pulumi.String("hjbktqbtg"),
Name: pulumi.String("losjjcujomepwhztzptrobavolc"),
},
Resources: &iotoperationsmq.VolumeClaimResourceRequirementsArgs{
Limits: nil,
Requests: nil,
},
Selector: interface{}{
MatchExpressions: iotoperationsmq.VolumeClaimSpecSelectorMatchExpressionsArray{
&iotoperationsmq.VolumeClaimSpecSelectorMatchExpressionsArgs{
Key: pulumi.String("d"),
Operator: pulumi.String("fcfvoarytcdbtccjervsmdis"),
Values: pulumi.StringArray{
pulumi.String("y"),
},
},
},
MatchLabels: nil,
},
StorageClassName: pulumi.String("etajfhrtgatxi"),
VolumeMode: pulumi.String("mipdeutsgidkzpxelbrqggjheplvmx"),
VolumeName: pulumi.String("dacuvlvuullautxjxwdctvzlmd"),
},
MaxSize: pulumi.String("gnwxgqjypylz"),
PersistentVolumeClaimSpec: interface{}{
AccessModes: pulumi.StringArray{
pulumi.String("cly"),
},
DataSource: &iotoperationsmq.VolumeClaimDataSourceArgs{
ApiGroup: pulumi.String("v"),
Kind: pulumi.String("pvzbnjebkoslzzucpaem"),
Name: pulumi.String("bgzdfwfpdrubbbnfwzyr"),
},
DataSourceRef: &iotoperationsmq.VolumeClaimDataSourceRefArgs{
ApiGroup: pulumi.String("e"),
Kind: pulumi.String("hjbktqbtg"),
Name: pulumi.String("losjjcujomepwhztzptrobavolc"),
},
Resources: &iotoperationsmq.VolumeClaimResourceRequirementsArgs{
Limits: nil,
Requests: nil,
},
Selector: interface{}{
MatchExpressions: iotoperationsmq.VolumeClaimSpecSelectorMatchExpressionsArray{
&iotoperationsmq.VolumeClaimSpecSelectorMatchExpressionsArgs{
Key: pulumi.String("d"),
Operator: pulumi.String("fcfvoarytcdbtccjervsmdis"),
Values: pulumi.StringArray{
pulumi.String("y"),
},
},
},
MatchLabels: nil,
},
StorageClassName: pulumi.String("etajfhrtgatxi"),
VolumeMode: pulumi.String("mipdeutsgidkzpxelbrqggjheplvmx"),
VolumeName: pulumi.String("dacuvlvuullautxjxwdctvzlmd"),
},
},
EncryptInternalTraffic: pulumi.Bool(true),
ExtendedLocation: &iotoperationsmq.ExtendedLocationPropertyArgs{
Name: pulumi.String("an"),
Type: pulumi.String("CustomLocation"),
},
HealthManagerImage: &iotoperationsmq.ContainerImageArgs{
PullPolicy: pulumi.String("imfuzvqxgbdwliqnn"),
PullSecrets: pulumi.String("klnqimxqsrdwhcqldjvdtsrs"),
Repository: pulumi.String("m"),
Tag: pulumi.String("jygfdiamhhm"),
},
HealthManagerNodeTolerations: &iotoperationsmq.NodeTolerationsArgs{
Effect: pulumi.String("eeswvciblqmmaeesjoflyvxqbz"),
Key: pulumi.String("wbrstdwxgm"),
Operator: pulumi.String("lbegegneekwnyodtzraarivtwhmzep"),
Value: pulumi.String("sfafsjdcezdmkwibxeluukxgl"),
},
InternalCerts: iotoperationsmq.CertManagerCertOptionsResponse{
Duration: pulumi.String("xjjmzq"),
PrivateKey: &iotoperationsmq.CertManagerPrivateKeyArgs{
Algorithm: pulumi.String("wwewfsddymjefuhxzqybwvay"),
RotationPolicy: pulumi.String("jxmpyvfneckopjiakjtous"),
Size: pulumi.Int(63427),
},
RenewBefore: pulumi.String("zkajhllevrxkfmfyzasmbllvd"),
},
Location: pulumi.String("ltzfwqzs"),
MemoryProfile: pulumi.String("tiny"),
Mode: pulumi.String("auto"),
MqName: pulumi.String("u229L1RZ5"),
ResourceGroupName: pulumi.String("rgiotoperationsmq"),
Tags: nil,
})
if err != nil {
return err
}
return nil
})
}
