


package v20191109

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:kusto/v20191109:getCluster", args, &rv, opts...)
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
	DataIngestionUri            string                               `pulumi:"dataIngestionUri"`
	EnableDiskEncryption        *bool                                `pulumi:"enableDiskEncryption"`
	EnableStreamingIngest       *bool                                `pulumi:"enableStreamingIngest"`
	Id                          string                               `pulumi:"id"`
	Identity                    *IdentityResponse                    `pulumi:"identity"`
	KeyVaultProperties          *KeyVaultPropertiesResponse          `pulumi:"keyVaultProperties"`
	Location                    string                               `pulumi:"location"`
	Name                        string                               `pulumi:"name"`
	OptimizedAutoscale          *OptimizedAutoscaleResponse          `pulumi:"optimizedAutoscale"`
	ProvisioningState           string                               `pulumi:"provisioningState"`
	Sku                         AzureSkuResponse                     `pulumi:"sku"`
	State                       string                               `pulumi:"state"`
	StateReason                 string                               `pulumi:"stateReason"`
	Tags                        map[string]string                    `pulumi:"tags"`
	TrustedExternalTenants      []TrustedExternalTenantResponse      `pulumi:"trustedExternalTenants"`
	Type                        string                               `pulumi:"type"`
	Uri                         string                               `pulumi:"uri"`
	VirtualNetworkConfiguration *VirtualNetworkConfigurationResponse `pulumi:"virtualNetworkConfiguration"`
	Zones                       []string                             `pulumi:"zones"`
}


func (val *LookupClusterResult) Defaults() *LookupClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableStreamingIngest) {
		enableStreamingIngest_ := false
		tmp.EnableStreamingIngest = &enableStreamingIngest_
	}
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

func (o LookupClusterResultOutput) DataIngestionUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.DataIngestionUri }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) EnableDiskEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EnableDiskEncryption }).(pulumi.BoolPtrOutput)
}

func (o LookupClusterResultOutput) EnableStreamingIngest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EnableStreamingIngest }).(pulumi.BoolPtrOutput)
}

func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupClusterResultOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

func (o LookupClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) OptimizedAutoscale() OptimizedAutoscaleResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *OptimizedAutoscaleResponse { return v.OptimizedAutoscale }).(OptimizedAutoscaleResponsePtrOutput)
}

func (o LookupClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Sku() AzureSkuResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) AzureSkuResponse { return v.Sku }).(AzureSkuResponseOutput)
}

func (o LookupClusterResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.StateReason }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupClusterResultOutput) TrustedExternalTenants() TrustedExternalTenantResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []TrustedExternalTenantResponse { return v.TrustedExternalTenants }).(TrustedExternalTenantResponseArrayOutput)
}

func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Uri }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) VirtualNetworkConfiguration() VirtualNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *VirtualNetworkConfigurationResponse { return v.VirtualNetworkConfiguration }).(VirtualNetworkConfigurationResponsePtrOutput)
}

func (o LookupClusterResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
