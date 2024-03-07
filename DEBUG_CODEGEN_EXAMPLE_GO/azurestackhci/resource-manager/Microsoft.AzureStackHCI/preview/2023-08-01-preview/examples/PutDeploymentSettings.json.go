package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := azurestackhci.NewDeploymentSetting(ctx, "deploymentSetting", &azurestackhci.DeploymentSettingArgs{
ArcNodeResourceIds: pulumi.StringArray{
pulumi.String("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1"),
pulumi.String("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2"),
},
ClusterName: pulumi.String("myCluster"),
DeploymentConfiguration: azurestackhci.DeploymentConfigurationResponse{
ScaleUnits: azurestackhci.ScaleUnitsArray{
interface{}{
DeploymentData: interface{}{
AdouPath: pulumi.String("OU=ms169,DC=ASZ1PLab8,DC=nttest,DC=microsoft,DC=com"),
Cluster: &azurestackhci.ClusterTypeArgs{
AzureServiceEndpoint: pulumi.String("core.windows.net"),
CloudAccountName: pulumi.String("myasestoragacct"),
Name: pulumi.String("testHCICluster"),
WitnessPath: pulumi.String("Cloud"),
WitnessType: pulumi.String("Cloud"),
},
DomainFqdn: pulumi.String("ASZ1PLab8.nttest.microsoft.com"),
HostNetwork: interface{}{
Intents: azurestackhci.IntentsArray{
interface{}{
Adapter: pulumi.StringArray{
pulumi.String("Port2"),
},
AdapterPropertyOverrides: &azurestackhci.AdapterPropertyOverridesArgs{
JumboPacket: pulumi.String("1514"),
NetworkDirect: pulumi.String("Enabled"),
NetworkDirectTechnology: pulumi.String("iWARP"),
},
Name: pulumi.String("Compute_Management"),
OverrideAdapterProperty: pulumi.Bool(false),
OverrideQosPolicy: pulumi.Bool(false),
OverrideVirtualSwitchConfiguration: pulumi.Bool(false),
QosPolicyOverrides: &azurestackhci.QosPolicyOverridesArgs{
BandwidthPercentageSMB: pulumi.String("50"),
PriorityValue8021ActionCluster: pulumi.String("7"),
PriorityValue8021ActionSMB: pulumi.String("3"),
},
TrafficType: pulumi.StringArray{
pulumi.String("Compute"),
pulumi.String("Management"),
},
VirtualSwitchConfigurationOverrides: &azurestackhci.VirtualSwitchConfigurationOverridesArgs{
EnableIov: pulumi.String("True"),
LoadBalancingAlgorithm: pulumi.String("HyperVPort"),
},
},
},
StorageConnectivitySwitchless: pulumi.Bool(true),
StorageNetworks: azurestackhci.StorageNetworksArray{
&azurestackhci.StorageNetworksArgs{
Name: pulumi.String("Storage1Network"),
NetworkAdapterName: pulumi.String("Port3"),
VlanId: pulumi.String("5"),
},
},
},
InfrastructureNetwork: azurestackhci.InfrastructureNetworkArray{
interface{}{
DnsServers: pulumi.StringArray{
pulumi.String("10.57.50.90"),
},
Gateway: pulumi.String("255.255.248.0"),
IpPools: azurestackhci.IpPoolsArray{
&azurestackhci.IpPoolsArgs{
EndingAddress: pulumi.String("10.57.48.66"),
StartingAddress: pulumi.String("10.57.48.60"),
},
},
SubnetMask: pulumi.String("255.255.248.0"),
},
},
NamingPrefix: pulumi.String("ms169"),
Observability: &azurestackhci.ObservabilityArgs{
EpisodicDataUpload: pulumi.Bool(true),
EuLocation: pulumi.Bool(false),
StreamingDataClient: pulumi.Bool(true),
},
OptionalServices: &azurestackhci.OptionalServicesArgs{
CustomLocation: pulumi.String("customLocationName"),
},
PhysicalNodes: azurestackhci.PhysicalNodesArray{
&azurestackhci.PhysicalNodesArgs{
Ipv4Address: pulumi.String("10.57.51.224"),
Name: pulumi.String("ms169host"),
},
&azurestackhci.PhysicalNodesArgs{
Ipv4Address: pulumi.String("10.57.53.236"),
Name: pulumi.String("ms154host"),
},
},
SecretsLocation: pulumi.String("/subscriptions/db4e2fdb-6d80-4e6e-b7cd-xxxxxxx/resourceGroups/test-rg/providers/Microsoft.KeyVault/vaults/abcd123"),
SecuritySettings: &azurestackhci.SecuritySettingsArgs{
BitlockerBootVolume: pulumi.Bool(true),
BitlockerDataVolumes: pulumi.Bool(true),
CredentialGuardEnforced: pulumi.Bool(false),
DriftControlEnforced: pulumi.Bool(true),
DrtmProtection: pulumi.Bool(true),
HvciProtection: pulumi.Bool(true),
SideChannelMitigationEnforced: pulumi.Bool(true),
SmbClusterEncryption: pulumi.Bool(false),
SmbSigningEnforced: pulumi.Bool(true),
WdacEnforced: pulumi.Bool(true),
},
Storage: &azurestackhci.StorageArgs{
ConfigurationMode: pulumi.String("Express"),
},
},
},
},
Version: pulumi.String("string"),
},
DeploymentMode: pulumi.String("Deploy"),
DeploymentSettingsName: pulumi.String("default"),
ResourceGroupName: pulumi.String("test-rg"),
})
if err != nil {
return err
}
return nil
})
}
