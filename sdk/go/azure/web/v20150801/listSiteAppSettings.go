


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteAppSettings(ctx *pulumi.Context, args *ListSiteAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListSiteAppSettingsResult, error) {
	var rv ListSiteAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteAppSettingsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSiteAppSettingsResult struct {
	Id         *string           `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Location   string            `pulumi:"location"`
	Name       *string           `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       *string           `pulumi:"type"`
}
