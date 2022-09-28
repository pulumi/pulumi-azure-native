


package v20170615preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupEventSubscription(ctx *pulumi.Context, args *LookupEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupEventSubscriptionResult, error) {
	var rv LookupEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid/v20170615preview:getEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEventSubscriptionArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	Scope                 string `pulumi:"scope"`
}


type LookupEventSubscriptionResult struct {
	Destination       *EventSubscriptionDestinationResponse `pulumi:"destination"`
	Filter            *EventSubscriptionFilterResponse      `pulumi:"filter"`
	Id                string                                `pulumi:"id"`
	Labels            []string                              `pulumi:"labels"`
	Name              string                                `pulumi:"name"`
	ProvisioningState string                                `pulumi:"provisioningState"`
	Topic             string                                `pulumi:"topic"`
	Type              string                                `pulumi:"type"`
}


func (val *LookupEventSubscriptionResult) Defaults() *LookupEventSubscriptionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Filter = tmp.Filter.Defaults()

	return &tmp
}

func LookupEventSubscriptionOutput(ctx *pulumi.Context, args LookupEventSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupEventSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventSubscriptionResult, error) {
			args := v.(LookupEventSubscriptionArgs)
			r, err := LookupEventSubscription(ctx, &args, opts...)
			var s LookupEventSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventSubscriptionResultOutput)
}

type LookupEventSubscriptionOutputArgs struct {
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	Scope                 pulumi.StringInput `pulumi:"scope"`
}

func (LookupEventSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventSubscriptionArgs)(nil)).Elem()
}


type LookupEventSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupEventSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventSubscriptionResult)(nil)).Elem()
}

func (o LookupEventSubscriptionResultOutput) ToLookupEventSubscriptionResultOutput() LookupEventSubscriptionResultOutput {
	return o
}

func (o LookupEventSubscriptionResultOutput) ToLookupEventSubscriptionResultOutputWithContext(ctx context.Context) LookupEventSubscriptionResultOutput {
	return o
}

func (o LookupEventSubscriptionResultOutput) Destination() EventSubscriptionDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) *EventSubscriptionDestinationResponse { return v.Destination }).(EventSubscriptionDestinationResponsePtrOutput)
}

func (o LookupEventSubscriptionResultOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) *EventSubscriptionFilterResponse { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

func (o LookupEventSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEventSubscriptionResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o LookupEventSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEventSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupEventSubscriptionResultOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) string { return v.Topic }).(pulumi.StringOutput)
}

func (o LookupEventSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventSubscriptionResultOutput{})
}
