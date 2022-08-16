


package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupChannel(ctx *pulumi.Context, args *LookupChannelArgs, opts ...pulumi.InvokeOption) (*LookupChannelResult, error) {
	var rv LookupChannelResult
	err := ctx.Invoke("azure-native:eventgrid:getChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupChannelArgs struct {
	ChannelName          string `pulumi:"channelName"`
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupChannelResult struct {
	ChannelType                     *string                                `pulumi:"channelType"`
	ExpirationTimeIfNotActivatedUtc *string                                `pulumi:"expirationTimeIfNotActivatedUtc"`
	Id                              string                                 `pulumi:"id"`
	MessageForActivation            *string                                `pulumi:"messageForActivation"`
	Name                            string                                 `pulumi:"name"`
	PartnerDestinationInfo          *WebhookPartnerDestinationInfoResponse `pulumi:"partnerDestinationInfo"`
	PartnerTopicInfo                *PartnerTopicInfoResponse              `pulumi:"partnerTopicInfo"`
	ProvisioningState               *string                                `pulumi:"provisioningState"`
	ReadinessState                  *string                                `pulumi:"readinessState"`
	SystemData                      SystemDataResponse                     `pulumi:"systemData"`
	Type                            string                                 `pulumi:"type"`
}


func (val *LookupChannelResult) Defaults() *LookupChannelResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PartnerDestinationInfo = tmp.PartnerDestinationInfo.Defaults()

	return &tmp
}

func LookupChannelOutput(ctx *pulumi.Context, args LookupChannelOutputArgs, opts ...pulumi.InvokeOption) LookupChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChannelResult, error) {
			args := v.(LookupChannelArgs)
			r, err := LookupChannel(ctx, &args, opts...)
			var s LookupChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChannelResultOutput)
}

type LookupChannelOutputArgs struct {
	ChannelName          pulumi.StringInput `pulumi:"channelName"`
	PartnerNamespaceName pulumi.StringInput `pulumi:"partnerNamespaceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelArgs)(nil)).Elem()
}


type LookupChannelResultOutput struct{ *pulumi.OutputState }

func (LookupChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelResult)(nil)).Elem()
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutput() LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutputWithContext(ctx context.Context) LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) ChannelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.ChannelType }).(pulumi.StringPtrOutput)
}

func (o LookupChannelResultOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

func (o LookupChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupChannelResultOutput) MessageForActivation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.MessageForActivation }).(pulumi.StringPtrOutput)
}

func (o LookupChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupChannelResultOutput) PartnerDestinationInfo() WebhookPartnerDestinationInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *WebhookPartnerDestinationInfoResponse { return v.PartnerDestinationInfo }).(WebhookPartnerDestinationInfoResponsePtrOutput)
}

func (o LookupChannelResultOutput) PartnerTopicInfo() PartnerTopicInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *PartnerTopicInfoResponse { return v.PartnerTopicInfo }).(PartnerTopicInfoResponsePtrOutput)
}

func (o LookupChannelResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupChannelResultOutput) ReadinessState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.ReadinessState }).(pulumi.StringPtrOutput)
}

func (o LookupChannelResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupChannelResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupChannelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelResultOutput{})
}
