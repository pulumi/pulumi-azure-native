


package v20180531

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReportConfigByResourceGroupName(ctx *pulumi.Context, args *LookupReportConfigByResourceGroupNameArgs, opts ...pulumi.InvokeOption) (*LookupReportConfigByResourceGroupNameResult, error) {
	var rv LookupReportConfigByResourceGroupNameResult
	err := ctx.Invoke("azure-native:costmanagement/v20180531:getReportConfigByResourceGroupName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportConfigByResourceGroupNameArgs struct {
	ReportConfigName  string `pulumi:"reportConfigName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupReportConfigByResourceGroupNameResult struct {
	Definition   ReportConfigDefinitionResponse   `pulumi:"definition"`
	DeliveryInfo ReportConfigDeliveryInfoResponse `pulumi:"deliveryInfo"`
	Format       *string                          `pulumi:"format"`
	Id           string                           `pulumi:"id"`
	Name         string                           `pulumi:"name"`
	Schedule     *ReportConfigScheduleResponse    `pulumi:"schedule"`
	Tags         map[string]string                `pulumi:"tags"`
	Type         string                           `pulumi:"type"`
}
