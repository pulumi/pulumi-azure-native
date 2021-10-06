


package v20180201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppFunctionKeysSlot(ctx *pulumi.Context, args *ListWebAppFunctionKeysSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppFunctionKeysSlotResult, error) {
	var rv ListWebAppFunctionKeysSlotResult
	err := ctx.Invoke("azure-native:web/v20180201:listWebAppFunctionKeysSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppFunctionKeysSlotArgs struct {
	FunctionName      string `pulumi:"functionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppFunctionKeysSlotResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}
