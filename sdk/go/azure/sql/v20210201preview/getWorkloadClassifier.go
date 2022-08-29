


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadClassifier(ctx *pulumi.Context, args *LookupWorkloadClassifierArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadClassifierResult, error) {
	var rv LookupWorkloadClassifierResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getWorkloadClassifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadClassifierArgs struct {
	DatabaseName           string `pulumi:"databaseName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerName             string `pulumi:"serverName"`
	WorkloadClassifierName string `pulumi:"workloadClassifierName"`
	WorkloadGroupName      string `pulumi:"workloadGroupName"`
}


type LookupWorkloadClassifierResult struct {
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

func LookupWorkloadClassifierOutput(ctx *pulumi.Context, args LookupWorkloadClassifierOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadClassifierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkloadClassifierResult, error) {
			args := v.(LookupWorkloadClassifierArgs)
			r, err := LookupWorkloadClassifier(ctx, &args, opts...)
			var s LookupWorkloadClassifierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkloadClassifierResultOutput)
}

type LookupWorkloadClassifierOutputArgs struct {
	DatabaseName           pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName             pulumi.StringInput `pulumi:"serverName"`
	WorkloadClassifierName pulumi.StringInput `pulumi:"workloadClassifierName"`
	WorkloadGroupName      pulumi.StringInput `pulumi:"workloadGroupName"`
}

func (LookupWorkloadClassifierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadClassifierArgs)(nil)).Elem()
}


type LookupWorkloadClassifierResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadClassifierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadClassifierResult)(nil)).Elem()
}

func (o LookupWorkloadClassifierResultOutput) ToLookupWorkloadClassifierResultOutput() LookupWorkloadClassifierResultOutput {
	return o
}

func (o LookupWorkloadClassifierResultOutput) ToLookupWorkloadClassifierResultOutputWithContext(ctx context.Context) LookupWorkloadClassifierResultOutput {
	return o
}

func (o LookupWorkloadClassifierResultOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) *string { return v.Context }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadClassifierResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadClassifierResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkloadClassifierResultOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) *string { return v.Importance }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadClassifierResultOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadClassifierResultOutput) MemberName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) string { return v.MemberName }).(pulumi.StringOutput)
}

func (o LookupWorkloadClassifierResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkloadClassifierResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadClassifierResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadClassifierResultOutput{})
}
