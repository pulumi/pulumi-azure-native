


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupQueueServiceProperties(ctx *pulumi.Context, args *LookupQueueServicePropertiesArgs, opts ...pulumi.InvokeOption) (*LookupQueueServicePropertiesResult, error) {
	var rv LookupQueueServicePropertiesResult
	err := ctx.Invoke("azure-native:storage/v20200801preview:getQueueServiceProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueServicePropertiesArgs struct {
	AccountName       string `pulumi:"accountName"`
	QueueServiceName  string `pulumi:"queueServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupQueueServicePropertiesResult struct {
	Cors *CorsRulesResponse `pulumi:"cors"`
	Id   string             `pulumi:"id"`
	Name string             `pulumi:"name"`
	Type string             `pulumi:"type"`
}

func LookupQueueServicePropertiesOutput(ctx *pulumi.Context, args LookupQueueServicePropertiesOutputArgs, opts ...pulumi.InvokeOption) LookupQueueServicePropertiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueueServicePropertiesResult, error) {
			args := v.(LookupQueueServicePropertiesArgs)
			r, err := LookupQueueServiceProperties(ctx, &args, opts...)
			var s LookupQueueServicePropertiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueueServicePropertiesResultOutput)
}

type LookupQueueServicePropertiesOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	QueueServiceName  pulumi.StringInput `pulumi:"queueServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupQueueServicePropertiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueServicePropertiesArgs)(nil)).Elem()
}


type LookupQueueServicePropertiesResultOutput struct{ *pulumi.OutputState }

func (LookupQueueServicePropertiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueServicePropertiesResult)(nil)).Elem()
}

func (o LookupQueueServicePropertiesResultOutput) ToLookupQueueServicePropertiesResultOutput() LookupQueueServicePropertiesResultOutput {
	return o
}

func (o LookupQueueServicePropertiesResultOutput) ToLookupQueueServicePropertiesResultOutputWithContext(ctx context.Context) LookupQueueServicePropertiesResultOutput {
	return o
}

func (o LookupQueueServicePropertiesResultOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v LookupQueueServicePropertiesResult) *CorsRulesResponse { return v.Cors }).(CorsRulesResponsePtrOutput)
}

func (o LookupQueueServicePropertiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueServicePropertiesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQueueServicePropertiesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueServicePropertiesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupQueueServicePropertiesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueServicePropertiesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueueServicePropertiesResultOutput{})
}
