


package v20200901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupAddon(ctx *pulumi.Context, args *LookupAddonArgs, opts ...pulumi.InvokeOption) (*LookupAddonResult, error) {
	var rv LookupAddonResult
	err := ctx.Invoke("azure-native:databoxedge/v20200901preview:getAddon", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAddonArgs struct {
	AddonName         string `pulumi:"addonName"`
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RoleName          string `pulumi:"roleName"`
}


type LookupAddonResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}
