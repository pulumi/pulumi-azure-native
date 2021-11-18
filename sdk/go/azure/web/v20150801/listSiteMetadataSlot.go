


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteMetadataSlot(ctx *pulumi.Context, args *ListSiteMetadataSlotArgs, opts ...pulumi.InvokeOption) (*ListSiteMetadataSlotResult, error) {
	var rv ListSiteMetadataSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteMetadataSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteMetadataSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListSiteMetadataSlotResult struct {
	Id         *string           `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Location   string            `pulumi:"location"`
	Name       *string           `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       *string           `pulumi:"type"`
}
