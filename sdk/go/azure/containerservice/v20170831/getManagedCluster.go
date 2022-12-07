


package v20170831

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:containerservice/v20170831:getManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedClusterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupManagedClusterResult struct {
	AgentPoolProfiles       []ContainerServiceAgentPoolProfileResponse       `pulumi:"agentPoolProfiles"`
	DnsPrefix               *string                                          `pulumi:"dnsPrefix"`
	Fqdn                    string                                           `pulumi:"fqdn"`
	Id                      string                                           `pulumi:"id"`
	KubernetesVersion       *string                                          `pulumi:"kubernetesVersion"`
	LinuxProfile            *ContainerServiceLinuxProfileResponse            `pulumi:"linuxProfile"`
	Location                string                                           `pulumi:"location"`
	Name                    string                                           `pulumi:"name"`
	ProvisioningState       string                                           `pulumi:"provisioningState"`
	ServicePrincipalProfile *ContainerServiceServicePrincipalProfileResponse `pulumi:"servicePrincipalProfile"`
	Tags                    map[string]string                                `pulumi:"tags"`
	Type                    string                                           `pulumi:"type"`
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

func (o LookupManagedClusterResultOutput) AgentPoolProfiles() ContainerServiceAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []ContainerServiceAgentPoolProfileResponse {
		return v.AgentPoolProfiles
	}).(ContainerServiceAgentPoolProfileResponseArrayOutput)
}

func (o LookupManagedClusterResultOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.DnsPrefix }).(pulumi.StringPtrOutput)
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

func (o LookupManagedClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) ServicePrincipalProfile() ContainerServiceServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ContainerServiceServicePrincipalProfileResponse {
		return v.ServicePrincipalProfile
	}).(ContainerServiceServicePrincipalProfileResponsePtrOutput)
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
