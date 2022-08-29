


package v20180531

import (
	"context"
	"reflect"

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

func LookupReportConfigByResourceGroupNameOutput(ctx *pulumi.Context, args LookupReportConfigByResourceGroupNameOutputArgs, opts ...pulumi.InvokeOption) LookupReportConfigByResourceGroupNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReportConfigByResourceGroupNameResult, error) {
			args := v.(LookupReportConfigByResourceGroupNameArgs)
			r, err := LookupReportConfigByResourceGroupName(ctx, &args, opts...)
			var s LookupReportConfigByResourceGroupNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReportConfigByResourceGroupNameResultOutput)
}

type LookupReportConfigByResourceGroupNameOutputArgs struct {
	ReportConfigName  pulumi.StringInput `pulumi:"reportConfigName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReportConfigByResourceGroupNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportConfigByResourceGroupNameArgs)(nil)).Elem()
}


type LookupReportConfigByResourceGroupNameResultOutput struct{ *pulumi.OutputState }

func (LookupReportConfigByResourceGroupNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportConfigByResourceGroupNameResult)(nil)).Elem()
}

func (o LookupReportConfigByResourceGroupNameResultOutput) ToLookupReportConfigByResourceGroupNameResultOutput() LookupReportConfigByResourceGroupNameResultOutput {
	return o
}

func (o LookupReportConfigByResourceGroupNameResultOutput) ToLookupReportConfigByResourceGroupNameResultOutputWithContext(ctx context.Context) LookupReportConfigByResourceGroupNameResultOutput {
	return o
}

func (o LookupReportConfigByResourceGroupNameResultOutput) Definition() ReportConfigDefinitionResponseOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) ReportConfigDefinitionResponse {
		return v.Definition
	}).(ReportConfigDefinitionResponseOutput)
}

func (o LookupReportConfigByResourceGroupNameResultOutput) DeliveryInfo() ReportConfigDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) ReportConfigDeliveryInfoResponse {
		return v.DeliveryInfo
	}).(ReportConfigDeliveryInfoResponseOutput)
}

func (o LookupReportConfigByResourceGroupNameResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LookupReportConfigByResourceGroupNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReportConfigByResourceGroupNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReportConfigByResourceGroupNameResultOutput) Schedule() ReportConfigScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) *ReportConfigScheduleResponse { return v.Schedule }).(ReportConfigScheduleResponsePtrOutput)
}

func (o LookupReportConfigByResourceGroupNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupReportConfigByResourceGroupNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportConfigByResourceGroupNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReportConfigByResourceGroupNameResultOutput{})
}
