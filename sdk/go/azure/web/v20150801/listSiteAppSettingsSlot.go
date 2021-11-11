


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteAppSettingsSlot(ctx *pulumi.Context, args *ListSiteAppSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListSiteAppSettingsSlotResult, error) {
	var rv ListSiteAppSettingsSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteAppSettingsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteAppSettingsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListSiteAppSettingsSlotResult struct {
	Id         *string           `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Location   string            `pulumi:"location"`
	Name       *string           `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       *string           `pulumi:"type"`
}
