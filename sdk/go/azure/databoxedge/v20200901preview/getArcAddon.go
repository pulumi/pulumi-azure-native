


package v20200901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupArcAddon(ctx *pulumi.Context, args *LookupArcAddonArgs, opts ...pulumi.InvokeOption) (*LookupArcAddonResult, error) {
	var rv LookupArcAddonResult
	err := ctx.Invoke("azure-native:databoxedge/v20200901preview:getArcAddon", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArcAddonArgs struct {
	AddonName         string `pulumi:"addonName"`
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RoleName          string `pulumi:"roleName"`
}


type LookupArcAddonResult struct {
	HostPlatform      string             `pulumi:"hostPlatform"`
	HostPlatformType  string             `pulumi:"hostPlatformType"`
	Id                string             `pulumi:"id"`
	Kind              string             `pulumi:"kind"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	ResourceLocation  string             `pulumi:"resourceLocation"`
	ResourceName      string             `pulumi:"resourceName"`
	SubscriptionId    string             `pulumi:"subscriptionId"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
	Version           string             `pulumi:"version"`
}
