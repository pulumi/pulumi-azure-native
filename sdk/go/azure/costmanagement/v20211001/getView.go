


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupView(ctx *pulumi.Context, args *LookupViewArgs, opts ...pulumi.InvokeOption) (*LookupViewResult, error) {
	var rv LookupViewResult
	err := ctx.Invoke("azure-native:costmanagement/v20211001:getView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewArgs struct {
	ViewName string `pulumi:"viewName"`
}


type LookupViewResult struct {
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
