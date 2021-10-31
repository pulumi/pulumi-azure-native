


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAddon(ctx *pulumi.Context, args *LookupAddonArgs, opts ...pulumi.InvokeOption) (*LookupAddonResult, error) {
	var rv LookupAddonResult
	err := ctx.Invoke("azure-native:avs/v20210601:getAddon", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAddonArgs struct {
	AddonName         string `pulumi:"addonName"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAddonResult struct {
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}
