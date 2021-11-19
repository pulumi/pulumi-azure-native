


package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSitePushSettingsSlot(ctx *pulumi.Context, args *ListWebAppSitePushSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppSitePushSettingsSlotResult, error) {
	var rv ListWebAppSitePushSettingsSlotResult
	err := ctx.Invoke("azure-native:web:listWebAppSitePushSettingsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSitePushSettingsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppSitePushSettingsSlotResult struct {
	DynamicTagsJson   *string `pulumi:"dynamicTagsJson"`
	Id                string  `pulumi:"id"`
	IsPushEnabled     bool    `pulumi:"isPushEnabled"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	TagWhitelistJson  *string `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
	Type              string  `pulumi:"type"`
}
