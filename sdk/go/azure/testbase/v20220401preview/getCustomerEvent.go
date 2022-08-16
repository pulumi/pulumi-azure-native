


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomerEvent(ctx *pulumi.Context, args *LookupCustomerEventArgs, opts ...pulumi.InvokeOption) (*LookupCustomerEventResult, error) {
	var rv LookupCustomerEventResult
	err := ctx.Invoke("azure-native:testbase/v20220401preview:getCustomerEvent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomerEventArgs struct {
	CustomerEventName   string `pulumi:"customerEventName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}


type LookupCustomerEventResult struct {
	EventName  string                              `pulumi:"eventName"`
	Id         string                              `pulumi:"id"`
	Name       string                              `pulumi:"name"`
	Receivers  []NotificationEventReceiverResponse `pulumi:"receivers"`
	SystemData SystemDataResponse                  `pulumi:"systemData"`
	Type       string                              `pulumi:"type"`
}

func LookupCustomerEventOutput(ctx *pulumi.Context, args LookupCustomerEventOutputArgs, opts ...pulumi.InvokeOption) LookupCustomerEventResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomerEventResult, error) {
			args := v.(LookupCustomerEventArgs)
			r, err := LookupCustomerEvent(ctx, &args, opts...)
			var s LookupCustomerEventResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomerEventResultOutput)
}

type LookupCustomerEventOutputArgs struct {
	CustomerEventName   pulumi.StringInput `pulumi:"customerEventName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (LookupCustomerEventOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomerEventArgs)(nil)).Elem()
}


type LookupCustomerEventResultOutput struct{ *pulumi.OutputState }

func (LookupCustomerEventResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomerEventResult)(nil)).Elem()
}

func (o LookupCustomerEventResultOutput) ToLookupCustomerEventResultOutput() LookupCustomerEventResultOutput {
	return o
}

func (o LookupCustomerEventResultOutput) ToLookupCustomerEventResultOutputWithContext(ctx context.Context) LookupCustomerEventResultOutput {
	return o
}

func (o LookupCustomerEventResultOutput) EventName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) string { return v.EventName }).(pulumi.StringOutput)
}

func (o LookupCustomerEventResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomerEventResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomerEventResultOutput) Receivers() NotificationEventReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) []NotificationEventReceiverResponse { return v.Receivers }).(NotificationEventReceiverResponseArrayOutput)
}

func (o LookupCustomerEventResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCustomerEventResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomerEventResultOutput{})
}
