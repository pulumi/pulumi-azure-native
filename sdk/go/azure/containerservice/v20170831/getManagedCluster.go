


package v20170831

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:containerservice/v20170831:getManagedCluster", args, &rv, opts...)
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
	AgentPoolProfiles       []ContainerServiceAgentPoolProfileResponse       `pulumi:"agentPoolProfiles"`
	DnsPrefix               *string                                          `pulumi:"dnsPrefix"`
	Fqdn                    string                                           `pulumi:"fqdn"`
	Id                      string                                           `pulumi:"id"`
	KubernetesVersion       *string                                          `pulumi:"kubernetesVersion"`
	LinuxProfile            *ContainerServiceLinuxProfileResponse            `pulumi:"linuxProfile"`
	Location                string                                           `pulumi:"location"`
	Name                    string                                           `pulumi:"name"`
	ProvisioningState       string                                           `pulumi:"provisioningState"`
	ServicePrincipalProfile *ContainerServiceServicePrincipalProfileResponse `pulumi:"servicePrincipalProfile"`
	Tags                    map[string]string                                `pulumi:"tags"`
	Type                    string                                           `pulumi:"type"`
}
