


package v20170401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOutput(ctx *pulumi.Context, args *LookupOutputArgs, opts ...pulumi.InvokeOption) (*LookupOutputResult, error) {
	var rv LookupOutputResult
	err := ctx.Invoke("azure-native:streamanalytics/v20170401preview:getOutput", args, &rv, opts...)
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

func LookupOutputOutput(ctx *pulumi.Context, args LookupOutputOutputArgs, opts ...pulumi.InvokeOption) LookupOutputResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOutputResult, error) {
			args := v.(LookupOutputArgs)
			r, err := LookupOutput(ctx, &args, opts...)
			var s LookupOutputResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOutputResultOutput)
}

type LookupOutputOutputArgs struct {
	JobName           pulumi.StringInput `pulumi:"jobName"`
	OutputName        pulumi.StringInput `pulumi:"outputName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOutputOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutputArgs)(nil)).Elem()
}


type LookupOutputResultOutput struct{ *pulumi.OutputState }

func (LookupOutputResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutputResult)(nil)).Elem()
}

func (o LookupOutputResultOutput) ToLookupOutputResultOutput() LookupOutputResultOutput {
	return o
}

func (o LookupOutputResultOutput) ToLookupOutputResultOutputWithContext(ctx context.Context) LookupOutputResultOutput {
	return o
}

func (o LookupOutputResultOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupOutputResult) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

func (o LookupOutputResultOutput) Diagnostics() DiagnosticsResponseOutput {
	return o.ApplyT(func(v LookupOutputResult) DiagnosticsResponse { return v.Diagnostics }).(DiagnosticsResponseOutput)
}

func (o LookupOutputResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutputResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupOutputResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutputResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOutputResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOutputResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupOutputResultOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupOutputResult) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

func (o LookupOutputResultOutput) SizeWindow() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupOutputResult) *float64 { return v.SizeWindow }).(pulumi.Float64PtrOutput)
}

func (o LookupOutputResultOutput) TimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOutputResult) *string { return v.TimeWindow }).(pulumi.StringPtrOutput)
}

func (o LookupOutputResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutputResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOutputResultOutput{})
}
