


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFileImport(ctx *pulumi.Context, args *LookupFileImportArgs, opts ...pulumi.InvokeOption) (*LookupFileImportResult, error) {
	var rv LookupFileImportResult
	err := ctx.Invoke("azure-native:securityinsights/v20220801preview:getFileImport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileImportArgs struct {
	FileImportId      string `pulumi:"fileImportId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupFileImportResult struct {
	ContentType             string                    `pulumi:"contentType"`
	CreatedTimeUTC          string                    `pulumi:"createdTimeUTC"`
	ErrorFile               FileMetadataResponse      `pulumi:"errorFile"`
	ErrorsPreview           []ValidationErrorResponse `pulumi:"errorsPreview"`
	FilesValidUntilTimeUTC  string                    `pulumi:"filesValidUntilTimeUTC"`
	Id                      string                    `pulumi:"id"`
	ImportFile              FileMetadataResponse      `pulumi:"importFile"`
	ImportValidUntilTimeUTC string                    `pulumi:"importValidUntilTimeUTC"`
	IngestedRecordCount     int                       `pulumi:"ingestedRecordCount"`
	IngestionMode           string                    `pulumi:"ingestionMode"`
	Name                    string                    `pulumi:"name"`
	Source                  string                    `pulumi:"source"`
	State                   string                    `pulumi:"state"`
	SystemData              SystemDataResponse        `pulumi:"systemData"`
	TotalRecordCount        int                       `pulumi:"totalRecordCount"`
	Type                    string                    `pulumi:"type"`
	ValidRecordCount        int                       `pulumi:"validRecordCount"`
}

func LookupFileImportOutput(ctx *pulumi.Context, args LookupFileImportOutputArgs, opts ...pulumi.InvokeOption) LookupFileImportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileImportResult, error) {
			args := v.(LookupFileImportArgs)
			r, err := LookupFileImport(ctx, &args, opts...)
			var s LookupFileImportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileImportResultOutput)
}

type LookupFileImportOutputArgs struct {
	FileImportId      pulumi.StringInput `pulumi:"fileImportId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupFileImportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileImportArgs)(nil)).Elem()
}


type LookupFileImportResultOutput struct{ *pulumi.OutputState }

func (LookupFileImportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileImportResult)(nil)).Elem()
}

func (o LookupFileImportResultOutput) ToLookupFileImportResultOutput() LookupFileImportResultOutput {
	return o
}

func (o LookupFileImportResultOutput) ToLookupFileImportResultOutputWithContext(ctx context.Context) LookupFileImportResultOutput {
	return o
}

func (o LookupFileImportResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o LookupFileImportResultOutput) CreatedTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.CreatedTimeUTC }).(pulumi.StringOutput)
}

func (o LookupFileImportResultOutput) ErrorFile() FileMetadataResponseOutput {
	return o.ApplyT(func(v LookupFileImportResult) FileMetadataResponse { return v.ErrorFile }).(FileMetadataResponseOutput)
}

func (o LookupFileImportResultOutput) ErrorsPreview() ValidationErrorResponseArrayOutput {
	return o.ApplyT(func(v LookupFileImportResult) []ValidationErrorResponse { return v.ErrorsPreview }).(ValidationErrorResponseArrayOutput)
}

func (o LookupFileImportResultOutput) FilesValidUntilTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.FilesValidUntilTimeUTC }).(pulumi.StringOutput)
}

func (o LookupFileImportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFileImportResultOutput) ImportFile() FileMetadataResponseOutput {
	return o.ApplyT(func(v LookupFileImportResult) FileMetadataResponse { return v.ImportFile }).(FileMetadataResponseOutput)
}

func (o LookupFileImportResultOutput) ImportValidUntilTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.ImportValidUntilTimeUTC }).(pulumi.StringOutput)
}

func (o LookupFileImportResultOutput) IngestedRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFileImportResult) int { return v.IngestedRecordCount }).(pulumi.IntOutput)
}

func (o LookupFileImportResultOutput) IngestionMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.IngestionMode }).(pulumi.StringOutput)
}

func (o LookupFileImportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFileImportResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.Source }).(pulumi.StringOutput)
}

func (o LookupFileImportResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupFileImportResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFileImportResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupFileImportResultOutput) TotalRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFileImportResult) int { return v.TotalRecordCount }).(pulumi.IntOutput)
}

func (o LookupFileImportResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupFileImportResultOutput) ValidRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFileImportResult) int { return v.ValidRecordCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileImportResultOutput{})
}
