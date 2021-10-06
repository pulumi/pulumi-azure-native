


package v20210115

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppHostKeysSlot(ctx *pulumi.Context, args *ListWebAppHostKeysSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppHostKeysSlotResult, error) {
	var rv ListWebAppHostKeysSlotResult
	err := ctx.Invoke("azure-native:web/v20210115:listWebAppHostKeysSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppHostKeysSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppHostKeysSlotResult struct {
	FunctionKeys map[string]string `pulumi:"functionKeys"`
	MasterKey    *string           `pulumi:"masterKey"`
	SystemKeys   map[string]string `pulumi:"systemKeys"`
}
