


package v20180101

import (
	"context"
	"reflect"

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

func ListTopicSharedAccessKeysOutput(ctx *pulumi.Context, args ListTopicSharedAccessKeysOutputArgs, opts ...pulumi.InvokeOption) ListTopicSharedAccessKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTopicSharedAccessKeysResult, error) {
			args := v.(ListTopicSharedAccessKeysArgs)
			r, err := ListTopicSharedAccessKeys(ctx, &args, opts...)
			var s ListTopicSharedAccessKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTopicSharedAccessKeysResultOutput)
}

type ListTopicSharedAccessKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TopicName         pulumi.StringInput `pulumi:"topicName"`
}

func (ListTopicSharedAccessKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTopicSharedAccessKeysArgs)(nil)).Elem()
}


type ListTopicSharedAccessKeysResultOutput struct{ *pulumi.OutputState }

func (ListTopicSharedAccessKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTopicSharedAccessKeysResult)(nil)).Elem()
}

func (o ListTopicSharedAccessKeysResultOutput) ToListTopicSharedAccessKeysResultOutput() ListTopicSharedAccessKeysResultOutput {
	return o
}

func (o ListTopicSharedAccessKeysResultOutput) ToListTopicSharedAccessKeysResultOutputWithContext(ctx context.Context) ListTopicSharedAccessKeysResultOutput {
	return o
}

func (o ListTopicSharedAccessKeysResultOutput) Key1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTopicSharedAccessKeysResult) *string { return v.Key1 }).(pulumi.StringPtrOutput)
}

func (o ListTopicSharedAccessKeysResultOutput) Key2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTopicSharedAccessKeysResult) *string { return v.Key2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTopicSharedAccessKeysResultOutput{})
}
