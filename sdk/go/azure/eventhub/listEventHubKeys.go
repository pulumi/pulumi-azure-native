


package eventhub

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEventHubKeys(ctx *pulumi.Context, args *ListEventHubKeysArgs, opts ...pulumi.InvokeOption) (*ListEventHubKeysResult, error) {
	var rv ListEventHubKeysResult
	err := ctx.Invoke("azure-native:eventhub:listEventHubKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEventHubKeysArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	EventHubName          string `pulumi:"eventHubName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListEventHubKeysResult struct {
	AliasPrimaryConnectionString   string `pulumi:"aliasPrimaryConnectionString"`
	AliasSecondaryConnectionString string `pulumi:"aliasSecondaryConnectionString"`
	KeyName                        string `pulumi:"keyName"`
	PrimaryConnectionString        string `pulumi:"primaryConnectionString"`
	PrimaryKey                     string `pulumi:"primaryKey"`
	SecondaryConnectionString      string `pulumi:"secondaryConnectionString"`
	SecondaryKey                   string `pulumi:"secondaryKey"`
}
