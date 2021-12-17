


package v20191001preview

type AzurePlanResponse struct {
	SkuDescription string  `pulumi:"skuDescription"`
	SkuId          *string `pulumi:"skuId"`
}

type InvoiceSectionWithCreateSubPermissionResponse struct {
	BillingProfileDisplayName      string              `pulumi:"billingProfileDisplayName"`
	BillingProfileId               string              `pulumi:"billingProfileId"`
	BillingProfileSpendingLimit    string              `pulumi:"billingProfileSpendingLimit"`
	BillingProfileStatus           string              `pulumi:"billingProfileStatus"`
	BillingProfileStatusReasonCode string              `pulumi:"billingProfileStatusReasonCode"`
	EnabledAzurePlans              []AzurePlanResponse `pulumi:"enabledAzurePlans"`
	InvoiceSectionDisplayName      string              `pulumi:"invoiceSectionDisplayName"`
	InvoiceSectionId               string              `pulumi:"invoiceSectionId"`
}

func init() {
}
