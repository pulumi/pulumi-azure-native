


package costmanagement

import (
	"context"
	"reflect"

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

func LookupReportByBillingAccountOutput(ctx *pulumi.Context, args LookupReportByBillingAccountOutputArgs, opts ...pulumi.InvokeOption) LookupReportByBillingAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReportByBillingAccountResult, error) {
			args := v.(LookupReportByBillingAccountArgs)
			r, err := LookupReportByBillingAccount(ctx, &args, opts...)
			var s LookupReportByBillingAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReportByBillingAccountResultOutput)
}

type LookupReportByBillingAccountOutputArgs struct {
	BillingAccountId pulumi.StringInput `pulumi:"billingAccountId"`
	ReportName       pulumi.StringInput `pulumi:"reportName"`
}

func (LookupReportByBillingAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportByBillingAccountArgs)(nil)).Elem()
}


type LookupReportByBillingAccountResultOutput struct{ *pulumi.OutputState }

func (LookupReportByBillingAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportByBillingAccountResult)(nil)).Elem()
}

func (o LookupReportByBillingAccountResultOutput) ToLookupReportByBillingAccountResultOutput() LookupReportByBillingAccountResultOutput {
	return o
}

func (o LookupReportByBillingAccountResultOutput) ToLookupReportByBillingAccountResultOutputWithContext(ctx context.Context) LookupReportByBillingAccountResultOutput {
	return o
}

func (o LookupReportByBillingAccountResultOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v LookupReportByBillingAccountResult) ReportDefinitionResponse { return v.Definition }).(ReportDefinitionResponseOutput)
}

func (o LookupReportByBillingAccountResultOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupReportByBillingAccountResult) ReportDeliveryInfoResponse { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

func (o LookupReportByBillingAccountResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReportByBillingAccountResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LookupReportByBillingAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByBillingAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReportByBillingAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByBillingAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReportByBillingAccountResultOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupReportByBillingAccountResult) *ReportScheduleResponse { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

func (o LookupReportByBillingAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReportByBillingAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupReportByBillingAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByBillingAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReportByBillingAccountResultOutput{})
}
