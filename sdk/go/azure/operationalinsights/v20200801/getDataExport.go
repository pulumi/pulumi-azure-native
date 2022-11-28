


package v20200801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataExport(ctx *pulumi.Context, args *LookupDataExportArgs, opts ...pulumi.InvokeOption) (*LookupDataExportResult, error) {
	var rv LookupDataExportResult
	err := ctx.Invoke("azure-native:operationalinsights/v20200801:getDataExport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataExportArgs struct {
	DataExportName    string `pulumi:"dataExportName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDataExportResult struct {
	CreatedDate      *string  `pulumi:"createdDate"`
	DataExportId     *string  `pulumi:"dataExportId"`
	Enable           *bool    `pulumi:"enable"`
	EventHubName     *string  `pulumi:"eventHubName"`
	Id               string   `pulumi:"id"`
	LastModifiedDate *string  `pulumi:"lastModifiedDate"`
	Name             string   `pulumi:"name"`
	ResourceId       string   `pulumi:"resourceId"`
	TableNames       []string `pulumi:"tableNames"`
	Type             string   `pulumi:"type"`
}

func LookupDataExportOutput(ctx *pulumi.Context, args LookupDataExportOutputArgs, opts ...pulumi.InvokeOption) LookupDataExportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataExportResult, error) {
			args := v.(LookupDataExportArgs)
			r, err := LookupDataExport(ctx, &args, opts...)
			var s LookupDataExportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataExportResultOutput)
}

type LookupDataExportOutputArgs struct {
	DataExportName    pulumi.StringInput `pulumi:"dataExportName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDataExportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataExportArgs)(nil)).Elem()
}


type LookupDataExportResultOutput struct{ *pulumi.OutputState }

func (LookupDataExportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataExportResult)(nil)).Elem()
}

func (o LookupDataExportResultOutput) ToLookupDataExportResultOutput() LookupDataExportResultOutput {
	return o
}

func (o LookupDataExportResultOutput) ToLookupDataExportResultOutputWithContext(ctx context.Context) LookupDataExportResultOutput {
	return o
}

func (o LookupDataExportResultOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataExportResult) *string { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

func (o LookupDataExportResultOutput) DataExportId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataExportResult) *string { return v.DataExportId }).(pulumi.StringPtrOutput)
}

func (o LookupDataExportResultOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDataExportResult) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o LookupDataExportResultOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataExportResult) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o LookupDataExportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataExportResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataExportResultOutput) LastModifiedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataExportResult) *string { return v.LastModifiedDate }).(pulumi.StringPtrOutput)
}

func (o LookupDataExportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataExportResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataExportResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataExportResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o LookupDataExportResultOutput) TableNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDataExportResult) []string { return v.TableNames }).(pulumi.StringArrayOutput)
}

func (o LookupDataExportResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataExportResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataExportResultOutput{})
}
