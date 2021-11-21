


package servicefabricmesh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSecretValue(ctx *pulumi.Context, args *ListSecretValueArgs, opts ...pulumi.InvokeOption) (*ListSecretValueResult, error) {
	var rv ListSecretValueResult
	err := ctx.Invoke("azure-native:servicefabricmesh:listSecretValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSecretValueArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	SecretResourceName      string `pulumi:"secretResourceName"`
	SecretValueResourceName string `pulumi:"secretValueResourceName"`
}


type ListSecretValueResult struct {
	Value *string `pulumi:"value"`
}
