package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := containerservice.NewManagedCluster(ctx, "managedCluster", &containerservice.ManagedClusterArgs{
AddonProfiles: nil,
AgentPoolProfiles: []containerservice.ManagedClusterAgentPoolProfileArgs{
{
Count: pulumi.Int(3),
EnableNodePublicIP: pulumi.Bool(true),
Mode: pulumi.String("System"),
Name: pulumi.String("nodepool1"),
OsType: pulumi.String("Linux"),
Type: pulumi.String("VirtualMachineScaleSets"),
VmSize: pulumi.String("Standard_DS2_v2"),
},
},
AutoScalerProfile: &containerservice.ManagedClusterPropertiesAutoScalerProfileArgs{
ScaleDownDelayAfterAdd: pulumi.String("15m"),
ScanInterval: pulumi.String("20s"),
},
DiskEncryptionSetID: pulumi.String("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
DnsPrefix: pulumi.String("dnsprefix1"),
EnablePodSecurityPolicy: pulumi.Bool(true),
EnableRBAC: pulumi.Bool(true),
HttpProxyConfig: &containerservice.ManagedClusterHTTPProxyConfigArgs{
HttpProxy: pulumi.String("http://myproxy.server.com:8080"),
HttpsProxy: pulumi.String("https://myproxy.server.com:8080"),
NoProxy: pulumi.StringArray{
pulumi.String("localhost"),
pulumi.String("127.0.0.1"),
},
TrustedCa: pulumi.String("Q29uZ3JhdHMhIFlvdSBoYXZlIGZvdW5kIGEgaGlkZGVuIG1lc3NhZ2U="),
},
KubernetesVersion: pulumi.String(""),
LinuxProfile: containerservice.ContainerServiceLinuxProfileResponse{
AdminUsername: pulumi.String("azureuser"),
Ssh: interface{}{
PublicKeys: containerservice.ContainerServiceSshPublicKeyArray{
&containerservice.ContainerServiceSshPublicKeyArgs{
KeyData: pulumi.String("keydata"),
},
},
},
},
Location: pulumi.String("location1"),
NetworkProfile: containerservice.ContainerServiceNetworkProfileResponse{
LoadBalancerProfile: interface{}{
ManagedOutboundIPs: &containerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs{
Count: pulumi.Int(2),
},
},
LoadBalancerSku: pulumi.String("standard"),
OutboundType: pulumi.String("loadBalancer"),
},
ResourceGroupName: pulumi.String("rg1"),
ResourceName: pulumi.String("clustername1"),
ServicePrincipalProfile: &containerservice.ManagedClusterServicePrincipalProfileArgs{
ClientId: pulumi.String("clientid"),
Secret: pulumi.String("secret"),
},
Sku: &containerservice.ManagedClusterSKUArgs{
Name: pulumi.String("Basic"),
Tier: pulumi.String("Free"),
},
Tags: pulumi.StringMap{
"archv2": pulumi.String(""),
"tier": pulumi.String("production"),
},
WindowsProfile: &containerservice.ManagedClusterWindowsProfileArgs{
AdminPassword: pulumi.String("replacePassword1234$"),
AdminUsername: pulumi.String("azureuser"),
},
})
if err != nil {
return err
}
return nil
})
}
