


package billing

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzurePlanResponse struct {
	SkuDescription string  `pulumi:"skuDescription"`
	SkuId          *string `pulumi:"skuId"`
}

type AzurePlanResponseOutput struct{ *pulumi.OutputState }

func (AzurePlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzurePlanResponse)(nil)).Elem()
}

func (o AzurePlanResponseOutput) ToAzurePlanResponseOutput() AzurePlanResponseOutput {
	return o
}

func (o AzurePlanResponseOutput) ToAzurePlanResponseOutputWithContext(ctx context.Context) AzurePlanResponseOutput {
	return o
}

func (o AzurePlanResponseOutput) SkuDescription() pulumi.StringOutput {
	return o.ApplyT(func(v AzurePlanResponse) string { return v.SkuDescription }).(pulumi.StringOutput)
}

func (o AzurePlanResponseOutput) SkuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePlanResponse) *string { return v.SkuId }).(pulumi.StringPtrOutput)
}

type AzurePlanResponseArrayOutput struct{ *pulumi.OutputState }

func (AzurePlanResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzurePlanResponse)(nil)).Elem()
}

func (o AzurePlanResponseArrayOutput) ToAzurePlanResponseArrayOutput() AzurePlanResponseArrayOutput {
	return o
}

func (o AzurePlanResponseArrayOutput) ToAzurePlanResponseArrayOutputWithContext(ctx context.Context) AzurePlanResponseArrayOutput {
	return o
}

func (o AzurePlanResponseArrayOutput) Index(i pulumi.IntInput) AzurePlanResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzurePlanResponse {
		return vs[0].([]AzurePlanResponse)[vs[1].(int)]
	}).(AzurePlanResponseOutput)
}

type InvoiceSectionWithCreateSubPermissionResponse struct {
	BillingProfileDisplayName      string              `pulumi:"billingProfileDisplayName"`
	BillingProfileId               string              `pulumi:"billingProfileId"`
	BillingProfileSpendingLimit    string              `pulumi:"billingProfileSpendingLimit"`
	BillingProfileStatus           string              `pulumi:"billingProfileStatus"`
	BillingProfileStatusReasonCode string              `pulumi:"billingProfileStatusReasonCode"`
	BillingProfileSystemId         string              `pulumi:"billingProfileSystemId"`
	EnabledAzurePlans              []AzurePlanResponse `pulumi:"enabledAzurePlans"`
	InvoiceSectionDisplayName      string              `pulumi:"invoiceSectionDisplayName"`
	InvoiceSectionId               string              `pulumi:"invoiceSectionId"`
	InvoiceSectionSystemId         string              `pulumi:"invoiceSectionSystemId"`
}

type InvoiceSectionWithCreateSubPermissionResponseOutput struct{ *pulumi.OutputState }

func (InvoiceSectionWithCreateSubPermissionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InvoiceSectionWithCreateSubPermissionResponse)(nil)).Elem()
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) ToInvoiceSectionWithCreateSubPermissionResponseOutput() InvoiceSectionWithCreateSubPermissionResponseOutput {
	return o
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) ToInvoiceSectionWithCreateSubPermissionResponseOutputWithContext(ctx context.Context) InvoiceSectionWithCreateSubPermissionResponseOutput {
	return o
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) BillingProfileDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.BillingProfileDisplayName }).(pulumi.StringOutput)
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) BillingProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.BillingProfileId }).(pulumi.StringOutput)
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) BillingProfileSpendingLimit() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.BillingProfileSpendingLimit }).(pulumi.StringOutput)
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) BillingProfileStatus() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.BillingProfileStatus }).(pulumi.StringOutput)
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) BillingProfileStatusReasonCode() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.BillingProfileStatusReasonCode }).(pulumi.StringOutput)
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) BillingProfileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.BillingProfileSystemId }).(pulumi.StringOutput)
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) EnabledAzurePlans() AzurePlanResponseArrayOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) []AzurePlanResponse { return v.EnabledAzurePlans }).(AzurePlanResponseArrayOutput)
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) InvoiceSectionDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.InvoiceSectionDisplayName }).(pulumi.StringOutput)
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) InvoiceSectionId() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.InvoiceSectionId }).(pulumi.StringOutput)
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) InvoiceSectionSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.InvoiceSectionSystemId }).(pulumi.StringOutput)
}

type InvoiceSectionWithCreateSubPermissionResponseArrayOutput struct{ *pulumi.OutputState }

func (InvoiceSectionWithCreateSubPermissionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InvoiceSectionWithCreateSubPermissionResponse)(nil)).Elem()
}

func (o InvoiceSectionWithCreateSubPermissionResponseArrayOutput) ToInvoiceSectionWithCreateSubPermissionResponseArrayOutput() InvoiceSectionWithCreateSubPermissionResponseArrayOutput {
	return o
}

func (o InvoiceSectionWithCreateSubPermissionResponseArrayOutput) ToInvoiceSectionWithCreateSubPermissionResponseArrayOutputWithContext(ctx context.Context) InvoiceSectionWithCreateSubPermissionResponseArrayOutput {
	return o
}

func (o InvoiceSectionWithCreateSubPermissionResponseArrayOutput) Index(i pulumi.IntInput) InvoiceSectionWithCreateSubPermissionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InvoiceSectionWithCreateSubPermissionResponse {
		return vs[0].([]InvoiceSectionWithCreateSubPermissionResponse)[vs[1].(int)]
	}).(InvoiceSectionWithCreateSubPermissionResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AzurePlanResponseOutput{})
	pulumi.RegisterOutputType(AzurePlanResponseArrayOutput{})
	pulumi.RegisterOutputType(InvoiceSectionWithCreateSubPermissionResponseOutput{})
	pulumi.RegisterOutputType(InvoiceSectionWithCreateSubPermissionResponseArrayOutput{})
}
