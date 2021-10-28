


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedCluster struct {
	pulumi.CustomResourceState

	AadProfile              ManagedClusterAADProfileResponsePtrOutput                  `pulumi:"aadProfile"`
	AddonProfiles           ManagedClusterAddonProfileResponseMapOutput                `pulumi:"addonProfiles"`
	AgentPoolProfiles       ManagedClusterAgentPoolProfileResponseArrayOutput          `pulumi:"agentPoolProfiles"`
	ApiServerAccessProfile  ManagedClusterAPIServerAccessProfileResponsePtrOutput      `pulumi:"apiServerAccessProfile"`
	AutoScalerProfile       ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput `pulumi:"autoScalerProfile"`
	AutoUpgradeProfile      ManagedClusterAutoUpgradeProfileResponsePtrOutput          `pulumi:"autoUpgradeProfile"`
	AzurePortalFQDN         pulumi.StringOutput                                        `pulumi:"azurePortalFQDN"`
	DisableLocalAccounts    pulumi.BoolPtrOutput                                       `pulumi:"disableLocalAccounts"`
	DiskEncryptionSetID     pulumi.StringPtrOutput                                     `pulumi:"diskEncryptionSetID"`
	DnsPrefix               pulumi.StringPtrOutput                                     `pulumi:"dnsPrefix"`
	EnablePodSecurityPolicy pulumi.BoolPtrOutput                                       `pulumi:"enablePodSecurityPolicy"`
	EnableRBAC              pulumi.BoolPtrOutput                                       `pulumi:"enableRBAC"`
	ExtendedLocation        ExtendedLocationResponsePtrOutput                          `pulumi:"extendedLocation"`
	Fqdn                    pulumi.StringOutput                                        `pulumi:"fqdn"`
	FqdnSubdomain           pulumi.StringPtrOutput                                     `pulumi:"fqdnSubdomain"`
	HttpProxyConfig         ManagedClusterHTTPProxyConfigResponsePtrOutput             `pulumi:"httpProxyConfig"`
	Identity                ManagedClusterIdentityResponsePtrOutput                    `pulumi:"identity"`
	IdentityProfile         ManagedClusterPropertiesResponseIdentityProfileMapOutput   `pulumi:"identityProfile"`
	KubernetesVersion       pulumi.StringPtrOutput                                     `pulumi:"kubernetesVersion"`
	LinuxProfile            ContainerServiceLinuxProfileResponsePtrOutput              `pulumi:"linuxProfile"`
	Location                pulumi.StringOutput                                        `pulumi:"location"`
	MaxAgentPools           pulumi.IntOutput                                           `pulumi:"maxAgentPools"`
	Name                    pulumi.StringOutput                                        `pulumi:"name"`
	NetworkProfile          ContainerServiceNetworkProfileResponsePtrOutput            `pulumi:"networkProfile"`
	NodeResourceGroup       pulumi.StringPtrOutput                                     `pulumi:"nodeResourceGroup"`
	PodIdentityProfile      ManagedClusterPodIdentityProfileResponsePtrOutput          `pulumi:"podIdentityProfile"`
	PowerState              PowerStateResponseOutput                                   `pulumi:"powerState"`
	PrivateFQDN             pulumi.StringOutput                                        `pulumi:"privateFQDN"`
	PrivateLinkResources    PrivateLinkResourceResponseArrayOutput                     `pulumi:"privateLinkResources"`
	ProvisioningState       pulumi.StringOutput                                        `pulumi:"provisioningState"`
	ServicePrincipalProfile ManagedClusterServicePrincipalProfileResponsePtrOutput     `pulumi:"servicePrincipalProfile"`
	Sku                     ManagedClusterSKUResponsePtrOutput                         `pulumi:"sku"`
	Tags                    pulumi.StringMapOutput                                     `pulumi:"tags"`
	Type                    pulumi.StringOutput                                        `pulumi:"type"`
	WindowsProfile          ManagedClusterWindowsProfileResponsePtrOutput              `pulumi:"windowsProfile"`
}


