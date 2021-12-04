


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppMetadataSlot(ctx *pulumi.Context, args *ListWebAppMetadataSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppMetadataSlotResult, error) {
	var rv ListWebAppMetadataSlotResult
	err := ctx.Invoke("azure-native:web/v20201201:listWebAppMetadataSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppMetadataSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppMetadataSlotResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}
