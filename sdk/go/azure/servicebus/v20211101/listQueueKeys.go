


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListQueueKeys(ctx *pulumi.Context, args *ListQueueKeysArgs, opts ...pulumi.InvokeOption) (*ListQueueKeysResult, error) {
	var rv ListQueueKeysResult
	err := ctx.Invoke("azure-native:servicebus/v20211101:listQueueKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListQueueKeysArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	QueueName             string `pulumi:"queueName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListQueueKeysResult struct {
	AliasPrimaryConnectionString   string `pulumi:"aliasPrimaryConnectionString"`
	AliasSecondaryConnectionString string `pulumi:"aliasSecondaryConnectionString"`
	KeyName                        string `pulumi:"keyName"`
	PrimaryConnectionString        string `pulumi:"primaryConnectionString"`
	PrimaryKey                     string `pulumi:"primaryKey"`
	SecondaryConnectionString      string `pulumi:"secondaryConnectionString"`
	SecondaryKey                   string `pulumi:"secondaryKey"`
}

func ListQueueKeysOutput(ctx *pulumi.Context, args ListQueueKeysOutputArgs, opts ...pulumi.InvokeOption) ListQueueKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListQueueKeysResult, error) {
			args := v.(ListQueueKeysArgs)
			r, err := ListQueueKeys(ctx, &args, opts...)
			var s ListQueueKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListQueueKeysResultOutput)
}

type ListQueueKeysOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	QueueName             pulumi.StringInput `pulumi:"queueName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListQueueKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListQueueKeysArgs)(nil)).Elem()
}


type ListQueueKeysResultOutput struct{ *pulumi.OutputState }

func (ListQueueKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListQueueKeysResult)(nil)).Elem()
}

func (o ListQueueKeysResultOutput) ToListQueueKeysResultOutput() ListQueueKeysResultOutput {
	return o
}

func (o ListQueueKeysResultOutput) ToListQueueKeysResultOutputWithContext(ctx context.Context) ListQueueKeysResultOutput {
	return o
}

func (o ListQueueKeysResultOutput) AliasPrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.AliasPrimaryConnectionString }).(pulumi.StringOutput)
}

func (o ListQueueKeysResultOutput) AliasSecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.AliasSecondaryConnectionString }).(pulumi.StringOutput)
}

func (o ListQueueKeysResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o ListQueueKeysResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

func (o ListQueueKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o ListQueueKeysResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

func (o ListQueueKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListQueueKeysResultOutput{})
}
