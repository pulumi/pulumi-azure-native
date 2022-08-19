


package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadGroup(ctx *pulumi.Context, args *LookupWorkloadGroupArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadGroupResult, error) {
	var rv LookupWorkloadGroupResult
	err := ctx.Invoke("azure-native:sql:getWorkloadGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadGroupArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	WorkloadGroupName string `pulumi:"workloadGroupName"`
}


type LookupWorkloadGroupResult struct {
	Id                           string   `pulumi:"id"`
	Importance                   *string  `pulumi:"importance"`
	MaxResourcePercent           int      `pulumi:"maxResourcePercent"`
	MaxResourcePercentPerRequest *float64 `pulumi:"maxResourcePercentPerRequest"`
	MinResourcePercent           int      `pulumi:"minResourcePercent"`
	MinResourcePercentPerRequest float64  `pulumi:"minResourcePercentPerRequest"`
	Name                         string   `pulumi:"name"`
	QueryExecutionTimeout        *int     `pulumi:"queryExecutionTimeout"`
	Type                         string   `pulumi:"type"`
}

func LookupWorkloadGroupOutput(ctx *pulumi.Context, args LookupWorkloadGroupOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkloadGroupResult, error) {
			args := v.(LookupWorkloadGroupArgs)
			r, err := LookupWorkloadGroup(ctx, &args, opts...)
			var s LookupWorkloadGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkloadGroupResultOutput)
}

type LookupWorkloadGroupOutputArgs struct {
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
	WorkloadGroupName pulumi.StringInput `pulumi:"workloadGroupName"`
}

func (LookupWorkloadGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadGroupArgs)(nil)).Elem()
}


type LookupWorkloadGroupResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadGroupResult)(nil)).Elem()
}

func (o LookupWorkloadGroupResultOutput) ToLookupWorkloadGroupResultOutput() LookupWorkloadGroupResultOutput {
	return o
}

func (o LookupWorkloadGroupResultOutput) ToLookupWorkloadGroupResultOutputWithContext(ctx context.Context) LookupWorkloadGroupResultOutput {
	return o
}

func (o LookupWorkloadGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkloadGroupResultOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) *string { return v.Importance }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadGroupResultOutput) MaxResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) int { return v.MaxResourcePercent }).(pulumi.IntOutput)
}

func (o LookupWorkloadGroupResultOutput) MaxResourcePercentPerRequest() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) *float64 { return v.MaxResourcePercentPerRequest }).(pulumi.Float64PtrOutput)
}

func (o LookupWorkloadGroupResultOutput) MinResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) int { return v.MinResourcePercent }).(pulumi.IntOutput)
}

func (o LookupWorkloadGroupResultOutput) MinResourcePercentPerRequest() pulumi.Float64Output {
	return o.ApplyT(func(v LookupWorkloadGroupResult) float64 { return v.MinResourcePercentPerRequest }).(pulumi.Float64Output)
}

func (o LookupWorkloadGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkloadGroupResultOutput) QueryExecutionTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) *int { return v.QueryExecutionTimeout }).(pulumi.IntPtrOutput)
}

func (o LookupWorkloadGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadGroupResultOutput{})
}
