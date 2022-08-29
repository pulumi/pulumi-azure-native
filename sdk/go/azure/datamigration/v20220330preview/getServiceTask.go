


package v20220330preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceTask(ctx *pulumi.Context, args *LookupServiceTaskArgs, opts ...pulumi.InvokeOption) (*LookupServiceTaskResult, error) {
	var rv LookupServiceTaskResult
	err := ctx.Invoke("azure-native:datamigration/v20220330preview:getServiceTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceTaskArgs struct {
	Expand      *string `pulumi:"expand"`
	GroupName   string  `pulumi:"groupName"`
	ServiceName string  `pulumi:"serviceName"`
	TaskName    string  `pulumi:"taskName"`
}


type LookupServiceTaskResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	Properties interface{}        `pulumi:"properties"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupServiceTaskOutput(ctx *pulumi.Context, args LookupServiceTaskOutputArgs, opts ...pulumi.InvokeOption) LookupServiceTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceTaskResult, error) {
			args := v.(LookupServiceTaskArgs)
			r, err := LookupServiceTask(ctx, &args, opts...)
			var s LookupServiceTaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceTaskResultOutput)
}

type LookupServiceTaskOutputArgs struct {
	Expand      pulumi.StringPtrInput `pulumi:"expand"`
	GroupName   pulumi.StringInput    `pulumi:"groupName"`
	ServiceName pulumi.StringInput    `pulumi:"serviceName"`
	TaskName    pulumi.StringInput    `pulumi:"taskName"`
}

func (LookupServiceTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceTaskArgs)(nil)).Elem()
}


type LookupServiceTaskResultOutput struct{ *pulumi.OutputState }

func (LookupServiceTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceTaskResult)(nil)).Elem()
}

func (o LookupServiceTaskResultOutput) ToLookupServiceTaskResultOutput() LookupServiceTaskResultOutput {
	return o
}

func (o LookupServiceTaskResultOutput) ToLookupServiceTaskResultOutputWithContext(ctx context.Context) LookupServiceTaskResultOutput {
	return o
}

func (o LookupServiceTaskResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceTaskResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupServiceTaskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceTaskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceTaskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceTaskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceTaskResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupServiceTaskResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupServiceTaskResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServiceTaskResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupServiceTaskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceTaskResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceTaskResultOutput{})
}
