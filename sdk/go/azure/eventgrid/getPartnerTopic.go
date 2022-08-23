


package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartnerTopic(ctx *pulumi.Context, args *LookupPartnerTopicArgs, opts ...pulumi.InvokeOption) (*LookupPartnerTopicResult, error) {
	var rv LookupPartnerTopicResult
	err := ctx.Invoke("azure-native:eventgrid:getPartnerTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerTopicArgs struct {
	PartnerTopicName  string `pulumi:"partnerTopicName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPartnerTopicResult struct {
	ActivationState                 *string                `pulumi:"activationState"`
	EventTypeInfo                   *EventTypeInfoResponse `pulumi:"eventTypeInfo"`
	ExpirationTimeIfNotActivatedUtc *string                `pulumi:"expirationTimeIfNotActivatedUtc"`
	Id                              string                 `pulumi:"id"`
	Identity                        *IdentityInfoResponse  `pulumi:"identity"`
	Location                        string                 `pulumi:"location"`
	MessageForActivation            *string                `pulumi:"messageForActivation"`
	Name                            string                 `pulumi:"name"`
	PartnerRegistrationImmutableId  *string                `pulumi:"partnerRegistrationImmutableId"`
	PartnerTopicFriendlyDescription *string                `pulumi:"partnerTopicFriendlyDescription"`
	ProvisioningState               string                 `pulumi:"provisioningState"`
	Source                          *string                `pulumi:"source"`
	SystemData                      SystemDataResponse     `pulumi:"systemData"`
	Tags                            map[string]string      `pulumi:"tags"`
	Type                            string                 `pulumi:"type"`
}

func LookupPartnerTopicOutput(ctx *pulumi.Context, args LookupPartnerTopicOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerTopicResult, error) {
			args := v.(LookupPartnerTopicArgs)
			r, err := LookupPartnerTopic(ctx, &args, opts...)
			var s LookupPartnerTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerTopicResultOutput)
}

type LookupPartnerTopicOutputArgs struct {
	PartnerTopicName  pulumi.StringInput `pulumi:"partnerTopicName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerTopicArgs)(nil)).Elem()
}


type LookupPartnerTopicResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerTopicResult)(nil)).Elem()
}

func (o LookupPartnerTopicResultOutput) ToLookupPartnerTopicResultOutput() LookupPartnerTopicResultOutput {
	return o
}

func (o LookupPartnerTopicResultOutput) ToLookupPartnerTopicResultOutputWithContext(ctx context.Context) LookupPartnerTopicResultOutput {
	return o
}

func (o LookupPartnerTopicResultOutput) ActivationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.ActivationState }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerTopicResultOutput) EventTypeInfo() EventTypeInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *EventTypeInfoResponse { return v.EventTypeInfo }).(EventTypeInfoResponsePtrOutput)
}

func (o LookupPartnerTopicResultOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPartnerTopicResultOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *IdentityInfoResponse { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

func (o LookupPartnerTopicResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPartnerTopicResultOutput) MessageForActivation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.MessageForActivation }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPartnerTopicResultOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerTopicResultOutput) PartnerTopicFriendlyDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.PartnerTopicFriendlyDescription }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerTopicResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPartnerTopicResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerTopicResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPartnerTopicResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPartnerTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerTopicResultOutput{})
}
