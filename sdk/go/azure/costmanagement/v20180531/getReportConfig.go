


package v20180531

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReportConfig(ctx *pulumi.Context, args *LookupReportConfigArgs, opts ...pulumi.InvokeOption) (*LookupReportConfigResult, error) {
	var rv LookupReportConfigResult
	err := ctx.Invoke("azure-native:costmanagement/v20180531:getReportConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportConfigArgs struct {
	ReportConfigName string `pulumi:"reportConfigName"`
}


type LookupReportConfigResult struct {
	Definition   ReportConfigDefinitionResponse   `pulumi:"definition"`
	DeliveryInfo ReportConfigDeliveryInfoResponse `pulumi:"deliveryInfo"`
	Format       *string                          `pulumi:"format"`
	Id           string                           `pulumi:"id"`
	Name         string                           `pulumi:"name"`
	Schedule     *ReportConfigScheduleResponse    `pulumi:"schedule"`
	Tags         map[string]string                `pulumi:"tags"`
	Type         string                           `pulumi:"type"`
}
