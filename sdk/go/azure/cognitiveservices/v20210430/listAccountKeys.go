


package v20210430

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAccountKeys(ctx *pulumi.Context, args *ListAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListAccountKeysResult, error) {
	var rv ListAccountKeysResult
	err := ctx.Invoke("azure-native:cognitiveservices/v20210430:listAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAccountKeysResult struct {
	Key1 *string `pulumi:"key1"`
	Key2 *string `pulumi:"key2"`
}
