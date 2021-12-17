


package v20180201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	AddOnFeatures                      pulumi.StringArrayOutput                         `pulumi:"addOnFeatures"`
	AvailableClusterVersions           ClusterVersionDetailsResponseArrayOutput         `pulumi:"availableClusterVersions"`
	AzureActiveDirectory               AzureActiveDirectoryResponsePtrOutput            `pulumi:"azureActiveDirectory"`
	Certificate                        CertificateDescriptionResponsePtrOutput          `pulumi:"certificate"`
	CertificateCommonNames             ServerCertificateCommonNamesResponsePtrOutput    `pulumi:"certificateCommonNames"`
	ClientCertificateCommonNames       ClientCertificateCommonNameResponseArrayOutput   `pulumi:"clientCertificateCommonNames"`
	ClientCertificateThumbprints       ClientCertificateThumbprintResponseArrayOutput   `pulumi:"clientCertificateThumbprints"`
	ClusterCodeVersion                 pulumi.StringPtrOutput                           `pulumi:"clusterCodeVersion"`
	ClusterEndpoint                    pulumi.StringOutput                              `pulumi:"clusterEndpoint"`
	ClusterId                          pulumi.StringOutput                              `pulumi:"clusterId"`
	ClusterState                       pulumi.StringOutput                              `pulumi:"clusterState"`
	DiagnosticsStorageAccountConfig    DiagnosticsStorageAccountConfigResponsePtrOutput `pulumi:"diagnosticsStorageAccountConfig"`
	FabricSettings                     SettingsSectionDescriptionResponseArrayOutput    `pulumi:"fabricSettings"`
	Location                           pulumi.StringOutput                              `pulumi:"location"`
	ManagementEndpoint                 pulumi.StringOutput                              `pulumi:"managementEndpoint"`
	Name                               pulumi.StringOutput                              `pulumi:"name"`
	NodeTypes                          NodeTypeDescriptionResponseArrayOutput           `pulumi:"nodeTypes"`
	ProvisioningState                  pulumi.StringOutput                              `pulumi:"provisioningState"`
	ReliabilityLevel                   pulumi.StringPtrOutput                           `pulumi:"reliabilityLevel"`
	ReverseProxyCertificate            CertificateDescriptionResponsePtrOutput          `pulumi:"reverseProxyCertificate"`
	ReverseProxyCertificateCommonNames ServerCertificateCommonNamesResponsePtrOutput    `pulumi:"reverseProxyCertificateCommonNames"`
	Tags                               pulumi.StringMapOutput                           `pulumi:"tags"`
	Type                               pulumi.StringOutput                              `pulumi:"type"`
	UpgradeDescription                 ClusterUpgradePolicyResponsePtrOutput            `pulumi:"upgradeDescription"`
	UpgradeMode                        pulumi.StringPtrOutput                           `pulumi:"upgradeMode"`
	VmImage                            pulumi.StringPtrOutput                           `pulumi:"vmImage"`
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
	upgradeDescriptionApplier := func(v ClusterUpgradePolicy) *ClusterUpgradePolicy { return v.Defaults() }
	if args.UpgradeDescription != nil {
		args.UpgradeDescription = args.UpgradeDescription.ToClusterUpgradePolicyPtrOutput().Elem().ApplyT(upgradeDescriptionApplier).(ClusterUpgradePolicyPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20160901:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20170701preview:Cluster"),
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
	err := ctx.RegisterResource("azure-native:servicefabric/v20180201:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:servicefabric/v20180201:Cluster", name, id, state, &resource, opts...)
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
	AddOnFeatures                      []string                         `pulumi:"addOnFeatures"`
	AzureActiveDirectory               *AzureActiveDirectory            `pulumi:"azureActiveDirectory"`
	Certificate                        *CertificateDescription          `pulumi:"certificate"`
	CertificateCommonNames             *ServerCertificateCommonNames    `pulumi:"certificateCommonNames"`
	ClientCertificateCommonNames       []ClientCertificateCommonName    `pulumi:"clientCertificateCommonNames"`
	ClientCertificateThumbprints       []ClientCertificateThumbprint    `pulumi:"clientCertificateThumbprints"`
	ClusterCodeVersion                 *string                          `pulumi:"clusterCodeVersion"`
	ClusterName                        *string                          `pulumi:"clusterName"`
	DiagnosticsStorageAccountConfig    *DiagnosticsStorageAccountConfig `pulumi:"diagnosticsStorageAccountConfig"`
	FabricSettings                     []SettingsSectionDescription     `pulumi:"fabricSettings"`
	Location                           *string                          `pulumi:"location"`
	ManagementEndpoint                 string                           `pulumi:"managementEndpoint"`
	NodeTypes                          []NodeTypeDescription            `pulumi:"nodeTypes"`
	ReliabilityLevel                   *string                          `pulumi:"reliabilityLevel"`
	ResourceGroupName                  string                           `pulumi:"resourceGroupName"`
	ReverseProxyCertificate            *CertificateDescription          `pulumi:"reverseProxyCertificate"`
	ReverseProxyCertificateCommonNames *ServerCertificateCommonNames    `pulumi:"reverseProxyCertificateCommonNames"`
	Tags                               map[string]string                `pulumi:"tags"`
	UpgradeDescription                 *ClusterUpgradePolicy            `pulumi:"upgradeDescription"`
	UpgradeMode                        *string                          `pulumi:"upgradeMode"`
	VmImage                            *string                          `pulumi:"vmImage"`
}


type ClusterArgs struct {
	AddOnFeatures                      pulumi.StringArrayInput
	AzureActiveDirectory               AzureActiveDirectoryPtrInput
	Certificate                        CertificateDescriptionPtrInput
	CertificateCommonNames             ServerCertificateCommonNamesPtrInput
	ClientCertificateCommonNames       ClientCertificateCommonNameArrayInput
	ClientCertificateThumbprints       ClientCertificateThumbprintArrayInput
	ClusterCodeVersion                 pulumi.StringPtrInput
	ClusterName                        pulumi.StringPtrInput
	DiagnosticsStorageAccountConfig    DiagnosticsStorageAccountConfigPtrInput
	FabricSettings                     SettingsSectionDescriptionArrayInput
	Location                           pulumi.StringPtrInput
	ManagementEndpoint                 pulumi.StringInput
	NodeTypes                          NodeTypeDescriptionArrayInput
	ReliabilityLevel                   pulumi.StringPtrInput
	ResourceGroupName                  pulumi.StringInput
	ReverseProxyCertificate            CertificateDescriptionPtrInput
	ReverseProxyCertificateCommonNames ServerCertificateCommonNamesPtrInput
	Tags                               pulumi.StringMapInput
	UpgradeDescription                 ClusterUpgradePolicyPtrInput
	UpgradeMode                        pulumi.StringPtrInput
	VmImage                            pulumi.StringPtrInput
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

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
