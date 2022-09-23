


package v20180801preview

import (
	"context"
	"reflect"

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

func LookupReportByDepartmentOutput(ctx *pulumi.Context, args LookupReportByDepartmentOutputArgs, opts ...pulumi.InvokeOption) LookupReportByDepartmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReportByDepartmentResult, error) {
			args := v.(LookupReportByDepartmentArgs)
			r, err := LookupReportByDepartment(ctx, &args, opts...)
			var s LookupReportByDepartmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReportByDepartmentResultOutput)
}

type LookupReportByDepartmentOutputArgs struct {
	DepartmentId pulumi.StringInput `pulumi:"departmentId"`
	ReportName   pulumi.StringInput `pulumi:"reportName"`
}

func (LookupReportByDepartmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportByDepartmentArgs)(nil)).Elem()
}


type LookupReportByDepartmentResultOutput struct{ *pulumi.OutputState }

func (LookupReportByDepartmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportByDepartmentResult)(nil)).Elem()
}

func (o LookupReportByDepartmentResultOutput) ToLookupReportByDepartmentResultOutput() LookupReportByDepartmentResultOutput {
	return o
}

func (o LookupReportByDepartmentResultOutput) ToLookupReportByDepartmentResultOutputWithContext(ctx context.Context) LookupReportByDepartmentResultOutput {
	return o
}

func (o LookupReportByDepartmentResultOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v LookupReportByDepartmentResult) ReportDefinitionResponse { return v.Definition }).(ReportDefinitionResponseOutput)
}

func (o LookupReportByDepartmentResultOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupReportByDepartmentResult) ReportDeliveryInfoResponse { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

func (o LookupReportByDepartmentResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReportByDepartmentResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LookupReportByDepartmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByDepartmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReportByDepartmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByDepartmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReportByDepartmentResultOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupReportByDepartmentResult) *ReportScheduleResponse { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

func (o LookupReportByDepartmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReportByDepartmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupReportByDepartmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByDepartmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReportByDepartmentResultOutput{})
}
