


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:servicefabric/v20200101preview:getManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupManagedClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagedClusterResult struct {
	AddonFeatures                []string                             `pulumi:"addonFeatures"`
	AdminPassword                *string                              `pulumi:"adminPassword"`
	AdminUserName                string                               `pulumi:"adminUserName"`
	AzureActiveDirectory         *AzureActiveDirectoryResponse        `pulumi:"azureActiveDirectory"`
	ClientConnectionPort         *int                                 `pulumi:"clientConnectionPort"`
	Clients                      []ClientCertificateResponse          `pulumi:"clients"`
	ClusterCertificateThumbprint string                               `pulumi:"clusterCertificateThumbprint"`
	ClusterCodeVersion           *string                              `pulumi:"clusterCodeVersion"`
	ClusterId                    string                               `pulumi:"clusterId"`
	ClusterState                 string                               `pulumi:"clusterState"`
	DnsName                      string                               `pulumi:"dnsName"`
	Etag                         string                               `pulumi:"etag"`
	FabricSettings               []SettingsSectionDescriptionResponse `pulumi:"fabricSettings"`
	Fqdn                         string                               `pulumi:"fqdn"`
	HttpGatewayConnectionPort    *int                                 `pulumi:"httpGatewayConnectionPort"`
	Id                           string                               `pulumi:"id"`
	LoadBalancingRules           []LoadBalancingRuleResponse          `pulumi:"loadBalancingRules"`
	Location                     string                               `pulumi:"location"`
	Name                         string                               `pulumi:"name"`
	ProvisioningState            string                               `pulumi:"provisioningState"`
	Sku                          *SkuResponse                         `pulumi:"sku"`
	Tags                         map[string]string                    `pulumi:"tags"`
	Type                         string                               `pulumi:"type"`
}


func (val *LookupManagedClusterResult) Defaults() *LookupManagedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientConnectionPort) {
		clientConnectionPort_ := 19000
		tmp.ClientConnectionPort = &clientConnectionPort_
	}
	if isZero(tmp.HttpGatewayConnectionPort) {
		httpGatewayConnectionPort_ := 19080
		tmp.HttpGatewayConnectionPort = &httpGatewayConnectionPort_
	}
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
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
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

func (o LookupManagedClusterResultOutput) AddonFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []string { return v.AddonFeatures }).(pulumi.StringArrayOutput)
}

func (o LookupManagedClusterResultOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o LookupManagedClusterResultOutput) AdminUserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.AdminUserName }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *AzureActiveDirectoryResponse { return v.AzureActiveDirectory }).(AzureActiveDirectoryResponsePtrOutput)
}

func (o LookupManagedClusterResultOutput) ClientConnectionPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *int { return v.ClientConnectionPort }).(pulumi.IntPtrOutput)
}

func (o LookupManagedClusterResultOutput) Clients() ClientCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []ClientCertificateResponse { return v.Clients }).(ClientCertificateResponseArrayOutput)
}

func (o LookupManagedClusterResultOutput) ClusterCertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ClusterCertificateThumbprint }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) ClusterCodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.ClusterCodeVersion }).(pulumi.StringPtrOutput)
}

func (o LookupManagedClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) ClusterState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ClusterState }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.DnsName }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) FabricSettings() SettingsSectionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []SettingsSectionDescriptionResponse { return v.FabricSettings }).(SettingsSectionDescriptionResponseArrayOutput)
}

func (o LookupManagedClusterResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) HttpGatewayConnectionPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *int { return v.HttpGatewayConnectionPort }).(pulumi.IntPtrOutput)
}

func (o LookupManagedClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedClusterResultOutput) LoadBalancingRules() LoadBalancingRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []LoadBalancingRuleResponse { return v.LoadBalancingRules }).(LoadBalancingRuleResponseArrayOutput)
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

func (o LookupManagedClusterResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
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
