


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PutAliasRequestProperties struct {
	BillingScope   string  `pulumi:"billingScope"`
	DisplayName    string  `pulumi:"displayName"`
	SubscriptionId *string `pulumi:"subscriptionId"`
	Workload       string  `pulumi:"workload"`
}





type PutAliasRequestPropertiesInput interface {
	pulumi.Input

	ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput
	ToPutAliasRequestPropertiesOutputWithContext(context.Context) PutAliasRequestPropertiesOutput
}

type PutAliasRequestPropertiesArgs struct {
	BillingScope   pulumi.StringInput    `pulumi:"billingScope"`
	DisplayName    pulumi.StringInput    `pulumi:"displayName"`
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
	Workload       pulumi.StringInput    `pulumi:"workload"`
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

func (o PutAliasRequestPropertiesOutput) BillingScope() pulumi.StringOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) string { return v.BillingScope }).(pulumi.StringOutput)
}

func (o PutAliasRequestPropertiesOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o PutAliasRequestPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) Workload() pulumi.StringOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) string { return v.Workload }).(pulumi.StringOutput)
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
		return &v.BillingScope
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
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
		return &v.Workload
	}).(pulumi.StringPtrOutput)
}

type PutAliasResponsePropertiesResponse struct {
	ProvisioningState *string `pulumi:"provisioningState"`
	SubscriptionId    string  `pulumi:"subscriptionId"`
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

func (o PutAliasResponsePropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasResponsePropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PutAliasResponsePropertiesResponseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v PutAliasResponsePropertiesResponse) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PutAliasRequestPropertiesOutput{})
	pulumi.RegisterOutputType(PutAliasRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PutAliasResponsePropertiesResponseOutput{})
}
