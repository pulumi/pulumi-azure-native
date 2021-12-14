


package v20180331

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:containerservice/v20180331:getManagedCluster", args, &rv, opts...)
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
	AadProfile              *ManagedClusterAADProfileResponse              `pulumi:"aadProfile"`
	AddonProfiles           map[string]ManagedClusterAddonProfileResponse  `pulumi:"addonProfiles"`
	AgentPoolProfiles       []ManagedClusterAgentPoolProfileResponse       `pulumi:"agentPoolProfiles"`
	DnsPrefix               *string                                        `pulumi:"dnsPrefix"`
	EnableRBAC              *bool                                          `pulumi:"enableRBAC"`
	Fqdn                    string                                         `pulumi:"fqdn"`
	Id                      string                                         `pulumi:"id"`
	KubernetesVersion       *string                                        `pulumi:"kubernetesVersion"`
	LinuxProfile            *ContainerServiceLinuxProfileResponse          `pulumi:"linuxProfile"`
	Location                string                                         `pulumi:"location"`
	Name                    string                                         `pulumi:"name"`
	NetworkProfile          *ContainerServiceNetworkProfileResponse        `pulumi:"networkProfile"`
	NodeResourceGroup       string                                         `pulumi:"nodeResourceGroup"`
	ProvisioningState       string                                         `pulumi:"provisioningState"`
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfileResponse `pulumi:"servicePrincipalProfile"`
	Tags                    map[string]string                              `pulumi:"tags"`
	Type                    string                                         `pulumi:"type"`
}


func (val *LookupManagedClusterResult) Defaults() *LookupManagedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkProfile = tmp.NetworkProfile.Defaults()

	return &tmp
}
