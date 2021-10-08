


package subscription

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PutAliasRequestProperties struct {
	BillingScope   *string `pulumi:"billingScope"`
	DisplayName    *string `pulumi:"displayName"`
	ResellerId     *string `pulumi:"resellerId"`
	SubscriptionId *string `pulumi:"subscriptionId"`
	Workload       *string `pulumi:"workload"`
}





type PutAliasRequestPropertiesInput interface {
	pulumi.Input

	ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput
	ToPutAliasRequestPropertiesOutputWithContext(context.Context) PutAliasRequestPropertiesOutput
}

type PutAliasRequestPropertiesArgs struct {
	BillingScope   pulumi.StringPtrInput `pulumi:"billingScope"`
	DisplayName    pulumi.StringPtrInput `pulumi:"displayName"`
	ResellerId     pulumi.StringPtrInput `pulumi:"resellerId"`
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
	Workload       pulumi.StringPtrInput `pulumi:"workload"`
}

func (PutAliasRequestPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasRequestProperties)(nil)).Elem()
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput {
	return i.ToPutAliasRequestPropertiesOutputWithContext(context.Background())
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesOutputWithContext(ctx context.Context) PutAliasRequestPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesOutput)
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return i.ToPutAliasRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesOutput).ToPutAliasRequestPropertiesPtrOutputWithContext(ctx)
}









type PutAliasRequestPropertiesPtrInput interface {
	pulumi.Input

	ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput
	ToPutAliasRequestPropertiesPtrOutputWithContext(context.Context) PutAliasRequestPropertiesPtrOutput
}

type putAliasRequestPropertiesPtrType PutAliasRequestPropertiesArgs

func PutAliasRequestPropertiesPtr(v *PutAliasRequestPropertiesArgs) PutAliasRequestPropertiesPtrInput {
	return (*putAliasRequestPropertiesPtrType)(v)
}

func (*putAliasRequestPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasRequestProperties)(nil)).Elem()
}

func (i *putAliasRequestPropertiesPtrType) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return i.ToPutAliasRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i *putAliasRequestPropertiesPtrType) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesPtrOutput)
}

type PutAliasRequestPropertiesOutput struct{ *pulumi.OutputState }

func (PutAliasRequestPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasRequestProperties)(nil)).Elem()
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput {
	return o
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesOutputWithContext(ctx context.Context) PutAliasRequestPropertiesOutput {
	return o
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return o.ToPutAliasRequestPropertiesPtrOutputWithContext(context.Background())
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PutAliasRequestProperties) *PutAliasRequestProperties {
		return &v
	}).(PutAliasRequestPropertiesPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) BillingScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.BillingScope }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) ResellerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.ResellerId }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) Workload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.Workload }).(pulumi.StringPtrOutput)
}

type PutAliasRequestPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PutAliasRequestPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasRequestProperties)(nil)).Elem()
}

func (o PutAliasRequestPropertiesPtrOutput) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return o
}

func (o PutAliasRequestPropertiesPtrOutput) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return o
}

func (o PutAliasRequestPropertiesPtrOutput) Elem() PutAliasRequestPropertiesOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) PutAliasRequestProperties {
		if v != nil {
			return *v
		}
		var ret PutAliasRequestProperties
		return ret
	}).(PutAliasRequestPropertiesOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) BillingScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.BillingScope
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) ResellerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResellerId
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) Workload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.Workload
	}).(pulumi.StringPtrOutput)
}

type PutAliasResponsePropertiesResponse struct {
	ProvisioningState *string `pulumi:"provisioningState"`
	SubscriptionId    string  `pulumi:"subscriptionId"`
}





type PutAliasResponsePropertiesResponseInput interface {
	pulumi.Input

	ToPutAliasResponsePropertiesResponseOutput() PutAliasResponsePropertiesResponseOutput
	ToPutAliasResponsePropertiesResponseOutputWithContext(context.Context) PutAliasResponsePropertiesResponseOutput
}

type PutAliasResponsePropertiesResponseArgs struct {
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	SubscriptionId    pulumi.StringInput    `pulumi:"subscriptionId"`
}

func (PutAliasResponsePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasResponsePropertiesResponse)(nil)).Elem()
}

func (i PutAliasResponsePropertiesResponseArgs) ToPutAliasResponsePropertiesResponseOutput() PutAliasResponsePropertiesResponseOutput {
	return i.ToPutAliasResponsePropertiesResponseOutputWithContext(context.Background())
}

func (i PutAliasResponsePropertiesResponseArgs) ToPutAliasResponsePropertiesResponseOutputWithContext(ctx context.Context) PutAliasResponsePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasResponsePropertiesResponseOutput)
}

