


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTable(ctx *pulumi.Context, args *LookupTableArgs, opts ...pulumi.InvokeOption) (*LookupTableResult, error) {
	var rv LookupTableResult
	err := ctx.Invoke("azure-native:operationalinsights/v20211201preview:getTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTableArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TableName         string `pulumi:"tableName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupTableResult struct {
	ArchiveRetentionInDays int                       `pulumi:"archiveRetentionInDays"`
	Id                     string                    `pulumi:"id"`
	LastPlanModifiedDate   string                    `pulumi:"lastPlanModifiedDate"`
	Name                   string                    `pulumi:"name"`
	Plan                   *string                   `pulumi:"plan"`
	ProvisioningState      string                    `pulumi:"provisioningState"`
	RestoredLogs           *RestoredLogsResponse     `pulumi:"restoredLogs"`
	ResultStatistics       *ResultStatisticsResponse `pulumi:"resultStatistics"`
	RetentionInDays        *int                      `pulumi:"retentionInDays"`
	Schema                 *SchemaResponse           `pulumi:"schema"`
	SearchResults          *SearchResultsResponse    `pulumi:"searchResults"`
	SystemData             SystemDataResponse        `pulumi:"systemData"`
	TotalRetentionInDays   *int                      `pulumi:"totalRetentionInDays"`
	Type                   string                    `pulumi:"type"`
}

func LookupTableOutput(ctx *pulumi.Context, args LookupTableOutputArgs, opts ...pulumi.InvokeOption) LookupTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTableResult, error) {
			args := v.(LookupTableArgs)
			r, err := LookupTable(ctx, &args, opts...)
			var s LookupTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTableResultOutput)
}

type LookupTableOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TableName         pulumi.StringInput `pulumi:"tableName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableArgs)(nil)).Elem()
}


type LookupTableResultOutput struct{ *pulumi.OutputState }

func (LookupTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableResult)(nil)).Elem()
}

func (o LookupTableResultOutput) ToLookupTableResultOutput() LookupTableResultOutput {
	return o
}

func (o LookupTableResultOutput) ToLookupTableResultOutputWithContext(ctx context.Context) LookupTableResultOutput {
	return o
}

func (o LookupTableResultOutput) ArchiveRetentionInDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTableResult) int { return v.ArchiveRetentionInDays }).(pulumi.IntOutput)
}

func (o LookupTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTableResultOutput) LastPlanModifiedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResult) string { return v.LastPlanModifiedDate }).(pulumi.StringOutput)
}

func (o LookupTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTableResultOutput) Plan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTableResult) *string { return v.Plan }).(pulumi.StringPtrOutput)
}

func (o LookupTableResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTableResultOutput) RestoredLogs() RestoredLogsResponsePtrOutput {
	return o.ApplyT(func(v LookupTableResult) *RestoredLogsResponse { return v.RestoredLogs }).(RestoredLogsResponsePtrOutput)
}

func (o LookupTableResultOutput) ResultStatistics() ResultStatisticsResponsePtrOutput {
	return o.ApplyT(func(v LookupTableResult) *ResultStatisticsResponse { return v.ResultStatistics }).(ResultStatisticsResponsePtrOutput)
}

func (o LookupTableResultOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTableResult) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o LookupTableResultOutput) Schema() SchemaResponsePtrOutput {
	return o.ApplyT(func(v LookupTableResult) *SchemaResponse { return v.Schema }).(SchemaResponsePtrOutput)
}

func (o LookupTableResultOutput) SearchResults() SearchResultsResponsePtrOutput {
	return o.ApplyT(func(v LookupTableResult) *SearchResultsResponse { return v.SearchResults }).(SearchResultsResponsePtrOutput)
}

func (o LookupTableResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTableResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTableResultOutput) TotalRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTableResult) *int { return v.TotalRetentionInDays }).(pulumi.IntPtrOutput)
}

func (o LookupTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTableResultOutput{})
}
