


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppFunctionKeys(ctx *pulumi.Context, args *ListWebAppFunctionKeysArgs, opts ...pulumi.InvokeOption) (*ListWebAppFunctionKeysResult, error) {
	var rv ListWebAppFunctionKeysResult
	err := ctx.Invoke("azure-native:web/v20190801:listWebAppFunctionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppFunctionKeysArgs struct {
	FunctionName      string `pulumi:"functionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppFunctionKeysResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}
