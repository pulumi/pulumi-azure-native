


package operationalinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:operationalinsights:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceResult struct {
	CreatedDate                     string                              `pulumi:"createdDate"`
	CustomerId                      string                              `pulumi:"customerId"`
	ETag                            *string                             `pulumi:"eTag"`
	Features                        *WorkspaceFeaturesResponse          `pulumi:"features"`
	ForceCmkForQuery                *bool                               `pulumi:"forceCmkForQuery"`
	Id                              string                              `pulumi:"id"`
	Location                        string                              `pulumi:"location"`
	ModifiedDate                    string                              `pulumi:"modifiedDate"`
	Name                            string                              `pulumi:"name"`
	PrivateLinkScopedResources      []PrivateLinkScopedResourceResponse `pulumi:"privateLinkScopedResources"`
	ProvisioningState               *string                             `pulumi:"provisioningState"`
	PublicNetworkAccessForIngestion *string                             `pulumi:"publicNetworkAccessForIngestion"`
	PublicNetworkAccessForQuery     *string                             `pulumi:"publicNetworkAccessForQuery"`
	RetentionInDays                 *int                                `pulumi:"retentionInDays"`
	Sku                             *WorkspaceSkuResponse               `pulumi:"sku"`
	Tags                            map[string]string                   `pulumi:"tags"`
	Type                            string                              `pulumi:"type"`
	WorkspaceCapping                *WorkspaceCappingResponse           `pulumi:"workspaceCapping"`
}
