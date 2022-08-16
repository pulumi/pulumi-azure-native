


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlPoolWorkloadGroup(ctx *pulumi.Context, args *LookupSqlPoolWorkloadGroupArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolWorkloadGroupResult, error) {
	var rv LookupSqlPoolWorkloadGroupResult
	err := ctx.Invoke("azure-native:synapse/v20210301:getSqlPoolWorkloadGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolWorkloadGroupArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SqlPoolName       string `pulumi:"sqlPoolName"`
	WorkloadGroupName string `pulumi:"workloadGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupSqlPoolWorkloadGroupResult struct {
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

func LookupSqlPoolWorkloadGroupOutput(ctx *pulumi.Context, args LookupSqlPoolWorkloadGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSqlPoolWorkloadGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlPoolWorkloadGroupResult, error) {
			args := v.(LookupSqlPoolWorkloadGroupArgs)
			r, err := LookupSqlPoolWorkloadGroup(ctx, &args, opts...)
			var s LookupSqlPoolWorkloadGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlPoolWorkloadGroupResultOutput)
}

type LookupSqlPoolWorkloadGroupOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlPoolName       pulumi.StringInput `pulumi:"sqlPoolName"`
	WorkloadGroupName pulumi.StringInput `pulumi:"workloadGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlPoolWorkloadGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolWorkloadGroupArgs)(nil)).Elem()
}


type LookupSqlPoolWorkloadGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSqlPoolWorkloadGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolWorkloadGroupResult)(nil)).Elem()
}

func (o LookupSqlPoolWorkloadGroupResultOutput) ToLookupSqlPoolWorkloadGroupResultOutput() LookupSqlPoolWorkloadGroupResultOutput {
	return o
}

func (o LookupSqlPoolWorkloadGroupResultOutput) ToLookupSqlPoolWorkloadGroupResultOutputWithContext(ctx context.Context) LookupSqlPoolWorkloadGroupResultOutput {
	return o
}

func (o LookupSqlPoolWorkloadGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlPoolWorkloadGroupResultOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) *string { return v.Importance }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolWorkloadGroupResultOutput) MaxResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) int { return v.MaxResourcePercent }).(pulumi.IntOutput)
}

func (o LookupSqlPoolWorkloadGroupResultOutput) MaxResourcePercentPerRequest() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) *float64 { return v.MaxResourcePercentPerRequest }).(pulumi.Float64PtrOutput)
}

func (o LookupSqlPoolWorkloadGroupResultOutput) MinResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) int { return v.MinResourcePercent }).(pulumi.IntOutput)
}

func (o LookupSqlPoolWorkloadGroupResultOutput) MinResourcePercentPerRequest() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) float64 { return v.MinResourcePercentPerRequest }).(pulumi.Float64Output)
}

func (o LookupSqlPoolWorkloadGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlPoolWorkloadGroupResultOutput) QueryExecutionTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) *int { return v.QueryExecutionTimeout }).(pulumi.IntPtrOutput)
}

func (o LookupSqlPoolWorkloadGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlPoolWorkloadGroupResultOutput{})
}
