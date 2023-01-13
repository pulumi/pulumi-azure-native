


package v20190201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:containerservice/v20190201:getManagedCluster", args, &rv, opts...)
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
	AadProfile                  *ManagedClusterAADProfileResponse              `pulumi:"aadProfile"`
	AddonProfiles               map[string]ManagedClusterAddonProfileResponse  `pulumi:"addonProfiles"`
	AgentPoolProfiles           []ManagedClusterAgentPoolProfileResponse       `pulumi:"agentPoolProfiles"`
	ApiServerAuthorizedIPRanges []string                                       `pulumi:"apiServerAuthorizedIPRanges"`
	DnsPrefix                   *string                                        `pulumi:"dnsPrefix"`
	EnablePodSecurityPolicy     *bool                                          `pulumi:"enablePodSecurityPolicy"`
	EnableRBAC                  *bool                                          `pulumi:"enableRBAC"`
	Fqdn                        string                                         `pulumi:"fqdn"`
	Id                          string                                         `pulumi:"id"`
	KubernetesVersion           *string                                        `pulumi:"kubernetesVersion"`
	LinuxProfile                *ContainerServiceLinuxProfileResponse          `pulumi:"linuxProfile"`
	Location                    string                                         `pulumi:"location"`
	Name                        string                                         `pulumi:"name"`
	NetworkProfile              *ContainerServiceNetworkProfileResponse        `pulumi:"networkProfile"`
	NodeResourceGroup           string                                         `pulumi:"nodeResourceGroup"`
	ProvisioningState           string                                         `pulumi:"provisioningState"`
	ServicePrincipalProfile     *ManagedClusterServicePrincipalProfileResponse `pulumi:"servicePrincipalProfile"`
	Tags                        map[string]string                              `pulumi:"tags"`
	Type                        string                                         `pulumi:"type"`
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

func (o LookupManagedClusterResultOutput) ApiServerAuthorizedIPRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []string { return v.ApiServerAuthorizedIPRanges }).(pulumi.StringArrayOutput)
}

func (o LookupManagedClusterResultOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.DnsPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupManagedClusterResultOutput) EnablePodSecurityPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.EnablePodSecurityPolicy }).(pulumi.BoolPtrOutput)
}

func (o LookupManagedClusterResultOutput) EnableRBAC() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.EnableRBAC }).(pulumi.BoolPtrOutput)
}

func (o LookupManagedClusterResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Id }).(pulumi.StringOutput)
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

func (o LookupManagedClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) NetworkProfile() ContainerServiceNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ContainerServiceNetworkProfileResponse { return v.NetworkProfile }).(ContainerServiceNetworkProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) NodeResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.NodeResourceGroup }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) ServicePrincipalProfile() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterServicePrincipalProfileResponse {
		return v.ServicePrincipalProfile
	}).(ManagedClusterServicePrincipalProfileResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupManagedClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedClusterResultOutput{})
}
