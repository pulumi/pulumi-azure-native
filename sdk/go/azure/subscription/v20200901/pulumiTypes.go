


package v20200901

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
	pulumi.RegisterOutputType(PutAliasResponsePropertiesResponseOutput{})
}
