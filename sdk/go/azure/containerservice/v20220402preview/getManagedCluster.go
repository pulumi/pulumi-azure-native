


package v20220402preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:containerservice/v20220402preview:getManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupManagedClusterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupManagedClusterResult struct {
	AadProfile               *ManagedClusterAADProfileResponse                  `pulumi:"aadProfile"`
	AddonProfiles            map[string]ManagedClusterAddonProfileResponse      `pulumi:"addonProfiles"`
	AgentPoolProfiles        []ManagedClusterAgentPoolProfileResponse           `pulumi:"agentPoolProfiles"`
	ApiServerAccessProfile   *ManagedClusterAPIServerAccessProfileResponse      `pulumi:"apiServerAccessProfile"`
	AutoScalerProfile        *ManagedClusterPropertiesResponseAutoScalerProfile `pulumi:"autoScalerProfile"`
	AutoUpgradeProfile       *ManagedClusterAutoUpgradeProfileResponse          `pulumi:"autoUpgradeProfile"`
	AzurePortalFQDN          string                                             `pulumi:"azurePortalFQDN"`
	CreationData             *CreationDataResponse                              `pulumi:"creationData"`
	CurrentKubernetesVersion string                                             `pulumi:"currentKubernetesVersion"`
	DisableLocalAccounts     *bool                                              `pulumi:"disableLocalAccounts"`
	DiskEncryptionSetID      *string                                            `pulumi:"diskEncryptionSetID"`
	DnsPrefix                *string                                            `pulumi:"dnsPrefix"`
	EnableNamespaceResources *bool                                              `pulumi:"enableNamespaceResources"`
	EnablePodSecurityPolicy  *bool                                              `pulumi:"enablePodSecurityPolicy"`
	EnableRBAC               *bool                                              `pulumi:"enableRBAC"`
	ExtendedLocation         *ExtendedLocationResponse                          `pulumi:"extendedLocation"`
	Fqdn                     string                                             `pulumi:"fqdn"`
	FqdnSubdomain            *string                                            `pulumi:"fqdnSubdomain"`
	HttpProxyConfig          *ManagedClusterHTTPProxyConfigResponse             `pulumi:"httpProxyConfig"`
	Id                       string                                             `pulumi:"id"`
	Identity                 *ManagedClusterIdentityResponse                    `pulumi:"identity"`
	IdentityProfile          map[string]UserAssignedIdentityResponse            `pulumi:"identityProfile"`
	IngressProfile           *ManagedClusterIngressProfileResponse              `pulumi:"ingressProfile"`
	KubernetesVersion        *string                                            `pulumi:"kubernetesVersion"`
	LinuxProfile             *ContainerServiceLinuxProfileResponse              `pulumi:"linuxProfile"`
	Location                 string                                             `pulumi:"location"`
	MaxAgentPools            int                                                `pulumi:"maxAgentPools"`
	Name                     string                                             `pulumi:"name"`
	NetworkProfile           *ContainerServiceNetworkProfileResponse            `pulumi:"networkProfile"`
	NodeResourceGroup        *string                                            `pulumi:"nodeResourceGroup"`
	OidcIssuerProfile        *ManagedClusterOIDCIssuerProfileResponse           `pulumi:"oidcIssuerProfile"`
	PodIdentityProfile       *ManagedClusterPodIdentityProfileResponse          `pulumi:"podIdentityProfile"`
	PowerState               PowerStateResponse                                 `pulumi:"powerState"`
	PrivateFQDN              string                                             `pulumi:"privateFQDN"`
	PrivateLinkResources     []PrivateLinkResourceResponse                      `pulumi:"privateLinkResources"`
	ProvisioningState        string                                             `pulumi:"provisioningState"`
	PublicNetworkAccess      *string                                            `pulumi:"publicNetworkAccess"`
	SecurityProfile          *ManagedClusterSecurityProfileResponse             `pulumi:"securityProfile"`
	ServicePrincipalProfile  *ManagedClusterServicePrincipalProfileResponse     `pulumi:"servicePrincipalProfile"`
	Sku                      *ManagedClusterSKUResponse                         `pulumi:"sku"`
	StorageProfile           *ManagedClusterStorageProfileResponse              `pulumi:"storageProfile"`
	SystemData               SystemDataResponse                                 `pulumi:"systemData"`
	Tags                     map[string]string                                  `pulumi:"tags"`
	Type                     string                                             `pulumi:"type"`
	WindowsProfile           *ManagedClusterWindowsProfileResponse              `pulumi:"windowsProfile"`
}


func (val *LookupManagedClusterResult) Defaults() *LookupManagedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkProfile = tmp.NetworkProfile.Defaults()

	return &tmp
}

func LookupManagedClusterOutput(ctx *pulumi.Context, args LookupManagedClusterOutputArgs, opts ...pulumi.InvokeOption) LookupManagedClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedClusterResult, error) {
			args := v.(LookupManagedClusterArgs)
			r, err := LookupManagedCluster(ctx, &args, opts...)
			var s LookupManagedClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedClusterResultOutput)
}

type LookupManagedClusterOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupManagedClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterArgs)(nil)).Elem()
}


type LookupManagedClusterResultOutput struct{ *pulumi.OutputState }

func (LookupManagedClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterResult)(nil)).Elem()
}

func (o LookupManagedClusterResultOutput) ToLookupManagedClusterResultOutput() LookupManagedClusterResultOutput {
	return o
}

