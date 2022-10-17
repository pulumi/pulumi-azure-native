


package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupQueue(ctx *pulumi.Context, args *LookupQueueArgs, opts ...pulumi.InvokeOption) (*LookupQueueResult, error) {
	var rv LookupQueueResult
	err := ctx.Invoke("azure-native:storage:getQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueArgs struct {
	AccountName       string `pulumi:"accountName"`
	QueueName         string `pulumi:"queueName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupQueueResult struct {
	ApproximateMessageCount int               `pulumi:"approximateMessageCount"`
	Id                      string            `pulumi:"id"`
	Metadata                map[string]string `pulumi:"metadata"`
	Name                    string            `pulumi:"name"`
	Type                    string            `pulumi:"type"`
}

func LookupQueueOutput(ctx *pulumi.Context, args LookupQueueOutputArgs, opts ...pulumi.InvokeOption) LookupQueueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueueResult, error) {
			args := v.(LookupQueueArgs)
			r, err := LookupQueue(ctx, &args, opts...)
			var s LookupQueueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueueResultOutput)
}

type LookupQueueOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	QueueName         pulumi.StringInput `pulumi:"queueName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupQueueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueArgs)(nil)).Elem()
}

type LookupQueueResultOutput struct{ *pulumi.OutputState }

func (LookupQueueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueResult)(nil)).Elem()
}

func (o LookupQueueResultOutput) ToLookupQueueResultOutput() LookupQueueResultOutput {
	return o
}

func (o LookupQueueResultOutput) ToLookupQueueResultOutputWithContext(ctx context.Context) LookupQueueResultOutput {
	return o
}

func (o LookupQueueResultOutput) ApproximateMessageCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQueueResult) int { return v.ApproximateMessageCount }).(pulumi.IntOutput)
}

func (o LookupQueueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQueueResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupQueueResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupQueueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupQueueResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueueResultOutput{})
}
