


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupTrigger(ctx *pulumi.Context, args *LookupTriggerArgs, opts ...pulumi.InvokeOption) (*LookupTriggerResult, error) {
	var rv LookupTriggerResult
	err := ctx.Invoke("azure-native:datashare/v20200901:getTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTriggerArgs struct {
	AccountName           string `pulumi:"accountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
	TriggerName           string `pulumi:"triggerName"`
}


type LookupTriggerResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
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
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
	TriggerName           pulumi.StringInput `pulumi:"triggerName"`
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

func (o LookupTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTriggerResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTriggerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTriggerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTriggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTriggerResultOutput{})
}
