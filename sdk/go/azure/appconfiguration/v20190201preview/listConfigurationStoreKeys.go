


package v20190201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConfigurationStoreKeys(ctx *pulumi.Context, args *ListConfigurationStoreKeysArgs, opts ...pulumi.InvokeOption) (*ListConfigurationStoreKeysResult, error) {
	var rv ListConfigurationStoreKeysResult
	err := ctx.Invoke("azure-native:appconfiguration/v20190201preview:listConfigurationStoreKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConfigurationStoreKeysArgs struct {
	ConfigStoreName   string  `pulumi:"configStoreName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SkipToken         *string `pulumi:"skipToken"`
}


type ListConfigurationStoreKeysResult struct {
	NextLink *string          `pulumi:"nextLink"`
	Value    []ApiKeyResponse `pulumi:"value"`
}
