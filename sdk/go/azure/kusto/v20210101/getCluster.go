


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:kusto/v20210101:getCluster", args, &rv, opts...)
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
	EnableDoubleEncryption      *bool                                `pulumi:"enableDoubleEncryption"`
	EnablePurge                 *bool                                `pulumi:"enablePurge"`
	EnableStreamingIngest       *bool                                `pulumi:"enableStreamingIngest"`
	EngineType                  *string                              `pulumi:"engineType"`
	Etag                        string                               `pulumi:"etag"`
	Id                          string                               `pulumi:"id"`
	Identity                    *IdentityResponse                    `pulumi:"identity"`
	KeyVaultProperties          *KeyVaultPropertiesResponse          `pulumi:"keyVaultProperties"`
	LanguageExtensions          LanguageExtensionsListResponse       `pulumi:"languageExtensions"`
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
	if isZero(tmp.EnableDiskEncryption) {
		enableDiskEncryption_ := false
		tmp.EnableDiskEncryption = &enableDiskEncryption_
	}
	if isZero(tmp.EnableDoubleEncryption) {
		enableDoubleEncryption_ := false
		tmp.EnableDoubleEncryption = &enableDoubleEncryption_
	}
	if isZero(tmp.EnablePurge) {
		enablePurge_ := false
		tmp.EnablePurge = &enablePurge_
	}
	if isZero(tmp.EnableStreamingIngest) {
		enableStreamingIngest_ := false
		tmp.EnableStreamingIngest = &enableStreamingIngest_
	}
	if isZero(tmp.EngineType) {
		engineType_ := "V3"
		tmp.EngineType = &engineType_
	}
	return &tmp
}
