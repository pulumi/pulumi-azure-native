


package v20180201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppApplicationSettingsSlot(ctx *pulumi.Context, args *ListWebAppApplicationSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppApplicationSettingsSlotResult, error) {
	var rv ListWebAppApplicationSettingsSlotResult
	err := ctx.Invoke("azure-native:web/v20180201:listWebAppApplicationSettingsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppApplicationSettingsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppApplicationSettingsSlotResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}