func (o LookupManagedClusterResultOutput) ToLookupManagedClusterResultOutputWithContext(ctx context.Context) LookupManagedClusterResultOutput {
	return o
}

func (o LookupManagedClusterResultOutput) AadProfile() ManagedClusterAADProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterAADProfileResponse { return v.AadProfile }).(ManagedClusterAADProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) AddonProfiles() ManagedClusterAddonProfileResponseMapOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) map[string]ManagedClusterAddonProfileResponse {
		return v.AddonProfiles
	}).(ManagedClusterAddonProfileResponseMapOutput)
}

func (o LookupManagedClusterResultOutput) AgentPoolProfiles() ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []ManagedClusterAgentPoolProfileResponse {
		return v.AgentPoolProfiles
	}).(ManagedClusterAgentPoolProfileResponseArrayOutput)
}

func (o LookupManagedClusterResultOutput) ApiServerAccessProfile() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterAPIServerAccessProfileResponse {
		return v.ApiServerAccessProfile
	}).(ManagedClusterAPIServerAccessProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) AutoScalerProfile() ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterPropertiesResponseAutoScalerProfile {
		return v.AutoScalerProfile
	}).(ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput)
}

func (o LookupManagedClusterResultOutput) AutoUpgradeProfile() ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterAutoUpgradeProfileResponse {
		return v.AutoUpgradeProfile
	}).(ManagedClusterAutoUpgradeProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) AzurePortalFQDN() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.AzurePortalFQDN }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) CreationData() CreationDataResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *CreationDataResponse { return v.CreationData }).(CreationDataResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) CurrentKubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.CurrentKubernetesVersion }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) DisableLocalAccounts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.DisableLocalAccounts }).(pulumi.BoolPtrOutput)
}

func (o LookupManagedClusterResultOutput) DiskEncryptionSetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.DiskEncryptionSetID }).(pulumi.StringPtrOutput)
}

func (o LookupManagedClusterResultOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.DnsPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupManagedClusterResultOutput) EnableNamespaceResources() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.EnableNamespaceResources }).(pulumi.BoolPtrOutput)
}

func (o LookupManagedClusterResultOutput) EnablePodSecurityPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.EnablePodSecurityPolicy }).(pulumi.BoolPtrOutput)
}

func (o LookupManagedClusterResultOutput) EnableRBAC() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.EnableRBAC }).(pulumi.BoolPtrOutput)
}

func (o LookupManagedClusterResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) FqdnSubdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.FqdnSubdomain }).(pulumi.StringPtrOutput)
}

func (o LookupManagedClusterResultOutput) HttpProxyConfig() ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterHTTPProxyConfigResponse { return v.HttpProxyConfig }).(ManagedClusterHTTPProxyConfigResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) Identity() ManagedClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterIdentityResponse { return v.Identity }).(ManagedClusterIdentityResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) IdentityProfile() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) map[string]UserAssignedIdentityResponse { return v.IdentityProfile }).(UserAssignedIdentityResponseMapOutput)
}

func (o LookupManagedClusterResultOutput) IngressProfile() ManagedClusterIngressProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterIngressProfileResponse { return v.IngressProfile }).(ManagedClusterIngressProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) KubernetesVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.KubernetesVersion }).(pulumi.StringPtrOutput)
}

func (o LookupManagedClusterResultOutput) LinuxProfile() ContainerServiceLinuxProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ContainerServiceLinuxProfileResponse { return v.LinuxProfile }).(ContainerServiceLinuxProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) MaxAgentPools() pulumi.IntOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) int { return v.MaxAgentPools }).(pulumi.IntOutput)
}

func (o LookupManagedClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) NetworkProfile() ContainerServiceNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ContainerServiceNetworkProfileResponse { return v.NetworkProfile }).(ContainerServiceNetworkProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) NodeResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.NodeResourceGroup }).(pulumi.StringPtrOutput)
}

func (o LookupManagedClusterResultOutput) OidcIssuerProfile() ManagedClusterOIDCIssuerProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterOIDCIssuerProfileResponse {
		return v.OidcIssuerProfile
	}).(ManagedClusterOIDCIssuerProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) PodIdentityProfile() ManagedClusterPodIdentityProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterPodIdentityProfileResponse {
		return v.PodIdentityProfile
	}).(ManagedClusterPodIdentityProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) PowerState() PowerStateResponseOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) PowerStateResponse { return v.PowerState }).(PowerStateResponseOutput)
}

func (o LookupManagedClusterResultOutput) PrivateFQDN() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.PrivateFQDN }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) PrivateLinkResources() PrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []PrivateLinkResourceResponse { return v.PrivateLinkResources }).(PrivateLinkResourceResponseArrayOutput)
}

func (o LookupManagedClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupManagedClusterResultOutput) SecurityProfile() ManagedClusterSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterSecurityProfileResponse { return v.SecurityProfile }).(ManagedClusterSecurityProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) ServicePrincipalProfile() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterServicePrincipalProfileResponse {
		return v.ServicePrincipalProfile
	}).(ManagedClusterServicePrincipalProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) Sku() ManagedClusterSKUResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterSKUResponse { return v.Sku }).(ManagedClusterSKUResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) StorageProfile() ManagedClusterStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterStorageProfileResponse { return v.StorageProfile }).(ManagedClusterStorageProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupManagedClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupManagedClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) WindowsProfile() ManagedClusterWindowsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterWindowsProfileResponse { return v.WindowsProfile }).(ManagedClusterWindowsProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedClusterResultOutput{})
}