func (i PutAliasResponsePropertiesResponseArgs) ToPutAliasResponsePropertiesResponsePtrOutput() PutAliasResponsePropertiesResponsePtrOutput {
	return i.ToPutAliasResponsePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PutAliasResponsePropertiesResponseArgs) ToPutAliasResponsePropertiesResponsePtrOutputWithContext(ctx context.Context) PutAliasResponsePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasResponsePropertiesResponseOutput).ToPutAliasResponsePropertiesResponsePtrOutputWithContext(ctx)
}









type PutAliasResponsePropertiesResponsePtrInput interface {
	pulumi.Input

	ToPutAliasResponsePropertiesResponsePtrOutput() PutAliasResponsePropertiesResponsePtrOutput
	ToPutAliasResponsePropertiesResponsePtrOutputWithContext(context.Context) PutAliasResponsePropertiesResponsePtrOutput
}

type putAliasResponsePropertiesResponsePtrType PutAliasResponsePropertiesResponseArgs

func PutAliasResponsePropertiesResponsePtr(v *PutAliasResponsePropertiesResponseArgs) PutAliasResponsePropertiesResponsePtrInput {
	return (*putAliasResponsePropertiesResponsePtrType)(v)
}

func (*putAliasResponsePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasResponsePropertiesResponse)(nil)).Elem()
}

func (i *putAliasResponsePropertiesResponsePtrType) ToPutAliasResponsePropertiesResponsePtrOutput() PutAliasResponsePropertiesResponsePtrOutput {
	return i.ToPutAliasResponsePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *putAliasResponsePropertiesResponsePtrType) ToPutAliasResponsePropertiesResponsePtrOutputWithContext(ctx context.Context) PutAliasResponsePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasResponsePropertiesResponsePtrOutput)
}

type PutAliasResponsePropertiesResponseOutput struct{ *pulumi.OutputState }

func (PutAliasResponsePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasResponsePropertiesResponse)(nil)).Elem()
}

func (o PutAliasResponsePropertiesResponseOutput) ToPutAliasResponsePropertiesResponseOutput() PutAliasResponsePropertiesResponseOutput {
	return o
}

func (o PutAliasResponsePropertiesResponseOutput) ToPutAliasResponsePropertiesResponseOutputWithContext(ctx context.Context) PutAliasResponsePropertiesResponseOutput {
	return o
}

func (o PutAliasResponsePropertiesResponseOutput) ToPutAliasResponsePropertiesResponsePtrOutput() PutAliasResponsePropertiesResponsePtrOutput {
	return o.ToPutAliasResponsePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PutAliasResponsePropertiesResponseOutput) ToPutAliasResponsePropertiesResponsePtrOutputWithContext(ctx context.Context) PutAliasResponsePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PutAliasResponsePropertiesResponse) *PutAliasResponsePropertiesResponse {
		return &v
	}).(PutAliasResponsePropertiesResponsePtrOutput)
}

func (o PutAliasResponsePropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasResponsePropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PutAliasResponsePropertiesResponseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v PutAliasResponsePropertiesResponse) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

type PutAliasResponsePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PutAliasResponsePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasResponsePropertiesResponse)(nil)).Elem()
}

func (o PutAliasResponsePropertiesResponsePtrOutput) ToPutAliasResponsePropertiesResponsePtrOutput() PutAliasResponsePropertiesResponsePtrOutput {
	return o
}

func (o PutAliasResponsePropertiesResponsePtrOutput) ToPutAliasResponsePropertiesResponsePtrOutputWithContext(ctx context.Context) PutAliasResponsePropertiesResponsePtrOutput {
	return o
}

func (o PutAliasResponsePropertiesResponsePtrOutput) Elem() PutAliasResponsePropertiesResponseOutput {
	return o.ApplyT(func(v *PutAliasResponsePropertiesResponse) PutAliasResponsePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PutAliasResponsePropertiesResponse
		return ret
	}).(PutAliasResponsePropertiesResponseOutput)
}

func (o PutAliasResponsePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasResponsePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasResponsePropertiesResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasResponsePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PutAliasRequestPropertiesInput)(nil)).Elem(), PutAliasRequestPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PutAliasRequestPropertiesPtrInput)(nil)).Elem(), PutAliasRequestPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PutAliasResponsePropertiesResponseInput)(nil)).Elem(), PutAliasResponsePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PutAliasResponsePropertiesResponsePtrInput)(nil)).Elem(), PutAliasResponsePropertiesResponseArgs{})
	pulumi.RegisterOutputType(PutAliasRequestPropertiesOutput{})
	pulumi.RegisterOutputType(PutAliasRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PutAliasResponsePropertiesResponseOutput{})
	pulumi.RegisterOutputType(PutAliasResponsePropertiesResponsePtrOutput{})
}
