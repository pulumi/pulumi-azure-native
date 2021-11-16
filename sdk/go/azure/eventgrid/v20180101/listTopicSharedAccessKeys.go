


package v20180101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTopicSharedAccessKeys(ctx *pulumi.Context, args *ListTopicSharedAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListTopicSharedAccessKeysResult, error) {
	var rv ListTopicSharedAccessKeysResult
	err := ctx.Invoke("azure-native:eventgrid/v20180101:listTopicSharedAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTopicSharedAccessKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TopicName         string `pulumi:"topicName"`
}


type ListTopicSharedAccessKeysResult struct {
	Key1 *string `pulumi:"key1"`
	Key2 *string `pulumi:"key2"`
}
