package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/batch/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := batch.NewPool(ctx, "pool", &batch.PoolArgs{
AccountName: pulumi.String("sampleacct"),
ApplicationLicenses: pulumi.StringArray{
pulumi.String("app-license0"),
pulumi.String("app-license1"),
},
ApplicationPackages: []batch.ApplicationPackageReferenceArgs{
{
Id: pulumi.String("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
Version: pulumi.String("asdf"),
},
},
Certificates: []batch.CertificateReferenceArgs{
{
Id: pulumi.String("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567"),
StoreLocation: batch.CertificateStoreLocationLocalMachine,
StoreName: pulumi.String("MY"),
Visibility: batch.CertificateVisibilityArray{
batch.CertificateVisibilityRemoteUser,
},
},
},
DeploymentConfiguration: batch.DeploymentConfigurationResponse{
CloudServiceConfiguration: &batch.CloudServiceConfigurationArgs{
OsFamily: pulumi.String("4"),
OsVersion: pulumi.String("WA-GUEST-OS-4.45_201708-01"),
},
},
DisplayName: pulumi.String("my-pool-name"),
InterNodeCommunication: batch.InterNodeCommunicationStateEnabled,
Metadata: []batch.MetadataItemArgs{
{
Name: pulumi.String("metadata-1"),
Value: pulumi.String("value-1"),
},
{
Name: pulumi.String("metadata-2"),
Value: pulumi.String("value-2"),
},
},
NetworkConfiguration: batch.NetworkConfigurationResponse{
PublicIPAddressConfiguration: &batch.PublicIPAddressConfigurationArgs{
IpAddressIds: pulumi.StringArray{
pulumi.String("/subscriptions/subid1/resourceGroups/rg13/providers/Microsoft.Network/publicIPAddresses/ip135"),
pulumi.String("/subscriptions/subid2/resourceGroups/rg24/providers/Microsoft.Network/publicIPAddresses/ip268"),
},
Provision: batch.IPAddressProvisioningTypeUserManaged,
},
SubnetId: pulumi.String("/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123"),
},
PoolName: pulumi.String("testpool"),
ResourceGroupName: pulumi.String("default-azurebatch-japaneast"),
ScaleSettings: batch.ScaleSettingsResponse{
FixedScale: &batch.FixedScaleSettingsArgs{
NodeDeallocationOption: batch.ComputeNodeDeallocationOptionTaskCompletion,
ResizeTimeout: pulumi.String("PT8M"),
TargetDedicatedNodes: pulumi.Int(6),
TargetLowPriorityNodes: pulumi.Int(28),
},
},
StartTask: batch.StartTaskResponse{
CommandLine: pulumi.String("cmd /c SET"),
EnvironmentSettings: batch.EnvironmentSettingArray{
&batch.EnvironmentSettingArgs{
Name: pulumi.String("MYSET"),
Value: pulumi.String("1234"),
},
},
MaxTaskRetryCount: pulumi.Int(6),
ResourceFiles: batch.ResourceFileArray{
&batch.ResourceFileArgs{
FileMode: pulumi.String("777"),
FilePath: pulumi.String("c:\\temp\\gohere"),
HttpUrl: pulumi.String("https://testaccount.blob.core.windows.net/example-blob-file"),
},
},
UserIdentity: interface{}{
AutoUser: &batch.AutoUserSpecificationArgs{
ElevationLevel: batch.ElevationLevelAdmin,
Scope: batch.AutoUserScopePool,
},
},
WaitForSuccess: pulumi.Bool(true),
},
TaskSchedulingPolicy: &batch.TaskSchedulingPolicyArgs{
NodeFillType: batch.ComputeNodeFillTypePack,
},
TaskSlotsPerNode: pulumi.Int(13),
UserAccounts: []batch.UserAccountArgs{
{
ElevationLevel: batch.ElevationLevelAdmin,
LinuxUserConfiguration: {
Gid: pulumi.Int(4567),
SshPrivateKey: pulumi.String("sshprivatekeyvalue"),
Uid: pulumi.Int(1234),
},
Name: pulumi.String("username1"),
Password: pulumi.String("<ExamplePassword>"),
},
},
VmSize: pulumi.String("STANDARD_D4"),
})
if err != nil {
return err
}
return nil
})
}
