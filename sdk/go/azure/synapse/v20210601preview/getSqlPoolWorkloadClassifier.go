


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlPoolWorkloadClassifier(ctx *pulumi.Context, args *LookupSqlPoolWorkloadClassifierArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolWorkloadClassifierResult, error) {
	var rv LookupSqlPoolWorkloadClassifierResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getSqlPoolWorkloadClassifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolWorkloadClassifierArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SqlPoolName            string `pulumi:"sqlPoolName"`
	WorkloadClassifierName string `pulumi:"workloadClassifierName"`
	WorkloadGroupName      string `pulumi:"workloadGroupName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type LookupSqlPoolWorkloadClassifierResult struct {
	Context    *string `pulumi:"context"`
	EndTime    *string `pulumi:"endTime"`
	Id         string  `pulumi:"id"`
	Importance *string `pulumi:"importance"`
	Label      *string `pulumi:"label"`
	MemberName string  `pulumi:"memberName"`
	Name       string  `pulumi:"name"`
	StartTime  *string `pulumi:"startTime"`
	Type       string  `pulumi:"type"`
}

func LookupSqlPoolWorkloadClassifierOutput(ctx *pulumi.Context, args LookupSqlPoolWorkloadClassifierOutputArgs, opts ...pulumi.InvokeOption) LookupSqlPoolWorkloadClassifierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlPoolWorkloadClassifierResult, error) {
			args := v.(LookupSqlPoolWorkloadClassifierArgs)
			r, err := LookupSqlPoolWorkloadClassifier(ctx, &args, opts...)
			var s LookupSqlPoolWorkloadClassifierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlPoolWorkloadClassifierResultOutput)
}

type LookupSqlPoolWorkloadClassifierOutputArgs struct {
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlPoolName            pulumi.StringInput `pulumi:"sqlPoolName"`
	WorkloadClassifierName pulumi.StringInput `pulumi:"workloadClassifierName"`
	WorkloadGroupName      pulumi.StringInput `pulumi:"workloadGroupName"`
	WorkspaceName          pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlPoolWorkloadClassifierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolWorkloadClassifierArgs)(nil)).Elem()
}


type LookupSqlPoolWorkloadClassifierResultOutput struct{ *pulumi.OutputState }

func (LookupSqlPoolWorkloadClassifierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolWorkloadClassifierResult)(nil)).Elem()
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) ToLookupSqlPoolWorkloadClassifierResultOutput() LookupSqlPoolWorkloadClassifierResultOutput {
	return o
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) ToLookupSqlPoolWorkloadClassifierResultOutputWithContext(ctx context.Context) LookupSqlPoolWorkloadClassifierResultOutput {
	return o
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) *string { return v.Context }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) *string { return v.Importance }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) MemberName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) string { return v.MemberName }).(pulumi.StringOutput)
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlPoolWorkloadClassifierResultOutput{})
}
