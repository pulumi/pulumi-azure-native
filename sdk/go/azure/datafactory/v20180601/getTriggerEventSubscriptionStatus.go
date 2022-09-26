


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTriggerEventSubscriptionStatus(ctx *pulumi.Context, args *GetTriggerEventSubscriptionStatusArgs, opts ...pulumi.InvokeOption) (*GetTriggerEventSubscriptionStatusResult, error) {
	var rv GetTriggerEventSubscriptionStatusResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getTriggerEventSubscriptionStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTriggerEventSubscriptionStatusArgs struct {
	FactoryName       string `pulumi:"factoryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TriggerName       string `pulumi:"triggerName"`
}


type GetTriggerEventSubscriptionStatusResult struct {
	Status      string `pulumi:"status"`
	TriggerName string `pulumi:"triggerName"`
}

func GetTriggerEventSubscriptionStatusOutput(ctx *pulumi.Context, args GetTriggerEventSubscriptionStatusOutputArgs, opts ...pulumi.InvokeOption) GetTriggerEventSubscriptionStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTriggerEventSubscriptionStatusResult, error) {
			args := v.(GetTriggerEventSubscriptionStatusArgs)
			r, err := GetTriggerEventSubscriptionStatus(ctx, &args, opts...)
			var s GetTriggerEventSubscriptionStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTriggerEventSubscriptionStatusResultOutput)
}

type GetTriggerEventSubscriptionStatusOutputArgs struct {
	FactoryName       pulumi.StringInput `pulumi:"factoryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TriggerName       pulumi.StringInput `pulumi:"triggerName"`
}

func (GetTriggerEventSubscriptionStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTriggerEventSubscriptionStatusArgs)(nil)).Elem()
}


type GetTriggerEventSubscriptionStatusResultOutput struct{ *pulumi.OutputState }

func (GetTriggerEventSubscriptionStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTriggerEventSubscriptionStatusResult)(nil)).Elem()
}

func (o GetTriggerEventSubscriptionStatusResultOutput) ToGetTriggerEventSubscriptionStatusResultOutput() GetTriggerEventSubscriptionStatusResultOutput {
	return o
}

func (o GetTriggerEventSubscriptionStatusResultOutput) ToGetTriggerEventSubscriptionStatusResultOutputWithContext(ctx context.Context) GetTriggerEventSubscriptionStatusResultOutput {
	return o
}

func (o GetTriggerEventSubscriptionStatusResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetTriggerEventSubscriptionStatusResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetTriggerEventSubscriptionStatusResultOutput) TriggerName() pulumi.StringOutput {
	return o.ApplyT(func(v GetTriggerEventSubscriptionStatusResult) string { return v.TriggerName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTriggerEventSubscriptionStatusResultOutput{})
}
