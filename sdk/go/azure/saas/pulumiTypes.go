


package saas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SaasCreationProperties struct {
	AutoRenew                *bool             `pulumi:"autoRenew"`
	OfferId                  *string           `pulumi:"offerId"`
	PaymentChannelMetadata   map[string]string `pulumi:"paymentChannelMetadata"`
	PaymentChannelType       *string           `pulumi:"paymentChannelType"`
	PublisherId              *string           `pulumi:"publisherId"`
	PublisherTestEnvironment *string           `pulumi:"publisherTestEnvironment"`
	Quantity                 *float64          `pulumi:"quantity"`
	SaasResourceName         *string           `pulumi:"saasResourceName"`
	SaasSessionId            *string           `pulumi:"saasSessionId"`
	SaasSubscriptionId       *string           `pulumi:"saasSubscriptionId"`
	SkuId                    *string           `pulumi:"skuId"`
	TermId                   *string           `pulumi:"termId"`
}





type SaasCreationPropertiesInput interface {
	pulumi.Input

	ToSaasCreationPropertiesOutput() SaasCreationPropertiesOutput
	ToSaasCreationPropertiesOutputWithContext(context.Context) SaasCreationPropertiesOutput
}

type SaasCreationPropertiesArgs struct {
	AutoRenew                pulumi.BoolPtrInput    `pulumi:"autoRenew"`
	OfferId                  pulumi.StringPtrInput  `pulumi:"offerId"`
	PaymentChannelMetadata   pulumi.StringMapInput  `pulumi:"paymentChannelMetadata"`
	PaymentChannelType       pulumi.StringPtrInput  `pulumi:"paymentChannelType"`
	PublisherId              pulumi.StringPtrInput  `pulumi:"publisherId"`
	PublisherTestEnvironment pulumi.StringPtrInput  `pulumi:"publisherTestEnvironment"`
	Quantity                 pulumi.Float64PtrInput `pulumi:"quantity"`
	SaasResourceName         pulumi.StringPtrInput  `pulumi:"saasResourceName"`
	SaasSessionId            pulumi.StringPtrInput  `pulumi:"saasSessionId"`
	SaasSubscriptionId       pulumi.StringPtrInput  `pulumi:"saasSubscriptionId"`
	SkuId                    pulumi.StringPtrInput  `pulumi:"skuId"`
	TermId                   pulumi.StringPtrInput  `pulumi:"termId"`
}

func (SaasCreationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SaasCreationProperties)(nil)).Elem()
}

func (i SaasCreationPropertiesArgs) ToSaasCreationPropertiesOutput() SaasCreationPropertiesOutput {
	return i.ToSaasCreationPropertiesOutputWithContext(context.Background())
}

func (i SaasCreationPropertiesArgs) ToSaasCreationPropertiesOutputWithContext(ctx context.Context) SaasCreationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SaasCreationPropertiesOutput)
}

func (i SaasCreationPropertiesArgs) ToSaasCreationPropertiesPtrOutput() SaasCreationPropertiesPtrOutput {
	return i.ToSaasCreationPropertiesPtrOutputWithContext(context.Background())
}

func (i SaasCreationPropertiesArgs) ToSaasCreationPropertiesPtrOutputWithContext(ctx context.Context) SaasCreationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SaasCreationPropertiesOutput).ToSaasCreationPropertiesPtrOutputWithContext(ctx)
}









type SaasCreationPropertiesPtrInput interface {
	pulumi.Input

	ToSaasCreationPropertiesPtrOutput() SaasCreationPropertiesPtrOutput
	ToSaasCreationPropertiesPtrOutputWithContext(context.Context) SaasCreationPropertiesPtrOutput
}

type saasCreationPropertiesPtrType SaasCreationPropertiesArgs

func SaasCreationPropertiesPtr(v *SaasCreationPropertiesArgs) SaasCreationPropertiesPtrInput {
	return (*saasCreationPropertiesPtrType)(v)
}

func (*saasCreationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SaasCreationProperties)(nil)).Elem()
}

func (i *saasCreationPropertiesPtrType) ToSaasCreationPropertiesPtrOutput() SaasCreationPropertiesPtrOutput {
	return i.ToSaasCreationPropertiesPtrOutputWithContext(context.Background())
}

func (i *saasCreationPropertiesPtrType) ToSaasCreationPropertiesPtrOutputWithContext(ctx context.Context) SaasCreationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SaasCreationPropertiesPtrOutput)
}

type SaasCreationPropertiesOutput struct{ *pulumi.OutputState }

func (SaasCreationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SaasCreationProperties)(nil)).Elem()
}

func (o SaasCreationPropertiesOutput) ToSaasCreationPropertiesOutput() SaasCreationPropertiesOutput {
	return o
}

func (o SaasCreationPropertiesOutput) ToSaasCreationPropertiesOutputWithContext(ctx context.Context) SaasCreationPropertiesOutput {
	return o
}

