


package streamanalytics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStreamingJob(ctx *pulumi.Context, args *LookupStreamingJobArgs, opts ...pulumi.InvokeOption) (*LookupStreamingJobResult, error) {
	var rv LookupStreamingJobResult
	err := ctx.Invoke("azure-native:streamanalytics:getStreamingJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingJobArgs struct {
	Expand            *string `pulumi:"expand"`
	JobName           string  `pulumi:"jobName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupStreamingJobResult struct {
	CompatibilityLevel                 *string                 `pulumi:"compatibilityLevel"`
	CreatedDate                        string                  `pulumi:"createdDate"`
	DataLocale                         *string                 `pulumi:"dataLocale"`
	Etag                               string                  `pulumi:"etag"`
	EventsLateArrivalMaxDelayInSeconds *int                    `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	EventsOutOfOrderMaxDelayInSeconds  *int                    `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	EventsOutOfOrderPolicy             *string                 `pulumi:"eventsOutOfOrderPolicy"`
	Functions                          []FunctionResponse      `pulumi:"functions"`
	Id                                 string                  `pulumi:"id"`
	Inputs                             []InputResponse         `pulumi:"inputs"`
	JobId                              string                  `pulumi:"jobId"`
	JobState                           string                  `pulumi:"jobState"`
	LastOutputEventTime                string                  `pulumi:"lastOutputEventTime"`
	Location                           *string                 `pulumi:"location"`
	Name                               string                  `pulumi:"name"`
	OutputErrorPolicy                  *string                 `pulumi:"outputErrorPolicy"`
	OutputStartMode                    *string                 `pulumi:"outputStartMode"`
	OutputStartTime                    *string                 `pulumi:"outputStartTime"`
	Outputs                            []OutputResponse        `pulumi:"outputs"`
	ProvisioningState                  string                  `pulumi:"provisioningState"`
	Sku                                *SkuResponse            `pulumi:"sku"`
	Tags                               map[string]string       `pulumi:"tags"`
	Transformation                     *TransformationResponse `pulumi:"transformation"`
	Type                               string                  `pulumi:"type"`
}

func LookupStreamingJobOutput(ctx *pulumi.Context, args LookupStreamingJobOutputArgs, opts ...pulumi.InvokeOption) LookupStreamingJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamingJobResult, error) {
			args := v.(LookupStreamingJobArgs)
			r, err := LookupStreamingJob(ctx, &args, opts...)
			var s LookupStreamingJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStreamingJobResultOutput)
}

type LookupStreamingJobOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	JobName           pulumi.StringInput    `pulumi:"jobName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupStreamingJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingJobArgs)(nil)).Elem()
}


type LookupStreamingJobResultOutput struct{ *pulumi.OutputState }

func (LookupStreamingJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingJobResult)(nil)).Elem()
}

func (o LookupStreamingJobResultOutput) ToLookupStreamingJobResultOutput() LookupStreamingJobResultOutput {
	return o
}

func (o LookupStreamingJobResultOutput) ToLookupStreamingJobResultOutputWithContext(ctx context.Context) LookupStreamingJobResultOutput {
	return o
}

func (o LookupStreamingJobResultOutput) CompatibilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.CompatibilityLevel }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingJobResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupStreamingJobResultOutput) DataLocale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.DataLocale }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingJobResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupStreamingJobResultOutput) EventsLateArrivalMaxDelayInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *int { return v.EventsLateArrivalMaxDelayInSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupStreamingJobResultOutput) EventsOutOfOrderMaxDelayInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *int { return v.EventsOutOfOrderMaxDelayInSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupStreamingJobResultOutput) EventsOutOfOrderPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.EventsOutOfOrderPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingJobResultOutput) Functions() FunctionResponseArrayOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) []FunctionResponse { return v.Functions }).(FunctionResponseArrayOutput)
}

func (o LookupStreamingJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStreamingJobResultOutput) Inputs() InputResponseArrayOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) []InputResponse { return v.Inputs }).(InputResponseArrayOutput)
}

func (o LookupStreamingJobResultOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.JobId }).(pulumi.StringOutput)
}

func (o LookupStreamingJobResultOutput) JobState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.JobState }).(pulumi.StringOutput)
}

func (o LookupStreamingJobResultOutput) LastOutputEventTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.LastOutputEventTime }).(pulumi.StringOutput)
}

func (o LookupStreamingJobResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStreamingJobResultOutput) OutputErrorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.OutputErrorPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingJobResultOutput) OutputStartMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.OutputStartMode }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingJobResultOutput) OutputStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.OutputStartTime }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingJobResultOutput) Outputs() OutputResponseArrayOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) []OutputResponse { return v.Outputs }).(OutputResponseArrayOutput)
}

func (o LookupStreamingJobResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupStreamingJobResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupStreamingJobResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStreamingJobResultOutput) Transformation() TransformationResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *TransformationResponse { return v.Transformation }).(TransformationResponsePtrOutput)
}

func (o LookupStreamingJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamingJobResultOutput{})
}
