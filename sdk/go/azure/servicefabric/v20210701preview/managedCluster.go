


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedCluster struct {
	pulumi.CustomResourceState

	AddonFeatures                        pulumi.StringArrayOutput                              `pulumi:"addonFeatures"`
	AdminPassword                        pulumi.StringPtrOutput                                `pulumi:"adminPassword"`
	AdminUserName                        pulumi.StringOutput                                   `pulumi:"adminUserName"`
	AllowRdpAccess                       pulumi.BoolPtrOutput                                  `pulumi:"allowRdpAccess"`
	ApplicationTypeVersionsCleanupPolicy ApplicationTypeVersionsCleanupPolicyResponsePtrOutput `pulumi:"applicationTypeVersionsCleanupPolicy"`
	AzureActiveDirectory                 AzureActiveDirectoryResponsePtrOutput                 `pulumi:"azureActiveDirectory"`
	ClientConnectionPort                 pulumi.IntPtrOutput                                   `pulumi:"clientConnectionPort"`
	Clients                              ClientCertificateResponseArrayOutput                  `pulumi:"clients"`
	ClusterCertificateThumbprints        pulumi.StringArrayOutput                              `pulumi:"clusterCertificateThumbprints"`
	ClusterCodeVersion                   pulumi.StringPtrOutput                                `pulumi:"clusterCodeVersion"`
	ClusterId                            pulumi.StringOutput                                   `pulumi:"clusterId"`
	ClusterState                         pulumi.StringOutput                                   `pulumi:"clusterState"`
	ClusterUpgradeCadence                pulumi.StringPtrOutput                                `pulumi:"clusterUpgradeCadence"`
	ClusterUpgradeMode                   pulumi.StringPtrOutput                                `pulumi:"clusterUpgradeMode"`
	DnsName                              pulumi.StringOutput                                   `pulumi:"dnsName"`
	EnableAutoOSUpgrade                  pulumi.BoolPtrOutput                                  `pulumi:"enableAutoOSUpgrade"`
	EnableIpv6                           pulumi.BoolPtrOutput                                  `pulumi:"enableIpv6"`
	Etag                                 pulumi.StringOutput                                   `pulumi:"etag"`
	FabricSettings                       SettingsSectionDescriptionResponseArrayOutput         `pulumi:"fabricSettings"`
	Fqdn                                 pulumi.StringOutput                                   `pulumi:"fqdn"`
	HttpGatewayConnectionPort            pulumi.IntPtrOutput                                   `pulumi:"httpGatewayConnectionPort"`
	IpTags                               IPTagResponseArrayOutput                              `pulumi:"ipTags"`
	Ipv4Address                          pulumi.StringOutput                                   `pulumi:"ipv4Address"`
	LoadBalancingRules                   LoadBalancingRuleResponseArrayOutput                  `pulumi:"loadBalancingRules"`
	Location                             pulumi.StringOutput                                   `pulumi:"location"`
	Name                                 pulumi.StringOutput                                   `pulumi:"name"`
	NetworkSecurityRules                 NetworkSecurityRuleResponseArrayOutput                `pulumi:"networkSecurityRules"`
	ProvisioningState                    pulumi.StringOutput                                   `pulumi:"provisioningState"`
	Sku                                  SkuResponsePtrOutput                                  `pulumi:"sku"`
	SubnetId                             pulumi.StringPtrOutput                                `pulumi:"subnetId"`
	SystemData                           SystemDataResponseOutput                              `pulumi:"systemData"`
	Tags                                 pulumi.StringMapOutput                                `pulumi:"tags"`
	Type                                 pulumi.StringOutput                                   `pulumi:"type"`
	ZonalResiliency                      pulumi.BoolPtrOutput                                  `pulumi:"zonalResiliency"`
}