func (o SaasCreationPropertiesOutput) ToSaasCreationPropertiesPtrOutput() SaasCreationPropertiesPtrOutput {
	return o.ToSaasCreationPropertiesPtrOutputWithContext(context.Background())
}

func (o SaasCreationPropertiesOutput) ToSaasCreationPropertiesPtrOutputWithContext(ctx context.Context) SaasCreationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SaasCreationProperties) *SaasCreationProperties {
		return &v
	}).(SaasCreationPropertiesPtrOutput)
}

func (o SaasCreationPropertiesOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SaasCreationProperties) *bool { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

func (o SaasCreationPropertiesOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasCreationProperties) *string { return v.OfferId }).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesOutput) PaymentChannelMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v SaasCreationProperties) map[string]string { return v.PaymentChannelMetadata }).(pulumi.StringMapOutput)
}

func (o SaasCreationPropertiesOutput) PaymentChannelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasCreationProperties) *string { return v.PaymentChannelType }).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasCreationProperties) *string { return v.PublisherId }).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesOutput) PublisherTestEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasCreationProperties) *string { return v.PublisherTestEnvironment }).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesOutput) Quantity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SaasCreationProperties) *float64 { return v.Quantity }).(pulumi.Float64PtrOutput)
}

func (o SaasCreationPropertiesOutput) SaasResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasCreationProperties) *string { return v.SaasResourceName }).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesOutput) SaasSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasCreationProperties) *string { return v.SaasSessionId }).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesOutput) SaasSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasCreationProperties) *string { return v.SaasSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesOutput) SkuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasCreationProperties) *string { return v.SkuId }).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesOutput) TermId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasCreationProperties) *string { return v.TermId }).(pulumi.StringPtrOutput)
}

type SaasCreationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SaasCreationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SaasCreationProperties)(nil)).Elem()
}

func (o SaasCreationPropertiesPtrOutput) ToSaasCreationPropertiesPtrOutput() SaasCreationPropertiesPtrOutput {
	return o
}

func (o SaasCreationPropertiesPtrOutput) ToSaasCreationPropertiesPtrOutputWithContext(ctx context.Context) SaasCreationPropertiesPtrOutput {
	return o
}

func (o SaasCreationPropertiesPtrOutput) Elem() SaasCreationPropertiesOutput {
	return o.ApplyT(func(v *SaasCreationProperties) SaasCreationProperties {
		if v != nil {
			return *v
		}
		var ret SaasCreationProperties
		return ret
	}).(SaasCreationPropertiesOutput)
}

func (o SaasCreationPropertiesPtrOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SaasCreationProperties) *bool {
		if v == nil {
			return nil
		}
		return v.AutoRenew
	}).(pulumi.BoolPtrOutput)
}

func (o SaasCreationPropertiesPtrOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.OfferId
	}).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesPtrOutput) PaymentChannelMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SaasCreationProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.PaymentChannelMetadata
	}).(pulumi.StringMapOutput)
}

func (o SaasCreationPropertiesPtrOutput) PaymentChannelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PaymentChannelType
	}).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesPtrOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublisherId
	}).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesPtrOutput) PublisherTestEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublisherTestEnvironment
	}).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesPtrOutput) Quantity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SaasCreationProperties) *float64 {
		if v == nil {
			return nil
		}
		return v.Quantity
	}).(pulumi.Float64PtrOutput)
}

func (o SaasCreationPropertiesPtrOutput) SaasResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.SaasResourceName
	}).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesPtrOutput) SaasSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.SaasSessionId
	}).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesPtrOutput) SaasSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.SaasSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesPtrOutput) SkuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.SkuId
	}).(pulumi.StringPtrOutput)
}

func (o SaasCreationPropertiesPtrOutput) TermId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.TermId
	}).(pulumi.StringPtrOutput)
}

type SaasPropertiesResponseTerm struct {
	EndDate   *string `pulumi:"endDate"`
	StartDate *string `pulumi:"startDate"`
	TermUnit  *string `pulumi:"termUnit"`
}

type SaasPropertiesResponseTermOutput struct{ *pulumi.OutputState }

func (SaasPropertiesResponseTermOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SaasPropertiesResponseTerm)(nil)).Elem()
}

func (o SaasPropertiesResponseTermOutput) ToSaasPropertiesResponseTermOutput() SaasPropertiesResponseTermOutput {
	return o
}

func (o SaasPropertiesResponseTermOutput) ToSaasPropertiesResponseTermOutputWithContext(ctx context.Context) SaasPropertiesResponseTermOutput {
	return o
}

func (o SaasPropertiesResponseTermOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasPropertiesResponseTerm) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o SaasPropertiesResponseTermOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasPropertiesResponseTerm) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o SaasPropertiesResponseTermOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasPropertiesResponseTerm) *string { return v.TermUnit }).(pulumi.StringPtrOutput)
}

type SaasPropertiesResponseTermPtrOutput struct{ *pulumi.OutputState }

func (SaasPropertiesResponseTermPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SaasPropertiesResponseTerm)(nil)).Elem()
}

