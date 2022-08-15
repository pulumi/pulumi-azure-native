


package v20180801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReportByResourceGroupName(ctx *pulumi.Context, args *LookupReportByResourceGroupNameArgs, opts ...pulumi.InvokeOption) (*LookupReportByResourceGroupNameResult, error) {
	var rv LookupReportByResourceGroupNameResult
	err := ctx.Invoke("azure-native:costmanagement/v20180801preview:getReportByResourceGroupName", args, &rv, opts...)
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

func LookupReportByResourceGroupNameOutput(ctx *pulumi.Context, args LookupReportByResourceGroupNameOutputArgs, opts ...pulumi.InvokeOption) LookupReportByResourceGroupNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReportByResourceGroupNameResult, error) {
			args := v.(LookupReportByResourceGroupNameArgs)
			r, err := LookupReportByResourceGroupName(ctx, &args, opts...)
			var s LookupReportByResourceGroupNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReportByResourceGroupNameResultOutput)
}

type LookupReportByResourceGroupNameOutputArgs struct {
	ReportName        pulumi.StringInput `pulumi:"reportName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReportByResourceGroupNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportByResourceGroupNameArgs)(nil)).Elem()
}


type LookupReportByResourceGroupNameResultOutput struct{ *pulumi.OutputState }

func (LookupReportByResourceGroupNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportByResourceGroupNameResult)(nil)).Elem()
}

func (o LookupReportByResourceGroupNameResultOutput) ToLookupReportByResourceGroupNameResultOutput() LookupReportByResourceGroupNameResultOutput {
	return o
}

func (o LookupReportByResourceGroupNameResultOutput) ToLookupReportByResourceGroupNameResultOutputWithContext(ctx context.Context) LookupReportByResourceGroupNameResultOutput {
	return o
}

func (o LookupReportByResourceGroupNameResultOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) ReportDefinitionResponse { return v.Definition }).(ReportDefinitionResponseOutput)
}

func (o LookupReportByResourceGroupNameResultOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) ReportDeliveryInfoResponse { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

func (o LookupReportByResourceGroupNameResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LookupReportByResourceGroupNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReportByResourceGroupNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReportByResourceGroupNameResultOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) *ReportScheduleResponse { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

func (o LookupReportByResourceGroupNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupReportByResourceGroupNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReportByResourceGroupNameResultOutput{})
}
