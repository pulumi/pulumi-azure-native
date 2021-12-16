


package v20191109

import (
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
