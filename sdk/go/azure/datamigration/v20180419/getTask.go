


package v20180419

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTask(ctx *pulumi.Context, args *LookupTaskArgs, opts ...pulumi.InvokeOption) (*LookupTaskResult, error) {
	var rv LookupTaskResult
	err := ctx.Invoke("azure-native:datamigration/v20180419:getTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTaskArgs struct {
	Expand      *string `pulumi:"expand"`
	GroupName   string  `pulumi:"groupName"`
	ProjectName string  `pulumi:"projectName"`
	ServiceName string  `pulumi:"serviceName"`
	TaskName    string  `pulumi:"taskName"`
}


type LookupTaskResult struct {
	Etag       *string     `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

func LookupTaskOutput(ctx *pulumi.Context, args LookupTaskOutputArgs, opts ...pulumi.InvokeOption) LookupTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTaskResult, error) {
			args := v.(LookupTaskArgs)
			r, err := LookupTask(ctx, &args, opts...)
			var s LookupTaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTaskResultOutput)
}

type LookupTaskOutputArgs struct {
	Expand      pulumi.StringPtrInput `pulumi:"expand"`
	GroupName   pulumi.StringInput    `pulumi:"groupName"`
	ProjectName pulumi.StringInput    `pulumi:"projectName"`
	ServiceName pulumi.StringInput    `pulumi:"serviceName"`
	TaskName    pulumi.StringInput    `pulumi:"taskName"`
}

func (LookupTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskArgs)(nil)).Elem()
}


type LookupTaskResultOutput struct{ *pulumi.OutputState }

func (LookupTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskResult)(nil)).Elem()
}

func (o LookupTaskResultOutput) ToLookupTaskResultOutput() LookupTaskResultOutput {
	return o
}

func (o LookupTaskResultOutput) ToLookupTaskResultOutputWithContext(ctx context.Context) LookupTaskResultOutput {
	return o
}

func (o LookupTaskResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupTaskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTaskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTaskResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTaskResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupTaskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTaskResultOutput{})
}
