


package costmanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReportByBillingAccount(ctx *pulumi.Context, args *LookupReportByBillingAccountArgs, opts ...pulumi.InvokeOption) (*LookupReportByBillingAccountResult, error) {
	var rv LookupReportByBillingAccountResult
	err := ctx.Invoke("azure-native:costmanagement:getReportByBillingAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportByBillingAccountArgs struct {
	BillingAccountId string `pulumi:"billingAccountId"`
	ReportName       string `pulumi:"reportName"`
}


type LookupReportByBillingAccountResult struct {
	Definition   ReportDefinitionResponse   `pulumi:"definition"`
	DeliveryInfo ReportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	Format       *string                    `pulumi:"format"`
	Id           string                     `pulumi:"id"`
	Name         string                     `pulumi:"name"`
	Schedule     *ReportScheduleResponse    `pulumi:"schedule"`
	Tags         map[string]string          `pulumi:"tags"`
	Type         string                     `pulumi:"type"`
}
