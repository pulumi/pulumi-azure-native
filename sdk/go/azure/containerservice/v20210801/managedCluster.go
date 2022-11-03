


package v20210801

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
	IdentityProfile         UserAssignedIdentityResponseMapOutput                      `pulumi:"identityProfile"`
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
	PublicNetworkAccess     pulumi.StringPtrOutput                                     `pulumi:"publicNetworkAccess"`
	SecurityProfile         ManagedClusterSecurityProfileResponsePtrOutput             `pulumi:"securityProfile"`
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
	if args.NetworkProfile != nil {
		args.NetworkProfile = args.NetworkProfile.ToContainerServiceNetworkProfilePtrOutput().ApplyT(func(v *ContainerServiceNetworkProfile) *ContainerServiceNetworkProfile { return v.Defaults() }).(ContainerServiceNetworkProfilePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20170831:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20180331:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20180801preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190801:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191001:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211001:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211101preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220102preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220302preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220803preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:ManagedCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedCluster
	err := ctx.RegisterResource("azure-native:containerservice/v20210801:ManagedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterState, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	var resource ManagedCluster
	err := ctx.ReadResource("azure-native:containerservice/v20210801:ManagedCluster", name, id, state, &resource, opts...)
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
	AadProfile              *ManagedClusterAADProfile                  `pulumi:"aadProfile"`
	AddonProfiles           map[string]ManagedClusterAddonProfile      `pulumi:"addonProfiles"`
	AgentPoolProfiles       []ManagedClusterAgentPoolProfile           `pulumi:"agentPoolProfiles"`
	ApiServerAccessProfile  *ManagedClusterAPIServerAccessProfile      `pulumi:"apiServerAccessProfile"`
	AutoScalerProfile       *ManagedClusterPropertiesAutoScalerProfile `pulumi:"autoScalerProfile"`
	AutoUpgradeProfile      *ManagedClusterAutoUpgradeProfile          `pulumi:"autoUpgradeProfile"`
	DisableLocalAccounts    *bool                                      `pulumi:"disableLocalAccounts"`
	DiskEncryptionSetID     *string                                    `pulumi:"diskEncryptionSetID"`
	DnsPrefix               *string                                    `pulumi:"dnsPrefix"`
	EnablePodSecurityPolicy *bool                                      `pulumi:"enablePodSecurityPolicy"`
	EnableRBAC              *bool                                      `pulumi:"enableRBAC"`
	ExtendedLocation        *ExtendedLocation                          `pulumi:"extendedLocation"`
	FqdnSubdomain           *string                                    `pulumi:"fqdnSubdomain"`
	HttpProxyConfig         *ManagedClusterHTTPProxyConfig             `pulumi:"httpProxyConfig"`
	Identity                *ManagedClusterIdentity                    `pulumi:"identity"`
	IdentityProfile         map[string]UserAssignedIdentity            `pulumi:"identityProfile"`
	KubernetesVersion       *string                                    `pulumi:"kubernetesVersion"`
	LinuxProfile            *ContainerServiceLinuxProfile              `pulumi:"linuxProfile"`
	Location                *string                                    `pulumi:"location"`
	NetworkProfile          *ContainerServiceNetworkProfile            `pulumi:"networkProfile"`
	NodeResourceGroup       *string                                    `pulumi:"nodeResourceGroup"`
	PodIdentityProfile      *ManagedClusterPodIdentityProfile          `pulumi:"podIdentityProfile"`
	PrivateLinkResources    []PrivateLinkResource                      `pulumi:"privateLinkResources"`
	PublicNetworkAccess     *string                                    `pulumi:"publicNetworkAccess"`
	ResourceGroupName       string                                     `pulumi:"resourceGroupName"`
	ResourceName            *string                                    `pulumi:"resourceName"`
	SecurityProfile         *ManagedClusterSecurityProfile             `pulumi:"securityProfile"`
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfile     `pulumi:"servicePrincipalProfile"`
	Sku                     *ManagedClusterSKU                         `pulumi:"sku"`
	Tags                    map[string]string                          `pulumi:"tags"`
	WindowsProfile          *ManagedClusterWindowsProfile              `pulumi:"windowsProfile"`
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
	IdentityProfile         UserAssignedIdentityMapInput
	KubernetesVersion       pulumi.StringPtrInput
	LinuxProfile            ContainerServiceLinuxProfilePtrInput
	Location                pulumi.StringPtrInput
	NetworkProfile          ContainerServiceNetworkProfilePtrInput
	NodeResourceGroup       pulumi.StringPtrInput
	PodIdentityProfile      ManagedClusterPodIdentityProfilePtrInput
	PrivateLinkResources    PrivateLinkResourceArrayInput
	PublicNetworkAccess     pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	ResourceName            pulumi.StringPtrInput
	SecurityProfile         ManagedClusterSecurityProfilePtrInput
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

func (o ManagedClusterOutput) AadProfile() ManagedClusterAADProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAADProfileResponsePtrOutput { return v.AadProfile }).(ManagedClusterAADProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) AddonProfiles() ManagedClusterAddonProfileResponseMapOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAddonProfileResponseMapOutput { return v.AddonProfiles }).(ManagedClusterAddonProfileResponseMapOutput)
}

