


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTrigger(ctx *pulumi.Context, args *LookupTriggerArgs, opts ...pulumi.InvokeOption) (*LookupTriggerResult, error) {
	var rv LookupTriggerResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTriggerArgs struct {
	FactoryName       string `pulumi:"factoryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TriggerName       string `pulumi:"triggerName"`
}


type LookupTriggerResult struct {
	Etag       string      `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

func LookupTriggerOutput(ctx *pulumi.Context, args LookupTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTriggerResult, error) {
			args := v.(LookupTriggerArgs)
			r, err := LookupTrigger(ctx, &args, opts...)
			var s LookupTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTriggerResultOutput)
}

type LookupTriggerOutputArgs struct {
	FactoryName       pulumi.StringInput `pulumi:"factoryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TriggerName       pulumi.StringInput `pulumi:"triggerName"`
}

func (LookupTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTriggerArgs)(nil)).Elem()
}


type LookupTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTriggerResult)(nil)).Elem()
}

func (o LookupTriggerResultOutput) ToLookupTriggerResultOutput() LookupTriggerResultOutput {
	return o
}

func (o LookupTriggerResultOutput) ToLookupTriggerResultOutputWithContext(ctx context.Context) LookupTriggerResultOutput {
	return o
}

func (o LookupTriggerResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTriggerResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTriggerResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupTriggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTriggerResultOutput{})
}
