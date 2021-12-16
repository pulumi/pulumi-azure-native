


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExport(ctx *pulumi.Context, args *LookupExportArgs, opts ...pulumi.InvokeOption) (*LookupExportResult, error) {
	var rv LookupExportResult
	err := ctx.Invoke("azure-native:costmanagement/v20210101:getExport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExportArgs struct {
	Expand     *string `pulumi:"expand"`
	ExportName string  `pulumi:"exportName"`
	Scope      string  `pulumi:"scope"`
}


type LookupExportResult struct {
	Definition          ExportDefinitionResponse           `pulumi:"definition"`
	DeliveryInfo        ExportDeliveryInfoResponse         `pulumi:"deliveryInfo"`
	ETag                *string                            `pulumi:"eTag"`
	Format              *string                            `pulumi:"format"`
	Id                  string                             `pulumi:"id"`
	Name                string                             `pulumi:"name"`
	NextRunTimeEstimate string                             `pulumi:"nextRunTimeEstimate"`
	PartitionData       *bool                              `pulumi:"partitionData"`
	RunHistory          *ExportExecutionListResultResponse `pulumi:"runHistory"`
	Schedule            *ExportScheduleResponse            `pulumi:"schedule"`
	Type                string                             `pulumi:"type"`
}
