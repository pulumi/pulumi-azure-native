


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventChannel(ctx *pulumi.Context, args *LookupEventChannelArgs, opts ...pulumi.InvokeOption) (*LookupEventChannelResult, error) {
	var rv LookupEventChannelResult
	err := ctx.Invoke("azure-native:eventgrid/v20210601preview:getEventChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEventChannelArgs struct {
	EventChannelName     string `pulumi:"eventChannelName"`
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupEventChannelResult struct {
	Destination                     *EventChannelDestinationResponse `pulumi:"destination"`
	ExpirationTimeIfNotActivatedUtc *string                          `pulumi:"expirationTimeIfNotActivatedUtc"`
	Filter                          *EventChannelFilterResponse      `pulumi:"filter"`
	Id                              string                           `pulumi:"id"`
	Name                            string                           `pulumi:"name"`
	PartnerTopicFriendlyDescription *string                          `pulumi:"partnerTopicFriendlyDescription"`
	PartnerTopicReadinessState      string                           `pulumi:"partnerTopicReadinessState"`
	ProvisioningState               string                           `pulumi:"provisioningState"`
	Source                          *EventChannelSourceResponse      `pulumi:"source"`
	SystemData                      SystemDataResponse               `pulumi:"systemData"`
	Type                            string                           `pulumi:"type"`
}


func (val *LookupEventChannelResult) Defaults() *LookupEventChannelResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Filter = tmp.Filter.Defaults()

	return &tmp
}

func LookupEventChannelOutput(ctx *pulumi.Context, args LookupEventChannelOutputArgs, opts ...pulumi.InvokeOption) LookupEventChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventChannelResult, error) {
			args := v.(LookupEventChannelArgs)
			r, err := LookupEventChannel(ctx, &args, opts...)
			var s LookupEventChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventChannelResultOutput)
}

type LookupEventChannelOutputArgs struct {
	EventChannelName     pulumi.StringInput `pulumi:"eventChannelName"`
	PartnerNamespaceName pulumi.StringInput `pulumi:"partnerNamespaceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventChannelArgs)(nil)).Elem()
}


type LookupEventChannelResultOutput struct{ *pulumi.OutputState }

func (LookupEventChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventChannelResult)(nil)).Elem()
}

func (o LookupEventChannelResultOutput) ToLookupEventChannelResultOutput() LookupEventChannelResultOutput {
	return o
}

func (o LookupEventChannelResultOutput) ToLookupEventChannelResultOutputWithContext(ctx context.Context) LookupEventChannelResultOutput {
	return o
}

func (o LookupEventChannelResultOutput) Destination() EventChannelDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupEventChannelResult) *EventChannelDestinationResponse { return v.Destination }).(EventChannelDestinationResponsePtrOutput)
}

func (o LookupEventChannelResultOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventChannelResult) *string { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

func (o LookupEventChannelResultOutput) Filter() EventChannelFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupEventChannelResult) *EventChannelFilterResponse { return v.Filter }).(EventChannelFilterResponsePtrOutput)
}

func (o LookupEventChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEventChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEventChannelResultOutput) PartnerTopicFriendlyDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventChannelResult) *string { return v.PartnerTopicFriendlyDescription }).(pulumi.StringPtrOutput)
}

func (o LookupEventChannelResultOutput) PartnerTopicReadinessState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventChannelResult) string { return v.PartnerTopicReadinessState }).(pulumi.StringOutput)
}

func (o LookupEventChannelResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventChannelResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupEventChannelResultOutput) Source() EventChannelSourceResponsePtrOutput {
	return o.ApplyT(func(v LookupEventChannelResult) *EventChannelSourceResponse { return v.Source }).(EventChannelSourceResponsePtrOutput)
}

func (o LookupEventChannelResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEventChannelResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEventChannelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventChannelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventChannelResultOutput{})
}