func NewManagedCluster(ctx *pulumi.Context,
	name string, args *ManagedClusterArgs, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20170831:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20170831:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20180331:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20180331:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20180801preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20180801preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190801:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190801:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191001:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20191001:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20191101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20201101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20201201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210501:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210801:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210901:ManagedCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedCluster
	err := ctx.RegisterResource("azure-native:containerservice/v20210301:ManagedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterState, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	var resource ManagedCluster
	err := ctx.ReadResource("azure-native:containerservice/v20210301:ManagedCluster", name, id, state, &resource, opts...)
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
	AadProfile              *ManagedClusterAADProfile                          `pulumi:"aadProfile"`
	AddonProfiles           map[string]ManagedClusterAddonProfile              `pulumi:"addonProfiles"`
	AgentPoolProfiles       []ManagedClusterAgentPoolProfile                   `pulumi:"agentPoolProfiles"`
	ApiServerAccessProfile  *ManagedClusterAPIServerAccessProfile              `pulumi:"apiServerAccessProfile"`
	AutoScalerProfile       *ManagedClusterPropertiesAutoScalerProfile         `pulumi:"autoScalerProfile"`
	AutoUpgradeProfile      *ManagedClusterAutoUpgradeProfile                  `pulumi:"autoUpgradeProfile"`
	DisableLocalAccounts    *bool                                              `pulumi:"disableLocalAccounts"`
	DiskEncryptionSetID     *string                                            `pulumi:"diskEncryptionSetID"`
	DnsPrefix               *string                                            `pulumi:"dnsPrefix"`
	EnablePodSecurityPolicy *bool                                              `pulumi:"enablePodSecurityPolicy"`
	EnableRBAC              *bool                                              `pulumi:"enableRBAC"`
	ExtendedLocation        *ExtendedLocation                                  `pulumi:"extendedLocation"`
	FqdnSubdomain           *string                                            `pulumi:"fqdnSubdomain"`
	HttpProxyConfig         *ManagedClusterHTTPProxyConfig                     `pulumi:"httpProxyConfig"`
	Identity                *ManagedClusterIdentity                            `pulumi:"identity"`
	IdentityProfile         map[string]ManagedClusterPropertiesIdentityProfile `pulumi:"identityProfile"`
	KubernetesVersion       *string                                            `pulumi:"kubernetesVersion"`
	LinuxProfile            *ContainerServiceLinuxProfile                      `pulumi:"linuxProfile"`
	Location                *string                                            `pulumi:"location"`
	NetworkProfile          *ContainerServiceNetworkProfile                    `pulumi:"networkProfile"`
	NodeResourceGroup       *string                                            `pulumi:"nodeResourceGroup"`
	PodIdentityProfile      *ManagedClusterPodIdentityProfile                  `pulumi:"podIdentityProfile"`
	PrivateLinkResources    []PrivateLinkResource                              `pulumi:"privateLinkResources"`
	ResourceGroupName       string                                             `pulumi:"resourceGroupName"`
	ResourceName            *string                                            `pulumi:"resourceName"`
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfile             `pulumi:"servicePrincipalProfile"`
	Sku                     *ManagedClusterSKU                                 `pulumi:"sku"`
	Tags                    map[string]string                                  `pulumi:"tags"`
	WindowsProfile          *ManagedClusterWindowsProfile                      `pulumi:"windowsProfile"`
}


type ManagedClusterArgs struct {
	AadProfile              ManagedClusterAADProfilePtrInput
	AddonProfiles           ManagedClusterAddonProfileMapInput
	AgentPoolProfiles       ManagedClusterAgentPoolProfileArrayInput
	ApiServerAccessProfile  ManagedClusterAPIServerAccessProfilePtrInput
	AutoScalerProfile       ManagedClusterPropertiesAutoScalerProfilePtrInput
	AutoUpgradeProfile      ManagedClusterAutoUpgradeProfilePtrInput
	DisableLocalAccounts    pulumi.BoolPtrInput
	DiskEncryptionSetID     pulumi.StringPtrInput
	DnsPrefix               pulumi.StringPtrInput
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	EnableRBAC              pulumi.BoolPtrInput
	ExtendedLocation        ExtendedLocationPtrInput
	FqdnSubdomain           pulumi.StringPtrInput
	HttpProxyConfig         ManagedClusterHTTPProxyConfigPtrInput
	Identity                ManagedClusterIdentityPtrInput
	IdentityProfile         ManagedClusterPropertiesIdentityProfileMapInput
	KubernetesVersion       pulumi.StringPtrInput
	LinuxProfile            ContainerServiceLinuxProfilePtrInput
	Location                pulumi.StringPtrInput
	NetworkProfile          ContainerServiceNetworkProfilePtrInput
	NodeResourceGroup       pulumi.StringPtrInput
	PodIdentityProfile      ManagedClusterPodIdentityProfilePtrInput
	PrivateLinkResources    PrivateLinkResourceArrayInput
	ResourceGroupName       pulumi.StringInput
	ResourceName            pulumi.StringPtrInput
	ServicePrincipalProfile ManagedClusterServicePrincipalProfilePtrInput
	Sku                     ManagedClusterSKUPtrInput
	Tags                    pulumi.StringMapInput
	WindowsProfile          ManagedClusterWindowsProfilePtrInput
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
	return reflect.TypeOf((*ManagedCluster)(nil))
}

func (i *ManagedCluster) ToManagedClusterOutput() ManagedClusterOutput {
	return i.ToManagedClusterOutputWithContext(context.Background())
}

func (i *ManagedCluster) ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterOutput)
}

type ManagedClusterOutput struct{ *pulumi.OutputState }

func (ManagedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedCluster)(nil))
}

func (o ManagedClusterOutput) ToManagedClusterOutput() ManagedClusterOutput {
	return o
}

func (o ManagedClusterOutput) ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedClusterInput)(nil)).Elem(), &ManagedCluster{})
	pulumi.RegisterOutputType(ManagedClusterOutput{})
}
