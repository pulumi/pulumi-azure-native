


package v20190121

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:kusto/v20190121:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupClusterResult struct {
	DataIngestionUri       string                          `pulumi:"dataIngestionUri"`
	Id                     string                          `pulumi:"id"`
	Location               string                          `pulumi:"location"`
	Name                   string                          `pulumi:"name"`
	ProvisioningState      string                          `pulumi:"provisioningState"`
	Sku                    AzureSkuResponse                `pulumi:"sku"`
	State                  string                          `pulumi:"state"`
	Tags                   map[string]string               `pulumi:"tags"`
	TrustedExternalTenants []TrustedExternalTenantResponse `pulumi:"trustedExternalTenants"`
	Type                   string                          `pulumi:"type"`
	Uri                    string                          `pulumi:"uri"`
}
