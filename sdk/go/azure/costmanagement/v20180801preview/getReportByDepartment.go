


package v20180801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReportByDepartment(ctx *pulumi.Context, args *LookupReportByDepartmentArgs, opts ...pulumi.InvokeOption) (*LookupReportByDepartmentResult, error) {
	var rv LookupReportByDepartmentResult
	err := ctx.Invoke("azure-native:costmanagement/v20180801preview:getReportByDepartment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportByDepartmentArgs struct {
	DepartmentId string `pulumi:"departmentId"`
	ReportName   string `pulumi:"reportName"`
}


type LookupReportByDepartmentResult struct {
	Definition   ReportDefinitionResponse   `pulumi:"definition"`
	DeliveryInfo ReportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	Format       *string                    `pulumi:"format"`
	Id           string                     `pulumi:"id"`
	Name         string                     `pulumi:"name"`
	Schedule     *ReportScheduleResponse    `pulumi:"schedule"`
	Tags         map[string]string          `pulumi:"tags"`
	Type         string                     `pulumi:"type"`
}
