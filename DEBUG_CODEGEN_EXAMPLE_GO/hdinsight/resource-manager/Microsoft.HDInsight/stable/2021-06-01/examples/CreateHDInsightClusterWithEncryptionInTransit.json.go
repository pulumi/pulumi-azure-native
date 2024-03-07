package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hdinsight/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := hdinsight.NewCluster(ctx, "cluster", &hdinsight.ClusterArgs{
ClusterName: pulumi.String("cluster1"),
Properties: hdinsight.ClusterGetPropertiesResponse{
ClusterDefinition: &hdinsight.ClusterDefinitionArgs{
Configurations: pulumi.Any{
Gateway: map[string]interface{}{
"restAuthCredential.isEnabled": true,
"restAuthCredential.password": "**********",
"restAuthCredential.username": "admin",
},
},
Kind: pulumi.String("Hadoop"),
},
ClusterVersion: pulumi.String("3.6"),
ComputeProfile: interface{}{
Roles: hdinsight.RoleArray{
interface{}{
HardwareProfile: &hdinsight.HardwareProfileArgs{
VmSize: pulumi.String("Large"),
},
Name: pulumi.String("headnode"),
OsProfile: interface{}{
LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
Password: pulumi.String("**********"),
Username: pulumi.String("sshuser"),
},
},
TargetInstanceCount: pulumi.Int(2),
},
interface{}{
HardwareProfile: &hdinsight.HardwareProfileArgs{
VmSize: pulumi.String("Large"),
},
Name: pulumi.String("workernode"),
OsProfile: interface{}{
LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
Password: pulumi.String("**********"),
Username: pulumi.String("sshuser"),
},
},
TargetInstanceCount: pulumi.Int(3),
},
interface{}{
HardwareProfile: &hdinsight.HardwareProfileArgs{
VmSize: pulumi.String("Small"),
},
Name: pulumi.String("zookeepernode"),
OsProfile: interface{}{
LinuxOperatingSystemProfile: &hdinsight.LinuxOperatingSystemProfileArgs{
Password: pulumi.String("**********"),
Username: pulumi.String("sshuser"),
},
},
TargetInstanceCount: pulumi.Int(3),
},
},
},
EncryptionInTransitProperties: &hdinsight.EncryptionInTransitPropertiesArgs{
IsEncryptionInTransitEnabled: pulumi.Bool(true),
},
OsType: pulumi.String("Linux"),
StorageProfile: interface{}{
Storageaccounts: hdinsight.StorageAccountArray{
&hdinsight.StorageAccountArgs{
Container: pulumi.String("default8525"),
IsDefault: pulumi.Bool(true),
Key: pulumi.String("storagekey"),
Name: pulumi.String("mystorage.blob.core.windows.net"),
},
},
},
Tier: pulumi.String("Standard"),
},
ResourceGroupName: pulumi.String("rg1"),
})
if err != nil {
return err
}
return nil
})
}