func (o ManagedClusterOutput) AgentPoolProfiles() ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAgentPoolProfileResponseArrayOutput { return v.AgentPoolProfiles }).(ManagedClusterAgentPoolProfileResponseArrayOutput)
}

func (o ManagedClusterOutput) ApiServerAccessProfile() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAPIServerAccessProfileResponsePtrOutput {
		return v.ApiServerAccessProfile
	}).(ManagedClusterAPIServerAccessProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) AutoScalerProfile() ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
		return v.AutoScalerProfile
	}).(ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput)
}

func (o ManagedClusterOutput) AutoUpgradeProfile() ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAutoUpgradeProfileResponsePtrOutput { return v.AutoUpgradeProfile }).(ManagedClusterAutoUpgradeProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) AzurePortalFQDN() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.AzurePortalFQDN }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) DisableLocalAccounts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.DisableLocalAccounts }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterOutput) DiskEncryptionSetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.DiskEncryptionSetID }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.DnsPrefix }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterOutput) EnablePodSecurityPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.EnablePodSecurityPolicy }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterOutput) EnableRBAC() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.EnableRBAC }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o ManagedClusterOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) FqdnSubdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.FqdnSubdomain }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterOutput) HttpProxyConfig() ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterHTTPProxyConfigResponsePtrOutput { return v.HttpProxyConfig }).(ManagedClusterHTTPProxyConfigResponsePtrOutput)
}

func (o ManagedClusterOutput) Identity() ManagedClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterIdentityResponsePtrOutput { return v.Identity }).(ManagedClusterIdentityResponsePtrOutput)
}

func (o ManagedClusterOutput) IdentityProfile() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedCluster) UserAssignedIdentityResponseMapOutput { return v.IdentityProfile }).(UserAssignedIdentityResponseMapOutput)
}

func (o ManagedClusterOutput) KubernetesVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.KubernetesVersion }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterOutput) LinuxProfile() ContainerServiceLinuxProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ContainerServiceLinuxProfileResponsePtrOutput { return v.LinuxProfile }).(ContainerServiceLinuxProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) MaxAgentPools() pulumi.IntOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.IntOutput { return v.MaxAgentPools }).(pulumi.IntOutput)
}

func (o ManagedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) NetworkProfile() ContainerServiceNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ContainerServiceNetworkProfileResponsePtrOutput { return v.NetworkProfile }).(ContainerServiceNetworkProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) NodeResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.NodeResourceGroup }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterOutput) PodIdentityProfile() ManagedClusterPodIdentityProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterPodIdentityProfileResponsePtrOutput { return v.PodIdentityProfile }).(ManagedClusterPodIdentityProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) PowerState() PowerStateResponseOutput {
	return o.ApplyT(func(v *ManagedCluster) PowerStateResponseOutput { return v.PowerState }).(PowerStateResponseOutput)
}

func (o ManagedClusterOutput) PrivateFQDN() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.PrivateFQDN }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) PrivateLinkResources() PrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) PrivateLinkResourceResponseArrayOutput { return v.PrivateLinkResources }).(PrivateLinkResourceResponseArrayOutput)
}

func (o ManagedClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterOutput) SecurityProfile() ManagedClusterSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterSecurityProfileResponsePtrOutput { return v.SecurityProfile }).(ManagedClusterSecurityProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) ServicePrincipalProfile() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterServicePrincipalProfileResponsePtrOutput {
		return v.ServicePrincipalProfile
	}).(ManagedClusterServicePrincipalProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) Sku() ManagedClusterSKUResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterSKUResponsePtrOutput { return v.Sku }).(ManagedClusterSKUResponsePtrOutput)
}

func (o ManagedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) WindowsProfile() ManagedClusterWindowsProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterWindowsProfileResponsePtrOutput { return v.WindowsProfile }).(ManagedClusterWindowsProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterOutput{})
}
