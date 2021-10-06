


package billing

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBillingAccountInvoiceSectionsByCreateSubscriptionPermission(ctx *pulumi.Context, args *ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionArgs, opts ...pulumi.InvokeOption) (*ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult, error) {
	var rv ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult
	err := ctx.Invoke("azure-native:billing:listBillingAccountInvoiceSectionsByCreateSubscriptionPermission", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionArgs struct {
	BillingAccountName string `pulumi:"billingAccountName"`
}


type ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult struct {
	NextLink string                                          `pulumi:"nextLink"`
	Value    []InvoiceSectionWithCreateSubPermissionResponse `pulumi:"value"`
}
