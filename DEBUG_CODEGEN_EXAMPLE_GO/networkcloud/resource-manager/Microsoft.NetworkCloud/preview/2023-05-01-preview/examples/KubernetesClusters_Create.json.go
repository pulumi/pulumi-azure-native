package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkcloud/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := networkcloud.NewKubernetesCluster(ctx, "kubernetesCluster", &networkcloud.KubernetesClusterArgs{
AadConfiguration: &networkcloud.AadConfigurationArgs{
AdminGroupObjectIds: pulumi.StringArray{
pulumi.String("ffffffff-ffff-ffff-ffff-ffffffffffff"),
},
},
AdministratorConfiguration: networkcloud.AdministratorConfigurationResponse{
AdminUsername: pulumi.String("azure"),
SshPublicKeys: networkcloud.SshPublicKeyArray{
&networkcloud.SshPublicKeyArgs{
KeyData: pulumi.String("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
},
},
},
ControlPlaneNodeConfiguration: networkcloud.ControlPlaneNodeConfigurationResponse{
AdministratorConfiguration: interface{}{
AdminUsername: pulumi.String("azure"),
SshPublicKeys: networkcloud.SshPublicKeyArray{
&networkcloud.SshPublicKeyArgs{
KeyData: pulumi.String("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
},
},
},
AvailabilityZones: pulumi.StringArray{
pulumi.String("1"),
pulumi.String("2"),
pulumi.String("3"),
},
Count: pulumi.Float64(3),
VmSkuName: pulumi.String("NC_G4_v1"),
},
ExtendedLocation: &networkcloud.ExtendedLocationArgs{
Name: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
Type: pulumi.String("CustomLocation"),
},
InitialAgentPoolConfigurations: []networkcloud.InitialAgentPoolConfigurationArgs{
{
AdministratorConfiguration: {
AdminUsername: pulumi.String("azure"),
SshPublicKeys: networkcloud.SshPublicKeyArray{
{
KeyData: pulumi.String("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
},
},
},
AgentOptions: {
HugepagesCount: pulumi.Float64(96),
HugepagesSize: pulumi.String("1G"),
},
AttachedNetworkConfiguration: {
L2Networks: networkcloud.L2NetworkAttachmentConfigurationArray{
{
NetworkId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName"),
PluginType: pulumi.String("DPDK"),
},
},
L3Networks: networkcloud.L3NetworkAttachmentConfigurationArray{
{
IpamEnabled: pulumi.String("False"),
NetworkId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
PluginType: pulumi.String("SRIOV"),
},
},
TrunkedNetworks: networkcloud.TrunkedNetworkAttachmentConfigurationArray{
{
NetworkId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName"),
PluginType: pulumi.String("MACVLAN"),
},
},
},
AvailabilityZones: pulumi.StringArray{
pulumi.String("1"),
pulumi.String("2"),
pulumi.String("3"),
},
Count: pulumi.Float64(3),
Labels: networkcloud.KubernetesLabelArray{
{
Key: pulumi.String("kubernetes.label"),
Value: pulumi.String("true"),
},
},
Mode: pulumi.String("System"),
Name: pulumi.String("SystemPool-1"),
Taints: networkcloud.KubernetesLabelArray{
{
Key: pulumi.String("kubernetes.taint"),
Value: pulumi.String("true"),
},
},
UpgradeSettings: {
MaxSurge: pulumi.String("1"),
},
VmSkuName: pulumi.String("NC_M16_v1"),
},
},
KubernetesClusterName: pulumi.String("kubernetesClusterName"),
KubernetesVersion: pulumi.String("1.24.12-1"),
Location: pulumi.String("location"),
ManagedResourceGroupConfiguration: &networkcloud.ManagedResourceGroupConfigurationArgs{
Location: pulumi.String("East US"),
Name: pulumi.String("my-managed-rg"),
},
NetworkConfiguration: networkcloud.NetworkConfigurationResponse{
AttachedNetworkConfiguration: interface{}{
L2Networks: networkcloud.L2NetworkAttachmentConfigurationArray{
&networkcloud.L2NetworkAttachmentConfigurationArgs{
NetworkId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName"),
PluginType: pulumi.String("DPDK"),
},
},
L3Networks: networkcloud.L3NetworkAttachmentConfigurationArray{
&networkcloud.L3NetworkAttachmentConfigurationArgs{
IpamEnabled: pulumi.String("False"),
NetworkId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
PluginType: pulumi.String("SRIOV"),
},
},
TrunkedNetworks: networkcloud.TrunkedNetworkAttachmentConfigurationArray{
&networkcloud.TrunkedNetworkAttachmentConfigurationArgs{
NetworkId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName"),
PluginType: pulumi.String("MACVLAN"),
},
},
},
BgpServiceLoadBalancerConfiguration: interface{}{
BgpAdvertisements: networkcloud.BgpAdvertisementArray{
&networkcloud.BgpAdvertisementArgs{
AdvertiseToFabric: pulumi.String("True"),
Communities: pulumi.StringArray{
pulumi.String("64512:100"),
},
IpAddressPools: pulumi.StringArray{
pulumi.String("pool1"),
},
Peers: pulumi.StringArray{
pulumi.String("peer1"),
},
},
},
BgpPeers: networkcloud.ServiceLoadBalancerBgpPeerArray{
&networkcloud.ServiceLoadBalancerBgpPeerArgs{
BfdEnabled: pulumi.String("False"),
BgpMultiHop: pulumi.String("False"),
HoldTime: pulumi.String("P300s"),
KeepAliveTime: pulumi.String("P300s"),
MyAsn: pulumi.Float64(64512),
Name: pulumi.String("peer1"),
PeerAddress: pulumi.String("203.0.113.254"),
PeerAsn: pulumi.Float64(64497),
PeerPort: pulumi.Float64(179),
},
},
FabricPeeringEnabled: pulumi.String("True"),
IpAddressPools: networkcloud.IpAddressPoolArray{
&networkcloud.IpAddressPoolArgs{
Addresses: pulumi.StringArray{
pulumi.String("198.51.102.0/24"),
},
AutoAssign: pulumi.String("True"),
Name: pulumi.String("pool1"),
OnlyUseHostIps: pulumi.String("True"),
},
},
},
CloudServicesNetworkId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName"),
CniNetworkId: pulumi.String("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
DnsServiceIp: pulumi.String("198.51.101.2"),
PodCidrs: pulumi.StringArray{
pulumi.String("198.51.100.0/24"),
},
ServiceCidrs: pulumi.StringArray{
pulumi.String("198.51.101.0/24"),
},
},
ResourceGroupName: pulumi.String("resourceGroupName"),
Tags: pulumi.StringMap{
"key1": pulumi.String("myvalue1"),
"key2": pulumi.String("myvalue2"),
},
})
if err != nil {
return err
}
return nil
})
}
