


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:media/v20220501preview:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	AccountName       string `pulumi:"accountName"`
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TransformName     string `pulumi:"transformName"`
}


type LookupJobResult struct {
	CorrelationData map[string]string        `pulumi:"correlationData"`
	Created         string                   `pulumi:"created"`
	Description     *string                  `pulumi:"description"`
	EndTime         string                   `pulumi:"endTime"`
	Id              string                   `pulumi:"id"`
	Input           interface{}              `pulumi:"input"`
	LastModified    string                   `pulumi:"lastModified"`
	Name            string                   `pulumi:"name"`
	Outputs         []JobOutputAssetResponse `pulumi:"outputs"`
	Priority        *string                  `pulumi:"priority"`
	StartTime       string                   `pulumi:"startTime"`
	State           string                   `pulumi:"state"`
	SystemData      SystemDataResponse       `pulumi:"systemData"`
	Type            string                   `pulumi:"type"`
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
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	JobName           pulumi.StringInput `pulumi:"jobName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TransformName     pulumi.StringInput `pulumi:"transformName"`
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

func (o LookupJobResultOutput) CorrelationData() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobResult) map[string]string { return v.CorrelationData }).(pulumi.StringMapOutput)
}

func (o LookupJobResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Input() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupJobResult) interface{} { return v.Input }).(pulumi.AnyOutput)
}

func (o LookupJobResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Outputs() JobOutputAssetResponseArrayOutput {
	return o.ApplyT(func(v LookupJobResult) []JobOutputAssetResponse { return v.Outputs }).(JobOutputAssetResponseArrayOutput)
}

func (o LookupJobResultOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupJobResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
