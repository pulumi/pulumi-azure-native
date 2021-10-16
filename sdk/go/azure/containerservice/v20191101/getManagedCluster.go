


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:containerservice/v20191101:getManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedClusterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupManagedClusterResult struct {
	AadProfile              *ManagedClusterAADProfileResponse                          `pulumi:"aadProfile"`
	AddonProfiles           map[string]ManagedClusterAddonProfileResponse              `pulumi:"addonProfiles"`
	AgentPoolProfiles       []ManagedClusterAgentPoolProfileResponse                   `pulumi:"agentPoolProfiles"`
	ApiServerAccessProfile  *ManagedClusterAPIServerAccessProfileResponse              `pulumi:"apiServerAccessProfile"`
	DnsPrefix               *string                                                    `pulumi:"dnsPrefix"`
	EnablePodSecurityPolicy *bool                                                      `pulumi:"enablePodSecurityPolicy"`
	EnableRBAC              *bool                                                      `pulumi:"enableRBAC"`
	Fqdn                    string                                                     `pulumi:"fqdn"`
	Id                      string                                                     `pulumi:"id"`
	Identity                *ManagedClusterIdentityResponse                            `pulumi:"identity"`
	IdentityProfile         map[string]ManagedClusterPropertiesResponseIdentityProfile `pulumi:"identityProfile"`
	KubernetesVersion       *string                                                    `pulumi:"kubernetesVersion"`
	LinuxProfile            *ContainerServiceLinuxProfileResponse                      `pulumi:"linuxProfile"`
	Location                string                                                     `pulumi:"location"`
	MaxAgentPools           int                                                        `pulumi:"maxAgentPools"`
	Name                    string                                                     `pulumi:"name"`
	NetworkProfile          *ContainerServiceNetworkProfileResponse                    `pulumi:"networkProfile"`
	NodeResourceGroup       *string                                                    `pulumi:"nodeResourceGroup"`
	PrivateFQDN             string                                                     `pulumi:"privateFQDN"`
	ProvisioningState       string                                                     `pulumi:"provisioningState"`
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfileResponse             `pulumi:"servicePrincipalProfile"`
	Tags                    map[string]string                                          `pulumi:"tags"`
	Type                    string                                                     `pulumi:"type"`
	WindowsProfile          *ManagedClusterWindowsProfileResponse                      `pulumi:"windowsProfile"`
}
