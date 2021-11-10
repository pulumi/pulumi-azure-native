


package v20201102preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationGroup(ctx *pulumi.Context, args *LookupApplicationGroupArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGroupResult, error) {
	var rv LookupApplicationGroupResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20201102preview:getApplicationGroup", args, &rv, opts...)
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
	ApplicationGroupType string            `pulumi:"applicationGroupType"`
	Description          *string           `pulumi:"description"`
	FriendlyName         *string           `pulumi:"friendlyName"`
	HostPoolArmPath      string            `pulumi:"hostPoolArmPath"`
	Id                   string            `pulumi:"id"`
	Location             string            `pulumi:"location"`
	Name                 string            `pulumi:"name"`
	Tags                 map[string]string `pulumi:"tags"`
	Type                 string            `pulumi:"type"`
	WorkspaceArmPath     string            `pulumi:"workspaceArmPath"`
}
