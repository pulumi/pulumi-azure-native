


package v20170701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Cluster struct {
	pulumi.CustomResourceState

	AddOnFeatures                   pulumi.StringArrayOutput                         `pulumi:"addOnFeatures"`
	AvailableClusterVersions        ClusterVersionDetailsResponseArrayOutput         `pulumi:"availableClusterVersions"`
	AzureActiveDirectory            AzureActiveDirectoryResponsePtrOutput            `pulumi:"azureActiveDirectory"`
	Certificate                     CertificateDescriptionResponsePtrOutput          `pulumi:"certificate"`
	ClientCertificateCommonNames    ClientCertificateCommonNameResponseArrayOutput   `pulumi:"clientCertificateCommonNames"`
	ClientCertificateThumbprints    ClientCertificateThumbprintResponseArrayOutput   `pulumi:"clientCertificateThumbprints"`
	ClusterCodeVersion              pulumi.StringPtrOutput                           `pulumi:"clusterCodeVersion"`
	ClusterEndpoint                 pulumi.StringOutput                              `pulumi:"clusterEndpoint"`
	ClusterId                       pulumi.StringOutput                              `pulumi:"clusterId"`
	ClusterState                    pulumi.StringPtrOutput                           `pulumi:"clusterState"`
	DiagnosticsStorageAccountConfig DiagnosticsStorageAccountConfigResponsePtrOutput `pulumi:"diagnosticsStorageAccountConfig"`
	FabricSettings                  SettingsSectionDescriptionResponseArrayOutput    `pulumi:"fabricSettings"`
	Location                        pulumi.StringOutput                              `pulumi:"location"`
	ManagementEndpoint              pulumi.StringOutput                              `pulumi:"managementEndpoint"`
	Name                            pulumi.StringOutput                              `pulumi:"name"`
	NodeTypes                       NodeTypeDescriptionResponseArrayOutput           `pulumi:"nodeTypes"`
	ProvisioningState               pulumi.StringOutput                              `pulumi:"provisioningState"`
	ReliabilityLevel                pulumi.StringPtrOutput                           `pulumi:"reliabilityLevel"`
	ReverseProxyCertificate         CertificateDescriptionResponsePtrOutput          `pulumi:"reverseProxyCertificate"`
	Tags                            pulumi.StringMapOutput                           `pulumi:"tags"`
	Type                            pulumi.StringOutput                              `pulumi:"type"`
	UpgradeDescription              ClusterUpgradePolicyResponsePtrOutput            `pulumi:"upgradeDescription"`
	UpgradeMode                     pulumi.StringPtrOutput                           `pulumi:"upgradeMode"`
	VmImage                         pulumi.StringPtrOutput                           `pulumi:"vmImage"`
}


func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ManagementEndpoint'")
	}
	if args.NodeTypes == nil {
		return nil, errors.New("invalid value for required argument 'NodeTypes'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20160901:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20180201:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190601preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20191101preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200301:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20201201preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210601:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:servicefabric/v20170701preview:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:servicefabric/v20170701preview:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	AddOnFeatures                   []string                         `pulumi:"addOnFeatures"`
	AvailableClusterVersions        []ClusterVersionDetails          `pulumi:"availableClusterVersions"`
	AzureActiveDirectory            *AzureActiveDirectory            `pulumi:"azureActiveDirectory"`
	Certificate                     *CertificateDescription          `pulumi:"certificate"`
	ClientCertificateCommonNames    []ClientCertificateCommonName    `pulumi:"clientCertificateCommonNames"`
	ClientCertificateThumbprints    []ClientCertificateThumbprint    `pulumi:"clientCertificateThumbprints"`
	ClusterCodeVersion              *string                          `pulumi:"clusterCodeVersion"`
	ClusterName                     *string                          `pulumi:"clusterName"`
	ClusterState                    *string                          `pulumi:"clusterState"`
	DiagnosticsStorageAccountConfig *DiagnosticsStorageAccountConfig `pulumi:"diagnosticsStorageAccountConfig"`
	FabricSettings                  []SettingsSectionDescription     `pulumi:"fabricSettings"`
	Location                        *string                          `pulumi:"location"`
	ManagementEndpoint              string                           `pulumi:"managementEndpoint"`
	NodeTypes                       []NodeTypeDescription            `pulumi:"nodeTypes"`
	ReliabilityLevel                *string                          `pulumi:"reliabilityLevel"`
	ResourceGroupName               string                           `pulumi:"resourceGroupName"`
	ReverseProxyCertificate         *CertificateDescription          `pulumi:"reverseProxyCertificate"`
	SubscriptionId                  *string                          `pulumi:"subscriptionId"`
	Tags                            map[string]string                `pulumi:"tags"`
	UpgradeDescription              *ClusterUpgradePolicy            `pulumi:"upgradeDescription"`
	UpgradeMode                     *string                          `pulumi:"upgradeMode"`
	VmImage                         *string                          `pulumi:"vmImage"`
}


type ClusterArgs struct {
	AddOnFeatures                   pulumi.StringArrayInput
	AvailableClusterVersions        ClusterVersionDetailsArrayInput
	AzureActiveDirectory            AzureActiveDirectoryPtrInput
	Certificate                     CertificateDescriptionPtrInput
	ClientCertificateCommonNames    ClientCertificateCommonNameArrayInput
	ClientCertificateThumbprints    ClientCertificateThumbprintArrayInput
	ClusterCodeVersion              pulumi.StringPtrInput
	ClusterName                     pulumi.StringPtrInput
	ClusterState                    pulumi.StringPtrInput
	DiagnosticsStorageAccountConfig DiagnosticsStorageAccountConfigPtrInput
	FabricSettings                  SettingsSectionDescriptionArrayInput
	Location                        pulumi.StringPtrInput
	ManagementEndpoint              pulumi.StringInput
	NodeTypes                       NodeTypeDescriptionArrayInput
	ReliabilityLevel                pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
	ReverseProxyCertificate         CertificateDescriptionPtrInput
	SubscriptionId                  pulumi.StringPtrInput
	Tags                            pulumi.StringMapInput
	UpgradeDescription              ClusterUpgradePolicyPtrInput
	UpgradeMode                     pulumi.StringPtrInput
	VmImage                         pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) AddOnFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.AddOnFeatures }).(pulumi.StringArrayOutput)
}

