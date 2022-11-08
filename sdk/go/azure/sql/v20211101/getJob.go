


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:sql/v20211101:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupJobArgs struct {
	JobAgentName      string `pulumi:"jobAgentName"`
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupJobResult struct {
	Description *string              `pulumi:"description"`
	Id          string               `pulumi:"id"`
	Name        string               `pulumi:"name"`
	Schedule    *JobScheduleResponse `pulumi:"schedule"`
	Type        string               `pulumi:"type"`
	Version     int                  `pulumi:"version"`
}


func (val *LookupJobResult) Defaults() *LookupJobResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Description) {
		description_ := ""
		tmp.Description = &description_
	}
	tmp.Schedule = tmp.Schedule.Defaults()

	return &tmp
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobResult, error) {
			args := v.(LookupJobArgs)
			r, err := LookupJob(ctx, &args, opts...)
			var s LookupJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	JobAgentName      pulumi.StringInput `pulumi:"jobAgentName"`
	JobName           pulumi.StringInput `pulumi:"jobName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}


type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Schedule() JobScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupJobResult) *JobScheduleResponse { return v.Schedule }).(JobScheduleResponsePtrOutput)
}

func (o LookupJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupJobResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
