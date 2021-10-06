


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteBuildFunctionAppSettings(ctx *pulumi.Context, args *ListStaticSiteBuildFunctionAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteBuildFunctionAppSettingsResult, error) {
	var rv ListStaticSiteBuildFunctionAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20200601:listStaticSiteBuildFunctionAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteBuildFunctionAppSettingsArgs struct {
	Name              string `pulumi:"name"`
	PrId              string `pulumi:"prId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteBuildFunctionAppSettingsResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}
