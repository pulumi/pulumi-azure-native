


package v20170701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:servicefabric/v20170701preview:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterName       string  `pulumi:"clusterName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SubscriptionId    *string `pulumi:"subscriptionId"`
}


type LookupClusterResult struct {
	AddOnFeatures                   []string                                 `pulumi:"addOnFeatures"`
	AvailableClusterVersions        []ClusterVersionDetailsResponse          `pulumi:"availableClusterVersions"`
	AzureActiveDirectory            *AzureActiveDirectoryResponse            `pulumi:"azureActiveDirectory"`
	Certificate                     *CertificateDescriptionResponse          `pulumi:"certificate"`
	ClientCertificateCommonNames    []ClientCertificateCommonNameResponse    `pulumi:"clientCertificateCommonNames"`
	ClientCertificateThumbprints    []ClientCertificateThumbprintResponse    `pulumi:"clientCertificateThumbprints"`
	ClusterCodeVersion              *string                                  `pulumi:"clusterCodeVersion"`
	ClusterEndpoint                 string                                   `pulumi:"clusterEndpoint"`
	ClusterId                       string                                   `pulumi:"clusterId"`
	ClusterState                    *string                                  `pulumi:"clusterState"`
	DiagnosticsStorageAccountConfig *DiagnosticsStorageAccountConfigResponse `pulumi:"diagnosticsStorageAccountConfig"`
	FabricSettings                  []SettingsSectionDescriptionResponse     `pulumi:"fabricSettings"`
	Id                              string                                   `pulumi:"id"`
	Location                        string                                   `pulumi:"location"`
	ManagementEndpoint              string                                   `pulumi:"managementEndpoint"`
	Name                            string                                   `pulumi:"name"`
	NodeTypes                       []NodeTypeDescriptionResponse            `pulumi:"nodeTypes"`
	ProvisioningState               string                                   `pulumi:"provisioningState"`
	ReliabilityLevel                *string                                  `pulumi:"reliabilityLevel"`
	ReverseProxyCertificate         *CertificateDescriptionResponse          `pulumi:"reverseProxyCertificate"`
	Tags                            map[string]string                        `pulumi:"tags"`
	Type                            string                                   `pulumi:"type"`
	UpgradeDescription              *ClusterUpgradePolicyResponse            `pulumi:"upgradeDescription"`
	UpgradeMode                     *string                                  `pulumi:"upgradeMode"`
	VmImage                         *string                                  `pulumi:"vmImage"`
}
