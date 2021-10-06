


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:servicefabric/v20210601:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupClusterResult struct {
	AddOnFeatures                        []string                                      `pulumi:"addOnFeatures"`
	ApplicationTypeVersionsCleanupPolicy *ApplicationTypeVersionsCleanupPolicyResponse `pulumi:"applicationTypeVersionsCleanupPolicy"`
	AvailableClusterVersions             []ClusterVersionDetailsResponse               `pulumi:"availableClusterVersions"`
	AzureActiveDirectory                 *AzureActiveDirectoryResponse                 `pulumi:"azureActiveDirectory"`
	Certificate                          *CertificateDescriptionResponse               `pulumi:"certificate"`
	CertificateCommonNames               *ServerCertificateCommonNamesResponse         `pulumi:"certificateCommonNames"`
	ClientCertificateCommonNames         []ClientCertificateCommonNameResponse         `pulumi:"clientCertificateCommonNames"`
	ClientCertificateThumbprints         []ClientCertificateThumbprintResponse         `pulumi:"clientCertificateThumbprints"`
	ClusterCodeVersion                   *string                                       `pulumi:"clusterCodeVersion"`
	ClusterEndpoint                      string                                        `pulumi:"clusterEndpoint"`
	ClusterId                            string                                        `pulumi:"clusterId"`
	ClusterState                         string                                        `pulumi:"clusterState"`
	DiagnosticsStorageAccountConfig      *DiagnosticsStorageAccountConfigResponse      `pulumi:"diagnosticsStorageAccountConfig"`
	Etag                                 string                                        `pulumi:"etag"`
	EventStoreServiceEnabled             *bool                                         `pulumi:"eventStoreServiceEnabled"`
	FabricSettings                       []SettingsSectionDescriptionResponse          `pulumi:"fabricSettings"`
	Id                                   string                                        `pulumi:"id"`
	InfrastructureServiceManager         *bool                                         `pulumi:"infrastructureServiceManager"`
	Location                             string                                        `pulumi:"location"`
	ManagementEndpoint                   string                                        `pulumi:"managementEndpoint"`
	Name                                 string                                        `pulumi:"name"`
	NodeTypes                            []NodeTypeDescriptionResponse                 `pulumi:"nodeTypes"`
	Notifications                        []NotificationResponse                        `pulumi:"notifications"`
	ProvisioningState                    string                                        `pulumi:"provisioningState"`
	ReliabilityLevel                     *string                                       `pulumi:"reliabilityLevel"`
	ReverseProxyCertificate              *CertificateDescriptionResponse               `pulumi:"reverseProxyCertificate"`
	ReverseProxyCertificateCommonNames   *ServerCertificateCommonNamesResponse         `pulumi:"reverseProxyCertificateCommonNames"`
	SfZonalUpgradeMode                   *string                                       `pulumi:"sfZonalUpgradeMode"`
	SystemData                           SystemDataResponse                            `pulumi:"systemData"`
	Tags                                 map[string]string                             `pulumi:"tags"`
	Type                                 string                                        `pulumi:"type"`
	UpgradeDescription                   *ClusterUpgradePolicyResponse                 `pulumi:"upgradeDescription"`
	UpgradeMode                          *string                                       `pulumi:"upgradeMode"`
	UpgradePauseEndTimestampUtc          *string                                       `pulumi:"upgradePauseEndTimestampUtc"`
	UpgradePauseStartTimestampUtc        *string                                       `pulumi:"upgradePauseStartTimestampUtc"`
	UpgradeWave                          *string                                       `pulumi:"upgradeWave"`
	VmImage                              *string                                       `pulumi:"vmImage"`
	VmssZonalUpgradeMode                 *string                                       `pulumi:"vmssZonalUpgradeMode"`
	WaveUpgradePaused                    *bool                                         `pulumi:"waveUpgradePaused"`
}
