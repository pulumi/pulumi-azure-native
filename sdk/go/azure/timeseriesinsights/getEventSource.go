


package timeseriesinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupEventSource(ctx *pulumi.Context, args *LookupEventSourceArgs, opts ...pulumi.InvokeOption) (*LookupEventSourceResult, error) {
	var rv LookupEventSourceResult
	err := ctx.Invoke("azure-native:timeseriesinsights:getEventSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventSourceArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	EventSourceName   string `pulumi:"eventSourceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEventSourceResult struct {
	Id       string            `pulumi:"id"`
	Kind     string            `pulumi:"kind"`
	Location string            `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}

func LookupEventSourceOutput(ctx *pulumi.Context, args LookupEventSourceOutputArgs, opts ...pulumi.InvokeOption) LookupEventSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventSourceResult, error) {
			args := v.(LookupEventSourceArgs)
			r, err := LookupEventSource(ctx, &args, opts...)
			var s LookupEventSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventSourceResultOutput)
}

type LookupEventSourceOutputArgs struct {
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	EventSourceName   pulumi.StringInput `pulumi:"eventSourceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventSourceArgs)(nil)).Elem()
}


type LookupEventSourceResultOutput struct{ *pulumi.OutputState }

func (LookupEventSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventSourceResult)(nil)).Elem()
}

func (o LookupEventSourceResultOutput) ToLookupEventSourceResultOutput() LookupEventSourceResultOutput {
	return o
}

func (o LookupEventSourceResultOutput) ToLookupEventSourceResultOutputWithContext(ctx context.Context) LookupEventSourceResultOutput {
	return o
}

func (o LookupEventSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEventSourceResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventSourceResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupEventSourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventSourceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupEventSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEventSourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEventSourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEventSourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventSourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventSourceResultOutput{})
}
