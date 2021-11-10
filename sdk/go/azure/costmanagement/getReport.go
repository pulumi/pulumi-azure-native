


package costmanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReport(ctx *pulumi.Context, args *LookupReportArgs, opts ...pulumi.InvokeOption) (*LookupReportResult, error) {
	var rv LookupReportResult
	err := ctx.Invoke("azure-native:costmanagement:getReport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportArgs struct {
	ReportName string `pulumi:"reportName"`
}


type LookupReportResult struct {
	Definition   ReportDefinitionResponse   `pulumi:"definition"`
	DeliveryInfo ReportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	Format       *string                    `pulumi:"format"`
	Id           string                     `pulumi:"id"`
	Name         string                     `pulumi:"name"`
	Schedule     *ReportScheduleResponse    `pulumi:"schedule"`
	Tags         map[string]string          `pulumi:"tags"`
	Type         string                     `pulumi:"type"`
}
