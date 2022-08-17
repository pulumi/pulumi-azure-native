


package v20191101

import (
	"context"
	"reflect"

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

func LookupExportOutput(ctx *pulumi.Context, args LookupExportOutputArgs, opts ...pulumi.InvokeOption) LookupExportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExportResult, error) {
			args := v.(LookupExportArgs)
			r, err := LookupExport(ctx, &args, opts...)
			var s LookupExportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExportResultOutput)
}

type LookupExportOutputArgs struct {
	ExportName pulumi.StringInput `pulumi:"exportName"`
	Scope      pulumi.StringInput `pulumi:"scope"`
}

func (LookupExportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportArgs)(nil)).Elem()
}


type LookupExportResultOutput struct{ *pulumi.OutputState }

func (LookupExportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportResult)(nil)).Elem()
}

func (o LookupExportResultOutput) ToLookupExportResultOutput() LookupExportResultOutput {
	return o
}

func (o LookupExportResultOutput) ToLookupExportResultOutputWithContext(ctx context.Context) LookupExportResultOutput {
	return o
}

func (o LookupExportResultOutput) Definition() ExportDefinitionResponseOutput {
	return o.ApplyT(func(v LookupExportResult) ExportDefinitionResponse { return v.Definition }).(ExportDefinitionResponseOutput)
}

func (o LookupExportResultOutput) DeliveryInfo() ExportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupExportResult) ExportDeliveryInfoResponse { return v.DeliveryInfo }).(ExportDeliveryInfoResponseOutput)
}

func (o LookupExportResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExportResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupExportResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExportResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LookupExportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExportResultOutput) Schedule() ExportScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupExportResult) *ExportScheduleResponse { return v.Schedule }).(ExportScheduleResponsePtrOutput)
}

func (o LookupExportResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExportResultOutput{})
}
