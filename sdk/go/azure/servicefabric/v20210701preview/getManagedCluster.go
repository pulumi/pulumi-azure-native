


package v20210701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:servicefabric/v20210701preview:getManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupManagedClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagedClusterResult struct {
	AddonFeatures                        []string                                      `pulumi:"addonFeatures"`
	AdminPassword                        *string                                       `pulumi:"adminPassword"`
	AdminUserName                        string                                        `pulumi:"adminUserName"`
	AllowRdpAccess                       *bool                                         `pulumi:"allowRdpAccess"`
	ApplicationTypeVersionsCleanupPolicy *ApplicationTypeVersionsCleanupPolicyResponse `pulumi:"applicationTypeVersionsCleanupPolicy"`
	AzureActiveDirectory                 *AzureActiveDirectoryResponse                 `pulumi:"azureActiveDirectory"`
	ClientConnectionPort                 *int                                          `pulumi:"clientConnectionPort"`
	Clients                              []ClientCertificateResponse                   `pulumi:"clients"`
	ClusterCertificateThumbprints        []string                                      `pulumi:"clusterCertificateThumbprints"`
	ClusterCodeVersion                   *string                                       `pulumi:"clusterCodeVersion"`
	ClusterId                            string                                        `pulumi:"clusterId"`
	ClusterState                         string                                        `pulumi:"clusterState"`
	ClusterUpgradeCadence                *string                                       `pulumi:"clusterUpgradeCadence"`
	ClusterUpgradeMode                   *string                                       `pulumi:"clusterUpgradeMode"`
	DnsName                              string                                        `pulumi:"dnsName"`
	EnableAutoOSUpgrade                  *bool                                         `pulumi:"enableAutoOSUpgrade"`
	EnableIpv6                           *bool                                         `pulumi:"enableIpv6"`
	Etag                                 string                                        `pulumi:"etag"`
	FabricSettings                       []SettingsSectionDescriptionResponse          `pulumi:"fabricSettings"`
	Fqdn                                 string                                        `pulumi:"fqdn"`
	HttpGatewayConnectionPort            *int                                          `pulumi:"httpGatewayConnectionPort"`
	Id                                   string                                        `pulumi:"id"`
	IpTags                               []IPTagResponse                               `pulumi:"ipTags"`
	Ipv4Address                          string                                        `pulumi:"ipv4Address"`
	LoadBalancingRules                   []LoadBalancingRuleResponse                   `pulumi:"loadBalancingRules"`
	Location                             string                                        `pulumi:"location"`
	Name                                 string                                        `pulumi:"name"`
	NetworkSecurityRules                 []NetworkSecurityRuleResponse                 `pulumi:"networkSecurityRules"`
	ProvisioningState                    string                                        `pulumi:"provisioningState"`
	Sku                                  *SkuResponse                                  `pulumi:"sku"`
	SubnetId                             *string                                       `pulumi:"subnetId"`
	SystemData                           SystemDataResponse                            `pulumi:"systemData"`
	Tags                                 map[string]string                             `pulumi:"tags"`
	Type                                 string                                        `pulumi:"type"`
	ZonalResiliency                      *bool                                         `pulumi:"zonalResiliency"`
}


func (val *LookupManagedClusterResult) Defaults() *LookupManagedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientConnectionPort) {
		clientConnectionPort_ := 19000
		tmp.ClientConnectionPort = &clientConnectionPort_
	}
	if isZero(tmp.HttpGatewayConnectionPort) {
		httpGatewayConnectionPort_ := 19080
		tmp.HttpGatewayConnectionPort = &httpGatewayConnectionPort_
	}
	if isZero(tmp.ZonalResiliency) {
		zonalResiliency_ := false
		tmp.ZonalResiliency = &zonalResiliency_
	}
	return &tmp
}
