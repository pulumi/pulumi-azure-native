


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTopicKeys(ctx *pulumi.Context, args *ListTopicKeysArgs, opts ...pulumi.InvokeOption) (*ListTopicKeysResult, error) {
	var rv ListTopicKeysResult
	err := ctx.Invoke("azure-native:servicebus/v20150801:listTopicKeys", args, &rv, opts...)
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
	KeyName                   *string `pulumi:"keyName"`
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}

func ListTopicKeysOutput(ctx *pulumi.Context, args ListTopicKeysOutputArgs, opts ...pulumi.InvokeOption) ListTopicKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTopicKeysResult, error) {
			args := v.(ListTopicKeysArgs)
			r, err := ListTopicKeys(ctx, &args, opts...)
			var s ListTopicKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTopicKeysResultOutput)
}

type ListTopicKeysOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TopicName             pulumi.StringInput `pulumi:"topicName"`
}

func (ListTopicKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTopicKeysArgs)(nil)).Elem()
}


type ListTopicKeysResultOutput struct{ *pulumi.OutputState }

func (ListTopicKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTopicKeysResult)(nil)).Elem()
}

func (o ListTopicKeysResultOutput) ToListTopicKeysResultOutput() ListTopicKeysResultOutput {
	return o
}

func (o ListTopicKeysResultOutput) ToListTopicKeysResultOutputWithContext(ctx context.Context) ListTopicKeysResultOutput {
	return o
}

func (o ListTopicKeysResultOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTopicKeysResult) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o ListTopicKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTopicKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListTopicKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTopicKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListTopicKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTopicKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListTopicKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListTopicKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTopicKeysResultOutput{})
}