func NewManagedCluster(ctx *pulumi.Context,
	name string, args *ManagedClusterArgs, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminUserName == nil {
		return nil, errors.New("invalid value for required argument 'AdminUserName'")
	}
	if args.DnsName == nil {
		return nil, errors.New("invalid value for required argument 'DnsName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.ClientConnectionPort) {
		args.ClientConnectionPort = pulumi.IntPtr(19000)
	}
	if isZero(args.HttpGatewayConnectionPort) {
		args.HttpGatewayConnectionPort = pulumi.IntPtr(19080)
	}
	if isZero(args.ZonalResiliency) {
		args.ZonalResiliency = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200101preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210101preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210501:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210901privatepreview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20211101preview:ManagedCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedCluster
	err := ctx.RegisterResource("azure-native:servicefabric/v20210701preview:ManagedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterState, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	var resource ManagedCluster
	err := ctx.ReadResource("azure-native:servicefabric/v20210701preview:ManagedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedClusterState struct {
}

type ManagedClusterState struct {
}

func (ManagedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterState)(nil)).Elem()
}

type managedClusterArgs struct {
	AddonFeatures                        []string                              `pulumi:"addonFeatures"`
	AdminPassword                        *string                               `pulumi:"adminPassword"`
	AdminUserName                        string                                `pulumi:"adminUserName"`
	AllowRdpAccess                       *bool                                 `pulumi:"allowRdpAccess"`
	ApplicationTypeVersionsCleanupPolicy *ApplicationTypeVersionsCleanupPolicy `pulumi:"applicationTypeVersionsCleanupPolicy"`
	AzureActiveDirectory                 *AzureActiveDirectory                 `pulumi:"azureActiveDirectory"`
	ClientConnectionPort                 *int                                  `pulumi:"clientConnectionPort"`
	Clients                              []ClientCertificate                   `pulumi:"clients"`
	ClusterCodeVersion                   *string                               `pulumi:"clusterCodeVersion"`
	ClusterName                          *string                               `pulumi:"clusterName"`
	ClusterUpgradeCadence                *string                               `pulumi:"clusterUpgradeCadence"`
	ClusterUpgradeMode                   *string                               `pulumi:"clusterUpgradeMode"`
	DnsName                              string                                `pulumi:"dnsName"`
	EnableAutoOSUpgrade                  *bool                                 `pulumi:"enableAutoOSUpgrade"`
	EnableIpv6                           *bool                                 `pulumi:"enableIpv6"`
	FabricSettings                       []SettingsSectionDescription          `pulumi:"fabricSettings"`
	HttpGatewayConnectionPort            *int                                  `pulumi:"httpGatewayConnectionPort"`
	IpTags                               []IPTag                               `pulumi:"ipTags"`
	LoadBalancingRules                   []LoadBalancingRule                   `pulumi:"loadBalancingRules"`
	Location                             *string                               `pulumi:"location"`
	NetworkSecurityRules                 []NetworkSecurityRule                 `pulumi:"networkSecurityRules"`
	ResourceGroupName                    string                                `pulumi:"resourceGroupName"`
	Sku                                  *Sku                                  `pulumi:"sku"`
	SubnetId                             *string                               `pulumi:"subnetId"`
	Tags                                 map[string]string                     `pulumi:"tags"`
	ZonalResiliency                      *bool                                 `pulumi:"zonalResiliency"`
}


type ManagedClusterArgs struct {
	AddonFeatures                        pulumi.StringArrayInput
	AdminPassword                        pulumi.StringPtrInput
	AdminUserName                        pulumi.StringInput
	AllowRdpAccess                       pulumi.BoolPtrInput
	ApplicationTypeVersionsCleanupPolicy ApplicationTypeVersionsCleanupPolicyPtrInput
	AzureActiveDirectory                 AzureActiveDirectoryPtrInput
	ClientConnectionPort                 pulumi.IntPtrInput
	Clients                              ClientCertificateArrayInput
	ClusterCodeVersion                   pulumi.StringPtrInput
	ClusterName                          pulumi.StringPtrInput
	ClusterUpgradeCadence                pulumi.StringPtrInput
	ClusterUpgradeMode                   pulumi.StringPtrInput
	DnsName                              pulumi.StringInput
	EnableAutoOSUpgrade                  pulumi.BoolPtrInput
	EnableIpv6                           pulumi.BoolPtrInput
	FabricSettings                       SettingsSectionDescriptionArrayInput
	HttpGatewayConnectionPort            pulumi.IntPtrInput
	IpTags                               IPTagArrayInput
	LoadBalancingRules                   LoadBalancingRuleArrayInput
	Location                             pulumi.StringPtrInput
	NetworkSecurityRules                 NetworkSecurityRuleArrayInput
	ResourceGroupName                    pulumi.StringInput
	Sku                                  SkuPtrInput
	SubnetId                             pulumi.StringPtrInput
	Tags                                 pulumi.StringMapInput
	ZonalResiliency                      pulumi.BoolPtrInput
}

func (ManagedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterArgs)(nil)).Elem()
}

type ManagedClusterInput interface {
	pulumi.Input

	ToManagedClusterOutput() ManagedClusterOutput
	ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput
}

func (*ManagedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCluster)(nil)).Elem()
}

func (i *ManagedCluster) ToManagedClusterOutput() ManagedClusterOutput {
	return i.ToManagedClusterOutputWithContext(context.Background())
}

func (i *ManagedCluster) ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterOutput)
}

type ManagedClusterOutput struct{ *pulumi.OutputState }

func (ManagedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCluster)(nil)).Elem()
}

func (o ManagedClusterOutput) ToManagedClusterOutput() ManagedClusterOutput {
	return o
}

func (o ManagedClusterOutput) ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterOutput{})
}
