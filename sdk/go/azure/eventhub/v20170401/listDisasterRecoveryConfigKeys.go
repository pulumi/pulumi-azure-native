


package v20170401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDisasterRecoveryConfigKeys(ctx *pulumi.Context, args *ListDisasterRecoveryConfigKeysArgs, opts ...pulumi.InvokeOption) (*ListDisasterRecoveryConfigKeysResult, error) {
	var rv ListDisasterRecoveryConfigKeysResult
	err := ctx.Invoke("azure-native:eventhub/v20170401:listDisasterRecoveryConfigKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDisasterRecoveryConfigKeysArgs struct {
	Alias                 string `pulumi:"alias"`
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListDisasterRecoveryConfigKeysResult struct {
	AliasPrimaryConnectionString   string `pulumi:"aliasPrimaryConnectionString"`
	AliasSecondaryConnectionString string `pulumi:"aliasSecondaryConnectionString"`
	KeyName                        string `pulumi:"keyName"`
	PrimaryConnectionString        string `pulumi:"primaryConnectionString"`
	PrimaryKey                     string `pulumi:"primaryKey"`
	SecondaryConnectionString      string `pulumi:"secondaryConnectionString"`
	SecondaryKey                   string `pulumi:"secondaryKey"`
}