func (o ClusterOutput) AvailableClusterVersions() ClusterVersionDetailsResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterVersionDetailsResponseArrayOutput { return v.AvailableClusterVersions }).(ClusterVersionDetailsResponseArrayOutput)
}

func (o ClusterOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) AzureActiveDirectoryResponsePtrOutput { return v.AzureActiveDirectory }).(AzureActiveDirectoryResponsePtrOutput)
}

func (o ClusterOutput) Certificate() CertificateDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) CertificateDescriptionResponsePtrOutput { return v.Certificate }).(CertificateDescriptionResponsePtrOutput)
}

func (o ClusterOutput) ClientCertificateCommonNames() ClientCertificateCommonNameResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) ClientCertificateCommonNameResponseArrayOutput { return v.ClientCertificateCommonNames }).(ClientCertificateCommonNameResponseArrayOutput)
}

func (o ClusterOutput) ClientCertificateThumbprints() ClientCertificateThumbprintResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) ClientCertificateThumbprintResponseArrayOutput { return v.ClientCertificateThumbprints }).(ClientCertificateThumbprintResponseArrayOutput)
}

func (o ClusterOutput) ClusterCodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ClusterCodeVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) ClusterEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterEndpoint }).(pulumi.StringOutput)
}

func (o ClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

func (o ClusterOutput) ClusterState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ClusterState }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) DiagnosticsStorageAccountConfig() DiagnosticsStorageAccountConfigResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) DiagnosticsStorageAccountConfigResponsePtrOutput {
		return v.DiagnosticsStorageAccountConfig
	}).(DiagnosticsStorageAccountConfigResponsePtrOutput)
}

func (o ClusterOutput) FabricSettings() SettingsSectionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) SettingsSectionDescriptionResponseArrayOutput { return v.FabricSettings }).(SettingsSectionDescriptionResponseArrayOutput)
}

func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ClusterOutput) ManagementEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ManagementEndpoint }).(pulumi.StringOutput)
}

func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterOutput) NodeTypes() NodeTypeDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) NodeTypeDescriptionResponseArrayOutput { return v.NodeTypes }).(NodeTypeDescriptionResponseArrayOutput)
}

func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ClusterOutput) ReliabilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ReliabilityLevel }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) ReverseProxyCertificate() CertificateDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) CertificateDescriptionResponsePtrOutput { return v.ReverseProxyCertificate }).(CertificateDescriptionResponsePtrOutput)
}

func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ClusterOutput) UpgradeDescription() ClusterUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterUpgradePolicyResponsePtrOutput { return v.UpgradeDescription }).(ClusterUpgradePolicyResponsePtrOutput)
}

func (o ClusterOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.UpgradeMode }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) VmImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.VmImage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
