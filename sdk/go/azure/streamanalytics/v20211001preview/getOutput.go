


package v20211001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOutput(ctx *pulumi.Context, args *LookupOutputArgs, opts ...pulumi.InvokeOption) (*LookupOutputResult, error) {
	var rv LookupOutputResult
	err := ctx.Invoke("azure-native:streamanalytics/v20211001preview:getOutput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOutputArgs struct {
	JobName           string `pulumi:"jobName"`
	OutputName        string `pulumi:"outputName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOutputResult struct {
	Datasource    interface{}         `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponse `pulumi:"diagnostics"`
	Etag          string              `pulumi:"etag"`
	Id            string              `pulumi:"id"`
	Name          *string             `pulumi:"name"`
	Serialization interface{}         `pulumi:"serialization"`
	SizeWindow    *float64            `pulumi:"sizeWindow"`
	TimeWindow    *string             `pulumi:"timeWindow"`
	Type          string              `pulumi:"type"`
}
