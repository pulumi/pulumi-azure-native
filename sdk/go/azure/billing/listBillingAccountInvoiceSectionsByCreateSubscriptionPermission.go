


package billing

import (
	"context"
	"reflect"

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

func ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionOutput(ctx *pulumi.Context, args ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionOutputArgs, opts ...pulumi.InvokeOption) ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult, error) {
			args := v.(ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionArgs)
			r, err := ListBillingAccountInvoiceSectionsByCreateSubscriptionPermission(ctx, &args, opts...)
			var s ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput)
}

type ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionOutputArgs struct {
	BillingAccountName pulumi.StringInput `pulumi:"billingAccountName"`
}

func (ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionArgs)(nil)).Elem()
}


type ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput struct{ *pulumi.OutputState }

func (ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult)(nil)).Elem()
}

func (o ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput) ToListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput() ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput {
	return o
}

func (o ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput) ToListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutputWithContext(ctx context.Context) ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput {
	return o
}

func (o ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult) string {
		return v.NextLink
	}).(pulumi.StringOutput)
}

func (o ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput) Value() InvoiceSectionWithCreateSubPermissionResponseArrayOutput {
	return o.ApplyT(func(v ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult) []InvoiceSectionWithCreateSubPermissionResponse {
		return v.Value
	}).(InvoiceSectionWithCreateSubPermissionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput{})
}
