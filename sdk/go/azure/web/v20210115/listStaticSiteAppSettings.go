


package v20210115

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteAppSettings(ctx *pulumi.Context, args *ListStaticSiteAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteAppSettingsResult, error) {
	var rv ListStaticSiteAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20210115:listStaticSiteAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteAppSettingsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteAppSettingsResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}
