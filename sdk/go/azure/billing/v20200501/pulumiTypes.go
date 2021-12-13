


package v20200501

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
	BillingProfileSystemId         string              `pulumi:"billingProfileSystemId"`
	EnabledAzurePlans              []AzurePlanResponse `pulumi:"enabledAzurePlans"`
	InvoiceSectionDisplayName      string              `pulumi:"invoiceSectionDisplayName"`
	InvoiceSectionId               string              `pulumi:"invoiceSectionId"`
	InvoiceSectionSystemId         string              `pulumi:"invoiceSectionSystemId"`
}

func init() {
}
