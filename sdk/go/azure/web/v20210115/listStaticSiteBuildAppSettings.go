


package v20210115

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteBuildAppSettings(ctx *pulumi.Context, args *ListStaticSiteBuildAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteBuildAppSettingsResult, error) {
	var rv ListStaticSiteBuildAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20210115:listStaticSiteBuildAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteBuildAppSettingsArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteBuildAppSettingsResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}
