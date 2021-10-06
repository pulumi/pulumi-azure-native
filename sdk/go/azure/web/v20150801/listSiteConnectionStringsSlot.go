


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteConnectionStringsSlot(ctx *pulumi.Context, args *ListSiteConnectionStringsSlotArgs, opts ...pulumi.InvokeOption) (*ListSiteConnectionStringsSlotResult, error) {
	var rv ListSiteConnectionStringsSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteConnectionStringsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteConnectionStringsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListSiteConnectionStringsSlotResult struct {
	Id         *string                                    `pulumi:"id"`
	Kind       *string                                    `pulumi:"kind"`
	Location   string                                     `pulumi:"location"`
	Name       *string                                    `pulumi:"name"`
	Properties map[string]ConnStringValueTypePairResponse `pulumi:"properties"`
	Tags       map[string]string                          `pulumi:"tags"`
	Type       *string                                    `pulumi:"type"`
}
