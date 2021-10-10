


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExport(ctx *pulumi.Context, args *LookupExportArgs, opts ...pulumi.InvokeOption) (*LookupExportResult, error) {
	var rv LookupExportResult
	err := ctx.Invoke("azure-native:costmanagement/v20191101:getExport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExportArgs struct {
	ExportName string `pulumi:"exportName"`
	Scope      string `pulumi:"scope"`
}


type LookupExportResult struct {
	Definition   ExportDefinitionResponse   `pulumi:"definition"`
	DeliveryInfo ExportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	ETag         *string                    `pulumi:"eTag"`
	Format       *string                    `pulumi:"format"`
	Id           string                     `pulumi:"id"`
	Name         string                     `pulumi:"name"`
	Schedule     *ExportScheduleResponse    `pulumi:"schedule"`
	Type         string                     `pulumi:"type"`
}
