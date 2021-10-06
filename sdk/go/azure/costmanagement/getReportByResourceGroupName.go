


package costmanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReportByResourceGroupName(ctx *pulumi.Context, args *LookupReportByResourceGroupNameArgs, opts ...pulumi.InvokeOption) (*LookupReportByResourceGroupNameResult, error) {
	var rv LookupReportByResourceGroupNameResult
	err := ctx.Invoke("azure-native:costmanagement:getReportByResourceGroupName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportByResourceGroupNameArgs struct {
	ReportName        string `pulumi:"reportName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupReportByResourceGroupNameResult struct {
	Definition   ReportDefinitionResponse   `pulumi:"definition"`
	DeliveryInfo ReportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	Format       *string                    `pulumi:"format"`
	Id           string                     `pulumi:"id"`
	Name         string                     `pulumi:"name"`
	Schedule     *ReportScheduleResponse    `pulumi:"schedule"`
	Tags         map[string]string          `pulumi:"tags"`
	Type         string                     `pulumi:"type"`
}
