


package v20180101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTopicKeys(ctx *pulumi.Context, args *ListTopicKeysArgs, opts ...pulumi.InvokeOption) (*ListTopicKeysResult, error) {
	var rv ListTopicKeysResult
	err := ctx.Invoke("azure-native:servicebus/v20180101preview:listTopicKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTopicKeysArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TopicName             string `pulumi:"topicName"`
}


type ListTopicKeysResult struct {
	AliasPrimaryConnectionString   string `pulumi:"aliasPrimaryConnectionString"`
	AliasSecondaryConnectionString string `pulumi:"aliasSecondaryConnectionString"`
	KeyName                        string `pulumi:"keyName"`
	PrimaryConnectionString        string `pulumi:"primaryConnectionString"`
	PrimaryKey                     string `pulumi:"primaryKey"`
	SecondaryConnectionString      string `pulumi:"secondaryConnectionString"`
	SecondaryKey                   string `pulumi:"secondaryKey"`
}
