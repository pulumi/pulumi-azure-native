


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:servicefabric/v20210601:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
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


func (val *LookupClusterResult) Defaults() *LookupClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.UpgradeDescription = tmp.UpgradeDescription.Defaults()

	return &tmp
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}


type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) AddOnFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.AddOnFeatures }).(pulumi.StringArrayOutput)
}

func (o LookupClusterResultOutput) ApplicationTypeVersionsCleanupPolicy() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ApplicationTypeVersionsCleanupPolicyResponse {
		return v.ApplicationTypeVersionsCleanupPolicy
	}).(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput)
}

func (o LookupClusterResultOutput) AvailableClusterVersions() ClusterVersionDetailsResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []ClusterVersionDetailsResponse { return v.AvailableClusterVersions }).(ClusterVersionDetailsResponseArrayOutput)
}

func (o LookupClusterResultOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *AzureActiveDirectoryResponse { return v.AzureActiveDirectory }).(AzureActiveDirectoryResponsePtrOutput)
}

func (o LookupClusterResultOutput) Certificate() CertificateDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *CertificateDescriptionResponse { return v.Certificate }).(CertificateDescriptionResponsePtrOutput)
}

func (o LookupClusterResultOutput) CertificateCommonNames() ServerCertificateCommonNamesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ServerCertificateCommonNamesResponse { return v.CertificateCommonNames }).(ServerCertificateCommonNamesResponsePtrOutput)
}

func (o LookupClusterResultOutput) ClientCertificateCommonNames() ClientCertificateCommonNameResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []ClientCertificateCommonNameResponse {
		return v.ClientCertificateCommonNames
	}).(ClientCertificateCommonNameResponseArrayOutput)
}

func (o LookupClusterResultOutput) ClientCertificateThumbprints() ClientCertificateThumbprintResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []ClientCertificateThumbprintResponse {
		return v.ClientCertificateThumbprints
	}).(ClientCertificateThumbprintResponseArrayOutput)
}

func (o LookupClusterResultOutput) ClusterCodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ClusterCodeVersion }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) ClusterEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterEndpoint }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ClusterState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterState }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) DiagnosticsStorageAccountConfig() DiagnosticsStorageAccountConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *DiagnosticsStorageAccountConfigResponse {
		return v.DiagnosticsStorageAccountConfig
	}).(DiagnosticsStorageAccountConfigResponsePtrOutput)
}

func (o LookupClusterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) EventStoreServiceEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EventStoreServiceEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupClusterResultOutput) FabricSettings() SettingsSectionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []SettingsSectionDescriptionResponse { return v.FabricSettings }).(SettingsSectionDescriptionResponseArrayOutput)
}

func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) InfrastructureServiceManager() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.InfrastructureServiceManager }).(pulumi.BoolPtrOutput)
}

func (o LookupClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ManagementEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ManagementEndpoint }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) NodeTypes() NodeTypeDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []NodeTypeDescriptionResponse { return v.NodeTypes }).(NodeTypeDescriptionResponseArrayOutput)
}

func (o LookupClusterResultOutput) Notifications() NotificationResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []NotificationResponse { return v.Notifications }).(NotificationResponseArrayOutput)
}

func (o LookupClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ReliabilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ReliabilityLevel }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) ReverseProxyCertificate() CertificateDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *CertificateDescriptionResponse { return v.ReverseProxyCertificate }).(CertificateDescriptionResponsePtrOutput)
}

func (o LookupClusterResultOutput) ReverseProxyCertificateCommonNames() ServerCertificateCommonNamesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ServerCertificateCommonNamesResponse {
		return v.ReverseProxyCertificateCommonNames
	}).(ServerCertificateCommonNamesResponsePtrOutput)
}

func (o LookupClusterResultOutput) SfZonalUpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.SfZonalUpgradeMode }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) UpgradeDescription() ClusterUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterUpgradePolicyResponse { return v.UpgradeDescription }).(ClusterUpgradePolicyResponsePtrOutput)
}

func (o LookupClusterResultOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.UpgradeMode }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) UpgradePauseEndTimestampUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.UpgradePauseEndTimestampUtc }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) UpgradePauseStartTimestampUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.UpgradePauseStartTimestampUtc }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) UpgradeWave() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.UpgradeWave }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) VmImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.VmImage }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) VmssZonalUpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.VmssZonalUpgradeMode }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) WaveUpgradePaused() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.WaveUpgradePaused }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
