


package v20180531

import (
	"context"
	"reflect"

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

func LookupReportConfigOutput(ctx *pulumi.Context, args LookupReportConfigOutputArgs, opts ...pulumi.InvokeOption) LookupReportConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReportConfigResult, error) {
			args := v.(LookupReportConfigArgs)
			r, err := LookupReportConfig(ctx, &args, opts...)
			var s LookupReportConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReportConfigResultOutput)
}

type LookupReportConfigOutputArgs struct {
	ReportConfigName pulumi.StringInput `pulumi:"reportConfigName"`
}

func (LookupReportConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportConfigArgs)(nil)).Elem()
}


type LookupReportConfigResultOutput struct{ *pulumi.OutputState }

func (LookupReportConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportConfigResult)(nil)).Elem()
}

func (o LookupReportConfigResultOutput) ToLookupReportConfigResultOutput() LookupReportConfigResultOutput {
	return o
}

func (o LookupReportConfigResultOutput) ToLookupReportConfigResultOutputWithContext(ctx context.Context) LookupReportConfigResultOutput {
	return o
}

func (o LookupReportConfigResultOutput) Definition() ReportConfigDefinitionResponseOutput {
	return o.ApplyT(func(v LookupReportConfigResult) ReportConfigDefinitionResponse { return v.Definition }).(ReportConfigDefinitionResponseOutput)
}

func (o LookupReportConfigResultOutput) DeliveryInfo() ReportConfigDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupReportConfigResult) ReportConfigDeliveryInfoResponse { return v.DeliveryInfo }).(ReportConfigDeliveryInfoResponseOutput)
}

func (o LookupReportConfigResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReportConfigResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LookupReportConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReportConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReportConfigResultOutput) Schedule() ReportConfigScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupReportConfigResult) *ReportConfigScheduleResponse { return v.Schedule }).(ReportConfigScheduleResponsePtrOutput)
}

func (o LookupReportConfigResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReportConfigResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupReportConfigResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportConfigResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReportConfigResultOutput{})
}