func (o SaasPropertiesResponseTermPtrOutput) ToSaasPropertiesResponseTermPtrOutput() SaasPropertiesResponseTermPtrOutput {
	return o
}

func (o SaasPropertiesResponseTermPtrOutput) ToSaasPropertiesResponseTermPtrOutputWithContext(ctx context.Context) SaasPropertiesResponseTermPtrOutput {
	return o
}

func (o SaasPropertiesResponseTermPtrOutput) Elem() SaasPropertiesResponseTermOutput {
	return o.ApplyT(func(v *SaasPropertiesResponseTerm) SaasPropertiesResponseTerm {
		if v != nil {
			return *v
		}
		var ret SaasPropertiesResponseTerm
		return ret
	}).(SaasPropertiesResponseTermOutput)
}

func (o SaasPropertiesResponseTermPtrOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasPropertiesResponseTerm) *string {
		if v == nil {
			return nil
		}
		return v.EndDate
	}).(pulumi.StringPtrOutput)
}

func (o SaasPropertiesResponseTermPtrOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasPropertiesResponseTerm) *string {
		if v == nil {
			return nil
		}
		return v.StartDate
	}).(pulumi.StringPtrOutput)
}

func (o SaasPropertiesResponseTermPtrOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SaasPropertiesResponseTerm) *string {
		if v == nil {
			return nil
		}
		return v.TermUnit
	}).(pulumi.StringPtrOutput)
}

type SaasResourceResponseProperties struct {
	AutoRenew                *bool                       `pulumi:"autoRenew"`
	Created                  string                      `pulumi:"created"`
	IsFreeTrial              *bool                       `pulumi:"isFreeTrial"`
	LastModified             *string                     `pulumi:"lastModified"`
	OfferId                  *string                     `pulumi:"offerId"`
	PaymentChannelMetadata   map[string]string           `pulumi:"paymentChannelMetadata"`
	PaymentChannelType       *string                     `pulumi:"paymentChannelType"`
	PublisherId              *string                     `pulumi:"publisherId"`
	PublisherTestEnvironment *string                     `pulumi:"publisherTestEnvironment"`
	Quantity                 *float64                    `pulumi:"quantity"`
	SaasResourceName         *string                     `pulumi:"saasResourceName"`
	SaasSessionId            *string                     `pulumi:"saasSessionId"`
	SaasSubscriptionId       *string                     `pulumi:"saasSubscriptionId"`
	SkuId                    *string                     `pulumi:"skuId"`
	Status                   *string                     `pulumi:"status"`
	Term                     *SaasPropertiesResponseTerm `pulumi:"term"`
	TermId                   *string                     `pulumi:"termId"`
}

type SaasResourceResponsePropertiesOutput struct{ *pulumi.OutputState }

func (SaasResourceResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SaasResourceResponseProperties)(nil)).Elem()
}

func (o SaasResourceResponsePropertiesOutput) ToSaasResourceResponsePropertiesOutput() SaasResourceResponsePropertiesOutput {
	return o
}

func (o SaasResourceResponsePropertiesOutput) ToSaasResourceResponsePropertiesOutputWithContext(ctx context.Context) SaasResourceResponsePropertiesOutput {
	return o
}

func (o SaasResourceResponsePropertiesOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *bool { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) string { return v.Created }).(pulumi.StringOutput)
}

func (o SaasResourceResponsePropertiesOutput) IsFreeTrial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *bool { return v.IsFreeTrial }).(pulumi.BoolPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) LastModified() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *string { return v.LastModified }).(pulumi.StringPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *string { return v.OfferId }).(pulumi.StringPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) PaymentChannelMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) map[string]string { return v.PaymentChannelMetadata }).(pulumi.StringMapOutput)
}

func (o SaasResourceResponsePropertiesOutput) PaymentChannelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *string { return v.PaymentChannelType }).(pulumi.StringPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *string { return v.PublisherId }).(pulumi.StringPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) PublisherTestEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *string { return v.PublisherTestEnvironment }).(pulumi.StringPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) Quantity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *float64 { return v.Quantity }).(pulumi.Float64PtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) SaasResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *string { return v.SaasResourceName }).(pulumi.StringPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) SaasSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *string { return v.SaasSessionId }).(pulumi.StringPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) SaasSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *string { return v.SaasSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) SkuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *string { return v.SkuId }).(pulumi.StringPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) Term() SaasPropertiesResponseTermPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *SaasPropertiesResponseTerm { return v.Term }).(SaasPropertiesResponseTermPtrOutput)
}

func (o SaasResourceResponsePropertiesOutput) TermId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SaasResourceResponseProperties) *string { return v.TermId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SaasCreationPropertiesOutput{})
	pulumi.RegisterOutputType(SaasCreationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SaasPropertiesResponseTermOutput{})
	pulumi.RegisterOutputType(SaasPropertiesResponseTermPtrOutput{})
	pulumi.RegisterOutputType(SaasResourceResponsePropertiesOutput{})
}
