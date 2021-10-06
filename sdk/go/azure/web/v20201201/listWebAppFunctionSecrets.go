


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppFunctionSecrets(ctx *pulumi.Context, args *ListWebAppFunctionSecretsArgs, opts ...pulumi.InvokeOption) (*ListWebAppFunctionSecretsResult, error) {
	var rv ListWebAppFunctionSecretsResult
	err := ctx.Invoke("azure-native:web/v20201201:listWebAppFunctionSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppFunctionSecretsArgs struct {
	FunctionName      string `pulumi:"functionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppFunctionSecretsResult struct {
	Key        *string `pulumi:"key"`
	TriggerUrl *string `pulumi:"triggerUrl"`
}
