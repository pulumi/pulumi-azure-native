


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:servicefabric/v20200101preview:getManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagedClusterResult struct {
	AddonFeatures                []string                             `pulumi:"addonFeatures"`
	AdminPassword                *string                              `pulumi:"adminPassword"`
	AdminUserName                string                               `pulumi:"adminUserName"`
	AzureActiveDirectory         *AzureActiveDirectoryResponse        `pulumi:"azureActiveDirectory"`
	ClientConnectionPort         *int                                 `pulumi:"clientConnectionPort"`
	Clients                      []ClientCertificateResponse          `pulumi:"clients"`
	ClusterCertificateThumbprint string                               `pulumi:"clusterCertificateThumbprint"`
	ClusterCodeVersion           *string                              `pulumi:"clusterCodeVersion"`
	ClusterId                    string                               `pulumi:"clusterId"`
	ClusterState                 string                               `pulumi:"clusterState"`
	DnsName                      string                               `pulumi:"dnsName"`
	Etag                         string                               `pulumi:"etag"`
	FabricSettings               []SettingsSectionDescriptionResponse `pulumi:"fabricSettings"`
	Fqdn                         string                               `pulumi:"fqdn"`
	HttpGatewayConnectionPort    *int                                 `pulumi:"httpGatewayConnectionPort"`
	Id                           string                               `pulumi:"id"`
	LoadBalancingRules           []LoadBalancingRuleResponse          `pulumi:"loadBalancingRules"`
	Location                     string                               `pulumi:"location"`
	Name                         string                               `pulumi:"name"`
	ProvisioningState            string                               `pulumi:"provisioningState"`
	Sku                          *SkuResponse                         `pulumi:"sku"`
	Tags                         map[string]string                    `pulumi:"tags"`
	Type                         string                               `pulumi:"type"`
}
