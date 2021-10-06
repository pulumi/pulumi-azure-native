


package v20180301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMetricAlert(ctx *pulumi.Context, args *LookupMetricAlertArgs, opts ...pulumi.InvokeOption) (*LookupMetricAlertResult, error) {
	var rv LookupMetricAlertResult
	err := ctx.Invoke("azure-native:insights/v20180301:getMetricAlert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetricAlertArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupMetricAlertResult struct {
	Actions              []MetricAlertActionResponse `pulumi:"actions"`
	AutoMitigate         *bool                       `pulumi:"autoMitigate"`
	Criteria             interface{}                 `pulumi:"criteria"`
	Description          *string                     `pulumi:"description"`
	Enabled              bool                        `pulumi:"enabled"`
	EvaluationFrequency  string                      `pulumi:"evaluationFrequency"`
	Id                   string                      `pulumi:"id"`
	IsMigrated           bool                        `pulumi:"isMigrated"`
	LastUpdatedTime      string                      `pulumi:"lastUpdatedTime"`
	Location             string                      `pulumi:"location"`
	Name                 string                      `pulumi:"name"`
	Scopes               []string                    `pulumi:"scopes"`
	Severity             int                         `pulumi:"severity"`
	Tags                 map[string]string           `pulumi:"tags"`
	TargetResourceRegion *string                     `pulumi:"targetResourceRegion"`
	TargetResourceType   *string                     `pulumi:"targetResourceType"`
	Type                 string                      `pulumi:"type"`
	WindowSize           string                      `pulumi:"windowSize"`
}
