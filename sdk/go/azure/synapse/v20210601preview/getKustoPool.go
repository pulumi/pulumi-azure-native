


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoPool(ctx *pulumi.Context, args *LookupKustoPoolArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolResult, error) {
	var rv LookupKustoPoolResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getKustoPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolArgs struct {
	KustoPoolName     string `pulumi:"kustoPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupKustoPoolResult struct {
	DataIngestionUri      string                         `pulumi:"dataIngestionUri"`
	EnablePurge           *bool                          `pulumi:"enablePurge"`
	EnableStreamingIngest *bool                          `pulumi:"enableStreamingIngest"`
	Etag                  string                         `pulumi:"etag"`
	Id                    string                         `pulumi:"id"`
	LanguageExtensions    LanguageExtensionsListResponse `pulumi:"languageExtensions"`
	Location              string                         `pulumi:"location"`
	Name                  string                         `pulumi:"name"`
	OptimizedAutoscale    *OptimizedAutoscaleResponse    `pulumi:"optimizedAutoscale"`
	ProvisioningState     string                         `pulumi:"provisioningState"`
	Sku                   AzureSkuResponse               `pulumi:"sku"`
	State                 string                         `pulumi:"state"`
	StateReason           string                         `pulumi:"stateReason"`
	SystemData            SystemDataResponse             `pulumi:"systemData"`
	Tags                  map[string]string              `pulumi:"tags"`
	Type                  string                         `pulumi:"type"`
	Uri                   string                         `pulumi:"uri"`
	WorkspaceUID          *string                        `pulumi:"workspaceUID"`
}
