


package v20170401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListHybridConnectionKeys(ctx *pulumi.Context, args *ListHybridConnectionKeysArgs, opts ...pulumi.InvokeOption) (*ListHybridConnectionKeysResult, error) {
	var rv ListHybridConnectionKeysResult
	err := ctx.Invoke("azure-native:relay/v20170401:listHybridConnectionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListHybridConnectionKeysArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	HybridConnectionName  string `pulumi:"hybridConnectionName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListHybridConnectionKeysResult struct {
	KeyName                   *string `pulumi:"keyName"`
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}
