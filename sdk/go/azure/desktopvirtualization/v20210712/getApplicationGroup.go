


package v20210712

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationGroup(ctx *pulumi.Context, args *LookupApplicationGroupArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGroupResult, error) {
	var rv LookupApplicationGroupResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20210712:getApplicationGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationGroupArgs struct {
	ApplicationGroupName string `pulumi:"applicationGroupName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupApplicationGroupResult struct {
	ApplicationGroupType string                                               `pulumi:"applicationGroupType"`
	CloudPcResource      bool                                                 `pulumi:"cloudPcResource"`
	Description          *string                                              `pulumi:"description"`
	Etag                 string                                               `pulumi:"etag"`
	FriendlyName         *string                                              `pulumi:"friendlyName"`
	HostPoolArmPath      string                                               `pulumi:"hostPoolArmPath"`
	Id                   string                                               `pulumi:"id"`
	Identity             *ResourceModelWithAllowedPropertySetResponseIdentity `pulumi:"identity"`
	Kind                 *string                                              `pulumi:"kind"`
	Location             *string                                              `pulumi:"location"`
	ManagedBy            *string                                              `pulumi:"managedBy"`
	MigrationRequest     *MigrationRequestPropertiesResponse                  `pulumi:"migrationRequest"`
	Name                 string                                               `pulumi:"name"`
	ObjectId             string                                               `pulumi:"objectId"`
	Plan                 *ResourceModelWithAllowedPropertySetResponsePlan     `pulumi:"plan"`
	Sku                  *ResourceModelWithAllowedPropertySetResponseSku      `pulumi:"sku"`
	Tags                 map[string]string                                    `pulumi:"tags"`
	Type                 string                                               `pulumi:"type"`
	WorkspaceArmPath     string                                               `pulumi:"workspaceArmPath"`
}
