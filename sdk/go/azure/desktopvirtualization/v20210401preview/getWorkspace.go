


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20210401preview:getWorkspace", args, &rv, opts...)
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
	ApplicationGroupReferences []string                                             `pulumi:"applicationGroupReferences"`
	CloudPcResource            bool                                                 `pulumi:"cloudPcResource"`
	Description                *string                                              `pulumi:"description"`
	Etag                       string                                               `pulumi:"etag"`
	FriendlyName               *string                                              `pulumi:"friendlyName"`
	Id                         string                                               `pulumi:"id"`
	Identity                   *ResourceModelWithAllowedPropertySetResponseIdentity `pulumi:"identity"`
	Kind                       *string                                              `pulumi:"kind"`
	Location                   *string                                              `pulumi:"location"`
	ManagedBy                  *string                                              `pulumi:"managedBy"`
	Name                       string                                               `pulumi:"name"`
	ObjectId                   string                                               `pulumi:"objectId"`
	Plan                       *ResourceModelWithAllowedPropertySetResponsePlan     `pulumi:"plan"`
	PublicNetworkAccess        *string                                              `pulumi:"publicNetworkAccess"`
	Sku                        *ResourceModelWithAllowedPropertySetResponseSku      `pulumi:"sku"`
	Tags                       map[string]string                                    `pulumi:"tags"`
	Type                       string                                               `pulumi:"type"`
}
