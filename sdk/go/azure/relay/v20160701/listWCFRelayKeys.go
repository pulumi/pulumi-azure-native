


package v20160701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWCFRelayKeys(ctx *pulumi.Context, args *ListWCFRelayKeysArgs, opts ...pulumi.InvokeOption) (*ListWCFRelayKeysResult, error) {
	var rv ListWCFRelayKeysResult
	err := ctx.Invoke("azure-native:relay/v20160701:listWCFRelayKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWCFRelayKeysArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	RelayName             string `pulumi:"relayName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListWCFRelayKeysResult struct {
	KeyName                   *string `pulumi:"keyName"`
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}
