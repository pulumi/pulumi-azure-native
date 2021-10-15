


package v20180201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppApplicationSettings(ctx *pulumi.Context, args *ListWebAppApplicationSettingsArgs, opts ...pulumi.InvokeOption) (*ListWebAppApplicationSettingsResult, error) {
	var rv ListWebAppApplicationSettingsResult
	err := ctx.Invoke("azure-native:web/v20180201:listWebAppApplicationSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppApplicationSettingsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppApplicationSettingsResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}
