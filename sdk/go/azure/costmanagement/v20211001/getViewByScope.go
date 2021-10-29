


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupViewByScope(ctx *pulumi.Context, args *LookupViewByScopeArgs, opts ...pulumi.InvokeOption) (*LookupViewByScopeResult, error) {
	var rv LookupViewByScopeResult
	err := ctx.Invoke("azure-native:costmanagement/v20211001:getViewByScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewByScopeArgs struct {
	Scope    string `pulumi:"scope"`
	ViewName string `pulumi:"viewName"`
}


type LookupViewByScopeResult struct {
	Accumulated *string                         `pulumi:"accumulated"`
	Chart       *string                         `pulumi:"chart"`
	CreatedOn   string                          `pulumi:"createdOn"`
	Dataset     *ReportConfigDatasetResponse    `pulumi:"dataset"`
	DisplayName *string                         `pulumi:"displayName"`
	ETag        *string                         `pulumi:"eTag"`
	Id          string                          `pulumi:"id"`
	Kpis        []KpiPropertiesResponse         `pulumi:"kpis"`
	Metric      *string                         `pulumi:"metric"`
	ModifiedOn  string                          `pulumi:"modifiedOn"`
	Name        string                          `pulumi:"name"`
	Pivots      []PivotPropertiesResponse       `pulumi:"pivots"`
	Scope       *string                         `pulumi:"scope"`
	TimePeriod  *ReportConfigTimePeriodResponse `pulumi:"timePeriod"`
	Timeframe   string                          `pulumi:"timeframe"`
	Type        string                          `pulumi:"type"`
}
