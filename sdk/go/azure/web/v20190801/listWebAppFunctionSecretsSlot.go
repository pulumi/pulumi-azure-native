


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppFunctionSecretsSlot(ctx *pulumi.Context, args *ListWebAppFunctionSecretsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppFunctionSecretsSlotResult, error) {
	var rv ListWebAppFunctionSecretsSlotResult
	err := ctx.Invoke("azure-native:web/v20190801:listWebAppFunctionSecretsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppFunctionSecretsSlotArgs struct {
	FunctionName      string `pulumi:"functionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppFunctionSecretsSlotResult struct {
	Id         string  `pulumi:"id"`
	Key        *string `pulumi:"key"`
	Kind       *string `pulumi:"kind"`
	Name       string  `pulumi:"name"`
	TriggerUrl *string `pulumi:"triggerUrl"`
	Type       string  `pulumi:"type"`
}
