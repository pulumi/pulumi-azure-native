


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppConnectionStringsSlot(ctx *pulumi.Context, args *ListWebAppConnectionStringsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppConnectionStringsSlotResult, error) {
	var rv ListWebAppConnectionStringsSlotResult
	err := ctx.Invoke("azure-native:web/v20190801:listWebAppConnectionStringsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppConnectionStringsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppConnectionStringsSlotResult struct {
	Id         string                                     `pulumi:"id"`
	Kind       *string                                    `pulumi:"kind"`
	Name       string                                     `pulumi:"name"`
	Properties map[string]ConnStringValueTypePairResponse `pulumi:"properties"`
	Type       string                                     `pulumi:"type"`
}
