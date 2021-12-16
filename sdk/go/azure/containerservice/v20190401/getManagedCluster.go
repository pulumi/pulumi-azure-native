


package v20190401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:containerservice/v20190401:getManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupManagedClusterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupManagedClusterResult struct {
	AadProfile                  *ManagedClusterAADProfileResponse              `pulumi:"aadProfile"`
	AddonProfiles               map[string]ManagedClusterAddonProfileResponse  `pulumi:"addonProfiles"`
	AgentPoolProfiles           []ManagedClusterAgentPoolProfileResponse       `pulumi:"agentPoolProfiles"`
	ApiServerAuthorizedIPRanges []string                                       `pulumi:"apiServerAuthorizedIPRanges"`
	DnsPrefix                   *string                                        `pulumi:"dnsPrefix"`
	EnablePodSecurityPolicy     *bool                                          `pulumi:"enablePodSecurityPolicy"`
	EnableRBAC                  *bool                                          `pulumi:"enableRBAC"`
	Fqdn                        string                                         `pulumi:"fqdn"`
	Id                          string                                         `pulumi:"id"`
	Identity                    *ManagedClusterIdentityResponse                `pulumi:"identity"`
	KubernetesVersion           *string                                        `pulumi:"kubernetesVersion"`
	LinuxProfile                *ContainerServiceLinuxProfileResponse          `pulumi:"linuxProfile"`
	Location                    string                                         `pulumi:"location"`
	MaxAgentPools               int                                            `pulumi:"maxAgentPools"`
	Name                        string                                         `pulumi:"name"`
	NetworkProfile              *ContainerServiceNetworkProfileResponse        `pulumi:"networkProfile"`
	NodeResourceGroup           *string                                        `pulumi:"nodeResourceGroup"`
	ProvisioningState           string                                         `pulumi:"provisioningState"`
	ServicePrincipalProfile     *ManagedClusterServicePrincipalProfileResponse `pulumi:"servicePrincipalProfile"`
	Tags                        map[string]string                              `pulumi:"tags"`
	Type                        string                                         `pulumi:"type"`
	WindowsProfile              *ManagedClusterWindowsProfileResponse          `pulumi:"windowsProfile"`
}


func (val *LookupManagedClusterResult) Defaults() *LookupManagedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkProfile = tmp.NetworkProfile.Defaults()

	return &tmp
}
